#!/bin/bash
# Snapshot both versions without changing the user's checkout.
set -euo pipefail
source "$(dirname "$0")/common.sh"
if [[ -e "$BUILD_ROOT" ]]; then
    printf 'Build root already exists: %s\n' "$BUILD_ROOT" >&2
    exit 1
fi
mkdir -p "$BUILD_ROOT" "$RESULTS"
python3 -I - "$REPO" "$EXP_DIR" "$BUILD_ROOT" "$BASELINE" <<'PY'
from pathlib import Path
import hashlib
import json
import subprocess
import sys
repo, experiment, build, baseline = Path(sys.argv[1]), Path(sys.argv[2]), Path(sys.argv[3]), sys.argv[4]
def git(*args):
    return subprocess.check_output(['git', '-C', str(repo), *args])
def old(name):
    return git('show', baseline + ':' + name)
names = [p.name for p in repo.glob('*.go') if not p.name.endswith('_test.go')]
names += ['go.mod', 'go.sum']
old_names = git('ls-tree', '--name-only', baseline).decode().splitlines()
old_names = [n for n in old_names if n.endswith('.go') and not n.endswith('_test.go') or n in ('go.mod', 'go.sum')]
facts = {'baseline': git('rev-parse', baseline).decode().strip(),
         'candidate_head': git('rev-parse', 'HEAD').decode().strip(), 'sources': {}}
for version in ('baseline', 'candidate'):
    root = build / version
    root.mkdir()
    hashes = {}
    source_names = old_names if version == 'baseline' else names
    for name in source_names:
        content = old(name) if version == 'baseline' else (repo / name).read_bytes()
        (root / name).write_bytes(content)
        hashes[name] = hashlib.sha256(content).hexdigest()
    for tool in ('annotate', 'pointlookup'):
        name = 'test/' + tool + '/main.go'
        content = old(name) if version == 'baseline' else (repo / name).read_bytes()
        path = root / 'cmd' / tool / 'main.go'
        path.parent.mkdir(parents=True)
        path.write_bytes(content)
        hashes['cmd/' + tool + '/main.go'] = hashlib.sha256(content).hexdigest()
    for tool in ('query-only', 'build-check'):
        content = (experiment / tool / 'main.go').read_text()
        if version == 'baseline':
            content = content.replace('immkv.StreamOptions{', 'immkv.Options{')
            content = content.replace('immkv.CreateOptions{', 'immkv.Options{')
            content = content.replace('got, found, err := db.Get(key[:])\n\t\t\tif err != nil {\n\t\t\t\treturn err\n\t\t\t}', 'got, found := db.Get(key[:])')
        path = root / 'cmd' / tool / 'main.go'
        path.parent.mkdir(parents=True)
        path.write_text(content)
        hashes['cmd/' + tool + '/main.go'] = hashlib.sha256(content.encode()).hexdigest()
    facts['sources'][version] = hashes
    digest = hashlib.sha256(json.dumps(hashes, sort_keys=True).encode()).hexdigest()
    facts[version + '_source_digest'] = digest
(experiment / 'results' / 'source-facts.json').write_text(json.dumps(facts, indent=2) + '\n')
(experiment / 'results' / 'candidate.patch').write_bytes(git('diff', 'HEAD', '--', '*.go', 'README.md', 'README.cn.md', 'docs/ingest-design.md'))
PY
for version in baseline candidate; do
    mkdir -p "$BUILD_ROOT/bin/$version"
    for tool in annotate pointlookup query-only build-check; do
        go -C "$BUILD_ROOT/$version" build -trimpath -o "$BUILD_ROOT/bin/$version/$tool" "./cmd/$tool"
    done
done
go build -trimpath -o "$BUILD_ROOT/bin/verifyvalues" "$REPO/test/verifyvalues"
{
    date -Is
    go version
    go env CGO_ENABLED GOOS GOARCH
    uname -a
    lscpu -e=CPU,CORE,SOCKET,ONLINE
    df -h "$REPO" "$FULL_LIB" "$NARROW_LIB"
    sha256sum "$QUERY"
    sha256sum "$(fields_file 17)" "$(fields_file 540)" "$(fields_file 568)"
    git -C "$REPO" status --short
} > "$RESULTS/environment.txt"
printf 'Built baseline and candidate in %s\n' "$BUILD_ROOT"
