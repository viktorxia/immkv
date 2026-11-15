#!/bin/bash
# Streaming outputs leave no large texts to remove; preserve every record.
set -euo pipefail
printf '%s\n' 'No retained annotation texts: timing uses /dev/null and correctness streams through FIFOs.' \
    'No files removed; release binaries, frozen sources, results, libraries, and archives are retained.'
