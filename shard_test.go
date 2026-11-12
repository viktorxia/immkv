package immkv

import (
	"bytes"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"testing"
)

// ---- 键编码（自 immdb 迁移） ----

func TestChromID(t *testing.T) {
	cases := []struct {
		in   string
		want uint8
		ok   bool
	}{
		{"1", 0, true}, {"chr1", 0, true}, {"CHR1", 0, true},
		{"22", 21, true}, {"chr22", 21, true},
		{"X", 22, true}, {"chrX", 22, true}, {"x", 22, true},
		{"Y", 23, true}, {"chrY", 23, true},
		{"M", 24, true}, {"MT", 24, true}, {"chrM", 24, true},
		{"0", 0, false}, {"23", 0, false}, {"25", 0, false},
		{"chrUn_KI270709v1_random", 0, false}, {"HLA-A*01:01", 0, false}, {"", 0, false},
	}
	for _, c := range cases {
		got, ok := ChromID(c.in)
		if ok != c.ok || got != c.want {
			t.Errorf("ChromID(%q) = (%d, %v)，期望 (%d, %v)", c.in, got, ok, c.want, c.ok)
		}
	}
}

func TestChromNameRoundTrip(t *testing.T) {
	for id := uint8(0); id <= 24; id++ {
		back, ok := ChromID(ChromName(id))
		if !ok || back != id {
			t.Fatalf("ChromName(%d) = %q 往返失败 (%d, %v)", id, ChromName(id), back, ok)
		}
	}
}

func TestVarKeyLayout(t *testing.T) {
	k := VarKey(0, 0x01020304, "A", "T")
	want := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 'A', 0x00, 'T'}
	if !bytes.Equal(k, want) {
		t.Fatalf("VarKey 布局 %x ≠ %x", k, want)
	}
	// 字节序 = 数值序：位置与染色体都靠前 → 键靠前
	if bytes.Compare(VarKey(0, 1, "A", "T"), VarKey(0, 2, "A", "T")) >= 0 {
		t.Fatal("pos 1 应排在 pos 2 前")
	}
	if bytes.Compare(VarKey(0, 1, "A", "T"), VarKey(1, 1, "A", "T")) >= 0 {
		t.Fatal("chrom 0 应排在 chrom 1 前")
	}
}

func TestPosKeyPrefix(t *testing.T) {
	p := PosKey(3, 100)
	if !bytes.HasPrefix(VarKey(3, 100, "AC", "G"), p) {
		t.Fatal("PosKey 应是同位置 VarKey 的前缀")
	}
	if bytes.HasPrefix(VarKey(3, 101, "A", "T"), p) {
		t.Fatal("不同位置的 VarKey 不应带这个前缀")
	}
}

// ---- 测试库 ----

// genLib 生成位置升序的一批记录（AddSorted 可直接按组吃）：每步前进 1–50，
// 每个位置 1–3 条记录，偶尔刻意制造同键重复。值为 genValue 口径（长短与
// 可压缩性都多样，覆盖裸存阈值两侧与零长）。
func genLib(n int, seed int64, c uint8) []kvRec {
	rnd := rand.New(rand.NewSource(seed))
	var recs []kvRec
	pos := uint32(0)
	for i := 0; i < n; i++ {
		pos += 1 + rnd.Uint32()%50
		for j := 0; j < 1+rnd.Intn(3); j++ {
			key := VarKey(c, pos, randBases(rnd, 1+rnd.Intn(2)), randBases(rnd, 1+rnd.Intn(2)))
			recs = append(recs, kvRec{key, genValue(rnd)})
			if rnd.Intn(50) == 0 {
				recs = append(recs, kvRec{bytes.Clone(key), genValue(rnd)})
			}
		}
	}
	return recs
}

