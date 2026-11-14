#!/usr/bin/env python3
"""Report complete recorded runs; missing or failed evidence remains explicit."""
import argparse
import json
from pathlib import Path
import re
import statistics

VERSIONS = ("old", "before", "after")
CELLS = ("narrow17", "full17", "full540", "full568", "full568_default")
ROOT = Path(__file__).resolve().parent
RESULTS = ROOT / "results"


def source(path):
    return str(path.relative_to(ROOT))


def time_report(path):
    values = {}
    for line in path.read_text().splitlines():
        text = line.strip()
        for prefix, key, cast in (
            ("User time (seconds):", "user_seconds", float),
            ("System time (seconds):", "system_seconds", float),
            ("Maximum resident set size (kbytes):", "rss_kib", int),
            ("Exit status:", "exit_status", int),
        ):
            if text.startswith(prefix):
                values[key] = cast(text[len(prefix):].strip())
        if text.startswith("Elapsed (wall clock) time"):
            total = 0.0
            for part in text.split("): ", 1)[1].split(":"):
                total = total * 60 + float(part)
            values["wall_seconds"] = total
    if values.get("exit_status") != 0:
        raise ValueError("Time report is not a successful run")
    return values


def completed(path):
    if not path.exists():
        return "pending"
    if not (path / "finished").exists():
        return "incomplete_or_failed"
    if (path / "exit-status").read_text().strip() != "0":
        return "failed"
    limits = (path / "cgroup.txt").read_text().splitlines()
    required = {"memory.max=17179869184", "memory.swap.max=0", "cpus=0,1,2,3", "GOMAXPROCS=4"}
    if not required.issubset(limits):
        return "invalid_constraints"
    return "complete"


def query_runs(tag, cell, version):
    output = {"runs": [], "status": "pending"}
    for repeat in (1, 2, 3):
        path = RESULTS / f"null.{tag}.{cell}.{version}.r{repeat}"
        status = completed(path)
        item = {"round": repeat, "status": status, "source": source(path)}
        if status == "complete":
            try:
                data = json.loads((path / "stdout").read_text())
                if data["profiled"] or data["digest_enabled"] or data["repeat"] < 1:
                    raise ValueError("Profile or digest is not a performance run")
                if data["threads"] != 4 or data["gomaxprocs"] != 4 or data["batch"] != 4096:
                    raise ValueError("Unexpected harness configuration")
                if data["explicit_columns"] != (cell != "full568_default"):
                    raise ValueError("Explicit column mode does not match workload")
                if len(data["passes"]) != data["repeat"] or data["total_keys"] != data["keys"] * data["repeat"]:
                    raise ValueError("Incomplete lookup pass counts")
                item.update(data)
                item.update(time_report(path / "time.txt"))
                item["lookup_seconds_per_pass"] = data["lookup_seconds"] / data["repeat"]
                item["hits_per_pass"] = data["hits"] // data["repeat"]
                item["time_source"] = source(path / "time.txt")
                item["lookup_source"] = source(path / "stdout")
            except (ValueError, KeyError, OSError) as exc:
                item["status"] = "invalid"
                item["error"] = str(exc)
        output["runs"].append(item)
    if all(item["status"] == "complete" for item in output["runs"]):
        identities = {(item["keys"], item["hits_per_pass"], item["repeat"], item["returned_bytes"]) for item in output["runs"]}
        if len(identities) != 1:
            output["status"] = "inconsistent_counts"
        else:
            output["status"] = "complete"
            output["median"] = {
                key: statistics.median(item[key] for item in output["runs"])
                for key in ("lookup_seconds_per_pass", "keys_per_second", "user_seconds", "system_seconds", "rss_kib", "wall_seconds")
            }
    elif any(item["status"] != "pending" for item in output["runs"]):
        output["status"] = "incomplete"
    return output


