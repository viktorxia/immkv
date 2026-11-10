// stream.go 是 immkv 的读侧之二：OpenStreaming 批量流式句柄。
//
// GetBatch 对一段升序键做两阶段查询：阶段一对每条染色体的键流做合并连接
// （顺扫键流窗口，光标只前进），收齐全部命中的值帧 (偏移, 长度, 裸存标记)；
// 阶段二把相邻命中合并成读段，Threads 路并行 pread + 解压回填结果。
// 磁盘读取因此是字节精确的：只读命中的值帧与扫过的键流，无缺页放大。

package immkv

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
)

const (
	streamWindow = 256 << 10 // 键流窗口目标字节数（一次 pread 覆盖若干组）
	mergeGap     = 4 << 10   // 值帧相邻合并的间隔阈值：间隔 ≤ 它并入同一读段；gnomAD 命中值平均间隔 ~114KB，阈值大了会顺读值文件（实测 256KB 时每轮多读 200GB）
	maxExtent    = 4 << 20   // 单个读段的上限，控制缓冲与单次 pread 尺寸
)

// ErrUnsorted 表示 GetBatch 的输入键不升序。
var ErrUnsorted = errors.New("immkv: GetBatch keys not ascending")

// StreamDB 是批量流式句柄：GetBatch / Close。方法面刻意只有这两个——
// 单条查询开 OpenEager；同一库可同时开两种句柄（只读无冲突）。
type StreamDB struct {
	shardSet
	threads int
	cols    [25]*colShard
	colSel  []int // 选中的列号（升序）；nil = 全部
	profile ioProfile

	forceCol bool // 测试用：跳过成本估算，强制列存路径
}

// OpenStreaming 打开 dir 下的批量流式句柄。Options.Threads 是唯一并发旋钮：
// 同时决定在途 pread 请求数与并行解压 goroutine 数（默认 1）；机械硬盘上
// 它不起作用（单磁头，引擎靠相邻合并转顺序读）。
//
// 库里有 base.col/manifest 时，顺序查按估算的 I/O 成本在记录序与列存之间
// 选路；Options.Columns 只选字段（nil/空 = 全部列），不决定路径。没有
// base.col 的旧库仍按记录序服务；此时给非空 Columns 会报错。
func OpenStreaming(dir string, opts Options) (*StreamDB, error) {
	if opts.Threads <= 0 {
		opts.Threads = 1
	}
	ss, err := openShardSet(dir, 0)
	if err != nil {
		return nil, err
	}
	d := &StreamDB{shardSet: *ss, threads: opts.Threads, profile: detectIOProfile(dir)}
	anyCol, allCol := false, true
	var buildID string
	for _, id := range d.chroms {
		cs, err := openColShard(filepath.Join(dir, "chr"+ChromName(id)), d.shards[id])
		if err != nil {
			d.Close()
			return nil, err
		}
		d.cols[id] = cs
		if cs == nil {
			allCol = false
			continue
		}
		anyCol = true
		if buildID == "" {
			buildID = cs.buildID
		} else if buildID != cs.buildID {
			d.Close()
			return nil, errors.New("immkv: 各染色体的 base.col build id 不一致")
		}
		if len(d.chroms) > 0 && d.cols[d.chroms[0]] != nil &&
			!sameColSchema(d.cols[d.chroms[0]].fields, cs.fields) {
			d.Close()
			return nil, errors.New("immkv: 各染色体的 base.col schema 不一致")
		}
	}
	if anyCol && !allCol {
		d.Close()
		return nil, errors.New("immkv: 部分染色体缺少 base.col/manifest")
	}
	if len(opts.Columns) > 0 {
		if !anyCol {
			d.Close()
			return nil, errors.New("immkv: 库没有 base.col/manifest，不能按列选择")
		}
		seen := make(map[int]bool, len(opts.Columns))
		sel := make([]int, 0, len(opts.Columns))
		for _, c := range opts.Columns {
			if c < 0 || c >= len(d.cols[d.chroms[0]].fields) {
				d.Close()
				return nil, fmt.Errorf("immkv: Columns 列号 %d 越界", c)
			}
			if seen[c] {
				d.Close()
				return nil, fmt.Errorf("immkv: Columns 列号 %d 重复", c)
			}
			seen[c] = true
			sel = append(sel, c)
		}
		sort.Ints(sel)
		d.colSel = sel
	}
	return d, nil
}

