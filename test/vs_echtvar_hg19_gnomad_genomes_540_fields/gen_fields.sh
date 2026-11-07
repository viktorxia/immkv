#!/bin/bash
# Regenerate the 540-field (all numeric) config and metadata table from the
# gnomAD header. Any sites file works as header source; chr22 is the smallest.
set -euo pipefail
source "$(dirname "$0")/common.sh"
cd "$REPO"
python3 "$REPO/test/genfields.py" --sites "$GNOMAD_SITES/gnomad.genomes.r2.1.1.sites.22.vcf.bgz" \
    --out "$EXP_DIR/fields540" --mode numeric
