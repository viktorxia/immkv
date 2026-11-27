#!/usr/bin/env python3
"""Record cold-cache scope constraints, persistent cgroup counters and outcomes."""
import argparse
import json
import os
from pathlib import Path
import select
import socket
import subprocess
import sys
import time

EXP_DIR = Path(__file__).resolve().parent
REPO = EXP_DIR.parents[1]
COUNTERS = ('memory.max', 'memory.swap.max', 'memory.peak', 'memory.current',
            'memory.events', 'memory.events.local', 'cpu.stat', 'io.stat')
REQUIRED = ('memory.max', 'memory.swap.max', 'memory.events', 'memory.peak', 'memory.current', 'cpu.stat')
OOM_KEYS = ('oom_kill', 'oom_group_kill')


def write_json(path, value, *, replace=False):
    text = json.dumps(value, indent=2, sort_keys=True) + '\n'
    if replace:
        partial = path.with_name(path.name + '.partial')
        with partial.open('x') as stream:
            stream.write(text)
        partial.replace(path)
    else:
        with path.open('x') as stream:
            stream.write(text)


def read_json(path):
    return json.loads(path.read_text())


def local_record(path):
    path = path.resolve()
    if not path.is_relative_to(EXP_DIR / 'results') or not path.is_dir():
        raise ValueError(f'Record must already exist inside {EXP_DIR / "results"}: {path}')
    return path


def cgroup_path():
    for line in Path('/proc/self/cgroup').read_text().splitlines():
        if line.startswith('0::'):
            return Path('/sys/fs/cgroup') / line[3:].lstrip('/')
    raise RuntimeError('Unified cgroup v2 is required')


def read_fd(fd):
    os.lseek(fd, 0, os.SEEK_SET)
    chunks = []
    while True:
        chunk = os.read(fd, 65536)
        if not chunk:
            return b''.join(chunks).decode().strip()
        chunks.append(chunk)


def counters(cg=None, fds=None):
    facts = {'timestamp_ns': time.time_ns(), 'monotonic_ns': time.monotonic_ns()}
    errors = {}
    for name in COUNTERS:
        try:
            facts[name] = read_fd(fds[name]) if fds is not None else (cg / name).read_text().strip()
        except (OSError, KeyError) as error:
            errors[name] = str(error)
    if errors:
        facts['read_errors'] = errors
    return facts


def values(text):
    return {line.split()[0]: int(line.split()[1]) for line in text.splitlines() if len(line.split()) == 2}


def io_values(text):
    result = {}
    for line in text.splitlines():
        device, *items = line.split()
        result[device] = {key: int(value) for key, value in (item.split('=', 1) for item in items)}
    return result


def deltas(before, after):
    result = {}
    for name in ('memory.events', 'memory.events.local', 'cpu.stat'):
        if name in before and name in after:
            old, fresh = values(before[name]), values(after[name])
            result[name] = {key: value - old.get(key, 0) for key, value in fresh.items()}
    if 'io.stat' in before and 'io.stat' in after:
        old, fresh = io_values(before['io.stat']), io_values(after['io.stat'])
        result['io.stat'] = {device: {key: value - old.get(device, {}).get(key, 0)
                                     for key, value in entries.items()} for device, entries in fresh.items()}
        result['io_totals'] = {key: sum(entries.get(key, 0) for entries in result['io.stat'].values())
                               for key in ('rbytes', 'wbytes', 'rios', 'wios', 'dbytes', 'dios')}
    return result


