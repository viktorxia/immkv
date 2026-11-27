#!/usr/bin/env python3
"""Read-only input validation and environment inventory; never run benchmarks."""
import argparse
from collections import Counter
from datetime import datetime, timezone
import gzip
import hashlib
import json
import os
from pathlib import Path
import platform
import shutil
import stat
import subprocess
import sys

NAME = "vs_immkv_hg19_hg001_4c16g_cold_postfix_20261125"
BASE_COMMIT = "904948f0db6f59d94e8a69d97d300fba82aae24b"
EXPECTED_ROWS = 3933714
EXPECTED_SHA256 = "0250956380e1f6910a83f92a579810671e1fe4dfeb0b1d855394021a87148fd1"
CHROMOSOMES = [str(n) for n in range(1, 23)] + ["X", "Y"]
EXP = Path(__file__).resolve().parent
REPO = EXP.parent.parent
BUILD = REPO / "tmp" / NAME
ROOTS = {"nvme": Path("/home/tian/immkv_test_nvme"),
         "hdd": Path("/mnt/data6t/tian/immkv_test_hdd")}


def now():
    return datetime.now(timezone.utc).isoformat()


def metadata(path):
    path = Path(path)
    s = path.stat()
    return {"path": str(path), "resolved_path": str(path.resolve()),
            "size_bytes": s.st_size, "device": s.st_dev, "inode": s.st_ino,
            "mtime_ns": s.st_mtime_ns, "ctime_ns": s.st_ctime_ns,
            "mode": stat.filemode(s.st_mode), "uid": s.st_uid, "gid": s.st_gid}


def sha256(path):
    h = hashlib.sha256()
    with Path(path).open("rb") as f:
        for block in iter(lambda: f.read(1048576), b""):
            h.update(block)
    return h.hexdigest()


def atomic_json(path, data):
    tmp = path.with_name(path.name + ".partial")
    with tmp.open("x") as f:
        json.dump(data, f, indent=2, sort_keys=True)
        f.write("\n")
    tmp.rename(path)


def read_text(path):
    try:
        return Path(path).read_text().strip()
    except OSError as exc:
        return {"unavailable": str(exc)}


def command(args):
    try:
        p = subprocess.run([str(a) for a in args], capture_output=True, text=True,
                           timeout=60, env=os.environ.copy())
        return {"argv": [str(a) for a in args], "exit_code": p.returncode,
                "stdout": p.stdout, "stderr": p.stderr}
    except (OSError, subprocess.TimeoutExpired) as exc:
        return {"argv": [str(a) for a in args], "error": str(exc)}


def unchanged(before, path):
    after = metadata(path)
    if before != after:
        raise ValueError(f"Input identity changed during read: {path}")


def query_facts(path):
    identity = metadata(path)
    rows = header_rows = size = 0
    data_hash = hashlib.sha256()
    full_hash = hashlib.sha256()
    chromosomes = Counter()
    header_chromosomes = []
    max_line = 0
    final_newline = False
    with gzip.open(path, "rb") as stream:
        for line in stream:
            size += len(line)
            max_line = max(max_line, len(line))
            full_hash.update(line)
            final_newline = line.endswith(b"\n")
            if line.startswith(b"#"):
                header_rows += 1
                if line.startswith(b"##contig=<ID="):
                    header_chromosomes.append(line.split(b"=", 2)[2].split(b",", 1)[0].split(b">", 1)[0].decode())
                continue
            columns = line.split(b"\t", 2)
            if len(columns) < 3:
                raise ValueError(f"Invalid VCF row {rows + 1}: {path}")
            chromosomes[columns[0].decode("ascii")] += 1
            rows += 1
            data_hash.update(line)
    unchanged(identity, path)
    digest = data_hash.hexdigest()
    if rows != EXPECTED_ROWS or digest != EXPECTED_SHA256:
        raise ValueError(f"Query baseline mismatch: {path}: rows={rows}, sha256={digest}")
    supported = {}
    unsupported = {}
    for name, n in chromosomes.items():
        canonical = name[3:] if name.startswith("chr") or name.startswith("CHR") else name
        canonical = canonical.upper()
        if canonical in CHROMOSOMES:
            supported[canonical] = supported.get(canonical, 0) + n
        else:
            unsupported[name] = n
    return {"identity": identity, "rows": rows, "data_rows": rows,
            "data_sha256": digest, "data_row_sha256": digest,
            "decompressed_sha256": full_hash.hexdigest(), "decompressed_bytes": size,
            "header_rows": header_rows, "header_contigs": header_chromosomes,
            "chromosome_rows": dict(chromosomes), "supported_chromosome_rows": supported,
            "unsupported_chromosome_rows": unsupported,
            "gzip_complete_eof": True, "gzip_crc_and_isize_validated": True,
            "maximum_line_bytes": max_line, "final_newline": final_newline}


def field_facts(count):
    path = REPO / f"test/vs_echtvar_hg19_gnomad_genomes_{count}_fields/fields{count}.tsv"
    rows = []
    for line in path.read_text().splitlines():
        if not line.strip() or line.startswith("#"):
            continue
        parts = line.split("\t")
        if len(parts) != 3 or parts[1] not in ("Integer", "Float", "String", "Flag"):
            raise ValueError(f"Invalid field TSV row in {path}: {line!r}")
        rows.append(parts)
    if len(rows) != count or len({r[0] for r in rows}) != count:
        raise ValueError(f"Field count or uniqueness mismatch: {path}")
    frozen_path = BUILD / "source" / path.relative_to(REPO)
    digest = sha256(path)
    if digest != sha256(frozen_path):
        raise ValueError(f"Field TSV differs from frozen source: {path}")
    return {"identity": metadata(path), "sha256": digest, "matches_frozen_source": True, "count": len(rows),
            "type_counts": dict(Counter(r[1] for r in rows)), "definitions": rows}