// Close 关闭全部 shard。调用方保证不与查询并发。
func (d *StreamDB) Close() error {
	for id := uint8(0); id < 25; id++ {
		if c := d.cols[id]; c != nil {
			c.Close()
			d.cols[id] = nil
		}
	}
	return d.close()
}

// directOK 报全部在场染色体的值文件都开出了 O_DIRECT 句柄（深队列读路径
// 的前提；任一失败就整体退回缓冲读，避免同一批里混两种路径）。
func (d *StreamDB) directOK() bool {
	for _, id := range d.chroms {
		if d.shards[id].vfDirect == nil {
			return false
		}
	}
	return len(d.chroms) > 0
}

// GetBatch 批量点查。keys 必须升序（字节序；VCF 按染色体自然序 + 位置排序
// 即天然满足），返回切片与 keys 等长同序，某项未命中为 nil。返回切片只保证
// 到下一次查询前有效，要保留自己拷。任何 I/O 或解压错误整批返回 error。
// 并发安全：句柄可被多个 goroutine 同时调用（每次调用状态独立）。
func (d *StreamDB) GetBatch(keys [][]byte) ([][]byte, error) {
	if len(keys) == 0 {
		return [][]byte{}, nil
	}
	for i := 1; i < len(keys); i++ {
		if bytes.Compare(keys[i-1], keys[i]) > 0 {
			return nil, ErrUnsorted
		}
	}
	out := make([][]byte, len(keys))
	// 三级流水线：生产者（本 goroutine，合并连接顺扫键流）→ 读池（Threads 路
	// 在途读）→ 解压池（Threads 路，CPU 由 GOMAXPROCS 兜底）。读与解压分池
	// 解耦，磁盘等待不占解压的 CPU 时间。
	var (
		extCh  = make(chan []extent, 4) // 生产者 → 读池
		decCh  = make(chan []extent, 8) // 读池 → 解压池
		abort  = make(chan struct{})
		wgRead sync.WaitGroup
		wgDec  sync.WaitGroup
		errMu  sync.Mutex
		ferr   error
	)
	fail := func(err error) {
		errMu.Lock()
		defer errMu.Unlock()
		if ferr == nil {
			ferr = err
			close(abort)
		}
	}
	failed := func() bool {
		errMu.Lock()
		defer errMu.Unlock()
		return ferr != nil
	}
	wgDec.Add(d.threads)
	if r, rerr := newIOUring(512); rerr == nil && d.directOK() {
		// O_DIRECT + io_uring：单驱动 goroutine 深队列异步读（绕开页缓存，
		// 在途读深度与线程数解耦；线程预算全部留给解压池）
		defer r.close()
		wgRead.Add(1)
		go func() {
			defer wgRead.Done()
			driveDirect(r, extCh, decCh, abort, fail)
		}()
	} else {
		wgRead.Add(d.threads)
		for w := 0; w < d.threads; w++ {
			go func() { // 读池（io_uring 或 O_DIRECT 不可用时的兜底）
				defer wgRead.Done()
				for {
					select {
					case ch, ok := <-extCh:
						if !ok {
							return
						}
						if err := readChunk(ch); err != nil {
							fail(err)
							return
						}
						select {
						case decCh <- ch:
						case <-abort:
							return
						}
					case <-abort:
						return
					}
				}
			}()
		}
	}
	for w := 0; w < d.threads; w++ {
		go func() { // 解压池
			defer wgDec.Done()
			for {
				select {
				case ch, ok := <-decCh:
					if !ok {
						return
					}
					if err := decodeChunk(ch, out); err != nil {
						fail(err)
						return
					}
				case <-abort:
					return
				}
			}
		}()
	}
	// 生产者：按染色体分段（键升序 ⇒ 同染色体键连续）合并连接、切读段投递
	i := 0
	for i < len(keys) && !failed() {
		if len(keys[i]) == 0 { // 空键排最前，直接未命中
			i++
			continue
		}
		id := keys[i][0]
		j := i + 1
		for j < len(keys) && len(keys[j]) > 0 && keys[j][0] == id {
			j++
		}
		if id <= 24 && d.present[id] {
			hits, err := mergeJoin(d.shards[id], keys[i:j], i)
			if err != nil {
				fail(err)
				break
			}
			if cs := d.cols[id]; cs != nil {
				useCol, cerr := d.preferColumnar(cs, hits)
				if cerr != nil {
					fail(cerr)
					break
				}
				if useCol {
					if cerr := cs.readColBatch(hits, selectedColumns(cs.fields, d.colSel), out, d.threads); cerr != nil {
						fail(cerr)
						break
					}
					i = j
					continue
				}
			}
			for _, ch := range chunkExtents(buildExtents(d.shards[id], hits)) {
				select {
				case extCh <- ch:
				case <-abort:
				}
				if failed() {
					break
				}
			}
		}
		i = j
	}
	// 统一收尾：关 extCh → 等读池 → 关 decCh → 等解压池，全部落定后判错——
	// 解压池的错误可能在生产者跑完之后才出现，提前判错会把它丢掉
	close(extCh)
	wgRead.Wait()
	close(decCh)
	wgDec.Wait()
	if ferr != nil {
		return nil, ferr
	}
	return out, nil
}

