#!/bin/bash
# Optional formatting comparison: full annotation text goes to /dev/null.
set -euo pipefail
source "$(dirname "$0")/common.sh"
[[ $# -gt 0 ]] || set -- old before after
versions=("$@")
tag=${RUN_TAG:-wgs}
valid_tag "$tag"
for version in "${versions[@]}"; do require_binary "$version" annotate; done
read -r -a cells <<< "${WORKLOADS:-narrow17 full17 full540 full568}"
for cell in "${cells[@]}"; do
    case "$cell" in narrow17|full17|full540|full568) ;; *) fail "No annotation workload: $cell" ;; esac
done
reserve "annotation.$tag.matrix"
record="$RESULTS/annotation.$tag.matrix"
sha256sum "$QUERY" "$(fields_file 17)" "$(fields_file 540)" "$(fields_file 568)" > "$record/inputs.sha256"
for cell in "${cells[@]}"; do
    for repeat in 1 2 3; do
        order=("${versions[@]}")
        if [[ "$repeat" == 2 ]]; then
            order=()
            for ((i=${#versions[@]}-1; i>=0; i--)); do order+=("${versions[i]}"); done
        fi
        for version in "${order[@]}"; do
            label="annotation.$tag.$cell.$version.r$repeat"
            printf '%s\n' "$label" >> "$record/order.txt"
            limited "$label" timed "$BUILD_ROOT/bin/$version/annotate" \
                -in "$QUERY" -lib "$(library_for "$cell")" -fields "$(fields_for "$cell")" \
                -threads "$THREADS" -batch "$BATCH" -out /dev/null
        done
    done
done
printf '%s\n' "$(date -Is)" > "$record/finished"
