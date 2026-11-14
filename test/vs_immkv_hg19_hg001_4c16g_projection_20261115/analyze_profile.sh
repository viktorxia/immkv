#!/bin/bash
# Analyze an existing profile separately from every performance measurement.
set -euo pipefail
source "$(dirname "$0")/common.sh"
[[ $# -ge 1 && $# -le 2 ]] || fail 'Usage: analyze_profile.sh before|after [run-tag]'
version=$1
tag=${2:-wgs}
valid_tag "$tag"
case "$version" in before|after) ;; *) fail 'Analyze before or after' ;; esac
require_binary "$version" query-only
label="profile.$tag.full540.$version"
profile="$RESULTS/$label/cpu.pprof"
[[ -f "$RESULTS/$label/finished" && -f "$profile" ]] || fail "Complete profile missing: $label"
reserve "analysis.$tag.full540.$version"
record="$RESULTS/analysis.$tag.full540.$version"
setup_go
export PPROF_TMPDIR="$BUILD_ROOT/cache/tmp"
go tool pprof -top -nodecount=60 "$BUILD_ROOT/bin/$version/query-only" "$profile" > "$record/top.txt" 2> "$record/top.stderr"
go tool pprof -top -cum -nodecount=60 "$BUILD_ROOT/bin/$version/query-only" "$profile" > "$record/cumulative.txt" 2> "$record/cumulative.stderr"
go tool pprof -list='projectRecord|projectRecordInto|GetBatch' \
    -source_path="$BUILD_ROOT/$version" -trim_path=github.com/viktorxia/immkv \
    "$BUILD_ROOT/bin/$version/query-only" "$profile" 2> "$record/projection.stderr" \
    | python3 -I -c 'import sys; sys.stdout.write(sys.stdin.read().encode("ascii", "backslashreplace").decode("ascii"))' \
    > "$record/projection.ascii.txt"
sha256sum "$profile" "$BUILD_ROOT/bin/$version/query-only" > "$record/inputs.sha256"
printf '%s\n' "$(date -Is)" > "$record/finished"