// buildTestLib 在 dir 里构建 chr1（9 千位置，约 2 万条记录，键流约 480KB，
// 跨约 20 组与 2 个流式窗口）与 chrX（1.5 千位置）。sorted=true 用 AddSorted
// 流式，false 用 Add 乱序。返回键→全部值（含重复）与全部键的升序列表（含重复）。
func buildTestLib(t *testing.T, dir string, opts Options, sortedMode bool) (map[string][][]byte, [][]byte) {
	t.Helper()
	if opts.Schema == nil {
		opts.Schema = testRecordSchema
	}
	libA := genLib(9000, 11, 0)
	libB := genLib(1500, 22, 22)
	put := func(c uint8, lib []kvRec) {
		t.Helper()
		w, err := Create(filepath.Join(dir, "chr"+ChromName(c)), opts)
		if err != nil {
			t.Fatal(err)
		}
		if sortedMode {
			for _, r := range lib {
				if err := w.AddSorted(r.key[:5], r.key, r.val); err != nil {
					t.Fatal(err)
				}
			}
		} else {
			idx := rand.New(rand.NewSource(7)).Perm(len(lib))
			for _, i := range idx {
				if err := w.Add(lib[i].key, lib[i].val); err != nil {
					t.Fatal(err)
				}
			}
		}
		if err := w.Finish(); err != nil {
			t.Fatal(err)
		}
	}
	put(0, libA)
	put(22, libB)
	byKey := map[string][][]byte{}
	var all [][]byte
	for _, lib := range [][]kvRec{libA, libB} {
		for _, r := range lib {
			byKey[string(r.key)] = append(byKey[string(r.key)], r.val)
			all = append(all, r.key)
		}
	}
	sort.Slice(all, func(i, j int) bool { return bytes.Compare(all[i], all[j]) < 0 })
	return byKey, all
}

// hitIn 判断 val 是否是指定键的期望值之一（重复键场景）。
func hitIn(byKey map[string][][]byte, key, val []byte) bool {
	for _, v := range byKey[string(key)] {
		if bytes.Equal(v, val) {
			return true
		}
	}
	return false
}

// missKey 生成一个不在期望集里的键。
func missKey(byKey map[string][][]byte, rnd *rand.Rand) []byte {
	for {
		k := VarKey(uint8(rnd.Intn(2))*22, rnd.Uint32(), randBases(rnd, 1), randBases(rnd, 1))
		if _, ok := byKey[string(k)]; !ok {
			return k
		}
	}
}

// ---- OpenEager ----

func TestEagerGetRoundTrip(t *testing.T) {
	dir := t.TempDir()
	byKey, all := buildTestLib(t, dir, Options{}, false)
	db, err := OpenEager(dir, 64)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	for _, k := range all {
		v, ok := db.Get(k)
		if !ok {
			t.Fatalf("键 %x 应命中", k)
		}
		if !hitIn(byKey, k, v) {
			t.Fatalf("键 %x 的值不在期望值列表里", k)
		}
	}
	rnd := rand.New(rand.NewSource(3))
	for i := 0; i < 500; i++ {
		if k := missKey(byKey, rnd); func() bool { _, ok := db.Get(k); return ok }() {
			t.Fatalf("键 %x 不应命中", k)
		}
	}
	if _, ok := db.Get(nil); ok {
		t.Fatal("空键不应命中")
	}
	if _, ok := db.Get([]byte{}); ok {
		t.Fatal("零长键不应命中")
	}
	// 缺席染色体（chr2 未构建）
	if _, ok := db.Get(VarKey(1, 100, "A", "T")); ok {
		t.Fatal("缺席染色体不应命中")
	}
}

