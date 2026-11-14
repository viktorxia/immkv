#!/bin/bash
set -euo pipefail
source "$(dirname "$0")/common.sh"
if [[ ! -d "$WORK/builds" ]]; then
    printf 'The authorized work directory must be created by run_full.sh first\n' >&2
    exit 1
fi
for framing in packed per_value; do
    pack=0
    [[ "$framing" == per_value ]] && pack=-1
    for repeat in 1 2 3; do
        versions='baseline candidate'
        [[ "$repeat" == 2 ]] && versions='candidate baseline'
        for version in $versions; do
            label="build.$framing.$version.r$repeat"
            base="$WORK/builds/$framing.$version.r$repeat/lib"
            limited "$label" "$BUILD_ROOT/bin/$version/build-check" \
                -mode build -base "$base" -n 200000 -pack "$pack"
            limited "verify.$label" "$BUILD_ROOT/bin/$version/build-check" \
                -mode verify -base "$base" -n 200000 -pack "$pack"
        done
    done
done
printf '%s\n' "$(date -Is)" > "$RESULTS/build.finished"
