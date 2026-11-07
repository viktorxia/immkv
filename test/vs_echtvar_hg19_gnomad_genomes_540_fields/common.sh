#!/bin/bash
# Shared path/variable setup for the 540-field experiment scripts.
# Every script sources this file; override any variable in the environment.

# Repository root (two levels up from this directory).
export REPO=$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)
export EXP_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

# Inputs.
export GNOMAD_SITES=${GNOMAD_SITES:-/home/tian/ssd_annotation_database/gnomad/GRCh37/v2.1.1}
export HG001_SOURCE=${HG001_SOURCE:-/mnt/data6t/tian/annotation_database/na12878/GRCh37/HG001_GRCh37_1_22_v4.2.1_benchmark.vcf.gz}
export ECHTVAR_BIN=${ECHTVAR_BIN:-$HOME/projects/immkv/tmp/echtvar/echtvar}
export BCFTOOLS=${BCFTOOLS:-$HOME/miniforge3/envs/transvar/bin/bcftools}

# immkv side: the pre-built 366 GiB full-load library (see ../buildfull for
# how it was built) + the 540-field whitelist. No narrow rebuild: the full
# library already contains every numeric field; the whitelist selects the
# same information the echtvar archives carry.
export FULL_LIB=${FULL_LIB:-$HOME/immkv_test_nvme/gnomad/immkv}

# Scratch and outputs live outside the repository (multi-GB artifacts).
export WORK=${WORK:-$HOME/immkv_test_nvme/test/540_fields}
export QUERY=$WORK/query/hg001.decomp.vcf.gz
export ARCHIVES=$WORK/archives
export DECOMP=$WORK/decomp

mkdir -p "$WORK" "$ARCHIVES" "$DECOMP" "$(dirname "$QUERY")"
