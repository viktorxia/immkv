# Projection experiment notebook

## Purpose and decision boundary

Measure the cost of explicit `StreamOptions.Columns` and verify an optimization
against both the previous frozen candidate and commit `9e72562`. Keep query
loading, INFO formatting, compression, CPU profiling, and correctness hashing
out of the default lookup timing. No point-lookup or build-throughput work is
planned here. Production changes and readiness are owned by the parent task.

The user authorized the new 4 CPU / 16 GiB / zero swap regime. All writes belong
to this experiment directory or its repository-local scratch directory. Existing
libraries, the previous experiment, and its frozen snapshots are read-only.
Never start a build, test, profile, digest, or verification process while a timed
series is running. The parent coordinates CPU and disk use.

## Status ledger

- Completed: old/before/after frozen builds, both-codec full production builds/tests, vet, race checks and cross-platform pure-Go builds.
- Completed: four complete whole-genome old/after annotation comparisons and all five before/after raw-record digest comparisons.
- Completed: before/after full540 whole-genome CPU profiles and their retained analysis.
- Completed: three-round whole-genome narrow17 timing; wide-library WGS and repeated chr22 disk matrices were interrupted when NVMe service time changed sharply, with all partial/noisy records retained.
- Completed: the separate memory-resident chr22 diagnostic, including preloaded byte hashes, all raw digests, five three-version query workloads and four old/after annotation workloads.

Exact measured values, sources and limitations are in `README.md`. Every
`finished`, `exit-status`, raw output and constraint record remains the
completion evidence; commands below describe reproduction, not additional runs.

## Build order

All commands below assume Bash. Paths are absolute so the working directory
does not affect source selection.

```bash
EXPERIMENT=/home/tian/projects/immkv/test/vs_immkv_hg19_hg001_4c16g_projection_20261115
SCRATCH=/home/tian/projects/immkv/tmp/vs_immkv_hg19_hg001_4c16g_projection_20261115
bash "$EXPERIMENT/build_versions.sh" old before
# Only after the parent has frozen the production changes:
bash "$EXPERIMENT/build_versions.sh" after
```

These builds were already completed in the initial preparation and will refuse
to replace their evidence. `before` imports the previous experiment's frozen
candidate and validates every original manifest hash, including the original
query and build tools, before accepting the required original digest. The new
source digest differs because this experiment has its own query harness and
includes only the needed executables. `old` checks its imported production and
annotation source bytes against the specified commit. `after` includes every
current root production `.go` file, including untracked files, and the current
`test/annotate` source. It is never built by default.

Source snapshots copy no production test files. Their tests cover the query
harness's batch grouping, repeat digests, mismatch handling, key filtering and
sorting. Full production verification remains a separate parent responsibility.
Dependencies are seeded from the existing read-only module cache into scratch;
all build caches, compiler temporary files, and telemetry configuration stay
local or disabled. Missing dependencies fail instead of downloading.

## Round 1: complete annotation correctness

Run this before interpreting timing results. No large text is retained, and
all annotation rows and selected fields are compared.

```bash
RUN_TAG=wgs RUN_TIMEOUT_SECONDS=1800 \
  bash "$EXPERIMENT/verify_stream.sh" old before after
# Or run only the comparison currently needed:
RUN_TAG=wgs-after RUN_TIMEOUT_SECONDS=1800 \
  bash "$EXPERIMENT/verify_stream.sh" old after
```

`WORKLOADS` can select `narrow17 full17 full540 full568`. Each pair of annotation
producers and the shared verifier occupies one aggregate 16 GiB / zero swap
scope on CPUs 0-3. This is an untimed correctness phase, so simultaneous producer
I/O cannot be used as a performance comparison. The verifier checks tolerance
zero, complete SHA-256 equality, independent output/query row counts,
rows-times-fields comparison counts, matching producer hits, and successful
producer exits. FIFOs are removed after completion or failure; logs remain.

`full568_default` is a query-only control. Existing annotation always supplies
explicit columns, so the FIFO script does not pretend to test default annotation.
Use the query digest control below for default-record equality.

## Diagnostic raw-record equality

```bash
RUN_TAG=wgs RUN_TIMEOUT_SECONDS=1800 \
  bash "$EXPERIMENT/run_digest.sh" before after
# Restrict a diagnostic without reusing existing result names:
RUN_TAG=wgs540 WORKLOADS=full540 RUN_TIMEOUT_SECONDS=1800 \
  bash "$EXPERIMENT/run_digest.sh" before after
```

The default covers all five query workloads. `DIAGNOSTIC_REPEAT` requests
additional passes on the same handle. Hashing is diagnostic overhead, and these
runs have no GNU time report. `check_digest.py` requires identical raw records,
counts, and returned bytes for before versus after and checks repeat consistency.
Frozen source and binary hashes are checked before and after. If old is included,
`full17` and `full540` raw byte equality is explicitly excluded because old can
return nonprojected records; counts still must agree and FIFO output still must
match.

## CPU profiles and separate analysis

```bash
RUN_TAG=wgs RUN_TIMEOUT_SECONDS=1800 bash "$EXPERIMENT/run_profile.sh" before 1
RUN_TAG=wgs RUN_TIMEOUT_SECONDS=1800 bash "$EXPERIMENT/run_profile.sh" after 1
bash "$EXPERIMENT/analyze_profile.sh" before wgs
bash "$EXPERIMENT/analyze_profile.sh" after wgs
```

