#!/usr/bin/env python3
"""Summarize retained release records without launching tools or reading datasets.

Examples:
  python3 -I summarize.py --format markdown
  python3 -I summarize.py --format json --output results/summary-intermediate.json

Output files are created exclusively; an existing file is never overwritten.
Missing, running, failed, OOM, timeout and deliberately skipped attempts remain
separate states. Only three fully valid rounds produce medians or paired ratios.
"""
import argparse
import json
import math
from pathlib import Path
import re
import statistics
import sys

HERE = Path(__file__).resolve().parent
ROWS, HITS = 3933714, 3875208
METRICS = ('wall_seconds', 'user_seconds', 'system_seconds', 'cpu_percent',
           'max_child_rss_kib', 'cgroup_peak_bytes')
EXPECTED = {'memory.max': '17179869184', 'memory.swap.max': '0',
            'cpus': [0, 1, 2, 3], 'GOMAXPROCS': '4'}


def evidence(path, value, field=None):
    return {'value': value, 'source': str(path), **({'field': field} if field else {})}


def read_json(path, issues):
    if not path.is_file():
        return None
    try:
        value = json.loads(path.read_text())
        if not isinstance(value, dict):
            raise ValueError('expected JSON object')
        return value
    except (OSError, ValueError) as exc:
        issues.append(f'{path}: {exc}')
        return None


def text(path):
    try:
        return path.read_text()
    except FileNotFoundError:
        return ''


def elapsed(value):
    """GNU time accepts m:ss, h:mm:ss, and decimal seconds."""
    parts = value.strip().split(':')
    if not 1 <= len(parts) <= 3:
        raise ValueError('invalid GNU elapsed time')
    result = 0.0
    for part in parts:
        result = result * 60 + float(part)
    if not math.isfinite(result) or result < 0:
        raise ValueError('invalid GNU elapsed time')
    return result


def go_duration(value):
    """Parse annotate's Go duration strings, including minute-form values."""
    match = re.fullmatch(r'(?:(\d+)h)?(?:(\d+)m)?([\d.]+)(ms|s)', value)
    if not match:
        raise ValueError('invalid Go duration')
    hours, minutes, seconds, unit = match.groups()
    result = int(hours or 0) * 3600 + int(minutes or 0) * 60
    result += float(seconds) / (1000 if unit == 'ms' else 1)
    if not math.isfinite(result) or result < 0:
        raise ValueError('invalid Go duration')
    return result


def parse_time(path, issues):
    patterns = {
        'wall_seconds': (r'Elapsed \(wall clock\) time \(h:mm:ss or m:ss\):\s*(\S+)', elapsed),
        'user_seconds': (r'User time \(seconds\):\s*(\S+)', float),
        'system_seconds': (r'System time \(seconds\):\s*(\S+)', float),
        'cpu_percent': (r'Percent of CPU this job got:\s*([\d.]+)%', float),
        'max_child_rss_kib': (r'Maximum resident set size \(kbytes\):\s*(\d+)', int),
        'filesystem_inputs_blocks': (r'File system inputs:\s*(\d+)', int),
    }
    result, data = {}, text(path)
    for name, (pattern, convert) in patterns.items():
        match = re.search(pattern, data)
        if match:
            try:
                value = convert(match[1])
                if not math.isfinite(value) or value < 0:
                    raise ValueError('negative or nonfinite value')
                result[name] = evidence(path, value, name)
            except ValueError as exc:
                issues.append(f'{path}: {name}: {exc}')
    match = re.search(r'Exit status:\s*(\d+)', data)
    if match:
        result['exit_code'] = evidence(path, int(match[1]), 'Exit status')
    return result


def constraints(record, issues):
    path = record / 'constraints.json'
    data = read_json(path, issues)
    mismatches = []
    if data is not None:
        for key, expected in EXPECTED.items():
            actual = data.get(key)
            if key in ('memory.max', 'memory.swap.max', 'GOMAXPROCS'):
                actual = str(actual)
            if actual != expected:
                mismatches.append(f'{key}: observed {actual!r}, required {expected!r}')
    return {'valid': data is not None and not mismatches, 'source': str(path),
            'observed': data, 'mismatches': mismatches}