// costs 返回读路的粗略成本参数：单次随机读耗时（秒）与顺序带宽（B/s）。
// 数值取自 2026-11-09/10 的 HDD/NVMe 实测量级，只用于二选一，不参与计时。
func (p ioProfile) costs() (randSec, seqBps float64) {
	if p.rotational {
		return 0.45e-3, 170e6
	}
	return 10e-6, 3e9
}

// preferColumnar 估算记录序（命中帧随机读）与列存（命中 bin 的选中列顺序
// 读）谁更省，返回是否走列存。hits 按记录号升序。
func (d *StreamDB) preferColumnar(cs *colShard, hits []sHit) (bool, error) {
	if len(hits) == 0 {
		return false, nil
	}
	// 实验开关：显式对比两条路，不改变默认选路。
	if os.Getenv("IMMKV_FORCE_RECORD") != "" {
		return false, nil
	}
	if os.Getenv("IMMKV_FORCE_COLUMNAR") != "" {
		return true, nil
	}
	if d.forceCol {
		return true, nil
	}
	sel := selectedColumns(cs.fields, d.colSel)
	colBytes, colRaw, err := cs.estimate(hits, sel)
	if err != nil {
		return false, err
	}
	var frames int
	var recBytes uint64
	var lastOff uint64
	var lastLen uint32
	have := false
	for _, h := range hits {
		if h.frameLen == 0 {
			continue
		}
		if have && h.valOff == lastOff && h.frameLen == lastLen {
			continue
		}
		have, lastOff, lastLen = true, h.valOff, h.frameLen
		frames++
		recBytes += uint64(h.frameLen)
	}
	randSec, seqBps := d.profile.costs()
	workers := float64(d.threads)
	if workers < 1 {
		workers = 1
	}
	const decodePerThread = 1.2e9 // zstd 解压吞吐量级（B/s/线程）
	decodeBps := decodePerThread * workers
	// 记录序只需解命中帧；原文大小按文本列存压缩比 ~8 估。
	recRaw := float64(recBytes) * 8
	recSec := float64(frames)*randSec + float64(recBytes)/seqBps + recRaw/decodeBps
	colSec := float64(colBytes)/seqBps + float64(colRaw)/decodeBps
	return colSec < recSec, nil
}

