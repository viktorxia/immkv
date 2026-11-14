#!/bin/bash
# Invoked inside the scope; record and assert the actual limits before work.
set -euo pipefail
set -o noclobber
record=$1
mode=$2
shift 2
[[ -d "$record" && $# -gt 0 ]] || exit 2
case "$mode" in timed|untimed) ;; *) exit 2 ;; esac
run_timeout=${RUN_TIMEOUT_SECONDS:-1800}
[[ "$run_timeout" =~ ^[1-9][0-9]*$ ]] || { printf 'Invalid RUN_TIMEOUT_SECONDS\n' >&2; exit 2; }
cg=$(python3 -I -c 'from pathlib import Path; print(next(line.split(":", 2)[2] for line in Path("/proc/self/cgroup").read_text().splitlines() if line.startswith("0::")))')
max=$(<"/sys/fs/cgroup$cg/memory.max")
swap=$(<"/sys/fs/cgroup$cg/memory.swap.max")
affinity=$(python3 -I -c 'import os; print(",".join(map(str, sorted(os.sched_getaffinity(0)))))')
if [[ "$max" != 17179869184 || "$swap" != 0 || "$affinity" != 0,1,2,3 || "${GOMAXPROCS:-}" != 4 ]]; then
    printf 'Unexpected constraints: memory=%s swap=%s affinity=%s GOMAXPROCS=%s\n' \
        "$max" "$swap" "$affinity" "${GOMAXPROCS:-unset}" >&2
    exit 1
fi
{
    printf 'cgroup=%s\nmemory.max=%s\nmemory.swap.max=%s\ncpus=%s\nGOMAXPROCS=%s\n' \
        "$cg" "$max" "$swap" "$affinity" "$GOMAXPROCS"
    printf 'mode=%s\ntimeout_seconds=%s\ncommand=' "$mode" "$run_timeout"
    printf '%q ' "$@"
    printf '\n'
} > "$record/cgroup.txt"
printf '%s\n' "$(date -Is)" > "$record/started"
record_io_context() {
    local phase=$1
    python3 -I - "$record" "$phase" <<'PY'
from pathlib import Path
import json
import sys
import time
record, phase = Path(sys.argv[1]), sys.argv[2]
context = {"timestamp_ns": time.time_ns()}
for name in ("/proc/pressure/io", "/proc/diskstats"):
    path = Path(name)
    if path.exists():
        context[name] = path.read_text()
with (record / ("io-context." + phase + ".json")).open("x") as output:
    json.dump(context, output, indent=2, sort_keys=True)
    output.write("\n")
PY
}
record_io_context before
set +e
if [[ "$mode" == timed ]]; then
    /usr/bin/time -v -o "$record/time.txt" \
        timeout --signal=TERM --kill-after=15s "$run_timeout" "$@" \
        > "$record/stdout" 2> "$record/stderr"
else
    timeout --signal=TERM --kill-after=15s "$run_timeout" "$@" \
        > "$record/stdout" 2> "$record/stderr"
fi
status=$?
set -e
printf '%s\n' "$status" > "$record/exit-status"
record_io_context after
for name in memory.peak memory.events; do
    if [[ -r "/sys/fs/cgroup$cg/$name" ]]; then
        printf '%s\n' "$name" >> "$record/cgroup.txt"
        /usr/bin/tee -a "$record/cgroup.txt" < "/sys/fs/cgroup$cg/$name" >/dev/null
    fi
done
if [[ "$status" == 0 ]]; then
    printf '%s\n' "$(date -Is)" > "$record/finished"
fi
exit "$status"
