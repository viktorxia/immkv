#!/bin/bash
# Round 2 of the two-round protocol: NO real output, timed.
# Round 1 (run_echtvar.sh + run_immkv.sh + verify.sh) proved output
# correctness with full outputs; this round isolates the lookup+format path
# by sending output to /dev/null:
#   - immkv:   -out /dev/null   (formatting runs; disk writes vanish)
#   - echtvar: anno ... /dev/null (bgzip compression still runs: that cost is
#     intrinsic to echtvar's gz-only output and stays in its numbers)
# echtvar still needs per-chromosome query slices (its 4-process requirement)
# and that orchestration stays inside its timed window, as in round 1.
# Two passes each for stability; .time reports land in $WORK/*_null.*.time.
set -euo pipefail
source "$(dirname "$0")/common.sh"
cd "$REPO"
go build -o "$WORK/annotate" ./test/annotate

one() {
    local c=$1
    "$BCFTOOLS" view -r "$c" -Oz -o "$WK/sl$c.vcf.gz" "$QUERY"
    "$ECHTVAR_BIN" anno -e "$ARCHIVES/chr$c.echtvar.zip" "$WK/sl$c.vcf.gz" /dev/null > "$WK/an$c.log" 2>&1
}
export -f one
export BCFTOOLS ECHTVAR_BIN ARCHIVES QUERY WK

for pass in pass1 pass2; do
    WK=$WORK/wk_null_echtvar.$pass
    export WK
    rm -rf "$WK" && mkdir -p "$WK"
    echo "== echtvar null $pass $(date -Is) =="
    /usr/bin/time -v -o "$WORK/echtvar_null.$pass.time" systemd-run --user --scope --collect -p MemoryMax=10G -p MemorySwapMax=0 bash -c 'seq 1 22 | xargs -P ${ECHTVAR_PAR:-4} -I{} bash -c "one {}"'
    grep -E "Elapsed|User time|System time" "$WORK/echtvar_null.$pass.time"
done
rm -rf "$WORK"/wk_null_echtvar.*

for pass in pass1 pass2; do
    echo "== immkv null $pass $(date -Is) =="
    /usr/bin/time -v -o "$WORK/immkv_null.$pass.time" \
        systemd-run --user --scope --collect -p MemoryMax=10G -p MemorySwapMax=0 \
        "$WORK/annotate" -in "$QUERY" -lib "$FULL_LIB" -out /dev/null \
            -fields "$EXP_DIR/fields540.tsv" -threads 4 -batch 4096
    grep -E "Elapsed|User time|System time" "$WORK/immkv_null.$pass.time"
done
echo "== run_null done $(date -Is) =="
