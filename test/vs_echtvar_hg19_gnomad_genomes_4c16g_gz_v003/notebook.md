# vs echtvar: v003 controlled rerun (4 cores / 16G / gzip output)

## Purpose

Repeat the 2026-11-11 controlled matrix on the v003 on-disk format: same
sample, same query, same 4-core / 16G cgroup protocol, same gzip-output
window. The v0.0.2 record of that matrix lives one directory up
(`../vs_echtvar_hg19_gnomad_genomes_4c16g_gz/`); this directory records the
v003 numbers and must never overwrite those files.

v003 changes what is stored, not the query surface: `Get` / `GetBatch` still
return one byte value per key, but the value is a typed record (presence
bitmap + typed payload), so `annotate` renders it back to normalized INFO
text. Floats therefore use `strconv.FormatFloat(..., 'g', -1, 64)` and are
compared numerically, not byte-for-byte, against echtvar.

## Libraries

Rebuilt in place before this run (v0.0.2 libraries deleted on both media):

- full 568-field library: `<media_root>/gnomad/immkv`
- 17-field library: `<media_root>/test/17_fields/lib17`

Both use per-value framing with a trained dictionary (`IMMKV_PACK_KB=-1`), the
same mode as the v0.0.2 baseline libraries, so the matrix compares on-disk
formats rather than framing choices. 540 and 568 share the full library; 540 is
the column subset selected by `fields540.tsv`.

## Script order

1. `run_all.sh` — builds `test/annotate`, then runs every timed cell:
   `{nvme,hdd} × {17,540,568} × immkv` and `{nvme,hdd} × {17,540} × echtvar`,
   two passes each, each cell in its own 16G / swap-off systemd scope pinned to
   CPUs 0-3. Pass 1 writes the full VCF; pass 2 writes the same output again.
   Resumable via `.done` markers under
   `<media_root>/test/4c16g_gz_v003/<field>_fields/`.
2. `verify.sh` — full per-row per-field comparison with `test/verifyvalues`
   against fresh echtvar outputs. 17/540 use the echtvar cells from the
   matrix; 568 gets one unconstrained echtvar reference (2 processes, ~48G
   peak) because 24G/process cannot fit the 16G cap.

## What to measure

- Wall time and pass-2 CPU / peak RSS per cell, from `/usr/bin/time -v`
  reports copied into `results/`.
- Automatic path decisions from `IMMKV_DEBUG_PATH=1` stderr logs in each work
  directory; the v003 model uses record-order extents and measured typed
  decode constants.
- Full-field verification logs under `results/verify/`.

## Environment

- 4 physical cores (CPUs 0-3; SMT siblings 8-11 excluded), `GOMAXPROCS=4`.
- `MemoryMax=16G`, `MemorySwapMax=0`; build phases are not memory-capped.
- `-threads 4 -batch 4096`, default path selection.
- Query: `hg001.decomp.vcf.gz` per field set (pre-existing per-field trees).

## Correction (2026-11-12, after the matrix)

The Libraries paragraph above is wrong for the 17-field baseline. The v0.0.2
17-field library was built by `test/buildnarrow` with
`immkv.Create(..., immkv.Options{})`, i.e. default 64KB packed frames; only
the v0.0.2 full library honoured `IMMKV_PACK_KB=-1`. The v003 17-field
library uses per-value framing, so the 17-field row mixes framing modes.
540/568 share per-value framing on both sides and carry the format
comparison. Evidence: first-window path logs 963 frames / 6,467,620 bytes
(old) vs 4,060 frames / 395,771 bytes (v003); see `README.md`.
