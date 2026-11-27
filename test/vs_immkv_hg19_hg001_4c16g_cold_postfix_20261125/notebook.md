# Cold-cache 4C16G post-fix performance verification

## Purpose

Verify full-output correctness and cold-cache 4C16G performance of `immkv`
across the six scenarios (NVMe SSD 17/540/568 fields and HDD 17/540/568 fields)
after fixing:
1. `iouring_linux.go`: `ioUring.inFlightIO` (`sqeTail - cqHead`) ring-depth
   throttling, immediate CQE reaping after `submitWait(1)`, and `ready[0] = nil`
   slice element cleanup.
2. `colread.go` / `stream.go`: counting non-contiguous `readRange` segments in
   `colShard.estimate` and adding `float64(colRanges) * randSec` to `colSec` in
   `StreamDB.preferColumnar`.

## Controlled protocol

- Pin each measured process group to physical cores 0,1,2,3 and set GOMAXPROCS=4.
- Use one systemd user scope per run with MemoryMax=16G and MemorySwapMax=0.
- Assert the effective cgroup limits and CPU affinity inside every scope.
- Use immkv Threads=4, batch=4096 and automatic read-path selection.
- Run complete-output correctness checks (against historical echtvar references
  via `test/verifyvalues` and exact `data_sha256` equality against `v0.0.5`
  `immkv` streams) before three cold-cache timing repetitions per scenario.
- Complete global sync and write 3 to `/proc/sys/vm/drop_caches` before each
  independently timed invocation via `cache_service.py`.

## Execution order

1. `bash build_tools.sh`
2. `python3 -I preflight.py`
3. Authenticate `cache_service.py` with `pkexec` and run `bash run_when_authenticated.sh`.
