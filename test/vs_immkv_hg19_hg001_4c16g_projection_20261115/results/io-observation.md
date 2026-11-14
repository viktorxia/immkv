# Runtime I/O observation

During the full17 query-only series, old lookup_seconds changed from
21.854s (rounded; see the exact value in null.wgs.full17.old.r1/stdout) to
52.470s (see null.wgs.full17.old.r2/stdout). These rounded values are
observations, not an adjustment to either sample. All samples are retained.

A parent observation of /proc/pressure/io during the next old run reported
some avg10=26.90 and full avg10=26.62. Previous snapshots during annotation
verification had near-zero pressure. These are whole-host observations and
do not attribute the cause to code or another process.

The timed run wrapper records per-run io-context.before.json and io-context.after.json
where available. Read the exact raw timing and I/O reports rather than
assuming the version median represents a stable code-only wall-time ratio.
Process CPU and lookup-only profiles are reported separately. No cache
reset or disk-isolation claim is made.

Old full17 round 3 subsequently reached 130.25 seconds of lookup time (exact
value in null.wgs.full17.old.r3/stdout), with host full I/O pressure around 32%.
The following before round 3 was stopped with SIGTERM after roughly two minutes
of disrupted execution. Its nonzero exit status and partial time report are
retained and excluded from medians. Later whole-genome timed cells were not run.
Full-genome annotation verification, raw-record digests and CPU profiles had
already completed successfully before this interruption.

A separate chr22 diagnostic uses the same library and field selections, the
same 4C16G constraints, and repeated complete query passes per process. Those
numbers are reported only for that subset, never substituted into the
whole-genome timing series. That repeated chr22 disk series also became
I/O-disrupted: full17 after per-pass times were 0.502s and 1.370s, while before
round 2 took about 7.406s/pass. Old round 2 was then interrupted with SIGTERM;
its partial report remains incomplete. Sources are null.chr22.full17.*.

A further memory-resident diagnostic loads the same chr22 keys/vals/col/manifest
bytes into Linux memfd files before running unchanged query binaries. The
16GiB scope includes the resident inputs, with swap disabled. This is an
explicit memory-resident workload, not an NVMe throughput result; it isolates
query CPU/allocation and syscall overhead from disk-service variability.