def constraints(record):
    cg = cgroup_path()
    facts = counters(cg)
    topology = []
    for cpu in sorted(os.sched_getaffinity(0)):
        base = Path(f'/sys/devices/system/cpu/cpu{cpu}/topology')
        topology.append({'cpu': cpu, 'physical_package_id': (base / 'physical_package_id').read_text().strip(),
                         'core_id': (base / 'core_id').read_text().strip()})
    facts.update(cgroup=str(cg), cpus=sorted(os.sched_getaffinity(0)), GOMAXPROCS=os.getenv('GOMAXPROCS'),
                 cpu_topology=topology,
                 io_stat_available='io.stat' in facts,
                 cgroup_controllers=(cg / 'cgroup.controllers').read_text().strip(),
                 io_unavailable_reason=None if 'io.stat' in facts else
                    'io controller is not delegated in the user cgroup hierarchy; IOAccounting=yes alone cannot expose io.stat; GNU time File system inputs remains available, cgroup byte deltas are unavailable.',
                 gnu_time_rss='Maximum child RSS, not simultaneous total RSS; aggregate memory uses cgroup memory.peak.')
    write_json(record / 'constraints.json', facts)
    if (facts.get('memory.max') != '17179869184' or facts.get('memory.swap.max') != '0'
            or facts['cpus'] != [0, 1, 2, 3] or facts['GOMAXPROCS'] != '4'
            or len({(item['physical_package_id'], item['core_id']) for item in topology}) != 4):
        raise RuntimeError('Actual cgroup/physical CPU constraints do not match the experiment')
    if any(name not in facts for name in REQUIRED):
        raise RuntimeError('Required cgroup counters unavailable')
    with (record / 'cgroup.path').open('x') as stream:
        stream.write(str(cg) + '\n')


def await_observer(record):
    deadline = time.monotonic() + 30
    while not (record / 'observer-ready.json').exists():
        if (record / 'observer-error.json').exists():
            raise RuntimeError(f'Observer failed: {read_json(record / "observer-error.json")}')
        if time.monotonic() >= deadline:
            raise TimeoutError('External observer did not open counters before workload start')
        time.sleep(0.02)
    ready = read_json(record / 'observer-ready.json')
    if ready['cgroup'] != (record / 'cgroup.path').read_text().strip() or not all(x in ready['opened'] for x in REQUIRED):
        raise RuntimeError('Observer handshake does not cover this cgroup')


def observe(record):
    fds, last, baseline = {}, {}, {}
    try:
        deadline = time.monotonic() + 60
        while not (record / 'cgroup.path').exists():
            if (record / 'scope-exit-status').exists():
                raise RuntimeError('Scope exited before publishing its cgroup')
            if time.monotonic() >= deadline:
                raise TimeoutError('Scope did not publish its cgroup')
            time.sleep(0.02)
        cg = Path((record / 'cgroup.path').read_text().strip())
        for name in COUNTERS:
            try:
                fds[name] = os.open(cg / name, os.O_RDONLY | os.O_CLOEXEC)
            except OSError:
                if name in REQUIRED:
                    raise
        baseline = counters(fds=fds)
        if any(name not in baseline for name in REQUIRED):
            raise RuntimeError('Cannot read required persistent counter descriptors')
        write_json(record / 'cgroup-before.json', baseline)
        write_json(record / 'observer-ready.json', {'cgroup': str(cg), 'opened': sorted(fds),
                   'pid': os.getpid(), 'timestamp_ns': time.time_ns(), 'persistent_descriptors': True})
        poller = select.poll()
        poller.register(fds['memory.events'], select.POLLPRI | select.POLLERR)
        acknowledged = False
        while True:
            sample = counters(fds=fds)
            for key, value in sample.items():
                if key != 'read_errors':
                    last[key] = value
            last['read_errors'] = sample.get('read_errors', {})
            write_json(record / 'memory-observer.json', last, replace=True)
            if (record / 'workload-done.json').exists() and not acknowledged:
                write_json(record / 'observer-workload-final.json', sample)
                with (record / 'observer-final-ack').open('x') as stream:
                    stream.write(str(time.time_ns()) + '\n')
                acknowledged = True
            if (record / 'scope-exit-status').exists():
                write_json(record / 'cgroup-observer-final.json', {'counters': last, 'last_read': sample,
                           'delta': deltas(baseline, last), 'after_scope_exit': True,
                           'complete_final_read': all(name in sample for name in REQUIRED),
                           'workload_final_handshake': acknowledged})
                break
            poller.poll(20)
    except Exception as error:
        write_json(record / 'observer-error.json', {'error': f'{type(error).__name__}: {error}',
                   'timestamp_ns': time.time_ns(), 'last_counters': last})
        raise
    finally:
        for fd in fds.values():
            os.close(fd)


def finalize(record, exit_code):
    facts = counters(cgroup_path())
    write_json(record / 'cgroup-final.json', facts)
    write_json(record / 'workload-done.json', {'exit_code': exit_code, 'timestamp_ns': time.time_ns()})
    deadline = time.monotonic() + 15
    while not (record / 'observer-final-ack').exists():
        if (record / 'observer-error.json').exists() or time.monotonic() >= deadline:
            raise RuntimeError('Observer did not acknowledge final workload counters')
        time.sleep(0.02)


