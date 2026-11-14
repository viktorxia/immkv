#!/usr/bin/env python3
"""Stream complete annotation outputs through the shared verifier and SHA-256."""
import argparse
import gzip
import hashlib
import json
import os
from pathlib import Path
import queue
import re
import signal
import subprocess
import threading
import time


def count_query(path):
    opener = gzip.open if path.suffix in (".gz", ".bgz") else open
    with opener(path, "rb") as stream:
        return sum(1 for line in stream if line.strip() and not line.startswith(b"#"))


def run(args):
    fields = [line.split("\t")[0] for line in args.fields.read_text().splitlines()
              if line and not line.startswith("#")]
    if not fields or len(fields) != len(set(fields)):
        raise ValueError("Empty or duplicate fields")
    expected_rows = count_query(args.query)
    if expected_rows <= 0:
        raise ValueError("Empty query")
    args.work.mkdir()
    processes, streams, threads = [], [], []
    outcomes = queue.Queue()
    stop = threading.Event()
    fifo_paths = []
    versions = (args.reference, args.candidate)

    def pump(version):
        digest = hashlib.sha256()
        rows = total_bytes = 0
        # Nonblocking opens plus a stop event keep error cleanup bounded even
        # when a producer fails before opening its FIFO.
        try:
            raw = args.work / ("raw." + version)
            checked = args.work / ("verify." + version)
            reader = os.open(raw, os.O_RDONLY | os.O_NONBLOCK)
            writer = None
            try:
                while not stop.is_set():
                    try:
                        writer = os.open(checked, os.O_WRONLY | os.O_NONBLOCK)
                        break
                    except OSError as exc:
                        if exc.errno != 6:  # ENXIO: verifier has not opened yet.
                            raise
                        stop.wait(0.01)
                if writer is None:
                    raise RuntimeError("Verification stopped before FIFO connected")
                # Read complete rows so each output's count is independent of
                # the legacy verifier's scanner and its exception categories.
                pending = bytearray()
                seen_data = False
                while not stop.is_set():
                    try:
                        block = os.read(reader, 1 << 20)
                    except BlockingIOError:
                        stop.wait(0.001)
                        continue
                    if not block:
                        producer = producers[version]
                        if producer.poll() is None:
                            stop.wait(0.01)
                            continue
                        if producer.returncode != 0:
                            raise RuntimeError(f"Producer {version} exited {producer.returncode}")
                        if pending:
                            raise RuntimeError(f"Unterminated final row from {version}")
                        if not seen_data:
                            raise RuntimeError(f"Empty output from {version}")
                        break
                    seen_data = True
                    digest.update(block)
                    total_bytes += len(block)
                    pending.extend(block)
                    last_newline = pending.rfind(b"\n")
                    if last_newline >= 0:
                        complete = pending[:last_newline + 1]
                        rows += sum(1 for line in complete.splitlines() if line and not line.startswith(b"#"))
                        del pending[:last_newline + 1]
                    view = memoryview(block)
                    while view and not stop.is_set():
                        try:
                            count = os.write(writer, view)
                            view = view[count:]
                        except BlockingIOError:
                            stop.wait(0.001)
                if stop.is_set():
                    raise RuntimeError("Verification stopped")
            finally:
                os.close(reader)
                if writer is not None:
                    os.close(writer)
            outcomes.put((version, {"sha256": digest.hexdigest(), "rows": rows, "bytes": total_bytes}, None))
        except BaseException as exc:
            outcomes.put((version, None, str(exc)))

    def interrupt(signum, frame):
        raise InterruptedError(f"Verification interrupted by signal {signum}")

    signal.signal(signal.SIGTERM, interrupt)
    signal.signal(signal.SIGINT, interrupt)
    try:
        for version in versions:
            for prefix in ("raw", "verify"):
                fifo = args.work / (prefix + "." + version)
                os.mkfifo(fifo)
                fifo_paths.append(fifo)
        log = (args.record / "verify.log").open("xb")
        streams.append(log)
        verify = subprocess.Popen([
            str(args.binaries / args.reference / "verifyvalues"),
            "-ev", str(args.work / ("verify." + args.reference)),
            "-iv", str(args.work / ("verify." + args.candidate)),
            "-fields", str(args.fields), "-tol", "0",
        ], stdout=log, stderr=subprocess.STDOUT)
        processes.append(verify)
        producers = {}
        for version in versions:
            stats = (args.record / (version + ".stats")).open("xb")
            stderr = (args.record / (version + ".stderr")).open("xb")
            streams.extend((stats, stderr))
            producer = subprocess.Popen([
                str(args.binaries / version / "annotate"), "-in", str(args.query),
                "-lib", str(args.library), "-fields", str(args.fields),
                "-threads", "4", "-batch", "4096",
                "-out", str(args.work / ("raw." + version)),
            ], stdout=stats, stderr=stderr)
            producers[version] = producer
            processes.append(producer)
        for version in versions:
            thread = threading.Thread(target=pump, args=(version,), daemon=True)
            thread.start()
            threads.append(thread)
        results = {}
        while len(results) < 2 or any(p.poll() is None for p in processes):
            for process in processes:
                if process.poll() not in (None, 0):
                    raise RuntimeError(f"Process exited {process.returncode}: {process.args[0]}")
            try:
                version, result, error = outcomes.get(timeout=0.05)
                if error:
                    raise RuntimeError(f"Stream {version}: {error}")
                results[version] = result
            except queue.Empty:
                pass
        for stream in streams:
            stream.flush()
        if any(item["rows"] != expected_rows for item in results.values()):
            raise ValueError(f"Stream row counts differ from query: {expected_rows}, {results}")
        if results[args.reference] != results[args.candidate]:
            raise ValueError(f"Complete output streams differ: {results}")
        for version in versions:
            stats = (args.record / (version + ".stats")).read_text()
            matches = re.findall(r"rows (\d+), hits (\d+)", stats)
            if len(matches) != 1 or int(matches[0][0]) != expected_rows:
                raise ValueError(f"Incomplete producer statistics: {version}")
            results[version]["producer_rows"] = int(matches[0][0])
            results[version]["producer_hits"] = int(matches[0][1])
            results[version]["producer_exit_status"] = producers[version].returncode
        if results[args.reference]["producer_hits"] != results[args.candidate]["producer_hits"]:
            raise ValueError("Producer hit counts differ")
        text = (args.record / "verify.log").read_text()
        matches = re.findall(r"rows (\d+), field comparisons (\d+), misses agreed (\d+), value mismatches (\d+), hit disagreements (\d+)", text)
        if len(matches) != 1:
            raise ValueError("Missing complete verifier summary")
        rows, comparisons, misses, mismatches, disagreements = map(int, matches[0])
        if (rows != expected_rows or comparisons != expected_rows * len(fields)
                or mismatches or disagreements or "VERIFY_VALUES OK" not in text):
            raise ValueError("Verifier counts or value comparisons failed")
        evidence = {
            "status": "passed", "query": str(args.query), "fields": str(args.fields),
            "library": str(args.library), "expected_rows": expected_rows,
            "field_count": len(fields), "field_comparisons": comparisons,
            "value_mismatches": mismatches, "hit_disagreements": disagreements,
            "tolerance": 0, "exact_stream_sha256_equal": True,
            "verifier_exit_status": verify.returncode, "streams": results,
        }
        with (args.record / "verification.json").open("x") as target:
            json.dump(evidence, target, indent=2, sort_keys=True)
            target.write("\n")
        for version, item in results.items():
            with (args.record / (version + ".sha256")).open("x") as target:
                target.write(item["sha256"] + "  complete-annotation-stream\n")
        print(f"EXACT_STREAM_SHA256_OK rows={expected_rows} field_comparisons={comparisons}")
    finally:
        stop.set()
        for process in processes:
            if process.poll() is None:
                process.terminate()
        for process in processes:
            try:
                process.wait(timeout=3)
            except subprocess.TimeoutExpired:
                process.kill()
                process.wait()
        for thread in threads:
            thread.join(timeout=3)
        for stream in streams:
            stream.close()
        for fifo in fifo_paths:
            fifo.unlink(missing_ok=True)
        args.work.rmdir()


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    for name in ("query", "library", "fields", "binaries", "record", "work"):
        parser.add_argument("--" + name, type=Path, required=True)
    for name in ("reference", "candidate"):
        parser.add_argument("--" + name, choices=("old", "before", "after"), required=True)
    args = parser.parse_args()
    if args.reference == args.candidate:
        parser.error("two different versions are required")
    run(args)


if __name__ == "__main__":
    main()
