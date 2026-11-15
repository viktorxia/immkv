#!/bin/bash
# Timing only: require full-output correctness evidence before each cell.
set -euo pipefail
source "$(dirname "$0")/common.sh"
require_binary
MEDIA=${MEDIA:-'nvme hdd'}
FIELDS=${FIELDS:-'17 540 568'}
RUN_TAG=${RUN_TAG:-release}
VERIFY_TAG=${VERIFY_TAG:-release}
valid_tag "$RUN_TAG"
valid_tag "$VERIFY_TAG"
mkdir -p "$RESULTS" "$BUILD_ROOT"
exec {matrix_lock}> "$BUILD_ROOT/matrix.lock"
flock -n "$matrix_lock" || fail 'Another matrix driver is running'

check_proof() {
    python3 -I - "$RESULTS" "$1" "$2" "$VERIFY_TAG" <<'PY'
import json, sys
from pathlib import Path
root, media, fields, tag = Path(sys.argv[1]), *sys.argv[2:]
def read(path):
    if not path.is_file():
        raise SystemExit(f'Missing correctness evidence: {path}')
    return json.loads(path.read_text())
p = root / f'verify-{media}-{fields}-{tag}'
if read(p / 'status.json')['status'] != 'ok' or read(p / 'verification.json')['status'] != 'passed':
    raise SystemExit(f'Full-output immkv correctness did not pass: {p}')
p = root / f'verify-echtvar-{media}-{fields}-{tag}'
status = read(p / 'status.json')['status']
if (p / 'status-resolution.json').exists():
    status = read(p / 'status-resolution.json')['status']
if status == 'oom':
    print(f'{media}/{fields}: fresh echtvar correctness attempt OOM; timing will measure the constrained failure once')
elif status != 'ok' or read(p / 'cell.json')['status'] != 'ok':
    raise SystemExit(f'Fresh echtvar correctness did not pass: {p}')
PY
}

for media in $MEDIA; do
    case "$media" in nvme|hdd) ;; *) fail "Unknown media: $media";; esac
    for fields in $FIELDS; do
        case "$fields" in 17|540|568) ;; *) fail "Unknown field set: $fields";; esac
        check_proof "$media" "$fields"
        for round in 1 2 3; do
            case "$round" in 2) engines='echtvar immkv';; *) engines='immkv echtvar';; esac
            for engine in $engines; do
                label="timing-$media-$fields-$engine-r$round-$RUN_TAG"
                if [[ -f "$RESULTS/$label/status.json" ]]; then
                    python3 -I - "$RESULTS/$label/status.json" <<'PY'
import json, sys
from pathlib import Path
def effective(path):
    path = Path(path)
    resolution = path.parent / 'status-resolution.json'
    return json.loads((resolution if resolution.exists() else path).read_text())['status']
s = effective(sys.argv[1])
if s not in ('ok', 'oom', 'skipped_after_oom'):
    raise SystemExit(f'Previous failed invocation requires investigation; preserve evidence: {sys.argv[1]}')
PY
                    printf 'Preserve existing result: %s\n' "$label"
                    continue
                fi
                # OOM retries are omitted; do not turn omitted work into samples.
                if [[ "$round" != 1 && "$engine" == echtvar ]] && \
                    python3 -I - "$RESULTS/timing-$media-$fields-echtvar-r1-$RUN_TAG/status.json" <<'PY'
import json, sys
from pathlib import Path
p = Path(sys.argv[1])
resolution = p.parent / 'status-resolution.json'
raise SystemExit(0 if p.exists() and json.loads((resolution if resolution.exists() else p).read_text())['status'] == 'oom' else 1)
PY
                then
                    mkdir "$RESULTS/$label" || fail "Result exists: $label"
                    python3 -I - "$RESULTS/$label" "timing-$media-$fields-echtvar-r1-$RUN_TAG" <<'PY'
import json, sys
from pathlib import Path
with (Path(sys.argv[1]) / 'status.json').open('x') as f:
    json.dump({'status': 'skipped_after_oom', 'source': sys.argv[2], 'timing_sample': False}, f, indent=2)
    f.write('\n')
PY
                    continue
                fi
                printf '%s start %s\n' "$(date -Is)" "$label"
                if limited "$label" "$media" python3 -I "$EXP_DIR/run_cell.py" \
                    --engine "$engine" --media "$media" --fields "$fields" --record "$RESULTS/$label"; then
                    :
                else
                    python3 -I - "$RESULTS/$label/status.json" <<'PY'
import json, sys
from pathlib import Path
p = Path(sys.argv[1])
resolution = p.parent / 'status-resolution.json'
status = json.loads((resolution if resolution.exists() else p).read_text())['status']
if status != 'oom':
    raise SystemExit(f'Cell failed ({status}); stopping matrix; evidence: {sys.argv[1]}')
PY
                fi
            done
        done
    done
done
