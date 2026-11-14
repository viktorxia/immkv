#!/bin/bash
# Run inside the scope, so time reports include the measured process and its
# workers, not just the systemd-run client.
set -euo pipefail
record=$1
shift
cg=$(python3 -I -c 'from pathlib import Path; print(next(line.split(":",2)[2] for line in Path("/proc/self/cgroup").read_text().splitlines() if line.startswith("0::")))')
max=$(<"/sys/fs/cgroup$cg/memory.max")
swap=$(<"/sys/fs/cgroup$cg/memory.swap.max")
if [[ "$max" != 10737418240 || "$swap" != 0 ]]; then
    printf 'Unexpected cgroup limits: memory=%s swap=%s\n' "$max" "$swap" >&2
    exit 1
fi
printf 'cgroup=%s\nmemory.max=%s\nmemory.swap.max=%s\n' "$cg" "$max" "$swap" > "$record.cgroup"
printf 'command=' >> "$record.cgroup"
printf '%q ' "$@" >> "$record.cgroup"
printf '\n' >> "$record.cgroup"
/usr/bin/time -v -o "$record.time" "$@" > "$record.stdout" 2> "$record.stderr"
for name in memory.peak memory.events; do
    if [[ -r "/sys/fs/cgroup$cg/$name" ]]; then
        printf '%s\n' "$name" >> "$record.cgroup"
        /usr/bin/tee -a "$record.cgroup" < "/sys/fs/cgroup$cg/$name" >/dev/null
    fi
done
