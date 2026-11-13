package immkv

import (
	"bufio"
	"bytes"
	"encoding/binary"
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

func mustIngestGet(t *testing.T, db *DB, key []byte, schema Schema) string {
	t.Helper()
	value, ok, err := db.Get(key)
	if err != nil {
		t.Fatalf("key %x: %v", key, err)
	}
	if !ok {
		t.Fatalf("key %x: not found", key)
	}
	text, err := RenderRecord(schema, value)
	if err != nil {
		t.Fatalf("key %x: RenderRecord: %v", key, err)
	}
	return string(text)
}

// ingestSchema 从 vcf.immkv.header 还原 v003 schema，模拟调用方按 manifest
// 取 schema 的路径。
func ingestSchema(t *testing.T, dir string) Schema {
	t.Helper()
	file, err := os.Open(filepath.Join(dir, vcfImmkvHeaderName))
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	header, _, err := readVCFHeader(bufio.NewReader(file), false)
	if err != nil {
		t.Fatal(err)
	}
	schema := make(Schema, 0, len(header.infoOrder))
	for _, id := range header.infoOrder {
		def := header.info[id]
		fieldType, ok := parseFieldType(def.typ)
		if !ok {
			t.Fatalf("INFO field %q has unsupported type %q", id, def.typ)
		}
		schema = append(schema, Field{Name: id, Type: fieldType, Number: def.number})
	}
	return schema
}

func mustIngestMiss(t *testing.T, db *DB, key []byte) {
	t.Helper()
	if v, ok, err := db.Get(key); err != nil || ok {
		t.Fatalf("key %x: unexpected hit %q (err %v)", key, v, err)
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
	if err := BuildFromTSVs(base, []string{p1, p2}, BuildOptions{}); err != nil {
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
		v, ok, err := db.Get([]byte(tc.key))
		if err != nil || !ok || string(v) != tc.want {
			t.Fatalf("Get(%q) = %q, %v, %v; want %q", tc.key, v, ok, err, tc.want)
		}
	}
	if _, ok, err := db.Get([]byte("zzz")); err != nil || ok {
		t.Fatalf("unexpected hit for zzz (err %v)", err)
	}
}

func TestBuildFromTSVsUnsorted(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "x.tsv")
	writeIngestFile(t, p, "c\tC\nb\tB\n")
	base := filepath.Join(dir, "genes")
	if err := BuildFromTSVs(base, []string{p}, BuildOptions{}); err == nil {
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
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	schema := ingestSchema(t, out)
	for _, tc := range []struct {
		key  []byte
		want string
	}{
		{VarKey(0, 100, "C", "A"), "AF=0.1;AC=3;AN=100;DB"},
		{VarKey(0, 100, "C", "G"), "AF=0.2;AC=4;AN=100;DB"},
		{VarKey(0, 100, "C", "T"), "AF=0.5;AC=10;AN=100"},
		{VarKey(1, 50, "G", "A"), "AF=0.3;AC=7;AN=100"},
	} {
		if got := mustIngestGet(t, db, tc.key, schema); got != tc.want {
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
	if err := BuildFromVCFs(out, []string{vcf}, fields, "IMMKV", BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	schema := ingestSchema(t, out)
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
		if got := mustIngestGet(t, db, tc.key, schema); got != tc.want {
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
	if err := BuildFromVCFs(out, []string{vcf}, nil, "IMMKV", BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	schema := ingestSchema(t, out)
	if got := mustIngestGet(t, db, VarKey(0, 10, "A", "C"), schema); got != "IMMKV_AF=0.1;IMMKV_AN=100;IMMKV_DB" {
		t.Fatalf("Get(A>C) = %q", got)
	}
	if got := mustIngestGet(t, db, VarKey(0, 20, "A", "G"), schema); got != "." {
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
	if err := BuildFromVCFs(out, []string{f1, f2}, []string{"AF"}, "IMMKV", BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	schema := ingestSchema(t, out)
	if got := mustIngestGet(t, db, VarKey(0, 100, "C", "A"), schema); got != "IMMKV_AF=0.1" {
		t.Fatalf("chr1 value = %q", got)
	}
	if got := mustIngestGet(t, db, VarKey(1, 50, "G", "A"), schema); got != "IMMKV_AF=0.2" {
		t.Fatalf("chr2 value = %q", got)
	}
}

func TestBuildFromVCFsMultiFileErrors(t *testing.T) {
	dir := t.TempDir()
	f1 := filepath.Join(dir, "1a.vcf")
	f2 := filepath.Join(dir, "1b.vcf")
	writeIngestFile(t, f1, ingestTestHeader+ingestRec("1", 100, "C", "A", "AF=0.1"))
	writeIngestFile(t, f2, ingestTestHeader+ingestRec("1", 200, "C", "T", "AF=0.2"))
	if err := BuildFromVCFs(filepath.Join(dir, "dup"), []string{f1, f2}, nil, "", BuildOptions{}); err == nil ||
		!strings.Contains(err.Error(), "more than one input") {
		t.Fatalf("duplicate chromosome error = %v", err)
	}

	f3 := filepath.Join(dir, "mixed.vcf")
	f4 := filepath.Join(dir, "3.vcf")
	writeIngestFile(t, f3, ingestTestHeader+
		ingestRec("1", 100, "C", "A", "AF=0.1")+
		ingestRec("2", 50, "G", "A", "AF=0.2"))
	writeIngestFile(t, f4, ingestTestHeader+ingestRec("3", 10, "A", "C", "AF=0.3"))
	if err := BuildFromVCFs(filepath.Join(dir, "mixed"), []string{f3, f4}, nil, "", BuildOptions{}); err == nil {
		t.Fatal("expected multi-chromosome file error")
	}

	hdr2 := "##fileformat=VCFv4.2\n" +
		"##INFO=<ID=AF,Number=1,Type=Float,Description=\"Allele frequency\">\n" +
		ingestTestSiteHeader
	f5 := filepath.Join(dir, "5.vcf")
	f6 := filepath.Join(dir, "6.vcf")
	writeIngestFile(t, f5, hdr2+ingestRec("1", 100, "C", "A", "AF=0.1"))
	writeIngestFile(t, f6, ingestTestHeader+ingestRec("2", 50, "G", "A", "AF=0.2"))
	if err := BuildFromVCFs(filepath.Join(dir, "hdr"), []string{f5, f6}, nil, "", BuildOptions{}); err == nil {
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
	if err := BuildFromVCFs(filepath.Join(dir, "order"), []string{f7, f8}, nil, "", BuildOptions{}); err == nil ||
		!strings.Contains(err.Error(), "order") {
		t.Fatalf("INFO order mismatch error = %v", err)
	}
}

func TestBuildFromVCFsTypedArrays(t *testing.T) {
	dir := t.TempDir()
	header := "##fileformat=VCFv4.2\n" +
		"##INFO=<ID=AD,Number=R,Type=Integer,Description=\"AD\">\n" +
		"##INFO=<ID=GL,Number=G,Type=Float,Description=\"GL\">\n" +
		"##INFO=<ID=XX,Number=.,Type=Integer,Description=\"XX\">\n" +
		"##INFO=<ID=STR,Number=.,Type=String,Description=\"STR\">\n" +
		ingestTestSiteHeader
	vcf := filepath.Join(dir, "in.vcf")
	writeIngestFile(t, vcf, header+
		ingestRec("1", 100, "C", "A", "AD=1,.,3;GL=0.5,.;XX=1,2;STR=a,b"))
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	raw, ok, err := db.Get(VarKey(0, 100, "C", "A"))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("typed array record missing")
	}
	schema := ingestSchema(t, out)
	record, err := NewRecord(schema, raw)
	if err != nil {
		t.Fatal(err)
	}
	values := map[string]Value{}
	for field, value := range record.Values() {
		values[field.Name] = value
	}
	ad := values["AD"]
	if ad.Type != TypeInteger || len(ad.Elems) != 3 || ad.Elems[0].Int != 1 || ad.Elems[1].Type != 0 || ad.Elems[2].Int != 3 {
		t.Fatalf("AD = %+v", ad)
	}
	gl := values["GL"]
	if gl.Type != TypeFloat || len(gl.Elems) != 2 || gl.Elems[0].Float != 0.5 || gl.Elems[1].Type != 0 {
		t.Fatalf("GL = %+v", gl)
	}
	xx := values["XX"]
	if xx.Type != TypeInteger || len(xx.Elems) != 2 || xx.Elems[0].Int != 1 || xx.Elems[1].Int != 2 {
		t.Fatalf("XX = %+v", xx)
	}
	if got := string(record.Bytes("STR")); got != "a,b" {
		t.Fatalf("STR = %q", got)
	}
	text, err := RenderRecord(schema, raw)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(text); got != "AD=1,.,3;GL=0.5,.;XX=1,2;STR=a,b" {
		t.Fatalf("RenderRecord = %q", got)
	}
}

func TestBuildFromVCFsTypedMissingPerALT(t *testing.T) {
	dir := t.TempDir()
	header := "##fileformat=VCFv4.2\n" +
		"##INFO=<ID=AF,Number=A,Type=Float,Description=\"AF\">\n" +
		"##INFO=<ID=AC,Number=A,Type=Integer,Description=\"AC\">\n" +
		ingestTestSiteHeader
	vcf := filepath.Join(dir, "in.vcf")
	writeIngestFile(t, vcf, header+
		ingestRec("1", 100, "C", "A,G", "AF=.,0.2;AC=1,2"))
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	schema := ingestSchema(t, out)
	for _, testCase := range []struct {
		key      []byte
		wantAF   bool
		wantAC   int64
		wantText string
	}{
		{VarKey(0, 100, "C", "A"), false, 1, "AC=1"},
		{VarKey(0, 100, "C", "G"), true, 2, "AF=0.2;AC=2"},
	} {
		raw, ok, err := db.Get(testCase.key)
		if err != nil {
			t.Fatal(err)
		}
		if !ok {
			t.Fatalf("key %x missing", testCase.key)
		}
		record, err := NewRecord(schema, raw)
		if err != nil {
			t.Fatal(err)
		}
		if got := record.HasValue("AF"); got != testCase.wantAF {
			t.Fatalf("key %x AF presence = %v", testCase.key, got)
		}
		if got := record.Int("AC"); got != testCase.wantAC {
			t.Fatalf("key %x AC = %d", testCase.key, got)
		}
		text, err := RenderRecord(schema, raw)
		if err != nil {
			t.Fatal(err)
		}
		if got := string(text); got != testCase.wantText {
			t.Fatalf("key %x render = %q", testCase.key, got)
		}
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
		{"bad-float", ingestRec("1", 100, "C", "A", "AF=abc"), nil, "", `field "AF"`},
		{"bad-int", ingestRec("1", 100, "C", "A", "AN=abc"), nil, "", `field "AN"`},
		{"int-overflow", ingestRec("1", 100, "C", "A", "AN=99999999999999999999"), nil, "", `field "AN"`},
		{"bad-prefix", okRec, nil, "A=B", "invalid prefix"},
		{"no-records", "", nil, "", "no buildable records"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			vcf := filepath.Join(dir, tc.name+".vcf")
			writeIngestFile(t, vcf, ingestTestHeader+tc.body)
			err := BuildFromVCFs(filepath.Join(dir, tc.name+"-lib"), []string{vcf}, tc.fields, tc.prefix, BuildOptions{})
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
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", BuildOptions{}); err == nil ||
		!strings.Contains(err.Error(), "already exists") {
		t.Fatalf("existing outdir error = %v", err)
	}
	if err := BuildFromVCFs(dir, nil, nil, "", BuildOptions{}); err == nil {
		t.Fatal("expected empty vcfs error")
	}
}

func TestBuildFromVCFsNonFiniteFloatsAreMissing(t *testing.T) {
	dir := t.TempDir()
	header := "##fileformat=VCFv4.2\n" +
		"##INFO=<ID=AF,Number=A,Type=Float,Description=\"AF\">\n" +
		"##INFO=<ID=AN,Number=1,Type=Integer,Description=\"AN\">\n" +
		"##INFO=<ID=AFR,Number=.,Type=Float,Description=\"array\">\n" +
		ingestTestSiteHeader
	body := ingestRec("1", 100, "C", "A", "AF=nan;AN=100") +
		ingestRec("1", 200, "C", "G", "AF=1.5;AN=100;AFR=0.1,nan,-inf,0.2") +
		ingestRec("1", 300, "C", "T", "AF=Inf;AN=100;AFR=1")
	vcf := filepath.Join(dir, "in.vcf")
	writeIngestFile(t, vcf, header+body)
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	schema := ingestSchema(t, out)
	for _, tc := range []struct {
		key  []byte
		want string
	}{
		{VarKey(0, 100, "C", "A"), "AN=100"},
		{VarKey(0, 200, "C", "G"), "AF=1.5;AN=100;AFR=0.1,.,.,0.2"},
		{VarKey(0, 300, "C", "T"), "AN=100;AFR=1"},
	} {
		if got := mustIngestGet(t, db, tc.key, schema); got != tc.want {
			t.Fatalf("key %x: rendered %q, want %q", tc.key, got, tc.want)
		}
	}
	manifest, err := ReadManifest(filepath.Join(out, "chr1.manifest"))
	if err != nil {
		t.Fatal(err)
	}
	if manifest.NonFiniteFloats != 4 {
		t.Fatalf("non_finite_floats = %d, want 4", manifest.NonFiniteFloats)
	}
}

func TestBuildFromVCFsBGZF(t *testing.T) {
	dir := t.TempDir()
	vcf := filepath.Join(dir, "in.vcf.gz")
	writeIngestBGZF(t, vcf, ingestTestHeader+ingestRec("1", 10, "A", "C", "AF=0.1;AC=1;AN=100"))
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, []string{"AF"}, "IMMKV", BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	schema := ingestSchema(t, out)
	if got := mustIngestGet(t, db, VarKey(0, 10, "A", "C"), schema); got != "IMMKV_AF=0.1" {
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

func TestBuildFromVCFsTypedRecord(t *testing.T) {
	dir := t.TempDir()
	header := "##fileformat=VCFv4.2\n" +
		"##INFO=<ID=AF,Number=A,Type=Float,Description=\"AF\">\n" +
		"##INFO=<ID=AC,Number=A,Type=Integer,Description=\"AC\">\n" +
		"##INFO=<ID=AN,Number=1,Type=Integer,Description=\"AN\">\n" +
		"##INFO=<ID=DB,Number=0,Type=Flag,Description=\"DB\">\n" +
		"##INFO=<ID=STR,Number=1,Type=String,Description=\"STR\">\n" +
		ingestTestSiteHeader
	vcf := filepath.Join(dir, "in.vcf")
	writeIngestFile(t, vcf, header+
		ingestRec("1", 100, "C", "A", "AF=0.25;AC=3;AN=100;DB;STR=foo"))
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	raw, ok, err := db.Get(VarKey(0, 100, "C", "A"))
	if err != nil {
		t.Fatal(err)
	}
	if !ok {
		t.Fatal("typed record missing")
	}
	schema := ingestSchema(t, out)
	record, err := NewRecord(schema, raw)
	if err != nil {
		t.Fatal(err)
	}
	if !record.HasValue("AF") || record.Float("AF") != 0.25 {
		t.Fatalf("AF = %v", record.Float("AF"))
	}
	if record.Int("AC") != 3 || record.Int("AN") != 100 {
		t.Fatalf("AC/AN = %d/%d", record.Int("AC"), record.Int("AN"))
	}
	if !record.HasValue("DB") {
		t.Fatal("DB should be present")
	}
	if got := string(record.Bytes("STR")); got != "foo" {
		t.Fatalf("STR = %q", got)
	}
	text, err := RenderRecord(schema, raw)
	if err != nil {
		t.Fatal(err)
	}
	if got := string(text); got != "AF=0.25;AC=3;AN=100;DB;STR=foo" {
		t.Fatalf("RenderRecord = %q", got)
	}

	valsFile, err := os.Open(filepath.Join(out, "chr1.vals"))
	if err != nil {
		t.Fatal(err)
	}
	defer valsFile.Close()
	var headerBytes [12]byte
	if _, err := valsFile.ReadAt(headerBytes[:], 0); err != nil {
		t.Fatal(err)
	}
	if got := string(headerBytes[:8]); got != valsMagic {
		t.Fatalf("vals magic = %q, want %q", got, valsMagic)
	}
	if version := binary.LittleEndian.Uint32(headerBytes[8:12]); version != 3 {
		t.Fatalf("vals version = %d, want 3", version)
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
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	db, err := OpenEager(out, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	schema := ingestSchema(t, out)
	if got := mustIngestGet(t, db, VarKey(0, 10, "A", "C"), schema); got != "AF=0.1;AN=100" {
		t.Fatalf("Get = %q; sample columns must not enter the value", got)
	}
}

func TestBuildFromVCFsRejectsExistingOutdir(t *testing.T) {
	dir := t.TempDir() // 已存在，哪怕是空目录也必须拒绝
	err := BuildFromVCFs(dir, []string{filepath.Join(dir, "unused.vcf")}, nil, "", BuildOptions{})
	if err == nil || !strings.Contains(err.Error(), "already exists") {
		t.Fatalf("已存在的 outdir 应报错，得到 %v", err)
	}
}

func TestBuildFromVCFsCleansFailedBuild(t *testing.T) {
	dir := t.TempDir()
	vcf := filepath.Join(dir, "empty.vcf")
	writeIngestFile(t, vcf, "") // 没有 #CHROM，准备输入阶段报错
	out := filepath.Join(dir, "lib")
	if err := BuildFromVCFs(out, []string{vcf}, nil, "", BuildOptions{}); err == nil {
		t.Fatal("空 VCF 应报错")
	}
	if _, err := os.Stat(out); !os.IsNotExist(err) {
		t.Fatalf("失败构建应删除本次创建的 outdir，得到 %v", err)
	}
}

func TestBuildFromTSVsRejectsExisting(t *testing.T) {
	dir := t.TempDir()
	tsv := filepath.Join(dir, "a.tsv")
	writeIngestFile(t, tsv, "a\tA\n")
	base := filepath.Join(dir, "lib")
	if err := BuildFromTSVs(base, []string{tsv}, BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	if err := BuildFromTSVs(base, []string{tsv}, BuildOptions{}); err == nil {
		t.Fatal("已有同名库应拒绝重建")
	}
	db, err := OpenSimpleKV(base, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	if v, ok, err := db.Get([]byte("a")); err != nil || !ok || string(v) != "A" {
		t.Fatalf("拒绝覆盖后原库应仍可读：%q %v %v", v, ok, err)
	}
}
