# vs echtvar: v003 controlled rerun (4 cores / 16G / gzip output)

## Purpose

Repeat the 2026-11-11 controlled matrix on the v003 on-disk format: same
sample, same query, same 4-core / 16G cgroup protocol, same gzip-output
window. The v0.0.2 record of that matrix lives one directory up
(`../vs_echtvar_hg19_gnomad_genomes_4c16g_gz/`); this directory records the
v003 numbers and does not overwrite those files.

v003 changes what is stored, not the query surface: `Get` / `GetBatch` still
return one byte value per key, but the value is a typed record (presence
bitmap + typed payload), so `test/annotate` renders it back to normalized
INFO text. Floats use `strconv.FormatFloat(..., 'g', -1, 64)` and are
compared numerically, not byte-for-byte, against echtvar.

## Protocol

- Query: the decomposed HG001 VCF, 3,933,714 data rows, 3,875,208 hits
  (98.51%). Source: `results/matrix.log`, `rows ...` line.
- CPU: `taskset -c 0-3`, `GOMAXPROCS=4`; CPUs 0-3 are four physical cores,
  their SMT siblings (8-11) are excluded. Source: `run_all.sh`, `common.sh`.
- Memory: `MemoryMax=16G`, `MemorySwapMax=0` on a per-cell systemd scope.
- Output: both tools write `.vcf.gz`; compression and the write are inside
  the timed window.
- immkv: `test/annotate -threads 4 -batch 4096`, default path selection,
  `IMMKV_DEBUG_PATH=1`.
- echtvar: four concurrent per-chromosome processes. 568-field echtvar is
  not attempted under the cap: its measured peak is 24.1 GB per process
  (`../vs_echtvar_hg19_gnomad_genomes_568_fields/README.md`).
- Two passes per cell. Sources:
  `results/<media>.<fields>.<tool>.passN.time`.

## Libraries

The full 568-field and 17-field libraries are rebuilt for v003 with per-value
framing (`IMMKV_PACK_KB=-1`, trained dictionary), then copied from NVMe to
HDD. Source: `build_v003.sh`.

| library | fields | records | build id | schema hash | non-finite floats | build wall | peak RSS |
|---|---:|---:|---|---|---:|---:|---:|
| full | 568 | 251,789,307 | `a6fa6f081196a359aec12e89ccdf8bcc` | `1e91a99722c59d5e` | 25 | 1:51:06 | 9.7 GB |
| 17-field | 17 | 251,789,307 | `7e9d2daf469b5ec1581da28cfc6b2894` | `99944e3316871f58` | 0 | 36:08 | 1.4 GB |

Byte totals (keys + vals + col = total):

| library | keys bytes | vals bytes | col bytes | total bytes |
|---|---:|---:|---:|---:|
| full | 4,312,533,441 | 248,261,546,086 | 108,194,223,184 | 360,768,302,711 |
| 17-field | 4,312,533,441 | 8,951,472,228 | 3,733,263,400 | 16,997,269,069 |

Sources: `results/library-facts.txt` (generated from each library's 22
`chrN.manifest` files), `results/build_full_nvme.log`,
`results/build_17_nvme.log`.

## Measured wall times (seconds)

| media | fields | v003 immkv p1 | v003 immkv p2 | v003 echtvar p1 | v003 echtvar p2 | v0.0.2 immkv p1 | v0.0.2 immkv p2 | v0.0.2 echtvar p1 | v0.0.2 echtvar p2 |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|
| nvme | 17 | 8.72 | 8.72 | 27.54 | 24.34 | 26.83 | 27.02 | 24.49 | 24.36 |
| nvme | 540 | 164.20 | 166.05 | 649.48 | 634.90 | 151.42 | 150.54 | 715.96 | 614.57 |
| nvme | 568 | 179.49 | 178.38 | — | — | 164.93 | 166.50 | — | — |
| hdd | 17 | 66.78 | 75.68 | 28.70 | 25.11 | 78.80 | 45.39 | 28.86 | 25.33 |
| hdd | 540 | 1640.64 | 1589.26 | 653.62 | 650.97 | 2092.78 | 2040.67 | 663.83 | 657.13 |
| hdd | 568 | 1660.06 | 1598.83 | — | — | 2114.83 | 2113.82 | — | — |

v003 sources: `results/<media>.<fields>.<tool>.passN.time`. v0.0.2 sources:
`../vs_echtvar_hg19_gnomad_genomes_4c16g_gz/results/<media>.<fields>.<tool>.passN.time`.

## immkv path selection (pass 2, 982 decisions per cell)

`982 = 961 windows of 4096 rows + 21 chromosome-boundary windows`. Sources:
`<work>/immkv.pass2.path.log` under
`<media_root>/test/4c16g_gz_v003/<fields>_fields/`; baseline decision counts
from `../vs_echtvar_hg19_gnomad_genomes_4c16g_gz/README.md`.

| media | fields | v003 columnar / record | v0.0.2 columnar / record |
|---|---:|---:|---:|
| nvme | 17 | 0 / 982 | 874 / 108 |
| nvme | 540 | 0 / 982 | 0 / 982 |
| nvme | 568 | 0 / 982 | 0 / 982 |
| hdd | 17 | 938 / 44 | 976 / 6 |
| hdd | 540 | 0 / 982 | 31 / 951 |
| hdd | 568 | 0 / 982 | 7 / 975 |

## Verification

