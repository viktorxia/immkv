// buildnarrow builds a narrow immkv library from raw gnomAD v2.1.1 sites
// files: only the fields listed in the TSV are ingested, one key per ALT
// element. Used for the 17-field library-to-library comparison where both
// sides carry identical information.
//
// Usage:
//
//	buildnarrow -gnomad <sitesDir> -out <libDir> -fields <fields.tsv>
//
// NARROW_ONLY=<chrom> (environment) restricts the build to one chromosome,
// e.g. NARROW_ONLY=22 for a smoke test.
//
// Input parsing notes (all of these were learned the hard way):
//   - gnomAD sites files have exactly 8 columns / 7 tabs. A parser that
//     expects 8 tabs finds none and silently ingests zero records.
//   - REF spans tabs[2]+1..tabs[3], ALT spans tabs[3]+1..tabs[4], INFO starts
//     after tabs[6]+1.
//
// Value semantics, mirroring the full library so annotate needs no changes:
// one "k=v;k=v" string per ALT element; Number=A fields contribute their
// element for that ALT index; fields absent from a record are omitted.
//
// The build is single-threaded and sequential over chromosomes 1..22.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/biogo/hts/bgzf"
	"github.com/viktorxia/immkv"
)

type field struct {
	name   string
	number string
}

func main() {
	gnomad := flag.String("gnomad", "", "directory holding gnomad.genomes.r2.1.1.sites.<N>.vcf.bgz")
	out := flag.String("out", "", "output library directory")
	fieldsPath := flag.String("fields", "", "field list TSV from genfields (name/Type/Number)")
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

func loadFields(path string) ([]field, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var out []field
	for _, ln := range bytes.Split(b, []byte("\n")) {
		if len(ln) == 0 || ln[0] == '#' {
			continue
		}
		cols := bytes.Split(ln, []byte("\t"))
		if len(cols) < 3 {
			return nil, fmt.Errorf("bad fields line %q", ln)
		}
		out = append(out, field{name: string(cols[0]), number: string(cols[2])})
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no fields in %s", path)
	}
	return out, nil
}

func run(gnomadDir, out, fieldsPath string) error {
	fields, err := loadFields(fieldsPath)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(out, 0o755); err != nil {
		return err
	}
	only := os.Getenv("NARROW_ONLY")
	total0 := time.Now()
	for c := 1; c <= 22; c++ {
		cs := strconv.Itoa(c)
		if only != "" && cs != only {
			continue
		}
		t0 := time.Now()
		path := filepath.Join(gnomadDir, "gnomad.genomes.r2.1.1.sites."+cs+".vcf.bgz")
		f, err := os.Open(path)
		if err != nil {
			return err
		}
		br, err := bgzf.NewReader(f, 0)
		if err != nil {
			return err
		}
		w, err := immkv.Create(filepath.Join(out, "chr"+cs), immkv.Options{})
		if err != nil {
			return err
		}
		cid, _ := immkv.ChromID(cs)
		sc := bufio.NewScanner(br)
		sc.Buffer(make([]byte, 1024*1024), 16*1024*1024)
		n := 0
		var valBuf []byte
		// elemLists[i] holds the comma-separated elements of field i's value
		// for the current record (nil if absent).
		elemLists := make([][][]byte, len(fields))
		for sc.Scan() {
			line := sc.Bytes()
			if len(line) == 0 || line[0] == '#' {
				continue
			}
			var tabs [8]int
			nt := 0
			for j := 0; j < len(line) && nt < 8; j++ {
				if line[j] == '\t' {
					tabs[nt] = j
					nt++
				}
			}
			if nt < 7 { // exactly 8 columns / 7 tabs in gnomAD sites files
				continue
			}
			pos, err := strconv.ParseUint(string(line[tabs[0]+1:tabs[1]]), 10, 32)
			if err != nil {
				continue
			}
			ref := line[tabs[2]+1 : tabs[3]]
			altCol := line[tabs[3]+1 : tabs[4]]
			info := line[tabs[6]+1:]
			if j := bytes.IndexByte(info, '\t'); j >= 0 {
				info = info[:j]
			}
			for i := range elemLists {
				elemLists[i] = elemLists[i][:0]
			}
			rest := info
			for len(rest) > 0 {
				var seg []byte
				if j := bytes.IndexByte(rest, ';'); j >= 0 {
					seg, rest = rest[:j], rest[j+1:]
				} else {
					seg, rest = rest, nil
				}
				eq := bytes.IndexByte(seg, '=')
				if eq < 0 {
					continue
				}
				for fi, fd := range fields {
					if string(seg[:eq]) == fd.name {
						lst := bytes.Split(seg[eq+1:], []byte(","))
						elemLists[fi] = append(elemLists[fi][:0], lst...)
						break
					}
				}
			}
			for i, a := range bytes.Split(altCol, []byte(",")) {
				valBuf = valBuf[:0]
				for fi, fd := range fields {
					lst := elemLists[fi]
					if len(lst) == 0 {
						continue // field absent from this record: omit
					}
					var v []byte
					if fd.number == "A" {
						if i >= len(lst) {
							continue // no element for this ALT index: omit
						}
						v = lst[i]
					} else {
						v = lst[0]
					}
					valBuf = append(valBuf, fd.name...)
					valBuf = append(valBuf, '=')
					valBuf = append(valBuf, v...)
					valBuf = append(valBuf, ';')
				}
				if len(valBuf) > 0 {
					valBuf = valBuf[:len(valBuf)-1]
				}
				if err := w.AddSorted(immkv.PosKey(cid, uint32(pos)), immkv.VarKey(cid, uint32(pos), string(ref), string(a)), valBuf); err != nil {
					return err
				}
			}
			n++
		}
		if err := sc.Err(); err != nil {
			return err
		}
		br.Close()
		f.Close()
		if err := w.Finish(); err != nil {
			return err
		}
		fmt.Printf("chr%s: %d records ingested in %s\n", cs, n, time.Since(t0).Round(time.Second))
	}
	fmt.Printf("all chromosomes done in %s\n", time.Since(total0).Round(time.Second))
	return nil
}
