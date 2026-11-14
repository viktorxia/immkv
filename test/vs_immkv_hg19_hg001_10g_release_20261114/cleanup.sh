#!/bin/bash
# Delete only regenerable diagnostic query text. Keep source snapshots, binaries,
# measured logs, compressed full outputs and built libraries.
set -euo pipefail
source "$(dirname "$0")/common.sh"
if [[ -e "$BUILD_ROOT/query-chr22.vcf" ]]; then
    rm "$BUILD_ROOT/query-chr22.vcf"
fi
