# Notebook: all-fields (568) comparison (vs echtvar)

## Purpose

Compare both systems under the FULL annotation load: all 568 INFO fields
including 16 String and 12 Flag. Versus the 540 experiment the added load is
the STRING fields -- a qualitative change, not a 5% increment: echtvar routes
them through categorical dictionaries that degenerate on real gnomAD columns
(pilot, chr22: vep 3.05M unique values, dp_hist_all_bin_freq 3.26M of 3.57M
variants), while immkv stores and returns raw bytes. This experiment measures
what the string load costs each side and whether values round-trip exactly.

## Controlled conditions

Identical to the 17-field notebook. Archive set ~68 GB (does not fit the 10G
budget); immkv library = 366G full-load + full 568-field list.

## Scripts, in execution order

1. `gen_fields.sh` — regenerate `fields568.tsv/.json` (all mode; census of
   Integer/Float/String/Flag counts is printed -- 368/172/16/12).
2. `probe_encode_chr22.sh` — OPTIONAL single-chromosome pilot (encode + full
   cardinality census + resources; leaves `pilot_encode_chr22.log`).
3. `prep_query.sh` — decompose + INDEX the query.
4. `build_echtvar.sh` — 22 archives (4-way, NOT 8: encode RSS scales with
   variant count, chr22 pilot peaked at 5.9G for the smallest chromosome).
5. `run_echtvar.sh` — ROUND 1 (full output, 2 passes; conversion untimed).
6. `run_immkv.sh` — ROUND 1 (2 passes; ~65 GB of text per pass).
7. `verify.sh` — hit-set comparison.
8. `verify.sh` includes the FULL value comparison (every row, every field).
   `verify_values.sh` is superseded (no-op kept for pipeline compatibility).
   Comparison rules are grounded in echtvar v0.2.4 SOURCE, fetched from the
   tagged release (keep the command in the README):
     curl -H "Accept: application/vnd.github.raw+json" \
       "https://api.github.com/repos/brentp/echtvar/contents/<path>?ref=v0.2.4"
   - Float quantization: `(val * multiplier) as i32` in
     src/commands/encoder_cmd.rs -- Rust `as i32` truncates toward zero;
     decode divides back. Error < one unit (5e-7 at multiplier 2e6), plus
     f32 arithmetic slack: tolerance 6e-7, DERIVED not guessed. The verifier
     also reports the max observed |echtvar-immkv| float difference; if it
     sits just under 5e-7 the truncation model is empirically confirmed
     across millions of values.
   - String fields are Categorical from encode time (same source file);
     miss sentinels: u32::MAX stored, rendered as missing_value (-1) for
     numbers and missing_string ("MISSING") for categoricals.
   - A real value colliding with a miss marker would surface as a hit
     disagreement in the full comparison -- adjudicated empirically.
9. `run_null.sh` — ROUND 2 (/dev/null).
10. `collect_results.sh`, `cleanup.sh`.

## Statistics to record

Same as the 17-field notebook, plus: value-check summary from
verify_values.sh (pairs, checks, mismatches) and the categorical cardinality
census from the build/pilot logs.

## Dress-rehearsal reference numbers (2026-11-06/07, NO cgroup)

echtvar 9m16s / 9m56s wall, CPU 3370s; immkv 4m44s / 5m10s wall, CPU 246s
(query 272-299s, format ~99s hidden); hits 3,875,208 both sides, zero
divergence; value sampling 29/30 clean, 1 float off by 3e-7 (within the
truncation unit). Caveat for the paper: echtvar's wall includes its forced
65GB bgzip output; ROUND 2 removes output from both sides and is the fair
lookup comparison -- the hypothesis that echtvar's compact 68G archive beats
immkv's 366G library on cold lookup speed is specifically tested there.
