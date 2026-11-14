// build-check measures deterministic sorted construction separately from full
// library reconstruction. The reader verifies every value outside build timing.
package main

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/viktorxia/immkv"
)

func main() {
	mode := flag.String("mode", "build", "build or verify")
	base := flag.String("base", "", "new library file prefix")
	count := flag.Int("n", 200000, "record count")
	pack := flag.Int("pack", 0, "PackTarget; -1 uses dictionary frames")
	flag.Parse()
	if *base == "" || *count < 1 {
		panic("-base and positive -n required")
	}
	if err := run(*mode, *base, *count, *pack); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func fillValue(dst []byte, i int) []byte {
	dst = dst[:0]
	dst = fmt.Appendf(dst, "gene-%08d|AF=0.0001|AN=240000|consequence=missense_variant|", i)
	for len(dst) < 512 {
		dst = append(dst, byte('A'+(i*7+len(dst)*3)%26))
	}
	return dst
}

func run(mode, base string, count, pack int) error {
	var key [8]byte
	value := make([]byte, 0, 512)
	switch mode {
	case "build":
		if err := os.MkdirAll(filepath.Dir(base), 0o755); err != nil {
			return err
		}
		started := time.Now()
		writer, err := immkv.Create(base, immkv.CreateOptions{PackTarget: pack})
		if err != nil {
			return err
		}
		for i := 0; i < count; i++ {
			binary.BigEndian.PutUint64(key[:], uint64(i))
			value = fillValue(value, i)
			if err := writer.AddSorted(key[:], key[:], value); err != nil {
				return err
			}
		}
		if err := writer.Finish(); err != nil {
			return err
		}
		elapsed := time.Since(started)
		return json.NewEncoder(os.Stdout).Encode(map[string]any{
			"records": count, "build_seconds": elapsed.Seconds(),
			"records_per_second": float64(count) / elapsed.Seconds(), "pack_target": pack,
		})
	case "verify":
		db, err := immkv.OpenSimpleKV(base, 0)
		if err != nil {
			return err
		}
		defer db.Close()
		for i := 0; i < count; i++ {
			binary.BigEndian.PutUint64(key[:], uint64(i))
			got, found, err := db.Get(key[:])
			if err != nil {
				return err
			}
			value = fillValue(value, i)
			if !found || !bytes.Equal(got, value) {
				return fmt.Errorf("record %d mismatch", i)
			}
		}
		fmt.Printf("BUILD_VALUES_OK records=%d pack_target=%d\n", count, pack)
		return nil
	default:
		return fmt.Errorf("invalid mode %q", mode)
	}
}
