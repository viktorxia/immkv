// samplecolumns is the read-only phase-1 sampler for the base.col design
// (docs/eureka/2026-11-09-hdd-columnar-bundle.md, section 10.7).
//
// It streams one gnomAD sites VCF (one chromosome), expands each record into
// one logical record per buildable ALT key (the same filtering as
// test/buildfull), and measures:
//
//  1. INFO field order: is the field sequence in each value consistent with
//     the ##INFO declaration order? The columnar path would emit fields in
//     schema order, so inversions are where record-order and columnar output
//     differ.
//  2. Lossless decimal round-trip rate for Integer/Float fields, using the
//     grammar "optional sign + digits + optional decimal point" that the
//     base.col value encoding assumes.
//  3. Per-field, per-bin compressed volume (zstd), presence-bitmap size, and
//     value-length table size, compared with the existing record-order
//     library's value-file size.
//
// The sampler is read-only: it opens the VCF and, optionally, the existing
// chrN.keys header for cross-validation, and writes only under -out.
//
// Usage:
//
//	samplecolumns -vcf <chr22.vcf.bgz> -out <dir> [-lib <full-load library dir>]
//	              [-chrom chr22] [-bin 1000000] [-level 9] [-max-examples 8]
//	              [-limit 0]
package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/biogo/hts/bgzf"
	"github.com/klauspost/compress/zstd"
	"github.com/viktorxia/immkv"
)

const keysMagicSample = "IMKVK002"

// infoDef is one ##INFO declaration in header order.
type infoDef struct {
	id     string
	number string
	typ    string
	isFlag bool
	order  int
}

type decExample struct {
	reason string
	value  string
}

type orderExample struct {
	pos     uint32
	known   int
	invs    int
	pair    string
	snippet string
}

// fieldStat accumulates one field's per-key and per-bin measurements.
type fieldStat struct {
	def *infoDef

	present     uint64 // keys where the field appears
	rawBytes    uint64 // raw value bytes over present keys (flags: 0)
	compressed  uint64 // zstd of per-bin value streams
	bitmap      uint64 // ceil(bin records/8) per field per bin
	lenVarint   uint64 // compact length table: uvarint per present value
	binFrames   uint64 // non-empty bins for this field
	decElements uint64
	decPass     uint64
	decFail     uint64
	decReasons  map[string]uint64
	decExamples []decExample

	buf          []byte // current bin's present values, in record order
	presentInBin uint64
}

type rowField struct {
	fi  int
	val []byte
}

type runStats struct {
	rows       uint64
	refDotRows uint64
	noKeysRows uint64
	keys       uint64
	limited    bool

	noInfoKeys   uint64
	unknownOcc   uint64
	duplicateOcc uint64
	emptySegOcc  uint64
	flagWithVal  uint64
	nonFlagBare  uint64

	keysKnown  uint64
	orderOK    uint64
	orderBad   uint64
	inversions uint64

	orderExamples []orderExample

	infoBytes uint64
	lengths   []uint32
	buckets   [7]uint64

	bins         uint64
	binRecords   uint64
	binRecordSum uint64
	binRecordMax uint64
	binDirBytes  uint64

	fieldCount int
	typeCounts map[string]int
	numCounts  map[string]int

	libRead     bool
	libKeys     uint64
	libKeysOK   bool
	libValsSize int64

	elapsed time.Duration
}

type sampleConfig struct {
	vcfPath   string
	libDir    string
	chromFlag string
	binWidth  uint64
	level     int
	limit     uint64
}

func main() {
	vcf := flag.String("vcf", "", "gnomAD sites VCF (one chromosome, .vcf.bgz)")
	out := flag.String("out", "", "output directory")
	lib := flag.String("lib", "", "optional full-load library directory for validation")
	chrom := flag.String("chrom", "", "optional expected chromosome, e.g. chr22")
	binWidth := flag.Uint64("bin", 1_000_000, "bin width in bp")
	level := flag.Int("level", 9, "zstd level for column volume estimates")
	maxExamples := flag.Int("max-examples", 8, "max examples per field or run")
	limit := flag.Uint64("limit", 0, "stop after this many VCF rows (0 = all)")
	flag.Parse()
	if *vcf == "" || *out == "" {
		flag.Usage()
		os.Exit(2)
	}
	start := time.Now()
	if err := run(*vcf, *out, *lib, *chrom, *binWidth, *level, *maxExamples, *limit); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stderr, "done in %s\n", time.Since(start).Round(time.Second))
}

