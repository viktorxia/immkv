// col.go 是 base.col 的写入侧：把一个已完成的记录序分片（base.keys/base.vals）
// 顺序扫一遍，按基因组区间切 bin，把每条 record 的字段按 schema 列拆开，
// 压缩后写成 base.col + base.manifest。格式契约见 docs/v003-format.md §5。
//
// 列存只服务 OpenStreaming 的顺序批量查；点查、前缀、区间扫描仍走记录序。
// 写入发生在记录序 Finish 之后，不改变键流与值帧格式。

package immkv

import (
	"bufio"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sync"

	"github.com/klauspost/compress/zstd"
)

const (
	colMagic           = "IMKVC003"
	colVersion         = 3
	colHeaderSize      = 64
	colBinEntrySize    = 32
	colColumnEntrySize = 32
	colBinWidth        = 1_000_000 // VCF 库固定 1Mb bin
	colZstdLevel       = 9         // 与抽样一致；值流按列独立压缩
)

// simpleKVSchema 是 BuildFromTSVs 的库 schema：单 String 字段 value。
func simpleKVSchema() Schema {
	return Schema{{Name: "value", Type: TypeString, Number: "1"}}
}

// colEntry 是列目录里的一项（32B）。
type colEntry struct {
	bitmapOff uint64
	valueOff  uint64
	bitmapLen uint32 // 位图 zstd 帧长度
	valueLen  uint32
	present   uint32
	rawLen    uint32 // 存在值的 typed payload 字节合计（饱和到 u32，仅用于选路估算）
}

// colBinRecord 是 bin 目录里的一项（32B）。
type colBinRecord struct {
	first   uint64
	count   uint32
	dirOff  uint64
	dataEnd uint64
	start   uint64 // 读侧推导的数据区起点（不落盘）
}

// colWriter 按 bin 攒列数据。bin 内每列：存在位图 + 一条压缩值流；值流是
// record 里该字段的编码 payload 按记录序直接拼接（自定界，不再套长度前缀）。
type colWriter struct {
	schema   Schema
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
	rawSum  []uint64 // 每列存在值的 payload 字节合计
}

// newColWriter 建写入器并预留 64B 头。binWidth 为 0 时全部记录进单个 bin
// （无位置语义的 KV 库）。
func newColWriter(f *os.File, schema Schema, binWidth uint32, buildID [16]byte) (*colWriter, error) {
	if err := schema.Validate(); err != nil {
		return nil, err
	}
	cw := &colWriter{
		schema: schema, binWidth: binWidth, buildID: buildID, f: f,
		w:    bufio.NewWriterSize(f, 1<<20),
		bits: make([][]byte, len(schema)), vals: make([][]byte, len(schema)),
		present: make([]int, len(schema)), rawSum: make([]uint64, len(schema)),
	}
	if _, err := cw.w.Write(make([]byte, colHeaderSize)); err != nil {
		return nil, err
	}
	cw.at = colHeaderSize
	return cw, nil
}

// add 投入一条记录：rec 是全局记录号，pos 是记录位置（binWidth 为 0 时忽略），
// record 是记录序里的 v003 record。字段 payload 原样进入对应列。
func (c *colWriter) add(rec uint64, pos uint32, record []byte) error {
	var bin uint32
	if c.binWidth != 0 {
		bin = pos / c.binWidth
	}
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
	err := walkRecord(c.schema, record, func(index int, payload []byte) error {
		c.bits[index][int(c.n)/8] |= 1 << (c.n % 8)
		c.present[index]++
		if c.schema[index].Type == TypeFlag {
			if len(payload) != 0 {
				return fmt.Errorf("immkv: base.col: Flag 字段 %q 带 %d 字节 payload",
					c.schema[index].Name, len(payload))
			}
			return nil
		}
		c.vals[index] = append(c.vals[index], payload...)
		c.rawSum[index] += uint64(len(payload))
		return nil
	})
	if err != nil {
		return fmt.Errorf("immkv: base.col 记录 %d: %w", rec, err)
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
	ncol := len(c.schema)
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
	for i := range c.schema {
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
		if c.schema[i].Type != TypeFlag {
			if len(compVal[i]) == 0 {
				return fmt.Errorf("immkv: base.col: 字段 %q 有形变值但值流为空", c.schema[i].Name)
			}
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
	binary.LittleEndian.PutUint32(h[40:44], uint32(len(c.schema)))
	binary.LittleEndian.PutUint32(h[44:48], colColumnEntrySize)
	binary.LittleEndian.PutUint64(h[48:56], binDirOff)
	hash := c.schema.Hash()
	binary.LittleEndian.PutUint64(h[56:64], binary.LittleEndian.Uint64(hash[:]))
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
// binWidth 为 0 表示无位置语义的单 bin 库；nonFinite 是构建期按缺失处理的
// 非有限 Float 次数，只写进 manifest 作诊断。失败时清掉半成品，不留 col 无
// manifest 的配对。
func writeColumnar(base string, schema Schema, buildID [16]byte, stats SourceInfoOrder, nonFinite int64, binWidth uint32) (err error) {
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
	cw, err := newColWriter(f, schema, binWidth, buildID)
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
	err = scanShardValues(s, vr, func(rec uint64, key, record []byte) error {
		var pos uint32
		if binWidth != 0 {
			if len(key) < 5 {
				return fmt.Errorf("immkv: base.col: 键 %x 不足 5 字节", key)
			}
			pos = binary.BigEndian.Uint32(key[1:5])
		}
		return cw.add(rec, pos, record)
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
	m, err := NewManifest(filepath.Base(base), schema)
	if err != nil {
		cleanup()
		return err
	}
	m.BuildID = hex.EncodeToString(buildID[:])
	m.BinWidth = binWidth
	m.RecordCount = s.count
	m.SourceInfoOrder = stats
	m.NonFiniteFloats = nonFinite
	stK, err := os.Stat(base + ".keys")
	if err != nil {
		cleanup()
		return err
	}
	stV, err := os.Stat(base + ".vals")
	if err != nil {
		cleanup()
		return err
	}
	stC, err := os.Stat(colPath)
	if err != nil {
		cleanup()
		return err
	}
	m.KeysBytes = stK.Size()
	m.ValsBytes = stV.Size()
	m.ColBytes = stC.Size()
	if err := WriteManifest(base+".manifest", m); err != nil {
		cleanup()
		return err
	}
	return nil
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