// emptyValue 是零长值的共享只读空切片（长度 0 无从写入，共享安全）。
var emptyValue = []byte{}

// sHit 是一条命中：值帧在 vals 文件里的位置 + 该键在整批里的回填下标；
// rec 是该记录在本染色体内的全局记录号（打包帧换算组内下标用）。
type sHit struct {
	valOff   uint64
	frameLen uint32
	raw      bool
	packed   bool
	rec      int64
	dst      int
}

// mergeJoin 对一段升序键做合并连接。键流按微块（64 条/块）顺扫：窗口一次
// 读入若干块（凑满 streamWindow 或到键流尾），块解析为零分配视图（数组是
// 窗口缓冲的子切片）；光标 (块号, 块内下标, 该下标的 valOff) 跨查询保持、
// 只前进；块内定位用二分（log₂64 = 6 次比较），命中时从光标累加推导 valOff
// （≤64 步）。返回的命中按值偏移升序（记录按 key 序、帧按记录序落盘）。
func mergeJoin(s *shard, keys [][]byte, base int) ([]sHit, error) {
	if s.nBlock == 0 || len(keys) == 0 {
		return nil, nil
	}
	var hits []sHit
	b := s.blockFor(keys[0])
	j := 0            // 块内下一待看下标
	jOff := uint64(0) // 第 j 条记录的 valOff（块内推导的运行值）
	var (
		blocks []blockView
		winB0  = -1
		winB1  = 0
	)
	loadWin := func(from int) error {
		b1 := from + 1
		for b1 < int(s.nBlock) && s.blkOff[b1]-s.blkOff[from] < streamWindow {
			b1++
		}
		hi := s.keysEnd
		if b1 < int(s.nBlock) {
			hi = s.blkOff[b1]
		}
		buf := make([]byte, hi-s.blkOff[from])
		if _, err := s.kf.ReadAt(buf, int64(s.blkOff[from])); err != nil {
			return err
		}
		blocks = blocks[:0]
		p := 0
		for bb := from; bb < b1; bb++ {
			v, used, err := parseBlockView(buf[p:], s.nRecs(bb))
			if err != nil {
				return err
			}
			blocks = append(blocks, v)
			p += used
		}
		winB0, winB1 = from, b1
		return nil
	}
	for ki := 0; ki < len(keys); ki++ {
		q := keys[ki]
		for {
			if b >= int(s.nBlock) {
				break // 数据扫完，剩余键全部未命中
			}
			if b >= winB1 {
				if err := loadWin(b); err != nil {
					return nil, err
				}
				j, jOff = 0, blocks[0].base
			}
			v := &blocks[b-winB0]
			if j == 0 {
				jOff = v.base
			}
			if j >= v.n {
				b++
				j = 0
				continue
			}
			jj := j + sort.Search(v.n-j, func(t int) bool {
				return bytes.Compare(v.keyAt(j+t), q) >= 0
			})
			if jj >= v.n {
				b++
				j = 0
				continue
			}
			c := bytes.Compare(v.keyAt(jj), q)
			if c > 0 {
				// 记录键 > 查询键：未命中，光标停在当前记录（先推偏移再挪下标）
				jOff = v.valOffAt(j, jj, jOff)
				j = jj
				break
			}
			// 命中：valOff 从光标累加推导
			vo := v.valOffAt(j, jj, jOff)
			l := v.lens[jj]
			hits = append(hits, sHit{
				valOff: vo, frameLen: l & frameLenMask,
				raw: l&flagRaw != 0, packed: l&flagPacked != 0,
				rec: int64(b)*int64(s.recsBlk) + int64(jj), dst: base + ki,
			})
			// 下一个输入键与此键相等时光标停在当前记录上，让同一记录
			// 再服务一次（批量点查对每个输入独立）
			nextOff := v.valOffAt(j, jj+1, jOff)
			j = jj + 1
			jOff = nextOff
			if ki+1 < len(keys) && bytes.Equal(keys[ki+1], q) {
				j = jj
				jOff = vo
			}
			break
		}
	}
	return hits, nil
}

