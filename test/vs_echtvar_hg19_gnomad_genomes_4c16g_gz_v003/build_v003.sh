#!/bin/bash
# Rebuild the v003 libraries in place on both media, then copy the NVMe build
# to HDD. Run from a clean state: the target library directories must not
# exist. The v0.0.2 libraries were deleted before this rebuild.
#
# Usage:
#
#	GNOMAD_SRC=<dir> build_v003.sh [nvme|hdd|both]
#
# GNOMAD_SRC defaults to /home/tian/ssd_annotation_database/gnomad/GRCh37/v2.1.1.
# The baseline v0.0.2 libraries were built in per-value framing mode with a
# trained dictionary (`IMMKV_PACK_KB=-1`); the v003 rebuild uses the same mode
# so the controlled matrix compares formats, not framing choices. Override with
# IMMKV_PACK_KB=64 (packed frames) for other experiments.
# The build itself is not memory-capped (per the repository experiment rules);
# only the timed matrix runs under 16G.
set -euo pipefail
source "$(dirname "$0")/common.sh"
cd "$REPO"

MEDIA=${1:-both}
SRC=${GNOMAD_SRC:-/home/tian/ssd_annotation_database/gnomad/GRCh37/v2.1.1}
export IMMKV_PACK_KB=${IMMKV_PACK_KB:--1}
BIN=$(media_root nvme)/test/4c16g_gz_v003/bin/buildvcf
mkdir -p "$(dirname "$BIN")"
go build -o "$BIN" ./test/buildvcf

build_one() { # root
    local root=$1
    local full=$root/gnomad/immkv
    local lib17=$root/test/17_fields/lib17
    if [ -e "$full" ] || [ -e "$lib17" ]; then
        echo "refusing: $full or $lib17 already exists" >&2
        exit 1
    fi
    echo "[$(date -Is)] building full library -> $full"
    "$BIN" -gnomad "$SRC" -out "$full"
    echo "[$(date -Is)] building 17-field library -> $lib17"
    "$BIN" -gnomad "$SRC" -out "$lib17" \
        -fields "$(fields_tsv 17)"
}

case "$MEDIA" in
nvme) build_one "$NVME_ROOT" ;;
hdd) build_one "$HDD_ROOT" ;;
both)
    build_one "$NVME_ROOT"
    echo "[$(date -Is)] copying NVMe libraries to HDD"
    mkdir -p "$HDD_ROOT/gnomad" "$HDD_ROOT/test/17_fields"
    rsync -a "$NVME_ROOT/gnomad/immkv" "$HDD_ROOT/gnomad/"
    rsync -a "$NVME_ROOT/test/17_fields/lib17" "$HDD_ROOT/test/17_fields/"
    ;;
*)
    echo "unknown media: $MEDIA (want nvme|hdd|both)" >&2
    exit 2
    ;;
esac
echo "[$(date -Is)] all builds done"
