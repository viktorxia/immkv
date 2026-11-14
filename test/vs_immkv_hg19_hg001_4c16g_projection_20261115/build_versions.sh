#!/bin/bash
# Build requested frozen versions; never build after implicitly.
set -euo pipefail
source "$(dirname "$0")/common.sh"
[[ $# -gt 0 ]] || set -- old before
versions=("$@")
for version in "${versions[@]}"; do
    valid_version "$version"
    [[ ! -e "$RESULTS/build.$version" && ! -e "$BUILD_ROOT/$version" && ! -e "$BUILD_ROOT/$version.partial" ]] ||
        fail "Build evidence already exists for $version; preserve it"
done
# Explicit `after` is only invoked after the parent reports production ready.
source_module_cache=$(go env GOMODCACHE)
setup_go
for version in "${versions[@]}"; do
    reserve "build.$version"
    record="$RESULTS/build.$version"
    printf '%s\n' "$(date -Is)" > "$record/started"
    python3 -I "$EXP_DIR/snapshot.py" "$version" --repo "$REPO" --experiment "$EXP_DIR" \
        --build "$BUILD_ROOT" --frozen "$FROZEN_ROOT" --facts "$FROZEN_FACTS" \
        --module-cache "$source_module_cache" --old-commit "$OLD_COMMIT" --before-digest "$BEFORE_DIGEST" \
        > "$record/snapshot.log" 2>&1
    python3 -I "$EXP_DIR/source_snapshot.py" retain "$version" > "$record/source-retention.log"
    mkdir -p "$BUILD_ROOT/bin/$version"
    {
        date -Is
        go version
        go env CGO_ENABLED GOOS GOARCH GOCACHE GOMODCACHE GOFLAGS
        uname -a
        lscpu -e=CPU,CORE,SOCKET,ONLINE
        findmnt -T "$FULL_LIB" -o SOURCE,FSTYPE,TARGET
        lsblk -d -o NAME,ROTA,MODEL,SIZE
        sha256sum "$QUERY" "$(fields_file 17)" "$(fields_file 540)" "$(fields_file 568)"
    } > "$record/environment.txt" 2>&1
    (
        set -x
        go -C "$BUILD_ROOT/$version" build -trimpath ./...
        go -C "$BUILD_ROOT/$version" test ./...
        for tool in query-only annotate verifyvalues; do
            target="$BUILD_ROOT/bin/$version/$tool"
            go -C "$BUILD_ROOT/$version" build -trimpath -o "$target.partial" "./cmd/$tool"
            mv -n "$target.partial" "$target"
        done
        CGO_ENABLED=0 go -C "$BUILD_ROOT/$version" build -trimpath ./...
        CGO_ENABLED=0 go -C "$BUILD_ROOT/$version" test ./...
        (cd "$BUILD_ROOT/$version" && sha256sum --check "$record/sources.sha256")
    ) > "$record/build.log" 2>&1
    (cd "$BUILD_ROOT/bin/$version" && sha256sum annotate query-only verifyvalues) > "$record/binaries.sha256"
    for tool in query-only annotate verifyvalues; do
        go version -m "$BUILD_ROOT/bin/$version/$tool"
    done > "$record/build-info.txt"
    printf '%s\n' "$(date -Is)" > "$record/finished"
    printf 'Built %s: %s/bin/%s\n' "$version" "$BUILD_ROOT" "$version"
done