func run(vcfPath, outDir, libDir, chromFlag string, binWidth uint64, level, maxExamples int, limit uint64) error {
	if binWidth == 0 {
		return errors.New("bin width must be > 0")
	}
	start := time.Now()
	cfg := sampleConfig{
		vcfPath: vcfPath, libDir: libDir, chromFlag: chromFlag,
		binWidth: binWidth, level: level, limit: limit,
	}

	f, err := os.Open(vcfPath)
	if err != nil {
		return err
	}
	defer f.Close()
	zr, err := bgzf.NewReader(f, 0)
	if err != nil {
		return err
	}
	defer zr.Close()
	br := bufio.NewReaderSize(zr, 1<<20)

	order, defs, err := readHeader(br)
	if err != nil {
		return err
	}
	if len(order) == 0 {
		return errors.New("no ##INFO declarations in header")
	}
	idx := make(map[string]int, len(order))
	fields := make([]*fieldStat, len(order))
	st := &runStats{
		fieldCount: len(order),
		typeCounts: make(map[string]int),
		numCounts:  make(map[string]int),
	}
	for i, id := range order {
		d := defs[id]
		fields[i] = &fieldStat{def: d, decReasons: make(map[string]uint64)}
		idx[id] = i
		st.typeCounts[d.typ]++
		st.numCounts[d.number]++
	}

	enc, err := zstd.NewWriter(nil,
		zstd.WithEncoderLevel(zstd.EncoderLevelFromZstd(level)),
		zstd.WithEncoderCRC(false))
	if err != nil {
		return err
	}
	defer enc.Close()

	var (
		curBin     uint64
		binInit    bool
		binRecords uint64
		chromID    uint8
		chromInit  bool
	)
	if chromFlag != "" {
		cid, ok := immkv.ChromID(chromFlag)
		if !ok {
			return fmt.Errorf("invalid -chrom %q", chromFlag)
		}
		chromID, chromInit = cid, true
	}

	flushBin := func() {
		if !binInit {
			return
		}
		st.bins++
		st.binDirBytes += 16 * uint64(len(fields))
		st.binRecordSum += binRecords
		if binRecords > st.binRecordMax {
			st.binRecordMax = binRecords
		}
		for _, fs := range fields {
			fs.bitmap += (binRecords + 7) / 8
			if fs.presentInBin > 0 && !fs.def.isFlag {
				comp := enc.EncodeAll(fs.buf, nil)
				fs.compressed += uint64(len(comp))
				fs.binFrames++
			}
			fs.buf = fs.buf[:0]
			fs.presentInBin = 0
		}
		binRecords = 0
	}

	seenGen := make([]uint32, len(order))
	rowPos := make([]int32, len(order))
	var gen uint32
	var rowFields []rowField
	var seq []int

	sc := bufio.NewScanner(br)
	sc.Buffer(make([]byte, 64<<10), 64<<20)
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		if line[len(line)-1] == '\r' {
			line = line[:len(line)-1]
		}
		cols := bytes.SplitN(line, []byte{'\t'}, 9)
		if len(cols) < 8 {
			return fmt.Errorf("row %d: fewer than 8 columns", st.rows+1)
		}
		cid, ok := immkv.ChromID(string(cols[0]))
		if !ok {
			continue // buildfull skips unknown contigs
		}
		if !chromInit {
			chromID, chromInit = cid, true
		}
		if cid != chromID {
			return fmt.Errorf("row %d: multiple chromosomes (%s then %s)",
				st.rows+1, immkv.ChromName(chromID), immkv.ChromName(cid))
		}
		pos, err := strconv.ParseUint(string(cols[1]), 10, 32)
		if err != nil {
			return fmt.Errorf("row %d: bad POS %q: %w", st.rows+1, cols[1], err)
		}
		if limit > 0 && st.rows >= limit {
			st.limited = true
			break
		}
		st.rows++
		if string(cols[3]) == "." {
			st.refDotRows++
			continue
		}
		var nKeys uint64
		for _, alt := range bytes.Split(cols[4], []byte{','}) {
			if skipAlt(alt) {
				continue
			}
			nKeys++
		}
		if nKeys == 0 {
			st.noKeysRows++
			continue
		}

		info := cols[7]
		noInfo := len(info) == 0 || bytes.Equal(info, []byte("."))
		gen++
		rowFields = rowFields[:0]
		seq = seq[:0]
		var unknown, dup, emptySeg int
		if !noInfo {
			rest := info
			for len(rest) > 0 {
				var seg []byte
				if i := bytes.IndexByte(rest, ';'); i >= 0 {
					seg, rest = rest[:i], rest[i+1:]
				} else {
					seg, rest = rest, nil
				}
				if len(seg) == 0 {
					emptySeg++
					continue
				}
				var id, val []byte
				eq := bytes.IndexByte(seg, '=')
				if eq >= 0 {
					id, val = seg[:eq], seg[eq+1:]
				} else {
					id = seg
				}
				if len(id) == 0 {
					emptySeg++
					continue
				}
				fi, ok := idx[string(id)]
				if !ok {
					unknown++
					continue
				}
				fs := fields[fi]
				if fs.def.isFlag {
					if eq >= 0 {
						st.flagWithVal += nKeys
					}
				} else if eq < 0 {
					st.nonFlagBare += nKeys
				}
				seq = append(seq, fi)
				if seenGen[fi] != gen {
					seenGen[fi] = gen
					rowPos[fi] = int32(len(rowFields))
					rowFields = append(rowFields, rowField{fi: fi, val: val})
				} else {
					dup++
					rowFields[rowPos[fi]].val = val
				}
			}
		}

		// INFO order: a record is order-consistent iff its known fields
		// appear in strictly increasing ##INFO order.
		inv := 0
		for i := 1; i < len(seq); i++ {
			if seq[i] <= seq[i-1] {
				inv++
			}
		}
		if len(seq) > 0 {
			st.keysKnown += nKeys
			if inv == 0 {
				st.orderOK += nKeys
			} else {
				st.orderBad += nKeys
				st.inversions += nKeys * uint64(inv)
				if len(st.orderExamples) < maxExamples {
					pair := ""
					for i := 1; i < len(seq); i++ {
						if seq[i] <= seq[i-1] {
							pair = order[seq[i-1]] + ">" + order[seq[i]]
							break
						}
					}
					st.orderExamples = append(st.orderExamples, orderExample{
						pos: uint32(pos), known: len(seq), invs: inv, pair: pair,
						snippet: truncate(info, 160),
					})
				}
			}
		}

		if noInfo {
			st.noInfoKeys += nKeys
		}
		st.unknownOcc += nKeys * uint64(unknown)
		st.duplicateOcc += nKeys * uint64(dup)
		st.emptySegOcc += nKeys * uint64(emptySeg)

		// Row length is per logical record (per ALT key), matching the
		// library's record count.
		infoLen := len(info)
		st.infoBytes += nKeys * uint64(infoLen)
		st.buckets[rowBucket(infoLen)] += nKeys
		for k := uint64(0); k < nKeys; k++ {
			st.lengths = append(st.lengths, uint32(infoLen))
		}

		for _, rf := range rowFields {
			fs := fields[rf.fi]
			fs.present += nKeys
			fs.presentInBin += nKeys
			if !fs.def.isFlag {
				fs.rawBytes += nKeys * uint64(len(rf.val))
				fs.lenVarint += nKeys * uint64(uvarintLen(len(rf.val)))
				for k := uint64(0); k < nKeys; k++ {
					fs.buf = append(fs.buf, rf.val...)
				}
			}
			if fs.def.typ == "Integer" || fs.def.typ == "Float" {
				for _, e := range bytes.Split(rf.val, []byte{','}) {
					ok, reason := decimalRoundTrip(e)
					fs.decElements += nKeys
					if ok {
						fs.decPass += nKeys
					} else {
						fs.decFail += nKeys
						fs.decReasons[reason] += nKeys
						if len(fs.decExamples) < maxExamples {
							fs.decExamples = append(fs.decExamples, decExample{
								reason: reason, value: truncate(e, 80),
							})
						}
					}
				}
			}
		}

		bin := pos / binWidth
		if !binInit {
			curBin, binInit = bin, true
		} else if bin != curBin {
			if bin < curBin {
				return fmt.Errorf("row %d: POS bin %d after %d (input not position-sorted)",
					st.rows, bin, curBin)
			}
			flushBin()
			curBin = bin
		}
		binRecords += nKeys
		st.keys += nKeys
		if st.rows%100000 == 0 {
			fmt.Fprintf(os.Stderr, "rows=%d keys=%d elapsed=%s\n",
				st.rows, st.keys, time.Since(start).Round(time.Second))
		}
	}
	if err := sc.Err(); err != nil {
		return err
	}
	flushBin()
	st.elapsed = time.Since(start)

	if libDir != "" && chromInit {
		if k, sz, err := readLibInfo(libDir, immkv.ChromName(chromID)); err != nil {
			fmt.Fprintf(os.Stderr, "warning: library validation skipped: %v\n", err)
		} else {
			st.libRead = true
			st.libKeys, st.libValsSize = k, sz
			st.libKeysOK = k == st.keys
		}
	}

	return writeOutputs(outDir, fields, st, cfg)
}

