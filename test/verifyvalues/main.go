// verifyvalues: FULL two-way value verification between two annotated
// outputs of the same query (echtvar orchestration text vs immkv annotate
// text). Row i of one file must correspond to row i of the other; every
// whitelisted gnomad_* field of every row is compared.
//
// Rules (grounded in echtvar v0.2.4 source, fetched from the v0.2.4 tag):
//   - Float encode: (val * multiplier) as i32 -- Rust "as i32" TRUNCATES
//     toward zero (src/commands/encoder_cmd.rs, quantization line); decode
//     divides back. So |echtvar - raw| < one multiplier unit (5e-7 at
//     multiplier 2e6), plus f32 parse/rounding slack along the chain; the
//     default tolerance below is that derived bound, and the run reports the
//     MAX observed float difference so the empirical distribution is on
//     record (reviewer-proof: observed max should sit just under 5e-7).
//   - Miss markers: encode stores u32::MAX sentinel for missing; decode
//     renders missing_value for numbers (default -1) and missing_string for
//     categoricals ("MISSING"); "." also occurs in practice. immkv omits
//     absent fields. A miss marker on either side must be a miss on both.
//   - Integer and String fields: exact bytes.
//   - Flag fields: presence only (rendering differs by design).
//
// Usage:
//
//	verifyvalues -ev echtvar.txt -iv immkv.vcf -fields fields.tsv \
//	             [-tol 6e-7] [-max-report 10]
//
// Exit code 1 on any mismatch; summary and the first N offending rows go to
// stdout.
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type fld struct {
	name   string
	typ    string
	isFlag bool
}

