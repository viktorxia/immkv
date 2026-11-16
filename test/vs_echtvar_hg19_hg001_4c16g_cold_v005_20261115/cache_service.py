#!/usr/bin/env python3
"""Finite, authenticated cache-drop service for this cold-cache experiment only.

Run with pkexec /usr/bin/python3 -I. The only privileged operation exposed
is sync followed by writing 3 to /proc/sys/vm/drop_caches. No client-supplied
paths or shell commands are accepted. The owner can stop the service.
"""
import argparse
import datetime
import json
import os
from pathlib import Path
import socket
import struct
import time


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--socket', required=True)
    parser.add_argument('--uid', required=True, type=int)
    args = parser.parse_args()
    if os.geteuid() != 0:
        raise SystemExit('root authentication is required')
    path = Path(args.socket)
    if path.exists():
        raise SystemExit('refusing to replace an existing socket')
    server = socket.socket(socket.AF_UNIX)
    server.bind(str(path))
    os.chown(path, args.uid, -1)
    os.chmod(path, 0o600)
    server.listen(1)
    server.settimeout(12 * 3600)
    deadline = time.monotonic() + 12 * 3600
    try:
        while time.monotonic() < deadline:
            try:
                connection, _ = server.accept()
            except socket.timeout:
                break
            with connection:
                connection.settimeout(5)
                _, uid, _ = struct.unpack('3i', connection.getsockopt(socket.SOL_SOCKET, socket.SO_PEERCRED, struct.calcsize('3i')))
                if uid != args.uid:
                    continue
                command = connection.recv(128)
                if command == b'stop\n':
                    connection.sendall(b'{"stopped": true}\n')
                    break
                if command != b'drop\n':
                    connection.sendall(b'{"error": "unsupported command"}\n')
                    continue
                report = {'command': 'os.sync(); write 3 to /proc/sys/vm/drop_caches',
                          'started_utc': datetime.datetime.now(datetime.timezone.utc).isoformat(),
                          'peer_uid': uid, 'effective_uid': os.geteuid(),
                          'meminfo_before': Path('/proc/meminfo').read_text()}
                started = time.monotonic()
                try:
                    os.sync()
                    with open('/proc/sys/vm/drop_caches', 'w') as stream:
                        stream.write('3\n')
                    report['exit_status'] = 0
                except OSError as error:
                    report.update(exit_status=1, error=str(error))
                report.update(elapsed_seconds=time.monotonic() - started,
                              finished_utc=datetime.datetime.now(datetime.timezone.utc).isoformat(),
                              meminfo_after=Path('/proc/meminfo').read_text())
                connection.sendall(json.dumps(report).encode() + b'\n')
    finally:
        server.close()
        path.unlink()


if __name__ == '__main__':
    main()
