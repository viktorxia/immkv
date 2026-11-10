#!/bin/bash
# Regenerate the 17-field config and metadata table from the gnomAD header.
set -euo pipefail
source "$(dirname "$0")/common.sh"
cd "$REPO"
python3 "$REPO/test/genfields.py" --sites "$GNOMAD_SITES/gnomad.genomes.r2.1.1.sites.22.vcf.bgz" \
    --out "$EXP_DIR/fields17" --mode numeric --keep "$EXP_DIR/keep17.txt"
