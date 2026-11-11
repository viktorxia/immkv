# vs echtvar: controlled 4-core / 16G / gzip-output rerun (hg19, gnomAD v2.1.1)

## Purpose

The earlier rounds mixed thread counts, memory caps, output forms and read
paths. This experiment holds one sample, four physical cores and 16G for every
timed cell, and makes both tools compress and write gzip inside the timed
window. The engine code is v0.0.2 (`bdfa14e`), unchanged.

## Protocol

- Query: the decomposed HG001 VCF, 3,933,714 data rows, 3,875,208 hits.
- CPU: `taskset -c 0-3`, `GOMAXPROCS=4`; CPUs 0-3 are four physical cores on
  this machine, their SMT siblings (8-11) are excluded.
- Memory: `MemoryMax=16G`, `MemorySwapMax=0` on a per-cell systemd scope.
- Output: both sides write `.vcf.gz`; compression and the write are inside the
  timed window, decompression and verification are outside it.
- immkv: `test/annotate -threads 4 -batch 4096`, default path selection (no
  `IMMKV_FORCE_*`), `IMMKV_DEBUG_PATH=1`.
- echtvar: four concurrent per-chromosome processes (`xargs -P 4`);
  `bcftools view` slicing stays inside the window. htslib read/write threads
  (2+2) are compiled into the binary.
- Two passes per cell; both wall times are reported.
- 568-field echtvar is not attempted: its measured peak is 24.1 GB per process
  (`../vs_echtvar_hg19_gnomad_genomes_568_fields/README.md`), so even one
  process does not fit a 16G cgroup.

## Measured wall times (seconds)

Sources: `results/<media>.<fields>.<tool>.passN.time` (copied from the work
directories by `run_all.sh`).

| media | fields | immkv pass1 | immkv pass2 | echtvar pass1 | echtvar pass2 |
|---|---:|---:|---:|---:|---:|
| nvme | 17 | 26.83 | 27.02 | 24.49 | 24.36 |
| nvme | 540 | 151.42 | 150.54 | 715.96 | 614.57 |
| nvme | 568 | 164.93 | 166.50 | cannot run in 16G | cannot run in 16G |
| hdd | 17 | 78.80 | 45.39 | 28.86 | 25.33 |
| hdd | 540 | 2092.78 | 2040.67 | 663.83 | 657.13 |
| hdd | 568 | 2114.83 | 2113.82 | cannot run in 16G | cannot run in 16G |

Output bytes of pass 2 (`.vcf.gz`):

| media | fields | immkv | echtvar (22 shards) |
|---|---:|---:|---:|
| nvme/hdd | 17 | 460,424,239 / 460,423,562 | 398,948,463 / 398,948,413 |
| nvme/hdd | 540 | 12,637,461,559 / 12,637,461,570 | 11,665,377,157 / 11,665,377,162 |
| nvme/hdd | 568 | 14,516,533,282 / 14,516,533,152 | — |

The few-byte media difference is the query file's `##bcftools_normCommand`
provenance line (each media tree has its own decompressed query); data rows
are identical.

## immkv path selection (pass 2, 982 decisions per cell)

`982 = 961 windows of 4096 rows + 21 chromosome-boundary windows`, one decision
per chromosome segment. Source: `<work>/immkv.pass2.path.log`.

| media | fields | columnar wins | record wins |
|---|---:|---:|---:|
| nvme | 17 | 874 | 108 |
| nvme | 540 | 0 | 982 |
| nvme | 568 | 0 | 982 |
| hdd | 17 | 976 | 6 |
| hdd | 540 | 31 | 951 |
| hdd | 568 | 7 | 975 |

At four threads the cost model picks record order for 540/568 on both media;
the 17-field narrow library is mostly columnar.

## Verification

- Data-row byte equality: all six cells' pass-2 outputs are byte-identical to
  the previously verified reference of the same field set
  (`results/verify/*.cmp.log`, headers excluded).
- Full field comparison (`test/verifyvalues`):
  - HDD 17: 66,873,138 field comparisons, 0 value mismatches, 0 hit
    disagreements, 0 echtvar overflow losses, 30,989,070 float comparisons,
    max |echtvar-immkv| 1.0e-6.
  - HDD 540: 2,124,205,560 field comparisons, 0 value mismatches, 0 hit
    disagreements, 808 echtvar i32-overflow losses (documented limitation),
    666,144,645 float comparisons, max |echtvar-immkv| 1.0e-6.
  - 568: data rows byte-identical to the reference that was verified against
    the unconstrained echtvar output (2,234,349,552 field comparisons in
    `../vs_echtvar_hg19_gnomad_genomes_568_fields/`).
  Sources: `results/verify/*.verifyvalues.log`, `results/verify.log`.

## Controlled facts vs interpretation

Facts:

- NVMe 540: immkv 151.42/150.54 s, echtvar 715.96/614.57 s.
- NVMe 568: immkv 164.93/166.50 s; echtvar cannot complete in 16G.
- NVMe 17: immkv 26.83/27.02 s, echtvar 24.49/24.36 s.
- HDD 540: immkv 2092.78/2040.67 s, echtvar 663.83/657.13 s.
- HDD 568: immkv 2114.83/2113.82 s; echtvar cannot complete in 16G.
- HDD 17: immkv 78.80/45.39 s, echtvar 28.86/25.33 s.

Interpretation:

- NVMe full-field: immkv is 4.08-4.76x faster at 540 and is the only tool
  that completes 568 under 16G.
- HDD full-field: echtvar is 3.1-3.2x faster at 540; 568 has no echtvar
  number because it cannot run.
- 17 fields: echtvar is ~10% faster on NVMe and ~1.8x faster on HDD.
- HDD pass 1 vs pass 2 differs only for the 14G 17-field narrow library
  (78.80 -> 45.39); the 474G full library and the archives exceed the cache,
  so 540/568 are stable across passes.
- The 540 immkv side still uses the 568-field full library plus a 540-field
  whitelist (more stored information than the echtvar 540 archive); no
  540-field narrow library exists.

## Reproduce

```sh
cd <repo>
bash test/vs_echtvar_hg19_gnomad_genomes_4c16g_gz/run_all.sh
bash test/vs_echtvar_hg19_gnomad_genomes_4c16g_gz/verify.sh
```

`MEDIA="nvme"` / `FIELDS="17 540"` restrict the matrix. Libraries, archives
and queries stay in the existing trees under `/mnt/data6t/tian/immkv_test_hdd`
and `$HOME/immkv_test_nvme`; new outputs go to `<root>/test/4c16g_gz/`.
