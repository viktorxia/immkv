// pointlookup measures immkv's OpenEager random point-lookup latency and
// throughput on a library built from a VCF. Keys are reservoir-sampled from a
// decomposed query VCF, so the sample is spread across the genome; misses are
// built by appending 0xFF to sampled keys (the trailing byte cannot occur in a
// real ALT). One process, -threads concurrent readers, optional decompression
// cache.
//
// Usage:
//
//	pointlookup -lib <immkvDir> -query <decomposed.vcf.gz> -n 100000 \
//	            -threads 4 -cache 512 -warmup 10000
//
// Output: one JSON object per phase (hits, misses) with count, QPS, mean and
// p50/p99/p999 latency in nanoseconds. Timer overhead is included in the
// per-call latencies; QPS is the primary throughput figure.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/biogo/hts/bgzf"
	"github.com/viktorxia/immkv"
)

func main() {
	lib := flag.String("lib", "", "immkv library directory")
	query := flag.String("query", "", "decomposed query VCF (.vcf or .vcf.gz)")
	n := flag.Int("n", 100000, "number of sampled keys")
	threads := flag.Int("threads", 1, "concurrent reader goroutines")
	cache := flag.Int("cache", 512, "OpenEager cacheMB")
	warmup := flag.Int("warmup", 10000, "warmup lookups per phase (0 = none)")
	seed := flag.Int64("seed", 1, "sampling/shuffle seed")
	flag.Parse()
	if *lib == "" || *query == "" {
		flag.Usage()
		os.Exit(2)
	}
	if err := run(*lib, *query, *n, *threads, *cache, *warmup, *seed); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func run(lib, query string, n, threads, cache, warmup int, seed int64) error {
	hits, err := sampleKeys(query, n, seed)
	if err != nil {
		return err
	}
	if len(hits) == 0 {
		return fmt.Errorf("no buildable query keys in %s", query)
	}
	misses := make([][]byte, len(hits))
	for i, k := range hits {
		misses[i] = append(bytes.Clone(k), 0xFF)
	}
	db, err := immkv.OpenEager(lib, cache)
	if err != nil {
		return err
	}
	defer db.Close()

	phase := func(name string, keys [][]byte, wantHit bool) error {
		if warmup > 0 {
			doLookups(db, keys, nThreads(threads), min(warmup, len(keys)), wantHit)
		}
		start := time.Now()
		lats, hitsN := doLookups(db, keys, nThreads(threads), len(keys), wantHit)
		elapsed := time.Since(start)
		sort.Slice(lats, func(i, j int) bool { return lats[i] < lats[j] })
		res := result{
			Phase: name, Count: len(lats), Hits: hitsN, Threads: threads,
			ElapsedSec: elapsed.Seconds(),
			QPS:        float64(len(lats)) / elapsed.Seconds(),
			MeanNS:     mean(lats),
			P50NS:      pct(lats, 0.50),
			P99NS:      pct(lats, 0.99),
			P999NS:     pct(lats, 0.999),
		}
		b, _ := json.Marshal(res)
		fmt.Println(string(b))
		return nil
	}
	if err := phase("sampled", hits, true); err != nil {
		return err
	}
	return phase("misses", misses, false)
}

type result struct {
	Phase      string  `json:"phase"`
	Count      int     `json:"count"`
	Hits       int     `json:"hits"`
	Threads    int     `json:"threads"`
	ElapsedSec float64 `json:"elapsed_sec"`
	QPS        float64 `json:"qps"`
	MeanNS     float64 `json:"mean_ns"`
	P50NS      int64   `json:"p50_ns"`
	P99NS      int64   `json:"p99_ns"`
	P999NS     int64   `json:"p999_ns"`
}

func nThreads(t int) int {
	if t < 1 {
		return 1
	}
	return t
}

// doLookups runs count lookups over keys with t goroutines and returns the
// pooled per-call latencies and hit count. Keys are interleaved across
// goroutines so each reader touches the whole sampled range. wantHit=false
// keys are constructed to be misses; a hit there is a build error and panics.
func doLookups(db *immkv.DB, keys [][]byte, t, count int, wantHit bool) ([]int64, int) {
	if count > len(keys) {
		count = len(keys)
	}
	lats := make([][]int64, t)
	hits := make([]int, t)
	var wg sync.WaitGroup
	for g := 0; g < t; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			own := make([]int64, 0, (count+t-1)/t)
			for i := g; i < count; i += t {
				k := keys[i]
				t0 := time.Now()
				v, ok := db.Get(k)
				d := time.Since(t0)
				if ok {
					hits[g]++
					if !wantHit {
						panic(fmt.Sprintf("miss key %x unexpectedly hit", k))
					}
				}
				_ = v
				own = append(own, int64(d))
			}
			lats[g] = own
		}(g)
	}
	wg.Wait()
	var all []int64
	total := 0
	for i, l := range lats {
		all = append(all, l...)
		total += hits[i]
	}
	return all, total
}

