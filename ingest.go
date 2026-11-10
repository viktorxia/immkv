// ingest.go 是 VCF/TSV 的建库入口：把输入格式解析、键编码、组序与落盘
// 收进库里，调用方只给路径、字段、前缀和 Options。设计见
// docs/ingest-design.md。
//
// 两个入口的产出形态沿用库的既有约定：
//   - BuildFromVCFs 产出分染色体库（outdir/chrN.keys + chrN.vals），
//     由 OpenEager/OpenStreaming 打开。
//   - BuildFromTSVs 产出单库（basePath.keys + basePath.vals），由
//     OpenSimpleKV 打开。
package immkv

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/biogo/hts/bgzf"
)

const (
	vcfRawHeaderName   = "vcf.raw.header"
	vcfImmkvHeaderName = "vcf.immkv.header"
	maxIngestLine      = 64 << 20 // 单行上限；超出返回错误，不静默截断
)

// BuildFromTSVs 从无表头的 TSV 建单库。输入按 \t 分隔：第一列是键，第二个
// 字段到行尾原样是值（列间 \t 保留）。多个文件按参数顺序拼接后，第一列必须
// 整体按字节序升序；内部走 AddSorted，内存有界。空行跳过，没有 \t 的行视为
// 键加空值，CRLF 的 \r 去掉。输出 basePath.keys + basePath.vals，用
// OpenSimpleKV 打开；basePath 的父目录不存在时创建。
func BuildFromTSVs(basePath string, tsvs []string, opts Options) error {
	if basePath == "" {
		return errors.New("immkv: BuildFromTSVs: empty basePath")
	}
	if len(tsvs) == 0 {
		return errors.New("immkv: BuildFromTSVs: no input files")
	}
	if err := os.MkdirAll(filepath.Dir(basePath), 0o755); err != nil {
		return err
	}
	w, err := Create(basePath, opts)
	if err != nil {
		return err
	}
	finished := false
	defer func() {
		if !finished {
			w.enc.Close()
			w.abort()
		}
	}()
	for _, path := range tsvs {
		if err := addTSVFile(w, path); err != nil {
			return err
		}
	}
	if err := w.Finish(); err != nil {
		return err
	}
	finished = true
	return nil
}

// addTSVFile 把一个 TSV 文件的记录按顺序喂给 w。行切片会被 AddSorted 拷贝，
// 调用方可复用扫描缓冲。
func addTSVFile(w *Writer, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 64*1024), maxIngestLine)
	lineNo := 0
	for sc.Scan() {
		lineNo++
		line := sc.Bytes()
		if n := len(line); n > 0 && line[n-1] == '\r' {
			line = line[:n-1]
		}
		if len(line) == 0 {
			continue
		}
		key := line
		var val []byte
		if i := bytes.IndexByte(line, '\t'); i >= 0 {
			key, val = line[:i], line[i+1:]
		}
		if err := w.AddSorted(key, key, val); err != nil {
			return fmt.Errorf("%s:%d: %w", path, lineNo, err)
		}
	}
	if err := sc.Err(); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	return nil
}

// BuildFromVCFs 从 VCF 建分染色体库：outdir/chrN.keys + chrN.vals，由
// OpenEager/OpenStreaming 打开。输入接受两种形态：
//
//  1. 单个 VCF：每条染色体的记录连续一块，块内 POS 升序；
//  2. 多个 VCF：每个文件只含一条染色体，文件内 POS 升序。
//
// 两种形态都允许同一 CHROM+POS 内 REF/ALT 乱序（组键吸收）。未知 contig、
// REF 为 "." 的记录、符号化/缺失 ALT 按 test/buildfull 口径跳过。
//
// fields 为 nil 或空表示全部 INFO 字段；非空时只存这些字段，fields 的顺序
// 只用于选择，值一律按 ##INFO 声明顺序拼；Number=A 的字段从 ##INFO 头读
// Number、按原始 ALT 下标取逗号元素。prefix 非空时每个 INFO 键写成 prefix_键。
// 源 INFO 的未知字段、重复字段、空段、Flag 带值、非 Flag 缺值都返回错误，
// 不静默丢字段；没有输出字段的记录值写成 "."。
//
// 构建成功后写 outdir/vcf.raw.header（第一个输入的完整头）与
// outdir/vcf.immkv.header（加前缀、按 fields 过滤并按 ##INFO 顺序排列的
// schema 头）。
func BuildFromVCFs(outdir string, vcfs []string, fields []string, prefix string, opts Options) error {
	if outdir == "" {
		return errors.New("immkv: BuildFromVCFs: empty outdir")
	}
	if len(vcfs) == 0 {
		return errors.New("immkv: BuildFromVCFs: no input files")
	}
	if err := validateVCFPrefix(prefix); err != nil {
		return err
	}
	if err := os.MkdirAll(outdir, 0o755); err != nil {
		return err
	}
	if err := checkVCFOutdir(outdir); err != nil {
		return err
	}
	// 头文件是构建成功的标记：先删旧的，失败时就不留头。
	for _, name := range []string{vcfRawHeaderName, vcfImmkvHeaderName} {
		if err := os.Remove(filepath.Join(outdir, name)); err != nil && !os.IsNotExist(err) {
			return err
		}
	}
	multi := len(vcfs) > 1
	plan, err := prepareVCFInputs(vcfs, fields, prefix, multi)
	if err != nil {
		return err
	}
	if multi {
		if err := buildVCFShardsParallel(outdir, plan, opts); err != nil {
			return err
		}
	} else {
		if err := buildVCFShardsSingle(outdir, vcfs[0], plan, opts); err != nil {
			return err
		}
	}
	if err := writeFileAtomic(filepath.Join(outdir, vcfRawHeaderName), plan.raw); err != nil {
		return err
	}
	return writeFileAtomic(filepath.Join(outdir, vcfImmkvHeaderName),
		buildImmkvHeader(plan.header.fileFormat, plan.sel))
}

