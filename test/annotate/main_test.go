package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/viktorxia/immkv"
)

func TestRunUnknownContigs(t *testing.T) {
	dir := t.TempDir()
	source := filepath.Join(dir, "source.vcf")
	writeAnnotationTestFile(t, source, "##fileformat=VCFv4.2\n"+
		"##INFO=<ID=AC,Number=A,Type=Integer,Description=\"Allele count\">\n"+
		"##INFO=<ID=AN,Number=1,Type=Integer,Description=\"Allele number\">\n"+
		"##INFO=<ID=LABEL,Number=1,Type=String,Description=\"Record label\">\n"+
		"#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO\n"+
		"1\t10\t.\tA\tC\t.\tPASS\tAC=11;AN=101;LABEL=firstC\n"+
		"1\t10\t.\tA\tG\t.\tPASS\tAC=12;AN=102;LABEL=firstG\n"+
		"1\t20\t.\tT\tA\t.\tPASS\tAC=21;AN=201;LABEL=secondA\n"+
		"1\t20\t.\tT\tC\t.\tPASS\tAC=22;AN=202;LABEL=secondC\n"+
		"1\t30\t.\tG\tA\t.\tPASS\tAC=31;AN=301;LABEL=thirdA\n"+
		"2\t10\t.\tA\tG\t.\tPASS\tAC=41;AN=401;LABEL=fourthG\n")
	libDir := filepath.Join(dir, "lib")
	if err := immkv.BuildFromVCFs(libDir, []string{source}, nil, "", immkv.BuildOptions{}); err != nil {
		t.Fatal(err)
	}
	fieldsPath := filepath.Join(dir, "fields.tsv")
	writeAnnotationTestFile(t, fieldsPath, "AN\tInteger\t1\nAC\tInteger\tA\nLABEL\tString\t1\n")

	const siteHeader = "##fileformat=VCFv4.2\n" +
		"##INFO=<ID=KEEP,Number=1,Type=String,Description=\"Original value\">\n" +
		"##INFO=<ID=FLAG,Number=0,Type=Flag,Description=\"Original flag\">\n"
	const sampleHeader = siteHeader +
		"##FORMAT=<ID=GT,Number=1,Type=String,Description=\"Genotype\">\n" +
		"##FORMAT=<ID=AD,Number=R,Type=Integer,Description=\"Allelic depths\">\n" +
		"#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO\tFORMAT\tsampleA\tsampleB\n"

	// Each unknown contig collides with a chr1 record if assigned chromosome ID 0.
	const unknown10 = "chrUn_gl000220\t10\tunknown10\tA\tG,T,C\t17.00\tq10\tKEEP=007;FLAG\tGT:AD\t1|3:01,02,03,04\t./.:.\n"
	const unknown20 = "NC_000001.11\t20\tunknown20\tT\tC,A\t.\tPASS\t.\tGT:AD\t1/2:3,4,5\t0/0:9,0,0\n"
	const unknown30 = "Chr1\t30\tunknown30\tG\tA\t.\tPASS\tKEEP=unchanged\tGT\t1\t.\n"
	const unknownSites = "chr1_alt\t10\tunknown-site10\tA\tG,C\t.\tPASS\tKEEP=008;FLAG\n" +
		"HLA-A\t20\tunknown-site20\tT\tA,C\t.\tPASS\t.\n"

	const known10 = "chr1\t10\tknown10\tA\tG,T,C\t19.00\tPASS\tKEEP=known10\tGT:AD\t1/3:1,2,3,4\t0/0:5,0,0,0\n"
	const wantKnown10 = "chr1\t10\tknown10\tA\tG,T,C\t19.00\tPASS\tKEEP=known10;ann_AN=101,102;ann_AC=12,.,11;ann_LABEL=firstC,firstG\tGT:AD\t1/3:1,2,3,4\t0/0:5,0,0,0\n"
	const known20 = "1\t20\tknown20\tT\tC,G,A\t.\tPASS\t.\tGT:AD\t1|3:2,3,0,4\t0/0:9,0,0,0\n"
	const wantKnown20 = "1\t20\tknown20\tT\tC,G,A\t.\tPASS\tann_AN=201,202;ann_AC=22,.,21;ann_LABEL=secondA,secondC\tGT:AD\t1|3:2,3,0,4\t0/0:9,0,0,0\n"
	const known30 = "CHR1\t30\tknown30\tG\tT,A\t.\tPASS\tKEEP=known30;FLAG\tGT:AD\t0/2:5,0,6\t0/0:7,0,0\n"
	const wantKnown30 = "CHR1\t30\tknown30\tG\tT,A\t.\tPASS\tKEEP=known30;FLAG;ann_AN=301;ann_AC=.,31;ann_LABEL=thirdA\tGT:AD\t0/2:5,0,6\t0/0:7,0,0\n"
	const known2 = "chr2\t10\tknown2\tA\tG\t.\tPASS\tKEEP=known2\tGT\t0/1\t1/1\n"
	const wantKnown2 = "chr2\t10\tknown2\tA\tG\t.\tPASS\tKEEP=known2;ann_AN=401;ann_AC=41;ann_LABEL=fourthG\tGT\t0/1\t1/1\n"
	const knownMiss = "chr1\t40\tknown-miss\tA\tT\t.\tPASS\tKEEP=miss\tGT\t0/1\t0/0\n"

	// With batch=2, the rows at positions 11 and 12 form an all-unknown window
	// between mixed windows; repeated positions also exercise shared ALT spans.
	const unknown11 = "chrUn_gl000220\t11\tunknown11\tA\tG,C\t.\tPASS\tKEEP=eleven\tGT\t1/2\t0/0\n"
	const unknown12 = "NC_000001.11\t12\tunknown12\tT\tA\t.\tPASS\tKEEP=twelve\tGT\t0/1\t./.\n"
	const unknown20After = "chr1_alt\t20\tunknown20-after\tT\tA,C\t.\tPASS\tKEEP=after;FLAG\tGT\t1/2\t0/0\n"
	const unknown21 = "HLA-A\t21\tunknown21\tT\tC\t.\tPASS\t.\tGT\t1/1\t0/0\n"
	const unknown40 = "chrUn_gl000220\t40\tunknown40\tA\tT\t.\tPASS\tKEEP=forty\tGT\t0/1\t0/0\n"

	cases := []struct {
		name   string
		header string
		input  string
		want   string
	}{
		{
			name:   "all_unknown_with_samples",
			header: sampleHeader,
			input:  unknown10 + unknown20 + unknown30,
			want:   unknown10 + unknown20 + unknown30,
		},
		{
			name:   "all_unknown_without_samples",
			header: siteHeader + "#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO\n",
			input:  unknownSites,
			want:   unknownSites,
		},
		{
			name:   "mixed_contigs_and_alt_expansions",
			header: sampleHeader,
			input: known10 + unknown10 + unknown11 + unknown12 + unknown20 +
				known20 + unknown20After + unknown21 + known30 + knownMiss + unknown40 + known2,
			want: wantKnown10 + unknown10 + unknown11 + unknown12 + unknown20 +
				wantKnown20 + unknown20After + unknown21 + wantKnown30 + knownMiss + unknown40 + wantKnown2,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for _, cfg := range []struct{ batch, threads int }{{1, 1}, {2, 3}, {64, 3}} {
				t.Run(fmt.Sprintf("batch_%d_threads_%d", cfg.batch, cfg.threads), func(t *testing.T) {
					workDir := t.TempDir()
					in := filepath.Join(workDir, "query.vcf")
					out := filepath.Join(workDir, "annotated.vcf")
					writeAnnotationTestFile(t, in, tc.header+tc.input)
					if err := run(in, libDir, out, fieldsPath, "ann_", cfg.batch, cfg.threads); err != nil {
						t.Fatal(err)
					}
					got, err := os.ReadFile(out)
					if err != nil {
						t.Fatal(err)
					}
					if want := []byte(tc.header + tc.want); !bytes.Equal(got, want) {
						t.Fatalf("annotation output differs:\ngot:\n%s\nwant:\n%s", got, want)
					}
				})
			}
		})
	}
}

func writeAnnotationTestFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
