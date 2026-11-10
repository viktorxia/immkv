#!/bin/bash
# Run the read-only base.col phase-1 sampler on gnomAD v2.1.1 chr22.
# All outputs go to results/; the source VCF and the full-load library are
# read-only inputs and are never modified.
set -euo pipefail

REPO=$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)
EXP_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
GNOMAD_SITES=${GNOMAD_SITES:-/home/tian/ssd_annotation_database/gnomad/GRCh37/v2.1.1}
FULL_LIB=${FULL_LIB:-$HOME/immkv_test_nvme/gnomad/immkv}
VCF=$GNOMAD_SITES/gnomad.genomes.r2.1.1.sites.22.vcf.bgz
OUT=$EXP_DIR/results

mkdir -p "$OUT"
cd "$REPO"

go run ./test/samplecolumns \
    -vcf "$VCF" \
    -lib "${FULL_LIB:-}" \
    -chrom chr22 \
    -bin 1000000 \
    -level 9 \
    -max-examples 8 \
    -out "$OUT" 2>&1 | tee "$OUT/run.log"
