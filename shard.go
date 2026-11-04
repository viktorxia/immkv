// shard.go 是 immkv 的读侧之一：分染色体分片、OpenEager 句柄、解压缓存。
//
// 每条染色体一个 shard：两个只读文件句柄（.keys/.vals），锚点索引打开时一次
// pread 全量读入常驻；组（1024 条记录）与值帧按需 pread，OpenEager 配 LRU
// 缓存组窗口与解压值。键首字节（chromID）路由到对应 shard，Scan 跨染色体按
// chromID 序拼接。

package immkv

import (
	"bytes"
	"container/list"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"unsafe"

	"github.com/klauspost/compress/zstd"
)

// ---- 解压缓存：分片 LRU（结构移植自 immdb，缓存对象从 []byte 泛化为 any） ----

const lruShards = 16

type lruEnt struct {
	key  int64
	blob any
	size int
}

type lruShard struct {
	mu   sync.Mutex
	cap  int
	used int
	m    map[int64]*list.Element
	l    *list.List
	_pad [64]byte // 隔离相邻分片 mutex 的缓存行（假共享）
}

type lruCache struct {
	shards [lruShards]lruShard
}

func newLRU(capBytes int) *lruCache {
	c := &lruCache{}
	per := capBytes / lruShards
	for i := range c.shards {
		c.shards[i].cap = per
		c.shards[i].m = make(map[int64]*list.Element)
		c.shards[i].l = list.New()
	}
	return c
}

func (c *lruCache) shard(k int64) *lruShard {
	// splitmix64 终结混合：键本身低位带结构（文件偏移），直接取模会全部
	// 挤进同一分片
	h := uint64(k)
	h ^= h >> 33
	h *= 0xff51afd7ed558ccd
	h ^= h >> 33
	return &c.shards[h%lruShards]
}

func (c *lruCache) get(k int64) (any, bool) {
	s := c.shard(k)
	s.mu.Lock()
	defer s.mu.Unlock()
	if el, ok := s.m[k]; ok {
		s.l.MoveToFront(el)
		return el.Value.(*lruEnt).blob, true
	}
	return nil, false
}

func (c *lruCache) add(k int64, v any, size int) {
	s := c.shard(k)
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.m[k]; ok {
		return
	}
	s.m[k] = s.l.PushFront(&lruEnt{k, v, size})
	s.used += size
	for s.used > s.cap && s.l.Len() > 1 {
		el := s.l.Back()
		ent := el.Value.(*lruEnt)
		s.l.Remove(el)
		delete(s.m, ent.key)
		s.used -= ent.size
	}
}

// frameDecoder 是值帧解压后端的窄接口：decode 把 src 的解压结果追加进 dst
// 返回（dst 容量够时零分配）；Close 释放后端资源。两个实现：libzstd CGO
// （zstdcgo_cgo.go，DDict 常驻、并发安全）与 klauspost 纯 Go（kpDecoder）。
// 帧格式标准互通，后端按 CGO 可用性与建器成败择一。
type frameDecoder interface {
	decode(src, dst []byte) ([]byte, error)
	Close() error
}

// kpDecoder 把 klauspost 解码器适配到 frameDecoder。
type kpDecoder struct {
	dec *zstd.Decoder
}

func (k *kpDecoder) decode(src, dst []byte) ([]byte, error) {
	return k.dec.DecodeAll(src, dst)
}

func (k *kpDecoder) Close() error {
	k.dec.Close()
	return nil
}

// ---- 单染色体 shard ----

type shard struct {
	kf, vf   *os.File
	vfDirect *os.File // O_DIRECT 值文件句柄；不可用为 nil（流式深队列读用）
	count    uint64
	nBlock   uint32
	recsBlk  uint64
	indexOff uint64
	blkPfx   []uint64 // 块首键 8B 前缀（大端装载，零填充），比较语义与全键一致
	blkOff   []uint64 // 每块在 keys 文件里的偏移
	keysEnd  uint64   // 键流数据区终点（= indexOff）
	dec      frameDecoder
	cache    *lruCache // OpenEager 用；nil = 不缓存（OpenStreaming 不用它）
}