def outcome(record, issues):
    path = record / 'status.json'
    data = read_json(path, issues)
    if data is not None:
        state = data.get('status', 'unknown')
    elif (record / 'started').exists() or (record / 'cell.json').exists():
        state = 'started'
    else:
        state = 'unknown'
    original = {'state': state, 'source': str(path), 'record': data}
    resolution_path = record / 'status-resolution.json'
    resolution = read_json(resolution_path, issues)
    source = path
    if resolution is not None:
        resolved = resolution.get('status')
        if isinstance(resolved, str) and resolved:
            state, source = resolved, resolution_path
        else:
            issues.append(f'{resolution_path}: missing or invalid status')
    return {'state': state, 'source': str(source), 'record': data,
            'original_status': original,
            'status_resolution': {
                'source': str(resolution_path), 'record': resolution,
                'evidence_sources': [str(record / name) for name in resolution.get('evidence', [])]
                if resolution else []},
            'directory_exists': record.is_dir(),
            'started_source': str(record / 'started') if (record / 'started').exists() else None}


def memory_peak(record, status, issues):
    """Retain a positive observed peak without claiming missing final counters."""
    data = status['record'] or {}
    final = read_json(record / 'cgroup-final.json', issues) or {}
    observed = data.get('last_observed_cgroup') or {}
    observed_source = record / 'status.json'
    if not observed:
        observed = read_json(record / 'memory-observer.json', issues) or {}
        observed_source = record / 'memory-observer.json'
    # The fallback wrapper copies observer counters to memory.peak. That copy
    # is still an observation, even when journal evidence confirms the OOM.
    candidates = [
        (data.get('memory_peak_bytes'), record / 'status.json', 'memory_peak_bytes', True),
        (final.get('memory.peak'), record / 'cgroup-final.json', 'memory.peak', True),
        (text(record / 'memory.peak').strip(), record / 'memory.peak', 'memory.peak',
         not bool(data.get('last_observed_cgroup'))),
        (observed.get('memory.peak'), observed_source, 'last observed memory.peak', False),
    ]
    for value, source, field, is_final in candidates:
        try:
            peak = int(value)
        except (TypeError, ValueError):
            continue
        if peak <= 0:
            continue  # Zero is an unavailable-counter sentinel, not a measured peak.
        result = evidence(source, peak, field)
        result['final_counters_available'] = is_final
        if not is_final:
            result['note'] = 'Last observed aggregate peak; final counters unavailable.'
        return result
    return None


def status_markdown(status):
    lines = []
    resolution = status['status_resolution']
    if resolution['record'] is not None:
        original = status['original_status']
        lines.extend([f"  - Original scope status={original['state']}; `{original['source']}`.",
                      f"  - Status resolution: `{resolution['source']}`; "
                      + json.dumps(resolution['record'], sort_keys=True) + '.'])
        for path in resolution['evidence_sources']:
            lines.append(f"  - Status resolution evidence: `{path}`.")
    return lines


def chromosome_records(record, cell, issues):
    records = cell.get('chromosomes') if cell else None
    if not isinstance(records, list):
        records = []
        for chromosome in range(1, 23):
            value = read_json(record / f'chr{chromosome}.json', issues)
            if value is not None:
                records.append(value)
    return records


def check_chromosomes(records):
    return (len(records) == 22 and
            sorted(x.get('chromosome', -1) for x in records) == list(range(1, 23)) and
            all(x.get('status') == 'ok' and x.get('slice_exit_code') == 0 and
                x.get('annotation_exit_code') == 0 for x in records))


