# vs echtvar on HDD: 540-field comparison (hg19, gnomAD v2.1.1 genomes)

Same controlled comparison as `../vs_echtvar_hg19_gnomad_genomes_540_fields`,
with the SAME artifacts copied onto a 7200 rpm HDD instead of the NVMe system
disk. Nothing is rebuilt: `copy_artifacts.sh` copies the library and archives
byte for byte and verifies file counts and byte totals.

The HDD is a Seagate ST6000NM0034 (6 TB, 7200 rpm) behind a USB bridge,
ext4, mounted at `/mnt/data6t`. All timed runs use the same 10G cgroup
(`MemoryMax=10G`, `MemorySwapMax=0`) and the same query; immkv runs one
process with 4 threads, echtvar 4 concurrent per-chromosome processes.

## Measured facts (all numbers from `results/`, 2 passes each, 10G cgroup)

Round 1 -- full plain-text output (echtvar's forced gz round trip converted
outside the timed window):

| side | pass1 | pass2 | CPU (user) | peak RSS |
|---|---|---|---|---|
| echtvar, 4 processes | 8:17.35 (`results/echtvar.pass1.time`) | 8:54.23 (`results/echtvar.pass2.time`) | 2257-2258 s | ~0.54 GB/process |
| immkv, full library + 540-field whitelist | 45:47.08 (`results/immkv.pass1.time`) | 43:15.83 (`results/immkv.pass2.time`) | 225.4-225.8 s | ~1.38 GB |

Round 2 -- output to /dev/null (lookup + format only):

| side | pass1 | pass2 | CPU (user) |
|---|---|---|---|
| echtvar | 7:43.73 (`results/echtvar_null.pass1.time`) | 7:45.71 (`results/echtvar_null.pass2.time`) | 2249-2253 s |
| immkv | 29:56.13 (`results/immkv_null.pass1.time`) | 29:26.11 (`results/immkv_null.pass2.time`) | 223.9-224.8 s |

Verification (`results/verify.log`): 2,124,205,560 field comparisons across
3,933,714 rows -- zero value mismatches, zero hit disagreements; 808 values
classified as echtvar i32-overflow losses (documented rival limitation),
same as the NVMe run.

Phase split (`results/run_immkv_console.log`): pass1 query 2724.19 s /
format 89.56 s / key 0.35 s; pass2 query 2576.20 s / format 89.30 s / key
0.36 s. Peak RSS 1,380,936 KB (`results/immkv.pass1.time`).

Artifacts copied (`copy_artifacts.sh`, verified): full-load library 46 files /
392,302,740,673 bytes; echtvar archives 22 files / 63,294,262,256 bytes.

## Interpretation (clearly labeled as such)

- The picture inverts relative to NVMe: there immkv was ~1.25x faster in
  round 1 and ~9x in round 2; here echtvar wins by ~5.3x (8:17 vs 45:47) and
  ~6.5x (7:44 vs 29:56).
- echtvar slows by only ~1/3 versus NVMe (6:14 -> 8:17), because it streams
  each chromosome's 59 GB archive sequentially; the extra time matches a
  sequential cold read.
- immkv's query stage (2724 s) equals a full pass over the ~366 GiB value
  file at ~180 MB/s. The sparse 4.8% hit pattern cannot beat that floor on a
  single rotating arm; the /dev/null round shows the same cost, so output
  writing is not the driver.
- immkv keeps a 10x CPU advantage (225 s vs 2258 s); on HDD that buys
  nothing because wall time is I/O-bound.

## Reproducing

Script order (see `notebook.md`):

1. `gen_fields.sh` -- regenerate `fields540.tsv/.json` (numeric mode).
2. `prep_query.sh` -- decompose + index the query.
3. `copy_artifacts.sh` -- copy library and archives from `NVME_ROOT`.
4. `run_echtvar.sh` -- round 1 (full output, 2 passes).
5. `run_immkv.sh` -- round 1 (2 passes).
6. `verify.sh` -- hit-set plus full per-row per-field comparison.
7. `run_null.sh` -- round 2 (/dev/null, 2 passes each side).
8. `collect_results.sh`, `cleanup.sh`.

## Columnar round (2026-11-11)

Library rebuilt with `test/buildvcf` (`BuildFromVCFs`, `IMMKV_PACK_KB=-1`):
record-order files are byte-identical to the old 366 GiB library (chr22
`cmp`; total `chrN.vals` 387,815,825,210 B, same as the old copy), plus
`chrN.col` 116,084,593,959 B and `chrN.manifest`. The 540 selected columns
measure 88.51 GB compressed / 673.6 GB raw over the 261,942,336 records
(`tmp/colstats.py`, parsing the column directories of all 23 `chrN.col`).

Run: `annotate -threads 16 -batch 4096` under the same 10G cgroup, columnar
path forced with `IMMKV_FORCE_COLUMNAR=1`; `-fields fields540.tsv` is mapped
to manifest column numbers.

| run | wall | query + format | peak RSS | source |
|---|---:|---:|---:|---|
| full output (66 GB) | 23:37.27 | 1385.82 + 90.28 s | 3.78 GB | `results/immkv_col.pass1.time` |
| /dev/null | 14:58.03 | 878.95 + 91.20 s | 4.82 GB | `results/immkv_col.null.time` |

Record-order on the same information (existing round): 45:47.08 / 43:15.83
full output, 29:56.13 / 29:26.11 null (`results/immkv.pass1.time`,
`immkv.pass2.time`, `immkv_null.*.time`). echtvar: 8:17.35 / 8:54.23 full,
7:43.73 / 7:45.71 null.

Correctness: `cmp` of the 66 GB columnar output against the record-order
`immkv.pass2.vcf` is byte-identical (`results/col_cmp_hdd.log`,
`540-identical`).

Interpretation (labeled): columnar is 1.94x faster than record-order on HDD
(23:37 vs 45:47) and 2.00x on the null round, but still 2.85x slower than
echtvar. The read is 88.5 GB of lossless column text plus 66 GB of text
output on one arm; echtvar streams a 59 GB quantized archive, so the
remaining gap is encoding density, not scheduling.

### gz-output timing (same protocol as echtvar)

`annotate -out immkv_col.pass1.vcf.gz` writes bgzf inside the timed window;
the plain-text expansion is then a separate step, exactly like echtvar.

| step | wall | source |
|---|---:|---|
| annotate -> gz (12.4 GB) | 16:36.03 | `results/immkv_col_gz.pass1.time` |
| zcat -> 66 GB text (untimed) | 416 s | `results/immkv_col_gz.convert.log` |
| same plain-text deliverable | **23:32** | sum of the two |

`cmp` of the expanded text against `immkv.pass2.vcf`: identical
(`results/immkv_col_gz.convert.log`, `540-gz-identical`). Writing gzip moves
about 7 minutes of text generation outside the timer; it does not change the
total work (23:32 vs 23:37 for direct plain text). echtvar's same-deliverable
total is 8:17 + 6:12 = 14:29.
