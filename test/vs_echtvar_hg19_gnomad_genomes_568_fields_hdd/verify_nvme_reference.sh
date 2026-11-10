#!/bin/bash
# Verify the HDD run against an NVMe reference run: same binary, same library
# contents, same query, output written to NVMe; then byte-compare.
#
# Why not against echtvar: at 568 fields echtvar cannot run under the 10G
# cgroup (OOM at any concurrency), so no rival text exists to compare with.
# The NVMe 568 output was itself verified against the unconstrained echtvar
# output in ../vs_echtvar_hg19_gnomad_genomes_568_fields
# (2,234,349,552 field comparisons, zero mismatches), so a byte-exact match
# with the NVMe run transfers that verification to the HDD run.
#
# Requires ~70 GB free on the NVMe filesystem for the reference text; the
# reference is deleted after the comparison. The reference .time report is
# copied into results/ by collect_results.sh.
set -euo pipefail
source "$(dirname "$0")/common.sh"
REF_DIR=${REF_DIR:-$NVME_ROOT/test/568_fields_ref}
mkdir -p "$REF_DIR"
ANNOTATE=$WORK/annotate
if [ ! -x "$ANNOTATE" ]; then
    (cd "$REPO" && go build -o "$ANNOTATE" ./test/annotate)
fi

echo "== NVMe reference run $(date -Is) =="
/usr/bin/time -v -o "$REF_DIR/nvme.time" \
    systemd-run --user --scope --collect -p MemoryMax=10G -p MemorySwapMax=0 \
    "$ANNOTATE" -in "$QUERY" -lib "$NVME_ROOT/gnomad/immkv" \
        -out "$REF_DIR/immkv.nvme.vcf" -fields "$EXP_DIR/fields568.tsv" \
        -threads 4 -batch 4096
grep -E "Elapsed|User time" "$REF_DIR/nvme.time"

if cmp "$REF_DIR/immkv.nvme.vcf" "$WORK/immkv.pass2.vcf"; then
    echo "CMP_OK: HDD output is byte-identical to the NVMe reference"
    rm -f "$REF_DIR/immkv.nvme.vcf"
else
    echo "CMP_FAIL: HDD and NVMe outputs differ" >&2
    exit 1
fi
