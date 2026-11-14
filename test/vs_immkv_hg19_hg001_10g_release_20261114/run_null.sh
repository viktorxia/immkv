#!/bin/bash
# This phase writes only small logs inside test/ and reads existing libraries.
set -euo pipefail
source "$(dirname "$0")/common.sh"
mkdir -p "$RESULTS"
if [[ ! -x "$BUILD_ROOT/bin/candidate/query-only" ]]; then
    printf 'Run build_versions.sh first\n' >&2
    exit 1
fi
if [[ -e "$RESULTS/null.started" ]]; then
    printf 'Null phase already started; preserve its records and use a new experiment directory\n' >&2
    exit 1
fi
printf '%s\n' "$(date -Is)" > "$RESULTS/null.started"
for cell in narrow17 full17 full540 full568 full568_default; do
    extra=()
    [[ "$cell" == full568_default ]] && extra=(-columns=false)
    for repeat in 1 2 3; do
        versions='baseline candidate'
        [[ "$repeat" == 2 ]] && versions='candidate baseline'
        for version in $versions; do
            label="null.$cell.$version.r$repeat"
            printf '%s %s\n' "$(date -Is)" "$label"
            limited "$label" "$BUILD_ROOT/bin/$version/query-only" \
                -query "$QUERY" -lib "$(library_for "$cell")" -fields "$(fields_for "$cell")" \
                -threads "$THREADS" -batch "$BATCH" "${extra[@]}"
        done
    done
done
printf '%s\n' "$(date -Is)" > "$RESULTS/null.finished"
