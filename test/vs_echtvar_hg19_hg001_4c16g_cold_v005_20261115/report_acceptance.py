#!/usr/bin/env python3
"""Audit cold-v005 logs; create an exclusive final report or an incomplete snapshot."""
import argparse
from datetime import datetime, timezone
import hashlib
import json
from pathlib import Path
import re
import shlex
import statistics
import sys

TAG, ROWS = 'cold-v005', 3933714
COMMIT = 'd2eac357f3fc126755b43ff8b44a64dc55bc0ec2'
ENGINES, CELLS = ('immkv', 'echtvar'), [(m, f) for m in ('nvme', 'hdd') for f in (17, 540, 568)]
REQUIRED = ('memory.max', 'memory.swap.max', 'memory.events', 'memory.peak', 'memory.current', 'cpu.stat')


def text(path):
    try:
        return path.read_text()
    except FileNotFoundError:
        return ''


def load(path):
    return json.loads(text(path) or '{}')


def counters(value):
    return {k: int(v) for k, v in (line.split() for line in value.splitlines())}


def time_values(value):
    result = {}
    for line in value.splitlines():
        line = line.strip()
        if line.startswith('Elapsed (wall clock) time'):
            wall = 0.0
            for part in line.split('): ', 1)[1].split(':'):
                wall = wall * 60 + float(part)
            result['wall_seconds'] = wall
        for prefix, key, cast in [('User time (seconds): ', 'user_seconds', float),
                                  ('System time (seconds): ', 'system_seconds', float),
                                  ('Maximum resident set size (kbytes): ', 'max_child_rss_kib', int),
                                  ('File system inputs: ', 'input_blocks', int), ('Exit status: ', 'exit_code', int)]:
            if line.startswith(prefix):
                result[key] = cast(line[len(prefix):])
    if 'user_seconds' in result and 'system_seconds' in result:
        result['cpu_seconds'] = result['user_seconds'] + result['system_seconds']
    return result


