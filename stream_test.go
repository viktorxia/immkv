package immkv

import (
	"bytes"
	"errors"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"testing"
)

func openStream(t *testing.T, dir string, threads int) *StreamDB {
	t.Helper()
	sdb, err := OpenStreaming(dir, StreamOptions{Threads: threads})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { sdb.Close() })
	return sdb
}

func TestGetBatchRoundTrip(t *testing.T) {
	dir := t.TempDir()
	byKey, all := buildTestLib(t, dir, CreateOptions{}, true) // 键流 ~1.5MB，跨多个窗口
	sdb := openStream(t, dir, 4)
	out, err := sdb.GetBatch(all)
	if err != nil {
		t.Fatal(err)
	}
	if len(out) != len(all) {
		t.Fatalf("返回长度 %d ≠ %d", len(out), len(all))
	}
	for i, k := range all {
		if out[i] == nil {
			t.Fatalf("键 %x 应命中", k)
		}
		if !hitIn(byKey, k, out[i]) {
			t.Fatalf("键 %x 值不符", k)
		}
	}
	// 混入未命中键（键后追加 0xFF 保证紧随其后、整批仍升序；真实键以碱基结尾不会撞 0xFF）
	var mixed [][]byte
	for i := 0; i < len(all); i += 3 {
		mixed = append(mixed, all[i], append(bytes.Clone(all[i]), 0xFF))
	}
	out2, err := sdb.GetBatch(mixed)
	if err != nil {
		t.Fatal(err)
	}
	for i, k := range mixed {
		unknown := k[len(k)-1] == 0xFF
		if unknown {
			if out2[i] != nil {
				t.Fatalf("拼造键 %x 不应命中", k)
			}
		} else if out2[i] == nil || !hitIn(byKey, k, out2[i]) {
			t.Fatalf("键 %x 应命中且值相符", k)
		}
	}
}

func TestGetBatchUnsorted(t *testing.T) {
	dir := t.TempDir()
	_, all := buildTestLib(t, dir, CreateOptions{}, false)
	sdb := openStream(t, dir, 1)
	bad := [][]byte{all[0], all[1], all[0]} // 第三个 < 第二个
	if _, err := sdb.GetBatch(bad); !errors.Is(err, ErrUnsorted) {
		t.Fatalf("乱序输入应返回 ErrUnsorted，得到 %v", err)
	}
}

func TestGetBatchEmptyInput(t *testing.T) {
	dir := t.TempDir()
	buildTestLib(t, dir, CreateOptions{}, false)
	sdb := openStream(t, dir, 1)
	out, err := sdb.GetBatch(nil)
	if err != nil || len(out) != 0 {
		t.Fatalf("空输入应返回空切片与 nil，得到 %v %v", out, err)
	}
	out, err = sdb.GetBatch([][]byte{})
	if err != nil || len(out) != 0 {
		t.Fatalf("零长输入应返回空切片与 nil，得到 %v %v", out, err)
	}
}

func TestGetBatchEmptyKeyAndUnknownChrom(t *testing.T) {
	dir := t.TempDir()
	_, all := buildTestLib(t, dir, CreateOptions{}, false)
	sdb := openStream(t, dir, 2)
	// 空键排最前；chrM（24）未构建、排在 chrX（22）之后
	batch := [][]byte{{}, all[0], all[1], VarKey(24, 100, "A", "T")}
	out, err := sdb.GetBatch(batch)
	if err != nil {
		t.Fatal(err)
	}
	if out[0] != nil {
		t.Fatal("空键应未命中")
	}
	if out[1] == nil || out[2] == nil {
		t.Fatal("正常键应命中")
	}
	if out[3] != nil {
		t.Fatal("未构建染色体应未命中")
	}
}

func TestGetBatchThreadsEquivalence(t *testing.T) {
	dir := t.TempDir()
	_, all := buildTestLib(t, dir, CreateOptions{}, false)
	s1 := openStream(t, dir, 1)
	s4 := openStream(t, dir, 4)
	o1, err := s1.GetBatch(all)
	if err != nil {
		t.Fatal(err)
	}
	o4, err := s4.GetBatch(all)
	if err != nil {
		t.Fatal(err)
	}
	for i := range o1 {
		if !bytes.Equal(o1[i], o4[i]) {
			t.Fatalf("下标 %d：threads 1/4 结果不一致", i)
		}
	}
}

