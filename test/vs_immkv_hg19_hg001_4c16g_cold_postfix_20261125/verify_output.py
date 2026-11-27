#!/usr/bin/env python3
"""Stream complete immkv output into the freshly frozen shared value verifier and check v0.0.5 digest equivalence."""
import argparse
import errno
import hashlib
import io
import json
import os
from pathlib import Path
import queue
import re
import select
import signal
import subprocess
import threading
import time
import zlib

ROWS, HISTORICAL_HITS, MAX_LINE = 3933714, 3875208, 32 << 20
EXP_DIR = Path(__file__).resolve().parent
REPO = EXP_DIR.parent.parent
SCRATCH = REPO / "tmp" / EXP_DIR.name
RESULTS = EXP_DIR / "results"
CHUNK = 1 << 20

V005_IMMKV_BASELINES = {
    17: {
        "rows": 3933714,
        "data_bytes": 3418342825,
        "data_sha256": "8e1c55459edc13e1b24236ef9d65b54f3215dc11dd55db97efa5979f5247af63",
    },
    540: {
        "rows": 3933714,
        "data_bytes": 67477268158,
        "data_sha256": "9e54e963a7a9db6063232afa59a58c48f37d999289116a43e16f97bfccd59dcd",
    },
    568: {
        "rows": 3933714,
        "data_bytes": 74419083677,
        "data_sha256": "c1ca1dd2d338f254df4db1e1e9c4de8c9cad927925a4cfc51111e9a0d8c8ee43",
    },
}


def sha(path):
    with Path(path).open("rb") as handle:
        return hashlib.file_digest(handle, "sha256").hexdigest()


def local_path(path, root):
    path = Path(path).resolve()
    if not path.is_relative_to(root.resolve()):
        raise ValueError(f"Writable path must be inside {root}: {path}")
    return path


def reference_paths(media, fields):
    roots = {"hdd": Path(os.environ.get("HDD_ROOT", "/mnt/data6t/tian/immkv_test_hdd")),
             "nvme": Path(os.environ.get("NVME_ROOT", "/home/tian/immkv_test_nvme"))}
    if fields == 568:
        return roots["hdd"] / "test/4c16g_gz_v003/568_fields/echtvar_ref.wk"
    return roots[media] / f"test/4c16g_gz_v003/{fields}_fields/echtvar.pass2.wk"


class FifoInput(io.RawIOBase):
    def __init__(self, path, producer_done, stop):
        super().__init__()
        self.fd = os.open(path, os.O_RDONLY | os.O_NONBLOCK)
        self.producer_done, self.stop = producer_done, stop

    def readable(self):
        return True

    def readinto(self, target):
        while not self.stop.is_set():
            try:
                data = os.read(self.fd, len(target))
                if data:
                    target[:len(data)] = data
                    return len(data)
                if self.producer_done.is_set():
                    return 0
            except BlockingIOError:
                pass
            select.select([self.fd], [], [], 0.05)
            self.stop.wait(0.01)
        raise RuntimeError("Cancelled before complete compressed producer EOF")

    def close(self):
        if not self.closed:
            os.close(self.fd)
        super().close()


class GzipMembers(io.RawIOBase):
    def __init__(self, source, stop):
        super().__init__()
        self.source, self.stop = source, stop
        self.pending, self.decoder = b"", None
        self.members, self.compressed_bytes = 0, 0
        self.compressed_sha256 = hashlib.sha256()
        self.complete = False

    def readable(self):
        return True

    def readinto(self, target):
        if self.complete:
            return 0
        while not self.stop.is_set():
            if not self.pending:
                self.pending = self.source.read(CHUNK)
                self.compressed_bytes += len(self.pending)
                self.compressed_sha256.update(self.pending)
                if not self.pending:
                    if self.decoder is not None or not self.members:
                        raise EOFError("Missing gzip member footer or empty compressed stream")
                    self.complete = True
                    return 0
            if self.decoder is None:
                self.pending = self.pending.lstrip(b"\0")
                if not self.pending:
                    continue
                self.decoder = zlib.decompressobj(16 + zlib.MAX_WBITS)
            data = self.decoder.decompress(self.pending, len(target))
            self.pending = (self.decoder.unused_data if self.decoder.eof
                            else self.decoder.unconsumed_tail)
            if self.decoder.eof:
                self.members += 1
                self.decoder = None
            if data:
                target[:len(data)] = data
                return len(data)
        raise RuntimeError("Cancelled before complete gzip member EOF")


def write_fifo(path, stop):
    while not stop.is_set():
        try:
            return os.open(path, os.O_WRONLY | os.O_NONBLOCK)
        except OSError as exc:
            if exc.errno != errno.ENXIO:
                raise
            stop.wait(0.05)
    raise RuntimeError("Cancelled while waiting for verifier FIFO reader")


