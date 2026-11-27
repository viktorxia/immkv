#!/usr/bin/env bash
# Freeze the current post-fix source snapshot, build annotate and verifyvalues, and run unit tests.
set -euo pipefail
export LC_ALL=C GOTELEMETRY=off GOTOOLCHAIN=local GOENV=off GOWORK=off
EXP_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
REPO=$(cd -- "$EXP_DIR/../.." && pwd)
NAME=vs_immkv_hg19_hg001_4c16g_cold_postfix_20261125
[[ "$EXP_DIR" == "$REPO/test/$NAME" ]] || { printf 'Unexpected experiment path\n' >&2; exit 1; }
BUILD_ROOT="$REPO/tmp/$NAME"
RESULTS="$EXP_DIR/results"
RECORD="$RESULTS/preparation"
BASE_COMMIT=904948f0db6f59d94e8a69d97d300fba82aae24b
ENGINE_BASE_COMMIT=d2eac357f3fc126755b43ff8b44a64dc55bc0ec2
OLD_CACHE="$REPO/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/cache/mod"
[[ $(git -C "$REPO" rev-parse HEAD) == "$BASE_COMMIT" ]] || { printf 'HEAD commit mismatch\n' >&2; exit 1; }
[[ ! -e "$RECORD" && ! -e "$BUILD_ROOT/source" && ! -e "$BUILD_ROOT/source.partial" ]] || { printf 'Preparation already reserved; preserve its evidence\n' >&2; exit 1; }
mkdir -p "$RESULTS" "$BUILD_ROOT/bin" "$BUILD_ROOT/cache/"{go-build,mod,gopath,tmp}
mkdir "$RECORD"
export GOCACHE="$BUILD_ROOT/cache/go-build" GOMODCACHE="$BUILD_ROOT/cache/mod"
export GOPATH="$BUILD_ROOT/cache/gopath" GOTMPDIR="$BUILD_ROOT/cache/tmp" TMPDIR="$BUILD_ROOT/cache/tmp"
export GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly
export EXP_DIR REPO BUILD_ROOT RECORD BASE_COMMIT ENGINE_BASE_COMMIT OLD_CACHE
trap 'rc=$?; printf "%s\n" "$rc" > "$RECORD/exit-status"; date -u +%FT%TZ > "$RECORD/finished"' EXIT
date -u +%FT%TZ > "$RECORD/started"
[[ -d "$OLD_CACHE" ]] || { printf 'Existing repository-local module cache missing\n' >&2; exit 1; }
cp -a --reflink=auto "$OLD_CACHE/." "$GOMODCACHE/"
printf '%s\n' "$OLD_CACHE" > "$RECORD/module-cache-origin"
git -C "$REPO" diff HEAD -- col_test.go colread.go iouring_linux.go iouring_linux_test.go stream.go > "$RECORD/working-tree.patch"
mkdir -p "$BUILD_ROOT/source.partial/test"
cp -a "$REPO"/go.mod "$REPO"/go.sum "$REPO"/*.go "$BUILD_ROOT/source.partial/"
for sub in annotate verifyvalues buildfull buildnarrow buildvcf pointlookup samplecolumns \
           vs_echtvar_hg19_gnomad_genomes_17_fields vs_echtvar_hg19_gnomad_genomes_540_fields vs_echtvar_hg19_gnomad_genomes_568_fields; do
    mkdir -p "$BUILD_ROOT/source.partial/test/$sub"
    cp -a "$REPO/test/$sub"/*.go "$BUILD_ROOT/source.partial/test/$sub/" 2>/dev/null || true
    cp -a "$REPO/test/$sub"/fields*.tsv "$BUILD_ROOT/source.partial/test/$sub/" 2>/dev/null || true
done
mv -- "$BUILD_ROOT/source.partial" "$BUILD_ROOT/source"
python3 -I - <<'PY'
import hashlib, json, os, subprocess
from pathlib import Path
repo, root, record = (Path(os.environ[k]) for k in ('REPO', 'BUILD_ROOT', 'RECORD'))
def digest(path):
    h = hashlib.sha256()
    with path.open('rb') as f:
        for b in iter(lambda: f.read(1048576), b''):
            h.update(b)
    return h.hexdigest()
source = root / 'source'
entries = []
for p in sorted(source.rglob('*')):
    if p.is_file():
        entries.append(f'{digest(p)}  {p.relative_to(source).as_posix()}\n')
manifest = ''.join(entries)
(record / 'sources.sha256').write_text(manifest)
patch_sha256 = digest(record / 'working-tree.patch')
sources_manifest_sha256 = hashlib.sha256(manifest.encode()).hexdigest()
facts = {
    'base_commit': os.environ['BASE_COMMIT'],
    'engine_base_commit': os.environ['ENGINE_BASE_COMMIT'],
    'working_tree_patch': 'working-tree.patch',
    'working_tree_patch_sha256': patch_sha256,
    'source_directory': str(source),
    'source_file_count': len(entries),
    'sources_manifest_sha256': sources_manifest_sha256,
    'module_cache_origin': os.environ['OLD_CACHE'],
    'module_cache_mode': 'private copy; original cache is read only',
    'build_flags': ['-trimpath', '-buildvcs=false'],
    'GOFLAGS': os.environ['GOFLAGS'],
    'timed_runs': False,
    'cache_clearing': False,
}
(record / 'source-facts.json').write_text(json.dumps(facts, indent=2) + '\n')
PY
{
    date -u +%FT%TZ
    uname -a
    go version
    go env -json
    command -v go gcc g++ ld python3 git tar
    gcc --version
    g++ --version
    ld --version
    python3 --version
    git --version
    tar --version
} > "$RECORD/environment.txt" 2>&1
[[ $(go env CGO_ENABLED) == 1 ]] || { printf 'Expected default CGO-enabled build\n' >&2; exit 1; }
(
    set -x
    go -C "$BUILD_ROOT/source" list -m
    go -C "$BUILD_ROOT/source" build -trimpath -buildvcs=false -o "$BUILD_ROOT/bin/annotate.partial" ./test/annotate
    mv -- "$BUILD_ROOT/bin/annotate.partial" "$BUILD_ROOT/bin/annotate"
    go -C "$BUILD_ROOT/source" build -trimpath -buildvcs=false -o "$BUILD_ROOT/bin/verifyvalues.partial" ./test/verifyvalues
    mv -- "$BUILD_ROOT/bin/verifyvalues.partial" "$BUILD_ROOT/bin/verifyvalues"
) > "$RECORD/build.log" 2>&1
(
    set -x
    go -C "$BUILD_ROOT/source" test -buildvcs=false -count=1 -v -run "TestDriveDirectInFlightRespectsRingDepth|TestPreferColumnarCountsColumnRanges" .
    go -C "$BUILD_ROOT/source" test -buildvcs=false -count=1 -v ./test/annotate ./test/verifyvalues
) > "$RECORD/unit-tests.log" 2>&1
(cd -- "$BUILD_ROOT/source" && sha256sum --check --status "$RECORD/sources.sha256")
(cd -- "$BUILD_ROOT/bin" && sha256sum annotate verifyvalues) > "$RECORD/binaries.sha256"
{
    go version -m "$BUILD_ROOT/bin/annotate"
    go version -m "$BUILD_ROOT/bin/verifyvalues"
} > "$RECORD/build-info.txt"
{
    sha256sum "$(command -v go)" "$(command -v gcc)" "$(command -v g++)" "$(command -v ld)"
    sha256sum "$(go env GOTOOLDIR)/compile" "$(go env GOTOOLDIR)/link" "$(go env GOTOOLDIR)/cgo"
} > "$RECORD/compiler-binaries.sha256"
"$BUILD_ROOT/bin/annotate" -h > "$RECORD/annotate-help.txt" 2>&1
"$BUILD_ROOT/bin/verifyvalues" -h > "$RECORD/verifyvalues-help.txt" 2>&1
printf 'Frozen annotate: %s\nFrozen verifyvalues: %s\nPreparation record: %s\n' "$BUILD_ROOT/bin/annotate" "$BUILD_ROOT/bin/verifyvalues" "$RECORD"
