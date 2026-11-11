# HDD (568 fields, all INFO fields): immkv under the full annotation load

Same library, query, cgroup and threads as
`../vs_echtvar_hg19_gnomad_genomes_568_fields`, with the artifacts copied
onto a 7200 rpm HDD instead of the NVMe system disk. Nothing is rebuilt:
`copy_artifacts.sh` copies the library and archives byte for byte and verifies
file counts and byte totals.

There is no echtvar arm in this directory. At 568 fields echtvar is
OOM-killed under the 10G cgroup at any concurrency (see the NVMe experiment's
README and `results/echtvar_null.pass1.time` below); that is a property of the
tool, not of the medium. The HDD question for this field count is only "what
does immkv's full load cost on a rotational disk", plus output correctness.

## Measured facts (all numbers from `results/`, 10G cgroup)

| round | pass1 | pass2 | CPU (user) | peak RSS |
|---|---|---|---|---|
| 1, full output | 44:38.83 (`results/immkv.pass1.time`) | 44:19.51 (`results/immkv.pass2.time`) | 226.9-229.4 s | ~1.41 GB |
| 2, /dev/null | 29:49.27 (`results/immkv_null.pass1.time`) | 29:26.10 (`results/immkv_null.pass2.time`) | 227.5-228.1 s | ~1.44 GB |

echtvar attempt under the same 10G cgroup: killed by signal 15 after 25.47 s
(`results/echtvar_null.pass1.time`; the file records the kill, exit status 0
is systemd's wrapper). This reproduces the NVMe finding.

Phase split (`results/run_immkv_console.log`): pass1 query 2665.68 s /
format 91.62 s / key 0.33 s; pass2 query 2643.01 s / format 91.38 s / key
0.33 s.

Verification (`results/chain_hdd.log`, line `CMP_OK`): the HDD pass2 output is
byte-identical to an NVMe reference run of the same binary on the same
library contents and query. The NVMe 568 output was itself verified against
the unconstrained echtvar output in the sibling NVMe experiment
(2,234,349,552 field comparisons, zero mismatches). Reference run:
2:12.09 wall / 201.6 s CPU (`results/nvme_reference.time`), deleted after the
comparison.

Artifacts copied (`copy_artifacts.sh`, verified): full-load library 46 files /
392,302,740,673 bytes; echtvar archives 22 files / 89,556,740,819 bytes.

## Interpretation (clearly labeled as such)

- Full-load immkv costs ~44 min with output and ~29.5 min without on HDD,
  against 5:07 / 39-48 s on NVMe. The /dev/null round is the physical floor:
  a sparse pass over the ~366 GiB value file, ~180 MB/s on one arm.
- Output writing adds ~15 min: 73 GB of text shares the same disk with the
  library reads, so the head is time-shared.
- The HDD run changes wall time only; the bytes returned are identical to
  the NVMe run, and the NVMe run's field-level agreement with echtvar
  transfers to this run.
- Where echtvar cannot run at all (10G budget), immkv still completes: that
  remains true on either medium.

## Reproducing

Script order (see `notebook.md`):

1. `gen_fields.sh` -- regenerate `fields568.tsv/.json` (all fields).
2. `prep_query.sh` -- decompose + index the query.
3. `copy_artifacts.sh` -- copy library and archives from `NVME_ROOT`.
4. `run_immkv.sh` -- round 1 (2 passes; needs ~75 GB free on the HDD).
5. `verify_nvme_reference.sh` -- NVMe reference run + byte compare
   (needs ~70 GB free on the NVMe filesystem).
6. `run_null.sh` -- round 2 (/dev/null, 2 passes; immkv only).
7. `collect_results.sh`, `cleanup.sh`.

## Columnar round (2026-11-11)

Library rebuilt with `test/buildvcf` (`BuildFromVCFs`, `IMMKV_PACK_KB=-1`):
record-order byte layout matches the old library (chr22 `cmp`; total
`chrN.vals` 387,815,825,210 B), plus `chrN.col` 116,084,593,959 B. The 568
selected columns measure 116.03 GB compressed / 978.2 GB raw over
261,942,336 records (`tmp/colstats.py`, all 23 column directories).

Run: `annotate -threads 16 -batch 4096`, 10G cgroup, columnar forced with
`IMMKV_FORCE_COLUMNAR=1`.

| run | wall | query + format | peak RSS | source |
|---|---:|---:|---:|---|
| full output (72 GB) | 25:11.04 | 1497.97 + 93.90 s | 5.30 GB | `results/immkv_col.pass1.time` |
| /dev/null | 17:10.44 | 1015.73 + 94.87 s | 5.18 GB | `results/immkv_col.null.time` |

Record-order (existing round): 44:38.83 / 44:19.51 full, 29:49.27 /
29:26.10 null. echtvar cannot complete 568 fields under the 10G cap.

Correctness: `cmp` of the 72 GB columnar output against the record-order
`immkv.pass2.vcf` is byte-identical (`results/col_cmp_hdd.log`,
`568-identical`).

Interpretation (labeled): columnar is 1.77x faster than record-order full
output (25:11 vs 44:39) and 1.74x on the null round. The absolute floor is
116 GB of lossless column data plus 72 GB of text output on one arm; the
String/Flag fields add ~27.5 GB over the 540 subset.

### gz-output timing (same protocol as echtvar)

`annotate -out immkv_col.pass1.vcf.gz` (bgzf, 4 compression workers) inside
the timed window; plain-text expansion outside.

| step | wall | source |
|---|---:|---|
| annotate -> gz (14 GB) | 18:47.25 | `results/immkv_col_gz.pass1.time` |
| zcat | cmp against the 72 GB record text (streamed) | 447 s | `results/immkv_col_gz.convert.log` (`rc=0`) |

The streamed comparison confirms byte identity. Same plain-text deliverable:
~18:47 + ~6.5 min ≈ 25:17, i.e. the same as direct plain text (25:11); gzip
moves text generation outside the timer but does not reduce total work.
