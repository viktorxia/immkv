# Notebook: 540-field controlled comparison (vs echtvar)

## Purpose

Compare both systems at echtvar's CEILING: all 540 numeric INFO fields
(368 Integer + 172 Float; echtvar cannot encode the 28 String/Flag fields).
Together with the 17-field experiment this forms the NUMERIC dose-response
pair (same kind of work, 32x more fields) that decides whether echtvar's
slowdown with field count is linear and whether its cache residency buys
anything measurable.

## Controlled conditions

Identical to the 17-field notebook: same query, 4 processes vs 1 process x
4 threads, 10G cgroup per side, plain-text output, two passes.

## Scripts, in execution order

1. `gen_fields.sh` — regenerate `fields540.tsv/.json` (numeric mode).
2. `prep_query.sh` — decompose + INDEX the query.
3. `build_echtvar.sh` — 22 archives (~62 GB total; 8-way; atomic writes;
   needs >=300G free).
4. `run_echtvar.sh` — ROUND 1 (full output, 2 passes).
5. `run_immkv.sh` — ROUND 1; library = pre-built 366G full-load library +
   540-field whitelist (same emitted information as the archives).
6. `verify.sh` — hit-set comparison PLUS full per-row per-field value
   comparison (see the 17-field notebook).
7. `run_null.sh` — ROUND 2 (/dev/null, 2 passes each side).
8. `collect_results.sh`, `cleanup.sh`.

## Statistics to record

Same list as the 17-field notebook (all .time reports, phase splits, verify
counts, conversion milliseconds, build account). Additionally note per-pass
page-cache asymmetry: the immkv library (366G) never fully caches; pass1 vs
pass2 spread IS data, record both.

## Dress-rehearsal notes (2026-11-06/07)

immkv (no cgroup): 2m48s / 5m13s (query 157-297s, format 98-106s -- pass2
slower than pass1: output writes evicted library pages). echtvar runs failed
three times on infrastructure (missing query index; 8 archives truncated by
a killed build); archives rebuilt 11-07 early morning. Official numbers come
from the re-run under the 10G cgroup.