def scope(path, tools, timing=False, engine=None, media=None, fields=None):
    status, cell = load(path / 'status.json'), load(path / 'cell.json')
    cache, limits, before = [load(path / n) for n in ('cache-clear.json', 'constraints.json', 'cgroup-before.json')]
    ready, final, done, observer = [load(path / n) for n in ('observer-ready.json', 'observer-workload-final.json', 'workload-done.json', 'cgroup-observer-final.json')]
    errors = []
    def require(ok, message):
        if not ok:
            errors.append(message)
    require(cache.get('status') == 'ok' and cache.get('response', {}).get('exit_status') == 0
            and cache.get('response', {}).get('effective_uid') == 0, 'cache clear not proven')
    require(limits.get('memory.max') == '17179869184' and limits.get('memory.swap.max') == '0'
            and limits.get('cpus') == [0, 1, 2, 3] and limits.get('GOMAXPROCS') == '4'
            and len({(x['physical_package_id'], x['core_id']) for x in limits.get('cpu_topology', [])}) == 4, 'resource limits not proven')
    require(all(k in before and k in ready.get('opened', []) for k in REQUIRED)
            and ready.get('persistent_descriptors') is True and ready.get('cgroup') == limits.get('cgroup'), 'observer baseline not proven')
    start = cell.get('started', {}).get('time_ns')
    children = cell.get('chromosomes') or [load(p) for p in sorted(path.glob('chr*.json'))]
    unit = text(path / 'scope-unit').strip()
    journal_sources = status.get('journal_sources', [])
    journal_oom = any(unit and unit in text(path / name) and any(marker in text(path / name) for marker in
                      ("Failed with result 'oom-kill'", 'killed by the OOM killer')) for name in journal_sources)
    oom_delta = status.get('cgroup_delta', {}).get('memory.events', {})
    metadata_lost_after_oom = bool(not timing and engine == 'echtvar' and fields == 568
                                  and status.get('status') == 'oom' and not cell and not children
                                  and journal_oom and any(oom_delta.get(k, 0) > 0 for k in ('oom_kill', 'oom_group_kill')))
    if metadata_lost_after_oom:
        argv = shlex.split(text(path / 'command.sh'))
        expected = ['python3', '-I', str(path.parents[1] / 'run_cell.py'), '--engine', 'echtvar',
                    '--media', media, '--fields', str(fields), '--record', str(path), '--verify-reference']
        launch_valid = (argv[:len(expected)] == expected and len(argv) == len(expected) + 3
                        and argv[-2:] == ['--expected-rows', str(ROWS)])
        launch_text = ' '.join(argv)
        require(launch_valid and launch_text in text(path / 'time-scope.txt')
                and f'--unit={unit.removesuffix(".scope")}' in text(path / 'time-scope.txt')
                and any(unit in text(path / name) and launch_text in text(path / name) for name in journal_sources),
                'OOM verification wrapper launch identity invalid')
    if not start:
        start = min((x.get('started', {}).get('time_ns', 0) for x in children if x.get('started')), default=0)
    boundary = ready.get('timestamp_ns', 0)
    require(0 < cache.get('requested_ns', 0) <= cache.get('recorded_ns', 0)
            <= limits.get('timestamp_ns', 0) <= before.get('timestamp_ns', 0) <= boundary,
            'cache/scope start boundary invalid')
    require(bool(text(path / 'started')) and bool(text(path / 'command.sh')), 'workload launch missing')
    if timing or engine:
        if not metadata_lost_after_oom:
            require(start > boundary, 'engine start timestamp missing or before observer readiness')
        if cell:
            require((cell.get('engine'), cell.get('media'), cell.get('fields')) == (engine, media, fields)
                    and cell.get('source_commit') == COMMIT and cell.get('compressed_output') is True,
                    'workload identity invalid')
            require(cell.get('timing_sample') is timing, 'workload timing_sample invalid')
        commands = ([('annotate', cell.get('command', []))] if engine == 'immkv' else
                    [(name, child.get(key, [])) for child in children for name, key in
                     (('echtvar', 'annotation_command'), ('bcftools', 'slice_command')) if child.get(key)])
        if not metadata_lost_after_oom:
            require(bool(commands) and all(cmd and cmd[0] == tools.get(name, {}).get('identity', {}).get('path')
                                          for name, cmd in commands), 'recorded binary command identity invalid')
    snapshots = [load(path / n) for n in ('cgroup-final.json', 'observer-workload-final.json', 'memory-observer.json')]
    snapshots.append(observer.get('counters', {}))
    event_oom = any(counters(s.get('memory.events', '')).get(k, 0) > counters(before.get('memory.events', '')).get(k, 0)
                    for s in snapshots for k in ('oom_kill', 'oom_group_kill')) if before else False
    outcome = status.get('status', 'incomplete')
    if outcome == 'oom':
        require(event_oom or journal_oom, 'OOM lacks recorded counter/journal proof')
    else:
        require(outcome == 'ok', 'scope unsuccessful or incomplete')
        require(status.get('exit_code') == status.get('scope_exit_code') == 0
                and text(path / 'exit-status').strip() == text(path / 'scope-exit-status').strip() == '0'
                and text(path / 'observer-exit-status').strip() == '0', 'successful exits missing')
        require(status.get('final_workload_handshake') is True and observer.get('workload_final_handshake') is True
                and all(k in final and k not in final.get('read_errors', {}) for k in REQUIRED)
                and done.get('exit_code') == 0 and 0 < final.get('timestamp_ns', 0)
                <= int(text(path / 'observer-final-ack') or '0') <= status.get('finished_ns', 0)
                and 0 < done.get('timestamp_ns', 0) <= int(text(path / 'observer-final-ack') or '0')
                and any(s.get('timestamp_ns', 0) >= done.get('timestamp_ns', 0)
                        and all(k in s for k in REQUIRED) for s in
                        (load(path / 'cgroup-final.json'), observer.get('counters', {}))), 'final observer handshake/counters missing')
        require(not event_oom and not journal_oom and not status.get('timeout_detected'), 'success conflicts with OOM/timeout evidence')
        if cell:
            require(cell.get('status') == 'ok' and cell.get('exit_code') == 0
                    and boundary < start <= cell.get('finished', {}).get('time_ns', 0) <= done.get('timestamp_ns', 0), 'engine completion invalid')
        if engine == 'echtvar':
            require(len(children) == 22 and {x.get('chromosome') for x in children} == set(range(1, 23))
                    and cell.get('workers') == 4 and all(x.get('status') == 'ok' and x.get('slice_exit_code') == x.get('annotation_exit_code') == 0
                    and start <= x.get('slice_started', {}).get('time_ns', 0) <= x.get('slice_finished', {}).get('time_ns', 0)
                    <= x.get('annotation_started', {}).get('time_ns', 0) <= x.get('annotation_finished', {}).get('time_ns', 0)
                    <= cell.get('finished', {}).get('time_ns', 0) for x in children), 'chromosome attempts/exits/timestamps incomplete')
    source = 'time.txt' if text(path / 'time.txt') else 'time-scope.txt'
    metrics = time_values(text(path / source))
    peak = max((int(s.get('memory.peak', 0)) for s in snapshots), default=0)
    metrics['cgroup_peak_bytes'] = peak
    require(peak > 0 and peak == status.get('memory_peak_bytes'), 'cgroup peak evidence missing or inconsistent')
    require(all(k in metrics for k in ('wall_seconds', 'cpu_seconds', 'max_child_rss_kib', 'input_blocks', 'exit_code')), 'GNU time metrics missing')
    if outcome == 'ok':
        require(source == 'time.txt' and metrics.get('exit_code') == 0, 'workload GNU time did not succeed')
    return dict(record=path.name, outcome=outcome, errors=errors, start_ns=start, boundary_ns=boundary,
                end_ns=status.get('finished_ns', 0), metrics=metrics, time_source=source,
                accepted_timing_sample=timing and outcome == 'ok' and not errors,
                launch_evidence='wrapper launch gated after observer readiness; exact engine timestamp and direct child argv unavailable'
                    if metadata_lost_after_oom else 'recorded workload timestamps and child argv',
                metadata_lost_after_oom=metadata_lost_after_oom,
                oom_counter_proven=event_oom, oom_journal_proven=journal_oom,
                status=status, cell=cell, sources=[f'{path.name}/{n}' for n in
                ('status.json', 'cell.json', 'cache-clear.json', 'constraints.json', 'cgroup-before.json', 'observer-ready.json',
                 'observer-workload-final.json', 'cgroup-observer-final.json', 'workload-done.json', 'observer-final-ack', source)
                if text(path / n)] + [f'{path.name}/{n}' for n in journal_sources])


