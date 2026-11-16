#!/usr/bin/env bash
# Wait for the already-requested desktop authentication, then execute once.
set -euo pipefail
source "$(dirname -- "${BASH_SOURCE[0]}")/common.sh"
label="orchestration-$(date +%s%N)"
record="$RESULTS/$label"
mkdir "$record"
exec > >(tee "$record/stdout") 2> >(tee "$record/stderr" >&2)
printf '%s\n' "$(date -Is)" > "$record/started"
printf '%s\n' 'Waiting up to 30 minutes for cache-service desktop authentication.'
python3 -I - "$CACHE_SOCKET" "$record" <<'PY'
import datetime, json, pathlib, sys, time
socket_path, record = pathlib.Path(sys.argv[1]), pathlib.Path(sys.argv[2])
deadline = time.monotonic() + 1800
while not socket_path.is_socket():
    if time.monotonic() >= deadline:
        with (record / 'authentication-timeout.json').open('x') as stream:
            json.dump({'status': 'blocked', 'reason': 'Cache-service authentication not completed; no performance samples started', 'utc': datetime.datetime.now(datetime.timezone.utc).isoformat()}, stream, indent=2)
            stream.write('\n')
        raise SystemExit('Authentication wait expired; no timed performance invocation started.')
    time.sleep(1)
with (record / 'authentication-ready.json').open('x') as stream:
    json.dump({'status': 'ready', 'socket': str(socket_path), 'utc': datetime.datetime.now(datetime.timezone.utc).isoformat()}, stream, indent=2)
    stream.write('\n')
PY
printf '%s\n' "$(date -Is) cache service ready; starting full matrix."
set +e
bash "$EXP_DIR/run_matrix.sh"
rc=$?
set -e
printf '%s\n' "$rc" > "$record/matrix-exit-status"
python3 -I "$EXP_DIR/summarize.py" --name "summary-$label"
printf '%s\n' "$(date -Is)" > "$record/finished"
# Shut down this finite privileged helper after the matrix, including failure.
python3 -I - "$CACHE_SOCKET" <<'PY'
import socket, sys
try:
    with socket.socket(socket.AF_UNIX) as connection:
        connection.settimeout(10)
        connection.connect(sys.argv[1])
        connection.sendall(b'stop\n')
        print(connection.recv(4096).decode().strip())
except OSError as error:
    print('Cache-service shutdown unavailable:', error, file=sys.stderr)
PY
exit "$rc"
