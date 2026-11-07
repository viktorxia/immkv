#!/bin/bash
# Build the echtvar side from raw gnomAD: per chromosome, decompose (bcftools
# norm -m -both) -> encode with fields568.json (ALL 568 INFO fields) -> delete
# the intermediate. 4 chromosomes in flight, NOT 8: the categorical
# dictionaries scale with variant count (chr22 pilot: 5.9 GiB peak RSS for
# 3.57M variants; ~3.5x that for chr1), and 8 concurrent large chromosomes
# would exceed the 96 GiB machine. Existing archives are skipped (resumable).
#
# chr22 is already seeded from the pilot run (see pilot_encode_chr22.log).
#
# Storage: archives ~90 GB total + ~100 GB transient decomp intermediates.
# The script refuses to start with less than 300 GB free on the work
# filesystem.
#
# Wall time ~2-3h; run under nohup. Per-chromosome timings land in the log.
set -euo pipefail
source "$(dirname "$0")/common.sh"

free_gb=$(df -BG --output=avail "$WORK" | tail -1 | tr -dc '0-9')
if [ "$free_gb" -lt 300 ]; then
    echo "only ${free_gb}G free under $WORK; need >=300G for the 568-field build" >&2
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
    if ! "$ECHTVAR_BIN" encode "$ARCHIVES/chr$c.echtvar.zip.part" "$EXP_DIR/fields568.json" \
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

echo "== build_echtvar (568) start $(date -Is), ${free_gb}G free =="
seq 1 22 | xargs -P 4 -I{} bash -c 'worker {}'
total=$(du -sh "$ARCHIVES" | cut -f1)
echo "== build_echtvar (568) done $(date -Is), archives total $total =="
