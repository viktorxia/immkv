#!/usr/bin/env bash
# Inside the constrained scope; GNU time measures the complete workload tree.
set -euo pipefail
record=$1
shift
[[ -d "$record" && $# -gt 0 ]] || exit 2
exp=$(cd -- "$(dirname -- "$0")" && pwd)
duration=${RUN_TIMEOUT_SECONDS:-3600}
[[ "$duration" =~ ^[1-9][0-9]*$ ]] || exit 2
python3 -I "$exp/env_tools.py" constraints "$record"
python3 -I "$exp/env_tools.py" await-observer "$record"
printf '%s\n' "$(date -Is)" > "$record/started"
printf '%q ' "$@" > "$record/command.sh"
printf '\n' >> "$record/command.sh"
if /usr/bin/time -v -o "$record/time.txt" \
    timeout --verbose --signal=TERM --kill-after=15s "$duration" "$@" \
    > "$record/stdout" 2> "$record/stderr"; then
    rc=0
else
    rc=$?
fi
printf '%s\n' "$rc" > "$record/exit-status"
python3 -I "$exp/env_tools.py" finalize "$record" "$rc"
exit "$rc"