def verification(root, media, fields, tools, manifest):
    p, e = root / f'verify-{media}-{fields}-{TAG}', root / f'verify-echtvar-{media}-{fields}-{TAG}'
    proof, fresh = load(p / 'verification.json'), load(e / 'cell.json')
    runs = [scope(p, tools), scope(e, tools, engine='echtvar', media=media, fields=fields)]
    errors = [f"{r['record']}: {message}" for r in runs for message in r['errors']]
    c, a = proof.get('comparison', {}), proof.get('annotate_statistics', {})
    if not (proof.get('status') == 'passed' and (proof.get('media'), proof.get('fields')) == (media, fields)
            and proof.get('timing_sample') is False and proof.get('selected_fields') == fields
            and (c.get('rows'), c.get('field_comparisons'), c.get('value_mismatches'), c.get('hit_disagreements')) == (ROWS, ROWS * fields, 0, 0)
            and all(isinstance(c.get(k), int) and c[k] >= 0 for k in ('misses_agreed', 'echtvar_i32_overflow_losses', 'float_comparisons'))
            and proof.get('float_tolerance') == 1.1e-6 and 0 <= c.get('max_float_difference', -1) <= 1.1e-6
            and a.get('rows') == ROWS and isinstance(a.get('hits'), int) and 0 <= a['hits'] <= ROWS
            and proof.get('exit_status') == {'annotate': 0, 'verifier': 0}
            and proof.get('frozen_binary_sha256') == manifest
            and all(proof.get('commands', {}).get(command, [None])[0] == tools.get(tool, {}).get('identity', {}).get('path')
                    for command, tool in (('annotate', 'annotate'), ('verifier', 'verifyvalues')))):
        errors.append('full comparison totals, exits or binary hashes invalid')
    for name in ENGINES:
        stream = proof.get('streams', {}).get(name, {})
        files = stream.get('files', [])
        if not (stream.get('status') == 'complete' and stream.get('rows') == ROWS and stream.get('gzip_crc_eof_ok')
                and len(files) == (1 if name == 'immkv' else 22) and sum(x.get('rows', 0) for x in files) == ROWS
                and all(x.get('gzip_crc_eof_ok') for x in files)):
            errors.append(f'{name} complete gzip stream proof missing')
    provenance = proof.get('reference_provenance', {})
    fallback = runs[1]['outcome'] == 'oom' and fields == 568
    if runs[1].get('metadata_lost_after_oom'):
        argv = shlex.split(text(e / 'command.sh'))
        if argv[argv.index('--verify-reference') + 1] != proof.get('reference'):
            errors.append('OOM wrapper reference does not match full comparison provenance')
    evidence = e / ('status.json' if fallback else 'cell.json')
    if not (provenance.get('path') == str(evidence) and provenance.get('sha256') == hashlib.sha256(text(evidence).encode()).hexdigest()
            and provenance.get('fresh_equivalence_proven') is (not fallback)
            and proof.get('reference_oom_evidence' if fallback else 'reference_evidence') == str(evidence)):
        errors.append('reference provenance invalid')
    if not fallback:
        chromosomes = fresh.get('chromosomes', [])
        if not (runs[1]['outcome'] == 'ok' and fresh.get('fresh_rows') == ROWS and len(chromosomes) == 22
                and fresh.get('verify_reference') == proof.get('reference')
                and sum(x.get('comparison', {}).get('fresh', {}).get('rows', 0) for x in chromosomes) == ROWS
                and all(x.get('comparison', {}).get('equal') is True
                        and x['comparison'].get('fresh') == x['comparison'].get('reference')
                        and x['comparison']['fresh'].get('gzip_crc_eof_ok')
                        and re.fullmatch('[0-9a-f]{64}', x['comparison']['fresh'].get('data_sha256', '')) for x in chromosomes)):
            errors.append('fresh echtvar digest equivalence incomplete')
    return dict(accepted=not errors, errors=errors, reference_mode='historical-only after proven fresh echtvar OOM' if fallback else 'fresh echtvar digest equivalence',
                proof=proof, scopes=runs, source=f'{p.name}/verification.json',
                metadata_note='Generic status.timing_sample may be true; workload verification metadata is false; excluded from timing.')


