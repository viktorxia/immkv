#!/bin/bash
# No large outputs: wait until pure lookup timing stops, then stream the complete
# output through correctness checks. No throughput claim is made for this phase.
set -euo pipefail
source "$(dirname "$0")/common.sh"
if [[ ! -e "$RESULTS/null.finished" ]]; then
    printf 'Wait for run_null.sh to finish before starting verification\n' >&2
    exit 1
fi
bash "$EXP_DIR/verify_stream.sh"
python3 -I "$EXP_DIR/summarize.py"
