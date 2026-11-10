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
// Value semantics: the INFO fields of every source record are reordered to the
// ##INFO declaration order and stored as the value of every ALT-derived key
// (Number=A fields keep all elements; annotate documents how they surface).
// Unknown/duplicate fields, empty segments, Flag-with-value, and
// non-Flag-without-value fail the build; a record with no INFO fields stores
// ".". Keys skip symbolic/missing alleles: alt starting with '<', '*', '.',
// or containing breakend brackets.
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
		os.Remove(base + ".keys.tmp")
		os.Remove(base + ".vals.tmp")
		return err
	}
	if err := w.Finish(); err != nil {
		return err
	}
	fmt.Printf("done %s: %d sites / %d keys in %s (source %s)\n",
		name, sites, keys, time.Since(start).Round(time.Second), humanBytes(fileSize(path)))
	return nil
}

// walkVCFFile streams a (bgzipped) VCF, reorders each record's INFO fields to
// the ##INFO declaration order (the same normalization as BuildFromVCFs), and
// splits each record into one emit call per ALT element: skip ref='.' and
// symbolic alleles ('<', '*', '.', breakend brackets). INFO is passed through
// as the normalized value.
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

	var (
		infoByID   = make(map[string]*infoDef)
		infoOrder  []string
		seenInfo   []uint32
		gen        uint32
		infoFields []infoField
		infoBuf    []byte
	)
	// sortInfo reorders one record's INFO fields to the ##INFO declaration
	// order and rebuilds the value. A record with no INFO fields becomes ".".
	// Unknown/duplicate fields, empty segments, Flag-with-value, and
	// non-Flag-without-value are errors.
	sortInfo := func(info []byte) ([]byte, error) {
		infoBuf = infoBuf[:0]
		if len(info) == 0 || (len(info) == 1 && info[0] == '.') {
			return []byte("."), nil
		}
		infoFields = infoFields[:0]
		gen++
		rest := info
		for len(rest) > 0 {
			var seg []byte
			if i := bytes.IndexByte(rest, ';'); i >= 0 {
				seg, rest = rest[:i], rest[i+1:]
			} else {
				seg, rest = rest, nil
			}
			if len(seg) == 0 {
				return nil, fmt.Errorf("empty INFO segment")
			}
			eq := bytes.IndexByte(seg, '=')
			key := seg
			var v []byte
			if eq >= 0 {
				key, v = seg[:eq], seg[eq+1:]
			}
			if len(key) == 0 {
				return nil, fmt.Errorf("empty INFO field name")
			}
			d, ok := infoByID[string(key)]
			if !ok {
				return nil, fmt.Errorf("unknown INFO field %q", key)
			}
			if seenInfo[d.order] == gen {
				return nil, fmt.Errorf("duplicate INFO field %q", key)
			}
			seenInfo[d.order] = gen
			if d.isFlag && eq >= 0 {
				return nil, fmt.Errorf("Flag INFO field %q has a value", key)
			}
			if !d.isFlag && eq < 0 {
				return nil, fmt.Errorf("non-Flag INFO field %q has no value", key)
			}
			infoFields = append(infoFields, infoField{
				order: d.order, id: string(key), val: v, isFlag: d.isFlag,
			})
		}
		sort.Slice(infoFields, func(i, j int) bool {
			return infoFields[i].order < infoFields[j].order
		})
		for _, f := range infoFields {
			if len(infoBuf) > 0 {
				infoBuf = append(infoBuf, ';')
			}
			infoBuf = append(infoBuf, f.id...)
			if f.isFlag {
				continue
			}
			infoBuf = append(infoBuf, '=')
			infoBuf = append(infoBuf, f.val...)
		}
		return infoBuf, nil
	}

	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 64*1024), 8*1024*1024)
	for sc.Scan() {
		line := sc.Bytes()
		if n := len(line); n > 0 && line[n-1] == '\r' {
			line = line[:n-1]
		}
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			if bytes.HasPrefix(line, []byte("##INFO=<")) {
				if id, d, ok := parseInfoLine(line); ok {
					if _, exists := infoByID[id]; exists {
						return sites, keys, fmt.Errorf("%s: duplicate ##INFO %q", filepath.Base(path), id)
					}
					d.order = len(infoOrder)
					infoByID[id] = d
					infoOrder = append(infoOrder, id)
					seenInfo = append(seenInfo, 0)
				}
			}
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
		nValid := 0
		for _, alt := range bytes.Split(cols[4], []byte(",")) {
			if skipAlt(alt) {
				continue
			}
			nValid++
		}
		if nValid == 0 {
			continue
		}
		value, err := sortInfo(cols[7])
		if err != nil {
			return sites, keys, fmt.Errorf("%s: %w", filepath.Base(path), err)
		}
		for _, alt := range bytes.Split(cols[4], []byte(",")) {
			if skipAlt(alt) {
				continue
			}
			if err := emit(c, uint32(pos), cols[3], alt, value); err != nil {
				return sites, keys, err
			}
			keys++
		}
	}
	return sites, keys, sc.Err()
}

// skipAlt is the ALT filter shared with BuildFromVCFs.
func skipAlt(alt []byte) bool {
	return len(alt) == 0 || alt[0] == '<' || alt[0] == '*' || alt[0] == '.' ||
		bytes.ContainsAny(alt, "[]")
}

// infoDef is one ##INFO declaration.
type infoDef struct {
	number string
	typ    string
	isFlag bool
	order  int
}

// infoField is one INFO field waiting to be re-emitted in ##INFO order.
type infoField struct {
	order  int
	id     string
	val    []byte
	isFlag bool
}

// parseInfoLine parses ##INFO=<ID=...,Number=...,Type=...,Description="...">.
func parseInfoLine(line []byte) (string, *infoDef, bool) {
	const prefix = "##INFO=<"
	if !bytes.HasPrefix(line, []byte(prefix)) || len(line) == 0 || line[len(line)-1] != '>' {
		return "", nil, false
	}
	content := line[len(prefix) : len(line)-1]
	var id, number, typ string
	start := 0
	inQuote := false
	for i := 0; i <= len(content); i++ {
		if i < len(content) {
			c := content[i]
			if inQuote {
				if c == '\\' {
					i++
				} else if c == '"' {
					inQuote = false
				}
				continue
			}
			if c == '"' {
				inQuote = true
				continue
			}
			if c != ',' {
				continue
			}
		}
		seg := content[start:i]
		if eq := bytes.IndexByte(seg, '='); eq >= 0 {
			switch string(seg[:eq]) {
			case "ID":
				id = string(headerRawValue(seg[eq+1:]))
			case "Number":
				number = string(headerRawValue(seg[eq+1:]))
			case "Type":
				typ = string(headerRawValue(seg[eq+1:]))
			}
		}
		start = i + 1
	}
	if id == "" {
		return "", nil, false
	}
	d := &infoDef{number: number, typ: typ}
	d.isFlag = number == "0" || typ == "Flag"
	return id, d, true
}

func headerRawValue(v []byte) []byte {
	if len(v) >= 2 && v[0] == '"' && v[len(v)-1] == '"' {
		return v[1 : len(v)-1]
	}
	return v
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