def collect(root):
    tools = load(root / 'preparation-environment.json').get('tools', {})
    manifest = {name.lstrip('*'): digest for digest, name in (line.split() for line in text(root / 'preparation/binaries.sha256').splitlines())}
    errors = []
    if not (load(root / 'preparation/source-facts.json').get('commit') == COMMIT
            and text(root / 'preparation/exit-status').strip() == '0'
            and set(manifest) == {'annotate', 'verifyvalues'}
            and all(tools.get(n, {}).get('sha256') == h and re.fullmatch('[0-9a-f]{64}', h) for n, h in manifest.items())
            and all(re.fullmatch('[0-9a-f]{64}', tools.get(n, {}).get('sha256', '')) for n in ('echtvar', 'bcftools'))):
        errors.append('frozen preparation and recorded binary identity incomplete')
    expected = {f'timing-{m}-{f}-{e}-r{r}-{TAG}' for m, f in CELLS for e in ENGINES for r in (1, 2, 3)}
    unexpected = sorted(p.name for p in root.glob(f'timing-*-{TAG}') if p.name not in expected)
    if unexpected:
        errors.append('unexpected timing attempts for tag: ' + ', '.join(unexpected))
    cells = []
    for media, fields in CELLS:
        proof = verification(root, media, fields, tools, manifest)
        records = [scope(root / f'timing-{media}-{fields}-{engine}-r{round_}-{TAG}', tools, True, engine, media, fields)
                   for round_ in (1, 2, 3) for engine in (ENGINES[::-1] if round_ == 2 else ENGINES)]
        issues = list(proof['errors'])
        for r in records:
            issues.extend(f"{r['record']}: {message}" for message in r['errors'])
            if r['outcome'] == 'oom' and not (fields == 568 and '-echtvar-' in r['record']):
                issues.append(f"{r['record']}: unexpected OOM")
        for previous, current in zip(records, records[1:]):
            cache_start = load(root / current['record'] / 'cache-clear.json').get('requested_ns', 0)
            if not (0 < previous['start_ns'] < previous['end_ns'] < cache_start < current['boundary_ns'] < current['start_ns']):
                issues.append(f"order/start boundary unproven: {previous['record']} then {current['record']}")
        summaries = {}
        for engine in ENGINES:
            group = [r for r in records if f'-{engine}-' in r['record']]
            good = [r for r in group if r['accepted_timing_sample']]
            summaries[engine] = dict(actual_attempts=sum(r['start_ns'] > 0 for r in group), accepted=len(good),
                cpu_total_seconds=sum(r['metrics'].get('cpu_seconds', 0) for r in good),
                medians={k: statistics.median(r['metrics'][k] for r in good) for k in
                         ('wall_seconds', 'cpu_seconds', 'max_child_rss_kib', 'cgroup_peak_bytes', 'input_blocks')} if good else {})
        cells.append(dict(media=media, fields=fields, complete=not issues, errors=issues, verification=proof, records=records, engines=summaries))
    return dict(tag=TAG, complete=not errors and all(c['complete'] for c in cells), errors=errors, cells=cells,
                generated_utc=datetime.now(timezone.utc).isoformat(), binary_identity=tools,
                identity_basis='Recorded command paths and preparation/verification SHA256 manifests; no live binary or input reads.',
                identity_sources=['preparation-environment.json', 'preparation/source-facts.json', 'preparation/binaries.sha256'])


