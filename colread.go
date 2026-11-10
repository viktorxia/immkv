// colread.go 是 base.col 的读取侧：给 OpenStreaming 的顺序批量查提供列存
// 路径。读侧按命中 bin 只读选中的列，解压后按 schema 顺序拼值；点查、
// 前缀、区间扫描不使用本文件。

package immkv

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/bits"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/klauspost/compress/zstd"
)

// colShard 是一个染色体的 base.col 读句柄。
type colShard struct {
	base     string
	f        *os.File
	size     int64
	binWidth uint32
	buildID  string
	fields   []colField
	bins     []colBinRecord
	dec      frameDecoder

	mu   sync.Mutex
	dirs map[int][]byte // bin 列目录缓存（非空 bin）
}

// openColShard 打开 base.col + base.manifest。两者都不存在返回 (nil, nil)；
// 只存在一个、或任何一个校验不过都返回错误，不回退。
func openColShard(base string, ks *shard) (*colShard, error) {
	colPath, manPath := base+".col", base+".manifest"
	_, errC := os.Stat(colPath)
	_, errM := os.Stat(manPath)
	colOK, manOK := errC == nil, errM == nil
	if !colOK && !manOK {
		return nil, nil
	}
	if colOK != manOK {
		return nil, fmt.Errorf("immkv: %s 的 base.col 与 base.manifest 不配对", base)
	}
	raw, err := os.ReadFile(manPath)
	if err != nil {
		return nil, err
	}
	var m colManifest
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("immkv: %s: 解析 manifest: %w", base, err)
	}
	if m.Format != colManifestFormat || m.Version != colManifestVersion {
		return nil, fmt.Errorf("immkv: %s: manifest format/version 不合法", base)
	}
	if m.Base != filepath.Base(base) {
		return nil, fmt.Errorf("immkv: %s: manifest base %q 不匹配", base, m.Base)
	}
	if len(m.Fields) == 0 {
		return nil, fmt.Errorf("immkv: %s: manifest 没有字段", base)
	}
	seen := make(map[string]bool, len(m.Fields))
	for i := range m.Fields {
		f := &m.Fields[i]
		if f.Column != i || f.Name == "" || f.Type == "" {
			return nil, fmt.Errorf("immkv: %s: manifest 字段 %d 不合法", base, i)
		}
		if f.Encoding != "raw" && f.Encoding != "bitmap" {
			return nil, fmt.Errorf("immkv: %s: 字段 %q 编码 %q 未知", base, f.Name, f.Encoding)
		}
		if seen[f.Name] {
			return nil, fmt.Errorf("immkv: %s: manifest 字段 %q 重复", base, f.Name)
		}
		seen[f.Name] = true
	}
	if colSchemaString(m.Fields) != m.SchemaHash {
		return nil, fmt.Errorf("immkv: %s: manifest schema hash 不匹配", base)
	}

	f, err := os.Open(colPath)
	if err != nil {
		return nil, err
	}
	fail := func(err error) (*colShard, error) {
		f.Close()
		return nil, err
	}
	st, err := f.Stat()
	if err != nil {
		return fail(err)
	}
	var h [colHeaderSize]byte
	if _, err := f.ReadAt(h[:], 0); err != nil {
		return fail(err)
	}
	if string(h[0:8]) != colMagic {
		return fail(fmt.Errorf("immkv: %s: bad base.col magic", base))
	}
	if v := binary.LittleEndian.Uint32(h[8:12]); v != colVersion {
		return fail(fmt.Errorf("immkv: %s: base.col version %d ≠ %d", base, v, colVersion))
	}
	if v := binary.LittleEndian.Uint32(h[12:16]); v != colHeaderSize {
		return fail(fmt.Errorf("immkv: %s: base.col header_size %d", base, v))
	}
	if buildID := hex.EncodeToString(h[16:32]); buildID != m.BuildID {
		return fail(fmt.Errorf("immkv: %s: build id 不匹配", base))
	}
	binWidth := binary.LittleEndian.Uint32(h[32:36])
	binCount := binary.LittleEndian.Uint32(h[36:40])
	colCount := binary.LittleEndian.Uint32(h[40:44])
	entrySize := binary.LittleEndian.Uint32(h[44:48])
	binDirOff := binary.LittleEndian.Uint64(h[48:56])
	schemaHash := binary.LittleEndian.Uint64(h[56:64])
	if binWidth != m.BinWidth || binWidth == 0 {
		return fail(fmt.Errorf("immkv: %s: bin_width %d 与 manifest 不符", base, binWidth))
	}
	if int(colCount) != len(m.Fields) {
		return fail(fmt.Errorf("immkv: %s: column_count %d ≠ manifest %d", base, colCount, len(m.Fields)))
	}
	if entrySize != colColumnEntrySize {
		return fail(fmt.Errorf("immkv: %s: column_entry_size %d", base, entrySize))
	}
	if schemaHash != colSchemaHash(m.Fields) {
		return fail(fmt.Errorf("immkv: %s: schema hash 与 manifest 不符", base))
	}
	if st.Size() != m.ColBytes || m.KeysBytes <= 0 || m.ValsBytes <= 0 {
		return fail(fmt.Errorf("immkv: %s: 文件字节数与 manifest 不符", base))
	}
	if ks.count != m.RecordCount {
		return fail(fmt.Errorf("immkv: %s: 记录数 %d ≠ manifest %d", base, ks.count, m.RecordCount))
	}
	if uint64(binDirOff)+uint64(binCount)*colBinEntrySize > uint64(st.Size()) {
		return fail(fmt.Errorf("immkv: %s: bin 目录越界", base))
	}
	dir := make([]byte, uint64(binCount)*colBinEntrySize)
	if _, err := f.ReadAt(dir, int64(binDirOff)); err != nil {
		return fail(err)
	}
	bins := make([]colBinRecord, binCount)
	var wantFirst uint64
	var sum uint32
	dataStart := uint64(colHeaderSize)
	dirBytes := uint64(colCount) * colColumnEntrySize
	for i := 0; i < int(binCount); i++ {
		e := dir[i*colBinEntrySize:]
		b := &bins[i]
		b.first = binary.LittleEndian.Uint64(e[0:8])
		b.count = binary.LittleEndian.Uint32(e[8:12])
		b.dirOff = binary.LittleEndian.Uint64(e[16:24])
		b.dataEnd = binary.LittleEndian.Uint64(e[24:32])
		if b.count == 0 {
			if b.first != 0 || b.dirOff != 0 || b.dataEnd != 0 {
				return fail(fmt.Errorf("immkv: %s: 空 bin %d 非法", base, i))
			}
			continue
		}
		if b.first != wantFirst {
			return fail(fmt.Errorf("immkv: %s: bin %d first_record %d ≠ %d", base, i, b.first, wantFirst))
		}
		b.start = dataStart
		if b.dataEnd < b.start || b.dataEnd > uint64(st.Size()) {
			return fail(fmt.Errorf("immkv: %s: bin %d 数据区越界", base, i))
		}
		if b.dirOff < b.start || b.dirOff+dirBytes > b.dataEnd {
			return fail(fmt.Errorf("immkv: %s: bin %d 列目录越界", base, i))
		}
		dataStart = b.dataEnd
		wantFirst += uint64(b.count)
		sum += b.count
	}
	if dataStart > uint64(binDirOff) {
		return fail(fmt.Errorf("immkv: %s: 数据区越过 bin 目录", base))
	}
	if uint64(sum) != m.RecordCount {
		return fail(fmt.Errorf("immkv: %s: bin 记录数合计 %d ≠ %d", base, sum, m.RecordCount))
	}

	var dec frameDecoder
	if d, derr := newCGODecoder(); derr == nil {
		dec = d
	} else {
		kd, kerr := zstd.NewReader(nil, zstd.IgnoreChecksum(true))
		if kerr != nil {
			return fail(kerr)
		}
		dec = &kpDecoder{dec: kd}
	}
	return &colShard{
		base: base, f: f, size: st.Size(), binWidth: binWidth,
		buildID: m.BuildID, fields: m.Fields, bins: bins, dec: dec,
		dirs: make(map[int][]byte),
	}, nil
}

