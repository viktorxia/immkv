package immkv

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"testing"
)

type simplePair struct {
	key string
	val []byte
}

// simplePairs 覆盖三种值帧：短值裸存、packed 帧、单值帧，外加零长值。
func simplePairs() []simplePair {
	return []simplePair{
		{"BRCA1", []byte(`{"gene":"BRCA1","length":81189}`)},
		{"EGFR", bytes.Repeat([]byte("x"), 512)},
		{"EMPTY", []byte{}},
		{"LONG", bytes.Repeat([]byte("y"), 40000)},
		{"TP53", []byte("tumor suppressor")},
	}
}

// getOK 调用 Get 并要求无错误；返回 (value, found)。仅用于期望读取成功的用例。
func getOK(t *testing.T, kv *SimpleKVDB, key []byte) ([]byte, bool) {
	t.Helper()
	v, ok, err := kv.Get(key)
	if err != nil {
		t.Fatalf("Get(%x): %v", key, err)
	}
	return v, ok
}

func buildSimpleLib(t *testing.T, base string, sortedMode bool) {
	t.Helper()
	w, err := Create(base, CreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	pairs := simplePairs()
	if sortedMode {
		for _, p := range pairs {
			if err := w.AddSorted([]byte(p.key), []byte(p.key), p.val); err != nil {
				t.Fatal(err)
			}
		}
	} else {
		for i := len(pairs) - 1; i >= 0; i-- {
			p := pairs[i]
			if err := w.Add([]byte(p.key), p.val); err != nil {
				t.Fatal(err)
			}
		}
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
}

func checkSimpleLib(t *testing.T, kv *SimpleKVDB) {
	t.Helper()
	for _, p := range simplePairs() {
		v, ok := getOK(t, kv, []byte(p.key))
		if !ok {
			t.Fatalf("Get(%q) miss", p.key)
		}
		if !bytes.Equal(v, p.val) {
			t.Fatalf("Get(%q) = %d bytes, want %d bytes", p.key, len(v), len(p.val))
		}
	}
	for _, k := range []string{"", "BRCA", "BRCA2", "ZZZ"} {
		if v, ok := getOK(t, kv, []byte(k)); ok {
			t.Fatalf("Get(%q) = %d bytes, want miss", k, len(v))
		}
	}
}

func TestSimpleKVRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		name   string
		sorted bool
	}{
		{"Add", false},
		{"AddSorted", true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			base := filepath.Join(t.TempDir(), "genes")
			buildSimpleLib(t, base, tc.sorted)
			kv, err := OpenSimpleKV(base, 0)
			if err != nil {
				t.Fatal(err)
			}
			defer kv.Close()
			checkSimpleLib(t, kv)
		})
	}
}

func TestSimpleKVCache(t *testing.T) {
	base := filepath.Join(t.TempDir(), "genes")
	buildSimpleLib(t, base, false)
	for _, cacheMB := range []int{-1, 0, 1} {
		kv, err := OpenSimpleKV(base, cacheMB)
		if err != nil {
			t.Fatalf("cacheMB %d: %v", cacheMB, err)
		}
		checkSimpleLib(t, kv)
		if err := kv.Close(); err != nil {
			t.Fatalf("cacheMB %d: Close: %v", cacheMB, err)
		}
	}
}

// TestSimpleKVDuplicateKey 验证重复键的行为：Get 返回排序后的第一条，
// 多次查询结果一致。
func TestSimpleKVDuplicateKey(t *testing.T) {
	base := filepath.Join(t.TempDir(), "dup")
	w, err := Create(base, CreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range [][]byte{[]byte("first"), []byte("second")} {
		if err := w.Add([]byte("DUP"), v); err != nil {
			t.Fatal(err)
		}
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
	kv, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer kv.Close()
	v, ok := getOK(t, kv, []byte("DUP"))
	if !ok {
		t.Fatal("Get(DUP) miss")
	}
	if !bytes.Equal(v, []byte("first")) && !bytes.Equal(v, []byte("second")) {
		t.Fatalf("Get(DUP) = %q, want one of the duplicates", v)
	}
	v2, _ := getOK(t, kv, []byte("DUP"))
	if !bytes.Equal(v, v2) {
		t.Fatalf("duplicate lookup not deterministic: %q vs %q", v, v2)
	}
}

func TestSimpleKVConcurrent(t *testing.T) {
	base := filepath.Join(t.TempDir(), "genes")
	buildSimpleLib(t, base, false)
	kv, err := OpenSimpleKV(base, 4)
	if err != nil {
		t.Fatal(err)
	}
	defer kv.Close()
	pairs := simplePairs()
	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for i := 0; i < 200; i++ {
				if i%10 == 0 {
					if _, ok, err := kv.Get([]byte("MISS")); err != nil || ok {
						t.Errorf("goroutine %d: Get(MISS) hit or error: %v", g, err)
						return
					}
					continue
				}
				p := pairs[(g+i)%len(pairs)]
				v, ok, err := kv.Get([]byte(p.key))
				if err != nil || !ok || !bytes.Equal(v, p.val) {
					t.Errorf("goroutine %d: Get(%q) miss or wrong value (err %v)", g, p.key, err)
					return
				}
			}
		}(g)
	}
	wg.Wait()
}

