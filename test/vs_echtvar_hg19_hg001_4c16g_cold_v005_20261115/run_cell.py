#!/usr/bin/env python3
"""Run one compressed-output cell; echtvar uses four chromosome workers."""
import argparse
from concurrent.futures import ThreadPoolExecutor, as_completed
import ctypes
import fcntl
import gzip
import hashlib
import io
import json
import os
from pathlib import Path
import select
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
    return {'query': source / 'query/hg001.decomp.vcf.gz', 'archives': source / 'archives',
            'library': root / 'test/17_fields/lib17' if fields == 17 else root / 'gnomad/immkv',
            'fields': REPO / f'test/vs_echtvar_hg19_gnomad_genomes_{fields}_fields/fields{fields}.tsv'}


def raw_command(engine, media, fields, output, *, chromosome=None, input_path=None):
    """Return the actual engine argv; callers may supply a streaming output path."""
    p = paths(media, fields)
    if engine == 'immkv':
        return [str(BUILD_ROOT / 'bin/annotate'), '-in', str(input_path or p['query']),
                '-lib', str(p['library']), '-out', str(output), '-fields', str(p['fields']),
                '-threads', '4', '-batch', '4096']
    if chromosome is None or input_path is None:
        raise ValueError('echtvar requires chromosome and sliced input')
    return [os.environ.get('ECHTVAR_BIN', str(REPO / 'tmp/echtvar/echtvar')), 'anno', '-e',
            str(p['archives'] / f'chr{chromosome}.echtvar.zip'), str(input_path), str(output)]


def timestamp():
    return {'time_ns': time.time_ns(), 'monotonic_ns': time.monotonic_ns()}


def write_new(path, value):
    with path.open('x') as f:
        json.dump(value, f, indent=2, sort_keys=True)
        f.write('\n')


def memfd(name):
    flags = getattr(os, 'MFD_CLOEXEC', 1) | getattr(os, 'MFD_ALLOW_SEALING', 2)
    if hasattr(os, 'memfd_create'):
        return os.memfd_create(name, flags)
    libc = ctypes.CDLL(None, use_errno=True)
    call = libc.memfd_create
    call.argtypes, call.restype = [ctypes.c_char_p, ctypes.c_uint], ctypes.c_int
    fd = call(name.encode(), flags)
    if fd < 0:
        error = ctypes.get_errno()
        raise OSError(error, os.strerror(error))
    return fd


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


class FifoInput(io.RawIOBase):
    """Read a FIFO without blocking forever when its producer dies before open."""
    def __init__(self, fd, done):
        super().__init__()
        self.fd, self.done = fd, done

    def readable(self):
        return True

    def readinto(self, target):
        while True:
            if STOP.is_set():
                raise RuntimeError('Run cancelled while reading compressed output')
            try:
                data = os.read(self.fd, len(target))
                if data:
                    target[:len(data)] = data
                    return len(data)
                if self.done.is_set():
                    return 0
            except BlockingIOError:
                pass
            select.select([self.fd], [], [], 0.1)
            # A FIFO without a writer is immediately readable at EOF.
            if not self.done.is_set():
                self.done.wait(0.01)

    def close(self):
        if not self.closed:
            os.close(self.fd)
        super().close()


def row_digest(stream):
    digest, rows = hashlib.sha256(), 0
    while True:
        line = stream.readline(32 << 20)
        if not line:
            break  # Complete gzip EOF validates CRC/ISIZE for every member.
        if len(line) >= 32 << 20 or not line.endswith(b'\n'):
            raise ValueError('Oversized or unterminated VCF line')
        if not line.startswith(b'#'):
            digest.update(line)
            rows += 1
    return {'rows': rows, 'data_sha256': digest.hexdigest(), 'gzip_crc_eof_ok': True}


