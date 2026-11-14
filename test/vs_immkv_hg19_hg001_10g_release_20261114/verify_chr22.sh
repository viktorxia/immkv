#!/bin/bash
# Complete-row verification on the fixed chr22 diagnostic subset.
set -euo pipefail
source "$(dirname "$0")/common.sh"
export QUERY="$BUILD_ROOT/query-chr22.vcf"
for cell in narrow17 full17 full540; do
    export VERIFY_TAG="chr22.$cell"
    if [[ -e "$RESULTS/verify.$VERIFY_TAG.log" ]]; then
        printf 'Verification record already exists: %s\n' "$VERIFY_TAG" >&2
        exit 1
    fi
    systemd-run --user --scope --collect -p MemoryMax=10G -p MemorySwapMax=0 \
        taskset -c "$CPUS" env GOMAXPROCS="$THREADS" \
        bash "$EXP_DIR/verify_stream_one.sh" "$cell"
done
# The first full568 stream verification used this same chr22 query; keep its
# original filenames and state its subset explicitly in the generated report.
printf '%s\n' "$(date -Is)" > "$RESULTS/verify-chr22.finished"
