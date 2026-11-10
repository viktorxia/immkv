// col.go 是 base.col 的写入侧：把一个已完成的记录序分片（base.keys/base.vals）
// 顺序扫一遍，按基因组区间切 bin，把每个字段的值拆成列，压缩后写成
// base.col + base.manifest。格式契约见 docs/base-col-format.md。
//
// 列存只服务 OpenStreaming 的顺序批量查；点查、前缀、区间扫描仍走记录序。
// 写入发生在记录序 Finish 之后，不改变现有 Writer 与盘上格式。

package immkv

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/klauspost/compress/zstd"
)

const (
	colMagic           = "IMKVC001"
	colVersion         = 1
	colHeaderSize      = 64
	colBinEntrySize    = 32
	colColumnEntrySize = 32
	colBinWidth        = 1_000_000 // v1 固定 1Mb bin
	colManifestFormat  = "immkv-col-manifest"
	colManifestVersion = 1
	colZstdLevel       = 9 // 与抽样一致；值流按列独立压缩
)

// colField 是 schema 里的一列，也是 manifest 的字段条目。Column == 下标。
type colField struct {
	Name     string `json:"name"`
	Column   int    `json:"column"`
	Type     string `json:"type"`
	Number   string `json:"number"`
	Encoding string `json:"encoding"` // raw | bitmap（dec 留第 3 阶段）
}

// colInfoStats 是构建期源 INFO 顺序的统计，只作诊断；入库已强制 header 序。
type colInfoStats struct {
	OK            int64 `json:"ok"`
	Bad           int64 `json:"bad"`
	Inversions    int64 `json:"inversions"`
	Unknown       int64 `json:"unknown"`
	Duplicate     int64 `json:"duplicate"`
	EmptySegments int64 `json:"empty_segments"`
}

// colManifest 是 base.manifest 的 JSON 形态。
type colManifest struct {
	Format          string       `json:"format"`
	Version         int          `json:"version"`
	Base            string       `json:"base"`
	BuildID         string       `json:"build_id"`
	BinWidth        uint32       `json:"bin_width"`
	RecordCount     uint64       `json:"record_count"`
	KeysBytes       int64        `json:"keys_bytes"`
	ValsBytes       int64        `json:"vals_bytes"`
	ColBytes        int64        `json:"col_bytes"`
	SchemaHash      string       `json:"schema_hash"`
	SourceInfoOrder colInfoStats `json:"source_info_order"`
	Fields          []colField   `json:"fields"`
}

// colSchemaHash 按 schema 顺序对 "name\ttype\tnumber\tencoding\n" 取 SHA-256
// 前 8 字节；col 头存 u64，manifest 存同样的 8 字节十六进制。
func colSchemaHash(fields []colField) uint64 {
	h := sha256.New()
	for i := range fields {
		fmt.Fprintf(h, "%s\t%s\t%s\t%s\n",
			fields[i].Name, fields[i].Type, fields[i].Number, fields[i].Encoding)
	}
	sum := h.Sum(nil)
	return binary.LittleEndian.Uint64(sum[:8])
}

// colSchemaString 返回 manifest 里 schema_hash 的十六进制形态。
func colSchemaString(fields []colField) string {
	h := sha256.New()
	for i := range fields {
		fmt.Fprintf(h, "%s\t%s\t%s\t%s\n",
			fields[i].Name, fields[i].Type, fields[i].Number, fields[i].Encoding)
	}
	return hex.EncodeToString(h.Sum(nil)[:8])
}

// colFieldsFromSel 把 ingest 的字段方案转成列 schema；v1 全部 raw/bitmap。
func colFieldsFromSel(sel []valueField) []colField {
	fields := make([]colField, len(sel))
	for i := range sel {
		typ := sel[i].typ
		if typ == "" {
			typ = "String"
		}
		enc := "raw"
		if sel[i].isFlag {
			enc = "bitmap"
		}
		fields[i] = colField{
			Name: sel[i].prefixed, Column: i, Type: typ,
			Number: sel[i].number, Encoding: enc,
		}
	}
	return fields
}

// colEntry 是列目录里的一项（32B）。
type colEntry struct {
	bitmapOff uint64
	valueOff  uint64
	bitmapLen uint32 // 位图 zstd 帧长度
	valueLen  uint32
	present   uint32
	rawLen    uint32 // 存在值原文字节合计（饱和到 u32，仅用于选路估算）
}

