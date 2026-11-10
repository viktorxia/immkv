# Notebook: 540-field comparison on HDD (vs echtvar)

## Purpose

Rerun the 540-field controlled comparison
(`../vs_echtvar_hg19_gnomad_genomes_540_fields`) with identical artifacts,
query, cgroup and thread counts; the only change is the storage medium (NVMe
system disk replaced by a 7200 rpm HDD).

At 540 fields the echtvar archive (59 GB) no longer fits the 10G budget's
page cache, so this is the regime where sequential-vs-sparse I/O behaviour of
the two layouts separates.

## Controlled conditions

- Query: GIAB HG001 GRCh37 v4.2.1, decomposed, 3,933,714 rows, identical file
  for both sides.
- Threads: echtvar = 4 concurrent per-chromosome processes; immkv = 1 process,
  4 threads. Orchestration inside echtvar's timed window.
- Memory: both sides under a 10G cgroup (`MemoryMax=10G`, `MemorySwapMax=0`).
- Output: plain text both sides; echtvar's gz conversion runs outside the
  timed window and is reported separately.
- Two passes each; page-cache state recorded as pass1/pass2.

## Environment (HDD)

- Disk: Seagate ST6000NM0034, 6 TB, 7200 rpm, USB bridge, ext4, mounted at
  `/mnt/data6t`.
- Library, archives, query and output share the same HDD; the full-output
  round therefore pays read/write head contention. The /dev/null round
  isolates lookup+format.
- Artifacts are byte-identical copies of the NVMe-built ones, verified by
  file count and byte total (`copy_artifacts.sh`); no rebuild is timed.

## Scripts, in execution order

1. `gen_fields.sh` -- regenerate `fields540.tsv/.json` (numeric mode).
2. `prep_query.sh` -- decompose + index the query.
3. `copy_artifacts.sh` -- rsync the archives and the full-load library from
   `NVME_ROOT`; verifies counts/bytes.
4. `run_echtvar.sh` -- round 1 (full output, 2 passes).
5. `run_immkv.sh` -- round 1 (2 passes; pass1 text deleted after pass2).
6. `verify.sh` -- hit-set plus full per-row per-field value comparison.
7. `run_null.sh` -- round 2 (/dev/null, 2 passes each side).
8. `collect_results.sh`, `cleanup.sh`.

Optional rebuild (not run for the recorded numbers, kept for provenance):
`build_echtvar.sh` builds the 22 archives from raw gnomAD (needs >=300 GB
free; build phases are not memory-capped).

## Statistics to record (and where)

- `results/*.time` -- wall / user / sys / max RSS for every pass.
- `results/run_*_console.log` -- annotate phase split (key / query / format).
- `results/verify.log` -- field comparisons, mismatches (expect 0), and the
  echtvar i32-overflow classification count.

## Recorded run

2026-11-09/10, 10G cgroup, HDD. Committed record: `results/`.
