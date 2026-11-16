#!/usr/bin/env python3
"""Resolve a unique scope from persistent counter deltas and unit journals."""
import argparse
import importlib.util
from pathlib import Path


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument('record', type=Path)
    args = parser.parse_args()
    # Explicit sibling import works with python -I without using the working
    # directory as an import source.
    source = Path(__file__).resolve().with_name('env_tools.py')
    spec = importlib.util.spec_from_file_location('cold_experiment_env_tools', source)
    tools = importlib.util.module_from_spec(spec)
    spec.loader.exec_module(tools)
    tools.resolve(tools.local_record(args.record))


if __name__ == '__main__':
    main()
