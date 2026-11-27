#!/usr/bin/env python3
"""Run one compressed-output immkv cell."""
import argparse
import json
import os
from pathlib import Path
import signal
import subprocess
import threading
import time

EXP_DIR = Path(__file__).resolve().parent
REPO = EXP_DIR.parent.parent
BUILD_ROOT = REPO / 'tmp' / EXP_DIR.name
STOP = threading.Event()
SIGNAL = 0


def paths(media, fields):
    root = Path(os.environ.get('HDD_ROOT', '/mnt/data6t/tian/immkv_test_hdd') if media == 'hdd'
                else os.environ.get('NVME_ROOT', '/home/tian/immkv_test_nvme'))
    source = root / 'test' / f'{fields}_fields'
    return {'query': source / 'query/hg001.decomp.vcf.gz',
            'library': root / 'test/17_fields/lib17' if fields == 17 else root / 'gnomad/immkv',
            'fields': REPO / f'test/vs_echtvar_hg19_gnomad_genomes_{fields}_fields/fields{fields}.tsv'}


def raw_command(media, fields, output, *, input_path=None):
    p = paths(media, fields)
    return [str(BUILD_ROOT / 'bin/annotate'), '-in', str(input_path or p['query']),
            '-lib', str(p['library']), '-out', str(output), '-fields', str(p['fields']),
            '-threads', '4', '-batch', '4096']


def timestamp():
    return {'time_ns': time.time_ns(), 'monotonic_ns': time.monotonic_ns()}


def write_new(path, value):
    with path.open('x') as f:
        json.dump(value, f, indent=2, sort_keys=True)
        f.write('\n')


def execute(command, log):
    if STOP.is_set():
        return None
    with log.open('xb') as stream:
        proc = subprocess.Popen(command, stdout=stream, stderr=subprocess.STDOUT, start_new_session=True)
        while proc.poll() is None:
            if STOP.wait(0.1):
                try:
                    os.killpg(proc.pid, signal.SIGTERM)
                    proc.wait(timeout=2)
                except subprocess.TimeoutExpired:
                    os.killpg(proc.pid, signal.SIGKILL)
                except ProcessLookupError:
                    pass
                break
        return proc.wait()


def handle_signal(number, frame):
    global SIGNAL
    SIGNAL = number
    STOP.set()


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--engine', choices=('immkv',), default='immkv')
    parser.add_argument('--media', choices=('nvme', 'hdd'), required=True)
    parser.add_argument('--fields', choices=(17, 540, 568), type=int, required=True)
    parser.add_argument('--record', type=Path, required=True)
    parser.add_argument('--scratch', type=Path)
    parser.add_argument('--output', type=Path, help='immkv output path')
    args = parser.parse_args()
    args.record = args.record.resolve()
    if not args.record.is_relative_to(EXP_DIR / 'results'):
        parser.error('--record must remain inside this experiment results tree')
    args.record.mkdir(parents=True, exist_ok=True)
    scratch = args.scratch or BUILD_ROOT / 'runs' / args.record.name
    if not scratch.resolve().is_relative_to(BUILD_ROOT.resolve()):
        parser.error('--scratch must remain inside the experiment scratch tree')
    scratch.mkdir(parents=True, exist_ok=False)
    for sig in (signal.SIGINT, signal.SIGTERM):
        signal.signal(sig, handle_signal)
    source_facts = json.loads((EXP_DIR / 'results/preparation/source-facts.json').read_text())
    summary = {'engine': 'immkv', 'media': args.media, 'fields': args.fields,
               'started': timestamp(), 'compressed_output': True, 'scratch': str(scratch),
               'base_commit': source_facts['base_commit'],
               'sources_manifest_sha256': source_facts['sources_manifest_sha256'],
               'timing_sample': not bool(args.output)}
    rc = 1
    try:
        output = args.output or scratch / 'output.vcf.gz'
        if not args.output:
            output.symlink_to('/dev/null')
        try:
            command = raw_command(args.media, args.fields, output)
            summary['command'] = command
            rc = execute(command, args.record / 'immkv.log')
            rc = rc if rc is not None and rc >= 0 else 128 - rc if rc is not None else 1
        finally:
            if not args.output:
                output.unlink(missing_ok=True)
        if SIGNAL:
            rc = 128 + SIGNAL
        summary['exit_code'], summary['status'] = rc, 'ok' if rc == 0 else 'failed'
    except Exception as exc:
        summary.update(status='failed', error=str(exc), exit_code=1)
        rc = 1
    finally:
        summary['finished'] = timestamp()
        write_new(args.record / 'cell.json', summary)
    return rc


if __name__ == '__main__':
    raise SystemExit(main())