def link(path):
    return f'[{path}]({path})'


def render(report):
    lines = ['# FINAL REPORT — ' + ('COMPLETE' if report['complete'] else 'INCOMPLETE / PARTIAL'), '',
             'Tag cold-v005; exactly six cells; three actual attempts per engine required; order I/E, E/I, I/E.', '',
             'I = immkv; E = echtvar; 4 physical CPUs, GOMAXPROCS=4, memory.max=17179869184, memory.swap.max=0.', '',
             'Only accepted timing-* workloads enter medians; verification scope timing_sample:true is a generic metadata bug, while workload verification metadata is false.', '',
             'GNU time CPU is user + system; RSS is maximum child RSS, not simultaneous group RSS; cgroup peak includes the scope. Input blocks are GNU time File system inputs, not inferred bytes. Scope-wrapper times are failure evidence only. OOM has no speed ratio.', '',
             'Recorded identity sources: ' + ', '.join(map(link, report['identity_sources'])) + '. No live binaries or inputs were read.', '',
             'Audit limits and OOM metadata exception: ' + link('final-matrix-audit.md') + '; correctness cleanup overlap and unavailable cgroup I/O counters are retained there.', '',
             '| Medium | Fields | I wall seconds r1/r2/r3 | I median | E wall seconds r1/r2/r3 | E median | Accepted cell |',
             '|---|---:|---|---:|---|---:|---|']
    for cell in report['cells']:
        values = []
        for engine in ENGINES:
            runs = [r for r in cell['records'] if f'-{engine}-' in r['record']]
            good = [r['metrics']['wall_seconds'] for r in runs if r['accepted_timing_sample']]
            values.extend([' / '.join(f"{r['metrics'].get('wall_seconds', 'missing')} ({r['outcome']}) " + link(f"{r['record']}/{r['time_source']}") for r in sorted(runs, key=lambda x: x['record'])),
                           str(statistics.median(good)) if len(good) == 3 else 'unavailable'])
        lines.append(f"| {cell['media']} | {cell['fields']} | " + ' | '.join(values) + f" | {cell['complete']} |")
    lines += ['', '## Actual attempts and resource measurements', '',
              '| Source | Outcome | Accepted timing | Wall s | CPU total s | Max child RSS KiB | Cgroup peak bytes | GNU input blocks |',
              '|---|---|---|---:|---:|---:|---:|---:|']
    for cell in report['cells']:
        for r in cell['records']:
            m = r['metrics']
            lines.append('| ' + link(f"{r['record']}/{r['time_source']}") + f"; {link(r['record'] + '/status.json')} | {r['outcome']} | {r['accepted_timing_sample']} | "
                         + ' | '.join(str(m.get(k, 'missing')) for k in ('wall_seconds', 'cpu_seconds', 'max_child_rss_kib', 'cgroup_peak_bytes', 'input_blocks')) + ' |')
    lines += ['', '## Verification, failures and evidence', '']
    for cell in report['cells']:
        v, runs = cell['verification'], cell['records']
        counts = {s: sum(r['outcome'] == s for r in runs) for s in ('ok', 'oom', 'failed', 'timeout', 'incomplete')}
        p = v['proof']
        lines += [f"### {cell['media']} / {cell['fields']} fields", '',
                  f"Timing outcomes (separate categories): {json.dumps(counts)}; correctness accepted: {v['accepted']}; reference: {v['reference_mode']}.", '',
                  f"Actual comparison {link(v['source'])}: {json.dumps(p.get('comparison', {}), sort_keys=True)}; annotate rows/hits: {json.dumps(p.get('annotate_statistics', {}), sort_keys=True)}.", '',
                  'Known categories: ' + p.get('marker_collision_accounting', 'missing') + ' True float -1 candidates: '
                  + str(p.get('streams', {}).get('immkv', {}).get('selected_float_true_minus_one_candidates', 'missing'))
                  + '; overflow losses and float comparisons remain separate in comparison totals; tolerance: '
                  + str(p.get('float_tolerance', 'missing')) + '; ' + p.get('float_tolerance_basis', 'missing') + '.', '']
        for r in runs + v['scopes']:
            lines += [f"- {r['record']}: start_ns={r['start_ns']}, observer_boundary_ns={r['boundary_ns']}, resolved_end_ns={r['end_ns']}; "
                      + r['launch_evidence'] + '; ' + ', '.join(map(link, r['sources'])) + '.']
        lines += ['', 'Acceptance gaps: ' + ('; '.join(cell['errors']) if cell['errors'] else 'none') + '.', '']
    lines += ['Global acceptance gaps: ' + ('; '.join(report['errors']) or 'none') + '.', '']
    return '\n'.join(lines)


