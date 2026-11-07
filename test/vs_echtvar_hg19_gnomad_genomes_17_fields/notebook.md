# Notebook: 17-field controlled comparison (vs echtvar)

## Purpose

Compare immkv and echtvar under IDENTICAL information load: the curated
17-field subset (AC/AN/AF + AC/AF for 7 populations). This is echtvar's
best-case regime (2.9 GB of archives, trivially cache-resident), and the
light end of the 17 -> 540 -> 568 dose-response series.

## Controlled conditions

- Query: GIAB HG001 GRCh37 v4.2.1, decomposed, 3,933,714 rows, identical file
  for both sides.
- Threads: echtvar = 4 concurrent per-chromosome processes (no thread flag);
  immkv = 1 process, 4 threads. Orchestration (slicing, process launches)
  inside echtvar's timed window.
- Memory: both sides run under a 10G cgroup (MemoryMax=10G, MemorySwapMax=0).
- Output: plain text on both sides; echtvar's forced gz round trip is reported
  as a separate untimed figure.
- Two passes each; page-cache state recorded as pass1/pass2.

## Scripts, in execution order

1. `gen_fields.sh` — regenerate `fields17.tsv/.json` from the gnomAD header
   (17-name keep list in `keep17.txt`).
2. `prep_query.sh` — decompose HG001 and INDEX the query (random access for
   per-chromosome slicing).
3. `build_echtvar.sh` — 22 archives from raw gnomAD (4-way norm+encode,
   atomic archive writes, resumable).
4. `build_immkv.sh` — 17-field narrow library (buildnarrow).
5. `run_echtvar.sh` — ROUND 1: timed annotation, 2 passes, plain-text
   conversion outside the timed window.
6. `run_immkv.sh` — ROUND 1: immkv annotate, 4 runs (round1 = full library +
   whitelist, round2 = narrow library), 2 passes each.
7. `verify.sh` — row-by-row hit-set comparison PLUS full per-row
   per-field value comparison (test/verifyvalues: every gnomad_* field of
   every row; floats within one multiplier-2e6 unit, Integer/String exact,
   Flag presence-only).
8. `run_null.sh` — ROUND 2: both sides output to /dev/null, 2 passes each;
   isolates lookup+format from output cost.
9. `collect_results.sh` — copy .time reports and logs into `results/`.
10. `cleanup.sh` — delete large texts/scratch, keep libraries and results.

## Statistics to record (and where)

- `results/*.time` — wall / user / sys / max RSS for every pass, both rounds.
- `results/run_*_console.log` — annotate's phase split (key / query / format)
  and rows/hits counts.
- `results/verify.log` — rows, per-side hits, diverging rows (expect 0).
- echtvar text-conversion milliseconds — printed by run_echtvar.sh (untimed).
- Build account (per-chromosome norm/encode times) — build logs.

## Dress-rehearsal reference numbers (2026-11-06/07, NO cgroup)

echtvar 18.3s / 18.1s; immkv narrow 35.1s / 29.8s; immkv full+whitelist
45.3s / 39.6s; hits 3,875,208 both sides, zero divergence. Official re-run
adds the 10G cgroup; expect differences in cache behavior.