func openShard(base string, cache *lruCache) (*shard, error) {
	kf, err := os.Open(base + ".keys")
	if err != nil {
		return nil, err
	}
	vf, err := os.Open(base + ".vals")
	if err != nil {
		kf.Close()
		return nil, err
	}
	s := &shard{kf: kf, vf: vf, cache: cache}
	// O_DIRECT 可选句柄：不支持（EINVAL 等）只是退回缓冲读，不算错
	if df, err := openDirect(base + ".vals"); err == nil {
		s.vfDirect = df
	}
	fail := func(err error) (*shard, error) {
		kf.Close()
		vf.Close()
		return nil, err
	}
	var kh [headerSize]byte
	if _, err := kf.ReadAt(kh[:], 0); err != nil {
		return fail(err)
	}
	if string(kh[0:8]) != keysMagic {
		return fail(errors.New("immkv: bad keys magic: " + base))
	}
	if v := binary.LittleEndian.Uint32(kh[8:12]); v != 1 {
		return fail(fmt.Errorf("immkv: keys version %d ≠ 1", v))
	}
	if v := binary.LittleEndian.Uint16(kh[20:22]); uint64(v) != recsPerBlock {
		return fail(fmt.Errorf("immkv: recsPerBlock %d ≠ %d", v, recsPerBlock))
	}
	s.count = binary.LittleEndian.Uint64(kh[12:20])
	s.recsBlk = recsPerBlock
	s.nBlock = binary.LittleEndian.Uint32(kh[24:28])
	s.indexOff = binary.LittleEndian.Uint64(kh[28:36])
	s.keysEnd = s.indexOff
	var vh [headerSize]byte
	if _, err := vf.ReadAt(vh[:], 0); err != nil {
		return fail(err)
	}
	if string(vh[0:8]) != valsMagic {
		return fail(errors.New("immkv: bad vals magic: " + base))
	}
	if v := binary.LittleEndian.Uint32(vh[8:12]); v != 1 {
		return fail(fmt.Errorf("immkv: vals version %d ≠ 1", v))
	}
	if v := binary.LittleEndian.Uint64(vh[12:20]); v != s.count {
		return fail(fmt.Errorf("immkv: vals count %d ≠ keys count %d", v, s.count))
	}
	// 块索引一次读入常驻：[prefix 8B ×N][blockOff 8B ×N]
	st, err := kf.Stat()
	if err != nil {
		return fail(err)
	}
	ib := make([]byte, int64(st.Size())-int64(s.indexOff))
	if _, err := kf.ReadAt(ib, int64(s.indexOff)); err != nil {
		return fail(err)
	}
	if uint64(len(ib)) != 16*uint64(s.nBlock) {
		return fail(fmt.Errorf("immkv: 块索引区 %d B ≠ 16×%d", len(ib), s.nBlock))
	}
	s.blkPfx = make([]uint64, s.nBlock)
	s.blkOff = make([]uint64, s.nBlock)
	for i := uint32(0); i < s.nBlock; i++ {
		s.blkPfx[i] = binary.BigEndian.Uint64(ib[8*i:])
		s.blkOff[i] = binary.LittleEndian.Uint64(ib[8*uint64(s.nBlock)+8*uint64(i):])
	}
	// 每值成帧构建的库在 vals 头里嵌训练字典；有字典的库 CGO 可用时解压走
	// libzstd DDict。无字典的 v2 打包库保留 klauspost：实测 gozstd 解 64KB
	// 打包帧即使 dst 给足也不快过 klauspost（4096/1：977ms 对 867ms；65536/4：
	// 586ms 对 562ms，ClinVar 键集 A/B 两轮），还省一个 96KB 缓冲桶；点查路径
	// 两者无差（p50 同 300ns）。
	var dict []byte
	decOpts := []zstd.DOption{zstd.IgnoreChecksum(true)} // 帧完整性由文件保证，跳过 xxhash 校验
	if dictLen := binary.LittleEndian.Uint32(vh[20:24]); dictLen > 0 {
		dictOff := binary.LittleEndian.Uint64(vh[24:32])
		dict = make([]byte, dictLen)
		if _, err := vf.ReadAt(dict, int64(dictOff)); err != nil {
			return fail(err)
		}
		decOpts = append(decOpts, zstd.WithDecoderDicts(dict))
	}
	var dec frameDecoder
	if len(dict) > 0 {
		dec, _ = newCGODictDecoder(dict)
	}
	if dec == nil {
		kd, derr := zstd.NewReader(nil, decOpts...)
		if derr != nil {
			return fail(derr)
		}
		dec = &kpDecoder{dec: kd}
	}
	s.dec = dec
	return s, nil
}

