# Notebook: 17-field comparison on HDD (vs echtvar)

## Purpose

Rerun the 17-field controlled comparison
(`../vs_echtvar_hg19_gnomad_genomes_17_fields`) with identical artifacts,
identical query, identical cgroup, identical thread counts -- the only change
is the storage medium: the NVMe system disk is replaced by a 7200 rpm HDD.

This is the light end of the 17 -> 540 -> 568 dose series and the regime where
echtvar is strongest (small archive, fully cache-resident).

## Controlled conditions

- Query: GIAB HG001 GRCh37 v4.2.1, decomposed, 3,933,714 rows, identical file
  for both sides.
- Threads: echtvar = 4 concurrent per-chromosome processes (no thread flag);
  immkv = 1 process, 4 threads. Orchestration inside echtvar's timed window.
- Memory: both sides under a 10G cgroup (`MemoryMax=10G`, `MemorySwapMax=0`).
- Output: plain text on both sides; echtvar's forced gz round trip is
  converted outside the timed window and reported separately.
- Two passes each; page-cache state recorded as pass1/pass2.

## Environment (HDD)

- Disk: Seagate ST6000NM0034, 6 TB, 7200 rpm, USB bridge, ext4, mounted at
  `/mnt/data6t`.
- Library, archives, query and output all live on the same HDD, so the
  full-output round pays read/write head contention; the /dev/null round
  isolates the lookup path.
- Artifacts are byte-identical copies of the NVMe-built ones, verified by
  file count and byte total (`copy_artifacts.sh`); no rebuild is timed.

## Scripts, in execution order

1. `gen_fields.sh` -- regenerate `fields17.tsv/.json` from the gnomAD header
   (17-name keep list in `keep17.txt`).
2. `prep_query.sh` -- decompose HG001 and index the query.
3. `copy_artifacts.sh` -- rsync the narrow library, the archives and the
   full-load library from `NVME_ROOT`; verifies counts/bytes.
4. `run_echtvar.sh` -- round 1, 2 passes.
5. `run_immkv.sh` -- round 1, 4 runs (round1 = full library + whitelist,
   round2 = narrow library), 2 passes each.
6. `verify.sh` -- hit-set comparison plus full per-row per-field value
   comparison (`test/verifyvalues`).
7. `run_null.sh` -- round 2 (/dev/null), 2 passes each.
8. `collect_results.sh` -- copy .time reports and logs into `results/`.
9. `cleanup.sh` -- delete large texts, keep libraries and results.

Optional rebuilds (not run for the recorded numbers, kept for provenance):
`build_echtvar.sh` and `build_immkv.sh` build the archives and the narrow
library from raw gnomAD.

## Statistics to record (and where)

- `results/*.time` -- wall / user / sys / max RSS for every pass.
- `results/run_*_console.log` -- annotate phase split (key / query / format).
- `results/verify.log` -- rows, field comparisons, mismatches (expect 0).
- `results/echtvar.pass*.time` -- echtvar wall and CPU per pass.

## Recorded run

2026-11-09/10, 10G cgroup, HDD. Raw chain log lived outside the repository;
the committed record is `results/`.
