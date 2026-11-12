// colread.go 是 base.col 的读取侧：给 OpenStreaming 的顺序批量查提供列存
// 路径。读侧按命中 bin 只读选中的列，解压后按 schema 顺序拼回 v003 record；
// 未选中的字段位为 0。点查、前缀、区间扫描不使用本文件。

package immkv

import (
	"encoding/binary"
	"encoding/hex"
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
	schema   Schema
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
	m, err := ReadManifest(manPath)
	if err != nil {
		return nil, err
	}
	if m.Base != filepath.Base(base) {
		return nil, fmt.Errorf("immkv: %s: manifest base %q 不匹配", base, m.Base)
	}
	schema, err := m.Schema()
	if err != nil {
		return nil, fmt.Errorf("immkv: %s: %w", base, err)
	}
	stK, err := os.Stat(base + ".keys")
	if err != nil {
		return nil, err
	}
	stV, err := os.Stat(base + ".vals")
	if err != nil {
		return nil, err
	}
	if stK.Size() != m.KeysBytes || stV.Size() != m.ValsBytes {
		return nil, fmt.Errorf("immkv: %s: keys/vals 字节数与 manifest 不符", base)
	}
	if ks.count != m.RecordCount {
		return nil, fmt.Errorf("immkv: %s: 记录数 %d ≠ manifest %d", base, ks.count, m.RecordCount)
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
	if binWidth != m.BinWidth {
		return fail(fmt.Errorf("immkv: %s: bin_width %d 与 manifest 不符", base, binWidth))
	}
	if binWidth == 0 {
		if binCount != 1 && !(binCount == 0 && m.RecordCount == 0) {
			return fail(fmt.Errorf("immkv: %s: bin_width 0 时 bin_count 应为 1，得到 %d", base, binCount))
		}
	}
	if int(colCount) != len(schema) {
		return fail(fmt.Errorf("immkv: %s: column_count %d ≠ manifest %d", base, colCount, len(schema)))
	}
	if entrySize != colColumnEntrySize {
		return fail(fmt.Errorf("immkv: %s: column_entry_size %d", base, entrySize))
	}
	hash := schema.Hash()
	if schemaHash != binary.LittleEndian.Uint64(hash[:]) {
		return fail(fmt.Errorf("immkv: %s: schema hash 与 manifest 不符", base))
	}
	if st.Size() != m.ColBytes {
		return fail(fmt.Errorf("immkv: %s: col 字节数与 manifest 不符", base))
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
		buildID: m.BuildID, schema: schema, bins: bins, dec: dec,
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
	n := len(c.schema) * colColumnEntrySize
	if b.dirOff+uint64(n) > uint64(c.size) {
		return nil, fmt.Errorf("immkv: %s: bin %d 列目录越界", c.base, bi)
	}
	d := make([]byte, n)
	if _, err := c.f.ReadAt(d, int64(b.dirOff)); err != nil {
		return nil, err
	}
	// 列目录必须落在本 bin 的数据区内，避免损坏的偏移把别的 bin 数据当成
	// 本 bin 的值而静默产出错误结果。
	for col := 0; col < len(c.schema); col++ {
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
		if c.schema[col].Type == TypeFlag {
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
func selectedColumns(schema Schema, sel []int) []int {
	if len(sel) == 0 {
		out := make([]int, len(schema))
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

// estimate 估算一次列存查询要读的压缩字节与要解压的 typed payload 字节：每个
// 命中 bin 的选中列位图、值流与 rawLen 合计（列目录常驻，不计）。
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
	lastPayload []byte
	lastPresent bool
	pos         int
}

// colSegs 是一列在一个装配分块里对每条命中的输出段（present=false = 缺失）。
type colSegs struct {
	buf     []byte
	off     []int32
	present []bool
}

// asmChunk 是单次装配的命中数上限，控制段缓冲与偏移数组的内存。
const asmChunk = 8192

// binLoad 是一个 bin 选中列的压缩数据：列定位 + 合并读段里的位图/值流视图。
type binLoad struct {
	bin       int
	meta      *colBinRecord
	parts     []colPart // 列号升序，装配按此顺序拼 payload
	bitFrames [][]byte
	frames    [][]byte
}

// loadBin 读一个 bin 的选中列：列目录 → 按偏移排序合并读段 → 顺序 pread。
func (c *colShard) loadBin(g hitBinGroup, sel []int) (*binLoad, error) {
	meta := &c.bins[g.bin]
	dir, err := c.loadDir(g.bin)
	if err != nil {
		return nil, err
	}
	parts := make([]colPart, 0, len(sel))
	for _, col := range sel {
		e := c.entry(dir, col)
		if e.bitmapOff == 0 {
			continue
		}
		p := colPart{
			col: col, flag: c.schema[col].Type == TypeFlag,
			present: int(e.present), bitmapAt: int64(e.bitmapOff),
			bitmapLen: int(e.bitmapLen),
		}
		if !p.flag {
			p.frameAt = int64(e.valueOff)
			p.frameLen = int(e.valueLen)
		}
		parts = append(parts, p)
	}
	ld := &binLoad{bin: g.bin, meta: meta, parts: parts}
	if len(parts) == 0 {
		return ld, nil
	}
	order := make([]int, len(parts))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool {
		return parts[order[i]].bitmapAt < parts[order[j]].bitmapAt
	})
	type readRange struct {
		start, end int64
		parts      []int
	}
	var ranges []readRange
	for _, pi := range order {
		p := &parts[pi]
		end := p.bitmapAt + int64(p.bitmapLen)
		if p.frameLen > 0 {
			end = p.frameAt + int64(p.frameLen)
		}
		if len(ranges) > 0 && p.bitmapAt <= ranges[len(ranges)-1].end {
			r := &ranges[len(ranges)-1]
			if end > r.end {
				r.end = end
			}
			r.parts = append(r.parts, pi)
			continue
		}
		ranges = append(ranges, readRange{start: p.bitmapAt, end: end, parts: []int{pi}})
	}
	ld.bitFrames = make([][]byte, len(parts))
	ld.frames = make([][]byte, len(parts))
	for ri := range ranges {
		r := &ranges[ri]
		buf := make([]byte, r.end-r.start)
		if _, err := c.f.ReadAt(buf, r.start); err != nil {
			return nil, fmt.Errorf("immkv: %s: 读 bin %d 列数据: %w", c.base, g.bin, err)
		}
		fadviseDontNeed(c.f, r.start, r.end-r.start)
		for _, pi := range r.parts {
			p := &parts[pi]
			off := p.bitmapAt - r.start
			ld.bitFrames[pi] = buf[off : off+int64(p.bitmapLen)]
			if p.frameLen > 0 {
				fo := p.frameAt - r.start
				ld.frames[pi] = buf[fo : fo+int64(p.frameLen)]
			}
		}
	}
	return ld, nil
}

// processBin 解压一个已读入的 bin 并按 schema 顺序拼出 record 回填 out。
func (c *colShard) processBin(ld *binLoad, g hitBinGroup, out [][]byte, threads int) error {
	meta, parts := ld.meta, ld.parts
	bitmapLen := recordBitmapLen(len(c.schema))
	if len(parts) == 0 {
		for _, h := range g.hits {
			out[h.dst] = make([]byte, bitmapLen)
		}
		return nil
	}
	bmLen := int((meta.count + 7) / 8)
	bits := make([][]byte, len(parts))
	decoded := make([][]byte, len(parts))
	workers := threads
	if workers > len(parts) {
		workers = len(parts)
	}
	if workers < 1 {
		workers = 1
	}
	var wg sync.WaitGroup
	errs := make(chan error, len(parts))
	jobs := make(chan int)
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := range jobs {
				b, err := c.dec.decode(ld.bitFrames[i], nil)
				if err != nil {
					errs <- fmt.Errorf("immkv: %s: 解压 bin %d 列 %d 位图: %w", c.base, g.bin, parts[i].col, err)
					continue
				}
				if len(b) != bmLen {
					errs <- fmt.Errorf("immkv: %s: bin %d 列 %d 位图长度 %d ≠ %d", c.base, g.bin, parts[i].col, len(b), bmLen)
					continue
				}
				if pc := popcountBytes(b); pc != parts[i].present {
					errs <- fmt.Errorf("immkv: %s: bin %d 列 %d 位图 %d 位 ≠ present_count %d",
						c.base, g.bin, parts[i].col, pc, parts[i].present)
					continue
				}
				bits[i] = b
				if parts[i].flag {
					continue
				}
				v, err := c.dec.decode(ld.frames[i], nil)
				if err != nil {
					errs <- fmt.Errorf("immkv: %s: 解压 bin %d 列 %d 值流: %w", c.base, g.bin, parts[i].col, err)
					continue
				}
				decoded[i] = v
			}
		}()
	}
	for i := range parts {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	close(errs)
	if err := <-errs; err != nil {
		return err
	}
	// 有列解码失败时不再装配（bits/decoded 可能不完整）。
	for i := range parts {
		if bits[i] == nil {
			return fmt.Errorf("immkv: %s: bin %d 列 %d 解码失败", c.base, g.bin, parts[i].col)
		}
		if !parts[i].flag && decoded[i] == nil {
			return fmt.Errorf("immkv: %s: bin %d 列 %d 值流解码失败", c.base, g.bin, parts[i].col)
		}
	}
	// 按列并行扫描：每列在自己的位图/值流上顺扫，产出「本分块每条命中的
	// payload 段」；最后按 schema 顺序合并成 record。
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
			field := &c.schema[p.col]
			cur := &cursors[pi]
			stream := decoded[pi]
			seg := colSegs{
				off:     make([]int32, len(hs)+1),
				present: make([]bool, len(hs)),
			}
			for hi, h := range hs {
				ridx := int(h.rec - int64(meta.first))
				if ridx < 0 || ridx >= int(meta.count) {
					return fmt.Errorf("immkv: %s: 记录号 %d 不在 bin %d 内", c.base, h.rec, g.bin)
				}
				if ridx != cur.lastRidx {
					for cur.lastBit < ridx {
						cur.lastBit++
						present := bitSet(bits[pi], cur.lastBit)
						var payload []byte
						if present && !p.flag {
							rest, err := skipValue(*field, stream[cur.pos:])
							if err != nil {
								return fmt.Errorf("immkv: %s: bin %d 列 %d 记录 %d: %w",
									c.base, g.bin, p.col, meta.first+uint64(cur.lastBit), err)
							}
							payload = stream[cur.pos : len(stream)-len(rest)]
							cur.pos += len(payload)
						}
						if cur.lastBit == ridx {
							cur.lastPayload, cur.lastPresent = payload, present
						}
					}
					cur.lastRidx = ridx
				}
				if cur.lastPresent {
					seg.present[hi] = true
					if !p.flag {
						seg.buf = append(seg.buf, cur.lastPayload...)
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
			total := bitmapLen
			for pi := range segs {
				if segs[pi].present[hi] {
					total += int(segs[pi].off[hi+1] - segs[pi].off[hi])
				}
			}
			rec := make([]byte, bitmapLen, total)
			for pi := range segs {
				if !segs[pi].present[hi] {
					continue
				}
				col := parts[pi].col
				rec[col>>3] |= 1 << (col & 7)
				rec = append(rec, segs[pi].buf[int(segs[pi].off[hi]):int(segs[pi].off[hi+1])]...)
			}
			out[h.dst] = rec
		}
	}
	return nil
}

// readColBatch 走列存路径回填 out；hits 按记录号升序。bin 级预取：读下一
// 个 bin 的压缩数据与解压/装配当前 bin 重叠，HDD 上不让磁头等 CPU。
func (c *colShard) readColBatch(hits []sHit, sel []int, out [][]byte, threads int) error {
	groups, err := c.groupHits(hits)
	if err != nil {
		return err
	}
	if len(groups) == 0 {
		return nil
	}
	type loadResult struct {
		ld  *binLoad
		err error
	}
	load := func(g hitBinGroup) <-chan loadResult {
		ch := make(chan loadResult, 1)
		go func() {
			ld, err := c.loadBin(g, sel)
			ch <- loadResult{ld: ld, err: err}
		}()
		return ch
	}
	ch := load(groups[0])
	for i, g := range groups {
		res := <-ch
		if res.err != nil {
			return res.err
		}
		if i+1 < len(groups) {
			ch = load(groups[i+1])
		}
		if err := c.processBin(res.ld, g, out, threads); err != nil {
			return err
		}
	}
	return nil
}

func bitSet(b []byte, i int) bool {
	return b[i>>3]&(1<<(i&7)) != 0
}

// sameColSchema 校验两个染色体的列 schema 完全一致（同一构建应如此）。
func sameColSchema(a, b Schema) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
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
