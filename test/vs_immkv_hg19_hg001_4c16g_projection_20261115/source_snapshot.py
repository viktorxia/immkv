#!/usr/bin/env python3
"""Retain or reconstruct exact source bytes using an ASCII JSON snapshot."""
import argparse
import hashlib
import json
from pathlib import Path


def sha(data):
    return hashlib.sha256(data).hexdigest()


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("operation", choices=("retain", "restore"))
    parser.add_argument("version", choices=("old", "before", "after"))
    args = parser.parse_args()
    experiment = Path(__file__).resolve().parent
    repo = experiment.parent.parent
    build = repo / "tmp" / experiment.name
    record = experiment / "results" / ("build." + args.version)
    facts = json.loads((record / "source-facts.json").read_text())
    hashes = facts["source_files"]
    bundle_path = record / "sources.json"
    if args.operation == "retain":
        if bundle_path.exists():
            raise SystemExit(f"Source snapshot exists: {bundle_path}")
        source = build / args.version
        contents = {}
        for name, expected in hashes.items():
            data = (source / name).read_bytes()
            if sha(data) != expected:
                raise SystemExit(f"Source changed: {source / name}")
            contents[name] = data.decode("utf-8")
        bundle = {"format": "utf8-source-files-v1", "source_digest": facts["source_digest"], "files": contents}
        partial = bundle_path.with_suffix(".json.partial")
        with partial.open("x", encoding="ascii") as stream:
            json.dump(bundle, stream, ensure_ascii=True, sort_keys=True, indent=2)
            stream.write("\n")
        # The build record is exclusively reserved; never replace its snapshot.
        partial.rename(bundle_path)
        with (record / "sources-json.sha256").open("x") as stream:
            stream.write(sha(bundle_path.read_bytes()) + "  sources.json\n")
        print(f"Retained exact source bytes: {bundle_path}")
    else:
        bundle = json.loads(bundle_path.read_text(encoding="ascii"))
        if bundle["format"] != "utf8-source-files-v1" or bundle["source_digest"] != facts["source_digest"]:
            raise SystemExit("Source snapshot identity mismatch")
        if set(bundle["files"]) != set(hashes):
            raise SystemExit("Source snapshot file list mismatch")
        destination = build / "restored" / args.version
        if destination.exists():
            raise SystemExit(f"Restore target exists: {destination}")
        partial = destination.with_name(args.version + ".partial")
        partial.mkdir(parents=True)
        for name, text in bundle["files"].items():
            relative = Path(name)
            if relative.is_absolute() or ".." in relative.parts:
                raise SystemExit("Unsafe source name")
            data = text.encode("utf-8")
            if sha(data) != hashes[name]:
                raise SystemExit(f"Source checksum mismatch: {name}")
            target = partial / relative
            target.parent.mkdir(parents=True, exist_ok=True)
            with target.open("xb") as stream:
                stream.write(data)
        partial.rename(destination)
        print(f"Restored verified source bytes: {destination}")


if __name__ == "__main__":
    main()
