# immkv

A standalone storage-engine library: **build once, immutable, sorted KV**.
The name comes from immutable sorted KV. It is the underlying engine for the
`storage_backend: kv-custom` access implementations in the annoplex catalog,
and can also be used independently in other projects.

## Philosophy

Three decisions define this engine:

1. **Explicit reads, no trust in mmap.** The predecessor immdb used mmap
   direct access and suffered severe read amplification with "large
   library + large values + sparse query positions" — both page-fault
   granularity and kernel readahead are outside the engine's control.
   immkv moves all reads to explicit on-demand pread/io_uring: how much to
   read, when to read, and how many segments to coalesce into — all in the
   engine's hands.
2. **The on-disk format is tailored to query patterns.** The key stream is
   cut into 64-record micro-blocks with the block index resident in memory;
   streaming batch lookups binary-search within a block instead of parsing
   variable-length records one by one; values are framed individually (with
   a zstd dictionary trained at build time), so a point lookup decompresses
   one value and reads only that value's bytes.
3. **Building may be slow, queries must be fast; CGO is the exception, not
   the default.** Pure Go by default; CGO is introduced — backed by
   measurements — only on the zstd codec hot path, and a fully
   functional pure-Go fallback is kept: `CGO_ENABLED=0` builds exactly as
   usual, just slower.

## API

Two handle types cover the two query shapes; both can be open on the same
library at once (read-only, no conflict):

```go
// Build: stream records in via AddSorted (groups ascending); Finish persists atomically
w, _ := immkv.Create(cacheDir+"/chr1", immkv.Options{})
w.AddSorted(immkv.PosKey(0, 1000), key, info)
w.Finish()

// Point-lookup handle: one key answered whenever it arrives (random keys, online serving)
db, _ := immkv.OpenEager(cacheDir, 512)          // cacheMB: decompressed-cache budget
v, ok := db.Get(key)                              // miss → ok=false
db.GetPrefix(prefix, func(k, v []byte) bool { … })
db.Scan(lo, hi, func(k, v []byte) bool { … })     // lo/hi may be nil
db.Close()

// Streaming handle: query a fully sorted VCF start to finish, then exit (batch annotation)
sdb, _ := immkv.OpenStreaming(cacheDir, immkv.Options{Threads: 9})
vals, _ := sdb.GetBatch(keys)   // keys ascending; returns equal length & order, nil = miss
sdb.Close()
```

The `GetBatch` contract: input must be ascending (VCF sorted by chromosome
natural order + position satisfies this naturally); the returned slice is
only guaranteed valid until the next query; the batch size is the prefetch
window (a few thousand to a few hundred thousand all work); `Threads` is
the only concurrency knob — it sets the decompression parallelism;
in-flight read depth is independent of it on the io_uring path, and it
doubles as the pread reader count only on the fallback path. The typical
annotation loop is "accumulate a batch of keys → one GetBatch per library →
consume and discard":

```go
const batch = 4096
keys := make([][]byte, 0, batch)
for vcf.Scan() {
    keys = append(keys, immkv.VarKey(chrom, pos, ref, alt))
    if len(keys) < batch { continue }
    gnomadVals, _ := gnomadDB.GetBatch(keys)   // vals[i] aligns with keys[i]
    clinvarVals, _ := clinvarDB.GetBatch(keys)
    keys = keys[:0]
}
```

Three key-encoding functions: `ChromID`/`ChromName` (chromosome name ↔
first-byte encoding; first-byte order = chromosome natural order), `VarKey`
(variant key: chromosome + big-endian position + ref + alt), and `PosKey`
(position prefix, used as the AddSorted group key or for prefix queries).

### GetBatch's ordering requirement and chromosome naming (FAQ)

**The ordering requirement is "ascending byte order of the encoded keys"**;
for variant keys this amounts to three rules:

1. Chromosomes in **natural order** 1–22, X, Y, M — not lexicographic.
   This is the easiest pitfall: the `ChromID` encoding guarantees that
   first-byte magnitude equals natural order (1→0 … 22→21, X→22, Y→23,
   M→24), but many tools emit lexicographically sorted VCFs (chr1, chr10,
   chr11, …, chr19, chr2, …); lexicographically chr10 comes before chr2
   while the keys compare 9 vs 1 — feeding that in is a descent;
2. within one chromosome, positions ascending (VCF's POS column is
   naturally so);
3. multiple records at the same position (keys split from multi-allelic
   variants) may come in any order — equal keys are allowed.