// colBinRecord 是 bin 目录里的一项（32B）。
type colBinRecord struct {
	first   uint64
	count   uint32
	dirOff  uint64
	dataEnd uint64
	start   uint64 // 读侧推导的数据区起点（不落盘）
}

// colWriter 按 bin 攒列数据。bin 内每列：存在位图 + 一条压缩值流；值流解压后
// 是「uvarint(len) value」按记录序重复。长度与值同帧压缩，避免单独的、压不动
// 的长度表（chr22 抽样里它占列存总量 52%）。
type colWriter struct {
	fields   []colField
	idx      map[string]int
	binWidth uint32
	buildID  [16]byte

	f  *os.File
	w  *bufio.Writer
	at int64

	bins []colBinRecord

	started bool
	curBin  uint32
	first   uint64
	n       uint32
	bits    [][]byte // 每列 ceil(n/8) 字节位图
	vals    [][]byte // 每列未压值流
	present []int
	rawSum  []uint64 // 每列存在值的原文字节合计
}

// newColWriter 建写入器并预留 64B 头。
func newColWriter(f *os.File, fields []colField, binWidth uint32, buildID [16]byte) (*colWriter, error) {
	if binWidth == 0 {
		binWidth = colBinWidth
	}
	cw := &colWriter{
		fields: fields, idx: make(map[string]int, len(fields)),
		binWidth: binWidth, buildID: buildID, f: f,
		w:    bufio.NewWriterSize(f, 1<<20),
		bits: make([][]byte, len(fields)), vals: make([][]byte, len(fields)),
		present: make([]int, len(fields)), rawSum: make([]uint64, len(fields)),
	}
	for i := range fields {
		cw.idx[fields[i].Name] = i
	}
	if _, err := cw.w.Write(make([]byte, colHeaderSize)); err != nil {
		return nil, err
	}
	cw.at = colHeaderSize
	return cw, nil
}

// add 投入一条记录：pos 是记录位置，val 是记录序值（字段已按 ##INFO 序拼好）。
func (c *colWriter) add(rec uint64, pos uint32, val []byte) error {
	bin := pos / c.binWidth
	switch {
	case !c.started:
		// 首条记录前的空 bin 也要占位，bin 目录下标必须等于基因组 bin 号。
		for i := uint32(0); i < bin; i++ {
			c.bins = append(c.bins, colBinRecord{})
		}
		c.beginBin(bin, rec)
	case bin != c.curBin:
		if bin < c.curBin {
			return fmt.Errorf("immkv: base.col bin 倒退：%d < %d", bin, c.curBin)
		}
		if err := c.flushBin(); err != nil {
			return err
		}
		for c.curBin+1 < bin {
			c.bins = append(c.bins, colBinRecord{})
			c.curBin++
		}
		c.beginBin(bin, rec)
	}
	if rec != c.first+uint64(c.n) {
		return fmt.Errorf("immkv: base.col 记录号不连续：%d，期望 %d", rec, c.first+uint64(c.n))
	}
	if c.n%8 == 0 {
		for i := range c.bits {
			c.bits[i] = append(c.bits[i], 0)
		}
	}
	if len(val) > 0 && !(len(val) == 1 && val[0] == '.') {
		rest := val
		for len(rest) > 0 {
			var seg []byte
			if i := bytes.IndexByte(rest, ';'); i >= 0 {
				seg, rest = rest[:i], rest[i+1:]
			} else {
				seg, rest = rest, nil
			}
			if len(seg) == 0 {
				return errors.New("immkv: base.col: empty INFO segment")
			}
			eq := bytes.IndexByte(seg, '=')
			name := seg
			var v []byte
			if eq >= 0 {
				name, v = seg[:eq], seg[eq+1:]
				if eq == 0 {
					return errors.New("immkv: base.col: empty INFO field name")
				}
			}
			col, ok := c.idx[string(name)]
			if !ok {
				return fmt.Errorf("immkv: base.col: field %q not in schema", name)
			}
			if c.fields[col].Encoding == "bitmap" {
				// Flag：只有存在位图。
			} else {
				if eq < 0 {
					return fmt.Errorf("immkv: base.col: non-Flag field %q has no value", name)
				}
				c.vals[col] = binary.AppendUvarint(c.vals[col], uint64(len(v)))
				c.vals[col] = append(c.vals[col], v...)
				c.rawSum[col] += uint64(len(v))
			}
			c.bits[col][int(c.n)/8] |= 1 << (c.n % 8)
			c.present[col]++
		}
	}
	c.n++
	return nil
}

