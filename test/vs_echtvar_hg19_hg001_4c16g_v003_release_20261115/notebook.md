# Release comparison notebook

Purpose: compare release immkv v0.0.3 (commit `840d36d80351f8b8651d9dc90abf1314bf5db21c`) with echtvar v0.2.4 for the HG001 hg19 query against gnomAD v2.1.1 genomes, on HDD and NVMe, with 17, 540, and 568 fields.

## Completed run ledger

The six-cell run finished. All six fresh immkv output checks pass; fresh
17/540 echtvar outputs pass on both media. All 18 immkv timing runs and all
12 echtvar 17/540 timing runs succeed. The two full-field echtvar timing
attempts OOM; later two repetitions per medium are explicitly skipped.
Independent full-field correctness attempts also OOM on both media.
Measured medians, resource metrics and cache limits are in README.md;
`results/summary-final.md` and `results/table-final.json` retain exact sources.

The initial NVMe full-field OOMs killed the wrapper before final counters were
read and before the unit journal had flushed. Original failed status records
are retained; `status-resolution.json` links authoritative journal/kernel OOM
evidence. `resolve_status.py` now waits briefly for explicit unit outcome
messages, without changing the engine command, resource limits or timing
boundary. The matrix can preserve these resolved results on resumption.

Actual staged execution started with NVMe/17 verification and timing, then
`run_remaining.sh` advanced through the remaining cells. After resolving the
NVMe OOM evidence, `MEDIA=nvme FIELDS=568` completed the remaining successful
immkv repetitions and `MEDIA=hdd` continued all HDD cells. Neither OOM failure
was retried without the cap. No cache reset was performed; the inherited cache
limitations documented in README.md are part of the measured regime.

## Controlled protocol

Each invocation uses one systemd user scope with `MemoryMax=16G` and `MemorySwapMax=0`, affinity CPUs 0–3, and `GOMAXPROCS=4`. These are four distinct physical cores on the current host. The wrapper asserts the effective limits before starting the workload. The 17-field case uses the narrow immkv library; 540 and 568 use the full library. Each medium reads its own existing query, library, and echtvar archives. Source data are read-only.

Both engines produce compressed output. The immkv timing output path is a scratch `output.vcf.gz` symlink to `/dev/null`; its suffix selects the standard annotate BGZF writer. Echtvar always uses its BGZF writer, including when its positional output is `/dev/null`. No large timing outputs are retained. Timing excludes verification.

Echtvar runs four chromosome tasks concurrently on CPUs 0–3. Each task first executes `bcftools view -r N -Oz -o SLICE QUERY`, then `echtvar anno -e chrN.echtvar.zip SLICE OUTPUT`. Slicing is inside the measured invocation. Slices are memfd-backed, referenced through scratch symlinks to the driver's `/proc/PID/fd/FD`, sealed after slicing, and closed after their chromosome finishes. All 22 chromosomes run. Four tasks does not mean four OS threads: each echtvar process may also create htslib reader/writer threads.

## Script order

1. Run `bash build_tools.sh` once to snapshot and build the release annotate tool.
2. Inspect the parent's `results/input-preflight.json` before workloads.
3. Run `bash verify.sh MEDIA FIELDS` for full-output immkv correctness using the parent's separate verifier.
4. Run the fresh echtvar correctness invocation below for the same cell.
5. Run `MEDIA=MEDIA FIELDS=FIELDS RUN_TAG=release bash run_matrix.sh` after correctness evidence is complete.

The build reads the existing module cache and copies the required dependencies into this experiment's scratch cache; network dependency downloads are disabled. The release source comes from an isolated `git archive`. Source archive, source hashes, binary hashes, Go version, CGO setting, and tool versions are recorded in `results/build/`. The separately owned verifier is not built by `build_tools.sh`.

Preflight inspects all 132 ZIP central directories without decompressing archive payloads, plus library manifests and input file facts. Query data identity and full correctness are established separately by the parent's verification tools. Do not interpret central-directory validation as a full archive CRC test.