// validateVCFPrefix 拒绝会破坏 INFO 语法的前缀。
func validateVCFPrefix(prefix string) error {
	if prefix == "" {
		return nil
	}
	if strings.ContainsAny(prefix, ";= \t\r\n") {
		return fmt.Errorf("immkv: BuildFromVCFs: invalid prefix %q", prefix)
	}
	return nil
}

// checkVCFOutdir 要求 outdir 里没有已存在的 chrN 库，避免新旧染色体混在一起。
func checkVCFOutdir(outdir string) error {
	for id := uint8(0); id <= 24; id++ {
		base := filepath.Join(outdir, "chr"+ChromName(id))
		for _, ext := range []string{".keys", ".vals", ".col", ".manifest"} {
			p := base + ext
			if _, err := os.Stat(p); err == nil {
				return fmt.Errorf("immkv: BuildFromVCFs: %s already exists; use an empty outdir", p)
			} else if !os.IsNotExist(err) {
				return err
			}
		}
	}
	return nil
}

// ---- VCF 头与字段 ----

// infoDef 是一条 ##INFO 声明的解析结果。
type infoDef struct {
	line   []byte // 完整声明行（不含换行）
	number string // Number 属性的原文
	typ    string // Type 属性的原文
	isFlag bool   // Number=0 或 Type=Flag
	order  int    // ##INFO 声明顺序（0 起）
}

// vcfHeader 是一个输入文件的头。
type vcfHeader struct {
	raw        []byte // 所有 ## 行 + #CHROM 行，逐字节
	fileFormat string // ##fileformat=... 整行
	info       map[string]*infoDef
	infoOrder  []string
}

// vcfInput 是多文件模式下待建的一个文件。
type vcfInput struct {
	path string
	cid  uint8
}

// valueField 是选中要存进值里的一个 INFO 字段。
type valueField struct {
	id       string // 原始 INFO ID
	prefixed string // prefix 后的键名
	number   string
	typ      string // Type 原文（列存 schema 用）
	isFlag   bool
	defLine  []byte // 原始 ##INFO 声明行
	order    int    // ##INFO 声明顺序（0 起）
}

// infoField 是一条待按 ##INFO 顺序重建的 INFO 字段。
type infoField struct {
	order  int
	id     string
	val    []byte
	isFlag bool
}

// vcfPlan 是一次 BuildFromVCFs 的输入与字段方案。
type vcfPlan struct {
	header  *vcfHeader
	raw     []byte
	sel     []valueField
	all     bool // fields 为空：值取整条 INFO
	prefix  string
	buildID [16]byte // 本次构建的记录序/列存配对代号
	inputs  []vcfInput
}