func TestEagerNoCacheRoundTrip(t *testing.T) {
	dir := t.TempDir()
	byKey, all := buildTestLib(t, dir, Options{}, true) // 换 AddSorted 构建路径
	db, err := OpenEager(dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	rnd := rand.New(rand.NewSource(5))
	for i := 0; i < 1000; i++ {
		k := all[rnd.Intn(len(all))]
		v, ok := db.Get(k)
		if !ok || !hitIn(byKey, k, v) {
			t.Fatalf("键 %x 未命中或值不符（%v）", k, ok)
		}
	}
}

func TestEagerGetPrefix(t *testing.T) {
	dir := t.TempDir()
	byKey, all := buildTestLib(t, dir, Options{}, false)
	db, err := OpenEager(dir, 64)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	rnd := rand.New(rand.NewSource(6))
	for i := 0; i < 50; i++ {
		p := all[rnd.Intn(len(all))][:5] // 位置前缀
		var wantN int
		for _, k := range all {
			if bytes.HasPrefix(k, p) {
				wantN++
			}
		}
		if wantN == 0 {
			continue
		}
		n := 0
		db.GetPrefix(p, func(k, v []byte) bool {
			if !bytes.HasPrefix(k, p) {
				t.Fatalf("回调给出了不带前缀的键 %x", k)
			}
			if !hitIn(byKey, k, v) {
				t.Fatalf("前缀 %x 的键 %x 值不符", p, k)
			}
			n++
			return true
		})
		if n != wantN {
			t.Fatalf("前缀 %x 命中 %d 条，期望 %d", p, n, wantN)
		}
	}
	// 提前停止：只回调一次
	n := 0
	db.GetPrefix(all[0][:5], func(k, v []byte) bool { n++; return false })
	if n != 1 {
		t.Fatalf("提前停止应只回调 1 次，得到 %d", n)
	}
	// 无命中：不回调
	n = 0
	db.GetPrefix(PosKey(0, 1<<30), func(k, v []byte) bool { n++; return true })
	if n != 0 {
		t.Fatalf("无命中前缀不应回调，得到 %d 次", n)
	}
}

func TestEagerScan(t *testing.T) {
	dir := t.TempDir()
	byKey, all := buildTestLib(t, dir, Options{}, false)
	db, err := OpenEager(dir, 64)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	// 全库扫描：条数、序、值
	var got [][]byte
	db.Scan(nil, nil, func(k, v []byte) bool {
		if !hitIn(byKey, k, v) {
			t.Fatalf("扫描键 %x 值不符", k)
		}
		got = append(got, bytes.Clone(k))
		return true
	})
	if len(got) != len(all) {
		t.Fatalf("全库扫描 %d 条，期望 %d", len(got), len(all))
	}
	for i := 1; i < len(got); i++ {
		if bytes.Compare(got[i-1], got[i]) > 0 {
			t.Fatalf("扫描序在 %d 处乱序", i)
		}
	}
	// 半开区间
	lo, hi := all[len(all)/3], all[2*len(all)/3]
	want := 0
	for _, k := range all {
		if bytes.Compare(k, lo) >= 0 && bytes.Compare(k, hi) < 0 {
			want++
		}
	}
	n := 0
	db.Scan(lo, hi, func(k, v []byte) bool {
		if bytes.Compare(k, lo) < 0 || bytes.Compare(k, hi) >= 0 {
			t.Fatalf("区间扫描越界键 %x", k)
		}
		n++
		return true
	})
	if n != want {
		t.Fatalf("区间扫描 %d 条，期望 %d", n, want)
	}
	// 单侧 nil：从 hi 扫到尾 = 全部键 ≥ hi 的条数
	fromHi := 0
	for _, k := range all {
		if bytes.Compare(k, hi) >= 0 {
			fromHi++
		}
	}
	n = 0
	db.Scan(hi, nil, func(k, v []byte) bool { n++; return true })
	if n != fromHi {
		t.Fatalf("Scan(hi, nil) %d 条，期望 %d", n, fromHi)
	}
	// 提前停止
	n = 0
	db.Scan(nil, nil, func(k, v []byte) bool { n++; return n < 10 })
	if n != 10 {
		t.Fatalf("提前停止应回调 10 次，得到 %d", n)
	}
	// 空区间
	n = 0
	db.Scan(all[len(all)-1], all[0], func(k, v []byte) bool { n++; return true })
	if n != 0 {
		t.Fatalf("空区间不应回调，得到 %d 次", n)
	}
}

func TestEagerSliceLifetime(t *testing.T) {
	dir := t.TempDir()
	byKey, all := buildTestLib(t, dir, Options{}, false)
	db, err := OpenEager(dir, 1) // 1MB 小缓存，扫荡必然逐出
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	k := all[len(all)/2]
	v1, ok := db.Get(k)
	if !ok {
		t.Fatal("应命中")
	}
	want := bytes.Clone(v1)
	for i := 0; i < 2000; i++ { // 大量后续查询把缓存翻几轮
		db.Get(all[(i*37)%len(all)])
	}
	if !bytes.Equal(v1, want) {
		t.Fatal("已返回切片在后续查询后仍应有效（契约：到下一次查询前）")
	}
	_ = byKey
}

func TestOpenEmptyDirPanic(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("空目录应 panic")
		}
	}()
	_, _ = OpenEager(t.TempDir(), 0)
}