## Fresh echtvar correctness invocation

From the repository, substitute the cell's medium and field count:

```bash
source /home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/common.sh
media=nvme
fields=17
label="verify-echtvar-$media-$fields-release"
reference="$(media_root "$media")/test/4c16g_gz_v003/${fields}_fields/echtvar.pass2.wk"
if [[ "$fields" == 568 ]]; then
    reference="$HDD_ROOT/test/4c16g_gz_v003/568_fields/echtvar_ref.wk"
fi
limited "$label" "$media" python3 -I "$EXP_DIR/run_cell.py" \
    --engine echtvar --media "$media" --fields "$fields" \
    --record "$RESULTS/$label" --verify-reference "$reference"
```

Each fresh chromosome output is compressed into a FIFO ending in `.vcf.gz`, then decompressed and compared with the historical `anN.vcf.gz`: every data row contributes to SHA-256 and row counts; header lines are excluded. Both counts and hashes must match for every chromosome, and the combined fresh count must be 3,933,714. This is a correctness invocation, not a timing sample. The immkv correctness driver performs the separate full per-row, per-field comparison through the shared verifier, including the rival's documented numeric and miss-marker accounting.

A fresh echtvar 568-field invocation must actually be attempted under the 16G scope independently on both media. Record a constrained OOM as such; do not retry without the cap. An immkv correctness pass against the historical 568-field reference is still required.

## Timing matrix

`run_matrix.sh` runs one cell at a time. Every cell has three paired rounds: immkv then echtvar, echtvar then immkv, immkv then echtvar. It requires `verify-MEDIA-FIELDS-VERIFY_TAG/{status.json,verification.json}` and `verify-echtvar-MEDIA-FIELDS-VERIFY_TAG/{status.json,cell.json}`; `VERIFY_TAG` defaults to `release`. An actual fresh echtvar OOM is accepted as a constrained failure, not a correctness pass.

The first timing OOM is measured and recorded independently per medium and field count. Later rounds for that same echtvar cell receive `skipped_after_oom` records, with no fabricated timing samples. Other failures stop the matrix. Existing successful or OOM records are preserved on resumption; incomplete directories and other failed invocations require investigation and a new tag rather than overwriting records. The matrix holds an experiment-local lock. The parent must also keep correctness runs and other storage workloads from overlapping timing.

Default timeout is 7,200 seconds for HDD and 1,800 for NVMe. Override with `HDD_TIMEOUT_SECONDS`, `NVME_TIMEOUT_SECONDS`, or `RUN_TIMEOUT_SECONDS`. Timeouts and signals produce nonzero results and are retained. These are ordinary paired runs; no cache-drop or cold-cache guarantee is made.

## Evidence and interpretation

`/usr/bin/time -v` executes inside the scope around the entire driver and its children. Its elapsed/user/system times include slicing, compression, and annotation. GNU maximum RSS is a maximum child RSS, not a sum of concurrently resident processes. `memory.peak` is aggregate scope memory, including charged file cache and memfd slices; use it for the controlled memory budget. `memory.events` identifies OOM kills. An outside-scope observer retains the last readable counters if the wrapper is killed; those fallback counters are explicitly marked rather than presented as final measurements.

Each result directory holds constraints, the command, wrapper and workload logs, GNU time, exit status, cgroup counters, status JSON, and per-chromosome timestamps and exit information where applicable. Results are exclusively reserved and are never silently overwritten. Libraries and archives remain outside the repository and are never modified. Scratch under `tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/` holds the frozen source, build caches, binary, and ephemeral FIFO/symlink directories.

Report each measured number with its exact source file in README.md; distinguish controlled facts, measured outcomes, and interpretation. Full correctness elapsed times are not lookup timing samples. Retain old experiment records and do not pool their timing protocol with these compressed-output null runs.
