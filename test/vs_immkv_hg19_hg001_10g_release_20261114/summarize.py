#!/usr/bin/env python3
"""Generate the report from raw measurements without treating missing runs as passes."""
import json
from pathlib import Path
import re
import statistics

ROOT = Path(__file__).resolve().parent
RESULTS = ROOT / "results"
VERSIONS = ("baseline", "candidate")


def records(path):
    return [json.loads(line) for line in path.read_text().splitlines() if line.startswith("{")]


def timing(path):
    result = {}
    if not path.exists():
        return result
    for line in path.read_text().splitlines():
        text = line.strip()
        for prefix, key, cast in (
            ("Maximum resident set size (kbytes):", "rss_kib", int),
            ("User time (seconds):", "user_seconds", float),
            ("System time (seconds):", "system_seconds", float),
            ("File system inputs:", "filesystem_inputs", int),
            ("Exit status:", "exit_status", int),
        ):
            if text.startswith(prefix):
                result[key] = cast(text[len(prefix):].strip())
        if text.startswith("Elapsed (wall clock) time"):
            total = 0.0
            for part in text.split("): ", 1)[1].split(":"):
                total = total * 60 + float(part)
            result["wall_seconds"] = total
        if "terminated by signal" in text.lower():
            result["interrupted"] = True
    return result


def paired(prefix, primary, annotation=False):
    output = {}
    for version in VERSIONS:
        runs = []
        for repeat in (1, 2, 3):
            path = RESULTS / f"{prefix}.{version}.r{repeat}.stdout"
            if not path.exists() or not path.stat().st_size:
                continue
            if annotation:
                text = path.read_text()
                match = re.search(r"rows (\d+), hits (\d+).*; wall ([0-9.]+)(ms|s)", text)
                if not match:
                    raise ValueError(f"incomplete annotation statistics: {path}")
                item = {
                    "rows": int(match[1]), "hits": int(match[2]),
                    "annotation_seconds": float(match[3]) / (1000 if match[4] == "ms" else 1),
                }
            else:
                rows = records(path)
                if len(rows) != 1:
                    raise ValueError(f"expected one JSON record: {path}")
                item = dict(rows[0])
            item.update(timing(path.with_suffix(".time")))
            if item.get("interrupted") or item.get("exit_status") != 0:
                continue
            item["source"] = str(path.relative_to(ROOT))
            runs.append(item)
        if runs:
            output[version] = {"runs": runs, "median": {}}
            for key in (primary, "keys_per_second", "build_seconds", "records_per_second", "wall_seconds", "rss_kib", "user_seconds", "system_seconds"):
                values = [run[key] for run in runs if key in run]
                if values:
                    output[version]["median"][key] = statistics.median(values)
    if all(version in output for version in VERSIONS):
        output["candidate_change_percent"] = 100 * (
            output["candidate"]["median"][primary] / output["baseline"]["median"][primary] - 1
        )
    return output