func TestSimpleKVCloseReopen(t *testing.T) {
	base := filepath.Join(t.TempDir(), "genes")
	buildSimpleLib(t, base, true)
	kv, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := getOK(t, kv, []byte("BRCA1")); !ok {
		t.Fatal("Get(BRCA1) miss before close")
	}
	if err := kv.Close(); err != nil {
		t.Fatal(err)
	}
	kv2, err := OpenSimpleKV(base, 1)
	if err != nil {
		t.Fatal(err)
	}
	defer kv2.Close()
	checkSimpleLib(t, kv2)
}

func TestOpenSimpleKVMissing(t *testing.T) {
	if _, err := OpenSimpleKV(filepath.Join(t.TempDir(), "nope"), 0); err == nil {
		t.Fatal("missing library should error")
	}
}

type simpleRec struct {
	key, val []byte
}

func buildSimpleRecords(t *testing.T, base string, opts CreateOptions, recs []simpleRec, sortedMode bool) {
	t.Helper()
	w, err := Create(base, opts)
	if err != nil {
		t.Fatal(err)
	}
	if sortedMode {
		for _, r := range recs {
			if err := w.AddSorted(r.key, r.key, r.val); err != nil {
				t.Fatal(err)
			}
		}
	} else {
		for i := len(recs) - 1; i >= 0; i-- {
			r := recs[i]
			if err := w.Add(r.key, r.val); err != nil {
				t.Fatal(err)
			}
		}
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
}

// removeSimpleArtifacts 删除单库的全部构建产物，供需要重建同一 base 的测试使用。
func removeSimpleArtifacts(t *testing.T, base string) {
	t.Helper()
	for _, ext := range []string{".keys", ".vals", ".keys.tmp", ".vals.tmp", ".col", ".manifest"} {
		if err := os.Remove(base + ext); err != nil && !os.IsNotExist(err) {
			t.Fatal(err)
		}
	}
}

func checkSimpleRecords(t *testing.T, kv *SimpleKVDB, recs []simpleRec) {
	t.Helper()
	for _, r := range recs {
		v, ok := getOK(t, kv, r.key)
		if !ok {
			t.Fatalf("Get(key len %d) miss", len(r.key))
		}
		if !bytes.Equal(v, r.val) {
			t.Fatalf("Get(key len %d) = %d bytes, want %d bytes", len(r.key), len(v), len(r.val))
		}
	}
}

func TestSimpleKVEmptyLib(t *testing.T) {
	base := filepath.Join(t.TempDir(), "empty")
	w, err := Create(base, CreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
	kv, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer kv.Close()
	if v, ok := getOK(t, kv, []byte("BRCA1")); ok {
		t.Fatalf("empty lib Get = %d bytes, want miss", len(v))
	}
	if v, ok := getOK(t, kv, nil); ok {
		t.Fatalf("empty lib Get(nil) = %d bytes, want miss", len(v))
	}
}

func TestSimpleKVMicroBlockBoundary(t *testing.T) {
	for _, n := range []int{63, 64, 65, 128, 129} {
		t.Run(fmt.Sprintf("n%d", n), func(t *testing.T) {
			recs := make([]simpleRec, n)
			for i := range recs {
				recs[i] = simpleRec{
					key: []byte(fmt.Sprintf("k%04d", i)),
					val: []byte(fmt.Sprintf("v%d", i)),
				}
			}
			base := filepath.Join(t.TempDir(), "lib")
			buildSimpleRecords(t, base, CreateOptions{}, recs, true)
			kv, err := OpenSimpleKV(base, 0)
			if err != nil {
				t.Fatal(err)
			}
			defer kv.Close()
			checkSimpleRecords(t, kv, recs)
			next := []byte(fmt.Sprintf("k%04d", n))
			if v, ok := getOK(t, kv, next); ok {
				t.Fatalf("Get(%q) = %d bytes, want miss", next, len(v))
			}
		})
	}
}

func TestSimpleKVDuplicateAcrossBlocks(t *testing.T) {
	var recs []simpleRec
	for i := 0; i < 62; i++ {
		recs = append(recs, simpleRec{key: []byte(fmt.Sprintf("A%04d", i)), val: []byte("A")})
	}
	for i := 0; i < 4; i++ {
		recs = append(recs, simpleRec{key: []byte("DUP"), val: []byte(fmt.Sprintf("d%d", i))})
	}
	for i := 0; i < 4; i++ {
		recs = append(recs, simpleRec{key: []byte(fmt.Sprintf("Z%04d", i)), val: []byte("Z")})
	}
	base := filepath.Join(t.TempDir(), "lib")
	buildSimpleRecords(t, base, CreateOptions{}, recs, true)
	kv, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer kv.Close()
	v, ok := getOK(t, kv, []byte("DUP"))
	if !ok {
		t.Fatal("Get(DUP) miss")
	}
	dupOK := false
	for i := 0; i < 4; i++ {
		if bytes.Equal(v, []byte(fmt.Sprintf("d%d", i))) {
			dupOK = true
			break
		}
	}
	if !dupOK {
		t.Fatalf("Get(DUP) = %q, want one of the duplicates", v)
	}
	v2, _ := getOK(t, kv, []byte("DUP"))
	if !bytes.Equal(v, v2) {
		t.Fatalf("duplicate lookup not deterministic: %q vs %q", v, v2)
	}
	for _, kvp := range []struct{ key, val string }{
		{"A0061", "A"},
		{"Z0003", "Z"},
	} {
		got, ok := getOK(t, kv, []byte(kvp.key))
		if !ok || !bytes.Equal(got, []byte(kvp.val)) {
			t.Fatalf("Get(%q) = %q, %v", kvp.key, got, ok)
		}
	}
}

func TestSimpleKVKeyBoundaries(t *testing.T) {
	recs := []simpleRec{
		{key: []byte{}, val: []byte("empty")},
		{key: []byte{0x00}, val: []byte("zero")},
		{key: []byte{0x18}, val: []byte("24")},
		{key: []byte{0x19}, val: []byte("25")},
		{key: []byte("A"), val: []byte("A")},
		{key: []byte("AA"), val: []byte("AA")},
		{key: []byte("AAA"), val: []byte("AAA")},
		{key: []byte{0xFF, 0x00, 0xFF}, val: []byte("bin")},
	}
	for _, mode := range []struct {
		name   string
		sorted bool
	}{
		{"Add", false},
		{"AddSorted", true},
	} {
		t.Run(mode.name, func(t *testing.T) {
			base := filepath.Join(t.TempDir(), "lib")
			buildSimpleRecords(t, base, CreateOptions{}, recs, mode.sorted)
			kv, err := OpenSimpleKV(base, 0)
			if err != nil {
				t.Fatal(err)
			}
			defer kv.Close()
			checkSimpleRecords(t, kv, recs)
			for _, miss := range [][]byte{{0x01}, []byte("AAAA"), []byte("B")} {
				if v, ok := getOK(t, kv, miss); ok {
					t.Fatalf("Get(%x) = %d bytes, want miss", miss, len(v))
				}
			}
			if v, ok := getOK(t, kv, nil); !ok || !bytes.Equal(v, []byte("empty")) {
				t.Fatalf("Get(nil) = %q, %v; want empty-key value", v, ok)
			}
		})
	}
}

func TestSimpleKVValueBoundaries(t *testing.T) {
	lens := []int{0, 1, 63, 64, 65, 32767, 32768, 32769}
	recs := make([]simpleRec, len(lens))
	for i, n := range lens {
		recs[i] = simpleRec{
			key: []byte(fmt.Sprintf("v%02d", i)),
			val: bytes.Repeat([]byte{byte('a' + i)}, n),
		}
	}
	base := filepath.Join(t.TempDir(), "lib")
	buildSimpleRecords(t, base, CreateOptions{}, recs, true)
	kv, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer kv.Close()
	checkSimpleRecords(t, kv, recs)
}

func TestSimpleKVPerValueFraming(t *testing.T) {
	// 小语料（64 条 320B）会让 zstd.BuildDict 返回 "0 literals found"；
	// 这里用 512 条 2048B（1MB）保证训练语料足够。
	const n = 512
	recs := make([]simpleRec, n)
	for i := range recs {
		val := make([]byte, 2048)
		prefix := fmt.Sprintf("gene-%04d:", i)
		copy(val, prefix)
		for j := len(prefix); j < len(val); j++ {
			val[j] = byte('A' + (i*7+j*3)%26)
		}
		recs[i] = simpleRec{
			key: []byte(fmt.Sprintf("d%03d", i)),
			val: val,
		}
	}
	base := filepath.Join(t.TempDir(), "lib")
	buildSimpleRecords(t, base, CreateOptions{PackTarget: -1}, recs, true)
	kv, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer kv.Close()
	checkSimpleRecords(t, kv, recs)
}

func TestSimpleKVRawThreshold(t *testing.T) {
	for _, tc := range []struct {
		name string
		opts CreateOptions
	}{
		{"threshold128", CreateOptions{RawValueBelow: 128}},
		{"allRaw", CreateOptions{RawValueBelow: 1 << 30}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			recs := []simpleRec{
				{key: []byte("a127"), val: bytes.Repeat([]byte("x"), 127)},
				{key: []byte("a128"), val: bytes.Repeat([]byte("x"), 128)},
				{key: []byte("a4096"), val: bytes.Repeat([]byte("y"), 4096)},
			}
			base := filepath.Join(t.TempDir(), "lib")
			buildSimpleRecords(t, base, tc.opts, recs, true)
			kv, err := OpenSimpleKV(base, 0)
			if err != nil {
				t.Fatal(err)
			}
			defer kv.Close()
			checkSimpleRecords(t, kv, recs)
		})
	}
}

func TestOpenSimpleKVBadLibrary(t *testing.T) {
	base := filepath.Join(t.TempDir(), "lib")
	recs := []simpleRec{{key: []byte("BRCA1"), val: []byte("v")}}
	buildSimpleRecords(t, base, CreateOptions{}, recs, true)

	f, err := os.OpenFile(base+".keys", os.O_WRONLY, 0)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteAt([]byte{0xFF}, 0); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenSimpleKV(base, 0); err == nil {
		t.Fatal("bad keys magic should error")
	}

	removeSimpleArtifacts(t, base)
	buildSimpleRecords(t, base, CreateOptions{}, recs, true)
	f, err = os.OpenFile(base+".vals", os.O_WRONLY, 0)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteAt([]byte{0xFF}, 0); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenSimpleKV(base, 0); err == nil {
		t.Fatal("bad vals magic should error")
	}

	if err := os.Remove(base + ".vals"); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenSimpleKV(base, 0); err == nil {
		t.Fatal("missing vals should error")
	}
}

func TestSimpleKVMaxKeyLen(t *testing.T) {
	key := bytes.Repeat([]byte("k"), maxKeyLen)
	val := []byte("max-key")
	base := filepath.Join(t.TempDir(), "lib")
	buildSimpleRecords(t, base, CreateOptions{}, []simpleRec{{key: key, val: val}}, true)
	kv, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer kv.Close()
	v, ok := getOK(t, kv, key)
	if !ok || !bytes.Equal(v, val) {
		t.Fatalf("Get(maxKeyLen key) = %d bytes, %v", len(v), ok)
	}
	miss := bytes.Clone(key)
	miss[len(miss)-1] = 'x'
	if _, ok := getOK(t, kv, miss); ok {
		t.Fatal("Get(maxKeyLen key + 1 change) should miss")
	}
}

func TestSimpleKVManyRecords(t *testing.T) {
	const n = 4096
	recs := make([]simpleRec, n)
	for i := range recs {
		recs[i] = simpleRec{
			key: []byte(fmt.Sprintf("k%05d", i)),
			val: []byte(fmt.Sprintf("v%d", i)),
		}
	}
	base := filepath.Join(t.TempDir(), "lib")
	buildSimpleRecords(t, base, CreateOptions{}, recs, true)
	kv, err := OpenSimpleKV(base, 1)
	if err != nil {
		t.Fatal(err)
	}
	defer kv.Close()
	for i := 0; i < n; i += 7 {
		v, ok := getOK(t, kv, recs[i].key)
		if !ok || !bytes.Equal(v, recs[i].val) {
			t.Fatalf("Get(%q) = %d bytes, %v", recs[i].key, len(v), ok)
		}
	}
	v, ok := getOK(t, kv, recs[n-1].key)
	if !ok || !bytes.Equal(v, recs[n-1].val) {
		t.Fatalf("Get(last) = %d bytes, %v", len(v), ok)
	}
}

func TestOpenSimpleKVHeaderValidation(t *testing.T) {
	base := filepath.Join(t.TempDir(), "lib")
	recs := []simpleRec{{key: []byte("BRCA1"), val: []byte("v")}}
	build := func() {
		removeSimpleArtifacts(t, base)
		buildSimpleRecords(t, base, CreateOptions{}, recs, true)
	}
	corrupt := func(file string, off int64, b []byte) {
		t.Helper()
		f, err := os.OpenFile(file, os.O_WRONLY, 0)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := f.WriteAt(b, off); err != nil {
			t.Fatal(err)
		}
		if err := f.Close(); err != nil {
			t.Fatal(err)
		}
	}
	cases := []struct {
		name string
		file string
		off  int64
		b    []byte
	}{
		{"keys version", base + ".keys", 8, []byte{2, 0, 0, 0}},
		{"keys recsPerBlock", base + ".keys", 20, []byte{32, 0}},
		{"keys index offset", base + ".keys", 28, []byte{65, 0, 0, 0, 0, 0, 0, 0}},
		{"vals version", base + ".vals", 8, []byte{2, 0, 0, 0}},
		{"vals count", base + ".vals", 12, []byte{2, 0, 0, 0, 0, 0, 0, 0}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			build()
			corrupt(tc.file, tc.off, tc.b)
			if _, err := OpenSimpleKV(base, 0); err == nil {
				t.Fatal("corrupt header should error")
			}
		})
	}
}

