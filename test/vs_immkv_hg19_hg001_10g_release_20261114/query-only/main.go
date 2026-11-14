// query-only consumes every query key but discards returned records without
// decoding them. Loading and sorting the query happen before the lookup timer.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/biogo/hts/bgzf"
	"github.com/viktorxia/immkv"
)

func main() {
	query := flag.String("query", "", "decomposed query VCF")
	library := flag.String("lib", "", "v003 library directory")
	fields := flag.String("fields", "", "field TSV")
	threads := flag.Int("threads", 4, "decompression workers")
	batch := flag.Int("batch", 4096, "keys per GetBatch")
	columns := flag.Bool("columns", true, "pass explicit field selection")
	flag.Parse()
	if *query == "" || *library == "" || *fields == "" || *batch < 1 {
		panic("-query, -lib, -fields and a positive -batch are required")
	}
	if err := run(*query, *library, *fields, *threads, *batch, *columns); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(query, library, fields string, threads, batch int, columns bool) error {
	started := time.Now()
	keys, err := loadKeys(query)
	if err != nil {
		return err
	}
	schema, err := immkv.ReadSchema(library)
	if err != nil {
		return err
	}
	fieldData, err := os.ReadFile(fields)
	if err != nil {
		return err
	}
	var selected []int
	for _, line := range strings.Split(string(fieldData), "\n") {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		name := strings.SplitN(line, "\t", 2)[0]
		index := schema.Index(name)
		if index < 0 {
			return fmt.Errorf("field %q missing from library", name)
		}
		selected = append(selected, index)
	}
	if len(selected) == 0 {
		return fmt.Errorf("no selected fields")
	}
	if !columns {
		selected = nil
	}
	db, err := immkv.OpenStreaming(library, immkv.StreamOptions{Threads: threads, Columns: selected})
	if err != nil {
		return err
	}
	defer db.Close()
	setup := time.Since(started)
	queryStart := time.Now()
	var hits, returnedBytes, batches uint64
	for start := 0; start < len(keys); {
		end := min(start+batch, len(keys))
		// Preserve same-position groups, as the annotation harness does.
		for end < len(keys) && bytes.Equal(keys[end-1][:5], keys[end][:5]) {
			end++
		}
		values, err := db.GetBatch(keys[start:end])
		if err != nil {
			return err
		}
		if len(values) != end-start {
			return fmt.Errorf("result length mismatch")
		}
		for _, value := range values {
			if value != nil {
				hits++
				returnedBytes += uint64(len(value))
			}
		}
		batches++
		start = end
	}
	elapsed := time.Since(queryStart)
	return json.NewEncoder(os.Stdout).Encode(struct {
		Keys            int     `json:"keys"`
		Hits            uint64  `json:"hits"`
		Batches         uint64  `json:"batches"`
		ReturnedBytes   uint64  `json:"returned_bytes"`
		ExplicitColumns bool    `json:"explicit_columns"`
		SetupSeconds    float64 `json:"setup_seconds"`
		LookupSeconds   float64 `json:"lookup_seconds"`
		KeysPerSecond   float64 `json:"keys_per_second"`
	}{len(keys), hits, batches, returnedBytes, columns, setup.Seconds(), elapsed.Seconds(), float64(len(keys)) / elapsed.Seconds()})
}

func loadKeys(path string) ([][]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var input io.Reader = file
	if strings.HasSuffix(path, ".gz") || strings.HasSuffix(path, ".bgz") {
		reader, err := bgzf.NewReader(file, 1)
		if err != nil {
			return nil, err
		}
		defer reader.Close()
		input = reader
	}
	scanner := bufio.NewScanner(input)
	scanner.Buffer(make([]byte, 1<<20), 64<<20)
	var keys [][]byte
	var arena []byte
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		parts := bytes.SplitN(line, []byte("\t"), 6)
		if len(parts) < 6 {
			return nil, fmt.Errorf("malformed VCF row")
		}
		chrom, ok := immkv.ChromID(string(parts[0]))
		if !ok || bytes.Equal(parts[3], []byte(".")) {
			continue
		}
		pos, err := strconv.ParseUint(string(parts[1]), 10, 32)
		if err != nil {
			return nil, err
		}
		for _, alt := range bytes.Split(parts[4], []byte(",")) {
			if len(alt) == 0 || alt[0] == '<' || alt[0] == '*' || alt[0] == '.' || bytes.ContainsAny(alt, "[]") {
				continue
			}
			key := immkv.VarKey(chrom, uint32(pos), string(parts[3]), string(alt))
			if cap(arena)-len(arena) < len(key) {
				arena = make([]byte, 0, 8<<20)
			}
			start := len(arena)
			arena = append(arena, key...)
			keys = append(keys, arena[start:len(arena):len(arena)])
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	sort.Slice(keys, func(i, j int) bool { return bytes.Compare(keys[i], keys[j]) < 0 })
	return keys, nil
}
