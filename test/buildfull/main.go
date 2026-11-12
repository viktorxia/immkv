// buildfull builds the full-load immkv library from raw gnomAD v2.1.1 sites
// files: every INFO field declared in the header, one key per non-symbolic
// ALT element. It is a thin wrapper over immkv.BuildFromVCFs, which writes
// v003 typed records plus the columnar copy and manifest per chromosome.
//
// Usage:
//
//	buildfull -gnomad <sitesDir> -out <libDir> [-only chrN]
//
// Environment knobs (same as buildvcf):
//
//	IMMKV_RAW=1             store all values uncompressed (framing experiment)
//	IMMKV_PACK_KB=N         PackTarget override; -1 trains dictionaries
//
// The output directory must be empty; existing chrN libraries are a build
// error rather than being skipped.
package main

import (
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
	only := flag.String("only", "", "restrict to one chromosome, e.g. chr22 (smoke test)")
	flag.Parse()
	if *gnomad == "" || *out == "" {
		flag.Usage()
		os.Exit(2)
	}
	if err := run(*gnomad, *out, *only); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(gnomadDir, outDir, only string) error {
	paths, err := filepath.Glob(filepath.Join(gnomadDir, "gnomad.genomes.r2.1.1.sites.*.vcf.bgz"))
	if err != nil {
		return err
	}
	type input struct {
		path string
		size int64
	}
	var inputs []input
	for _, path := range paths {
		base := filepath.Base(path)
		parts := strings.Split(strings.TrimSuffix(base, ".vcf.bgz"), ".")
		chrom := parts[len(parts)-1]
		id, ok := immkv.ChromID(chrom)
		if !ok {
			fmt.Printf("skipping non-standard contig file: %s\n", base)
			continue
		}
		if only != "" && "chr"+immkv.ChromName(id) != only {
			continue
		}
		st, err := os.Stat(path)
		if err != nil {
			return err
		}
		inputs = append(inputs, input{path: path, size: st.Size()})
	}
	if len(inputs) == 0 {
		return fmt.Errorf("no buildable files (only=%q)", only)
	}
	// Largest first: BuildFromVCFs' fixed worker pool then starts the longest
	// jobs earliest, so the tail is filled by small chromosomes.
	sort.Slice(inputs, func(i, j int) bool { return inputs[i].size > inputs[j].size })
	vcfs := make([]string, len(inputs))
	for i := range inputs {
		vcfs[i] = inputs[i].path
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
	fmt.Printf("building %d chromosome(s) -> %s\n", len(vcfs), outDir)
	return immkv.BuildFromVCFs(outDir, vcfs, nil, "", opts)
}
