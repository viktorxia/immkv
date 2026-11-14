#!/bin/bash
# Paired query-only runs; the second of three rounds reverses version order.
set -euo pipefail
source "$(dirname "$0")/common.sh"
[[ $# -gt 0 ]] || set -- old before after
versions=("$@")
tag=${RUN_TAG:-wgs}
valid_tag "$tag"
query_repeat=${QUERY_REPEAT:-1}
[[ "$query_repeat" =~ ^[1-9][0-9]*$ ]] || fail 'QUERY_REPEAT must be positive'
for version in "${versions[@]}"; do require_binary "$version" query-only; done
read -r -a cells <<< "${WORKLOADS:-narrow17 full17 full540 full568 full568_default}"
for cell in "${cells[@]}"; do library_for "$cell" >/dev/null; done
reserve "null.$tag.matrix"
record="$RESULTS/null.$tag.matrix"
sha256sum "$QUERY" "$(fields_file 17)" "$(fields_file 540)" "$(fields_file 568)" > "$record/inputs.sha256"
for cell in "${cells[@]}"; do
    extra=()
    [[ "$cell" == full568_default ]] && extra=(-columns=false)
    for repeat in 1 2 3; do
        order=("${versions[@]}")
        if [[ "$repeat" == 2 ]]; then
            order=()
            for ((i=${#versions[@]}-1; i>=0; i--)); do order+=("${versions[i]}"); done
        fi
        for version in "${order[@]}"; do
            label="null.$tag.$cell.$version.r$repeat"
            printf '%s\n' "$label" >> "$record/order.txt"
            limited "$label" timed "$BUILD_ROOT/bin/$version/query-only" \
                -query "$QUERY" -lib "$(library_for "$cell")" -fields "$(fields_for "$cell")" \
                -threads "$THREADS" -batch "$BATCH" -repeat "$query_repeat" "${extra[@]}"
        done
    done
done
printf '%s\n' "$(date -Is)" > "$record/finished"
