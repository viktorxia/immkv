# vs echtvar on HDD: 17-field controlled comparison (hg19, gnomAD v2.1.1 genomes)

Same controlled comparison as `../vs_echtvar_hg19_gnomad_genomes_17_fields`,
with the SAME artifacts copied onto a 7200 rpm HDD instead of the NVMe system
disk. Nothing is rebuilt: `copy_artifacts.sh` copies the libraries and
archives byte for byte and verifies file counts and byte totals.

The HDD is a Seagate ST6000NM0034 (6 TB, 7200 rpm) behind a USB bridge,
ext4, mounted at `/mnt/data6t`. All timed runs use the same 10G cgroup
(`MemoryMax=10G`, `MemorySwapMax=0`) and the same query as the NVMe
experiment; immkv runs one process with 4 threads, echtvar 4 concurrent
per-chromosome processes.

## Measured facts (all numbers from `results/`, 2 passes each, 10G cgroup)

Round 1 -- full plain-text output:

| side | pass1 | pass2 | CPU (user) |
|---|---|---|---|
| echtvar, 4 processes | 23.60 s (`results/echtvar.pass1.time`) | 19.53 s (`results/echtvar.pass2.time`) | 107.9-108.8 s |
| immkv, full library + 17-field whitelist | 31:50.87 (`results/immkv.round1.pass1.time`) | 30:18.32 (`results/immkv.round1.pass2.time`) | 162.5-162.8 s |
| immkv, 17-field narrow library | 1:34.82 (`results/immkv.round2.pass1.time`) | 1:18.23 (`results/immkv.round2.pass2.time`) | 47.5-47.8 s |

Round 2 -- output to /dev/null (lookup + format only):

| side | pass1 | pass2 | CPU (user) |
|---|---|---|---|
| echtvar | 25.19 s (`results/echtvar_null.pass1.time`) | 18.78 s (`results/echtvar_null.pass2.time`) | 108.2-109.7 s |
| immkv, full library | 29:50.67 (`results/immkv_null.round1.pass1.time`) | 29:26.90 (`results/immkv_null.round1.pass2.time`) | 162.9-163.1 s |
| immkv, narrow library | 1:04.90 (`results/immkv_null.round2.pass1.time`) | 56.00 s (`results/immkv_null.round2.pass2.time`) | 47.4-47.5 s |

Verification (`results/verify.log`): 66,873,138 field comparisons across
3,933,714 rows -- zero value mismatches, zero hit disagreements; max float
deviation 1.000e-06 (tolerance 1.1e-6) -- identical to the NVMe run.

Phase split (`results/run_immkv_console.log`): full-library pass1 query
1894.68 s / format 27.85 s / key 0.32 s; narrow-library pass2 query 75.99 s /
format 3.73 s. Peak RSS: 1,348,208 KB (full library, `results/immkv.round1.pass1.time`),
308,308 KB (narrow, `results/immkv.round2.pass1.time`).

Artifacts copied (`copy_artifacts.sh`, verified): full-load library 46 files /
392,302,740,673 bytes; narrow library 44 files / 10,077,280,800 bytes;
echtvar archives 22 files / 3,021,239,374 bytes.

## Interpretation (clearly labeled as such)

- echtvar is almost unaffected by the medium: 23.6 / 19.5 s on HDD vs
  17.8 / 18.0 s on NVMe. Its 2.9 GB archive fits the 10G budget's page cache,
  and each chromosome is streamed sequentially.
- The full-load immkv library collapses on HDD: 30-32 min against 1:35-2:22
  on NVMe. The value file is ~366 GiB read in key order with 4.8% hit
  density; measured query time (1894 s) matches a full sequential pass at
  ~180 MB/s, i.e. the drive's sequential floor.
- The 9.4 GB narrow library stays usable (1:04-1:35): its size fits the
  page-cache allowance, and repeated passes read it mostly from memory
  (second pass 56 s).
- CPU stays with immkv (48-163 s vs 108 s), but on HDD wall time is set by
  bytes read, not CPU.

## Reproducing

Script order (see `notebook.md` for details):

1. `gen_fields.sh` -- regenerate `fields17.tsv/.json` from the gnomAD header.
2. `prep_query.sh` -- decompose + index the GIAB HG001 query.
3. `copy_artifacts.sh` -- copy libraries and archives from `NVME_ROOT`.
4. `run_echtvar.sh`, `run_immkv.sh` -- round 1, 2 passes each.
5. `verify.sh` -- hit-set plus full per-row per-field comparison.
6. `run_null.sh` -- round 2 (/dev/null), 2 passes each.
7. `collect_results.sh`, `cleanup.sh`.

## Columnar round (2026-11-11)

Library rebuilt with `test/buildvcf` (`BuildFromVCFs`): full library
record-order matches the old bytes (chr22 `cmp`) plus `chrN.col`
116,084,593,959 B; the 17-field narrow library has 6,018,755,545 B of values
and 3,874,932,029 B of columns over 261,942,336 records. The 17 selected
columns in the full library measure 3.87 GB compressed / 26.7 GB raw.

Run: `annotate -threads 16 -batch 4096`, 10G cgroup, columnar forced.

| library / run | wall | query + format | source |
|---|---:|---:|---|
| full library, full output (3.5 GB) | 3:30.24 | 205.72 + 4.77 s | console log |
| full library, /dev/null | 3:32.25 | 208.98 + 4.77 s | `results/immkv_col.full.null.time` |
| narrow library, full output | 1:31.78 | 88.64 + 3.64 s | console log |
| narrow library, /dev/null | 1:04.11 | 60.80 + 3.56 s | `results/immkv_col.narrow.null.time` |

Record-order (existing rounds): full library 31:50.87 / 30:18.32 full,
29:50.67 / 29:26.90 null; narrow library 1:34.82 / 1:18.23 full, 1:04.90 /
56.00 null. echtvar: 23.60 / 19.53 full, 25.19 / 18.78 null.

Correctness: columnar outputs of both libraries are byte-identical to the
record-order outputs (`results/col_cmp_hdd.log`).

Interpretation (labeled): on the full library columnar cuts 17-field HDD
time from ~30 min to 3.5 min (9.1x); on the narrow library columnar and
record-order are the same (~1-1.5 min), because both read the 3.9 GB of
column data / 6.0 GB of values plus the 4.5 GB key stream and write 3.5 GB
of text. Neither reaches echtvar's 19-25 s: its 2.9 GB archive is both
smaller than immkv's 3.9 GB columnar values and does not carry a separate
4.5 GB key stream.
