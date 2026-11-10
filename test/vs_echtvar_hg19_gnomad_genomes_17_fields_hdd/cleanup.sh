#!/bin/bash
# Free the multi-GB run artifacts after verification: text outputs, per-pass
# scratch, decompose leftovers. Keeps .time reports, logs, and the archives
# (re-verification of hit sets against the archives stays possible; the texts
# are reproducible from the run scripts).
set -euo pipefail
source "$(dirname "$0")/common.sh"

"$(dirname "$0")"/collect_results.sh

rm -rf "$WORK/wk_echtvar.pass1" "$WORK/wk_echtvar.pass2" "$WORK/decomp"/*
rm -f "$WORK"/immkv.round*.vcf "$WORK/annotate" "$WORK/chr22.slice.vcf.gz" "$WORK/chr22.new.vcf" "$WORK/chr22.old.vcf"
echo "cleaned; remaining under $WORK:"
du -sh "$WORK"/* 2>/dev/null
df -BG --output=avail "$WORK" | tail -1