func (c *colWriter) beginBin(bin uint32, rec uint64) {
	c.started = true
	c.curBin = bin
	c.first = rec
	c.n = 0
	for i := range c.bits {
		c.bits[i] = c.bits[i][:0]
		c.vals[i] = c.vals[i][:0]
		c.present[i] = 0
		c.rawSum[i] = 0
	}
}

// flushBin 压缩并写当前 bin：每列 [位图][值流]，随后写列目录；bin 目录最后
// 统一落盘。
func (c *colWriter) flushBin() error {
	if !c.started || c.n == 0 {
		return nil
	}
	ncol := len(c.fields)
	compBit := make([][]byte, ncol)
	compVal := make([][]byte, ncol)
	err := parallelFor(ncol, func(i int) error {
		if c.present[i] == 0 {
			return nil
		}
		compBit[i] = colCompress(c.bits[i])
		if len(c.vals[i]) > 0 {
			compVal[i] = colCompress(c.vals[i])
		}
		return nil
	})
	if err != nil {
		return err
	}
	bitmapLen := int((c.n + 7) / 8)
	entries := make([]colEntry, ncol)
	for i := range c.fields {
		if c.present[i] == 0 {
			continue
		}
		if len(c.bits[i]) != bitmapLen {
			return fmt.Errorf("immkv: base.col: 位图长度 %d ≠ %d", len(c.bits[i]), bitmapLen)
		}
		entries[i].bitmapOff = uint64(c.at)
		entries[i].bitmapLen = uint32(len(compBit[i]))
		if err := c.write(compBit[i]); err != nil {
			return err
		}
		if len(compVal[i]) > 0 {
			entries[i].valueOff = uint64(c.at)
			entries[i].valueLen = uint32(len(compVal[i]))
			if err := c.write(compVal[i]); err != nil {
				return err
			}
		}
		entries[i].present = uint32(c.present[i])
		raw := c.rawSum[i]
		if raw > uint64(^uint32(0)) {
			raw = uint64(^uint32(0))
		}
		entries[i].rawLen = uint32(raw)
	}
	dirOff := uint64(c.at)
	var dir [colColumnEntrySize]byte
	for i := range entries {
		binary.LittleEndian.PutUint64(dir[0:8], entries[i].bitmapOff)
		binary.LittleEndian.PutUint64(dir[8:16], entries[i].valueOff)
		binary.LittleEndian.PutUint32(dir[16:20], entries[i].bitmapLen)
		binary.LittleEndian.PutUint32(dir[20:24], entries[i].valueLen)
		binary.LittleEndian.PutUint32(dir[24:28], entries[i].present)
		binary.LittleEndian.PutUint32(dir[28:32], entries[i].rawLen)
		if err := c.write(dir[:]); err != nil {
			return err
		}
	}
	c.bins = append(c.bins, colBinRecord{
		first: c.first, count: c.n, dirOff: dirOff, dataEnd: uint64(c.at),
	})
	c.started = false
	return nil
}

func (c *colWriter) write(p []byte) error {
	n, err := c.w.Write(p)
	c.at += int64(n)
	return err
}

// finish 落 bin 目录与文件头。
func (c *colWriter) finish() error {
	if c.started {
		if err := c.flushBin(); err != nil {
			return err
		}
	}
	binDirOff := uint64(c.at)
	var ent [colBinEntrySize]byte
	for i := range c.bins {
		b := c.bins[i]
		binary.LittleEndian.PutUint64(ent[0:8], b.first)
		binary.LittleEndian.PutUint32(ent[8:12], b.count)
		binary.LittleEndian.PutUint64(ent[16:24], b.dirOff)
		binary.LittleEndian.PutUint64(ent[24:32], b.dataEnd)
		if err := c.write(ent[:]); err != nil {
			return err
		}
	}
	if err := c.w.Flush(); err != nil {
		return err
	}
	h := make([]byte, colHeaderSize)
	copy(h[0:8], colMagic)
	binary.LittleEndian.PutUint32(h[8:12], colVersion)
	binary.LittleEndian.PutUint32(h[12:16], colHeaderSize)
	copy(h[16:32], c.buildID[:])
	binary.LittleEndian.PutUint32(h[32:36], c.binWidth)
	binary.LittleEndian.PutUint32(h[36:40], uint32(len(c.bins)))
	binary.LittleEndian.PutUint32(h[40:44], uint32(len(c.fields)))
	binary.LittleEndian.PutUint32(h[44:48], colColumnEntrySize)
	binary.LittleEndian.PutUint64(h[48:56], binDirOff)
	binary.LittleEndian.PutUint64(h[56:64], colSchemaHash(c.fields))
	_, err := c.f.WriteAt(h, 0)
	return err
}

