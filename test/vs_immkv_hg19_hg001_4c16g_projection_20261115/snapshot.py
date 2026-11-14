#!/usr/bin/env python3
"""Freeze one version and record every input; never replace existing evidence."""
import argparse
import hashlib
import json
from pathlib import Path
import re
import shutil
import subprocess


def sha(data):
    return hashlib.sha256(data).hexdigest()


def manifest_digest(hashes):
    # Matches the previous experiment's exact digest serialization.
    return sha(json.dumps(hashes, sort_keys=True).encode())


def write_new(path, data):
    path.parent.mkdir(parents=True, exist_ok=True)
    with path.open("xb") as stream:
        stream.write(data)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("version", choices=("old", "before", "after"))
    parser.add_argument("--repo", type=Path, required=True)
    parser.add_argument("--experiment", type=Path, required=True)
    parser.add_argument("--build", type=Path, required=True)
    parser.add_argument("--frozen", type=Path, required=True)
    parser.add_argument("--facts", type=Path, required=True)
    parser.add_argument("--module-cache", type=Path, required=True)
    parser.add_argument("--old-commit", required=True)
    parser.add_argument("--before-digest", required=True)
    args = parser.parse_args()
    repo, experiment, build = args.repo, args.experiment, args.build
    version = args.version
    destination = build / version
    if destination.exists() or (build / (version + ".partial")).exists():
        raise SystemExit(f"Snapshot exists: {destination}")
    facts = {"version": version, "source_files": {}, "harness_files": {}}

    def git(*cmd):
        return subprocess.check_output(["git", "-C", str(repo), *cmd])

    if version == "after":
        source_paths = sorted(p.name for p in repo.glob("*.go") if not p.name.endswith("_test.go"))
        source_paths += ["go.mod", "go.sum"]
        source = {name: (repo / name).read_bytes() for name in source_paths}
        source["cmd/annotate/main.go"] = (repo / "test/annotate/main.go").read_bytes()
        facts["origin"] = "working production source, explicitly requested after parent readiness"
        facts["head"] = git("rev-parse", "HEAD").decode().strip()
        # Capture provenance beside the snapshot; production comments can be
        # non-English, so source copies and patch stay in scratch, not test/.
        patch = git("diff", "HEAD", "--", "*.go", "go.mod", "go.sum", "test/annotate/main.go")
    else:
        previous = json.loads(args.facts.read_text())
        previous_version = "baseline" if version == "old" else "candidate"
        origin = args.frozen / previous_version
        expected = previous["sources"][previous_version]
        actual = {name: sha((origin / name).read_bytes()) for name in expected}
        if actual != expected:
            raise SystemExit(f"Frozen {previous_version} files no longer match their recorded hashes")
        actual_digest = manifest_digest(actual)
        if actual_digest != previous[previous_version + "_source_digest"]:
            raise SystemExit("Previous manifest digest does not match its per-file hashes")
        if version == "before" and actual_digest != args.before_digest:
            raise SystemExit("Frozen candidate digest differs from the required before digest")
        source_paths = sorted(name for name in expected if "/" not in name)
        source_paths.append("cmd/annotate/main.go")
        source = {name: (origin / name).read_bytes() for name in source_paths}
        if version == "old":
            commit = git("rev-parse", args.old_commit).decode().strip()
            for name, data in source.items():
                git_name = "test/annotate/main.go" if name == "cmd/annotate/main.go" else name
                if data != git("show", commit + ":" + git_name):
                    raise SystemExit(f"Frozen baseline differs from {commit}: {name}")
            facts["commit"] = commit
        facts["origin"] = str(origin)
        facts["frozen_manifest_digest"] = actual_digest
        facts["frozen_manifest_files"] = actual
        facts["frozen_facts_sha256"] = sha(args.facts.read_bytes())
        patch = None

    partial = build / (version + ".partial")
    partial.mkdir()
    for name, data in source.items():
        write_new(partial / name, data)
        facts["source_files"][name] = sha(data)
    for path in sorted((experiment / "query-only").glob("*.go")):
        data = path.read_bytes()
        facts["harness_files"]["query-only/" + path.name] = sha(data)
        if version == "old":
            data = data.replace(b"immkv.StreamOptions{", b"immkv.Options{")
        name = "cmd/query-only/" + path.name
        write_new(partial / name, data)
        facts["source_files"][name] = sha(data)
    verifier = (repo / "test/verifyvalues/main.go").read_bytes()
    shared_verifier = build / "shared/verifyvalues.go"
    if shared_verifier.exists():
        if verifier != shared_verifier.read_bytes():
            raise SystemExit("Shared verifier changed since the first build")
    else:
        write_new(shared_verifier, verifier)
    facts["verifyvalues_sha256"] = sha(verifier)
    write_new(partial / "cmd/verifyvalues/main.go", verifier)
    facts["source_files"]["cmd/verifyvalues/main.go"] = sha(verifier)
    facts["source_digest"] = manifest_digest(facts["source_files"])
    facts["query_harness_adaptation"] = (
        "Only immkv.StreamOptions is renamed to immkv.Options" if version == "old" else "none"
    )

    # Seed the experiment-local module cache by reading the existing cache.
    # Missing dependencies fail rather than downloading or writing elsewhere.
    mod_text = source["go.mod"].decode()
    dependencies = re.findall(r"(?m)^\s*(?:require\s+)?([\w./-]+)\s+(v\S+)", mod_text)
    for module, version_number in dependencies:
        relative = module + "@" + version_number
        origin = args.module_cache / relative
        target = build / "cache/mod" / relative
        if not origin.is_dir():
            raise SystemExit(f"Required cached dependency is missing: {origin}")
        if not target.exists():
            target.parent.mkdir(parents=True, exist_ok=True)
            shutil.copytree(origin, target)
        download_relative = Path("cache/download") / module / "@v"
        for suffix in (".mod", ".info", ".ziphash", ".zip"):
            origin_metadata = args.module_cache / download_relative / (version_number + suffix)
            target_metadata = build / "cache/mod" / download_relative / (version_number + suffix)
            if origin_metadata.exists() and not target_metadata.exists():
                target_metadata.parent.mkdir(parents=True, exist_ok=True)
                shutil.copyfile(origin_metadata, target_metadata)
    if patch is not None:
        write_new(build / "after.patch", patch)
    destination.parent.mkdir(parents=True, exist_ok=True)
    partial.rename(destination)
    record = experiment / "results" / ("build." + args.version)
    write_new(record / "source-facts.json", (json.dumps(facts, indent=2, sort_keys=True) + "\n").encode())
    write_new(record / "sources.sha256", "".join(
        f"{digest}  {name}\n" for name, digest in sorted(facts["source_files"].items())
    ).encode())
    print(f"Snapshotted {args.version}: {facts['source_digest']}")
    if "frozen_manifest_digest" in facts:
        print(f"Verified frozen manifest: {facts['frozen_manifest_digest']}")


if __name__ == "__main__":
    main()
