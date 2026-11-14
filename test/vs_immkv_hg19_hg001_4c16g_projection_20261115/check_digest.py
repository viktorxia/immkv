#!/usr/bin/env python3
"""Check raw-record diagnostics without claiming legacy subset bytes must match."""
import argparse
import json
from pathlib import Path


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--tag", required=True)
    parser.add_argument("--workloads", nargs="+", required=True)
    parser.add_argument("--versions", nargs="+", choices=("old", "before", "after"), required=True)
    args = parser.parse_args()
    results = Path(__file__).resolve().parent / "results"
    output = {}
    for cell in args.workloads:
        records = {}
        for version in args.versions:
            root = results / f"digest.{args.tag}.{cell}.{version}"
            if not (root / "finished").exists() or (root / "exit-status").read_text().strip() != "0":
                raise ValueError(f"Diagnostic did not complete: {root}")
            data = json.loads((root / "stdout").read_text())
            if not data["digest_enabled"] or data["profiled"] or not data["passes"]:
                raise ValueError(f"Not an unprofiled digest run: {root}")
            first = data["passes"][0]
            identity = {key: first[key] for key in ("raw_sha256", "hits", "batches", "returned_bytes")}
            if len(identity["raw_sha256"]) != 64:
                raise ValueError("Missing complete raw-record digest")
            for item in data["passes"]:
                if any(item[key] != value for key, value in identity.items()):
                    raise ValueError(f"Inconsistent repeat: {root}")
            if data["repeat"] != len(data["passes"]) or data["total_keys"] != data["keys"] * data["repeat"]:
                raise ValueError(f"Inconsistent query counts: {root}")
            records[version] = {"keys": data["keys"], **identity, "source": str(root.relative_to(results.parent) / "stdout")}
        signatures = {(item["keys"], item["hits"], item["batches"]) for item in records.values()}
        if len(signatures) != 1:
            raise ValueError(f"Query or hit counts differ: {cell}")
        exact_versions = [v for v in args.versions if not (v == "old" and cell in ("full17", "full540"))]
        raw_signatures = {(records[v]["raw_sha256"], records[v]["returned_bytes"]) for v in exact_versions}
        if len(raw_signatures) > 1:
            raise ValueError(f"Raw records differ between {exact_versions}: {cell}")
        output[cell] = {
            "records": records, "exact_comparison_versions": exact_versions,
            "exact_raw_records_equal": True if len(exact_versions) >= 2 else None,
            "legacy_subset_comparison_excluded": "old" in records and cell in ("full17", "full540"),
        }
    print(json.dumps(output, indent=2, sort_keys=True))


if __name__ == "__main__":
    main()
