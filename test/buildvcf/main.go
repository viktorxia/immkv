// buildvcf builds an immkv library from bgzipped gnomAD-style per-chromosome
// VCFs through immkv.BuildFromVCFs. Unlike buildfull (which writes only the
// record-order shards via Create), this program produces the record-order
// library and its base.col/base.manifest companion in one pass, so it is the
// reproducible entry point for the columnar experiments.
//
// Usage:
//
//	buildvcf -gnomad <sitesDir> -out <libDir> [-only chrN] [-fields fields.tsv] [-prefix p]
//
// -fields is an optional TSV whose first column is the INFO name (extra
// columns are ignored); empty means all INFO fields. The output values always
// follow the ##INFO declaration order. Options.PackTarget can be overridden
// with IMMKV_PACK_KB (negative = per-value framing + trained dictionary);
// IMMKV_RAW=1 stores values uncompressed, matching buildfull's knobs.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/viktorxia/immkv"
)

func main() {
	gnomad := flag.String("gnomad", "", "directory holding gnomad.genomes.r2.1.1.sites.<N>.vcf.bgz")
	out := flag.String("out", "", "output library directory (must be empty)")
	only := flag.String("only", "", "restrict to one chromosome, e.g. chr22")
	fieldsPath := flag.String("fields", "", "optional TSV of INFO names (first column)")
	prefix := flag.String("prefix", "", "prefix for stored INFO names")
	flag.Parse()
	if *gnomad == "" || *out == "" {
		flag.Usage()
		os.Exit(2)
	}
	if err := run(*gnomad, *out, *only, *fieldsPath, *prefix); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(gnomadDir, outDir, only, fieldsPath, prefix string) error {
	files, err := filepath.Glob(filepath.Join(gnomadDir, "gnomad.genomes.r2.1.1.sites.*.vcf.bgz"))
	if err != nil {
		return err
	}
	sort.Strings(files)
	var inputs []string
	for _, p := range files {
		base := filepath.Base(p)
		parts := strings.Split(strings.TrimSuffix(base, ".vcf.bgz"), ".")
		chrom := parts[len(parts)-1]
		c, ok := immkv.ChromID(chrom)
		if !ok {
			fmt.Printf("skipping non-standard contig file: %s\n", base)
			continue
		}
		if only != "" && "chr"+immkv.ChromName(c) != only {
			continue
		}
		inputs = append(inputs, p)
	}
	if len(inputs) == 0 {
		return fmt.Errorf("no buildable files (only=%q)", only)
	}
	// Largest first: the fixed worker pool then starts the longest jobs earliest,
	// so the tail is filled by small chromosomes instead of one giant file.
	sort.Slice(inputs, func(i, j int) bool {
		a, errA := os.Stat(inputs[i])
		b, errB := os.Stat(inputs[j])
		if errA != nil || errB != nil {
			return inputs[i] > inputs[j]
		}
		return a.Size() > b.Size()
	})
	var fields []string
	if fieldsPath != "" {
		fields, err = loadFields(fieldsPath)
		if err != nil {
			return err
		}
	}
	opts := immkv.Options{}
	if os.Getenv("IMMKV_RAW") != "" {
		opts.RawValueBelow = 1 << 30
	}
	if v := os.Getenv("IMMKV_PACK_KB"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n != 0 {
			if n > 0 {
				opts.PackTarget = n << 10
			} else {
				opts.PackTarget = -1
			}
		}
	}
	fmt.Printf("building %d input(s) -> %s (fields=%d, prefix=%q)\n",
		len(inputs), outDir, len(fields), prefix)
	if err := immkv.BuildFromVCFs(outDir, inputs, fields, prefix, opts); err != nil {
		return err
	}
	fmt.Println("done")
	return nil
}

func loadFields(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var out []string
	for sc.Scan() {
		ln := bytes.TrimSpace(sc.Bytes())
		if len(ln) == 0 || ln[0] == '#' {
			continue
		}
		cols := bytes.Split(ln, []byte("\t"))
		name := string(bytes.TrimSpace(cols[0]))
		if name != "" {
			out = append(out, name)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no fields in %s", path)
	}
	return out, nil
}
