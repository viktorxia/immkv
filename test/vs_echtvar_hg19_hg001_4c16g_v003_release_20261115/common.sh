#!/bin/bash
# Shared paths for the authorized 4 CPU / 16 GiB / zero-swap experiment.
set -euo pipefail
export LC_ALL=C GOTELEMETRY=off
EXP_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)
REPO=$(cd "$EXP_DIR/../.." && pwd)
NAME=$(basename "$EXP_DIR")
BUILD_ROOT="$REPO/tmp/$NAME"
RESULTS="$EXP_DIR/results"
RELEASE_COMMIT=840d36d80351f8b8651d9dc90abf1314bf5db21c
ANNOTATE="$BUILD_ROOT/bin/annotate"
export ECHTVAR_BIN=${ECHTVAR_BIN:-/home/tian/projects/immkv/tmp/echtvar/echtvar}
export BCFTOOLS=${BCFTOOLS:-/home/tian/miniforge3/envs/transvar/bin/bcftools}
export HDD_ROOT=${HDD_ROOT:-/mnt/data6t/tian/immkv_test_hdd}
export NVME_ROOT=${NVME_ROOT:-/home/tian/immkv_test_nvme}
export THREADS=4 CPUS=0-3 BATCH=4096
export EXP_DIR REPO BUILD_ROOT RESULTS RELEASE_COMMIT ANNOTATE
EXP="$EXP_DIR"
export EXP
fail() { printf '%s\n' "$*" >&2; exit 1; }
valid_tag() { [[ "$1" =~ ^[a-zA-Z0-9_-]+$ ]] || fail "Invalid tag: $1"; }
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
    [[ -f "$RESULTS/build/finished" && -x "$ANNOTATE" ]] || fail 'Run build_tools.sh first'
    (cd "$BUILD_ROOT/bin" && sha256sum --check --status "$RESULTS/build/binaries.sha256") || fail 'Binary changed since build'
}
# Reserve first; preserve failed and successful invocations equally.
limited() {
    local label=$1 media=$2
    shift 2
    valid_tag "$label"
    mkdir -p "$RESULTS" "$BUILD_ROOT/runs"
    mkdir "$RESULTS/$label" || fail "Result exists: $label"
    local record="$RESULTS/$label" duration scope rc
    case "$media" in hdd) duration=${HDD_TIMEOUT_SECONDS:-7200};; nvme) duration=${NVME_TIMEOUT_SECONDS:-1800};; *) fail "Unknown media: $media";; esac
    duration=${RUN_TIMEOUT_SECONDS:-$duration}
    scope="immkv-echtvar-${label}-$(date +%s)-$$"
    printf '%s\n' "$scope.scope" > "$record/scope-unit"
    python3 -I "$EXP_DIR/env_tools.py" watch "$record" &
    local observer=$!
    set +e
    systemd-run --user --scope --collect --unit="$scope" -p MemoryMax=16G -p MemorySwapMax=0 \
        taskset -c "$CPUS" env -u IMMKV_FORCE_RECORD -u IMMKV_FORCE_COLUMNAR \
        -u IMMKV_DEBUG_PATH -u GODEBUG -u GOGC -u GOMEMLIMIT \
        GOMAXPROCS=4 RUN_TIMEOUT_SECONDS="$duration" \
        bash "$EXP_DIR/run_limited.sh" "$record" "$@" \
        > "$record/scope.stdout" 2> "$record/scope.stderr"
    rc=$?
    set -e
    printf '%s\n' "$rc" > "$record/scope-exit-status"
    wait "$observer" || true
    # A killed wrapper may not be able to finalize; retain a truthful failure.
    if [[ ! -f "$record/status.json" ]]; then
        python3 -I - "$record" "$rc" <<'PY'
import json, subprocess, sys
from pathlib import Path
p, rc = Path(sys.argv[1]), int(sys.argv[2])
observed = p / 'memory-observer.json'
facts = json.loads(observed.read_text()) if observed.exists() else {}
events = dict(line.split() for line in facts.get('memory.events', '').splitlines())
oom = int(events.get('oom_kill', '0')) > 0 or int(events.get('oom_group_kill', '0')) > 0
unit = (p / 'scope-unit').read_text().strip()
journal = subprocess.run(['journalctl', '--user', '-u', unit, '--no-pager'], capture_output=True, text=True)
with (p / 'systemd-unit.log').open('x') as log:
    log.write(journal.stdout)
oom = oom or "Failed with result 'oom-kill'" in journal.stdout or 'killed by the OOM killer' in journal.stdout
for name in ('memory.events', 'memory.events.local', 'memory.peak'):
    if name in facts and not (p / name).exists():
        (p / name).write_text(facts[name] + '\n')
with (p / 'status.json').open('x') as f:
    json.dump({'status': 'oom' if oom else 'scope_failed', 'exit_code': rc,
               'memory_events_unavailable': not bool(events), 'last_observed_cgroup': facts,
               'systemd_journal_source': 'systemd-unit.log'}, f, indent=2)
    f.write('\n')
PY
    fi
    if [[ "$rc" != 0 ]]; then
        python3 -I "$EXP_DIR/resolve_status.py" "$record" > "$record/resolved-status"
    fi
    return "$rc"
}
