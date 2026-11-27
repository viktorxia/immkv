#!/usr/bin/env bash
# Shared interfaces for the cold-cache 4 physical CPU / 16 GiB post-fix experiment.
set -euo pipefail
export LC_ALL=C GOTELEMETRY=off
EXP_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
REPO=$(cd -- "$EXP_DIR/../.." && pwd)
NAME=$(basename -- "$EXP_DIR")
BUILD_ROOT="$REPO/tmp/$NAME"
RESULTS="$EXP_DIR/results"
BASE_COMMIT=904948f0db6f59d94e8a69d97d300fba82aae24b
ENGINE_BASE_COMMIT=d2eac357f3fc126755b43ff8b44a64dc55bc0ec2
ANNOTATE="$BUILD_ROOT/bin/annotate"
VERIFIER="$BUILD_ROOT/bin/verifyvalues"
export BCFTOOLS=${BCFTOOLS:-/home/tian/miniforge3/envs/transvar/bin/bcftools}
export HDD_ROOT=${HDD_ROOT:-/mnt/data6t/tian/immkv_test_hdd}
export NVME_ROOT=${NVME_ROOT:-/home/tian/immkv_test_nvme}
export CACHE_SOCKET=${CACHE_SOCKET:-$BUILD_ROOT/cache.sock}
export THREADS=4 CPUS=0-3 BATCH=4096
export EXP_DIR REPO BUILD_ROOT RESULTS BASE_COMMIT ENGINE_BASE_COMMIT ANNOTATE VERIFIER
EXP="$EXP_DIR"
export EXP
fail() { printf '%s\n' "$*" >&2; exit 1; }
valid_tag() { [[ "$1" =~ ^[a-zA-Z0-9_-]{1,96}$ ]] || fail "Invalid tag: $1"; }
media_root() {
    case "$1" in hdd) printf '%s' "$HDD_ROOT";; nvme) printf '%s' "$NVME_ROOT";; *) fail "Unknown media: $1";; esac
}
source_dir() { printf '%s/test/%s_fields' "$(media_root "$1")" "$2"; }
query_for() { printf '%s/query/hg001.decomp.vcf.gz' "$(source_dir "$1" "$2")"; }
fields_file() { printf '%s/test/vs_echtvar_hg19_gnomad_genomes_%s_fields/fields%s.tsv' "$REPO" "$1" "$1"; }
fields_tsv() { fields_file "$@"; }
library() {
    case "$2" in 17) printf '%s/test/17_fields/lib17' "$(media_root "$1")";;
        540|568) printf '%s/gnomad/immkv' "$(media_root "$1")";; *) fail "Unknown field set: $2";; esac
}
setup_go() {
    mkdir -p "$BUILD_ROOT/cache/"{go-build,mod,gopath,tmp}
    export GOCACHE="$BUILD_ROOT/cache/go-build" GOMODCACHE="$BUILD_ROOT/cache/mod"
    export GOPATH="$BUILD_ROOT/cache/gopath" GOTMPDIR="$BUILD_ROOT/cache/tmp" TMPDIR="$BUILD_ROOT/cache/tmp"
    export GOTOOLCHAIN=local GOWORK=off GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly
}
require_binary() {
    [[ -f "$RESULTS/preparation/finished" && -f "$RESULTS/preparation/exit-status" &&
       $(<"$RESULTS/preparation/exit-status") == 0 && -x "$ANNOTATE" && -x "$VERIFIER" ]] || fail 'Successful build_tools.sh preparation is required'
    (cd -- "$BUILD_ROOT/bin" && sha256sum --check --status "$RESULTS/preparation/binaries.sha256") || fail 'Binary changed since preparation'
    (cd -- "$BUILD_ROOT/source" && sha256sum --check --status "$RESULTS/preparation/sources.sha256") || fail 'Frozen source changed since preparation'
}
limited() {
    local label=$1 media=$2
    shift 2
    valid_tag "$label"
    [[ $# -gt 0 ]] || fail 'limited requires a workload command'
    local duration
    case "$media" in hdd) duration=${HDD_TIMEOUT_SECONDS:-3600};; nvme) duration=${NVME_TIMEOUT_SECONDS:-3600};; *) fail "Unknown media: $media";; esac
    duration=${RUN_TIMEOUT_SECONDS:-$duration}
    [[ "$duration" =~ ^[1-9][0-9]*$ ]] || fail 'Timeout must be a positive number of seconds'
    mkdir -p "$RESULTS" "$BUILD_ROOT/runs"
    local invocation_lock
    exec {invocation_lock}> "$BUILD_ROOT/invocation.lock"
    flock -n "$invocation_lock" || fail 'Another measured invocation owns the cache/scope lock'
    mkdir "$RESULTS/$label" || fail "Result exists; preserve it: $label"
    local record="$RESULTS/$label" scope rc observer observer_rc
    scope="immkv-cold-${label}-$(date +%s%N)-$BASHPID"
    printf '%s\n' "$scope.scope" > "$record/scope-unit"
    printf '%s\n' "$duration" > "$record/timeout-seconds"
    if ! python3 -I "$EXP_DIR/env_tools.py" cache-drop "$record"; then
        flock -u "$invocation_lock"
        exec {invocation_lock}>&-
        return 1
    fi
    python3 -I "$EXP_DIR/env_tools.py" watch "$record" > "$record/observer.stdout" 2> "$record/observer.stderr" &
    observer=$!
    if /usr/bin/time -v -o "$record/time-scope.txt" \
        systemd-run --user --scope --collect --unit="$scope" -p MemoryMax=16G -p MemorySwapMax=0 -p IOAccounting=yes \
        taskset -c "$CPUS" env -u IMMKV_FORCE_RECORD -u IMMKV_FORCE_COLUMNAR \
        -u IMMKV_DEBUG_PATH -u GODEBUG -u GOGC -u GOMEMLIMIT \
        GOMAXPROCS=4 RUN_TIMEOUT_SECONDS="$duration" \
        bash "$EXP_DIR/run_limited.sh" "$record" "$@" \
        > "$record/scope.stdout" 2> "$record/scope.stderr"; then
        rc=0
    else
        rc=$?
    fi
    printf '%s\n' "$rc" > "$record/scope-exit-status"
    if wait "$observer"; then observer_rc=0; else observer_rc=$?; fi
    printf '%s\n' "$observer_rc" > "$record/observer-exit-status"
    local status
    if ! status=$(python3 -I "$EXP_DIR/resolve_status.py" "$record"); then
        flock -u "$invocation_lock"
        exec {invocation_lock}>&-
        return 1
    fi
    printf '%s\n' "$status" > "$record/resolved-status"
    flock -u "$invocation_lock"
    exec {invocation_lock}>&-
    [[ "$status" == ok && "$rc" == 0 && "$observer_rc" == 0 ]]
}