// Close 关闭文件句柄与解码器。shard 不可再用。
func (s *shard) Close() error {
	s.dec.Close()
	err := s.kf.Close()
	if ce := s.vf.Close(); err == nil {
		err = ce
	}
	if s.vfDirect != nil {
		if ce := s.vfDirect.Close(); err == nil {
			err = ce
		}
	}
	return err
}

// blockEnt 是一个微块的解析结果：视图指向窗口/缓存持有的 buf；cum[i] 是
// 第 i 条记录相对 baseValOff 的偏移（share 位不计增量），valOff_i = base+cum[i]。
type blockEnt struct {
	n       int
	keyEnds []uint32
	lens    []uint32
	base    uint64
	keys    []byte
	cum     []uint64
	buf     []byte
}

// keyAt 返回块内第 j 条键（视图）。
func (e *blockEnt) keyAt(j int) []byte {
	lo := uint32(0)
	if j > 0 {
		lo = e.keyEnds[j-1]
	}
	return e.keys[lo:e.keyEnds[j]]
}

// valLoc 返回块内第 j 条的值帧定位。
func (e *blockEnt) valLoc(j int) (valOff uint64, frameLen uint32, raw, packed bool) {
	l := e.lens[j]
	return e.base + e.cum[j], l & frameLenMask, l&flagRaw != 0, l&flagPacked != 0
}

// parseBlockView 从 buf 起点解析微块的零分配视图：keyEnds/lens/keys 都是 buf
// 的子切片，不含 cum（随机访问的偏移前缀和由块持有方另行构建）。流式顺扫用
// 这个——全量一轮 410 万块，逐块分配会把 GC 打爆。
type blockView struct {
	n       int
	base    uint64
	keyEnds []uint32
	lens    []uint32
	keys    []byte
}

func (v *blockView) keyAt(j int) []byte {
	lo := uint32(0)
	if j > 0 {
		lo = v.keyEnds[j-1]
	}
	return v.keys[lo:v.keyEnds[j]]
}

// valOffAt 从块内第 from 条（已知其 valOff）推导第 to 条的 valOff：
// to 条带 share 位则偏移不动，否则累加前一条帧长（≤64 步）。
func (v *blockView) valOffAt(from, to int, fromOff uint64) uint64 {
	off := fromOff
	for i := from; i < to && i+1 < v.n; i++ {
		if v.lens[i+1]&flagShare == 0 {
			off += uint64(v.lens[i] & frameLenMask)
		}
	}
	return off
}

func parseBlockView(buf []byte, n int) (blockView, int, error) {
	if n < 1 {
		return blockView{}, 0, errors.New("immkv: 空微块")
	}
	if len(buf) < 8*n+8 {
		return blockView{}, 0, errors.New("immkv: 微块头不足")
	}
	v := blockView{n: n}
	v.keyEnds = unsafe.Slice((*uint32)(unsafe.Add(unsafe.Pointer(&buf[0]), 0)), n)
	lensAt := 4 * n
	v.lens = unsafe.Slice((*uint32)(unsafe.Add(unsafe.Pointer(&buf[0]), lensAt)), n)
	v.base = binary.LittleEndian.Uint64(buf[8*n:])
	keyLen := int(v.keyEnds[n-1])
	if len(buf) < 8*n+8+keyLen {
		return blockView{}, 0, errors.New("immkv: 微块键区不足")
	}
	v.keys = buf[8*n+8 : 8*n+8+keyLen]
	return v, 8*n + 8 + keyLen, nil
}

// parseBlock 从 buf 起点解析一个 n 条记录的微块（含 cum），返回块与占用字节数。
// eager 点查用（块进 LRU，解析成本一次性）。
func parseBlock(buf []byte, n int) (*blockEnt, int, error) {
	v, used, err := parseBlockView(buf, n)
	if err != nil {
		return nil, 0, err
	}
	e := &blockEnt{n: v.n, base: v.base, keyEnds: v.keyEnds, lens: v.lens, keys: v.keys, buf: buf}
	// cum[i] = 第 i 条相对 base 的偏移：第 i 条带 share 位（与上一条同帧）
	// 则不加第 i−1 条的帧长，否则加（零长帧自然推进 0）
	e.cum = make([]uint64, n+1)
	for i := 1; i <= n; i++ {
		e.cum[i] = e.cum[i-1]
		if i < n && e.lens[i]&flagShare == 0 {
			e.cum[i] += uint64(e.lens[i-1] & frameLenMask)
		}
	}
	return e, used, nil
}