func readHeader(br *bufio.Reader) ([]string, map[string]*infoDef, error) {
	defs := make(map[string]*infoDef)
	var order []string
	for {
		line, err := br.ReadBytes('\n')
		if len(line) > 0 {
			t := bytes.TrimRight(line, "\r\n")
			switch {
			case bytes.HasPrefix(t, []byte("##INFO=<")):
				if id, d, ok := parseInfoLine(t); ok {
					if _, exists := defs[id]; !exists {
						d.order = len(order)
						order = append(order, id)
					}
					defs[id] = d
				}
			case bytes.HasPrefix(t, []byte("#CHROM")):
				return order, defs, nil
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
	d := &infoDef{id: id, number: number, typ: typ}
	d.isFlag = number == "0" || typ == "Flag"
	return id, d, true
}

func headerRawValue(v []byte) []byte {
	if len(v) >= 2 && v[0] == '"' && v[len(v)-1] == '"' {
		return v[1 : len(v)-1]
	}
	return v
}

func skipAlt(alt []byte) bool {
	return len(alt) == 0 || alt[0] == '<' || alt[0] == '*' || alt[0] == '.' ||
		bytes.ContainsAny(alt, "[]")
}

func rowBucket(n int) int {
	switch {
	case n < 64:
		return 0
	case n < 256:
		return 1
	case n < 1024:
		return 2
	case n < 4096:
		return 3
	case n < 16<<10:
		return 4
	case n < 64<<10:
		return 5
	default:
		return 6
	}
}

func uvarintLen(n int) int {
	switch {
	case n < 1<<7:
		return 1
	case n < 1<<14:
		return 2
	case n < 1<<21:
		return 3
	case n < 1<<28:
		return 4
	default:
		return 5
	}
}

// decimalRoundTrip parses the plain-decimal grammar the base.col design
// assumes and reconstructs it byte for byte. Leading/trailing zeros, "+",
// "-0", ".5", and "5." are preserved; exponents, NaN/Inf, whitespace, and
// empty elements are rejected with a reason.
func decimalRoundTrip(s []byte) (bool, string) {
	if len(s) == 0 {
		return false, "empty"
	}
	if len(s) == 3 && bytes.EqualFold(s, []byte("nan")) {
		return false, "nan"
	}
	if (len(s) == 3 && bytes.EqualFold(s, []byte("inf"))) ||
		(len(s) == 4 && (bytes.EqualFold(s, []byte("+inf")) || bytes.EqualFold(s, []byte("-inf")))) {
		return false, "inf"
	}
	i := 0
	var sign byte
	if s[i] == '+' || s[i] == '-' {
		sign = s[i]
		i++
	}
	intStart := i
	for i < len(s) && isDigit(s[i]) {
		i++
	}
	intLen := i - intStart
	hasPoint := false
	fracLen := 0
	if i < len(s) && s[i] == '.' {
		hasPoint = true
		i++
		fracStart := i
		for i < len(s) && isDigit(s[i]) {
			i++
		}
		fracLen = i - fracStart
	}
	if i != len(s) {
		switch s[i] {
		case 'e', 'E':
			return false, "exponent"
		case ' ', '\t', '\r', '\n':
			return false, "whitespace"
		default:
			return false, "nondigit"
		}
	}
	if intLen == 0 && fracLen == 0 {
		return false, "empty-digits"
	}
	out := make([]byte, 0, len(s))
	if sign != 0 {
		out = append(out, sign)
	}
	if hasPoint {
		out = append(out, s[intStart:intStart+intLen]...)
		out = append(out, '.')
		out = append(out, s[intStart+intLen+1:intStart+intLen+1+fracLen]...)
	} else {
		out = append(out, s[intStart:intStart+intLen]...)
	}
	if !bytes.Equal(out, s) {
		return false, "roundtrip"
	}
	return true, ""
}

func isDigit(c byte) bool { return c >= '0' && c <= '9' }

func truncate(b []byte, n int) string {
	if len(b) <= n {
		return string(b)
	}
	return string(b[:n]) + "..."
}

func readLibInfo(libDir, chromName string) (uint64, int64, error) {
	base := filepath.Join(libDir, "chr"+chromName)
	kf, err := os.Open(base + ".keys")
	if err != nil {
		return 0, 0, err
	}
	defer kf.Close()
	var h [64]byte
	if _, err := kf.ReadAt(h[:], 0); err != nil {
		return 0, 0, err
	}
	if string(h[0:8]) != keysMagicSample {
		return 0, 0, fmt.Errorf("bad keys magic in %s.keys", base)
	}
	count := binary.LittleEndian.Uint64(h[12:20])
	st, err := os.Stat(base + ".vals")
	if err != nil {
		return count, 0, err
	}
	return count, st.Size(), nil
}

func writeOutputs(outDir string, fields []*fieldStat, st *runStats, cfg sampleConfig) error {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}
	if err := writeFields(filepath.Join(outDir, "fields.tsv"), fields, st); err != nil {
		return err
	}
	if err := writeFailures(filepath.Join(outDir, "dec_failures.tsv"), fields); err != nil {
		return err
	}
	if err := writeOrderExamples(filepath.Join(outDir, "order_examples.tsv"), st); err != nil {
		return err
	}
	return writeSummary(filepath.Join(outDir, "summary.txt"), fields, st, cfg)
}

func writeFields(path string, fields []*fieldStat, st *runStats) error {
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	fmt.Fprintln(w, "# field\tType\tNumber\theader_order\tpresent_keys\tpresent_pct\tavg_value_bytes\traw_value_bytes\tcompressed_bytes\tbitmap_bytes\tlen_varint_bytes\tlen_u32_bytes\tbin_frames\tdec_elements\tdec_pass\tdec_fail\tdec_all_pass\tdec_reasons\texample")
	for _, fs := range fields {
		pct := "-"
		if st.keys > 0 {
			pct = fmt.Sprintf("%.2f", 100*float64(fs.present)/float64(st.keys))
		}
		avg := "-"
		if fs.present > 0 {
			avg = fmt.Sprintf("%.1f", float64(fs.rawBytes)/float64(fs.present))
		}
		var lenU32 uint64
		if !fs.def.isFlag {
			lenU32 = 4 * fs.present
		}
		allPass := "-"
		if fs.def.typ == "Integer" || fs.def.typ == "Float" {
			switch {
			case fs.decElements == 0:
				allPass = "na"
			case fs.decFail == 0:
				allPass = "1"
			default:
				allPass = "0"
			}
		}
		ex := ""
		if len(fs.decExamples) > 0 {
			ex = fs.decExamples[0].reason + ":" + fs.decExamples[0].value
		}
		if _, err := fmt.Fprintf(w, "%s\t%s\t%s\t%d\t%d\t%s\t%s\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%d\t%s\t%s\t%s\n",
			fs.def.id, fs.def.typ, fs.def.number, fs.def.order+1,
			fs.present, pct, avg, fs.rawBytes, fs.compressed, fs.bitmap,
			fs.lenVarint, lenU32, fs.binFrames,
			fs.decElements, fs.decPass, fs.decFail, allPass,
			reasonString(fs.decReasons), ex); err != nil {
			w.Close()
			return err
		}
	}
	return w.Close()
}

func reasonString(m map[string]uint64) string {
	if len(m) == 0 {
		return "-"
	}
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var b strings.Builder
	for i, k := range keys {
		if i > 0 {
			b.WriteByte(',')
		}
		fmt.Fprintf(&b, "%s=%d", k, m[k])
	}
	return b.String()
}

func writeFailures(path string, fields []*fieldStat) error {
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	fmt.Fprintln(w, "# field\treason\tvalue")
	for _, fs := range fields {
		for _, ex := range fs.decExamples {
			if _, err := fmt.Fprintf(w, "%s\t%s\t%s\n", fs.def.id, ex.reason, ex.value); err != nil {
				w.Close()
				return err
			}
		}
	}
	return w.Close()
}

func writeOrderExamples(path string, st *runStats) error {
	w, err := os.Create(path)
	if err != nil {
		return err
	}
	fmt.Fprintln(w, "# pos\tknown_fields\tinversions\tfirst_inversion\tinfo_snippet")
	for _, ex := range st.orderExamples {
		if _, err := fmt.Fprintf(w, "%d\t%d\t%d\t%s\t%s\n",
			ex.pos, ex.known, ex.invs, ex.pair, ex.snippet); err != nil {
			w.Close()
			return err
		}
	}
	return w.Close()
}

func writeSummary(path string, fields []*fieldStat, st *runStats, cfg sampleConfig) error {
	w, err := os.Create(path)
	if err != nil {
		return err
	}

	sort.Slice(st.lengths, func(i, j int) bool { return st.lengths[i] < st.lengths[j] })
	percentile := func(p int) uint32 {
		if len(st.lengths) == 0 {
			return 0
		}
		i := len(st.lengths) * p / 100
		if i >= len(st.lengths) {
			i = len(st.lengths) - 1
		}
		return st.lengths[i]
	}
	var rowMax uint32
	if len(st.lengths) > 0 {
		rowMax = st.lengths[len(st.lengths)-1]
	}
	lt1k := st.buckets[0] + st.buckets[1] + st.buckets[2]

	var (
		totalRaw, totalCompressed, totalBitmap uint64
		totalLenVarint, totalLenU32            uint64
		totalBinFrames                         uint64
		numericFields, numericAllPass          uint64
		numericFail, numericNoValues           uint64
		numericElements, numericPass           uint64
		numericFailElements                    uint64
	)
	decReasons := make(map[string]uint64)
	for _, fs := range fields {
		totalRaw += fs.rawBytes
		totalCompressed += fs.compressed
		totalBitmap += fs.bitmap
		totalLenVarint += fs.lenVarint
		totalBinFrames += fs.binFrames
		if !fs.def.isFlag {
			totalLenU32 += 4 * fs.present
		}
		if fs.def.typ != "Integer" && fs.def.typ != "Float" {
			continue
		}
		numericFields++
		numericElements += fs.decElements
		numericPass += fs.decPass
		numericFailElements += fs.decFail
		switch {
		case fs.decElements == 0:
			numericNoValues++
		case fs.decFail == 0:
			numericAllPass++
		default:
			numericFail++
		}
		for k, v := range fs.decReasons {
			decReasons[k] += v
		}
	}
	totalVarint := totalCompressed + totalBitmap + totalLenVarint + st.binDirBytes
	totalU32 := totalCompressed + totalBitmap + totalLenU32 + st.binDirBytes

	writeRatio := func(name string, v uint64) {
		if st.libRead && st.libValsSize > 0 {
			fmt.Fprintf(w, "%s\t%.4f\n", name, float64(v)/float64(st.libValsSize))
		} else {
			fmt.Fprintf(w, "%s\t-\n", name)
		}
	}

	fmt.Fprintln(w, "# samplecolumns summary")
	fmt.Fprintf(w, "command\t%s\n", strings.Join(os.Args, " "))
	fmt.Fprintf(w, "vcf\t%s\n", cfg.vcfPath)
	libLabel := cfg.libDir
	if libLabel == "" {
		libLabel = "-"
	}
	fmt.Fprintf(w, "lib\t%s\n", libLabel)
	fmt.Fprintf(w, "chrom\t%s\n", cfg.chromFlag)
	fmt.Fprintf(w, "bin_width\t%d\n", cfg.binWidth)
	fmt.Fprintf(w, "zstd_level\t%d\n", cfg.level)
	fmt.Fprintf(w, "limit\t%d\n", cfg.limit)
	fmt.Fprintf(w, "limited\t%t\n", st.limited)
	fmt.Fprintf(w, "elapsed_seconds\t%.1f\n", st.elapsed.Seconds())
	fmt.Fprintf(w, "rows\t%d\n", st.rows)
	fmt.Fprintf(w, "ref_dot_rows\t%d\n", st.refDotRows)
	fmt.Fprintf(w, "rows_without_keys\t%d\n", st.noKeysRows)
	fmt.Fprintf(w, "keys\t%d\n", st.keys)
	fmt.Fprintf(w, "keys_with_known_fields\t%d\n", st.keysKnown)
	fmt.Fprintf(w, "keys_without_known_fields\t%d\n", st.keys-st.keysKnown)
	fmt.Fprintf(w, "info_dot_keys\t%d\n", st.noInfoKeys)
	fmt.Fprintf(w, "info_unknown_occurrences\t%d\n", st.unknownOcc)
	fmt.Fprintf(w, "info_duplicate_occurrences\t%d\n", st.duplicateOcc)
	fmt.Fprintf(w, "info_empty_segments\t%d\n", st.emptySegOcc)
	fmt.Fprintf(w, "info_flag_with_value\t%d\n", st.flagWithVal)
	fmt.Fprintf(w, "info_nonflag_bare\t%d\n", st.nonFlagBare)
	fmt.Fprintf(w, "info_order_ok\t%d\n", st.orderOK)
	fmt.Fprintf(w, "info_order_bad\t%d\n", st.orderBad)
	fmt.Fprintf(w, "info_order_inversions\t%d\n", st.inversions)
	fmt.Fprintf(w, "header_fields\t%d\n", st.fieldCount)
	for _, t := range []string{"Integer", "Float", "String", "Flag"} {
		fmt.Fprintf(w, "header_type_%s\t%d\n", t, st.typeCounts[t])
	}
	for _, n := range []string{"A", "1", ".", "R", "G", "0"} {
		fmt.Fprintf(w, "header_number_%s\t%d\n", n, st.numCounts[n])
	}
	fmt.Fprintf(w, "row_total_info_bytes\t%d\n", st.infoBytes)
	if st.keys > 0 {
		fmt.Fprintf(w, "row_mean\t%.1f\n", float64(st.infoBytes)/float64(st.keys))
	} else {
		fmt.Fprintf(w, "row_mean\t-\n")
	}
	fmt.Fprintf(w, "row_p50\t%d\n", percentile(50))
	fmt.Fprintf(w, "row_p90\t%d\n", percentile(90))
	fmt.Fprintf(w, "row_p99\t%d\n", percentile(99))
	fmt.Fprintf(w, "row_max\t%d\n", rowMax)
	fmt.Fprintf(w, "row_lt_1k\t%d\n", lt1k)
	if st.keys > 0 {
		fmt.Fprintf(w, "row_lt_1k_pct\t%.2f\n", 100*float64(lt1k)/float64(st.keys))
	} else {
		fmt.Fprintf(w, "row_lt_1k_pct\t-\n")
	}
	for i, label := range []string{"lt_64", "lt_256", "lt_1k", "lt_4k", "lt_16k", "lt_64k", "ge_64k"} {
		fmt.Fprintf(w, "row_bucket_%s\t%d\n", label, st.buckets[i])
	}
	fmt.Fprintf(w, "bins\t%d\n", st.bins)
	if st.bins > 0 {
		fmt.Fprintf(w, "bin_records_avg\t%.1f\n", float64(st.binRecordSum)/float64(st.bins))
	} else {
		fmt.Fprintf(w, "bin_records_avg\t-\n")
	}
	fmt.Fprintf(w, "bin_records_max\t%d\n", st.binRecordMax)
	fmt.Fprintf(w, "bin_dir_bytes\t%d\n", st.binDirBytes)
	fmt.Fprintf(w, "columnar_raw_value_bytes\t%d\n", totalRaw)
	fmt.Fprintf(w, "columnar_compressed_value_bytes\t%d\n", totalCompressed)
	fmt.Fprintf(w, "columnar_bitmap_bytes\t%d\n", totalBitmap)
	fmt.Fprintf(w, "columnar_len_varint_bytes\t%d\n", totalLenVarint)
	fmt.Fprintf(w, "columnar_len_u32_bytes\t%d\n", totalLenU32)
	fmt.Fprintf(w, "columnar_bin_dir_bytes\t%d\n", st.binDirBytes)
	fmt.Fprintf(w, "columnar_value_frames\t%d\n", totalBinFrames)
	fmt.Fprintf(w, "columnar_total_varint\t%d\n", totalVarint)
	fmt.Fprintf(w, "columnar_total_u32\t%d\n", totalU32)
	if totalRaw > 0 {
		fmt.Fprintf(w, "columnar_value_compression_ratio\t%.4f\n", float64(totalCompressed)/float64(totalRaw))
	} else {
		fmt.Fprintf(w, "columnar_value_compression_ratio\t-\n")
	}
	if st.libRead {
		fmt.Fprintf(w, "library_keys\t%d\n", st.libKeys)
		fmt.Fprintf(w, "library_keys_match\t%t\n", st.libKeysOK)
		fmt.Fprintf(w, "record_order_vals_bytes\t%d\n", st.libValsSize)
	} else {
		fmt.Fprintf(w, "library_keys\t-\n")
		fmt.Fprintf(w, "library_keys_match\t-\n")
		fmt.Fprintf(w, "record_order_vals_bytes\t-\n")
	}
	writeRatio("columnar_ratio_varint", totalVarint)
	writeRatio("columnar_ratio_u32", totalU32)
	fmt.Fprintf(w, "numeric_fields\t%d\n", numericFields)
	fmt.Fprintf(w, "numeric_all_pass_fields\t%d\n", numericAllPass)
	fmt.Fprintf(w, "numeric_fail_fields\t%d\n", numericFail)
	fmt.Fprintf(w, "numeric_no_values_fields\t%d\n", numericNoValues)
	fmt.Fprintf(w, "numeric_elements\t%d\n", numericElements)
	fmt.Fprintf(w, "numeric_pass\t%d\n", numericPass)
	fmt.Fprintf(w, "numeric_fail\t%d\n", numericFailElements)
	if numericElements > 0 {
		fmt.Fprintf(w, "numeric_pass_rate\t%.6f\n", float64(numericPass)/float64(numericElements))
	} else {
		fmt.Fprintf(w, "numeric_pass_rate\t-\n")
	}
	reasons := make([]string, 0, len(decReasons))
	for k := range decReasons {
		reasons = append(reasons, k)
	}
	sort.Strings(reasons)
	for _, k := range reasons {
		fmt.Fprintf(w, "dec_reason_%s\t%d\n", k, decReasons[k])
	}
	return w.Close()
}
