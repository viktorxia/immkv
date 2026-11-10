// annotate: end-to-end annotation of a VCF query against an immkv library
// built from gnomAD v2.1.1 sites (see buildfull / buildnarrow for the two
// library flavors).
//
// Usage:
//
//	annotate -in query.vcf[.gz] -lib <immkvLibDir> -out out.vcf \
//	         -fields <fields.tsv> [-prefix gnomad_] [-batch 4096] [-threads 4]
//
// -fields is the whitelist AND the output order: a TSV of
// "name<Tab>Type<Tab>Number" lines (see genfields). Only listed fields are
// emitted, so the same program serves the 17-field and 540-field controlled
// comparisons by swapping the file.
//
// Matching semantics: one VarKey per ALT element of each query row
// (chrom/pos/ref/alt); a row with several ALTs can hit several library
// entries. Query rows should be decomposed (bcftools norm -m -both) to match
// the libraries, which are also built decomposed.
//
// Output semantics (the appended INFO segment, <prefix>-prefixed):
//   - Row bytes are passed through verbatim except for the appended segment.
//     Values are emitted raw, byte-for-byte as stored in the library -- no
//     float reformatting (the legacy vcfanno-comparison harness re-ordered
//     floats like vcfgo; this tool deliberately does not).
//   - Single hit: only the fields present in that library value.
//   - Multiple hits: every whitelist field; Flag fields render as "name="
//     (comma-joined empty values); Number=A fields expand across the query
//     row's ALT positions (unhit positions get "."); other fields join the
//     hit values with commas in ALT byte order.
//   - Miss: row emitted unchanged.
//
// Concurrency: three-stage pipeline. The main thread reads lines, builds keys
// and issues StreamDB.GetBatch; values are copied into a per-window arena
// (engine values are only valid until the next query) and handed to -threads
// formatting workers; a collector writes windows in order. jobs/resCh channel
// capacities provide backpressure: ~3x threads windows in flight.
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"github.com/biogo/hts/bgzf"
	"github.com/viktorxia/immkv"
)

type field struct {
	name      string
	isFlag    bool
	isNumberA bool
}

func main() {
	in := flag.String("in", "", "query VCF (.vcf or .vcf.gz, decomposed)")
	libDir := flag.String("lib", "", "immkv library directory")
	out := flag.String("out", "", "output plain-text VCF")
	fieldsPath := flag.String("fields", "", "field whitelist TSV from genfields (name/Type/Number)")
	prefix := flag.String("prefix", "gnomad_", "prefix for appended INFO keys")
	batch := flag.Int("batch", 4096, "rows per GetBatch window")
	threads := flag.Int("threads", 4, "format worker threads == immkv Options.Threads")
	flag.Parse()
	if *in == "" || *libDir == "" || *out == "" || *fieldsPath == "" {
		flag.Usage()
		os.Exit(2)
	}
	if err := run(*in, *libDir, *out, *fieldsPath, *prefix, *batch, *threads); err != nil {
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
			return nil, fmt.Errorf("bad fields line %q (want name\\tType\\tNumber)", ln)
		}
		out = append(out, field{
			name:      string(cols[0]),
			isFlag:    string(cols[1]) == "Flag",
			isNumberA: string(cols[2]) == "A",
		})
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("no fields in %s", path)
	}
	return out, nil
}

// resolveColumns maps the whitelist to base.manifest column numbers so the
// engine can read only the needed columns. Without a manifest (old library)
// it returns nil and the engine keeps the record-order contract. Whitelist
// names absent from the schema are simply not selected.
func resolveColumns(libDir string, fields []field) ([]int, error) {
	paths, err := filepath.Glob(filepath.Join(libDir, "chr*.manifest"))
	if err != nil {
		return nil, err
	}
	if len(paths) == 0 {
		return nil, nil
	}
	raw, err := os.ReadFile(paths[0])
	if err != nil {
		return nil, err
	}
	var m struct {
		Fields []struct {
			Name   string `json:"name"`
			Column int    `json:"column"`
		} `json:"fields"`
	}
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("%s: %w", paths[0], err)
	}
	want := make(map[string]bool, len(fields))
	for _, fd := range fields {
		want[fd.name] = true
	}
	var cols []int
	for _, f := range m.Fields {
		if want[f.Name] {
			cols = append(cols, f.Column)
		}
	}
	if len(cols) == 0 {
		return nil, nil
	}
	return cols, nil
}

// row is one query line awaiting formatting.
type row struct {
	line    []byte // whole line, no newline
	infoOff int    // INFO column start
	infoEnd int    // INFO column end ('\t' or end of line)
	refOff  int
	refEnd  int
	altOff  int
	altEnd  int
	chrom   uint8
	pos     uint32
}