func TestSimpleKVCorruptValueFrame(t *testing.T) {
	base := filepath.Join(t.TempDir(), "lib")
	w, err := Create(base, CreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	key := []byte("BRCA1")
	if err := w.Add(key, bytes.Repeat([]byte("val"), 300)); err != nil {
		t.Fatal(err)
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}
	recs, _, _ := parseKeys(t, base+".keys")
	if len(recs) != 1 || !recs[0].packed {
		t.Fatalf("want one packed record, got %d", len(recs))
	}
	f, err := os.OpenFile(base+".vals", os.O_WRONLY, 0)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteAt([]byte{0xFF, 0xFF, 0xFF, 0xFF}, int64(recs[0].valOff)); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	kv, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer kv.Close()
	if v, ok, err := kv.Get(key); err == nil || ok {
		t.Fatalf("corrupt frame Get = %d bytes, ok=%v, err=%v; want error", len(v), ok, err)
	}
}

func TestSimpleKVRecordWrapping(t *testing.T) {
	base := filepath.Join(t.TempDir(), "lib")
	value := []byte("hello")
	w, err := Create(base, CreateOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if err := w.Add([]byte("k"), value); err != nil {
		t.Fatal(err)
	}
	if err := w.Finish(); err != nil {
		t.Fatal(err)
	}

	shard, err := openShard(base, nil)
	if err != nil {
		t.Fatal(err)
	}
	defer shard.Close()
	raw, ok, err := shard.get([]byte("k"))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("raw record missing")
	}
	schema := Schema{{Name: "value", Type: TypeString, Number: "1"}}
	record, err := NewRecord(schema, raw)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(record.Bytes("value")); got != string(value) {
		t.Fatalf("record value = %q, want %q", got, value)
	}

	kv, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer kv.Close()
	got, ok := getOK(t, kv, []byte("k"))
	if !ok || !bytes.Equal(got, value) {
		t.Fatalf("Get = %q, %v; want %q", got, ok, value)
	}
}
