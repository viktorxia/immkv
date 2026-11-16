# 4C16G cold-cache comparison: v0.0.5

Status: complete. Six scenarios, all full-output verifications and all 36
scheduled timing attempts finished; 30 timing samples succeeded and six
were evidence-backed echtvar / 568-field OOM outcomes. This directory is
committed with tag `v0.0.5`; the tested engine code is the unchanged commit
`d2eac357f3fc126755b43ff8b44a64dc55bc0ec2`.

## Committed scope

This commit contains the reusable scripts, this README, `notebook.md` and
the final summary in `results/final-summary.md` / `results/final-summary.json`.
Per the selected scope, raw per-invocation evidence in `results/` is kept
on the test machine and is not committed. It includes the preparation and
preflight records and every `verify-*` / `timing-*` record with its
cache-clear report, resource constraints, cgroup counters, GNU time report,
per-chromosome logs and journals.

Consequently, the committed numbers can be traced to the local summary and
raw files on this machine, but those raw files are not part of Git history.
This is a deliberate departure from the repository rule that experiment
`results/` are committed, chosen because further test rounds are planned.
The committed summary records that limitation explicitly.

## Results

Wall times are medians of three successful compressed-output lookup runs,
in seconds. echtvar OOM has no successful median or speed ratio.

| Medium | Fields | immkv | echtvar | Result |
|---|---:|---:|---:|---|
| NVMe SSD | 17 | 13.20 | 32.27 | immkv 2.44 times faster |
| NVMe SSD | 540 | 158.86 | 619.24 | immkv 3.90 times faster |
| NVMe SSD | 568 | 172.47 | OOM in all three attempts | immkv completes within 16 GiB |
| HDD | 17 | 96.46 | 27.60 | echtvar 3.49 times faster |
| HDD | 540 | 1468.32 | 640.94 | echtvar 2.29 times faster |
| HDD | 568 | 1469.52 | OOM in all three attempts | immkv completes within 16 GiB |

Median CPU totals (user+system seconds) and cgroup memory peaks:

| Medium | Fields | immkv CPU | echtvar CPU | immkv peak | echtvar peak |
|---|---:|---:|---:|---:|---:|
| NVMe SSD | 17 | 33.15 | 102.59 | 4.47 GiB | 3.00 GiB |
| NVMe SSD | 540 | 616.02 | 2419.55 | 6.47 GiB | 16.00 GiB |
| NVMe SSD | 568 | 670.46 | OOM | 7.41 GiB | 16.00 GiB per OOM attempt |
| HDD | 17 | 204.67 | 96.57 | 6.99 GiB | 3.01 GiB |
| HDD | 540 | 674.01 | 2410.14 | 4.98 GiB | 16.00 GiB |
| HDD | 568 | 719.20 | OOM | 5.15 GiB | 16.00 GiB per OOM attempt |

These are current-version observations under this protocol. Historical
experiments used different cache and budget conditions, so the comparison
does not isolate the effect of a code change. The tools use different
library and archive layouts; this experiment compares their existing
end-to-end annotation implementations without rebuilding either layout.

`memory.peak` is whole-cgroup charged memory including page cache, memfd
and kernel memory. GNU time RSS, retained in the local evidence, is the
maximum individual process RSS, not simultaneous group RSS.

## Correctness

All six comparisons processed 3,933,714 rows and observed 3,875,208 immkv
hits, with zero value mismatches and zero hit disagreements. Float
tolerance was 1.1e-6; maximum observed difference was 1e-6. Values below
apply to each medium.

| Fields | Field comparisons | echtvar i32 overflow losses | True float -1 candidates | Reference mode |
|---:|---:|---:|---:|---|
| 17 | 66,873,138 | 0 | 0 | fresh echtvar digest equivalent to historical reference |
| 540 | 2,124,205,560 | 808 | 663 | fresh echtvar digest equivalent to historical reference |
| 568 | 2,234,349,552 | 808 | 663 | historical-only after fresh echtvar OOM |

Total across six comparisons: 8,850,856,500 field comparisons. Overflow
losses and true-minus-one candidates remain in their existing separate
categories; candidates are included in agreed misses, not added as a new
mismatch class. Complete producer exits, gzip-member CRC/EOF and stream
row counts are in the local per-run records. Tool-output agreement does
not independently prove every semantic detail of the original VCF.

