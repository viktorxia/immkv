// query-only loads and sorts all keys before measuring GetBatch and consuming
// its returned slices; INFO rendering and output compression are excluded.
package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
	"runtime"
	"runtime/pprof"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/biogo/hts/bgzf"
	"github.com/viktorxia/immkv"
)

type options struct {
	query, library, fields, cpuProfile string
	threads, batch, repeat             int
	columns, digest                    bool
}

type passResult struct {
	Hits          uint64  `json:"hits"`
	Batches       uint64  `json:"batches"`
	ReturnedBytes uint64  `json:"returned_bytes"`
	LookupSeconds float64 `json:"lookup_seconds"`
	RawSHA256     string  `json:"raw_sha256,omitempty"`
}

type batchGetter interface {
	GetBatch([][]byte) ([][]byte, error)
}

func main() {
	var o options
	flag.StringVar(&o.query, "query", "", "decomposed query VCF (.gz/.bgz or plain text)")
	flag.StringVar(&o.library, "lib", "", "v003 library directory")
	flag.StringVar(&o.fields, "fields", "", "field TSV")
	flag.IntVar(&o.threads, "threads", 4, "decompression workers")
	flag.IntVar(&o.batch, "batch", 4096, "keys per GetBatch, extending same-position groups")
	flag.IntVar(&o.repeat, "repeat", 1, "complete lookup passes on one handle; setup happens once")
	flag.BoolVar(&o.columns, "columns", true, "pass explicit field selection; false passes Columns=nil")
	flag.BoolVar(&o.digest, "digest", false, "diagnostic raw-record SHA-256 per pass, included in lookup time")
	flag.StringVar(&o.cpuProfile, "cpuprofile", "", "new CPU profile path; profile ONLY the lookup loop")
	flag.Parse()
	if o.query == "" || o.library == "" || o.fields == "" || o.batch < 1 || o.threads < 1 || o.repeat < 1 || flag.NArg() != 0 {
		fmt.Fprintln(os.Stderr, "-query, -lib, -fields and positive -threads, -batch, -repeat are required")
		os.Exit(2)
	}
	if err := run(o); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(o options) error {
	started := time.Now()
	keys, err := loadKeys(o.query)
	if err != nil {
		return err
	}
	if len(keys) == 0 {
		return fmt.Errorf("query contains no buildable keys")
	}
	schema, err := immkv.ReadSchema(o.library)
	if err != nil {
		return err
	}
	fieldData, err := os.ReadFile(o.fields)
	if err != nil {
		return err
	}
	var selected []int
	seen := make(map[int]bool)
	for _, line := range strings.Split(string(fieldData), "\n") {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		name := strings.SplitN(line, "\t", 2)[0]
		index := schema.Index(name)
		if index < 0 || seen[index] {
			return fmt.Errorf("field %q missing from library or duplicated", name)
		}
		seen[index] = true
		selected = append(selected, index)
	}
	if len(selected) == 0 {
		return fmt.Errorf("no selected fields")
	}
	selectedFields := len(selected)
	if !o.columns {
		selected = nil
	}
	db, err := immkv.OpenStreaming(o.library, immkv.StreamOptions{Threads: o.threads, Columns: selected})
	if err != nil {
		return err
	}
	defer db.Close()
	setup := time.Since(started)

	var profile *os.File
	if o.cpuProfile != "" {
		if _, err := os.Lstat(o.cpuProfile); !os.IsNotExist(err) {
			return fmt.Errorf("profile target must not exist: %s", o.cpuProfile)
		}
		profile, err = os.OpenFile(o.cpuProfile+".partial", os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
		if err != nil {
			return err
		}
		defer profile.Close()
		if err := pprof.StartCPUProfile(profile); err != nil {
			return err
		}
	}
	// StartCPUProfile, StopCPUProfile, setup, Close, and JSON output are outside
	// lookup_seconds. Profiling is enabled only while lookupPass executes.
	passes := make([]passResult, 0, o.repeat)
	var elapsed time.Duration
	var hits, returnedBytes, batches uint64
	var lookupErr error
	for iteration := 0; iteration < o.repeat; iteration++ {
		var result passResult
		result, lookupErr = lookupPass(db, keys, o.batch, o.digest)
		if lookupErr != nil {
			break
		}
		elapsed += time.Duration(result.LookupSeconds * float64(time.Second))
		hits += result.Hits
		returnedBytes += result.ReturnedBytes
		batches += result.Batches
		passes = append(passes, result)
	}
	if profile != nil {
		pprof.StopCPUProfile()
		if err := profile.Close(); lookupErr == nil && err != nil {
			lookupErr = err
		}
	}
	if lookupErr != nil {
		return lookupErr
	}
	for i := 1; i < len(passes); i++ {
		a, b := passes[0], passes[i]
		if a.Hits != b.Hits || a.Batches != b.Batches || a.ReturnedBytes != b.ReturnedBytes || a.RawSHA256 != b.RawSHA256 {
			return fmt.Errorf("lookup pass %d differs from pass 1", i+1)
		}
	}
	if profile != nil {
		// Link refuses replacement, then remove the temporary name on success.
		if err := os.Link(o.cpuProfile+".partial", o.cpuProfile); err != nil {
			return err
		}
		if err := os.Remove(o.cpuProfile + ".partial"); err != nil {
			return err
		}
	}
	return json.NewEncoder(os.Stdout).Encode(struct {
		Keys            int          `json:"keys"`
		TotalKeys       uint64       `json:"total_keys"`
		Repeat          int          `json:"repeat"`
		Hits            uint64       `json:"hits"`
		Batches         uint64       `json:"batches"`
		ReturnedBytes   uint64       `json:"returned_bytes"`
		ExplicitColumns bool         `json:"explicit_columns"`
		SelectedFields  int          `json:"selected_fields"`
		SchemaFields    int          `json:"schema_fields"`
		Threads         int          `json:"threads"`
		Batch           int          `json:"batch"`
		GOMAXPROCS      int          `json:"gomaxprocs"`
		Profiled        bool         `json:"profiled"`
		DigestEnabled   bool         `json:"digest_enabled"`
		SetupSeconds    float64      `json:"setup_seconds"`
		LookupSeconds   float64      `json:"lookup_seconds"`
		KeysPerSecond   float64      `json:"keys_per_second"`
		Passes          []passResult `json:"passes"`
	}{len(keys), uint64(len(keys)) * uint64(o.repeat), o.repeat, hits, batches, returnedBytes,
		o.columns, selectedFields, len(schema), o.threads, o.batch, runtime.GOMAXPROCS(0),
		profile != nil, o.digest, setup.Seconds(), elapsed.Seconds(),
		float64(len(keys)) * float64(o.repeat) / elapsed.Seconds(), passes})
}

func lookupPass(db batchGetter, keys [][]byte, batch int, digest bool) (passResult, error) {
	var result passResult
	var h hash.Hash
	if digest {
		h = sha256.New()
	}
	started := time.Now()
	for start := 0; start < len(keys); {
		end := min(start+batch, len(keys))
		// Preserve same-position groups, as the annotation harness does.
		for end < len(keys) && bytes.Equal(keys[end-1][:5], keys[end][:5]) {
			end++
		}
		values, err := db.GetBatch(keys[start:end])
		if err != nil {
			return result, err
		}
		if len(values) != end-start {
			return result, fmt.Errorf("result length mismatch")
		}
		for i, value := range values {
			if value != nil {
				result.Hits++
				result.ReturnedBytes += uint64(len(value))
			}
			if h != nil {
				// Unambiguous key length, value length, presence, key, and bytes.
				// This hashes actual records, without assuming projection matches
				// across versions; old may return full records for a subset.
				var header [17]byte
				key := keys[start+i]
				binary.LittleEndian.PutUint64(header[:8], uint64(len(key)))
				binary.LittleEndian.PutUint64(header[8:16], uint64(len(value)))
				if value != nil {
					header[16] = 1
				}
				h.Write(header[:])
				h.Write(key)
				h.Write(value)
			}
		}
		result.Batches++
		start = end
	}
	result.LookupSeconds = time.Since(started).Seconds()
	if h != nil {
		result.RawSHA256 = hex.EncodeToString(h.Sum(nil))
	}
	return result, nil
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
				arena = make([]byte, 0, max(8<<20, len(key)))
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