def sample(root, media, fields, engine, round_number):
    record = root / f'timing-{media}-{fields}-{engine}-r{round_number}-release'
    issues = []
    status = outcome(record, issues)
    scope = constraints(record, issues)
    cell = read_json(record / 'cell.json', issues)
    metrics = parse_time(record / 'time.txt', issues)
    peak_is_final = False
    status_data = status['record'] or {}
    # A skipped round has no quantitative sample, even if stray files exist.
    if status['state'] != 'skipped_after_oom':
        peak = memory_peak(record, status, issues)
        if peak is not None:
            metrics['cgroup_peak_bytes'] = peak
            peak_is_final = peak['final_counters_available']
    else:
        metrics = {}
    count_proof, phases = {}, {}
    if engine == 'immkv':
        path = record / 'immkv.log'
        data = text(path)
        match = re.search(r'rows (\d+), hits (\d+) ', data)
        valid_counts = bool(match and tuple(map(int, match.groups())) == (ROWS, HITS))
        if match:
            count_proof = {name: evidence(path, int(value), name)
                           for name, value in zip(('rows', 'hits'), match.groups())}
        phase = re.search(r'read ([\d.]+)GiB; key ([\d.]+)s / query ([\d.]+)s / format ([\d.]+)s', data)
        if phase:
            phases = {name: evidence(path, float(value), name) for name, value in zip(
                ('read_gib_rounded', 'key_seconds', 'query_seconds', 'format_seconds'), phase.groups())}
            phases['read_gib_rounded']['note'] = (
                'Rounded tool-internal read accounting, not physical I/O; '
                'GNU time filesystem_inputs_blocks is retained separately.')
        wall = re.search(r'\bwall (\S+)', data)
        if wall:
            try:
                phases['wall_seconds'] = evidence(path, go_duration(wall[1]), 'wall_seconds')
            except ValueError:
                pass  # Optional tool-internal phase statistics do not invalidate GNU timing.
    else:
        records = chromosome_records(record, cell, issues)
        valid_counts = check_chromosomes(records)
        count_proof = {'all_22_chromosomes_ok': valid_counts,
                       'source': str(record / 'cell.json'),
                       'chromosome_statuses': records,
                       'timed_output_rows_checked': False,
                       'note': 'Compressed timing output is discarded; fresh output proof is separate.'}
        summaries = []
        for chromosome in range(1, 23):
            path = record / f'chr{chromosome}.anno.log'
            match = re.search(r'\[echtvar\] evaluated (\d+) variants .*?wrote (\d+) variants\.', text(path))
            if match:
                summaries.append({'chromosome': chromosome,
                                  'evaluated': evidence(path, int(match[1]), 'evaluated variants'),
                                  'wrote': evidence(path, int(match[2]), 'wrote variants')})
        count_proof['annotation_log_statistics'] = summaries
        if len(summaries) == 22:
            count_proof['evaluated_total'] = {
                'value': sum(x['evaluated']['value'] for x in summaries),
                'sources': [x['evaluated'] for x in summaries]}
            # These statistics are evidence, not a replacement for fresh row proof.
            if count_proof['evaluated_total']['value'] != ROWS:
                valid_counts = False
                issues.append('All chromosome evaluated counts exist but total differs from expected rows')
    reasons = []
    if status['state'] != 'ok':
        reasons.append(f"scope status is {status['state']}")
    if status_data.get('exit_code') != 0:
        reasons.append('scope exit_code is absent or nonzero')
    if not cell or cell.get('status') != 'ok' or cell.get('exit_code') != 0:
        reasons.append('cell status/exit_code is absent or unsuccessful')
    if cell and any(cell.get(k) != v for k, v in (('engine', engine), ('media', media), ('fields', fields))):
        reasons.append('cell identity does not match matrix label')
    if not scope['valid']:
        reasons.append('actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed')
    if not valid_counts:
        reasons.append('required annotation counts or chromosome completion are unconfirmed')
    if any(k not in metrics for k in METRICS):
        reasons.append('one or more GNU time / cgroup metrics are absent')
    if not peak_is_final:
        reasons.append('final aggregate cgroup peak is unconfirmed')
    if metrics.get('exit_code', {}).get('value') != 0:
        reasons.append('GNU time exit status is absent or nonzero')
    return {'round': round_number, 'record_path': str(record), 'status': status,
            'valid': not reasons and not issues, 'invalid_reasons': reasons,
            'issues': issues, 'constraints': scope, 'metrics': metrics,
            'count_proof': count_proof, 'immkv_phase_statistics': phases,
            'cell_source': str(record / 'cell.json')}


def verifier_counters(path):
    data, result = text(path), {}
    # Some retained verifier versions emit an additional JSON object. Preserve
    # every key rather than guessing which exception counters are relevant.
    try:
        optional_json = json.loads(data)
    except ValueError:
        optional_json = None
    if isinstance(optional_json, dict):
        result['optional_json'] = evidence(path, optional_json, 'verifier stdout JSON')
    patterns = {
        'rows': r'rows (\d+), field comparisons',
        'field_comparisons': r'field comparisons (\d+)',
        'misses_agreed': r'misses agreed (\d+)',
        'value_mismatches': r'value mismatches (\d+)',
        'hit_disagreements': r'hit disagreements (\d+)',
        'echtvar_i32_overflow_losses': r'echtvar i32-overflow value losses \(documented limitation\): (\d+)',
        'float_comparisons': r'float comparisons (\d+)',
        'max_float_absolute_difference': r'max \|echtvar-immkv\| observed:\s*([\d.eE+-]+)',
        'float_tolerance': r'\(tolerance ([\d.eE+-]+)',
    }
    for name, pattern in patterns.items():
        match = re.search(pattern, data)
        if match:
            convert = float if name in ('max_float_absolute_difference', 'float_tolerance') else int
            result[name] = evidence(path, convert(match[1]), name)
    result['verify_values_ok'] = evidence(path, 'VERIFY_VALUES OK:' in data, 'VERIFY_VALUES OK')
    return result