// nRecs 返回第 b 块的记录数（末块可短）。
func (s *shard) nRecs(b int) int {
	n := int(s.count - uint64(b)*s.recsBlk)
	if n > int(s.recsBlk) {
		n = int(s.recsBlk)
	}
	return n
}

// block 取第 b 块（LRU 缓存，键取块偏移正值域）。
func (s *shard) block(b int) *blockEnt {
	ck := int64(s.blkOff[b])
	if s.cache != nil {
		if v, ok := s.cache.get(ck); ok {
			return v.(*blockEnt)
		}
	}
	hi := s.keysEnd
	if b+1 < int(s.nBlock) {
		hi = s.blkOff[b+1]
	}
	buf := make([]byte, hi-s.blkOff[b])
	if _, err := s.kf.ReadAt(buf, int64(s.blkOff[b])); err != nil {
		return nil
	}
	e, _, err := parseBlock(buf, s.nRecs(b))
	if err != nil {
		return nil
	}
	if s.cache != nil {
		s.cache.add(ck, e, len(buf)+len(e.cum)*8+64)
	}
	return e
}

// key8 把键前 8 字节按大端装进 u64（不足补零），供块索引二分。
func key8(key []byte) uint64 {
	var b [8]byte
	copy(b[:], key)
	return binary.BigEndian.Uint64(b[:])
}

// blockFor 返回首个「块内可能含 ≥ key 记录」的块号：最后一个首键前缀 < key
// 前缀的块的下一块——与旧锚点语义一致，重复键跨块时能落到 run 的首块。
func (s *shard) blockFor(key []byte) int {
	b := sort.Search(int(s.nBlock), func(i int) bool {
		return s.blkPfx[i] >= key8(key)
	}) - 1
	if b < 0 {
		return 0
	}
	return b
}

// seek 返回首个 key ≥ 目标 的位置 (块号, 块内下标)；超过数据末尾返回
// (nBlock, 0)。
func (s *shard) seek(key []byte) (int, int) {
	for b := s.blockFor(key); b < int(s.nBlock); b++ {
		e := s.block(b)
		if e == nil {
			return 0, 0
		}
		n := s.nRecs(b)
		j := sort.Search(n, func(j int) bool {
			return bytes.Compare(e.keyAt(j), key) >= 0
		})
		if j < n {
			return b, j
		}
	}
	return int(s.nBlock), 0
}

// get 精确点查，命中返回 lowerBound 那条（重复键取 run 的第一条）。
func (s *shard) get(key []byte) ([]byte, bool) {
	for b := s.blockFor(key); b < int(s.nBlock); b++ {
		e := s.block(b)
		if e == nil {
			return nil, false
		}
		n := s.nRecs(b)
		j := sort.Search(n, func(j int) bool {
			return bytes.Compare(e.keyAt(j), key) >= 0
		})
		if j >= n {
			continue // 本块全 < key：下一块首键可能恰为 key（run 跨块）
		}
		if !bytes.Equal(e.keyAt(j), key) {
			return nil, false
		}
		vo, fl, raw, packed := e.valLoc(j)
		v := s.value(b*int(s.recsBlk)+j, vo, fl, raw, packed)
		if v == nil {
			return nil, false
		}
		return v, true
	}
	return nil, false
}

// getPrefix 依序回调所有以 prefix 开头的键值对，fn 返回 false 提前停止。
func (s *shard) getPrefix(prefix []byte, fn func(k, v []byte) bool) {
	for b, j := s.seek(prefix); b < int(s.nBlock); b, j = b+1, 0 {
		e := s.block(b)
		if e == nil {
			return
		}
		n := s.nRecs(b)
		for ; j < n; j++ {
			k := e.keyAt(j)
			if !bytes.HasPrefix(k, prefix) {
				return
			}
			vo, fl, raw, packed := e.valLoc(j)
			v := s.value(b*int(s.recsBlk)+j, vo, fl, raw, packed)
			if v == nil || !fn(k, v) {
				return
			}
		}
	}
}