Each profile covers only the full540 lookup loop and its workers. Profiles are
retained in the result directory after success; partial profiles and the scratch
copy remain available if a process fails. `retain_profile.py before wgs` safely
retains an already completed earlier invocation. Flat and cumulative `pprof`
reports, plus matching projection source listings, are saved independently.
Do not interpret a profiled lookup time as an unprofiled matrix sample.

The existing read-only chr22 fixture supports short repeat diagnostics:

```bash
QUERY=/home/tian/projects/immkv/tmp/vs_immkv_hg19_hg001_10g_release_20261114/query-chr22.vcf \
  RUN_TAG=chr22 DIAGNOSTIC_REPEAT=10 RUN_TIMEOUT_SECONDS=600 \
  bash "$EXPERIMENT/run_digest.sh" before after
QUERY=/home/tian/projects/immkv/tmp/vs_immkv_hg19_hg001_10g_release_20261114/query-chr22.vcf \
  RUN_TAG=chr22 RUN_TIMEOUT_SECONDS=600 \
  bash "$EXPERIMENT/run_profile.sh" before 10
```

Keep chr22 evidence separate from whole-genome evidence. Repetition reduces
startup noise but reuses a handle and the OS cache state; it does not establish
a cold-cache performance result.

## Round 2: unprofiled query-only timing

```bash
RUN_TAG=wgs RUN_TIMEOUT_SECONDS=1800 \
  bash "$EXPERIMENT/run_null.sh" old before after
```

The default matrix runs `narrow17`, `full17`, `full540`, `full568`, and
`full568_default`, each with three rounds. Within each workload the order is old,
before, after; after, before, old; old, before, after. The records themselves go
to no output sink; only a small JSON statistics record is retained. No profile or
digest is enabled. GNU time measures the entire process in the scope and the
harness separately reports setup and lookup time.

For a short control with repeated complete passes:

```bash
QUERY=/home/tian/projects/immkv/tmp/vs_immkv_hg19_hg001_10g_release_20261114/query-chr22.vcf \
  RUN_TAG=chr22 QUERY_REPEAT=10 WORKLOADS='full568 full568_default' \
  RUN_TIMEOUT_SECONDS=600 bash "$EXPERIMENT/run_null.sh" old before after
```

The report normalizes lookup seconds per pass and retains total process CPU/RSS.
Do not compare different repeat counts or query tags as paired observations.

## Optional annotation timing

```bash
QUERY=/home/tian/projects/immkv/tmp/vs_immkv_hg19_hg001_10g_release_20261114/query-chr22.vcf \
  RUN_TAG=chr22 RUN_TIMEOUT_SECONDS=1800 \
  bash "$EXPERIMENT/run_annotation.sh" old after
```

This additional three-round paired chr22 matrix uses existing `test/annotate` with
`-out /dev/null` for the four explicit workloads. It includes parsing and typed
formatting and excludes gzip and stored text I/O. Run after query-only timing so
the main measurement remains identifiable.

## Resident chr22 diagnostic after disk variability

The disk runs encountered changing NVMe service time; see
`results/io-observation.md`. The separate resident diagnostic copies both
chr22 libraries, the same query and fields into sealed memfds, verifies full
byte hashes and library metadata, and then runs the unchanged binaries. All
resident input memory is charged to the same asserted 16GiB/swap0 scope.
Source loading is outside every query timer. Proxy symlinks are temporary and
are removed together with the anonymous memfds on completion.

```bash
RUN_TIMEOUT_SECONDS=180 MEMORY_TIMEOUT_SECONDS=1800 \
  bash "$EXPERIMENT/run_memory_chr22.sh"
python3 -I "$EXPERIMENT/summarize_memory.py"
```

The initial run completed under `results/memory.chr22/`; it cannot be overwritten.
Use a new RUN_TAG for a later run and explicitly adapt the summary input path.
The driver performs raw-record checks for all three versions, excluding old
subset raw equality, then three rounds of ten query passes for all five
workloads, followed by three old/after annotation rounds for four workloads.
This is a memory-resident diagnostic and is never pooled with NVMe results.
`results/summary.memory.chr22.json` contains the validated medians and source
paths. Whole-genome output correctness remains the earlier FIFO evidence.

## Reporting and reconstruction

```bash
python3 -I "$EXPERIMENT/summarize.py" --tag wgs --format json \
  --output "$EXPERIMENT/results/summary.wgs.final.json"
python3 -I "$EXPERIMENT/summarize.py" --tag wgs --format markdown \
  --output "$EXPERIMENT/results/report.wgs.final.md"
python3 -I "$EXPERIMENT/source_snapshot.py" restore before
```

Reporting never overwrites an existing report or README. For progress snapshots,
use stdout or choose another output filename. Missing and failed runs are named
explicitly; timing medians and paired changes require three complete rounds with
matching query and hit counts and validated scope constraints. CPU profiles,
raw digests and annotation verification retain their own evidence paths.

Source restoration creates only
`$SCRATCH/restored/<version>` and validates every retained byte against the
manifest. Build a restored module with experiment-local caches using the same
`go build -trimpath` commands recorded in `results/build.<version>/build.log`;
write binaries to a new scratch target. Source JSON retains exact bytes,
including the old options adaptation and untracked after files, so deleting the
original ignored scratch snapshot does not lose source provenance.

Report interpretation must distinguish the lookup timer from whole-process
metrics and profile samples. NVMe service time and cache state are not held
constant by the CPU/memory limits. Preserve all samples and scope failures;
never turn a noisy wall-time median into a CPU-only optimization claim.
