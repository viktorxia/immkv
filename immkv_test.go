package immkv

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/klauspost/compress/zstd"
)

// ---- 独立解析器：按格式规范从零解析，不复用 Writer 内部 ----
// 子任务 1 的验收依赖这份独立实现交叉核对 Writer 的产物。

type tRec struct {
	off      uint64 // 本条记录在 keys 文件里的起点（keyLen 字段处）
	key      []byte
	valOff   uint64
	frameLen uint32
	raw      bool
	packed   bool
}

type tBlock struct {
	off      uint64 // 块在 keys 文件里的偏移
	pfx      uint64 // 首键 8B 前缀（大端）
	firstKey []byte
}

func parseKeys(t *testing.T, path string) (recs []tRec, blocks []tBlock, count uint64) {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if len(b) < headerSize || string(b[0:8]) != keysMagic {
		t.Fatalf("keys 头不合法")
	}
	if v := binary.LittleEndian.Uint32(b[8:12]); v != 1 {
		t.Fatalf("version = %d", v)
	}
	count = binary.LittleEndian.Uint64(b[12:20])
	rpb := uint64(binary.LittleEndian.Uint16(b[20:22]))
	if rpb != recsPerBlock {
		t.Fatalf("recsPerBlock = %d", rpb)
	}
	nBlock := binary.LittleEndian.Uint32(b[24:28])
	indexOff := binary.LittleEndian.Uint64(b[28:36])
	// 块区顺序解析（独立实现，不复用引擎 parseBlock）
	p := uint64(headerSize)
	var winB0, prevEnd uint64
	for bi := uint32(0); bi < nBlock; bi++ {
		n := count - uint64(bi)*rpb
		if n > rpb {
			n = rpb
		}
		blk := tBlock{off: p}
		keyEnds := make([]uint32, n)
		for i := uint64(0); i < n; i++ {
			keyEnds[i] = binary.LittleEndian.Uint32(b[p+4*i:])
		}
		lensAt := p + 4*n
		lens := make([]uint32, n)
		for i := uint64(0); i < n; i++ {
			lens[i] = binary.LittleEndian.Uint32(b[lensAt+4*i:])
		}
		base := binary.LittleEndian.Uint64(b[p+8*n:])
		keysAt := p + 8*n + 8
		keyLen := uint64(keyEnds[n-1])
		keys := b[keysAt : keysAt+keyLen]
		_ = winB0
		_ = prevEnd
		prev := uint32(0)
		valOff := base
		for i := uint64(0); i < n; i++ {
			if i > 0 && lens[i]&flagShare == 0 {
				valOff += uint64(lens[i-1] & frameLenMask)
			}
			// i=0：valOff = base；i>0 且 share：同上一条；否则前推
			// i=0 用 base；i>0 share 则同值，否则前推
			r := tRec{
				key:      append([]byte{}, keys[prev:keyEnds[i]]...),
				valOff:   valOff,
				frameLen: lens[i] & frameLenMask,
				raw:      lens[i]&flagRaw != 0,
				packed:   lens[i]&flagPacked != 0,
			}
			recs = append(recs, r)
			prev = keyEnds[i]
		}
		blk.firstKey = append([]byte{}, keys[:keyEnds[0]]...)
		var p8 [8]byte
		copy(p8[:], blk.firstKey)
		blk.pfx = binary.BigEndian.Uint64(p8[:])
		blocks = append(blocks, blk)
		p = keysAt + keyLen
	}
	if p != indexOff {
		t.Fatalf("块区终点 %d ≠ indexOff %d", p, indexOff)
	}
	if uint64(len(recs)) != count {
		t.Fatalf("count 头 %d 与块区 %d 条不符", count, len(recs))
	}
	// 索引区：prefix 大端 ×N + offset 小端 ×N，与块区对照
	if uint64(len(b))-indexOff != 16*uint64(nBlock) {
		t.Fatalf("索引区大小不符")
	}
	for i := uint32(0); i < nBlock; i++ {
		if got := binary.BigEndian.Uint64(b[indexOff+8*uint64(i):]); got != blocks[i].pfx {
			t.Fatalf("块 %d 索引前缀不符", i)
		}
		if got := binary.LittleEndian.Uint64(b[indexOff+8*uint64(nBlock)+8*uint64(i):]); got != blocks[i].off {
			t.Fatalf("块 %d 索引偏移不符", i)
		}
	}
	return recs, blocks, count
}

