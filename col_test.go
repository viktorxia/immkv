package immkv

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const colTestHeader = "##fileformat=VCFv4.2\n" +
	"##INFO=<ID=AF,Number=A,Type=Float,Description=\"AF\">\n" +
	"##INFO=<ID=AC,Number=A,Type=Integer,Description=\"AC\">\n" +
	"##INFO=<ID=AN,Number=1,Type=Integer,Description=\"AN\">\n" +
	"##INFO=<ID=DB,Number=0,Type=Flag,Description=\"DB\">\n" +
	"##INFO=<ID=STR,Number=1,Type=String,Description=\"STR\">\n" +
	ingestTestSiteHeader

// buildColTestLib 建一个覆盖 Number=A、Flag、缺失字段与 "." 的小库。
func buildColTestLib(t *testing.T) (string, [][]byte) {
	t.Helper()
	dir := t.TempDir()
	vcf := filepath.Join(dir, "in.vcf")
	writeIngestFile(t, vcf, colTestHeader+
		ingestRec("1", 100, "C", "A,G", "STR=foo;DB;AN=100;AC=3,4;AF=0.1,0.2")+
		ingestRec("1", 100, "C", "T", "STR=bar;AN=100;AF=0.5;AC=10")+
		ingestRec("1", 200, "A", "C", "AC=1")+
		ingestRec("1", 300, "A", "G", ".")+
		ingestRec("1", 2500000, "A", "T", "AC=2;AN=10")+ // 前两个 bin 为空
		ingestRec("2", 50, "G", "A", "STR=baz;DB"))
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", Options{}); err != nil {
		t.Fatal(err)
	}
	keys := [][]byte{
		VarKey(0, 100, "C", "A"), VarKey(0, 100, "C", "G"), VarKey(0, 100, "C", "T"),
		VarKey(0, 200, "A", "C"), VarKey(0, 300, "A", "G"),
		VarKey(0, 2500000, "A", "T"), VarKey(1, 50, "G", "A"),
	}
	return out, keys
}

func openStreamCol(t *testing.T, dir string, threads int, sel []int) *StreamDB {
	t.Helper()
	sdb, err := OpenStreaming(dir, Options{Threads: threads, Columns: sel})
	if err != nil {
		t.Fatal(err)
	}
	sdb.forceCol = true
	t.Cleanup(func() { sdb.Close() })
	return sdb
}