def main():
    facts = json.loads((RESULTS / "source-facts.json").read_text())
    result = {"sources": facts, "point": {}, "batch": {}, "diagnostic": {}, "annotation": {}, "verification": {}}
    lines = [
        "# Release performance regression: measured results", "",
        "**Do not publish this candidate without addressing the explicit-column batch-query regression.**", "",
        "## Controlled facts", "",
        f"Baseline: `{facts['baseline']}` (before the API/error-channel refactor).",
        f"Candidate: uncommitted working-tree snapshot, digest `{facts['candidate_source_digest']}`.",
        "Source: `results/source-facts.json`; tracked changes: `results/candidate.patch`.", "",
        "Both versions read the same existing hg19 / gnomAD v2.1.1 v003 libraries.",
        "All measured commands used CPUs 0-3, four workers, MemoryMax=10G and MemorySwapMax=0.",
        "Sources: `common.sh`, `run_limited.sh`, `results/*.cgroup`.",
        "Compiler, CPU topology and exact input hashes: `results/environment.txt`.", "",
        "The full HG001 query has 3,933,714 rows, 3,933,369 buildable keys and 3,875,208 hits.",
        "Sources: prior `../vs_echtvar_hg19_gnomad_genomes_4c16g_gz_v003/results/matrix.log` (input rows), and this run's `results/null.*.stdout` (buildable keys and hits). The key-only harness filters unbuildable alleles.", "",
        "## Random point lookup", "",
        "Median of three paired runs with the same seed and warmups; sampled frames are warm in the OS page cache.", "",
        "| cache MiB | phase | baseline QPS | candidate QPS | QPS change | baseline p50 us | candidate p50 us | baseline p99 us | candidate p99 us |",
        "|---:|---|---:|---:|---:|---:|---:|---:|---:|",
    ]
    for cache in (0, 512):
        for phase in ("sampled", "misses"):
            cell = {}
            for version in VERSIONS:
                runs = []
                for repeat in (1, 2, 3):
                    path = RESULTS / f"point.cache{cache}.{version}.r{repeat}.stdout"
                    if not path.exists():
                        continue
                    for entry in records(path):
                        if entry["phase"] == phase:
                            runs.append({**entry, "source": str(path.relative_to(ROOT))})
                if runs:
                    if any(run["count"] != 100000 or run["hits"] != (98486 if phase == "sampled" else 0) for run in runs):
                        raise ValueError("point count drift")
                    cell[version] = {"runs": runs, "median": {
                        key: statistics.median(run[key] for run in runs)
                        for key in ("qps", "p50_ns", "p99_ns", "p999_ns")
                    }}
            result["point"][f"cache{cache}.{phase}"] = cell
            if all(version in cell for version in VERSIONS):
                a, b = (cell[version]["median"] for version in VERSIONS)
                delta = 100 * (b["qps"] / a["qps"] - 1)
                lines.append(f"| {cache} | {phase} | {a['qps']:.0f} | {b['qps']:.0f} | {delta:+.2f}% | {a['p50_ns']/1000:.3f} | {b['p50_ns']/1000:.3f} | {a['p99_ns']/1000:.3f} | {b['p99_ns']/1000:.3f} |")
    lines += ["", "Sources: `results/point.cache*.r*.stdout`; setup runs are excluded.", "", "## Whole-genome GetBatch only", "",
              "No INFO formatting or compression; query loading/sorting is excluded from lookup time.", "",
              "| workload | pairs | baseline lookup s | candidate lookup s | time change | baseline peak RSS MiB | candidate peak RSS MiB |",
              "|---|---:|---:|---:|---:|---:|---:|"]
    for name in ("narrow17", "full17", "full540", "full568", "full568_default"):
        cell = paired("null." + name, "lookup_seconds")
        result["batch"][name] = cell
        if "candidate_change_percent" not in cell:
            lines.append(f"| {name} | incomplete | - | - | - | - | - |")
            continue
        a, b = (cell[version] for version in VERSIONS)
        for run in a["runs"] + b["runs"]:
            if run["keys"] != 3933369 or run["hits"] != 3875208:
                raise ValueError("batch count drift")
        pairs = min(len(a["runs"]), len(b["runs"]))
        suffix = " *" if name == "full568" else ""
        lines.append(f"| {name}{suffix} | {pairs} | {a['median']['lookup_seconds']:.3f} | {b['median']['lookup_seconds']:.3f} | {cell['candidate_change_percent']:+.2f}% | {a['median']['rss_kib']/1024:.1f} | {b['median']['rss_kib']/1024:.1f} |")
    lines += ["", "Sources: `results/null.*.stdout` and matching `.time` files; RSS includes the complete process.",
              "* Full568 timings include a changed NVMe I/O regime; retain the raw samples but do not interpret the median ratio as stable code-only slowdown.",
              "The following default-path whole-genome run was interrupted (signal 2) because disk throughput collapsed; its partial time is not a completed result.",
              "Source: `results/io-regime-note.txt` and `results/null.full568_default.baseline.r1.time`.", "",
              "Explicit subset output bytes differ: the old record-order path ignored Columns and returned full records, while the candidate performs the promised projection.",
              "Selected annotation values are checked separately below.", "",
              "| workload | baseline user CPU s | candidate user CPU s | CPU change |",
              "|---|---:|---:|---:|"]
    for name in ("narrow17", "full17", "full540", "full568"):
        cell = result["batch"][name]
        a, b = (cell[version]["median"]["user_seconds"] for version in VERSIONS)
        lines.append(f"| {name} | {a:.2f} | {b:.2f} | {(b/a-1)*100:+.2f}% |")
    lines += ["", "CPU source: median of `User time (seconds)` in `results/null.<workload>.<version>.r*.time`.", "",
              "## Chr22 diagnostic", "",
              "This separate subset has 49,540 rows, 49,536 buildable keys and 48,709 hits.",
              "Sources: `results/chr22-query.sha256`, `results/diagnostic.chr22.*.stdout` and verification logs.", "",
              "| call | baseline lookup s | candidate lookup s | time change | baseline user CPU s | candidate user CPU s |",
              "|---|---:|---:|---:|---:|---:|"]
    for mode in ("default", "explicit"):
        cell = paired("diagnostic.chr22." + mode, "lookup_seconds")
        result["diagnostic"][mode] = cell
        a, b = (cell[version]["median"] for version in VERSIONS)
        lines.append(f"| {mode} | {a['lookup_seconds']:.3f} | {b['lookup_seconds']:.3f} | {cell['candidate_change_percent']:+.2f}% | {a['user_seconds']:.2f} | {b['user_seconds']:.2f} |")
    lines += ["", "Sources: `results/diagnostic.chr22.*.stdout` and `.time`. Default passes Columns=nil; explicit selects all 568 fields.",
              "These subsecond default results have substantial run-to-run variability; do not extrapolate their percentage to the whole genome.", "",
              "## Chr22 annotation including typed formatting", "",
              "Text output goes to /dev/null; this includes VCF parsing, lookup and formatting but excludes gzip and stored-output I/O.", "",
              "| workload | baseline annotation s | candidate annotation s | time change |",
              "|---|---:|---:|---:|"]
    for name in ("narrow17", "full17", "full540", "full568"):
        cell = paired("annotation.chr22." + name, "annotation_seconds", annotation=True)
        result["annotation"][name] = cell
        a, b = (cell[version]["median"]["annotation_seconds"] for version in VERSIONS)
        lines.append(f"| {name} | {a:.3f} | {b:.3f} | {cell['candidate_change_percent']:+.2f}% |")
    lines += ["", "Sources: median `wall` statistic in `results/annotation.chr22.*.stdout`; whole-process CPU/RSS in matching `.time` files.", "",
              "## Correctness", "",
              "Both complete chr22 annotation streams were checked for every row and field using `test/verifyvalues` with tolerance zero, plus whole-stream SHA-256.", "",
              "| workload | rows | field comparisons | mismatches | full-stream hashes equal | source |",
              "|---|---:|---:|---:|---|---|"]
    for name in ("narrow17", "full17", "full540", "full568"):
        path = RESULTS / (f"verify.chr22.{name}.log" if name != "full568" else "verify.full568.log")
        text = path.read_text()
        match = re.search(r"rows (\d+), field comparisons (\d+).*value mismatches (\d+), hit disagreements (\d+)", text)
        good_hash = "EXACT_STREAM_SHA256_OK" in text
        item = {"rows": int(match[1]), "comparisons": int(match[2]), "mismatches": int(match[3])+int(match[4]),
                "sha256_equal": good_hash, "source": str(path.relative_to(ROOT))}
        result["verification"][name] = item
        if item["rows"] != 49540 or item["mismatches"] or not good_hash:
            raise ValueError("correctness verification failed")
        lines.append(f"| {name} | {item['rows']} | {item['comparisons']} | {item['mismatches']} | yes | `{item['source']}` |")
    lines += ["", "The initial full568 verification log lacks chr22 in its filename but used the same fixed chr22 query; its row count and producer statistics identify the scope.",
              "Full-genome field-by-field verification has not been completed.", "",
              "## Interpretation and release decision", "",
              "- Point-lookup throughput is close to the baseline for this warm-cache sample; the no-cache miss tail is noisier and is retained in the table.",
              "- Explicit Columns causes a large batch CPU regression, including the all-selected case where returned bytes are identical.",
              "- The candidate StreamDB.GetBatch performs projection after joining all read/decompression workers; projectRecord walks fields and allocates additional buffers even when every field is selected.",
              "- The chr22 explicit-versus-default diagnostic and higher user CPU support projection as a major contributor; no profile was collected, so the exact fraction is not established.",
              "- Annotation with formatting still regresses in every tested chr22 workload; this is not solely an artifact of the query-only harness.",
              "- Keep this candidate unreleased until the batch regression is addressed and the same controlled tests are repeated.", "",
              "## Limits and unexecuted phases", "",
              "- This is NVMe testing; no HDD performance claim is made.",
              "- Cold point-lookup performance was not compared under controlled cold-cache conditions.",
              "- Whole-genome round 3 experienced I/O disruption; raw files remain intact and CPU times are reported separately.",
              "- Full stored-output gzip timing and synthetic build throughput were not run because authorization to write the external work directory was not received.",
              "- No production library was modified, rebuilt or deleted; the library source code was not changed during measurement.",
              "- Scripts and observed data are preserved here; historical experiments are unchanged.", ""]
    (RESULTS / "summary.json").write_text(json.dumps(result, indent=2) + "\n")
    (ROOT / "README.md").write_text("\n".join(lines))
    print(ROOT / "README.md")


if __name__ == "__main__":
    main()