// prepareVCFInputs 预读所有输入的头，校验一致性，并确定字段方案。
func prepareVCFInputs(vcfs []string, fields []string, prefix string, multi bool) (*vcfPlan, error) {
	plan := &vcfPlan{all: len(fields) == 0, prefix: prefix}
	var seen [25]bool
	for _, path := range vcfs {
		r, err := openVCF(path, 1)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		hdr, firstData, err := readVCFHeader(r.br, multi)
		cerr := r.Close()
		if err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		if cerr != nil {
			return nil, fmt.Errorf("%s: %w", path, cerr)
		}
		if plan.header == nil {
			plan.header = hdr
			plan.raw = hdr.raw
		} else if err := sameVCFInfo(plan.header, hdr); err != nil {
			return nil, fmt.Errorf("%s: %w", path, err)
		}
		if multi {
			if firstData == nil {
				return nil, fmt.Errorf("%s: no records", path)
			}
			cid, err := firstDataChrom(firstData)
			if err != nil {
				return nil, fmt.Errorf("%s: %w", path, err)
			}
			if seen[cid] {
				return nil, fmt.Errorf("%s: chromosome %s appears in more than one input", path, ChromName(cid))
			}
			seen[cid] = true
			plan.inputs = append(plan.inputs, vcfInput{path: path, cid: cid})
		}
	}
	sel, err := selectValueFields(plan.header, fields, prefix)
	if err != nil {
		return nil, err
	}
	plan.sel = sel
	id, err := newBuildID()
	if err != nil {
		return nil, err
	}
	plan.buildID = id
	return plan, nil
}

// selectValueFields 把 fields 解析成 valueField 列表；fields 为空时取全部
// 字段；非空时只做选择，结果统一按 ##INFO 声明顺序排列（fields 的顺序不再
// 决定值序）。
func selectValueFields(h *vcfHeader, fields []string, prefix string) ([]valueField, error) {
	if len(fields) == 0 {
		sel := make([]valueField, 0, len(h.infoOrder))
		for _, id := range h.infoOrder {
			d := h.info[id]
			sel = append(sel, valueField{
				id: id, prefixed: prefixVCFID(prefix, id),
				number: d.number, typ: d.typ, isFlag: d.isFlag, defLine: d.line,
				order: d.order,
			})
		}
		return sel, nil
	}
	dup := make(map[string]bool, len(fields))
	sel := make([]valueField, 0, len(fields))
	for _, name := range fields {
		if name == "" {
			return nil, errors.New("immkv: BuildFromVCFs: empty field name")
		}
		if dup[name] {
			return nil, fmt.Errorf("immkv: BuildFromVCFs: duplicate field %q", name)
		}
		dup[name] = true
		d, ok := h.info[name]
		if !ok {
			return nil, fmt.Errorf("immkv: BuildFromVCFs: field %q not declared in ##INFO", name)
		}
		sel = append(sel, valueField{
			id: name, prefixed: prefixVCFID(prefix, name),
			number: d.number, typ: d.typ, isFlag: d.isFlag, defLine: d.line,
			order: d.order,
		})
	}
	sort.Slice(sel, func(i, j int) bool { return sel[i].order < sel[j].order })
	return sel, nil
}

// prefixVCFID 返回 prefix_字段名；prefix 为空时原样返回。
func prefixVCFID(prefix, id string) string {
	if prefix == "" {
		return id
	}
	return prefix + "_" + id
}

// sameVCFInfo 校验两个输入的头在 fileformat、全部 ##INFO 声明与声明顺序上一致。
func sameVCFInfo(a, b *vcfHeader) error {
	if a.fileFormat != b.fileFormat {
		return fmt.Errorf("##fileformat differs: %q vs %q", a.fileFormat, b.fileFormat)
	}
	if len(a.info) != len(b.info) {
		return errors.New("##INFO definitions differ")
	}
	for id, da := range a.info {
		db, ok := b.info[id]
		if !ok || da.number != db.number || !bytes.Equal(da.line, db.line) {
			return fmt.Errorf("##INFO definition %q differs", id)
		}
	}
	if len(a.infoOrder) != len(b.infoOrder) {
		return errors.New("##INFO order differs")
	}
	for i := range a.infoOrder {
		if a.infoOrder[i] != b.infoOrder[i] {
			return fmt.Errorf("##INFO order differs: %q vs %q", a.infoOrder[i], b.infoOrder[i])
		}
	}
	return nil
}

// ---- 文件读取 ----

// vcfReader 统一纯文本与 bgzf 输入，Close 同时关解压器和文件。
type vcfReader struct {
	f  *os.File
	br *bufio.Reader
	bg *bgzf.Reader
}

func (r *vcfReader) Close() error {
	var first error
	if r.bg != nil {
		if err := r.bg.Close(); err != nil {
			first = err
		}
	}
	if err := r.f.Close(); err != nil && first == nil {
		first = err
	}
	return first
}