func TestColumnarMatchesRecordOrder(t *testing.T) {
	dir, keys := buildColTestLib(t)
	db, err := OpenEager(dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	want := make([][]byte, len(keys))
	for i, k := range keys {
		v, ok := db.Get(k)
		if !ok {
			t.Fatalf("记录序未命中 %x", k)
		}
		want[i] = bytes.Clone(v)
	}
	sdb := openStreamCol(t, dir, 4, nil)
	got, err := sdb.GetBatch(keys)
	if err != nil {
		t.Fatal(err)
	}
	for i := range keys {
		if !bytes.Equal(got[i], want[i]) {
			t.Fatalf("键 %x：列存 %q ≠ 记录序 %q", keys[i], got[i], want[i])
		}
	}
	// 混入未命中键（追加 0xFF 保持升序）。
	var mixed [][]byte
	for _, k := range keys {
		mixed = append(mixed, k, append(bytes.Clone(k), 0xFF))
	}
	out, err := sdb.GetBatch(mixed)
	if err != nil {
		t.Fatal(err)
	}
	for i, k := range mixed {
		if k[len(k)-1] == 0xFF {
			if out[i] != nil {
				t.Fatalf("拼造键 %x 不应命中", k)
			}
			continue
		}
	}
	// 重复输入键。
	dup := [][]byte{keys[0], bytes.Clone(keys[0]), keys[0]}
	out, err = sdb.GetBatch(dup)
	if err != nil {
		t.Fatal(err)
	}
	for i := range dup {
		if !bytes.Equal(out[i], want[0]) {
			t.Fatalf("重复键第 %d 项 %q ≠ %q", i, out[i], want[0])
		}
	}
}

func TestColumnarSelectedColumns(t *testing.T) {
	dir, keys := buildColTestLib(t)
	// schema 序：AF=0, AC=1, AN=2, DB=3, STR=4。输入顺序故意倒置。
	cases := []struct {
		sel  []int
		want []string
	}{
		{[]int{4, 1}, []string{"AC=3,4;STR=foo", "AC=3,4;STR=foo", "AC=10;STR=bar", "AC=1", ".", "AC=2", "STR=baz"}},
		{[]int{3}, []string{"DB", "DB", ".", ".", ".", ".", "DB"}},
		{[]int{0}, []string{"AF=0.1,0.2", "AF=0.1,0.2", "AF=0.5", ".", ".", ".", "."}},
	}
	for ci, tc := range cases {
		sdb := openStreamCol(t, dir, 3, tc.sel)
		out, err := sdb.GetBatch(keys)
		if err != nil {
			t.Fatalf("case %d: %v", ci, err)
		}
		for i := range keys {
			if string(out[i]) != tc.want[i] {
				t.Fatalf("case %d 键 %x：%q ≠ %q", ci, keys[i], out[i], tc.want[i])
			}
		}
	}
}

func TestColumnarSelectedFieldsBuild(t *testing.T) {
	dir := t.TempDir()
	vcf := filepath.Join(dir, "in.vcf")
	writeIngestFile(t, vcf, colTestHeader+
		ingestRec("1", 100, "C", "A,G", "DB;AN=100;AC=3,4;AF=0.1,0.2;STR=foo")+
		ingestRec("1", 200, "A", "C", "AC=1")+
		ingestRec("1", 300, "A", "G", "."))
	out := filepath.Join(dir, "lib")
	fields := []string{"DB", "AN", "AC", "AF"} // 选择顺序故意倒置
	if err := BuildFromVCFs(out, []string{vcf}, fields, "X", Options{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	keys := [][]byte{
		VarKey(0, 100, "C", "A"), VarKey(0, 100, "C", "G"),
		VarKey(0, 200, "A", "C"), VarKey(0, 300, "A", "G"),
	}
	want := []string{
		"X_AF=0.1;X_AC=3;X_AN=100;X_DB",
		"X_AF=0.2;X_AC=4;X_AN=100;X_DB",
		"X_AC=1",
		".",
	}
	for i, k := range keys {
		v, ok := db.Get(k)
		if !ok || string(v) != want[i] {
			t.Fatalf("记录序 键 %x = %q, %v; want %q", k, v, ok, want[i])
		}
	}
	sdb := openStreamCol(t, out, 3, nil)
	got, err := sdb.GetBatch(keys)
	if err != nil {
		t.Fatal(err)
	}
	for i := range keys {
		if string(got[i]) != want[i] {
			t.Fatalf("列存 键 %x = %q; want %q", keys[i], got[i], want[i])
		}
	}
	sub := openStreamCol(t, out, 3, []int{1}) // AC（schema 序 AF,AC,AN,DB）
	got, err = sub.GetBatch(keys)
	if err != nil {
		t.Fatal(err)
	}
	for i, wantSub := range []string{"X_AC=3", "X_AC=4", "X_AC=1", "."} {
		if string(got[i]) != wantSub {
			t.Fatalf("子集 键 %x = %q; want %q", keys[i], got[i], wantSub)
		}
	}
}

func TestColumnarManifestSchema(t *testing.T) {
	dir, _ := buildColTestLib(t)
	raw, err := os.ReadFile(filepath.Join(dir, "chr1.manifest"))
	if err != nil {
		t.Fatal(err)
	}
	var m colManifest
	if err := json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	if m.Format != colManifestFormat || m.Base != "chr1" || m.RecordCount != 6 {
		t.Fatalf("manifest: %+v", m)
	}
	wantNames := []string{"AF", "AC", "AN", "DB", "STR"}
	if len(m.Fields) != len(wantNames) {
		t.Fatalf("fields = %d", len(m.Fields))
	}
	for i, name := range wantNames {
		if m.Fields[i].Name != name || m.Fields[i].Column != i {
			t.Fatalf("field %d = %+v", i, m.Fields[i])
		}
	}
	if m.Fields[3].Encoding != "bitmap" || m.Fields[0].Encoding != "raw" {
		t.Fatalf("encoding: %+v", m.Fields)
	}
	// 源顺序故意倒置：三条单字段/无字段记录算 OK，两条多字段记录算 Bad。
	if m.SourceInfoOrder.OK != 3 || m.SourceInfoOrder.Bad != 2 || m.SourceInfoOrder.Inversions != 6 {
		t.Fatalf("source_info_order = %+v", m.SourceInfoOrder)
	}
	hdr, err := os.ReadFile(filepath.Join(dir, "vcf.immkv.header"))
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range wantNames {
		if !strings.Contains(string(hdr), "ID="+name+",") {
			t.Fatalf("immkv header 缺 %s", name)
		}
	}
}

func TestColumnarCorrupt(t *testing.T) {
	t.Run("magic", func(t *testing.T) {
		dir, _ := buildColTestLib(t)
		p := filepath.Join(dir, "chr1.col")
		f, err := os.OpenFile(p, os.O_WRONLY, 0)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := f.WriteAt([]byte("X"), 0); err != nil {
			t.Fatal(err)
		}
		f.Close()
		if _, err := OpenStreaming(dir, Options{}); err == nil {
			t.Fatal("坏 magic 应报错")
		}
	})
	t.Run("truncate", func(t *testing.T) {
		dir, _ := buildColTestLib(t)
		p := filepath.Join(dir, "chr1.col")
		st, err := os.Stat(p)
		if err != nil {
			t.Fatal(err)
		}
		if err := os.Truncate(p, st.Size()-1); err != nil {
			t.Fatal(err)
		}
		if _, err := OpenStreaming(dir, Options{}); err == nil {
			t.Fatal("截断应报错")
		}
	})
	t.Run("manifest-missing", func(t *testing.T) {
		dir, _ := buildColTestLib(t)
		if err := os.Remove(filepath.Join(dir, "chr1.manifest")); err != nil {
			t.Fatal(err)
		}
		if _, err := OpenStreaming(dir, Options{}); err == nil {
			t.Fatal("缺 manifest 应报错")
		}
	})
	t.Run("build-id", func(t *testing.T) {
		dir, _ := buildColTestLib(t)
		p := filepath.Join(dir, "chr1.manifest")
		raw, err := os.ReadFile(p)
		if err != nil {
			t.Fatal(err)
		}
		var m colManifest
		if err := json.Unmarshal(raw, &m); err != nil {
			t.Fatal(err)
		}
		m.BuildID = strings.Repeat("0", 32)
		b, _ := json.Marshal(m)
		if err := os.WriteFile(p, b, 0o644); err != nil {
			t.Fatal(err)
		}
		if _, err := OpenStreaming(dir, Options{}); err == nil {
			t.Fatal("build id 不一致应报错")
		}
	})
	t.Run("value-frame", func(t *testing.T) {
		dir, keys := buildColTestLib(t)
		p := filepath.Join(dir, "chr1.col")
		off := firstColValueOff(t, p)
		f, err := os.OpenFile(p, os.O_WRONLY, 0)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := f.WriteAt([]byte{0xFF, 0xFF, 0xFF, 0xFF}, off); err != nil {
			t.Fatal(err)
		}
		f.Close()
		sdb := openStreamCol(t, dir, 2, nil)
		if _, err := sdb.GetBatch(keys); err == nil {
			t.Fatal("坏值流应报错")
		}
	})
}

// firstColValueOff 找第一个非空 bin 的第一列值流偏移（测试用）。
func firstColValueOff(t *testing.T, path string) int64 {
	t.Helper()
	b, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	binCount := binary.LittleEndian.Uint32(b[36:40])
	colCount := binary.LittleEndian.Uint32(b[40:44])
	binDir := binary.LittleEndian.Uint64(b[48:56])
	for i := 0; i < int(binCount); i++ {
		e := b[binDir+uint64(i*colBinEntrySize):]
		count := binary.LittleEndian.Uint32(e[8:12])
		if count == 0 {
			continue
		}
		dirOff := binary.LittleEndian.Uint64(e[16:24])
		for c := 0; c < int(colCount); c++ {
			ce := b[dirOff+uint64(c*colColumnEntrySize):]
			if v := binary.LittleEndian.Uint64(ce[8:16]); v != 0 {
				return int64(v)
			}
		}
	}
	t.Fatal("找不到值流")
	return 0
}

// TestColumnarLeadingEmptyBins：首个 bin 之前全空时，bin 目录下标仍须等于
// 基因组 bin 号。
func TestColumnarLeadingEmptyBins(t *testing.T) {
	dir := t.TempDir()
	vcf := filepath.Join(dir, "in.vcf")
	writeIngestFile(t, vcf, colTestHeader+
		ingestRec("1", 2500000, "A", "T", "AC=2;AN=10")+
		ingestRec("1", 5200000, "C", "G", "AF=0.3"))
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", Options{}); err != nil {
		t.Fatal(err)
	}
	keys := [][]byte{VarKey(0, 2500000, "A", "T"), VarKey(0, 5200000, "C", "G")}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	sdb := openStreamCol(t, out, 2, nil)
	got, err := sdb.GetBatch(keys)
	if err != nil {
		t.Fatal(err)
	}
	for i, k := range keys {
		want, ok := db.Get(k)
		if !ok || !bytes.Equal(got[i], want) {
			t.Fatalf("键 %x：列存 %q，记录序 %q (hit %v)", k, got[i], want, ok)
		}
	}
}

func TestColumnarColumnsRequireBaseCol(t *testing.T) {
	dir := t.TempDir()
	buildTestLib(t, dir, Options{}, false)
	if _, err := OpenStreaming(dir, Options{Columns: []int{0}}); err == nil {
		t.Fatal("无 base.col 时 Columns 应报错")
	}
	if _, err := OpenStreaming(dir, Options{}); err != nil {
		t.Fatalf("无 base.col 的旧库应仍可按记录序打开：%v", err)
	}
}
