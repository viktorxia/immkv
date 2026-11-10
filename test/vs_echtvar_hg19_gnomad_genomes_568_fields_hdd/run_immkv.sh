#!/bin/bash
# Time the immkv side under the controlled protocol: one process, 4 threads,
# plain-text output, same decomposed query as run_echtvar.sh. Library = the
# pre-built 366 GiB full-load library with the FULL 568-field list: every
# stored byte is emitted (the heaviest immkv output load; ~70 GB of text per
# pass). Two passes, cold-ish then warm. Pass-1 output is deleted once pass-2
# has been written and verified row counts -- the disk cost of keeping both
# ~70 GB texts is not worth it.
#
# Output: $WORK/immkv.<pass>.vcf + .time
set -euo pipefail
source "$(dirname "$0")/common.sh"
exec > >(tee -a "$WORK/run_immkv_console.log") 2>&1
cd "$REPO"

go build -o "$WORK/annotate" ./test/annotate

for pass in pass1 pass2; do
    echo "== immkv $pass $(date -Is) =="
    /usr/bin/time -v -o "$WORK/immkv.$pass.time" \
        systemd-run --user --scope --collect -p MemoryMax=10G -p MemorySwapMax=0 \
        "$WORK/annotate" -in "$QUERY" -lib "$FULL_LIB" \
            -out "$WORK/immkv.$pass.vcf" -fields "$EXP_DIR/fields568.tsv" \
            -threads 4 -batch 4096
    grep -E "Elapsed|User time|System time|Percent|Maximum resident" "$WORK/immkv.$pass.time"
done
# Keep only the pass-2 text for verification; timings live in the .time logs.
rm -f "$WORK/immkv.pass1.vcf"
echo "== run_immkv done $(date -Is) =="