// openVCF 打开输入：.gz/.bgz 走 bgzf（必须是 bgzf，不是普通 gzip），其余纯文本。
// bgzfWorkers=0 表示解压内部并行，适合单文件顺序建库；多文件模式传 1，由文件
// 级并行提供吞吐。
func openVCF(path string, bgzfWorkers int) (*vcfReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	if strings.HasSuffix(path, ".gz") || strings.HasSuffix(path, ".bgz") {
		bg, err := bgzf.NewReader(f, bgzfWorkers)
		if err != nil {
			f.Close()
			return nil, err
		}
		return &vcfReader{f: f, br: bufio.NewReaderSize(bg, 1<<20), bg: bg}, nil
	}
	return &vcfReader{f: f, br: bufio.NewReaderSize(f, 1<<20)}, nil
}

// readVCFHeader 读 ## 行与 #CHROM 行；needFirst 时再读第一条数据行（跳空行）。
// 返回的 raw 是头的逐字节副本。
func readVCFHeader(br *bufio.Reader, needFirst bool) (*vcfHeader, []byte, error) {
	h := &vcfHeader{info: make(map[string]*infoDef)}
	for {
		line, err := br.ReadBytes('\n')
		if len(line) > 0 {
			trimmed := bytes.TrimRight(line, "\r\n")
			switch {
			case len(trimmed) == 0:
				// 空行：跳过
			case bytes.HasPrefix(trimmed, []byte("##")):
				h.raw = append(h.raw, line...)
				if bytes.HasPrefix(trimmed, []byte("##INFO=<")) {
					if id, d, ok := parseInfoLine(trimmed); ok {
						if _, exists := h.info[id]; exists {
							return nil, nil, fmt.Errorf("duplicate ##INFO %q", id)
						}
						d.order = len(h.infoOrder)
						h.info[id] = d
						h.infoOrder = append(h.infoOrder, id)
					}
				} else if bytes.HasPrefix(trimmed, []byte("##fileformat=")) {
					h.fileFormat = string(trimmed)
				}
			case bytes.HasPrefix(trimmed, []byte("#CHROM")):
				if len(line) == len(trimmed) { // 末行没有换行
					h.raw = append(h.raw, line...)
					h.raw = append(h.raw, '\n')
				} else {
					h.raw = append(h.raw, line...)
				}
				if !needFirst {
					return h, nil, nil
				}
				first, err := readFirstDataLine(br)
				return h, first, err
			default:
				return nil, nil, errors.New("data before #CHROM line")
			}
		}
		if err != nil {
			if err == io.EOF {
				return nil, nil, errors.New("header missing #CHROM line")
			}
			return nil, nil, err
		}
	}
}

// readFirstDataLine 返回第一条非空数据行（去换行、拷贝）。
func readFirstDataLine(br *bufio.Reader) ([]byte, error) {
	for {
		line, err := br.ReadBytes('\n')
		if len(line) > 0 {
			trimmed := bytes.TrimRight(line, "\r\n")
			if len(trimmed) > 0 {
				return append([]byte(nil), trimmed...), nil
			}
		}
		if err != nil {
			if err == io.EOF {
				return nil, nil
			}
			return nil, err
		}
	}
}

// parseInfoLine 解析 ##INFO=<ID=...,Number=...,Type=...,Description="..."> 的
// ID、Number、Type；Description 里的逗号与转义引号不影响解析。
func parseInfoLine(line []byte) (string, *infoDef, bool) {
	const prefix = "##INFO=<"
	if !bytes.HasPrefix(line, []byte(prefix)) || line[len(line)-1] != '>' {
		return "", nil, false
	}
	content := line[len(prefix) : len(line)-1]
	var id, number, typ string
	start := 0
	inQuote := false
	for i := 0; i <= len(content); i++ {
		if i < len(content) {
			c := content[i]
			if inQuote {
				if c == '\\' {
					i++
				} else if c == '"' {
					inQuote = false
				}
				continue
			}
			if c == '"' {
				inQuote = true
				continue
			}
			if c != ',' {
				continue
			}
		}
		seg := content[start:i]
		if eq := bytes.IndexByte(seg, '='); eq >= 0 {
			switch string(seg[:eq]) {
			case "ID":
				id = string(headerRawValue(seg[eq+1:]))
			case "Number":
				number = string(headerRawValue(seg[eq+1:]))
			case "Type":
				typ = string(headerRawValue(seg[eq+1:]))
			}
		}
		start = i + 1
	}
	if id == "" {
		return "", nil, false
	}
	d := &infoDef{line: append([]byte(nil), line...), number: number, typ: typ}
	d.isFlag = number == "0" || typ == "Flag"
	return id, d, true
}

// headerRawValue 去掉属性值两侧的双引号。
func headerRawValue(v []byte) []byte {
	if len(v) >= 2 && v[0] == '"' && v[len(v)-1] == '"' {
		return v[1 : len(v)-1]
	}
	return v
}

