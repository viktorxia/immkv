package main

import (
	"bufio"
	"errors"
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func testRow(info string) string {
	return "1\t100\t.\tA\tC\t.\tPASS\t" + info + "\n"
}

func writeCompareFiles(t *testing.T, ev, iv string) (string, string) {
	t.Helper()
	dir := t.TempDir()
	evPath := filepath.Join(dir, "echtvar.vcf")
	ivPath := filepath.Join(dir, "immkv.vcf")
	if err := os.WriteFile(evPath, []byte(ev), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(ivPath, []byte(iv), 0o600); err != nil {
		t.Fatal(err)
	}
	return evPath, ivPath
}

func compareText(t *testing.T, ev, iv string, fields []fld, tol float64) (int, *stats, error) {
	t.Helper()
	evPath, ivPath := writeCompareFiles(t, ev, iv)
	return compare(evPath, ivPath, fields, tol, 0)
}

func TestCompareFiniteFloats(t *testing.T) {
	fields := []fld{{name: "AF", typ: "Float"}}
	for _, tt := range []struct {
		name       string
		ev, iv     string
		tol, diff  float64
		mismatches int
	}{
		{name: "equality with zero tolerance", ev: "0.125", iv: "0.125"},
		{name: "equality with negative zero tolerance", ev: "0", iv: "-0", tol: math.Copysign(0, -1)},
		{name: "equal largest finite values", ev: "1.7976931348623157e308", iv: "1.7976931348623157e308"},
		{name: "within default tolerance", ev: "0", iv: "0.000001", tol: 1.1e-6, diff: 1e-6},
		{name: "beyond default tolerance", ev: "0", iv: "0.0000012", tol: 1.1e-6, diff: 1.2e-6, mismatches: 1},
		{name: "at tolerance", ev: "0.5", iv: "0.625", tol: 0.125, diff: 0.125},
		{name: "below tolerance", ev: "0.5", iv: "0.5625", tol: 0.125, diff: 0.0625},
		{name: "above tolerance", ev: "0.5", iv: "0.75", tol: 0.125, diff: 0.25, mismatches: 1},
		{name: "negative difference", ev: "-0.5", iv: "-0.625", tol: 0.125, diff: 0.125},
		{name: "zero tolerance rejects difference", ev: "0", iv: "5e-324", diff: math.SmallestNonzeroFloat64, mismatches: 1},
		{name: "finite subtraction overflows", ev: "1.7976931348623157e308", iv: "-1.7976931348623157e308", tol: math.MaxFloat64, diff: math.Inf(1), mismatches: 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			n, st, err := compareText(t, testRow("gnomad_AF="+tt.ev), testRow("gnomad_AF="+tt.iv), fields, tt.tol)
			want := stats{cmp: 1, valErr: tt.mismatches, floatN: 1, maxDiff: tt.diff}
			if err != nil || n != 1 || st == nil || *st != want {
				t.Fatalf("compare = (%d, %+v, %v), want (1, %+v, nil)", n, st, err, want)
			}
		})
	}
}

func TestCompareNonFiniteFloats(t *testing.T) {
	fields := []fld{{name: "AF", typ: "Float"}}
	for _, value := range []string{"NaN", "+Inf", "-Inf", "Infinity"} {
		t.Run(value, func(t *testing.T) {
			bad := "gnomad_AF=" + value
			for _, tt := range []struct {
				name   string
				ev, iv string
			}{
				{name: "echtvar value", ev: bad, iv: "gnomad_AF=0.5"},
				{name: "immkv value", ev: "gnomad_AF=0.5", iv: bad},
				{name: "same value on both sides", ev: bad, iv: bad},
				{name: "immkv absent", ev: bad, iv: "."},
				{name: "echtvar absent", ev: ".", iv: bad},
				{name: "echtvar miss marker", ev: "gnomad_AF=-1", iv: bad},
				{name: "immkv true minus one", ev: bad, iv: "gnomad_AF=-1"},
			} {
				t.Run(tt.name, func(t *testing.T) {
					n, st, err := compareText(t, testRow(tt.ev), testRow(tt.iv), fields, 1.1e-6)
					want := stats{cmp: 1, valErr: 1}
					if err != nil || n != 1 || st == nil || *st != want {
						t.Fatalf("compare = (%d, %+v, %v), want (1, %+v, nil)", n, st, err, want)
					}
				})
			}
		})
	}
}

func TestCompareUnparseableFloats(t *testing.T) {
	fields := []fld{{name: "AF", typ: "Float"}}
	for _, value := range []string{"invalid", "1e309", "-1e309"} {
		t.Run(value, func(t *testing.T) {
			for _, tt := range []struct{ ev, iv string }{
				{ev: value, iv: "0.5"},
				{ev: "0.5", iv: value},
				{ev: "-1", iv: value},
			} {
				n, st, err := compareText(t, testRow("gnomad_AF="+tt.ev), testRow("gnomad_AF="+tt.iv), fields, 1.1e-6)
				want := stats{cmp: 1, valErr: 1}
				if err != nil || n != 1 || st == nil || *st != want {
					t.Fatalf("values (%q, %q): compare = (%d, %+v, %v), want (1, %+v, nil)", tt.ev, tt.iv, n, st, err, want)
				}
			}
		})
	}
}

func TestCompareInvalidTolerance(t *testing.T) {
	for _, tol := range []float64{math.NaN(), math.Inf(1), math.Inf(-1), -1, -math.SmallestNonzeroFloat64} {
		n, st, err := compare("", "", nil, tol, 0)
		if err == nil || !strings.Contains(err.Error(), "tolerance must be finite and non-negative") || n != 0 || st != nil {
			t.Fatalf("tolerance %g: compare = (%d, %+v, %v), want a tolerance error before opening files", tol, n, st, err)
		}
	}
}

func TestCompareOverflowAndMinusOne(t *testing.T) {
	fields := []fld{{name: "AF", typ: "Float"}}
	for _, tt := range []struct {
		name   string
		ev, iv string
		want   stats
	}{
		{name: "old positive cutoff", ev: "-1", iv: "1073.7411", want: stats{cmp: 1, hitErr: 1}},
		{name: "old negative cutoff", ev: "-1", iv: "-1073.7411", want: stats{cmp: 1, hitErr: 1}},
		{name: "last safe positive f32", ev: "-1", iv: "1073.74169921875", want: stats{cmp: 1, hitErr: 1}},
		{name: "first overflowing positive f32", ev: "-1", iv: "1073.7418212890625", want: stats{cmp: 1, overflowLoss: 1}},
		{name: "positive f32 rounding", ev: "-1", iv: "1073.7418", want: stats{cmp: 1, overflowLoss: 1}},
		{name: "exact positive i32 ratio", ev: "-1", iv: "1073.7418235", want: stats{cmp: 1, overflowLoss: 1}},
		{name: "last safe negative f32", ev: "-1", iv: "-1073.7418212890625", want: stats{cmp: 1, hitErr: 1}},
		{name: "exact negative i32 ratio", ev: "-1", iv: "-1073.741824", want: stats{cmp: 1, hitErr: 1}},
		{name: "first overflowing negative f32", ev: "-1", iv: "-1073.741943359375", want: stats{cmp: 1, overflowLoss: 1}},
		{name: "positive overflow", ev: "-1", iv: "2000", want: stats{cmp: 1, overflowLoss: 1}},
		{name: "negative overflow", ev: "-1", iv: "-2000", want: stats{cmp: 1, overflowLoss: 1}},
		{name: "largest finite positive overflow", ev: "-1", iv: "1.7976931348623157e308", want: stats{cmp: 1, overflowLoss: 1}},
		{name: "largest finite negative overflow", ev: "-1", iv: "-1.7976931348623157e308", want: stats{cmp: 1, overflowLoss: 1}},
		{name: "true minus one", ev: "-1", iv: "-1", want: stats{cmp: 1, miss: 1}},
		{name: "decimal true minus one", ev: "-1", iv: "-1.0", want: stats{cmp: 1, miss: 1}},
		{name: "exponent true minus one", ev: "-1", iv: "-1e0", want: stats{cmp: 1, miss: 1}},
		{name: "both missing", ev: "-1", iv: ".", want: stats{cmp: 1, miss: 1}},
		{name: "ordinary miss disagreement", ev: "-1", iv: "0.5", want: stats{cmp: 1, hitErr: 1}},
		{name: "other miss marker is not overflow", ev: ".", iv: "2000", want: stats{cmp: 1, hitErr: 1}},
		{name: "reverse miss is not overflow", ev: "2000", iv: ".", want: stats{cmp: 1, hitErr: 1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			n, st, err := compareText(t, testRow("gnomad_AF="+tt.ev), testRow("gnomad_AF="+tt.iv), fields, 1.1e-6)
			if err != nil || n != 1 || st == nil || *st != tt.want {
				t.Fatalf("compare = (%d, %+v, %v), want (1, %+v, nil)", n, st, err, tt.want)
			}
		})
	}
}

func TestCompareRows(t *testing.T) {
	fields := []fld{{name: "AF", typ: "Float"}}
	row := testRow("gnomad_AF=0.5")
	for _, tt := range []struct {
		name    string
		ev, iv  string
		rows    int
		wantErr string
	}{
		{name: "empty"},
		{name: "headers only", ev: "##fileformat=VCFv4.2\n#CHROM\tPOS\n", iv: "#different header\n\n#trailing comment"},
		{name: "echtvar headers versus empty", ev: "#header\n"},
		{name: "immkv headers versus empty", iv: "#header\n"},
		{name: "blank versus empty", ev: "\n\r\n"},
		{name: "empty versus blank", iv: "\n\r\n"},
		{name: "different trailing comments", ev: "#header\n\n" + row + "#end\n", iv: row + "\n#another end\n#last comment", rows: 1},
		{name: "echtvar trailing blanks", ev: row + "\n\n", iv: row, rows: 1},
		{name: "immkv trailing blanks", ev: row, iv: row + "\n\n", rows: 1},
		{name: "comments and blanks between rows", ev: row + "\n#middle\n" + row + "#end\n", iv: "#header\n" + row + row + "\n", rows: 2},
		{name: "no final newline", ev: strings.TrimSuffix(row, "\n"), iv: row, rows: 1},
		{name: "echtvar extra row", ev: row + "#middle\n" + row, iv: row + "\n#end\n", rows: 1, wantErr: "row count mismatch at row 2"},
		{name: "immkv extra row", ev: row + "#end\n", iv: row + "\n#middle\n" + row, rows: 1, wantErr: "row count mismatch at row 2"},
		{name: "echtvar data versus headers", ev: row, iv: "#header\n", wantErr: "row count mismatch at row 1"},
		{name: "immkv data versus headers", ev: "#header\n", iv: row, wantErr: "row count mismatch at row 1"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			n, st, err := compareText(t, tt.ev, tt.iv, fields, 1.1e-6)
			if (tt.wantErr == "" && err != nil) || (tt.wantErr != "" && (err == nil || !strings.Contains(err.Error(), tt.wantErr))) {
				t.Fatalf("error = %v, want %q", err, tt.wantErr)
			}
			want := stats{cmp: tt.rows, floatN: tt.rows}
			if n != tt.rows || st == nil || *st != want {
				t.Fatalf("compare = (%d, %+v, %v), want rows %d and stats %+v", n, st, err, tt.rows, want)
			}
		})
	}
}

