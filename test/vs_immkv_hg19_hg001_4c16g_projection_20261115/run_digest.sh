#!/bin/bash
# Untimed raw-record diagnostics; full annotation remains the correctness gate.
set -euo pipefail
source "$(dirname "$0")/common.sh"
[[ $# -gt 0 ]] || set -- before after
versions=("$@")
tag=${RUN_TAG:-wgs}
valid_tag "$tag"
repeat=${DIAGNOSTIC_REPEAT:-1}
[[ "$repeat" =~ ^[1-9][0-9]*$ ]] || fail 'DIAGNOSTIC_REPEAT must be positive'
for version in "${versions[@]}"; do require_binary "$version" query-only; done
read -r -a cells <<< "${WORKLOADS:-narrow17 full17 full540 full568 full568_default}"
for cell in "${cells[@]}"; do library_for "$cell" >/dev/null; done
reserve "digest.$tag.check"
record="$RESULTS/digest.$tag.check"
sha256sum "$QUERY" "$(fields_file 17)" "$(fields_file 540)" "$(fields_file 568)" > "$record/inputs.sha256"
for cell in "${cells[@]}"; do
    extra=()
    [[ "$cell" == full568_default ]] && extra=(-columns=false)
    for version in "${versions[@]}"; do
        limited "digest.$tag.$cell.$version" untimed "$BUILD_ROOT/bin/$version/query-only" \
            -query "$QUERY" -lib "$(library_for "$cell")" -fields "$(fields_for "$cell")" \
            -threads "$THREADS" -batch "$BATCH" -repeat "$repeat" -digest "${extra[@]}"
    done
done
for version in "${versions[@]}"; do require_binary "$version" query-only; done
python3 -I "$EXP_DIR/check_digest.py" --tag "$tag" --workloads "${cells[@]}" --versions "${versions[@]}" \
    > "$record/comparison.json"
printf '%s\n' "$(date -Is)" > "$record/finished"
