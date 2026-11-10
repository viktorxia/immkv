#!/bin/bash
# Verify the controlled comparison: compare hit sets row by row (identical
# logic to the 17-field experiment; see that verify.sh for the rationale).
#
# Usage: verify.sh [echtvar.txt] [immkv.vcf]
#   defaults: $WORK/wk_echtvar.pass2/echtvar.pass2.txt and
#             $WORK/immkv.pass2.vcf
set -euo pipefail
source "$(dirname "$0")/common.sh"

exec > >(tee "$WORK/verify.log") 2>&1

EV=${1:-$WORK/wk_echtvar.pass2/echtvar.pass2.txt}
IV=${2:-$WORK/immkv.pass2.vcf}

echo "echtvar: $EV"
echo "immkv:   $IV"

awk -v ev="$EV" -v iv="$IV" '
BEGIN {
    while ((getline line < ev) > 0) {
        if (line ~ /^#/) continue
        n++
        eh = (line ~ /gnomad_AC=-1(;|$)/) ? 0 : (line ~ /gnomad_AC=/ ? 1 : 0)
        ehit[n] = eh
        if (eh) ec++
        eline[n] = line
    }
    m = 0
    while ((getline line < iv) > 0) {
        if (line ~ /^#/) continue
        m++
        ih = (line ~ /gnomad_/) ? 1 : 0
        if (ih) ic++
        if (m <= n && ih != ehit[m] && bad < 10) {
            bad++
            printf "DIVERGENCE row %d:\n  echtvar: %.200s\n  immkv:   %.200s\n", m, eline[m], line
        }
        if (m <= n && ih != ehit[m]) badtotal++
    }
    printf "rows: echtvar %d, immkv %d\n", n, m
    printf "hits: echtvar %d, immkv %d\n", ec, ic
    printf "diverging rows: %d\n", badtotal
    exit (n != m || badtotal != 0) ? 1 : 0
}' || { echo "VERIFY FAILED"; exit 1; }
echo "VERIFY OK: row universes match, hit sets identical"

# Full per-row, per-field value verification (test/verifyvalues): every
# whitelisted gnomad_* field of every row compared between the two outputs.
# Float fields within one multiplier-2e6 unit (5.2e-7, echtvar truncates);
# Integer/String exact; Flag presence-only.
cd "$REPO" && go build -o "$WORK/verifyvalues" ./test/verifyvalues
"$WORK/verifyvalues" -ev "$EV" -iv "$IV" -fields "$EXP_DIR/fields540.tsv" \
    || { echo "VERIFY FAILED (value level)"; exit 1; }