// firstDataChrom 从第一行数据取 CHROM 并编码；未知 contig 报错。
func firstDataChrom(line []byte) (uint8, error) {
	chrom, _, _, _, _, err := splitVCFLine(line)
	if err != nil {
		return 0, err
	}
	cid, ok := ChromID(string(chrom))
	if !ok {
		return 0, fmt.Errorf("unknown contig %q", chrom)
	}
	return cid, nil
}

// splitVCFLine 切出 CHROM/POS/REF/ALT/INFO 五列。只认前 8 列：INFO 是第 8
// 列，有样本列时忽略；少于 8 列报错。
func splitVCFLine(line []byte) (chrom, pos, ref, alt, info []byte, err error) {
	var tabs [8]int
	n := 0
	for i := 0; i < len(line) && n < 8; i++ {
		if line[i] == '\t' {
			tabs[n] = i
			n++
		}
	}
	if n < 7 {
		return nil, nil, nil, nil, nil, errors.New("fewer than 8 columns")
	}
	chrom = line[:tabs[0]]
	pos = line[tabs[0]+1 : tabs[1]]
	ref = line[tabs[2]+1 : tabs[3]]
	alt = line[tabs[3]+1 : tabs[4]]
	if n >= 8 {
		info = line[tabs[6]+1 : tabs[7]]
	} else {
		info = line[tabs[6]+1:]
	}
	return chrom, pos, ref, alt, info, nil
}

// skipVCFAlt 是 test/buildfull 的 ALT 过滤口径。
func skipVCFAlt(alt []byte) bool {
	return len(alt) == 0 || alt[0] == '<' || alt[0] == '*' || alt[0] == '.' ||
		bytes.ContainsAny(alt, "[]")
}

// ---- 分染色体落盘 ----

// buildState 是多文件模式的取消标志与首个错误。
type buildState struct {
	canceled atomic.Bool
	mu       sync.Mutex
	err      error
}

func (s *buildState) fail(err error) {
	s.mu.Lock()
	if s.err == nil {
		s.err = err
	}
	s.mu.Unlock()
	s.canceled.Store(true)
}

func (s *buildState) firstErr() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.err
}

// vcfShardBuilder 管理一个输入文件的记录到分染色体 writer 的推进。
type vcfShardBuilder struct {
	outdir string
	opts   Options
	sel    []valueField
	all    bool
	prefix string
	fixed  int // -1：跟随 CHROM 块；>=0：多文件模式只允许该染色体

	cur  int
	w    *Writer
	seen [25]bool

	hdr        *vcfHeader
	seenInfo   []uint32
	gen        uint32
	infoFields []infoField

	buildID [16]byte
	stats   colInfoStats
	srcLast int
	srcBad  bool

	idx     map[string]int
	vals    [][]byte
	has     []bool
	touched []int
	valBuf  []byte
	infoBuf []byte
	lineNo  int
}

func newVCFShardBuilder(outdir string, plan *vcfPlan, opts Options, fixed int) *vcfShardBuilder {
	b := &vcfShardBuilder{
		outdir: outdir, opts: opts, sel: plan.sel, all: plan.all,
		prefix: plan.prefix, fixed: fixed, cur: -1,
		hdr: plan.header, seenInfo: make([]uint32, len(plan.header.infoOrder)),
		buildID: plan.buildID,
	}
	if !b.all {
		b.idx = make(map[string]int, len(b.sel))
		for i := range b.sel {
			b.idx[b.sel[i].id] = i
		}
		b.vals = make([][]byte, len(b.sel))
		b.has = make([]bool, len(b.sel))
	}
	return b
}