def profile_analysis(path):
    if not path.exists():
        return {"status": "pending", "source": source(path)}
    text = path.read_text()
    headline = re.search(r"Duration: ([0-9.]+)s, Total samples = ([0-9.]+)s \(([0-9.]+)%\)", text)
    rows = []
    for line in text.splitlines():
        parts = line.split()
        if len(parts) >= 6 and parts[1].endswith("%") and parts[4].endswith("%"):
            try:
                rows.append({
                    "function": " ".join(parts[5:]),
                    "flat_seconds": float(parts[0].removesuffix("s")),
                    "flat_percent": float(parts[1].removesuffix("%")),
                    "cumulative_seconds": float(parts[3].removesuffix("s")),
                    "cumulative_percent": float(parts[4].removesuffix("%")),
                })
            except ValueError:
                continue
    entry = {"status": "recorded", "source": source(path), "functions": rows}
    if headline:
        entry.update(profile_duration_seconds=float(headline[1]), sampled_cpu_seconds=float(headline[2]),
                     sampled_cpu_percent_of_elapsed=float(headline[3]))
    return entry


def report(tag):
    output = {"tag": tag, "versions": {}, "query_only": {}, "verification": {}, "profiles": {}, "annotation": {}, "digests": {},
              "interpretation_limits": [
                  "Lookup time excludes setup; GNU time CPU/RSS include the complete process.",
                  "Profile and digest invocations are diagnostic and excluded from timing medians.",
                  "CPU and memory limits do not fix NVMe service time or OS cache state.",
                  "A wall-time ratio alone does not identify the CPU contribution to an optimization.",
              ]}
    for version in VERSIONS:
        facts_path = RESULTS / f"build.{version}" / "source-facts.json"
        output["versions"][version] = {
            "status": "built" if (facts_path.parent / "finished").exists() else "pending",
            "source": source(facts_path),
        }
        if facts_path.exists():
            facts = json.loads(facts_path.read_text())
            output["versions"][version]["source_digest"] = facts["source_digest"]
            if "frozen_manifest_digest" in facts:
                output["versions"][version]["frozen_manifest_digest"] = facts["frozen_manifest_digest"]
    for cell in CELLS:
        records = {version: query_runs(tag, cell, version) for version in VERSIONS}
        comparisons = {}
        for reference, candidate in (("old", "before"), ("old", "after"), ("before", "after")):
            a, b = records[reference], records[candidate]
            if a["status"] != "complete" or b["status"] != "complete":
                comparisons[reference + "-" + candidate] = {"status": "pending_complete_three_pairs"}
                continue
            for left, right in zip(a["runs"], b["runs"]):
                if any(left[key] != right[key] for key in ("keys", "hits_per_pass", "repeat")):
                    raise ValueError(f"Cross-version query/hit counts differ: {cell}")
            matrix = RESULTS / f"null.{tag}.matrix"
            if not (matrix / "inputs.sha256").exists() or not (matrix / "order.txt").exists():
                comparisons[reference + "-" + candidate] = {"status": "missing_matrix_inputs_or_order"}
                continue
            order = (matrix / "order.txt").read_text().splitlines()
            expected = [f"null.{tag}.{cell}.{v}.r{r}" for r in (1, 2, 3)
                        for v in ((reference, candidate) if r != 2 else (candidate, reference))]
            pair_entries = [name for name in order if name in expected]
            if pair_entries != expected:
                comparisons[reference + "-" + candidate] = {"status": "invalid_paired_order"}
                continue
            deltas = [100 * (right["lookup_seconds_per_pass"] / left["lookup_seconds_per_pass"] - 1)
                      for left, right in zip(a["runs"], b["runs"])]
            comparisons[reference + "-" + candidate] = {
                "status": "complete", "paired_lookup_change_percent": deltas,
                "median_paired_lookup_change_percent": statistics.median(deltas),
                "sources": [r["lookup_source"] for version in (a, b) for r in version["runs"]],
                "input_hash_source": source(matrix / "inputs.sha256"), "order_source": source(matrix / "order.txt"),
            }
        output["query_only"][cell] = {"versions": records, "comparisons": comparisons}
    for path in sorted(RESULTS.glob(f"verify.{tag}.*")):
        entry = {"status": completed(path), "source": source(path)}
        if entry["status"] == "complete":
            evidence = json.loads((path / "verification.json").read_text())
            if not evidence["exact_stream_sha256_equal"] or evidence["value_mismatches"] or evidence["hit_disagreements"]:
                raise ValueError(f"Verification did not pass: {path}")
            entry["evidence"] = evidence
        output["verification"][path.name] = entry
    for version in ("before", "after"):
        path = RESULTS / f"profile.{tag}.full540.{version}"
        entry = {"status": completed(path), "source": source(path), "performance_sample": False}
        if entry["status"] == "complete":
            entry["statistics"] = json.loads((path / "stdout").read_text())
            entry["cpu_profile_retained"] = (path / "cpu.pprof").exists()
        analysis = RESULTS / f"analysis.{tag}.full540.{version}"
        entry["flat_profile"] = profile_analysis(analysis / "top.txt")
        entry["cumulative_profile"] = profile_analysis(analysis / "cumulative.txt")
        listing = analysis / "projection.ascii.txt"
        entry["projection_listing"] = source(listing) if listing.exists() else None
        output["profiles"][version] = entry
    for path in sorted(RESULTS.glob(f"annotation.{tag}.*.r[123]")):
        entry = {"status": completed(path), "source": source(path)}
        if entry["status"] == "complete":
            text = (path / "stdout").read_text()
            matches = re.findall(r"rows (\d+), hits (\d+)", text)
            if len(matches) != 1:
                raise ValueError(f"Incomplete annotation statistics: {path}")
            entry.update(rows=int(matches[0][0]), hits=int(matches[0][1]), **time_report(path / "time.txt"))
        output["annotation"][path.name] = entry
    output["annotation_summary"] = {}
    for cell in CELLS[:-1]:
        entries = {}
        for version in VERSIONS:
            runs = [output["annotation"].get(f"annotation.{tag}.{cell}.{version}.r{r}") for r in (1, 2, 3)]
            entry = {"status": "pending"}
            if all(run and run["status"] == "complete" for run in runs):
                if len({(run["rows"], run["hits"]) for run in runs}) != 1:
                    raise ValueError(f"Annotation counts drift: {cell}, {version}")
                entry = {"status": "complete", "rows": runs[0]["rows"], "hits": runs[0]["hits"],
                         "sources": [run["source"] for run in runs],
                         "median": {key: statistics.median(run[key] for run in runs)
                                    for key in ("wall_seconds", "user_seconds", "system_seconds", "rss_kib")}}
            elif any(run is not None for run in runs):
                entry["status"] = "incomplete"
            entries[version] = entry
        output["annotation_summary"][cell] = entries
    digest = RESULTS / f"digest.{tag}.check"
    if (digest / "finished").exists():
        output["digests"] = {"status": "complete", "source": source(digest / "comparison.json"),
                             "comparisons": json.loads((digest / "comparison.json").read_text())}
    else:
        output["digests"] = {"status": "pending_or_incomplete"}
    return output