// verifyBuild 交叉核对一次构建：记录有序、帧可还原成期望值、帧区无缝、
// 锚点数与对位正确。
func verifyBuild(t *testing.T, base string, expected map[string][][]byte) {
	t.Helper()
	recs, blocks, count := parseKeys(t, base+".keys")
	if count == 0 {
		t.Fatal("count = 0，应先验证空库用例")
	}
	for i := 1; i < len(recs); i++ {
		if bytes.Compare(recs[i-1].key, recs[i].key) > 0 {
			t.Fatalf("记录 %d 乱序", i)
		}
	}
	vb, err := os.ReadFile(base + ".vals")
	if err != nil {
		t.Fatal(err)
	}
	if len(vb) < headerSize || string(vb[0:8]) != valsMagic {
		t.Fatalf("vals 头不合法")
	}
	if got := binary.LittleEndian.Uint64(vb[12:20]); got != count {
		t.Fatalf("vals count %d ≠ keys count %d", got, count)
	}
	dec, err := zstd.NewReader(nil)
	if err != nil {
		t.Fatal(err)
	}
	defer dec.Close()
	exp := make(map[string][][]byte, len(expected))
	for k, vs := range expected {
		exp[k] = append([][]byte{}, vs...)
	}
	// 帧区按记录序顺序消费：零长裸存帧不占字节、与后续帧共享 valOff；打包帧
	// 多条记录共享一个 valOff，成员按记录序出现、顺序消耗组内前缀。
	frameEnd := uint64(headerSize)
	curOff := ^uint64(0) // 当前帧的 valOff（哨兵 = 还没消费过帧）
	var pack struct {
		first uint32
		pfx   []int32
		buf   []byte
		idx   int // 已消费的成员数
	}
	check := func(i int, r tRec, got []byte) {
		t.Helper()
		vs := exp[string(r.key)]
		found := false
		for j, v := range vs {
			if bytes.Equal(v, got) {
				vs = append(vs[:j], vs[j+1:]...)
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("键 %q 的值（%d 字节）不在期望值列表里", r.key, len(got))
		}
		exp[string(r.key)] = vs
	}
	for i, r := range recs {
		var got []byte
		switch {
		case r.raw && r.frameLen == 0:
			if r.valOff != frameEnd {
				t.Fatalf("记录 %d 零长帧 valOff=%d ≠ 帧尾 %d", i, r.valOff, frameEnd)
			}
			got = []byte{}
		case r.packed && r.valOff == curOff:
			// 同一打包帧的后续成员
			if pack.idx >= len(pack.pfx)-1 {
				t.Fatalf("记录 %d 超出打包帧成员数", i)
			}
			got = pack.buf[pack.pfx[pack.idx]:pack.pfx[pack.idx+1]]
			pack.idx++
			if want := pack.first + uint32(pack.idx-1); want != uint32(i) {
				t.Fatalf("打包帧成员错位：帧内第 %d 个成员应为记录 %d，实际 %d", pack.idx-1, want, i)
			}
		default:
			if r.valOff != frameEnd {
				t.Fatalf("记录 %d valOff=%d，期望帧区无缝续接于 %d", i, r.valOff, frameEnd)
			}
			curOff = r.valOff
			frame := vb[r.valOff : r.valOff+uint64(r.frameLen)]
			frameEnd += uint64(r.frameLen)
			switch {
			case r.raw:
				got = frame
			case r.packed:
				buf, err := dec.DecodeAll(frame, nil)
				if err != nil {
					t.Fatalf("记录 %d 打包帧解压失败: %v", i, err)
				}
				if len(buf) < 8 {
					t.Fatalf("记录 %d 打包帧头不足", i)
				}
				n := int(binary.LittleEndian.Uint32(buf[4:8]))
				if n < 1 || len(buf) < 8+4*n {
					t.Fatalf("记录 %d 打包帧头不合法", i)
				}
				pack.first = binary.LittleEndian.Uint32(buf[0:4])
				pack.pfx = pack.pfx[:0]
				o := int32(8 + 4*n)
				for m := 0; m < n; m++ {
					pack.pfx = append(pack.pfx, o)
					o += int32(binary.LittleEndian.Uint32(buf[8+4*m:]))
				}
				pack.pfx = append(pack.pfx, o) // 末项哨兵 = 值区终点
				if int(o) != len(buf) {
					t.Fatalf("记录 %d 打包帧长度表与值区不符", i)
				}
				pack.buf = buf
				pack.idx = 1
				if pack.first != uint32(i) {
					t.Fatalf("打包帧首记录号 %d ≠ 实际 %d", pack.first, i)
				}
				got = pack.buf[pack.pfx[0]:pack.pfx[1]]
			default:
				if got, err = dec.DecodeAll(frame, nil); err != nil {
					t.Fatalf("记录 %d（键 %q）解压失败: %v", i, r.key, err)
				}
			}
		}
		check(i, r, got)
	}
	if frameEnd != uint64(len(vb)) {
		t.Fatalf("vals 文件尾部剩余 %d 字节", uint64(len(vb))-frameEnd)
	}
	wantBlocks := (count + recsPerBlock - 1) / recsPerBlock
	if uint64(len(blocks)) != wantBlocks {
		t.Fatalf("块数 %d，期望 %d", len(blocks), wantBlocks)
	}
	for i, blk := range blocks {
		j := uint64(i) * recsPerBlock
		if !bytes.Equal(blk.firstKey, recs[j].key) {
			t.Fatalf("块 %d 首键与记录 %d 的键不符", i, j)
		}
		if i > 0 && bytes.Compare(recs[(uint64(i)-1)*recsPerBlock].key, blk.firstKey) > 0 {
			t.Fatalf("块 %d 首键小于上一块首键（索引序破坏）", i)
		}
	}
}

// ---- 测试数据 ----

func randBases(rnd *rand.Rand, n int) string {
	const bases = "ACGTN"
	b := make([]byte, n)
	for i := range b {
		b[i] = bases[rnd.Intn(len(bases))]
	}
	return string(b)
}

// genValue 生成长度与可压缩性都多样的值：覆盖裸存阈值两侧、打包阈值两侧
// （≥ packBelow 的走单值帧）与零长。
func genValue(rnd *rand.Rand) []byte {
	n := rnd.Intn(6000)
	if rnd.Intn(4) == 0 {
		n = rnd.Intn(100)
	}
	if rnd.Intn(60) == 0 {
		n = packBelow + rnd.Intn(6000)
	}
	buf := make([]byte, n)
	if rnd.Intn(2) == 0 {
		for i := range buf {
			buf[i] = byte('A' + rnd.Intn(4))
		}
	} else {
		rnd.Read(buf)
	}
	return buf
}

// ---- 用例 ----

func TestWriterAddRoundTrip(t *testing.T) {
	base := filepath.Join(t.TempDir(), "chr1")
	w, err := Create(base, Options{})
	if err != nil {
		t.Fatal(err)
	}
	rnd := rand.New(rand.NewSource(1))
	expected := map[string][][]byte{}
	for i := 0; i < 3000; i++ {
		key := VarKey(uint8(rnd.Intn(25)), rnd.Uint32(), randBases(rnd, rnd.Intn(3)+1), randBases(rnd, rnd.Intn(3)+1))
		val := genValue(rnd)
		expected[string(key)] = append(expected[string(key)], val)
		if err := w.Add(key, val); err != nil {
			t.Fatal(err)
		}
	}
	// 刻意制造重复键：同一键多个值都要在库里
	for i := 0; i < 8; i++ {
		key := VarKey(1, 42, "A", "T")
		v1, v2 := []byte{byte(i)}, []byte("dup")
		expected[string(key)] = append(expected[string(key)], v1, v2)
		if err := w.Add(key, v1); err != nil {
			t.Fatal(err)
		}
		if err := w.Add(key, v2); err != nil {
			t.Fatal(err)
		}
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
	verifyBuild(t, base, expected)
	for _, tmp := range []string{base + ".keys.tmp", base + ".vals.tmp"} {
		if _, err := os.Stat(tmp); !os.IsNotExist(err) {
			t.Fatalf("%s 残留", tmp)
		}
	}
}

func TestWriterAddSortedRoundTrip(t *testing.T) {
	base := filepath.Join(t.TempDir(), "chr2")
	w, err := Create(base, Options{})
	if err != nil {
		t.Fatal(err)
	}
	rnd := rand.New(rand.NewSource(2))
	expected := map[string][][]byte{}
	// 位置升序（跨两条染色体考验组序），同一位置 1–3 条记录乱序到达
	for pos := uint32(1); pos <= 2500; pos++ {
		c := uint8(0)
		if pos >= 1200 {
			c = 7
		}
		type kv struct{ key, val []byte }
		grp := make([]kv, 0, 3)
		for j := 0; j < 1+rnd.Intn(3); j++ {
			key := VarKey(c, pos, randBases(rnd, rnd.Intn(2)+1), randBases(rnd, rnd.Intn(2)+1))
			grp = append(grp, kv{key, genValue(rnd)})
		}
		rnd.Shuffle(len(grp), func(i, j int) { grp[i], grp[j] = grp[j], grp[i] })
		for _, r := range grp {
			expected[string(r.key)] = append(expected[string(r.key)], r.val)
			if err := w.AddSorted(PosKey(c, pos), r.key, r.val); err != nil {
				t.Fatal(err)
			}
		}
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
	verifyBuild(t, base, expected)
}

func TestWriterOrderingErrors(t *testing.T) {
	dir := t.TempDir()
	v := []byte("v")

	// group 倒序
	w, err := Create(filepath.Join(dir, "a"), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if err := w.AddSorted(PosKey(0, 2), VarKey(0, 2, "A", "T"), v); err != nil {
		t.Fatal(err)
	}
	err = w.AddSorted(PosKey(0, 1), VarKey(0, 1, "A", "T"), v)
	if err == nil || !strings.Contains(err.Error(), "out of order") {
		t.Fatalf("group 倒序应报错，得到 %v", err)
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}

	// 组间键倒挂：组键升序但键低于已进入缓冲的最大键
	w, err = Create(filepath.Join(dir, "b"), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if err := w.AddSorted(PosKey(0, 100), VarKey(0, 100, "A", "T"), v); err != nil {
		t.Fatal(err)
	}
	err = w.AddSorted(PosKey(0, 200), VarKey(0, 90, "C", "G"), v)
	if err == nil || !strings.Contains(err.Error(), "below last emitted") {
		t.Fatalf("组间倒挂应报错，得到 %v", err)
	}

	// 两种模式混用
	w, err = Create(filepath.Join(dir, "c"), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if err := w.Add(VarKey(0, 1, "A", "T"), v); err != nil {
		t.Fatal(err)
	}
	if err := w.AddSorted(PosKey(0, 2), VarKey(0, 2, "A", "T"), v); err == nil || !strings.Contains(err.Error(), "AddSorted after Add") {
		t.Fatalf("Add 后 AddSorted 应报错，得到 %v", err)
	}
	w, err = Create(filepath.Join(dir, "d"), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if err := w.AddSorted(PosKey(0, 2), VarKey(0, 2, "A", "T"), v); err != nil {
		t.Fatal(err)
	}
	if err := w.Add(VarKey(0, 1, "A", "T"), v); err == nil || !strings.Contains(err.Error(), "Add after AddSorted") {
		t.Fatalf("AddSorted 后 Add 应报错，得到 %v", err)
	}

	// Finish 两次
	w, err = Create(filepath.Join(dir, "e"), Options{})
	if err != nil {
		t.Fatal(err)
	}
	if err := w.Add(VarKey(0, 1, "A", "T"), v); err != nil {
		t.Fatal(err)
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
	if err := w.Finish(); err == nil {
		t.Fatal("二次 Finish 应报错")
	}
}

func TestWriterRawThreshold(t *testing.T) {
	base := filepath.Join(t.TempDir(), "chr3")
	w, err := Create(base, Options{RawValueBelow: 128})
	if err != nil {
		t.Fatal(err)
	}
	tiny := []byte("short")                // 5B < 128 → 裸存
	empty := []byte{}                      // 0B < 128 → 裸存
	big := bytes.Repeat([]byte("A"), 1000) // 128 ≤ len < packBelow → 打包帧
	def63 := bytes.Repeat([]byte("z"), 63) // 默认阈值 64 下应裸存，但本构建阈值 128
	expect := map[string][][]byte{}
	put := func(k []byte, v []byte) {
		expect[string(k)] = append(expect[string(k)], v)
		if err := w.Add(k, v); err != nil {
			t.Fatal(err)
		}
	}
	put(VarKey(0, 1, "A", "T"), tiny)
	put(VarKey(0, 2, "A", "T"), empty)
	put(VarKey(0, 3, "A", "T"), big)
	put(VarKey(0, 4, "A", "T"), def63)
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
	verifyBuild(t, base, expect)
	recs, _, _ := parseKeys(t, base+".keys")
	wantRaw := map[int]bool{0: true, 1: true, 2: false, 3: true}
	for i, r := range recs {
		if r.raw != wantRaw[i] {
			t.Fatalf("记录 %d 裸存标记 = %v，期望 %v", i, r.raw, wantRaw[i])
		}
	}
	// 头里记录的构建参数
	kb, _ := os.ReadFile(base + ".keys")
	if got := binary.LittleEndian.Uint32(kb[36:40]); got != 128 {
		t.Fatalf("rawBelow 头 = %d，期望 128", got)
	}
}

// TestWriterDefaultThreshold 验证 RawValueBelow=0 取默认 64：63 裸存、64 压缩。
func TestWriterDefaultThreshold(t *testing.T) {
	base := filepath.Join(t.TempDir(), "chr4")
	w, err := Create(base, Options{})
	if err != nil {
		t.Fatal(err)
	}
	if err := w.Add(VarKey(0, 1, "A", "T"), bytes.Repeat([]byte("z"), 63)); err != nil {
		t.Fatal(err)
	}
	if err := w.Add(VarKey(0, 2, "A", "T"), bytes.Repeat([]byte("z"), 64)); err != nil {
		t.Fatal(err)
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
	recs, _, _ := parseKeys(t, base+".keys")
	if !recs[0].raw || recs[1].raw {
		t.Fatalf("默认阈值 64 下裸存标记应为 [true false]，得到 [%v %v]", recs[0].raw, recs[1].raw)
	}
}

func TestWriterEmpty(t *testing.T) {
	base := filepath.Join(t.TempDir(), "chr5")
	w, err := Create(base, Options{})
	if err != nil {
		t.Fatal(err)
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
	kb, err := os.ReadFile(base + ".keys")
	if err != nil {
		t.Fatal(err)
	}
	if len(kb) != headerSize {
		t.Fatalf("空库 keys 文件应为 %dB，得到 %dB", headerSize, len(kb))
	}
	vb, err := os.ReadFile(base + ".vals")
	if err != nil {
		t.Fatal(err)
	}
	if len(vb) != headerSize {
		t.Fatalf("空库 vals 文件应为 %dB，得到 %dB", headerSize, len(vb))
	}
	if got := binary.LittleEndian.Uint64(kb[12:20]); got != 0 {
		t.Fatalf("空库 count = %d", got)
	}
	if got := binary.LittleEndian.Uint32(kb[24:28]); got != 0 {
		t.Fatalf("空库锚点数 = %d", got)
	}
}

// TestWriterPacked 验证小值打包帧：多条记录共享一个 valOff、组头带首记录号与
// 长度表、组内上下文把共享长前缀的小值压到远小于原文；OpenEager 与
// OpenStreaming 都能逐条读回。
func TestWriterPacked(t *testing.T) {
	base := filepath.Join(t.TempDir(), "chr1")
	w, err := Create(base, Options{})
	if err != nil {
		t.Fatal(err)
	}
	prefix := bytes.Repeat([]byte("CLNSIG=pathogenic_or_likely_pathogenic;ORIGIN=germline;"), 5) // 325B 公共前缀
	expect := map[string][][]byte{}
	var rawTotal uint64
	for i := 0; i < 2000; i++ {
		val := append(append([]byte{}, prefix...), []byte(fmt.Sprintf("N=%06d", i))...)
		key := VarKey(0, uint32(i), "A", "T")
		expect[string(key)] = append(expect[string(key)], val)
		if err := w.Add(key, val); err != nil {
			t.Fatal(err)
		}
		rawTotal += uint64(len(val))
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
	verifyBuild(t, base, expect)
	recs, _, _ := parseKeys(t, base+".keys")
	frames := map[uint64]uint32{} // valOff → 帧长（打包帧多记录共享，只计一次）
	for i, r := range recs {
		if r.raw || !r.packed {
			t.Fatalf("记录 %d 应为打包帧（raw=%v packed=%v）", i, r.raw, r.packed)
		}
		frames[r.valOff] = r.frameLen
	}
	if len(frames) >= 20 {
		t.Fatalf("2000 条 331B 值应攒成少数几个打包帧，实际 %d 个", len(frames))
	}
	// 打包帧合计远小于原文（组内上下文 + 长度表分摊）
	var frameTotal uint64
	for _, l := range frames {
		frameTotal += uint64(l)
	}
	if frameTotal > rawTotal/6 {
		t.Fatalf("打包帧合计 %d，超过原文 %d 的 1/6", frameTotal, rawTotal)
	}
	// 端到端：两种句柄读回全部值
	db, err := OpenEager(filepath.Dir(base), 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	sdb, err := OpenStreaming(filepath.Dir(base), Options{Threads: 2})
	if err != nil {
		t.Fatal(err)
	}
	defer sdb.Close()
	keys := make([][]byte, 0, len(recs))
	for _, r := range recs {
		keys = append(keys, r.key)
	}
	batch, err := sdb.GetBatch(keys)
	if err != nil {
		t.Fatal(err)
	}
	for i, r := range recs {
		want := append(append([]byte{}, prefix...), []byte(fmt.Sprintf("N=%06d", i))...)
		if v, ok := db.Get(r.key); !ok || !bytes.Equal(v, want) {
			t.Fatalf("OpenEager 记录 %d 读回不符", i)
		}
		if !bytes.Equal(batch[i], want) {
			t.Fatalf("GetBatch 记录 %d 读回不符", i)
		}
	}
}

// TestWriterKeyTooLong 验证键超 u16 上限时 panic（格式装不下，属数据问题）。
func TestWriterKeyTooLong(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("超长键应 panic")
		}
	}()
	base := filepath.Join(t.TempDir(), "chr6")
	w, err := Create(base, Options{})
	if err != nil {
		t.Fatal(err)
	}
	if err := w.Add(make([]byte, maxKeyLen+1), []byte("v")); err != nil {
		t.Fatal(err)
	}
	_ = w.Finish()
}
