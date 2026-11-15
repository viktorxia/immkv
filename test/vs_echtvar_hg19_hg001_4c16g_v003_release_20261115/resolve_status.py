#!/usr/bin/env python3
"""Resolve killed scopes from their unit journal while retaining raw status."""
import argparse
import json
from pathlib import Path
import subprocess
import time


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('record', type=Path)
    args = parser.parse_args()
    p = args.record
    status = json.loads((p / 'status.json').read_text())['status']
    resolved = p / 'status-resolution.json'
    if resolved.exists():
        print(json.loads(resolved.read_text())['status'])
        return
    if status in ('ok', 'oom', 'skipped_after_oom'):
        print(status)
        return
    unit = (p / 'scope-unit').read_text().strip()
    journal = ''
    for _ in range(40):
        result = subprocess.run(['journalctl', '--user', '-u', unit, '--no-pager'], capture_output=True, text=True)
        journal = result.stdout
        if "Failed with result 'oom-kill'" in journal or 'killed by the OOM killer' in journal:
            break
        time.sleep(0.5)
    log = p / 'systemd-final.log'
    with log.open('x') as handle:
        handle.write(journal)
    if "Failed with result 'oom-kill'" in journal or 'killed by the OOM killer' in journal:
        with resolved.open('x') as handle:
            json.dump({'status': 'oom', 'original_status_source': 'status.json',
                       'evidence': ['systemd-final.log'],
                       'reason': 'Unit journal explicitly records the OOM killer; raw observer status is retained.'}, handle, indent=2)
            handle.write('\n')
        print('oom')
    else:
        print(status)


if __name__ == '__main__':
    main()
