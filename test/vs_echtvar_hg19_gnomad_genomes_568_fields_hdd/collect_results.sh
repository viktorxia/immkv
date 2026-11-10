#!/bin/bash
# Copy the small result records from WORK into results/ so they live in the
# repository: /usr/bin/time reports, run console logs, the null-round log, the
# NVMe reference timing, and the chain log that carries the CMP_OK line.
# Multi-GB VCF texts stay in WORK and never enter the repository.
set -euo pipefail
source "$(dirname "$0")/common.sh"

mkdir -p "$EXP_DIR/results"
cp -f "$WORK"/*.time "$EXP_DIR/results/" 2>/dev/null || true
for f in run_immkv_console.log null_rerun.log; do
    [ -f "$WORK/$f" ] && cp -f "$WORK/$f" "$EXP_DIR/results/"
done
# Chain log of the recorded run (prep -> immkv -> null -> NVMe compare).
[ -f "$HDD_ROOT/test/568_fields_watcher.log" ] && \
    cp -f "$HDD_ROOT/test/568_fields_watcher.log" "$EXP_DIR/results/chain_hdd.log"
# NVMe reference timing produced by verify_nvme_reference.sh.
[ -f "$NVME_ROOT/test/568_fields_ref/nvme.time" ] && \
    cp -f "$NVME_ROOT/test/568_fields_ref/nvme.time" "$EXP_DIR/results/nvme_reference.time"
ls -la "$EXP_DIR/results/"