// scan 依序回调键在 [lo, hi) 内的记录；lo 为 nil 从头开始，hi 为 nil 到结尾。
func (s *shard) scan(lo, hi []byte, fn func(k, v []byte) bool) {
	b, j := 0, 0
	if lo != nil {
		b, j = s.seek(lo)
	}
	for ; b < int(s.nBlock); b, j = b+1, 0 {
		e := s.block(b)
		if e == nil {
			return
		}
		n := s.nRecs(b)
		for ; j < n; j++ {
			k := e.keyAt(j)
			if hi != nil && bytes.Compare(k, hi) >= 0 {
				return
			}
			vo, fl, raw, packed := e.valLoc(j)
			v := s.value(b*int(s.recsBlk)+j, vo, fl, raw, packed)
			if v == nil || !fn(k, v) {
				return
			}
		}
	}
}

// packEnt 是一个已解压的打包帧：组头 [firstRec u32][n u32][lens u32×n][values]，
// pfx[i] 是第 i 个成员值在 buf 里的起点（pfx[n] = 值区终点）。
type packEnt struct {
	first uint32
	n     uint32
	pfx   []int32
	buf   []byte
}

func (p *packEnt) value(i uint32) []byte {
	return p.buf[p.pfx[i]:p.pfx[i+1]]
}

// decodePack 解一个打包帧并建组内前缀索引；帧不合法返回 error。dst 可传入
// 复用缓冲（容量够时零分配；流式路径用池，eager 路径传 nil 交给缓存持有）。
func decodePack(dec frameDecoder, frame, dst []byte) (*packEnt, error) {
	buf, err := dec.decode(frame, dst[:0])
	if err != nil {
		return nil, err
	}
	if len(buf) < 8 {
		return nil, errors.New("immkv: 打包帧头不足 8B")
	}
	p := &packEnt{first: binary.LittleEndian.Uint32(buf[0:4]), n: binary.LittleEndian.Uint32(buf[4:8])}
	if int(p.n) < 1 || len(buf) < 8+4*int(p.n) {
		return nil, errors.New("immkv: 打包帧头不合法")
	}
	p.pfx = make([]int32, p.n+1)
	o := int32(8 + 4*int(p.n))
	for i := uint32(0); i < p.n; i++ {
		p.pfx[i] = o
		o += int32(binary.LittleEndian.Uint32(buf[8+4*i:]))
	}
	p.pfx[p.n] = o
	if int(o) != len(buf) {
		return nil, errors.New("immkv: 打包帧长度表与值区不符")
	}
	p.buf = buf
	return p, nil
}

// value 取一条记录的值：帧按 valOff 查 LRU（键取负值，与组窗口的正值键隔离；
// 零长裸存帧不入缓存——它与后一帧共享 valOff，会顶掉别人的缓存位）。打包帧
// 缓存解压后的 packEnt，单值帧缓存解压值。返回切片指向缓存里的缓冲，生命周期
// 到它被逐出为止（契约口径：到下一次查询前）。失败返回 nil（未命中口径）。
func (s *shard) value(rec int, valOff uint64, frameLen uint32, raw, packed bool) []byte {
	if raw && frameLen == 0 {
		return []byte{} // 零长值不入缓存，直接给空切片
	}
	ck := -int64(valOff)
	if s.cache != nil {
		if v, ok := s.cache.get(ck); ok {
			if packed {
				return v.(*packEnt).value(uint32(rec) - v.(*packEnt).first)
			}
			return v.([]byte)
		}
	}
	frame := make([]byte, frameLen)
	if _, err := s.vf.ReadAt(frame, int64(valOff)); err != nil {
		return nil
	}
	if packed {
		p, err := decodePack(s.dec, frame, nil) // 解出进 LRU 缓存，缓冲由缓存持有
		if err != nil {
			return nil
		}
		if uint32(rec) < p.first || uint32(rec) >= p.first+p.n {
			return nil // 数据不一致：记录号不在该帧覆盖的 run 内
		}
		if s.cache != nil {
			s.cache.add(ck, p, len(p.buf)+len(p.pfx)*4+64)
		}
		return p.value(uint32(rec) - p.first)
	}
	var v []byte
	if raw {
		v = frame
	} else {
		var err error
		if v, err = s.dec.decode(frame, nil); err != nil {
			return nil
		}
	}
	if s.cache != nil {
		s.cache.add(ck, v, len(v)+64)
	}
	return v
}

