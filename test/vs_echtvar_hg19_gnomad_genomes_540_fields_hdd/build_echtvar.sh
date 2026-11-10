#!/bin/bash
# Build the echtvar side from raw gnomAD: per chromosome, decompose (bcftools
# norm -m -both) -> encode with fields540.json (all 540 numeric INFO fields,
# echtvar's ceiling) -> delete the intermediate. 8 chromosomes in flight.
# Existing archives are skipped, so the script is resumable after a reboot.
#
# Storage: intermediate decomposed files are ~10-25 GB per chromosome; with 8
# in flight the peak is ~150-200 GB on top of the ~62 GB of archives. The
# script refuses to start with less than 300 GB free on the work filesystem.
#
# Wall time ~3-5h; run under nohup. Per-chromosome timings land in the log.
set -euo pipefail
source "$(dirname "$0")/common.sh"

free_gb=$(df -BG --output=avail "$WORK" | tail -1 | tr -dc '0-9')
if [ "$free_gb" -lt 300 ]; then
    echo "only ${free_gb}G free under $WORK; need >=300G for the 540-field build" >&2
    exit 1
fi

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
    if ! "$ECHTVAR_BIN" encode "$ARCHIVES/chr$c.echtvar.zip.part" "$EXP_DIR/fields540.json" \
        "$DECOMP/chr$c.decomp.vcf.gz" > "$DECOMP/chr$c.encode.log" 2>&1; then
        echo "chr$c: encode FAILED (see $DECOMP/chr$c.encode.log)"
        return 1
    fi
    local e=$(date +%s)
    mv "$ARCHIVES/chr$c.echtvar.zip.part" "$ARCHIVES/chr$c.echtvar.zip"   # atomic: a killed encode never leaves a truncated "complete" archive
    rm -f "$DECOMP/chr$c.decomp.vcf.gz" "$DECOMP/chr$c.decomp.vcf.gz.csi" "$DECOMP/chr$c.decomp.vcf.gz.tbi"
    echo "chr$c: norm $((m-s))s encode $((e-m))s $(du -h "$ARCHIVES/chr$c.echtvar.zip" | cut -f1)"
}
export -f worker
export GNOMAD_SITES ECHTVAR_BIN ARCHIVES DECOMP EXP_DIR BCFTOOLS

echo "== build_echtvar (540) start $(date -Is), ${free_gb}G free =="
seq 1 22 | xargs -P 8 -I{} bash -c 'worker {}'
total=$(du -sh "$ARCHIVES" | cut -f1)
echo "== build_echtvar (540) done $(date -Is), archives total $total =="