Wrong ordering **never fails silently**: the engine checks adjacent keys
pairwise and returns `ErrUnsorted` immediately on a descent; it never
derives wrong misses from out-of-order input. If unsure, just feed the
keys in — an error means the source needs re-sorting.

**How to sort a VCF into natural order**: `bcftools sort` sorts by the
header's `##contig` declaration order + position, and mainstream projects'
headers are themselves in natural order (gnomAD, 1000G, dbSNP, etc.) —
sort and use directly:

```bash
bcftools view -h in.vcf.gz | grep contig   # check the header order first
bcftools sort -O z -o sorted.vcf.gz in.vcf.gz
```

Note that a few VCFs declare MT **first** in the header — `bcftools sort`
faithfully follows it, while M ranks last in our encoding; feeding that in
triggers `ErrUnsorted`. Fix the header order, then re-sort.

**Chromosome spelling normalization**: the key's first byte is the
canonical encoding, not the raw name; `chr1` and `1` mixed within one file
do not affect ordering. Spellings `ChromID` accepts:

| Spelling | Recognized | Encoding |
| --- | --- | --- |
| `1`…`22`, `chr1`…`chr22`, `CHR1` | ✓ | 0…21 |
| `X`/`x`/`chrX`/`CHRX`, same for Y | ✓ (case-insensitive) | 22, 23 |
| `M`/`MT`/`chrM`/`chrMT` (any case) | ✓ | 24 |
| `Chr1` (mixed form with a capital first letter) | ✗ | — |
| `NC_000001.11` and other accessions | ✗ | — |
| `chrUn_*`, `*_alt`, `HLA-*`, etc. | ✗ | — |

Unrecognized contigs return false and **the caller skips them** — alt/decoy
sequences carry no annotation value; that is the designed behavior, not a
defect.

## Performance

Controlled comparison experiments live in test/ (rival echtvar v0.2.4,
hg19 / gnomAD v2.1.1 genomes, 10G cgroup, two-round protocol, 17→540→568
field dose series); highlights:

- **17-field light load** (echtvar's home turf, archive fully page-cache
  resident): wall clock 18s vs 30s for the immkv narrow library; CPU 42s
  vs 109s (~2.6× saved); 66,873,138 field comparisons, zero mismatches;
- **568-field full load** (16 String + 12 Flag): echtvar OOMs at any
  concurrency under the 10G budget (unconstrained reference: 24.1GB peak
  RSS/process); immkv finishes entirely within 10G — ~5 minutes with full
  output, 39–48s null round, 3.9GB peak RSS; 2,234,349,552 field
  comparisons, zero value mismatches, strings byte-exact;
- **Field dose contrast** (540→568, adding the string fields): echtvar
  goes from 6.5 min within 10G to cannot run; immkv 5:28→5:07, null round
  43→39s — string fields cost almost nothing (raw-byte storage, no
  per-field decode).

## Contracts

- **Lifecycle**: Finish goes through temp file + fsync + rename; callers
  never see a half-built library; handles already open keep reading the old
  data. The same data files can be opened by any number of processes at
  once (no locks), and within one handle, concurrent use by multiple
  goroutines is safe.
- **Returned slices**: guaranteed valid only until the next query; copy
  anything you want to keep.
- **Duplicate keys**: allowed; `Get` returns the first one in sorted order,
  `GetPrefix`/`Scan` see all of them.
- **Streaming input**: a descent returns `ErrUnsorted` (guarding against
  silent wrong results); equal keys are allowed.

### How Threads is spent and split (FAQ)

`Threads` governs **the decompression parallelism inside each GetBatch
call**: every call spins up a three-stage pipeline (1 producer + 1 io_uring
read driver + N decompression workers) and tears it down when the call
ends. Note the read driver is fixed at a single goroutine — **in-flight
read depth has nothing to do with Threads**; that is the whole point of
io_uring. Only on the fallback path, when io_uring is unavailable, does
Threads double as the pread reader count.

Thread-scaling conclusion from measurement: **four threads is the
ceiling** — more threads only burn CPU without going faster
(decompression-pool contention + worsening GC amortization).

**One handle with 4 threads, or two handles with 2 threads each?** In the
single-stream shape, throughput is equivalent (CPU is not the bottleneck;
splitting 4 decompression workers across pipelines does not change the
total live capacity). Choose by scenario: fastest for a single sample →
1 handle × 4 threads; multiple samples in parallel → split by
orchestration convenience. **But concurrency has one real disk-contention
pitfall**: when two processes read the same disk with high queue depth at
the same time, a DRAM-less low-end NVMe (such as this machine's Kingston
NV3, which also hosts the system disk) collapses across the board —
aggregate bandwidth falls far below single-stream levels, with no CPU or
memory pressure. In production, when multiple tasks on one node share one
NVMe,
"bandwidth splits roughly evenly" does not hold for such disks;
countermeasures: cap the node's total in-flight read depth (drop per-task
threads to 2, or make the io_uring depth a node-level quota), or give the
node disks with DRAM caches.