// colEncoderPool 复用 zstd 编码器；EncodeAll 单次独占一个编码器。
var colEncoderPool = sync.Pool{New: func() any {
	enc, err := zstd.NewWriter(nil,
		zstd.WithEncoderLevel(zstd.EncoderLevelFromZstd(colZstdLevel)),
		zstd.WithEncoderConcurrency(1),
		zstd.WithEncoderCRC(false))
	if err != nil {
		panic(err)
	}
	return enc
}}

func colCompress(src []byte) []byte {
	enc := colEncoderPool.Get().(*zstd.Encoder)
	defer colEncoderPool.Put(enc)
	return enc.EncodeAll(src, nil)
}

// parallelFor 用 GOMAXPROCS 路并发执行 fn；返回首个错误。
func parallelFor(n int, fn func(int) error) error {
	return parallelForN(n, runtime.GOMAXPROCS(0), fn)
}

// parallelForN 用至多 workers 路并发执行 fn；返回首个错误。
func parallelForN(n, workers int, fn func(int) error) error {
	if workers > n {
		workers = n
	}
	if workers <= 1 {
		for i := 0; i < n; i++ {
			if err := fn(i); err != nil {
				return err
			}
		}
		return nil
	}
	var (
		wg   sync.WaitGroup
		mu   sync.Mutex
		ferr error
	)
	jobs := make(chan int)
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				if err := fn(i); err != nil {
					mu.Lock()
					if ferr == nil {
						ferr = err
					}
					mu.Unlock()
				}
			}
		}()
	}
	for i := 0; i < n; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	return ferr
}

// writeColumnar 为已完成的记录序分片生成 base.col 与 base.manifest。
// 失败时清掉半成品，不留 col 无 manifest 的配对。
func writeColumnar(base string, sel []valueField, buildID [16]byte, stats colInfoStats) (err error) {
	fields := colFieldsFromSel(sel)
	s, err := openShard(base, nil)
	if err != nil {
		return err
	}
	defer s.Close()

	colPath := base + ".col"
	tmp := colPath + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	cleanup := func() {
		f.Close()
		os.Remove(tmp)
		os.Remove(colPath)
		os.Remove(base + ".manifest")
		os.Remove(base + ".manifest.tmp")
	}
	cw, err := newColWriter(f, fields, colBinWidth, buildID)
	if err != nil {
		cleanup()
		return err
	}
	vr, err := newSeqValReader(base+".vals", s.dec)
	if err != nil {
		cleanup()
		return err
	}
	defer vr.Close()
	err = scanShardValues(s, vr, func(rec uint64, key, val []byte) error {
		if len(key) < 5 {
			return fmt.Errorf("immkv: base.col: 键 %x 不足 5 字节", key)
		}
		pos := binary.BigEndian.Uint32(key[1:5])
		return cw.add(rec, pos, val)
	})
	if err == nil {
		err = cw.finish()
	}
	if err == nil {
		err = f.Sync()
	}
	if cerr := f.Close(); err == nil {
		err = cerr
	}
	if err != nil {
		cleanup()
		return err
	}
	if err := os.Rename(tmp, colPath); err != nil {
		cleanup()
		return err
	}
	man, err := buildColManifest(base, fields, buildID, s.count, stats)
	if err != nil {
		cleanup()
		return err
	}
	if err := writeFileAtomic(base+".manifest", man); err != nil {
		cleanup()
		return err
	}
	return nil
}

