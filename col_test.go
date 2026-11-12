package immkv

import (
	"bytes"
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"
)

const colTestHeader = "##fileformat=VCFv4.2\n" +
	"##INFO=<ID=AF,Number=A,Type=Float,Description=\"Allele frequency\">\n" +
	"##INFO=<ID=AC,Number=A,Type=Integer,Description=\"Allele count\">\n" +
	"##INFO=<ID=AN,Number=1,Type=Integer,Description=\"Allele number\">\n" +
	"##INFO=<ID=DB,Number=0,Type=Flag,Description=\"dbSNP membership\">\n" +
	"##INFO=<ID=AD,Number=R,Type=Integer,Description=\"Allelic depths\">\n" +
	"##INFO=<ID=AFR,Number=.,Type=Float,Description=\"Array frequency\">\n" +
	"##INFO=<ID=ANN,Number=.,Type=String,Description=\"Annotations\">\n" +
	"#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO\n"

// buildColTestLib 建一个覆盖多 bin、Number=A/R/.、Flag、String 与整条缺值的
// 小库，返回目录、schema 与按升序排列的键。
func buildColTestLib(t *testing.T) (string, Schema, [][]byte) {
	t.Helper()
	dir := t.TempDir()
	vcf := filepath.Join(dir, "in.vcf")
	writeIngestFile(t, vcf, colTestHeader+
		ingestRec("1", 100, "C", "A,G",
			"AF=0.1,0.2;AC=3,4;AN=100;DB;AD=10,5,6;AFR=0.01,.;ANN=a|b,c|d")+
		ingestRec("1", 1_500_000, "G", "T", "AN=200")+
		ingestRec("2", 50, "T", "C", "."))
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", Options{}); err != nil {
		t.Fatal(err)
	}
	schema := ingestSchema(t, out)
	keys := [][]byte{
		VarKey(0, 100, "C", "A"),
		VarKey(0, 100, "C", "G"),
		VarKey(0, 1_500_000, "G", "T"),
		VarKey(1, 50, "T", "C"),
	}
	return out, schema, keys
}

func TestBuildFromVCFsWritesColumnar(t *testing.T) {
	out, schema, keys := buildColTestLib(t)
	manifest, err := ReadManifest(filepath.Join(out, "chr1.manifest"))
	if err != nil {
		t.Fatal(err)
	}
	if manifest.Format != manifestFormat || manifest.Version != manifestVersion {
		t.Fatalf("manifest format/version = %q/%d", manifest.Format, manifest.Version)
	}
	if manifest.Base != "chr1" || manifest.BinWidth != colBinWidth {
		t.Fatalf("manifest base/bin_width = %q/%d", manifest.Base, manifest.BinWidth)
	}
	if manifest.RecordCount != 3 {
		t.Fatalf("manifest record_count = %d, want 3", manifest.RecordCount)
	}
	if len(manifest.Fields) != len(schema) {
		t.Fatalf("manifest fields %d ≠ schema %d", len(manifest.Fields), len(schema))
	}
	for i := range schema {
		mf := manifest.Fields[i]
		if mf.Column != i || mf.Name != schema[i].Name ||
			mf.Type != schema[i].Type.String() || mf.Number != schema[i].Number {
			t.Fatalf("manifest field %d = %+v, want %+v", i, mf, schema[i])
		}
	}
	raw, err := os.ReadFile(filepath.Join(out, "chr1.col"))
	if err != nil {
		t.Fatal(err)
	}
	if string(raw[:8]) != colMagic {
		t.Fatalf("col magic = %q, want %q", raw[:8], colMagic)
	}

	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	want := make([][]byte, len(keys))
	for i, key := range keys {
		value, ok := db.Get(key)
		if !ok {
			t.Fatalf("record-order miss for %x", key)
		}
		want[i] = bytes.Clone(value)
	}
	sdb, err := OpenStreaming(out, Options{Threads: 2})
	if err != nil {
		t.Fatal(err)
	}
	defer sdb.Close()
	sdb.forceCol = true
	got, err := sdb.GetBatch(keys)
	if err != nil {
		t.Fatal(err)
	}
	for i := range keys {
		if !bytes.Equal(got[i], want[i]) {
			rendered, rerr := RenderRecord(schema, got[i])
			t.Fatalf("key %x: columnar record %x (%s, err %v) ≠ record order %x",
				keys[i], got[i], rendered, rerr, want[i])
		}
	}
}

func TestColumnarSelectedColumns(t *testing.T) {
	out, schema, keys := buildColTestLib(t)
	sdb, err := OpenStreaming(out, Options{Threads: 2, Columns: []int{0, 3, 5}})
	if err != nil {
		t.Fatal(err)
	}
	defer sdb.Close()
	sdb.forceCol = true
	got, err := sdb.GetBatch(keys)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"AF=0.1;DB;AFR=0.01,.",
		"AF=0.2;DB;AFR=0.01,.",
		".",
		".",
	}
	for i := range keys {
		text, err := RenderRecord(schema, got[i])
		if err != nil {
			t.Fatalf("key %x: RenderRecord: %v", keys[i], err)
		}
		if string(text) != want[i] {
			t.Fatalf("key %x: rendered %q, want %q", keys[i], text, want[i])
		}
		if len(got[i]) < recordBitmapLen(len(schema)) {
			t.Fatalf("key %x: record 位图短：%d < %d", keys[i], len(got[i]), recordBitmapLen(len(schema)))
		}
	}
}