def chromosome_task(chrom, args, scratch):
    record = {'chromosome': chrom, 'started': timestamp()}
    if STOP.is_set():
        record.update(status='cancelled', finished=timestamp())
        write_new(args.record / f'chr{chrom}.json', record)
        return record
    fd, link, fifo, reader = None, scratch / f'sl{chrom}.vcf.gz', None, None
    producer_done, comparison = threading.Event(), {}
    try:
        fd = memfd(f'echtvar-chr{chrom}')
        link.symlink_to(f'/proc/{os.getpid()}/fd/{fd}')
        query = paths(args.media, args.fields)['query']
        command = [os.environ.get('BCFTOOLS', '/home/tian/miniforge3/envs/transvar/bin/bcftools'),
                   'view', '-r', str(chrom), '-Oz', '-o', str(link), str(query)]
        record['slice_command'], record['slice_started'] = command, timestamp()
        rc = execute(command, args.record / f'chr{chrom}.slice.log')
        record['slice_exit_code'], record['slice_finished'] = rc, timestamp()
        if rc != 0:
            raise RuntimeError(f'bcftools exited {rc}')
        record['slice_compressed_bytes'] = os.fstat(fd).st_size
        # Finished slices are immutable until their descriptor is closed.
        fcntl.fcntl(fd, getattr(fcntl, 'F_ADD_SEALS', 1033), 15)
        fifo = scratch / f'an{chrom}.vcf.gz'
        if args.verify_reference:
            os.mkfifo(fifo, 0o600)
            input_fd = os.open(fifo, os.O_RDONLY | os.O_NONBLOCK)
            output = str(fifo)

            def compare():
                try:
                    with gzip.open(args.verify_reference / f'an{chrom}.vcf.gz', 'rb') as old:
                        comparison['reference'] = row_digest(old)
                    with io.BufferedReader(FifoInput(input_fd, producer_done), buffer_size=1 << 20) as raw:
                        with gzip.GzipFile(fileobj=raw, mode='rb') as fresh:
                            comparison['fresh'] = row_digest(fresh)
                    comparison['equal'] = comparison['fresh'] == comparison['reference']
                    if not comparison['equal']:
                        raise RuntimeError('Fresh and historical decompressed data rows differ')
                except Exception as exc:
                    comparison['error'] = str(exc)
                    STOP.set()
                    # The fd belongs to this thread even when reference opening fails.
                    try:
                        os.close(input_fd)
                    except OSError:
                        pass

            reader = threading.Thread(target=compare, name=f'verify-chr{chrom}', daemon=True)
            reader.start()
        else:
            fifo.symlink_to('/dev/null')
            output = str(fifo)
        command = raw_command('echtvar', args.media, args.fields, output, chromosome=chrom, input_path=link)
        record['annotation_command'], record['annotation_started'] = command, timestamp()
        rc = execute(command, args.record / f'chr{chrom}.anno.log')
        record['annotation_exit_code'], record['annotation_finished'] = rc, timestamp()
        producer_done.set()
        if reader:
            reader.join()
            record['comparison'] = comparison
        if rc != 0:
            raise RuntimeError(f'echtvar exited {rc}')
        if args.verify_reference and not comparison.get('equal'):
            raise RuntimeError(comparison.get('error', 'Output verification incomplete'))
        record['status'] = 'ok'
    except Exception as exc:
        record.update(status='failed', error=str(exc))
        STOP.set()
    finally:
        producer_done.set()
        if reader and reader.is_alive():
            reader.join()
        if fd is not None:
            os.close(fd)
        link.unlink(missing_ok=True)
        if fifo:
            fifo.unlink(missing_ok=True)
        record['finished'] = timestamp()
        write_new(args.record / f'chr{chrom}.json', record)
    return record


def handle_signal(number, frame):
    global SIGNAL
    SIGNAL = number
    STOP.set()


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--engine', choices=('immkv', 'echtvar'), required=True)
    parser.add_argument('--media', choices=('nvme', 'hdd'), required=True)
    parser.add_argument('--fields', choices=(17, 540, 568), type=int, required=True)
    parser.add_argument('--record', type=Path, required=True)
    parser.add_argument('--scratch', type=Path)
    parser.add_argument('--output', type=Path, help='immkv output path, including a caller-owned FIFO.gz')
    parser.add_argument('--verify-reference', type=Path, help='echtvar historical anN.vcf.gz directory')
    parser.add_argument('--expected-rows', type=int, default=3933714)
    args = parser.parse_args()
    if args.output and args.engine != 'immkv':
        parser.error('--output is supported for immkv only')
    if args.verify_reference and args.engine != 'echtvar':
        parser.error('--verify-reference is supported for echtvar only')
    args.record = args.record.resolve()
    if not args.record.is_relative_to(EXP_DIR / 'results'):
        parser.error('--record must remain inside this experiment results tree')
    if args.output and not args.output.resolve().is_relative_to(BUILD_ROOT.resolve()):
        parser.error('--output must remain inside the experiment scratch tree')
    args.record.mkdir(parents=True, exist_ok=True)
    scratch = args.scratch or BUILD_ROOT / 'runs' / args.record.name
    if not scratch.resolve().is_relative_to(BUILD_ROOT.resolve()):
        parser.error('--scratch must remain inside the experiment scratch tree')
    scratch.mkdir(parents=True, exist_ok=False)
    for sig in (signal.SIGINT, signal.SIGTERM):
        signal.signal(sig, handle_signal)
    summary = {'engine': args.engine, 'media': args.media, 'fields': args.fields,
               'started': timestamp(), 'compressed_output': True, 'scratch': str(scratch),
               'source_commit': 'd2eac357f3fc126755b43ff8b44a64dc55bc0ec2',
               'verify_reference': str(args.verify_reference.resolve()) if args.verify_reference else None,
               'reference_directory': str(args.verify_reference) if args.verify_reference else None,
               'timing_sample': not bool(args.verify_reference or args.output)}
    rc = 1
    try:
        if args.engine == 'immkv':
            output = args.output or scratch / 'output.vcf.gz'
            if not args.output:
                output.symlink_to('/dev/null')
            try:
                command = raw_command('immkv', args.media, args.fields, output)
                summary['command'] = command
                rc = execute(command, args.record / 'immkv.log')
                rc = rc if rc is not None and rc >= 0 else 128 - rc if rc is not None else 1
            finally:
                if not args.output:
                    output.unlink(missing_ok=True)
        else:
            with ThreadPoolExecutor(max_workers=4) as pool:
                futures = [pool.submit(chromosome_task, chrom, args, scratch) for chrom in range(1, 23)]
                records = [f.result() for f in as_completed(futures)]
            summary['chromosomes'] = sorted(records, key=lambda x: x['chromosome'])
            summary['workers'] = 4
            summary['worker_note'] = 'Four chromosome subprocess tasks on CPUs 0-3; each echtvar process may create additional htslib threads.'
            rc = 0 if all(x['status'] == 'ok' for x in records) else 1
            if args.verify_reference:
                total = sum(x.get('comparison', {}).get('fresh', {}).get('rows', 0) for x in records)
                summary['fresh_rows'], summary['expected_rows'] = total, args.expected_rows
                if total != args.expected_rows:
                    summary['error'] = 'Combined fresh data-row count differs from expected query rows'
                    rc = 1
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