func mean(v []int64) float64 {
	if len(v) == 0 {
		return 0
	}
	var s int64
	for _, x := range v {
		s += x
	}
	return float64(s) / float64(len(v))
}

func pct(sorted []int64, p float64) int64 {
	if len(sorted) == 0 {
		return 0
	}
	i := int(p * float64(len(sorted)))
	if i >= len(sorted) {
		i = len(sorted) - 1
	}
	return sorted[i]
}

// sampleKeys reservoir-samples buildable VarKeys from the query VCF.
func sampleKeys(path string, n int, seed int64) ([][]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var r *bufio.Reader
	if strings.HasSuffix(path, ".gz") || strings.HasSuffix(path, ".bgz") {
		br, err := bgzf.NewReader(f, 0)
		if err != nil {
			return nil, err
		}
		defer br.Close()
		r = bufio.NewReaderSize(br, 1<<20)
	} else {
		r = bufio.NewReaderSize(f, 1<<20)
	}
	rnd := newRand(seed)
	var sample [][]byte
	var seen int
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 1<<20), 64<<20)
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) == 0 || line[0] == '#' {
			continue
		}
		var tabs [8]int
		nt := 0
		for i := 0; i < len(line) && nt < 8; i++ {
			if line[i] == '\t' {
				tabs[nt] = i
				nt++
			}
		}
		if nt < 5 {
			continue
		}
		chrom, ok := immkv.ChromID(string(line[:tabs[0]]))
		if !ok {
			continue
		}
		pos, err := strconv.ParseUint(string(line[tabs[0]+1:tabs[1]]), 10, 32)
		if err != nil {
			continue
		}
		ref := string(line[tabs[2]+1 : tabs[3]])
		if ref == "." {
			continue
		}
		for _, alt := range bytes.Split(line[tabs[3]+1:tabs[4]], []byte(",")) {
			if len(alt) == 0 || alt[0] == '<' || alt[0] == '*' || alt[0] == '.' ||
				bytes.ContainsAny(alt, "[]") {
				continue
			}
			k := immkv.VarKey(chrom, uint32(pos), ref, string(alt))
			if len(sample) < n {
				sample = append(sample, k)
			} else {
				j := rnd.Intn(seen + 1)
				if j < n {
					sample[j] = k
				}
			}
			seen++
		}
	}
	if err := sc.Err(); err != nil {
		return nil, err
	}
	// Shuffle so consecutive lookups are not position-ordered.
	for i := len(sample) - 1; i > 0; i-- {
		j := rnd.Intn(i + 1)
		sample[i], sample[j] = sample[j], sample[i]
	}
	return sample, nil
}

// newRand 是 math/rand 的最小封装，避免测试工具依赖全局种子。
type smallRand struct{ state uint64 }

func newRand(seed int64) *smallRand { return &smallRand{state: uint64(seed)} }

func (r *smallRand) next() uint64 {
	// splitmix64
	r.state += 0x9e3779b97f4a7c15
	z := r.state
	z = (z ^ (z >> 30)) * 0xbf58476d1ce4e5b9
	z = (z ^ (z >> 27)) * 0x94d049bb133111eb
	return z ^ (z >> 31)
}

func (r *smallRand) Intn(n int) int {
	if n <= 0 {
		return 0
	}
	return int(r.next() % uint64(n))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