// addLine 处理一行 VCF 数据。
func (b *vcfShardBuilder) addLine(line []byte) error {
	b.lineNo++
	if n := len(line); n > 0 && line[n-1] == '\r' {
		line = line[:n-1]
	}
	if len(line) == 0 {
		return nil
	}
	chrom, pos, ref, alt, info, err := splitVCFLine(line)
	if err != nil {
		return fmt.Errorf("line %d: %w", b.lineNo, err)
	}
	cid, ok := ChromID(string(chrom))
	if !ok {
		return nil // 未知 contig：跳过
	}
	if bytes.Equal(ref, []byte(".")) {
		return nil
	}
	p, err := strconv.ParseUint(string(pos), 10, 32)
	if err != nil {
		return fmt.Errorf("line %d: bad POS %q: %w", b.lineNo, pos, err)
	}
	nValid := 0
	for _, a := range bytes.Split(alt, []byte(",")) {
		if !skipVCFAlt(a) {
			nValid++
		}
	}
	if nValid == 0 {
		return nil
	}
	if b.fixed >= 0 && int(cid) != b.fixed {
		return fmt.Errorf("line %d: expected chromosome %s, found %s",
			b.lineNo, ChromName(uint8(b.fixed)), ChromName(cid))
	}
	if int(cid) != b.cur {
		if b.fixed < 0 {
			if b.seen[cid] {
				return fmt.Errorf("line %d: chromosome %s block is not contiguous",
					b.lineNo, ChromName(cid))
			}
			b.seen[cid] = true
		}
		if err := b.finishCurrent(); err != nil {
			return err
		}
		b.stats = colInfoStats{}
		w, err := Create(filepath.Join(b.outdir, "chr"+ChromName(cid)), b.opts)
		if err != nil {
			return err
		}
		b.w = w
		b.cur = int(cid)
	}
	var common []byte
	if b.all {
		v, err := b.sortInfo(info)
		if err != nil {
			return fmt.Errorf("line %d: %w", b.lineNo, err)
		}
		common = v
	} else if err := b.parseInfo(info); err != nil {
		return fmt.Errorf("line %d: %w", b.lineNo, err)
	}
	group := PosKey(cid, uint32(p))
	for altIdx, a := range bytes.Split(alt, []byte(",")) {
		if skipVCFAlt(a) {
			continue
		}
		var val []byte
		if b.all {
			val = common
		} else {
			val = b.buildValue(altIdx)
		}
		key := VarKey(cid, uint32(p), string(ref), string(a))
		if err := b.w.AddSorted(group, key, val); err != nil {
			return fmt.Errorf("line %d: %w", b.lineNo, err)
		}
	}
	return nil
}

// parseInfo 把当前记录的 INFO 解析进 b.vals/b.has（只关心选中字段），同时
// 校验整条 INFO：未知/重复字段、空段、Flag 带值、非 Flag 缺值都返回错误。
func (b *vcfShardBuilder) parseInfo(info []byte) error {
	for _, fi := range b.touched {
		b.vals[fi] = nil
		b.has[fi] = false
	}
	b.touched = b.touched[:0]
	b.gen++
	b.srcLast, b.srcBad = -1, false
	if len(info) == 0 || (len(info) == 1 && info[0] == '.') {
		b.endRecordOrder()
		return nil
	}
	rest := info
	for len(rest) > 0 {
		var seg []byte
		if i := bytes.IndexByte(rest, ';'); i >= 0 {
			seg, rest = rest[:i], rest[i+1:]
		} else {
			seg, rest = rest, nil
		}
		if len(seg) == 0 {
			return errors.New("empty INFO segment")
		}
		eq := bytes.IndexByte(seg, '=')
		key := seg
		var v []byte
		if eq >= 0 {
			key, v = seg[:eq], seg[eq+1:]
		}
		if len(key) == 0 {
			return errors.New("empty INFO field name")
		}
		d, ok := b.hdr.info[string(key)]
		if !ok {
			return fmt.Errorf("unknown INFO field %q", key)
		}
		b.noteInfoOrder(d)
		if b.seenInfo[d.order] == b.gen {
			return fmt.Errorf("duplicate INFO field %q", key)
		}
		b.seenInfo[d.order] = b.gen
		if d.isFlag && eq >= 0 {
			return fmt.Errorf("Flag INFO field %q has a value", key)
		}
		if !d.isFlag && eq < 0 {
			return fmt.Errorf("non-Flag INFO field %q has no value", key)
		}
		if fi, ok := b.idx[string(key)]; ok {
			if !b.has[fi] {
				b.touched = append(b.touched, fi)
			}
			b.has[fi] = true
			b.vals[fi] = v
		}
	}
	b.endRecordOrder()
	return nil
}

// noteInfoOrder 记录源 INFO 顺序是否与 ##INFO 声明序一致；只作诊断统计。
func (b *vcfShardBuilder) noteInfoOrder(d *infoDef) {
	if d.order < b.srcLast {
		b.srcBad = true
		b.stats.Inversions++
	}
	b.srcLast = d.order
}

// endRecordOrder 结算当前记录的源顺序统计。
func (b *vcfShardBuilder) endRecordOrder() {
	if b.srcBad {
		b.stats.Bad++
	} else {
		b.stats.OK++
	}
}

