#!/usr/bin/env python3
"""Summarize raw cold-cache measurements without overwriting past reports."""
import argparse
import json
from pathlib import Path
import statistics


def load(path):
    return json.loads(path.read_text()) if path.exists() else None


def elapsed_seconds(text):
    values = [float(value) for value in text.split(':')]
    result = 0.0
    for value in values:
        result = result * 60 + value
    return result


def time_values(path):
    result = {}
    if not path.exists():
        return result
    for line in path.read_text().splitlines():
        line = line.strip()
        if line.startswith('Elapsed (wall clock) time'):
            result['wall_seconds'] = elapsed_seconds(line.split('): ', 1)[1])
        else:
            for prefix, key, conversion in (
                ('User time (seconds): ', 'user_seconds', float),
                ('System time (seconds): ', 'system_seconds', float),
                ('Maximum resident set size (kbytes): ', 'max_child_rss_kib', int),
                ('File system inputs: ', 'file_system_inputs', int),
                ('File system outputs: ', 'file_system_outputs', int),
                ('Exit status: ', 'time_exit_status', int),
            ):
                if line.startswith(prefix):
                    result[key] = conversion(line[len(prefix):])
    if 'user_seconds' in result and 'system_seconds' in result:
        result['cpu_seconds'] = result['user_seconds'] + result['system_seconds']
    return result


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument('--results', type=Path, default=Path(__file__).resolve().parent / 'results')
    parser.add_argument('--name', default='summary-final')
    args = parser.parse_args()
    if Path(args.name).name != args.name:
        parser.error('--name must be a filename stem')
    records = []
    for directory in sorted(args.results.glob('timing-*')):
        if not directory.is_dir():
            continue
        parts = directory.name.split('-')
        if len(parts) < 5:
            continue
        status = load(directory / 'status.json')
        cell = load(directory / 'cell.json')
        cache = load(directory / 'cache-clear.json')
        constraints = load(directory / 'constraints.json')
        time_path = directory / 'time.txt'
        if not time_path.exists() and (directory / 'time-scope.txt').exists():
            time_path = directory / 'time-scope.txt'
        metrics = time_values(time_path)
        if status and 'memory_peak_bytes' in status:
            metrics['cgroup_memory_peak_bytes'] = status['memory_peak_bytes']
        record = {'record': directory.name, 'media': parts[1], 'fields': int(parts[2]),
                  'engine': parts[3], 'round': parts[4], 'status': status,
                  'cell': cell, 'cache_drop': cache, 'constraints': constraints,
                  'time_source': time_path.name,
                  'time_boundary': 'workload tree' if time_path.name == 'time.txt' else 'scope and wrapper; failure evidence only',
                  'time': metrics}
        record['accepted_timing_sample'] = bool(
            status and status.get('status') == 'ok' and status.get('exit_code') == 0
            and status.get('scope_exit_code') == 0 and status.get('final_workload_handshake')
            and cell and cell.get('status') == 'ok' and cell.get('exit_code') == 0
            and cell.get('timing_sample') and cell.get('compressed_output')
            and cell.get('source_commit') == 'd2eac357f3fc126755b43ff8b44a64dc55bc0ec2'
            and cache and cache.get('status') == 'ok' and cache.get('response', {}).get('exit_status') == 0
            and constraints and constraints.get('memory.max') == '17179869184'
            and constraints.get('memory.swap.max') == '0' and constraints.get('cpus') == [0, 1, 2, 3]
            and constraints.get('GOMAXPROCS') == '4' and metrics.get('time_exit_status') == 0
            and time_path.name == 'time.txt')
        record['cgroup_io_available'] = status.get('cgroup_io_available') if status else None
        record['cgroup_io_unavailable_reason'] = status.get('cgroup_io_unavailable_reason') if status else None
        records.append(record)
    aggregates = []
    for media in ('nvme', 'hdd'):
        for fields in (17, 540, 568):
            for engine in ('immkv', 'echtvar'):
                group = [r for r in records if (r['media'], r['fields'], r['engine']) == (media, fields, engine)]
                successful = [r for r in group if r['accepted_timing_sample']]
                medians = {}
                for key in ('wall_seconds', 'cpu_seconds', 'max_child_rss_kib', 'file_system_inputs', 'cgroup_memory_peak_bytes'):
                    values = [r['time'][key] for r in successful if key in r['time']]
                    if values:
                        medians[key] = statistics.median(values)
                aggregates.append({'media': media, 'fields': fields, 'engine': engine,
                                   'attempts': len(group), 'successful': len(successful),
                                   'medians': medians, 'records': [r['record'] for r in group]})
    report = {'records': records, 'aggregates': aggregates,
              'note': 'Only successful complete runs enter medians. GNU time RSS is maximum child RSS, not simultaneous group RSS.'}
    output_json = args.results / (args.name + '.json')
    output_md = args.results / (args.name + '.md')
    if output_json.exists() or output_md.exists():
        raise SystemExit('refusing to overwrite an existing summary; choose a new --name')
    lines = ['# Cold-cache timing summary', '',
             'Only successful complete runs enter medians; unsuccessful runs remain in the raw records.', '',
             '| Medium | Fields | Engine | Complete / attempts | Median wall (s) | Median CPU (s) | Median max child RSS (KiB) | Median cgroup peak (bytes) | Sources |',
             '|---|---:|---|---:|---:|---:|---:|---:|---|']
    for aggregate in aggregates:
        medians = aggregate['medians']
        def number(key):
            return f'{medians[key]:.3f}' if key in medians else '—'
        sources = ', '.join(f'[{name}]({name}/status.json)' for name in aggregate['records']) or 'Not run'
        lines.append(f"| {aggregate['media']} | {aggregate['fields']} | {aggregate['engine']} | {aggregate['successful']} / {aggregate['attempts']} | {number('wall_seconds')} | {number('cpu_seconds')} | {number('max_child_rss_kib')} | {number('cgroup_memory_peak_bytes')} | {sources} |")
    lines.extend(['', '## Raw repetitions', '',
                  '| Record | Status | Wall (s) | CPU (s) | Max child RSS (KiB) | Cgroup peak (bytes) | Input blocks |',
                  '|---|---|---:|---:|---:|---:|---:|'])
    for record in records:
        values = record['time']
        state = record['status'].get('status', 'unknown') if record['status'] else 'incomplete'
        name = record['record']
        lines.append(f"| [{name}]({name}/{record['time_source']}) | {state} | {values.get('wall_seconds', '—')} | {values.get('cpu_seconds', '—')} | {values.get('max_child_rss_kib', '—')} | {values.get('cgroup_memory_peak_bytes', '—')} | {values.get('file_system_inputs', '—')} |")
    lines.extend(['', 'The JSON report preserves each cell, cache-drop report and cgroup status alongside GNU time metrics.', ''])
    with output_json.open('x') as stream:
        json.dump(report, stream, indent=2)
        stream.write('\n')
    with output_md.open('x') as stream:
        stream.write('\n'.join(lines))
    print(output_md)


if __name__ == '__main__':
    main()
