# immkv v0.0.3 versus echtvar v0.2.4: HDD / NVMe, 4C16G

The six-cell comparison is complete. In this compressed-output annotation
workload, immkv leads on NVMe at 17 and 540 fields, while echtvar leads on HDD
at those field counts. At all 568 fields, immkv completes on both media and
echtvar is killed by the 16 GiB memory cgroup on both. Cache state is inherited
and uncontrolled; these are not cold-cache or isolated device benchmarks.

## Controlled facts

| Item | Configuration | Source |
| --- | --- | --- |
| immkv | Release `v0.0.3`, commit `840d36d80351f8b8651d9dc90abf1314bf5db21c` | `results/build/source-facts.json` |
| Build | Go 1.27.1, default CGO enabled; isolated release source | `results/build/environment.txt`, `results/build/build-info.txt` |
| immkv executable SHA-256 | `4179f6f9bf774615ce68e1794dd3c46f2c5e4b2316011b117191584f5ec16391` | `results/build/binaries.sha256` |
| Rival | echtvar 0.2.4 | `results/build/environment.txt` |
| Query slicing | bcftools 1.25 / htslib 1.25 | `results/build/environment.txt` |
| Limits | Four physical cores, CPUs 0-3, GOMAXPROCS=4, MemoryMax=16G, MemorySwapMax=0 | `common.sh`, every actual run's `constraints.json` |
| Effective memory | 17,179,869,184 bytes, zero swap | `constraints.json` in each run directory |
| Dataset | hg19 / gnomAD v2.1.1 genomes; decomposed HG001 query | `results/input-preflight.json`, `common.sh` |
| Query scope | 3,933,714 rows, 3,875,208 immkv hits | `results/verify-*-release/verification.json`, each `immkv.log` |
| Data-row identity | All six query copies have non-header SHA-256 `0250956380e1f6910a83f92a579810671e1fe4dfeb0b1d855394021a87148fd1`; their headers differ | `results/input-preflight.json` |
| Field definitions | The existing 17 / 540 / 568 TSVs; 17 uses the narrow library, 540/568 use the full library | `common.sh`, `results/input-preflight.json` |
| NVMe | KINGSTON SNV3S2000G, ext4, `/home/tian/immkv_test_nvme` | `results/hardware-final.txt` |
| HDD | SEAGATE ST6000NM0034, ext4, `/mnt/data6t/tian/immkv_test_hdd` | `results/hardware-final.txt` |
| Preflight | All 132 archive central directories inspected; library manifests and file sizes checked | `results/input-preflight.json` |

The explicit 4C16G user requirement overrides the repository's generic 10G
experiment setting for this experiment. Build phases are outside the cap;
every actual correctness or timing invocation asserts its effective limits.
Production source, existing libraries, archives and old results remain unchanged.

The executable build metadata contains `vcs.modified=true`: the isolated
`git archive` build ran beneath the surrounding working repository. Provenance
therefore rests on the exact release source hashes and executable hash, not on
that dirty flag. The root engine files, annotate source and module files were
checked against the recorded source snapshot and release commit.

## Timing boundary

This is an end-to-end annotation comparison with compression and compressed
bytes discarded, not a raw KV lookup benchmark. immkv reads the query, builds
batches, queries, formats typed INFO, and writes BGZF using its standard
annotate tool with Threads=4 and batch=4096. A scratch pathname ending in
`.vcf.gz` points to `/dev/null`, preserving the BGZF writer selected by suffix.
Echtvar's writer remains BGZF-compressed when its output is `/dev/null`.

Echtvar runs four chromosome tasks on the same four-core CPU affinity. Every
task slices its chromosome using `bcftools view -r N -Oz`, then calls
`echtvar anno -e chrN.echtvar.zip SLICE /dev/null`. Slicing, slice compression,
annotation and output compression are inside the timed invocation. The slice
is a short-lived memfd accessed through a scratch symlink; it is not a stored
intermediate on the source disk. Each echtvar process also creates HTSlib
threads, so four tasks is not a claim of four total OS threads.

Correctness runs precede each cell's timing. Successful cells have three rounds
in orders immkv/echtvar, echtvar/immkv, immkv/echtvar. All runs are sequential,
and timing contains no reference decompression, field comparison or hashing.
Full-field echtvar OOM is attempted independently on each medium during both
correctness and timing; later timing repetitions are explicitly skipped.

`/usr/bin/time -v` is inside each scope around the entire driver process tree.
Wall, user and system CPU times include child work. GNU maximum RSS is a maximum
child measurement, not aggregate simultaneous RSS; aggregate cgroup memory.peak
is reported separately and includes charged file cache and memfd slices.

Protocol implementation: `run_cell.py`, `run_limited.sh`, `common.sh`,
`run_matrix.sh`. This differs from historical tests retaining full gzip outputs
and disk-backed chromosome slices; historical and current absolute timings
must not be pooled. It also fixes the older null-phase compression asymmetry:
plain immkv `/dev/null` would disable BGZF, while echtvar would still compress.