func appendOversizedLine(t *testing.T, path string) {
	t.Helper()
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString("1\t101\t.\tA\tC\t.\tPASS\tgnomad_AF="); err != nil {
		t.Fatal(err)
	}
	chunk := strings.Repeat("1", 1<<20)
	for i := 0; i < 33; i++ {
		if _, err := f.WriteString(chunk); err != nil {
			t.Fatal(err)
		}
	}
	if _, err := f.WriteString("\n"); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
}

func TestCompareScannerErrors(t *testing.T) {
	fields := []fld{{name: "AF", typ: "Float"}}
	for _, prefix := range []struct {
		name, text string
		rows       int
	}{
		{name: "before data", text: "#header\n\n"},
		{name: "after data", text: "#header\n" + testRow("gnomad_AF=0.5") + "\n#comment\n", rows: 1},
	} {
		t.Run(prefix.name, func(t *testing.T) {
			for _, tt := range []struct {
				name           string
				evLong, ivLong bool
			}{
				{name: "echtvar", evLong: true},
				{name: "immkv", ivLong: true},
				{name: "both", evLong: true, ivLong: true},
			} {
				t.Run(tt.name, func(t *testing.T) {
					evPath, ivPath := writeCompareFiles(t, prefix.text, prefix.text)
					if tt.evLong {
						appendOversizedLine(t, evPath)
					}
					if tt.ivLong {
						appendOversizedLine(t, ivPath)
					}
					n, st, err := compare(evPath, ivPath, fields, 1.1e-6, 0)
					if !errors.Is(err, bufio.ErrTooLong) {
						t.Fatalf("error = %v, want bufio.ErrTooLong", err)
					}
					path := ivPath
					if tt.evLong {
						path = evPath
					}
					if !strings.Contains(err.Error(), path) {
						t.Fatalf("error %q does not identify %q", err, path)
					}
					want := stats{cmp: prefix.rows, floatN: prefix.rows}
					if n != prefix.rows || st == nil || *st != want {
						t.Fatalf("compare = (%d, %+v, %v), want rows %d and stats %+v", n, st, err, prefix.rows, want)
					}
				})
			}
		})
	}
}