// ---- 分染色体集合与 OpenEager ----

type shardSet struct {
	present [25]bool
	chroms  []uint8
	shards  [25]*shard
}

// openShardSet 打开 dir 下全部成对存在的 chr*.keys/chr*.vals；目录为空 panic。
// cacheMB > 0 时按染色体数均分建 LRU，0 表示不缓存。
func openShardSet(dir string, cacheMB int) (*shardSet, error) {
	ss := &shardSet{}
	for id := uint8(0); id <= 24; id++ {
		base := filepath.Join(dir, "chr"+ChromName(id))
		if _, err := os.Stat(base + ".keys"); err == nil {
			ss.chroms = append(ss.chroms, id)
			ss.present[id] = true
		}
	}
	if len(ss.chroms) == 0 {
		panic(fmt.Sprintf("immkv: %s 里没有任何 chr*.keys/chr*.vals", dir))
	}
	for _, id := range ss.chroms {
		base := filepath.Join(dir, "chr"+ChromName(id))
		// 每 shard 各持一份独立 LRU：缓存键是 shard 内局部编号，
		// 共享实例会跨染色体撞键互相污染
		var cache *lruCache
		if cacheMB > 0 {
			cache = newLRU(cacheMB << 20 / len(ss.chroms))
		}
		s, err := openShard(base, cache)
		if err != nil {
			ss.close()
			return nil, err
		}
		ss.shards[id] = s
	}
	return ss, nil
}

func (ss *shardSet) close() error {
	var firstErr error
	for id := uint8(0); id < 25; id++ {
		if s := ss.shards[id]; s != nil {
			if err := s.Close(); err != nil && firstErr == nil {
				firstErr = err
			}
			ss.shards[id] = nil
		}
	}
	return firstErr
}

// DB 是 OpenEager 的分染色体路由读器；多 goroutine 并发安全。
type DB struct {
	shardSet
}

// OpenEager 打开 dir 下全部存在的染色体（chr1..chr22、chrX、chrY、chrM），
// 锚点索引即时全量载入，不预读数据。不存在的染色体视为无数据：查询路由过去
// 即未命中。cacheMB 是全局解压缓存预算（组窗口 + 值帧），0 不缓存。
// 同一目录可被任意多进程同时打开。
func OpenEager(dir string, cacheMB int) (*DB, error) {
	ss, err := openShardSet(dir, cacheMB)
	if err != nil {
		return nil, err
	}
	return &DB{shardSet: *ss}, nil
}

// Close 关闭全部 shard。调用方保证不与查询并发。
func (d *DB) Close() error { return d.close() }

// Get 精确点查；未命中 ok=false。空键不命中。
func (d *DB) Get(key []byte) ([]byte, bool) {
	if len(key) == 0 || key[0] > 24 || !d.present[key[0]] {
		return nil, false
	}
	return d.shards[key[0]].get(key)
}

// GetPrefix 前缀查询：按 prefix[0] 路由。前缀为空直接返回。
func (d *DB) GetPrefix(prefix []byte, fn func(k, v []byte) bool) {
	if len(prefix) == 0 {
		return
	}
	if prefix[0] <= 24 && d.present[prefix[0]] {
		d.shards[prefix[0]].getPrefix(prefix, fn)
	}
}

// Scan 区间扫描：跨染色体按 chromID 序拼接；lo/hi 可为 nil（nil = 全库扫描），
// hi 为独占上界。lo/hi 原样传给内层文件——只有命中首尾染色体的那一侧才需要设界。
func (d *DB) Scan(lo, hi []byte, fn func(k, v []byte) bool) {
	cLo, cHi := 0, 25
	if len(lo) > 0 {
		cLo = int(lo[0])
	}
	if len(hi) > 0 {
		if int(hi[0]) < cLo {
			return
		}
		cHi = int(hi[0])
	}
	for c := cLo; c <= cHi && c <= 24; c++ {
		s := d.shards[c]
		if s == nil {
			continue
		}
		var l, h []byte
		if len(lo) > 0 && c == cLo {
			l = lo
		}
		if len(hi) > 0 && c == cHi {
			h = hi
		}
		stop := false
		s.scan(l, h, func(k, v []byte) bool {
			if !fn(k, v) {
				stop = true
				return false
			}
			return true
		})
		if stop {
			return
		}
	}
}
