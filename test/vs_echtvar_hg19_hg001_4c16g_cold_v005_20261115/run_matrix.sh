#!/usr/bin/env bash
# From-zero six-cell matrix: full-output correctness then three cold repeats.
set -euo pipefail
source "$(dirname -- "${BASH_SOURCE[0]}")/common.sh"
require_binary
MEDIA=${MEDIA:-'nvme hdd'}
FIELDS=${FIELDS:-'17 540 568'}
RUN_TAG=${RUN_TAG:-cold-v005}
VERIFY_TAG=${VERIFY_TAG:-cold-v005}
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
if proof['status'] != 'passed' or proof['media'] != media or int(proof['fields']) != int(fields):
    raise SystemExit(f'Full-output proof does not match this cell: {p}')
comparison = proof['comparison']
if (comparison['rows'] != 3933714 or comparison['field_comparisons'] != 3933714 * int(fields)
        or comparison['value_mismatches'] != 0 or comparison['hit_disagreements'] != 0):
    raise SystemExit(f'Full-output comparison totals are invalid: {p}')
e = root / f'verify-echtvar-{media}-{fields}-{tag}'
status = json.loads((e / 'status.json').read_text())['status']
if status == 'ok':
    tools.validate_success(e)
    fresh = json.loads((e / 'cell.json').read_text())
    if fresh['engine'] != 'echtvar' or fresh['media'] != media or fresh['fields'] != int(fields) or fresh.get('fresh_rows') != 3933714:
        raise SystemExit(f'Fresh echtvar equivalence evidence is invalid: {e}')
elif status == 'oom' and fields == '568':
    # verify_output.py checks explicit OOM evidence and historical provenance;
    # this is correctness evidence only, never a successful performance sample.
    if (not proof.get('reference_kind') or
            Path(proof.get('reference_oom_evidence', '')).resolve() != (e / 'status.json').resolve()):
        raise SystemExit(f'Historical fallback lacks explicit OOM provenance: {p}')
else:
    raise SystemExit(f'Fresh echtvar proof did not succeed: {e} ({status})')
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
            case "$round" in 2) engines='echtvar immkv';; *) engines='immkv echtvar';; esac
            for engine in $engines; do
                label="timing-$media-$fields-$engine-r$round-$RUN_TAG"
                record="$RESULTS/$label"
                if [[ -e "$record" ]]; then
                    # Only complete, validated successful records are resumable.
                    # OOM/failure/incomplete records remain preserved and require
                    # a new RUN_TAG; no old attempt is silently skipped.
                    python3 -I "$EXP_DIR/env_tools.py" validate-success "$record"
                    python3 -I - "$record/cell.json" "$engine" "$media" "$fields" <<'PY'
import json, sys
from pathlib import Path
cell = json.loads(Path(sys.argv[1]).read_text())
if (cell['engine'], cell['media'], cell['fields']) != (sys.argv[2], sys.argv[3], int(sys.argv[4])) or not cell.get('timing_sample'):
    raise SystemExit('Existing successful record belongs to a different invocation')
PY
                    printf 'Preserve validated success: %s\n' "$label"
                    continue
                fi
                printf '%s start %s\n' "$(date -Is)" "$label"
                if limited "$label" "$media" python3 -I "$EXP_DIR/run_cell.py" \
                    --engine "$engine" --media "$media" --fields "$fields" --record "$record"; then
                    :
                else
                    # A genuine OOM is an observed constrained failure. All
                    # three scheduled attempts still run; no skip-after-OOM.
                    python3 -I - "$record/status.json" <<'PY'
import json, sys
from pathlib import Path
status = json.loads(Path(sys.argv[1]).read_text())['status']
if status != 'oom':
    raise SystemExit(f'Invocation failed ({status}); preserve evidence and stop: {sys.argv[1]}')
PY
                fi
            done
        done
    done
done