def write_all(fd, data, stop):
    remaining = memoryview(data)
    while remaining:
        if stop.is_set():
            raise RuntimeError("Cancelled while writing verifier input")
        try:
            count = os.write(fd, remaining)
            if count == 0:
                raise BrokenPipeError("Verifier input accepted no bytes")
            remaining = remaining[count:]
        except BlockingIOError:
            select.select([], [fd], [], 0.05)


def check_frozen_tools(annotate, verifier):
    for path, name in ((annotate, "annotate"), (verifier, "verifyvalues")):
        expected = SCRATCH / "bin" / name
        if path.resolve() != expected.resolve() or not path.is_file():
            raise ValueError(f"Use the experiment's frozen tool: {expected}")
    preparation = RESULTS / "preparation"
    if (preparation / "exit-status").read_text().strip() != "0":
        raise ValueError("Tool preparation did not succeed")
    manifest = {}
    for line in (preparation / "binaries.sha256").read_text().splitlines():
        digest, name = line.split(maxsplit=1)
        manifest[name.lstrip("*")] = digest
    actual = {"annotate": sha(annotate), "verifyvalues": sha(verifier)}
    for name, digest in actual.items():
        if manifest.get(name) != digest:
            raise ValueError(f"Frozen {name} binary differs from preparation manifest")
    return actual


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--media", required=True, choices=("hdd", "nvme"))
    parser.add_argument("--fields", required=True, type=int, choices=(17, 540, 568))
    parser.add_argument("--result", required=True, type=Path)
    parser.add_argument("--annotate", type=Path, default=SCRATCH / "bin/annotate")
    parser.add_argument("--verifier", type=Path, default=SCRATCH / "bin/verifyvalues")
    parser.add_argument("--reference", type=Path, help="Historical anN.vcf.gz directory")
    args = parser.parse_args()
    result = local_path(args.result, RESULTS)
    result.mkdir(parents=True, exist_ok=True)
    evidence = result / "verification.json"
    with evidence.open("x") as handle:
        json.dump({"status": "started"}, handle)
    facts = {"status": "failed", "media": args.media, "fields": args.fields,
             "expected_rows": ROWS, "expected_field_comparisons": ROWS * args.fields,
             "historical_hits_baseline": HISTORICAL_HITS, "float_tolerance": 1.1e-6,
             "reference_release": "echtvar v0.2.4", "timing_sample": False,
             "v005_immkv_baseline": V005_IMMKV_BASELINES[args.fields]}
    processes, threads, logs, fifos = {}, [], [], []
    errors, streams = queue.Queue(), {}
    stop, producer_done = threading.Event(), threading.Event()
    started = time.monotonic()

    def interrupted(signum, frame):
        raise SystemExit(f"Received signal {signum}")

    signal.signal(signal.SIGTERM, interrupted)
    signal.signal(signal.SIGINT, interrupted)
    try:
        work = local_path(SCRATCH / "verify" / result.name, SCRATCH)
        work.mkdir(parents=True, exist_ok=False)
        roots = {"hdd": Path(os.environ.get("HDD_ROOT", "/mnt/data6t/tian/immkv_test_hdd")),
                 "nvme": Path(os.environ.get("NVME_ROOT", "/home/tian/immkv_test_nvme"))}
        root = roots[args.media]
        query = root / f"test/{args.fields}_fields/query/hg001.decomp.vcf.gz"
        library = root / "test/17_fields/lib17" if args.fields == 17 else root / "gnomad/immkv"
        fields = REPO / f"test/vs_echtvar_hg19_gnomad_genomes_{args.fields}_fields/fields{args.fields}.tsv"
        reference = (args.reference or reference_paths(args.media, args.fields)).resolve()
        references = [reference / f"an{i}.vcf.gz" for i in range(1, 23)]
        facts.update(query=str(query), library=str(library), fields_tsv=str(fields),
                     reference=str(reference), reference_kind="historical echtvar reference + v0.0.5 immkv exact data_sha256")
        for path in (query, fields, *references):
            if not path.is_file():
                raise ValueError(f"Missing input: {path}")
        facts["frozen_binary_sha256"] = check_frozen_tools(args.annotate, args.verifier)
        selected = [line.split(b"\t") for line in fields.read_bytes().splitlines()
                    if line and not line.startswith(b"#")]
        if (len(selected) != args.fields or any(len(cols) < 2 for cols in selected)
                or len({cols[0] for cols in selected}) != args.fields):
            raise ValueError(f"Malformed/duplicate field whitelist or field count differs from {args.fields}")
        facts["selected_fields"] = len(selected)
        float_fields = {b"gnomad_" + cols[0] for cols in selected if cols[1] == b"Float"}
        facts["sha256"] = {str(path): sha(path) for path in (args.annotate, args.verifier, fields, Path(__file__))}
        gz_fifo, iv_fifo, ev_fifo = [work / name for name in ("output.vcf.gz", "immkv.fifo", "echtvar.fifo")]
        for path in (gz_fifo, iv_fifo, ev_fifo):
            os.mkfifo(path, 0o600)
            fifos.append(path)
        minus_one_pattern = re.compile(rb"(?:^|[;\t])(gnomad_[^;=\t]+)=(-1(?:\.0*)?)(?=;|\t|\r?\n|$)")

        def drain(label, inputs, destination):
            records, digest, total, total_bytes, minus_one, fd = [], hashlib.sha256(), 0, 0, 0, None
            streams[label] = {"status": "started", "files": records}
            try:
                raw_fifo = FifoInput(inputs[0], producer_done, stop) if label == "immkv" else None
                try:
                    fd = write_fifo(destination, stop)
                    for path in inputs:
                        count, size, file_digest = 0, 0, hashlib.sha256()
                        record = {"path": str(path), "gzip_crc_eof_ok": False}
                        records.append(record)
                        raw = raw_fifo if raw_fifo is not None else path.open("rb", buffering=0)
                        with raw:
                            decoder = GzipMembers(raw, stop)
                            with io.BufferedReader(decoder, buffer_size=CHUNK) as source:
                                while True:
                                    line = source.readline(MAX_LINE)
                                    if not line:
                                        break
                                    if len(line) >= MAX_LINE or not line.endswith(b"\n"):
                                        raise ValueError(f"Oversized or unterminated line: {path}")
                                    if line.startswith(b"#") or not line.strip():
                                        continue
                                    count += 1
                                    if label == "immkv" and b"=-1" in line:
                                        minus_one += sum(m[1] in float_fields for m in minus_one_pattern.finditer(line))
                                    size += len(line)
                                    digest.update(line)
                                    file_digest.update(line)
                                    write_all(fd, line, stop)
                            if not decoder.complete:
                                raise EOFError(f"Incomplete gzip member EOF: {path}")
                            record.update(rows=count, data_bytes=size, data_sha256=file_digest.hexdigest(),
                                          gzip_members=decoder.members, gzip_crc_eof_ok=True,
                                          compressed_bytes=decoder.compressed_bytes,
                                          compressed_sha256=decoder.compressed_sha256.hexdigest())
                        total += count
                        total_bytes += size
                    if total != ROWS:
                        raise ValueError(f"{label}: {total} rows, expected {ROWS}")
                    streams[label].update(status="complete", rows=total, data_bytes=total_bytes,
                                          data_sha256=digest.hexdigest(), gzip_crc_eof_ok=True,
                                          gzip_members=sum(r["gzip_members"] for r in records))
                    if label == "immkv":
                        streams[label]["selected_float_true_minus_one_candidates"] = minus_one
                finally:
                    if raw_fifo is not None:
                        raw_fifo.close()
            except BaseException as exc:
                streams[label].update(status="failed", rows=total, error=f"{type(exc).__name__}: {exc}")
                errors.put(f"{label}: {type(exc).__name__}: {exc}")
            finally:
                if fd is not None:
                    os.close(fd)

        commands = {
            "annotate": [str(args.annotate), "-in", str(query), "-lib", str(library), "-fields", str(fields),
                         "-threads", "4", "-batch", "4096", "-out", str(gz_fifo)],
            "verifier": [str(args.verifier), "-ev", str(ev_fifo), "-iv", str(iv_fifo),
                         "-fields", str(fields), "-tol", "1.1e-6", "-max-report", "20"]}
        facts["commands"] = commands
        for label, inputs, output in (("immkv", [gz_fifo], iv_fifo), ("echtvar", references, ev_fifo)):
            thread = threading.Thread(target=drain, args=(label, inputs, output), name=f"verify-{label}", daemon=True)
            threads.append(thread)
            thread.start()
        for label, command in commands.items():
            stdout, stderr = [open(result / f"{label}.{suffix}", "xb") for suffix in ("stdout", "stderr")]
            logs.extend((stdout, stderr))
            processes[label] = subprocess.Popen(command, stdout=stdout, stderr=stderr, start_new_session=True)
        deadline = started + int(os.environ.get("RUN_TIMEOUT_SECONDS", "7200"))
        while True:
            if processes["annotate"].poll() is not None:
                producer_done.set()
            if not errors.empty():
                raise RuntimeError(errors.get())
            failed = [(label, process.returncode) for label, process in processes.items()
                      if process.poll() not in (None, 0)]
            if failed:
                raise RuntimeError(f"Subprocess failed: {failed}")
            if all(process.poll() == 0 for process in processes.values()) and not any(t.is_alive() for t in threads):
                break
            if time.monotonic() > deadline:
                raise TimeoutError("Full-output verification deadline exceeded")
            stop.wait(0.05)
        if not errors.empty():
            raise RuntimeError(errors.get())
        for handle in logs:
            handle.flush()
        annotate = (result / "annotate.stdout").read_text()
        match = re.search(r"rows (\d+), hits (\d+) ", annotate)
        if not match or int(match[1]) != ROWS or not 0 <= int(match[2]) <= ROWS:
            raise ValueError("Annotation rows/hits statistics missing or inconsistent")
        facts["annotate_statistics"] = {"rows": int(match[1]), "hits": int(match[2]),
                                        "hits_equal_historical_baseline": int(match[2]) == HISTORICAL_HITS}
        report = (result / "verifier.stdout").read_text()
        match = re.search(r"rows (\d+), field comparisons (\d+), misses agreed (\d+), value mismatches (\d+), hit disagreements (\d+)", report)
        overflow = re.search(r"echtvar i32-overflow value losses \(documented limitation\): (\d+)", report)
        floats = re.search(r"float comparisons (\d+), max \|echtvar-immkv\| observed: ([\deE.+-]+)", report)
        if not match or not overflow or not floats or "VERIFY_VALUES OK:" not in report:
            raise ValueError("Shared verifier summary missing")
        rows, comparisons, misses, mismatches, disagreements = map(int, match.groups())
        facts["comparison"] = dict(rows=rows, field_comparisons=comparisons, misses_agreed=misses,
                                   value_mismatches=mismatches, hit_disagreements=disagreements,
                                   echtvar_i32_overflow_losses=int(overflow[1]), float_comparisons=int(floats[1]),
                                   max_float_difference=float(floats[2]))
        if (rows, comparisons, mismatches, disagreements) != (ROWS, ROWS * args.fields, 0, 0):
            raise ValueError("Shared verifier totals or disagreement counts unexpected")
        if any(streams.get(label, {}).get("status") != "complete" for label in ("immkv", "echtvar")):
            raise ValueError("Independent complete-stream checks missing")
        v005 = V005_IMMKV_BASELINES[args.fields]
        immkv_stream = streams["immkv"]
        digest_match = (immkv_stream.get("data_sha256") == v005["data_sha256"]
                        and immkv_stream.get("data_bytes") == v005["data_bytes"]
                        and immkv_stream.get("rows") == v005["rows"])
        facts["v005_digest_equal"] = digest_match
        if not digest_match:
            raise ValueError(f"immkv output stream differs from v0.0.5 baseline: got sha256={immkv_stream.get('data_sha256')} bytes={immkv_stream.get('data_bytes')}")
        facts["status"] = "passed"
    except BaseException as exc:
        facts["error"] = f"{type(exc).__name__}: {exc}"
    finally:
        stop.set()
        producer_done.set()
        for process in processes.values():
            if process.poll() is None:
                try:
                    os.killpg(process.pid, signal.SIGTERM)
                except ProcessLookupError:
                    pass
        for process in processes.values():
            try:
                process.wait(timeout=3)
            except subprocess.TimeoutExpired:
                try:
                    os.killpg(process.pid, signal.SIGKILL)
                except ProcessLookupError:
                    pass
                process.wait(timeout=3)
        for thread in threads:
            thread.join(timeout=5)
        remaining = [thread.name for thread in threads if thread.is_alive()]
        if remaining:
            facts.update(status="failed", error=f"Reader threads did not terminate: {remaining}")
        while not errors.empty():
            facts.setdefault("stream_errors", []).append(errors.get())
        if facts.get("stream_errors"):
            facts["status"] = "failed"
        for handle in logs:
            handle.close()
        facts["exit_status"] = {label: process.returncode for label, process in processes.items()}
        facts["streams"] = streams
        facts["elapsed_seconds_correctness_only"] = time.monotonic() - started
        for path in fifos:
            path.unlink(missing_ok=True)
        partial = result / "verification.json.partial"
        with partial.open("x") as handle:
            json.dump(facts, handle, indent=2)
            handle.write("\n")
        partial.replace(evidence)
    print(json.dumps({"status": facts["status"], "evidence": str(evidence), "error": facts.get("error")}))
    return 0 if facts["status"] == "passed" else 1


if __name__ == "__main__":
    raise SystemExit(main())
