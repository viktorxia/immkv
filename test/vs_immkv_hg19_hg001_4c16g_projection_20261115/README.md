# Query projection optimization: design and measured results

The field-selection design is retained. The optimized implementation removes
the explicit-column batch regression in the resident chr22 comparison and
outperforms the pre-refactor version in all five measured query workloads.
Full-genome annotation outputs and before/after raw records are identical.
Whole-genome wide-library NVMe wall-time conclusions remain limited by severe
changes in device service time; interrupted and noisy runs are retained.

## Design and implementation

`StreamOptions.Columns` remains a logical selection independent of record-order
versus columnar path. Both return the same schema-shaped v003 record, with
unselected fields absent. This is a sound API contract; the previous serial
post-decompression projection was the avoidable cost.

The optimized path compiles the selection once, projects inside decompression
workers, copies contiguous selected payload spans once, bypasses projection for
all-field selection, decodes directly into call-owned output storage, and reuses
safe scratch/input buffers. Duplicate keys reuse read-only results. Columnar
assembly avoids intermediate payload copies. Buffer capacity/alignment errors,
selected-column tail validation, ring cleanup and the shared Threads budget are
also addressed. Details and validation boundaries: `design-review.md`.

Public function signatures, disk layout, format magic and codec frame format
are unchanged. Existing v003 libraries are the inputs to all recorded runs.
Unprojected record-order results, including explicit all-field selection, follow
the default query path and receive typed validation through `NewRecord`.
Partial projection still validates skipped payloads and trailing bytes.

## Controlled facts

| Item | Value | Source |
| --- | --- | --- |
| old | `9e72562966c4caa1b2c813719aadb92033eede85`, before the API/error-channel refactor | `results/build.old/source-facts.json` |
| before | Original frozen working-tree digest `da913ce83951b2e00998343af092321fab16fcd433c7328b58cd91b60376c620` | `results/build.before/source-facts.json` |
| after | Exact source snapshot digest `cffc74c7403d8b89072d2a1b3e61e7fa3606fd37719e70d49c585c761efc9172` | `results/build.after/source-facts.json` |
| CPU / memory | Physical cores 0-3, GOMAXPROCS=4, Threads=4, MemoryMax=16G, MemorySwapMax=0 | `common.sh`, each invocation's `cgroup.txt` |
| Batch | 4,096 keys, extending same-position groups | `query-only/main.go` |
| Dataset | hg19 / gnomAD v2.1.1, decomposed HG001 query, existing narrow17 and full568 libraries | `common.sh`, `results/build.*/environment.txt` |
| Whole-genome scope | 3,933,714 rows; 3,933,369 buildable query keys; 3,875,208 hits | `results/verify.wgs.*/verification.json`, `results/digest.wgs.*/stdout` |
| Chr22 scope | 49,540 rows; 49,536 query keys; 48,709 hits | `results/memory.chr22/inputs.json`, `results/memory.chr22/query.*/stdout` |
| Compiler and hardware | Exact compiler, CPU topology/model, filesystem and NVMe model retained | `results/build.*/environment.txt`, `results/environment-context.txt` |

The explicit user requirement of 4C16G overrides the repository's usual 10G
experiment limit for this experiment. Setup/builds are not performance samples.
All query, profile, correctness and memory-diagnostic runs assert their actual
resource limits. Each result directory is exclusive; earlier records are never
replaced. Exact compiled sources, including untracked production files, are
retained as ASCII JSON with per-file hashes in `results/build.*/sources.json`.

## Resident chr22 query comparison

Three alternating rounds per version, ten complete query passes per process.
The table reports the median per-pass lookup time; setup, source preload,
formatting, hashing and CPU profiling are excluded. All eight chr22 library
files plus query/field inputs are copied and SHA-256 verified into sealed Linux
memfds before querying. The same unchanged binaries and automatic path selection
read the resident bytes through filesystem handles. This is a memory-resident
query diagnostic, not an NVMe throughput result.

| Workload | old ms | before ms | after ms | after vs old | after vs before |
| --- | ---: | ---: | ---: | ---: | ---: |
| narrow17: all 17 fields | 42.426 | 67.987 | 33.415 | -21.24% | -50.85% |
| full17: select 17 of 568 | 166.175 | 621.020 | 147.868 | -11.02% | -76.19% |
| full540: select 540 of 568 | 166.176 | 796.735 | 154.570 | -6.98% | -80.60% |
| full568: explicit all fields | 167.901 | 847.561 | 140.802 | -16.14% | -83.39% |
| full568_default: Columns=nil | 167.517 | 170.109 | 139.280 | -16.86% | -18.12% |

Sources: `results/summary.memory.chr22.json` query entries, with every underlying
sample linked to `results/memory.chr22/query.<workload>.<version>.r{1,2,3}/stdout`.
The optimized explicit-full range is 135.83-140.81 ms; default is 137.73-142.24 ms,
consistent with both using the same path. Source: the same three-round entries.

Old record-order subset queries returned unselected fields too. Thus old subset
measurements provide historical context, while before versus after is the
strict same-result optimization comparison. The old and optimized annotation
outputs still match exactly; before/after raw-record hashes also match.

Resident inputs total 5,233,767,230 bytes. Scope peak memory was 11,986,907,136
bytes (about 11.16 GiB), including memfds and source page-cache charges; OOM and
OOM-kill counters are zero. Sources: `results/memory.chr22/inputs.json` and
`results/memory.chr22/cgroup.txt`. Memfds are released and proxy links removed
on completion (`resident.cleaned.json`); there are no retained multi-GB copies.

## Resident chr22 annotation including formatting