// TestGetBatchMatchesGet：批量与单条语义一致——命中与否一致；命中时两者都取
// lowerBound 那条记录（重复键下是同一条），值必须相等。
func TestGetBatchMatchesGet(t *testing.T) {
	dir := t.TempDir()
	byKey, all := buildTestLib(t, dir, CreateOptions{}, false)
	db, err := OpenEager(dir, 32)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	sdb := openStream(t, dir, 3)
	rnd := rand.New(rand.NewSource(9))
	var sample [][]byte
	for i := 0; i < 1500; i++ {
		sample = append(sample, all[rnd.Intn(len(all))])
	}
	for i := 0; i < 500; i++ {
		sample = append(sample, missKey(byKey, rnd))
	}
	// 排序后成批（等值键相邻是允许的）
	sorted := make([][]byte, len(sample))
	copy(sorted, sample)
	sort.Slice(sorted, func(i, j int) bool { return bytes.Compare(sorted[i], sorted[j]) < 0 })
	batch, err := sdb.GetBatch(sorted)
	if err != nil {
		t.Fatal(err)
	}
	for i, k := range sorted {
		sv, ok, err := db.Get(k)
		if err != nil {
			t.Fatal(err)
		}
		if ok != (batch[i] != nil) {
			t.Fatalf("键 %x：Get 命中 %v，GetBatch 命中 %v", k, ok, batch[i] != nil)
		}
		if ok && !bytes.Equal(sv, batch[i]) {
			t.Fatalf("键 %x：Get 与 GetBatch 取值不一致", k)
		}
	}
}

func TestGetBatchDuplicateInputKeys(t *testing.T) {
	dir := t.TempDir()
	byKey, all := buildTestLib(t, dir, CreateOptions{}, false)
	sdb := openStream(t, dir, 2)
	k := all[len(all)/2]
	out, err := sdb.GetBatch([][]byte{k, bytes.Clone(k), k})
	if err != nil {
		t.Fatal(err)
	}
	for i := range out {
		if out[i] == nil || !hitIn(byKey, k, out[i]) {
			t.Fatalf("重复输入键第 %d 项应命中且值相符", i)
		}
	}
}

func TestGetBatchSingleKey(t *testing.T) {
	dir := t.TempDir()
	_, all := buildTestLib(t, dir, CreateOptions{}, false)
	sdb := openStream(t, dir, 1)
	out, err := sdb.GetBatch([][]byte{all[7]})
	if err != nil || len(out) != 1 || out[0] == nil {
		t.Fatalf("单键批应命中：%v %v", out, err)
	}
	out, err = sdb.GetBatch([][]byte{VarKey(0, 1<<30, "A", "T")})
	if err != nil || len(out) != 1 || out[0] != nil {
		t.Fatalf("单键未命中批：%v %v", out, err)
	}
}

// TestGetBatchCorruptFrame：值帧损坏时 GetBatch 整批报错（对比：Get 按未命中
// 口径静默处理）。
func TestGetBatchCorruptFrame(t *testing.T) {
	dir := t.TempDir()
	base := filepath.Join(dir, "chr1")
	w, err := Create(base, CreateOptions{Schema: testRecordSchema})
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < 50; i++ {
		if err := w.Add(VarKey(0, uint32(i), "A", "T"), bytes.Repeat([]byte("val"), 300)); err != nil {
			t.Fatal(err)
		}
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
	recs, _, _ := parseKeys(t, base+".keys")
	var victim []byte
	var vOff uint64
	for _, r := range recs {
		if !r.raw { // 挑一个 zstd 帧
			victim = r.key
			vOff = r.valOff
			break
		}
	}
	if victim == nil {
		t.Fatal("找不到 zstd 帧")
	}
	f, err := os.OpenFile(base+".vals", os.O_WRONLY, 0)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteAt([]byte{0xFF, 0xFF, 0xFF, 0xFF}, int64(vOff)); err != nil {
		t.Fatal(err)
	}
	f.Close()
	sdb := openStream(t, dir, 2)
	if _, err := sdb.GetBatch([][]byte{victim}); err == nil {
		t.Fatal("损坏帧应使 GetBatch 报错")
	}
}

func TestGetBatchConcurrent(t *testing.T) {
	dir := t.TempDir()
	_, all := buildTestLib(t, dir, CreateOptions{}, false)
	sdb := openStream(t, dir, 4)
	var wg sync.WaitGroup
	errs := make(chan error, 8)
	for g := 0; g < 4; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			lo := g * len(all) / 4
			hi := lo + len(all)/4
			for i := 0; i < 5; i++ {
				if _, err := sdb.GetBatch(all[lo:hi]); err != nil {
					errs <- err
					return
				}
			}
		}(g)
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Fatalf("并发 GetBatch: %v", err)
	}
}