def cache_drop(record):
    socket_path = Path(os.environ.get('CACHE_SOCKET', str(REPO / 'tmp' / EXP_DIR.name / 'cache.sock')))
    facts = {'socket': str(socket_path), 'request': 'drop\\n', 'requested_ns': time.time_ns()}
    try:
        response = bytearray()
        with socket.socket(socket.AF_UNIX, socket.SOCK_STREAM) as client:
            client.settimeout(float(os.environ.get('CACHE_TIMEOUT_SECONDS', '180')))
            client.connect(str(socket_path))
            client.sendall(b'drop\n')
            while b'\n' not in response:
                chunk = client.recv(65536)
                if not chunk:
                    raise RuntimeError('Cache service closed without a complete JSON line')
                response.extend(chunk)
                if len(response) > 1 << 20:
                    raise RuntimeError('Cache response exceeds one MiB')
        report = json.loads(bytes(response).split(b'\n', 1)[0])
        facts['response'] = report
        if type(report.get('exit_status')) is not int or report['exit_status'] != 0:
            raise RuntimeError('Cache service did not report exit_status 0')
        facts['status'] = 'ok'
    except Exception as error:
        facts.update(status='failed', error=f'{type(error).__name__}: {error}')
    facts['recorded_ns'] = time.time_ns()
    write_json(record / 'cache-clear.json', facts)
    if facts['status'] != 'ok':
        write_json(record / 'status.json', {'status': 'cache_clear_failed', 'timing_sample': False,
                   'error': facts['error'], 'finished_ns': time.time_ns()})
        raise RuntimeError(facts['error'])


def collect_journals(record):
    unit = (record / 'scope-unit').read_text().strip()
    attempt = record / 'journals'
    attempt.mkdir(exist_ok=True)
    index = 1
    while True:
        directory = attempt / f'attempt-{index:03d}'
        try:
            directory.mkdir()
            break
        except FileExistsError:
            index += 1
    reports, oom = [], False
    for kind, options in (('user', ['--user']), ('system', [])):
        for number in range(3):
            command = ['journalctl', *options, '--no-pager', '--quiet', '-o', 'short-iso-precise', '-u', unit]
            try:
                result = subprocess.run(command, capture_output=True, text=True, timeout=10)
                text, error, rc = result.stdout, result.stderr, result.returncode
            except (OSError, subprocess.TimeoutExpired) as exception:
                text, error, rc = '', str(exception), -1
            log = directory / f'{kind}-{number + 1}.log'
            with log.open('x') as stream:
                stream.write(text)
            reports.append({'kind': kind, 'argv': command, 'returncode': rc,
                            'stdout': str(log.relative_to(record)), 'stderr': error})
            explicit = ("Failed with result 'oom-kill'" in text or 'killed by the OOM killer' in text
                        or 'A process of this unit has been killed by the OOM killer' in text)
            oom = oom or explicit
            if explicit or number == 2:
                break
            time.sleep(0.25)
    write_json(directory / 'commands.json', reports)
    return oom, [item['stdout'] for item in reports]


