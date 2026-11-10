#!/bin/bash
# Prepare the query: decompose (split multiallelics) the GIAB HG001 benchmark
# VCF. Both echtvar (hard requirement) and immkv (libraries are decomposed)
# consume the decomposed form. Takes seconds.
set -euo pipefail
source "$(dirname "$0")/common.sh"

if [ ! -f "$HG001_SOURCE" ]; then
    echo "HG001 source not found: $HG001_SOURCE" >&2
    exit 1
fi
echo "== decomposing HG001 ($(date -Is)) =="
"$BCFTOOLS" norm -m -both -Oz --threads 4 -o "$QUERY" "$HG001_SOURCE" 2> "$WORK/query/norm.log"
"$BCFTOOLS" index -c -f "$QUERY"   # -r random access needs an index
rows=$(zcat "$QUERY" | grep -vc '^#')
echo "query rows: $rows (expect 3,933,714)"
