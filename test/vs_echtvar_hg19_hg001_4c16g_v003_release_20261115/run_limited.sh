#!/bin/bash
# This wrapper runs inside the scope; time covers the complete driver tree.
set -euo pipefail
record=$1
shift
[[ -d "$record" && $# -gt 0 ]] || exit 2
exp=$(cd "$(dirname "$0")" && pwd)
duration=${RUN_TIMEOUT_SECONDS:-1800}
[[ "$duration" =~ ^[1-9][0-9]*$ ]] || exit 2
python3 -I "$exp/env_tools.py" constraints "$record"
printf '%s\n' "$(date -Is)" > "$record/started"
printf '%q ' "$@" > "$record/command.sh"
printf '\n' >> "$record/command.sh"
# The wrapper and GNU time remain available to record child OOM failures.
set +e
/usr/bin/time -v -o "$record/time.txt" \
    timeout --signal=TERM --kill-after=15s "$duration" "$@" \
    > "$record/stdout" 2> "$record/stderr"
rc=$?
set -e
printf '%s\n' "$rc" > "$record/exit-status"
python3 -I "$exp/env_tools.py" finalize "$record" "$rc"
exit "$rc"