def correctness(root, media, fields, fresh=False):
    prefix = 'verify-echtvar' if fresh else 'verify'
    record = root / f'{prefix}-{media}-{fields}-release'
    issues = []
    status = outcome(record, issues)
    scope = constraints(record, issues)
    path = record / ('cell.json' if fresh else 'verification.json')
    data = read_json(path, issues)
    passed = False
    if data and status['state'] == 'ok' and scope['valid']:
        if fresh:
            chromosomes = chromosome_records(record, data, issues)
            passed = (data.get('status') == 'ok' and data.get('exit_code') == 0 and
                      data.get('fresh_rows') == ROWS and check_chromosomes(chromosomes) and
                      all(x.get('comparison', {}).get('equal') is True and
                          x['comparison'].get('fresh') == x['comparison'].get('reference')
                          for x in chromosomes))
        else:
            comp = data.get('comparison', {})
            passed = (data.get('status') == 'passed' and comp.get('rows') == ROWS and
                      comp.get('field_comparisons') == ROWS * fields and
                      comp.get('value_mismatches') == 0 and comp.get('hit_disagreements') == 0 and
                      data.get('annotate_statistics') == {'rows': ROWS, 'hits': HITS})
    return {'passed': passed and not issues, 'status': status, 'constraints': scope,
            'cgroup_peak': memory_peak(record, status, issues),
            'source': str(path), 'evidence': data, 'issues': issues,
            'verifier_stdout': verifier_counters(record / 'verifier.stdout') if not fresh else {},
            'note': 'Correctness phase only; excluded from timing medians.'}


def aggregate(rounds):
    if not all(r['valid'] for r in rounds):
        return None
    return {metric: {'median': statistics.median([r['metrics'][metric]['value'] for r in rounds]),
                     'minimum': min(r['metrics'][metric]['value'] for r in rounds),
                     'maximum': max(r['metrics'][metric]['value'] for r in rounds),
                     'values': [r['metrics'][metric]['value'] for r in rounds],
                     'sources': [r['metrics'][metric] for r in rounds], 'n': 3}
            for metric in METRICS}


def provenance(root):
    result = {}
    for relative in ('build/source-facts.json', 'input-preflight.json'):
        issues = []
        value = read_json(root / relative, issues)
        if relative.endswith('source-facts.json') and value:
            value = {k: v for k, v in value.items() if k != 'source_files'}
        result[relative] = {'source': str(root / relative), 'record': value, 'issues': issues}
    for relative in ('build/binaries.sha256', 'build/environment.txt', 'build/build-info.txt'):
        path = root / relative
        result[relative] = {'source': str(path), 'text': text(path) if path.is_file() else None}
    return result


