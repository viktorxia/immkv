# vs echtvar: 17-field controlled comparison (hg19, gnomAD v2.1.1 genomes)

The light end of the 17 → 540 → 568 dose-response series. Both sides carry
exactly the same information: the curated 17-field subset (AC/AN/AF + AC/AF
for 7 populations). This is echtvar's best-case regime -- 2.9 GB of archives,
trivially cache-resident even inside the 10G budget.

## Measured facts (all numbers from `results/`, 2 passes each, 10G cgroup)

Round 1 -- full plain-text output:

| side | pass1 | pass2 | CPU (user) |
|---|---|---|---|
| echtvar, 4 processes | 17.76 s | 18.04 s | 107-109 s |
| immkv, full library + 17-field whitelist | 2:22.64 | 1:35.36 | 114-118 s |
| immkv, 17-field narrow library | **29.74 s** | 35.17 s | **42 s** |

Round 2 -- output to /dev/null (lookup + format only):

| side | pass1 | pass2 | CPU (user) |
|---|---|---|---|
| echtvar | 18.16 s | 18.15 s | 109 s |
| immkv full library | 49.14 s | 1:26.28 | 115 s |
| immkv narrow library | 40.48 s | 39.43 s | 43 s |

Verification: 66,873,138 field comparisons across all 3,933,714 rows --
zero value mismatches, zero hit disagreements; max float deviation
1.000e-06, inside the derived tolerance 1.1e-6 (`results/verify.log`).

Build account: echtvar 22 archives, 2.9 GB, ~2.5 h (4-way); immkv narrow
library 9.4 GB, 1 h 53 m single-threaded; full library referenced, not
rebuilt (see `../buildfull`).

## Interpretation (clearly labeled as such)

- In its home regime echtvar wins the wall clock (18 s vs 30 s) -- its
  archive sits entirely in page cache and every lookup is a memory read.
- immkv still wins CPU by ~2.6x in the library-to-library pairing (42 s vs
  109 s): one process, one decompression thread pool, no process
  orchestration, no text re-encoding of the query.
- The full-library rows quantify decode amplification: the same hits cost
  49-86 s (null round) when values carry all 568 stored fields but only 17
  are emitted -- scanning 568-field frames to pick 17.

## Reproducing

See `notebook.md` for the script order. Everything runs unattended in ~6
minutes once the libraries exist (`build_echtvar.sh`, `build_immkv.sh`).

## Columnar round (2026-11-11, NVMe)

Rebuilt full library, 17 selected columns = 3.87 GB compressed / 26.7 GB raw.
`annotate -threads 16`, 10G cgroup (`results/col_nvme.txt`):

- default 23.35 s and forced columnar 23.50 s on `/dev/null` (query ~21.3 s
  + format 3.35 s) - the model selects columnar here;
- forced columnar with full output: 20.66 s (query 18.56 s + format
  3.39 s), byte-identical to `results/immkv.round1.pass1.vcf`.

Existing record-order: full library 2:22.64 / 1:35.36 full, 49.14 / 86.28 s
null; narrow library 29.74 / 35.17 s full, 40.48 / 39.43 s null; echtvar
17.76 / 18.04 s. Columnar cuts the full-library 17-field query to ~21 s,
near echtvar but not below it.