func run(in, libDir, out, fieldsPath, prefix string, batch, threads int) error {
	fields, err := loadFields(fieldsPath)
	if err != nil {
		return err
	}
	idxOf := make(map[string]int, len(fields))
	for i, fd := range fields {
		idxOf[fd.name] = i
	}
	cols, err := resolveColumns(libDir, fields)
	if err != nil {
		return err
	}

	start := time.Now()
	f, err := os.Open(in)
	if err != nil {
		return err
	}
	defer f.Close()
	var r io.Reader = bufio.NewReaderSize(f, 1<<20)
	if strings.HasSuffix(in, ".gz") {
		br, err := bgzf.NewReader(f, 0)
		if err != nil {
			return err
		}
		defer br.Close()
		r = br
	}
	sdb, err := immkv.OpenStreaming(libDir, immkv.Options{Threads: threads, Columns: cols})
	if err != nil {
		return err
	}
	defer sdb.Close()

	var total, hits atomic.Uint64
	var tKey, tQuery, tFmt atomic.Int64
	read0 := readBytes()

	type window struct {
		idx       int
		rows      []row
		byFlat    [][]byte
		spanStart []int
		spanN     []int
		header    []byte // only the first window carries the header
	}
	type result struct {
		idx int
		buf []byte
	}
	jobs := make(chan window, threads)
	resCh := make(chan result, threads)
	bufPool := make(chan []byte, threads*3)
	getBuf := func(need int) []byte {
		select {
		case b := <-bufPool:
			if cap(b) >= need {
				return b[:0]
			}
		default:
		}
		return make([]byte, 0, need)
	}
	var wgp sync.WaitGroup
	for w := 0; w < threads; w++ {
		wgp.Add(1)
		go func() {
			defer wgp.Done()
			sc := &fmtScratch{}
			for jw := range jobs {
				t0 := time.Now()
				var need int
				for i := range jw.rows {
					need += len(jw.rows[i].line) + 2
				}
				for _, v := range jw.byFlat {
					need += len(v) + len(v)/8 + 16
				}
				need += len(jw.header)
				buf := bytes.NewBuffer(getBuf(need))
				buf.Write(jw.header)
				h := formatWindow(buf, jw.rows, jw.byFlat, jw.spanStart, jw.spanN, fields, idxOf, prefix, sc)
				hits.Add(h)
				tFmt.Add(int64(time.Since(t0)))
				b := buf.Bytes()
				if len(b) == 0 {
					select {
					case bufPool <- b[:0]:
					default:
					}
					continue
				}
				resCh <- result{jw.idx, b}
			}
		}()
	}
	w, err := os.Create(out)
	if err != nil {
		return err
	}
	bw := bufio.NewWriterSize(w, 4<<20)
	collectorDone := make(chan error, 1)
	go func() {
		var werr error
		pending := map[int][]byte{}
		next := 0
		for r := range resCh {
			pending[r.idx] = r.buf
			for {
				b, ok := pending[next]
				if !ok {
					break
				}
				delete(pending, next)
				if werr == nil {
					if _, werr = bw.Write(b); werr != nil {
						fmt.Fprintln(os.Stderr, "write failed:", werr)
					}
				}
				select {
				case bufPool <- b[:0]:
				default:
				}
				next++
			}
		}
		if werr == nil {
			werr = bw.Flush()
		}
		if cerr := w.Close(); werr == nil {
			werr = cerr
		}
		collectorDone <- werr
	}()

	rows := make([]row, 0, batch)
	var flat, sorted [][]byte // main-thread scratch
	var order []int
	var header bytes.Buffer
	winIdx := 0
	headerQueued := false
	queueHeader := func() {
		if headerQueued {
			return
		}
		headerQueued = true
		if header.Len() == 0 {
			return
		}
		jobs <- window{idx: winIdx, header: append([]byte(nil), header.Bytes()...)}
		winIdx++
	}
	emitWindow := func() error {
		if len(rows) == 0 {
			return nil
		}
		queueHeader()
		t0 := time.Now()
		flat = flat[:0]
		spanStart := make([]int, len(rows))
		spanN := make([]int, len(rows))
		for i := range rows {
			ref := string(rows[i].line[rows[i].refOff:rows[i].refEnd])
			spanStart[i] = len(flat)
			for _, a := range bytes.Split(rows[i].line[rows[i].altOff:rows[i].altEnd], []byte(",")) {
				flat = append(flat, immkv.VarKey(rows[i].chrom, rows[i].pos, ref, string(a)))
				spanN[i]++
			}
		}
		if cap(order) < len(flat) {
			order = make([]int, len(flat))
		}
		order = order[:len(flat)]
		for i := range order {
			order[i] = i
		}
		sort.Slice(order, func(a, b int) bool { return bytes.Compare(flat[order[a]], flat[order[b]]) < 0 })
		if cap(sorted) < len(flat) {
			sorted = make([][]byte, len(flat))
		}
		sorted = sorted[:len(flat)]
		for i, o := range order {
			sorted[i] = flat[o]
		}
		tKey.Add(int64(time.Since(t0)))
		t0 = time.Now()
		vals, err := sdb.GetBatch(sorted)
		if err != nil {
			return fmt.Errorf("GetBatch: %w", err)
		}
		tQuery.Add(int64(time.Since(t0)))
		// Engine values are only valid until the next query: copy the whole
		// window into an owned arena before handing it to workers.
		var need int
		for _, v := range vals {
			need += len(v)
		}
		arena := make([]byte, need)
		byFlat := make([][]byte, len(flat))
		off := 0
		for i, o := range order {
			v := vals[i]
			if len(v) > 0 {
				copy(arena[off:], v)
				byFlat[o] = arena[off : off+len(v)]
				off += len(v)
			}
		}
		total.Add(uint64(len(rows)))
		jobs <- window{
			idx:       winIdx,
			rows:      rows,
			byFlat:    byFlat,
			spanStart: spanStart,
			spanN:     spanN,
		}
		winIdx++
		rows = make([]row, 0, batch) // previous window now owned by a worker
		return nil
	}

	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 64*1024), 16*1024*1024)
	for sc.Scan() {
		line := sc.Bytes()
		if len(line) > 0 && line[0] == '#' {
			header.Write(line)
			header.WriteByte('\n')
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
		if nt < 7 {
			return fmt.Errorf("malformed data line: %q", line[:min(80, len(line))])
		}
		c, ok := immkv.ChromID(string(line[:tabs[0]]))
		if !ok {
			c = 0 // unknown contig: no keys, line passes through unchanged
		}
		pos, err := strconv.ParseUint(string(line[tabs[0]+1:tabs[1]]), 10, 32)
		if err != nil {
			return fmt.Errorf("POS: %w", err)
		}
		infoEnd := len(line)
		if nt == 8 {
			infoEnd = tabs[7]
		}
		// Split windows only where position strictly advances, so rows at the
		// same site never span windows and batch keys stay sorted.
		if len(rows) >= batch {
			last := &rows[len(rows)-1]
			if last.chrom != c || last.pos < uint32(pos) {
				if err := emitWindow(); err != nil {
					return err
				}
			}
		}
		rows = append(rows, row{
			line:    append([]byte(nil), line...),
			infoOff: tabs[6] + 1,
			infoEnd: infoEnd,
			refOff:  tabs[2] + 1,
			refEnd:  tabs[3],
			altOff:  tabs[3] + 1,
			altEnd:  altEnd(nt, tabs, infoEnd),
			chrom:   c,
			pos:     uint32(pos),
		})
	}
	if err := sc.Err(); err != nil {
		return err
	}
	if err := emitWindow(); err != nil {
		return err
	}
	queueHeader()
	close(jobs)
	wgp.Wait()
	close(resCh)
	if err := <-collectorDone; err != nil {
		return err
	}
	fmt.Printf("rows %d, hits %d (%.2f%%), read %.1fGiB; key %.2fs / query %.2fs / format %.2fs (%d threads); wall %s\n",
		total.Load(), hits.Load(), float64(hits.Load())*100/float64(total.Load()), float64(readBytes()-read0)/(1<<30),
		float64(tKey.Load())/1e9, float64(tQuery.Load())/1e9, float64(tFmt.Load())/1e9, threads,
		time.Since(start).Round(time.Millisecond))
	return nil
}

