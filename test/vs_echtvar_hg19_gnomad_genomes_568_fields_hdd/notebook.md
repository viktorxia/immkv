# Notebook: 568-field (all INFO fields) full load on HDD

## Purpose

Measure immkv's full annotation load (all 568 INFO fields, every stored byte
emitted) with the library on a 7200 rpm HDD, and verify that the returned
bytes are identical to an NVMe run.

No echtvar arm: at 568 fields echtvar is OOM-killed under the 10G cgroup at
any concurrency (see `../vs_echtvar_hg19_gnomad_genomes_568_fields`). The
single 25-second attempt recorded here reproduces that under HDD conditions
and is committed as evidence, not as a performance number.

## Controlled conditions

- Query: GIAB HG001 GRCh37 v4.2.1, decomposed, 3,933,714 rows.
- immkv: 1 process, 4 threads, 10G cgroup (`MemoryMax=10G`,
  `MemorySwapMax=0`).
- Two rounds: round 1 full plain-text output (~73 GB per pass), round 2
  `/dev/null` (lookup + format only). Two passes per round.

## Environment (HDD)

- Disk: Seagate ST6000NM0034, 6 TB, 7200 rpm, USB bridge, ext4, mounted at
  `/mnt/data6t`.
- Library, query and output share the same HDD; round 1 therefore pays
  read/write head contention, round 2 isolates lookup+format.
- Artifacts are byte-identical copies of the NVMe-built ones, verified by
  file count and byte total (`copy_artifacts.sh`); no rebuild is timed.

## Scripts, in execution order

1. `gen_fields.sh` -- regenerate `fields568.tsv/.json` (all fields; the
   Integer/Float/String/Flag census is printed).
2. `prep_query.sh` -- decompose + index the query.
3. `copy_artifacts.sh` -- rsync the archives and the full-load library from
   `NVME_ROOT`; verifies counts/bytes.
4. `run_immkv.sh` -- round 1 (2 passes; pass1 text deleted after pass2 is
   written). Needs ~75 GB free on the HDD.
5. `verify_nvme_reference.sh` -- run the same binary on the NVMe copy of the
   library, then byte-compare with the HDD pass2 output. Needs ~70 GB free on
   the NVMe filesystem; deletes the reference after the comparison.
6. `run_null.sh` -- round 2 (`/dev/null`, 2 passes; immkv only).
7. `collect_results.sh`, `cleanup.sh`.

Optional rebuild (not run for the recorded numbers, kept for provenance):
`build_echtvar.sh` builds the 22 archives from raw gnomAD (4-way; encode RSS
scales with variant count).

## Statistics to record (and where)

- `results/*.time` -- wall / user / sys / max RSS for every pass.
- `results/run_immkv_console.log` -- phase split (key / query / format).
- `results/null_rerun.log` -- the two null passes with their Elapsed lines.
- `results/chain_hdd.log` -- chain log ending in `CMP_OK` (byte-identical to
  the NVMe reference).
- `results/nvme_reference.time` -- the reference run's timing.
- `results/echtvar_null.pass1.time` -- the OOM evidence (killed, 25.47 s).

## Recorded run

2026-11-09/10, 10G cgroup, HDD. Committed record: `results/`.
