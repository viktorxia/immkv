# Release regression: immkv versus immkv

This experiment compares commit `9e72562` (before the API/error-channel refactor)
with an exact snapshot of the current working tree. Existing libraries and
historical experiment records are read-only inputs.

## Protocol

- Reference: hg19 / gnomAD v2.1.1 genomes, existing v003 libraries.
- Query: the same decomposed HG001 VCF for both versions and all workloads.
- CPUs: physical cores 0-3, `GOMAXPROCS=4`, four query workers, batch size 4096.
- Every measured process runs in its own systemd scope with `MemoryMax=10G`
  and `MemorySwapMax=0`; the runner checks the actual cgroup files.
- Round 1 writes complete annotation output to a temporary gzip filename and
  renames it only on success. Both versions are verified row by row and field
  by field with `test/verifyvalues`, plus an exact decompressed-byte comparison.
- Round 2 runs `GetBatch` without record decoding, INFO formatting, or gzip
  compression. Results are consumed and discarded; JSON reports lookup time,
  total time, keys, hits, and returned bytes. This is distinct from running the
  annotation formatter with `-out /dev/null`.
- Three measured null rounds alternate execution order. Point lookup uses the
  existing `test/pointlookup` tool, identical samples/seed, warmups, cache budgets
  of 0 and 512 MiB, and three paired repeats after untimed setup runs.
- Build throughput uses a deterministic 200,000-record fixture, preserves each
  built library, and verifies every value after building. This is a synthetic
  build check, not a full gnomAD reconstruction.
- Existing v003 files are reused; no production library is rebuilt or removed.

## Workloads

1. The 17-field narrow library, all 17 fields.
2. The 568-field library, selecting 17 fields.
3. The 568-field library, selecting 540 fields.
4. The 568-field library, all 568 fields.
5. Random point lookups in the full library, 100,000 keys, seed 1.
6. Deterministic sorted build with default packed and per-value framing.

## Script order

1. `bash build_versions.sh`: snapshot sources and build all tools inside the
   repository's ignored `tmp/` directory; this stage is not memory-capped.
2. `bash run_all.sh`: create the new external work directory, run complete-output
   correctness passes, then null rounds, point queries, and build checks.
3. `python3 -I summarize.py`: collect raw timing reports into JSON and Markdown.

The scripts stop on failures, retain their logs, and never overwrite an existing
experiment work directory. The candidate source hashes are recorded to bind the
results to the uncommitted candidate. Large outputs live outside the repository;
only the scripts and result logs belong here. Measurements and limitations are
reported in `README.md` with references to the raw files.

## Current run notes

The external output-directory authorization was not answered, so this run did
not create it. The stored-output gzip and synthetic-build scripts were prepared
but not executed. Point and whole-genome query-only timings ran first, writing
only small reports inside this experiment.

The full HG001 query contains 3,933,714 rows (recorded in the prior controlled
experiment's `results/matrix.log`). A query-only pass consumes 3,933,369
buildable encoded keys and reports 3,875,208 hits; the harness filters unknown
contigs, missing REF and unbuildable ALTs. Sources for this run are
`results/null.*.stdout`; the gap between row and key counts is not interpreted
as missing lookup results. Explicit column selection was tested separately from
the default (`Columns=nil`) path. During full568 round 3 the NVMe I/O regime
changed; the following
whole-genome default-path pass was interrupted. The incomplete pass and all
preceding raw measurements are preserved in `results/`; see
`results/io-regime-note.txt`.

After stopping whole-genome timing, `run_chr22_diagnostic.sh` extracted exactly
chr22 from the same query into the ignored build directory and measured default
versus explicit all-column calls with a 90-second per-process timeout.
`run_annotation_chr22.sh` then measured parsing, lookup and formatting with text
output discarded. These diagnostic/subset numbers are never combined with the
whole-genome timing series.

Correctness was checked for every one of the 49,540 chr22 rows and every selected
field in all four library/field scenarios. `verify_stream_one.sh` and
`verify_chr22.sh` produced complete annotation streams through named pipes,
compared fields with `test/verifyvalues` at zero tolerance and hashed every output
byte. No large output was retained. This phase is not used as performance data:
both producers share a scope and run under verifier backpressure. The first
full568 correctness log retained the filename `verify.full568.log`, but its
49,540-row count identifies it as chr22, not whole-genome verification.

Reproduce the diagnostic/correctness phases after the existing source snapshot:

```bash
bash run_chr22_diagnostic.sh
QUERY="$PWD/../../tmp/vs_immkv_hg19_hg001_10g_release_20261114/query-chr22.vcf" \
  systemd-run --user --scope -p MemoryMax=10G -p MemorySwapMax=0 \
  taskset -c 0-3 env GOMAXPROCS=4 bash verify_stream_one.sh full568
bash verify_chr22.sh
bash run_annotation_chr22.sh
python3 -I summarize.py
```

Use a fresh experiment directory for another run; existing result files are not
permission to overwrite earlier evidence. Whole-genome correctness verification,
stored gzip timing, controlled cold-cache timing, HDD timing and full-size build
performance remain unmeasured in this run.
