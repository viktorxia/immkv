#!/bin/bash
# Verify the v003 controlled rerun with the shared full field comparator.
#
# v0.0.2 stored raw INFO text; v003 renders typed records back to normalized
# text (floats use strconv 'g'), so byte comparison against the old outputs is
# not a valid v003 check. Instead every field of every row is compared with
# test/verifyvalues against a fresh echtvar output:
#   - 17 / 540: echtvar ran in the timed matrix under the 16G cap;
#   - 568: echtvar needs ~24 GB per process, so one unconstrained reference is
#     generated outside the matrix (verification-only, not a timed cell).
set -euo pipefail
source "$(dirname "$0")/common.sh"
cd "$REPO"

V=$EXP_DIR/results/verify
mkdir -p "$V"
# Large concatenated texts stay on HDD; only the small logs live in the repo.
TMP=$(work_dir hdd 568)/../verify_tmp
mkdir -p "$TMP"
VBIN=$(media_root nvme)/test/4c16g_gz_v003/bin/verifyvalues
mkdir -p "$(dirname "$VBIN")"
go build -o "$VBIN" ./test/verifyvalues

log() { echo "[$(date -Is)] $*"; }

# echtvar_568_ref builds the unconstrained 568 reference once on HDD. Two
# processes keep peak RSS near 48G on this 92G machine.
echtvar_568_ref() {
    local work wk q arch
    work=$(work_dir hdd 568)
    wk=$work/echtvar_ref.wk
    if [ -f "$work/echtvar_ref.done" ]; then
        log "skip echtvar 568 reference"
        return
    fi
    q=$(source_dir hdd 568)/query/hg001.decomp.vcf.gz
    arch=$(source_dir hdd 568)/archives
    rm -rf "$wk"
    mkdir -p "$wk"
    log "start echtvar 568 reference procs=2 (unconstrained)"
    export BCFTOOLS ECHTVAR_BIN
    export QUERY=$q ARCHIVES=$arch WK=$wk
    seq 1 22 | xargs -P 2 -I{} bash -c '
        c=$1
        "$BCFTOOLS" view -r "$c" -Oz -o "$WK/sl$c.vcf.gz" "$QUERY"
        "$ECHTVAR_BIN" anno -e "$ARCHIVES/chr$c.echtvar.zip" \
            "$WK/sl$c.vcf.gz" "$WK/an$c.vcf.gz" >"$WK/an$c.log" 2>&1
    ' _ {}
    touch "$work/echtvar_ref.done"
    log "echtvar 568 reference done"
}

# verify_pair media field echtvar_workdir
verify_pair() {
    local media=$1 field=$2 ewk=$3
    local work src ev iv tsv log
    work=$(work_dir "$media" "$field")
    src=$(source_dir "$media" "$field")
    tsv=$(fields_tsv "$field")
    ev=$TMP/$media.$field.echtvar.txt
    iv=$TMP/$media.$field.immkv.txt
    log=$V/$media.$field.verifyvalues.log
    if [ -f "$log" ] && grep -q 'VERIFY_VALUES OK' "$log"; then
        log "skip verifyvalues $media $field"
        return
    fi
    log "build text $media $field"
    "$BCFTOOLS" view -h "$src/query/hg001.decomp.vcf.gz" >"$ev"
    local c
    for c in $(seq 1 22); do
        zcat "$ewk/an$c.vcf.gz" | grep -v '^#' >>"$ev"
    done
    zcat "$work/immkv.pass2.vcf.gz" >"$iv"
    "$VBIN" -ev "$ev" -iv "$iv" -fields "$tsv" | tee "$log"
    rm -f "$ev" "$iv"
}

verify_pair nvme 17 "$(work_dir nvme 17)/echtvar.pass2.wk"
verify_pair nvme 540 "$(work_dir nvme 540)/echtvar.pass2.wk"
echtvar_568_ref
verify_pair hdd 568 "$(work_dir hdd 568)/echtvar_ref.wk"

log "all verification done"