func (c *colShard) Close() error {
	var err error
	if c.dec != nil {
		err = c.dec.Close()
	}
	if ce := c.f.Close(); err == nil {
		err = ce
	}
	return err
}

// loadDir 返回一个 bin 的列目录（column_count × 32B）；空 bin 返回 nil。
func (c *colShard) loadDir(bi int) ([]byte, error) {
	b := &c.bins[bi]
	if b.count == 0 {
		return nil, nil
	}
	c.mu.Lock()
	if d, ok := c.dirs[bi]; ok {
		c.mu.Unlock()
		return d, nil
	}
	c.mu.Unlock()
	n := len(c.fields) * colColumnEntrySize
	if b.dirOff+uint64(n) > uint64(c.size) {
		return nil, fmt.Errorf("immkv: %s: bin %d 列目录越界", c.base, bi)
	}
	d := make([]byte, n)
	if _, err := c.f.ReadAt(d, int64(b.dirOff)); err != nil {
		return nil, err
	}
	// 列目录必须落在本 bin 的数据区内，避免损坏的偏移把别的 bin 数据当成
	// 本 bin 的值而静默产出错误结果。
	for col := 0; col < len(c.fields); col++ {
		e := c.entry(d, col)
		if e.bitmapOff == 0 {
			if e.bitmapLen != 0 || e.valueOff != 0 || e.valueLen != 0 || e.present != 0 || e.rawLen != 0 {
				return nil, fmt.Errorf("immkv: %s: bin %d 列 %d 空条目非法", c.base, bi, col)
			}
			continue
		}
		if e.bitmapLen == 0 || e.bitmapOff < b.start ||
			e.bitmapOff+uint64(e.bitmapLen) > b.dataEnd {
			return nil, fmt.Errorf("immkv: %s: bin %d 列 %d 位图越界", c.base, bi, col)
		}
		if e.present == 0 || e.present > b.count {
			return nil, fmt.Errorf("immkv: %s: bin %d 列 %d present_count %d 非法", c.base, bi, col, e.present)
		}
		if c.fields[col].Encoding == "bitmap" {
			if e.valueOff != 0 || e.valueLen != 0 || e.rawLen != 0 {
				return nil, fmt.Errorf("immkv: %s: bin %d 列 %d Flag 不应有值流", c.base, bi, col)
			}
			continue
		}
		if e.valueOff != e.bitmapOff+uint64(e.bitmapLen) || e.valueLen == 0 ||
			e.valueOff+uint64(e.valueLen) > b.dataEnd {
			return nil, fmt.Errorf("immkv: %s: bin %d 列 %d 值流越界", c.base, bi, col)
		}
	}
	c.mu.Lock()
	if old, ok := c.dirs[bi]; ok {
		d = old
	} else {
		c.dirs[bi] = d
	}
	c.mu.Unlock()
	return d, nil
}

