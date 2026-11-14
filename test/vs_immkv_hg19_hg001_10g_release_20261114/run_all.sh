#!/bin/bash
# Full suite: call only after authorizing the new external WORK directory.
set -euo pipefail
source "$(dirname "$0")/common.sh"
if [[ ! -x "$BUILD_ROOT/bin/candidate/query-only" ]]; then
    bash "$EXP_DIR/build_versions.sh"
fi
bash "$EXP_DIR/run_full.sh"
if [[ ! -e "$RESULTS/point.started" ]]; then
    bash "$EXP_DIR/run_point.sh"
fi
if [[ ! -e "$RESULTS/null.started" ]]; then
    bash "$EXP_DIR/run_null.sh"
fi
bash "$EXP_DIR/run_build.sh"
python3 -I "$EXP_DIR/summarize.py"
