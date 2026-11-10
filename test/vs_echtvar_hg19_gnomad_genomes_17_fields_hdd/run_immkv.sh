#!/bin/bash
# Time the immkv side under the controlled protocol: one process, 4 threads,
# plain-text output, same decomposed query as run_echtvar.sh. Two passes per
# library (cold-ish then warm; the 366 GiB full library never fits the page
# cache fully -- that asymmetry is structural and recorded as-is).
#
#   round1 = full-load library + 17-field whitelist (information superset on
#            disk, same emitted fields; the previous baseline's round 1)
#   round2 = 17-field narrow library (library-to-library, information
#            identical to the echtvar archives)
#
# The annotate binary is built once up front so /usr/bin/time measures only
# the annotation run.
#
# Output: $WORK/immkv.<round>.<pass>.vcf + .time
set -euo pipefail
source "$(dirname "$0")/common.sh"
exec > >(tee -a "$WORK/run_immkv_console.log") 2>&1
cd "$REPO"

go build -o "$WORK/annotate" ./test/annotate

run_round() {
    local round=$1 lib=$2
    for pass in pass1 pass2; do
        echo "== immkv $round $pass $(date -Is) =="
        /usr/bin/time -v -o "$WORK/immkv.$round.$pass.time" \
            systemd-run --user --scope --collect -p MemoryMax=10G -p MemorySwapMax=0 \
            "$WORK/annotate" -in "$QUERY" -lib "$lib" \
                -out "$WORK/immkv.$round.$pass.vcf" -fields "$EXP_DIR/fields17.tsv" \
                -threads 4 -batch 4096
        grep -E "Elapsed|User time|System time|Percent|Maximum resident" "$WORK/immkv.$round.$pass.time"
    done
}

run_round round1 "$FULL_LIB"
run_round round2 "$LIB17"
echo "== run_immkv done $(date -Is) =="
