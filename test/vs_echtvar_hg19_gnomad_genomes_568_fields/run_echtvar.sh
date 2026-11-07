#!/bin/bash
# Time the echtvar side under the controlled protocol (same as the other
# experiments): 4 concurrent per-chromosome processes (echtvar has no thread
# flag), slicing and orchestration inside the measured wall time, plain-text
# conversion outside it (reported separately), two passes.
#
# At 568 fields the archives total ~90 GB -- far beyond the page cache; the
# annotation-rate regime this run measures is the experiment's subject.
#
# Output: $WORK/echtvar.<pass>.txt (plain text) + .time (time -v report).
set -euo pipefail
source "$(dirname "$0")/common.sh"
# CGROUP_MEM=none disables the memory scope (verification-only, untimed runs);
# default 10G is the production-budget protocol.
if [ "${CGROUP_MEM:-10G}" = none ]; then CG_WRAP=""; else CG_WRAP="systemd-run --user --scope --collect -p MemoryMax=${CGROUP_MEM:-10G} -p MemorySwapMax=0"; fi
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
    /usr/bin/time -v -o "$WORK/echtvar.$label.time" ${CG_WRAP} bash -c '
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