// buildValue 按 ##INFO 声明顺序拼一个 ALT 的值；Number=A 取该 ALT 下标的
// 元素；没有输出字段时返回 "."。
func (b *vcfShardBuilder) buildValue(altIdx int) []byte {
	b.valBuf = b.valBuf[:0]
	for fi := range b.sel {
		if !b.has[fi] {
			continue
		}
		v := b.vals[fi]
		if b.sel[fi].number == "A" {
			var ok bool
			v, ok = commaElem(v, altIdx)
			if !ok {
				continue
			}
		}
		if len(b.valBuf) > 0 {
			b.valBuf = append(b.valBuf, ';')
		}
		b.valBuf = append(b.valBuf, b.sel[fi].prefixed...)
		if b.sel[fi].isFlag {
			continue
		}
		b.valBuf = append(b.valBuf, '=')
		b.valBuf = append(b.valBuf, v...)
	}
	if len(b.valBuf) == 0 {
		b.valBuf = append(b.valBuf, '.')
	}
	return b.valBuf
}

// commaElem 返回逗号分隔值的第 n 个元素；元素可以是空串。
func commaElem(v []byte, n int) ([]byte, bool) {
	start := 0
	for i := 0; i <= len(v); i++ {
		if i == len(v) || v[i] == ',' {
			if n == 0 {
				return v[start:i], true
			}
			n--
			start = i + 1
		}
	}
	return nil, false
}

// sortInfo 把整条 INFO 按 ##INFO 声明顺序重排并重建值；prefix 非空时每个
// 字段名加 prefix_。INFO 为 "." 或空时返回 "."。未知字段、重复字段、空段、
// Flag 带值、非 Flag 缺值都返回错误，不静默丢字段。
func (b *vcfShardBuilder) sortInfo(info []byte) ([]byte, error) {
	b.infoBuf = b.infoBuf[:0]
	b.srcLast, b.srcBad = -1, false
	if len(info) == 0 || (len(info) == 1 && info[0] == '.') {
		b.endRecordOrder()
		return []byte("."), nil
	}
	b.infoFields = b.infoFields[:0]
	b.gen++
	rest := info
	for len(rest) > 0 {
		var seg []byte
		if i := bytes.IndexByte(rest, ';'); i >= 0 {
			seg, rest = rest[:i], rest[i+1:]
		} else {
			seg, rest = rest, nil
		}
		if len(seg) == 0 {
			return nil, errors.New("empty INFO segment")
		}
		eq := bytes.IndexByte(seg, '=')
		key := seg
		var v []byte
		if eq >= 0 {
			key, v = seg[:eq], seg[eq+1:]
		}
		if len(key) == 0 {
			return nil, errors.New("empty INFO field name")
		}
		d, ok := b.hdr.info[string(key)]
		if !ok {
			return nil, fmt.Errorf("unknown INFO field %q", key)
		}
		b.noteInfoOrder(d)
		if b.seenInfo[d.order] == b.gen {
			return nil, fmt.Errorf("duplicate INFO field %q", key)
		}
		b.seenInfo[d.order] = b.gen
		if d.isFlag && eq >= 0 {
			return nil, fmt.Errorf("Flag INFO field %q has a value", key)
		}
		if !d.isFlag && eq < 0 {
			return nil, fmt.Errorf("non-Flag INFO field %q has no value", key)
		}
		b.infoFields = append(b.infoFields, infoField{
			order: d.order, id: string(key), val: v, isFlag: d.isFlag,
		})
	}
	sort.Slice(b.infoFields, func(i, j int) bool {
		return b.infoFields[i].order < b.infoFields[j].order
	})
	for _, f := range b.infoFields {
		if len(b.infoBuf) > 0 {
			b.infoBuf = append(b.infoBuf, ';')
		}
		if b.prefix != "" {
			b.infoBuf = append(b.infoBuf, b.prefix...)
			b.infoBuf = append(b.infoBuf, '_')
		}
		b.infoBuf = append(b.infoBuf, f.id...)
		if f.isFlag {
			continue
		}
		b.infoBuf = append(b.infoBuf, '=')
		b.infoBuf = append(b.infoBuf, f.val...)
	}
	b.endRecordOrder()
	return b.infoBuf, nil
}

// finishCurrent 落盘当前染色体的 writer。
func (b *vcfShardBuilder) finishCurrent() error {
	if b.w == nil {
		return nil
	}
	w := b.w
	b.w = nil
	if err := w.Finish(); err != nil {
		return err
	}
	base := filepath.Join(b.outdir, "chr"+ChromName(uint8(b.cur)))
	return writeColumnar(base, b.sel, b.buildID, b.stats)
}

// finish 结束构建；一条可建记录都没有时报错。
func (b *vcfShardBuilder) finish() error {
	if b.w == nil {
		return errors.New("no buildable records")
	}
	return b.finishCurrent()
}

