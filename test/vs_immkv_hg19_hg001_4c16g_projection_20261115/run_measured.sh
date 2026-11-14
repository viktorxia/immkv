#!/bin/bash
# Parent-coordinated continuation: never overlap lookup measurements with
# correctness producers or compilation. Every launcher asserts 4C16G/swap0.
set -euo pipefail
source "$(dirname "$0")/common.sh"
for cell in narrow17 full17 full540 full568; do
    [[ -f "$RESULTS/verify.wgs.$cell.old-after/finished" ]] || fail "Complete annotation verification first: $cell"
done
[[ -f "$RESULTS/code-checks/finished" ]] || fail 'Complete code verification first'
RUN_TAG=wgs bash "$EXP_DIR/run_profile.sh" after 1
bash "$EXP_DIR/analyze_profile.sh" after wgs
RUN_TAG=wgs WORKLOADS='narrow17 full17 full540 full568 full568_default' bash "$EXP_DIR/run_digest.sh" before after
RUN_TAG=wgs WORKLOADS='narrow17 full17 full540 full568 full568_default' bash "$EXP_DIR/run_null.sh" old before after
QUERY="$REPO/tmp/$PREVIOUS/query-chr22.vcf" RUN_TAG=chr22 WORKLOADS='narrow17 full17 full540 full568' \
    bash "$EXP_DIR/run_annotation.sh" old after