def main(argv=None):
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('--results', type=Path, default=Path(__file__).resolve().parent / 'results')
    parser.add_argument('--name', default='acceptance-final', help='Exclusive output filename stem')
    parser.add_argument('--partial', action='store_true', help='Allow an explicitly incomplete snapshot')
    args = parser.parse_args(argv)
    if not re.fullmatch('[A-Za-z0-9_-]+', args.name):
        parser.error('--name must be a simple filename stem')
    root = args.results.resolve()
    if not root.is_relative_to(Path(__file__).resolve().parents[2]):
        parser.error('--results must remain inside this repository')
    report = collect(root)
    if not report['complete'] and not args.partial:
        print('INCOMPLETE: refusing final report; use --partial --name NEW_STEM for an incomplete snapshot.', file=sys.stderr)
        return 1
    paths = [root / (args.name + suffix) for suffix in ('.json', '.md')]
    if any(p.exists() for p in paths):
        parser.error('report already exists; choose a new --name')
    # Reserve both outputs before writing; never overwrite or unlink an existing report.
    with paths[0].open('x') as output_json:
        with paths[1].open('x') as output_md:
            json.dump(report, output_json, indent=2)
            output_json.write('\n')
            output_md.write(render(report))
    print(('COMPLETE' if report['complete'] else 'INCOMPLETE / PARTIAL') + ': ' + str(paths[1]))
    return 0


if __name__ == '__main__':
    raise SystemExit(main())
