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
