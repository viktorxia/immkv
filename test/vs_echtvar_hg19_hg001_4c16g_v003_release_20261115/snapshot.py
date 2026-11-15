#!/usr/bin/env python3
"""Snapshot the release with git archive and seed a private offline Go cache."""
import argparse
import hashlib
import io
import json
from pathlib import Path
import re
import shutil
import subprocess
import tarfile


def main():
    p = argparse.ArgumentParser(description=__doc__)
    p.add_argument('--repo', type=Path, required=True)
    p.add_argument('--build', type=Path, required=True)
    p.add_argument('--record', type=Path, required=True)
    p.add_argument('--commit', required=True)
    p.add_argument('--module-cache', type=Path, required=True)
    a = p.parse_args()
    commit = subprocess.check_output(['git', '-C', str(a.repo), 'rev-parse', a.commit + '^{commit}'], text=True).strip()
    if commit != a.commit:
        raise SystemExit('Exact release commit required')
    source = a.build / 'source'
    source.mkdir(parents=True)
    archive = subprocess.check_output(['git', '-C', str(a.repo), 'archive', '--format=tar', commit])
    with (a.build / 'source.tar').open('xb') as f:
        f.write(archive)
    with tarfile.open(fileobj=io.BytesIO(archive)) as tar:
        tar.extractall(source, filter='data')
    # The archive contains the standard test/annotate tool unchanged.
    hashes = {str(p.relative_to(source)): hashlib.sha256(p.read_bytes()).hexdigest()
              for p in sorted(source.rglob('*')) if p.is_file()}
    with (a.record / 'sources.sha256').open('x') as f:
        for name, digest in hashes.items():
            f.write(f'{digest}  {name}\n')
    dependencies = re.findall(r'(?m)^\s*(?:require\s+)?([\w./-]+)\s+(v\S+)', (source / 'go.mod').read_text())
    for module, version in dependencies:
        rel = module + '@' + version
        origin, target = a.module_cache / rel, a.build / 'cache/mod' / rel
        if not origin.is_dir():
            raise SystemExit(f'Missing cached dependency: {origin}')
        target.parent.mkdir(parents=True, exist_ok=True)
        shutil.copytree(origin, target)
        for suffix in ('.mod', '.info', '.ziphash', '.zip'):
            relmeta = Path('cache/download') / module / '@v' / (version + suffix)
            originmeta, targetmeta = a.module_cache / relmeta, a.build / 'cache/mod' / relmeta
            if originmeta.exists():
                targetmeta.parent.mkdir(parents=True, exist_ok=True)
                shutil.copyfile(originmeta, targetmeta)
    facts = {'commit': commit, 'archive_sha256': hashlib.sha256(archive).hexdigest(),
             'source_files': hashes, 'annotate': 'Unchanged test/annotate from release git archive',
             'module_cache_origin': str(a.module_cache), 'dependencies': dependencies}
    with (a.record / 'source-facts.json').open('x') as f:
        json.dump(facts, f, indent=2, sort_keys=True)
        f.write('\n')
    print(f'Snapshotted {commit} to {source}')


if __name__ == '__main__':
    main()
