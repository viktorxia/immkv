#!/bin/bash
# Untimed correctness phase: stream every annotated byte through the shared
# verifier, hash both complete outputs, and retain only logs and FIFO endpoints.
set -euo pipefail
source "$(dirname "$0")/common.sh"
cell=$1
tag=${VERIFY_TAG:-$cell}
pipe_dir="$BUILD_ROOT/verify-$tag"
if [[ -e "$pipe_dir" ]]; then
    printf 'Verification work directory already exists: %s\n' "$pipe_dir" >&2
    exit 1
fi
mkdir "$pipe_dir"
mkfifo "$pipe_dir/raw.baseline" "$pipe_dir/raw.candidate" \
    "$pipe_dir/verify.baseline" "$pipe_dir/verify.candidate"
pids=()
cleanup() {
    for pid in "${pids[@]}"; do
        kill "$pid" 2>/dev/null || true
    done
}
trap cleanup EXIT
for version in baseline candidate; do
    "$BUILD_ROOT/bin/$version/annotate" -in "$QUERY" -lib "$(library_for "$cell")" \
        -fields "$(fields_for "$cell")" -threads "$THREADS" -batch "$BATCH" \
        -out "$pipe_dir/raw.$version" \
        > "$RESULTS/verify-stream.$tag.$version.stats" \
        2> "$RESULTS/verify-stream.$tag.$version.stderr" &
    pids+=("$!")
    (tee "$pipe_dir/verify.$version" < "$pipe_dir/raw.$version" \
        | sha256sum > "$RESULTS/verify-stream.$tag.$version.sha256") &
    pids+=("$!")
done
"$BUILD_ROOT/bin/verifyvalues" -ev "$pipe_dir/verify.baseline" \
    -iv "$pipe_dir/verify.candidate" -fields "$(fields_for "$cell")" -tol 0 \
    > "$RESULTS/verify.$tag.log" 2>&1
for pid in "${pids[@]}"; do
    wait "$pid"
done
pids=()
cmp "$RESULTS/verify-stream.$tag.baseline.sha256" "$RESULTS/verify-stream.$tag.candidate.sha256"
printf 'EXACT_STREAM_SHA256_OK\n' >> "$RESULTS/verify.$tag.log"
for name in raw.baseline raw.candidate verify.baseline verify.candidate; do
    rm "$pipe_dir/$name"
done
rmdir "$pipe_dir"
