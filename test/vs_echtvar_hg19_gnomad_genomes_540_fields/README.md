# vs echtvar: 540-field comparison (hg19, gnomAD v2.1.1 genomes)

echtvar's ceiling: all 540 NUMERIC INFO fields (368 Integer + 172 Float) --
it cannot encode the 28 String/Flag fields at all (see the 568 experiment).
Together with the 17-field experiment this is the numeric dose-response pair
(same kind of work, 32x more fields).

## Measured facts (all numbers from `results/`, 2 passes each, 10G cgroup)

Round 1 -- full plain-text output (echtvar's forced gz round trip reported
separately, outside the timed window: 334-426 s per pass):

| side | pass1 | pass2 | CPU (user) | peak RSS |
|---|---|---|---|---|
| echtvar, 4 processes | 6:14.27 | 6:52.36 | 2298 s | ~0.54 GB/proc |
| immkv, full library + 540-field whitelist | 5:28.08 | 5:15.94 | 205 s | ~3.9 GB |

Round 2 -- output to /dev/null (lookup + format only):

| side | pass1 | pass2 | CPU (user) |
|---|---|---|---|
| echtvar | 6:43.02 | 6:39.87 | 2288 s |
| immkv | **43.16 s** | **47.41 s** | **191 s** |

Verification (`results/verify.log`): 2,124,205,560 field comparisons across
3,933,714 rows -- zero value mismatches, zero hit disagreements. 808 values
classified as echtvar i32-overflow losses: fields like VQSLOD whose
magnitude exceeds i32max/multiplier (~1073 at multiplier 2e6) are silently
encoded as missing (source: `(val * multiplier) as i32` saturates, echtvar
v0.2.4 encoder, see ../568/sources/). Float agreement within the derived
tolerance 1.1e-6 (quantization 5e-7 + 6-decimal print truncation 5e-7);
observed max exactly 1.000e-06 over 666 M float comparisons.

Build account: 22 archives, 59 GB, 8-way pipeline, resumable.

## Interpretation (clearly labeled as such)

- Under a 10G production budget the picture inverts from the 17-field case:
  round 1 immkv is ~1.25x faster on wall and 11x on CPU; round 2 (pure
  lookup+format) immkv is ~9x faster on wall and 12x on CPU.
- echtvar's archive (~59 GB) no longer fits the 10G page-cache allowance, so
  its lookups go to disk; its CPU also scales with field count (it decodes
  540 separate field streams per variant), keeping ~5.7 cores busy.
- immkv's round-1 wall is dominated by its own query stage on the 366 GB
  library (I/O-bound); removing output drops it to ~45 s.

## Reproducing

See `notebook.md`. Requires >=300 GB free during the build; runs are
unattended (~40 min round 1 + round 2 once libraries exist).
