#!/usr/bin/env bash
# From-zero six-cell immkv matrix: full-output correctness then three cold repeats.
set -euo pipefail
source "$(dirname -- "${BASH_SOURCE[0]}")/common.sh"
require_binary
MEDIA=${MEDIA:-'nvme hdd'}
FIELDS=${FIELDS:-'17 540 568'}
RUN_TAG=${RUN_TAG:-cold-postfix}
VERIFY_TAG=${VERIFY_TAG:-cold-postfix}
valid_tag "$RUN_TAG"
valid_tag "$VERIFY_TAG"
mkdir -p "$RESULTS" "$BUILD_ROOT"
exec {matrix_lock}> "$BUILD_ROOT/matrix.lock"
flock -n "$matrix_lock" || fail 'Another matrix driver is running'

check_proof() {
    python3 -I - "$RESULTS" "$1" "$2" "$VERIFY_TAG" "$EXP_DIR/env_tools.py" <<'PY'
import importlib.util, json, sys
from pathlib import Path
root, media, fields, tag, source = Path(sys.argv[1]), *sys.argv[2:]
spec = importlib.util.spec_from_file_location('cold_environment', source)
tools = importlib.util.module_from_spec(spec)
spec.loader.exec_module(tools)
p = root / f'verify-{media}-{fields}-{tag}'
tools.validate_success(p)
proof = json.loads((p / 'verification.json').read_text())
if proof['status'] != 'passed' or not proof.get('v005_digest_equal') or proof['media'] != media or int(proof['fields']) != int(fields):
    raise SystemExit(f'Full-output proof does not match this cell: {p}')
comparison = proof['comparison']
if (comparison['rows'] != 3933714 or comparison['field_comparisons'] != 3933714 * int(fields)
        or comparison['value_mismatches'] != 0 or comparison['hit_disagreements'] != 0):
    raise SystemExit(f'Full-output comparison totals are invalid: {p}')
PY
}

for media in $MEDIA; do
    case "$media" in nvme|hdd) ;; *) fail "Unknown media: $media";; esac
    for fields in $FIELDS; do
        case "$fields" in 17|540|568) ;; *) fail "Unknown field set: $fields";; esac
        proof="$RESULTS/verify-$media-$fields-$VERIFY_TAG"
        if [[ ! -d "$proof" ]]; then
            bash "$EXP_DIR/verify.sh" "$media" "$fields" "$VERIFY_TAG"
        fi
        check_proof "$media" "$fields"
        for round in 1 2 3; do
            label="timing-$media-$fields-immkv-r$round-$RUN_TAG"
            record="$RESULTS/$label"
            if [[ -e "$record" ]]; then
                python3 -I "$EXP_DIR/env_tools.py" validate-success "$record"
                printf 'Preserve validated success: %s\n' "$label"
                continue
            fi
            printf '%s start %s\n' "$(date -Is)" "$label"
            limited "$label" "$media" python3 -I "$EXP_DIR/run_cell.py" \
                --engine immkv --media "$media" --fields "$fields" --record "$record"
        done
    done
done