## Implementation

### File layout

Two files per chromosome (`Create`'s basePath, e.g. `…/chr1`):

```text
chr1.keys: [64B header][micro-blocks ×N][block index]
  micro-block (64 records): [keyEnds u32×n][lens u32×n][baseValOff u64][key area]
                     lens flag bits: bit31 raw / bit30 packed frame / bit29 same-frame
  block index (resident in memory): [first-key 8B prefix u64 big-endian ×N][block offsets u64 ×N]
chr1.vals: [64B header][dictionary (per-value mode)][value frames ×N]
```

- A record's valOff is not stored per record: the micro-block stores the
  first record's absolute offset and the rest are derived from lens (the
  bit29 same-frame bit adds no previous frame length); the key stream is
  markedly smaller than with per-record absolute offsets.
- Three kinds of value frames: values shorter than `RawValueBelow`
  (default 64B) stored raw; in packed mode, values below `packBelow` are
  grouped by `packTarget` (default 64KB) into one frame (the group header
  holds the first record number and a length table; small values get their
  compression ratio from in-group context); `Options.PackTarget < 0`
  enters **per-value framing mode** — every value becomes its own frame,
  and a zstd dictionary trained at build time from the first batch of
  values (512KB corpus, 5000 samples, libzstd L9) is embedded in the vals
  header. gnomAD-like large-value libraries use per-value mode; ClinVar-like
  small-value libraries use the default packed mode.
- On open: the block index is read in once and stays resident (the
  point-lookup handle additionally gets an LRU of CacheMB, caching
  micro-blocks and decompressed values); the streaming handle caches
  nothing (the batch is the prefetch window).

### Read path

- **Streaming (GetBatch)** three-stage pipeline: a merge-join producer
  (forward-scans the key stream in micro-block windows, 256KB per window,
  in-block binary search, zero-allocation views) → read-segment driver →
  Threads-way decompression pool. A read segment = adjacent hits merged
  (gaps ≤4KB join into one segment); each read segment is fetched with one
  4KB-aligned O_DIRECT read, and an EOF short read past the file end counts
  as success when it covers the bytes needed.
- **Read driver**: io_uring (512 deep, a single-goroutine event loop
  submitting and reaping; in-flight read depth decoupled from thread
  count); falls back to a Threads-way pread read pool when io_uring or
  O_DIRECT is unavailable. After reading, fadvise(DONTNEED) on the covered
  range — values are read only once; leaving them in the page cache would
  only blow the cgroup limit.
- **Point lookup (Get)**: in-memory block-index binary search → read one
  micro-block (LRU) → in-block binary search → read the value frame
  and decompress. Zero syscalls on cache hits.

### Codec (hybrid tiers)

| Path | Backend | Notes |
| --- | --- | --- |
| Build compression (per-value mode) | CGO libzstd L9 | klauspost trains the dictionary (once) + gozstd CDict compresses value by value |
| Query decompression (dictionary libraries) | CGO libzstd DDict | standard frame format, interoperable with pure Go |
| Query decompression (packed libraries, no dictionary) | klauspost pure Go | gozstd is not faster than klauspost on packed frames even with dst fully provisioned, and holds an extra resident decode buffer — not used |

Under `CGO_ENABLED=0` everything falls back to klauspost; the read/write
format is unchanged.
Selection logic: at build time, CGO availability; at query time, whether
the library header carries a dictionary.

### Platforms

Full functionality on linux (io_uring/O_DIRECT/fadvise are linux-only;
other platforms automatically degrade to pread). Files are explicitly
little-endian, with absolute offsets and no address dependence — usable
directly across machines and architectures.

### Dependencies

- `github.com/valyala/gozstd` (CGO, vendored libzstd): main decoding path
  for dictionary libraries; under `CGO_ENABLED=0` falls back to
  `github.com/klauspost/compress` (pure Go).
- `golang.org/x/sys`: linux-only wrappers for fadvise/DirectIO and other
  syscalls.
- `github.com/biogo/hts`: used by the **`test/` comparison tools**
  (annotate/buildnarrow/buildfull) to read bgzip VCFs; the library itself
  does not depend on it. The referencers live in the repository, so
  `go mod tidy` keeps it.
