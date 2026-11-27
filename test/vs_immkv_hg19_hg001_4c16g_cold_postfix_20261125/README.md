# 4C16G cold-cache post-fix verification

Status: complete. All six full-output correctness verifications and all 18
cold-cache timing runs (three rounds per scenario across NVMe SSD 17/540/568
and HDD 17/540/568) succeeded with zero value mismatches, zero hit
disagreements, and exact decompressed data-row SHA-256 equality with the
`v0.0.5` (`d2eac357f3fc126755b43ff8b44a64dc55bc0ec2`) `immkv` baseline.

## Tested changes

Base commit: `904948f0db6f59d94e8a69d97d300fba82aae24b` (tag `v0.0.5`, engine
commit `d2eac357f3fc126755b43ff8b44a64dc55bc0ec2`).
Patch: `results/preparation/working-tree.patch`; frozen source manifest:
`results/preparation/sources.sha256` (`9775500f8502e62d9c0d7fa02453de8e9fce696544a6a9f524b1ab4752d97944`).

1. `iouring_linux.go`: added `ioUring.inFlightIO` (`sqeTail - cqHead`) to bound
   unharvested kernel in-flight read requests to the ring depth in `driveDirect`,
   reaped CQEs immediately after `submitWait(1)`, and cleared `ready[0] = nil`
   when dequeuing completed chunks (`TestDriveDirectInFlightRespectsRingDepth` in
   `iouring_linux_test.go`).
2. `colread.go` / `stream.go`: counted non-contiguous `readRange` segments in
   `colShard.estimate` using the same adjacency rule as `colShard.loadBin` and
   added `float64(colRanges) * randSec` to `colSec` in `StreamDB.preferColumnar`
   (`TestPreferColumnarCountsColumnRanges` in `col_test.go`).

## Controlled facts and timing results

Protocol: CPUs 0-3 (four distinct physical cores), `GOMAXPROCS=4`,
`MemoryMax=16G`, `MemorySwapMax=0`, `Threads=4`, `batch=4096`, automatic
read-path selection, and root `os.sync()` + write `3` to
`/proc/sys/vm/drop_caches` before every independently timed invocation.
Sources: `common.sh`, `run_limited.sh`, `results/timing-*/constraints.json`,
`results/timing-*/cache-clear.json`.

Medians of three cold-cache compressed-output timing rounds compared with
`v0.0.5` (`../vs_echtvar_hg19_hg001_4c16g_cold_v005_20261115/results/final-summary.json`):

| Medium | Fields | Post-fix wall (s) | v0.0.5 wall (s) | Post-fix CPU (s) | v0.0.5 CPU (s) | Post-fix max child RSS (KiB) | v0.0.5 max child RSS (KiB) | Post-fix cgroup peak (GiB) | v0.0.5 cgroup peak (GiB) | Source records |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---:|---|
| NVMe SSD | 17 | 11.57 | 13.20 | 35.14 | 33.15 | 268,964 | 288,500 | 4.44 | 4.47 | `results/timing-nvme-17-immkv-r{1,2,3}-cold-postfix/time.txt` |
| NVMe SSD | 540 | 171.01 | 158.86 | 660.11 | 616.02 | 2,340,680 | 2,331,400 | 6.47 | 6.47 | `results/timing-nvme-540-immkv-r{1,2,3}-cold-postfix/time.txt` |
| NVMe SSD | 568 | 192.98 | 172.47 | 731.70 | 670.46 | 3,067,360 | 3,317,232 | 7.17 | 7.41 | `results/timing-nvme-568-immkv-r{1,2,3}-cold-postfix/time.txt` |
| HDD | 17 | 99.77 | 96.46 | 215.96 | 204.67 | 255,796 | 250,988 | 6.98 | 6.99 | `results/timing-hdd-17-immkv-r{1,2,3}-cold-postfix/time.txt` |
| HDD | 540 | 1468.42 | 1468.32 | 725.88 | 674.01 | 682,192 | 675,300 | 4.98 | 4.98 | `results/timing-hdd-540-immkv-r{1,2,3}-cold-postfix/time.txt` |
| HDD | 568 | 1469.80 | 1469.52 | 776.11 | 719.20 | 863,952 | 859,164 | 5.12 | 5.15 | `results/timing-hdd-568-immkv-r{1,2,3}-cold-postfix/time.txt` |

