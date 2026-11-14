#!/bin/bash
# All experiment files and comments are in English.
set -euo pipefail
EXP_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
REPO=$(cd "$EXP_DIR/../.." && pwd)
NAME=$(basename "$EXP_DIR")
BUILD_ROOT="$REPO/tmp/$NAME"
WORK="/home/tian/immkv_test_nvme/test/$NAME"
QUERY=${QUERY:-/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz}
FULL_LIB=/home/tian/immkv_test_nvme/gnomad/immkv
NARROW_LIB=/home/tian/immkv_test_nvme/test/17_fields/lib17
RESULTS="$EXP_DIR/results"
THREADS=4
BATCH=4096
CPUS=0-3
BASELINE=9e72562

fields_file() {
    printf '%s/test/vs_echtvar_hg19_gnomad_genomes_%s_fields/fields%s.tsv' "$REPO" "$1" "$1"
}

library_for() {
    case "$1" in
        narrow17) printf '%s' "$NARROW_LIB" ;;
        full17|full540|full568|full568_default) printf '%s' "$FULL_LIB" ;;
        *) return 1 ;;
    esac
}

fields_for() {
    case "$1" in
        narrow17|full17) fields_file 17 ;;
        full540) fields_file 540 ;;
        full568|full568_default) fields_file 568 ;;
        *) return 1 ;;
    esac
}

limited() {
    local label=$1
    shift
    systemd-run --user --scope --collect -p MemoryMax=10G -p MemorySwapMax=0 \
        taskset -c "$CPUS" env -u IMMKV_FORCE_RECORD -u IMMKV_FORCE_COLUMNAR \
        -u IMMKV_DEBUG_PATH GOMAXPROCS="$THREADS" \
        bash "$EXP_DIR/run_limited.sh" "$RESULTS/$label" "$@"
}