## Protocol

- CPUs 0-3 on four distinct physical cores; GOMAXPROCS=4.
- `systemd-run --user --scope` with MemoryMax=16G and MemorySwapMax=0.
- Root `sync` and a successful write of `3` to
  `/proc/sys/vm/drop_caches` before every independently timed invocation.
- immkv Threads=4, batch=4096 and automatic read-path selection.
- echtvar four chromosome tasks inside the same constrained scope.
- Full-output correctness before three alternating timing pairs:
  immkv/echtvar, echtvar/immkv, immkv/echtvar.
- Existing libraries and archives reused; no format, magic, structure or
  file-count changes. The tested tag is not moved and no release is created
  by the runtime experiment itself.

The explicit user-selected 16 GiB budget supersedes the repository's
generic 10 GiB experiment budget. The measured devices were KINGSTON
SNV3S2000G on ext4 for NVMe and SEAGATE ST6000NM0034 on ext4 for HDD.
Only the OS page cache was cleared; device/controller cache emptiness is
not claimed.

## Recorded limitations

- The host does not delegate cgroup `io.stat`; cgroup I/O byte deltas are
  unavailable. GNU time file-system input blocks are retained instead of
  invented bytes.
- Six echtvar / 568-field timing attempts OOMed within 16 GiB. Their
  failed-scope elapsed values are not lookup speed samples.
- Two fresh echtvar / 568-field correctness attempts OOMed. Historical
  complete outputs serve only as correctness references; no fresh-output
  equivalence is claimed for 568 fields.
- During NVMe / 568-field fresh correctness OOM termination, direct-child
  completion metadata was not persisted. The final audit validates cache,
  limits, baseline, counter and journal OOM evidence and the exact
  wrapper/scope/journal launch identity; precise engine-start time and
  child argv remain unavailable and are not invented.
- One HDD / 568-field failed correctness comparison/driver cleanup extended
  0.607921622 seconds beyond the next correctness cache completion. All old
  annotation subprocesses had exited before that cache request. There was
  no overlap among the 36 timing invocations; strict all-driver correctness
  quiescence is not claimed.
- Successful correctness scope status has a generic `timing_sample:true`
  metadata bug; timing summaries use only `timing-*` paths plus
  workload-specific metadata.
- Retained cgroup counter descriptors can return `ENODEV` after systemd
  removes a scope. Successful runs retain complete acknowledged
  pre-removal counters.
- Generic status files and partial audit snapshots from the run are kept
  locally rather than rewritten.

## Scripts and reproduction order

| Script | Purpose |
|---|---|
| `build_tools.sh` | Freeze the tagged source, build annotate and verifyvalues, record hashes and run targeted tests |
| `preflight.py` | Validate the six queries, archive central directories, field tables and environment |
| `cache_service.py` | Finite authenticated helper for root sync and global cache drop |
| `common.sh`, `env_tools.py`, `run_limited.sh` | Scope constraints, observer handshake, counters, cache proof and status resolution |
| `run_cell.py` | One immkv or four-worker echtvar cell with bounded compressed output |
| `verify_output.py`, `verify.sh` | Full row/field verification and fresh echtvar digest proof |
| `run_matrix.sh` | Six cells, correctness then three alternating timing rounds |
| `run_when_authenticated.sh` | Wait for the authenticated cache service and run the matrix once |
| `summarize.py`, `report_acceptance.py` | Aggregate local evidence and audit completion |
| `verify_rss_accounting.py` | Reproduce the non-benchmark RSS accounting control on CPUs 4-7 |

See [notebook.md](notebook.md) for the detailed execution order. The final
summary is available as [Markdown](results/final-summary.md) and
[JSON](results/final-summary.json).

## Checks

Both tools were built from the tagged source with Go 1.27.1 and
CGO_ENABLED=1; targeted annotate and verifyvalues tests passed. Default
`go build ./...` and `go test ./...` passed after the experiment. The
engine source was not modified; this commit only adds experiment tooling
and records. The full codec/platform functional matrix for the exact
commit had already passed before this performance task and was not
repeated.