## Measured wall time

Median of three successful complete runs; ranges contain all three samples.
The ratio is computed from the two median wall times. OOM has no completed
throughput or speed ratio.

| Medium | Fields | immkv seconds, range | echtvar seconds, range | Result |
| --- | ---: | ---: | ---: | --- |
| NVMe | 17 | **8.17**, 8.16-8.37 | 24.52, 24.42-24.52 | immkv 3.00x faster |
| NVMe | 540 | **157.57**, 157.55-157.76 | 606.40, 603.36-643.97 | immkv 3.85x faster |
| NVMe | 568 | **168.65**, 168.05-168.76 | OOM at 16 GiB | immkv completes |
| HDD | 17 | 64.66, 64.57-64.76 | **24.53**, 24.53-24.55 | echtvar 2.64x faster |
| HDD | 540 | 1,416.31, 1,410.04-1,428.45 | **621.93**, 616.84-631.26 | echtvar 2.28x faster |
| HDD | 568 | **1,413.10**, 1,412.24-1,422.31 | OOM at 16 GiB | immkv completes |

Sources: `results/table-final.json`, `results/summary-final.md`, and
`results/timing-MEDIA-FIELDS-ENGINE-r{1,2,3}-release/time.txt` with successful
`status.json` and `cell.json`. The generated summary links every number to its
individual source and retains per-pair ratios as well as median ratios.

There are 30 successful timing runs, two OOM timing attempts and four
`skipped_after_oom` slots. Skipped slots are not samples. All six immkv cells
have three successful timing runs; echtvar has three for each 17/540 cell.
Sources: per-run `status.json`, the NVMe OOM `status-resolution.json`, and
`results/summary-final.md`.

## CPU and aggregate memory

CPU values are the median of user+system CPU per run, not the sum of separately
rounded median columns. Memory is median aggregate cgroup peak per successful
run; it is not private RSS or a claim of total corpus cache residency.

| Medium | Fields | immkv CPU seconds | echtvar CPU seconds | immkv cgroup peak GiB | echtvar cgroup peak GiB |
| --- | ---: | ---: | ---: | ---: | ---: |
| NVMe | 17 | 30.02 | 95.18 | 0.490 | 0.127 |
| NVMe | 540 | 608.52 | 2,393.47 | 2.400 | 16.000 |
| NVMe | 568 | 655.56 | No completed run | 3.820 | OOM at cap |
| HDD | 17 | 200.51 | 95.38 | 0.331 | 0.127 |
| HDD | 540 | 664.53 | 2,389.84 | 1.503 | 16.000 |
| HDD | 568 | 707.05 | No completed run | 1.052 | OOM at cap |

Sources: `results/table-final.json`, underlying `time.txt`, `status.json`,
`memory.peak` and `cgroup-final.json`. An aggregate 16 GiB peak on successful
540-field echtvar runs does not mean 16 GiB of anonymous application memory:
charged file cache can reach the cap and be reclaimed without an OOM kill.

Both wide HDD immkv workloads use much less CPU than echtvar's 540-field run,
but spend far longer waiting for storage. The HDD narrow immkv workload also
uses more CPU than its NVMe counterpart; the read-path and media selection
need separate profiling before attributing that extra CPU to a specific cause.
No implementation was changed in response to these measurements.

## Correctness

Fresh immkv BGZF output is fully decompressed and compared with retained echtvar
references through the shared `test/verifyvalues`, at float tolerance 1.1e-6.
Every row and selected field is compared. Independent stream row counts,
complete gzip member CRC/EOF checks, producer/verifier exit checks and stream
hashes guard against truncated delivery or scanner limitations.

| Medium | Fields | Rows | Field comparisons | Extra value differences | Hit disagreements | echtvar overflow losses |
| --- | ---: | ---: | ---: | ---: | ---: | ---: |
| NVMe | 17 | 3,933,714 | 66,873,138 | 0 | 0 | 0 |
| NVMe | 540 | 3,933,714 | 2,124,205,560 | 0 | 0 | 808 |
| NVMe | 568 | 3,933,714 | 2,234,349,552 | 0 | 0 | 808 |
| HDD | 17 | 3,933,714 | 66,873,138 | 0 | 0 | 0 |
| HDD | 540 | 3,933,714 | 2,124,205,560 | 0 | 0 | 808 |
| HDD | 568 | 3,933,714 | 2,234,349,552 | 0 | 0 | 808 |

Total: 8,850,856,500 field comparisons. Sources:
`results/verify-MEDIA-FIELDS-release/verification.json` and `verifier.stdout`.
Maximum observed float absolute difference is 1e-6 in all six runs, within
1.1e-6. The tolerance derives from echtvar v0.2.4's quantization truncation,
printed-decimal truncation and f32 slack; integers and strings compare exactly,
Flags by presence. See `test/verifyvalues/main.go` and the verifier logs.

