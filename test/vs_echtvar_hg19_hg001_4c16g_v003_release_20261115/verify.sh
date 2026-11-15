#!/bin/bash
# Full-output correctness; elapsed time includes decoding and comparison, not lookup timing.
set -euo pipefail
source "$(dirname "${BASH_SOURCE[0]}")/common.sh"
[[ $# -ge 2 ]] || fail 'Usage: verify.sh MEDIA FIELDS [TAG] [--annotate PATH] [--verifier PATH]'
media=$1 fields=$2
shift 2
case "$media" in hdd|nvme) ;; *) fail "Unknown media: $media";; esac
case "$fields" in 17|540|568) ;; *) fail "Unknown field set: $fields";; esac
tag="verify-$media-$fields-release"
if [[ $# -gt 0 && $1 != --* ]]; then tag=$1; shift; fi
valid_tag "$tag"
require_binary
limited "$tag" "$media" python3 -I "$EXP_DIR/verify_output.py" \
    --media "$media" --fields "$fields" --result "$RESULTS/$tag" \
    --annotate "$ANNOTATE" --verifier "$BUILD_ROOT/bin/verifyvalues" "$@"
