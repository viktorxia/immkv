#!/bin/bash
# Identical random samples, cache budgets, workers and warmups for both versions.
set -euo pipefail
source "$(dirname "$0")/common.sh"
mkdir -p "$RESULTS"
if [[ -e "$RESULTS/point.started" ]]; then
    printf 'Point phase already started; preserve these records\n' >&2
    exit 1
fi
printf '%s\n' "$(date -Is)" > "$RESULTS/point.started"
for cache in 0 512; do
    for version in baseline candidate; do
        limited "point.cache$cache.$version.setup" "$BUILD_ROOT/bin/$version/pointlookup" \
            -lib "$FULL_LIB" -query "$QUERY" -n 100000 -threads "$THREADS" \
            -cache "$cache" -warmup 10000 -seed 1
    done
    for repeat in 1 2 3; do
        versions='baseline candidate'
        [[ "$repeat" == 2 ]] && versions='candidate baseline'
        for version in $versions; do
            label="point.cache$cache.$version.r$repeat"
            printf '%s %s\n' "$(date -Is)" "$label"
            limited "$label" "$BUILD_ROOT/bin/$version/pointlookup" \
                -lib "$FULL_LIB" -query "$QUERY" -n 100000 -threads "$THREADS" \
                -cache "$cache" -warmup 10000 -seed 1
        done
    done
done
printf '%s\n' "$(date -Is)" > "$RESULTS/point.finished"