def library_facts(path, count, required_chromosomes):
    files = [metadata(p) for p in sorted(path.iterdir()) if p.is_file()]
    by_name = {Path(f["path"]).name: f for f in files}
    manifests = []
    available = set()
    for p in sorted(path.glob("chr*.manifest")):
        identity = metadata(p)
        manifest = json.loads(p.read_text())
        chromosome = p.stem[3:]
        available.add(chromosome)
        expected_count = 17 if count == 17 else 568
        if len(manifest["fields"]) != expected_count:
            raise ValueError(f"Library schema field count mismatch: {p}")
        if manifest["format"] != "immkv-v003" or manifest["version"] != 3:
            raise ValueError(f"Unexpected library manifest format: {p}")
        for ext in ("keys", "vals", "col"):
            file_name = p.stem + "." + ext
            if file_name not in by_name or by_name[file_name]["size_bytes"] != manifest[ext + "_bytes"]:
                raise ValueError(f"Library file size mismatch: {path / file_name}")
        manifests.append({"identity": identity, "sha256": sha256(p),
                          "chromosome": chromosome, "build_id": manifest["build_id"],
                          "format": manifest["format"], "schema_hash": manifest["schema_hash"],
                          "field_count": len(manifest["fields"]), "record_count": manifest["record_count"]})
    if set(required_chromosomes) - available:
        raise ValueError(f"Missing required library chromosomes: {path}: {set(required_chromosomes) - available}")
    return {"path": str(path), "files": files, "manifests": manifests,
            "available_chromosomes": sorted(available, key=CHROMOSOMES.index),
            "total_size_bytes": sum(f["size_bytes"] for f in files),
            "large_file_content_hashing": False}


def reference_facts(media, count):
    root = ROOTS["hdd"] if count == 568 else ROOTS[media]
    sub = "test/4c16g_gz_v003/568_fields/echtvar_ref.wk" if count == 568 else f"test/4c16g_gz_v003/{count}_fields/echtvar.pass2.wk"
    ref_dir = root / sub
    files = []
    for i in range(1, 23):
        p = ref_dir / f"an{i}.vcf.gz"
        if not p.is_file():
            raise ValueError(f"Missing historical reference file: {p}")
        files.append(metadata(p))
    return {"directory": str(ref_dir), "files": files}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--label", default="", help="Optional alphanumeric suffix")
    args = parser.parse_args()
    if EXP != REPO / "test" / NAME:
        parser.error("Unexpected experiment path")
    suffix = "-" + args.label if args.label else ""
    results = EXP / "results"
    results.mkdir(exist_ok=True)
    report_path = results / f"input-preflight{suffix}.json"
    log_path = results / f"input-preflight{suffix}.log"
    for p in (report_path, log_path):
        if p.exists() or p.with_name(p.name + ".partial").exists():
            parser.error(f"Result already exists; preserve evidence: {p}")
    report = {"started_at": now(), "expected_rows": EXPECTED_ROWS,
              "expected_data_sha256": EXPECTED_SHA256, "base_commit": BASE_COMMIT,
              "scenarios": [], "fields": {}, "errors": []}
    with log_path.open("x", buffering=1) as log:
        def note(text):
            value = f"{now()} {text}"
            print(value, flush=True)
            log.write(value + "\n")
        note("Preparation input preflight started")
        build_record = results / "preparation"
        try:
            if (build_record / "exit-status").read_text().strip() != "0":
                raise ValueError("Frozen build did not finish successfully")
            for line in (build_record / "binaries.sha256").read_text().splitlines():
                digest, name = line.split(None, 1)
                if sha256(BUILD / "bin" / name.strip()) != digest:
                    raise ValueError(f"Frozen binary hash mismatch: {name}")
            report["preparation_record"] = str(build_record)
            report["frozen_binaries_validated"] = True
        except Exception as exc:
            report["errors"].append(str(exc))
        for n in (17, 540, 568):
            try:
                report["fields"][str(n)] = field_facts(n)
                note(f"Validated {n}-field TSV")
            except Exception as exc:
                report["errors"].append(str(exc))
        library_cache = {}
        for media, root in ROOTS.items():
            for n in (17, 540, 568):
                scenario = {"media": media, "fields": n, "errors": []}
                report["scenarios"].append(scenario)
                base = root / f"test/{n}_fields"
                query = base / "query/hg001.decomp.vcf.gz"
                note(f"Reading complete query: {media} {n}")
                try:
                    scenario["query"] = query_facts(query)
                    required = list(scenario["query"]["supported_chromosome_rows"])
                    scenario["required_chromosomes"] = [c for c in CHROMOSOMES if c in required]
                    library = root / ("test/17_fields/lib17" if n == 17 else "gnomad/immkv")
                    if str(library) not in library_cache:
                        library_cache[str(library)] = library_facts(library, n, required)
                    scenario["library"] = library_cache[str(library)]
                    scenario["reference"] = reference_facts(media, n)
                    scenario["valid"] = True
                    note(f"Scenario passed: {media} {n}")
                except Exception as exc:
                    scenario["errors"].append(str(exc))
                    report["errors"].append(f"{media} {n}: {exc}")
                    scenario["valid"] = False
                    note(f"FAILED {media} {n}: {exc}")
        report["finished_at"] = now()
        report["queries_validated"] = sum(s.get("query", {}).get("gzip_complete_eof", False) for s in report["scenarios"])
        report["valid"] = not report["errors"] and len(report["scenarios"]) == 6
        atomic_json(report_path, report)
        note(f"Preflight complete: valid={report['valid']}; queries={report['queries_validated']}")
    return 0 if report["valid"] else 1


if __name__ == "__main__":
    sys.exit(main())