Summary files: `results/final-summary.md`, `results/final-summary.json`.

## Correctness

All six scenarios processed 3,933,714 rows and observed 3,875,208 hits, with
zero value mismatches, zero hit disagreements, and exact decompressed data-row
SHA-256 and byte-length equality with the `v0.0.5` `immkv` outputs:

| Medium | Fields | Field comparisons | Value mismatches | Hit disagreements | `immkv` data-row SHA-256 (`v005_digest_equal=true`) | Source |
|---|---:|---:|---:|---:|---|---|
| NVMe SSD | 17 | 66,873,138 | 0 | 0 | `8e1c55459edc13e1b24236ef9d65b54f3215dc11dd55db97efa5979f5247af63` | `results/verify-nvme-17-cold-postfix/verification.json` |
| NVMe SSD | 540 | 2,124,205,560 | 0 | 0 | `9e54e963a7a9db6063232afa59a58c48f37d999289116a43e16f97bfccd59dcd` | `results/verify-nvme-540-cold-postfix/verification.json` |
| NVMe SSD | 568 | 2,234,349,552 | 0 | 0 | `c1ca1dd2d338f254df4db1e1e9c4de8c9cad927925a4cfc51111e9a0d8c8ee43` | `results/verify-nvme-568-cold-postfix/verification.json` |
| HDD | 17 | 66,873,138 | 0 | 0 | `8e1c55459edc13e1b24236ef9d65b54f3215dc11dd55db97efa5979f5247af63` | `results/verify-hdd-17-cold-postfix/verification.json` |
| HDD | 540 | 2,124,205,560 | 0 | 0 | `9e54e963a7a9db6063232afa59a58c48f37d999289116a43e16f97bfccd59dcd` | `results/verify-hdd-540-cold-postfix/verification.json` |
| HDD | 568 | 2,234,349,552 | 0 | 0 | `c1ca1dd2d338f254df4db1e1e9c4de8c9cad927925a4cfc51111e9a0d8c8ee43` | `results/verify-hdd-568-cold-postfix/verification.json` |

Total field comparisons across the six verifications: 8,850,856,500.

## Interpretation

- Full-output correctness is byte-for-byte identical to `v0.0.5` across all six
  scenarios (both on the NVMe record-order `io_uring` path and on the HDD
  columnar `base.col` path).
- On HDD, where `StreamDB.preferColumnar` selects the columnar path, adding
  `float64(colRanges) * randSec` preserved the columnar choice across all three
  field sets: HDD 540 (`1468.42s` vs `1468.32s`) and HDD 568 (`1469.80s` vs
  `1469.52s`) match `v0.0.5` within `0.3s` (`< 0.02%`).
- On NVMe, enforcing the 512-entry `io_uring` unharvested in-flight cap and
  `ready` slice pointer cleanup improved NVMe 17 (`11.57s` vs `13.20s`, `-12.3%`)
  and reduced NVMe 568 peak child RSS (`3,067,360 KiB` vs `3,317,232 KiB`,
  `-244 MiB`) and cgroup memory peak (`7.17 GiB` vs `7.41 GiB`, `-0.24 GiB`).
  NVMe 540 (`171.01s` vs `158.86s`, still `3.62×` faster than `echtvar`'s
  `619.24s`) and NVMe 568 (`192.98s` vs `172.47s`, while `echtvar` OOMs) show a
  small wall-time increase (`+7.6%` to `+11.9%`) accompanied by a proportional
  user-CPU increase (`660.11s` vs `616.02s`, `731.70s` vs `670.46s`) from
  stricter 512-request ring batching and immediate `reap` cycles per `submitWait(1)`.
