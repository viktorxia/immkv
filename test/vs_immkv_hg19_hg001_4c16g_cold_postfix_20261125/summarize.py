#!/usr/bin/env python3
"""Summarize raw cold-cache measurements and compare with v0.0.5 medians."""
import argparse
import json
from pathlib import Path
import statistics

V005_MEDIANS = {
    ('nvme', 17): {'wall_seconds': 13.20, 'cpu_seconds': 33.15, 'max_child_rss_kib': 288500, 'cgroup_memory_peak_bytes': 4798017536},
    ('nvme', 540): {'wall_seconds': 158.86, 'cpu_seconds': 616.02, 'max_child_rss_kib': 2331400, 'cgroup_memory_peak_bytes': 6947467264},
    ('nvme', 568): {'wall_seconds': 172.47, 'cpu_seconds': 670.46, 'max_child_rss_kib': 3317232, 'cgroup_memory_peak_bytes': 7959490560},
    ('hdd', 17): {'wall_seconds': 96.46, 'cpu_seconds': 204.67, 'max_child_rss_kib': 250988, 'cgroup_memory_peak_bytes': 7506931712},
    ('hdd', 540): {'wall_seconds': 1468.32, 'cpu_seconds': 674.01, 'max_child_rss_kib': 675300, 'cgroup_memory_peak_bytes': 5348339712},
    ('hdd', 568): {'wall_seconds': 1469.52, 'cpu_seconds': 719.20, 'max_child_rss_kib': 859164, 'cgroup_memory_peak_bytes': 5525741568},
}


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
    source_facts = load(args.results / 'preparation/source-facts.json') or {}
    expected_manifest = source_facts.get('sources_manifest_sha256')
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
                  'time_source': time_path.name, 'time': metrics}
        record['accepted_timing_sample'] = bool(
            status and status.get('status') == 'ok' and status.get('exit_code') == 0
            and status.get('scope_exit_code') == 0 and status.get('final_workload_handshake')
            and cell and cell.get('status') == 'ok' and cell.get('exit_code') == 0
            and cell.get('timing_sample') and cell.get('compressed_output')
            and cell.get('sources_manifest_sha256') == expected_manifest
            and cache and cache.get('status') == 'ok' and cache.get('response', {}).get('exit_status') == 0
            and constraints and constraints.get('memory.max') == '17179869184'
            and constraints.get('memory.swap.max') == '0' and constraints.get('cpus') == [0, 1, 2, 3]
            and constraints.get('GOMAXPROCS') == '4' and metrics.get('time_exit_status') == 0
            and time_path.name == 'time.txt')
        records.append(record)
    aggregates = []
    for media in ('nvme', 'hdd'):
        for fields in (17, 540, 568):
            group = [r for r in records if (r['media'], r['fields'], r['engine']) == (media, fields, 'immkv')]
            successful = [r for r in group if r['accepted_timing_sample']]
            medians = {}
            for key in ('wall_seconds', 'cpu_seconds', 'max_child_rss_kib', 'file_system_inputs', 'cgroup_memory_peak_bytes'):
                values = [r['time'][key] for r in successful if key in r['time']]
                if values:
                    medians[key] = statistics.median(values)
            baseline = V005_MEDIANS[(media, fields)]
            aggregates.append({'media': media, 'fields': fields, 'engine': 'immkv',
                               'attempts': len(group), 'successful': len(successful),
                               'medians': medians, 'v005_baseline': baseline,
                               'records': [r['record'] for r in group]})
    report = {'records': records, 'aggregates': aggregates}
    output_json = args.results / (args.name + '.json')
    output_md = args.results / (args.name + '.md')
    if output_json.exists() or output_md.exists():
        raise SystemExit('refusing to overwrite an existing summary; choose a new --name')
    lines = ['# Cold-cache timing summary (post-fix vs v0.0.5)', '',
             '| Medium | Fields | Complete / attempts | Post-fix wall (s) | v0.0.5 wall (s) | Post-fix CPU (s) | v0.0.5 CPU (s) | Post-fix cgroup peak (bytes) | v0.0.5 cgroup peak (bytes) | Sources |',
             '|---|---:|---:|---:|---:|---:|---:|---:|---:|---|']
    for aggregate in aggregates:
        medians = aggregate['medians']
        base = aggregate['v005_baseline']
        def number(key):
            return f'{medians[key]:.2f}' if key in medians else '—'
        sources = ', '.join(f'[{name}]({name}/status.json)' for name in aggregate['records']) or 'Not run'
        lines.append(f"| {aggregate['media']} | {aggregate['fields']} | {aggregate['successful']} / {aggregate['attempts']} | {number('wall_seconds')} | {base['wall_seconds']:.2f} | {number('cpu_seconds')} | {base['cpu_seconds']:.2f} | {medians.get('cgroup_memory_peak_bytes', '—')} | {base['cgroup_memory_peak_bytes']} | {sources} |")
    lines.extend(['', '## Raw repetitions', '',
                  '| Record | Status | Wall (s) | CPU (s) | Max child RSS (KiB) | Cgroup peak (bytes) | Input blocks |',
                  '|---|---|---:|---:|---:|---:|---:|'])
    for record in records:
        values = record['time']
        state = record['status'].get('status', 'unknown') if record['status'] else 'incomplete'
        name = record['record']
        lines.append(f"| [{name}]({name}/{record['time_source']}) | {state} | {values.get('wall_seconds', '—')} | {values.get('cpu_seconds', '—')} | {values.get('max_child_rss_kib', '—')} | {values.get('cgroup_memory_peak_bytes', '—')} | {values.get('file_system_inputs', '—')} |")
    with output_json.open('x') as stream:
        json.dump(report, stream, indent=2)
        stream.write('\n')
    with output_md.open('x') as stream:
        stream.write('\n'.join(lines) + '\n')
    print(output_md)


if __name__ == '__main__':
    main()
