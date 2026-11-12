package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/biogo/hts/bgzf"
	"github.com/viktorxia/immkv"
)

const buildfullTestHeader = "##fileformat=VCFv4.2\n" +
	"##INFO=<ID=AF,Number=A,Type=Float,Description=\"AF\">\n" +
	"##INFO=<ID=AC,Number=A,Type=Integer,Description=\"AC\">\n" +
	"##INFO=<ID=AN,Number=1,Type=Integer,Description=\"AN\">\n" +
	"##INFO=<ID=DB,Number=0,Type=Flag,Description=\"DB\">\n" +
	"#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO\n"

func writeTestBGZF(t *testing.T, path, content string) {
	t.Helper()
	var buf bytes.Buffer
	bw := bgzf.NewWriter(&buf, 1)
	if _, err := bw.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := bw.Close(); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		t.Fatal(err)
	}
}

func TestRunBuildsChromosome(t *testing.T) {
	dir := t.TempDir()
	gnomadDir := filepath.Join(dir, "gnomad")
	if err := os.MkdirAll(gnomadDir, 0o755); err != nil {
		t.Fatal(err)
	}
	writeTestBGZF(t, filepath.Join(gnomadDir, "gnomad.genomes.r2.1.1.sites.22.vcf.bgz"),
		buildfullTestHeader+
			"22\t10\t.\tA\tC\t.\tPASS\tDB;AN=100;AC=1;AF=0.1\n"+
			"22\t20\t.\tA\tG\t.\tPASS\t.\n")
	outDir := filepath.Join(dir, "lib")
	if err := run(gnomadDir, outDir, "chr22"); err != nil {
		t.Fatal(err)
	}
	manifest, err := immkv.ReadManifest(filepath.Join(outDir, "chr22.manifest"))
	if err != nil {
		t.Fatal(err)
	}
	if manifest.RecordCount != 2 {
		t.Fatalf("record_count = %d, want 2", manifest.RecordCount)
	}
	for _, ext := range []string{".keys", ".vals", ".col", ".manifest"} {
		if _, err := os.Stat(filepath.Join(outDir, "chr22"+ext)); err != nil {
			t.Fatalf("missing chr22%s: %v", ext, err)
		}
	}
}
