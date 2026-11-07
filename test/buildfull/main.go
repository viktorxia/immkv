// buildfull builds the full-load immkv library from raw gnomAD v2.1.1 sites
// files: every INFO byte of every record, one key per non-symbolic ALT
// element. This is the program that produced the 366 GiB library used by the
// 540-field experiments; it lives here for provenance and re-buildability,
// not because the test suite rebuilds the library (a full build takes ~2h).
//
// Usage:
//
//	buildfull -gnomad <sitesDir> -out <libDir> [-only chrN]
//
// Environment knobs (same as the original builder):
//
//	IMMKV_BUILD_WORKERS=N   parallel per-chromosome workers (default 8)
//	IMMKV_RAW=1             store all values uncompressed (framing experiment)
//	IMMKV_PACK_KB=N         PackTarget override; -1 trains dictionaries
//
// Value semantics: the whole INFO string of the source record is the value of
// every ALT-derived key of that record (Number=A fields keep all elements;
// annotate documents how they surface). Keys skip symbolic/missing alleles:
// alt starting with '<', '*', '.', or containing breakend brackets.
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
	"sync"
	"time"

	"github.com/biogo/hts/bgzf"
	"github.com/viktorxia/immkv"
)

func main() {
	gnomad := flag.String("gnomad", "", "directory holding gnomad.genomes.r2.1.1.sites.<N>.vcf.bgz")
	out := flag.String("out", "", "output library directory")
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
	bgzs, err := filepath.Glob(filepath.Join(gnomadDir, "gnomad.genomes.r2.1.1.sites.*.vcf.bgz"))
	if err != nil {
		return err
	}
	type job struct {
		path string
		c    uint8
		size int64
	}
	var jobs []job
	for _, p := range bgzs {
		base := filepath.Base(p)
		// gnomad.genomes.r2.1.1.sites.<N>.vcf.bgz -> chromosome <N>
		parts := strings.Split(strings.TrimSuffix(strings.TrimSuffix(base, ".vcf.bgz"), ".bgz"), ".")
		chrom := parts[len(parts)-1]
		c, ok := immkv.ChromID(chrom)
		if !ok {
			fmt.Printf("skipping non-standard contig file: %s\n", base)
			continue
		}
		if only != "" && "chr"+immkv.ChromName(c) != only {
			continue
		}
		st, err := os.Stat(p)
		if err != nil {
			return err
		}
		jobs = append(jobs, job{path: p, c: c, size: st.Size()})
	}
	if len(jobs) == 0 {
		return fmt.Errorf("no buildable files (only=%q)", only)
	}
	// Largest chromosome first: with a fixed worker pool this minimizes the
	// tail where one big file runs alone.
	sort.Slice(jobs, func(i, j int) bool { return jobs[i].size > jobs[j].size })
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return err
	}
	workers := 8
	if v := os.Getenv("IMMKV_BUILD_WORKERS"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 {
			workers = n
		}
	}
	fmt.Printf("building %d chromosomes -> %s (%d workers, largest first)\n", len(jobs), outDir, workers)
	totalStart := time.Now()
	q := make(chan job)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range q {
				if err := buildChrom(outDir, j.c, j.path); err != nil {
					fmt.Printf("!! chr%s failed: %v\n", immkv.ChromName(j.c), err)
					os.Exit(1)
				}
			}
		}()
	}
	for _, j := range jobs {
		q <- j
	}
	close(q)
	wg.Wait()
	fmt.Printf("all done in %s\n", time.Since(totalStart).Round(time.Second))
	return nil
}

func buildChrom(outDir string, c uint8, path string) error {
	name := "chr" + immkv.ChromName(c)
	base := filepath.Join(outDir, name)
	if fi, err := os.Stat(base + ".keys"); err == nil && fi.Size() > 64 {
		fmt.Printf("skip existing: %s\n", name)
		return nil
	}
	start := time.Now()
	opts := immkv.Options{}
	if os.Getenv("IMMKV_RAW") != "" {
		opts.RawValueBelow = 1 << 30 // store everything uncompressed
	}
	if v := os.Getenv("IMMKV_PACK_KB"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n != 0 {
			if n > 0 {
				opts.PackTarget = n << 10
			} else {
				opts.PackTarget = -1 // frame each value + trained dictionaries
			}
		}
	}
	w, err := immkv.Create(base, opts)
	if err != nil {
		return err
	}
	sites, keys, err := walkVCFFile(path, func(c2 uint8, pos uint32, ref, alt, info []byte) error {
		key := immkv.VarKey(c2, pos, string(ref), string(alt))
		return w.AddSorted(immkv.PosKey(c2, pos), key, info)
	})
	if err != nil {
		return err
	}
	if err := w.Finish(); err != nil {
		return err
	}
	fmt.Printf("done %s: %d sites / %d keys in %s (source %s)\n",
		name, sites, keys, time.Since(start).Round(time.Second), humanBytes(fileSize(path)))
	return nil
}

// walkVCFFile streams a (bgzipped) VCF, splitting each record into one emit
// call per ALT element: skip ref='.' and symbolic alleles ('<', '*', '.',
// breakend brackets). INFO is passed through raw as the value.
func walkVCFFile(path string, emit func(c uint8, pos uint32, ref, alt, info []byte) error) (sites, keys uint64, err error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, 0, err
	}
	defer f.Close()
	r, err := bgzf.NewReader(f, 0)
	if err != nil {
		return 0, 0, err
	}
	defer r.Close()
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 64*1024), 8*1024*1024)
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		var cols [8][]byte
		rest := line
		for i := 0; i < 8; i++ {
			if i == 7 {
				cols[i] = rest
				break
			}
			j := bytes.IndexByte(rest, '\t')
			if j < 0 {
				return sites, keys, fmt.Errorf("%s: malformed line: %q", filepath.Base(path), line)
			}
			cols[i] = rest[:j]
			rest = rest[j+1:]
		}
		c, ok := immkv.ChromID(string(cols[0]))
		if !ok {
			continue
		}
		pos, err := strconv.ParseUint(string(cols[1]), 10, 32)
		if err != nil {
			return sites, keys, fmt.Errorf("POS %q: %w", cols[1], err)
		}
		sites++
		if string(cols[3]) == "." {
			continue
		}
		for _, alt := range bytes.Split(cols[4], []byte(",")) {
			if len(alt) == 0 || alt[0] == '<' || alt[0] == '*' || alt[0] == '.' ||
				bytes.ContainsAny(alt, "[]") {
				continue
			}
			if err := emit(c, uint32(pos), cols[3], alt, cols[7]); err != nil {
				return sites, keys, err
			}
			keys++
		}
	}
	return sites, keys, sc.Err()
}

func fileSize(path string) int64 {
	st, err := os.Stat(path)
	if err != nil {
		return 0
	}
	return st.Size()
}

func humanBytes(n int64) string {
	switch {
	case n >= 1<<30:
		return fmt.Sprintf("%.1fG", float64(n)/(1<<30))
	case n >= 1<<20:
		return fmt.Sprintf("%.1fM", float64(n)/(1<<20))
	default:
		return fmt.Sprintf("%.1fK", float64(n)/(1<<10))
	}
}
