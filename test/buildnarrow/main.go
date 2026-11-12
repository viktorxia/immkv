// buildnarrow builds a narrow immkv library from raw gnomAD v2.1.1 sites
// files: only the fields listed in the TSV are ingested, one key per ALT
// element. It is a thin wrapper over immkv.BuildFromVCFs, which writes v003
// typed records plus the columnar copy and manifest per chromosome.
//
// Usage:
//
//	buildnarrow -gnomad <sitesDir> -out <libDir> -fields <fields.tsv>
//
// The TSV's first column is the INFO name; extra columns are ignored. Schema
// order is the ##INFO declaration order, not the TSV order. NARROW_ONLY=<N>
// (environment) restricts the build to one chromosome number, e.g.
// NARROW_ONLY=22 for a smoke test. The output directory must be empty.
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

	"github.com/viktorxia/immkv"
)

func main() {
	gnomad := flag.String("gnomad", "", "directory holding gnomad.genomes.r2.1.1.sites.<N>.vcf.bgz")
	out := flag.String("out", "", "output library directory (must be empty)")
	fieldsPath := flag.String("fields", "", "field list TSV from genfields (first column is the name)")
	flag.Parse()
	if *gnomad == "" || *out == "" || *fieldsPath == "" {
		flag.Usage()
		os.Exit(2)
	}
	if err := run(*gnomad, *out, *fieldsPath); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func loadFields(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	var fields []string
	for sc.Scan() {
		line := bytes.TrimSpace(sc.Bytes())
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		name := string(bytes.TrimSpace(bytes.SplitN(line, []byte("\t"), 2)[0]))
		if name != "" {
			fields = append(fields, name)
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	if len(fields) == 0 {
		return nil, fmt.Errorf("no fields in %s", path)
	}
	return fields, nil
}

func run(gnomadDir, outDir, fieldsPath string) error {
	fields, err := loadFields(fieldsPath)
	if err != nil {
		return err
	}
	only := os.Getenv("NARROW_ONLY")
	var inputs []string
	for c := 1; c <= 22; c++ {
		chrom := strconv.Itoa(c)
		if only != "" && chrom != only {
			continue
		}
		path := filepath.Join(gnomadDir, "gnomad.genomes.r2.1.1.sites."+chrom+".vcf.bgz")
		if _, err := os.Stat(path); err != nil {
			return err
		}
		inputs = append(inputs, path)
	}
	if len(inputs) == 0 {
		return fmt.Errorf("no buildable files (NARROW_ONLY=%q)", only)
	}
	// Largest first, matching buildvcf/buildfull scheduling.
	sort.Slice(inputs, func(i, j int) bool {
		a, errA := os.Stat(inputs[i])
		b, errB := os.Stat(inputs[j])
		if errA != nil || errB != nil {
			return inputs[i] > inputs[j]
		}
		return a.Size() > b.Size()
	})
	fmt.Printf("building %d chromosome(s) -> %s (fields=%d)\n", len(inputs), outDir, len(fields))
	return immkv.BuildFromVCFs(outDir, inputs, fields, "", immkv.Options{})
}