func TestOpenBadMagic(t *testing.T) {
	dir := t.TempDir()
	buildTestLib(t, dir, Options{}, false)
	// 破坏 keys 魔数
	f, err := os.OpenFile(filepath.Join(dir, "chr1.keys"), os.O_WRONLY, 0)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteAt([]byte{0xFF}, 0); err != nil {
		t.Fatal(err)
	}
	f.Close()
	if _, err := OpenEager(dir, 0); err == nil {
		t.Fatal("坏魔数应报错")
	}
	// 破坏 vals 魔数（keys 先复原）
	f, _ = os.OpenFile(filepath.Join(dir, "chr1.keys"), os.O_WRONLY, 0)
	f.WriteAt([]byte{keysMagic[0]}, 0)
	f.Close()
	buildTestLib(t, dir, Options{}, false) // 重建干净的
	f, _ = os.OpenFile(filepath.Join(dir, "chrX.vals"), os.O_WRONLY, 0)
	f.WriteAt([]byte{0xFF}, 0)
	f.Close()
	if _, err := OpenEager(dir, 0); err == nil {
		t.Fatal("vals 坏魔数应报错")
	}
}

func TestOpenIncompletePair(t *testing.T) {
	dir := t.TempDir()
	buildTestLib(t, dir, Options{}, false)
	if err := os.Remove(filepath.Join(dir, "chr1.vals")); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenEager(dir, 0); err == nil {
		t.Fatal("缺 .vals 应报错")
	}
}

func TestIgnoresForeignFiles(t *testing.T) {
	dir := t.TempDir()
	byKey, all := buildTestLib(t, dir, Options{}, false)
	for _, name := range []string{"notes.txt", "chrY.keys.bak", "chrM.vals.tmp"} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte("x"), 0o644); err != nil {
			t.Fatal(err)
		}
	}
	db, err := OpenEager(dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if _, ok := db.Get(VarKey(23, 100, "A", "T")); ok { // chrY 未构建
		t.Fatal("chrY 不应命中")
	}
	if v, ok := db.Get(all[0]); !ok || !hitIn(byKey, all[0], v) {
		t.Fatal("正常键应照常命中")
	}
}

