#!/bin/bash
# Controlled v003 rerun: immkv vs echtvar, one sample, 4 physical cores,
# 16G cgroup, gzip output on both sides (compression and the write are inside
# the timed window; decompression and verification are outside it).
#
# Cells:
#   nvme/hdd x {17, 540, 568} x immkv, two passes each
#   nvme/hdd x {17, 540}      x echtvar, two passes each
# 568-field echtvar is not run under the 16G cap: its measured peak is 24.1 GB
# per process (vs_echtvar_hg19_gnomad_genomes_568_fields/README.md), so a 16G
# cgroup cannot hold even one process. verify.sh generates one unconstrained
# echtvar reference for 568 outside the timed matrix.
#
# The script is sequential and resumable: a pass is skipped when its .done
# marker exists. Pass-1 outputs are removed once pass 2 has succeeded; pass-2
# outputs are kept for verification.
set -euo pipefail
source "$(dirname "$0")/common.sh"
cd "$REPO"

MEDIA=${MEDIA:-"nvme hdd"}
FIELDS=${FIELDS:-"17 540 568"}

BIN=$(media_root nvme)/test/4c16g_gz_v003/bin/annotate
mkdir -p "$(dirname "$BIN")" "$EXP_DIR/results"
go build -o "$BIN" ./test/annotate

log() { echo "[$(date -Is)] $*"; }

# Refuse to start if another annotation process is running: a concurrent
# reader would contaminate both the wall time and the page cache.
if pgrep -af 'test/annotate|annotate -in|echtvar anno|bcftools view -r' | grep -v pgrep >/dev/null; then
    echo "refusing to start: another annotation process is running" >&2
    pgrep -af 'test/annotate|annotate -in|echtvar anno|bcftools view -r' >&2 || true
    exit 1
fi

for media in $MEDIA; do
    free=$(df -BG --output=avail "$(media_root "$media")" | tail -1 | tr -dc '0-9')
    log "$media free space: ${free}G"
    if [ "$free" -lt 150 ]; then
        echo "refusing to start: ${media} has less than 150G free" >&2
        exit 1
    fi
done

immkv_cell() { # media field pass
    local media=$1 field=$2 pass=$3
    local work q tsv lib out
    work=$(work_dir "$media" "$field")
    mkdir -p "$work"
    [ -f "$work/immkv.$pass.done" ] && {
        log "skip immkv $media $field $pass"
        return
    }
    q=$(source_dir "$media" "$field")/query/hg001.decomp.vcf.gz
    tsv=$(fields_tsv "$field")
    lib=$(library "$media" "$field")
    out=$work/immkv.$pass.vcf.gz

    log "start immkv $media $field $pass threads=$THREADS batch=$BATCH"
    /usr/bin/time -v -o "$work/immkv.$pass.time" \
        systemd-run --user --scope --collect \
        -p MemoryMax="$MEM_MAX" -p MemorySwapMax=0 \
        taskset -c "$CPUS" env GOMAXPROCS="$THREADS" IMMKV_DEBUG_PATH=1 \
        "$BIN" -in "$q" -lib "$lib" -out "$out" -fields "$tsv" \
        -threads "$THREADS" -batch "$BATCH" \
        2> "$work/immkv.$pass.path.log"
    touch "$work/immkv.$pass.done"
    cp "$work/immkv.$pass.time" "$EXP_DIR/results/$media.$field.immkv.$pass.time"
    grep -E "Elapsed|User time|System time|Percent|Maximum resident|File system inputs|File system outputs" \
        "$work/immkv.$pass.time" || true
}

echtvar_cell() { # media field pass
    local media=$1 field=$2 pass=$3
    local work q arch wk
    work=$(work_dir "$media" "$field")
    mkdir -p "$work"
    [ -f "$work/echtvar.$pass.done" ] && {
        log "skip echtvar $media $field $pass"
        return
    }
    q=$(source_dir "$media" "$field")/query/hg001.decomp.vcf.gz
    arch=$(source_dir "$media" "$field")/archives
    wk=$work/echtvar.$pass.wk
    rm -rf "$wk"
    mkdir -p "$wk"

    export BCFTOOLS ECHTVAR_BIN
    export QUERY=$q ARCHIVES=$arch WK=$wk
    log "start echtvar $media $field $pass procs=$ECHTVAR_PAR"
    /usr/bin/time -v -o "$work/echtvar.$pass.time" \
        systemd-run --user --scope --collect \
        -p MemoryMax="$MEM_MAX" -p MemorySwapMax=0 \
        taskset -c "$CPUS" bash -c '
            export GOMAXPROCS='"$THREADS"'
            one() {
                local c=$1
                "$BCFTOOLS" view -r "$c" -Oz -o "$WK/sl$c.vcf.gz" "$QUERY"
                "$ECHTVAR_BIN" anno -e "$ARCHIVES/chr$c.echtvar.zip" \
                    "$WK/sl$c.vcf.gz" "$WK/an$c.vcf.gz" >"$WK/an$c.log" 2>&1
            }
            export -f one
            seq 1 22 | xargs -P '"$ECHTVAR_PAR"' -I{} bash -c "one {}"
        '
    touch "$work/echtvar.$pass.done"
    cp "$work/echtvar.$pass.time" "$EXP_DIR/results/$media.$field.echtvar.$pass.time"
    grep -E "Elapsed|User time|System time|Percent|Maximum resident|File system inputs|File system outputs" \
        "$work/echtvar.$pass.time" || true
}

for media in $MEDIA; do
    for field in $FIELDS; do
        immkv_cell "$media" "$field" pass1
        immkv_cell "$media" "$field" pass2
        if [ "$field" != 568 ]; then
            echtvar_cell "$media" "$field" pass1
            echtvar_cell "$media" "$field" pass2
        fi
        # Pass-1 artifacts are not needed after pass 2; keep the .time files.
        work=$(work_dir "$media" "$field")
        rm -f "$work/immkv.pass1.vcf.gz"
        [ "$field" != 568 ] && rm -rf "$work/echtvar.pass1.wk"
    done
done

log "all cells done"