def report(root):
    cells = []
    for media in ('nvme', 'hdd'):
        for fields in (17, 540, 568):
            engines = {}
            for engine in ('immkv', 'echtvar'):
                rounds = [sample(root, media, fields, engine, n) for n in (1, 2, 3)]
                engines[engine] = {'rounds': rounds, 'valid_rounds': sum(r['valid'] for r in rounds),
                                   'summary': aggregate(rounds)}
            ratios = None
            if all(engines[e]['summary'] is not None for e in engines):
                paired = []
                for left, right in zip(engines['immkv']['rounds'], engines['echtvar']['rounds']):
                    a, b = left['metrics']['wall_seconds'], right['metrics']['wall_seconds']
                    if a['value'] <= 0:
                        break
                    paired.append({'round': left['round'], 'value': b['value'] / a['value'],
                                   'numerator': b, 'denominator': a})
                if len(paired) == 3:
                    values = [x['value'] for x in paired]
                    ratios = {'definition': 'echtvar wall / immkv wall, paired by round',
                              'n': 3, 'rounds': paired, 'median': statistics.median(values),
                              'minimum': min(values), 'maximum': max(values),
                              'range_over_median': (max(values) - min(values)) / statistics.median(values)}
            cells.append({'media': media, 'fields': fields, 'engines': engines,
                          'paired_wall_ratios': ratios,
                          'correctness': {'immkv': correctness(root, media, fields),
                                          'fresh_echtvar': correctness(root, media, fields, True)}})
    return {'schema_version': 1, 'experiment': HERE.name, 'results': str(root),
            'expected_query': {'rows': ROWS, 'hits_immkv': HITS,
                               'source': str(HERE / 'verify_output.py')},
            'required_constraints': {'values': EXPECTED, 'source': str(HERE / 'env_tools.py')},
            'provenance': provenance(root), 'cells': cells,
            'notes': [
                'Every metric retains its source path; medians use exactly three complete valid rounds.',
                'OOM, timeout, failed, skipped_after_oom, started and unknown are independent outcomes; no missing result is estimated.',
                'GNU time maximum RSS is maximum child RSS in KiB, not simultaneous aggregate RSS.',
                'Cgroup memory.peak is aggregate charged memory in bytes, including page cache and memfd slices; observer fallback may precede final peak.',
                'CPU percent is GNU time percent (100 percent equals one CPU); four allowed CPUs need not imply only four htslib threads.',
                'Timing includes the complete Python driver tree. Echtvar includes bcftools slicing/BGZF into sealed memfd slices and four chromosome tasks; immkv includes reading, lookup, formatting and compressed output.',
                'Compressed output bytes are discarded during timing; memfd input slices are not stored on the tested medium. These runs do not measure full-output gzip disk-write performance.',
                'Fresh echtvar decompressed-row proof and immkv full-field verification are separate correctness phases; historical output references are not current timing samples.',
                'The release annotate source and recorded binary hashes identify the unchanged implementation and inputs; no current binary is executed or rehashed by this summary.',
                'Cache state is uncontrolled. Three paired raw values and their range describe observed spread only; no p-value or statistical certainty is inferred.',
                'read_gib_rounded is rounded tool-internal accounting, not physical I/O. GNU time filesystem_inputs_blocks is a separate kernel filesystem-input block counter; neither is a cgroup io.stat byte measurement.',
                'Round engine order is immkv/echtvar, echtvar/immkv, immkv/echtvar (run_matrix.sh). Phase statistics from immkv.log are rounded internal timings and are not additive GNU wall components.'
            ]}


def fmt(value):
    if value is None:
        return '—'
    return str(value) if isinstance(value, int) else f'{value:.6g}'


