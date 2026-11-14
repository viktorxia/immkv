#!/bin/bash
# Same 10G memory/swap/CPU limits as timed runs; this phase is not timed.
set -euo pipefail
source "$(dirname "$0")/common.sh"
for cell in narrow17 full17 full540 full568; do
    if [[ -e "$RESULTS/verify.$cell.log" ]]; then
        printf 'Verification record exists: %s\n' "$cell" >&2
        exit 1
    fi
    printf '%s verify-stream %s\n' "$(date -Is)" "$cell"
    systemd-run --user --scope --collect -p MemoryMax=10G -p MemorySwapMax=0 \
        taskset -c "$CPUS" env -u IMMKV_FORCE_RECORD -u IMMKV_FORCE_COLUMNAR \
        -u IMMKV_DEBUG_PATH GOMAXPROCS="$THREADS" \
        bash "$EXP_DIR/verify_stream_one.sh" "$cell"
done
printf '%s\n' "$(date -Is)" > "$RESULTS/verify-stream.finished"
