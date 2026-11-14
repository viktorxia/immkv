# Query design review and implementation rationale

## Decision

Keep StreamOptions.Columns as a logical field-selection contract and keep the
schema-shaped v003 record as the result representation. Record-order and
columnar paths must return the same selected fields. The design does not need
a public API or on-disk format change. Fixing the old record-order path that
ignored Columns remains necessary.

## Implemented query changes

- Normalize a validated explicit all-field selection to the default path.
- Compile an immutable recordProjection once per streaming handle.
- Project subsets in the existing decompression workers, iterating only present
  bitmap bits and copying contiguous selected payload runs once.
- Validate skipped payloads and trailing bytes during subset projection.
- Decode single-value frames directly into unused output storage when no
  projection is needed; use separate reusable scratch for subset projection.
- Retain each call's output independently; cap returned views and reuse duplicate
  keys' already decoded/projected result without modifying their shared source.
- Return compressed input buffers when no raw result view refers to them.
- Copy columnar decoded spans directly into final records and validate selected
  column tails plus bitmap padding.
- Keep record-order and columnar decompression within the same Threads budget.
- Avoid creating an io_uring when the direct-read prerequisites are absent.
- Correct O_DIRECT aligned-slice capacity bounds and reject undersized pool
  returns, including nil or decoder-grown capacities that do not match a bucket.

Implementation: stream.go, projection.go, colread.go, iouring_linux.go.
Removed the serial post-decompression projectRecord pass from record.go.
README.md and README.cn.md document read-only call-owned GetBatch results,
complete encoded-key sorting, and the default/all-selected fast path.

## Validation boundary

Unprojected record-order results follow the existing default-query behavior:
typed record validation is performed when the caller constructs NewRecord.
Explicit all-field selection now has exactly that behavior. Partial projection
still validates even unselected typed payloads because record boundaries must
be parsed correctly. I/O, frame decompression, key structure, projection and
columnar decoding errors are propagated rather than converted into misses.

## Regression coverage

projection_test.go covers wide schemas, all field types, missing/empty values,
malformed skipped payloads, deterministic fuzz seeds, raw/per-value/packed
framing, duplicate keys, concurrent queries, retained output, large values and
buffer-pool capacity safety. colread_projection_test.go covers duplicate spans,
missing values, array/Flag/String fields, tail corruption, and padding bits.
iouring_linux_test.go covers buffer alignment/capacity including large buffers.

Both default and CGO_ENABLED=0 full builds and tests, go vet, the race suite,
and Darwin/Windows/FreeBSD pure-Go builds passed. Sources for these claims are
results/code-checks/*.log and results/code-checks/finished. Runtime results bind
to the exact retained sources in results/build.after/; the production source
was not modified during measurement. No codec frame-format or magic changes
were introduced.

## Limits

This implementation does not establish a theoretical maximum throughput.
The optimized full540 profile is dominated by zstd decompression and I/O
syscalls, with a smaller residual projection cost. Large-library wall-clock
comparisons on this machine encountered changing NVMe service time; the raw
observations remain in results/io-observation.md. A separate memory-resident
chr22 diagnostic isolates query work from that device variability.
