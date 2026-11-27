#!/usr/bin/env bash
# Full per-row/per-field comparison against historical references plus v0.0.5 immkv digest equality.
set -euo pipefail
source "$(dirname "${BASH_SOURCE[0]}")/common.sh"
[[ $# -ge 2 ]] || fail 'Usage: verify.sh MEDIA FIELDS [TAG]'
media=$1 fields=$2
shift 2
case "$media" in hdd|nvme) ;; *) fail "Unknown media: $media";; esac
case "$fields" in 17|540|568) ;; *) fail "Unknown field set: $fields";; esac
tag=${VERIFY_TAG:-cold-postfix}
if [[ $# -gt 0 && $1 != --* ]]; then tag=$1; shift; fi
valid_tag "$tag"
if [[ "$fields" == 568 ]]; then
    reference="$HDD_ROOT/test/4c16g_gz_v003/568_fields/echtvar_ref.wk"
else
    reference="$(media_root "$media")/test/4c16g_gz_v003/${fields}_fields/echtvar.pass2.wk"
fi
require_binary
comparison_tag="verify-$media-$fields-$tag"
valid_tag "$comparison_tag"
[[ ! -e "$RESULTS/$comparison_tag" ]] || fail 'Verification result already exists; choose a new tag and preserve old evidence'
limited "$comparison_tag" "$media" python3 -I "$EXP_DIR/verify_output.py" \
    --media "$media" --fields "$fields" --result "$RESULTS/$comparison_tag" \
    --annotate "$ANNOTATE" --verifier "$VERIFIER" --reference "$reference"
