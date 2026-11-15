#!/bin/bash
# Build only the standard release annotate; the verifier is owned separately.
set -euo pipefail
source "$(dirname "$0")/common.sh"
[[ ! -e "$RESULTS/build" && ! -e "$BUILD_ROOT/source" ]] || fail 'Build already reserved; preserve existing evidence'
source_module_cache=$(go env GOMODCACHE)
mkdir -p "$RESULTS/build" "$BUILD_ROOT/bin"
record="$RESULTS/build"
date -Is > "$record/started"
setup_go
python3 -I "$EXP_DIR/snapshot.py" --repo "$REPO" --build "$BUILD_ROOT" \
    --record "$record" --commit "$RELEASE_COMMIT" --module-cache "$source_module_cache" \
    > "$record/snapshot.log" 2>&1
{
    date -Is
    go version
    go env CGO_ENABLED GOOS GOARCH GOCACHE GOMODCACHE GOFLAGS
    uname -a
    "$ECHTVAR_BIN" --version
    "$BCFTOOLS" --version
    sha256sum "$ECHTVAR_BIN" "$BCFTOOLS"
} > "$record/environment.txt" 2>&1
[[ $(go env CGO_ENABLED) == 1 ]] || fail 'Expected default CGO-enabled Go build'
(
    set -x
    go -C "$BUILD_ROOT/source" build -trimpath -o "$ANNOTATE.partial" ./test/annotate
    mv -n "$ANNOTATE.partial" "$ANNOTATE"
    cd "$BUILD_ROOT/source"
    sha256sum --check --status "$record/sources.sha256"
) > "$record/build.log" 2>&1
(cd "$BUILD_ROOT/bin" && sha256sum annotate) > "$record/binaries.sha256"
go version -m "$ANNOTATE" > "$record/build-info.txt"
date -Is > "$record/finished"
printf 'Release annotate ready: %s\n' "$ANNOTATE"