// abort 清理未完成的 writer：关编码器、删 .tmp。
func (b *vcfShardBuilder) abort() {
	if b.w != nil {
		b.w.enc.Close()
		b.w.abort()
		b.w = nil
	}
}

// buildVCFShardsSingle 是单文件模式：顺序读，CHROM 块变化时切 writer。
func buildVCFShardsSingle(outdir, path string, plan *vcfPlan, opts Options) error {
	r, err := openVCF(path, 0)
	if err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	defer r.Close()
	if _, _, err := readVCFHeader(r.br, false); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	b := newVCFShardBuilder(outdir, plan, opts, -1)
	defer b.abort()
	sc := bufio.NewScanner(r.br)
	sc.Buffer(make([]byte, 64*1024), maxIngestLine)
	for sc.Scan() {
		if err := b.addLine(sc.Bytes()); err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
	}
	if err := sc.Err(); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	if err := b.finish(); err != nil {
		return fmt.Errorf("%s: %w", path, err)
	}
	return nil
}

// buildVCFShardsParallel 是多文件模式：一个文件一条染色体，按文件并行。
func buildVCFShardsParallel(outdir string, plan *vcfPlan, opts Options) error {
	var state buildState
	var wg sync.WaitGroup
	for _, in := range plan.inputs {
		wg.Add(1)
		go func(in vcfInput) {
			defer wg.Done()
			if err := buildOneVCFShard(outdir, in, plan, opts, &state); err != nil {
				state.fail(err)
			}
		}(in)
	}
	wg.Wait()
	return state.firstErr()
}

// buildOneVCFShard 建一个文件的染色体库。
func buildOneVCFShard(outdir string, in vcfInput, plan *vcfPlan, opts Options, state *buildState) error {
	r, err := openVCF(in.path, 1)
	if err != nil {
		return fmt.Errorf("%s: %w", in.path, err)
	}
	defer r.Close()
	if _, _, err := readVCFHeader(r.br, false); err != nil {
		return fmt.Errorf("%s: %w", in.path, err)
	}
	b := newVCFShardBuilder(outdir, plan, opts, int(in.cid))
	defer b.abort()
	sc := bufio.NewScanner(r.br)
	sc.Buffer(make([]byte, 64*1024), maxIngestLine)
	for sc.Scan() {
		if state.canceled.Load() {
			return nil
		}
		if err := b.addLine(sc.Bytes()); err != nil {
			return fmt.Errorf("%s: %w", in.path, err)
		}
	}
	if err := sc.Err(); err != nil {
		return fmt.Errorf("%s: %w", in.path, err)
	}
	if state.canceled.Load() {
		return nil
	}
	if err := b.finish(); err != nil {
		return fmt.Errorf("%s: %w", in.path, err)
	}
	return nil
}

// ---- header 落盘 ----

// buildImmkvHeader 生成 vcf.immkv.header：fileformat + 选中字段的 ##INFO
// （ID 加前缀，按 ##INFO 声明顺序）+ 无样本的 #CHROM 行。
func buildImmkvHeader(fileFormat string, sel []valueField) []byte {
	if fileFormat == "" {
		fileFormat = "##fileformat=VCFv4.2"
	}
	out := make([]byte, 0, len(fileFormat)+64*len(sel)+64)
	out = append(out, fileFormat...)
	out = append(out, '\n')
	for i := range sel {
		out = append(out, prefixInfoLine(sel[i].defLine, sel[i].prefixed)...)
		out = append(out, '\n')
	}
	out = append(out, "#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO\n"...)
	return out
}

// prefixInfoLine 把 ##INFO 声明里的 ID 属性值替换为 prefixed。
func prefixInfoLine(line []byte, prefixed string) []byte {
	for i := 0; i+4 <= len(line); i++ {
		if (line[i] == '<' || line[i] == ',') && bytes.HasPrefix(line[i+1:], []byte("ID=")) {
			start := i + 1 + 3
			end := start
			for end < len(line) && line[end] != ',' && line[end] != '>' {
				end++
			}
			out := make([]byte, 0, len(line)+len(prefixed))
			out = append(out, line[:start]...)
			out = append(out, prefixed...)
			out = append(out, line[end:]...)
			return out
		}
	}
	return line
}

// writeFileAtomic 写临时文件、fsync、rename。
func writeFileAtomic(path string, data []byte) error {
	tmp := path + ".tmp"
	f, err := os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	_, err = f.Write(data)
	if err == nil {
		err = f.Sync()
	}
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		os.Remove(tmp)
		return err
	}
	if err := os.Rename(tmp, path); err != nil {
		os.Remove(tmp)
		return err
	}
	return nil
}