func TestCompareIntegerStringAndFlag(t *testing.T) {
	for _, tt := range []struct {
		name, typ string
		ev, iv    string
		want      stats
	}{
		{name: "integer exact", typ: "Integer", ev: "gnomad_X=01", iv: "gnomad_X=01", want: stats{cmp: 1}},
		{name: "integer different bytes", typ: "Integer", ev: "gnomad_X=01", iv: "gnomad_X=1", want: stats{cmp: 1, valErr: 1}},
		{name: "integer minus one marker", typ: "Integer", ev: "gnomad_X=-1", iv: "gnomad_X=-1", want: stats{cmp: 1, hitErr: 1}},
		{name: "string exact", typ: "String", ev: "gnomad_X=A,B|C", iv: "gnomad_X=A,B|C", want: stats{cmp: 1}},
		{name: "string different bytes", typ: "String", ev: "gnomad_X=A", iv: "gnomad_X=a", want: stats{cmp: 1, valErr: 1}},
		{name: "string NaN remains text", typ: "String", ev: "gnomad_X=NaN", iv: "gnomad_X=NaN", want: stats{cmp: 1}},
		{name: "flag present", typ: "Flag", ev: "gnomad_X", iv: "gnomad_X=NaN", want: stats{cmp: 1}},
		{name: "flag absent", typ: "Flag", ev: ".", iv: ".", want: stats{cmp: 1, miss: 1}},
		{name: "flag echtvar only", typ: "Flag", ev: "gnomad_X=0", iv: ".", want: stats{cmp: 1, hitErr: 1}},
		{name: "flag immkv only", typ: "Flag", ev: ".", iv: "gnomad_X", want: stats{cmp: 1, hitErr: 1}},
	} {
		t.Run(tt.name, func(t *testing.T) {
			fields := []fld{{name: "X", typ: tt.typ, isFlag: tt.typ == "Flag"}}
			n, st, err := compareText(t, testRow(tt.ev), testRow(tt.iv), fields, 1.1e-6)
			if err != nil || n != 1 || st == nil || *st != tt.want {
				t.Fatalf("compare = (%d, %+v, %v), want (1, %+v, nil)", n, st, err, tt.want)
			}
		})
	}
}
