#!/usr/bin/env python3
"""Summarize the completed resident diagnostic without mixing disk results."""
import json
from pathlib import Path
import re
import statistics

ROOT = Path(__file__).resolve().parent
RUN = ROOT / "results" / "memory.chr22"
CELLS = ("narrow17", "full17", "full540", "full568", "full568_default")
VERSIONS = ("old", "before", "after")


def process_time(path):
    out = {}
    for line in path.read_text().splitlines():
        for label, key in (("User time (seconds):", "user_seconds"),
                           ("System time (seconds):", "system_seconds"),
                           ("Maximum resident set size (kbytes):", "rss_kib"),
                           ("Exit status:", "exit_status")):
            if line.strip().startswith(label):
                out[key] = float(line.strip()[len(label):])
    if out.get("exit_status") != 0:
        raise ValueError(f"Unsuccessful process: {path}")
    return out


def read_run(kind, cell, version, repeat):
    path = RUN / f"{kind}.{cell}.{version}.r{repeat}"
    if not (path / "finished").exists() or (path / "exit-status").read_text().strip() != "0":
        raise ValueError(f"Incomplete run: {path}")
    constraints = set((path / "cgroup.txt").read_text().splitlines())
    if not {"memory.max=17179869184", "memory.swap.max=0", "cpus=0,1,2,3", "GOMAXPROCS=4"}.issubset(constraints):
        raise ValueError(f"Wrong constraints: {path}")
    item = {"source": str(path.relative_to(ROOT)), **process_time(path / "time.txt")}
    text = (path / "stdout").read_text()
    if kind == "query":
        data = json.loads(text)
        if (data["keys"] != 49536 or data["repeat"] != 10 or data["hits"] != 487090
                or data["profiled"] or data["digest_enabled"]):
            raise ValueError(f"Wrong query scope/counts: {path}")
        item["seconds"] = data["lookup_seconds"] / data["repeat"]
        item["user_seconds_per_pass"] = item["user_seconds"] / data["repeat"]
    else:
        count = re.search(r"rows (\d+), hits (\d+)", text)
        wall = re.search(r"; wall ([0-9.]+)(ms|s)", text)
        if not count or not wall or tuple(map(int, count.groups())) != (49540, 48709):
            raise ValueError(f"Wrong annotation scope/counts: {path}")
        item["seconds"] = float(wall[1]) / (1000 if wall[2] == "ms" else 1)
    return item


def summarize(kind, cells, versions):
    result = {}
    for cell in cells:
        entry = {}
        for version in versions:
            runs = [read_run(kind, cell, version, repeat) for repeat in (1, 2, 3)]
            entry[version] = {
                "runs": runs,
                "median_seconds": statistics.median(r["seconds"] for r in runs),
                "min_seconds": min(r["seconds"] for r in runs),
                "max_seconds": max(r["seconds"] for r in runs),
                "median_process_user_seconds": statistics.median(r["user_seconds"] for r in runs),
                "median_process_rss_mib": statistics.median(r["rss_kib"] for r in runs) / 1024,
            }
        entry["after_change_vs_old_percent"] = 100 * (entry["after"]["median_seconds"] / entry["old"]["median_seconds"] - 1)
        if "before" in entry:
            entry["after_change_vs_before_percent"] = 100 * (entry["after"]["median_seconds"] / entry["before"]["median_seconds"] - 1)
        result[cell] = entry
    return result


def main():
    if not (RUN / "finished").exists():
        raise ValueError("Resident driver did not finish")
    result = {
        "regime": "memory-resident chr22, 4 physical cores, 16 GiB, swap disabled",
        "query": summarize("query", CELLS, VERSIONS),
        "annotation": summarize("annotation", CELLS[:4], ("old", "after")),
        "inputs": json.loads((RUN / "inputs.json").read_text()),
        "digest_evidence": "results/memory.chr22/digest-comparison.json",
        "cgroup_evidence": "results/memory.chr22/cgroup.txt",
    }
    destination = ROOT / "results" / "summary.memory.chr22.json"
    with destination.open("x") as stream:
        json.dump(result, stream, indent=2, sort_keys=True)
        stream.write("\n")
    print(destination)


if __name__ == "__main__":
    main()
