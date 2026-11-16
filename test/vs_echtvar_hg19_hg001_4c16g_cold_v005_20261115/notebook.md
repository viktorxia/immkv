# Cold-cache 4C16G performance acceptance of v0.0.5-perf-untested.1

## Purpose

Compare the unchanged existing immkv libraries and echtvar v0.2.4 archives
for HG001 against hg19 / gnomAD v2.1.1 genomes on NVMe SSD and HDD with
17, 540 and 568 selected fields. The tested commit is
`d2eac357f3fc126755b43ff8b44a64dc55bc0ec2`.

## Controlled protocol

- Pin each measured process group to physical cores 0,1,2,3 and set GOMAXPROCS=4.
- Use one systemd user scope per run with MemoryMax=16G and MemorySwapMax=0.
- Assert the effective cgroup limits and CPU affinity inside every scope.
- Use immkv Threads=4, batch=4096 and automatic read-path selection.
- Run four chromosome workers for echtvar inside the same scope.
- Use identical query data rows, selected fields and compressed VCF output.
- Run complete-output correctness checks before timing-only repetitions.
- Run three repetitions in the orders immkv/echtvar, echtvar/immkv, immkv/echtvar.
- Complete global sync and write 3 to /proc/sys/vm/drop_caches before each independently timed invocation.
- Keep cache clearing outside the measured window and record its result and /proc/meminfo.
- Abort a run if cache clearing fails and never report that run as cold-cache timing.
- Serialize all measured invocations and do not pre-read libraries or archives after clearing caches.
- Preserve failures, OOMs and incomplete results without treating them as successful timing samples.

The explicit user request for 16 GiB overrides the repository's generic
10 GiB comparison budget. Clearing the OS page cache does not establish
that disk-controller or device caches are empty.

## Preparation and execution order

1. Build annotate and verifyvalues from the exact tagged source snapshot.
2. Record the source identity, compiler environment and executable SHA-256 hashes.
3. Run the annotate and verifyvalues regression tests from that snapshot.
4. Read all six gzip query streams to EOF, checking CRCs, row counts and data-row SHA-256.
5. Check archive central directories, field tables, source metadata and hardware facts.
6. Authenticate the finite cache service with pkexec using the desktop system dialog.
7. Verify full output streams for each scenario and then run its three alternating repetitions.
8. Report raw repeats, medians, CPU, GNU time child RSS, cgroup memory peak and I/O observations.

9. For an unattended authenticated start, run `bash run_when_authenticated.sh`; it waits up to thirty minutes for the cache service and retains a unique orchestration log.

All generated experiment records are new files in this directory. Source
snapshots, binaries and transient FIFOs stay under the matching repository
tmp/ directory. Existing libraries, archives and historical records remain
read-only. Full correctness output is consumed through bounded pipelines.

## Cache-service authentication

`cache_service.py` runs as root for at most twelve hours after explicit
system authentication. Its owner-only Unix socket accepts exactly `drop`
and `stop`; it executes no client-supplied command or pathname. `drop`
performs `os.sync()` followed by a write of `3` to
`/proc/sys/vm/drop_caches`, returning timestamps, exit status and meminfo.
The service exits and removes its socket on `stop` or timeout.
