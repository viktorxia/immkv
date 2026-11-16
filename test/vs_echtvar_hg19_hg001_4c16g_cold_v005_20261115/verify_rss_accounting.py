#!/usr/bin/env python3
"""Reproduce the RSS accounting control, separately from benchmark timing.

This control touches only 100 MiB of anonymous memory in one child at a time,
uses CPUs 4-7, and never opens benchmark libraries or queries or clears caches.
The original source and raw logs are preserved in results/rss-accounting-control.
Run with: python3 -I verify_rss_accounting.py
Each invocation creates a new directory; existing evidence is never replaced.
"""
import argparse
import json
import os
from pathlib import Path
import shutil
import subprocess
import sys
import time


def main():
    experiment = Path(__file__).resolve().parent
    archive = experiment / 'results' / 'rss-accounting-control'
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--output', type=Path,
                        default=archive / 'reproductions' / str(time.time_ns()),
                        help='New output directory inside this experiment results tree')
    args = parser.parse_args()
    output = args.output.resolve()
    if not output.is_relative_to((experiment / 'results').resolve()):
        parser.error('--output must remain inside this experiment results tree')
    cpus = {4, 5, 6, 7}
    if not cpus.issubset(os.sched_getaffinity(0)):
        parser.error('CPUs 4-7 must all be available; this control must avoid CPUs 0-3')
    os.sched_setaffinity(0, cpus)
    output.mkdir(parents=True, exist_ok=False)
    results = output / 'results'
    results.mkdir()
    source = output / 'smoke.py'
    shutil.copyfile(archive / 'smoke.py', source)
    cases = ('main', 'thread', 'thread-posix-spawn', 'thread-direct')
    commands = []
    for case in cases:
        command = ['/usr/bin/time', '-v', '-o', str(results / (case + '.outer.time.txt')),
                   'timeout', '10s', sys.executable, '-I', str(source), case]
        commands.append(command)
        with (results / (case + '.stdout')).open('x') as stdout:
            subprocess.run(command, stdout=stdout, check=True)
    with (output / 'commands.json').open('x') as stream:
        json.dump({'nonbenchmark_control': True, 'affinity': sorted(cpus),
                   'allocation_bytes_per_child': 100 * 1024 * 1024,
                   'commands': commands}, stream, indent=2)
        stream.write('\n')
    print(output)


if __name__ == '__main__':
    main()