// formatWindow formats one window's rows into buf and returns the hit count.
func formatWindow(buf *bytes.Buffer, rows []row, byFlat [][]byte, spanStart, spanN []int, fields []field, idxOf map[string]int, prefix string, sc *fmtScratch) uint64 {
	var hits uint64
	for i := range rows {
		rw := &rows[i]
		buf.Write(rw.line[:rw.infoOff])
		info := rw.line[rw.infoOff:rw.infoEnd]
		sc.perAlt = sc.perAlt[:0]
		for e := 0; e < spanN[i]; e++ {
			sc.perAlt = append(sc.perAlt, byFlat[spanStart[i]+e])
		}
		has := false
		for _, v := range sc.perAlt {
			if len(v) > 0 {
				has = true
				break
			}
		}
		if has {
			hits++
			// A bare "." INFO placeholder is dropped before appending, like
			// vcfanno; anything else is kept and ';' terminates it.
			if !(len(info) == 1 && info[0] == '.') {
				buf.Write(info)
				buf.WriteByte(';')
			}
			writeFields(buf, sc.perAlt, fields, idxOf, rw.line[rw.altOff:rw.altEnd], prefix, sc)
		} else {
			buf.Write(info)
		}
		buf.Write(rw.line[rw.infoEnd:])
		buf.WriteByte('\n')
	}
	return hits
}

