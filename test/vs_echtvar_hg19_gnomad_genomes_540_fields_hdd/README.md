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
