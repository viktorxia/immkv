#!/bin/bash
# Sequential continuation after the first NVMe/17 timing cell.
# Every invocation preserves prior evidence and runs in its own 4C16G scope.
set -euo pipefail
source "$(dirname "$0")/common.sh"
for round in 1 2 3; do
    for engine in immkv echtvar; do
        [[ -f "$RESULTS/timing-nvme-17-$engine-r$round-release/finished" ]] || fail 'Complete the initial NVMe/17 timing cell first'
    done
done
for media in ${MEDIA:-nvme hdd}; do
    for fields in ${FIELDS:-17 540 568}; do
        [[ "$media" == nvme && "$fields" == 17 ]] && continue
        proof="$RESULTS/verify-$media-$fields-release"
        if [[ ! -f "$proof/finished" ]]; then
            bash "$EXP_DIR/verify.sh" "$media" "$fields"
        fi
        label="verify-echtvar-$media-$fields-release"
        if [[ "$fields" == 568 ]]; then
            reference="$HDD_ROOT/test/4c16g_gz_v003/568_fields/echtvar_ref.wk"
        else
            reference="$(media_root "$media")/test/4c16g_gz_v003/${fields}_fields/echtvar.pass2.wk"
        fi
        if [[ -f "$RESULTS/$label/status.json" ]]; then
            python3 -I - "$RESULTS/$label" <<'PY'
import json, sys
from pathlib import Path
p = Path(sys.argv[1])
source = p / 'status-resolution.json'
status = json.loads((source if source.exists() else p / 'status.json').read_text())['status']
if status not in ('ok', 'oom'):
    raise SystemExit(f'Existing correctness attempt requires investigation: {p}')
PY
        elif limited "$label" "$media" python3 -I "$EXP_DIR/run_cell.py" \
            --engine echtvar --media "$media" --fields "$fields" --record "$RESULTS/$label" \
            --verify-reference "$reference"; then
            :
        else
            python3 -I - "$RESULTS/$label/status.json" <<'PY'
import json, sys
from pathlib import Path
record = Path(sys.argv[1])
resolution = record.parent / 'status-resolution.json'
if json.loads((resolution if resolution.exists() else record).read_text())["status"] != "oom":
    raise SystemExit(f"Correctness attempt failed without OOM; inspect {record}")
PY
        fi
        MEDIA="$media" FIELDS="$fields" RUN_TAG=release bash "$EXP_DIR/run_matrix.sh"
    done
done
