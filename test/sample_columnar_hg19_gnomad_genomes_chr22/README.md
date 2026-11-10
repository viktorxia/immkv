# Sampling: `base.col` phase-1 on gnomAD v2.1.1 chr22

Read-only sampling for the HDD columnar design in
`docs/eureka/2026-11-09-hdd-columnar-bundle.md` section 10.7. This directory
records the chr22 pilot; it changes neither the library nor the on-disk format.

## Data and run

- Source VCF: `gnomad.genomes.r2.1.1.sites.22.vcf.bgz`, 6,968,754,974 bytes
  (`results/summary.txt`, `vcf`; size from `ls -l`).
- Full-load library used for validation: `.../gnomad/immkv`; `chr22.keys`
  record count 3,567,085 and `chr22.vals` 4,288,824,537 bytes
  (`results/summary.txt`, `library_keys`, `record_order_vals_bytes`).
- Run: `bash run_sample.sh`; 168.1 s (`summary.txt`, `elapsed_seconds`);
  bin 1,000,000 bp, zstd level 9 (`summary.txt`, `bin_width`, `zstd_level`);
  no cgroup, no timing protocol; this is a sampling run, not an experiment
  round.
- Record unit: one logical record per buildable ALT key, same filtering as
  `test/buildfull`. 3,567,085 VCF rows produced 3,567,085 keys, matching the
  library (`summary.txt`, `rows`, `keys`, `library_keys`,
  `library_keys_match`); 0 `REF="."` rows and 0 rows without a buildable ALT
  (`ref_dot_rows`, `rows_without_keys`).

## Measured facts

Byte counts below are decimal (GB = 10^9 bytes); the design's `G` figures are
quoted as written (the 366 G library is 366 GiB).

### Header and INFO order

- 568 `##INFO` fields: 368 Integer, 172 Float, 16 String, 12 Flag; Number:
  424 `A`, 131 `1`, 1 `.`, 12 `0` (`summary.txt`, `header_fields`,
  `header_type_*`, `header_number_*`).
- All 3,567,085 keys have their known fields in strict `##INFO` order:
  0 order-bad keys, 0 inversions (`info_order_ok`, `info_order_bad`,
  `info_order_inversions`); 0 unknown fields, duplicate fields, empty
  segments, Flag-with-value, or non-Flag-without-`=`
  (`info_unknown_occurrences`, `info_duplicate_occurrences`,
  `info_empty_segments`, `info_flag_with_value`, `info_nonflag_bare`);
  0 INFO `.` keys (`info_dot_keys`).
- `results/order_examples.tsv` has no data rows.

### Row length per key

| metric | bytes |
| --- | ---: |
| mean | 14,869.6 |
| p50 | 14,521 |
| p90 | 16,458 |
| p99 | 19,271 |
| max | 64,016 |
| keys <1 KB | 0 |

(`summary.txt`, `row_mean`, `row_p50`, `row_p90`, `row_p99`, `row_max`,
`row_lt_1k`.) All keys are ≥4 KB (`row_bucket_lt_4k` 0); 3,188,076 keys are
<16 KB and 379,009 are <64 KB (`row_bucket_lt_16k`, `row_bucket_lt_64k`).

### Columnar volume (raw-text columns)

| component | bytes |
| --- | ---: |
| compressed value stream | 1,552,948,968 |
| presence bitmap | 253,271,200 |
| length table (uvarint) | 1,964,504,380 |
| length table (u32) | 7,845,874,048 |
| bin directory | 327,168 |
| total, uvarint lengths | 3,771,051,716 |
| total, u32 lengths | 9,652,421,384 |
| record-order chr22 vals | 4,288,824,537 |

(`summary.txt`, `columnar_compressed_value_bytes`,
`columnar_bitmap_bytes`, `columnar_len_varint_bytes`,
`columnar_len_u32_bytes`, `columnar_bin_dir_bytes`,
`columnar_total_varint`, `columnar_total_u32`,
`record_order_vals_bytes`.) The estimate has 20,016 field-bin value frames
over 36 bins (`columnar_value_frames`, `bins`); the value compression ratio
is 0.1084 (`columnar_value_compression_ratio`); the total ratios are
`columnar_ratio_varint` 0.8793 and `columnar_ratio_u32` 2.2506.

### Per-type totals

| Type | fields | present keys | raw bytes | compressed | bitmap | len varint | len u32 | dec elements | dec pass | dec fail |
| --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |
| Integer | 368 | 1,302,089,420 | 2,490,846,550 | 491,694,800 | 164,091,200 | 1,302,089,420 | 5,208,357,680 | 1,302,089,420 | 1,302,089,420 | 0 |
| Float | 172 | 608,833,055 | 6,705,996,939 | 651,892,248 | 76,694,800 | 608,833,055 | 2,435,332,220 | 608,833,055 | 0 | 608,833,055 |
| String | 16 | 50,546,037 | 5,132,340,677 | 409,361,920 | 7,134,400 | 53,581,905 | 202,184,148 | 0 | 0 | 0 |
| Flag | 12 | 3,406,364 | 0 | 0 | 5,350,800 | 0 | 0 | 0 | 0 | 0 |

(Summed over `results/fields.tsv` columns `Type`, `present_keys`,
`raw_value_bytes`, `compressed_bytes`, `bitmap_bytes`, `len_varint_bytes`,
`len_u32_bytes`, `dec_elements`, `dec_pass`, `dec_fail`.)

### Largest compressed fields