def resolve(record):
    journal_oom, journals = collect_journals(record)
    baseline = read_json(record / 'cgroup-before.json') if (record / 'cgroup-before.json').exists() else {}
    candidates = []
    for name in ('cgroup-final.json', 'observer-workload-final.json', 'memory-observer.json'):
        if (record / name).exists():
            candidates.append((name, read_json(record / name)))
    if (record / 'cgroup-observer-final.json').exists():
        candidates.append(('cgroup-observer-final.json', read_json(record / 'cgroup-observer-final.json')['counters']))
    final = max(candidates, key=lambda item: item[1].get('timestamp_ns', 0)) if candidates else ('unavailable', {})
    delta = deltas(baseline, final[1])
    event_oom = any(value > 0 for key, value in delta.get('memory.events', {}).items() if key in OOM_KEYS)
    oom = event_oom or journal_oom
    rc_path = record / 'exit-status'
    scope_rc = int((record / 'scope-exit-status').read_text())
    rc = int(rc_path.read_text()) if rc_path.exists() else scope_rc
    stderr = (record / 'stderr').read_text(errors='replace') if (record / 'stderr').exists() else ''
    timeout = rc == 124 or ('timeout: sending signal ' in stderr)
    constraints_ok = (record / 'constraints.json').exists() and (record / 'observer-ready.json').exists()
    cache_ok = (record / 'cache-clear.json').exists() and read_json(record / 'cache-clear.json').get('status') == 'ok'
    final_ack = (record / 'observer-final-ack').exists()
    status = ('oom' if oom else 'timeout' if timeout else 'ok'
              if rc == 0 and scope_rc == 0 and constraints_ok and cache_ok and final_ack else 'failed')
    facts = {'status': status, 'exit_code': rc, 'scope_exit_code': scope_rc, 'finished_ns': time.time_ns(),
             'oom_kill_detected': oom, 'oom_from_counter_delta': event_oom, 'oom_from_unit_journal': journal_oom,
             'timeout_detected': timeout, 'counters_source': final[0], 'cgroup_delta': delta,
             'memory_peak_bytes': max((int(item[1].get('memory.peak', '0')) for item in candidates), default=0),
             'cgroup_io_available': 'io.stat' in delta,
             'cgroup_io_unavailable_reason': None if 'io.stat' in delta else 'io controller not delegated; use GNU time File system inputs; no cgroup byte delta available',
             'journal_sources': journals,
             'counter_baseline_available': bool(baseline), 'final_workload_handshake': final_ack,
             'timing_sample': status == 'ok' and not (record / 'verification.json').exists(),
             'exit137_is_not_oom_evidence': True}
    if not (record / 'status.json').exists():
        write_json(record / 'status.json', facts)
    else:
        directory = record / 'resolutions'
        directory.mkdir(exist_ok=True)
        index = 1
        while (directory / f'{index:03d}.json').exists():
            index += 1
        write_json(directory / f'{index:03d}.json', facts)
    if status == 'ok' and not (record / 'finished').exists():
        with (record / 'finished').open('x') as stream:
            stream.write(str(time.time_ns()) + '\n')
    print(status)
    return facts


def validate_success(record):
    facts = read_json(record / 'status.json')
    constraints_facts = read_json(record / 'constraints.json')
    cache = read_json(record / 'cache-clear.json')
    before = read_json(record / 'cgroup-before.json')
    if (facts['status'] != 'ok' or facts['exit_code'] != 0 or facts['scope_exit_code'] != 0
            or int((record / 'exit-status').read_text()) != 0
            or int((record / 'scope-exit-status').read_text()) != 0
            or not (record / 'finished').exists() or not (record / 'observer-final-ack').exists()
            or cache['status'] != 'ok' or cache['response']['exit_status'] != 0
            or constraints_facts['memory.max'] != '17179869184' or constraints_facts['memory.swap.max'] != '0'
            or constraints_facts['cpus'] != [0, 1, 2, 3] or constraints_facts['GOMAXPROCS'] != '4'
            or not all(name in before for name in REQUIRED) or not (record / 'time.txt').is_file()):
        raise RuntimeError(f'Existing record is not a validated successful invocation: {record}')
    if (record / 'cell.json').exists():
        cell = read_json(record / 'cell.json')
        if cell['status'] != 'ok' or cell['exit_code'] != 0:
            raise RuntimeError('Cell result did not succeed')
    if (record / 'verification.json').exists() and read_json(record / 'verification.json')['status'] != 'passed':
        raise RuntimeError('Full-output verification did not pass')


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('mode', choices=('constraints', 'await-observer', 'watch', 'finalize', 'cache-drop', 'resolve', 'validate-success'))
    parser.add_argument('record', type=Path)
    parser.add_argument('exit_code', nargs='?', type=int)
    args = parser.parse_args()
    record = local_record(args.record)
    if args.mode == 'constraints':
        constraints(record)
    elif args.mode == 'await-observer':
        await_observer(record)
    elif args.mode == 'watch':
        observe(record)
    elif args.mode == 'finalize':
        finalize(record, args.exit_code)
    elif args.mode == 'cache-drop':
        cache_drop(record)
    elif args.mode == 'resolve':
        resolve(record)
    else:
        validate_success(record)


if __name__ == '__main__':
    main()