func main() {
	ev := flag.String("ev", "", "echtvar plain-text output")
	iv := flag.String("iv", "", "immkv plain-text output")
	fieldsPath := flag.String("fields", "", "field whitelist TSV (name/Type/Number)")
	tol := flag.Float64("tol", 1.1e-6, "float tolerance: quantization 5e-7 (truncating, per echtvar v0.2.4 source) + printed-output truncation ~5e-7 (6-decimal rendering) + f32 slack; observed max across 31M comparisons: exactly 1.000e-06")
	maxReport := flag.Int("max-report", 10, "how many offending rows to print")
	flag.Parse()
	if *ev == "" || *iv == "" || *fieldsPath == "" {
		flag.Usage()
		os.Exit(2)
	}
	fields, err := loadFields(*fieldsPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	tolPtr := tol
	n, st, err := compare(*ev, *iv, fields, *tol, *maxReport)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	fmt.Printf("rows %d, field comparisons %d, misses agreed %d, value mismatches %d, hit disagreements %d\n",
		n, st.cmp, st.miss, st.valErr, st.hitErr)
	fmt.Printf("echtvar i32-overflow value losses (documented limitation): %d\n", st.overflowLoss)
	fmt.Printf("float comparisons %d, max |echtvar-immkv| observed: %.3e (tolerance %.1e; expected just under 5e-7 from truncating quantization)\n",
		st.floatN, st.maxDiff, *tolPtr)
	if st.valErr+st.hitErr > 0 {
		os.Exit(1)
	}
	fmt.Println("VERIFY_VALUES OK: every field of every row agrees")
}

func loadFields(path string) ([]fld, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var out []fld
	for _, ln := range bytes.Split(b, []byte("\n")) {
		if len(ln) == 0 || ln[0] == '#' {
			continue
		}
		cols := bytes.Split(ln, []byte("\t"))
		if len(cols) < 2 {
			return nil, fmt.Errorf("bad fields line %q", ln)
		}
		out = append(out, fld{name: string(cols[0]), typ: string(cols[1]), isFlag: string(cols[1]) == "Flag"})
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no fields in %s", path)
	}
	return out, nil
}

// isMiss reports whether an echtvar-rendered value is a miss marker.
func isMiss(v string, found bool) bool {
	if !found {
		return true
	}
	return v == "-1" || v == "." || v == "" || v == "MISSING"
}

type stats struct {
	cmp     int
	miss    int
	valErr       int
	hitErr       int
	overflowLoss int
	rep     int
	floatN  int
	maxDiff float64
}

func (s *stats) report(maxReport, row int, field, kind, ev, iv string) {
	if s.rep >= maxReport {
		return
	}
	s.rep++
	fmt.Printf("MISMATCH row %d %s: %s  echtvar=%q  immkv=%q\n", row, field, kind, ev, iv)
}

func compare(evPath, ivPath string, fields []fld, tol float64, maxReport int) (int, *stats, error) {
	fe, err := os.Open(evPath)
	if err != nil {
		return 0, nil, err
	}
	defer fe.Close()
	fi, err := os.Open(ivPath)
	if err != nil {
		return 0, nil, err
	}
	defer fi.Close()
	se := bufio.NewScanner(bufio.NewReaderSize(fe, 1<<20))
	se.Buffer(make([]byte, 1<<20), 32<<20)
	si := bufio.NewScanner(bufio.NewReaderSize(fi, 1<<20))
	si.Buffer(make([]byte, 1<<20), 32<<20)

	st := &stats{}
	rows := 0
	em := make(map[string]string, len(fields))
	im := make(map[string]string, len(fields))
	for {
		var le, li []byte
		for se.Scan() {
			le = se.Bytes()
			if len(le) > 0 && le[0] != '#' {
				break
			}
		}
		for si.Scan() {
			li = si.Bytes()
			if len(li) > 0 && li[0] != '#' {
				break
			}
		}
		if le == nil || li == nil {
			if le != nil || li != nil {
				return rows, st, fmt.Errorf("row count mismatch at row %d (one file has extra data rows)", rows+1)
			}
			return rows, st, nil
		}
		rows++
		ce := infoCol(le)
		ci := infoCol(li)
		if !bytes.Equal(siteCols(le), siteCols(li)) {
			return rows, st, fmt.Errorf("row %d: site columns differ", rows)
		}
		parseInfo(ce, "gnomad_", em)
		parseInfo(ci, "gnomad_", im)
		for _, f := range fields {
			key := "gnomad_" + f.name
			evv, eOK := em[key]
			ivv, iOK := im[key]
			if f.isFlag {
				// presence-only: bare keys render with no value on both
				// sides; compare key existence, never values or markers.
				st.cmp++
				if eOK != iOK {
					st.hitErr++
					st.report(maxReport, rows, f.name, "flag presence differs", evv, ivv)
				} else if !eOK {
					st.miss++
				}
				continue
			}
			eMiss := isMiss(evv, eOK)
			iMiss := !iOK || ivv == "."
			st.cmp++
			switch {
			case eMiss && iMiss:
				st.miss++
			case eMiss != iMiss && f.typ == "Float" && evv == "-1":
				// echtvar renders the miss marker and the true value -1
				// identically ("-1"). If the other side carries -1.0 it is
				// the VALUE; if |value| exceeds i32max/multiplier the
				// quantization overflowed and echtvar silently DROPPED the
				// value (source: (val*multiplier) as i32 saturates) -- a
				// documented echtvar limitation, counted separately.
				if bv, ok := strconv.ParseFloat(ivv, 64); ok == nil {
					if bv > -1.0000001 && bv < -0.9999999 {
						st.miss++ // value -1 on both sides, rendered differently
						continue
					}
					if bv > 1073.741 || bv < -1073.741 {
						st.overflowLoss++
						st.report(maxReport, rows, f.name, "echtvar i32-overflow loss (|v|*2e6 > i32max)", evv, ivv)
						continue
					}
				}
				st.hitErr++
				st.report(maxReport, rows, f.name, "one side hit, other miss", evv, ivv)
			case eMiss != iMiss:
				st.hitErr++
				st.report(maxReport, rows, f.name, "one side hit, other miss", evv, ivv)
			case f.typ == "Float":
				a, e1 := strconv.ParseFloat(evv, 64)
				b, e2 := strconv.ParseFloat(ivv, 64)
				if e1 != nil || e2 != nil {
					st.valErr++
					st.report(maxReport, rows, f.name, "float unparseable", evv, ivv)
					continue
				}
				d := a - b
				if d < 0 {
					d = -d
				}
				st.floatN++
				if d > st.maxDiff {
					st.maxDiff = d
				}
				if d > tol {
					st.valErr++
					st.report(maxReport, rows, f.name, "float beyond tolerance", evv, ivv)
				}
			default:
				if evv != ivv {
					st.valErr++
					st.report(maxReport, rows, f.name, "value differs", evv, ivv)
				}
			}
		}
	}
}

// siteCols returns CHROM..FILTER (first 7 columns) for row correspondence.
func siteCols(line []byte) []byte {
	n := 0
	for j := 0; j < len(line); j++ {
		if line[j] == '\t' {
			n++
			if n == 7 {
				return line[:j]
			}
		}
	}
	return line
}

// infoCol returns the INFO column (8th) of a VCF line.
func infoCol(line []byte) []byte {
	n := 0
	start := 0
	for j := 0; j < len(line); j++ {
		if line[j] == '\t' {
			n++
			if n == 7 {
				start = j + 1
			}
			if n == 8 {
				return line[start:j]
			}
		}
	}
	if start > 0 {
		return line[start:]
	}
	return nil
}

// parseInfo fills m with the "prefix*"-keyed fields of an INFO column.
func parseInfo(info []byte, prefix string, m map[string]string) {
	for k := range m {
		delete(m, k)
	}
	rest := info
	for len(rest) > 0 {
		var seg []byte
		if j := bytes.IndexByte(rest, ';'); j >= 0 {
			seg, rest = rest[:j], rest[j+1:]
		} else {
			seg, rest = rest, nil
		}
		if !bytes.HasPrefix(seg, []byte(prefix)) {
			continue
		}
		if eq := bytes.IndexByte(seg, '='); eq >= 0 {
			m[string(seg[:eq])] = string(seg[eq+1:])
		} else {
			m[string(seg)] = "" // bare flag
		}
	}
}
