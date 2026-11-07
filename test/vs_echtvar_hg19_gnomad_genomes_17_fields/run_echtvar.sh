#!/bin/bash
# Time the echtvar side under the controlled protocol:
#   - echtvar has no thread flag; 4 threads means 4 concurrent processes,
#     one per chromosome, orchestrated by this script.
#   - Slicing and process orchestration are part of the measured wall time --
#     that is the only honest way to attribute process-parallelism costs.
#   - echtvar writes gz only; the plain-text conversion runs OUTSIDE the timed
#     window and is reported as a separate figure (previous-baseline protocol:
#     the forced gz round trip is transparent, not buried in the headline).
#   - Two passes: pass1 (page cache as-is after the builds) and pass2 (warm).
#     echtvar archives are small (3.02 GB total) and effectively always warm.
#
# Output: $WORK/echtvar.<pass>.txt (plain text) + .time (time -v report).
set -euo pipefail
source "$(dirname "$0")/common.sh"
exec > >(tee -a "$WORK/run_echtvar_console.log") 2>&1

one() {
    local c=$1
    "$BCFTOOLS" view -r "$c" -Oz -o "$WK/sl$c.vcf.gz" "$QUERY"
    "$ECHTVAR_BIN" anno -e "$ARCHIVES/chr$c.echtvar.zip" "$WK/sl$c.vcf.gz" \
        "$WK/an$c.vcf.gz" > "$WK/an$c.log" 2>&1
}
export -f one
export BCFTOOLS ECHTVAR_BIN ARCHIVES QUERY WK

run_pass() {
    local label=$1
    WK=$WORK/wk_echtvar.$label
    export WK
    rm -rf "$WK" && mkdir -p "$WK"
    echo "== echtvar pass $label $(date -Is) =="
    # Timed window: per-chromosome slicing + 4-process annotation + gz writes
    # (echtvar has no plain-text mode; see below for the conversion cost).
    /usr/bin/time -v -o "$WORK/echtvar.$label.time" systemd-run --user --scope --collect -p MemoryMax=10G -p MemorySwapMax=0 bash -c '
        seq 1 22 | xargs -P ${ECHTVAR_PAR:-4} -I{} bash -c "one {}"
    '
    grep -E "Elapsed|User time|System time|Percent|Maximum resident" "$WORK/echtvar.$label.time"
    # Plain-text conversion (header + zcat of the 22 gz outputs), OUTSIDE the
    # timed window -- previous-baseline protocol: echtvar's forced gz round
    # trip is real but is reported as a separate figure, not buried in the
    # headline wall time.
    local t0=$(date +%s%3N)
    "$BCFTOOLS" view -h "$QUERY" > "$WK/echtvar.$label.txt"
    for c in $(seq 1 22); do
        zcat "$WK/an$c.vcf.gz" | grep -v "^#" >> "$WK/echtvar.$label.txt"
    done
    echo "text conversion (untimed): $(( $(date +%s%3N) - t0 )) ms"
    echo "output: $(wc -l < "$WK/echtvar.$label.txt") lines"
}

run_pass pass1
run_pass pass2
echo "== run_echtvar done $(date -Is) =="
