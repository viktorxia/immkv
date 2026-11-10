#!/bin/bash
# Shared path/variable setup for the 540-field HDD experiment.
# Every script sources this file; override any variable in the environment.
#
# This experiment reruns the 540-field comparison of
# ../vs_echtvar_hg19_gnomad_genomes_540_fields with the SAME artifacts copied
# onto a 7200 rpm HDD. Nothing is rebuilt for the timed runs;
# copy_artifacts.sh copies the NVMe-built library and archives byte for byte
# and verifies file counts and byte totals.

# Repository root (two levels up from this directory).
export REPO=$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd)
export EXP_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)

# Inputs.
export GNOMAD_SITES=${GNOMAD_SITES:-/home/tian/ssd_annotation_database/gnomad/GRCh37/v2.1.1}
export HG001_SOURCE=${HG001_SOURCE:-/mnt/data6t/tian/annotation_database/na12878/GRCh37/HG001_GRCh37_1_22_v4.2.1_benchmark.vcf.gz}
export ECHTVAR_BIN=${ECHTVAR_BIN:-$HOME/projects/immkv/tmp/echtvar/echtvar}
export BCFTOOLS=${BCFTOOLS:-$HOME/miniforge3/envs/transvar/bin/bcftools}

# HDD work root and the NVMe source of the copied artifacts.
export HDD_ROOT=${HDD_ROOT:-/mnt/data6t/tian/immkv_test_hdd}
export NVME_ROOT=${NVME_ROOT:-$HOME/immkv_test_nvme}

# immkv side: the pre-built 366 GiB full-load library (copy under HDD_ROOT)
# plus the 540-field whitelist; no narrow rebuild.
export FULL_LIB=${FULL_LIB:-$HDD_ROOT/gnomad/immkv}

# Scratch and outputs live outside the repository (multi-GB artifacts).
export WORK=${WORK:-$HDD_ROOT/test/540_fields}
export QUERY=$WORK/query/hg001.decomp.vcf.gz
export ARCHIVES=$WORK/archives
export DECOMP=$WORK/decomp

mkdir -p "$WORK" "$ARCHIVES" "$DECOMP" "$(dirname "$QUERY")"
