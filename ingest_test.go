package immkv

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/biogo/hts/bgzf"
)

const ingestTestHeader = "##fileformat=VCFv4.2\n" +
	"##INFO=<ID=AF,Number=A,Type=Float,Description=\"Allele frequency\">\n" +
	"##INFO=<ID=AC,Number=A,Type=Integer,Description=\"Allele count\">\n" +
	"##INFO=<ID=AN,Number=1,Type=Integer,Description=\"Allele number\">\n" +
	"##INFO=<ID=DB,Number=0,Type=Flag,Description=\"dbSNP membership\">\n" +
	"##FILTER=<ID=PASS,Description=\"All filters passed\">\n" +
	"#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO\n"

const ingestTestSiteHeader = "#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO\n"

func ingestRec(chrom string, pos int, ref, alt, info string) string {
	return fmt.Sprintf("%s\t%d\t.\t%s\t%s\t.\tPASS\t%s\n", chrom, pos, ref, alt, info)
}

func writeIngestFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

func writeIngestBGZF(t *testing.T, path, content string) {
	t.Helper()
	var buf bytes.Buffer
	bw := bgzf.NewWriter(&buf, 1)
	if _, err := bw.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := bw.Close(); err != nil {
		t.Fatal(err)
	}
	writeIngestFile(t, path, buf.String())
}

func mustIngestGet(t *testing.T, db *DB, key []byte) string {
	t.Helper()
	v, ok := db.Get(key)
	if !ok {
		t.Fatalf("key %x: not found", key)
	}
	return string(v)
}

func mustIngestMiss(t *testing.T, db *DB, key []byte) {
	t.Helper()
	if v, ok := db.Get(key); ok {
		t.Fatalf("key %x: unexpected hit %q", key, v)
	}
}

func countShardKeys(t *testing.T, dir string) int {
	t.Helper()
	ents, err := os.ReadDir(dir)
	if err != nil {
		t.Fatal(err)
	}
	n := 0
	for _, e := range ents {
		if strings.HasPrefix(e.Name(), "chr") && strings.HasSuffix(e.Name(), ".keys") {
			n++
		}
	}
	return n
}

