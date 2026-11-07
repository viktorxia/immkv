#!/usr/bin/env python3
"""genfields.py -- derive field lists from a gnomAD sites VCF header.

Reads the ##INFO declarations (in declared order) from one gnomAD sites file
and emits two artifacts used by the experiment scripts:

  <prefix>.tsv   field metadata table: one line per field, tab-separated:
                 name, Type (Integer|Float|String|Flag), Number (A, 1, ., ...).
                 Consumed by annotate (flag/Number=A semantics) and
                 buildnarrow (field selection + Number=A element picking).
  <prefix>.json  echtvar encode config.

Modes (--mode):
  numeric   only Integer/Float fields. Integer -> bare entry; Float ->
            multiplier 2e6 + zigzag (the quantization used by all previous
            echtvar baselines; error <= 5e-7 by construction).
  all       every field. String/Flag fields get bare entries too; echtvar
            infers the encoding from the data itself and routes non-numeric
            values into its categorical dictionary path.

--keep restricts output to the field names listed in a file (one per line),
preserving header order; used for the curated 17-field subset.

Aliases are always "gnomad_" + field name, matching the output prefix of
annotate, so annotated columns line up across tools. Requires only python3
(bgzf blocks are standard gzip members, so gzip.open reads them fine).
"""

import argparse
import gzip
import json
import sys


def load_header(path):
    fields = []
    # Detect gzip (covers plain .gz and bgzf .bgz, both start with 1f 8b) by
    # magic bytes; extension matching is fragile.
    with open(path, "rb") as probe:
        gzipped = probe.read(2) == b"\x1f\x8b"
    op = gzip.open if gzipped else open
    with op(path, "rt") as f:
        for line in f:
            if not line.startswith("#"):
                break
            if not line.startswith("##INFO=<ID="):
                continue
            rest = line[len("##INFO=<ID="):].rstrip().rstrip(">")
            eq = rest.find(",")
            if eq < 0:
                continue
            info = {"name": rest[:eq], "typ": "", "number": ""}
            for attr in split_attrs(rest[eq + 1:]):
                if attr.startswith("Number="):
                    info["number"] = attr[len("Number="):]
                elif attr.startswith("Type="):
                    info["typ"] = attr[len("Type="):]
            fields.append(info)
    return fields


def split_attrs(s):
    """Split 'Number=A,Type=Integer,Description="a, b"' on commas that are
    not inside a quoted string."""
    out, depth, start = [], False, 0
    for i, ch in enumerate(s):
        if ch == '"':
            depth = not depth
        elif ch == "," and not depth:
            out.append(s[start:i])
            start = i + 1
    out.append(s[start:])
    return out


def main():
    ap = argparse.ArgumentParser(description=__doc__)
    ap.add_argument("--sites", required=True, help="path to one gnomAD sites vcf.bgz (header source)")
    ap.add_argument("--out", required=True, help="output path prefix; writes <prefix>.tsv and <prefix>.json")
    ap.add_argument("--mode", choices=["numeric", "all"], default="numeric")
    ap.add_argument("--keep", help="optional file of field names to keep (one per line)")
    args = ap.parse_args()

    fields = load_header(args.sites)
    if args.keep:
        want = set()
        with open(args.keep) as f:
            for ln in f:
                ln = ln.strip()
                if ln and not ln.startswith("#"):
                    want.add(ln)
        fields = [fd for fd in fields if fd["name"] in want]
        if len(fields) != len(want):
            print(f"error: --keep lists {len(want)} names but only {len(fields)} matched the header", file=sys.stderr)
            sys.exit(1)

    # Mode filtering applies to BOTH artifacts: the json is echtvar's encode
    # config and the tsv is the annotate/buildnarrow whitelist -- they must
    # describe the same field set. (The first version wrote the full census
    # to the tsv regardless of mode, silently widening the immkv side.)
    cfg = []
    excluded = []
    selected = []
    for fd in fields:
        entry = {"field": fd["name"], "alias": "gnomad_" + fd["name"]}
        if fd["typ"] == "Float":
            entry["multiplier"] = 2000000
            entry["zigzag"] = True
        if fd["typ"] in ("Integer", "Float") or args.mode == "all":
            cfg.append(entry)
            selected.append(fd)
        else:
            excluded.append(fd["name"])

    with open(args.out + ".tsv", "w") as f:
        f.write("# field\tType\tNumber\n")
        for fd in selected:
            f.write(f"{fd['name']}\t{fd['typ']}\t{fd['number']}\n")
    with open(args.out + ".json", "w") as f:
        json.dump(cfg, f, separators=(",", ":"))
        f.write("\n")

    counts = {}
    for fd in fields:
        counts[fd["typ"]] = counts.get(fd["typ"], 0) + 1
    print(f"header fields: {len(fields)}")
    for t in ("Integer", "Float", "String", "Flag"):
        if counts.get(t):
            print(f"  {t:<8} {counts[t]}")
    if excluded:
        print(f"excluded from numeric config: {len(excluded)}")
        for n in sorted(excluded):
            print(f"  {n}")
    print(f"wrote {args.out}.tsv and {args.out}.json ({len(cfg)} entries, mode={args.mode})")


if __name__ == "__main__":
    main()