`verify.sh` compares every field of every row with `test/verifyvalues`.
17/540 use the echtvar outputs from the timed matrix; 568 uses one
unconstrained echtvar reference (2 processes) generated on HDD, because
24.1 GB/process cannot fit the 16G cap.

| cell | field comparisons | misses agreed | value mismatches | echtvar i32-overflow losses | float max diff | result |
|---|---:|---:|---:|---:|---:|---|
| nvme 17 | 66,873,138 | 1,007,196 | 0 | 0 | 1.000e-06 | OK |
| nvme 540 | 2,124,205,560 | 32,050,262 | 0 | 808 | 1.000e-06 | OK |
| hdd 568 | 2,234,349,552 | 72,156,854 | 0 | 808 | 1.000e-06 | OK |

Sources: `results/verify/nvme.17.verifyvalues.log`,
`results/verify/nvme.540.verifyvalues.log`,
`results/verify/hdd.568.verifyvalues.log`.

The 808 i32-overflow losses are echtvar's documented limitation
(`|VQSLOD|*2e6 > i32max` becomes `-1`); they are counted separately, not as
value mismatches.

## Point lookup (NVMe, full 568-field library)

Same tool and key sample as the v0.0.2 baseline
(`../vs_echtvar_hg19_gnomad_genomes_17_fields/results/pointlookup_nvme.jsonl`;
its header records that it measured the full 366 GiB library). 100,000 keys
sampled from the decomposed HG001 query (98,486 hits), 4 reader goroutines,
warmup 10,000, seed 1. Sources: `results/pointlookup_nvme.jsonl` (v003) and
the baseline file above.

| library | cache | run | QPS | mean | p50 | p99 | p999 |
|---|---:|---|---:|---:|---:|---:|---:|
| v0.0.2 full | 512 | 4t | 66,521 | 59.9 µs | 26.3 µs | 321.7 µs | 403.8 µs |
| v0.0.2 full | 512 | 4t repeat | 87,255 | 45.6 µs | 25.3 µs | 286.4 µs | 382.7 µs |
| v003 full | 512 | 4t cold | 15,258 | 261.5 µs | 298.3 µs | 406.9 µs | 556.2 µs |
| v003 full | 512 | 4t warm | 325,204 | 12.2 µs | 11.9 µs | 30.9 µs | 129.4 µs |
| v003 full | 0 | 4t warm | 378,722 | 10.4 µs | 9.7 µs | 20.5 µs | 79.0 µs |
| v003 full | 2048 | 4t warm | 301,090 | 13.0 µs | 12.3 µs | 35.8 µs | 155.5 µs |

- Only v003's first cache-512 run starts with a cold page cache; every later
  run, and both v0.0.2 4-thread runs, have the sampled keys' frames in the OS
  page cache. Compare the v003 cache-512 warm run with the v0.0.2 4-thread
  runs.
- Under the same warm condition, v003 lowers p50 by 2.1–2.2×, raises QPS by
  3.7–4.9×, and lowers p99 by 9–10×. Typed per-value frames are smaller than
  raw INFO text frames, so decompression and cache residency both improve.
- Misses: v003 p50 1.0–2.4 µs and 1.5–3.4 M QPS; v0.0.2 p50 2.3 µs and
  1.5 M QPS.
- The cold v003 run is not comparable to the baseline's warm runs: it reads
  the sampled frames from NVMe at random, p50 298 µs.

## Interpretation

- 540/568 share per-value framing on both sides, so those rows compare
  formats. On NVMe v003 is slower: 540 +8.4%/+10.3% wall and +12.8% user CPU
  (pass 1: 609.83 s vs 540.53 s); 568 +8.8%/+7.1% wall and +12.1% user CPU
  (666.13 s vs 594.21 s). NVMe is CPU-bound, and typed decode costs more than
  text passthrough. Sources: the time files above.
- On HDD the same rows are faster under v003: 540 −21.6%/−22.1% wall with
  −18.4% filesystem inputs (53,699,560 vs 65,829,072 blocks); 568
  −21.5%/−24.4% wall with −15.7% inputs (53,587,112 vs 63,564,032), while
  user CPU stays within +6%. HDD is bytes-read-bound, and typed record frames
  are smaller than raw INFO text. Sources: the time files above.
- The 17-field row is not a format comparison: the v0.0.2 17-field library
  was built by `test/buildnarrow` with `immkv.Create(..., immkv.Options{})`
  (default 64KB packed frames; source:
  `git show bdfa14e:test/buildnarrow/main.go`), while the v003 17-field
  library uses per-value framing. First-window path logs: old 963 frames /
  6,467,620 bytes vs new 4,060 frames / 395,771 bytes. Treat 8.72 s vs
  26.83 s as packed-versus-per-value, not v003-versus-v0.0.2. Source:
  `<work>/immkv.pass2.path.log` for both experiments.

## Environment

- NVMe library root: `/home/tian/immkv_test_nvme`; HDD library root:
  `/mnt/data6t/tian/immkv_test_hdd`. Both are set in `common.sh`.
- Libraries are outside the repository
  (`<media_root>/gnomad/immkv` and `<media_root>/test/17_fields/lib17`).
- Result texts and verification temporaries live under
  `<media_root>/test/4c16g_gz_v003/`; only times and logs are committed.

## Reproduction

```bash
# rebuild both libraries (refuses to overwrite existing ones)
GNOMAD_SRC=/path/to/gnomad/GRCh37/v2.1.1 ./build_v003.sh both

# controlled matrix (resumable via .done markers)
./run_all.sh

# full per-field verification, including the unconstrained 568 reference
./verify.sh
```