func TestBuildFromTSVs(t *testing.T) {
	dir := t.TempDir()
	p1 := filepath.Join(dir, "a.tsv")
	p2 := filepath.Join(dir, "b.tsv")
	writeIngestFile(t, p1, "b\tB1\nb\tB1\nc\tC1\tC2\n\n")
	writeIngestFile(t, p2, "d\tD1\n")
	base := filepath.Join(dir, "sub", "genes")
	if err := BuildFromTSVs(base, []string{p1, p2}, Options{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	for _, tc := range []struct{ key, want string }{
		{"b", "B1"},
		{"c", "C1\tC2"},
		{"d", "D1"},
	} {
		v, ok := db.Get([]byte(tc.key))
		if !ok || string(v) != tc.want {
			t.Fatalf("Get(%q) = %q, %v; want %q", tc.key, v, ok, tc.want)
		}
	}
	if _, ok := db.Get([]byte("zzz")); ok {
		t.Fatal("unexpected hit for zzz")
	}
}

func TestBuildFromTSVsUnsorted(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "x.tsv")
	writeIngestFile(t, p, "c\tC\nb\tB\n")
	base := filepath.Join(dir, "genes")
	if err := BuildFromTSVs(base, []string{p}, Options{}); err == nil {
		t.Fatal("expected unsorted error")
	}
	for _, suffix := range []string{".keys", ".vals", ".keys.tmp", ".vals.tmp"} {
		if _, err := os.Stat(base + suffix); !os.IsNotExist(err) {
			t.Fatalf("%s should not exist after failure (stat err %v)", base+suffix, err)
		}
	}
}

func TestBuildFromVCFsAllFieldsSorted(t *testing.T) {
	dir := t.TempDir()
	vcf := filepath.Join(dir, "in.vcf")
	writeIngestFile(t, vcf, ingestTestHeader+
		ingestRec("1", 100, "C", "A,G", "DB;AN=100;AC=3,4;AF=0.1,0.2")+
		ingestRec("1", 100, "C", "T", "AN=100;AF=0.5;AC=10")+
		ingestRec("2", 50, "G", "A", "AF=0.3;AC=7;AN=100"))
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", Options{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	for _, tc := range []struct {
		key  []byte
		want string
	}{
		{VarKey(0, 100, "C", "A"), "AF=0.1,0.2;AC=3,4;AN=100;DB"},
		{VarKey(0, 100, "C", "G"), "AF=0.1,0.2;AC=3,4;AN=100;DB"},
		{VarKey(0, 100, "C", "T"), "AF=0.5;AC=10;AN=100"},
		{VarKey(1, 50, "G", "A"), "AF=0.3;AC=7;AN=100"},
	} {
		if got := mustIngestGet(t, db, tc.key); got != tc.want {
			t.Fatalf("Get(%x) = %q; want %q", tc.key, got, tc.want)
		}
	}
	raw, err := os.ReadFile(filepath.Join(out, "vcf.raw.header"))
	if err != nil {
		t.Fatal(err)
	}
	if string(raw) != ingestTestHeader {
		t.Fatalf("raw header mismatch:\n%s", raw)
	}
	wantHeader := "##fileformat=VCFv4.2\n" +
		"##INFO=<ID=AF,Number=A,Type=Float,Description=\"Allele frequency\">\n" +
		"##INFO=<ID=AC,Number=A,Type=Integer,Description=\"Allele count\">\n" +
		"##INFO=<ID=AN,Number=1,Type=Integer,Description=\"Allele number\">\n" +
		"##INFO=<ID=DB,Number=0,Type=Flag,Description=\"dbSNP membership\">\n" +
		ingestTestSiteHeader
	hdr, err := os.ReadFile(filepath.Join(out, "vcf.immkv.header"))
	if err != nil {
		t.Fatal(err)
	}
	if string(hdr) != wantHeader {
		t.Fatalf("immkv header mismatch:\n%s", hdr)
	}
}

