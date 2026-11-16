# Cold-cache 4C16G performance summary

Tag: `v0.0.5`. Tested engine commit: `d2eac357f3fc126755b43ff8b44a64dc55bc0ec2`.

Status: six scenarios complete; 30 successful timing samples and six evidence-backed echtvar 568-field OOM attempts; all six full comparisons passed.

## Wall-time medians

Seconds; median of three successful compressed-output lookup runs. OOM has no median or speed ratio.

| Medium | Fields | immkv | echtvar | Result |
|---|---:|---:|---:|---|
| nvme | 17 | 13.20 | 32.27 | echtvar/immkv ratio 2.445 |
| nvme | 540 | 158.86 | 619.24 | echtvar/immkv ratio 3.898 |
| nvme | 568 | 172.47 | OOM | echtvar OOM in all three attempts |
| hdd | 17 | 96.46 | 27.60 | echtvar/immkv ratio 0.286 |
| hdd | 540 | 1468.32 | 640.94 | echtvar/immkv ratio 0.437 |
| hdd | 568 | 1469.52 | OOM | echtvar OOM in all three attempts |

## CPU and whole-scope memory

Median GNU time user+system seconds and median cgroup memory.peak. `memory.peak` includes charged file cache, memfd and kernel memory; it is not process RSS.

| Medium | Fields | immkv CPU s | echtvar CPU s | immkv cgroup peak bytes | echtvar cgroup peak bytes |
|---|---:|---:|---:|---:|---:|
| nvme | 17 | 33.15 | 102.59 | 4798017536 | 3219034112 |
| nvme | 540 | 616.02 | 2419.55 | 6947467264 | 17179869184 |
| nvme | 568 | 670.46 | OOM | 7959490560 | 17179869184 |
| hdd | 17 | 204.67 | 96.57 | 7506931712 | 3231981568 |
| hdd | 540 | 674.01 | 2410.14 | 5348339712 | 17179869184 |
| hdd | 568 | 719.20 | OOM | 5525741568 | 17179869184 |

For OOM rows the echtvar cgroup peak is the 16 GiB limit reached by each failed attempt, not a successful median.

## Full correctness verification

Each medium processed 3,933,714 rows and observed 3,875,208 immkv hits. Values below apply independently to each medium.

| Fields | Field comparisons | Value mismatches | Hit disagreements | echtvar i32 overflow losses | True float -1 candidates | Reference |
|---:|---:|---:|---:|---:|---:|---|
| 17 | 66,873,138 | 0 | 0 | 0 | 0 | fresh digest equivalent to historical reference |
| 540 | 2,124,205,560 | 0 | 0 | 808 | 663 | fresh digest equivalent to historical reference |
| 568 | 2,234,349,552 | 0 | 0 | 808 | 663 | historical-only correctness reference after fresh echtvar OOM |

Total across six comparisons: **4,425,428,250 field comparisons**, zero value mismatches, zero hit disagreements. Float tolerance 1.1e-6; maximum observed difference 1e-6.

## Protocol and evidence limits

- CPUs 0-3 on four distinct physical cores, GOMAXPROCS=4, memory.max=17179869184, memory.swap.max=0.
- Root global sync and `drop_caches=3` succeeded before every independently timed invocation.
- Three alternating rounds per cell: immkv/echtvar, echtvar/immkv, immkv/echtvar.
- Existing libraries and echtvar archives were reused; no format, magic, structure or file-count change.
- Raw per-invocation logs are intentionally not committed and remain in the local `results/` directory.
- Cgroup `io.stat` is unavailable; GNU time file-system input blocks are retained.
- GNU time RSS is maximum individual-process RSS; cgroup peak is a separate whole-scope metric.
- Six echtvar/568-field timing attempts OOMed within 16 GiB; no speed ratio is computed.
- Two fresh echtvar/568-field correctness attempts OOMed; historical outputs serve only as correctness references.
- One HDD/568-field failed correctness cleanup overlapped the next cache completion by 0.607921622 seconds; all old annotation children had exited before the next cache request. No timing invocation overlapped.
- Successful correctness scope status has a generic `timing_sample:true` metadata bug; summaries exclude verification records from timing medians.
- Tool-output agreement does not independently prove every semantic detail of the original VCF.

The machine-readable form is [final-summary.json](final-summary.json). Detailed local-only evidence paths and raw timings are in the uncommitted `results/` directory, including `acceptance-final.md`, `final-matrix-audit.md`, `summary-final-audited.md` and each `verify-*`/`timing-*` record.
