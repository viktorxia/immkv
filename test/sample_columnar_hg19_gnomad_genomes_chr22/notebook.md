# Notebook: `base.col` phase-1 sampling on gnomAD v2.1.1 chr22

## Purpose

Read-only sampling for the HDD columnar design in
`docs/eureka/2026-11-09-hdd-columnar-bundle.md` section 10.7. The sampler
tests three assumptions before any on-disk format is written:

1. INFO field order in the source records is consistent with the `##INFO`
   declaration order (the columnar path emits fields in schema order).
2. Integer/Float values can be round-tripped losslessly through the proposed
   "sign + digits + decimal point position" encoding, per column.
3. The per-field, per-bin column volume (zstd-compressed values + presence
   bitmap + length table) compared with the existing record-order value file.

The sampler also measures the average INFO row length and the fraction below
1 KB, so the `<1KB` build-time exception can be discussed with data instead of
a guess.

## Data

- Source: `gnomad.genomes.r2.1.1.sites.22.vcf.bgz` (6,968,754,974 bytes) from
  the gnomAD v2.1.1 GRCh37 sites directory.
- Validation: the 366 GiB full-load immkv library's `chr22.keys` header
  (record count) and `chr22.vals` size (4,288,824,537 bytes).
- Sample unit: one gnomAD VCF row is expanded into one logical record per
  buildable ALT key, exactly as `test/buildfull` does (skip `REF="."`, skip
  empty/`<`/`*`/`.`/breakend ALTs).

## Method

- Tool: `test/samplecolumns` (Go, read-only).
- Bin width: 1,000,000 bp; zstd level 9; per-field per-bin frames.
- Presence bitmap: `ceil(bin records / 8)` bytes per field per bin.
- Length table: both a compact uvarint per present value and a 4-byte u32
  per present value are reported; the actual `base.col` format is not fixed.
- Value stream: present values concatenated in record order and compressed
  per field per bin with zstd; Flag fields have no value stream.
- Decimal round trip: optional sign, digits, optional decimal point with
  digits on either side; leading/trailing zeros, `+`, and `-0` are preserved;
  exponents, `NaN`, `Inf`, whitespace, and empty elements fail.
- INFO order: the record's known-field sequence is compared with the
  `##INFO` order; any non-increasing adjacent pair is an inversion. A record
  with no inversions is reproduced in schema order by the columnar path.
- Row length is the INFO value length per logical record; a multi-ALT row
  duplicates the INFO value per key, matching the library's record count.

## Script order

1. `run_sample.sh` — runs the sampler on chr22 and writes `results/`.

## What is measured

- `results/summary.txt`: record/key counts, library validation, INFO order,
  row-length distribution, columnar volume totals, numeric round-trip totals.
- `results/fields.tsv`: per-field present count, raw/compressed volume,
  bitmap, length-table variants, decimal pass/fail and reasons.
- `results/dec_failures.tsv`: examples of decimal failures.
- `results/order_examples.tsv`: records with INFO order inversions.
- `results/run.log`: progress log and the exact command line.

## Not measured

- Full genome or other chromosomes; only chr22.
- On-disk `base.col` bytes (the format is not implemented).
- Wall-clock and I/O on HDD or NVMe (this is not a timed run).
- Point-lookup behavior; the columnar path is for sequential scans.
- ClinVar-like <1 KB libraries; the sampler is generic and can be rerun on
  another VCF, but this record is gnomAD chr22.
- Decimal-encoded compressed volume; the volume here is raw-text columns.

## Assumptions

- zstd level 9 and no frame CRC; the build may choose another level.
- Per-field per-bin frames; the bin directory cost is approximated at 16
  bytes per field per bin.
- The length table is either uvarint or u32; the format is not fixed.
- Multi-ALT rows duplicate the INFO value per key, matching the full-load
  library's record count.
