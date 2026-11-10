#!/bin/bash
# Round 2 of the two-round protocol: NO real output, timed (see the 540/568
# run_null.sh for the rationale). immkv runs BOTH round-1 variants:
#   round1 = full-load library + 17-field whitelist
#   round2 = 17-field narrow library (library-to-library)
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

run_round() {
    local round=$1 lib=$2
    for pass in pass1 pass2; do
        echo "== immkv null $round $pass $(date -Is) =="
        /usr/bin/time -v -o "$WORK/immkv_null.$round.$pass.time" \
            systemd-run --user --scope --collect -p MemoryMax=10G -p MemorySwapMax=0 \
            "$WORK/annotate" -in "$QUERY" -lib "$lib" -out /dev/null \
                -fields "$EXP_DIR/fields17.tsv" -threads 4 -batch 4096
        grep -E "Elapsed|User time|System time" "$WORK/immkv_null.$round.$pass.time"
    done
}
run_round round1 "$FULL_LIB"
run_round round2 "$LIB17"
echo "== run_null done $(date -Is) =="
