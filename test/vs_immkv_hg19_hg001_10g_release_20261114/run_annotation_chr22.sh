#!/bin/bash
# Annotation includes VCF parsing and typed field formatting but discards its
# text output. This is deliberately separate from query-only lookup timing.
set -euo pipefail
source "$(dirname "$0")/common.sh"
query="$BUILD_ROOT/query-chr22.vcf"
for cell in narrow17 full17 full540 full568; do
    for repeat in 1 2 3; do
        versions='baseline candidate'
        [[ "$repeat" == 2 ]] && versions='candidate baseline'
        for version in $versions; do
            label="annotation.chr22.$cell.$version.r$repeat"
            if [[ -e "$RESULTS/$label.time" ]]; then
                printf 'Timing record already exists: %s\n' "$label" >&2
                exit 1
            fi
            limited "$label" timeout 90s "$BUILD_ROOT/bin/$version/annotate" \
                -in "$query" -lib "$(library_for "$cell")" -fields "$(fields_for "$cell")" \
                -out /dev/null -threads "$THREADS" -batch "$BATCH"
        done
    done
done
printf '%s\n' "$(date -Is)" > "$RESULTS/annotation-chr22.finished"
