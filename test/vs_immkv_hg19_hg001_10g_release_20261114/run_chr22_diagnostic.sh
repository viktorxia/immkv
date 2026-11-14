#!/bin/bash
# Diagnostic only: the whole-genome run encountered a changed NVMe I/O regime.
# Use exactly the same chr22 rows for both binaries and do not mix these numbers
# with the whole-genome results.
set -euo pipefail
source "$(dirname "$0")/common.sh"
query22="$BUILD_ROOT/query-chr22.vcf"
if [[ -e "$query22" ]]; then
    printf 'Diagnostic query already exists: %s\n' "$query22" >&2
    exit 1
fi
/home/tian/miniforge3/envs/transvar/bin/bcftools view -r 22 -Ov \
    -o "$query22.tmp" "$QUERY"
mv "$query22.tmp" "$query22"
sha256sum "$query22" > "$RESULTS/chr22-query.sha256"
for mode in default explicit; do
    for repeat in 1 2 3; do
        versions='baseline candidate'
        [[ "$repeat" == 2 ]] && versions='candidate baseline'
        for version in $versions; do
            flag=true
            [[ "$mode" == default ]] && flag=false
            label="diagnostic.chr22.$mode.$version.r$repeat"
            printf '%s %s\n' "$(date -Is)" "$label"
            limited "$label" timeout 90s "$BUILD_ROOT/bin/$version/query-only" \
                -query "$query22" -lib "$FULL_LIB" -fields "$(fields_file 568)" \
                -threads "$THREADS" -batch "$BATCH" "-columns=$flag"
        done
    done
done
printf '%s\n' "$(date -Is)" > "$RESULTS/diagnostic-chr22.finished"
