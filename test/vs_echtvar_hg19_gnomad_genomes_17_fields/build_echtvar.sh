#!/bin/bash
# Build the echtvar side from raw gnomAD: per chromosome, decompose (bcftools
# norm -m -both) -> encode with fields17.json -> delete the intermediate.
# 4 chromosomes in flight (xargs -P 4), so up to 4 concurrent norm (each
# --threads 4) or encode (single-threaded) processes. Existing archives are
# skipped, so the script is resumable after a reboot.
#
# Wall time ~5-8h; run under nohup. Per-chromosome timings land in the log:
# they are the "library preparation account" recorded in the README.
set -euo pipefail
source "$(dirname "$0")/common.sh"

worker() {
    local c=$1
    if [ -f "$ARCHIVES/chr$c.echtvar.zip" ]; then
        echo "chr$c: archive exists, skip"
        return 0
    fi
    local s=$(date +%s)
    "$BCFTOOLS" norm -m -both -Oz --threads 4 -o "$DECOMP/chr$c.decomp.vcf.gz" \
        "$GNOMAD_SITES/gnomad.genomes.r2.1.1.sites.$c.vcf.bgz" 2> "$DECOMP/chr$c.norm.log"
    local m=$(date +%s)
    if ! "$ECHTVAR_BIN" encode "$ARCHIVES/chr$c.echtvar.zip" "$EXP_DIR/fields17.json" \
        "$DECOMP/chr$c.decomp.vcf.gz" > "$DECOMP/chr$c.encode.log" 2>&1; then
        echo "chr$c: encode FAILED (see $DECOMP/chr$c.encode.log)"
        return 1
    fi
    local e=$(date +%s)
    rm -f "$DECOMP/chr$c.decomp.vcf.gz" "$DECOMP/chr$c.decomp.vcf.gz.csi" "$DECOMP/chr$c.decomp.vcf.gz.tbi"
    echo "chr$c: norm $((m-s))s encode $((e-m))s $(du -h "$ARCHIVES/chr$c.echtvar.zip" | cut -f1)"
}
export -f worker
export GNOMAD_SITES ECHTVAR_BIN ARCHIVES DECOMP EXP_DIR BCFTOOLS

echo "== build_echtvar start $(date -Is) =="
seq 1 22 | xargs -P 4 -I{} bash -c 'worker {}'
total=$(du -sh "$ARCHIVES" | cut -f1)
echo "== build_echtvar done $(date -Is), archives total $total =="
