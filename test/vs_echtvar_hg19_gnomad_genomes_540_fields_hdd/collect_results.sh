#!/bin/bash
# Copy the small result records from WORK into results/ so they live in the
# repository: /usr/bin/time reports, verify logs, run console logs, and the
# build account. The multi-GB VCF texts stay in WORK and never enter the
# repository -- they are reproducible from the run scripts.
set -euo pipefail
source "$(dirname "$0")/common.sh"

mkdir -p "$EXP_DIR/results"
cp -f "$WORK"/*.time "$EXP_DIR/results/" 2>/dev/null || true
for f in verify.log verify_values.log run_echtvar_console.log run_immkv_console.log; do
    [ -f "$WORK/$f" ] && cp -f "$WORK/$f" "$EXP_DIR/results/"
done
ls -la "$EXP_DIR/results/"
