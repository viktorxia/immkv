#!/usr/bin/env python3
"""Atomically retain a completed profile without replacing an existing file."""
import argparse
import hashlib
import os
from pathlib import Path


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("version", choices=("before", "after"))
    parser.add_argument("tag")
    args = parser.parse_args()
    if not args.tag or any(c not in "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-" for c in args.tag):
        parser.error("Invalid tag")
    experiment = Path(__file__).resolve().parent
    scratch = experiment.parent.parent / "tmp" / experiment.name
    label = f"profile.{args.tag}.full540.{args.version}"
    record = experiment / "results" / label
    if not (record / "finished").exists():
        raise SystemExit("Profile process has not finished successfully")
    origin = scratch / "profiles" / label / "cpu.pprof"
    target = record / "cpu.pprof"
    data = origin.read_bytes()
    if not data:
        raise SystemExit("Empty profile")
    if target.exists():
        if data != target.read_bytes():
            raise SystemExit("Existing profile differs from scratch")
        print(f"Profile already retained: {target}")
        return
    temporary = target.with_name(target.name + ".partial")
    with temporary.open("xb") as output:
        output.write(data)
    os.link(temporary, target)
    temporary.unlink()
    with (record / "cpu-profile.sha256").open("x") as output:
        output.write(hashlib.sha256(data).hexdigest() + "  cpu.pprof\n")
    print(f"Profile retained: {target}")


if __name__ == "__main__":
    main()