// altEnd returns the end offset of the ALT column.
func altEnd(nt int, tabs [8]int, infoEnd int) int {
	if nt >= 5 {
		return tabs[4]
	}
	return infoEnd
}

type spanT struct{ off, end int }

// fmtScratch is per-worker scratch reused across rows and windows. Without
// reuse, every row would allocate a [ALT][nFields] span array -- tens of GiB
// of garbage over a full query, and GC assist slows the query stage several
// times over.
type fmtScratch struct {
	found  [][]spanT // [ALT element][field]
	idxs   []int
	elems  [][]byte
	perAlt [][]byte
}

func (s *fmtScratch) resetFound(k, nFields int) {
	if len(s.found) < k {
		s.found = append(s.found, make([][]spanT, k-len(s.found))...)
	}
	s.found = s.found[:k]
	for e := 0; e < k; e++ {
		if len(s.found[e]) != nFields {
			s.found[e] = make([]spanT, nFields)
		}
		fs := s.found[e]
		for i := range fs {
			fs[i].off = -1
		}
	}
}

// writeFields appends the hit annotation segment for one row, following the
// merge rules described in the header comment. See also the note on raw value
// passthrough there.
func writeFields(w io.Writer, perAlt [][]byte, fields []field, idxOf map[string]int, altCol []byte, prefix string, sc *fmtScratch) {
	k := len(perAlt)
	m := 0 // hit elements
	for _, v := range perAlt {
		if len(v) > 0 {
			m++
		}
	}
	all := m >= 2
	sc.resetFound(k, len(fields))
	for e, info := range perAlt {
		if len(info) == 0 {
			continue
		}
		fs := sc.found[e]
		rest := info
		off := 0
		for len(rest) > 0 {
			var seg []byte
			if j := bytes.IndexByte(rest, ';'); j >= 0 {
				seg = rest[:j]
				rest = rest[j+1:]
			} else {
				seg = rest
				rest = nil
			}
			eq := bytes.IndexByte(seg, '=')
			key := seg
			if eq >= 0 {
				key = seg[:eq]
			}
			if idx, ok := idxOf[string(key)]; ok && fs[idx].off < 0 {
				fs[idx] = spanT{off, off + len(seg)}
			}
			off += len(seg) + 1
		}
	}
	value := func(fd field, i, e int) []byte {
		seg := perAlt[e][sc.found[e][i].off:sc.found[e][i].end]
		eq := bytes.IndexByte(seg, '=')
		return seg[eq+1:] // raw bytes as stored, no reformatting
	}
	// Multi-hit join order is the ALT byte order of the hitting records
	// (matching how a tabix stream meets same-site records).
	sc.elems = append(sc.elems[:0], bytes.Split(altCol, []byte(","))...)
	byByte := func(idxs []int) []int {
		out := append(sc.idxs[:0], idxs...)
		sort.Slice(out, func(a, b int) bool { return bytes.Compare(sc.elems[out[a]], sc.elems[out[b]]) < 0 })
		return out
	}
	first := true
	var idxs []int
	for i, fd := range fields {
		idxs = idxs[:0]
		for e := range sc.found {
			if sc.found[e][i].off >= 0 {
				idxs = append(idxs, e)
			}
		}
		if len(idxs) == 0 && !(all && fd.isFlag) {
			continue
		}
		if !first {
			io.WriteString(w, ";")
		}
		first = false
		io.WriteString(w, prefix)
		io.WriteString(w, fd.name)
		if fd.isFlag {
			if all {
				io.WriteString(w, "=")
				for j := 1; j < m; j++ {
					io.WriteString(w, ",")
				}
			}
			continue
		}
		io.WriteString(w, "=")
		switch {
		case fd.isNumberA && k > 1:
			for j := 0; j < k; j++ {
				if j > 0 {
					io.WriteString(w, ",")
				}
				hitHere := false
				for _, e := range idxs {
					if e == j {
						hitHere = true
						break
					}
				}
				if hitHere {
					w.Write(value(fd, i, j))
				} else {
					io.WriteString(w, ".")
				}
			}
		case len(idxs) == 1:
			w.Write(value(fd, i, idxs[0]))
		case len(idxs) > 1:
			for n, e := range byByte(idxs) {
				if n > 0 {
					io.WriteString(w, ",")
				}
				w.Write(value(fd, i, e))
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// readBytes reports bytes actually read from disk (proc/self/io), separating
// cold-cache I/O from page-cache hits across runs.
func readBytes() int64 {
	b, err := os.ReadFile("/proc/self/io")
	if err != nil {
		return 0
	}
	for _, ln := range bytes.Split(b, []byte("\n")) {
		if bytes.HasPrefix(ln, []byte("read_bytes:")) {
			v, _ := strconv.ParseInt(string(ln[len("read_bytes:"):]), 10, 64)
			return v
		}
	}
	return 0
}
