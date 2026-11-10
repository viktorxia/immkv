#!/bin/bash
# Copy the NVMe-built artifacts onto the HDD work root. Nothing is rebuilt.
# Each copy is verified by file count and byte total; a mismatch aborts.
# The 366 GiB full-load library is shared by all three HDD experiments, so it
# may already exist here when this script runs for a different field count.
set -euo pipefail
source "$(dirname "$0")/common.sh"

copy_tree() {
    local src=$1 dst=$2
    if [ -d "$dst" ]; then
        echo "exists, skip: $dst"
    else
        mkdir -p "$(dirname "$dst")"
        rsync -a --info=stats2 "$src" "$dst"
    fi
    local sc sb dc db
    sc=$(find "$src" -type f | wc -l); sb=$(find "$src" -type f -printf '%s\n' | awk '{a+=$1} END{printf "%.0f", a}')
    dc=$(find "$dst" -type f | wc -l); db=$(find "$dst" -type f -printf '%s\n' | awk '{a+=$1} END{printf "%.0f", a}')
    if [ "$sc" = "$dc" ] && [ "$sb" = "$db" ]; then
        echo "verified: $dst files=$dc bytes=$db"
    else
        echo "MISMATCH: $src -> $dst (files $sc/$dc, bytes $sb/$db)" >&2
        return 1
    fi
}

copy_tree "$NVME_ROOT/test/17_fields/lib17" "$LIB17"
copy_tree "$NVME_ROOT/test/17_fields/archives" "$ARCHIVES"
copy_tree "$NVME_ROOT/gnomad/immkv" "$HDD_ROOT/gnomad/immkv"
