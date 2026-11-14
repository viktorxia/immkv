#!/bin/bash
# Round 1: complete annotation output via FIFOs, no large retained texts.
set -euo pipefail
source "$(dirname "$0")/common.sh"
[[ $# -gt 0 ]] || set -- old before after
[[ $# -ge 2 ]] || fail 'Usage: verify_stream.sh reference candidate [candidate ...]'
reference=$1
shift
versions=("$@")
tag=${RUN_TAG:-wgs}
valid_tag "$tag"
require_binary "$reference" annotate
require_binary "$reference" verifyvalues
for version in "${versions[@]}"; do
    require_binary "$version" annotate
    [[ "$reference" != "$version" ]] || fail 'Reference and candidate must differ'
done
read -r -a cells <<< "${WORKLOADS:-narrow17 full17 full540 full568}"
for cell in "${cells[@]}"; do
    case "$cell" in narrow17|full17|full540|full568) ;; *) fail "No annotation workload: $cell" ;; esac
done
for cell in "${cells[@]}"; do
    for version in "${versions[@]}"; do
        label="verify.$tag.$cell.$reference-$version"
        limited "$label" untimed python3 -I "$EXP_DIR/verify_stream.py" \
            --query "$QUERY" --library "$(library_for "$cell")" --fields "$(fields_for "$cell")" \
            --binaries "$BUILD_ROOT/bin" --reference "$reference" --candidate "$version" \
            --record "$RESULTS/$label" --work "$BUILD_ROOT/$label.fifos"
        sha256sum "$QUERY" "$(fields_for "$cell")" > "$RESULTS/$label/inputs.sha256"
    done
done
