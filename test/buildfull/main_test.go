package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/biogo/hts/bgzf"
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

func TestWalkVCFFileSortsINFO(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "in.vcf.gz")
	writeTestBGZF(t, path, buildfullTestHeader+
		"1\t10\t.\tA\tC\t.\tPASS\tDB;AN=100;AC=1;AF=0.1\n"+
		"1\t20\t.\tA\tG\t.\tPASS\t.\n")

	var got []string
	sites, keys, err := walkVCFFile(path, func(c uint8, pos uint32, ref, alt, info []byte) error {
		got = append(got, fmt.Sprintf("%d\t%d\t%s\t%s\t%s", c, pos, ref, alt, info))
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if sites != 2 || keys != 2 {
		t.Fatalf("sites/keys = %d/%d; want 2/2", sites, keys)
	}
	want := []string{
		"0\t10\tA\tC\tAF=0.1;AC=1;AN=100;DB",
		"0\t20\tA\tG\t.",
	}
	if len(got) != len(want) {
		t.Fatalf("emitted %d records; want %d: %v", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("record %d = %q; want %q", i, got[i], want[i])
		}
	}
}

func TestWalkVCFFileINFOErrors(t *testing.T) {
	cases := []struct{ name, info, want string }{
		{"unknown", "AF=0.1;NOPE=1", "unknown INFO field"},
		{"duplicate", "AF=0.1;AF=0.2", "duplicate INFO field"},
		{"empty", "AF=0.1;;AN=100", "empty INFO segment"},
		{"empty-name", "=1", "empty INFO field name"},
		{"flag-value", "DB=1", "Flag INFO field"},
		{"bare", "AF", "non-Flag INFO field"},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			path := filepath.Join(dir, "in.vcf.gz")
			writeTestBGZF(t, path, buildfullTestHeader+
				"1\t10\t.\tA\tC\t.\tPASS\t"+tc.info+"\n")
			_, _, err := walkVCFFile(path, func(uint8, uint32, []byte, []byte, []byte) error {
				return nil
			})
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("error = %v; want substring %q", err, tc.want)
			}
		})
	}
}
