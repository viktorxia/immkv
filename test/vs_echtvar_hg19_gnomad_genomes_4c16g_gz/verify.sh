#!/bin/bash
# Verify the controlled rerun.
#
# The engine code is unchanged from v0.0.2, so every pass-2 output must be
# byte-identical to the previously verified reference of the same field set.
# We first prove that by streaming the new gzip output through cmp (no 66 GB
# text files), then run the shared verifier test/verifyvalues on one media per
# field set (HDD 17 / HDD 540) to satisfy the per-row per-field protocol.
#
# 568 has no echtvar side (16G cannot hold it): its reference was verified
# against the unconstrained echtvar output in the earlier 568 experiment.
set -euo pipefail
source "$(dirname "$0")/common.sh"
cd "$REPO"

V=$EXP_DIR/results/verify
mkdir -p "$V"
VBIN=$(media_root nvme)/test/4c16g_gz/bin/verifyvalues
mkdir -p "$(dirname "$VBIN")"
go build -o "$VBIN" ./test/verifyvalues

log() { echo "[$(date -Is)] $*"; }

cmp_stream() { # media field reference
    local media=$1 field=$2 ref=$3
    local work log
    work=$(work_dir "$media" "$field")
    log=$V/$media.$field.cmp.log
    if [ -f "$log" ] && grep -q '^CMP OK' "$log"; then
        log "skip cmp $media $field"
        return
    fi
    log "cmp $media $field"
    # Data rows only: the NVMe and HDD query copies carry different
    # ##bcftools_normCommand provenance lines, everything else must match.
    if zcat "$work/immkv.pass2.vcf.gz" | grep -v '^#' | cmp - <(grep -v '^#' "$ref") >"$log" 2>&1; then
        echo "CMP OK $media $field" >>"$log"
    else
        echo "CMP FAILED $media $field" >>"$log"
        return 1
    fi
}

# 1. All immkv outputs against the verified references.
for media in nvme hdd; do
    cmp_stream "$media" 17 "$HDD_ROOT/test/17_fields/immkv.round2.pass2.vcf"
    cmp_stream "$media" 540 "$HDD_ROOT/test/540_fields/immkv.pass2.vcf"
    cmp_stream "$media" 568 "$HDD_ROOT/test/568_fields/immkv.pass2.vcf"
done

# 2. Full field comparison with the fresh echtvar output (HDD 17 and 540).
verify_pair() { # media field
    local media=$1 field=$2
    local work src ev iv tsv log
    work=$(work_dir "$media" "$field")
    src=$(source_dir "$media" "$field")
    tsv=$(fields_tsv "$field")
    ev=$V/$media.$field.echtvar.txt
    iv=$V/$media.$field.immkv.txt
    log=$V/$media.$field.verifyvalues.log
    if [ -f "$log" ] && grep -q 'VERIFY_VALUES OK' "$log"; then
        log "skip verifyvalues $media $field"
        return
    fi
    log "build text $media $field"
    "$BCFTOOLS" view -h "$src/query/hg001.decomp.vcf.gz" >"$ev"
    local c
    for c in $(seq 1 22); do
        zcat "$work/echtvar.pass2.wk/an$c.vcf.gz" | grep -v '^#' >>"$ev"
    done
    zcat "$work/immkv.pass2.vcf.gz" >"$iv"
    "$VBIN" -ev "$ev" -iv "$iv" -fields "$tsv" | tee "$log"
    rm -f "$ev" "$iv"
}

verify_pair hdd 17
verify_pair hdd 540

log "all verification done"
