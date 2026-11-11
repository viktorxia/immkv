# vs echtvar: 4-core / 16G / gzip-output controlled rerun (hg19, gnomAD v2.1.1)

## Purpose

Re-measure the hg19 gnomAD comparison with the deployment variables fixed.
The earlier numbers mixed thread counts, memory caps and output forms: the
gzip rows used 16 logical CPUs, a 10G cgroup, and immkv forced onto the
columnar path. This experiment holds one sample, four physical cores and 16G
for every timed cell, and makes both tools compress and write gzip inside the
timed window.

## Protocol

- Query: the same decomposed HG001 VCF used by every earlier round,
  3,933,714 data rows. One sample at a time.
- CPU: each timed scope is pinned with `taskset -c 0-3`; CPUs 0-3 are four
  distinct physical cores on this machine (SMT siblings 8-11 are excluded).
  `GOMAXPROCS=4`.
- Memory: `MemoryMax=16G`, `MemorySwapMax=0` on the scope.
- Output: both sides write `.vcf.gz`; compression and the write are inside
  the timed window. Decompression and verification happen outside it.
- immkv: `test/annotate -threads 4 -batch 4096`, default path selection (no
  `IMMKV_FORCE_COLUMNAR` / `IMMKV_FORCE_RECORD`); `IMMKV_DEBUG_PATH=1`
  records the per-batch cost estimate and the chosen path.
- echtvar: four concurrent per-chromosome processes (`xargs -P 4`), as its
  only parallelism; `bcftools view` slicing stays inside the window
  (echtvar's archive layout requires it). The htslib read/write thread
  counts are compiled into the binary.
- Two passes per cell; both wall times are reported.
- 568-field echtvar is not attempted: its measured peak is 24.1 GB per
  process, so no concurrency fits a 16G cgroup (recorded as "cannot run").

## Script order

1. `run_all.sh` builds `test/annotate` once, checks free space and that no
   other annotation process is running, then runs every cell.
2. `verify.sh` compares pass-2 outputs after all timed cells are finished.

## What is recorded

- Wall clock per pass (`/usr/bin/time -v`, scope inside the timer).
- gzip output bytes per pass.
- immkv path decisions (`IMMKV_DEBUG_PATH` stderr logs).
- Verification: row/hit agreement and full per-field comparison.

## Artifacts

Libraries, echtvar archives and the query stay in the existing per-field
trees under `/mnt/data6t/tian/immkv_test_hdd` and
`$HOME/immkv_test_nvme`. New outputs go to `<root>/test/4c16g_gz/`; old
records are not overwritten.
