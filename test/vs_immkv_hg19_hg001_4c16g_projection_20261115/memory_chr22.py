#!/usr/bin/env python3
"""Resident chr22 diagnostic; invoke through run_memory_chr22.sh, never directly.

All corpus bytes are copied and checked before any digest or timed invocation.
The parent retains sealed memfds; children open /proc/<parent>/fd through local
symlinks, so descriptors need not be inheritable. No source library is writable.
Only corpus reads are resident: executable loading and result logging retain
normal filesystem behavior. This is not an NVMe benchmark or a cold-cache run.
"""
import argparse
import ctypes
import fcntl
import hashlib
import json
import os
from pathlib import Path
import signal
import stat
import struct
import subprocess
import time

EXP = Path(__file__).resolve().parent
CELLS = ("narrow17", "full17", "full540", "full568", "full568_default")
VERSIONS = ("old", "before", "after")
BLOCK = 4 << 20


def require(condition, message):
    if not condition:
        raise RuntimeError(message)


def save(path, data):
    with path.open("x") as stream:
        json.dump(data, stream, indent=2, sort_keys=True)
        stream.write("\n")


def resources(expected=None):
    group = next(line[3:] for line in Path("/proc/self/cgroup").read_text().splitlines() if line.startswith("0::"))
    root = Path("/sys/fs/cgroup" + group)
    data = {"cgroup": group, "cpus": sorted(os.sched_getaffinity(0)), "gomaxprocs": os.environ.get("GOMAXPROCS")}
    for name in ("memory.max", "memory.swap.max", "memory.current", "memory.peak", "memory.stat", "memory.events", "io.stat"):
        if (root / name).exists():
            data[name] = (root / name).read_text().strip()
    require(data["memory.max"] == "17179869184" and data["memory.swap.max"] == "0", "Expected 16 GiB and zero swap")
    require(data["cpus"] == [0, 1, 2, 3] and data["gomaxprocs"] == "4", "Expected CPUs 0-3 and GOMAXPROCS=4")
    require(expected is None or group == expected, "Child scope changed")
    return data


def preload(source, target, descriptors, links):
    name = "immkv-" + target.parent.name + "-" + target.name
    if hasattr(os, "memfd_create"):
        fd = os.memfd_create(name, 0x1 | 0x2)  # MFD_CLOEXEC | MFD_ALLOW_SEALING
    else:
        create = ctypes.CDLL(None, use_errno=True).memfd_create
        create.argtypes, create.restype = [ctypes.c_char_p, ctypes.c_uint], ctypes.c_int
        fd = create(name.encode(), 0x1 | 0x2)
        if fd < 0:
            error = ctypes.get_errno()
            raise OSError(error, os.strerror(error))
    descriptors.append(fd)
    digest = hashlib.sha256()
    with source.open("rb", buffering=0) as stream:
        before = os.fstat(stream.fileno())
        require(stat.S_ISREG(before.st_mode), f"Not a regular input: {source}")
        while block := os.read(stream.fileno(), BLOCK):
            digest.update(block)
            remaining = memoryview(block)
            while remaining:
                written = os.write(fd, remaining)
                require(written > 0, "Short memfd write")
                remaining = remaining[written:]
        after = os.fstat(stream.fileno())
    identity = lambda st: (st.st_dev, st.st_ino, st.st_size, st.st_mtime_ns, st.st_ctime_ns)
    require(identity(before) == identity(after), f"Input changed during preload: {source}")
    require(os.fstat(fd).st_size == before.st_size, f"Wrong memfd length: {source}")
    os.lseek(fd, 0, os.SEEK_SET)
    check = hashlib.sha256()
    while block := os.read(fd, BLOCK):
        check.update(block)
    require(check.digest() == digest.digest(), f"Memfd digest mismatch: {source}")
    os.lseek(fd, 0, os.SEEK_SET)
    os.fchmod(fd, 0o444)
    seals = 0x0F  # F_SEAL_SEAL | F_SEAL_SHRINK | F_SEAL_GROW | F_SEAL_WRITE
    fcntl.fcntl(fd, getattr(fcntl, "F_ADD_SEALS", 1033), seals)
    target.symlink_to(f"/proc/{os.getpid()}/fd/{fd}")
    links.append(target)
    print(f"Preloaded {source}: {before.st_size} bytes", flush=True)
    return {"source": str(source), "proxy": str(target), "bytes": before.st_size,
            "sha256": digest.hexdigest(), "memfd_sha256": check.hexdigest(), "seals": seals}


