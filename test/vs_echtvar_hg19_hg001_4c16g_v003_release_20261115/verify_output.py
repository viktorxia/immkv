#!/usr/bin/env python3
"""Full-output correctness only; run through verify.sh's authorized limited scope."""
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

ROWS, HITS, MAX_LINE = 3933714, 3875208, 32 << 20
REPO = Path(__file__).resolve().parents[2]
NAME = Path(__file__).resolve().parent.name
SCRATCH = REPO / "tmp" / NAME
RETAINED = REPO / "tmp/vs_immkv_hg19_hg001_4c16g_projection_20261115/bin/after/verifyvalues"


def sha(path):
    with open(path, "rb") as handle:
        return hashlib.file_digest(handle, "sha256").hexdigest()


def local_path(path, root):
    path = Path(path).resolve()
    if not path.is_relative_to(root.resolve()):
        raise ValueError(f"Writable path must be inside {root}: {path}")
    return path


def main():
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--media", required=True, choices=("hdd", "nvme"))
    parser.add_argument("--fields", required=True, type=int, choices=(17, 540, 568))
    parser.add_argument("--result", required=True, type=Path)
    parser.add_argument("--annotate", type=Path, default=SCRATCH / "bin/annotate")
    parser.add_argument("--verifier", type=Path, default=SCRATCH / "bin/verifyvalues")
    args = parser.parse_args()
    result = local_path(args.result, REPO / "test" / NAME / "results")
    result.mkdir(parents=True, exist_ok=True)  # limited() reserves the parent record.
    evidence = result / "verification.json"
    with evidence.open("x") as handle:
        json.dump({"status": "started"}, handle)
    work = local_path(SCRATCH / "verify" / result.name, SCRATCH)
    work.mkdir(parents=True, exist_ok=False)
    facts = {"status": "failed", "media": args.media, "fields": args.fields,
             "expected_rows": ROWS, "expected_hits": HITS, "float_tolerance": 1.1e-6,
             "reference_release": "echtvar v0.2.4", "timing_sample": False,
             "marker_collision_accounting": "Shared verifier includes true float -1 marker collisions in misses agreed.",
             "float_tolerance_basis": "echtvar v0.2.4 quantization truncation plus six-decimal print truncation and f32 slack"}
    processes, threads, logs, fifos = {}, [], [], []
    errors, streams = queue.Queue(), {}
    stop = threading.Event()
    started = time.monotonic()
    def interrupted(signum, frame):
        raise SystemExit(f"Received signal {signum}")
    signal.signal(signal.SIGTERM, interrupted)
    signal.signal(signal.SIGINT, interrupted)
    try:
        roots = {"hdd": Path(os.environ.get("HDD_ROOT", "/mnt/data6t/tian/immkv_test_hdd")),
                 "nvme": Path(os.environ.get("NVME_ROOT", "/home/tian/immkv_test_nvme"))}
        root = roots[args.media]
        query = root / f"test/{args.fields}_fields/query/hg001.decomp.vcf.gz"
        library = root / "test/17_fields/lib17" if args.fields == 17 else root / "gnomad/immkv"
        fields = REPO / f"test/vs_echtvar_hg19_gnomad_genomes_{args.fields}_fields/fields{args.fields}.tsv"
        reference = (roots["hdd"] / "test/4c16g_gz_v003/568_fields/echtvar_ref.wk" if args.fields == 568
                     else root / f"test/4c16g_gz_v003/{args.fields}_fields/echtvar.pass2.wk")
        references = [reference / f"an{i}.vcf.gz" for i in range(1, 23)]
        facts.update(query=str(query), library=str(library), fields_tsv=str(fields),
                     reference_kind="historical correctness reference; not a current performance sample",
                     reference_media="hdd" if args.fields == 568 else args.media)
        for path in [query, fields, args.annotate, args.verifier, *references]:
            if not path.is_file():
                raise ValueError(f"Missing input: {path}")
        selected = [line for line in fields.read_bytes().splitlines() if line and not line.startswith(b"#")]
        float_fields = {b"gnomad_" + line.split(b"\t")[0] for line in selected
                        if line.split(b"\t")[1] == b"Float"}
        if len(selected) != args.fields:
            raise ValueError(f"Field whitelist has {len(selected)} entries, expected {args.fields}")
        facts["sha256"] = {str(path): sha(path) for path in (args.annotate, args.verifier, fields, Path(__file__))}
        gz_fifo, iv_fifo, ev_fifo = [work / name for name in ("output.vcf.gz", "immkv.fifo", "echtvar.fifo")]
        for path in (gz_fifo, iv_fifo, ev_fifo):
            os.mkfifo(path, 0o600)
            fifos.append(path)

        minus_one_pattern = re.compile(rb"(?:^|[;\t])(gnomad_[^;=\t]+)=(-1(?:\.0*)?)(?=;|\t|\r?\n|$)")

        def drain(label, inputs, destination):
            records, digest, total, minus_one = [], hashlib.sha256(), 0, 0
            try:
                # Blocking opens are paired by verifier and annotate; daemon threads plus
                # process supervision handle failed startup without an open() deadlock.
                with destination.open("wb", buffering=1 << 20) as output:
                    for path in inputs:
                        count, size = 0, 0
                        file_digest = hashlib.sha256()
                        with path.open("rb") as raw, gzip.GzipFile(fileobj=raw) as source:
                            while not stop.is_set():
                                line = source.readline(MAX_LINE)
                                if not line:
                                    break  # GzipFile reached EOF and checked every member's CRC/ISIZE.
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
                                output.write(line)
                        if stop.is_set():
                            raise RuntimeError("Cancelled before complete gzip EOF")
                        total += count
                        records.append({"path": str(path), "rows": count, "data_bytes": size,
                                        "data_sha256": file_digest.hexdigest(), "gzip_crc_eof_ok": True})
                streams[label] = {"rows": total, "data_sha256": digest.hexdigest(), "files": records}
                if label == "immkv":
                    streams[label]["selected_float_true_minus_one_candidates"] = minus_one
                if total != ROWS:
                    raise ValueError(f"{label}: {total} rows, expected {ROWS}")
            except BaseException as exc:
                errors.put(f"{label}: {type(exc).__name__}: {exc}")

        commands = {
            "annotate": [str(args.annotate), "-in", str(query), "-lib", str(library), "-fields", str(fields),
                         "-threads", "4", "-batch", "4096", "-out", str(gz_fifo)],
            "verifier": [str(args.verifier), "-ev", str(ev_fifo), "-iv", str(iv_fifo),
                         "-fields", str(fields), "-tol", "1.1e-6", "-max-report", "20"]}
        facts["commands"] = commands
        for label, inputs, output in (("immkv", [gz_fifo], iv_fifo), ("echtvar", references, ev_fifo)):
            thread = threading.Thread(target=drain, args=(label, inputs, output), daemon=True)
            threads.append(thread)
            thread.start()
        for label, command in commands.items():
            stdout, stderr = [open(result / f"{label}.{suffix}", "xb") for suffix in ("stdout", "stderr")]
            logs.extend((stdout, stderr))
            processes[label] = subprocess.Popen(command, stdout=stdout, stderr=stderr, start_new_session=True)
        deadline = started + int(os.environ.get("RUN_TIMEOUT_SECONDS", "7200"))
        while True:
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
            time.sleep(0.1)
        if not errors.empty():
            raise RuntimeError(errors.get())
        for handle in logs:
            handle.flush()
        annotate = (result / "annotate.stdout").read_text()
        match = re.search(r"rows (\d+), hits (\d+) ", annotate)
        if not match or tuple(map(int, match.groups())) != (ROWS, HITS):
            raise ValueError("Annotation rows/hits statistics missing or unexpected")
        facts["annotate_statistics"] = {"rows": ROWS, "hits": HITS}
        report = (result / "verifier.stdout").read_text()
        match = re.search(r"rows (\d+), field comparisons (\d+), misses agreed (\d+), value mismatches (\d+), hit disagreements (\d+)", report)
        overflow = re.search(r"echtvar i32-overflow value losses \(documented limitation\): (\d+)", report)
        if not match or not overflow or "VERIFY_VALUES OK:" not in report:
            raise ValueError("Shared verifier summary missing")
        rows, comparisons, misses, mismatches, disagreements = map(int, match.groups())
        facts["comparison"] = dict(rows=rows, field_comparisons=comparisons, misses_agreed=misses,
                                   value_mismatches=mismatches, hit_disagreements=disagreements,
                                   echtvar_i32_overflow_losses=int(overflow[1]))
        if (rows, comparisons, mismatches, disagreements) != (ROWS, ROWS * args.fields, 0, 0):
            raise ValueError("Shared verifier totals or disagreement counts unexpected")
        facts["status"] = "passed"
    except BaseException as exc:
        facts["error"] = f"{type(exc).__name__}: {exc}"
    finally:
        stop.set()
        for process in processes.values():
            if process.poll() is None:
                os.killpg(process.pid, signal.SIGTERM)
        for process in processes.values():
            try:
                process.wait(timeout=3)
            except subprocess.TimeoutExpired:
                os.killpg(process.pid, signal.SIGKILL)
                process.wait()
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