def markdown(data):
    lines = ["# Projection experiment report", "", f"Run tag: `{data['tag']}`.", "",
             "Only complete three-round pairs receive medians; profile and digest runs are excluded.", "",
             "| Workload | Version | Status | Lookup seconds/pass | Process user seconds | Process peak RSS MiB | Sources |",
             "| --- | --- | --- | ---: | ---: | ---: | --- |"]
    for cell, item in data["query_only"].items():
        for version, result in item["versions"].items():
            if result["status"] == "complete":
                median = result["median"]
                sources = ", ".join(f"`{r['lookup_source']}`, `{r['time_source']}`" for r in result["runs"])
                lines.append(f"| {cell} | {version} | complete | {median['lookup_seconds_per_pass']:.6f} | {median['user_seconds']:.3f} | {median['rss_kib']/1024:.2f} | {sources} |")
            else:
                lines.append(f"| {cell} | {version} | {result['status']} | - | - | - | - |")
    lines += ["", "## Paired lookup changes", "",
              "Negative values mean less lookup time; each change uses matching rounds.", "",
              "| Workload | Comparison | Status | Median paired change | Sources |",
              "| --- | --- | --- | ---: | --- |"]
    for cell, item in data["query_only"].items():
        for comparison, result in item["comparisons"].items():
            if result["status"] == "complete":
                sources = ", ".join(f"`{path}`" for path in result["sources"])
                lines.append(f"| {cell} | {comparison} | complete | {result['median_paired_lookup_change_percent']:+.2f}% | {sources} |")
            else:
                lines.append(f"| {cell} | {comparison} | {result['status']} | - | - |")
    lines += ["", "## CPU profile diagnostics", "",
              "Profiles are excluded from the timing matrix; cumulative samples include child functions.", "",
              "| Version | Function | Cumulative CPU seconds | Cumulative sample share | Source |",
              "| --- | --- | ---: | ---: | --- |"]
    for version, profile in data["profiles"].items():
        cumulative = profile["cumulative_profile"]
        functions = [row for row in cumulative.get("functions", [])
                     if row["function"].endswith((".projectRecord", ".projectRecordInto", ".walkRecord", ".decodeValue", ".cgocall", ".scanObject"))]
        if not functions:
            lines.append(f"| {version} | pending analysis or function absent from recorded top nodes | - | - | `{cumulative['source']}` |")
        for row in functions:
            lines.append(f"| {version} | `{row['function']}` | {row['cumulative_seconds']:.2f} | {row['cumulative_percent']:.2f}% | `{cumulative['source']}` |")
    lines += ["", "## Complete annotation verification", ""]
    if not data["verification"]:
        lines.append("Pending; no completed annotation verification is recorded for this tag.")
    for name, entry in data["verification"].items():
        if entry["status"] == "complete":
            fact = entry["evidence"]
            lines.append(f"- `{name}`: {fact['expected_rows']} rows, {fact['field_comparisons']} field comparisons, zero value/hit differences and equal full-stream SHA-256 (`{entry['source']}/verification.json`).")
        else:
            lines.append(f"- `{name}`: {entry['status']} (`{entry['source']}`).")
    lines += ["", "## Raw-record digest checks", ""]
    if data["digests"]["status"] == "complete":
        lines.append(f"Complete diagnostic comparisons: `{data['digests']['source']}`.")
        for cell, evidence in data["digests"]["comparisons"].items():
            versions = ", ".join(evidence["exact_comparison_versions"])
            equality = evidence["exact_raw_records_equal"]
            lines.append(f"- `{cell}`: versions {versions}; exact raw-record equality = {equality}; legacy subset excluded = {evidence['legacy_subset_comparison_excluded']}.")
    else:
        lines.append("Pending or incomplete; no digest pass is inferred.")
    lines += ["", "## Annotation to /dev/null", "",
              f"Scope is run tag `{data['tag']}`; query input hashes are retained in `results/annotation.{data['tag']}.matrix/inputs.sha256` when present.", "",
              "| Workload | Version | Status | Process wall seconds | User CPU seconds | Peak RSS MiB | Sources |",
              "| --- | --- | --- | ---: | ---: | ---: | --- |"]
    for cell, entries in data["annotation_summary"].items():
        for version, entry in entries.items():
            if entry["status"] == "complete":
                median = entry["median"]
                sources = ", ".join(f"`{path}/time.txt`" for path in entry["sources"])
                lines.append(f"| {cell} | {version} | complete | {median['wall_seconds']:.3f} | {median['user_seconds']:.3f} | {median['rss_kib']/1024:.2f} | {sources} |")
            else:
                lines.append(f"| {cell} | {version} | {entry['status']} | - | - | - | - |")
    lines += ["", "CPU/RSS cover the whole process; lookup time excludes setup and output formatting.",
              "Raw-record subset bytes from old are not required to match before/after; selected annotation fields are verified separately.",
              "Runtime path selection is automatic; these records do not establish per-invocation io_uring or O_DIRECT use.",
              "CPU and memory limits do not control NVMe service time or OS cache state; wall-time ratios alone cannot isolate CPU improvements.", ""]
    return "\n".join(lines)


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--tag", default="wgs")
    parser.add_argument("--format", choices=("json", "markdown"), default="json")
    parser.add_argument("--output", type=Path, help="new file under this experiment's results directory")
    args = parser.parse_args()
    if not re.fullmatch(r"[a-zA-Z0-9_-]+", args.tag):
        parser.error("Invalid tag")
    data = report(args.tag)
    text = json.dumps(data, indent=2, sort_keys=True) + "\n" if args.format == "json" else markdown(data)
    if args.output:
        output = args.output.resolve()
        if not output.is_relative_to(RESULTS):
            parser.error("Output must be inside this experiment's results directory")
        with output.open("x") as stream:
            stream.write(text)
        print(output)
    else:
        print(text, end="")


if __name__ == "__main__":
    main()