// extent 是一次 pread 的读段：从 off 起读到段内最后一条命中的帧尾。
type extent struct {
	s      *shard
	off    uint64
	hits   []sHit // 段内命中，值偏移升序
	buf    []byte // 命中帧字节视图（O_DIRECT 路径是 dbuf 的子切片）
	dbuf   []byte // 读入缓冲本体（对齐或普通），解压完归还池；raw/单值段被 out 引用不还
	size   uint64 // 段的盘上字节数（readChunk / driveDirect 填）
	dNeed  int32  // dbuf 内需要覆盖到的字节数（O_DIRECT 对齐读可能越过文件尾，EOF 短读合法）
	direct bool   // dbuf 走对齐缓冲池（driveDirect 填；归还各回各池）
}

// buildExtents 把命中按相邻合并成读段：间隔 ≤ mergeGap 且段长 ≤ maxExtent
// 时并入，否则另起一段。稀疏位点在值文件里的平均间隔 ~160KB（gnomAD 全量库
// 实测口径），合并后每次 syscall 取回数条帧。
func buildExtents(s *shard, hits []sHit) []extent {
	var exts []extent
	for i := 0; i < len(hits); {
		j := i
		last := hits[i].valOff + uint64(hits[i].frameLen)
		for j+1 < len(hits) {
			h := hits[j+1]
			end := h.valOff + uint64(h.frameLen)
			// 同帧成员（含零长裸存帧）的 valOff 不大于 last，间隔按 0 计；
			// 直接相减会无符号下溢成巨值，把段无谓打碎
			var gap int64
			if h.valOff > last {
				gap = int64(h.valOff - last)
			}
			if gap > mergeGap || end-hits[i].valOff > maxExtent {
				break
			}
			last = end
			j++
		}
		exts = append(exts, extent{s: s, off: hits[i].valOff, hits: hits[i : j+1]})
		i = j + 1
	}
	return exts
}

// readChunk 把一块读段全部读入（每 reader 一路在途读，threads 个 reader 即
// Options.Threads 路并发读），读完对该块覆盖的文件区间做 fadvise(DONTNEED)
// ——值只读一次，页缓存留着只会顶爆 cgroup 限额、拖慢后续读。
func readChunk(chunk []extent) error {
	for i := range chunk {
		e := &chunk[i]
		last := e.hits[len(e.hits)-1]
		e.size = last.valOff + uint64(last.frameLen) - e.off
		e.buf = getExtentBuf(int(e.size))
		e.dbuf = e.buf
		if _, err := e.s.vf.ReadAt(e.buf, int64(e.off)); err != nil {
			return err
		}
	}
	first, lastE := chunk[0], chunk[len(chunk)-1]
	fadviseDontNeed(first.s.vf, int64(first.off), int64(lastE.off+lastE.size-first.off))
	return nil
}

// decodeChunk 解一块读段并回填。打包帧解出的成员值拷进本块 slab——结果自
// 包含、段读缓冲可即时回收复用，全量注释一轮的分配量从「每命中一整帧」降到
// 「每命中一个值」。
func decodeChunk(chunk []extent, out [][]byte) error {
	slab := make([]byte, 0, 1<<16)
	for i := range chunk {
		var err error
		if slab, err = decodeExtent(&chunk[i], out, slab); err != nil {
			return err
		}
	}
	return nil
}

