#!/bin/bash
# Shared setup for the controlled rerun: 4 physical cores, 16G cgroup,
# gzip output on both sides, one sample at a time.
#
# Every timed cell runs as its own systemd scope (MemoryMax=16G,
# MemorySwapMax=0) and is pinned to CPUs 0-3. On this machine CPUs 0-3 are
# four distinct physical cores; their SMT siblings are 8-11.
#
# Artifacts (libraries, archives, query) are the pre-existing trees under
# HDD_ROOT and NVME_ROOT; nothing is rebuilt. New outputs go under
# <root>/test/4c16g_gz/ and never overwrite the old records.
set -euo pipefail

export REPO=$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)
export EXP_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

export ECHTVAR_BIN=${ECHTVAR_BIN:-$HOME/projects/immkv/tmp/echtvar/echtvar}
export BCFTOOLS=${BCFTOOLS:-$HOME/miniforge3/envs/transvar/bin/bcftools}

export HDD_ROOT=${HDD_ROOT:-/mnt/data6t/tian/immkv_test_hdd}
export NVME_ROOT=${NVME_ROOT:-$HOME/immkv_test_nvme}

# Controlled protocol.
export MEM_MAX=${MEM_MAX:-16G}
export CPUS=${CPUS:-0-3}
export THREADS=${THREADS:-4}
export BATCH=${BATCH:-4096}
export ECHTVAR_PAR=${ECHTVAR_PAR:-4}

media_root() {
    case "$1" in
    hdd) printf '%s' "$HDD_ROOT" ;;
    nvme) printf '%s' "$NVME_ROOT" ;;
    *)
        echo "unknown media: $1" >&2
        return 1
        ;;
    esac
}

# Existing per-field experiment tree: query and archives live there.
source_dir() { printf '%s/test/%s_fields' "$(media_root "$1")" "$2"; }

# New work tree for this experiment.
work_dir() { printf '%s/test/4c16g_gz/%s_fields' "$(media_root "$1")" "$2"; }

fields_tsv() {
    printf '%s/../vs_echtvar_hg19_gnomad_genomes_%s_fields/fields%s.tsv' \
        "$EXP_DIR" "$1" "$1"
}

library() {
    case "$2" in
    17) printf '%s/test/17_fields/lib17' "$(media_root "$1")" ;;
    540 | 568) printf '%s/gnomad/immkv' "$(media_root "$1")" ;;
    *)
        echo "unknown field set: $2" >&2
        return 1
        ;;
    esac
}
