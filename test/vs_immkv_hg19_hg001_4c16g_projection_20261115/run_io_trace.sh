#!/bin/bash
# Optional bounded, untimed backend diagnostic; never run beside measurements.
set -euo pipefail
source "$(dirname "$0")/common.sh"
[[ $# -le 1 ]] || fail 'Usage: run_io_trace.sh [old|before|after]'
version=${1:-after}
require_binary "$version" query-only
command -v strace >/dev/null || fail 'strace is required'
tag=${RUN_TAG:-chr22}
valid_tag "$tag"
trace_query=${TRACE_QUERY:-$FROZEN_ROOT/query-chr22.vcf}
[[ -f "$trace_query" ]] || fail "Missing trace query: $trace_query"
label="io-trace.$tag.full17.$version"
limited "$label" untimed strace -f -qq -e trace=io_uring_setup,io_uring_enter \
    -o "$RESULTS/$label/syscalls.log" "$BUILD_ROOT/bin/$version/query-only" \
    -query "$trace_query" -lib "$FULL_LIB" -fields "$(fields_file 17)" \
    -threads "$THREADS" -batch "$BATCH" -repeat 1
sha256sum "$trace_query" "$(fields_file 17)" > "$RESULTS/$label/inputs.sha256"
printf 'Inspect successful io_uring_setup/enter calls in %s; this trace does not prove O_DIRECT flags or other runs.\n' \
    "$RESULTS/$label/syscalls.log"
