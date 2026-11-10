#!/bin/bash
# Regenerate the all-fields config and metadata table from the gnomAD header:
# all 568 fields (368 Integer + 172 Float + 16 String + 12 Flag). echtvar
# routes non-numeric values through its categorical dictionary path; the
# unique-value counts observed on real data are recorded in the README.
set -euo pipefail
source "$(dirname "$0")/common.sh"
cd "$REPO"
python3 "$REPO/test/genfields.py" --sites "$GNOMAD_SITES/gnomad.genomes.r2.1.1.sites.22.vcf.bgz" \
    --out "$EXP_DIR/fields568" --mode all