// buildColManifest 汇总配对信息与 schema，写成 JSON。
func buildColManifest(base string, fields []colField, buildID [16]byte, count uint64, stats colInfoStats) ([]byte, error) {
	stK, err := os.Stat(base + ".keys")
	if err != nil {
		return nil, err
	}
	stV, err := os.Stat(base + ".vals")
	if err != nil {
		return nil, err
	}
	stC, err := os.Stat(base + ".col")
	if err != nil {
		return nil, err
	}
	m := colManifest{
		Format: colManifestFormat, Version: colManifestVersion,
		Base: filepath.Base(base), BuildID: hex.EncodeToString(buildID[:]),
		BinWidth: colBinWidth, RecordCount: count,
		KeysBytes: stK.Size(), ValsBytes: stV.Size(), ColBytes: stC.Size(),
		SchemaHash: colSchemaString(fields), SourceInfoOrder: stats,
		Fields: fields,
	}
	return json.MarshalIndent(m, "", "  ")
}

// newBuildID 产生 16 字节随机构建代号。
func newBuildID() ([16]byte, error) {
	var id [16]byte
	if _, err := rand.Read(id[:]); err != nil {
		return id, err
	}
	return id, nil
}

// scanShardValues 顺序扫记录序分片，把每条记录的键与值交给 fn。值只在
// fn 返回前有效。
func scanShardValues(s *shard, vr *seqValReader, fn func(rec uint64, key, val []byte) error) error {
	for b := 0; b < int(s.nBlock); b++ {
		e := s.block(b)
		if e == nil {
			return fmt.Errorf("immkv: base.col: 读微块 %d 失败", b)
		}
		n := s.nRecs(b)
		for j := 0; j < n; j++ {
			rec := uint64(b)*s.recsBlk + uint64(j)
			vo, fl, raw, packed := e.valLoc(j)
			val, err := vr.value(rec, vo, fl, raw, packed)
			if err != nil {
				return err
			}
			if err := fn(rec, e.keyAt(j), val); err != nil {
				return err
			}
		}
	}
	return nil
}

// seqValReader 顺序读 vals：记录序的帧按记录序物理排列，按 valOff 前进即可；
// 打包帧只解一次，供同帧的后续记录复用。
type seqValReader struct {
	f       *os.File
	r       *bufio.Reader
	dec     frameDecoder
	off     uint64
	buf     []byte
	pack    *packEnt
	packOff uint64
}

func newSeqValReader(path string, dec frameDecoder) (*seqValReader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return &seqValReader{f: f, r: bufio.NewReaderSize(f, 1<<20), dec: dec}, nil
}

func (v *seqValReader) Close() error { return v.f.Close() }

func (v *seqValReader) value(rec uint64, valOff uint64, frameLen uint32, raw, packed bool) ([]byte, error) {
	if frameLen == 0 {
		return emptyValue, nil
	}
	if packed && v.pack != nil && valOff == v.packOff {
		idx := rec - uint64(v.pack.first)
		if idx >= uint64(v.pack.n) {
			return nil, fmt.Errorf("immkv: 打包帧记录号 %d 越界", rec)
		}
		return v.pack.value(uint32(idx)), nil
	}
	if valOff != v.off {
		if valOff < v.off {
			return nil, fmt.Errorf("immkv: 值偏移倒退：%d < %d", valOff, v.off)
		}
		if _, err := v.f.Seek(int64(valOff), io.SeekStart); err != nil {
			return nil, err
		}
		v.r.Reset(v.f)
		v.off = valOff
	}
	if uint64(cap(v.buf)) < uint64(frameLen) {
		v.buf = make([]byte, frameLen)
	}
	frame := v.buf[:frameLen]
	if _, err := io.ReadFull(v.r, frame); err != nil {
		return nil, err
	}
	v.off += uint64(frameLen)
	v.pack = nil
	if packed {
		p, err := decodePack(v.dec, frame, nil)
		if err != nil {
			return nil, err
		}
		v.pack, v.packOff = p, valOff
		idx := rec - uint64(p.first)
		if idx >= uint64(p.n) {
			return nil, fmt.Errorf("immkv: 打包帧记录号 %d 越界", rec)
		}
		return p.value(uint32(idx)), nil
	}
	if raw {
		return frame, nil
	}
	// 单值帧：用分片的解码器（可能带构建期字典）。
	return v.dec.decode(frame, nil)
}