| field | Type | raw bytes | compressed |
| --- | --- | ---: | ---: |
| vep | String | 4,005,300,084 | 149,506,437 |
| gq_hist_all_bin_freq | String | 264,921,326 | 106,883,473 |
| dp_hist_all_bin_freq | String | 221,928,910 | 72,220,170 |
| ab_hist_alt_bin_freq | String | 145,034,454 | 22,901,550 |
| dp_hist_alt_bin_freq | String | 144,676,172 | 20,563,060 |
| gq_hist_alt_bin_freq | String | 144,172,765 | 15,606,564 |
| rf_tp_probability | Float | 39,237,935 | 11,907,853 |
| age_hist_het_bin_freq | String | 71,109,912 | 11,501,110 |
| AF | Float | 39,222,854 | 10,311,715 |
| non_topmed_AF | Float | 39,222,205 | 9,988,642 |

(`results/fields.tsv`, sorted by `compressed_bytes`.)

### Decimal round trip

- 540 numeric fields; 368 Integer fields all pass, 172 Float fields all fail
  (`summary.txt`, `numeric_fields`, `numeric_all_pass_fields`,
  `numeric_fail_fields`).
- 1,910,922,475 numeric elements; 1,302,089,420 pass, 608,833,055 fail;
  pass rate 0.681393 (`numeric_elements`, `numeric_pass`, `numeric_fail`,
  `numeric_pass_rate`).
- The only failure reason is `exponent` (`dec_reason_exponent` 608,833,055);
  examples include `AF=4.44444e-01`, `BaseQRankSum=7.36000e-01`,
  `MQRankSum=-1.02600e+00` (`results/dec_failures.tsv`).

## Interpretation

- **Record unit and INFO order are validated.** The 3,567,085 keys match the
  library exactly, and the source field order matches the `##INFO` order for
  every key; the columnar schema order would reproduce the record-order field
  order on this dataset. The manifest's INFO-order check would be a no-op
  here.
- **The 0-inversion result is a property of gnomAD, not a design guarantee.**
  The 2026-11-10 decision makes ingest rebuild every value in `##INFO` order
  (`BuildFromVCFs` and `test/buildfull`), so the columnar schema order matches
  the record-order value by construction; this sample's 0 is no longer a design
  dependency.
- **The stated decimal grammar does not cover gnomAD Float formatting.** The
  Float columns are printed in scientific notation (`4.44444e-01`), which the
  "sign + digits + point position" grammar rejects. All 172 Float columns
  (42% of the compressed value bytes; 6.71 GB raw, 0.652 GB compressed) would
  fall back to raw text; only the 368 Integer columns (2.49 GB raw, 0.492 GB
  compressed) can use the binary form as stated. An exponent-aware grammar
  (mantissa + decimal position + exponent) is required to bring Float into
  the binary path.
- **The design's 60–85 G columnar estimate is not supported by raw-text
  columns.** On chr22 the raw-text columnar total is 3.77 GB with uvarint
  lengths (0.879× the 4.29 GB record-order value file); with u32 lengths it is
  9.65 GB (2.25×). The compressed value stream alone is 1.55 GB (0.362×), so
  field-name removal does shrink the values, but the length table (1.96 GB,
  52% of the total) and the presence bitmap (0.253 GB, 6.7%) close the gap.
  The length table is the largest single component; the design text says
  "column = presence bitmap + value sequence" but does not specify lengths.
  Reaching the design's 60–85 G full-genome target needs a compact or
  compressed length table, fixed-width columns, or both; the binary decimal
  encoding alone does not remove this overhead.
- **The 1 MB bin target is larger than the design's 20–30 MB.** The 36 bins
  average 99,085.7 keys and 43.1 MB of compressed values (1.55 GB / 36), plus
  54.6 MB of uvarint lengths and 7.0 MB of bitmap, for ~104.8 MB per bin; the
  design's 20–30 MB/bin target at 1 Mb is not met on gnomAD chr22. This is a
  bin-granularity calibration issue, not a format blocker.
- **The `<1KB` exception is not exercised.** Every chr22 key is ≥1 KB (mean
  14,869.6, p50 14,521); the gnomAD sample says nothing about the ClinVar-like
  case. The sampler is generic; rerunning it on a ClinVar VCF would test the
  exception directly. The design's own premise (build decides fields, query
  decides subset, layout does not preset query patterns) still argues against
  a build-time average-row threshold, but this sample does not settle the
  ClinVar volume.
- **String fields are not uniform.** `vep` is 4.0 GB raw but 149.5 MB
  compressed (3.7%); the top six String fields (`vep` plus five histogram
  fields) total 0.388 GB compressed, and the whole String class is 0.409 GB.
  The Float class is 0.652 GB compressed and the Integer class 0.492 GB.
  Field selection is where the columnar path wins; the full-field scan still
  reads all of it.

## Reproducing

```bash
GNOMAD_SITES=/path/to/gnomad/v2.1.1 \
FULL_LIB=/path/to/full-load/library \
bash run_sample.sh
```

The script runs `go run ./test/samplecolumns` from the repository root and
writes `results/`. The source VCF and the library are read-only inputs. The
sampler is generic: point it at another chromosome or another VCF; `-bin`,
`-level`, and `-limit` control the estimate. The committed `results/` are
from the 2026-11-10 chr22 run (168.1 s).

## Limitations

- chr22 only; no full-genome extrapolation beyond the recorded ratio.
- Raw-text columns; the volume does not include a decimal-encoded or
  exponent-aware binary value stream, nor a compressed length table.
- The length table is counted uncompressed; the uvarint and u32 variants are
  bounds, not the final format.
- zstd level 9, no frame CRC, 1 Mb bins, 16 B/field/bin directory estimate.
- The gnomAD chr22 rows are single-ALT; the sampler supports multi-ALT rows
  (the value is duplicated per key), but this sample had none.
- No HDD/NVMe timing, no `base.col` implementation, no point-lookup path.
