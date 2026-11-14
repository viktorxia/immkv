package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/viktorxia/immkv"
)

type fakeBatch struct {
	calls [][]byte
}

func (f *fakeBatch) GetBatch(keys [][]byte) ([][]byte, error) {
	f.calls = append(f.calls, append([]byte(nil), keys[0]...))
	values := make([][]byte, len(keys))
	for i, key := range keys {
		if bytes.HasSuffix(key, []byte("T")) {
			values[i] = []byte("value")
		}
	}
	return values, nil
}

func TestLookupSamePositionAndRepeatDigest(t *testing.T) {
	keys := [][]byte{
		immkv.VarKey(0, 1, "A", "C"), immkv.VarKey(0, 1, "A", "T"),
		immkv.VarKey(0, 2, "A", "T"),
	}
	getter := new(fakeBatch)
	first, err := lookupPass(getter, keys, 1, true)
	if err != nil {
		t.Fatal(err)
	}
	if first.Hits != 2 || first.ReturnedBytes != 10 || first.Batches != 2 || first.RawSHA256 == "" {
		t.Fatalf("unexpected result: %+v", first)
	}
	if !bytes.Equal(getter.calls[1], keys[2]) {
		t.Fatal("same-position keys split across batches")
	}
	second, err := lookupPass(getter, keys, 1, true)
	if err != nil || second.RawSHA256 != first.RawSHA256 {
		t.Fatalf("repeat changed digest: %+v, %v", second, err)
	}
	plain, err := lookupPass(getter, keys, 2, false)
	if err != nil || plain.RawSHA256 != "" || plain.Hits != first.Hits || plain.ReturnedBytes != first.ReturnedBytes {
		t.Fatalf("digest altered lookup results: %+v, %v", plain, err)
	}
}

type shortBatch struct{}

func (shortBatch) GetBatch([][]byte) ([][]byte, error) { return nil, nil }

func TestLookupRejectsMissingResults(t *testing.T) {
	if _, err := lookupPass(shortBatch{}, [][]byte{immkv.VarKey(0, 1, "A", "T")}, 1, false); err == nil {
		t.Fatal("accepted a result count mismatch")
	}
}

func TestLoadKeysSortAndFilter(t *testing.T) {
	path := filepath.Join(t.TempDir(), "query.vcf")
	data := "#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO\n"
	for _, row := range []string{"chr22\t8\t.\tA\tT,C\t.\tPASS\t.", "chr1\t9\t.\tC\tA\t.\tPASS\t.", "chr1\t1\t.\tA\t<DEL>,*,.,A]2:1]\t.\tPASS\t.", "chrUn\t1\t.\tA\tT\t.\tPASS\t."} {
		data += row + "\n"
	}
	if err := os.WriteFile(path, []byte(data), 0o600); err != nil {
		t.Fatal(err)
	}
	keys, err := loadKeys(path)
	if err != nil {
		t.Fatal(err)
	}
	want := [][]byte{immkv.VarKey(0, 9, "C", "A"), immkv.VarKey(21, 8, "A", "C"), immkv.VarKey(21, 8, "A", "T")}
	if fmt.Sprint(keys) != fmt.Sprint(want) {
		t.Fatalf("keys %v; want %v", keys, want)
	}
}
