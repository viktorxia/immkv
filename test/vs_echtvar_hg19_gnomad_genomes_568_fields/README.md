# vs echtvar: all-fields (568) comparison (hg19, gnomAD v2.1.1 genomes)

The full annotation load: all 568 INFO fields including 16 String and 12
Flag. Versus the 540 experiment the added load is the STRING fields -- a
qualitative change, not a 5% increment. This directory also records two
corrected claims:

1. "echtvar cannot encode all fields" (previous session) is WRONG at the
   encode stage: it encodes everything, routing strings through categorical
   dictionaries (source: `TagType::String => FieldType::Categorical`,
   echtvar v0.2.4). On real chr22 those dictionaries DEGENERATE -- near
   record-count cardinalities: vep 3.05 M unique values, gq_hist_all 3.49 M,
   dp_hist_all 3.26 M of 3.57 M variants (`pilot_encode_chr22.log`).
2. The real wall is RUNTIME MEMORY, not encode feasibility (facts below).

## Measured facts (all numbers from `results/` or cited logs)

Under the 10G budget echtvar cannot annotate at ANY concurrency:

| attempt | outcome |
|---|---|
| 4 processes | OOM-killed after 12 s (journal: `oom-kill`) |
| 2 processes | OOM-killed (rc=143) |
| 1 process | OOM-killed (rc=143) |

Unconstrained reference runs (NOT budget-comparable; machine has 92 G):
4 processes, no limit: 9:16 / 9:56 wall, 3370 s CPU, **24.1 GB peak RSS per
process** (the string dictionaries live in the heap). 2 processes
(CGROUP_MEM=none, verification-only run): 16:45 / 17:01.

immkv under the 10G budget:

| round | pass1 | pass2 | CPU (user) | peak RSS |
|---|---|---|---|---|
| 1, full output | 5:06.55 | 4:54.66 | 207 s | **3.9 GB** |
| 2, /dev/null | 38.82 s | 47.54 s | 191 s | -- |

Verification (`results/verify.log`, against the unconstrained echtvar
output): 2,234,349,552 field comparisons -- zero value mismatches, zero hit
disagreements, all string fields (vep, histograms, popmax family) round-trip
exactly; 808 echtvar i32-overflow losses (VQSLOD-class magnitudes, same
limitation as the 540 experiment); floats within 1.1e-6, observed max
exactly 1.000e-06.

Build account: 22 archives, 84 GB, 4-way (encode RSS scales with variant
count; 8-way would exceed the machine), chr22 seeded from the pilot.

## Interpretation (clearly labeled as such)

- echtvar's speed privilege rests on two memory extravagances: page cache
  when the archive is small (killed by the 10G budget at 540 fields) and
  heap for degenerate string dictionaries at full load (~24 GB/process --
  infeasible in any production-sized container).
- immkv completes the same workload with 3.9 GB resident, in ~5 minutes,
  and every string value byte-exact.
- The 540-vs-568 contrast isolates the string load: echtvar goes from
  "6.5 min in 10G" to "cannot run in 10G"; immkv goes 5:28 → 5:07 wall and
  43 → 39 s null -- the string fields cost immkv almost nothing (raw-byte
  storage, no per-field decode).

## Reproducing

See `notebook.md`. `probe_encode_chr22.sh` reproduces the cardinality
census in ~30 min. echtvar's v0.2.4 encoder/decoder sources cited by every
comparison rule are in `sources/` (fetched from the tagged release via
api.github.com).

## Columnar round (2026-11-11, NVMe)

Rebuilt full library: 568 selected columns = 116.03 GB compressed / 978.2 GB
raw. `annotate -threads 16`, `/dev/null`: default path selection 37.61 s wall
(query 26.64 s + format 119.87 s; source `results/col_nvme.txt`). The cost
model picks record-order, matching the existing record null times (38.82 /
47.54 s) and full-output times (5:06.55 / 4:54.66 s).