def check_library(root, fields):
    manifest = json.loads((root / "chr22.manifest").read_text())
    require(manifest["format"] == "immkv-v003" and manifest["version"] == 3 and manifest["base"] == "chr22", "Invalid chr22 manifest")
    require(len(manifest["fields"]) == fields and manifest["record_count"] > 0, "Invalid schema or record count")
    for suffix, magic, version in (("keys", b"IMKVK002", 1), ("vals", b"IMKVV003", 3)):
        path = root / ("chr22." + suffix)
        with path.open("rb") as stream:
            header = stream.read(64)
        require(header[:8] == magic and struct.unpack_from("<I", header, 8)[0] == version, f"Invalid {suffix} header")
        require(struct.unpack_from("<Q", header, 12)[0] == manifest["record_count"], f"{suffix} count differs from manifest")
        require(path.stat().st_size == manifest[suffix + "_bytes"], f"{suffix} size differs from manifest")
    with (root / "chr22.col").open("rb") as stream:
        header = stream.read(64)
    require(header[:8] == b"IMKVC003" and header[16:32].hex() == manifest["build_id"], "Columnar build ID mismatch")
    require((root / "chr22.col").stat().st_size == manifest["col_bytes"], "Columnar size mismatch")
    return {key: manifest[key] for key in ("record_count", "build_id", "schema_hash")}


def run_child(record, name, command, timed, group, order):
    child = record / name
    child.mkdir()  # Refuse replacement, including failed results.
    save(child / "resident.before.json", resources(group))
    with order.open("a") as stream:
        stream.write(name + "\n")
    print(f"Starting {name}", flush=True)
    result = subprocess.run(["bash", str(EXP / "run_limited.sh"), str(child),
                             "timed" if timed else "untimed", *map(str, command)], check=False)
    save(child / "resident.after.json", resources(group))
    require(result.returncode == 0 and (child / "finished").exists(), f"Failed invocation: {child}")
    require((child / "cgroup.txt").read_text().splitlines()[0] == "cgroup=" + group, "Invocation escaped resident scope")
    return child


