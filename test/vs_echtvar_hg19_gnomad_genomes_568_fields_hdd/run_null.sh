#!/bin/bash
# Round 2 of the two-round protocol for the all-fields (568) HDD experiment:
# no real output, timed. immkv runs with -out /dev/null (formatting still
# runs, disk writes vanish); two passes.
#
# The echtvar side is intentionally absent. At 568 fields it is OOM-killed
# under the 10G cgroup at any concurrency (see the NVMe experiment
# ../vs_echtvar_hg19_gnomad_genomes_568_fields/README.md); a single attempt is
# kept as results/echtvar_null.pass1.time with the kill evidence.
set -euo pipefail
source "$(dirname "$0")/common.sh"
cd "$REPO"
go build -o "$WORK/annotate" ./test/annotate

for pass in pass1 pass2; do
    echo "== immkv null $pass $(date -Is) =="
    /usr/bin/time -v -o "$WORK/immkv_null.$pass.time" \
        systemd-run --user --scope --collect -p MemoryMax=10G -p MemorySwapMax=0 \
        "$WORK/annotate" -in "$QUERY" -lib "$FULL_LIB" -out /dev/null \
            -fields "$EXP_DIR/fields568.tsv" -threads 4 -batch 4096
    grep -E "Elapsed|User time|System time" "$WORK/immkv_null.$pass.time"
done
echo "== run_null done $(date -Is) =="
