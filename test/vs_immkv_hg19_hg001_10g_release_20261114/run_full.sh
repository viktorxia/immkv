#!/bin/bash
# Start only after explicit authorization to create the external work directory.
# A fresh work directory preserves all prior experiments and outputs.
set -euo pipefail
source "$(dirname "$0")/common.sh"
if [[ -e "$WORK" ]]; then
    printf 'Work directory already exists; do not overwrite: %s\n' "$WORK" >&2
    exit 1
fi
mkdir "$WORK"
mkdir "$WORK/output" "$WORK/builds"
for cell in narrow17 full17 full540 full568; do
    for version in baseline candidate; do
        label="full.$cell.$version"
        output="$WORK/output/$cell.$version.vcf.gz"
        temporary="$WORK/output/$cell.$version.tmp.vcf.gz"
        printf '%s %s\n' "$(date -Is)" "$label"
        limited "$label" "$BUILD_ROOT/bin/$version/annotate" \
            -in "$QUERY" -lib "$(library_for "$cell")" -out "$temporary" \
            -fields "$(fields_for "$cell")" -threads "$THREADS" -batch "$BATCH"
        gzip -t "$temporary"
        mv "$temporary" "$output"
    done
    baseline="$WORK/output/$cell.baseline.vcf.gz"
    candidate="$WORK/output/$cell.candidate.vcf.gz"
    verify_log="$RESULTS/verify.$cell.log"
    # The shared verifier requires plain streams; process substitution avoids
    # keeping another 200+ GiB of uncompressed annotation text.
    "$BUILD_ROOT/bin/verifyvalues" -ev <(gzip -dc "$baseline") \
        -iv <(gzip -dc "$candidate") -fields "$(fields_for "$cell")" -tol 0 \
        > "$verify_log" 2>&1
    cmp <(gzip -dc "$baseline") <(gzip -dc "$candidate")
    printf 'EXACT_DECOMPRESSED_BYTES_OK\n' >> "$verify_log"
    for output in "$baseline" "$candidate"; do
        sha256sum "$output"
        stat --printf='%n %s bytes\n' "$output"
    done >> "$RESULTS/output-facts.txt"
done
printf '%s\n' "$(date -Is)" > "$RESULTS/full.finished"
