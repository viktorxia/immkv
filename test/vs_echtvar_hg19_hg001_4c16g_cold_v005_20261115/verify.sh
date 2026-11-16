#!/usr/bin/env bash
# Fresh echtvar digest equivalence first, then full per-row/per-field comparison.
set -euo pipefail
source "$(dirname "${BASH_SOURCE[0]}")/common.sh"
[[ $# -ge 2 ]] || fail 'Usage: verify.sh MEDIA FIELDS [TAG] [--reference DIR] [--annotate PATH] [--verifier PATH]'
media=$1 fields=$2
shift 2
case "$media" in hdd|nvme) ;; *) fail "Unknown media: $media";; esac
case "$fields" in 17|540|568) ;; *) fail "Unknown field set: $fields";; esac
tag=${VERIFY_TAG:-cold-v005}
if [[ $# -gt 0 && $1 != --* ]]; then tag=$1; shift; fi
valid_tag "$tag"
if [[ "$fields" == 568 ]]; then
    reference="$HDD_ROOT/test/4c16g_gz_v003/568_fields/echtvar_ref.wk"
else
    reference="$(media_root "$media")/test/4c16g_gz_v003/${fields}_fields/echtvar.pass2.wk"
fi
annotate="$BUILD_ROOT/bin/annotate"
verifier="$BUILD_ROOT/bin/verifyvalues"
while [[ $# -gt 0 ]]; do
    [[ $# -ge 2 ]] || fail "Missing value for $1"
    case "$1" in
        --reference) reference=$2;;
        --annotate) annotate=$2;;
        --verifier) verifier=$2;;
        *) fail "Unknown argument: $1";;
    esac
    shift 2
done
require_binary
fresh_tag="verify-echtvar-$media-$fields-$tag"
comparison_tag="verify-$media-$fields-$tag"
valid_tag "$fresh_tag"
valid_tag "$comparison_tag"
[[ ! -e "$RESULTS/$fresh_tag" && ! -e "$RESULTS/$comparison_tag" ]] || fail 'Verification result already exists; choose a new tag and preserve old evidence'
if limited "$fresh_tag" "$media" python3 -I "$EXP_DIR/run_cell.py" \
        --engine echtvar --media "$media" --fields "$fields" --record "$RESULTS/$fresh_tag" \
        --verify-reference "$reference" --expected-rows 3933714; then
    proof=(--reference-evidence "$RESULTS/$fresh_tag/cell.json")
else
    # A cap-limited full-field OOM permits historical correctness only; every
    # other producer, integrity, timeout, or infrastructure failure stops here.
    python3 -I - "$RESULTS/$fresh_tag/status.json" <<'PY'
import json, sys
from pathlib import Path
path = Path(sys.argv[1])
if not path.is_file() or json.loads(path.read_text()).get('status') != 'oom':
    raise SystemExit('Fresh echtvar verification failed without recorded OOM; comparison not started')
print('Fresh echtvar OOM recorded; historical-only correctness fallback, no fresh digest equivalence claim')
PY
    proof=(--reference-oom-evidence "$RESULTS/$fresh_tag/status.json")
fi
limited "$comparison_tag" "$media" python3 -I "$EXP_DIR/verify_output.py" \
    --media "$media" --fields "$fields" --result "$RESULTS/$comparison_tag" \
    --annotate "$annotate" --verifier "$verifier" --reference "$reference" "${proof[@]}"