def markdown(data):
    lines = ['# Retained release matrix summary', '',
             'Six media/field cells; 568 means all selected fields. Values below come only from retained records.', '',
             '| Media | Fields | Engine | Valid rounds | Wall s | User s | System s | CPU % | Max child RSS KiB | Cgroup peak bytes |',
             '|---|---:|---|---:|---:|---:|---:|---:|---:|---:|']
    for cell in data['cells']:
        for engine, entry in cell['engines'].items():
            summary = entry['summary']
            values = [fmt(summary[k]['median']) if summary else '—' for k in METRICS]
            lines.append(f"| {cell['media']} | {cell['fields']} | {engine} | {entry['valid_rounds']}/3 | " + ' | '.join(values) + ' |')
    lines.extend(['', '## Raw rounds and provenance', ''])
    for cell in data['cells']:
        lines.extend([f"### {cell['media']} / {cell['fields']} fields", ''])
        for engine, entry in cell['engines'].items():
            for item in entry['rounds']:
                state = item['status']
                lines.append(f"- {engine} r{item['round']}: {state['state']}; valid={item['valid']}; status source `{state['source']}`.")
                lines.extend(status_markdown(state))
                for metric, ev in item['metrics'].items():
                    lines.append(f"  - {metric}={fmt(ev['value'])}; `{ev['source']}` ({ev.get('field', metric)}).")
                    if ev.get('final_counters_available') is False:
                        lines.append('  - Last observed aggregate peak; final counters unavailable.')
                if item['invalid_reasons'] or item['issues']:
                    lines.append('  - Exclusion: ' + '; '.join(item['invalid_reasons'] + item['issues']) + '.')
                lines.append(f"  - Actual scope constraints: `{item['constraints']['source']}`; confirmed={item['constraints']['valid']}.")
                if engine == 'immkv':
                    for name, ev in item['count_proof'].items():
                        lines.append(f"  - {name}={ev['value']}; `{ev['source']}`.")
                    for name, ev in item['immkv_phase_statistics'].items():
                        lines.append(f"  - Internal {name}={fmt(ev['value'])}; `{ev['source']}`.")
                        if ev.get('note'):
                            lines.append('  - ' + ev['note'])
                else:
                    proof = item['count_proof']
                    lines.append(f"  - All chromosome statuses successful={proof.get('all_22_chromosomes_ok')}; `{proof.get('source')}`.")
                    for row in proof.get('annotation_log_statistics', []):
                        lines.append(f"  - chr{row['chromosome']} evaluated={row['evaluated']['value']}, wrote={row['wrote']['value']}; `{row['evaluated']['source']}`.")
            if entry['summary']:
                for name, metric in entry['summary'].items():
                    sources = ', '.join(f"`{x['source']}`" for x in metric['sources'])
                    lines.append(f"- {engine} {name}: median={fmt(metric['median'])}, range={fmt(metric['minimum'])}–{fmt(metric['maximum'])}; sources {sources}.")
        ratios = cell['paired_wall_ratios']
        if ratios:
            sources = '; '.join(f"r{x['round']}: `{x['numerator']['source']}` / `{x['denominator']['source']}`" for x in ratios['rounds'])
            values = ', '.join(fmt(x['value']) for x in ratios['rounds'])
            lines.append(f"- Paired echtvar/immkv wall ratios: [{values}], median={fmt(ratios['median'])}, range={fmt(ratios['minimum'])}–{fmt(ratios['maximum'])}, range/median={fmt(ratios['range_over_median'])}; {sources}.")
        else:
            lines.append('- Paired wall ratios unavailable: both engines require all three valid rounds.')
        for name, proof in cell['correctness'].items():
            lines.append(f"- Correctness {name}: scope={proof['status']['state']}, passed={proof['passed']}; `{proof['source']}`; scope source `{proof['status']['source']}`.")
            lines.extend(status_markdown(proof['status']))
            peak = proof['cgroup_peak']
            if peak:
                lines.append(f"  - Aggregate cgroup peak bytes={peak['value']}; `{peak['source']}`; final counters available={peak['final_counters_available']}.")
            # Keep every exception counter, row/field count, stream digest and source
            # visible even when optional verifier JSON fields evolve.
            if proof['evidence'] is not None:
                lines.extend(['', f"Evidence from `{proof['source']}`:", '', '```json',
                              json.dumps(proof['evidence'], indent=2, sort_keys=True), '```', ''])
            for counter, ev in proof['verifier_stdout'].items():
                lines.append(f"- {name} {counter}={ev['value']}; `{ev['source']}`.")
        lines.append('')
    lines.extend(['## Binary, source and input provenance', '', '```json',
                  json.dumps(data['provenance'], indent=2, sort_keys=True), '```', '',
                  '## Interpretation limits', ''])
    lines.extend('- ' + note for note in data['notes'])
    return '\n'.join(lines) + '\n'


def main():
    parser = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    parser.add_argument('--results', type=Path, default=HERE / 'results')
    parser.add_argument('--format', choices=('json', 'markdown', 'both'), default='both')
    parser.add_argument('--output', type=Path, help='New output file inside this experiment; existing paths refused')
    args = parser.parse_args()
    if args.output:
        destination = args.output.resolve()
        if not destination.is_relative_to(HERE):
            parser.error('--output must stay inside the experiment directory')
        if args.output.exists() or args.output.is_symlink():
            parser.error(f'output already exists: {args.output}')
    data = report(args.results.resolve())
    payload = json.dumps(data, indent=2, sort_keys=True, allow_nan=False) + '\n'
    rendered = payload if args.format == 'json' else markdown(data)
    if args.format == 'both':
        rendered += '\n## Machine-readable report\n\n```json\n' + payload + '```\n'
    if args.output:
        try:
            with args.output.open('x') as stream:
                stream.write(rendered)
        except OSError as exc:
            parser.error(str(exc))
    else:
        sys.stdout.write(rendered)
    return 0


if __name__ == '__main__':
    raise SystemExit(main())
