# Release performance regression: measured results

**Do not publish this candidate without addressing the explicit-column batch-query regression.**

## Controlled facts

Baseline: `9e72562966c4caa1b2c813719aadb92033eede85` (before the API/error-channel refactor).
Candidate: uncommitted working-tree snapshot, digest `da913ce83951b2e00998343af092321fab16fcd433c7328b58cd91b60376c620`.
Source: `results/source-facts.json`; tracked changes: `results/candidate.patch`.

Both versions read the same existing hg19 / gnomAD v2.1.1 v003 libraries.
All measured commands used CPUs 0-3, four workers, MemoryMax=10G and MemorySwapMax=0.
Sources: `common.sh`, `run_limited.sh`, `results/*.cgroup`.
Compiler, CPU topology and exact input hashes: `results/environment.txt`.

The full HG001 query has 3,933,714 rows, 3,933,369 buildable keys and 3,875,208 hits.
Sources: prior `../vs_echtvar_hg19_gnomad_genomes_4c16g_gz_v003/results/matrix.log` (input rows), and this run's `results/null.*.stdout` (buildable keys and hits). The key-only harness filters unbuildable alleles.

## Random point lookup

Median of three paired runs with the same seed and warmups; sampled frames are warm in the OS page cache.

| cache MiB | phase | baseline QPS | candidate QPS | QPS change | baseline p50 us | candidate p50 us | baseline p99 us | candidate p99 us |
|---:|---|---:|---:|---:|---:|---:|---:|---:|
| 0 | sampled | 350698 | 343357 | -2.09% | 9.767 | 9.838 | 20.146 | 19.425 |
| 0 | misses | 1537147 | 1496491 | -2.64% | 2.224 | 2.265 | 4.889 | 7.273 |
| 512 | sampled | 280635 | 282239 | +0.57% | 12.012 | 12.222 | 31.386 | 33.871 |
| 512 | misses | 2115186 | 2143165 | +1.32% | 1.162 | 1.162 | 8.095 | 7.975 |

Sources: `results/point.cache*.r*.stdout`; setup runs are excluded.

## Whole-genome GetBatch only

No INFO formatting or compression; query loading/sorting is excluded from lookup time.

| workload | pairs | baseline lookup s | candidate lookup s | time change | baseline peak RSS MiB | candidate peak RSS MiB |
|---|---:|---:|---:|---:|---:|---:|
| narrow17 | 3 | 4.451 | 6.342 | +42.50% | 528.4 | 507.4 |
| full17 | 3 | 19.627 | 55.575 | +183.16% | 671.0 | 641.9 |
| full540 | 3 | 19.546 | 66.967 | +242.62% | 675.2 | 672.9 |
| full568 * | 3 | 30.303 | 86.554 | +185.62% | 664.5 | 659.7 |
| full568_default | incomplete | - | - | - | - | - |

Sources: `results/null.*.stdout` and matching `.time` files; RSS includes the complete process.
* Full568 timings include a changed NVMe I/O regime; retain the raw samples but do not interpret the median ratio as stable code-only slowdown.
The following default-path whole-genome run was interrupted (signal 2) because disk throughput collapsed; its partial time is not a completed result.
Source: `results/io-regime-note.txt` and `results/null.full568_default.baseline.r1.time`.

Explicit subset output bytes differ: the old record-order path ignored Columns and returned full records, while the candidate performs the promised projection.
Selected annotation values are checked separately below.

| workload | baseline user CPU s | candidate user CPU s | CPU change |
|---|---:|---:|---:|
| narrow17 | 9.81 | 12.54 | +27.83% |
| full17 | 44.32 | 84.15 | +89.87% |
| full540 | 44.34 | 100.60 | +126.88% |
| full568 | 44.33 | 106.33 | +139.86% |

CPU source: median of `User time (seconds)` in `results/null.<workload>.<version>.r*.time`.

## Chr22 diagnostic

This separate subset has 49,540 rows, 49,536 buildable keys and 48,709 hits.
Sources: `results/chr22-query.sha256`, `results/diagnostic.chr22.*.stdout` and verification logs.

| call | baseline lookup s | candidate lookup s | time change | baseline user CPU s | candidate user CPU s |
|---|---:|---:|---:|---:|---:|
| default | 0.232 | 0.376 | +61.79% | 0.39 | 0.42 |
| explicit | 0.265 | 1.061 | +299.73% | 0.41 | 1.01 |

Sources: `results/diagnostic.chr22.*.stdout` and `.time`. Default passes Columns=nil; explicit selects all 568 fields.
These subsecond default results have substantial run-to-run variability; do not extrapolate their percentage to the whole genome.

## Chr22 annotation including typed formatting

Text output goes to /dev/null; this includes VCF parsing, lookup and formatting but excludes gzip and stored-output I/O.

| workload | baseline annotation s | candidate annotation s | time change |
|---|---:|---:|---:|
| narrow17 | 0.155 | 0.176 | +13.55% |
| full17 | 0.684 | 1.004 | +46.78% |
| full540 | 1.281 | 1.620 | +26.46% |
| full568 | 1.221 | 1.650 | +35.14% |

Sources: median `wall` statistic in `results/annotation.chr22.*.stdout`; whole-process CPU/RSS in matching `.time` files.

## Correctness

Both complete chr22 annotation streams were checked for every row and field using `test/verifyvalues` with tolerance zero, plus whole-stream SHA-256.

| workload | rows | field comparisons | mismatches | full-stream hashes equal | source |
|---|---:|---:|---:|---|---|
| narrow17 | 49540 | 842180 | 0 | yes | `results/verify.chr22.narrow17.log` |
| full17 | 49540 | 842180 | 0 | yes | `results/verify.chr22.full17.log` |
| full540 | 49540 | 26751600 | 0 | yes | `results/verify.chr22.full540.log` |
| full568 | 49540 | 28138720 | 0 | yes | `results/verify.full568.log` |

The initial full568 verification log lacks chr22 in its filename but used the same fixed chr22 query; its row count and producer statistics identify the scope.
Full-genome field-by-field verification has not been completed.

## Interpretation and release decision

- Point-lookup throughput is close to the baseline for this warm-cache sample; the no-cache miss tail is noisier and is retained in the table.
- Explicit Columns causes a large batch CPU regression, including the all-selected case where returned bytes are identical.
- The candidate StreamDB.GetBatch performs projection after joining all read/decompression workers; projectRecord walks fields and allocates additional buffers even when every field is selected.
- The chr22 explicit-versus-default diagnostic and higher user CPU support projection as a major contributor; no profile was collected, so the exact fraction is not established.
- Annotation with formatting still regresses in every tested chr22 workload; this is not solely an artifact of the query-only harness.
- Keep this candidate unreleased until the batch regression is addressed and the same controlled tests are repeated.

## Limits and unexecuted phases

- This is NVMe testing; no HDD performance claim is made.
- Cold point-lookup performance was not compared under controlled cold-cache conditions.
- Whole-genome round 3 experienced I/O disruption; raw files remain intact and CPU times are reported separately.
- Full stored-output gzip timing and synthetic build throughput were not run because authorization to write the external work directory was not received.
- No production library was modified, rebuilt or deleted; the library source code was not changed during measurement.
- Scripts and observed data are preserved here; historical experiments are unchanged.