Three alternating old/after runs, complete query parsing, lookup and typed INFO
formatting; output goes to `/dev/null`. Preload is excluded; gzip is absent.

| Workload | old seconds | after seconds | change |
| --- | ---: | ---: | ---: |
| narrow17 | 0.091 | 0.087 | -4.40% |
| full17 | 0.528 | 0.210 | -60.23% |
| full540 | 1.032 | 1.043 | +1.07% |
| full568 | 1.187 | 1.092 | -8.00% |

Sources: `results/summary.memory.chr22.json` annotation entries and
`results/memory.chr22/annotation.<workload>.<version>.r{1,2,3}/stdout`.
Full540 formatting is effectively unchanged within this small observed spread;
the experiment does not claim every complete annotation workload got faster.

## Whole-genome NVMe measurements and limits

The narrow17 series completed all three paired rounds with consistent timings.

| Version | lookup median seconds | process user CPU seconds | peak RSS MiB |
| --- | ---: | ---: | ---: |
| old | 4.844 | 9.32 | 522.0 |
| before | 6.920 | 11.75 | 503.1 |
| after | 3.987 | 6.60 | 510.4 |

Sources: `results/summary.wgs.disk.json` narrow17 entries and
`results/null.wgs.narrow17.<version>.r{1,2,3}/{stdout,time.txt}`. The optimized
lookup median is 17.70% below old and 42.39% below before.

The full17 series encountered severe NVMe variability: old lookup rounds were
21.854, 52.470 and 130.250 seconds. The first two after rounds were 16.333 and
16.380 seconds; before was 66.555 and 68.318 seconds. The following before round
3 was interrupted with SIGTERM and is not a completed sample. Later wide-library
whole-genome timing cells were not run. Sources: `results/null.wgs.full17.*`,
`results/summary.wgs.disk.json`, `results/io-observation.md`.

The first two full17 process user CPU readings were stable: old 41.94/42.21s,
before 80.28/80.19s, after 32.88/32.90s (`results/null.wgs.full17.*.r{1,2}/time.txt`).
These support a CPU improvement independently of the unstable disk wall times.

A separate repeated chr22 disk series also encountered I/O variability and was
interrupted. Its narrow17 rounds completed, but full17 and later workloads are
incomplete (`results/summary.chr22.disk.json`). It is not substituted into the
resident comparison. Per-run system I/O pressure/diskstats snapshots are in
`io-context.*.json`; they describe host context, not per-process attribution.

## Lookup-only CPU profiles

Both full540 profiles consume the whole-genome query and return identical key,
hit and byte counts. They are diagnostic runs, excluded from timing medians.

| Metric | before CPU sample seconds | after CPU sample seconds | reduction |
| --- | ---: | ---: | ---: |
| Projection, cumulative | 43.92 | 6.93 | 84.22% |
| Total sampled CPU | 105.90 | 41.60 | 60.72% |

Sources: `results/analysis.wgs.full540.before/cumulative.txt` (`projectRecord`),
`results/analysis.wgs.full540.after/cumulative.txt` (`recordProjection.appendRecord`)
and profile headers. The after profile attributes 19.23s to dictionary decoding,
10.23s to the direct-read driver and 6.93s to projection; cumulative values must
not be added across ancestor/descendant functions. Wall times differed with
I/O conditions and are not used as a clean optimization ratio.

The profiles show `driveDirect`/`ioUring.submitWait` activity in these specific
runs. Other automatic-path invocations were not individually traced; the report
does not assume every run used the same kernel path.

## Correctness and code verification

Full old/after annotation streams were compared at tolerance zero using the
shared verifier, independent row counts, producer exit checks and whole-stream
SHA-256. All four complete whole-genome scenarios passed.

| Workload | rows | field comparisons | value/hit differences | SHA-256 equal |
| --- | ---: | ---: | ---: | --- |
| narrow17 | 3,933,714 | 66,873,138 | 0 | yes |
| full17 | 3,933,714 | 66,873,138 | 0 | yes |
| full540 | 3,933,714 | 2,124,205,560 | 0 | yes |
| full568 | 3,933,714 | 2,234,349,552 | 0 | yes |

Sources: `results/verify.wgs.<workload>.old-after/verification.json` and matching
`verify.log`. Total: 4,492,301,388 field comparisons. Full output is streamed via
FIFOs and is not retained. Independent hashes/counts guard the shared verifier's
legacy rival-specific exceptions and scanner limitations.

Before/after raw-record digests match for all five complete whole-genome
workloads (`results/digest.wgs.check/comparison.json`) and all five resident
workloads (`results/memory.chr22/digest-comparison.json`). Missing and zero-valued
records are hashed distinctly, with keys, lengths and presence markers included.

Default and `CGO_ENABLED=0` full builds/tests, `go vet`, the race suite, and
Darwin/Windows/FreeBSD pure-Go builds passed (`results/code-checks/*.log`,
`results/code-checks/finished`). The frozen per-version harness build checks
are separate evidence (`results/build.*/build.log`).

## Conclusion and scope

The avoidable projection regression is removed in the controlled resident
comparison, and the optimized query implementation exceeds the old query
performance there. The full-genome narrow-library NVMe comparison agrees.
Wide-library whole-genome CPU and correctness evidence is strong, while stable
wide-library NVMe wall-clock comparisons remain unestablished on this device.

The dominant remaining after-profile costs are zstd decompression and I/O;
this result is not a claim of a theoretical maximum or that every format/render
workload improves. No HDD, controlled cold-cache, point-lookup, gzip-output or
build-throughput claim is made in this experiment. Source and old experiment
records remain available; reproduction order and commands are in `notebook.md`.