// decodeExtent 解一个读段并回填：raw 帧零拷（out 引用 buf，本块内该段缓冲
// 不回收）；单值 zstd 帧整帧解压成新缓冲；打包帧整帧解压后把成员值拷进
// slab，段缓冲随后回收。返回推进后的 slab。
func decodeExtent(e *extent, out [][]byte, slab []byte) ([]byte, error) {
	reusable := true
	defer func() {
		if reusable {
			if e.direct {
				putDirectBuf(e.dbuf)
			} else {
				putExtentBuf(e.dbuf)
			}
			e.buf, e.dbuf = nil, nil
		}
	}()
	for i := 0; i < len(e.hits); {
		j := i
		h := e.hits[i]
		// 只把同一打包帧的成员并成组（valOff、frameLen、packed 三者全同）；
		// 零长裸存帧与共址的下一帧不能并——它不是那帧的成员
		for j+1 < len(e.hits) &&
			e.hits[j+1].valOff == h.valOff &&
			e.hits[j+1].frameLen == h.frameLen &&
			h.packed && e.hits[j+1].packed {
			j++
		}
		frame := e.buf[h.valOff-e.off : h.valOff-e.off+uint64(h.frameLen)]
		switch {
		case h.raw:
			if h.frameLen == 0 {
				out[h.dst] = emptyValue // 零长值：非 nil 空切片（nil 会被当成未命中）
				i = j + 1
				continue
			}
			out[h.dst] = frame
			reusable = false // out 引用 buf，本调用结束前不能复用
		case h.packed:
			p, err := decodePack(e.s.dec, frame, getExtentBuf(64<<10))
			if err != nil {
				return slab, err
			}
			for k := i; k <= j; k++ {
				hk := e.hits[k]
				idx := hk.rec - int64(p.first)
				if idx < 0 || idx >= int64(p.n) {
					return slab, fmt.Errorf("immkv: 打包帧记录号 %d 不在 [%d, %d) 内", hk.rec, p.first, p.first+p.n)
				}
				v := p.value(uint32(idx))
				t := len(slab)
				slab = append(slab, v...)
				out[hk.dst] = slab[t:]
			}
			putExtentBuf(p.buf) // 成员值已拷进 slab，解码缓冲即可复用
		default:
			v, err := e.s.dec.decode(frame, nil)
			if err != nil {
				return slab, err
			}
			out[h.dst] = v
			reusable = false // 解出的整帧缓冲由 out 引用，交 GC
		}
		i = j + 1
	}
	return slab, nil
}

// extentBufPool 复用段读缓冲。按尺寸分桶避免大段把池撑爆：≤64KB（gnomAD
// 段主流尺寸）走池，更大的现配。
var extentBufPools = [2]sync.Pool{}

func getExtentBuf(n int) []byte {
	if n <= 16<<10 {
		if bp := extentBufPools[0].Get(); bp != nil {
			return (*bp.(*[]byte))[:n]
		}
		b := make([]byte, 16<<10)
		return b[:n]
	}
	if n <= 64<<10 {
		if bp := extentBufPools[1].Get(); bp != nil {
			return (*bp.(*[]byte))[:n]
		}
		b := make([]byte, 64<<10)
		return b[:n]
	}
	return make([]byte, n)
}

func putExtentBuf(b []byte) {
	switch {
	case cap(b) <= 16<<10:
		extentBufPools[0].Put(&b)
	case cap(b) <= 64<<10:
		extentBufPools[1].Put(&b)
	}
}

// chunkExtents 把一条染色体的读段切成投递块：≤64 段一块。
func chunkExtents(exts []extent) [][]extent {
	if len(exts) == 0 {
		return nil
	}
	const n = 64
	var chunks [][]extent
	for i := 0; i < len(exts); i += n {
		j := i + n
		if j > len(exts) {
			j = len(exts)
		}
		chunks = append(chunks, exts[i:j])
	}
	return chunks
}
