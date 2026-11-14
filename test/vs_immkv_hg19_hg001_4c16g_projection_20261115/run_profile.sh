#!/bin/bash
# CPU profiling is diagnostic and has no /usr/bin/time measurement.
set -euo pipefail
source "$(dirname "$0")/common.sh"
[[ $# -ge 1 && $# -le 2 ]] || fail 'Usage: run_profile.sh before|after [repeat]'
version=$1
case "$version" in before|after) ;; *) fail 'Profile before or after' ;; esac
repeat=${2:-1}
[[ "$repeat" =~ ^[1-9][0-9]*$ ]] || fail 'repeat must be positive'
require_binary "$version" query-only
tag=${RUN_TAG:-wgs}
valid_tag "$tag"
label="profile.$tag.full540.$version"
profile_dir="$BUILD_ROOT/profiles/$label"
[[ ! -e "$profile_dir" && ! -e "$RESULTS/$label" ]] || fail "Profile already exists: $label"
mkdir -p "$BUILD_ROOT/profiles"
mkdir "$profile_dir"
limited "$label" untimed "$BUILD_ROOT/bin/$version/query-only" \
    -query "$QUERY" -lib "$FULL_LIB" -fields "$(fields_file 540)" \
    -threads "$THREADS" -batch "$BATCH" -repeat "$repeat" \
    -cpuprofile "$profile_dir/cpu.pprof"
python3 -I "$EXP_DIR/retain_profile.py" "$version" "$tag"
sha256sum "$QUERY" "$(fields_file 540)" "$RESULTS/$label/cpu.pprof" > "$RESULTS/$label/inputs-and-profile.sha256"
printf 'Profile ready: %s\n' "$RESULTS/$label/cpu.pprof"
printf 'Analyze later with: bash %q %q %q\n' "$EXP_DIR/analyze_profile.sh" "$version" "$tag"