func TestReopenAfterReplace(t *testing.T) {
	dir := t.TempDir()
	build := func(tag string) {
		t.Helper()
		w, err := Create(filepath.Join(dir, "chr1"), Options{Schema: testRecordSchema})
		if err != nil {
			t.Fatal(err)
		}
		for i := 0; i < 100; i++ {
			if err := w.Add(VarKey(0, uint32(i), "A", "T"), []byte(tag)); err != nil {
				t.Fatal(err)
			}
		}
		if err := w.Finish(); err != nil {
			t.Fatal(err)
		}
	}
	build("v1")
	db1, err := OpenEager(dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	k := VarKey(0, 50, "A", "T")
	if v, _ := db1.Get(k); string(v) != "v1" {
		t.Fatalf("替换前应读到 v1，得到 %q", v)
	}
	build("v2") // 原子替换
	if v, _ := db1.Get(k); string(v) != "v1" {
		t.Fatalf("旧句柄应继续读旧数据，得到 %q", v)
	}
	db2, err := OpenEager(dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db2.Close()
	if v, _ := db2.Get(k); string(v) != "v2" {
		t.Fatalf("新句柄应读新数据，得到 %q", v)
	}
	db1.Close()
}

func TestEagerMultiDBSameDir(t *testing.T) {
	dir := t.TempDir()
	byKey, all := buildTestLib(t, dir, Options{}, false)
	d1, err := OpenEager(dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer d1.Close()
	d2, err := OpenEager(dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer d2.Close()
	for _, db := range []*DB{d1, d2} {
		v, ok := db.Get(all[0])
		if !ok || !hitIn(byKey, all[0], v) {
			t.Fatal("同目录多句柄应各自正确")
		}
	}
}

func TestConcurrentMixedReads(t *testing.T) {
	dir := t.TempDir()
	_, all := buildTestLib(t, dir, Options{}, false)
	db, err := OpenEager(dir, 8)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	sdb, err := OpenStreaming(dir, Options{Threads: 4})
	if err != nil {
		t.Fatal(err)
	}
	defer sdb.Close()
	var wg sync.WaitGroup
	for g := 0; g < 4; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			rnd := rand.New(rand.NewSource(int64(100 + g)))
			for i := 0; i < 200; i++ {
				switch rnd.Intn(3) {
				case 0:
					db.Get(all[rnd.Intn(len(all))])
				case 1:
					lo := rnd.Intn(len(all) - 1)
					hi := lo + 1 + rnd.Intn(100)
					if hi >= len(all) {
						hi = len(all) - 1
					}
					db.Scan(all[lo], all[hi], func(k, v []byte) bool { return true })
				default:
					span := 200
					if len(all) < span+1 {
						span = len(all)/2 - 1
					}
					lo := rnd.Intn(len(all) - span)
					if _, err := sdb.GetBatch(all[lo : lo+span]); err != nil {
						t.Errorf("并发 GetBatch: %v", err)
						return
					}
				}
			}
		}(g)
	}
	wg.Wait()
}

// TestEmptyChromLib：单染色体空库（count=0）与其他染色体共存。
func TestEmptyChromLib(t *testing.T) {
	dir := t.TempDir()
	w, err := Create(filepath.Join(dir, "chr1"), Options{Schema: testRecordSchema})
	if err != nil {
		t.Fatal(err)
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
	w2, err := Create(filepath.Join(dir, "chrX"), Options{Schema: testRecordSchema})
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 10; i++ {
		if err := w2.Add(VarKey(22, uint32(i), "A", "T"), []byte("x")); err != nil {
			t.Fatal(err)
		}
	}
	if err := w2.Finish(); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if _, ok := db.Get(VarKey(0, 1, "A", "T")); ok {
		t.Fatal("空染色体不应命中")
	}
	n := 0
	db.GetPrefix(PosKey(0, 1), func(k, v []byte) bool { n++; return true })
	if n != 0 {
		t.Fatal("空染色体前缀不应回调")
	}
	sdb, err := OpenStreaming(dir, Options{Threads: 2})
	if err != nil {
		t.Fatal(err)
	}
	defer sdb.Close()
	out, err := sdb.GetBatch([][]byte{VarKey(0, 1, "A", "T"), VarKey(22, 1, "A", "T")})
	if err != nil {
		t.Fatal(err)
	}
	if out[0] != nil || string(out[1]) != "x" {
		t.Fatalf("空染色体 nil、正常染色体命中：得到 %v %q", out[0], out[1])
	}
}
