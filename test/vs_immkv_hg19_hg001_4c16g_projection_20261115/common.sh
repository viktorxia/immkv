#!/bin/bash
# Shared paths and the user-authorized 4 CPU / 16 GiB / zero swap protocol.
set -euo pipefail
set -o noclobber
export LC_ALL=C GOTELEMETRY=off
EXP_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
REPO=$(cd "$EXP_DIR/../.." && pwd)
NAME=$(basename "$EXP_DIR")
BUILD_ROOT="$REPO/tmp/$NAME"
RESULTS="$EXP_DIR/results"
PREVIOUS=vs_immkv_hg19_hg001_10g_release_20261114
FROZEN_ROOT="$REPO/tmp/$PREVIOUS"
FROZEN_FACTS="$REPO/test/$PREVIOUS/results/source-facts.json"
QUERY=${QUERY:-/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz}
FULL_LIB=/home/tian/immkv_test_nvme/gnomad/immkv
NARROW_LIB=/home/tian/immkv_test_nvme/test/17_fields/lib17
THREADS=4
BATCH=4096
CPUS=0-3
OLD_COMMIT=9e72562966c4caa1b2c813719aadb92033eede85
BEFORE_DIGEST=da913ce83951b2e00998343af092321fab16fcd433c7328b58cd91b60376c620

fail() { printf '%s\n' "$*" >&2; exit 1; }

fields_file() {
    printf '%s/test/vs_echtvar_hg19_gnomad_genomes_%s_fields/fields%s.tsv' "$REPO" "$1" "$1"
}

library_for() {
    case "$1" in
        narrow17) printf '%s' "$NARROW_LIB" ;;
        full17|full540|full568|full568_default) printf '%s' "$FULL_LIB" ;;
        *) fail "Unknown workload: $1" ;;
    esac
}

fields_for() {
    case "$1" in
        narrow17|full17) fields_file 17 ;;
        full540) fields_file 540 ;;
        full568|full568_default) fields_file 568 ;;
        *) fail "Unknown workload: $1" ;;
    esac
}

valid_version() {
    case "$1" in old|before|after) ;; *) fail "Unknown version: $1" ;; esac
}

valid_tag() {
    [[ "$1" =~ ^[a-zA-Z0-9_-]+$ ]] || fail "Invalid tag: $1"
}

require_binary() {
    valid_version "$1"
    [[ -f "$RESULTS/build.$1/finished" && -x "$BUILD_ROOT/bin/$1/$2" ]] ||
        fail "Build $1 first with build_versions.sh"
    (cd "$BUILD_ROOT/$1" && sha256sum --check --status "$RESULTS/build.$1/sources.sha256") ||
        fail "Frozen source hash changed for $1"
    (cd "$BUILD_ROOT/bin/$1" && sha256sum --check --status "$RESULTS/build.$1/binaries.sha256") ||
        fail "Binary hash changed for $1"
}

reserve() {
    mkdir -p "$RESULTS" "$BUILD_ROOT"
    mkdir "$RESULTS/$1" || fail "Result already exists; choose a new RUN_TAG to preserve it: $1"
}

limited() {
    local label=$1 mode=$2
    shift 2
    reserve "$label"
    printf '%s %s\n' "$(date -Is)" "$label"
    systemd-run --user --scope --collect -p MemoryMax=16G -p MemorySwapMax=0 \
        taskset -c "$CPUS" env -u IMMKV_FORCE_RECORD -u IMMKV_FORCE_COLUMNAR \
        -u IMMKV_DEBUG_PATH -u GODEBUG -u GOGC -u GOMEMLIMIT \
        GOMAXPROCS="$THREADS" LC_ALL=C \
        bash "$EXP_DIR/run_limited.sh" "$RESULTS/$label" "$mode" "$@" \
        > "$RESULTS/$label/scope.stdout" 2> "$RESULTS/$label/scope.stderr"
}

# Go must write only inside this experiment's scratch directory.
setup_go() {
    mkdir -p "$BUILD_ROOT/cache/go-build" "$BUILD_ROOT/cache/mod" \
        "$BUILD_ROOT/cache/gopath" "$BUILD_ROOT/cache/tmp"
    export GOCACHE="$BUILD_ROOT/cache/go-build" GOMODCACHE="$BUILD_ROOT/cache/mod"
    export GOPATH="$BUILD_ROOT/cache/gopath" GOTMPDIR="$BUILD_ROOT/cache/tmp"
    export TMPDIR="$BUILD_ROOT/cache/tmp" GOTOOLCHAIN=local GOWORK=off
    export GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly
    export GOTELEMETRY=off
}