The 808 wide-field losses are the rival's documented i32 quantization overflow,
reported separately from extra differences. Each wide immkv output contains 663
selected Float values equal to -1, versus zero in the narrow set. These are
marker-collision candidates: the legacy verifier folds true -1/missing-marker
ambiguity into misses agreed, so this count is not an independently reconciled
collision count. Source: `streams.immkv.selected_float_true_minus_one_candidates`
and `marker_collision_accounting` in each `verification.json`.

Fresh echtvar 17/540 outputs also pass per-chromosome full data-row SHA-256 and
row-count checks against retained references, on both media: 22 chromosomes and
3,933,714 rows per proof. Sources:
`results/verify-echtvar-MEDIA-{17,540}-release/cell.json` and `chrN.json`.
Fresh full-field echtvar fails under the current cap; those attempts are not
correctness passes. The historical 568 reference was generated without the
current cap and is used solely as a correctness reference, never as a timing
sample. Both media's fresh immkv data-row hashes agree for each field set;
header bytes and numerical representation differ from echtvar by design.

## Full-field OOM evidence

Four independent echtvar full-field attempts fail under the actual 16 GiB
scope: correctness and timing on NVMe, then correctness and timing on HDD.
The corresponding unit journals explicitly report an OOM kill. NVMe kernel
logs also identify `CONSTRAINT_MEMCG`, the matching scope and task `echtvar`.

Sources:

- `results/verify-echtvar-nvme-568-release/{status.json,status-resolution.json,systemd-unit.log,kernel-oom.log}`.
- `results/timing-nvme-568-echtvar-r1-release/{status.json,status-resolution.json,systemd-final.log}`.
- `results/verify-echtvar-hdd-568-release/{status.json,systemd-unit.log}`.
- `results/timing-hdd-568-echtvar-r1-release/{status.json,systemd-unit.log}`.

Some OOMs terminate the wrapper before it can read final counters. The outside
observer's last memory.events can still show oom_kill=0; it was sampled before
the kill, and is not treated as a final counter. The explicit unit journal and
kernel evidence resolve that ambiguity. Original `scope_failed`/143 records
are retained with separate status-resolution evidence, rather than overwritten.
OOM peak observations reach the cap but have no completed-time median.

## Cache, I/O and interpretation limits

Cache state is not reset or isolated. Correctness precedes timing; OS page cache
can retain data across processes and scope boundaries. File pages charged to an
earlier cgroup are not necessarily reflected as new private memory in a later
scope. Thus a 16 GiB current scope limit is not proof that all useful cache data
fit inside that invocation's measured memory peak.

This matters especially for 17 fields: echtvar reports zero GNU filesystem-input
blocks in all six NVMe/HDD timing runs. HDD immkv reports 3,070,360 blocks per
run; NVMe immkv reports 11,693,048. These are warm/inherited-cache results with
asymmetric actual storage work, not a cold HDD versus cold SSD comparison.
Sources: `filesystem_inputs_blocks` in `results/summary-final.md` and matching
`time.txt`. The small echtvar peak must not be interpreted as the total
application-plus-cached-archive memory needed from a cold start.

For 540 fields, GNU filesystem-input block samples also vary as cache state
changes. NVMe echtvar: 90,444,408 / 59,958,072 / 39,519,488 blocks; HDD echtvar:
96,294,872 / 73,387,984 / 58,645,824. HDD immkv: 46,185,256 / 47,778,944 /
44,801,872. These values are retained as I/O context, without claiming exact
per-library physical bytes or identical cache residency.

The annotate tool's internal `read 0.0GiB` statistic is not physical I/O and is
not used to conclude that a run read nothing; GNU inputs show otherwise.
Runtime record-order/columnar and kernel read-path selection remain automatic,
not forced for this comparison. No additional path profiling was performed.

The timing excludes retained output disk writes and uses memfd slice
intermediates; it includes each tool's own compression and formatting, which
are not identical implementations. Accordingly the ratios compare these
complete annotation workflows under this protocol. They do not establish a
universal engine speed ratio, a controlled cold-cache device ratio or a
full-output-to-disk time ratio. No new build-throughput or random-point-query
claim is made.

## Conclusion and reproduction

For this observed workload and cache regime, immkv v0.0.3 is faster in the two
completing NVMe paired cells and uses less CPU for wide-field annotation;
echtvar is faster in the two completing HDD paired cells. immkv completes
all-field annotation under 16 GiB on both media, while echtvar does not.
The HDD read-path cost is a concrete performance limitation for future work.

The scripts, source and executable hashes, raw GNU reports, cgroup records,
complete verification logs, OOM evidence and generated summary are retained
here. Launch order and resumption instructions are in `notebook.md`; detailed
per-number evidence is in `results/summary-final.md` and `table-final.json`.
Old experiments remain unchanged, production code was not modified, and this
new experiment has not been committed or pushed.
