#!/usr/bin/env bash
# Freeze the authorized tag, build both tools, and run their unit tests only.
set -euo pipefail
export LC_ALL=C GOTELEMETRY=off GOTOOLCHAIN=local GOENV=off GOWORK=off
EXP_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd)
REPO=$(cd -- "$EXP_DIR/../.." && pwd)
NAME=vs_echtvar_hg19_hg001_4c16g_cold_v005_20261115
[[ "$EXP_DIR" == "$REPO/test/$NAME" ]] || { printf 'Unexpected experiment path\n' >&2; exit 1; }
BUILD_ROOT="$REPO/tmp/$NAME"
RESULTS="$EXP_DIR/results"
RECORD="$RESULTS/preparation"
TAG=v0.0.5-perf-untested.1
COMMIT=d2eac357f3fc126755b43ff8b44a64dc55bc0ec2
OLD_CACHE="$REPO/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/cache/mod"
[[ $(git -C "$REPO" rev-parse "$TAG^{commit}") == "$COMMIT" ]] || { printf 'Tag target mismatch\n' >&2; exit 1; }
[[ ! -e "$RECORD" && ! -e "$BUILD_ROOT/source" && ! -e "$BUILD_ROOT/source.partial" ]] || { printf 'Preparation already reserved; preserve its evidence\n' >&2; exit 1; }
mkdir -p "$RESULTS" "$BUILD_ROOT/bin" "$BUILD_ROOT/cache/"{go-build,mod,gopath,tmp}
mkdir "$RECORD"
export GOCACHE="$BUILD_ROOT/cache/go-build" GOMODCACHE="$BUILD_ROOT/cache/mod"
export GOPATH="$BUILD_ROOT/cache/gopath" GOTMPDIR="$BUILD_ROOT/cache/tmp" TMPDIR="$BUILD_ROOT/cache/tmp"
export GOPROXY=off GOSUMDB=off GOFLAGS=-mod=readonly
export EXP_DIR REPO BUILD_ROOT RECORD TAG COMMIT OLD_CACHE
trap 'rc=$?; printf "%s\n" "$rc" > "$RECORD/exit-status"; date -u +%FT%TZ > "$RECORD/finished"' EXIT
date -u +%FT%TZ > "$RECORD/started"
# A private cache copy prevents even module lock or compiler writes in old dirs.
[[ -d "$OLD_CACHE" ]] || { printf 'Existing repository-local module cache missing\n' >&2; exit 1; }
cp -a --reflink=auto "$OLD_CACHE/." "$GOMODCACHE/"
printf '%s\n' "$OLD_CACHE" > "$RECORD/module-cache-origin"
git -C "$REPO" archive --format=tar "$COMMIT" > "$BUILD_ROOT/source.tar.partial"
mv -- "$BUILD_ROOT/source.tar.partial" "$BUILD_ROOT/source.tar"
mkdir "$BUILD_ROOT/source.partial"
tar -xf "$BUILD_ROOT/source.tar" -C "$BUILD_ROOT/source.partial"
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
facts = {
    'tag': os.environ['TAG'], 'commit': os.environ['COMMIT'],
    'git_tree': subprocess.check_output(['git', '-C', str(repo), 'rev-parse', os.environ['COMMIT'] + '^{tree}'], text=True).strip(),
    'archive': str(root / 'source.tar'), 'archive_sha256': digest(root / 'source.tar'),
    'source_directory': str(source), 'source_file_count': len(entries),
    'sources_manifest_sha256': hashlib.sha256(manifest.encode()).hexdigest(),
    'module_cache_origin': os.environ['OLD_CACHE'],
    'module_cache_mode': 'private copy; original cache is read only',
    'build_flags': ['-trimpath', '-buildvcs=false'], 'GOFLAGS': os.environ['GOFLAGS'],
    'functional_full_suite': 'Previously passed; not rerun during preparation',
    'timed_runs': False, 'cache_clearing': False,
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