func (c *colShard) entry(dir []byte, col int) colEntry {
	e := dir[col*colColumnEntrySize:]
	return colEntry{
		bitmapOff: binary.LittleEndian.Uint64(e[0:8]),
		valueOff:  binary.LittleEndian.Uint64(e[8:16]),
		bitmapLen: binary.LittleEndian.Uint32(e[16:20]),
		valueLen:  binary.LittleEndian.Uint32(e[20:24]),
		present:   binary.LittleEndian.Uint32(e[24:28]),
		rawLen:    binary.LittleEndian.Uint32(e[28:32]),
	}
}

// selectedColumns 把 Options.Columns 展开为 schema 序的列号；nil/空 = 全部。
func selectedColumns(fields []colField, sel []int) []int {
	if len(sel) == 0 {
		out := make([]int, len(fields))
		for i := range out {
			out[i] = i
		}
		return out
	}
	out := append([]int(nil), sel...)
	sort.Ints(out)
	return out
}

// hitBinGroup 是一批命中里落在同一个 bin 的记录。
type hitBinGroup struct {
	bin  int
	hits []sHit
}

// groupHits 把按记录号升序的命中切成 bin 组。hits 必须已按记录号非降序。
func (c *colShard) groupHits(hits []sHit) ([]hitBinGroup, error) {
	var out []hitBinGroup
	bi := 0
	for i := 0; i < len(hits); {
		rec := hits[i].rec
		for bi < len(c.bins) {
			b := &c.bins[bi]
			if b.count == 0 || rec >= int64(b.first)+int64(b.count) {
				bi++
				continue
			}
			break
		}
		if bi >= len(c.bins) {
			return nil, fmt.Errorf("immkv: %s: 记录号 %d 不在任何 bin 内", c.base, rec)
		}
		b := &c.bins[bi]
		if rec < int64(b.first) || rec >= int64(b.first)+int64(b.count) {
			return nil, fmt.Errorf("immkv: %s: 记录号 %d 不在 bin %d 内", c.base, rec, bi)
		}
		end := int64(b.first) + int64(b.count)
		j := i
		for j < len(hits) && hits[j].rec < end {
			j++
		}
		out = append(out, hitBinGroup{bin: bi, hits: hits[i:j]})
		i = j
	}
	return out, nil
}

// estimate 估算一次列存查询要读的压缩字节与要解压的原文字节：每个命中 bin
// 的选中列位图、值流与 rawLen 合计（列目录常驻，不计）。
func (c *colShard) estimate(hits []sHit, sel []int) (compressed, raw uint64, err error) {
	groups, err := c.groupHits(hits)
	if err != nil {
		return 0, 0, err
	}
	for _, g := range groups {
		dir, err := c.loadDir(g.bin)
		if err != nil {
			return 0, 0, err
		}
		for _, col := range sel {
			e := c.entry(dir, col)
			if e.bitmapOff == 0 {
				continue
			}
			compressed += uint64(e.bitmapLen) + uint64(e.valueLen)
			raw += uint64(e.rawLen)
		}
	}
	return compressed, raw, nil
}