def check_query(data, count, digest):
    require(data["threads"] == 4 and data["gomaxprocs"] == 4 and data["batch"] == 4096, "Query resource mismatch")
    require(data["repeat"] == count and len(data["passes"]) == count and not data["profiled"], "Invalid query passes")
    require(data["digest_enabled"] == digest and data["keys"] > 0 and data["total_keys"] == data["keys"] * count, "Invalid query counts")
    first = data["passes"][0]
    for item in data["passes"]:
        require(all(item[key] == first[key] for key in ("hits", "batches", "returned_bytes", "raw_sha256") if key in first), "Query passes differ")
    if digest:
        require(len(first["raw_sha256"]) == 64, "Missing raw-record digest")
    return {"keys": data["keys"], **{key: first[key] for key in ("hits", "batches", "returned_bytes", "raw_sha256") if key in first}}


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    for name in ("record", "build-root", "query", "full-lib", "narrow-lib", "fields17", "fields540", "fields568"):
        parser.add_argument("--" + name, required=True, type=Path)
    args = parser.parse_args()
    record = args.record.resolve()
    require(record.parent == EXP / "results" and record.name.startswith("memory.chr22"), "Expected a separate memory.chr22 result directory")
    require(record.is_dir(), "Launcher must reserve the result directory")
    for name in ("IMMKV_FORCE_RECORD", "IMMKV_FORCE_COLUMNAR", "IMMKV_DEBUG_PATH", "IMMKV_NO_DIRECT", "GODEBUG", "GOGC", "GOMEMLIMIT", "LD_PRELOAD", "LD_AUDIT"):
        require(not os.environ.get(name), f"Unexpected runtime override: {name}")
    initial = resources()
    group = initial["cgroup"]
    save(record / "resident.initial.json", initial)
    scratch = record / "resident-inputs"
    scratch.mkdir()
    descriptors, links, dirs = [], [], [scratch]
    entries, metadata, summaries, digests = [], {}, {}, {}
    order = record / "order.txt"
    order.touch(exist_ok=False)
    try:
        for name, source, field_count in (("full", args.full_lib, 568), ("narrow", args.narrow_lib, 17)):
            root = scratch / name
            root.mkdir()
            dirs.append(root)
            for suffix in ("keys", "vals", "col", "manifest"):
                filename = "chr22." + suffix
                entries.append(preload(source / filename, root / filename, descriptors, links))
            metadata[name] = check_library(root, field_count)
        require(metadata["full"]["record_count"] == metadata["narrow"]["record_count"], "Full and narrow chr22 record counts differ")
        for name, source in (("query.vcf", args.query), ("fields17.tsv", args.fields17), ("fields540.tsv", args.fields540), ("fields568.tsv", args.fields568)):
            entries.append(preload(source, scratch / name, descriptors, links))
        rows = 0
        with (scratch / "query.vcf").open("rb") as stream:
            for line in stream:
                if line.startswith(b"#") or not line.strip():
                    continue
                require(line.split(b"\t", 1)[0] in (b"22", b"chr22", b"CHR22"), "Fixture contains non-chr22 data")
                rows += 1
        require(rows > 0, "Empty chr22 fixture")
        total = sum(entry["bytes"] for entry in entries)
        save(record / "inputs.json", {"kind": "memfd-resident-chr22", "total_preloaded_bytes": total,
             "preload_finished_ns": time.time_ns(), "query_rows": rows, "libraries": metadata, "files": entries,
             "io_policy": "Unmodified automatic path selection; no claim of O_DIRECT/io_uring use."})
        ready = resources(group)
        save(record / "resident.preloaded.json", ready)
        shmem = int(dict(line.split() for line in ready["memory.stat"].splitlines())["shmem"])
        require(shmem >= total - (16 << 20), "Resident input pages are not charged to this cgroup")

        def command(cell, version, annotation=False):
            fields = 17 if cell in ("narrow17", "full17") else 540 if cell == "full540" else 568
            return [args.build_root / "bin" / version / ("annotate" if annotation else "query-only"),
                    "-in" if annotation else "-query", scratch / "query.vcf", "-lib",
                    scratch / ("narrow" if cell == "narrow17" else "full"), "-fields", scratch / f"fields{fields}.tsv",
                    "-threads", "4", "-batch", "4096"] + (["-columns=false"] if cell == "full568_default" else [])

        # Complete every raw-record check before starting any timing sample.
        for cell in CELLS:
            digests[cell] = {}
            for version in VERSIONS:
                child = run_child(record, f"digest.{cell}.{version}", command(cell, version) + ["-repeat", "1", "-digest=true"], False, group, order)
                data = json.loads((child / "stdout").read_text())
                digests[cell][version] = check_query(data, 1, True)
            values = digests[cell]
            require(len({(d["keys"], d["hits"], d["batches"]) for d in values.values()}) == 1, f"Query/hit counts differ: {cell}")
            exact = VERSIONS[1:] if cell in ("full17", "full540") else VERSIONS
            require(len({(values[v]["returned_bytes"], values[v]["raw_sha256"]) for v in exact}) == 1, f"Raw records differ: {cell}")
        save(record / "digest-comparison.json", {"records": digests, "before_after_all_five_equal": True,
             "old_raw_subset_comparison_excluded": ["full17", "full540"]})
        for annotation in (False, True):
            cells, versions = (CELLS[:4], ("old", "after")) if annotation else (CELLS, VERSIONS)
            for cell in cells:
                for round_number in (1, 2, 3):
                    for version in versions[::-1] if round_number == 2 else versions:
                        kind = "annotation" if annotation else "query"
                        name = f"{kind}.{cell}.{version}.r{round_number}"
                        extra = ["-out", "/dev/null"] if annotation else ["-repeat", "10"]
                        child = run_child(record, name, command(cell, version, annotation) + extra, True, group, order)
                        summaries[name] = {"stdout": name + "/stdout", "stderr": name + "/stderr", "time": name + "/time.txt"}
                        if not annotation:
                            data = json.loads((child / "stdout").read_text())
                            identity = check_query(data, 10, False)
                            require(all(value == digests[cell][version][key] for key, value in identity.items()), f"Timed counts differ: {name}")
                            summaries[name]["query"] = data
        save(record / "memory-summary.json", {"kind": "memfd-resident-chr22", "disk_benchmark": False,
             "query_rounds": 3, "query_repeat": 10, "annotation_rounds": 3, "annotation_output": "/dev/null",
             "preload_in_timed_interval": False, "results": summaries})
    finally:
        # Only remove proxy entries created by this process; retain all evidence.
        for link in reversed(links):
            link.unlink()
        for fd in descriptors:
            os.close(fd)
        for directory in reversed(dirs):
            directory.rmdir()
        save(record / "resident.cleaned.json", resources(group))


if __name__ == "__main__":
    # The enclosing timeout signals the entire process group; unwind the driver
    # after its active child terminates so memfds and proxy names are released.
    def terminate(signum, frame):
        raise SystemExit(128 + signum)
    signal.signal(signal.SIGTERM, terminate)
    main()
