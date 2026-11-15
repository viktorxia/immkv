#!/usr/bin/env python3
"""Capture actual scope limits, aggregate memory counters, and run outcome."""
import argparse
import json
import os
from pathlib import Path
import sys
import time


def write_json(path, value):
    with path.open('x') as f:
        json.dump(value, f, indent=2, sort_keys=True)
        f.write('\n')


def cgroup_path():
    return Path('/sys/fs/cgroup' + next(x.split(':', 2)[2] for x in Path('/proc/self/cgroup').read_text().splitlines() if x.startswith('0::')))


def counters(cg):
    result = {'timestamp_ns': time.time_ns()}
    for name in ('memory.max', 'memory.swap.max', 'memory.peak', 'memory.current', 'memory.events', 'memory.events.local', 'cpu.stat', 'io.stat'):
        try:
            result[name] = (cg / name).read_text().strip()
        except FileNotFoundError:
            pass
    return result


def main():
    p = argparse.ArgumentParser(description=__doc__)
    p.add_argument('mode', choices=('constraints', 'finalize', 'watch'))
    p.add_argument('record', type=Path)
    p.add_argument('exit_code', nargs='?', type=int)
    a = p.parse_args()
    record = a.record
    if a.mode == 'constraints':
        cg = cgroup_path()
        facts = counters(cg)
        facts.update(cgroup=str(cg), cpus=sorted(os.sched_getaffinity(0)), GOMAXPROCS=os.getenv('GOMAXPROCS'),
                     gnu_time_rss='Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.')
        write_json(record / 'constraints.json', facts)
        (record / 'cgroup.path').write_text(str(cg) + '\n')
        for name in ('memory.events', 'memory.events.local', 'memory.peak'):
            if name in facts:
                (record / (name + '.before')).write_text(facts[name] + '\n')
        if (facts.get('memory.max') != '17179869184' or facts.get('memory.swap.max') != '0'
                or facts['cpus'] != [0, 1, 2, 3] or facts['GOMAXPROCS'] != '4'):
            raise SystemExit('Actual cgroup/CPU constraints do not match the experiment')
    elif a.mode == 'watch':
        # Outside the scope: retain last readable counters even if the wrapper dies.
        while not (record / 'scope-exit-status').exists():
            try:
                cg = Path((record / 'cgroup.path').read_text().strip())
                facts = counters(cg)
                if 'memory.events' in facts:
                    tmp = record / 'memory-observer.json.partial'
                    tmp.write_text(json.dumps(facts, indent=2, sort_keys=True) + '\n')
                    tmp.replace(record / 'memory-observer.json')
            except FileNotFoundError:
                pass
            time.sleep(0.05)
    else:
        cg = cgroup_path()
        facts = counters(cg)
        for name in ('memory.events', 'memory.events.local', 'memory.peak'):
            if name in facts:
                (record / name).write_text(facts[name] + '\n')
        write_json(record / 'cgroup-final.json', facts)
        events = dict(line.split() for line in facts.get('memory.events', '').splitlines())
        oom = int(events.get('oom_kill', '0')) > 0 or int(events.get('oom_group_kill', '0')) > 0
        before = record / 'memory.events.before'
        if before.exists():
            baseline = dict(line.split() for line in before.read_text().splitlines())
            oom = any(int(events.get(k, '0')) > int(baseline.get(k, '0')) for k in ('oom_kill', 'oom_group_kill'))
        rc = a.exit_code
        status = 'oom' if oom else 'ok' if rc == 0 else 'timeout' if rc in (124, 137) else 'failed'
        write_json(record / 'status.json', {'status': status, 'exit_code': rc, 'finished_ns': time.time_ns(),
                   'memory_peak_bytes': int(facts.get('memory.peak', '0')), 'oom_kill_detected': oom})
        if status == 'ok':
            (record / 'finished').write_text(str(time.time_ns()) + '\n')


if __name__ == '__main__':
    main()
