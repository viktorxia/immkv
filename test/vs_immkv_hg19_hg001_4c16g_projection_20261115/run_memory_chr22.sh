#!/bin/bash
# Resident chr22 diagnostic: one scope owns preload, digests, and both matrices.
# This is a memory diagnostic; never feed these results to the disk-run parser.
set -euo pipefail
source "$(dirname "$0")/common.sh"
[[ $# == 0 ]] || fail 'Usage: bash run_memory_chr22.sh (optional RUN_TAG suffix)'
label=memory.chr22
if [[ -n "${RUN_TAG:-}" ]]; then
    valid_tag "$RUN_TAG"
    label+=".$RUN_TAG"
fi
query="$FROZEN_ROOT/query-chr22.vcf"
[[ -f "$query" ]] || fail "Missing existing plain chr22 fixture: $query"
for version in old before after; do require_binary "$version" query-only; done
for version in old after; do require_binary "$version" annotate; done
child_timeout=${RUN_TIMEOUT_SECONDS:-1800}
driver_timeout=${MEMORY_TIMEOUT_SECONDS:-14400}
[[ "$child_timeout" =~ ^[1-9][0-9]*$ && "$driver_timeout" =~ ^[1-9][0-9]*$ ]] ||
    fail 'RUN_TIMEOUT_SECONDS and MEMORY_TIMEOUT_SECONDS must be positive integers'
# limited() creates the only systemd scope. Each child reuses run_limited.sh
# within that scope, so the resident memfd pages count against the same 16 GiB.
# Do not inherit I/O overrides or loader injection into this automatic-path run.
RUN_TIMEOUT_SECONDS="$driver_timeout" limited "$label" untimed \
    env -u IMMKV_NO_DIRECT -u LD_PRELOAD -u LD_AUDIT \
    RUN_TIMEOUT_SECONDS="$child_timeout" \
    python3 -I "$EXP_DIR/memory_chr22.py" \
    --record "$RESULTS/$label" --build-root "$BUILD_ROOT" \
    --query "$query" --full-lib "$FULL_LIB" --narrow-lib "$NARROW_LIB" \
    --fields17 "$(fields_file 17)" --fields540 "$(fields_file 540)" --fields568 "$(fields_file 568)"
printf 'Resident chr22 results: %s\n' "$RESULTS/$label"