func TestColumnarPresentCountMismatch(t *testing.T) {
	out, _, keys := buildColTestLib(t)
	path := filepath.Join(out, "chr1.col")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	binCount := binary.LittleEndian.Uint32(raw[36:40])
	if binCount == 0 {
		t.Fatal("bin_count = 0")
	}
	binDirOff := binary.LittleEndian.Uint64(raw[48:56])
	dirOff := binary.LittleEndian.Uint64(raw[binDirOff+16:])
	// 第一列第一个非空 bin 的 present_count 改成大于 bin 记录数。
	entry := dirOff
	binary.LittleEndian.PutUint32(raw[entry+24:], 99)
	if err := os.WriteFile(path, raw, 0o644); err != nil {
		t.Fatal(err)
	}
	sdb, err := OpenStreaming(out, Options{Threads: 2})
	if err != nil {
		t.Fatal(err)
	}
	defer sdb.Close()
	sdb.forceCol = true
	if _, err := sdb.GetBatch(keys); err == nil {
		t.Fatal("present_count 不符应报错")
	}
}

func TestColumnarValueStreamTruncated(t *testing.T) {
	out, _, keys := buildColTestLib(t)
	path := filepath.Join(out, "chr1.col")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	binDirOff := binary.LittleEndian.Uint64(raw[48:56])
	dirOff := binary.LittleEndian.Uint64(raw[binDirOff+16:])
	// 把第一列值流长度砍半：帧本身仍能解压，但后续记录的值会越界。
	valueLen := binary.LittleEndian.Uint32(raw[dirOff+20:])
	binary.LittleEndian.PutUint32(raw[dirOff+20:], valueLen/2)
	if err := os.WriteFile(path, raw, 0o644); err != nil {
		t.Fatal(err)
	}
	sdb, err := OpenStreaming(out, Options{Threads: 2})
	if err != nil {
		t.Fatal(err)
	}
	defer sdb.Close()
	sdb.forceCol = true
	if _, err := sdb.GetBatch(keys); err == nil {
		t.Fatal("值流截断应报错")
	}
}

func TestColumnarManifestMismatch(t *testing.T) {
	out, _, _ := buildColTestLib(t)
	path := filepath.Join(out, "chr1.manifest")
	manifest, err := ReadManifest(path)
	if err != nil {
		t.Fatal(err)
	}
	manifest.ValsBytes++
	if err := WriteManifest(path, manifest); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenStreaming(out, Options{Threads: 1}); err == nil {
		t.Fatal("vals 字节数与 manifest 不符应报错")
	}
}

func TestColumnarOldMagicRejected(t *testing.T) {
	out, _, _ := buildColTestLib(t)
	path := filepath.Join(out, "chr1.col")
	raw, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	copy(raw[:8], "IMKVC001")
	if err := os.WriteFile(path, raw, 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenStreaming(out, Options{Threads: 1}); err == nil {
		t.Fatal("旧 col magic 应报错")
	}
}

func TestColumnarPairMissing(t *testing.T) {
	out, _, _ := buildColTestLib(t)
	if err := os.Remove(filepath.Join(out, "chr1.manifest")); err != nil {
		t.Fatal(err)
	}
	if _, err := OpenStreaming(out, Options{Threads: 1}); err == nil {
		t.Fatal("col 缺 manifest 应报错")
	}
}

func TestColumnarColumnsValidation(t *testing.T) {
	out, _, _ := buildColTestLib(t)
	if _, err := OpenStreaming(out, Options{Threads: 1, Columns: []int{0, 0}}); err == nil {
		t.Fatal("重复列号应报错")
	}
	if _, err := OpenStreaming(out, Options{Threads: 1, Columns: []int{-1}}); err == nil {
		t.Fatal("负列号应报错")
	}
	if _, err := OpenStreaming(out, Options{Threads: 1, Columns: []int{99}}); err == nil {
		t.Fatal("越界列号应报错")
	}
}

func TestBuildFromTSVsWritesColumnar(t *testing.T) {
	dir := t.TempDir()
	tsv := filepath.Join(dir, "in.tsv")
	writeIngestFile(t, tsv, "b\tB1\nc\tC1\tC2\nd\tD1\n")
	base := filepath.Join(dir, "genes")
	if err := BuildFromTSVs(base, []string{tsv}, Options{}); err != nil {
		t.Fatal(err)
	}
	manifest, err := ReadManifest(base + ".manifest")
	if err != nil {
		t.Fatal(err)
	}
	if manifest.BinWidth != 0 || manifest.RecordCount != 3 {
		t.Fatalf("manifest bin_width/record_count = %d/%d", manifest.BinWidth, manifest.RecordCount)
	}
	if len(manifest.Fields) != 1 || manifest.Fields[0].Name != "value" ||
		manifest.Fields[0].Type != TypeString.String() {
		t.Fatalf("manifest fields = %+v", manifest.Fields)
	}
	raw, err := os.ReadFile(base + ".col")
	if err != nil {
		t.Fatal(err)
	}
	if string(raw[:8]) != colMagic {
		t.Fatalf("col magic = %q", raw[:8])
	}
	if binCount := binary.LittleEndian.Uint32(raw[36:40]); binCount != 1 {
		t.Fatalf("bin_count = %d, want 1", binCount)
	}
	db, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if v, ok := db.Get([]byte("c")); !ok || string(v) != "C1\tC2" {
		t.Fatalf("Get(c) = %q, %v", v, ok)
	}
}