func TestBuildFromVCFsSelectedFields(t *testing.T) {
	dir := t.TempDir()
	vcf := filepath.Join(dir, "in.vcf")
	writeIngestFile(t, vcf, ingestTestHeader+
		ingestRec("1", 100, "C", "A,G", "AF=0.1,0.2;AC=3,4;AN=100;DB")+
		ingestRec("1", 100, "C", "T", "AF=0.5;AC=10;AN=100")+
		ingestRec("1", 200, "C", "<DEL>,A,*", "AF=0.9,0.4,0.0;AC=1,2,0;AN=100")+
		ingestRec("1", 300, ".", "A", "AF=0.7;AC=5;AN=100")+
		ingestRec("chrUn_gl000220", 10, "A", "C", "AF=0.8;AC=6;AN=100")+
		ingestRec("1", 400, "C", "A", ".")+
		ingestRec("2", 50, "G", "A", "AF=0.3;AC=7;AN=100"))
	out := filepath.Join(dir, "lib")
	fields := []string{"DB", "AN", "AC", "AF"}
	if err := BuildFromVCFs(out, []string{vcf}, fields, "IMMKV", Options{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	for _, tc := range []struct {
		key  []byte
		want string
	}{
		{VarKey(0, 100, "C", "A"), "IMMKV_AF=0.1;IMMKV_AC=3;IMMKV_AN=100;IMMKV_DB"},
		{VarKey(0, 100, "C", "G"), "IMMKV_AF=0.2;IMMKV_AC=4;IMMKV_AN=100;IMMKV_DB"},
		{VarKey(0, 100, "C", "T"), "IMMKV_AF=0.5;IMMKV_AC=10;IMMKV_AN=100"},
		{VarKey(0, 200, "C", "A"), "IMMKV_AF=0.4;IMMKV_AC=2;IMMKV_AN=100"},
		{VarKey(0, 400, "C", "A"), "."},
		{VarKey(1, 50, "G", "A"), "IMMKV_AF=0.3;IMMKV_AC=7;IMMKV_AN=100"},
	} {
		if got := mustIngestGet(t, db, tc.key); got != tc.want {
			t.Fatalf("Get(%x) = %q; want %q", tc.key, got, tc.want)
		}
	}
	mustIngestMiss(t, db, VarKey(0, 300, ".", "A"))
	mustIngestMiss(t, db, VarKey(0, 200, "C", "<DEL>"))
	mustIngestMiss(t, db, VarKey(0, 200, "C", "*"))
	if n := countShardKeys(t, out); n != 2 {
		t.Fatalf("shard count = %d; want 2", n)
	}
	wantHeader := "##fileformat=VCFv4.2\n" +
		"##INFO=<ID=IMMKV_AF,Number=A,Type=Float,Description=\"Allele frequency\">\n" +
		"##INFO=<ID=IMMKV_AC,Number=A,Type=Integer,Description=\"Allele count\">\n" +
		"##INFO=<ID=IMMKV_AN,Number=1,Type=Integer,Description=\"Allele number\">\n" +
		"##INFO=<ID=IMMKV_DB,Number=0,Type=Flag,Description=\"dbSNP membership\">\n" +
		ingestTestSiteHeader
	hdr, err := os.ReadFile(filepath.Join(out, "vcf.immkv.header"))
	if err != nil {
		t.Fatal(err)
	}
	if string(hdr) != wantHeader {
		t.Fatalf("immkv header mismatch:\n%s", hdr)
	}
}

func TestBuildFromVCFsAllFieldsPrefix(t *testing.T) {
	dir := t.TempDir()
	vcf := filepath.Join(dir, "in.vcf")
	writeIngestFile(t, vcf, ingestTestHeader+
		ingestRec("1", 10, "A", "C", "DB;AN=100;AF=0.1,0.2")+
		ingestRec("1", 20, "A", "G", "."))
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, nil, "IMMKV", Options{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if got := mustIngestGet(t, db, VarKey(0, 10, "A", "C")); got != "IMMKV_AF=0.1,0.2;IMMKV_AN=100;IMMKV_DB" {
		t.Fatalf("Get(A>C) = %q", got)
	}
	if got := mustIngestGet(t, db, VarKey(0, 20, "A", "G")); got != "." {
		t.Fatalf("Get(A>G) = %q; want %q", got, ".")
	}
}

func TestBuildFromVCFsMultiFile(t *testing.T) {
	dir := t.TempDir()
	f1 := filepath.Join(dir, "1.vcf")
	f2 := filepath.Join(dir, "2.vcf")
	writeIngestFile(t, f1, ingestTestHeader+ingestRec("1", 100, "C", "A", "AF=0.1;AC=1;AN=100"))
	writeIngestFile(t, f2, ingestTestHeader+ingestRec("2", 50, "G", "A", "AF=0.2;AC=2;AN=100"))
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{f1, f2}, []string{"AF"}, "IMMKV", Options{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if got := mustIngestGet(t, db, VarKey(0, 100, "C", "A")); got != "IMMKV_AF=0.1" {
		t.Fatalf("chr1 value = %q", got)
	}
	if got := mustIngestGet(t, db, VarKey(1, 50, "G", "A")); got != "IMMKV_AF=0.2" {
		t.Fatalf("chr2 value = %q", got)
	}
}

func TestBuildFromVCFsMultiFileErrors(t *testing.T) {
	dir := t.TempDir()
	f1 := filepath.Join(dir, "1a.vcf")
	f2 := filepath.Join(dir, "1b.vcf")
	writeIngestFile(t, f1, ingestTestHeader+ingestRec("1", 100, "C", "A", "AF=0.1"))
	writeIngestFile(t, f2, ingestTestHeader+ingestRec("1", 200, "C", "T", "AF=0.2"))
	if err := BuildFromVCFs(filepath.Join(dir, "dup"), []string{f1, f2}, nil, "", Options{}); err == nil ||
		!strings.Contains(err.Error(), "more than one input") {
		t.Fatalf("duplicate chromosome error = %v", err)
	}

	f3 := filepath.Join(dir, "mixed.vcf")
	f4 := filepath.Join(dir, "3.vcf")
	writeIngestFile(t, f3, ingestTestHeader+
		ingestRec("1", 100, "C", "A", "AF=0.1")+
		ingestRec("2", 50, "G", "A", "AF=0.2"))
	writeIngestFile(t, f4, ingestTestHeader+ingestRec("3", 10, "A", "C", "AF=0.3"))
	if err := BuildFromVCFs(filepath.Join(dir, "mixed"), []string{f3, f4}, nil, "", Options{}); err == nil {
		t.Fatal("expected multi-chromosome file error")
	}

	hdr2 := "##fileformat=VCFv4.2\n" +
		"##INFO=<ID=AF,Number=1,Type=Float,Description=\"Allele frequency\">\n" +
		ingestTestSiteHeader
	f5 := filepath.Join(dir, "5.vcf")
	f6 := filepath.Join(dir, "6.vcf")
	writeIngestFile(t, f5, hdr2+ingestRec("1", 100, "C", "A", "AF=0.1"))
	writeIngestFile(t, f6, ingestTestHeader+ingestRec("2", 50, "G", "A", "AF=0.2"))
	if err := BuildFromVCFs(filepath.Join(dir, "hdr"), []string{f5, f6}, nil, "", Options{}); err == nil {
		t.Fatal("expected header mismatch error")
	}

	hdrOrder := "##fileformat=VCFv4.2\n" +
		"##INFO=<ID=DB,Number=0,Type=Flag,Description=\"dbSNP membership\">\n" +
		"##INFO=<ID=AN,Number=1,Type=Integer,Description=\"Allele number\">\n" +
		"##INFO=<ID=AC,Number=A,Type=Integer,Description=\"Allele count\">\n" +
		"##INFO=<ID=AF,Number=A,Type=Float,Description=\"Allele frequency\">\n" +
		ingestTestSiteHeader
	f7 := filepath.Join(dir, "7.vcf")
	f8 := filepath.Join(dir, "8.vcf")
	writeIngestFile(t, f7, hdrOrder+ingestRec("1", 100, "C", "A", "AF=0.1"))
	writeIngestFile(t, f8, ingestTestHeader+ingestRec("2", 50, "G", "A", "AF=0.2"))
	if err := BuildFromVCFs(filepath.Join(dir, "order"), []string{f7, f8}, nil, "", Options{}); err == nil ||
		!strings.Contains(err.Error(), "order") {
		t.Fatalf("INFO order mismatch error = %v", err)
	}
}

func TestBuildFromVCFsErrors(t *testing.T) {
	dir := t.TempDir()
	okRec := ingestRec("1", 100, "C", "A", "AF=0.1")
	cases := []struct {
		name   string
		body   string
		fields []string
		prefix string
		want   string
	}{
		{
			"pos-desc",
			ingestRec("1", 200, "C", "A", "AF=0.1") + ingestRec("1", 100, "C", "T", "AF=0.2"),
			nil, "", "out of order",
		},
		{
			"chrom-block",
			ingestRec("1", 100, "C", "A", "AF=0.1") +
				ingestRec("2", 50, "G", "A", "AF=0.2") +
				ingestRec("1", 200, "C", "T", "AF=0.3"),
			nil, "", "not contiguous",
		},
		{"unknown-field", okRec, []string{"NOPE"}, "", "not declared"},
		{"dup-field", okRec, []string{"AF", "AF"}, "", "duplicate"},
		{"record-unknown", ingestRec("1", 100, "C", "A", "AF=0.1;NOPE=1"), nil, "", "unknown INFO field"},
		{"selected-unknown", ingestRec("1", 100, "C", "A", "AF=0.1;NOPE=1"), []string{"AF"}, "", "unknown INFO field"},
		{"record-duplicate", ingestRec("1", 100, "C", "A", "AF=0.1;AF=0.2"), nil, "", "duplicate INFO field"},
		{"record-empty", ingestRec("1", 100, "C", "A", "AF=0.1;;AN=100"), nil, "", "empty INFO segment"},
		{"record-empty-name", ingestRec("1", 100, "C", "A", "=1"), nil, "", "empty INFO field name"},
		{"flag-value", ingestRec("1", 100, "C", "A", "DB=1"), nil, "", "Flag INFO field"},
		{"bare-nonflag", ingestRec("1", 100, "C", "A", "AF"), nil, "", "non-Flag INFO field"},
		{"bad-prefix", okRec, nil, "A=B", "invalid prefix"},
		{"no-records", "", nil, "", "no buildable records"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			vcf := filepath.Join(dir, tc.name+".vcf")
			writeIngestFile(t, vcf, ingestTestHeader+tc.body)
			err := BuildFromVCFs(filepath.Join(dir, tc.name+"-lib"), []string{vcf}, tc.fields, tc.prefix, Options{})
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("error = %v; want substring %q", err, tc.want)
			}
		})
	}

	vcf := filepath.Join(dir, "ok.vcf")
	writeIngestFile(t, vcf, ingestTestHeader+okRec)
	out := filepath.Join(dir, "existing")
	if err := os.MkdirAll(out, 0o755); err != nil {
		t.Fatal(err)
	}
	writeIngestFile(t, filepath.Join(out, "chr1.keys"), "")
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", Options{}); err == nil ||
		!strings.Contains(err.Error(), "already exists") {
		t.Fatalf("existing outdir error = %v", err)
	}
	if err := BuildFromVCFs(dir, nil, nil, "", Options{}); err == nil {
		t.Fatal("expected empty vcfs error")
	}
}

func TestBuildFromVCFsBGZF(t *testing.T) {
	dir := t.TempDir()
	vcf := filepath.Join(dir, "in.vcf.gz")
	writeIngestBGZF(t, vcf, ingestTestHeader+ingestRec("1", 10, "A", "C", "AF=0.1;AC=1;AN=100"))
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, []string{"AF"}, "IMMKV", Options{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if got := mustIngestGet(t, db, VarKey(0, 10, "A", "C")); got != "IMMKV_AF=0.1" {
		t.Fatalf("Get = %q", got)
	}
	raw, err := os.ReadFile(filepath.Join(out, "vcf.raw.header"))
	if err != nil {
		t.Fatal(err)
	}
	if string(raw) != ingestTestHeader {
		t.Fatalf("raw header mismatch:\n%s", raw)
	}
}

func TestBuildFromVCFsSampleColumns(t *testing.T) {
	dir := t.TempDir()
	header := "##fileformat=VCFv4.2\n" +
		"##INFO=<ID=AF,Number=A,Type=Float,Description=\"AF\">\n" +
		"##INFO=<ID=AN,Number=1,Type=Integer,Description=\"AN\">\n" +
		"#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO\tFORMAT\tS1\n"
	vcf := filepath.Join(dir, "in.vcf")
	writeIngestFile(t, vcf, header+
		"1\t10\t.\tA\tC\t.\tPASS\tAF=0.1;AN=100\tGT\t0/1\n")
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", Options{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if got := mustIngestGet(t, db, VarKey(0, 10, "A", "C")); got != "AF=0.1;AN=100" {
		t.Fatalf("Get = %q; sample columns must not enter the value", got)
	}
}
