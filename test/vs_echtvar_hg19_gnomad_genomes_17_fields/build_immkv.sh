#!/bin/bash
# Build the immkv side from raw gnomAD: a 17-field narrow library (exactly the
# same information content as the echtvar archives -> library-to-library
# comparison). Single-threaded, ~55 min for all 22 chromosomes.
#
# NARROW_ONLY=<chrom> restricts to one chromosome (smoke test).
set -euo pipefail
source "$(dirname "$0")/common.sh"
cd "$REPO"

echo "== build_immkv (narrow 17-field) start $(date -Is) =="
go run ./test/buildnarrow -gnomad "$GNOMAD_SITES" -out "$LIB17" -fields "$EXP_DIR/fields17.tsv"
echo "== build_immkv done $(date -Is), $(du -sh "$LIB17" | cut -f1) =="
