#!/bin/bash
# Pilot: encode ONE real gnomAD chromosome (chr22, the smallest) with all 568
# fields, capturing cardinalities and resources. This is what established the
# experiment's facts before committing to the genome-wide build:
#
#   - encode does NOT fail: rc=0, 22 min, 5.9 GiB peak RSS, 1.36 GB archive
#     (pilot_encode_chr22.log is the verbatim record)
#   - the categorical "dictionary" degenerates on high-cardinality columns:
#     dp_hist_all_bin_freq 3.26M unique values, vep 3.05M, gq_hist_all 3.49M
#     on chr22 alone (3.57M variants) -- near record-count, i.e. stored
#     verbatim; only popmax/allele_type/VQSR_culprit style fields are true
#     dictionaries (3-8 values)
#
# The pilot archive seeds the genome-wide build (chr22.echtvar.zip).
set -euo pipefail

HERE=$(cd "$(dirname "$0")" && pwd)
WORK=${WORK:-$HOME/immkv_test_nvme/test/568_fields}
SITES22=${SITES22:-/home/tian/ssd_annotation_database/gnomad/GRCh37/v2.1.1/gnomad.genomes.r2.1.1.sites.22.vcf.bgz}
ECHTVAR=${ECHTVAR:-$HOME/projects/immkv/tmp/echtvar/echtvar}
BCFTOOLS=${BCFTOOLS:-$HOME/miniforge3/envs/transvar/bin/bcftools}

mkdir -p "$WORK" && cd "$WORK"
echo "== probe_encode_chr22 $(date -Is) =="

if [ ! -f chr22.decomp.vcf.gz ]; then
  echo "== decomposing chr22 (bcftools norm -m -both) =="
  "$BCFTOOLS" norm -m -both -Oz --threads 4 -o chr22.decomp.vcf.gz "$SITES22" 2> norm.log
fi

rm -f chr22.all.echtvar.zip probe.stdout probe.stderr probe.time
set +e
timeout 3600 /usr/bin/time -v -o probe.time \
  "$ECHTVAR" encode chr22.all.echtvar.zip "$HERE/fields568.json" chr22.decomp.vcf.gz \
  > probe.stdout 2> probe.stderr
rc=$?
set -e

{
  echo "command: echtvar encode chr22.all.echtvar.zip fields568.json chr22.decomp.vcf.gz"
  echo "echtvar version: $("$ECHTVAR" --version)"
  echo "exit code: $rc   (124 = killed by 1h timeout)"
  echo
  echo "-- stderr (verbatim, all unique-value lines are the cardinality census) --"
  cat probe.stderr
  echo
  echo "-- resource usage (/usr/bin/time -v) --"
  grep -E "Elapsed|User time|System time|Maximum resident" probe.time || true
  echo
  echo "-- archive on disk --"
  ls -la chr22.all.echtvar.zip 2>/dev/null || echo "no archive produced"
} | tee "$HERE/pilot_encode_chr22.log"

# hand the pilot archive to the genome-wide build
cp -f chr22.all.echtvar.zip "$WORK/archives/chr22.echtvar.zip"
echo "== probe done rc=$rc $(date -Is) =="