// colPart 是一次 bin 读取里的一列：压缩位图与值流在合并读段中的位置。
type colPart struct {
	col       int
	flag      bool
	present   int
	bitmapAt  int64
	bitmapLen int
	frameAt   int64
	frameLen  int
}

// colCursor 是列内「记录号 → 值流位置」的游标；跨装配分块保持，保证每个
// 列流每 bin 只顺扫一遍。
type colCursor struct {
	lastBit     int
	lastRidx    int
	lastVal     []byte
	lastPresent bool
	pos         int
}

// colSegs 是一列在一个装配分块里对每条命中的输出段（空 = 字段缺失）。
type colSegs struct {
	buf []byte
	off []int32
}

// asmChunk 是单次装配的命中数上限，控制段缓冲与偏移数组的内存。
const asmChunk = 8192

// readBin 读一个 bin 的选中列并回填 out。读段按偏移合并，顺序 pread；
// 位图与值流都用 threads 路并行解压。
func (c *colShard) readBin(g hitBinGroup, sel []int, out [][]byte, threads int) error {
	meta := &c.bins[g.bin]
	dir, err := c.loadDir(g.bin)
	if err != nil {
		return err
	}
	bmLen := int((meta.count + 7) / 8)
	parts := make([]colPart, 0, len(sel))
	for _, col := range sel {
		e := c.entry(dir, col)
		if e.bitmapOff == 0 {
			continue
		}
		p := colPart{
			col: col, flag: c.fields[col].Encoding == "bitmap",
			present: int(e.present), bitmapAt: int64(e.bitmapOff),
			bitmapLen: int(e.bitmapLen),
		}
		if !p.flag {
			p.frameAt = int64(e.valueOff)
			p.frameLen = int(e.valueLen)
		}
		parts = append(parts, p)
	}
	if len(parts) == 0 {
		for _, h := range g.hits {
			out[h.dst] = []byte(".")
		}
		return nil
	}
	sort.Slice(parts, func(i, j int) bool { return parts[i].bitmapAt < parts[j].bitmapAt })
	type readRange struct {
		start, end int64
		parts      []int
	}
	var ranges []readRange
	for i := range parts {
		p := &parts[i]
		end := p.bitmapAt + int64(p.bitmapLen)
		if p.frameLen > 0 {
			end = p.frameAt + int64(p.frameLen)
		}
		if len(ranges) > 0 && p.bitmapAt <= ranges[len(ranges)-1].end {
			r := &ranges[len(ranges)-1]
			if end > r.end {
				r.end = end
			}
			r.parts = append(r.parts, i)
			continue
		}
		ranges = append(ranges, readRange{start: p.bitmapAt, end: end, parts: []int{i}})
	}
	bitFrames := make([][]byte, len(parts))
	frames := make([][]byte, len(parts))
	bits := make([][]byte, len(parts))
	decoded := make([][]byte, len(parts))
	for ri := range ranges {
		r := &ranges[ri]
		buf := make([]byte, r.end-r.start)
		if _, err := c.f.ReadAt(buf, r.start); err != nil {
			return fmt.Errorf("immkv: %s: 读 bin %d 列数据: %w", c.base, g.bin, err)
		}
		fadviseDontNeed(c.f, r.start, r.end-r.start)
		for _, pi := range r.parts {
			p := &parts[pi]
			off := p.bitmapAt - r.start
			bitFrames[pi] = buf[off : off+int64(p.bitmapLen)]
			if p.frameLen > 0 {
				fo := p.frameAt - r.start
				frames[pi] = buf[fo : fo+int64(p.frameLen)]
			}
		}
	}
	var wg sync.WaitGroup
	errs := make(chan error, len(parts))
	sem := make(chan struct{}, threads)
	for i := range parts {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int) {
			defer wg.Done()
			defer func() { <-sem }()
			b, err := c.dec.decode(bitFrames[i], nil)
			if err != nil {
				errs <- fmt.Errorf("immkv: %s: 解压 bin %d 列 %d 位图: %w", c.base, g.bin, parts[i].col, err)
				return
			}
			if len(b) != bmLen {
				errs <- fmt.Errorf("immkv: %s: bin %d 列 %d 位图长度 %d ≠ %d", c.base, g.bin, parts[i].col, len(b), bmLen)
				return
			}
			if pc := popcountBytes(b); pc != parts[i].present {
				errs <- fmt.Errorf("immkv: %s: bin %d 列 %d 位图 %d 位 ≠ present_count %d",
					c.base, g.bin, parts[i].col, pc, parts[i].present)
				return
			}
			bits[i] = b
			if parts[i].flag {
				return
			}
			v, err := c.dec.decode(frames[i], nil)
			if err != nil {
				errs <- fmt.Errorf("immkv: %s: 解压 bin %d 列 %d 值流: %w", c.base, g.bin, parts[i].col, err)
				return
			}
			decoded[i] = v
		}(i)
	}
	wg.Wait()
	close(errs)
	if err := <-errs; err != nil {
		return err
	}
	// 按列并行扫描：每列在自己的位图/值流上顺扫，产出「本分块每条命中的
	// 输出段」；最后按 schema 顺序合并。列内值顺序 = bin 内存在记录顺序。
	cursors := make([]colCursor, len(parts))
	for i := range cursors {
		cursors[i].lastBit = -1
		cursors[i].lastRidx = -2
	}
	for start := 0; start < len(g.hits); start += asmChunk {
		end := start + asmChunk
		if end > len(g.hits) {
			end = len(g.hits)
		}
		hs := g.hits[start:end]
		segs := make([]colSegs, len(parts))
		err := parallelForN(len(parts), threads, func(pi int) error {
			p := &parts[pi]
			field := &c.fields[p.col]
			cur := &cursors[pi]
			stream := decoded[pi]
			seg := colSegs{off: make([]int32, len(hs)+1)}
			for hi, h := range hs {
				ridx := int(h.rec - int64(meta.first))
				if ridx < 0 || ridx >= int(meta.count) {
					return fmt.Errorf("immkv: %s: 记录号 %d 不在 bin %d 内", c.base, h.rec, g.bin)
				}
				if ridx != cur.lastRidx {
					for cur.lastBit < ridx {
						cur.lastBit++
						present := bitSet(bits[pi], cur.lastBit)
						var v []byte
						if present && !p.flag {
							l, n := binary.Uvarint(stream[cur.pos:])
							if n <= 0 {
								return fmt.Errorf("immkv: %s: bin %d 列 %d 值流长度非法", c.base, g.bin, p.col)
							}
							cur.pos += n
							if l > uint64(len(stream)-cur.pos) {
								return fmt.Errorf("immkv: %s: bin %d 列 %d 值超出值流", c.base, g.bin, p.col)
							}
							v = stream[cur.pos : cur.pos+int(l)]
							cur.pos += int(l)
						}
						if cur.lastBit == ridx {
							cur.lastVal, cur.lastPresent = v, present
						}
					}
					cur.lastRidx = ridx
				}
				if cur.lastPresent {
					seg.buf = append(seg.buf, field.Name...)
					if !p.flag {
						seg.buf = append(seg.buf, '=')
						seg.buf = append(seg.buf, cur.lastVal...)
					}
				}
				seg.off[hi+1] = int32(len(seg.buf))
			}
			segs[pi] = seg
			return nil
		})
		if err != nil {
			return err
		}
		for hi, h := range hs {
			var o []byte
			for pi := range segs {
				s := segs[pi].buf[segs[pi].off[hi]:segs[pi].off[hi+1]]
				if len(s) == 0 {
					continue
				}
				if len(o) > 0 {
					o = append(o, ';')
				}
				o = append(o, s...)
			}
			if o == nil {
				o = []byte(".")
			}
			out[h.dst] = o
		}
	}
	return nil
}

// readColBatch 走列存路径回填 out；hits 按记录号升序。
func (c *colShard) readColBatch(hits []sHit, sel []int, out [][]byte, threads int) error {
	groups, err := c.groupHits(hits)
	if err != nil {
		return err
	}
	for _, g := range groups {
		if err := c.readBin(g, sel, out, threads); err != nil {
			return err
		}
	}
	return nil
}

func bitSet(b []byte, i int) bool {
	return b[i>>3]&(1<<(i&7)) != 0
}

// sameColSchema 校验两个染色体的列 schema 完全一致（同一构建应如此）。
func sameColSchema(a, b []colField) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Name != b[i].Name || a[i].Type != b[i].Type ||
			a[i].Number != b[i].Number || a[i].Encoding != b[i].Encoding ||
			a[i].Column != b[i].Column {
			return false
		}
	}
	return true
}

func popcountBytes(b []byte) int {
	n := 0
	for _, v := range b {
		n += bits.OnesCount8(v)
	}
	return n
}
