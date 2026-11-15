# Retained release matrix summary

Six media/field cells; 568 means all selected fields. Values below come only from retained records.

| Media | Fields | Engine | Valid rounds | Wall s | User s | System s | CPU % | Max child RSS KiB | Cgroup peak bytes |
|---|---:|---|---:|---:|---:|---:|---:|---:|---:|
| nvme | 17 | immkv | 3/3 | 8.17 | 27.84 | 2.15 | 367 | 419928 | 526442496 |
| nvme | 17 | echtvar | 3/3 | 24.52 | 85.96 | 9.15 | 388 | 31756 | 136314880 |
| nvme | 540 | immkv | 1/3 | — | — | — | — | — | — |
| nvme | 540 | echtvar | 1/3 | — | — | — | — | — | — |
| nvme | 568 | immkv | 0/3 | — | — | — | — | — | — |
| nvme | 568 | echtvar | 0/3 | — | — | — | — | — | — |
| hdd | 17 | immkv | 0/3 | — | — | — | — | — | — |
| hdd | 17 | echtvar | 0/3 | — | — | — | — | — | — |
| hdd | 540 | immkv | 0/3 | — | — | — | — | — | — |
| hdd | 540 | echtvar | 0/3 | — | — | — | — | — | — |
| hdd | 568 | immkv | 0/3 | — | — | — | — | — | — |
| hdd | 568 | echtvar | 0/3 | — | — | — | — | — | — |

## Raw rounds and provenance

### nvme / 17 fields

- immkv r1: ok; valid=True; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/status.json`.
  - wall_seconds=8.37; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt` (wall_seconds).
  - user_seconds=27.49; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt` (user_seconds).
  - system_seconds=2.15; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt` (system_seconds).
  - cpu_percent=354; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt` (cpu_percent).
  - max_child_rss_kib=413424; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt` (max_child_rss_kib).
  - exit_code=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt` (Exit status).
  - cgroup_peak_bytes=520769536; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/status.json` (memory_peak_bytes).
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/constraints.json`; confirmed=True.
  - rows=3933714; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log`.
  - hits=3875208; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log`.
  - Internal read_gib_rounded=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log`.
  - Internal key_seconds=0.29; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log`.
  - Internal query_seconds=4.46; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log`.
  - Internal format_seconds=6.99; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log`.
  - Internal wall_seconds=8.24; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log`.
- immkv r2: ok; valid=True; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/status.json`.
  - wall_seconds=8.17; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt` (wall_seconds).
  - user_seconds=27.96; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt` (user_seconds).
  - system_seconds=2.1; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt` (system_seconds).
  - cpu_percent=367; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt` (cpu_percent).
  - max_child_rss_kib=419928; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt` (max_child_rss_kib).
  - exit_code=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt` (Exit status).
  - cgroup_peak_bytes=526442496; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/status.json` (memory_peak_bytes).
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/constraints.json`; confirmed=True.
  - rows=3933714; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log`.
  - hits=3875208; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log`.
  - Internal read_gib_rounded=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log`.
  - Internal key_seconds=0.3; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log`.
  - Internal query_seconds=4.12; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log`.
  - Internal format_seconds=7.12; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log`.
  - Internal wall_seconds=8.049; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log`.
- immkv r3: ok; valid=True; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/status.json`.
  - wall_seconds=8.16; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt` (wall_seconds).
  - user_seconds=27.84; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt` (user_seconds).
  - system_seconds=2.18; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt` (system_seconds).
  - cpu_percent=367; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt` (cpu_percent).
  - max_child_rss_kib=423004; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt` (max_child_rss_kib).
  - exit_code=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt` (Exit status).
  - cgroup_peak_bytes=529625088; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/status.json` (memory_peak_bytes).
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/constraints.json`; confirmed=True.
  - rows=3933714; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log`.
  - hits=3875208; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log`.
  - Internal read_gib_rounded=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log`.
  - Internal key_seconds=0.3; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log`.
  - Internal query_seconds=4.04; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log`.
  - Internal format_seconds=7.06; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log`.
  - Internal wall_seconds=7.989; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log`.
- immkv wall_seconds: median=8.17, range=8.16–8.37; sources `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt`.
- immkv user_seconds: median=27.84, range=27.49–27.96; sources `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt`.
- immkv system_seconds: median=2.15, range=2.1–2.18; sources `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt`.
- immkv cpu_percent: median=367, range=354–367; sources `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt`.
- immkv max_child_rss_kib: median=419928, range=413424–423004; sources `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt`.
- immkv cgroup_peak_bytes: median=526442496, range=520769536–529625088; sources `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/status.json`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/status.json`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/status.json`.
- echtvar r1: ok; valid=True; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/status.json`.
  - wall_seconds=24.42; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt` (wall_seconds).
  - user_seconds=85.66; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt` (user_seconds).
  - system_seconds=9.15; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt` (system_seconds).
  - cpu_percent=388; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt` (cpu_percent).
  - max_child_rss_kib=31756; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt` (max_child_rss_kib).
  - exit_code=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt` (Exit status).
  - cgroup_peak_bytes=136314880; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/status.json` (memory_peak_bytes).
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/constraints.json`; confirmed=True.
  - All chromosome statuses successful=True; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/cell.json`.
  - chr1 evaluated=308298, wrote=308297; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr1.anno.log`.
  - chr2 evaluated=329821, wrote=329821; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr2.anno.log`.
  - chr3 evaluated=275651, wrote=275647; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr3.anno.log`.
  - chr4 evaluated=296434, wrote=296434; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr4.anno.log`.
  - chr5 evaluated=253843, wrote=253841; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr5.anno.log`.
  - chr6 evaluated=264414, wrote=264084; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr6.anno.log`.
  - chr7 evaluated=222733, wrote=222732; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr7.anno.log`.
  - chr8 evaluated=213367, wrote=213367; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr8.anno.log`.
  - chr9 evaluated=168542, wrote=168542; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr9.anno.log`.
  - chr10 evaluated=198281, wrote=198281; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr10.anno.log`.
  - chr11 evaluated=200574, wrote=200574; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr11.anno.log`.
  - chr12 evaluated=189356, wrote=189356; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr12.anno.log`.
  - chr13 evaluated=156939, wrote=156938; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr13.anno.log`.
  - chr14 evaluated=125631, wrote=125631; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr14.anno.log`.
  - chr15 evaluated=114117, wrote=114117; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr15.anno.log`.
  - chr16 evaluated=118437, wrote=118437; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr16.anno.log`.
  - chr17 evaluated=102461, wrote=102461; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr17.anno.log`.
  - chr18 evaluated=119403, wrote=119401; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr18.anno.log`.
  - chr19 evaluated=86442, wrote=86442; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr19.anno.log`.
  - chr20 evaluated=83452, wrote=83452; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr20.anno.log`.
  - chr21 evaluated=55978, wrote=55978; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr21.anno.log`.
  - chr22 evaluated=49540, wrote=49536; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr22.anno.log`.
- echtvar r2: ok; valid=True; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/status.json`.
  - wall_seconds=24.52; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt` (wall_seconds).
  - user_seconds=85.96; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt` (user_seconds).
  - system_seconds=9.26; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt` (system_seconds).
  - cpu_percent=388; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt` (cpu_percent).
  - max_child_rss_kib=31736; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt` (max_child_rss_kib).
  - exit_code=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt` (Exit status).
  - cgroup_peak_bytes=136704000; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/status.json` (memory_peak_bytes).
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/constraints.json`; confirmed=True.
  - All chromosome statuses successful=True; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/cell.json`.
  - chr1 evaluated=308298, wrote=308297; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr1.anno.log`.
  - chr2 evaluated=329821, wrote=329821; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr2.anno.log`.
  - chr3 evaluated=275651, wrote=275647; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr3.anno.log`.
  - chr4 evaluated=296434, wrote=296434; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr4.anno.log`.
  - chr5 evaluated=253843, wrote=253841; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr5.anno.log`.
  - chr6 evaluated=264414, wrote=264084; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr6.anno.log`.
  - chr7 evaluated=222733, wrote=222732; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr7.anno.log`.
  - chr8 evaluated=213367, wrote=213367; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr8.anno.log`.
  - chr9 evaluated=168542, wrote=168542; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr9.anno.log`.
  - chr10 evaluated=198281, wrote=198281; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr10.anno.log`.
  - chr11 evaluated=200574, wrote=200574; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr11.anno.log`.
  - chr12 evaluated=189356, wrote=189356; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr12.anno.log`.
  - chr13 evaluated=156939, wrote=156938; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr13.anno.log`.
  - chr14 evaluated=125631, wrote=125631; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr14.anno.log`.
  - chr15 evaluated=114117, wrote=114117; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr15.anno.log`.
  - chr16 evaluated=118437, wrote=118437; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr16.anno.log`.
  - chr17 evaluated=102461, wrote=102461; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr17.anno.log`.
  - chr18 evaluated=119403, wrote=119401; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr18.anno.log`.
  - chr19 evaluated=86442, wrote=86442; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr19.anno.log`.
  - chr20 evaluated=83452, wrote=83452; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr20.anno.log`.
  - chr21 evaluated=55978, wrote=55978; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr21.anno.log`.
  - chr22 evaluated=49540, wrote=49536; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr22.anno.log`.
- echtvar r3: ok; valid=True; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/status.json`.
  - wall_seconds=24.52; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt` (wall_seconds).
  - user_seconds=86.12; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt` (user_seconds).
  - system_seconds=9.06; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt` (system_seconds).
  - cpu_percent=388; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt` (cpu_percent).
  - max_child_rss_kib=31760; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt` (max_child_rss_kib).
  - exit_code=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt` (Exit status).
  - cgroup_peak_bytes=136237056; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/status.json` (memory_peak_bytes).
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/constraints.json`; confirmed=True.
  - All chromosome statuses successful=True; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/cell.json`.
  - chr1 evaluated=308298, wrote=308297; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr1.anno.log`.
  - chr2 evaluated=329821, wrote=329821; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr2.anno.log`.
  - chr3 evaluated=275651, wrote=275647; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr3.anno.log`.
  - chr4 evaluated=296434, wrote=296434; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr4.anno.log`.
  - chr5 evaluated=253843, wrote=253841; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr5.anno.log`.
  - chr6 evaluated=264414, wrote=264084; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr6.anno.log`.
  - chr7 evaluated=222733, wrote=222732; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr7.anno.log`.
  - chr8 evaluated=213367, wrote=213367; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr8.anno.log`.
  - chr9 evaluated=168542, wrote=168542; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr9.anno.log`.
  - chr10 evaluated=198281, wrote=198281; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr10.anno.log`.
  - chr11 evaluated=200574, wrote=200574; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr11.anno.log`.
  - chr12 evaluated=189356, wrote=189356; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr12.anno.log`.
  - chr13 evaluated=156939, wrote=156938; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr13.anno.log`.
  - chr14 evaluated=125631, wrote=125631; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr14.anno.log`.
  - chr15 evaluated=114117, wrote=114117; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr15.anno.log`.
  - chr16 evaluated=118437, wrote=118437; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr16.anno.log`.
  - chr17 evaluated=102461, wrote=102461; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr17.anno.log`.
  - chr18 evaluated=119403, wrote=119401; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr18.anno.log`.
  - chr19 evaluated=86442, wrote=86442; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr19.anno.log`.
  - chr20 evaluated=83452, wrote=83452; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr20.anno.log`.
  - chr21 evaluated=55978, wrote=55978; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr21.anno.log`.
  - chr22 evaluated=49540, wrote=49536; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr22.anno.log`.
- echtvar wall_seconds: median=24.52, range=24.42–24.52; sources `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt`.
- echtvar user_seconds: median=85.96, range=85.66–86.12; sources `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt`.
- echtvar system_seconds: median=9.15, range=9.06–9.26; sources `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt`.
- echtvar cpu_percent: median=388, range=388–388; sources `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt`.
- echtvar max_child_rss_kib: median=31756, range=31736–31760; sources `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt`.
- echtvar cgroup_peak_bytes: median=136314880, range=136237056–136704000; sources `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/status.json`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/status.json`, `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/status.json`.
- Paired echtvar/immkv wall ratios: [2.91756, 3.00122, 3.0049], median=3.00122, range=2.91756–3.0049, range/median=0.0291012; r1: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt` / `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt`; r2: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt` / `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt`; r3: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt` / `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt`.
- Correctness immkv: scope=ok, passed=True; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verification.json`; scope source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/status.json`.

Evidence from `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verification.json`:

```json
{
  "annotate_statistics": {
    "hits": 3875208,
    "rows": 3933714
  },
  "commands": {
    "annotate": [
      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/annotate",
      "-in",
      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz",
      "-lib",
      "/home/tian/immkv_test_nvme/test/17_fields/lib17",
      "-fields",
      "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_17_fields/fields17.tsv",
      "-threads",
      "4",
      "-batch",
      "4096",
      "-out",
      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-17-release/output.vcf.gz"
    ],
    "verifier": [
      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/verifyvalues",
      "-ev",
      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-17-release/echtvar.fifo",
      "-iv",
      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-17-release/immkv.fifo",
      "-fields",
      "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_17_fields/fields17.tsv",
      "-tol",
      "1.1e-6",
      "-max-report",
      "20"
    ]
  },
  "comparison": {
    "echtvar_i32_overflow_losses": 0,
    "field_comparisons": 66873138,
    "hit_disagreements": 0,
    "misses_agreed": 1007196,
    "rows": 3933714,
    "value_mismatches": 0
  },
  "elapsed_seconds_correctness_only": 40.02600424699631,
  "exit_status": {
    "annotate": 0,
    "verifier": 0
  },
  "expected_hits": 3875208,
  "expected_rows": 3933714,
  "fields": 17,
  "fields_tsv": "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_17_fields/fields17.tsv",
  "float_tolerance": 1.1e-06,
  "float_tolerance_basis": "echtvar v0.2.4 quantization truncation plus six-decimal print truncation and f32 slack",
  "library": "/home/tian/immkv_test_nvme/test/17_fields/lib17",
  "marker_collision_accounting": "Shared verifier includes true float -1 marker collisions in misses agreed.",
  "media": "nvme",
  "query": "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz",
  "reference_kind": "historical correctness reference; not a current performance sample",
  "reference_media": "nvme",
  "reference_release": "echtvar v0.2.4",
  "sha256": {
    "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_17_fields/fields17.tsv": "db2c8bc29926b5f4265e9077443a19571a55c8f99d2d30906e0a126c6c0db030",
    "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify_output.py": "2466d4c8fb67067d5180a83f44344057ebd420726a174019956225a1f98cef78",
    "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/annotate": "4179f6f9bf774615ce68e1794dd3c46f2c5e4b2316011b117191584f5ec16391",
    "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/verifyvalues": "8e307e09177310a07ec9615734493f9e423a779f3b6f3a2a5dd82f4364343400"
  },
  "status": "passed",
  "streams": {
    "echtvar": {
      "data_sha256": "ec008db079adf59cb5f84928a5096821694fddef1ded9757fff1043725fe6720",
      "files": [
        {
          "data_bytes": 269824780,
          "data_sha256": "1f44447b194ead91f936221f2bb3222981dd0926c80a98045774a787721fb477",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an1.vcf.gz",
          "rows": 308298
        },
        {
          "data_bytes": 289150520,
          "data_sha256": "d9209bf52f85e8d53e6456f6090005d145b8c723c8eb438525ce62e5d699bb25",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an2.vcf.gz",
          "rows": 329821
        },
        {
          "data_bytes": 241537645,
          "data_sha256": "a8d37d11be87aa98cbdc60f41451c38cd3f180fdc713f36a3a0403d4e0d03b06",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an3.vcf.gz",
          "rows": 275651
        },
        {
          "data_bytes": 259661090,
          "data_sha256": "c70e6ff200808692bb15357b95933f255ef22c8784c172135bf869aeada9b192",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an4.vcf.gz",
          "rows": 296434
        },
        {
          "data_bytes": 222394381,
          "data_sha256": "f423059e24425dbb679170527a887a13a1677ccec48ce1e2d6aa4947b5961a3f",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an5.vcf.gz",
          "rows": 253843
        },
        {
          "data_bytes": 217988880,
          "data_sha256": "dd63acbe440fe990831807637a7741570970f50e1ccdead0cf2b18ebf92ecd67",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an6.vcf.gz",
          "rows": 264414
        },
        {
          "data_bytes": 194980426,
          "data_sha256": "ff2efa1a6c2b20a49037768efaf307afeebcbb3536f2a9f5670a7e744fc067ef",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an7.vcf.gz",
          "rows": 222733
        },
        {
          "data_bytes": 187116360,
          "data_sha256": "1adf70d93eeab1d8ee988b0556b21d89313f85123cacd243cb5981cf880c7b0d",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an8.vcf.gz",
          "rows": 213367
        },
        {
          "data_bytes": 147736697,
          "data_sha256": "325be35c2cc1be533573151313f57e5331ece22d9ebd65cb8ae769ae11cf9a06",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an9.vcf.gz",
          "rows": 168542
        },
        {
          "data_bytes": 173832355,
          "data_sha256": "63e7d0f919c529977ebeed9ea8826093402f53a3e23ff4df3e9b3486a55b0213",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an10.vcf.gz",
          "rows": 198281
        },
        {
          "data_bytes": 175692515,
          "data_sha256": "58763a4b3db30bf77b4c2fcd43de752d36868ef90cd9ec7703c148d0b86ebdc9",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an11.vcf.gz",
          "rows": 200574
        },
        {
          "data_bytes": 165790011,
          "data_sha256": "dc8d9d2acfb5bf8cd68a7e1c4db6dad160b920f88cc8a774d6ddcac27ba266a2",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an12.vcf.gz",
          "rows": 189356
        },
        {
          "data_bytes": 137625829,
          "data_sha256": "25e1505b346e6f86f120186736e83d344c4fa729ac1372068e791264ea2a3b65",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an13.vcf.gz",
          "rows": 156939
        },
        {
          "data_bytes": 110046147,
          "data_sha256": "952a7a7b4ea03254e567232feaedf7108156810b47894e1071eb5093051a08ed",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an14.vcf.gz",
          "rows": 125631
        },
        {
          "data_bytes": 100094646,
          "data_sha256": "0e7f24f259fa29545b6a2e3d0c266c18ef7d04675e17c88cdf85ee4b15a70ead",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an15.vcf.gz",
          "rows": 114117
        },
        {
          "data_bytes": 103879536,
          "data_sha256": "68c1a3c0ffd48e36e7583ba722f4af1f2d400ded3454af3c1cdfba59ace359b0",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an16.vcf.gz",
          "rows": 118437
        },
        {
          "data_bytes": 89523814,
          "data_sha256": "851a7e861003ed1ba8e021ad28f7a87dfc99f0b38159b7d0cfecdaaf0725c020",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an17.vcf.gz",
          "rows": 102461
        },
        {
          "data_bytes": 104836551,
          "data_sha256": "07f618d6065c8f4998b9e03fbd3fdb1436ae81ac475a6205a33fe214cd286a6e",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an18.vcf.gz",
          "rows": 119403
        },
        {
          "data_bytes": 75169025,
          "data_sha256": "f088ed123c5b73dd87beea4691ae6559a17a91cfc8ef6deb94ebe810db8c1710",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an19.vcf.gz",
          "rows": 86442
        },
        {
          "data_bytes": 73159040,
          "data_sha256": "45f3e17c5f4670ab27b47b3b496fe2a37f79068afdbf5366b3c1db91ce54a640",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an20.vcf.gz",
          "rows": 83452
        },
        {
          "data_bytes": 49088311,
          "data_sha256": "1d4a87c093eb5bf1dc897ad5e6c5f94538f316b868d37462fb6514bb5fee9de2",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an21.vcf.gz",
          "rows": 55978
        },
        {
          "data_bytes": 43278460,
          "data_sha256": "c4e02794b0aa94ca0a0eae05006399e26e108da4d8d7ce5aae62e855faf31ea7",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an22.vcf.gz",
          "rows": 49540
        }
      ],
      "rows": 3933714
    },
    "immkv": {
      "data_sha256": "8e1c55459edc13e1b24236ef9d65b54f3215dc11dd55db97efa5979f5247af63",
      "files": [
        {
          "data_bytes": 3418342825,
          "data_sha256": "8e1c55459edc13e1b24236ef9d65b54f3215dc11dd55db97efa5979f5247af63",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-17-release/output.vcf.gz",
          "rows": 3933714
        }
      ],
      "rows": 3933714,
      "selected_float_true_minus_one_candidates": 0
    }
  },
  "timing_sample": false
}
```

- immkv rows=3933714; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout`.
- immkv field_comparisons=66873138; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout`.
- immkv misses_agreed=1007196; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout`.
- immkv value_mismatches=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout`.
- immkv hit_disagreements=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout`.
- immkv echtvar_i32_overflow_losses=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout`.
- immkv float_comparisons=30989070; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout`.
- immkv max_float_absolute_difference=1e-06; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout`.
- immkv float_tolerance=1.1e-06; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout`.
- immkv verify_values_ok=True; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout`.
- Correctness fresh_echtvar: scope=ok, passed=True; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-17-release/cell.json`; scope source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-17-release/status.json`.

Evidence from `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-17-release/cell.json`:

```json
{
  "chromosomes": [
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr1.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl1.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an1.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58857612954211,
        "time_ns": 1794684680803278795
      },
      "annotation_started": {
        "monotonic_ns": 58849454555341,
        "time_ns": 1794684672644879785
      },
      "chromosome": 1,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "1f44447b194ead91f936221f2bb3222981dd0926c80a98045774a787721fb477",
          "rows": 308298
        },
        "reference": {
          "data_sha256": "1f44447b194ead91f936221f2bb3222981dd0926c80a98045774a787721fb477",
          "rows": 308298
        }
      },
      "finished": {
        "monotonic_ns": 58857622844052,
        "time_ns": 1794684680813168727
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "1",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl1.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 10999772,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58849453672913,
        "time_ns": 1794684672643996826
      },
      "slice_started": {
        "monotonic_ns": 58848351059493,
        "time_ns": 1794684671541384047
      },
      "started": {
        "monotonic_ns": 58848350632034,
        "time_ns": 1794684671540956308
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr2.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl2.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an2.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58858117841837,
        "time_ns": 1794684681308165610
      },
      "annotation_started": {
        "monotonic_ns": 58849656654371,
        "time_ns": 1794684672846979056
      },
      "chromosome": 2,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "d9209bf52f85e8d53e6456f6090005d145b8c723c8eb438525ce62e5d699bb25",
          "rows": 329821
        },
        "reference": {
          "data_sha256": "d9209bf52f85e8d53e6456f6090005d145b8c723c8eb438525ce62e5d699bb25",
          "rows": 329821
        }
      },
      "finished": {
        "monotonic_ns": 58858128074900,
        "time_ns": 1794684681318398903
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "2",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl2.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 11750291,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58849655718373,
        "time_ns": 1794684672846042366
      },
      "slice_started": {
        "monotonic_ns": 58848351318366,
        "time_ns": 1794684671541642921
      },
      "started": {
        "monotonic_ns": 58848350909553,
        "time_ns": 1794684671541234047
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr3.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl3.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an3.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58856594949838,
        "time_ns": 1794684679785273882
      },
      "annotation_started": {
        "monotonic_ns": 58849455373830,
        "time_ns": 1794684672645698564
      },
      "chromosome": 3,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "a8d37d11be87aa98cbdc60f41451c38cd3f180fdc713f36a3a0403d4e0d03b06",
          "rows": 275651
        },
        "reference": {
          "data_sha256": "a8d37d11be87aa98cbdc60f41451c38cd3f180fdc713f36a3a0403d4e0d03b06",
          "rows": 275651
        }
      },
      "finished": {
        "monotonic_ns": 58856605112139,
        "time_ns": 1794684679795436182
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "3",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl3.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 9815087,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58849454905174,
        "time_ns": 1794684672645229769
      },
      "slice_started": {
        "monotonic_ns": 58848352222455,
        "time_ns": 1794684671542547169
      },
      "started": {
        "monotonic_ns": 58848351397745,
        "time_ns": 1794684671541722199
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr4.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl4.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an4.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58856807554780,
        "time_ns": 1794684679997878924
      },
      "annotation_started": {
        "monotonic_ns": 58849455824722,
        "time_ns": 1794684672646149276
      },
      "chromosome": 4,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "c70e6ff200808692bb15357b95933f255ef22c8784c172135bf869aeada9b192",
          "rows": 296434
        },
        "reference": {
          "data_sha256": "c70e6ff200808692bb15357b95933f255ef22c8784c172135bf869aeada9b192",
          "rows": 296434
        }
      },
      "finished": {
        "monotonic_ns": 58856811310585,
        "time_ns": 1794684680001634788
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "4",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl4.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 10519953,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58849455007115,
        "time_ns": 1794684672645331399
      },
      "slice_started": {
        "monotonic_ns": 58848352079818,
        "time_ns": 1794684671542404583
      },
      "started": {
        "monotonic_ns": 58848351722241,
        "time_ns": 1794684671542046815
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr5.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl5.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an5.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58863646196183,
        "time_ns": 1794684686836520157
      },
      "annotation_started": {
        "monotonic_ns": 58857612927200,
        "time_ns": 1794684680803251745
      },
      "chromosome": 5,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "f423059e24425dbb679170527a887a13a1677ccec48ce1e2d6aa4947b5961a3f",
          "rows": 253843
        },
        "reference": {
          "data_sha256": "f423059e24425dbb679170527a887a13a1677ccec48ce1e2d6aa4947b5961a3f",
          "rows": 253843
        }
      },
      "finished": {
        "monotonic_ns": 58863650387251,
        "time_ns": 1794684686840711715
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "5",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl5.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 9032175,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58857612383806,
        "time_ns": 1794684680802707679
      },
      "slice_started": {
        "monotonic_ns": 58856605442506,
        "time_ns": 1794684679795767270
      },
      "started": {
        "monotonic_ns": 58856605330486,
        "time_ns": 1794684679795655201
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr6.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl6.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an6.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58863855780172,
        "time_ns": 1794684687046104035
      },
      "annotation_started": {
        "monotonic_ns": 58857819106140,
        "time_ns": 1794684681009430394
      },
      "chromosome": 6,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "dd63acbe440fe990831807637a7741570970f50e1ccdead0cf2b18ebf92ecd67",
          "rows": 264414
        },
        "reference": {
          "data_sha256": "dd63acbe440fe990831807637a7741570970f50e1ccdead0cf2b18ebf92ecd67",
          "rows": 264414
        }
      },
      "finished": {
        "monotonic_ns": 58863860302508,
        "time_ns": 1794684687050626992
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "6",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl6.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 8584380,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58857817057904,
        "time_ns": 1794684681007382068
      },
      "slice_started": {
        "monotonic_ns": 58856812873403,
        "time_ns": 1794684680003198158
      },
      "started": {
        "monotonic_ns": 58856811506130,
        "time_ns": 1794684680001830844
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr7.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl7.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an7.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58864273911914,
        "time_ns": 1794684687464235877
      },
      "annotation_started": {
        "monotonic_ns": 58858541707379,
        "time_ns": 1794684681732031944
      },
      "chromosome": 7,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "ff2efa1a6c2b20a49037768efaf307afeebcbb3536f2a9f5670a7e744fc067ef",
          "rows": 222733
        },
        "reference": {
          "data_sha256": "ff2efa1a6c2b20a49037768efaf307afeebcbb3536f2a9f5670a7e744fc067ef",
          "rows": 222733
        }
      },
      "finished": {
        "monotonic_ns": 58864280160644,
        "time_ns": 1794684687470485079
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "7",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl7.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 7924912,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58858541415193,
        "time_ns": 1794684681731739187
      },
      "slice_started": {
        "monotonic_ns": 58857624769819,
        "time_ns": 1794684680815094534
      },
      "started": {
        "monotonic_ns": 58857624560428,
        "time_ns": 1794684680814885163
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr8.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl8.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an8.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58864500111625,
        "time_ns": 1794684687690435658
      },
      "annotation_started": {
        "monotonic_ns": 58859048603306,
        "time_ns": 1794684682238927480
      },
      "chromosome": 8,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "1adf70d93eeab1d8ee988b0556b21d89313f85123cacd243cb5981cf880c7b0d",
          "rows": 213367
        },
        "reference": {
          "data_sha256": "1adf70d93eeab1d8ee988b0556b21d89313f85123cacd243cb5981cf880c7b0d",
          "rows": 213367
        }
      },
      "finished": {
        "monotonic_ns": 58864503390609,
        "time_ns": 1794684687693715304
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "8",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl8.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 7556435,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58859048280002,
        "time_ns": 1794684682238603936
      },
      "slice_started": {
        "monotonic_ns": 58858128651146,
        "time_ns": 1794684681318975881
      },
      "started": {
        "monotonic_ns": 58858128519129,
        "time_ns": 1794684681318843854
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr9.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl9.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an9.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58868281699497,
        "time_ns": 1794684691472023460
      },
      "annotation_started": {
        "monotonic_ns": 58864456238814,
        "time_ns": 1794684687646563298
      },
      "chromosome": 9,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "325be35c2cc1be533573151313f57e5331ece22d9ebd65cb8ae769ae11cf9a06",
          "rows": 168542
        },
        "reference": {
          "data_sha256": "325be35c2cc1be533573151313f57e5331ece22d9ebd65cb8ae769ae11cf9a06",
          "rows": 168542
        }
      },
      "finished": {
        "monotonic_ns": 58868287700765,
        "time_ns": 1794684691478024919
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "9",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl9.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 5978980,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58864455774567,
        "time_ns": 1794684687646098510
      },
      "slice_started": {
        "monotonic_ns": 58863651813665,
        "time_ns": 1794684686842138360
      },
      "started": {
        "monotonic_ns": 58863650563059,
        "time_ns": 1794684686840887784
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr10.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl10.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an10.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58869297187635,
        "time_ns": 1794684692487511939
      },
      "annotation_started": {
        "monotonic_ns": 58864767747697,
        "time_ns": 1794684687958072392
      },
      "chromosome": 10,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "63e7d0f919c529977ebeed9ea8826093402f53a3e23ff4df3e9b3486a55b0213",
          "rows": 198281
        },
        "reference": {
          "data_sha256": "63e7d0f919c529977ebeed9ea8826093402f53a3e23ff4df3e9b3486a55b0213",
          "rows": 198281
        }
      },
      "finished": {
        "monotonic_ns": 58869301659677,
        "time_ns": 1794684692491984061
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "10",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl10.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 7017518,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58864767187872,
        "time_ns": 1794684687957511715
      },
      "slice_started": {
        "monotonic_ns": 58863860584315,
        "time_ns": 1794684687050909069
      },
      "started": {
        "monotonic_ns": 58863860476643,
        "time_ns": 1794684687050801368
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr11.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl11.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an11.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58870323351221,
        "time_ns": 1794684693513675244
      },
      "annotation_started": {
        "monotonic_ns": 58865190669595,
        "time_ns": 1794684688380993969
      },
      "chromosome": 11,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "58763a4b3db30bf77b4c2fcd43de752d36868ef90cd9ec7703c148d0b86ebdc9",
          "rows": 200574
        },
        "reference": {
          "data_sha256": "58763a4b3db30bf77b4c2fcd43de752d36868ef90cd9ec7703c148d0b86ebdc9",
          "rows": 200574
        }
      },
      "finished": {
        "monotonic_ns": 58870326745150,
        "time_ns": 1794684693517069604
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "11",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl11.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 7109132,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58865184499271,
        "time_ns": 1794684688374823104
      },
      "slice_started": {
        "monotonic_ns": 58864280440567,
        "time_ns": 1794684687470765342
      },
      "started": {
        "monotonic_ns": 58864280332395,
        "time_ns": 1794684687470657140
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr12.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl12.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an12.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58870059090133,
        "time_ns": 1794684693249414417
      },
      "annotation_started": {
        "monotonic_ns": 58865217512049,
        "time_ns": 1794684688407836583
      },
      "chromosome": 12,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "dc8d9d2acfb5bf8cd68a7e1c4db6dad160b920f88cc8a774d6ddcac27ba266a2",
          "rows": 189356
        },
        "reference": {
          "data_sha256": "dc8d9d2acfb5bf8cd68a7e1c4db6dad160b920f88cc8a774d6ddcac27ba266a2",
          "rows": 189356
        }
      },
      "finished": {
        "monotonic_ns": 58870064540182,
        "time_ns": 1794684693254864326
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "12",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl12.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 6749307,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58865214950735,
        "time_ns": 1794684688405274879
      },
      "slice_started": {
        "monotonic_ns": 58864503672446,
        "time_ns": 1794684687693997170
      },
      "started": {
        "monotonic_ns": 58864503567139,
        "time_ns": 1794684687693891744
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr13.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl13.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an13.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58872739642855,
        "time_ns": 1794684695929966968
      },
      "annotation_started": {
        "monotonic_ns": 58869098099610,
        "time_ns": 1794684692288424144
      },
      "chromosome": 13,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "25e1505b346e6f86f120186736e83d344c4fa729ac1372068e791264ea2a3b65",
          "rows": 156939
        },
        "reference": {
          "data_sha256": "25e1505b346e6f86f120186736e83d344c4fa729ac1372068e791264ea2a3b65",
          "rows": 156939
        }
      },
      "finished": {
        "monotonic_ns": 58872742972373,
        "time_ns": 1794684695933296807
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "13",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl13.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 5531589,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58869090410048,
        "time_ns": 1794684692280733851
      },
      "slice_started": {
        "monotonic_ns": 58868288026143,
        "time_ns": 1794684691478350888
      },
      "started": {
        "monotonic_ns": 58868287892433,
        "time_ns": 1794684691478217028
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr14.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl14.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an14.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58872933498447,
        "time_ns": 1794684696123822630
      },
      "annotation_started": {
        "monotonic_ns": 58869904927284,
        "time_ns": 1794684693095251989
      },
      "chromosome": 14,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "952a7a7b4ea03254e567232feaedf7108156810b47894e1071eb5093051a08ed",
          "rows": 125631
        },
        "reference": {
          "data_sha256": "952a7a7b4ea03254e567232feaedf7108156810b47894e1071eb5093051a08ed",
          "rows": 125631
        }
      },
      "finished": {
        "monotonic_ns": 58872943508462,
        "time_ns": 1794684696133832946
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "14",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl14.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 4486824,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58869904693057,
        "time_ns": 1794684693095017251
      },
      "slice_started": {
        "monotonic_ns": 58869302055907,
        "time_ns": 1794684692492380622
      },
      "started": {
        "monotonic_ns": 58869301847858,
        "time_ns": 1794684692492172603
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr15.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl15.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an15.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58873596614131,
        "time_ns": 1794684696786938154
      },
      "annotation_started": {
        "monotonic_ns": 58870673549422,
        "time_ns": 1794684693863873917
      },
      "chromosome": 15,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "0e7f24f259fa29545b6a2e3d0c266c18ef7d04675e17c88cdf85ee4b15a70ead",
          "rows": 114117
        },
        "reference": {
          "data_sha256": "0e7f24f259fa29545b6a2e3d0c266c18ef7d04675e17c88cdf85ee4b15a70ead",
          "rows": 114117
        }
      },
      "finished": {
        "monotonic_ns": 58873602585683,
        "time_ns": 1794684696792910057
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "15",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl15.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 4074508,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58870673273026,
        "time_ns": 1794684693863597000
      },
      "slice_started": {
        "monotonic_ns": 58870064815296,
        "time_ns": 1794684693255140031
      },
      "started": {
        "monotonic_ns": 58870064708837,
        "time_ns": 1794684693255033502
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr16.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl16.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an16.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58873957666875,
        "time_ns": 1794684697147990868
      },
      "annotation_started": {
        "monotonic_ns": 58870939338895,
        "time_ns": 1794684694129663159
      },
      "chromosome": 16,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "68c1a3c0ffd48e36e7583ba722f4af1f2d400ded3454af3c1cdfba59ace359b0",
          "rows": 118437
        },
        "reference": {
          "data_sha256": "68c1a3c0ffd48e36e7583ba722f4af1f2d400ded3454af3c1cdfba59ace359b0",
          "rows": 118437
        }
      },
      "finished": {
        "monotonic_ns": 58873960372448,
        "time_ns": 1794684697150696641
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "16",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl16.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 4212957,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58870937399563,
        "time_ns": 1794684694127723476
      },
      "slice_started": {
        "monotonic_ns": 58870327021085,
        "time_ns": 1794684693517345820
      },
      "started": {
        "monotonic_ns": 58870326909476,
        "time_ns": 1794684693517234211
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr17.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl17.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an17.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58875768363532,
        "time_ns": 1794684698958687165
      },
      "annotation_started": {
        "monotonic_ns": 58873251704897,
        "time_ns": 1794684696442029422
      },
      "chromosome": 17,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "851a7e861003ed1ba8e021ad28f7a87dfc99f0b38159b7d0cfecdaaf0725c020",
          "rows": 102461
        },
        "reference": {
          "data_sha256": "851a7e861003ed1ba8e021ad28f7a87dfc99f0b38159b7d0cfecdaaf0725c020",
          "rows": 102461
        }
      },
      "finished": {
        "monotonic_ns": 58875771319833,
        "time_ns": 1794684698961644188
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "17",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl17.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 3711658,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58873251457816,
        "time_ns": 1794684696441781659
      },
      "slice_started": {
        "monotonic_ns": 58872747802594,
        "time_ns": 1794684695938127139
      },
      "started": {
        "monotonic_ns": 58872743177085,
        "time_ns": 1794684695933501720
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr18.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl18.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an18.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58876472392459,
        "time_ns": 1794684699662716632
      },
      "annotation_started": {
        "monotonic_ns": 58873554142367,
        "time_ns": 1794684696744466841
      },
      "chromosome": 18,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "07f618d6065c8f4998b9e03fbd3fdb1436ae81ac475a6205a33fe214cd286a6e",
          "rows": 119403
        },
        "reference": {
          "data_sha256": "07f618d6065c8f4998b9e03fbd3fdb1436ae81ac475a6205a33fe214cd286a6e",
          "rows": 119403
        }
      },
      "finished": {
        "monotonic_ns": 58876475817676,
        "time_ns": 1794684699666142210
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "18",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl18.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 4239925,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58873550024826,
        "time_ns": 1794684696740348779
      },
      "slice_started": {
        "monotonic_ns": 58872943795298,
        "time_ns": 1794684696134120073
      },
      "started": {
        "monotonic_ns": 58872943689170,
        "time_ns": 1794684696134013894
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr19.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl19.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an19.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58876222171564,
        "time_ns": 1794684699412495728
      },
      "annotation_started": {
        "monotonic_ns": 58874007738343,
        "time_ns": 1794684697198062817
      },
      "chromosome": 19,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "f088ed123c5b73dd87beea4691ae6559a17a91cfc8ef6deb94ebe810db8c1710",
          "rows": 86442
        },
        "reference": {
          "data_sha256": "f088ed123c5b73dd87beea4691ae6559a17a91cfc8ef6deb94ebe810db8c1710",
          "rows": 86442
        }
      },
      "finished": {
        "monotonic_ns": 58876226436169,
        "time_ns": 1794684699416760513
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "19",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl19.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 3140431,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58874005803439,
        "time_ns": 1794684697196127372
      },
      "slice_started": {
        "monotonic_ns": 58873602867329,
        "time_ns": 1794684696793192094
      },
      "started": {
        "monotonic_ns": 58873602764277,
        "time_ns": 1794684696793089002
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr20.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl20.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an20.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58876482708716,
        "time_ns": 1794684699673032980
      },
      "annotation_started": {
        "monotonic_ns": 58874364530029,
        "time_ns": 1794684697554854503
      },
      "chromosome": 20,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "45f3e17c5f4670ab27b47b3b496fe2a37f79068afdbf5366b3c1db91ce54a640",
          "rows": 83452
        },
        "reference": {
          "data_sha256": "45f3e17c5f4670ab27b47b3b496fe2a37f79068afdbf5366b3c1db91ce54a640",
          "rows": 83452
        }
      },
      "finished": {
        "monotonic_ns": 58876488670010,
        "time_ns": 1794684699678994544
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "20",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl20.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 2951322,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58874364289409,
        "time_ns": 1794684697554613533
      },
      "slice_started": {
        "monotonic_ns": 58873960657901,
        "time_ns": 1794684697150982616
      },
      "started": {
        "monotonic_ns": 58873960552264,
        "time_ns": 1794684697150876988
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr21.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl21.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an21.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58877285346914,
        "time_ns": 1794684700475670747
      },
      "annotation_started": {
        "monotonic_ns": 58876175609190,
        "time_ns": 1794684699365933815
      },
      "chromosome": 21,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "1d4a87c093eb5bf1dc897ad5e6c5f94538f316b868d37462fb6514bb5fee9de2",
          "rows": 55978
        },
        "reference": {
          "data_sha256": "1d4a87c093eb5bf1dc897ad5e6c5f94538f316b868d37462fb6514bb5fee9de2",
          "rows": 55978
        }
      },
      "finished": {
        "monotonic_ns": 58877289337978,
        "time_ns": 1794684700479662412
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "21",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl21.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 1982476,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58876175373099,
        "time_ns": 1794684699365697133
      },
      "slice_started": {
        "monotonic_ns": 58875773236954,
        "time_ns": 1794684698963561399
      },
      "started": {
        "monotonic_ns": 58875773089529,
        "time_ns": 1794684698963414144
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/17_fields/archives/chr22.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl22.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an22.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 58877434140083,
        "time_ns": 1794684700624464176
      },
      "annotation_started": {
        "monotonic_ns": 58876532647586,
        "time_ns": 1794684699722971980
      },
      "chromosome": 22,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "c4e02794b0aa94ca0a0eae05006399e26e108da4d8d7ce5aae62e855faf31ea7",
          "rows": 49540
        },
        "reference": {
          "data_sha256": "c4e02794b0aa94ca0a0eae05006399e26e108da4d8d7ce5aae62e855faf31ea7",
          "rows": 49540
        }
      },
      "finished": {
        "monotonic_ns": 58877438460272,
        "time_ns": 1794684700628784816
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "22",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl22.vcf.gz",
        "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 1765325,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 58876532314975,
        "time_ns": 1794684699722638888
      },
      "slice_started": {
        "monotonic_ns": 58876227568394,
        "time_ns": 1794684699417893088
      },
      "started": {
        "monotonic_ns": 58876227418223,
        "time_ns": 1794684699417742778
      },
      "status": "ok"
    }
  ],
  "compressed_output": true,
  "engine": "echtvar",
  "exit_code": 0,
  "expected_rows": 3933714,
  "fields": 17,
  "finished": {
    "monotonic_ns": 58877438991183,
    "time_ns": 1794684700629315587
  },
  "fresh_rows": 3933714,
  "media": "nvme",
  "scratch": "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release",
  "started": {
    "monotonic_ns": 58848350363843,
    "time_ns": 1794684671540683118
  },
  "status": "ok",
  "worker_note": "Four chromosome subprocess tasks on CPUs 0-3; each echtvar process may create additional htslib threads.",
  "workers": 4
}
```


### nvme / 540 fields

- immkv r1: ok; valid=True; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/status.json`.
  - wall_seconds=157.76; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/time.txt` (wall_seconds).
  - user_seconds=596.51; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/time.txt` (user_seconds).
  - system_seconds=12.37; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/time.txt` (system_seconds).
  - cpu_percent=385; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/time.txt` (cpu_percent).
  - max_child_rss_kib=2407068; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/time.txt` (max_child_rss_kib).
  - exit_code=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/time.txt` (Exit status).
  - cgroup_peak_bytes=2572652544; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/status.json` (memory_peak_bytes).
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/constraints.json`; confirmed=True.
  - rows=3933714; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/immkv.log`.
  - hits=3875208; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/immkv.log`.
- immkv r2: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r2-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r2-release/constraints.json`; confirmed=False.
- immkv r3: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r3-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r3-release/constraints.json`; confirmed=False.
- echtvar r1: ok; valid=True; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/status.json`.
  - wall_seconds=643.97; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/time.txt` (wall_seconds).
  - user_seconds=2190.82; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/time.txt` (user_seconds).
  - system_seconds=322.45; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/time.txt` (system_seconds).
  - cpu_percent=390; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/time.txt` (cpu_percent).
  - max_child_rss_kib=540148; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/time.txt` (max_child_rss_kib).
  - exit_code=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/time.txt` (Exit status).
  - cgroup_peak_bytes=17179869184; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/status.json` (memory_peak_bytes).
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/constraints.json`; confirmed=True.
  - All chromosome statuses successful=True; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/cell.json`.
  - chr1 evaluated=308298, wrote=308297; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr1.anno.log`.
  - chr2 evaluated=329821, wrote=329821; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr2.anno.log`.
  - chr3 evaluated=275651, wrote=275647; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr3.anno.log`.
  - chr4 evaluated=296434, wrote=296434; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr4.anno.log`.
  - chr5 evaluated=253843, wrote=253841; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr5.anno.log`.
  - chr6 evaluated=264414, wrote=264084; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr6.anno.log`.
  - chr7 evaluated=222733, wrote=222732; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr7.anno.log`.
  - chr8 evaluated=213367, wrote=213367; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr8.anno.log`.
  - chr9 evaluated=168542, wrote=168542; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr9.anno.log`.
  - chr10 evaluated=198281, wrote=198281; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr10.anno.log`.
  - chr11 evaluated=200574, wrote=200574; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr11.anno.log`.
  - chr12 evaluated=189356, wrote=189356; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr12.anno.log`.
  - chr13 evaluated=156939, wrote=156938; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr13.anno.log`.
  - chr14 evaluated=125631, wrote=125631; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr14.anno.log`.
  - chr15 evaluated=114117, wrote=114117; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr15.anno.log`.
  - chr16 evaluated=118437, wrote=118437; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr16.anno.log`.
  - chr17 evaluated=102461, wrote=102461; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr17.anno.log`.
  - chr18 evaluated=119403, wrote=119401; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr18.anno.log`.
  - chr19 evaluated=86442, wrote=86442; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr19.anno.log`.
  - chr20 evaluated=83452, wrote=83452; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr20.anno.log`.
  - chr21 evaluated=55978, wrote=55978; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr21.anno.log`.
  - chr22 evaluated=49540, wrote=49536; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr22.anno.log`.
- echtvar r2: started; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/status.json`.
  - cgroup_peak_bytes=17179869184; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/memory-observer.json` (last observed memory.peak (may precede final peak)).
  - Exclusion: scope status is started; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/constraints.json`; confirmed=True.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/cell.json`.
  - chr1 evaluated=308298, wrote=308297; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/chr1.anno.log`.
  - chr2 evaluated=329821, wrote=329821; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/chr2.anno.log`.
  - chr3 evaluated=275651, wrote=275647; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/chr3.anno.log`.
  - chr4 evaluated=296434, wrote=296434; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/chr4.anno.log`.
- echtvar r3: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r3-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r3-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r3-release/cell.json`.
- Paired wall ratios unavailable: both engines require all three valid rounds.
- Correctness immkv: scope=ok, passed=True; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verification.json`; scope source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/status.json`.

Evidence from `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verification.json`:

```json
{
  "annotate_statistics": {
    "hits": 3875208,
    "rows": 3933714
  },
  "commands": {
    "annotate": [
      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/annotate",
      "-in",
      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz",
      "-lib",
      "/home/tian/immkv_test_nvme/gnomad/immkv",
      "-fields",
      "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_540_fields/fields540.tsv",
      "-threads",
      "4",
      "-batch",
      "4096",
      "-out",
      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-540-release/output.vcf.gz"
    ],
    "verifier": [
      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/verifyvalues",
      "-ev",
      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-540-release/echtvar.fifo",
      "-iv",
      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-540-release/immkv.fifo",
      "-fields",
      "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_540_fields/fields540.tsv",
      "-tol",
      "1.1e-6",
      "-max-report",
      "20"
    ]
  },
  "comparison": {
    "echtvar_i32_overflow_losses": 808,
    "field_comparisons": 2124205560,
    "hit_disagreements": 0,
    "misses_agreed": 32050262,
    "rows": 3933714,
    "value_mismatches": 0
  },
  "elapsed_seconds_correctness_only": 975.7349105379981,
  "exit_status": {
    "annotate": 0,
    "verifier": 0
  },
  "expected_hits": 3875208,
  "expected_rows": 3933714,
  "fields": 540,
  "fields_tsv": "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_540_fields/fields540.tsv",
  "float_tolerance": 1.1e-06,
  "float_tolerance_basis": "echtvar v0.2.4 quantization truncation plus six-decimal print truncation and f32 slack",
  "library": "/home/tian/immkv_test_nvme/gnomad/immkv",
  "marker_collision_accounting": "Shared verifier includes true float -1 marker collisions in misses agreed.",
  "media": "nvme",
  "query": "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz",
  "reference_kind": "historical correctness reference; not a current performance sample",
  "reference_media": "nvme",
  "reference_release": "echtvar v0.2.4",
  "sha256": {
    "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_540_fields/fields540.tsv": "fa1f21636fcd186308e3d070a65b0c087a8813f6763a78f8db70dc96c65f7643",
    "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify_output.py": "5ff9a41ae7be63fbe25c78ecb16519f3b3cec517953792d68e4552325243467a",
    "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/annotate": "4179f6f9bf774615ce68e1794dd3c46f2c5e4b2316011b117191584f5ec16391",
    "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/verifyvalues": "8e307e09177310a07ec9615734493f9e423a779f3b6f3a2a5dd82f4364343400"
  },
  "status": "passed",
  "streams": {
    "echtvar": {
      "data_sha256": "37d131239563d193d5f802dbff870547bbeb4cad795951b2dd71f7fe2a618bf4",
      "files": [
        {
          "data_bytes": 5358350269,
          "data_sha256": "59ef94490a873f8ceb5139499102b66658a2c22d25a4e1c4fdc5166cb04bcdea",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an1.vcf.gz",
          "rows": 308298
        },
        {
          "data_bytes": 5732240580,
          "data_sha256": "e4bdc61476b8d46eb502e0b45c3393067cec0e6b2e847b9d0e688b6aa4603c40",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an2.vcf.gz",
          "rows": 329821
        },
        {
          "data_bytes": 4792633475,
          "data_sha256": "d8baaf493daa7bcaeab011f8d6f3f3e2b7c1cf3cb994d5edf14b1e2c77e6de25",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an3.vcf.gz",
          "rows": 275651
        },
        {
          "data_bytes": 5153980926,
          "data_sha256": "b0acf0f50c7a6a6abfda371fb79763cb60f506cb517699eaf71b603d25b6b1be",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an4.vcf.gz",
          "rows": 296434
        },
        {
          "data_bytes": 4412841424,
          "data_sha256": "3ffd54981a331bfd12e5913e58afb4af35f9f9eb94ae652d229b5303055e4f31",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an5.vcf.gz",
          "rows": 253843
        },
        {
          "data_bytes": 4574212905,
          "data_sha256": "6777b4b677680e6a2c5ff730a17947644bf36e14b079268ee9f93c8cdd85f81f",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an6.vcf.gz",
          "rows": 264414
        },
        {
          "data_bytes": 3871325750,
          "data_sha256": "b8d05e8429a1fd1d9cc85ca0293b1887a0cc26838a7a4afe31598cf79b028494",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an7.vcf.gz",
          "rows": 222733
        },
        {
          "data_bytes": 3710141684,
          "data_sha256": "222c4768dd63f7ac4d2394c04ce44287ceff2703c1ea5670318509aa7388ec12",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an8.vcf.gz",
          "rows": 213367
        },
        {
          "data_bytes": 2929629765,
          "data_sha256": "2eb2612e426858cc861337211691e8e303f3a181950c3e9c0c7a388c7cdddfb2",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an9.vcf.gz",
          "rows": 168542
        },
        {
          "data_bytes": 3446687952,
          "data_sha256": "258bd49bb17c322c57f1b8252de26f36133f81e918991e2fd1a1d3af07039651",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an10.vcf.gz",
          "rows": 198281
        },
        {
          "data_bytes": 3486144839,
          "data_sha256": "849c192386068b10f5da434d41a878d9813290486fa267a6c28997e147d82706",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an11.vcf.gz",
          "rows": 200574
        },
        {
          "data_bytes": 3290577146,
          "data_sha256": "8e8ad77655aecf5b143da65c43e1eaccc11a749fe5c12611490c7096655f0cf9",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an12.vcf.gz",
          "rows": 189356
        },
        {
          "data_bytes": 2727987517,
          "data_sha256": "ce4d59d8da2af38e2db8933f9ccebdcf63c45e91b2f6ac2558c9adb0505266bf",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an13.vcf.gz",
          "rows": 156939
        },
        {
          "data_bytes": 2183774964,
          "data_sha256": "ca822fb4d18bfe4626577736dbae9831c8e8de44d71f22a2b69cbcf7ea1db381",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an14.vcf.gz",
          "rows": 125631
        },
        {
          "data_bytes": 1983196172,
          "data_sha256": "de1808e7b7c482288acb0672139756800e4ea47120a9590d8cc06ef4ad4838d7",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an15.vcf.gz",
          "rows": 114117
        },
        {
          "data_bytes": 2058578699,
          "data_sha256": "ce19990fefc7c57c5241cb715ddce8d75294c8044391873945cabcc2b5e5c243",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an16.vcf.gz",
          "rows": 118437
        },
        {
          "data_bytes": 1780261584,
          "data_sha256": "6b1ab9961447b7b87d9dd796602b7b1a8aec05fe1bc838ef241932cf72851e8e",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an17.vcf.gz",
          "rows": 102461
        },
        {
          "data_bytes": 2075298740,
          "data_sha256": "30cfca328b607afd97c5d13eb3b130b1788bd449d3ba5f2ae8fd5879748666a7",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an18.vcf.gz",
          "rows": 119403
        },
        {
          "data_bytes": 1501634495,
          "data_sha256": "559bdfc5144aabecdb11ef8d67ea8a5dd71e9b7b06ee57fd0be15f8214adc371",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an19.vcf.gz",
          "rows": 86442
        },
        {
          "data_bytes": 1450566921,
          "data_sha256": "784a6057e388e9ccce45a2366647193caa076f04c270283ee912eb24cf4feea9",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an20.vcf.gz",
          "rows": 83452
        },
        {
          "data_bytes": 972937987,
          "data_sha256": "c67c4a591c15d64d9c4bf5352c88028a5e07c916354e54025297c6103f9c4407",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an21.vcf.gz",
          "rows": 55978
        },
        {
          "data_bytes": 860563548,
          "data_sha256": "5636e7ed9b75c0666497ee77366c60d2922c32c79d60c968cb519a06d2b936ac",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an22.vcf.gz",
          "rows": 49540
        }
      ],
      "rows": 3933714
    },
    "immkv": {
      "data_sha256": "9e54e963a7a9db6063232afa59a58c48f37d999289116a43e16f97bfccd59dcd",
      "files": [
        {
          "data_bytes": 67477268158,
          "data_sha256": "9e54e963a7a9db6063232afa59a58c48f37d999289116a43e16f97bfccd59dcd",
          "gzip_crc_eof_ok": true,
          "path": "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-540-release/output.vcf.gz",
          "rows": 3933714
        }
      ],
      "rows": 3933714,
      "selected_float_true_minus_one_candidates": 663
    }
  },
  "timing_sample": false
}
```

- immkv rows=3933714; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout`.
- immkv field_comparisons=2124205560; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout`.
- immkv misses_agreed=32050262; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout`.
- immkv value_mismatches=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout`.
- immkv hit_disagreements=0; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout`.
- immkv echtvar_i32_overflow_losses=808; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout`.
- immkv float_comparisons=666144645; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout`.
- immkv max_float_absolute_difference=1e-06; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout`.
- immkv float_tolerance=1.1e-06; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout`.
- immkv verify_values_ok=True; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout`.
- Correctness fresh_echtvar: scope=ok, passed=True; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-540-release/cell.json`; scope source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-540-release/status.json`.

Evidence from `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-540-release/cell.json`:

```json
{
  "chromosomes": [
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr1.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl1.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an1.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60213517147958,
        "time_ns": 1794686036707472312
      },
      "annotation_started": {
        "monotonic_ns": 59978069300611,
        "time_ns": 1794685801259625176
      },
      "chromosome": 1,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "59ef94490a873f8ceb5139499102b66658a2c22d25a4e1c4fdc5166cb04bcdea",
          "rows": 308298
        },
        "reference": {
          "data_sha256": "59ef94490a873f8ceb5139499102b66658a2c22d25a4e1c4fdc5166cb04bcdea",
          "rows": 308298
        }
      },
      "finished": {
        "monotonic_ns": 60213523918677,
        "time_ns": 1794686036714243141
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "1",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl1.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 10999771,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 59978068676555,
        "time_ns": 1794685801259000779
      },
      "slice_started": {
        "monotonic_ns": 59976865617832,
        "time_ns": 1794685800055942226
      },
      "started": {
        "monotonic_ns": 59976865269055,
        "time_ns": 1794685800055593329
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr2.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl2.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an2.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60226951375671,
        "time_ns": 1794686050141699494
      },
      "annotation_started": {
        "monotonic_ns": 59978074330122,
        "time_ns": 1794685801264654736
      },
      "chromosome": 2,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "e4bdc61476b8d46eb502e0b45c3393067cec0e6b2e847b9d0e688b6aa4603c40",
          "rows": 329821
        },
        "reference": {
          "data_sha256": "e4bdc61476b8d46eb502e0b45c3393067cec0e6b2e847b9d0e688b6aa4603c40",
          "rows": 329821
        }
      },
      "finished": {
        "monotonic_ns": 60226955317460,
        "time_ns": 1794686050145641804
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "2",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl2.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 11750290,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 59978070169519,
        "time_ns": 1794685801260494134
      },
      "slice_started": {
        "monotonic_ns": 59976866134105,
        "time_ns": 1794685800056458519
      },
      "started": {
        "monotonic_ns": 59976865752045,
        "time_ns": 1794685800056076529
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr3.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl3.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an3.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60185448028221,
        "time_ns": 1794686008638352404
      },
      "annotation_started": {
        "monotonic_ns": 59977870272607,
        "time_ns": 1794685801060597311
      },
      "chromosome": 3,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "d8baaf493daa7bcaeab011f8d6f3f3e2b7c1cf3cb994d5edf14b1e2c77e6de25",
          "rows": 275651
        },
        "reference": {
          "data_sha256": "d8baaf493daa7bcaeab011f8d6f3f3e2b7c1cf3cb994d5edf14b1e2c77e6de25",
          "rows": 275651
        }
      },
      "finished": {
        "monotonic_ns": 60185450850578,
        "time_ns": 1794686008641174992
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "3",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl3.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 9815086,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 59977869739913,
        "time_ns": 1794685801060063856
      },
      "slice_started": {
        "monotonic_ns": 59976866858140,
        "time_ns": 1794685800057182815
      },
      "started": {
        "monotonic_ns": 59976866408823,
        "time_ns": 1794685800056733327
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr4.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl4.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an4.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60197305801714,
        "time_ns": 1794686020496125638
      },
      "annotation_started": {
        "monotonic_ns": 59977973111441,
        "time_ns": 1794685801163436095
      },
      "chromosome": 4,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "b0acf0f50c7a6a6abfda371fb79763cb60f506cb517699eaf71b603d25b6b1be",
          "rows": 296434
        },
        "reference": {
          "data_sha256": "b0acf0f50c7a6a6abfda371fb79763cb60f506cb517699eaf71b603d25b6b1be",
          "rows": 296434
        }
      },
      "finished": {
        "monotonic_ns": 60197316128410,
        "time_ns": 1794686020506452653
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "4",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl4.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 10519954,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 59977968617340,
        "time_ns": 1794685801158941614
      },
      "slice_started": {
        "monotonic_ns": 59976867201247,
        "time_ns": 1794685800057525821
      },
      "started": {
        "monotonic_ns": 59976866997803,
        "time_ns": 1794685800057322277
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr5.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl5.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an5.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60368941157374,
        "time_ns": 1794686192131481558
      },
      "annotation_started": {
        "monotonic_ns": 60186563782881,
        "time_ns": 1794686009754107545
      },
      "chromosome": 5,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "3ffd54981a331bfd12e5913e58afb4af35f9f9eb94ae652d229b5303055e4f31",
          "rows": 253843
        },
        "reference": {
          "data_sha256": "3ffd54981a331bfd12e5913e58afb4af35f9f9eb94ae652d229b5303055e4f31",
          "rows": 253843
        }
      },
      "finished": {
        "monotonic_ns": 60368945155121,
        "time_ns": 1794686192135479525
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "5",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl5.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 9032175,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60186563476955,
        "time_ns": 1794686009753800899
      },
      "slice_started": {
        "monotonic_ns": 60185452023934,
        "time_ns": 1794686008642348659
      },
      "started": {
        "monotonic_ns": 60185451901724,
        "time_ns": 1794686008642226228
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr6.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl6.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an6.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60381987387110,
        "time_ns": 1794686205177711333
      },
      "annotation_started": {
        "monotonic_ns": 60198326955843,
        "time_ns": 1794686021517280437
      },
      "chromosome": 6,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "6777b4b677680e6a2c5ff730a17947644bf36e14b079268ee9f93c8cdd85f81f",
          "rows": 264414
        },
        "reference": {
          "data_sha256": "6777b4b677680e6a2c5ff730a17947644bf36e14b079268ee9f93c8cdd85f81f",
          "rows": 264414
        }
      },
      "finished": {
        "monotonic_ns": 60381990448808,
        "time_ns": 1794686205180773282
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "6",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl6.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 8584380,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60198326645990,
        "time_ns": 1794686021516970104
      },
      "slice_started": {
        "monotonic_ns": 60197316554882,
        "time_ns": 1794686020506879616
      },
      "started": {
        "monotonic_ns": 60197316324047,
        "time_ns": 1794686020506648792
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr7.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl7.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an7.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60375981197353,
        "time_ns": 1794686199171521506
      },
      "annotation_started": {
        "monotonic_ns": 60214433447966,
        "time_ns": 1794686037623772570
      },
      "chromosome": 7,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "b8d05e8429a1fd1d9cc85ca0293b1887a0cc26838a7a4afe31598cf79b028494",
          "rows": 222733
        },
        "reference": {
          "data_sha256": "b8d05e8429a1fd1d9cc85ca0293b1887a0cc26838a7a4afe31598cf79b028494",
          "rows": 222733
        }
      },
      "finished": {
        "monotonic_ns": 60375985356622,
        "time_ns": 1794686199175681066
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "7",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl7.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 7924912,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60214432305138,
        "time_ns": 1794686037622629361
      },
      "slice_started": {
        "monotonic_ns": 60213525743607,
        "time_ns": 1794686036716068282
      },
      "started": {
        "monotonic_ns": 60213525027892,
        "time_ns": 1794686036715352597
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr8.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl8.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an8.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60381817108219,
        "time_ns": 1794686205007432182
      },
      "annotation_started": {
        "monotonic_ns": 60227860024246,
        "time_ns": 1794686051050348941
      },
      "chromosome": 8,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "222c4768dd63f7ac4d2394c04ce44287ceff2703c1ea5670318509aa7388ec12",
          "rows": 213367
        },
        "reference": {
          "data_sha256": "222c4768dd63f7ac4d2394c04ce44287ceff2703c1ea5670318509aa7388ec12",
          "rows": 213367
        }
      },
      "finished": {
        "monotonic_ns": 60381824113592,
        "time_ns": 1794686205014437946
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "8",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl8.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 7556433,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60227859629484,
        "time_ns": 1794686051049953347
      },
      "slice_started": {
        "monotonic_ns": 60226956090804,
        "time_ns": 1794686050146415569
      },
      "started": {
        "monotonic_ns": 60226955554165,
        "time_ns": 1794686050145878930
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr9.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl9.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an9.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60489694515719,
        "time_ns": 1794686312884839913
      },
      "annotation_started": {
        "monotonic_ns": 60369752958186,
        "time_ns": 1794686192943282891
      },
      "chromosome": 9,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "2eb2612e426858cc861337211691e8e303f3a181950c3e9c0c7a388c7cdddfb2",
          "rows": 168542
        },
        "reference": {
          "data_sha256": "2eb2612e426858cc861337211691e8e303f3a181950c3e9c0c7a388c7cdddfb2",
          "rows": 168542
        }
      },
      "finished": {
        "monotonic_ns": 60489697371127,
        "time_ns": 1794686312887694980
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "9",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl9.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 5978980,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60369752694612,
        "time_ns": 1794686192943018505
      },
      "slice_started": {
        "monotonic_ns": 60368945549822,
        "time_ns": 1794686192135874476
      },
      "started": {
        "monotonic_ns": 60368945403207,
        "time_ns": 1794686192135727841
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr10.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl10.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an10.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60521040171064,
        "time_ns": 1794686344230495128
      },
      "annotation_started": {
        "monotonic_ns": 60376794044130,
        "time_ns": 1794686199984368574
      },
      "chromosome": 10,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "258bd49bb17c322c57f1b8252de26f36133f81e918991e2fd1a1d3af07039651",
          "rows": 198281
        },
        "reference": {
          "data_sha256": "258bd49bb17c322c57f1b8252de26f36133f81e918991e2fd1a1d3af07039651",
          "rows": 198281
        }
      },
      "finished": {
        "monotonic_ns": 60521042971488,
        "time_ns": 1794686344233296032
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "10",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl10.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 7017518,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60376792234272,
        "time_ns": 1794686199982558445
      },
      "slice_started": {
        "monotonic_ns": 60375987170187,
        "time_ns": 1794686199177494661
      },
      "started": {
        "monotonic_ns": 60375987034753,
        "time_ns": 1794686199177359287
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr11.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl11.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an11.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60530833787072,
        "time_ns": 1794686354024110985
      },
      "annotation_started": {
        "monotonic_ns": 60382629434130,
        "time_ns": 1794686205819758775
      },
      "chromosome": 11,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "849c192386068b10f5da434d41a878d9813290486fa267a6c28997e147d82706",
          "rows": 200574
        },
        "reference": {
          "data_sha256": "849c192386068b10f5da434d41a878d9813290486fa267a6c28997e147d82706",
          "rows": 200574
        }
      },
      "finished": {
        "monotonic_ns": 60530840209640,
        "time_ns": 1794686354030534315
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "11",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl11.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 7109132,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60382629186726,
        "time_ns": 1794686205819510669
      },
      "slice_started": {
        "monotonic_ns": 60381826009982,
        "time_ns": 1794686205016334557
      },
      "started": {
        "monotonic_ns": 60381825843510,
        "time_ns": 1794686205016168044
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr12.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl12.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an12.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60524136591196,
        "time_ns": 1794686347326915289
      },
      "annotation_started": {
        "monotonic_ns": 60382796963409,
        "time_ns": 1794686205987287622
      },
      "chromosome": 12,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "8e8ad77655aecf5b143da65c43e1eaccc11a749fe5c12611490c7096655f0cf9",
          "rows": 189356
        },
        "reference": {
          "data_sha256": "8e8ad77655aecf5b143da65c43e1eaccc11a749fe5c12611490c7096655f0cf9",
          "rows": 189356
        }
      },
      "finished": {
        "monotonic_ns": 60524141814916,
        "time_ns": 1794686347332139390
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "12",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl12.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 6749307,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60382796511530,
        "time_ns": 1794686205986835373
      },
      "slice_started": {
        "monotonic_ns": 60381990757438,
        "time_ns": 1794686205181082182
      },
      "started": {
        "monotonic_ns": 60381990623516,
        "time_ns": 1794686205180948301
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr13.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl13.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an13.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60599369222246,
        "time_ns": 1794686422559546520
      },
      "annotation_started": {
        "monotonic_ns": 60490408340337,
        "time_ns": 1794686313598664991
      },
      "chromosome": 13,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "ce4d59d8da2af38e2db8933f9ccebdcf63c45e91b2f6ac2558c9adb0505266bf",
          "rows": 156939
        },
        "reference": {
          "data_sha256": "ce4d59d8da2af38e2db8933f9ccebdcf63c45e91b2f6ac2558c9adb0505266bf",
          "rows": 156939
        }
      },
      "finished": {
        "monotonic_ns": 60599372025824,
        "time_ns": 1794686422562349727
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "13",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl13.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 5531589,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60490405216786,
        "time_ns": 1794686313595540839
      },
      "slice_started": {
        "monotonic_ns": 60489698585044,
        "time_ns": 1794686312888909689
      },
      "started": {
        "monotonic_ns": 60489697569539,
        "time_ns": 1794686312887894094
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr14.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl14.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an14.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60613663567966,
        "time_ns": 1794686436853892070
      },
      "annotation_started": {
        "monotonic_ns": 60521649645055,
        "time_ns": 1794686344839969740
      },
      "chromosome": 14,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "ca822fb4d18bfe4626577736dbae9831c8e8de44d71f22a2b69cbcf7ea1db381",
          "rows": 125631
        },
        "reference": {
          "data_sha256": "ca822fb4d18bfe4626577736dbae9831c8e8de44d71f22a2b69cbcf7ea1db381",
          "rows": 125631
        }
      },
      "finished": {
        "monotonic_ns": 60613669244221,
        "time_ns": 1794686436859568535
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "14",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl14.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 4486825,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60521649349641,
        "time_ns": 1794686344839673624
      },
      "slice_started": {
        "monotonic_ns": 60521043266802,
        "time_ns": 1794686344233591547
      },
      "started": {
        "monotonic_ns": 60521043163778,
        "time_ns": 1794686344233488523
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr15.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl15.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an15.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60609324732844,
        "time_ns": 1794686432515056838
      },
      "annotation_started": {
        "monotonic_ns": 60524751195425,
        "time_ns": 1794686347941519058
      },
      "chromosome": 15,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "de1808e7b7c482288acb0672139756800e4ea47120a9590d8cc06ef4ad4838d7",
          "rows": 114117
        },
        "reference": {
          "data_sha256": "de1808e7b7c482288acb0672139756800e4ea47120a9590d8cc06ef4ad4838d7",
          "rows": 114117
        }
      },
      "finished": {
        "monotonic_ns": 60609328353404,
        "time_ns": 1794686432518677718
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "15",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl15.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 4074508,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60524747289587,
        "time_ns": 1794686347937613681
      },
      "slice_started": {
        "monotonic_ns": 60524142133323,
        "time_ns": 1794686347332457998
      },
      "started": {
        "monotonic_ns": 60524141989704,
        "time_ns": 1794686347332314459
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr16.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl16.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an16.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60617904692483,
        "time_ns": 1794686441095016456
      },
      "annotation_started": {
        "monotonic_ns": 60531449723511,
        "time_ns": 1794686354640048156
      },
      "chromosome": 16,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "ce19990fefc7c57c5241cb715ddce8d75294c8044391873945cabcc2b5e5c243",
          "rows": 118437
        },
        "reference": {
          "data_sha256": "ce19990fefc7c57c5241cb715ddce8d75294c8044391873945cabcc2b5e5c243",
          "rows": 118437
        }
      },
      "finished": {
        "monotonic_ns": 60617908251206,
        "time_ns": 1794686441098575580
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "16",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl16.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 4212957,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60531449492177,
        "time_ns": 1794686354639816311
      },
      "slice_started": {
        "monotonic_ns": 60530840510004,
        "time_ns": 1794686354030834689
      },
      "started": {
        "monotonic_ns": 60530840365332,
        "time_ns": 1794686354030690107
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr17.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl17.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an17.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60675360863979,
        "time_ns": 1794686498551187812
      },
      "annotation_started": {
        "monotonic_ns": 60599879383745,
        "time_ns": 1794686423069708059
      },
      "chromosome": 17,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "6b1ab9961447b7b87d9dd796602b7b1a8aec05fe1bc838ef241932cf72851e8e",
          "rows": 102461
        },
        "reference": {
          "data_sha256": "6b1ab9961447b7b87d9dd796602b7b1a8aec05fe1bc838ef241932cf72851e8e",
          "rows": 102461
        }
      },
      "finished": {
        "monotonic_ns": 60675363809992,
        "time_ns": 1794686498554134597
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "17",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl17.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 3711657,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60599879143946,
        "time_ns": 1794686423069467949
      },
      "slice_started": {
        "monotonic_ns": 60599372298706,
        "time_ns": 1794686422562623470
      },
      "started": {
        "monotonic_ns": 60599372192546,
        "time_ns": 1794686422562517071
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr18.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl18.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an18.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60691654748764,
        "time_ns": 1794686514845072567
      },
      "annotation_started": {
        "monotonic_ns": 60609933370620,
        "time_ns": 1794686433123695194
      },
      "chromosome": 18,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "30cfca328b607afd97c5d13eb3b130b1788bd449d3ba5f2ae8fd5879748666a7",
          "rows": 119403
        },
        "reference": {
          "data_sha256": "30cfca328b607afd97c5d13eb3b130b1788bd449d3ba5f2ae8fd5879748666a7",
          "rows": 119403
        }
      },
      "finished": {
        "monotonic_ns": 60691657301319,
        "time_ns": 1794686514847625854
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "18",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl18.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 4239924,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60609932327754,
        "time_ns": 1794686433122651958
      },
      "slice_started": {
        "monotonic_ns": 60609328651183,
        "time_ns": 1794686432518975957
      },
      "started": {
        "monotonic_ns": 60609328537059,
        "time_ns": 1794686432518861733
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr19.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl19.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an19.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60680132272346,
        "time_ns": 1794686503322596520
      },
      "annotation_started": {
        "monotonic_ns": 60614073366609,
        "time_ns": 1794686437263690783
      },
      "chromosome": 19,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "559bdfc5144aabecdb11ef8d67ea8a5dd71e9b7b06ee57fd0be15f8214adc371",
          "rows": 86442
        },
        "reference": {
          "data_sha256": "559bdfc5144aabecdb11ef8d67ea8a5dd71e9b7b06ee57fd0be15f8214adc371",
          "rows": 86442
        }
      },
      "finished": {
        "monotonic_ns": 60680134459777,
        "time_ns": 1794686503324784171
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "19",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl19.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 3140430,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60614070204760,
        "time_ns": 1794686437260528944
      },
      "slice_started": {
        "monotonic_ns": 60613669576975,
        "time_ns": 1794686436859901720
      },
      "started": {
        "monotonic_ns": 60613669438776,
        "time_ns": 1794686436859763361
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr20.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl20.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an20.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60679850403122,
        "time_ns": 1794686503040727035
      },
      "annotation_started": {
        "monotonic_ns": 60618314403307,
        "time_ns": 1794686441504728062
      },
      "chromosome": 20,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "784a6057e388e9ccce45a2366647193caa076f04c270283ee912eb24cf4feea9",
          "rows": 83452
        },
        "reference": {
          "data_sha256": "784a6057e388e9ccce45a2366647193caa076f04c270283ee912eb24cf4feea9",
          "rows": 83452
        }
      },
      "finished": {
        "monotonic_ns": 60679852698545,
        "time_ns": 1794686503043022929
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "20",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl20.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 2951322,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60618313989662,
        "time_ns": 1794686441504313845
      },
      "slice_started": {
        "monotonic_ns": 60617908587657,
        "time_ns": 1794686441098912432
      },
      "started": {
        "monotonic_ns": 60617908451792,
        "time_ns": 1794686441098776467
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr21.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl21.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an21.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60708600372198,
        "time_ns": 1794686531790696332
      },
      "annotation_started": {
        "monotonic_ns": 60675766634401,
        "time_ns": 1794686498956958925
      },
      "chromosome": 21,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "c67c4a591c15d64d9c4bf5352c88028a5e07c916354e54025297c6103f9c4407",
          "rows": 55978
        },
        "reference": {
          "data_sha256": "c67c4a591c15d64d9c4bf5352c88028a5e07c916354e54025297c6103f9c4407",
          "rows": 55978
        }
      },
      "finished": {
        "monotonic_ns": 60708603315075,
        "time_ns": 1794686531793639499
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "21",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl21.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 1982474,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60675766413897,
        "time_ns": 1794686498956737961
      },
      "slice_started": {
        "monotonic_ns": 60675364307125,
        "time_ns": 1794686498554631859
      },
      "started": {
        "monotonic_ns": 60675364144159,
        "time_ns": 1794686498554468513
      },
      "status": "ok"
    },
    {
      "annotation_command": [
        "/home/tian/projects/immkv/tmp/echtvar/echtvar",
        "anno",
        "-e",
        "/home/tian/immkv_test_nvme/test/540_fields/archives/chr22.echtvar.zip",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl22.vcf.gz",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an22.vcf.gz"
      ],
      "annotation_exit_code": 0,
      "annotation_finished": {
        "monotonic_ns": 60709865367835,
        "time_ns": 1794686533055691287
      },
      "annotation_started": {
        "monotonic_ns": 60680155352908,
        "time_ns": 1794686503345677302
      },
      "chromosome": 22,
      "comparison": {
        "equal": true,
        "fresh": {
          "data_sha256": "5636e7ed9b75c0666497ee77366c60d2922c32c79d60c968cb519a06d2b936ac",
          "rows": 49540
        },
        "reference": {
          "data_sha256": "5636e7ed9b75c0666497ee77366c60d2922c32c79d60c968cb519a06d2b936ac",
          "rows": 49540
        }
      },
      "finished": {
        "monotonic_ns": 60709869724061,
        "time_ns": 1794686533060048235
      },
      "slice_command": [
        "/home/tian/miniforge3/envs/transvar/bin/bcftools",
        "view",
        "-r",
        "22",
        "-Oz",
        "-o",
        "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl22.vcf.gz",
        "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
      ],
      "slice_compressed_bytes": 1765323,
      "slice_exit_code": 0,
      "slice_finished": {
        "monotonic_ns": 60680155132064,
        "time_ns": 1794686503345456308
      },
      "slice_started": {
        "monotonic_ns": 60679852979742,
        "time_ns": 1794686503043304397
      },
      "started": {
        "monotonic_ns": 60679852873262,
        "time_ns": 1794686503043197847
      },
      "status": "ok"
    }
  ],
  "compressed_output": true,
  "engine": "echtvar",
  "exit_code": 0,
  "expected_rows": 3933714,
  "fields": 540,
  "finished": {
    "monotonic_ns": 60709870599803,
    "time_ns": 1794686533060924447
  },
  "fresh_rows": 3933714,
  "media": "nvme",
  "scratch": "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release",
  "started": {
    "monotonic_ns": 59976865033180,
    "time_ns": 1794685800055354068
  },
  "status": "ok",
  "worker_note": "Four chromosome subprocess tasks on CPUs 0-3; each echtvar process may create additional htslib threads.",
  "workers": 4
}
```


### nvme / 568 fields

- immkv r1: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r1-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r1-release/constraints.json`; confirmed=False.
- immkv r2: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r2-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r2-release/constraints.json`; confirmed=False.
- immkv r3: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r3-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r3-release/constraints.json`; confirmed=False.
- echtvar r1: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r1-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r1-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r1-release/cell.json`.
- echtvar r2: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r2-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r2-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r2-release/cell.json`.
- echtvar r3: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r3-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r3-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r3-release/cell.json`.
- Paired wall ratios unavailable: both engines require all three valid rounds.
- Correctness immkv: scope=unknown, passed=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-568-release/verification.json`; scope source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-568-release/status.json`.
- immkv verify_values_ok=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-568-release/verifier.stdout`.
- Correctness fresh_echtvar: scope=unknown, passed=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-568-release/cell.json`; scope source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-568-release/status.json`.

### hdd / 17 fields

- immkv r1: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r1-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r1-release/constraints.json`; confirmed=False.
- immkv r2: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r2-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r2-release/constraints.json`; confirmed=False.
- immkv r3: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r3-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r3-release/constraints.json`; confirmed=False.
- echtvar r1: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r1-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r1-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r1-release/cell.json`.
- echtvar r2: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r2-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r2-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r2-release/cell.json`.
- echtvar r3: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r3-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r3-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r3-release/cell.json`.
- Paired wall ratios unavailable: both engines require all three valid rounds.
- Correctness immkv: scope=unknown, passed=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-17-release/verification.json`; scope source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-17-release/status.json`.
- immkv verify_values_ok=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-17-release/verifier.stdout`.
- Correctness fresh_echtvar: scope=unknown, passed=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-17-release/cell.json`; scope source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-17-release/status.json`.

### hdd / 540 fields

- immkv r1: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r1-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r1-release/constraints.json`; confirmed=False.
- immkv r2: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r2-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r2-release/constraints.json`; confirmed=False.
- immkv r3: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r3-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r3-release/constraints.json`; confirmed=False.
- echtvar r1: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r1-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r1-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r1-release/cell.json`.
- echtvar r2: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r2-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r2-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r2-release/cell.json`.
- echtvar r3: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r3-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r3-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r3-release/cell.json`.
- Paired wall ratios unavailable: both engines require all three valid rounds.
- Correctness immkv: scope=unknown, passed=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-540-release/verification.json`; scope source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-540-release/status.json`.
- immkv verify_values_ok=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-540-release/verifier.stdout`.
- Correctness fresh_echtvar: scope=unknown, passed=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-540-release/cell.json`; scope source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-540-release/status.json`.

### hdd / 568 fields

- immkv r1: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r1-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r1-release/constraints.json`; confirmed=False.
- immkv r2: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r2-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r2-release/constraints.json`; confirmed=False.
- immkv r3: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r3-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r3-release/constraints.json`; confirmed=False.
- echtvar r1: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r1-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r1-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r1-release/cell.json`.
- echtvar r2: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r2-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r2-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r2-release/cell.json`.
- echtvar r3: unknown; valid=False; status source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r3-release/status.json`.
  - Exclusion: scope status is unknown; scope exit_code is absent or nonzero; cell status/exit_code is absent or unsuccessful; actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed; required annotation counts or chromosome completion are unconfirmed; one or more GNU time / cgroup metrics are absent; final aggregate cgroup peak is unconfirmed; GNU time exit status is absent or nonzero.
  - Actual scope constraints: `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r3-release/constraints.json`; confirmed=False.
  - All chromosome statuses successful=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r3-release/cell.json`.
- Paired wall ratios unavailable: both engines require all three valid rounds.
- Correctness immkv: scope=unknown, passed=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-568-release/verification.json`; scope source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-568-release/status.json`.
- immkv verify_values_ok=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-568-release/verifier.stdout`.
- Correctness fresh_echtvar: scope=unknown, passed=False; `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-568-release/cell.json`; scope source `/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-568-release/status.json`.

## Binary, source and input provenance

```json
{
  "build/binaries.sha256": {
    "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/build/binaries.sha256",
    "text": "4179f6f9bf774615ce68e1794dd3c46f2c5e4b2316011b117191584f5ec16391  annotate\n"
  },
  "build/build-info.txt": {
    "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/build/build-info.txt",
    "text": "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/annotate: go1.27.1\n\tpath\tgithub.com/viktorxia/immkv/test/annotate\n\tmod\tgithub.com/viktorxia/immkv\tv0.0.3+dirty\t\n\tdep\tgithub.com/biogo/hts\tv1.4.5\th1:mhVCpZaTYlAhBjMaAATGWBnauioBtmvOb0ApLdU4/+0=\n\tdep\tgithub.com/klauspost/compress\tv1.20.3\th1:zWNoFqq/lSuYQRim3dUXlXeA0IqwuZGicdG5ZTfY5j4=\n\tdep\tgithub.com/valyala/gozstd\tv1.26.0\th1:vvQxZ4FANU2Q6FG4dUMpuQ8PUvkq9MdbU2jFkF+brm8=\n\tdep\tgolang.org/x/sys\tv0.49.0\th1:XbzkgYJHdqh/8m2Uu0W/dQv8nktxx4BFHp1M0gROTXA=\n\tbuild\t-buildmode=exe\n\tbuild\t-compiler=gc\n\tbuild\t-trimpath=true\n\tbuild\tDefaultGODEBUG=tracebacklabels=0,x509sslcertoverrideplatform=0\n\tbuild\tCGO_ENABLED=1\n\tbuild\tGOARCH=amd64\n\tbuild\tGOOS=linux\n\tbuild\tGOAMD64=v1\n\tbuild\tvcs=git\n\tbuild\tvcs.revision=840d36d80351f8b8651d9dc90abf1314bf5db21c\n\tbuild\tvcs.time=2026-11-14T19:17:41Z\n\tbuild\tvcs.modified=true\n"
  },
  "build/environment.txt": {
    "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/build/environment.txt",
    "text": "2026-11-15T03:26:55+08:00\ngo version go1.27.1 linux/amd64\n1\nlinux\namd64\n/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/cache/go-build\n/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/cache/mod\n-mod=readonly\nLinux tian-mini 7.0.0-42-generic #42~24.04.1-Ubuntu SMP PREEMPT_DYNAMIC Mon Oct 19 14:43:01 UTC 2 x86_64 x86_64 x86_64 GNU/Linux\nechtvar 0.2.4\nbcftools 1.25\nUsing htslib 1.25\nCopyright (C) 2026 Genome Research Ltd.\nLicense GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>\nThis is free software: you are free to change and redistribute it.\nThere is NO WARRANTY, to the extent permitted by law.\n2c425bad9a4e6aeeade5eeec229e7f887d7385642a0aada3a03909b6303f1e09  /home/tian/projects/immkv/tmp/echtvar/echtvar\n728cc52cd0a82eed97995b78622f882d9e19c9b947b687c8bbb295065aca9c58  /home/tian/miniforge3/envs/transvar/bin/bcftools\n"
  },
  "build/source-facts.json": {
    "issues": [],
    "record": {
      "annotate": "Unchanged test/annotate from release git archive",
      "archive_sha256": "d5bd8a2545d080bbb8992cb7ff4bf5518480193f131899fabae11dc33c06941e",
      "commit": "840d36d80351f8b8651d9dc90abf1314bf5db21c",
      "dependencies": [
        [
          "github.com/klauspost/compress",
          "v1.20.3"
        ],
        [
          "github.com/valyala/gozstd",
          "v1.26.0"
        ],
        [
          "golang.org/x/sys",
          "v0.49.0"
        ],
        [
          "github.com/biogo/hts",
          "v1.4.5"
        ]
      ],
      "module_cache_origin": "/home/tian/go/pkg/mod"
    },
    "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/build/source-facts.json"
  },
  "input-preflight.json": {
    "issues": [],
    "record": {
      "created_ns": 1794684432117851506,
      "media": {
        "hdd": {
          "17": {
            "archives": [
              {
                "bytes": 244234654,
                "central_directory_members": 4163,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr1.echtvar.zip",
                "uncompressed_bytes": 578567882
              },
              {
                "bytes": 258657713,
                "central_directory_members": 4391,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr2.echtvar.zip",
                "uncompressed_bytes": 620538239
              },
              {
                "bytes": 210659980,
                "central_directory_members": 3555,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr3.echtvar.zip",
                "uncompressed_bytes": 505087392
              },
              {
                "bytes": 208828208,
                "central_directory_members": 3441,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr4.echtvar.zip",
                "uncompressed_bytes": 494714168
              },
              {
                "bytes": 192352387,
                "central_directory_members": 3251,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr5.echtvar.zip",
                "uncompressed_bytes": 461270809
              },
              {
                "bytes": 183934502,
                "central_directory_members": 3080,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr6.echtvar.zip",
                "uncompressed_bytes": 432604103
              },
              {
                "bytes": 178674759,
                "central_directory_members": 2852,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr7.echtvar.zip",
                "uncompressed_bytes": 420231928
              },
              {
                "bytes": 164725132,
                "central_directory_members": 2624,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr8.echtvar.zip",
                "uncompressed_bytes": 397702935
              },
              {
                "bytes": 138388260,
                "central_directory_members": 2263,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr9.echtvar.zip",
                "uncompressed_bytes": 325968828
              },
              {
                "bytes": 147108442,
                "central_directory_members": 2434,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr10.echtvar.zip",
                "uncompressed_bytes": 345447116
              },
              {
                "bytes": 146764897,
                "central_directory_members": 2415,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr11.echtvar.zip",
                "uncompressed_bytes": 350042016
              },
              {
                "bytes": 144193515,
                "central_directory_members": 2396,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr12.echtvar.zip",
                "uncompressed_bytes": 339777686
              },
              {
                "bytes": 104972306,
                "central_directory_members": 1750,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr13.echtvar.zip",
                "uncompressed_bytes": 247521444
              },
              {
                "bytes": 99273295,
                "central_directory_members": 1617,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr14.echtvar.zip",
                "uncompressed_bytes": 233719970
              },
              {
                "bytes": 93455566,
                "central_directory_members": 1503,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr15.echtvar.zip",
                "uncompressed_bytes": 219684326
              },
              {
                "bytes": 102348503,
                "central_directory_members": 1465,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr16.echtvar.zip",
                "uncompressed_bytes": 244290082
              },
              {
                "bytes": 89862907,
                "central_directory_members": 1446,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr17.echtvar.zip",
                "uncompressed_bytes": 210528323
              },
              {
                "bytes": 83163449,
                "central_directory_members": 1389,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr18.echtvar.zip",
                "uncompressed_bytes": 196324184
              },
              {
                "bytes": 75601696,
                "central_directory_members": 1047,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr19.echtvar.zip",
                "uncompressed_bytes": 172173473
              },
              {
                "bytes": 66291586,
                "central_directory_members": 1123,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr20.echtvar.zip",
                "uncompressed_bytes": 157162354
              },
              {
                "bytes": 43014863,
                "central_directory_members": 686,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr21.echtvar.zip",
                "uncompressed_bytes": 99628454
              },
              {
                "bytes": 44732754,
                "central_directory_members": 648,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr22.echtvar.zip",
                "uncompressed_bytes": 102376807
              }
            ],
            "fields_sha256": "db2c8bc29926b5f4265e9077443a19571a55c8f99d2d30906e0a126c6c0db030",
            "index_bytes": 1367830,
            "library": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/lib17",
            "manifests": [
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 1,
                "manifest_sha256": "b828cfa17a2df725f18e9907cee4cf0e9907c433d0ab5992ad6a25651f59852d",
                "records": 20370890,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 2,
                "manifest_sha256": "c99e8c2b512d651fefdba28fd4bc47ebc1ca96b816393fd8c77d08245a868505",
                "records": 21892245,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 3,
                "manifest_sha256": "88c35cffd4cc14950f915a4b39d1f466ccd0bcf1d49890a29beb47436dbf4f55",
                "records": 17834355,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 4,
                "manifest_sha256": "e310f6b01626703f969a903e630f029d97d3fdb089de0c8e3a5a1cb4b3f2a23c",
                "records": 17423639,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 5,
                "manifest_sha256": "96bfbf63bfb6181918baa760bdb8ff6ec675ef1459e2241ce0eb6eee1f8490f4",
                "records": 16290654,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 6,
                "manifest_sha256": "9a77beac2449ac85a8cf028e7769b9568fc7dc168c7f82134761233c8a6983ca",
                "records": 15190442,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 7,
                "manifest_sha256": "3e6c4b0b97e2e8acc707f9001fba3c96b7583aa761f285fb9187726b0737467f",
                "records": 14752321,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 8,
                "manifest_sha256": "a0de0210e73dbd886181a9c6e3b39e17854efabb53aa95d2f39114ab5de8ed64",
                "records": 14073574,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 9,
                "manifest_sha256": "90d4ded4deb27f1c1ffadebe47f5c65cf45455e657191676a4c23c557b61f8a9",
                "records": 11480614,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 10,
                "manifest_sha256": "12d27572091e8cb3fbb27386b5e777f73e330643f72d18ac8ea79618f465d547",
                "records": 12118750,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 11,
                "manifest_sha256": "c9343209324cedb76d8336d0b162e13adbde8ac07357906e101d816042827b3c",
                "records": 12345492,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 12,
                "manifest_sha256": "9c93500cf3fdbdde0991afd4a9a42602419f56be02fc0ede49e2268dec415de8",
                "records": 11933435,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 13,
                "manifest_sha256": "ee367aa983c8d31b9ef6485966954bc84c840c308f9ec10673955eea0ba8f052",
                "records": 8693573,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 14,
                "manifest_sha256": "2336388c4f0d6324d1d391bcbac4e81f0e32e95a343bb1a6872d1879f8446f2c",
                "records": 8220638,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 15,
                "manifest_sha256": "ca21551ad8d8d7b5e92ff40d830393ba374c9aaa8e18a45017433de9ea1c9b29",
                "records": 7723667,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 16,
                "manifest_sha256": "43209ec7835c225072e05d34144afba16d6bba7d83b747fccceb6f018d837cc7",
                "records": 8618657,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 17,
                "manifest_sha256": "57583441aab447fa2243d6939f679e50098ab45bb207b93820a679c6d36b6488",
                "records": 7373819,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 18,
                "manifest_sha256": "9a8bb1ae9a643ee3388533082705d2e26bcd09ab03cda90b64e8965d509cd369",
                "records": 6893751,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 19,
                "manifest_sha256": "c905eb1ca428e5fc53d8d67e0c5d4fdc0ee719800d8e089e3f3b794d6c80510a",
                "records": 5991870,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 20,
                "manifest_sha256": "fc4513940ce0caed90006e84a126eb2660c11f6e07a28c1cf70d352e89d529a3",
                "records": 5516836,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 21,
                "manifest_sha256": "5a064dcd7b265753c4e4e49d31e6aec772357bba315257e7d1b0a8e78d6c1772",
                "records": 3483000,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 22,
                "manifest_sha256": "d5987ea72aaf88492de717ab4f935c2108bec30b957d63357268d66bafff178c",
                "records": 3567085,
                "schema_hash": "99944e3316871f58"
              }
            ],
            "query": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/query/hg001.decomp.vcf.gz",
            "query_bytes": 139085616,
            "query_sha256": "b9afb10b724c9b85c59e723be330a3ea756928c9ca2092692db3b1faa64472d0"
          },
          "540": {
            "archives": [
              {
                "bytes": 5123651180,
                "central_directory_members": 118700,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr1.echtvar.zip",
                "uncompressed_bytes": 17527475456
              },
              {
                "bytes": 5398570861,
                "central_directory_members": 125204,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr2.echtvar.zip",
                "uncompressed_bytes": 18815801585
              },
              {
                "bytes": 4396709608,
                "central_directory_members": 101356,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr3.echtvar.zip",
                "uncompressed_bytes": 15334578110
              },
              {
                "bytes": 4372074758,
                "central_directory_members": 98104,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr4.echtvar.zip",
                "uncompressed_bytes": 15000414640
              },
              {
                "bytes": 4015346448,
                "central_directory_members": 92684,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr5.echtvar.zip",
                "uncompressed_bytes": 14002770510
              },
              {
                "bytes": 3834498842,
                "central_directory_members": 87806,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr6.echtvar.zip",
                "uncompressed_bytes": 13088773102
              },
              {
                "bytes": 3752875084,
                "central_directory_members": 81302,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr7.echtvar.zip",
                "uncompressed_bytes": 12699326179
              },
              {
                "bytes": 3445232398,
                "central_directory_members": 74798,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr8.echtvar.zip",
                "uncompressed_bytes": 12093741209
              },
              {
                "bytes": 2927934457,
                "central_directory_members": 64500,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr9.echtvar.zip",
                "uncompressed_bytes": 9883896838
              },
              {
                "bytes": 3075555846,
                "central_directory_members": 69378,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr10.echtvar.zip",
                "uncompressed_bytes": 10438645523
              },
              {
                "bytes": 3075222135,
                "central_directory_members": 68836,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr11.echtvar.zip",
                "uncompressed_bytes": 10620876194
              },
              {
                "bytes": 3019225963,
                "central_directory_members": 68294,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr12.echtvar.zip",
                "uncompressed_bytes": 10271390853
              },
              {
                "bytes": 2190906161,
                "central_directory_members": 49866,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr13.echtvar.zip",
                "uncompressed_bytes": 7487102196
              },
              {
                "bytes": 2082208118,
                "central_directory_members": 46072,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr14.echtvar.zip",
                "uncompressed_bytes": 7077438370
              },
              {
                "bytes": 1966885592,
                "central_directory_members": 42820,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr15.echtvar.zip",
                "uncompressed_bytes": 6648377385
              },
              {
                "bytes": 2159856477,
                "central_directory_members": 41736,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr16.echtvar.zip",
                "uncompressed_bytes": 7407307841
              },
              {
                "bytes": 1890442581,
                "central_directory_members": 41194,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr17.echtvar.zip",
                "uncompressed_bytes": 6342630113
              },
              {
                "bytes": 1730167927,
                "central_directory_members": 39568,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr18.echtvar.zip",
                "uncompressed_bytes": 5936324589
              },
              {
                "bytes": 1612041534,
                "central_directory_members": 29812,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr19.echtvar.zip",
                "uncompressed_bytes": 5165403019
              },
              {
                "bytes": 1376525969,
                "central_directory_members": 31980,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr20.echtvar.zip",
                "uncompressed_bytes": 4746336425
              },
              {
                "bytes": 902589299,
                "central_directory_members": 19514,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr21.echtvar.zip",
                "uncompressed_bytes": 3004608164
              },
              {
                "bytes": 945741018,
                "central_directory_members": 18430,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr22.echtvar.zip",
                "uncompressed_bytes": 3075770914
              }
            ],
            "fields_sha256": "fa1f21636fcd186308e3d070a65b0c087a8813f6763a78f8db70dc96c65f7643",
            "index_bytes": 1367830,
            "library": "/mnt/data6t/tian/immkv_test_hdd/gnomad/immkv",
            "manifests": [
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 1,
                "manifest_sha256": "e6502a10540c839e768f615df8081d70b4a71e61b2a0a88afa0821d687343e4d",
                "records": 20370890,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 2,
                "manifest_sha256": "0be51c08eec7941cd997de60a6fef8e3f72f9f38c3a4af9d889e395a7de99ea7",
                "records": 21892245,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 3,
                "manifest_sha256": "4fb8ff096c20282a966b468355676dc8e476b8400d7c200f4a42bd882d587e9c",
                "records": 17834355,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 4,
                "manifest_sha256": "85e3fdd517978eefa71480c09b88efbfe179188da72392025e1c995802ff3e37",
                "records": 17423639,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 5,
                "manifest_sha256": "88624fbba2bbc84b65c8ef80d78420d0b7a35928ed8ae9ac523b9f2e8f0f91a3",
                "records": 16290654,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 6,
                "manifest_sha256": "bdf9562a2110e95c010b344e601f16f8a543a51dca79fa2ced4031c53896fb91",
                "records": 15190442,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 7,
                "manifest_sha256": "77e5990e85572cbbbfce7871d79ee2fde5f7fc18cff2fe9a8f9c9900e2d824b5",
                "records": 14752321,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 8,
                "manifest_sha256": "56f8047c1da4c09bdb6e15f44acd99c448132ae6230e32ef22cfa20239567b5e",
                "records": 14073574,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 9,
                "manifest_sha256": "48b7a1d79c9a4a057bc7513eb930b850d7354f2abe380e3768aadaf87b45714f",
                "records": 11480614,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 10,
                "manifest_sha256": "db671e979c2a949f8c0433112f9053bcb02dcf4bc09a30e3407cdffa9378788a",
                "records": 12118750,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 11,
                "manifest_sha256": "0c8f3ed0470504385960e748aba1bc8854b806e3dd5044768739ef3d3a19a7f6",
                "records": 12345492,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 12,
                "manifest_sha256": "f0d71794035581a52f8a9c49187e7c5d77f79f5300aa76fb147efa0afa65ff88",
                "records": 11933435,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 13,
                "manifest_sha256": "563bb2b82c86652d08e862b28befc5be90d24897cf921aaad34621e79744f5d1",
                "records": 8693573,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 14,
                "manifest_sha256": "0471ec2e516c76f6b89410c89d8825aebfa36bfc9b15378c0045f9b46f7035c7",
                "records": 8220638,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 15,
                "manifest_sha256": "f9210d1c3b5ca3d8b36087dcc025a44f9e9934158ef77f78d85b5bb90c391667",
                "records": 7723667,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 16,
                "manifest_sha256": "3cf4786236d60d3c5402f496bd1cd8cfbe2ac3a2773e32f7cb1e680ef96c9a77",
                "records": 8618657,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 17,
                "manifest_sha256": "7f18009ed6de25e04fe6781dd2b181ad310ca3ecadeb5d65fd5f881be16eff8b",
                "records": 7373819,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 18,
                "manifest_sha256": "e7d6244b50c2e3b2752e25ab393dee0d5135133f087b10a0178d83a0bdb72ccb",
                "records": 6893751,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 19,
                "manifest_sha256": "df905df6fec7a78b7a65108551deb1e8c29bd3a3dca06edf68f5488bf8e7d360",
                "records": 5991870,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 20,
                "manifest_sha256": "5ffef94f04de22a04f846f3f06f2b33d1c6e828cc099efd15abeb6d6196a1db5",
                "records": 5516836,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 21,
                "manifest_sha256": "869bb863c33345bfcedba5de9265f1acac3c5033460ef83cb11d432d7143cdf4",
                "records": 3483000,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 22,
                "manifest_sha256": "872946e081799b83f30444581a23e5d7a4d3a4f46dba7e12a5185f8109e77895",
                "records": 3567085,
                "schema_hash": "1e91a99722c59d5e"
              }
            ],
            "query": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/query/hg001.decomp.vcf.gz",
            "query_bytes": 139085616,
            "query_sha256": "4fed97e319de887d251a7b1c603f2f842dfbc1247c1740eb7140172b67b22fe8"
          },
          "568": {
            "archives": [
              {
                "bytes": 7257511949,
                "central_directory_members": 124848,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr1.echtvar.zip",
                "uncompressed_bytes": 38780286499
              },
              {
                "bytes": 7624945470,
                "central_directory_members": 131688,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr2.echtvar.zip",
                "uncompressed_bytes": 40899531669
              },
              {
                "bytes": 6263830958,
                "central_directory_members": 106608,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr3.echtvar.zip",
                "uncompressed_bytes": 34926259821
              },
              {
                "bytes": 6074278192,
                "central_directory_members": 103188,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr4.echtvar.zip",
                "uncompressed_bytes": 30168583462
              },
              {
                "bytes": 5622847357,
                "central_directory_members": 97488,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr5.echtvar.zip",
                "uncompressed_bytes": 28696987035
              },
              {
                "bytes": 5333501460,
                "central_directory_members": 92358,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr6.echtvar.zip",
                "uncompressed_bytes": 26290904570
              },
              {
                "bytes": 5284983462,
                "central_directory_members": 85518,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr7.echtvar.zip",
                "uncompressed_bytes": 27981956434
              },
              {
                "bytes": 4867485483,
                "central_directory_members": 78678,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr8.echtvar.zip",
                "uncompressed_bytes": 26040552400
              },
              {
                "bytes": 4063797988,
                "central_directory_members": 67848,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr9.echtvar.zip",
                "uncompressed_bytes": 19890246125
              },
              {
                "bytes": 4299658931,
                "central_directory_members": 72978,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr10.echtvar.zip",
                "uncompressed_bytes": 21839763121
              },
              {
                "bytes": 4415331558,
                "central_directory_members": 72408,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr11.echtvar.zip",
                "uncompressed_bytes": 24849737455
              },
              {
                "bytes": 4324593804,
                "central_directory_members": 71838,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr12.echtvar.zip",
                "uncompressed_bytes": 24288159769
              },
              {
                "bytes": 2982695276,
                "central_directory_members": 52458,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr13.echtvar.zip",
                "uncompressed_bytes": 13200835719
              },
              {
                "bytes": 2978458193,
                "central_directory_members": 48468,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr14.echtvar.zip",
                "uncompressed_bytes": 16672855058
              },
              {
                "bytes": 2821168599,
                "central_directory_members": 45048,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr15.echtvar.zip",
                "uncompressed_bytes": 16007926228
              },
              {
                "bytes": 3147128737,
                "central_directory_members": 43908,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr16.echtvar.zip",
                "uncompressed_bytes": 18820464097
              },
              {
                "bytes": 2804825536,
                "central_directory_members": 43338,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr17.echtvar.zip",
                "uncompressed_bytes": 17740105448
              },
              {
                "bytes": 2434370983,
                "central_directory_members": 41628,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr18.echtvar.zip",
                "uncompressed_bytes": 12818068081
              },
              {
                "bytes": 2393428912,
                "central_directory_members": 31368,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr19.echtvar.zip",
                "uncompressed_bytes": 14771573347
              },
              {
                "bytes": 1939275218,
                "central_directory_members": 33648,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr20.echtvar.zip",
                "uncompressed_bytes": 10042845922
              },
              {
                "bytes": 1259700453,
                "central_directory_members": 20538,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr21.echtvar.zip",
                "uncompressed_bytes": 6372577608
              },
              {
                "bytes": 1362922300,
                "central_directory_members": 19398,
                "central_directory_valid": true,
                "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr22.echtvar.zip",
                "uncompressed_bytes": 7774267602
              }
            ],
            "fields_sha256": "28e7720f959b1cf583af444f1290a9bad857e3d733ea8cd6d0b75442b78ebf86",
            "index_bytes": 1367830,
            "library": "/mnt/data6t/tian/immkv_test_hdd/gnomad/immkv",
            "manifests": [
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 1,
                "manifest_sha256": "e6502a10540c839e768f615df8081d70b4a71e61b2a0a88afa0821d687343e4d",
                "records": 20370890,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 2,
                "manifest_sha256": "0be51c08eec7941cd997de60a6fef8e3f72f9f38c3a4af9d889e395a7de99ea7",
                "records": 21892245,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 3,
                "manifest_sha256": "4fb8ff096c20282a966b468355676dc8e476b8400d7c200f4a42bd882d587e9c",
                "records": 17834355,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 4,
                "manifest_sha256": "85e3fdd517978eefa71480c09b88efbfe179188da72392025e1c995802ff3e37",
                "records": 17423639,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 5,
                "manifest_sha256": "88624fbba2bbc84b65c8ef80d78420d0b7a35928ed8ae9ac523b9f2e8f0f91a3",
                "records": 16290654,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 6,
                "manifest_sha256": "bdf9562a2110e95c010b344e601f16f8a543a51dca79fa2ced4031c53896fb91",
                "records": 15190442,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 7,
                "manifest_sha256": "77e5990e85572cbbbfce7871d79ee2fde5f7fc18cff2fe9a8f9c9900e2d824b5",
                "records": 14752321,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 8,
                "manifest_sha256": "56f8047c1da4c09bdb6e15f44acd99c448132ae6230e32ef22cfa20239567b5e",
                "records": 14073574,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 9,
                "manifest_sha256": "48b7a1d79c9a4a057bc7513eb930b850d7354f2abe380e3768aadaf87b45714f",
                "records": 11480614,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 10,
                "manifest_sha256": "db671e979c2a949f8c0433112f9053bcb02dcf4bc09a30e3407cdffa9378788a",
                "records": 12118750,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 11,
                "manifest_sha256": "0c8f3ed0470504385960e748aba1bc8854b806e3dd5044768739ef3d3a19a7f6",
                "records": 12345492,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 12,
                "manifest_sha256": "f0d71794035581a52f8a9c49187e7c5d77f79f5300aa76fb147efa0afa65ff88",
                "records": 11933435,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 13,
                "manifest_sha256": "563bb2b82c86652d08e862b28befc5be90d24897cf921aaad34621e79744f5d1",
                "records": 8693573,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 14,
                "manifest_sha256": "0471ec2e516c76f6b89410c89d8825aebfa36bfc9b15378c0045f9b46f7035c7",
                "records": 8220638,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 15,
                "manifest_sha256": "f9210d1c3b5ca3d8b36087dcc025a44f9e9934158ef77f78d85b5bb90c391667",
                "records": 7723667,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 16,
                "manifest_sha256": "3cf4786236d60d3c5402f496bd1cd8cfbe2ac3a2773e32f7cb1e680ef96c9a77",
                "records": 8618657,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 17,
                "manifest_sha256": "7f18009ed6de25e04fe6781dd2b181ad310ca3ecadeb5d65fd5f881be16eff8b",
                "records": 7373819,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 18,
                "manifest_sha256": "e7d6244b50c2e3b2752e25ab393dee0d5135133f087b10a0178d83a0bdb72ccb",
                "records": 6893751,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 19,
                "manifest_sha256": "df905df6fec7a78b7a65108551deb1e8c29bd3a3dca06edf68f5488bf8e7d360",
                "records": 5991870,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 20,
                "manifest_sha256": "5ffef94f04de22a04f846f3f06f2b33d1c6e828cc099efd15abeb6d6196a1db5",
                "records": 5516836,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 21,
                "manifest_sha256": "869bb863c33345bfcedba5de9265f1acac3c5033460ef83cb11d432d7143cdf4",
                "records": 3483000,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 22,
                "manifest_sha256": "872946e081799b83f30444581a23e5d7a4d3a4f46dba7e12a5185f8109e77895",
                "records": 3567085,
                "schema_hash": "1e91a99722c59d5e"
              }
            ],
            "query": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/query/hg001.decomp.vcf.gz",
            "query_bytes": 139085616,
            "query_sha256": "4c61188a183910ae4b80fa038139718d0ccf3b3fe3911a7c14917109b3ceee06"
          }
        },
        "nvme": {
          "17": {
            "archives": [
              {
                "bytes": 244234654,
                "central_directory_members": 4163,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr1.echtvar.zip",
                "uncompressed_bytes": 578567882
              },
              {
                "bytes": 258657713,
                "central_directory_members": 4391,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr2.echtvar.zip",
                "uncompressed_bytes": 620538239
              },
              {
                "bytes": 210659980,
                "central_directory_members": 3555,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr3.echtvar.zip",
                "uncompressed_bytes": 505087392
              },
              {
                "bytes": 208828208,
                "central_directory_members": 3441,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr4.echtvar.zip",
                "uncompressed_bytes": 494714168
              },
              {
                "bytes": 192352387,
                "central_directory_members": 3251,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr5.echtvar.zip",
                "uncompressed_bytes": 461270809
              },
              {
                "bytes": 183934502,
                "central_directory_members": 3080,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr6.echtvar.zip",
                "uncompressed_bytes": 432604103
              },
              {
                "bytes": 178674759,
                "central_directory_members": 2852,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr7.echtvar.zip",
                "uncompressed_bytes": 420231928
              },
              {
                "bytes": 164725132,
                "central_directory_members": 2624,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr8.echtvar.zip",
                "uncompressed_bytes": 397702935
              },
              {
                "bytes": 138388260,
                "central_directory_members": 2263,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr9.echtvar.zip",
                "uncompressed_bytes": 325968828
              },
              {
                "bytes": 147108442,
                "central_directory_members": 2434,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr10.echtvar.zip",
                "uncompressed_bytes": 345447116
              },
              {
                "bytes": 146764897,
                "central_directory_members": 2415,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr11.echtvar.zip",
                "uncompressed_bytes": 350042016
              },
              {
                "bytes": 144193515,
                "central_directory_members": 2396,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr12.echtvar.zip",
                "uncompressed_bytes": 339777686
              },
              {
                "bytes": 104972306,
                "central_directory_members": 1750,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr13.echtvar.zip",
                "uncompressed_bytes": 247521444
              },
              {
                "bytes": 99273295,
                "central_directory_members": 1617,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr14.echtvar.zip",
                "uncompressed_bytes": 233719970
              },
              {
                "bytes": 93455566,
                "central_directory_members": 1503,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr15.echtvar.zip",
                "uncompressed_bytes": 219684326
              },
              {
                "bytes": 102348503,
                "central_directory_members": 1465,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr16.echtvar.zip",
                "uncompressed_bytes": 244290082
              },
              {
                "bytes": 89862907,
                "central_directory_members": 1446,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr17.echtvar.zip",
                "uncompressed_bytes": 210528323
              },
              {
                "bytes": 83163449,
                "central_directory_members": 1389,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr18.echtvar.zip",
                "uncompressed_bytes": 196324184
              },
              {
                "bytes": 75601696,
                "central_directory_members": 1047,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr19.echtvar.zip",
                "uncompressed_bytes": 172173473
              },
              {
                "bytes": 66291586,
                "central_directory_members": 1123,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr20.echtvar.zip",
                "uncompressed_bytes": 157162354
              },
              {
                "bytes": 43014863,
                "central_directory_members": 686,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr21.echtvar.zip",
                "uncompressed_bytes": 99628454
              },
              {
                "bytes": 44732754,
                "central_directory_members": 648,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr22.echtvar.zip",
                "uncompressed_bytes": 102376807
              }
            ],
            "fields_sha256": "db2c8bc29926b5f4265e9077443a19571a55c8f99d2d30906e0a126c6c0db030",
            "index_bytes": 1367854,
            "library": "/home/tian/immkv_test_nvme/test/17_fields/lib17",
            "manifests": [
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 1,
                "manifest_sha256": "b828cfa17a2df725f18e9907cee4cf0e9907c433d0ab5992ad6a25651f59852d",
                "records": 20370890,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 2,
                "manifest_sha256": "c99e8c2b512d651fefdba28fd4bc47ebc1ca96b816393fd8c77d08245a868505",
                "records": 21892245,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 3,
                "manifest_sha256": "88c35cffd4cc14950f915a4b39d1f466ccd0bcf1d49890a29beb47436dbf4f55",
                "records": 17834355,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 4,
                "manifest_sha256": "e310f6b01626703f969a903e630f029d97d3fdb089de0c8e3a5a1cb4b3f2a23c",
                "records": 17423639,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 5,
                "manifest_sha256": "96bfbf63bfb6181918baa760bdb8ff6ec675ef1459e2241ce0eb6eee1f8490f4",
                "records": 16290654,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 6,
                "manifest_sha256": "9a77beac2449ac85a8cf028e7769b9568fc7dc168c7f82134761233c8a6983ca",
                "records": 15190442,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 7,
                "manifest_sha256": "3e6c4b0b97e2e8acc707f9001fba3c96b7583aa761f285fb9187726b0737467f",
                "records": 14752321,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 8,
                "manifest_sha256": "a0de0210e73dbd886181a9c6e3b39e17854efabb53aa95d2f39114ab5de8ed64",
                "records": 14073574,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 9,
                "manifest_sha256": "90d4ded4deb27f1c1ffadebe47f5c65cf45455e657191676a4c23c557b61f8a9",
                "records": 11480614,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 10,
                "manifest_sha256": "12d27572091e8cb3fbb27386b5e777f73e330643f72d18ac8ea79618f465d547",
                "records": 12118750,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 11,
                "manifest_sha256": "c9343209324cedb76d8336d0b162e13adbde8ac07357906e101d816042827b3c",
                "records": 12345492,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 12,
                "manifest_sha256": "9c93500cf3fdbdde0991afd4a9a42602419f56be02fc0ede49e2268dec415de8",
                "records": 11933435,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 13,
                "manifest_sha256": "ee367aa983c8d31b9ef6485966954bc84c840c308f9ec10673955eea0ba8f052",
                "records": 8693573,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 14,
                "manifest_sha256": "2336388c4f0d6324d1d391bcbac4e81f0e32e95a343bb1a6872d1879f8446f2c",
                "records": 8220638,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 15,
                "manifest_sha256": "ca21551ad8d8d7b5e92ff40d830393ba374c9aaa8e18a45017433de9ea1c9b29",
                "records": 7723667,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 16,
                "manifest_sha256": "43209ec7835c225072e05d34144afba16d6bba7d83b747fccceb6f018d837cc7",
                "records": 8618657,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 17,
                "manifest_sha256": "57583441aab447fa2243d6939f679e50098ab45bb207b93820a679c6d36b6488",
                "records": 7373819,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 18,
                "manifest_sha256": "9a8bb1ae9a643ee3388533082705d2e26bcd09ab03cda90b64e8965d509cd369",
                "records": 6893751,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 19,
                "manifest_sha256": "c905eb1ca428e5fc53d8d67e0c5d4fdc0ee719800d8e089e3f3b794d6c80510a",
                "records": 5991870,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 20,
                "manifest_sha256": "fc4513940ce0caed90006e84a126eb2660c11f6e07a28c1cf70d352e89d529a3",
                "records": 5516836,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 21,
                "manifest_sha256": "5a064dcd7b265753c4e4e49d31e6aec772357bba315257e7d1b0a8e78d6c1772",
                "records": 3483000,
                "schema_hash": "99944e3316871f58"
              },
              {
                "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                "chromosome": 22,
                "manifest_sha256": "d5987ea72aaf88492de717ab4f935c2108bec30b957d63357268d66bafff178c",
                "records": 3567085,
                "schema_hash": "99944e3316871f58"
              }
            ],
            "query": "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz",
            "query_bytes": 139085620,
            "query_sha256": "014dc5c07f967a71cfdcc650c57148aeaac34ceb12e7d9f3a513b9720a88a9c6"
          },
          "540": {
            "archives": [
              {
                "bytes": 5123651180,
                "central_directory_members": 118700,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr1.echtvar.zip",
                "uncompressed_bytes": 17527475456
              },
              {
                "bytes": 5398570861,
                "central_directory_members": 125204,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr2.echtvar.zip",
                "uncompressed_bytes": 18815801585
              },
              {
                "bytes": 4396709608,
                "central_directory_members": 101356,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr3.echtvar.zip",
                "uncompressed_bytes": 15334578110
              },
              {
                "bytes": 4372074758,
                "central_directory_members": 98104,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr4.echtvar.zip",
                "uncompressed_bytes": 15000414640
              },
              {
                "bytes": 4015346448,
                "central_directory_members": 92684,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr5.echtvar.zip",
                "uncompressed_bytes": 14002770510
              },
              {
                "bytes": 3834498842,
                "central_directory_members": 87806,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr6.echtvar.zip",
                "uncompressed_bytes": 13088773102
              },
              {
                "bytes": 3752875084,
                "central_directory_members": 81302,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr7.echtvar.zip",
                "uncompressed_bytes": 12699326179
              },
              {
                "bytes": 3445232398,
                "central_directory_members": 74798,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr8.echtvar.zip",
                "uncompressed_bytes": 12093741209
              },
              {
                "bytes": 2927934457,
                "central_directory_members": 64500,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr9.echtvar.zip",
                "uncompressed_bytes": 9883896838
              },
              {
                "bytes": 3075555846,
                "central_directory_members": 69378,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr10.echtvar.zip",
                "uncompressed_bytes": 10438645523
              },
              {
                "bytes": 3075222135,
                "central_directory_members": 68836,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr11.echtvar.zip",
                "uncompressed_bytes": 10620876194
              },
              {
                "bytes": 3019225963,
                "central_directory_members": 68294,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr12.echtvar.zip",
                "uncompressed_bytes": 10271390853
              },
              {
                "bytes": 2190906161,
                "central_directory_members": 49866,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr13.echtvar.zip",
                "uncompressed_bytes": 7487102196
              },
              {
                "bytes": 2082208118,
                "central_directory_members": 46072,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr14.echtvar.zip",
                "uncompressed_bytes": 7077438370
              },
              {
                "bytes": 1966885592,
                "central_directory_members": 42820,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr15.echtvar.zip",
                "uncompressed_bytes": 6648377385
              },
              {
                "bytes": 2159856477,
                "central_directory_members": 41736,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr16.echtvar.zip",
                "uncompressed_bytes": 7407307841
              },
              {
                "bytes": 1890442581,
                "central_directory_members": 41194,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr17.echtvar.zip",
                "uncompressed_bytes": 6342630113
              },
              {
                "bytes": 1730167927,
                "central_directory_members": 39568,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr18.echtvar.zip",
                "uncompressed_bytes": 5936324589
              },
              {
                "bytes": 1612041534,
                "central_directory_members": 29812,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr19.echtvar.zip",
                "uncompressed_bytes": 5165403019
              },
              {
                "bytes": 1376525969,
                "central_directory_members": 31980,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr20.echtvar.zip",
                "uncompressed_bytes": 4746336425
              },
              {
                "bytes": 902589299,
                "central_directory_members": 19514,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr21.echtvar.zip",
                "uncompressed_bytes": 3004608164
              },
              {
                "bytes": 945741018,
                "central_directory_members": 18430,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr22.echtvar.zip",
                "uncompressed_bytes": 3075770914
              }
            ],
            "fields_sha256": "fa1f21636fcd186308e3d070a65b0c087a8813f6763a78f8db70dc96c65f7643",
            "index_bytes": 1367925,
            "library": "/home/tian/immkv_test_nvme/gnomad/immkv",
            "manifests": [
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 1,
                "manifest_sha256": "e6502a10540c839e768f615df8081d70b4a71e61b2a0a88afa0821d687343e4d",
                "records": 20370890,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 2,
                "manifest_sha256": "0be51c08eec7941cd997de60a6fef8e3f72f9f38c3a4af9d889e395a7de99ea7",
                "records": 21892245,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 3,
                "manifest_sha256": "4fb8ff096c20282a966b468355676dc8e476b8400d7c200f4a42bd882d587e9c",
                "records": 17834355,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 4,
                "manifest_sha256": "85e3fdd517978eefa71480c09b88efbfe179188da72392025e1c995802ff3e37",
                "records": 17423639,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 5,
                "manifest_sha256": "88624fbba2bbc84b65c8ef80d78420d0b7a35928ed8ae9ac523b9f2e8f0f91a3",
                "records": 16290654,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 6,
                "manifest_sha256": "bdf9562a2110e95c010b344e601f16f8a543a51dca79fa2ced4031c53896fb91",
                "records": 15190442,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 7,
                "manifest_sha256": "77e5990e85572cbbbfce7871d79ee2fde5f7fc18cff2fe9a8f9c9900e2d824b5",
                "records": 14752321,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 8,
                "manifest_sha256": "56f8047c1da4c09bdb6e15f44acd99c448132ae6230e32ef22cfa20239567b5e",
                "records": 14073574,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 9,
                "manifest_sha256": "48b7a1d79c9a4a057bc7513eb930b850d7354f2abe380e3768aadaf87b45714f",
                "records": 11480614,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 10,
                "manifest_sha256": "db671e979c2a949f8c0433112f9053bcb02dcf4bc09a30e3407cdffa9378788a",
                "records": 12118750,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 11,
                "manifest_sha256": "0c8f3ed0470504385960e748aba1bc8854b806e3dd5044768739ef3d3a19a7f6",
                "records": 12345492,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 12,
                "manifest_sha256": "f0d71794035581a52f8a9c49187e7c5d77f79f5300aa76fb147efa0afa65ff88",
                "records": 11933435,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 13,
                "manifest_sha256": "563bb2b82c86652d08e862b28befc5be90d24897cf921aaad34621e79744f5d1",
                "records": 8693573,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 14,
                "manifest_sha256": "0471ec2e516c76f6b89410c89d8825aebfa36bfc9b15378c0045f9b46f7035c7",
                "records": 8220638,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 15,
                "manifest_sha256": "f9210d1c3b5ca3d8b36087dcc025a44f9e9934158ef77f78d85b5bb90c391667",
                "records": 7723667,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 16,
                "manifest_sha256": "3cf4786236d60d3c5402f496bd1cd8cfbe2ac3a2773e32f7cb1e680ef96c9a77",
                "records": 8618657,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 17,
                "manifest_sha256": "7f18009ed6de25e04fe6781dd2b181ad310ca3ecadeb5d65fd5f881be16eff8b",
                "records": 7373819,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 18,
                "manifest_sha256": "e7d6244b50c2e3b2752e25ab393dee0d5135133f087b10a0178d83a0bdb72ccb",
                "records": 6893751,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 19,
                "manifest_sha256": "df905df6fec7a78b7a65108551deb1e8c29bd3a3dca06edf68f5488bf8e7d360",
                "records": 5991870,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 20,
                "manifest_sha256": "5ffef94f04de22a04f846f3f06f2b33d1c6e828cc099efd15abeb6d6196a1db5",
                "records": 5516836,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 21,
                "manifest_sha256": "869bb863c33345bfcedba5de9265f1acac3c5033460ef83cb11d432d7143cdf4",
                "records": 3483000,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 22,
                "manifest_sha256": "872946e081799b83f30444581a23e5d7a4d3a4f46dba7e12a5185f8109e77895",
                "records": 3567085,
                "schema_hash": "1e91a99722c59d5e"
              }
            ],
            "query": "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz",
            "query_bytes": 139085621,
            "query_sha256": "21b02c83b020ef6d97ad39068ce21deb28c496aaec3ba984365b4799bc2f29da"
          },
          "568": {
            "archives": [
              {
                "bytes": 7257511949,
                "central_directory_members": 124848,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr1.echtvar.zip",
                "uncompressed_bytes": 38780286499
              },
              {
                "bytes": 7624945470,
                "central_directory_members": 131688,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr2.echtvar.zip",
                "uncompressed_bytes": 40899531669
              },
              {
                "bytes": 6263830958,
                "central_directory_members": 106608,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr3.echtvar.zip",
                "uncompressed_bytes": 34926259821
              },
              {
                "bytes": 6074278192,
                "central_directory_members": 103188,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr4.echtvar.zip",
                "uncompressed_bytes": 30168583462
              },
              {
                "bytes": 5622847357,
                "central_directory_members": 97488,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr5.echtvar.zip",
                "uncompressed_bytes": 28696987035
              },
              {
                "bytes": 5333501460,
                "central_directory_members": 92358,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr6.echtvar.zip",
                "uncompressed_bytes": 26290904570
              },
              {
                "bytes": 5284983462,
                "central_directory_members": 85518,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr7.echtvar.zip",
                "uncompressed_bytes": 27981956434
              },
              {
                "bytes": 4867485483,
                "central_directory_members": 78678,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr8.echtvar.zip",
                "uncompressed_bytes": 26040552400
              },
              {
                "bytes": 4063797988,
                "central_directory_members": 67848,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr9.echtvar.zip",
                "uncompressed_bytes": 19890246125
              },
              {
                "bytes": 4299658931,
                "central_directory_members": 72978,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr10.echtvar.zip",
                "uncompressed_bytes": 21839763121
              },
              {
                "bytes": 4415331558,
                "central_directory_members": 72408,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr11.echtvar.zip",
                "uncompressed_bytes": 24849737455
              },
              {
                "bytes": 4324593804,
                "central_directory_members": 71838,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr12.echtvar.zip",
                "uncompressed_bytes": 24288159769
              },
              {
                "bytes": 2982695276,
                "central_directory_members": 52458,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr13.echtvar.zip",
                "uncompressed_bytes": 13200835719
              },
              {
                "bytes": 2978458193,
                "central_directory_members": 48468,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr14.echtvar.zip",
                "uncompressed_bytes": 16672855058
              },
              {
                "bytes": 2821168599,
                "central_directory_members": 45048,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr15.echtvar.zip",
                "uncompressed_bytes": 16007926228
              },
              {
                "bytes": 3147128737,
                "central_directory_members": 43908,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr16.echtvar.zip",
                "uncompressed_bytes": 18820464097
              },
              {
                "bytes": 2804825536,
                "central_directory_members": 43338,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr17.echtvar.zip",
                "uncompressed_bytes": 17740105448
              },
              {
                "bytes": 2434370983,
                "central_directory_members": 41628,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr18.echtvar.zip",
                "uncompressed_bytes": 12818068081
              },
              {
                "bytes": 2393428912,
                "central_directory_members": 31368,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr19.echtvar.zip",
                "uncompressed_bytes": 14771573347
              },
              {
                "bytes": 1939275218,
                "central_directory_members": 33648,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr20.echtvar.zip",
                "uncompressed_bytes": 10042845922
              },
              {
                "bytes": 1259700453,
                "central_directory_members": 20538,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr21.echtvar.zip",
                "uncompressed_bytes": 6372577608
              },
              {
                "bytes": 1362922300,
                "central_directory_members": 19398,
                "central_directory_valid": true,
                "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr22.echtvar.zip",
                "uncompressed_bytes": 7774267602
              }
            ],
            "fields_sha256": "28e7720f959b1cf583af444f1290a9bad857e3d733ea8cd6d0b75442b78ebf86",
            "index_bytes": 1367925,
            "library": "/home/tian/immkv_test_nvme/gnomad/immkv",
            "manifests": [
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 1,
                "manifest_sha256": "e6502a10540c839e768f615df8081d70b4a71e61b2a0a88afa0821d687343e4d",
                "records": 20370890,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 2,
                "manifest_sha256": "0be51c08eec7941cd997de60a6fef8e3f72f9f38c3a4af9d889e395a7de99ea7",
                "records": 21892245,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 3,
                "manifest_sha256": "4fb8ff096c20282a966b468355676dc8e476b8400d7c200f4a42bd882d587e9c",
                "records": 17834355,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 4,
                "manifest_sha256": "85e3fdd517978eefa71480c09b88efbfe179188da72392025e1c995802ff3e37",
                "records": 17423639,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 5,
                "manifest_sha256": "88624fbba2bbc84b65c8ef80d78420d0b7a35928ed8ae9ac523b9f2e8f0f91a3",
                "records": 16290654,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 6,
                "manifest_sha256": "bdf9562a2110e95c010b344e601f16f8a543a51dca79fa2ced4031c53896fb91",
                "records": 15190442,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 7,
                "manifest_sha256": "77e5990e85572cbbbfce7871d79ee2fde5f7fc18cff2fe9a8f9c9900e2d824b5",
                "records": 14752321,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 8,
                "manifest_sha256": "56f8047c1da4c09bdb6e15f44acd99c448132ae6230e32ef22cfa20239567b5e",
                "records": 14073574,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 9,
                "manifest_sha256": "48b7a1d79c9a4a057bc7513eb930b850d7354f2abe380e3768aadaf87b45714f",
                "records": 11480614,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 10,
                "manifest_sha256": "db671e979c2a949f8c0433112f9053bcb02dcf4bc09a30e3407cdffa9378788a",
                "records": 12118750,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 11,
                "manifest_sha256": "0c8f3ed0470504385960e748aba1bc8854b806e3dd5044768739ef3d3a19a7f6",
                "records": 12345492,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 12,
                "manifest_sha256": "f0d71794035581a52f8a9c49187e7c5d77f79f5300aa76fb147efa0afa65ff88",
                "records": 11933435,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 13,
                "manifest_sha256": "563bb2b82c86652d08e862b28befc5be90d24897cf921aaad34621e79744f5d1",
                "records": 8693573,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 14,
                "manifest_sha256": "0471ec2e516c76f6b89410c89d8825aebfa36bfc9b15378c0045f9b46f7035c7",
                "records": 8220638,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 15,
                "manifest_sha256": "f9210d1c3b5ca3d8b36087dcc025a44f9e9934158ef77f78d85b5bb90c391667",
                "records": 7723667,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 16,
                "manifest_sha256": "3cf4786236d60d3c5402f496bd1cd8cfbe2ac3a2773e32f7cb1e680ef96c9a77",
                "records": 8618657,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 17,
                "manifest_sha256": "7f18009ed6de25e04fe6781dd2b181ad310ca3ecadeb5d65fd5f881be16eff8b",
                "records": 7373819,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 18,
                "manifest_sha256": "e7d6244b50c2e3b2752e25ab393dee0d5135133f087b10a0178d83a0bdb72ccb",
                "records": 6893751,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 19,
                "manifest_sha256": "df905df6fec7a78b7a65108551deb1e8c29bd3a3dca06edf68f5488bf8e7d360",
                "records": 5991870,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 20,
                "manifest_sha256": "5ffef94f04de22a04f846f3f06f2b33d1c6e828cc099efd15abeb6d6196a1db5",
                "records": 5516836,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 21,
                "manifest_sha256": "869bb863c33345bfcedba5de9265f1acac3c5033460ef83cb11d432d7143cdf4",
                "records": 3483000,
                "schema_hash": "1e91a99722c59d5e"
              },
              {
                "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                "chromosome": 22,
                "manifest_sha256": "872946e081799b83f30444581a23e5d7a4d3a4f46dba7e12a5185f8109e77895",
                "records": 3567085,
                "schema_hash": "1e91a99722c59d5e"
              }
            ],
            "query": "/home/tian/immkv_test_nvme/test/568_fields/query/hg001.decomp.vcf.gz",
            "query_bytes": 139085621,
            "query_sha256": "ab247503db51111482d0b54b5cc76e3c0a94d5b2310746dd85f51d92b48f937b"
          }
        }
      },
      "query_data_sha256_from_readonly_preflight": "0250956380e1f6910a83f92a579810671e1fe4dfeb0b1d855394021a87148fd1"
    },
    "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/input-preflight.json"
  }
}
```

## Interpretation limits

- Every metric retains its source path; medians use exactly three complete valid rounds.
- OOM, timeout, failed, skipped_after_oom, started and unknown are independent outcomes; no missing result is estimated.
- GNU time maximum RSS is maximum child RSS in KiB, not simultaneous aggregate RSS.
- Cgroup memory.peak is aggregate charged memory in bytes, including page cache and memfd slices; observer fallback may precede final peak.
- CPU percent is GNU time percent (100 percent equals one CPU); four allowed CPUs need not imply only four htslib threads.
- Timing includes the complete Python driver tree. Echtvar includes bcftools slicing/BGZF into sealed memfd slices and four chromosome tasks; immkv includes reading, lookup, formatting and compressed output.
- Compressed output bytes are discarded during timing; memfd input slices are not stored on the tested medium. These runs do not measure full-output gzip disk-write performance.
- Fresh echtvar decompressed-row proof and immkv full-field verification are separate correctness phases; historical output references are not current timing samples.
- The release annotate source and recorded binary hashes identify the unchanged implementation and inputs; no current binary is executed or rehashed by this summary.
- Cache state is uncontrolled. Three paired raw values and their range describe observed spread only; no p-value or statistical certainty is inferred.
- Round engine order is immkv/echtvar, echtvar/immkv, immkv/echtvar (run_matrix.sh). Phase statistics from immkv.log are rounded internal timings and are not additive GNU wall components.

## Machine-readable report

```json
{
  "cells": [
    {
      "correctness": {
        "fresh_echtvar": {
          "constraints": {
            "mismatches": [],
            "observed": {
              "GOMAXPROCS": "4",
              "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-verify-echtvar-nvme-17-release-1794684671-90594.scope",
              "cpu.stat": "usage_usec 30658\nuser_usec 24724\nsystem_usec 5933\nnice_usec 0\ncore_sched.force_idle_usec 0",
              "cpus": [
                0,
                1,
                2,
                3
              ],
              "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
              "memory.current": "9805824",
              "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
              "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
              "memory.max": "17179869184",
              "memory.peak": "9805824",
              "memory.swap.max": "0",
              "timestamp_ns": 1794684671497735014
            },
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-17-release/constraints.json",
            "valid": true
          },
          "evidence": {
            "chromosomes": [
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr1.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl1.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an1.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58857612954211,
                  "time_ns": 1794684680803278795
                },
                "annotation_started": {
                  "monotonic_ns": 58849454555341,
                  "time_ns": 1794684672644879785
                },
                "chromosome": 1,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "1f44447b194ead91f936221f2bb3222981dd0926c80a98045774a787721fb477",
                    "rows": 308298
                  },
                  "reference": {
                    "data_sha256": "1f44447b194ead91f936221f2bb3222981dd0926c80a98045774a787721fb477",
                    "rows": 308298
                  }
                },
                "finished": {
                  "monotonic_ns": 58857622844052,
                  "time_ns": 1794684680813168727
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "1",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl1.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 10999772,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58849453672913,
                  "time_ns": 1794684672643996826
                },
                "slice_started": {
                  "monotonic_ns": 58848351059493,
                  "time_ns": 1794684671541384047
                },
                "started": {
                  "monotonic_ns": 58848350632034,
                  "time_ns": 1794684671540956308
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr2.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl2.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an2.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58858117841837,
                  "time_ns": 1794684681308165610
                },
                "annotation_started": {
                  "monotonic_ns": 58849656654371,
                  "time_ns": 1794684672846979056
                },
                "chromosome": 2,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "d9209bf52f85e8d53e6456f6090005d145b8c723c8eb438525ce62e5d699bb25",
                    "rows": 329821
                  },
                  "reference": {
                    "data_sha256": "d9209bf52f85e8d53e6456f6090005d145b8c723c8eb438525ce62e5d699bb25",
                    "rows": 329821
                  }
                },
                "finished": {
                  "monotonic_ns": 58858128074900,
                  "time_ns": 1794684681318398903
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "2",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl2.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 11750291,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58849655718373,
                  "time_ns": 1794684672846042366
                },
                "slice_started": {
                  "monotonic_ns": 58848351318366,
                  "time_ns": 1794684671541642921
                },
                "started": {
                  "monotonic_ns": 58848350909553,
                  "time_ns": 1794684671541234047
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr3.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl3.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an3.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58856594949838,
                  "time_ns": 1794684679785273882
                },
                "annotation_started": {
                  "monotonic_ns": 58849455373830,
                  "time_ns": 1794684672645698564
                },
                "chromosome": 3,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "a8d37d11be87aa98cbdc60f41451c38cd3f180fdc713f36a3a0403d4e0d03b06",
                    "rows": 275651
                  },
                  "reference": {
                    "data_sha256": "a8d37d11be87aa98cbdc60f41451c38cd3f180fdc713f36a3a0403d4e0d03b06",
                    "rows": 275651
                  }
                },
                "finished": {
                  "monotonic_ns": 58856605112139,
                  "time_ns": 1794684679795436182
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "3",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl3.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 9815087,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58849454905174,
                  "time_ns": 1794684672645229769
                },
                "slice_started": {
                  "monotonic_ns": 58848352222455,
                  "time_ns": 1794684671542547169
                },
                "started": {
                  "monotonic_ns": 58848351397745,
                  "time_ns": 1794684671541722199
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr4.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl4.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an4.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58856807554780,
                  "time_ns": 1794684679997878924
                },
                "annotation_started": {
                  "monotonic_ns": 58849455824722,
                  "time_ns": 1794684672646149276
                },
                "chromosome": 4,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "c70e6ff200808692bb15357b95933f255ef22c8784c172135bf869aeada9b192",
                    "rows": 296434
                  },
                  "reference": {
                    "data_sha256": "c70e6ff200808692bb15357b95933f255ef22c8784c172135bf869aeada9b192",
                    "rows": 296434
                  }
                },
                "finished": {
                  "monotonic_ns": 58856811310585,
                  "time_ns": 1794684680001634788
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "4",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl4.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 10519953,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58849455007115,
                  "time_ns": 1794684672645331399
                },
                "slice_started": {
                  "monotonic_ns": 58848352079818,
                  "time_ns": 1794684671542404583
                },
                "started": {
                  "monotonic_ns": 58848351722241,
                  "time_ns": 1794684671542046815
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr5.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl5.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an5.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58863646196183,
                  "time_ns": 1794684686836520157
                },
                "annotation_started": {
                  "monotonic_ns": 58857612927200,
                  "time_ns": 1794684680803251745
                },
                "chromosome": 5,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "f423059e24425dbb679170527a887a13a1677ccec48ce1e2d6aa4947b5961a3f",
                    "rows": 253843
                  },
                  "reference": {
                    "data_sha256": "f423059e24425dbb679170527a887a13a1677ccec48ce1e2d6aa4947b5961a3f",
                    "rows": 253843
                  }
                },
                "finished": {
                  "monotonic_ns": 58863650387251,
                  "time_ns": 1794684686840711715
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "5",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl5.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 9032175,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58857612383806,
                  "time_ns": 1794684680802707679
                },
                "slice_started": {
                  "monotonic_ns": 58856605442506,
                  "time_ns": 1794684679795767270
                },
                "started": {
                  "monotonic_ns": 58856605330486,
                  "time_ns": 1794684679795655201
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr6.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl6.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an6.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58863855780172,
                  "time_ns": 1794684687046104035
                },
                "annotation_started": {
                  "monotonic_ns": 58857819106140,
                  "time_ns": 1794684681009430394
                },
                "chromosome": 6,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "dd63acbe440fe990831807637a7741570970f50e1ccdead0cf2b18ebf92ecd67",
                    "rows": 264414
                  },
                  "reference": {
                    "data_sha256": "dd63acbe440fe990831807637a7741570970f50e1ccdead0cf2b18ebf92ecd67",
                    "rows": 264414
                  }
                },
                "finished": {
                  "monotonic_ns": 58863860302508,
                  "time_ns": 1794684687050626992
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "6",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl6.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 8584380,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58857817057904,
                  "time_ns": 1794684681007382068
                },
                "slice_started": {
                  "monotonic_ns": 58856812873403,
                  "time_ns": 1794684680003198158
                },
                "started": {
                  "monotonic_ns": 58856811506130,
                  "time_ns": 1794684680001830844
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr7.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl7.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an7.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58864273911914,
                  "time_ns": 1794684687464235877
                },
                "annotation_started": {
                  "monotonic_ns": 58858541707379,
                  "time_ns": 1794684681732031944
                },
                "chromosome": 7,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "ff2efa1a6c2b20a49037768efaf307afeebcbb3536f2a9f5670a7e744fc067ef",
                    "rows": 222733
                  },
                  "reference": {
                    "data_sha256": "ff2efa1a6c2b20a49037768efaf307afeebcbb3536f2a9f5670a7e744fc067ef",
                    "rows": 222733
                  }
                },
                "finished": {
                  "monotonic_ns": 58864280160644,
                  "time_ns": 1794684687470485079
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "7",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl7.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 7924912,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58858541415193,
                  "time_ns": 1794684681731739187
                },
                "slice_started": {
                  "monotonic_ns": 58857624769819,
                  "time_ns": 1794684680815094534
                },
                "started": {
                  "monotonic_ns": 58857624560428,
                  "time_ns": 1794684680814885163
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr8.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl8.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an8.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58864500111625,
                  "time_ns": 1794684687690435658
                },
                "annotation_started": {
                  "monotonic_ns": 58859048603306,
                  "time_ns": 1794684682238927480
                },
                "chromosome": 8,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "1adf70d93eeab1d8ee988b0556b21d89313f85123cacd243cb5981cf880c7b0d",
                    "rows": 213367
                  },
                  "reference": {
                    "data_sha256": "1adf70d93eeab1d8ee988b0556b21d89313f85123cacd243cb5981cf880c7b0d",
                    "rows": 213367
                  }
                },
                "finished": {
                  "monotonic_ns": 58864503390609,
                  "time_ns": 1794684687693715304
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "8",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl8.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 7556435,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58859048280002,
                  "time_ns": 1794684682238603936
                },
                "slice_started": {
                  "monotonic_ns": 58858128651146,
                  "time_ns": 1794684681318975881
                },
                "started": {
                  "monotonic_ns": 58858128519129,
                  "time_ns": 1794684681318843854
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr9.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl9.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an9.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58868281699497,
                  "time_ns": 1794684691472023460
                },
                "annotation_started": {
                  "monotonic_ns": 58864456238814,
                  "time_ns": 1794684687646563298
                },
                "chromosome": 9,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "325be35c2cc1be533573151313f57e5331ece22d9ebd65cb8ae769ae11cf9a06",
                    "rows": 168542
                  },
                  "reference": {
                    "data_sha256": "325be35c2cc1be533573151313f57e5331ece22d9ebd65cb8ae769ae11cf9a06",
                    "rows": 168542
                  }
                },
                "finished": {
                  "monotonic_ns": 58868287700765,
                  "time_ns": 1794684691478024919
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "9",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl9.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 5978980,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58864455774567,
                  "time_ns": 1794684687646098510
                },
                "slice_started": {
                  "monotonic_ns": 58863651813665,
                  "time_ns": 1794684686842138360
                },
                "started": {
                  "monotonic_ns": 58863650563059,
                  "time_ns": 1794684686840887784
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr10.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl10.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an10.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58869297187635,
                  "time_ns": 1794684692487511939
                },
                "annotation_started": {
                  "monotonic_ns": 58864767747697,
                  "time_ns": 1794684687958072392
                },
                "chromosome": 10,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "63e7d0f919c529977ebeed9ea8826093402f53a3e23ff4df3e9b3486a55b0213",
                    "rows": 198281
                  },
                  "reference": {
                    "data_sha256": "63e7d0f919c529977ebeed9ea8826093402f53a3e23ff4df3e9b3486a55b0213",
                    "rows": 198281
                  }
                },
                "finished": {
                  "monotonic_ns": 58869301659677,
                  "time_ns": 1794684692491984061
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "10",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl10.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 7017518,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58864767187872,
                  "time_ns": 1794684687957511715
                },
                "slice_started": {
                  "monotonic_ns": 58863860584315,
                  "time_ns": 1794684687050909069
                },
                "started": {
                  "monotonic_ns": 58863860476643,
                  "time_ns": 1794684687050801368
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr11.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl11.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an11.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58870323351221,
                  "time_ns": 1794684693513675244
                },
                "annotation_started": {
                  "monotonic_ns": 58865190669595,
                  "time_ns": 1794684688380993969
                },
                "chromosome": 11,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "58763a4b3db30bf77b4c2fcd43de752d36868ef90cd9ec7703c148d0b86ebdc9",
                    "rows": 200574
                  },
                  "reference": {
                    "data_sha256": "58763a4b3db30bf77b4c2fcd43de752d36868ef90cd9ec7703c148d0b86ebdc9",
                    "rows": 200574
                  }
                },
                "finished": {
                  "monotonic_ns": 58870326745150,
                  "time_ns": 1794684693517069604
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "11",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl11.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 7109132,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58865184499271,
                  "time_ns": 1794684688374823104
                },
                "slice_started": {
                  "monotonic_ns": 58864280440567,
                  "time_ns": 1794684687470765342
                },
                "started": {
                  "monotonic_ns": 58864280332395,
                  "time_ns": 1794684687470657140
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr12.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl12.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an12.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58870059090133,
                  "time_ns": 1794684693249414417
                },
                "annotation_started": {
                  "monotonic_ns": 58865217512049,
                  "time_ns": 1794684688407836583
                },
                "chromosome": 12,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "dc8d9d2acfb5bf8cd68a7e1c4db6dad160b920f88cc8a774d6ddcac27ba266a2",
                    "rows": 189356
                  },
                  "reference": {
                    "data_sha256": "dc8d9d2acfb5bf8cd68a7e1c4db6dad160b920f88cc8a774d6ddcac27ba266a2",
                    "rows": 189356
                  }
                },
                "finished": {
                  "monotonic_ns": 58870064540182,
                  "time_ns": 1794684693254864326
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "12",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl12.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 6749307,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58865214950735,
                  "time_ns": 1794684688405274879
                },
                "slice_started": {
                  "monotonic_ns": 58864503672446,
                  "time_ns": 1794684687693997170
                },
                "started": {
                  "monotonic_ns": 58864503567139,
                  "time_ns": 1794684687693891744
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr13.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl13.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an13.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58872739642855,
                  "time_ns": 1794684695929966968
                },
                "annotation_started": {
                  "monotonic_ns": 58869098099610,
                  "time_ns": 1794684692288424144
                },
                "chromosome": 13,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "25e1505b346e6f86f120186736e83d344c4fa729ac1372068e791264ea2a3b65",
                    "rows": 156939
                  },
                  "reference": {
                    "data_sha256": "25e1505b346e6f86f120186736e83d344c4fa729ac1372068e791264ea2a3b65",
                    "rows": 156939
                  }
                },
                "finished": {
                  "monotonic_ns": 58872742972373,
                  "time_ns": 1794684695933296807
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "13",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl13.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 5531589,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58869090410048,
                  "time_ns": 1794684692280733851
                },
                "slice_started": {
                  "monotonic_ns": 58868288026143,
                  "time_ns": 1794684691478350888
                },
                "started": {
                  "monotonic_ns": 58868287892433,
                  "time_ns": 1794684691478217028
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr14.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl14.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an14.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58872933498447,
                  "time_ns": 1794684696123822630
                },
                "annotation_started": {
                  "monotonic_ns": 58869904927284,
                  "time_ns": 1794684693095251989
                },
                "chromosome": 14,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "952a7a7b4ea03254e567232feaedf7108156810b47894e1071eb5093051a08ed",
                    "rows": 125631
                  },
                  "reference": {
                    "data_sha256": "952a7a7b4ea03254e567232feaedf7108156810b47894e1071eb5093051a08ed",
                    "rows": 125631
                  }
                },
                "finished": {
                  "monotonic_ns": 58872943508462,
                  "time_ns": 1794684696133832946
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "14",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl14.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 4486824,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58869904693057,
                  "time_ns": 1794684693095017251
                },
                "slice_started": {
                  "monotonic_ns": 58869302055907,
                  "time_ns": 1794684692492380622
                },
                "started": {
                  "monotonic_ns": 58869301847858,
                  "time_ns": 1794684692492172603
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr15.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl15.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an15.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58873596614131,
                  "time_ns": 1794684696786938154
                },
                "annotation_started": {
                  "monotonic_ns": 58870673549422,
                  "time_ns": 1794684693863873917
                },
                "chromosome": 15,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "0e7f24f259fa29545b6a2e3d0c266c18ef7d04675e17c88cdf85ee4b15a70ead",
                    "rows": 114117
                  },
                  "reference": {
                    "data_sha256": "0e7f24f259fa29545b6a2e3d0c266c18ef7d04675e17c88cdf85ee4b15a70ead",
                    "rows": 114117
                  }
                },
                "finished": {
                  "monotonic_ns": 58873602585683,
                  "time_ns": 1794684696792910057
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "15",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl15.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 4074508,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58870673273026,
                  "time_ns": 1794684693863597000
                },
                "slice_started": {
                  "monotonic_ns": 58870064815296,
                  "time_ns": 1794684693255140031
                },
                "started": {
                  "monotonic_ns": 58870064708837,
                  "time_ns": 1794684693255033502
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr16.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl16.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an16.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58873957666875,
                  "time_ns": 1794684697147990868
                },
                "annotation_started": {
                  "monotonic_ns": 58870939338895,
                  "time_ns": 1794684694129663159
                },
                "chromosome": 16,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "68c1a3c0ffd48e36e7583ba722f4af1f2d400ded3454af3c1cdfba59ace359b0",
                    "rows": 118437
                  },
                  "reference": {
                    "data_sha256": "68c1a3c0ffd48e36e7583ba722f4af1f2d400ded3454af3c1cdfba59ace359b0",
                    "rows": 118437
                  }
                },
                "finished": {
                  "monotonic_ns": 58873960372448,
                  "time_ns": 1794684697150696641
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "16",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl16.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 4212957,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58870937399563,
                  "time_ns": 1794684694127723476
                },
                "slice_started": {
                  "monotonic_ns": 58870327021085,
                  "time_ns": 1794684693517345820
                },
                "started": {
                  "monotonic_ns": 58870326909476,
                  "time_ns": 1794684693517234211
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr17.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl17.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an17.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58875768363532,
                  "time_ns": 1794684698958687165
                },
                "annotation_started": {
                  "monotonic_ns": 58873251704897,
                  "time_ns": 1794684696442029422
                },
                "chromosome": 17,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "851a7e861003ed1ba8e021ad28f7a87dfc99f0b38159b7d0cfecdaaf0725c020",
                    "rows": 102461
                  },
                  "reference": {
                    "data_sha256": "851a7e861003ed1ba8e021ad28f7a87dfc99f0b38159b7d0cfecdaaf0725c020",
                    "rows": 102461
                  }
                },
                "finished": {
                  "monotonic_ns": 58875771319833,
                  "time_ns": 1794684698961644188
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "17",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl17.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 3711658,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58873251457816,
                  "time_ns": 1794684696441781659
                },
                "slice_started": {
                  "monotonic_ns": 58872747802594,
                  "time_ns": 1794684695938127139
                },
                "started": {
                  "monotonic_ns": 58872743177085,
                  "time_ns": 1794684695933501720
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr18.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl18.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an18.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58876472392459,
                  "time_ns": 1794684699662716632
                },
                "annotation_started": {
                  "monotonic_ns": 58873554142367,
                  "time_ns": 1794684696744466841
                },
                "chromosome": 18,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "07f618d6065c8f4998b9e03fbd3fdb1436ae81ac475a6205a33fe214cd286a6e",
                    "rows": 119403
                  },
                  "reference": {
                    "data_sha256": "07f618d6065c8f4998b9e03fbd3fdb1436ae81ac475a6205a33fe214cd286a6e",
                    "rows": 119403
                  }
                },
                "finished": {
                  "monotonic_ns": 58876475817676,
                  "time_ns": 1794684699666142210
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "18",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl18.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 4239925,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58873550024826,
                  "time_ns": 1794684696740348779
                },
                "slice_started": {
                  "monotonic_ns": 58872943795298,
                  "time_ns": 1794684696134120073
                },
                "started": {
                  "monotonic_ns": 58872943689170,
                  "time_ns": 1794684696134013894
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr19.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl19.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an19.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58876222171564,
                  "time_ns": 1794684699412495728
                },
                "annotation_started": {
                  "monotonic_ns": 58874007738343,
                  "time_ns": 1794684697198062817
                },
                "chromosome": 19,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "f088ed123c5b73dd87beea4691ae6559a17a91cfc8ef6deb94ebe810db8c1710",
                    "rows": 86442
                  },
                  "reference": {
                    "data_sha256": "f088ed123c5b73dd87beea4691ae6559a17a91cfc8ef6deb94ebe810db8c1710",
                    "rows": 86442
                  }
                },
                "finished": {
                  "monotonic_ns": 58876226436169,
                  "time_ns": 1794684699416760513
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "19",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl19.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 3140431,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58874005803439,
                  "time_ns": 1794684697196127372
                },
                "slice_started": {
                  "monotonic_ns": 58873602867329,
                  "time_ns": 1794684696793192094
                },
                "started": {
                  "monotonic_ns": 58873602764277,
                  "time_ns": 1794684696793089002
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr20.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl20.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an20.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58876482708716,
                  "time_ns": 1794684699673032980
                },
                "annotation_started": {
                  "monotonic_ns": 58874364530029,
                  "time_ns": 1794684697554854503
                },
                "chromosome": 20,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "45f3e17c5f4670ab27b47b3b496fe2a37f79068afdbf5366b3c1db91ce54a640",
                    "rows": 83452
                  },
                  "reference": {
                    "data_sha256": "45f3e17c5f4670ab27b47b3b496fe2a37f79068afdbf5366b3c1db91ce54a640",
                    "rows": 83452
                  }
                },
                "finished": {
                  "monotonic_ns": 58876488670010,
                  "time_ns": 1794684699678994544
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "20",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl20.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 2951322,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58874364289409,
                  "time_ns": 1794684697554613533
                },
                "slice_started": {
                  "monotonic_ns": 58873960657901,
                  "time_ns": 1794684697150982616
                },
                "started": {
                  "monotonic_ns": 58873960552264,
                  "time_ns": 1794684697150876988
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr21.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl21.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an21.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58877285346914,
                  "time_ns": 1794684700475670747
                },
                "annotation_started": {
                  "monotonic_ns": 58876175609190,
                  "time_ns": 1794684699365933815
                },
                "chromosome": 21,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "1d4a87c093eb5bf1dc897ad5e6c5f94538f316b868d37462fb6514bb5fee9de2",
                    "rows": 55978
                  },
                  "reference": {
                    "data_sha256": "1d4a87c093eb5bf1dc897ad5e6c5f94538f316b868d37462fb6514bb5fee9de2",
                    "rows": 55978
                  }
                },
                "finished": {
                  "monotonic_ns": 58877289337978,
                  "time_ns": 1794684700479662412
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "21",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl21.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 1982476,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58876175373099,
                  "time_ns": 1794684699365697133
                },
                "slice_started": {
                  "monotonic_ns": 58875773236954,
                  "time_ns": 1794684698963561399
                },
                "started": {
                  "monotonic_ns": 58875773089529,
                  "time_ns": 1794684698963414144
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/17_fields/archives/chr22.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl22.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/an22.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 58877434140083,
                  "time_ns": 1794684700624464176
                },
                "annotation_started": {
                  "monotonic_ns": 58876532647586,
                  "time_ns": 1794684699722971980
                },
                "chromosome": 22,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "c4e02794b0aa94ca0a0eae05006399e26e108da4d8d7ce5aae62e855faf31ea7",
                    "rows": 49540
                  },
                  "reference": {
                    "data_sha256": "c4e02794b0aa94ca0a0eae05006399e26e108da4d8d7ce5aae62e855faf31ea7",
                    "rows": 49540
                  }
                },
                "finished": {
                  "monotonic_ns": 58877438460272,
                  "time_ns": 1794684700628784816
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "22",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release/sl22.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 1765325,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 58876532314975,
                  "time_ns": 1794684699722638888
                },
                "slice_started": {
                  "monotonic_ns": 58876227568394,
                  "time_ns": 1794684699417893088
                },
                "started": {
                  "monotonic_ns": 58876227418223,
                  "time_ns": 1794684699417742778
                },
                "status": "ok"
              }
            ],
            "compressed_output": true,
            "engine": "echtvar",
            "exit_code": 0,
            "expected_rows": 3933714,
            "fields": 17,
            "finished": {
              "monotonic_ns": 58877438991183,
              "time_ns": 1794684700629315587
            },
            "fresh_rows": 3933714,
            "media": "nvme",
            "scratch": "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-17-release",
            "started": {
              "monotonic_ns": 58848350363843,
              "time_ns": 1794684671540683118
            },
            "status": "ok",
            "worker_note": "Four chromosome subprocess tasks on CPUs 0-3; each echtvar process may create additional htslib threads.",
            "workers": 4
          },
          "issues": [],
          "note": "Correctness phase only; excluded from timing medians.",
          "passed": true,
          "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-17-release/cell.json",
          "status": {
            "directory_exists": true,
            "record": {
              "exit_code": 0,
              "finished_ns": 1794684700664759285,
              "memory_peak_bytes": 3054837760,
              "oom_kill_detected": false,
              "status": "ok"
            },
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-17-release/status.json",
            "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-17-release/started",
            "state": "ok"
          },
          "verifier_stdout": {}
        },
        "immkv": {
          "constraints": {
            "mismatches": [],
            "observed": {
              "GOMAXPROCS": "4",
              "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-verify-nvme-17-release-1794684617-90457.scope",
              "cpu.stat": "usage_usec 31795\nuser_usec 20865\nsystem_usec 10929\nnice_usec 0\ncore_sched.force_idle_usec 0",
              "cpus": [
                0,
                1,
                2,
                3
              ],
              "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
              "memory.current": "10055680",
              "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
              "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
              "memory.max": "17179869184",
              "memory.peak": "10055680",
              "memory.swap.max": "0",
              "timestamp_ns": 1794684617562083162
            },
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/constraints.json",
            "valid": true
          },
          "evidence": {
            "annotate_statistics": {
              "hits": 3875208,
              "rows": 3933714
            },
            "commands": {
              "annotate": [
                "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/annotate",
                "-in",
                "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz",
                "-lib",
                "/home/tian/immkv_test_nvme/test/17_fields/lib17",
                "-fields",
                "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_17_fields/fields17.tsv",
                "-threads",
                "4",
                "-batch",
                "4096",
                "-out",
                "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-17-release/output.vcf.gz"
              ],
              "verifier": [
                "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/verifyvalues",
                "-ev",
                "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-17-release/echtvar.fifo",
                "-iv",
                "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-17-release/immkv.fifo",
                "-fields",
                "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_17_fields/fields17.tsv",
                "-tol",
                "1.1e-6",
                "-max-report",
                "20"
              ]
            },
            "comparison": {
              "echtvar_i32_overflow_losses": 0,
              "field_comparisons": 66873138,
              "hit_disagreements": 0,
              "misses_agreed": 1007196,
              "rows": 3933714,
              "value_mismatches": 0
            },
            "elapsed_seconds_correctness_only": 40.02600424699631,
            "exit_status": {
              "annotate": 0,
              "verifier": 0
            },
            "expected_hits": 3875208,
            "expected_rows": 3933714,
            "fields": 17,
            "fields_tsv": "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_17_fields/fields17.tsv",
            "float_tolerance": 1.1e-06,
            "float_tolerance_basis": "echtvar v0.2.4 quantization truncation plus six-decimal print truncation and f32 slack",
            "library": "/home/tian/immkv_test_nvme/test/17_fields/lib17",
            "marker_collision_accounting": "Shared verifier includes true float -1 marker collisions in misses agreed.",
            "media": "nvme",
            "query": "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz",
            "reference_kind": "historical correctness reference; not a current performance sample",
            "reference_media": "nvme",
            "reference_release": "echtvar v0.2.4",
            "sha256": {
              "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_17_fields/fields17.tsv": "db2c8bc29926b5f4265e9077443a19571a55c8f99d2d30906e0a126c6c0db030",
              "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify_output.py": "2466d4c8fb67067d5180a83f44344057ebd420726a174019956225a1f98cef78",
              "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/annotate": "4179f6f9bf774615ce68e1794dd3c46f2c5e4b2316011b117191584f5ec16391",
              "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/verifyvalues": "8e307e09177310a07ec9615734493f9e423a779f3b6f3a2a5dd82f4364343400"
            },
            "status": "passed",
            "streams": {
              "echtvar": {
                "data_sha256": "ec008db079adf59cb5f84928a5096821694fddef1ded9757fff1043725fe6720",
                "files": [
                  {
                    "data_bytes": 269824780,
                    "data_sha256": "1f44447b194ead91f936221f2bb3222981dd0926c80a98045774a787721fb477",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an1.vcf.gz",
                    "rows": 308298
                  },
                  {
                    "data_bytes": 289150520,
                    "data_sha256": "d9209bf52f85e8d53e6456f6090005d145b8c723c8eb438525ce62e5d699bb25",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an2.vcf.gz",
                    "rows": 329821
                  },
                  {
                    "data_bytes": 241537645,
                    "data_sha256": "a8d37d11be87aa98cbdc60f41451c38cd3f180fdc713f36a3a0403d4e0d03b06",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an3.vcf.gz",
                    "rows": 275651
                  },
                  {
                    "data_bytes": 259661090,
                    "data_sha256": "c70e6ff200808692bb15357b95933f255ef22c8784c172135bf869aeada9b192",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an4.vcf.gz",
                    "rows": 296434
                  },
                  {
                    "data_bytes": 222394381,
                    "data_sha256": "f423059e24425dbb679170527a887a13a1677ccec48ce1e2d6aa4947b5961a3f",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an5.vcf.gz",
                    "rows": 253843
                  },
                  {
                    "data_bytes": 217988880,
                    "data_sha256": "dd63acbe440fe990831807637a7741570970f50e1ccdead0cf2b18ebf92ecd67",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an6.vcf.gz",
                    "rows": 264414
                  },
                  {
                    "data_bytes": 194980426,
                    "data_sha256": "ff2efa1a6c2b20a49037768efaf307afeebcbb3536f2a9f5670a7e744fc067ef",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an7.vcf.gz",
                    "rows": 222733
                  },
                  {
                    "data_bytes": 187116360,
                    "data_sha256": "1adf70d93eeab1d8ee988b0556b21d89313f85123cacd243cb5981cf880c7b0d",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an8.vcf.gz",
                    "rows": 213367
                  },
                  {
                    "data_bytes": 147736697,
                    "data_sha256": "325be35c2cc1be533573151313f57e5331ece22d9ebd65cb8ae769ae11cf9a06",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an9.vcf.gz",
                    "rows": 168542
                  },
                  {
                    "data_bytes": 173832355,
                    "data_sha256": "63e7d0f919c529977ebeed9ea8826093402f53a3e23ff4df3e9b3486a55b0213",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an10.vcf.gz",
                    "rows": 198281
                  },
                  {
                    "data_bytes": 175692515,
                    "data_sha256": "58763a4b3db30bf77b4c2fcd43de752d36868ef90cd9ec7703c148d0b86ebdc9",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an11.vcf.gz",
                    "rows": 200574
                  },
                  {
                    "data_bytes": 165790011,
                    "data_sha256": "dc8d9d2acfb5bf8cd68a7e1c4db6dad160b920f88cc8a774d6ddcac27ba266a2",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an12.vcf.gz",
                    "rows": 189356
                  },
                  {
                    "data_bytes": 137625829,
                    "data_sha256": "25e1505b346e6f86f120186736e83d344c4fa729ac1372068e791264ea2a3b65",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an13.vcf.gz",
                    "rows": 156939
                  },
                  {
                    "data_bytes": 110046147,
                    "data_sha256": "952a7a7b4ea03254e567232feaedf7108156810b47894e1071eb5093051a08ed",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an14.vcf.gz",
                    "rows": 125631
                  },
                  {
                    "data_bytes": 100094646,
                    "data_sha256": "0e7f24f259fa29545b6a2e3d0c266c18ef7d04675e17c88cdf85ee4b15a70ead",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an15.vcf.gz",
                    "rows": 114117
                  },
                  {
                    "data_bytes": 103879536,
                    "data_sha256": "68c1a3c0ffd48e36e7583ba722f4af1f2d400ded3454af3c1cdfba59ace359b0",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an16.vcf.gz",
                    "rows": 118437
                  },
                  {
                    "data_bytes": 89523814,
                    "data_sha256": "851a7e861003ed1ba8e021ad28f7a87dfc99f0b38159b7d0cfecdaaf0725c020",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an17.vcf.gz",
                    "rows": 102461
                  },
                  {
                    "data_bytes": 104836551,
                    "data_sha256": "07f618d6065c8f4998b9e03fbd3fdb1436ae81ac475a6205a33fe214cd286a6e",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an18.vcf.gz",
                    "rows": 119403
                  },
                  {
                    "data_bytes": 75169025,
                    "data_sha256": "f088ed123c5b73dd87beea4691ae6559a17a91cfc8ef6deb94ebe810db8c1710",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an19.vcf.gz",
                    "rows": 86442
                  },
                  {
                    "data_bytes": 73159040,
                    "data_sha256": "45f3e17c5f4670ab27b47b3b496fe2a37f79068afdbf5366b3c1db91ce54a640",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an20.vcf.gz",
                    "rows": 83452
                  },
                  {
                    "data_bytes": 49088311,
                    "data_sha256": "1d4a87c093eb5bf1dc897ad5e6c5f94538f316b868d37462fb6514bb5fee9de2",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an21.vcf.gz",
                    "rows": 55978
                  },
                  {
                    "data_bytes": 43278460,
                    "data_sha256": "c4e02794b0aa94ca0a0eae05006399e26e108da4d8d7ce5aae62e855faf31ea7",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/17_fields/echtvar.pass2.wk/an22.vcf.gz",
                    "rows": 49540
                  }
                ],
                "rows": 3933714
              },
              "immkv": {
                "data_sha256": "8e1c55459edc13e1b24236ef9d65b54f3215dc11dd55db97efa5979f5247af63",
                "files": [
                  {
                    "data_bytes": 3418342825,
                    "data_sha256": "8e1c55459edc13e1b24236ef9d65b54f3215dc11dd55db97efa5979f5247af63",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-17-release/output.vcf.gz",
                    "rows": 3933714
                  }
                ],
                "rows": 3933714,
                "selected_float_true_minus_one_candidates": 0
              }
            },
            "timing_sample": false
          },
          "issues": [],
          "note": "Correctness phase only; excluded from timing medians.",
          "passed": true,
          "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verification.json",
          "status": {
            "directory_exists": true,
            "record": {
              "exit_code": 0,
              "finished_ns": 1794684657662554011,
              "memory_peak_bytes": 970469376,
              "oom_kill_detected": false,
              "status": "ok"
            },
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/status.json",
            "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/started",
            "state": "ok"
          },
          "verifier_stdout": {
            "echtvar_i32_overflow_losses": {
              "field": "echtvar_i32_overflow_losses",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout",
              "value": 0
            },
            "field_comparisons": {
              "field": "field_comparisons",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout",
              "value": 66873138
            },
            "float_comparisons": {
              "field": "float_comparisons",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout",
              "value": 30989070
            },
            "float_tolerance": {
              "field": "float_tolerance",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout",
              "value": 1.1e-06
            },
            "hit_disagreements": {
              "field": "hit_disagreements",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout",
              "value": 0
            },
            "max_float_absolute_difference": {
              "field": "max_float_absolute_difference",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout",
              "value": 1e-06
            },
            "misses_agreed": {
              "field": "misses_agreed",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout",
              "value": 1007196
            },
            "rows": {
              "field": "rows",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout",
              "value": 3933714
            },
            "value_mismatches": {
              "field": "value_mismatches",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout",
              "value": 0
            },
            "verify_values_ok": {
              "field": "VERIFY_VALUES OK",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-17-release/verifier.stdout",
              "value": true
            }
          }
        }
      },
      "engines": {
        "echtvar": {
          "rounds": [
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": {
                  "GOMAXPROCS": "4",
                  "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-timing-nvme-17-echtvar-r1-release-1794684732-90895.scope",
                  "cpu.stat": "usage_usec 27775\nuser_usec 17239\nsystem_usec 10535\nnice_usec 0\ncore_sched.force_idle_usec 0",
                  "cpus": [
                    0,
                    1,
                    2,
                    3
                  ],
                  "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
                  "memory.current": "9461760",
                  "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.max": "17179869184",
                  "memory.peak": "9461760",
                  "memory.swap.max": "0",
                  "timestamp_ns": 1794684732485746278
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/constraints.json",
                "valid": true
              },
              "count_proof": {
                "all_22_chromosomes_ok": true,
                "annotation_log_statistics": [
                  {
                    "chromosome": 1,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr1.anno.log",
                      "value": 308298
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr1.anno.log",
                      "value": 308297
                    }
                  },
                  {
                    "chromosome": 2,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr2.anno.log",
                      "value": 329821
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr2.anno.log",
                      "value": 329821
                    }
                  },
                  {
                    "chromosome": 3,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr3.anno.log",
                      "value": 275651
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr3.anno.log",
                      "value": 275647
                    }
                  },
                  {
                    "chromosome": 4,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr4.anno.log",
                      "value": 296434
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr4.anno.log",
                      "value": 296434
                    }
                  },
                  {
                    "chromosome": 5,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr5.anno.log",
                      "value": 253843
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr5.anno.log",
                      "value": 253841
                    }
                  },
                  {
                    "chromosome": 6,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr6.anno.log",
                      "value": 264414
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr6.anno.log",
                      "value": 264084
                    }
                  },
                  {
                    "chromosome": 7,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr7.anno.log",
                      "value": 222733
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr7.anno.log",
                      "value": 222732
                    }
                  },
                  {
                    "chromosome": 8,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr8.anno.log",
                      "value": 213367
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr8.anno.log",
                      "value": 213367
                    }
                  },
                  {
                    "chromosome": 9,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr9.anno.log",
                      "value": 168542
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr9.anno.log",
                      "value": 168542
                    }
                  },
                  {
                    "chromosome": 10,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr10.anno.log",
                      "value": 198281
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr10.anno.log",
                      "value": 198281
                    }
                  },
                  {
                    "chromosome": 11,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr11.anno.log",
                      "value": 200574
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr11.anno.log",
                      "value": 200574
                    }
                  },
                  {
                    "chromosome": 12,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr12.anno.log",
                      "value": 189356
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr12.anno.log",
                      "value": 189356
                    }
                  },
                  {
                    "chromosome": 13,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr13.anno.log",
                      "value": 156939
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr13.anno.log",
                      "value": 156938
                    }
                  },
                  {
                    "chromosome": 14,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr14.anno.log",
                      "value": 125631
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr14.anno.log",
                      "value": 125631
                    }
                  },
                  {
                    "chromosome": 15,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr15.anno.log",
                      "value": 114117
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr15.anno.log",
                      "value": 114117
                    }
                  },
                  {
                    "chromosome": 16,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr16.anno.log",
                      "value": 118437
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr16.anno.log",
                      "value": 118437
                    }
                  },
                  {
                    "chromosome": 17,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr17.anno.log",
                      "value": 102461
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr17.anno.log",
                      "value": 102461
                    }
                  },
                  {
                    "chromosome": 18,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr18.anno.log",
                      "value": 119403
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr18.anno.log",
                      "value": 119401
                    }
                  },
                  {
                    "chromosome": 19,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr19.anno.log",
                      "value": 86442
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr19.anno.log",
                      "value": 86442
                    }
                  },
                  {
                    "chromosome": 20,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr20.anno.log",
                      "value": 83452
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr20.anno.log",
                      "value": 83452
                    }
                  },
                  {
                    "chromosome": 21,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr21.anno.log",
                      "value": 55978
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr21.anno.log",
                      "value": 55978
                    }
                  },
                  {
                    "chromosome": 22,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr22.anno.log",
                      "value": 49540
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr22.anno.log",
                      "value": 49536
                    }
                  }
                ],
                "chromosome_statuses": [
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr1.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl1.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58916851625847,
                      "time_ns": 1794684740041950051
                    },
                    "annotation_started": {
                      "monotonic_ns": 58910539897599,
                      "time_ns": 1794684733730222274
                    },
                    "chromosome": 1,
                    "finished": {
                      "monotonic_ns": 58916853086064,
                      "time_ns": 1794684740043410529
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "1",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl1.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 10999774,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58910539712033,
                      "time_ns": 1794684733730035936
                    },
                    "slice_started": {
                      "monotonic_ns": 58909337384886,
                      "time_ns": 1794684732527709591
                    },
                    "started": {
                      "monotonic_ns": 58909337131253,
                      "time_ns": 1794684732527455947
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr2.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl2.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58917354560574,
                      "time_ns": 1794684740544884698
                    },
                    "annotation_started": {
                      "monotonic_ns": 58910540790146,
                      "time_ns": 1794684733731114801
                    },
                    "chromosome": 2,
                    "finished": {
                      "monotonic_ns": 58917356144843,
                      "time_ns": 1794684740546469438
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "2",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl2.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 11750293,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58910540632762,
                      "time_ns": 1794684733730957397
                    },
                    "slice_started": {
                      "monotonic_ns": 58909337792077,
                      "time_ns": 1794684732528116772
                    },
                    "started": {
                      "monotonic_ns": 58909337571204,
                      "time_ns": 1794684732527895659
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr3.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl3.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58915858759678,
                      "time_ns": 1794684739049083371
                    },
                    "annotation_started": {
                      "monotonic_ns": 58910340581281,
                      "time_ns": 1794684733530905856
                    },
                    "chromosome": 3,
                    "finished": {
                      "monotonic_ns": 58915860365167,
                      "time_ns": 1794684739050689892
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "3",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl3.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 9815089,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58910340392389,
                      "time_ns": 1794684733530716262
                    },
                    "slice_started": {
                      "monotonic_ns": 58909338139706,
                      "time_ns": 1794684732528464441
                    },
                    "started": {
                      "monotonic_ns": 58909337821332,
                      "time_ns": 1794684732528144634
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr4.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl4.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58916149166613,
                      "time_ns": 1794684739339490797
                    },
                    "annotation_started": {
                      "monotonic_ns": 58910441591025,
                      "time_ns": 1794684733631915770
                    },
                    "chromosome": 4,
                    "finished": {
                      "monotonic_ns": 58916150244126,
                      "time_ns": 1794684739340568931
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "4",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl4.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 10519955,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58910441444271,
                      "time_ns": 1794684733631768334
                    },
                    "slice_started": {
                      "monotonic_ns": 58909338307450,
                      "time_ns": 1794684732528632214
                    },
                    "started": {
                      "monotonic_ns": 58909338056421,
                      "time_ns": 1794684732528381105
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr5.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl5.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58921978308085,
                      "time_ns": 1794684745168632028
                    },
                    "annotation_started": {
                      "monotonic_ns": 58916965070405,
                      "time_ns": 1794684740155395119
                    },
                    "chromosome": 5,
                    "finished": {
                      "monotonic_ns": 58921979501244,
                      "time_ns": 1794684745169825568
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "5",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl5.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 9032177,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58916964937126,
                      "time_ns": 1794684740155261339
                    },
                    "slice_started": {
                      "monotonic_ns": 58915860788237,
                      "time_ns": 1794684739051112942
                    },
                    "started": {
                      "monotonic_ns": 58915860566954,
                      "time_ns": 1794684739050891538
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr6.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl6.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58922272047253,
                      "time_ns": 1794684745462371928
                    },
                    "annotation_started": {
                      "monotonic_ns": 58917256195050,
                      "time_ns": 1794684740446519665
                    },
                    "chromosome": 6,
                    "finished": {
                      "monotonic_ns": 58922273127861,
                      "time_ns": 1794684745463452616
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "6",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl6.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 8584382,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58917256011046,
                      "time_ns": 1794684740446335230
                    },
                    "slice_started": {
                      "monotonic_ns": 58916150690029,
                      "time_ns": 1794684739341014743
                    },
                    "started": {
                      "monotonic_ns": 58916150399025,
                      "time_ns": 1794684739340723760
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr7.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl7.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58922270855918,
                      "time_ns": 1794684745461179801
                    },
                    "annotation_started": {
                      "monotonic_ns": 58917759350279,
                      "time_ns": 1794684740949674883
                    },
                    "chromosome": 7,
                    "finished": {
                      "monotonic_ns": 58922273169990,
                      "time_ns": 1794684745463494605
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "7",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl7.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7924914,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58917759208835,
                      "time_ns": 1794684740949532918
                    },
                    "slice_started": {
                      "monotonic_ns": 58916853364314,
                      "time_ns": 1794684740043688938
                    },
                    "started": {
                      "monotonic_ns": 58916853249279,
                      "time_ns": 1794684740043574024
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr8.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl8.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58922671211340,
                      "time_ns": 1794684745861535373
                    },
                    "annotation_started": {
                      "monotonic_ns": 58918361823872,
                      "time_ns": 1794684741552148637
                    },
                    "chromosome": 8,
                    "finished": {
                      "monotonic_ns": 58922672753370,
                      "time_ns": 1794684745863078165
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "8",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl8.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7556437,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58918361692657,
                      "time_ns": 1794684741552016560
                    },
                    "slice_started": {
                      "monotonic_ns": 58917356465161,
                      "time_ns": 1794684740546789936
                    },
                    "started": {
                      "monotonic_ns": 58917356336141,
                      "time_ns": 1794684740546660635
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr9.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl9.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58926190676257,
                      "time_ns": 1794684749381000271
                    },
                    "annotation_started": {
                      "monotonic_ns": 58922683667784,
                      "time_ns": 1794684745873992529
                    },
                    "chromosome": 9,
                    "finished": {
                      "monotonic_ns": 58926191325610,
                      "time_ns": 1794684749381650365
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "9",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl9.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 5978982,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58922683570282,
                      "time_ns": 1794684745873894696
                    },
                    "slice_started": {
                      "monotonic_ns": 58921979821983,
                      "time_ns": 1794684745170146737
                    },
                    "started": {
                      "monotonic_ns": 58921979696178,
                      "time_ns": 1794684745170020912
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr10.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl10.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58926992363093,
                      "time_ns": 1794684750182687497
                    },
                    "annotation_started": {
                      "monotonic_ns": 58923076347381,
                      "time_ns": 1794684746266671975
                    },
                    "chromosome": 10,
                    "finished": {
                      "monotonic_ns": 58926994020168,
                      "time_ns": 1794684750184344742
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "10",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl10.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7017520,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58923076182843,
                      "time_ns": 1794684746266506626
                    },
                    "slice_started": {
                      "monotonic_ns": 58922273534471,
                      "time_ns": 1794684745463859215
                    },
                    "started": {
                      "monotonic_ns": 58922273353042,
                      "time_ns": 1794684745463677757
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr11.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl11.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58926989231334,
                      "time_ns": 1794684750179555237
                    },
                    "annotation_started": {
                      "monotonic_ns": 58923078338359,
                      "time_ns": 1794684746268663124
                    },
                    "chromosome": 11,
                    "finished": {
                      "monotonic_ns": 58926990454579,
                      "time_ns": 1794684750180779384
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "11",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl11.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7109134,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58923078246127,
                      "time_ns": 1794684746268570711
                    },
                    "slice_started": {
                      "monotonic_ns": 58922273599192,
                      "time_ns": 1794684745463923936
                    },
                    "started": {
                      "monotonic_ns": 58922273434334,
                      "time_ns": 1794684745463759119
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr12.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl12.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58927287252108,
                      "time_ns": 1794684750477576312
                    },
                    "annotation_started": {
                      "monotonic_ns": 58923475490067,
                      "time_ns": 1794684746665814832
                    },
                    "chromosome": 12,
                    "finished": {
                      "monotonic_ns": 58927287951705,
                      "time_ns": 1794684750478276490
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "12",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl12.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 6749309,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58923475370714,
                      "time_ns": 1794684746665694708
                    },
                    "slice_started": {
                      "monotonic_ns": 58922673020609,
                      "time_ns": 1794684745863345364
                    },
                    "started": {
                      "monotonic_ns": 58922672907438,
                      "time_ns": 1794684745863232172
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr13.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl13.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58930003832989,
                      "time_ns": 1794684753194156912
                    },
                    "annotation_started": {
                      "monotonic_ns": 58926994548795,
                      "time_ns": 1794684750184873570
                    },
                    "chromosome": 13,
                    "finished": {
                      "monotonic_ns": 58930004426528,
                      "time_ns": 1794684753194751222
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "13",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl13.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 5531591,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58926994472362,
                      "time_ns": 1794684750184797117
                    },
                    "slice_started": {
                      "monotonic_ns": 58926191654324,
                      "time_ns": 1794684749381978999
                    },
                    "started": {
                      "monotonic_ns": 58926191511858,
                      "time_ns": 1794684749381836593
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr14.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl14.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58930200633858,
                      "time_ns": 1794684753390957861
                    },
                    "annotation_started": {
                      "monotonic_ns": 58927592566083,
                      "time_ns": 1794684750782890838
                    },
                    "chromosome": 14,
                    "finished": {
                      "monotonic_ns": 58930201715648,
                      "time_ns": 1794684753392040133
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "14",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl14.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4486826,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58927592471146,
                      "time_ns": 1794684750782795390
                    },
                    "slice_started": {
                      "monotonic_ns": 58926990740172,
                      "time_ns": 1794684750181064917
                    },
                    "started": {
                      "monotonic_ns": 58926990604248,
                      "time_ns": 1794684750180928923
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr15.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl15.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58929909668742,
                      "time_ns": 1794684753099992845
                    },
                    "annotation_started": {
                      "monotonic_ns": 58927501296913,
                      "time_ns": 1794684750691621688
                    },
                    "chromosome": 15,
                    "finished": {
                      "monotonic_ns": 58929910295693,
                      "time_ns": 1794684753100620437
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "15",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl15.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4074510,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58927501206394,
                      "time_ns": 1794684750691530457
                    },
                    "slice_started": {
                      "monotonic_ns": 58926994268772,
                      "time_ns": 1794684750184593487
                    },
                    "started": {
                      "monotonic_ns": 58926994159127,
                      "time_ns": 1794684750184483862
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr16.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl16.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58930304055475,
                      "time_ns": 1794684753494379528
                    },
                    "annotation_started": {
                      "monotonic_ns": 58927794388431,
                      "time_ns": 1794684750984713085
                    },
                    "chromosome": 16,
                    "finished": {
                      "monotonic_ns": 58930304856651,
                      "time_ns": 1794684753495181366
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "16",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl16.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4212959,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58927794242077,
                      "time_ns": 1794684750984566001
                    },
                    "slice_started": {
                      "monotonic_ns": 58927288225266,
                      "time_ns": 1794684750478550060
                    },
                    "started": {
                      "monotonic_ns": 58927288117574,
                      "time_ns": 1794684750478442309
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr17.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl17.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58932722533467,
                      "time_ns": 1794684755912857751
                    },
                    "annotation_started": {
                      "monotonic_ns": 58930513669174,
                      "time_ns": 1794684753703993858
                    },
                    "chromosome": 17,
                    "finished": {
                      "monotonic_ns": 58932722934265,
                      "time_ns": 1794684755913259070
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "17",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl17.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 3711659,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58930513526116,
                      "time_ns": 1794684753703849829
                    },
                    "slice_started": {
                      "monotonic_ns": 58929910626070,
                      "time_ns": 1794684753100950804
                    },
                    "started": {
                      "monotonic_ns": 58929910492971,
                      "time_ns": 1794684753100817716
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr18.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl18.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58932923675863,
                      "time_ns": 1794684756114000278
                    },
                    "annotation_started": {
                      "monotonic_ns": 58930611849381,
                      "time_ns": 1794684753802174156
                    },
                    "chromosome": 18,
                    "finished": {
                      "monotonic_ns": 58932924138167,
                      "time_ns": 1794684756114462871
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "18",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl18.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4239926,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58930611746449,
                      "time_ns": 1794684753802070443
                    },
                    "slice_started": {
                      "monotonic_ns": 58930008603088,
                      "time_ns": 1794684753198927683
                    },
                    "started": {
                      "monotonic_ns": 58930007414749,
                      "time_ns": 1794684753197739002
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr19.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl19.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58932612406136,
                      "time_ns": 1794684755802730299
                    },
                    "annotation_started": {
                      "monotonic_ns": 58930704310308,
                      "time_ns": 1794684753894635083
                    },
                    "chromosome": 19,
                    "finished": {
                      "monotonic_ns": 58932612838904,
                      "time_ns": 1794684755803163578
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "19",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl19.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 3140432,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58930704198038,
                      "time_ns": 1794684753894521951
                    },
                    "slice_started": {
                      "monotonic_ns": 58930202013044,
                      "time_ns": 1794684753392337809
                    },
                    "started": {
                      "monotonic_ns": 58930201893040,
                      "time_ns": 1794684753392217774
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr20.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl20.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58932515515286,
                      "time_ns": 1794684755705839299
                    },
                    "annotation_started": {
                      "monotonic_ns": 58930707692855,
                      "time_ns": 1794684753898017630
                    },
                    "chromosome": 20,
                    "finished": {
                      "monotonic_ns": 58932515866311,
                      "time_ns": 1794684755706191086
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "20",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl20.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 2951324,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58930707611664,
                      "time_ns": 1794684753897936208
                    },
                    "slice_started": {
                      "monotonic_ns": 58930305159216,
                      "time_ns": 1794684753495483951
                    },
                    "started": {
                      "monotonic_ns": 58930305022070,
                      "time_ns": 1794684753495346644
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr21.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl21.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58933620071418,
                      "time_ns": 1794684756810395171
                    },
                    "annotation_started": {
                      "monotonic_ns": 58932818609284,
                      "time_ns": 1794684756008933909
                    },
                    "chromosome": 21,
                    "finished": {
                      "monotonic_ns": 58933620633618,
                      "time_ns": 1794684756810958203
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "21",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl21.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 1982478,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58932818484000,
                      "time_ns": 1794684756008807643
                    },
                    "slice_started": {
                      "monotonic_ns": 58932516217918,
                      "time_ns": 1794684755706542553
                    },
                    "started": {
                      "monotonic_ns": 58932516069170,
                      "time_ns": 1794684755706393835
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr22.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl22.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58933715701754,
                      "time_ns": 1794684756906025176
                    },
                    "annotation_started": {
                      "monotonic_ns": 58932914061928,
                      "time_ns": 1794684756104386523
                    },
                    "chromosome": 22,
                    "finished": {
                      "monotonic_ns": 58933716597968,
                      "time_ns": 1794684756906922051
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "22",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r1-release/sl22.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 1765327,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58932913921245,
                      "time_ns": 1794684756104244918
                    },
                    "slice_started": {
                      "monotonic_ns": 58932613178438,
                      "time_ns": 1794684755803503193
                    },
                    "started": {
                      "monotonic_ns": 58932613034289,
                      "time_ns": 1794684755803359054
                    },
                    "status": "ok"
                  }
                ],
                "evaluated_total": {
                  "sources": [
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr1.anno.log",
                      "value": 308298
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr2.anno.log",
                      "value": 329821
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr3.anno.log",
                      "value": 275651
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr4.anno.log",
                      "value": 296434
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr5.anno.log",
                      "value": 253843
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr6.anno.log",
                      "value": 264414
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr7.anno.log",
                      "value": 222733
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr8.anno.log",
                      "value": 213367
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr9.anno.log",
                      "value": 168542
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr10.anno.log",
                      "value": 198281
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr11.anno.log",
                      "value": 200574
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr12.anno.log",
                      "value": 189356
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr13.anno.log",
                      "value": 156939
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr14.anno.log",
                      "value": 125631
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr15.anno.log",
                      "value": 114117
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr16.anno.log",
                      "value": 118437
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr17.anno.log",
                      "value": 102461
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr18.anno.log",
                      "value": 119403
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr19.anno.log",
                      "value": 86442
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr20.anno.log",
                      "value": 83452
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr21.anno.log",
                      "value": 55978
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/chr22.anno.log",
                      "value": 49540
                    }
                  ],
                  "value": 3933714
                },
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [],
              "issues": [],
              "metrics": {
                "cgroup_peak_bytes": {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/status.json",
                  "value": 136314880
                },
                "cpu_percent": {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt",
                  "value": 388.0
                },
                "exit_code": {
                  "field": "Exit status",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt",
                  "value": 0
                },
                "max_child_rss_kib": {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt",
                  "value": 31756
                },
                "system_seconds": {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt",
                  "value": 9.15
                },
                "user_seconds": {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt",
                  "value": 85.66
                },
                "wall_seconds": {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt",
                  "value": 24.42
                }
              },
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release",
              "round": 1,
              "status": {
                "directory_exists": true,
                "record": {
                  "exit_code": 0,
                  "finished_ns": 1794684756943147518,
                  "memory_peak_bytes": 136314880,
                  "oom_kill_detected": false,
                  "status": "ok"
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/status.json",
                "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/started",
                "state": "ok"
              },
              "valid": true
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": {
                  "GOMAXPROCS": "4",
                  "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-timing-nvme-17-echtvar-r2-release-1794684757-90895.scope",
                  "cpu.stat": "usage_usec 27762\nuser_usec 20822\nsystem_usec 6940\nnice_usec 0\ncore_sched.force_idle_usec 0",
                  "cpus": [
                    0,
                    1,
                    2,
                    3
                  ],
                  "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
                  "memory.current": "9486336",
                  "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.max": "17179869184",
                  "memory.peak": "9486336",
                  "memory.swap.max": "0",
                  "timestamp_ns": 1794684757049607850
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/constraints.json",
                "valid": true
              },
              "count_proof": {
                "all_22_chromosomes_ok": true,
                "annotation_log_statistics": [
                  {
                    "chromosome": 1,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr1.anno.log",
                      "value": 308298
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr1.anno.log",
                      "value": 308297
                    }
                  },
                  {
                    "chromosome": 2,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr2.anno.log",
                      "value": 329821
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr2.anno.log",
                      "value": 329821
                    }
                  },
                  {
                    "chromosome": 3,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr3.anno.log",
                      "value": 275651
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr3.anno.log",
                      "value": 275647
                    }
                  },
                  {
                    "chromosome": 4,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr4.anno.log",
                      "value": 296434
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr4.anno.log",
                      "value": 296434
                    }
                  },
                  {
                    "chromosome": 5,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr5.anno.log",
                      "value": 253843
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr5.anno.log",
                      "value": 253841
                    }
                  },
                  {
                    "chromosome": 6,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr6.anno.log",
                      "value": 264414
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr6.anno.log",
                      "value": 264084
                    }
                  },
                  {
                    "chromosome": 7,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr7.anno.log",
                      "value": 222733
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr7.anno.log",
                      "value": 222732
                    }
                  },
                  {
                    "chromosome": 8,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr8.anno.log",
                      "value": 213367
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr8.anno.log",
                      "value": 213367
                    }
                  },
                  {
                    "chromosome": 9,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr9.anno.log",
                      "value": 168542
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr9.anno.log",
                      "value": 168542
                    }
                  },
                  {
                    "chromosome": 10,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr10.anno.log",
                      "value": 198281
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr10.anno.log",
                      "value": 198281
                    }
                  },
                  {
                    "chromosome": 11,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr11.anno.log",
                      "value": 200574
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr11.anno.log",
                      "value": 200574
                    }
                  },
                  {
                    "chromosome": 12,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr12.anno.log",
                      "value": 189356
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr12.anno.log",
                      "value": 189356
                    }
                  },
                  {
                    "chromosome": 13,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr13.anno.log",
                      "value": 156939
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr13.anno.log",
                      "value": 156938
                    }
                  },
                  {
                    "chromosome": 14,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr14.anno.log",
                      "value": 125631
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr14.anno.log",
                      "value": 125631
                    }
                  },
                  {
                    "chromosome": 15,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr15.anno.log",
                      "value": 114117
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr15.anno.log",
                      "value": 114117
                    }
                  },
                  {
                    "chromosome": 16,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr16.anno.log",
                      "value": 118437
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr16.anno.log",
                      "value": 118437
                    }
                  },
                  {
                    "chromosome": 17,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr17.anno.log",
                      "value": 102461
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr17.anno.log",
                      "value": 102461
                    }
                  },
                  {
                    "chromosome": 18,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr18.anno.log",
                      "value": 119403
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr18.anno.log",
                      "value": 119401
                    }
                  },
                  {
                    "chromosome": 19,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr19.anno.log",
                      "value": 86442
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr19.anno.log",
                      "value": 86442
                    }
                  },
                  {
                    "chromosome": 20,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr20.anno.log",
                      "value": 83452
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr20.anno.log",
                      "value": 83452
                    }
                  },
                  {
                    "chromosome": 21,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr21.anno.log",
                      "value": 55978
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr21.anno.log",
                      "value": 55978
                    }
                  },
                  {
                    "chromosome": 22,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr22.anno.log",
                      "value": 49540
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr22.anno.log",
                      "value": 49536
                    }
                  }
                ],
                "chromosome_statuses": [
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr1.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl1.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58941424449010,
                      "time_ns": 1794684764614773143
                    },
                    "annotation_started": {
                      "monotonic_ns": 58935104171802,
                      "time_ns": 1794684758294496566
                    },
                    "chromosome": 1,
                    "finished": {
                      "monotonic_ns": 58941425785656,
                      "time_ns": 1794684764616110030
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "1",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl1.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 10999773,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58935104010400,
                      "time_ns": 1794684758294334414
                    },
                    "slice_started": {
                      "monotonic_ns": 58933901517495,
                      "time_ns": 1794684757091842190
                    },
                    "started": {
                      "monotonic_ns": 58933901309808,
                      "time_ns": 1794684757091634152
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr2.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl2.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58941924733974,
                      "time_ns": 1794684765115057887
                    },
                    "annotation_started": {
                      "monotonic_ns": 58935105272818,
                      "time_ns": 1794684758295597573
                    },
                    "chromosome": 2,
                    "finished": {
                      "monotonic_ns": 58941926398813,
                      "time_ns": 1794684765116723257
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "2",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl2.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 11750292,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58935105121195,
                      "time_ns": 1794684758295445589
                    },
                    "slice_started": {
                      "monotonic_ns": 58933902670619,
                      "time_ns": 1794684757092995274
                    },
                    "started": {
                      "monotonic_ns": 58933901787410,
                      "time_ns": 1794684757092111884
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr3.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl3.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58940418110147,
                      "time_ns": 1794684763608434331
                    },
                    "annotation_started": {
                      "monotonic_ns": 58934906457047,
                      "time_ns": 1794684758096781751
                    },
                    "chromosome": 3,
                    "finished": {
                      "monotonic_ns": 58940419311922,
                      "time_ns": 1794684763609636677
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "3",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl3.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 9815088,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58934906078870,
                      "time_ns": 1794684758096402984
                    },
                    "slice_started": {
                      "monotonic_ns": 58933902498137,
                      "time_ns": 1794684757092822852
                    },
                    "started": {
                      "monotonic_ns": 58933902254061,
                      "time_ns": 1794684757092578736
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr4.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl4.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58940824785160,
                      "time_ns": 1794684764015109253
                    },
                    "annotation_started": {
                      "monotonic_ns": 58935105935756,
                      "time_ns": 1794684758296260521
                    },
                    "chromosome": 4,
                    "finished": {
                      "monotonic_ns": 58940825881768,
                      "time_ns": 1794684764016206522
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "4",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl4.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 10519955,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58935105818527,
                      "time_ns": 1794684758296143272
                    },
                    "slice_started": {
                      "monotonic_ns": 58933902756159,
                      "time_ns": 1794684757093080904
                    },
                    "started": {
                      "monotonic_ns": 58933902436572,
                      "time_ns": 1794684757092761267
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr5.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl5.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58946638526186,
                      "time_ns": 1794684769828850380
                    },
                    "annotation_started": {
                      "monotonic_ns": 58941526740456,
                      "time_ns": 1794684764717065140
                    },
                    "chromosome": 5,
                    "finished": {
                      "monotonic_ns": 58946639490608,
                      "time_ns": 1794684769829815312
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "5",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl5.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 9032176,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58941526561872,
                      "time_ns": 1794684764716885846
                    },
                    "slice_started": {
                      "monotonic_ns": 58940419607143,
                      "time_ns": 1794684763609931898
                    },
                    "started": {
                      "monotonic_ns": 58940419494473,
                      "time_ns": 1794684763609819218
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr6.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl6.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58947037714205,
                      "time_ns": 1794684770228038398
                    },
                    "annotation_started": {
                      "monotonic_ns": 58941929670674,
                      "time_ns": 1794684765119995438
                    },
                    "chromosome": 6,
                    "finished": {
                      "monotonic_ns": 58947039332597,
                      "time_ns": 1794684770229657352
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "6",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl6.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 8584381,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58941929554376,
                      "time_ns": 1794684765119878991
                    },
                    "slice_started": {
                      "monotonic_ns": 58940826220631,
                      "time_ns": 1794684764016545365
                    },
                    "started": {
                      "monotonic_ns": 58940826080799,
                      "time_ns": 1794684764016405454
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr7.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl7.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58946838648938,
                      "time_ns": 1794684770028973031
                    },
                    "annotation_started": {
                      "monotonic_ns": 58942331586307,
                      "time_ns": 1794684765521911082
                    },
                    "chromosome": 7,
                    "finished": {
                      "monotonic_ns": 58946839486552,
                      "time_ns": 1794684770029811287
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "7",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl7.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7924913,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58942331457496,
                      "time_ns": 1794684765521781600
                    },
                    "slice_started": {
                      "monotonic_ns": 58941426488679,
                      "time_ns": 1794684764616813384
                    },
                    "started": {
                      "monotonic_ns": 58941426203116,
                      "time_ns": 1794684764616527620
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr8.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl8.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58947237900354,
                      "time_ns": 1794684770428224348
                    },
                    "annotation_started": {
                      "monotonic_ns": 58942829061934,
                      "time_ns": 1794684766019386708
                    },
                    "chromosome": 8,
                    "finished": {
                      "monotonic_ns": 58947239720324,
                      "time_ns": 1794684770430045058
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "8",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl8.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7556436,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58942828927362,
                      "time_ns": 1794684766019251215
                    },
                    "slice_started": {
                      "monotonic_ns": 58941926714923,
                      "time_ns": 1794684765117039678
                    },
                    "started": {
                      "monotonic_ns": 58941926585251,
                      "time_ns": 1794684765116910006
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr9.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl9.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58950749823967,
                      "time_ns": 1794684773940148100
                    },
                    "annotation_started": {
                      "monotonic_ns": 58947441209828,
                      "time_ns": 1794684770631534372
                    },
                    "chromosome": 9,
                    "finished": {
                      "monotonic_ns": 58950750464082,
                      "time_ns": 1794684773940788877
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "9",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl9.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 5978981,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58947441052113,
                      "time_ns": 1794684770631376106
                    },
                    "slice_started": {
                      "monotonic_ns": 58946639809643,
                      "time_ns": 1794684769830134368
                    },
                    "started": {
                      "monotonic_ns": 58946639678248,
                      "time_ns": 1794684769830003003
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr10.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl10.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58951550634428,
                      "time_ns": 1794684774740958502
                    },
                    "annotation_started": {
                      "monotonic_ns": 58947643728544,
                      "time_ns": 1794684770834053289
                    },
                    "chromosome": 10,
                    "finished": {
                      "monotonic_ns": 58951551372837,
                      "time_ns": 1794684774741697622
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "10",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl10.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7017520,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58947643611365,
                      "time_ns": 1794684770833935469
                    },
                    "slice_started": {
                      "monotonic_ns": 58946839761125,
                      "time_ns": 1794684770030085860
                    },
                    "started": {
                      "monotonic_ns": 58946839655077,
                      "time_ns": 1794684770029979842
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr11.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl11.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58951850947160,
                      "time_ns": 1794684775041271515
                    },
                    "annotation_started": {
                      "monotonic_ns": 58947942810277,
                      "time_ns": 1794684771133135052
                    },
                    "chromosome": 11,
                    "finished": {
                      "monotonic_ns": 58951852595339,
                      "time_ns": 1794684775042920093
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "11",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl11.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7109134,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58947942690534,
                      "time_ns": 1794684771133014497
                    },
                    "slice_started": {
                      "monotonic_ns": 58947039630333,
                      "time_ns": 1794684770229955058
                    },
                    "started": {
                      "monotonic_ns": 58947039496363,
                      "time_ns": 1794684770229821108
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr12.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl12.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58951850905723,
                      "time_ns": 1794684775041229556
                    },
                    "annotation_started": {
                      "monotonic_ns": 58948042879493,
                      "time_ns": 1794684771233204268
                    },
                    "chromosome": 12,
                    "finished": {
                      "monotonic_ns": 58951852389504,
                      "time_ns": 1794684775042714129
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "12",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl12.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 6749309,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58948042763356,
                      "time_ns": 1794684771233087610
                    },
                    "slice_started": {
                      "monotonic_ns": 58947240021857,
                      "time_ns": 1794684770430346611
                    },
                    "started": {
                      "monotonic_ns": 58947239887145,
                      "time_ns": 1794684770430211890
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr13.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl13.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58954564553202,
                      "time_ns": 1794684777754877486
                    },
                    "annotation_started": {
                      "monotonic_ns": 58951555405850,
                      "time_ns": 1794684774745730544
                    },
                    "chromosome": 13,
                    "finished": {
                      "monotonic_ns": 58954565143825,
                      "time_ns": 1794684777755468570
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "13",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl13.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 5531590,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58951555287278,
                      "time_ns": 1794684774745611772
                    },
                    "slice_started": {
                      "monotonic_ns": 58950750763762,
                      "time_ns": 1794684773941088527
                    },
                    "started": {
                      "monotonic_ns": 58950750646603,
                      "time_ns": 1794684773940971278
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr14.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl14.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58954760744521,
                      "time_ns": 1794684777951068604
                    },
                    "annotation_started": {
                      "monotonic_ns": 58952155701257,
                      "time_ns": 1794684775346025952
                    },
                    "chromosome": 14,
                    "finished": {
                      "monotonic_ns": 58954761536079,
                      "time_ns": 1794684777951860884
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "14",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl14.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4486826,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58952155544374,
                      "time_ns": 1794684775345868147
                    },
                    "slice_started": {
                      "monotonic_ns": 58951551656727,
                      "time_ns": 1794684774741981422
                    },
                    "started": {
                      "monotonic_ns": 58951551522296,
                      "time_ns": 1794684774741847051
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr15.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl15.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58954770278627,
                      "time_ns": 1794684777960603011
                    },
                    "annotation_started": {
                      "monotonic_ns": 58952360548542,
                      "time_ns": 1794684775550873287
                    },
                    "chromosome": 15,
                    "finished": {
                      "monotonic_ns": 58954770838593,
                      "time_ns": 1794684777961163317
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "15",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl15.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4074510,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58952360441773,
                      "time_ns": 1794684775550765877
                    },
                    "slice_started": {
                      "monotonic_ns": 58951852846428,
                      "time_ns": 1794684775043171182
                    },
                    "started": {
                      "monotonic_ns": 58951852637087,
                      "time_ns": 1794684775042961821
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr16.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl16.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58955061307089,
                      "time_ns": 1794684778251630822
                    },
                    "annotation_started": {
                      "monotonic_ns": 58952456247485,
                      "time_ns": 1794684775646572280
                    },
                    "chromosome": 16,
                    "finished": {
                      "monotonic_ns": 58955062326623,
                      "time_ns": 1794684778252651137
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "16",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl16.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4212959,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58952456144493,
                      "time_ns": 1794684775646468416
                    },
                    "slice_started": {
                      "monotonic_ns": 58951852988072,
                      "time_ns": 1794684775043312817
                    },
                    "started": {
                      "monotonic_ns": 58951852875211,
                      "time_ns": 1794684775043199946
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr17.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl17.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58957284273521,
                      "time_ns": 1794684780474597604
                    },
                    "annotation_started": {
                      "monotonic_ns": 58955071147187,
                      "time_ns": 1794684778261471932
                    },
                    "chromosome": 17,
                    "finished": {
                      "monotonic_ns": 58957284678337,
                      "time_ns": 1794684780475003101
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "17",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl17.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 3711659,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58955071069692,
                      "time_ns": 1794684778261394327
                    },
                    "slice_started": {
                      "monotonic_ns": 58954565420572,
                      "time_ns": 1794684777755745337
                    },
                    "started": {
                      "monotonic_ns": 58954565309835,
                      "time_ns": 1794684777755634490
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr18.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl18.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58957575085518,
                      "time_ns": 1794684780765410112
                    },
                    "annotation_started": {
                      "monotonic_ns": 58955267289363,
                      "time_ns": 1794684778457614068
                    },
                    "chromosome": 18,
                    "finished": {
                      "monotonic_ns": 58957575713872,
                      "time_ns": 1794684780766038596
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "18",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl18.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4239926,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58955267151606,
                      "time_ns": 1794684778457475689
                    },
                    "slice_started": {
                      "monotonic_ns": 58954764716589,
                      "time_ns": 1794684777955041204
                    },
                    "started": {
                      "monotonic_ns": 58954764590554,
                      "time_ns": 1794684777954915188
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr19.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl19.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58957182924455,
                      "time_ns": 1794684780373248389
                    },
                    "annotation_started": {
                      "monotonic_ns": 58955274553821,
                      "time_ns": 1794684778464878516
                    },
                    "chromosome": 19,
                    "finished": {
                      "monotonic_ns": 58957183292583,
                      "time_ns": 1794684780373617368
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "19",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl19.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 3140432,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58955274441010,
                      "time_ns": 1794684778464765334
                    },
                    "slice_started": {
                      "monotonic_ns": 58954771055698,
                      "time_ns": 1794684777961380433
                    },
                    "started": {
                      "monotonic_ns": 58954770962674,
                      "time_ns": 1794684777961287429
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr20.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl20.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58957273071149,
                      "time_ns": 1794684780463395212
                    },
                    "annotation_started": {
                      "monotonic_ns": 58955465735453,
                      "time_ns": 1794684778656060228
                    },
                    "chromosome": 20,
                    "finished": {
                      "monotonic_ns": 58957273401225,
                      "time_ns": 1794684780463726030
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "20",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl20.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 2951324,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58955465638351,
                      "time_ns": 1794684778655962184
                    },
                    "slice_started": {
                      "monotonic_ns": 58955062782925,
                      "time_ns": 1794684778253107600
                    },
                    "started": {
                      "monotonic_ns": 58955062610733,
                      "time_ns": 1794684778252935488
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr21.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl21.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58958288265845,
                      "time_ns": 1794684781478589838
                    },
                    "annotation_started": {
                      "monotonic_ns": 58957486285460,
                      "time_ns": 1794684780676610154
                    },
                    "chromosome": 21,
                    "finished": {
                      "monotonic_ns": 58958288500152,
                      "time_ns": 1794684781478824927
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "21",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl21.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 1982478,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58957486170195,
                      "time_ns": 1794684780676493917
                    },
                    "slice_started": {
                      "monotonic_ns": 58957183585209,
                      "time_ns": 1794684780373909974
                    },
                    "started": {
                      "monotonic_ns": 58957183473701,
                      "time_ns": 1794684780373798456
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr22.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl22.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58958376418754,
                      "time_ns": 1794684781566742216
                    },
                    "annotation_started": {
                      "monotonic_ns": 58957574540240,
                      "time_ns": 1794684780764864864
                    },
                    "chromosome": 22,
                    "finished": {
                      "monotonic_ns": 58958376754701,
                      "time_ns": 1794684781567079366
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "22",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r2-release/sl22.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 1765327,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58957574430244,
                      "time_ns": 1794684780764754017
                    },
                    "slice_started": {
                      "monotonic_ns": 58957273708609,
                      "time_ns": 1794684780464033384
                    },
                    "started": {
                      "monotonic_ns": 58957273574569,
                      "time_ns": 1794684780463899303
                    },
                    "status": "ok"
                  }
                ],
                "evaluated_total": {
                  "sources": [
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr1.anno.log",
                      "value": 308298
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr2.anno.log",
                      "value": 329821
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr3.anno.log",
                      "value": 275651
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr4.anno.log",
                      "value": 296434
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr5.anno.log",
                      "value": 253843
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr6.anno.log",
                      "value": 264414
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr7.anno.log",
                      "value": 222733
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr8.anno.log",
                      "value": 213367
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr9.anno.log",
                      "value": 168542
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr10.anno.log",
                      "value": 198281
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr11.anno.log",
                      "value": 200574
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr12.anno.log",
                      "value": 189356
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr13.anno.log",
                      "value": 156939
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr14.anno.log",
                      "value": 125631
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr15.anno.log",
                      "value": 114117
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr16.anno.log",
                      "value": 118437
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr17.anno.log",
                      "value": 102461
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr18.anno.log",
                      "value": 119403
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr19.anno.log",
                      "value": 86442
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr20.anno.log",
                      "value": 83452
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr21.anno.log",
                      "value": 55978
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/chr22.anno.log",
                      "value": 49540
                    }
                  ],
                  "value": 3933714
                },
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [],
              "issues": [],
              "metrics": {
                "cgroup_peak_bytes": {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/status.json",
                  "value": 136704000
                },
                "cpu_percent": {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt",
                  "value": 388.0
                },
                "exit_code": {
                  "field": "Exit status",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt",
                  "value": 0
                },
                "max_child_rss_kib": {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt",
                  "value": 31736
                },
                "system_seconds": {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt",
                  "value": 9.26
                },
                "user_seconds": {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt",
                  "value": 85.96
                },
                "wall_seconds": {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt",
                  "value": 24.52
                }
              },
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release",
              "round": 2,
              "status": {
                "directory_exists": true,
                "record": {
                  "exit_code": 0,
                  "finished_ns": 1794684781604865327,
                  "memory_peak_bytes": 136704000,
                  "oom_kill_detected": false,
                  "status": "ok"
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/status.json",
                "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/started",
                "state": "ok"
              },
              "valid": true
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": {
                  "GOMAXPROCS": "4",
                  "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-timing-nvme-17-echtvar-r3-release-1794684798-90895.scope",
                  "cpu.stat": "usage_usec 28801\nuser_usec 20856\nsystem_usec 7945\nnice_usec 0\ncore_sched.force_idle_usec 0",
                  "cpus": [
                    0,
                    1,
                    2,
                    3
                  ],
                  "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
                  "memory.current": "9474048",
                  "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.max": "17179869184",
                  "memory.peak": "9474048",
                  "memory.swap.max": "0",
                  "timestamp_ns": 1794684798256141953
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/constraints.json",
                "valid": true
              },
              "count_proof": {
                "all_22_chromosomes_ok": true,
                "annotation_log_statistics": [
                  {
                    "chromosome": 1,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr1.anno.log",
                      "value": 308298
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr1.anno.log",
                      "value": 308297
                    }
                  },
                  {
                    "chromosome": 2,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr2.anno.log",
                      "value": 329821
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr2.anno.log",
                      "value": 329821
                    }
                  },
                  {
                    "chromosome": 3,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr3.anno.log",
                      "value": 275651
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr3.anno.log",
                      "value": 275647
                    }
                  },
                  {
                    "chromosome": 4,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr4.anno.log",
                      "value": 296434
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr4.anno.log",
                      "value": 296434
                    }
                  },
                  {
                    "chromosome": 5,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr5.anno.log",
                      "value": 253843
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr5.anno.log",
                      "value": 253841
                    }
                  },
                  {
                    "chromosome": 6,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr6.anno.log",
                      "value": 264414
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr6.anno.log",
                      "value": 264084
                    }
                  },
                  {
                    "chromosome": 7,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr7.anno.log",
                      "value": 222733
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr7.anno.log",
                      "value": 222732
                    }
                  },
                  {
                    "chromosome": 8,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr8.anno.log",
                      "value": 213367
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr8.anno.log",
                      "value": 213367
                    }
                  },
                  {
                    "chromosome": 9,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr9.anno.log",
                      "value": 168542
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr9.anno.log",
                      "value": 168542
                    }
                  },
                  {
                    "chromosome": 10,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr10.anno.log",
                      "value": 198281
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr10.anno.log",
                      "value": 198281
                    }
                  },
                  {
                    "chromosome": 11,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr11.anno.log",
                      "value": 200574
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr11.anno.log",
                      "value": 200574
                    }
                  },
                  {
                    "chromosome": 12,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr12.anno.log",
                      "value": 189356
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr12.anno.log",
                      "value": 189356
                    }
                  },
                  {
                    "chromosome": 13,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr13.anno.log",
                      "value": 156939
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr13.anno.log",
                      "value": 156938
                    }
                  },
                  {
                    "chromosome": 14,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr14.anno.log",
                      "value": 125631
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr14.anno.log",
                      "value": 125631
                    }
                  },
                  {
                    "chromosome": 15,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr15.anno.log",
                      "value": 114117
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr15.anno.log",
                      "value": 114117
                    }
                  },
                  {
                    "chromosome": 16,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr16.anno.log",
                      "value": 118437
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr16.anno.log",
                      "value": 118437
                    }
                  },
                  {
                    "chromosome": 17,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr17.anno.log",
                      "value": 102461
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr17.anno.log",
                      "value": 102461
                    }
                  },
                  {
                    "chromosome": 18,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr18.anno.log",
                      "value": 119403
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr18.anno.log",
                      "value": 119401
                    }
                  },
                  {
                    "chromosome": 19,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr19.anno.log",
                      "value": 86442
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr19.anno.log",
                      "value": 86442
                    }
                  },
                  {
                    "chromosome": 20,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr20.anno.log",
                      "value": 83452
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr20.anno.log",
                      "value": 83452
                    }
                  },
                  {
                    "chromosome": 21,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr21.anno.log",
                      "value": 55978
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr21.anno.log",
                      "value": 55978
                    }
                  },
                  {
                    "chromosome": 22,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr22.anno.log",
                      "value": 49540
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr22.anno.log",
                      "value": 49536
                    }
                  }
                ],
                "chromosome_statuses": [
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr1.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl1.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58982627644911,
                      "time_ns": 1794684805817969115
                    },
                    "annotation_started": {
                      "monotonic_ns": 58976311898324,
                      "time_ns": 1794684799502223099
                    },
                    "chromosome": 1,
                    "finished": {
                      "monotonic_ns": 58982628781975,
                      "time_ns": 1794684805819106669
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "1",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl1.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 10999774,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58976311725652,
                      "time_ns": 1794684799502049415
                    },
                    "slice_started": {
                      "monotonic_ns": 58975108704497,
                      "time_ns": 1794684798299028931
                    },
                    "started": {
                      "monotonic_ns": 58975108359042,
                      "time_ns": 1794684798298683376
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr2.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl2.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58983128318865,
                      "time_ns": 1794684806318642738
                    },
                    "annotation_started": {
                      "monotonic_ns": 58976412398263,
                      "time_ns": 1794684799602723038
                    },
                    "chromosome": 2,
                    "finished": {
                      "monotonic_ns": 58983130096094,
                      "time_ns": 1794684806320420338
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "2",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl2.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 11750293,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58976412217195,
                      "time_ns": 1794684799602540918
                    },
                    "slice_started": {
                      "monotonic_ns": 58975109399405,
                      "time_ns": 1794684798299723619
                    },
                    "started": {
                      "monotonic_ns": 58975108796699,
                      "time_ns": 1794684798299121003
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr3.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl3.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58981822477591,
                      "time_ns": 1794684805012801825
                    },
                    "annotation_started": {
                      "monotonic_ns": 58976211927043,
                      "time_ns": 1794684799402251597
                    },
                    "chromosome": 3,
                    "finished": {
                      "monotonic_ns": 58981823511582,
                      "time_ns": 1794684805013836197
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "3",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl3.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 9815089,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58976211707973,
                      "time_ns": 1794684799402032347
                    },
                    "slice_started": {
                      "monotonic_ns": 58975109581104,
                      "time_ns": 1794684798299905819
                    },
                    "started": {
                      "monotonic_ns": 58975109211153,
                      "time_ns": 1794684798299535648
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr4.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl4.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58981922547787,
                      "time_ns": 1794684805112871680
                    },
                    "annotation_started": {
                      "monotonic_ns": 58976212070080,
                      "time_ns": 1794684799402394514
                    },
                    "chromosome": 4,
                    "finished": {
                      "monotonic_ns": 58981923642712,
                      "time_ns": 1794684805113967487
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "4",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl4.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 10519955,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58976211674771,
                      "time_ns": 1794684799401998815
                    },
                    "slice_started": {
                      "monotonic_ns": 58975109796516,
                      "time_ns": 1794684798300121261
                    },
                    "started": {
                      "monotonic_ns": 58975109709003,
                      "time_ns": 1794684798300033698
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr5.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl5.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58988041539464,
                      "time_ns": 1794684811231863447
                    },
                    "annotation_started": {
                      "monotonic_ns": 58982929385968,
                      "time_ns": 1794684806119710713
                    },
                    "chromosome": 5,
                    "finished": {
                      "monotonic_ns": 58988042866242,
                      "time_ns": 1794684811233190786
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "5",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl5.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 9032177,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58982929265252,
                      "time_ns": 1794684806119589617
                    },
                    "slice_started": {
                      "monotonic_ns": 58981825368740,
                      "time_ns": 1794684805015693415
                    },
                    "started": {
                      "monotonic_ns": 58981823716274,
                      "time_ns": 1794684805014040989
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr6.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl6.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58987942591143,
                      "time_ns": 1794684811132915147
                    },
                    "annotation_started": {
                      "monotonic_ns": 58982927730336,
                      "time_ns": 1794684806118055061
                    },
                    "chromosome": 6,
                    "finished": {
                      "monotonic_ns": 58987943520569,
                      "time_ns": 1794684811133845304
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "6",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl6.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 8584382,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58982927607927,
                      "time_ns": 1794684806117932191
                    },
                    "slice_started": {
                      "monotonic_ns": 58981923940238,
                      "time_ns": 1794684805114264993
                    },
                    "started": {
                      "monotonic_ns": 58981923824071,
                      "time_ns": 1794684805114148796
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr7.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl7.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58988143295068,
                      "time_ns": 1794684811333619261
                    },
                    "annotation_started": {
                      "monotonic_ns": 58983533173321,
                      "time_ns": 1794684806723497875
                    },
                    "chromosome": 7,
                    "finished": {
                      "monotonic_ns": 58988144126991,
                      "time_ns": 1794684811334451656
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "7",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl7.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7924914,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58983532979329,
                      "time_ns": 1794684806723303032
                    },
                    "slice_started": {
                      "monotonic_ns": 58982629090481,
                      "time_ns": 1794684805819415096
                    },
                    "started": {
                      "monotonic_ns": 58982628930342,
                      "time_ns": 1794684805819255077
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr8.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl8.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58988543398400,
                      "time_ns": 1794684811733722433
                    },
                    "annotation_started": {
                      "monotonic_ns": 58984136242102,
                      "time_ns": 1794684807326566816
                    },
                    "chromosome": 8,
                    "finished": {
                      "monotonic_ns": 58988545153798,
                      "time_ns": 1794684811735478533
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "8",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl8.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7556437,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58984136089316,
                      "time_ns": 1794684807326413220
                    },
                    "slice_started": {
                      "monotonic_ns": 58983130435057,
                      "time_ns": 1794684806320759762
                    },
                    "started": {
                      "monotonic_ns": 58983130287041,
                      "time_ns": 1794684806320611615
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr9.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl9.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58992156946221,
                      "time_ns": 1794684815347270115
                    },
                    "annotation_started": {
                      "monotonic_ns": 58988748066788,
                      "time_ns": 1794684811938391342
                    },
                    "chromosome": 9,
                    "finished": {
                      "monotonic_ns": 58992157817539,
                      "time_ns": 1794684815348142273
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "9",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl9.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 5978982,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58988747900598,
                      "time_ns": 1794684811938224431
                    },
                    "slice_started": {
                      "monotonic_ns": 58987943824797,
                      "time_ns": 1794684811134149522
                    },
                    "started": {
                      "monotonic_ns": 58987943709071,
                      "time_ns": 1794684811134033826
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr10.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl10.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58992853316004,
                      "time_ns": 1794684816043640117
                    },
                    "annotation_started": {
                      "monotonic_ns": 58988845876752,
                      "time_ns": 1794684812036201477
                    },
                    "chromosome": 10,
                    "finished": {
                      "monotonic_ns": 58992854385361,
                      "time_ns": 1794684816044710106
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "10",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl10.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7017520,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58988845760135,
                      "time_ns": 1794684812036084409
                    },
                    "slice_started": {
                      "monotonic_ns": 58988043221726,
                      "time_ns": 1794684811233546341
                    },
                    "started": {
                      "monotonic_ns": 58988043047330,
                      "time_ns": 1794684811233372085
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr11.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl11.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58992857372490,
                      "time_ns": 1794684816047697014
                    },
                    "annotation_started": {
                      "monotonic_ns": 58988945953390,
                      "time_ns": 1794684812136278125
                    },
                    "chromosome": 11,
                    "finished": {
                      "monotonic_ns": 58992858095119,
                      "time_ns": 1794684816048419904
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "11",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl11.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7109134,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58988945835129,
                      "time_ns": 1794684812136159383
                    },
                    "slice_started": {
                      "monotonic_ns": 58988144383951,
                      "time_ns": 1794684811334708706
                    },
                    "started": {
                      "monotonic_ns": 58988144274717,
                      "time_ns": 1794684811334599472
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr12.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl12.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58993255136863,
                      "time_ns": 1794684816445461067
                    },
                    "annotation_started": {
                      "monotonic_ns": 58989348692024,
                      "time_ns": 1794684812539016769
                    },
                    "chromosome": 12,
                    "finished": {
                      "monotonic_ns": 58993255853692,
                      "time_ns": 1794684816446178446
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "12",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl12.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 6749309,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58989348578833,
                      "time_ns": 1794684812538902986
                    },
                    "slice_started": {
                      "monotonic_ns": 58988545463967,
                      "time_ns": 1794684811735788702
                    },
                    "started": {
                      "monotonic_ns": 58988545308818,
                      "time_ns": 1794684811735633532
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr13.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl13.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58995875428940,
                      "time_ns": 1794684819065753244
                    },
                    "annotation_started": {
                      "monotonic_ns": 58992863334074,
                      "time_ns": 1794684816053658798
                    },
                    "chromosome": 13,
                    "finished": {
                      "monotonic_ns": 58995876282023,
                      "time_ns": 1794684819066606798
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "13",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl13.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 5531591,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58992863245378,
                      "time_ns": 1794684816053569953
                    },
                    "slice_started": {
                      "monotonic_ns": 58992158165418,
                      "time_ns": 1794684815348490103
                    },
                    "started": {
                      "monotonic_ns": 58992158005971,
                      "time_ns": 1794684815348330715
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr14.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl14.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58996065300663,
                      "time_ns": 1794684819255624416
                    },
                    "annotation_started": {
                      "monotonic_ns": 58993456146548,
                      "time_ns": 1794684816646471283
                    },
                    "chromosome": 14,
                    "finished": {
                      "monotonic_ns": 58996066304307,
                      "time_ns": 1794684819256628882
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "14",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl14.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4486826,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58993456037134,
                      "time_ns": 1794684816646361337
                    },
                    "slice_started": {
                      "monotonic_ns": 58992854700880,
                      "time_ns": 1794684816045025515
                    },
                    "started": {
                      "monotonic_ns": 58992854535641,
                      "time_ns": 1794684816044860406
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr15.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl15.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58995865362880,
                      "time_ns": 1794684819055686793
                    },
                    "annotation_started": {
                      "monotonic_ns": 58993461719206,
                      "time_ns": 1794684816652043730
                    },
                    "chromosome": 15,
                    "finished": {
                      "monotonic_ns": 58995866182341,
                      "time_ns": 1794684819056507115
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "15",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl15.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4074510,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58993461599923,
                      "time_ns": 1794684816651924137
                    },
                    "slice_started": {
                      "monotonic_ns": 58992858328655,
                      "time_ns": 1794684816048653390
                    },
                    "started": {
                      "monotonic_ns": 58992858234500,
                      "time_ns": 1794684816048559254
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr16.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl16.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58996273797070,
                      "time_ns": 1794684819464121154
                    },
                    "annotation_started": {
                      "monotonic_ns": 58993764631128,
                      "time_ns": 1794684816954955782
                    },
                    "chromosome": 16,
                    "finished": {
                      "monotonic_ns": 58996274789854,
                      "time_ns": 1794684819465114449
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "16",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl16.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4212959,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58993764537022,
                      "time_ns": 1794684816954861086
                    },
                    "slice_started": {
                      "monotonic_ns": 58993256165644,
                      "time_ns": 1794684816446490379
                    },
                    "started": {
                      "monotonic_ns": 58993256029500,
                      "time_ns": 1794684816446354114
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr17.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl17.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58998579528998,
                      "time_ns": 1794684821769853112
                    },
                    "annotation_started": {
                      "monotonic_ns": 58996369578276,
                      "time_ns": 1794684819559902880
                    },
                    "chromosome": 17,
                    "finished": {
                      "monotonic_ns": 58998579949914,
                      "time_ns": 1794684821770274689
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "17",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl17.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 3711659,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58996369455677,
                      "time_ns": 1794684819559779720
                    },
                    "slice_started": {
                      "monotonic_ns": 58995866785968,
                      "time_ns": 1794684819057110532
                    },
                    "started": {
                      "monotonic_ns": 58995866359101,
                      "time_ns": 1794684819056683845
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr18.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl18.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58998794718245,
                      "time_ns": 1794684821985042088
                    },
                    "annotation_started": {
                      "monotonic_ns": 58996485493153,
                      "time_ns": 1794684819675817798
                    },
                    "chromosome": 18,
                    "finished": {
                      "monotonic_ns": 58998795656247,
                      "time_ns": 1794684821985980812
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "18",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl18.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4239926,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58996485355586,
                      "time_ns": 1794684819675679610
                    },
                    "slice_started": {
                      "monotonic_ns": 58995878663240,
                      "time_ns": 1794684819068987945
                    },
                    "started": {
                      "monotonic_ns": 58995878537566,
                      "time_ns": 1794684819068862290
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr19.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl19.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58998480937695,
                      "time_ns": 1794684821671261859
                    },
                    "annotation_started": {
                      "monotonic_ns": 58996569904113,
                      "time_ns": 1794684819760228838
                    },
                    "chromosome": 19,
                    "finished": {
                      "monotonic_ns": 58998481303648,
                      "time_ns": 1794684821671628373
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "19",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl19.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 3140432,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58996569805389,
                      "time_ns": 1794684819760129502
                    },
                    "slice_started": {
                      "monotonic_ns": 58996066672435,
                      "time_ns": 1794684819256997160
                    },
                    "started": {
                      "monotonic_ns": 58996066530059,
                      "time_ns": 1794684819256854784
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr20.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl20.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58998483767540,
                      "time_ns": 1794684821674091934
                    },
                    "annotation_started": {
                      "monotonic_ns": 58996677574573,
                      "time_ns": 1794684819867899168
                    },
                    "chromosome": 20,
                    "finished": {
                      "monotonic_ns": 58998484171184,
                      "time_ns": 1794684821674495899
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "20",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl20.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 2951324,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58996677450732,
                      "time_ns": 1794684819867774575
                    },
                    "slice_started": {
                      "monotonic_ns": 58996275301229,
                      "time_ns": 1794684819465625994
                    },
                    "started": {
                      "monotonic_ns": 58996275088692,
                      "time_ns": 1794684819465413237
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr21.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl21.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58999584540021,
                      "time_ns": 1794684822774863484
                    },
                    "annotation_started": {
                      "monotonic_ns": 58998782535582,
                      "time_ns": 1794684821972860257
                    },
                    "chromosome": 21,
                    "finished": {
                      "monotonic_ns": 58999585405698,
                      "time_ns": 1794684822775729771
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "21",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl21.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 1982477,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58998782394920,
                      "time_ns": 1794684821972718492
                    },
                    "slice_started": {
                      "monotonic_ns": 58998481627373,
                      "time_ns": 1794684821671951957
                    },
                    "started": {
                      "monotonic_ns": 58998481476702,
                      "time_ns": 1794684821671801336
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/17_fields/archives/chr22.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl22.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 58999586686590,
                      "time_ns": 1794684822777011115
                    },
                    "annotation_started": {
                      "monotonic_ns": 58998785070336,
                      "time_ns": 1794684821975395081
                    },
                    "chromosome": 22,
                    "finished": {
                      "monotonic_ns": 58999586971633,
                      "time_ns": 1794684822777296357
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "22",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-17-echtvar-r3-release/sl22.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 1765326,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 58998784990527,
                      "time_ns": 1794684821975314521
                    },
                    "slice_started": {
                      "monotonic_ns": 58998484398288,
                      "time_ns": 1794684821674723023
                    },
                    "started": {
                      "monotonic_ns": 58998484309372,
                      "time_ns": 1794684821674634117
                    },
                    "status": "ok"
                  }
                ],
                "evaluated_total": {
                  "sources": [
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr1.anno.log",
                      "value": 308298
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr2.anno.log",
                      "value": 329821
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr3.anno.log",
                      "value": 275651
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr4.anno.log",
                      "value": 296434
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr5.anno.log",
                      "value": 253843
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr6.anno.log",
                      "value": 264414
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr7.anno.log",
                      "value": 222733
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr8.anno.log",
                      "value": 213367
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr9.anno.log",
                      "value": 168542
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr10.anno.log",
                      "value": 198281
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr11.anno.log",
                      "value": 200574
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr12.anno.log",
                      "value": 189356
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr13.anno.log",
                      "value": 156939
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr14.anno.log",
                      "value": 125631
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr15.anno.log",
                      "value": 114117
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr16.anno.log",
                      "value": 118437
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr17.anno.log",
                      "value": 102461
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr18.anno.log",
                      "value": 119403
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr19.anno.log",
                      "value": 86442
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr20.anno.log",
                      "value": 83452
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr21.anno.log",
                      "value": 55978
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/chr22.anno.log",
                      "value": 49540
                    }
                  ],
                  "value": 3933714
                },
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [],
              "issues": [],
              "metrics": {
                "cgroup_peak_bytes": {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/status.json",
                  "value": 136237056
                },
                "cpu_percent": {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt",
                  "value": 388.0
                },
                "exit_code": {
                  "field": "Exit status",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt",
                  "value": 0
                },
                "max_child_rss_kib": {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt",
                  "value": 31760
                },
                "system_seconds": {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt",
                  "value": 9.06
                },
                "user_seconds": {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt",
                  "value": 86.12
                },
                "wall_seconds": {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt",
                  "value": 24.52
                }
              },
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release",
              "round": 3,
              "status": {
                "directory_exists": true,
                "record": {
                  "exit_code": 0,
                  "finished_ns": 1794684822814081599,
                  "memory_peak_bytes": 136237056,
                  "oom_kill_detected": false,
                  "status": "ok"
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/status.json",
                "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/started",
                "state": "ok"
              },
              "valid": true
            }
          ],
          "summary": {
            "cgroup_peak_bytes": {
              "maximum": 136704000,
              "median": 136314880,
              "minimum": 136237056,
              "n": 3,
              "sources": [
                {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/status.json",
                  "value": 136314880
                },
                {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/status.json",
                  "value": 136704000
                },
                {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/status.json",
                  "value": 136237056
                }
              ],
              "values": [
                136314880,
                136704000,
                136237056
              ]
            },
            "cpu_percent": {
              "maximum": 388.0,
              "median": 388.0,
              "minimum": 388.0,
              "n": 3,
              "sources": [
                {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt",
                  "value": 388.0
                },
                {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt",
                  "value": 388.0
                },
                {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt",
                  "value": 388.0
                }
              ],
              "values": [
                388.0,
                388.0,
                388.0
              ]
            },
            "max_child_rss_kib": {
              "maximum": 31760,
              "median": 31756,
              "minimum": 31736,
              "n": 3,
              "sources": [
                {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt",
                  "value": 31756
                },
                {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt",
                  "value": 31736
                },
                {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt",
                  "value": 31760
                }
              ],
              "values": [
                31756,
                31736,
                31760
              ]
            },
            "system_seconds": {
              "maximum": 9.26,
              "median": 9.15,
              "minimum": 9.06,
              "n": 3,
              "sources": [
                {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt",
                  "value": 9.15
                },
                {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt",
                  "value": 9.26
                },
                {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt",
                  "value": 9.06
                }
              ],
              "values": [
                9.15,
                9.26,
                9.06
              ]
            },
            "user_seconds": {
              "maximum": 86.12,
              "median": 85.96,
              "minimum": 85.66,
              "n": 3,
              "sources": [
                {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt",
                  "value": 85.66
                },
                {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt",
                  "value": 85.96
                },
                {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt",
                  "value": 86.12
                }
              ],
              "values": [
                85.66,
                85.96,
                86.12
              ]
            },
            "wall_seconds": {
              "maximum": 24.52,
              "median": 24.52,
              "minimum": 24.42,
              "n": 3,
              "sources": [
                {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt",
                  "value": 24.42
                },
                {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt",
                  "value": 24.52
                },
                {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt",
                  "value": 24.52
                }
              ],
              "values": [
                24.42,
                24.52,
                24.52
              ]
            }
          },
          "valid_rounds": 3
        },
        "immkv": {
          "rounds": [
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": {
                  "GOMAXPROCS": "4",
                  "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-timing-nvme-17-immkv-r1-release-1794684723-90895.scope",
                  "cpu.stat": "usage_usec 29803\nuser_usec 20862\nsystem_usec 8941\nnice_usec 0\ncore_sched.force_idle_usec 0",
                  "cpus": [
                    0,
                    1,
                    2,
                    3
                  ],
                  "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
                  "memory.current": "9834496",
                  "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.max": "17179869184",
                  "memory.peak": "9834496",
                  "memory.swap.max": "0",
                  "timestamp_ns": 1794684723970939370
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/constraints.json",
                "valid": true
              },
              "count_proof": {
                "hits": {
                  "field": "hits",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log",
                  "value": 3875208
                },
                "rows": {
                  "field": "rows",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log",
                  "value": 3933714
                }
              },
              "immkv_phase_statistics": {
                "format_seconds": {
                  "field": "format_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log",
                  "value": 6.99
                },
                "key_seconds": {
                  "field": "key_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log",
                  "value": 0.29
                },
                "query_seconds": {
                  "field": "query_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log",
                  "value": 4.46
                },
                "read_gib_rounded": {
                  "field": "read_gib_rounded",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log",
                  "value": 0.0
                },
                "wall_seconds": {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/immkv.log",
                  "value": 8.24
                }
              },
              "invalid_reasons": [],
              "issues": [],
              "metrics": {
                "cgroup_peak_bytes": {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/status.json",
                  "value": 520769536
                },
                "cpu_percent": {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt",
                  "value": 354.0
                },
                "exit_code": {
                  "field": "Exit status",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt",
                  "value": 0
                },
                "max_child_rss_kib": {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt",
                  "value": 413424
                },
                "system_seconds": {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt",
                  "value": 2.15
                },
                "user_seconds": {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt",
                  "value": 27.49
                },
                "wall_seconds": {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt",
                  "value": 8.37
                }
              },
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release",
              "round": 1,
              "status": {
                "directory_exists": true,
                "record": {
                  "exit_code": 0,
                  "finished_ns": 1794684732379237053,
                  "memory_peak_bytes": 520769536,
                  "oom_kill_detected": false,
                  "status": "ok"
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/status.json",
                "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/started",
                "state": "ok"
              },
              "valid": true
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": {
                  "GOMAXPROCS": "4",
                  "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-timing-nvme-17-immkv-r2-release-1794684781-90895.scope",
                  "cpu.stat": "usage_usec 28744\nuser_usec 22797\nsystem_usec 5947\nnice_usec 0\ncore_sched.force_idle_usec 0",
                  "cpus": [
                    0,
                    1,
                    2,
                    3
                  ],
                  "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
                  "memory.current": "9584640",
                  "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.max": "17179869184",
                  "memory.peak": "9584640",
                  "memory.swap.max": "0",
                  "timestamp_ns": 1794684781666519967
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/constraints.json",
                "valid": true
              },
              "count_proof": {
                "hits": {
                  "field": "hits",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log",
                  "value": 3875208
                },
                "rows": {
                  "field": "rows",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log",
                  "value": 3933714
                }
              },
              "immkv_phase_statistics": {
                "format_seconds": {
                  "field": "format_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log",
                  "value": 7.12
                },
                "key_seconds": {
                  "field": "key_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log",
                  "value": 0.3
                },
                "query_seconds": {
                  "field": "query_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log",
                  "value": 4.12
                },
                "read_gib_rounded": {
                  "field": "read_gib_rounded",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log",
                  "value": 0.0
                },
                "wall_seconds": {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/immkv.log",
                  "value": 8.049
                }
              },
              "invalid_reasons": [],
              "issues": [],
              "metrics": {
                "cgroup_peak_bytes": {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/status.json",
                  "value": 526442496
                },
                "cpu_percent": {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt",
                  "value": 367.0
                },
                "exit_code": {
                  "field": "Exit status",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt",
                  "value": 0
                },
                "max_child_rss_kib": {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt",
                  "value": 419928
                },
                "system_seconds": {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt",
                  "value": 2.1
                },
                "user_seconds": {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt",
                  "value": 27.96
                },
                "wall_seconds": {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt",
                  "value": 8.17
                }
              },
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release",
              "round": 2,
              "status": {
                "directory_exists": true,
                "record": {
                  "exit_code": 0,
                  "finished_ns": 1794684789873877599,
                  "memory_peak_bytes": 526442496,
                  "oom_kill_detected": false,
                  "status": "ok"
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/status.json",
                "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/started",
                "state": "ok"
              },
              "valid": true
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": {
                  "GOMAXPROCS": "4",
                  "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-timing-nvme-17-immkv-r3-release-1794684789-90895.scope",
                  "cpu.stat": "usage_usec 28755\nuser_usec 20823\nsystem_usec 7932\nnice_usec 0\ncore_sched.force_idle_usec 0",
                  "cpus": [
                    0,
                    1,
                    2,
                    3
                  ],
                  "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
                  "memory.current": "9617408",
                  "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.max": "17179869184",
                  "memory.peak": "9617408",
                  "memory.swap.max": "0",
                  "timestamp_ns": 1794684789951080908
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/constraints.json",
                "valid": true
              },
              "count_proof": {
                "hits": {
                  "field": "hits",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log",
                  "value": 3875208
                },
                "rows": {
                  "field": "rows",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log",
                  "value": 3933714
                }
              },
              "immkv_phase_statistics": {
                "format_seconds": {
                  "field": "format_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log",
                  "value": 7.06
                },
                "key_seconds": {
                  "field": "key_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log",
                  "value": 0.3
                },
                "query_seconds": {
                  "field": "query_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log",
                  "value": 4.04
                },
                "read_gib_rounded": {
                  "field": "read_gib_rounded",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log",
                  "value": 0.0
                },
                "wall_seconds": {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/immkv.log",
                  "value": 7.989
                }
              },
              "invalid_reasons": [],
              "issues": [],
              "metrics": {
                "cgroup_peak_bytes": {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/status.json",
                  "value": 529625088
                },
                "cpu_percent": {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt",
                  "value": 367.0
                },
                "exit_code": {
                  "field": "Exit status",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt",
                  "value": 0
                },
                "max_child_rss_kib": {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt",
                  "value": 423004
                },
                "system_seconds": {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt",
                  "value": 2.18
                },
                "user_seconds": {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt",
                  "value": 27.84
                },
                "wall_seconds": {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt",
                  "value": 8.16
                }
              },
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release",
              "round": 3,
              "status": {
                "directory_exists": true,
                "record": {
                  "exit_code": 0,
                  "finished_ns": 1794684798148682115,
                  "memory_peak_bytes": 529625088,
                  "oom_kill_detected": false,
                  "status": "ok"
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/status.json",
                "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/started",
                "state": "ok"
              },
              "valid": true
            }
          ],
          "summary": {
            "cgroup_peak_bytes": {
              "maximum": 529625088,
              "median": 526442496,
              "minimum": 520769536,
              "n": 3,
              "sources": [
                {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/status.json",
                  "value": 520769536
                },
                {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/status.json",
                  "value": 526442496
                },
                {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/status.json",
                  "value": 529625088
                }
              ],
              "values": [
                520769536,
                526442496,
                529625088
              ]
            },
            "cpu_percent": {
              "maximum": 367.0,
              "median": 367.0,
              "minimum": 354.0,
              "n": 3,
              "sources": [
                {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt",
                  "value": 354.0
                },
                {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt",
                  "value": 367.0
                },
                {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt",
                  "value": 367.0
                }
              ],
              "values": [
                354.0,
                367.0,
                367.0
              ]
            },
            "max_child_rss_kib": {
              "maximum": 423004,
              "median": 419928,
              "minimum": 413424,
              "n": 3,
              "sources": [
                {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt",
                  "value": 413424
                },
                {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt",
                  "value": 419928
                },
                {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt",
                  "value": 423004
                }
              ],
              "values": [
                413424,
                419928,
                423004
              ]
            },
            "system_seconds": {
              "maximum": 2.18,
              "median": 2.15,
              "minimum": 2.1,
              "n": 3,
              "sources": [
                {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt",
                  "value": 2.15
                },
                {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt",
                  "value": 2.1
                },
                {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt",
                  "value": 2.18
                }
              ],
              "values": [
                2.15,
                2.1,
                2.18
              ]
            },
            "user_seconds": {
              "maximum": 27.96,
              "median": 27.84,
              "minimum": 27.49,
              "n": 3,
              "sources": [
                {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt",
                  "value": 27.49
                },
                {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt",
                  "value": 27.96
                },
                {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt",
                  "value": 27.84
                }
              ],
              "values": [
                27.49,
                27.96,
                27.84
              ]
            },
            "wall_seconds": {
              "maximum": 8.37,
              "median": 8.17,
              "minimum": 8.16,
              "n": 3,
              "sources": [
                {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt",
                  "value": 8.37
                },
                {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt",
                  "value": 8.17
                },
                {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt",
                  "value": 8.16
                }
              ],
              "values": [
                8.37,
                8.17,
                8.16
              ]
            }
          },
          "valid_rounds": 3
        }
      },
      "fields": 17,
      "media": "nvme",
      "paired_wall_ratios": {
        "definition": "echtvar wall / immkv wall, paired by round",
        "maximum": 3.0049019607843137,
        "median": 3.001223990208078,
        "minimum": 2.9175627240143376,
        "n": 3,
        "range_over_median": 0.02910120572637459,
        "rounds": [
          {
            "denominator": {
              "field": "wall_seconds",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r1-release/time.txt",
              "value": 8.37
            },
            "numerator": {
              "field": "wall_seconds",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r1-release/time.txt",
              "value": 24.42
            },
            "round": 1,
            "value": 2.9175627240143376
          },
          {
            "denominator": {
              "field": "wall_seconds",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r2-release/time.txt",
              "value": 8.17
            },
            "numerator": {
              "field": "wall_seconds",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r2-release/time.txt",
              "value": 24.52
            },
            "round": 2,
            "value": 3.001223990208078
          },
          {
            "denominator": {
              "field": "wall_seconds",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-immkv-r3-release/time.txt",
              "value": 8.16
            },
            "numerator": {
              "field": "wall_seconds",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-17-echtvar-r3-release/time.txt",
              "value": 24.52
            },
            "round": 3,
            "value": 3.0049019607843137
          }
        ]
      }
    },
    {
      "correctness": {
        "fresh_echtvar": {
          "constraints": {
            "mismatches": [],
            "observed": {
              "GOMAXPROCS": "4",
              "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-verify-echtvar-nvme-540-release-1794685799-91662.scope",
              "cpu.stat": "usage_usec 30854\nuser_usec 22892\nsystem_usec 7962\nnice_usec 0\ncore_sched.force_idle_usec 0",
              "cpus": [
                0,
                1,
                2,
                3
              ],
              "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
              "memory.current": "9682944",
              "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
              "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
              "memory.max": "17179869184",
              "memory.peak": "9682944",
              "memory.swap.max": "0",
              "timestamp_ns": 1794685800010292943
            },
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-540-release/constraints.json",
            "valid": true
          },
          "evidence": {
            "chromosomes": [
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr1.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl1.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an1.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60213517147958,
                  "time_ns": 1794686036707472312
                },
                "annotation_started": {
                  "monotonic_ns": 59978069300611,
                  "time_ns": 1794685801259625176
                },
                "chromosome": 1,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "59ef94490a873f8ceb5139499102b66658a2c22d25a4e1c4fdc5166cb04bcdea",
                    "rows": 308298
                  },
                  "reference": {
                    "data_sha256": "59ef94490a873f8ceb5139499102b66658a2c22d25a4e1c4fdc5166cb04bcdea",
                    "rows": 308298
                  }
                },
                "finished": {
                  "monotonic_ns": 60213523918677,
                  "time_ns": 1794686036714243141
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "1",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl1.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 10999771,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 59978068676555,
                  "time_ns": 1794685801259000779
                },
                "slice_started": {
                  "monotonic_ns": 59976865617832,
                  "time_ns": 1794685800055942226
                },
                "started": {
                  "monotonic_ns": 59976865269055,
                  "time_ns": 1794685800055593329
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr2.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl2.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an2.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60226951375671,
                  "time_ns": 1794686050141699494
                },
                "annotation_started": {
                  "monotonic_ns": 59978074330122,
                  "time_ns": 1794685801264654736
                },
                "chromosome": 2,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "e4bdc61476b8d46eb502e0b45c3393067cec0e6b2e847b9d0e688b6aa4603c40",
                    "rows": 329821
                  },
                  "reference": {
                    "data_sha256": "e4bdc61476b8d46eb502e0b45c3393067cec0e6b2e847b9d0e688b6aa4603c40",
                    "rows": 329821
                  }
                },
                "finished": {
                  "monotonic_ns": 60226955317460,
                  "time_ns": 1794686050145641804
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "2",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl2.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 11750290,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 59978070169519,
                  "time_ns": 1794685801260494134
                },
                "slice_started": {
                  "monotonic_ns": 59976866134105,
                  "time_ns": 1794685800056458519
                },
                "started": {
                  "monotonic_ns": 59976865752045,
                  "time_ns": 1794685800056076529
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr3.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl3.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an3.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60185448028221,
                  "time_ns": 1794686008638352404
                },
                "annotation_started": {
                  "monotonic_ns": 59977870272607,
                  "time_ns": 1794685801060597311
                },
                "chromosome": 3,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "d8baaf493daa7bcaeab011f8d6f3f3e2b7c1cf3cb994d5edf14b1e2c77e6de25",
                    "rows": 275651
                  },
                  "reference": {
                    "data_sha256": "d8baaf493daa7bcaeab011f8d6f3f3e2b7c1cf3cb994d5edf14b1e2c77e6de25",
                    "rows": 275651
                  }
                },
                "finished": {
                  "monotonic_ns": 60185450850578,
                  "time_ns": 1794686008641174992
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "3",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl3.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 9815086,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 59977869739913,
                  "time_ns": 1794685801060063856
                },
                "slice_started": {
                  "monotonic_ns": 59976866858140,
                  "time_ns": 1794685800057182815
                },
                "started": {
                  "monotonic_ns": 59976866408823,
                  "time_ns": 1794685800056733327
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr4.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl4.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an4.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60197305801714,
                  "time_ns": 1794686020496125638
                },
                "annotation_started": {
                  "monotonic_ns": 59977973111441,
                  "time_ns": 1794685801163436095
                },
                "chromosome": 4,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "b0acf0f50c7a6a6abfda371fb79763cb60f506cb517699eaf71b603d25b6b1be",
                    "rows": 296434
                  },
                  "reference": {
                    "data_sha256": "b0acf0f50c7a6a6abfda371fb79763cb60f506cb517699eaf71b603d25b6b1be",
                    "rows": 296434
                  }
                },
                "finished": {
                  "monotonic_ns": 60197316128410,
                  "time_ns": 1794686020506452653
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "4",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl4.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 10519954,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 59977968617340,
                  "time_ns": 1794685801158941614
                },
                "slice_started": {
                  "monotonic_ns": 59976867201247,
                  "time_ns": 1794685800057525821
                },
                "started": {
                  "monotonic_ns": 59976866997803,
                  "time_ns": 1794685800057322277
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr5.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl5.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an5.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60368941157374,
                  "time_ns": 1794686192131481558
                },
                "annotation_started": {
                  "monotonic_ns": 60186563782881,
                  "time_ns": 1794686009754107545
                },
                "chromosome": 5,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "3ffd54981a331bfd12e5913e58afb4af35f9f9eb94ae652d229b5303055e4f31",
                    "rows": 253843
                  },
                  "reference": {
                    "data_sha256": "3ffd54981a331bfd12e5913e58afb4af35f9f9eb94ae652d229b5303055e4f31",
                    "rows": 253843
                  }
                },
                "finished": {
                  "monotonic_ns": 60368945155121,
                  "time_ns": 1794686192135479525
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "5",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl5.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 9032175,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60186563476955,
                  "time_ns": 1794686009753800899
                },
                "slice_started": {
                  "monotonic_ns": 60185452023934,
                  "time_ns": 1794686008642348659
                },
                "started": {
                  "monotonic_ns": 60185451901724,
                  "time_ns": 1794686008642226228
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr6.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl6.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an6.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60381987387110,
                  "time_ns": 1794686205177711333
                },
                "annotation_started": {
                  "monotonic_ns": 60198326955843,
                  "time_ns": 1794686021517280437
                },
                "chromosome": 6,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "6777b4b677680e6a2c5ff730a17947644bf36e14b079268ee9f93c8cdd85f81f",
                    "rows": 264414
                  },
                  "reference": {
                    "data_sha256": "6777b4b677680e6a2c5ff730a17947644bf36e14b079268ee9f93c8cdd85f81f",
                    "rows": 264414
                  }
                },
                "finished": {
                  "monotonic_ns": 60381990448808,
                  "time_ns": 1794686205180773282
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "6",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl6.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 8584380,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60198326645990,
                  "time_ns": 1794686021516970104
                },
                "slice_started": {
                  "monotonic_ns": 60197316554882,
                  "time_ns": 1794686020506879616
                },
                "started": {
                  "monotonic_ns": 60197316324047,
                  "time_ns": 1794686020506648792
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr7.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl7.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an7.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60375981197353,
                  "time_ns": 1794686199171521506
                },
                "annotation_started": {
                  "monotonic_ns": 60214433447966,
                  "time_ns": 1794686037623772570
                },
                "chromosome": 7,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "b8d05e8429a1fd1d9cc85ca0293b1887a0cc26838a7a4afe31598cf79b028494",
                    "rows": 222733
                  },
                  "reference": {
                    "data_sha256": "b8d05e8429a1fd1d9cc85ca0293b1887a0cc26838a7a4afe31598cf79b028494",
                    "rows": 222733
                  }
                },
                "finished": {
                  "monotonic_ns": 60375985356622,
                  "time_ns": 1794686199175681066
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "7",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl7.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 7924912,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60214432305138,
                  "time_ns": 1794686037622629361
                },
                "slice_started": {
                  "monotonic_ns": 60213525743607,
                  "time_ns": 1794686036716068282
                },
                "started": {
                  "monotonic_ns": 60213525027892,
                  "time_ns": 1794686036715352597
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr8.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl8.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an8.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60381817108219,
                  "time_ns": 1794686205007432182
                },
                "annotation_started": {
                  "monotonic_ns": 60227860024246,
                  "time_ns": 1794686051050348941
                },
                "chromosome": 8,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "222c4768dd63f7ac4d2394c04ce44287ceff2703c1ea5670318509aa7388ec12",
                    "rows": 213367
                  },
                  "reference": {
                    "data_sha256": "222c4768dd63f7ac4d2394c04ce44287ceff2703c1ea5670318509aa7388ec12",
                    "rows": 213367
                  }
                },
                "finished": {
                  "monotonic_ns": 60381824113592,
                  "time_ns": 1794686205014437946
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "8",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl8.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 7556433,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60227859629484,
                  "time_ns": 1794686051049953347
                },
                "slice_started": {
                  "monotonic_ns": 60226956090804,
                  "time_ns": 1794686050146415569
                },
                "started": {
                  "monotonic_ns": 60226955554165,
                  "time_ns": 1794686050145878930
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr9.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl9.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an9.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60489694515719,
                  "time_ns": 1794686312884839913
                },
                "annotation_started": {
                  "monotonic_ns": 60369752958186,
                  "time_ns": 1794686192943282891
                },
                "chromosome": 9,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "2eb2612e426858cc861337211691e8e303f3a181950c3e9c0c7a388c7cdddfb2",
                    "rows": 168542
                  },
                  "reference": {
                    "data_sha256": "2eb2612e426858cc861337211691e8e303f3a181950c3e9c0c7a388c7cdddfb2",
                    "rows": 168542
                  }
                },
                "finished": {
                  "monotonic_ns": 60489697371127,
                  "time_ns": 1794686312887694980
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "9",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl9.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 5978980,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60369752694612,
                  "time_ns": 1794686192943018505
                },
                "slice_started": {
                  "monotonic_ns": 60368945549822,
                  "time_ns": 1794686192135874476
                },
                "started": {
                  "monotonic_ns": 60368945403207,
                  "time_ns": 1794686192135727841
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr10.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl10.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an10.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60521040171064,
                  "time_ns": 1794686344230495128
                },
                "annotation_started": {
                  "monotonic_ns": 60376794044130,
                  "time_ns": 1794686199984368574
                },
                "chromosome": 10,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "258bd49bb17c322c57f1b8252de26f36133f81e918991e2fd1a1d3af07039651",
                    "rows": 198281
                  },
                  "reference": {
                    "data_sha256": "258bd49bb17c322c57f1b8252de26f36133f81e918991e2fd1a1d3af07039651",
                    "rows": 198281
                  }
                },
                "finished": {
                  "monotonic_ns": 60521042971488,
                  "time_ns": 1794686344233296032
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "10",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl10.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 7017518,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60376792234272,
                  "time_ns": 1794686199982558445
                },
                "slice_started": {
                  "monotonic_ns": 60375987170187,
                  "time_ns": 1794686199177494661
                },
                "started": {
                  "monotonic_ns": 60375987034753,
                  "time_ns": 1794686199177359287
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr11.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl11.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an11.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60530833787072,
                  "time_ns": 1794686354024110985
                },
                "annotation_started": {
                  "monotonic_ns": 60382629434130,
                  "time_ns": 1794686205819758775
                },
                "chromosome": 11,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "849c192386068b10f5da434d41a878d9813290486fa267a6c28997e147d82706",
                    "rows": 200574
                  },
                  "reference": {
                    "data_sha256": "849c192386068b10f5da434d41a878d9813290486fa267a6c28997e147d82706",
                    "rows": 200574
                  }
                },
                "finished": {
                  "monotonic_ns": 60530840209640,
                  "time_ns": 1794686354030534315
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "11",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl11.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 7109132,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60382629186726,
                  "time_ns": 1794686205819510669
                },
                "slice_started": {
                  "monotonic_ns": 60381826009982,
                  "time_ns": 1794686205016334557
                },
                "started": {
                  "monotonic_ns": 60381825843510,
                  "time_ns": 1794686205016168044
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr12.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl12.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an12.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60524136591196,
                  "time_ns": 1794686347326915289
                },
                "annotation_started": {
                  "monotonic_ns": 60382796963409,
                  "time_ns": 1794686205987287622
                },
                "chromosome": 12,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "8e8ad77655aecf5b143da65c43e1eaccc11a749fe5c12611490c7096655f0cf9",
                    "rows": 189356
                  },
                  "reference": {
                    "data_sha256": "8e8ad77655aecf5b143da65c43e1eaccc11a749fe5c12611490c7096655f0cf9",
                    "rows": 189356
                  }
                },
                "finished": {
                  "monotonic_ns": 60524141814916,
                  "time_ns": 1794686347332139390
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "12",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl12.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 6749307,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60382796511530,
                  "time_ns": 1794686205986835373
                },
                "slice_started": {
                  "monotonic_ns": 60381990757438,
                  "time_ns": 1794686205181082182
                },
                "started": {
                  "monotonic_ns": 60381990623516,
                  "time_ns": 1794686205180948301
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr13.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl13.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an13.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60599369222246,
                  "time_ns": 1794686422559546520
                },
                "annotation_started": {
                  "monotonic_ns": 60490408340337,
                  "time_ns": 1794686313598664991
                },
                "chromosome": 13,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "ce4d59d8da2af38e2db8933f9ccebdcf63c45e91b2f6ac2558c9adb0505266bf",
                    "rows": 156939
                  },
                  "reference": {
                    "data_sha256": "ce4d59d8da2af38e2db8933f9ccebdcf63c45e91b2f6ac2558c9adb0505266bf",
                    "rows": 156939
                  }
                },
                "finished": {
                  "monotonic_ns": 60599372025824,
                  "time_ns": 1794686422562349727
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "13",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl13.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 5531589,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60490405216786,
                  "time_ns": 1794686313595540839
                },
                "slice_started": {
                  "monotonic_ns": 60489698585044,
                  "time_ns": 1794686312888909689
                },
                "started": {
                  "monotonic_ns": 60489697569539,
                  "time_ns": 1794686312887894094
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr14.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl14.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an14.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60613663567966,
                  "time_ns": 1794686436853892070
                },
                "annotation_started": {
                  "monotonic_ns": 60521649645055,
                  "time_ns": 1794686344839969740
                },
                "chromosome": 14,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "ca822fb4d18bfe4626577736dbae9831c8e8de44d71f22a2b69cbcf7ea1db381",
                    "rows": 125631
                  },
                  "reference": {
                    "data_sha256": "ca822fb4d18bfe4626577736dbae9831c8e8de44d71f22a2b69cbcf7ea1db381",
                    "rows": 125631
                  }
                },
                "finished": {
                  "monotonic_ns": 60613669244221,
                  "time_ns": 1794686436859568535
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "14",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl14.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 4486825,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60521649349641,
                  "time_ns": 1794686344839673624
                },
                "slice_started": {
                  "monotonic_ns": 60521043266802,
                  "time_ns": 1794686344233591547
                },
                "started": {
                  "monotonic_ns": 60521043163778,
                  "time_ns": 1794686344233488523
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr15.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl15.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an15.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60609324732844,
                  "time_ns": 1794686432515056838
                },
                "annotation_started": {
                  "monotonic_ns": 60524751195425,
                  "time_ns": 1794686347941519058
                },
                "chromosome": 15,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "de1808e7b7c482288acb0672139756800e4ea47120a9590d8cc06ef4ad4838d7",
                    "rows": 114117
                  },
                  "reference": {
                    "data_sha256": "de1808e7b7c482288acb0672139756800e4ea47120a9590d8cc06ef4ad4838d7",
                    "rows": 114117
                  }
                },
                "finished": {
                  "monotonic_ns": 60609328353404,
                  "time_ns": 1794686432518677718
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "15",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl15.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 4074508,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60524747289587,
                  "time_ns": 1794686347937613681
                },
                "slice_started": {
                  "monotonic_ns": 60524142133323,
                  "time_ns": 1794686347332457998
                },
                "started": {
                  "monotonic_ns": 60524141989704,
                  "time_ns": 1794686347332314459
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr16.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl16.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an16.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60617904692483,
                  "time_ns": 1794686441095016456
                },
                "annotation_started": {
                  "monotonic_ns": 60531449723511,
                  "time_ns": 1794686354640048156
                },
                "chromosome": 16,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "ce19990fefc7c57c5241cb715ddce8d75294c8044391873945cabcc2b5e5c243",
                    "rows": 118437
                  },
                  "reference": {
                    "data_sha256": "ce19990fefc7c57c5241cb715ddce8d75294c8044391873945cabcc2b5e5c243",
                    "rows": 118437
                  }
                },
                "finished": {
                  "monotonic_ns": 60617908251206,
                  "time_ns": 1794686441098575580
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "16",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl16.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 4212957,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60531449492177,
                  "time_ns": 1794686354639816311
                },
                "slice_started": {
                  "monotonic_ns": 60530840510004,
                  "time_ns": 1794686354030834689
                },
                "started": {
                  "monotonic_ns": 60530840365332,
                  "time_ns": 1794686354030690107
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr17.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl17.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an17.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60675360863979,
                  "time_ns": 1794686498551187812
                },
                "annotation_started": {
                  "monotonic_ns": 60599879383745,
                  "time_ns": 1794686423069708059
                },
                "chromosome": 17,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "6b1ab9961447b7b87d9dd796602b7b1a8aec05fe1bc838ef241932cf72851e8e",
                    "rows": 102461
                  },
                  "reference": {
                    "data_sha256": "6b1ab9961447b7b87d9dd796602b7b1a8aec05fe1bc838ef241932cf72851e8e",
                    "rows": 102461
                  }
                },
                "finished": {
                  "monotonic_ns": 60675363809992,
                  "time_ns": 1794686498554134597
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "17",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl17.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 3711657,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60599879143946,
                  "time_ns": 1794686423069467949
                },
                "slice_started": {
                  "monotonic_ns": 60599372298706,
                  "time_ns": 1794686422562623470
                },
                "started": {
                  "monotonic_ns": 60599372192546,
                  "time_ns": 1794686422562517071
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr18.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl18.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an18.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60691654748764,
                  "time_ns": 1794686514845072567
                },
                "annotation_started": {
                  "monotonic_ns": 60609933370620,
                  "time_ns": 1794686433123695194
                },
                "chromosome": 18,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "30cfca328b607afd97c5d13eb3b130b1788bd449d3ba5f2ae8fd5879748666a7",
                    "rows": 119403
                  },
                  "reference": {
                    "data_sha256": "30cfca328b607afd97c5d13eb3b130b1788bd449d3ba5f2ae8fd5879748666a7",
                    "rows": 119403
                  }
                },
                "finished": {
                  "monotonic_ns": 60691657301319,
                  "time_ns": 1794686514847625854
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "18",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl18.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 4239924,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60609932327754,
                  "time_ns": 1794686433122651958
                },
                "slice_started": {
                  "monotonic_ns": 60609328651183,
                  "time_ns": 1794686432518975957
                },
                "started": {
                  "monotonic_ns": 60609328537059,
                  "time_ns": 1794686432518861733
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr19.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl19.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an19.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60680132272346,
                  "time_ns": 1794686503322596520
                },
                "annotation_started": {
                  "monotonic_ns": 60614073366609,
                  "time_ns": 1794686437263690783
                },
                "chromosome": 19,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "559bdfc5144aabecdb11ef8d67ea8a5dd71e9b7b06ee57fd0be15f8214adc371",
                    "rows": 86442
                  },
                  "reference": {
                    "data_sha256": "559bdfc5144aabecdb11ef8d67ea8a5dd71e9b7b06ee57fd0be15f8214adc371",
                    "rows": 86442
                  }
                },
                "finished": {
                  "monotonic_ns": 60680134459777,
                  "time_ns": 1794686503324784171
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "19",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl19.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 3140430,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60614070204760,
                  "time_ns": 1794686437260528944
                },
                "slice_started": {
                  "monotonic_ns": 60613669576975,
                  "time_ns": 1794686436859901720
                },
                "started": {
                  "monotonic_ns": 60613669438776,
                  "time_ns": 1794686436859763361
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr20.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl20.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an20.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60679850403122,
                  "time_ns": 1794686503040727035
                },
                "annotation_started": {
                  "monotonic_ns": 60618314403307,
                  "time_ns": 1794686441504728062
                },
                "chromosome": 20,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "784a6057e388e9ccce45a2366647193caa076f04c270283ee912eb24cf4feea9",
                    "rows": 83452
                  },
                  "reference": {
                    "data_sha256": "784a6057e388e9ccce45a2366647193caa076f04c270283ee912eb24cf4feea9",
                    "rows": 83452
                  }
                },
                "finished": {
                  "monotonic_ns": 60679852698545,
                  "time_ns": 1794686503043022929
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "20",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl20.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 2951322,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60618313989662,
                  "time_ns": 1794686441504313845
                },
                "slice_started": {
                  "monotonic_ns": 60617908587657,
                  "time_ns": 1794686441098912432
                },
                "started": {
                  "monotonic_ns": 60617908451792,
                  "time_ns": 1794686441098776467
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr21.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl21.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an21.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60708600372198,
                  "time_ns": 1794686531790696332
                },
                "annotation_started": {
                  "monotonic_ns": 60675766634401,
                  "time_ns": 1794686498956958925
                },
                "chromosome": 21,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "c67c4a591c15d64d9c4bf5352c88028a5e07c916354e54025297c6103f9c4407",
                    "rows": 55978
                  },
                  "reference": {
                    "data_sha256": "c67c4a591c15d64d9c4bf5352c88028a5e07c916354e54025297c6103f9c4407",
                    "rows": 55978
                  }
                },
                "finished": {
                  "monotonic_ns": 60708603315075,
                  "time_ns": 1794686531793639499
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "21",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl21.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 1982474,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60675766413897,
                  "time_ns": 1794686498956737961
                },
                "slice_started": {
                  "monotonic_ns": 60675364307125,
                  "time_ns": 1794686498554631859
                },
                "started": {
                  "monotonic_ns": 60675364144159,
                  "time_ns": 1794686498554468513
                },
                "status": "ok"
              },
              {
                "annotation_command": [
                  "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                  "anno",
                  "-e",
                  "/home/tian/immkv_test_nvme/test/540_fields/archives/chr22.echtvar.zip",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl22.vcf.gz",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/an22.vcf.gz"
                ],
                "annotation_exit_code": 0,
                "annotation_finished": {
                  "monotonic_ns": 60709865367835,
                  "time_ns": 1794686533055691287
                },
                "annotation_started": {
                  "monotonic_ns": 60680155352908,
                  "time_ns": 1794686503345677302
                },
                "chromosome": 22,
                "comparison": {
                  "equal": true,
                  "fresh": {
                    "data_sha256": "5636e7ed9b75c0666497ee77366c60d2922c32c79d60c968cb519a06d2b936ac",
                    "rows": 49540
                  },
                  "reference": {
                    "data_sha256": "5636e7ed9b75c0666497ee77366c60d2922c32c79d60c968cb519a06d2b936ac",
                    "rows": 49540
                  }
                },
                "finished": {
                  "monotonic_ns": 60709869724061,
                  "time_ns": 1794686533060048235
                },
                "slice_command": [
                  "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                  "view",
                  "-r",
                  "22",
                  "-Oz",
                  "-o",
                  "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release/sl22.vcf.gz",
                  "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                ],
                "slice_compressed_bytes": 1765323,
                "slice_exit_code": 0,
                "slice_finished": {
                  "monotonic_ns": 60680155132064,
                  "time_ns": 1794686503345456308
                },
                "slice_started": {
                  "monotonic_ns": 60679852979742,
                  "time_ns": 1794686503043304397
                },
                "started": {
                  "monotonic_ns": 60679852873262,
                  "time_ns": 1794686503043197847
                },
                "status": "ok"
              }
            ],
            "compressed_output": true,
            "engine": "echtvar",
            "exit_code": 0,
            "expected_rows": 3933714,
            "fields": 540,
            "finished": {
              "monotonic_ns": 60709870599803,
              "time_ns": 1794686533060924447
            },
            "fresh_rows": 3933714,
            "media": "nvme",
            "scratch": "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/verify-echtvar-nvme-540-release",
            "started": {
              "monotonic_ns": 59976865033180,
              "time_ns": 1794685800055354068
            },
            "status": "ok",
            "worker_note": "Four chromosome subprocess tasks on CPUs 0-3; each echtvar process may create additional htslib threads.",
            "workers": 4
          },
          "issues": [],
          "note": "Correctness phase only; excluded from timing medians.",
          "passed": true,
          "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-540-release/cell.json",
          "status": {
            "directory_exists": true,
            "record": {
              "exit_code": 0,
              "finished_ns": 1794686533095304668,
              "memory_peak_bytes": 17179869184,
              "oom_kill_detected": false,
              "status": "ok"
            },
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-540-release/status.json",
            "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-540-release/started",
            "state": "ok"
          },
          "verifier_stdout": {}
        },
        "immkv": {
          "constraints": {
            "mismatches": [],
            "observed": {
              "GOMAXPROCS": "4",
              "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-verify-nvme-540-release-1794684824-91792.scope",
              "cpu.stat": "usage_usec 28734\nuser_usec 17835\nsystem_usec 10899\nnice_usec 0\ncore_sched.force_idle_usec 0",
              "cpus": [
                0,
                1,
                2,
                3
              ],
              "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
              "memory.current": "9801728",
              "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
              "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
              "memory.max": "17179869184",
              "memory.peak": "9801728",
              "memory.swap.max": "0",
              "timestamp_ns": 1794684824090967166
            },
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/constraints.json",
            "valid": true
          },
          "evidence": {
            "annotate_statistics": {
              "hits": 3875208,
              "rows": 3933714
            },
            "commands": {
              "annotate": [
                "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/annotate",
                "-in",
                "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz",
                "-lib",
                "/home/tian/immkv_test_nvme/gnomad/immkv",
                "-fields",
                "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_540_fields/fields540.tsv",
                "-threads",
                "4",
                "-batch",
                "4096",
                "-out",
                "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-540-release/output.vcf.gz"
              ],
              "verifier": [
                "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/verifyvalues",
                "-ev",
                "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-540-release/echtvar.fifo",
                "-iv",
                "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-540-release/immkv.fifo",
                "-fields",
                "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_540_fields/fields540.tsv",
                "-tol",
                "1.1e-6",
                "-max-report",
                "20"
              ]
            },
            "comparison": {
              "echtvar_i32_overflow_losses": 808,
              "field_comparisons": 2124205560,
              "hit_disagreements": 0,
              "misses_agreed": 32050262,
              "rows": 3933714,
              "value_mismatches": 0
            },
            "elapsed_seconds_correctness_only": 975.7349105379981,
            "exit_status": {
              "annotate": 0,
              "verifier": 0
            },
            "expected_hits": 3875208,
            "expected_rows": 3933714,
            "fields": 540,
            "fields_tsv": "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_540_fields/fields540.tsv",
            "float_tolerance": 1.1e-06,
            "float_tolerance_basis": "echtvar v0.2.4 quantization truncation plus six-decimal print truncation and f32 slack",
            "library": "/home/tian/immkv_test_nvme/gnomad/immkv",
            "marker_collision_accounting": "Shared verifier includes true float -1 marker collisions in misses agreed.",
            "media": "nvme",
            "query": "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz",
            "reference_kind": "historical correctness reference; not a current performance sample",
            "reference_media": "nvme",
            "reference_release": "echtvar v0.2.4",
            "sha256": {
              "/home/tian/projects/immkv/test/vs_echtvar_hg19_gnomad_genomes_540_fields/fields540.tsv": "fa1f21636fcd186308e3d070a65b0c087a8813f6763a78f8db70dc96c65f7643",
              "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify_output.py": "5ff9a41ae7be63fbe25c78ecb16519f3b3cec517953792d68e4552325243467a",
              "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/annotate": "4179f6f9bf774615ce68e1794dd3c46f2c5e4b2316011b117191584f5ec16391",
              "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/verifyvalues": "8e307e09177310a07ec9615734493f9e423a779f3b6f3a2a5dd82f4364343400"
            },
            "status": "passed",
            "streams": {
              "echtvar": {
                "data_sha256": "37d131239563d193d5f802dbff870547bbeb4cad795951b2dd71f7fe2a618bf4",
                "files": [
                  {
                    "data_bytes": 5358350269,
                    "data_sha256": "59ef94490a873f8ceb5139499102b66658a2c22d25a4e1c4fdc5166cb04bcdea",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an1.vcf.gz",
                    "rows": 308298
                  },
                  {
                    "data_bytes": 5732240580,
                    "data_sha256": "e4bdc61476b8d46eb502e0b45c3393067cec0e6b2e847b9d0e688b6aa4603c40",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an2.vcf.gz",
                    "rows": 329821
                  },
                  {
                    "data_bytes": 4792633475,
                    "data_sha256": "d8baaf493daa7bcaeab011f8d6f3f3e2b7c1cf3cb994d5edf14b1e2c77e6de25",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an3.vcf.gz",
                    "rows": 275651
                  },
                  {
                    "data_bytes": 5153980926,
                    "data_sha256": "b0acf0f50c7a6a6abfda371fb79763cb60f506cb517699eaf71b603d25b6b1be",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an4.vcf.gz",
                    "rows": 296434
                  },
                  {
                    "data_bytes": 4412841424,
                    "data_sha256": "3ffd54981a331bfd12e5913e58afb4af35f9f9eb94ae652d229b5303055e4f31",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an5.vcf.gz",
                    "rows": 253843
                  },
                  {
                    "data_bytes": 4574212905,
                    "data_sha256": "6777b4b677680e6a2c5ff730a17947644bf36e14b079268ee9f93c8cdd85f81f",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an6.vcf.gz",
                    "rows": 264414
                  },
                  {
                    "data_bytes": 3871325750,
                    "data_sha256": "b8d05e8429a1fd1d9cc85ca0293b1887a0cc26838a7a4afe31598cf79b028494",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an7.vcf.gz",
                    "rows": 222733
                  },
                  {
                    "data_bytes": 3710141684,
                    "data_sha256": "222c4768dd63f7ac4d2394c04ce44287ceff2703c1ea5670318509aa7388ec12",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an8.vcf.gz",
                    "rows": 213367
                  },
                  {
                    "data_bytes": 2929629765,
                    "data_sha256": "2eb2612e426858cc861337211691e8e303f3a181950c3e9c0c7a388c7cdddfb2",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an9.vcf.gz",
                    "rows": 168542
                  },
                  {
                    "data_bytes": 3446687952,
                    "data_sha256": "258bd49bb17c322c57f1b8252de26f36133f81e918991e2fd1a1d3af07039651",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an10.vcf.gz",
                    "rows": 198281
                  },
                  {
                    "data_bytes": 3486144839,
                    "data_sha256": "849c192386068b10f5da434d41a878d9813290486fa267a6c28997e147d82706",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an11.vcf.gz",
                    "rows": 200574
                  },
                  {
                    "data_bytes": 3290577146,
                    "data_sha256": "8e8ad77655aecf5b143da65c43e1eaccc11a749fe5c12611490c7096655f0cf9",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an12.vcf.gz",
                    "rows": 189356
                  },
                  {
                    "data_bytes": 2727987517,
                    "data_sha256": "ce4d59d8da2af38e2db8933f9ccebdcf63c45e91b2f6ac2558c9adb0505266bf",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an13.vcf.gz",
                    "rows": 156939
                  },
                  {
                    "data_bytes": 2183774964,
                    "data_sha256": "ca822fb4d18bfe4626577736dbae9831c8e8de44d71f22a2b69cbcf7ea1db381",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an14.vcf.gz",
                    "rows": 125631
                  },
                  {
                    "data_bytes": 1983196172,
                    "data_sha256": "de1808e7b7c482288acb0672139756800e4ea47120a9590d8cc06ef4ad4838d7",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an15.vcf.gz",
                    "rows": 114117
                  },
                  {
                    "data_bytes": 2058578699,
                    "data_sha256": "ce19990fefc7c57c5241cb715ddce8d75294c8044391873945cabcc2b5e5c243",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an16.vcf.gz",
                    "rows": 118437
                  },
                  {
                    "data_bytes": 1780261584,
                    "data_sha256": "6b1ab9961447b7b87d9dd796602b7b1a8aec05fe1bc838ef241932cf72851e8e",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an17.vcf.gz",
                    "rows": 102461
                  },
                  {
                    "data_bytes": 2075298740,
                    "data_sha256": "30cfca328b607afd97c5d13eb3b130b1788bd449d3ba5f2ae8fd5879748666a7",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an18.vcf.gz",
                    "rows": 119403
                  },
                  {
                    "data_bytes": 1501634495,
                    "data_sha256": "559bdfc5144aabecdb11ef8d67ea8a5dd71e9b7b06ee57fd0be15f8214adc371",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an19.vcf.gz",
                    "rows": 86442
                  },
                  {
                    "data_bytes": 1450566921,
                    "data_sha256": "784a6057e388e9ccce45a2366647193caa076f04c270283ee912eb24cf4feea9",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an20.vcf.gz",
                    "rows": 83452
                  },
                  {
                    "data_bytes": 972937987,
                    "data_sha256": "c67c4a591c15d64d9c4bf5352c88028a5e07c916354e54025297c6103f9c4407",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an21.vcf.gz",
                    "rows": 55978
                  },
                  {
                    "data_bytes": 860563548,
                    "data_sha256": "5636e7ed9b75c0666497ee77366c60d2922c32c79d60c968cb519a06d2b936ac",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/immkv_test_nvme/test/4c16g_gz_v003/540_fields/echtvar.pass2.wk/an22.vcf.gz",
                    "rows": 49540
                  }
                ],
                "rows": 3933714
              },
              "immkv": {
                "data_sha256": "9e54e963a7a9db6063232afa59a58c48f37d999289116a43e16f97bfccd59dcd",
                "files": [
                  {
                    "data_bytes": 67477268158,
                    "data_sha256": "9e54e963a7a9db6063232afa59a58c48f37d999289116a43e16f97bfccd59dcd",
                    "gzip_crc_eof_ok": true,
                    "path": "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify/verify-nvme-540-release/output.vcf.gz",
                    "rows": 3933714
                  }
                ],
                "rows": 3933714,
                "selected_float_true_minus_one_candidates": 663
              }
            },
            "timing_sample": false
          },
          "issues": [],
          "note": "Correctness phase only; excluded from timing medians.",
          "passed": true,
          "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verification.json",
          "status": {
            "directory_exists": true,
            "record": {
              "exit_code": 0,
              "finished_ns": 1794685799898830770,
              "memory_peak_bytes": 14087831552,
              "oom_kill_detected": false,
              "status": "ok"
            },
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/status.json",
            "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/started",
            "state": "ok"
          },
          "verifier_stdout": {
            "echtvar_i32_overflow_losses": {
              "field": "echtvar_i32_overflow_losses",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout",
              "value": 808
            },
            "field_comparisons": {
              "field": "field_comparisons",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout",
              "value": 2124205560
            },
            "float_comparisons": {
              "field": "float_comparisons",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout",
              "value": 666144645
            },
            "float_tolerance": {
              "field": "float_tolerance",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout",
              "value": 1.1e-06
            },
            "hit_disagreements": {
              "field": "hit_disagreements",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout",
              "value": 0
            },
            "max_float_absolute_difference": {
              "field": "max_float_absolute_difference",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout",
              "value": 1e-06
            },
            "misses_agreed": {
              "field": "misses_agreed",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout",
              "value": 32050262
            },
            "rows": {
              "field": "rows",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout",
              "value": 3933714
            },
            "value_mismatches": {
              "field": "value_mismatches",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout",
              "value": 0
            },
            "verify_values_ok": {
              "field": "VERIFY_VALUES OK",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-540-release/verifier.stdout",
              "value": true
            }
          }
        }
      },
      "engines": {
        "echtvar": {
          "rounds": [
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": {
                  "GOMAXPROCS": "4",
                  "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-timing-nvme-540-echtvar-r1-release-1794686691-92423.scope",
                  "cpu.stat": "usage_usec 27784\nuser_usec 18853\nsystem_usec 8930\nnice_usec 0\ncore_sched.force_idle_usec 0",
                  "cpus": [
                    0,
                    1,
                    2,
                    3
                  ],
                  "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
                  "memory.current": "9707520",
                  "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.max": "17179869184",
                  "memory.peak": "9707520",
                  "memory.swap.max": "0",
                  "timestamp_ns": 1794686691091889804
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/constraints.json",
                "valid": true
              },
              "count_proof": {
                "all_22_chromosomes_ok": true,
                "annotation_log_statistics": [
                  {
                    "chromosome": 1,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr1.anno.log",
                      "value": 308298
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr1.anno.log",
                      "value": 308297
                    }
                  },
                  {
                    "chromosome": 2,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr2.anno.log",
                      "value": 329821
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr2.anno.log",
                      "value": 329821
                    }
                  },
                  {
                    "chromosome": 3,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr3.anno.log",
                      "value": 275651
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr3.anno.log",
                      "value": 275647
                    }
                  },
                  {
                    "chromosome": 4,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr4.anno.log",
                      "value": 296434
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr4.anno.log",
                      "value": 296434
                    }
                  },
                  {
                    "chromosome": 5,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr5.anno.log",
                      "value": 253843
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr5.anno.log",
                      "value": 253841
                    }
                  },
                  {
                    "chromosome": 6,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr6.anno.log",
                      "value": 264414
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr6.anno.log",
                      "value": 264084
                    }
                  },
                  {
                    "chromosome": 7,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr7.anno.log",
                      "value": 222733
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr7.anno.log",
                      "value": 222732
                    }
                  },
                  {
                    "chromosome": 8,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr8.anno.log",
                      "value": 213367
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr8.anno.log",
                      "value": 213367
                    }
                  },
                  {
                    "chromosome": 9,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr9.anno.log",
                      "value": 168542
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr9.anno.log",
                      "value": 168542
                    }
                  },
                  {
                    "chromosome": 10,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr10.anno.log",
                      "value": 198281
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr10.anno.log",
                      "value": 198281
                    }
                  },
                  {
                    "chromosome": 11,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr11.anno.log",
                      "value": 200574
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr11.anno.log",
                      "value": 200574
                    }
                  },
                  {
                    "chromosome": 12,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr12.anno.log",
                      "value": 189356
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr12.anno.log",
                      "value": 189356
                    }
                  },
                  {
                    "chromosome": 13,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr13.anno.log",
                      "value": 156939
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr13.anno.log",
                      "value": 156938
                    }
                  },
                  {
                    "chromosome": 14,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr14.anno.log",
                      "value": 125631
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr14.anno.log",
                      "value": 125631
                    }
                  },
                  {
                    "chromosome": 15,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr15.anno.log",
                      "value": 114117
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr15.anno.log",
                      "value": 114117
                    }
                  },
                  {
                    "chromosome": 16,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr16.anno.log",
                      "value": 118437
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr16.anno.log",
                      "value": 118437
                    }
                  },
                  {
                    "chromosome": 17,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr17.anno.log",
                      "value": 102461
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr17.anno.log",
                      "value": 102461
                    }
                  },
                  {
                    "chromosome": 18,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr18.anno.log",
                      "value": 119403
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr18.anno.log",
                      "value": 119401
                    }
                  },
                  {
                    "chromosome": 19,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr19.anno.log",
                      "value": 86442
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr19.anno.log",
                      "value": 86442
                    }
                  },
                  {
                    "chromosome": 20,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr20.anno.log",
                      "value": 83452
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr20.anno.log",
                      "value": 83452
                    }
                  },
                  {
                    "chromosome": 21,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr21.anno.log",
                      "value": 55978
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr21.anno.log",
                      "value": 55978
                    }
                  },
                  {
                    "chromosome": 22,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr22.anno.log",
                      "value": 49540
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr22.anno.log",
                      "value": 49536
                    }
                  }
                ],
                "chromosome_statuses": [
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr1.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl1.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61071844354881,
                      "time_ns": 1794686895034678995
                    },
                    "annotation_started": {
                      "monotonic_ns": 60869149499869,
                      "time_ns": 1794686692339824303
                    },
                    "chromosome": 1,
                    "finished": {
                      "monotonic_ns": 61071845771995,
                      "time_ns": 1794686895036096480
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "1",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl1.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 10999773,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 60869149189888,
                      "time_ns": 1794686692339514242
                    },
                    "slice_started": {
                      "monotonic_ns": 60867945261867,
                      "time_ns": 1794686691135586532
                    },
                    "started": {
                      "monotonic_ns": 60867944870363,
                      "time_ns": 1794686691135194707
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr2.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl2.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61083103192486,
                      "time_ns": 1794686906293516059
                    },
                    "annotation_started": {
                      "monotonic_ns": 60869149369314,
                      "time_ns": 1794686692339693979
                    },
                    "chromosome": 2,
                    "finished": {
                      "monotonic_ns": 61083104795849,
                      "time_ns": 1794686906295120313
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "2",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl2.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 11750292,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 60869149143271,
                      "time_ns": 1794686692339467304
                    },
                    "slice_started": {
                      "monotonic_ns": 60867945668619,
                      "time_ns": 1794686691135993374
                    },
                    "started": {
                      "monotonic_ns": 60867945401899,
                      "time_ns": 1794686691135726364
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr3.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl3.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61043474458119,
                      "time_ns": 1794686866664781952
                    },
                    "annotation_started": {
                      "monotonic_ns": 60868950828424,
                      "time_ns": 1794686692141153109
                    },
                    "chromosome": 3,
                    "finished": {
                      "monotonic_ns": 61043476132916,
                      "time_ns": 1794686866666457641
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "3",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl3.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 9815088,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 60868950689294,
                      "time_ns": 1794686692141013558
                    },
                    "slice_started": {
                      "monotonic_ns": 60867947866288,
                      "time_ns": 1794686691138190982
                    },
                    "started": {
                      "monotonic_ns": 60867945718994,
                      "time_ns": 1794686691136043478
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr4.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl4.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61048164880487,
                      "time_ns": 1794686871355204671
                    },
                    "annotation_started": {
                      "monotonic_ns": 60869050379861,
                      "time_ns": 1794686692240704375
                    },
                    "chromosome": 4,
                    "finished": {
                      "monotonic_ns": 61048166217060,
                      "time_ns": 1794686871356541745
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "4",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl4.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 10519956,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 60869050183703,
                      "time_ns": 1794686692240507717
                    },
                    "slice_started": {
                      "monotonic_ns": 60867947981203,
                      "time_ns": 1794686691138305938
                    },
                    "started": {
                      "monotonic_ns": 60867947774716,
                      "time_ns": 1794686691138099391
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr5.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl5.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61219750316470,
                      "time_ns": 1794687042940640564
                    },
                    "annotation_started": {
                      "monotonic_ns": 61044581889511,
                      "time_ns": 1794686867772214095
                    },
                    "chromosome": 5,
                    "finished": {
                      "monotonic_ns": 61219751585607,
                      "time_ns": 1794687042941909761
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "5",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl5.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 9032176,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61044581685929,
                      "time_ns": 1794686867772009622
                    },
                    "slice_started": {
                      "monotonic_ns": 61043476455180,
                      "time_ns": 1794686866666779924
                    },
                    "started": {
                      "monotonic_ns": 61043476313835,
                      "time_ns": 1794686866666638560
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr6.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl6.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61218535661547,
                      "time_ns": 1794687041725985430
                    },
                    "annotation_started": {
                      "monotonic_ns": 61049272633494,
                      "time_ns": 1794686872462958269
                    },
                    "chromosome": 6,
                    "finished": {
                      "monotonic_ns": 61218536913792,
                      "time_ns": 1794687041727238106
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "6",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl6.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 8584382,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61049272505464,
                      "time_ns": 1794686872462829588
                    },
                    "slice_started": {
                      "monotonic_ns": 61048166533443,
                      "time_ns": 1794686871356858097
                    },
                    "started": {
                      "monotonic_ns": 61048166415742,
                      "time_ns": 1794686871356740467
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr7.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl7.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61229665090322,
                      "time_ns": 1794687052855414436
                    },
                    "annotation_started": {
                      "monotonic_ns": 61072854319347,
                      "time_ns": 1794686896044643851
                    },
                    "chromosome": 7,
                    "finished": {
                      "monotonic_ns": 61229666105714,
                      "time_ns": 1794687052856429978
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "7",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl7.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7924914,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61072854180537,
                      "time_ns": 1794686896044504651
                    },
                    "slice_started": {
                      "monotonic_ns": 61071846201159,
                      "time_ns": 1794686895036525704
                    },
                    "started": {
                      "monotonic_ns": 61071846010883,
                      "time_ns": 1794686895036335377
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr8.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl8.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61228422430886,
                      "time_ns": 1794687051612754799
                    },
                    "annotation_started": {
                      "monotonic_ns": 61084013350072,
                      "time_ns": 1794686907203674577
                    },
                    "chromosome": 8,
                    "finished": {
                      "monotonic_ns": 61228423275377,
                      "time_ns": 1794687051613599601
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "8",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl8.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7556437,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61084013217774,
                      "time_ns": 1794686907203541848
                    },
                    "slice_started": {
                      "monotonic_ns": 61083105217829,
                      "time_ns": 1794686906295542454
                    },
                    "started": {
                      "monotonic_ns": 61083105015921,
                      "time_ns": 1794686906295340405
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr9.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl9.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61336092935700,
                      "time_ns": 1794687159283259984
                    },
                    "annotation_started": {
                      "monotonic_ns": 61219344898091,
                      "time_ns": 1794687042535222515
                    },
                    "chromosome": 9,
                    "finished": {
                      "monotonic_ns": 61336093570309,
                      "time_ns": 1794687159283895083
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "9",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl9.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 5978981,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61219344778537,
                      "time_ns": 1794687042535102420
                    },
                    "slice_started": {
                      "monotonic_ns": 61218537304594,
                      "time_ns": 1794687041727629168
                    },
                    "started": {
                      "monotonic_ns": 61218537150836,
                      "time_ns": 1794687041727475340
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr10.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl10.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61352221192344,
                      "time_ns": 1794687175411516317
                    },
                    "annotation_started": {
                      "monotonic_ns": 61220656381112,
                      "time_ns": 1794687043846705647
                    },
                    "chromosome": 10,
                    "finished": {
                      "monotonic_ns": 61352221929805,
                      "time_ns": 1794687175412254549
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "10",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl10.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7017519,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61220656248344,
                      "time_ns": 1794687043846572227
                    },
                    "slice_started": {
                      "monotonic_ns": 61219752013078,
                      "time_ns": 1794687042942337442
                    },
                    "started": {
                      "monotonic_ns": 61219751817161,
                      "time_ns": 1794687042942141595
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr11.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl11.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61363457011560,
                      "time_ns": 1794687186647335553
                    },
                    "annotation_started": {
                      "monotonic_ns": 61229429797216,
                      "time_ns": 1794687052620121800
                    },
                    "chromosome": 11,
                    "finished": {
                      "monotonic_ns": 61363457911855,
                      "time_ns": 1794687186648236460
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "11",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl11.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 7109133,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61229429669938,
                      "time_ns": 1794687052619993891
                    },
                    "slice_started": {
                      "monotonic_ns": 61228423674194,
                      "time_ns": 1794687051613998338
                    },
                    "started": {
                      "monotonic_ns": 61228423489208,
                      "time_ns": 1794687051613813301
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr12.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl12.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61358347952038,
                      "time_ns": 1794687181538276041
                    },
                    "annotation_started": {
                      "monotonic_ns": 61230469968614,
                      "time_ns": 1794687053660293128
                    },
                    "chromosome": 12,
                    "finished": {
                      "monotonic_ns": 61358348683147,
                      "time_ns": 1794687181539007942
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "12",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl12.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 6749308,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61230469842668,
                      "time_ns": 1794687053660166571
                    },
                    "slice_started": {
                      "monotonic_ns": 61229666520601,
                      "time_ns": 1794687052856845095
                    },
                    "started": {
                      "monotonic_ns": 61229666342848,
                      "time_ns": 1794687052856666791
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr13.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl13.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61429607515574,
                      "time_ns": 1794687252797839667
                    },
                    "annotation_started": {
                      "monotonic_ns": 61336800586321,
                      "time_ns": 1794687159990910885
                    },
                    "chromosome": 13,
                    "finished": {
                      "monotonic_ns": 61429608132539,
                      "time_ns": 1794687252798457284
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "13",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl13.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 5531591,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61336800370447,
                      "time_ns": 1794687159990693929
                    },
                    "slice_started": {
                      "monotonic_ns": 61336094017105,
                      "time_ns": 1794687159284341820
                    },
                    "started": {
                      "monotonic_ns": 61336093881682,
                      "time_ns": 1794687159284206386
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr14.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl14.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61430733183187,
                      "time_ns": 1794687253923507582
                    },
                    "annotation_started": {
                      "monotonic_ns": 61352824030620,
                      "time_ns": 1794687176014355385
                    },
                    "chromosome": 14,
                    "finished": {
                      "monotonic_ns": 61430733787859,
                      "time_ns": 1794687253924112654
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "14",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl14.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4486826,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61352823923058,
                      "time_ns": 1794687176014246942
                    },
                    "slice_started": {
                      "monotonic_ns": 61352222737567,
                      "time_ns": 1794687175413062312
                    },
                    "started": {
                      "monotonic_ns": 61352222128877,
                      "time_ns": 1794687175412453612
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr15.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl15.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61427434303313,
                      "time_ns": 1794687250624627587
                    },
                    "annotation_started": {
                      "monotonic_ns": 61358955185136,
                      "time_ns": 1794687182145509901
                    },
                    "chromosome": 15,
                    "finished": {
                      "monotonic_ns": 61427434814320,
                      "time_ns": 1794687250625139105
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "15",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl15.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4074509,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61358955083415,
                      "time_ns": 1794687182145407499
                    },
                    "slice_started": {
                      "monotonic_ns": 61358349007885,
                      "time_ns": 1794687181539332669
                    },
                    "started": {
                      "monotonic_ns": 61358348877681,
                      "time_ns": 1794687181539202426
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr16.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl16.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61437751953879,
                      "time_ns": 1794687260942278023
                    },
                    "annotation_started": {
                      "monotonic_ns": 61364065563501,
                      "time_ns": 1794687187255888306
                    },
                    "chromosome": 16,
                    "finished": {
                      "monotonic_ns": 61437752414121,
                      "time_ns": 1794687260942738906
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "16",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl16.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4212959,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61364065469626,
                      "time_ns": 1794687187255793789
                    },
                    "slice_started": {
                      "monotonic_ns": 61363458192971,
                      "time_ns": 1794687186648517726
                    },
                    "started": {
                      "monotonic_ns": 61363458082104,
                      "time_ns": 1794687186648406768
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr17.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl17.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61494206848326,
                      "time_ns": 1794687317397172139
                    },
                    "annotation_started": {
                      "monotonic_ns": 61427937314968,
                      "time_ns": 1794687251127639743
                    },
                    "chromosome": 17,
                    "finished": {
                      "monotonic_ns": 61494207743843,
                      "time_ns": 1794687317398068607
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "17",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl17.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 3711656,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61427937227404,
                      "time_ns": 1794687251127551458
                    },
                    "slice_started": {
                      "monotonic_ns": 61427435098502,
                      "time_ns": 1794687250625423277
                    },
                    "started": {
                      "monotonic_ns": 61427434978588,
                      "time_ns": 1794687250625303262
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr18.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl18.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61499718597846,
                      "time_ns": 1794687322908921960
                    },
                    "annotation_started": {
                      "monotonic_ns": 61430215615172,
                      "time_ns": 1794687253405939947
                    },
                    "chromosome": 18,
                    "finished": {
                      "monotonic_ns": 61499719070851,
                      "time_ns": 1794687322909395536
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "18",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl18.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 4239926,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61430215516237,
                      "time_ns": 1794687253405840140
                    },
                    "slice_started": {
                      "monotonic_ns": 61429611501065,
                      "time_ns": 1794687252801825820
                    },
                    "started": {
                      "monotonic_ns": 61429611374468,
                      "time_ns": 1794687252801699113
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr19.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl19.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61486684624849,
                      "time_ns": 1794687309874948952
                    },
                    "annotation_started": {
                      "monotonic_ns": 61431238833439,
                      "time_ns": 1794687254429158224
                    },
                    "chromosome": 19,
                    "finished": {
                      "monotonic_ns": 61486685073108,
                      "time_ns": 1794687309875397853
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "19",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl19.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 3140431,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61431238752627,
                      "time_ns": 1794687254429076701
                    },
                    "slice_started": {
                      "monotonic_ns": 61430734133136,
                      "time_ns": 1794687253924457831
                    },
                    "started": {
                      "monotonic_ns": 61430733974679,
                      "time_ns": 1794687253924299384
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr20.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl20.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61488406951738,
                      "time_ns": 1794687311597275611
                    },
                    "annotation_started": {
                      "monotonic_ns": 61438256159811,
                      "time_ns": 1794687261446484556
                    },
                    "chromosome": 20,
                    "finished": {
                      "monotonic_ns": 61488407305691,
                      "time_ns": 1794687311597630285
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "20",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl20.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 2951323,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61438256073550,
                      "time_ns": 1794687261446397764
                    },
                    "slice_started": {
                      "monotonic_ns": 61437752708312,
                      "time_ns": 1794687260943033067
                    },
                    "started": {
                      "monotonic_ns": 61437752569452,
                      "time_ns": 1794687260942894197
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr21.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl21.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61511457609910,
                      "time_ns": 1794687334647933854
                    },
                    "annotation_started": {
                      "monotonic_ns": 61487089639406,
                      "time_ns": 1794687310279964181
                    },
                    "chromosome": 21,
                    "finished": {
                      "monotonic_ns": 61511458140093,
                      "time_ns": 1794687334648464597
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "21",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl21.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 1982478,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61487089552954,
                      "time_ns": 1794687310279877058
                    },
                    "slice_started": {
                      "monotonic_ns": 61486685364714,
                      "time_ns": 1794687309875689479
                    },
                    "started": {
                      "monotonic_ns": 61486685256331,
                      "time_ns": 1794687309875581076
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr22.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl22.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61511869549499,
                      "time_ns": 1794687335059872982
                    },
                    "annotation_started": {
                      "monotonic_ns": 61488812165849,
                      "time_ns": 1794687312002490534
                    },
                    "chromosome": 22,
                    "finished": {
                      "monotonic_ns": 61511870273755,
                      "time_ns": 1794687335060598380
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "22",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r1-release/sl22.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 1765326,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61488812074688,
                      "time_ns": 1794687312002398612
                    },
                    "slice_started": {
                      "monotonic_ns": 61488407605172,
                      "time_ns": 1794687311597929936
                    },
                    "started": {
                      "monotonic_ns": 61488407489675,
                      "time_ns": 1794687311597814430
                    },
                    "status": "ok"
                  }
                ],
                "evaluated_total": {
                  "sources": [
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr1.anno.log",
                      "value": 308298
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr2.anno.log",
                      "value": 329821
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr3.anno.log",
                      "value": 275651
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr4.anno.log",
                      "value": 296434
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr5.anno.log",
                      "value": 253843
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr6.anno.log",
                      "value": 264414
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr7.anno.log",
                      "value": 222733
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr8.anno.log",
                      "value": 213367
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr9.anno.log",
                      "value": 168542
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr10.anno.log",
                      "value": 198281
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr11.anno.log",
                      "value": 200574
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr12.anno.log",
                      "value": 189356
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr13.anno.log",
                      "value": 156939
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr14.anno.log",
                      "value": 125631
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr15.anno.log",
                      "value": 114117
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr16.anno.log",
                      "value": 118437
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr17.anno.log",
                      "value": 102461
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr18.anno.log",
                      "value": 119403
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr19.anno.log",
                      "value": 86442
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr20.anno.log",
                      "value": 83452
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr21.anno.log",
                      "value": 55978
                    },
                    {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/chr22.anno.log",
                      "value": 49540
                    }
                  ],
                  "value": 3933714
                },
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [],
              "issues": [],
              "metrics": {
                "cgroup_peak_bytes": {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/status.json",
                  "value": 17179869184
                },
                "cpu_percent": {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/time.txt",
                  "value": 390.0
                },
                "exit_code": {
                  "field": "Exit status",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/time.txt",
                  "value": 0
                },
                "max_child_rss_kib": {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/time.txt",
                  "value": 540148
                },
                "system_seconds": {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/time.txt",
                  "value": 322.45
                },
                "user_seconds": {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/time.txt",
                  "value": 2190.82
                },
                "wall_seconds": {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/time.txt",
                  "value": 643.97
                }
              },
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release",
              "round": 1,
              "status": {
                "directory_exists": true,
                "record": {
                  "exit_code": 0,
                  "finished_ns": 1794687335098164046,
                  "memory_peak_bytes": 17179869184,
                  "oom_kill_detected": false,
                  "status": "ok"
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/status.json",
                "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r1-release/started",
                "state": "ok"
              },
              "valid": true
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": {
                  "GOMAXPROCS": "4",
                  "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-timing-nvme-540-echtvar-r2-release-1794687335-92423.scope",
                  "cpu.stat": "usage_usec 29755\nuser_usec 21820\nsystem_usec 7934\nnice_usec 0\ncore_sched.force_idle_usec 0",
                  "cpus": [
                    0,
                    1,
                    2,
                    3
                  ],
                  "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
                  "memory.current": "9904128",
                  "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.max": "17179869184",
                  "memory.peak": "9904128",
                  "memory.swap.max": "0",
                  "timestamp_ns": 1794687335219248758
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/constraints.json",
                "valid": true
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [
                  {
                    "chromosome": 1,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/chr1.anno.log",
                      "value": 308298
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/chr1.anno.log",
                      "value": 308297
                    }
                  },
                  {
                    "chromosome": 2,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/chr2.anno.log",
                      "value": 329821
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/chr2.anno.log",
                      "value": 329821
                    }
                  },
                  {
                    "chromosome": 3,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/chr3.anno.log",
                      "value": 275651
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/chr3.anno.log",
                      "value": 275647
                    }
                  },
                  {
                    "chromosome": 4,
                    "evaluated": {
                      "field": "evaluated variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/chr4.anno.log",
                      "value": 296434
                    },
                    "wrote": {
                      "field": "wrote variants",
                      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/chr4.anno.log",
                      "value": 296434
                    }
                  }
                ],
                "chromosome_statuses": [
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr1.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r2-release/sl1.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61708144025184,
                      "time_ns": 1794687531334349428
                    },
                    "annotation_started": {
                      "monotonic_ns": 61513274172491,
                      "time_ns": 1794687336464497145
                    },
                    "chromosome": 1,
                    "finished": {
                      "monotonic_ns": 61708145436378,
                      "time_ns": 1794687531335760933
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "1",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r2-release/sl1.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 10999773,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61513273948541,
                      "time_ns": 1794687336464272505
                    },
                    "slice_started": {
                      "monotonic_ns": 61512070521988,
                      "time_ns": 1794687335260846422
                    },
                    "started": {
                      "monotonic_ns": 61512070192171,
                      "time_ns": 1794687335260516424
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr2.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r2-release/sl2.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61721215098899,
                      "time_ns": 1794687544405422902
                    },
                    "annotation_started": {
                      "monotonic_ns": 61513275076473,
                      "time_ns": 1794687336465401258
                    },
                    "chromosome": 2,
                    "finished": {
                      "monotonic_ns": 61721216612718,
                      "time_ns": 1794687544406937303
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "2",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r2-release/sl2.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 11750292,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61513274933435,
                      "time_ns": 1794687336465257890
                    },
                    "slice_started": {
                      "monotonic_ns": 61512071435188,
                      "time_ns": 1794687335261759862
                    },
                    "started": {
                      "monotonic_ns": 61512070622947,
                      "time_ns": 1794687335260947441
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr3.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r2-release/sl3.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61678427449933,
                      "time_ns": 1794687501617774106
                    },
                    "annotation_started": {
                      "monotonic_ns": 61513175310873,
                      "time_ns": 1794687336365635307
                    },
                    "chromosome": 3,
                    "finished": {
                      "monotonic_ns": 61678428490331,
                      "time_ns": 1794687501618814986
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "3",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r2-release/sl3.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 9815088,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61513174965587,
                      "time_ns": 1794687336365289971
                    },
                    "slice_started": {
                      "monotonic_ns": 61512071687781,
                      "time_ns": 1794687335262012485
                    },
                    "started": {
                      "monotonic_ns": 61512071173057,
                      "time_ns": 1794687335261497541
                    },
                    "status": "ok"
                  },
                  {
                    "annotation_command": [
                      "/home/tian/projects/immkv/tmp/echtvar/echtvar",
                      "anno",
                      "-e",
                      "/home/tian/immkv_test_nvme/test/540_fields/archives/chr4.echtvar.zip",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r2-release/sl4.vcf.gz",
                      "/dev/null"
                    ],
                    "annotation_exit_code": 0,
                    "annotation_finished": {
                      "monotonic_ns": 61684676458910,
                      "time_ns": 1794687507866783124
                    },
                    "annotation_started": {
                      "monotonic_ns": 61513175174558,
                      "time_ns": 1794687336365499223
                    },
                    "chromosome": 4,
                    "finished": {
                      "monotonic_ns": 61684677562763,
                      "time_ns": 1794687507867887457
                    },
                    "slice_command": [
                      "/home/tian/miniforge3/envs/transvar/bin/bcftools",
                      "view",
                      "-r",
                      "4",
                      "-Oz",
                      "-o",
                      "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/runs/timing-nvme-540-echtvar-r2-release/sl4.vcf.gz",
                      "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz"
                    ],
                    "slice_compressed_bytes": 10519956,
                    "slice_exit_code": 0,
                    "slice_finished": {
                      "monotonic_ns": 61513174923608,
                      "time_ns": 1794687336365247612
                    },
                    "slice_started": {
                      "monotonic_ns": 61512071735610,
                      "time_ns": 1794687335262060154
                    },
                    "started": {
                      "monotonic_ns": 61512071540886,
                      "time_ns": 1794687335261865350
                    },
                    "status": "ok"
                  }
                ],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is started",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {
                "cgroup_peak_bytes": {
                  "field": "last observed memory.peak (may precede final peak)",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/memory-observer.json",
                  "value": 17179869184
                }
              },
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release",
              "round": 2,
              "status": {
                "directory_exists": true,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/status.json",
                "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r2-release/started",
                "state": "started"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r3-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r3-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r3-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r3-release",
              "round": 3,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-echtvar-r3-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            }
          ],
          "summary": null,
          "valid_rounds": 1
        },
        "immkv": {
          "rounds": [
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": {
                  "GOMAXPROCS": "4",
                  "cgroup": "/sys/fs/cgroup/user.slice/user-1000.slice/user@1000.service/app.slice/immkv-echtvar-timing-nvme-540-immkv-r1-release-1794686533-92423.scope",
                  "cpu.stat": "usage_usec 29182\nuser_usec 19119\nsystem_usec 10062\nnice_usec 0\ncore_sched.force_idle_usec 0",
                  "cpus": [
                    0,
                    1,
                    2,
                    3
                  ],
                  "gnu_time_rss": "Maximum resident set size is a maximum child RSS, not aggregate simultaneous RSS; use cgroup memory.peak for aggregate memory.",
                  "memory.current": "9596928",
                  "memory.events": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.events.local": "low 0\nhigh 0\nmax 0\noom 0\noom_kill 0\noom_group_kill 0\nsock_throttled 0",
                  "memory.max": "17179869184",
                  "memory.peak": "9596928",
                  "memory.swap.max": "0",
                  "timestamp_ns": 1794686533220214157
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/constraints.json",
                "valid": true
              },
              "count_proof": {
                "hits": {
                  "field": "hits",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/immkv.log",
                  "value": 3875208
                },
                "rows": {
                  "field": "rows",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/immkv.log",
                  "value": 3933714
                }
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [],
              "issues": [],
              "metrics": {
                "cgroup_peak_bytes": {
                  "field": "memory_peak_bytes",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/status.json",
                  "value": 2572652544
                },
                "cpu_percent": {
                  "field": "cpu_percent",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/time.txt",
                  "value": 385.0
                },
                "exit_code": {
                  "field": "Exit status",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/time.txt",
                  "value": 0
                },
                "max_child_rss_kib": {
                  "field": "max_child_rss_kib",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/time.txt",
                  "value": 2407068
                },
                "system_seconds": {
                  "field": "system_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/time.txt",
                  "value": 12.37
                },
                "user_seconds": {
                  "field": "user_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/time.txt",
                  "value": 596.51
                },
                "wall_seconds": {
                  "field": "wall_seconds",
                  "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/time.txt",
                  "value": 157.76
                }
              },
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release",
              "round": 1,
              "status": {
                "directory_exists": true,
                "record": {
                  "exit_code": 0,
                  "finished_ns": 1794686691015602188,
                  "memory_peak_bytes": 2572652544,
                  "oom_kill_detected": false,
                  "status": "ok"
                },
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/status.json",
                "started_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r1-release/started",
                "state": "ok"
              },
              "valid": true
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r2-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r2-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r2-release",
              "round": 2,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r2-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r3-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r3-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r3-release",
              "round": 3,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-540-immkv-r3-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            }
          ],
          "summary": null,
          "valid_rounds": 1
        }
      },
      "fields": 540,
      "media": "nvme",
      "paired_wall_ratios": null
    },
    {
      "correctness": {
        "fresh_echtvar": {
          "constraints": {
            "mismatches": [],
            "observed": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-568-release/constraints.json",
            "valid": false
          },
          "evidence": null,
          "issues": [],
          "note": "Correctness phase only; excluded from timing medians.",
          "passed": false,
          "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-568-release/cell.json",
          "status": {
            "directory_exists": false,
            "record": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-nvme-568-release/status.json",
            "started_source": null,
            "state": "unknown"
          },
          "verifier_stdout": {}
        },
        "immkv": {
          "constraints": {
            "mismatches": [],
            "observed": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-568-release/constraints.json",
            "valid": false
          },
          "evidence": null,
          "issues": [],
          "note": "Correctness phase only; excluded from timing medians.",
          "passed": false,
          "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-568-release/verification.json",
          "status": {
            "directory_exists": false,
            "record": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-568-release/status.json",
            "started_source": null,
            "state": "unknown"
          },
          "verifier_stdout": {
            "verify_values_ok": {
              "field": "VERIFY_VALUES OK",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-nvme-568-release/verifier.stdout",
              "value": false
            }
          }
        }
      },
      "engines": {
        "echtvar": {
          "rounds": [
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r1-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r1-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r1-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r1-release",
              "round": 1,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r1-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r2-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r2-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r2-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r2-release",
              "round": 2,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r2-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r3-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r3-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r3-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r3-release",
              "round": 3,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-echtvar-r3-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            }
          ],
          "summary": null,
          "valid_rounds": 0
        },
        "immkv": {
          "rounds": [
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r1-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r1-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r1-release",
              "round": 1,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r1-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r2-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r2-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r2-release",
              "round": 2,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r2-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r3-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r3-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r3-release",
              "round": 3,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-nvme-568-immkv-r3-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            }
          ],
          "summary": null,
          "valid_rounds": 0
        }
      },
      "fields": 568,
      "media": "nvme",
      "paired_wall_ratios": null
    },
    {
      "correctness": {
        "fresh_echtvar": {
          "constraints": {
            "mismatches": [],
            "observed": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-17-release/constraints.json",
            "valid": false
          },
          "evidence": null,
          "issues": [],
          "note": "Correctness phase only; excluded from timing medians.",
          "passed": false,
          "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-17-release/cell.json",
          "status": {
            "directory_exists": false,
            "record": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-17-release/status.json",
            "started_source": null,
            "state": "unknown"
          },
          "verifier_stdout": {}
        },
        "immkv": {
          "constraints": {
            "mismatches": [],
            "observed": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-17-release/constraints.json",
            "valid": false
          },
          "evidence": null,
          "issues": [],
          "note": "Correctness phase only; excluded from timing medians.",
          "passed": false,
          "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-17-release/verification.json",
          "status": {
            "directory_exists": false,
            "record": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-17-release/status.json",
            "started_source": null,
            "state": "unknown"
          },
          "verifier_stdout": {
            "verify_values_ok": {
              "field": "VERIFY_VALUES OK",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-17-release/verifier.stdout",
              "value": false
            }
          }
        }
      },
      "engines": {
        "echtvar": {
          "rounds": [
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r1-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r1-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r1-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r1-release",
              "round": 1,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r1-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r2-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r2-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r2-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r2-release",
              "round": 2,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r2-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r3-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r3-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r3-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r3-release",
              "round": 3,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-echtvar-r3-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            }
          ],
          "summary": null,
          "valid_rounds": 0
        },
        "immkv": {
          "rounds": [
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r1-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r1-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r1-release",
              "round": 1,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r1-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r2-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r2-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r2-release",
              "round": 2,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r2-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r3-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r3-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r3-release",
              "round": 3,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-17-immkv-r3-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            }
          ],
          "summary": null,
          "valid_rounds": 0
        }
      },
      "fields": 17,
      "media": "hdd",
      "paired_wall_ratios": null
    },
    {
      "correctness": {
        "fresh_echtvar": {
          "constraints": {
            "mismatches": [],
            "observed": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-540-release/constraints.json",
            "valid": false
          },
          "evidence": null,
          "issues": [],
          "note": "Correctness phase only; excluded from timing medians.",
          "passed": false,
          "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-540-release/cell.json",
          "status": {
            "directory_exists": false,
            "record": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-540-release/status.json",
            "started_source": null,
            "state": "unknown"
          },
          "verifier_stdout": {}
        },
        "immkv": {
          "constraints": {
            "mismatches": [],
            "observed": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-540-release/constraints.json",
            "valid": false
          },
          "evidence": null,
          "issues": [],
          "note": "Correctness phase only; excluded from timing medians.",
          "passed": false,
          "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-540-release/verification.json",
          "status": {
            "directory_exists": false,
            "record": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-540-release/status.json",
            "started_source": null,
            "state": "unknown"
          },
          "verifier_stdout": {
            "verify_values_ok": {
              "field": "VERIFY_VALUES OK",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-540-release/verifier.stdout",
              "value": false
            }
          }
        }
      },
      "engines": {
        "echtvar": {
          "rounds": [
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r1-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r1-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r1-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r1-release",
              "round": 1,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r1-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r2-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r2-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r2-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r2-release",
              "round": 2,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r2-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r3-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r3-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r3-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r3-release",
              "round": 3,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-echtvar-r3-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            }
          ],
          "summary": null,
          "valid_rounds": 0
        },
        "immkv": {
          "rounds": [
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r1-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r1-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r1-release",
              "round": 1,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r1-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r2-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r2-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r2-release",
              "round": 2,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r2-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r3-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r3-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r3-release",
              "round": 3,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-540-immkv-r3-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            }
          ],
          "summary": null,
          "valid_rounds": 0
        }
      },
      "fields": 540,
      "media": "hdd",
      "paired_wall_ratios": null
    },
    {
      "correctness": {
        "fresh_echtvar": {
          "constraints": {
            "mismatches": [],
            "observed": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-568-release/constraints.json",
            "valid": false
          },
          "evidence": null,
          "issues": [],
          "note": "Correctness phase only; excluded from timing medians.",
          "passed": false,
          "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-568-release/cell.json",
          "status": {
            "directory_exists": false,
            "record": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-echtvar-hdd-568-release/status.json",
            "started_source": null,
            "state": "unknown"
          },
          "verifier_stdout": {}
        },
        "immkv": {
          "constraints": {
            "mismatches": [],
            "observed": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-568-release/constraints.json",
            "valid": false
          },
          "evidence": null,
          "issues": [],
          "note": "Correctness phase only; excluded from timing medians.",
          "passed": false,
          "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-568-release/verification.json",
          "status": {
            "directory_exists": false,
            "record": null,
            "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-568-release/status.json",
            "started_source": null,
            "state": "unknown"
          },
          "verifier_stdout": {
            "verify_values_ok": {
              "field": "VERIFY_VALUES OK",
              "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/verify-hdd-568-release/verifier.stdout",
              "value": false
            }
          }
        }
      },
      "engines": {
        "echtvar": {
          "rounds": [
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r1-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r1-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r1-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r1-release",
              "round": 1,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r1-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r2-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r2-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r2-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r2-release",
              "round": 2,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r2-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r3-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r3-release/constraints.json",
                "valid": false
              },
              "count_proof": {
                "all_22_chromosomes_ok": false,
                "annotation_log_statistics": [],
                "chromosome_statuses": [],
                "note": "Compressed timing output is discarded; fresh output proof is separate.",
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r3-release/cell.json",
                "timed_output_rows_checked": false
              },
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r3-release",
              "round": 3,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-echtvar-r3-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            }
          ],
          "summary": null,
          "valid_rounds": 0
        },
        "immkv": {
          "rounds": [
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r1-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r1-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r1-release",
              "round": 1,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r1-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r2-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r2-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r2-release",
              "round": 2,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r2-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            },
            {
              "cell_source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r3-release/cell.json",
              "constraints": {
                "mismatches": [],
                "observed": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r3-release/constraints.json",
                "valid": false
              },
              "count_proof": {},
              "immkv_phase_statistics": {},
              "invalid_reasons": [
                "scope status is unknown",
                "scope exit_code is absent or nonzero",
                "cell status/exit_code is absent or unsuccessful",
                "actual 4 CPU / 16 GiB / zero swap constraints are unconfirmed",
                "required annotation counts or chromosome completion are unconfirmed",
                "one or more GNU time / cgroup metrics are absent",
                "final aggregate cgroup peak is unconfirmed",
                "GNU time exit status is absent or nonzero"
              ],
              "issues": [],
              "metrics": {},
              "record_path": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r3-release",
              "round": 3,
              "status": {
                "directory_exists": false,
                "record": null,
                "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/timing-hdd-568-immkv-r3-release/status.json",
                "started_source": null,
                "state": "unknown"
              },
              "valid": false
            }
          ],
          "summary": null,
          "valid_rounds": 0
        }
      },
      "fields": 568,
      "media": "hdd",
      "paired_wall_ratios": null
    }
  ],
  "expected_query": {
    "hits_immkv": 3875208,
    "rows": 3933714,
    "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/verify_output.py"
  },
  "experiment": "vs_echtvar_hg19_hg001_4c16g_v003_release_20261115",
  "notes": [
    "Every metric retains its source path; medians use exactly three complete valid rounds.",
    "OOM, timeout, failed, skipped_after_oom, started and unknown are independent outcomes; no missing result is estimated.",
    "GNU time maximum RSS is maximum child RSS in KiB, not simultaneous aggregate RSS.",
    "Cgroup memory.peak is aggregate charged memory in bytes, including page cache and memfd slices; observer fallback may precede final peak.",
    "CPU percent is GNU time percent (100 percent equals one CPU); four allowed CPUs need not imply only four htslib threads.",
    "Timing includes the complete Python driver tree. Echtvar includes bcftools slicing/BGZF into sealed memfd slices and four chromosome tasks; immkv includes reading, lookup, formatting and compressed output.",
    "Compressed output bytes are discarded during timing; memfd input slices are not stored on the tested medium. These runs do not measure full-output gzip disk-write performance.",
    "Fresh echtvar decompressed-row proof and immkv full-field verification are separate correctness phases; historical output references are not current timing samples.",
    "The release annotate source and recorded binary hashes identify the unchanged implementation and inputs; no current binary is executed or rehashed by this summary.",
    "Cache state is uncontrolled. Three paired raw values and their range describe observed spread only; no p-value or statistical certainty is inferred.",
    "Round engine order is immkv/echtvar, echtvar/immkv, immkv/echtvar (run_matrix.sh). Phase statistics from immkv.log are rounded internal timings and are not additive GNU wall components."
  ],
  "provenance": {
    "build/binaries.sha256": {
      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/build/binaries.sha256",
      "text": "4179f6f9bf774615ce68e1794dd3c46f2c5e4b2316011b117191584f5ec16391  annotate\n"
    },
    "build/build-info.txt": {
      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/build/build-info.txt",
      "text": "/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/bin/annotate: go1.27.1\n\tpath\tgithub.com/viktorxia/immkv/test/annotate\n\tmod\tgithub.com/viktorxia/immkv\tv0.0.3+dirty\t\n\tdep\tgithub.com/biogo/hts\tv1.4.5\th1:mhVCpZaTYlAhBjMaAATGWBnauioBtmvOb0ApLdU4/+0=\n\tdep\tgithub.com/klauspost/compress\tv1.20.3\th1:zWNoFqq/lSuYQRim3dUXlXeA0IqwuZGicdG5ZTfY5j4=\n\tdep\tgithub.com/valyala/gozstd\tv1.26.0\th1:vvQxZ4FANU2Q6FG4dUMpuQ8PUvkq9MdbU2jFkF+brm8=\n\tdep\tgolang.org/x/sys\tv0.49.0\th1:XbzkgYJHdqh/8m2Uu0W/dQv8nktxx4BFHp1M0gROTXA=\n\tbuild\t-buildmode=exe\n\tbuild\t-compiler=gc\n\tbuild\t-trimpath=true\n\tbuild\tDefaultGODEBUG=tracebacklabels=0,x509sslcertoverrideplatform=0\n\tbuild\tCGO_ENABLED=1\n\tbuild\tGOARCH=amd64\n\tbuild\tGOOS=linux\n\tbuild\tGOAMD64=v1\n\tbuild\tvcs=git\n\tbuild\tvcs.revision=840d36d80351f8b8651d9dc90abf1314bf5db21c\n\tbuild\tvcs.time=2026-11-14T19:17:41Z\n\tbuild\tvcs.modified=true\n"
    },
    "build/environment.txt": {
      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/build/environment.txt",
      "text": "2026-11-15T03:26:55+08:00\ngo version go1.27.1 linux/amd64\n1\nlinux\namd64\n/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/cache/go-build\n/home/tian/projects/immkv/tmp/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/cache/mod\n-mod=readonly\nLinux tian-mini 7.0.0-42-generic #42~24.04.1-Ubuntu SMP PREEMPT_DYNAMIC Mon Oct 19 14:43:01 UTC 2 x86_64 x86_64 x86_64 GNU/Linux\nechtvar 0.2.4\nbcftools 1.25\nUsing htslib 1.25\nCopyright (C) 2026 Genome Research Ltd.\nLicense GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>\nThis is free software: you are free to change and redistribute it.\nThere is NO WARRANTY, to the extent permitted by law.\n2c425bad9a4e6aeeade5eeec229e7f887d7385642a0aada3a03909b6303f1e09  /home/tian/projects/immkv/tmp/echtvar/echtvar\n728cc52cd0a82eed97995b78622f882d9e19c9b947b687c8bbb295065aca9c58  /home/tian/miniforge3/envs/transvar/bin/bcftools\n"
    },
    "build/source-facts.json": {
      "issues": [],
      "record": {
        "annotate": "Unchanged test/annotate from release git archive",
        "archive_sha256": "d5bd8a2545d080bbb8992cb7ff4bf5518480193f131899fabae11dc33c06941e",
        "commit": "840d36d80351f8b8651d9dc90abf1314bf5db21c",
        "dependencies": [
          [
            "github.com/klauspost/compress",
            "v1.20.3"
          ],
          [
            "github.com/valyala/gozstd",
            "v1.26.0"
          ],
          [
            "golang.org/x/sys",
            "v0.49.0"
          ],
          [
            "github.com/biogo/hts",
            "v1.4.5"
          ]
        ],
        "module_cache_origin": "/home/tian/go/pkg/mod"
      },
      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/build/source-facts.json"
    },
    "input-preflight.json": {
      "issues": [],
      "record": {
        "created_ns": 1794684432117851506,
        "media": {
          "hdd": {
            "17": {
              "archives": [
                {
                  "bytes": 244234654,
                  "central_directory_members": 4163,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr1.echtvar.zip",
                  "uncompressed_bytes": 578567882
                },
                {
                  "bytes": 258657713,
                  "central_directory_members": 4391,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr2.echtvar.zip",
                  "uncompressed_bytes": 620538239
                },
                {
                  "bytes": 210659980,
                  "central_directory_members": 3555,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr3.echtvar.zip",
                  "uncompressed_bytes": 505087392
                },
                {
                  "bytes": 208828208,
                  "central_directory_members": 3441,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr4.echtvar.zip",
                  "uncompressed_bytes": 494714168
                },
                {
                  "bytes": 192352387,
                  "central_directory_members": 3251,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr5.echtvar.zip",
                  "uncompressed_bytes": 461270809
                },
                {
                  "bytes": 183934502,
                  "central_directory_members": 3080,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr6.echtvar.zip",
                  "uncompressed_bytes": 432604103
                },
                {
                  "bytes": 178674759,
                  "central_directory_members": 2852,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr7.echtvar.zip",
                  "uncompressed_bytes": 420231928
                },
                {
                  "bytes": 164725132,
                  "central_directory_members": 2624,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr8.echtvar.zip",
                  "uncompressed_bytes": 397702935
                },
                {
                  "bytes": 138388260,
                  "central_directory_members": 2263,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr9.echtvar.zip",
                  "uncompressed_bytes": 325968828
                },
                {
                  "bytes": 147108442,
                  "central_directory_members": 2434,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr10.echtvar.zip",
                  "uncompressed_bytes": 345447116
                },
                {
                  "bytes": 146764897,
                  "central_directory_members": 2415,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr11.echtvar.zip",
                  "uncompressed_bytes": 350042016
                },
                {
                  "bytes": 144193515,
                  "central_directory_members": 2396,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr12.echtvar.zip",
                  "uncompressed_bytes": 339777686
                },
                {
                  "bytes": 104972306,
                  "central_directory_members": 1750,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr13.echtvar.zip",
                  "uncompressed_bytes": 247521444
                },
                {
                  "bytes": 99273295,
                  "central_directory_members": 1617,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr14.echtvar.zip",
                  "uncompressed_bytes": 233719970
                },
                {
                  "bytes": 93455566,
                  "central_directory_members": 1503,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr15.echtvar.zip",
                  "uncompressed_bytes": 219684326
                },
                {
                  "bytes": 102348503,
                  "central_directory_members": 1465,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr16.echtvar.zip",
                  "uncompressed_bytes": 244290082
                },
                {
                  "bytes": 89862907,
                  "central_directory_members": 1446,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr17.echtvar.zip",
                  "uncompressed_bytes": 210528323
                },
                {
                  "bytes": 83163449,
                  "central_directory_members": 1389,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr18.echtvar.zip",
                  "uncompressed_bytes": 196324184
                },
                {
                  "bytes": 75601696,
                  "central_directory_members": 1047,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr19.echtvar.zip",
                  "uncompressed_bytes": 172173473
                },
                {
                  "bytes": 66291586,
                  "central_directory_members": 1123,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr20.echtvar.zip",
                  "uncompressed_bytes": 157162354
                },
                {
                  "bytes": 43014863,
                  "central_directory_members": 686,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr21.echtvar.zip",
                  "uncompressed_bytes": 99628454
                },
                {
                  "bytes": 44732754,
                  "central_directory_members": 648,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/archives/chr22.echtvar.zip",
                  "uncompressed_bytes": 102376807
                }
              ],
              "fields_sha256": "db2c8bc29926b5f4265e9077443a19571a55c8f99d2d30906e0a126c6c0db030",
              "index_bytes": 1367830,
              "library": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/lib17",
              "manifests": [
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 1,
                  "manifest_sha256": "b828cfa17a2df725f18e9907cee4cf0e9907c433d0ab5992ad6a25651f59852d",
                  "records": 20370890,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 2,
                  "manifest_sha256": "c99e8c2b512d651fefdba28fd4bc47ebc1ca96b816393fd8c77d08245a868505",
                  "records": 21892245,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 3,
                  "manifest_sha256": "88c35cffd4cc14950f915a4b39d1f466ccd0bcf1d49890a29beb47436dbf4f55",
                  "records": 17834355,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 4,
                  "manifest_sha256": "e310f6b01626703f969a903e630f029d97d3fdb089de0c8e3a5a1cb4b3f2a23c",
                  "records": 17423639,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 5,
                  "manifest_sha256": "96bfbf63bfb6181918baa760bdb8ff6ec675ef1459e2241ce0eb6eee1f8490f4",
                  "records": 16290654,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 6,
                  "manifest_sha256": "9a77beac2449ac85a8cf028e7769b9568fc7dc168c7f82134761233c8a6983ca",
                  "records": 15190442,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 7,
                  "manifest_sha256": "3e6c4b0b97e2e8acc707f9001fba3c96b7583aa761f285fb9187726b0737467f",
                  "records": 14752321,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 8,
                  "manifest_sha256": "a0de0210e73dbd886181a9c6e3b39e17854efabb53aa95d2f39114ab5de8ed64",
                  "records": 14073574,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 9,
                  "manifest_sha256": "90d4ded4deb27f1c1ffadebe47f5c65cf45455e657191676a4c23c557b61f8a9",
                  "records": 11480614,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 10,
                  "manifest_sha256": "12d27572091e8cb3fbb27386b5e777f73e330643f72d18ac8ea79618f465d547",
                  "records": 12118750,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 11,
                  "manifest_sha256": "c9343209324cedb76d8336d0b162e13adbde8ac07357906e101d816042827b3c",
                  "records": 12345492,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 12,
                  "manifest_sha256": "9c93500cf3fdbdde0991afd4a9a42602419f56be02fc0ede49e2268dec415de8",
                  "records": 11933435,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 13,
                  "manifest_sha256": "ee367aa983c8d31b9ef6485966954bc84c840c308f9ec10673955eea0ba8f052",
                  "records": 8693573,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 14,
                  "manifest_sha256": "2336388c4f0d6324d1d391bcbac4e81f0e32e95a343bb1a6872d1879f8446f2c",
                  "records": 8220638,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 15,
                  "manifest_sha256": "ca21551ad8d8d7b5e92ff40d830393ba374c9aaa8e18a45017433de9ea1c9b29",
                  "records": 7723667,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 16,
                  "manifest_sha256": "43209ec7835c225072e05d34144afba16d6bba7d83b747fccceb6f018d837cc7",
                  "records": 8618657,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 17,
                  "manifest_sha256": "57583441aab447fa2243d6939f679e50098ab45bb207b93820a679c6d36b6488",
                  "records": 7373819,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 18,
                  "manifest_sha256": "9a8bb1ae9a643ee3388533082705d2e26bcd09ab03cda90b64e8965d509cd369",
                  "records": 6893751,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 19,
                  "manifest_sha256": "c905eb1ca428e5fc53d8d67e0c5d4fdc0ee719800d8e089e3f3b794d6c80510a",
                  "records": 5991870,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 20,
                  "manifest_sha256": "fc4513940ce0caed90006e84a126eb2660c11f6e07a28c1cf70d352e89d529a3",
                  "records": 5516836,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 21,
                  "manifest_sha256": "5a064dcd7b265753c4e4e49d31e6aec772357bba315257e7d1b0a8e78d6c1772",
                  "records": 3483000,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 22,
                  "manifest_sha256": "d5987ea72aaf88492de717ab4f935c2108bec30b957d63357268d66bafff178c",
                  "records": 3567085,
                  "schema_hash": "99944e3316871f58"
                }
              ],
              "query": "/mnt/data6t/tian/immkv_test_hdd/test/17_fields/query/hg001.decomp.vcf.gz",
              "query_bytes": 139085616,
              "query_sha256": "b9afb10b724c9b85c59e723be330a3ea756928c9ca2092692db3b1faa64472d0"
            },
            "540": {
              "archives": [
                {
                  "bytes": 5123651180,
                  "central_directory_members": 118700,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr1.echtvar.zip",
                  "uncompressed_bytes": 17527475456
                },
                {
                  "bytes": 5398570861,
                  "central_directory_members": 125204,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr2.echtvar.zip",
                  "uncompressed_bytes": 18815801585
                },
                {
                  "bytes": 4396709608,
                  "central_directory_members": 101356,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr3.echtvar.zip",
                  "uncompressed_bytes": 15334578110
                },
                {
                  "bytes": 4372074758,
                  "central_directory_members": 98104,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr4.echtvar.zip",
                  "uncompressed_bytes": 15000414640
                },
                {
                  "bytes": 4015346448,
                  "central_directory_members": 92684,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr5.echtvar.zip",
                  "uncompressed_bytes": 14002770510
                },
                {
                  "bytes": 3834498842,
                  "central_directory_members": 87806,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr6.echtvar.zip",
                  "uncompressed_bytes": 13088773102
                },
                {
                  "bytes": 3752875084,
                  "central_directory_members": 81302,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr7.echtvar.zip",
                  "uncompressed_bytes": 12699326179
                },
                {
                  "bytes": 3445232398,
                  "central_directory_members": 74798,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr8.echtvar.zip",
                  "uncompressed_bytes": 12093741209
                },
                {
                  "bytes": 2927934457,
                  "central_directory_members": 64500,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr9.echtvar.zip",
                  "uncompressed_bytes": 9883896838
                },
                {
                  "bytes": 3075555846,
                  "central_directory_members": 69378,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr10.echtvar.zip",
                  "uncompressed_bytes": 10438645523
                },
                {
                  "bytes": 3075222135,
                  "central_directory_members": 68836,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr11.echtvar.zip",
                  "uncompressed_bytes": 10620876194
                },
                {
                  "bytes": 3019225963,
                  "central_directory_members": 68294,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr12.echtvar.zip",
                  "uncompressed_bytes": 10271390853
                },
                {
                  "bytes": 2190906161,
                  "central_directory_members": 49866,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr13.echtvar.zip",
                  "uncompressed_bytes": 7487102196
                },
                {
                  "bytes": 2082208118,
                  "central_directory_members": 46072,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr14.echtvar.zip",
                  "uncompressed_bytes": 7077438370
                },
                {
                  "bytes": 1966885592,
                  "central_directory_members": 42820,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr15.echtvar.zip",
                  "uncompressed_bytes": 6648377385
                },
                {
                  "bytes": 2159856477,
                  "central_directory_members": 41736,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr16.echtvar.zip",
                  "uncompressed_bytes": 7407307841
                },
                {
                  "bytes": 1890442581,
                  "central_directory_members": 41194,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr17.echtvar.zip",
                  "uncompressed_bytes": 6342630113
                },
                {
                  "bytes": 1730167927,
                  "central_directory_members": 39568,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr18.echtvar.zip",
                  "uncompressed_bytes": 5936324589
                },
                {
                  "bytes": 1612041534,
                  "central_directory_members": 29812,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr19.echtvar.zip",
                  "uncompressed_bytes": 5165403019
                },
                {
                  "bytes": 1376525969,
                  "central_directory_members": 31980,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr20.echtvar.zip",
                  "uncompressed_bytes": 4746336425
                },
                {
                  "bytes": 902589299,
                  "central_directory_members": 19514,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr21.echtvar.zip",
                  "uncompressed_bytes": 3004608164
                },
                {
                  "bytes": 945741018,
                  "central_directory_members": 18430,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/archives/chr22.echtvar.zip",
                  "uncompressed_bytes": 3075770914
                }
              ],
              "fields_sha256": "fa1f21636fcd186308e3d070a65b0c087a8813f6763a78f8db70dc96c65f7643",
              "index_bytes": 1367830,
              "library": "/mnt/data6t/tian/immkv_test_hdd/gnomad/immkv",
              "manifests": [
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 1,
                  "manifest_sha256": "e6502a10540c839e768f615df8081d70b4a71e61b2a0a88afa0821d687343e4d",
                  "records": 20370890,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 2,
                  "manifest_sha256": "0be51c08eec7941cd997de60a6fef8e3f72f9f38c3a4af9d889e395a7de99ea7",
                  "records": 21892245,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 3,
                  "manifest_sha256": "4fb8ff096c20282a966b468355676dc8e476b8400d7c200f4a42bd882d587e9c",
                  "records": 17834355,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 4,
                  "manifest_sha256": "85e3fdd517978eefa71480c09b88efbfe179188da72392025e1c995802ff3e37",
                  "records": 17423639,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 5,
                  "manifest_sha256": "88624fbba2bbc84b65c8ef80d78420d0b7a35928ed8ae9ac523b9f2e8f0f91a3",
                  "records": 16290654,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 6,
                  "manifest_sha256": "bdf9562a2110e95c010b344e601f16f8a543a51dca79fa2ced4031c53896fb91",
                  "records": 15190442,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 7,
                  "manifest_sha256": "77e5990e85572cbbbfce7871d79ee2fde5f7fc18cff2fe9a8f9c9900e2d824b5",
                  "records": 14752321,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 8,
                  "manifest_sha256": "56f8047c1da4c09bdb6e15f44acd99c448132ae6230e32ef22cfa20239567b5e",
                  "records": 14073574,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 9,
                  "manifest_sha256": "48b7a1d79c9a4a057bc7513eb930b850d7354f2abe380e3768aadaf87b45714f",
                  "records": 11480614,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 10,
                  "manifest_sha256": "db671e979c2a949f8c0433112f9053bcb02dcf4bc09a30e3407cdffa9378788a",
                  "records": 12118750,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 11,
                  "manifest_sha256": "0c8f3ed0470504385960e748aba1bc8854b806e3dd5044768739ef3d3a19a7f6",
                  "records": 12345492,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 12,
                  "manifest_sha256": "f0d71794035581a52f8a9c49187e7c5d77f79f5300aa76fb147efa0afa65ff88",
                  "records": 11933435,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 13,
                  "manifest_sha256": "563bb2b82c86652d08e862b28befc5be90d24897cf921aaad34621e79744f5d1",
                  "records": 8693573,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 14,
                  "manifest_sha256": "0471ec2e516c76f6b89410c89d8825aebfa36bfc9b15378c0045f9b46f7035c7",
                  "records": 8220638,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 15,
                  "manifest_sha256": "f9210d1c3b5ca3d8b36087dcc025a44f9e9934158ef77f78d85b5bb90c391667",
                  "records": 7723667,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 16,
                  "manifest_sha256": "3cf4786236d60d3c5402f496bd1cd8cfbe2ac3a2773e32f7cb1e680ef96c9a77",
                  "records": 8618657,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 17,
                  "manifest_sha256": "7f18009ed6de25e04fe6781dd2b181ad310ca3ecadeb5d65fd5f881be16eff8b",
                  "records": 7373819,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 18,
                  "manifest_sha256": "e7d6244b50c2e3b2752e25ab393dee0d5135133f087b10a0178d83a0bdb72ccb",
                  "records": 6893751,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 19,
                  "manifest_sha256": "df905df6fec7a78b7a65108551deb1e8c29bd3a3dca06edf68f5488bf8e7d360",
                  "records": 5991870,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 20,
                  "manifest_sha256": "5ffef94f04de22a04f846f3f06f2b33d1c6e828cc099efd15abeb6d6196a1db5",
                  "records": 5516836,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 21,
                  "manifest_sha256": "869bb863c33345bfcedba5de9265f1acac3c5033460ef83cb11d432d7143cdf4",
                  "records": 3483000,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 22,
                  "manifest_sha256": "872946e081799b83f30444581a23e5d7a4d3a4f46dba7e12a5185f8109e77895",
                  "records": 3567085,
                  "schema_hash": "1e91a99722c59d5e"
                }
              ],
              "query": "/mnt/data6t/tian/immkv_test_hdd/test/540_fields/query/hg001.decomp.vcf.gz",
              "query_bytes": 139085616,
              "query_sha256": "4fed97e319de887d251a7b1c603f2f842dfbc1247c1740eb7140172b67b22fe8"
            },
            "568": {
              "archives": [
                {
                  "bytes": 7257511949,
                  "central_directory_members": 124848,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr1.echtvar.zip",
                  "uncompressed_bytes": 38780286499
                },
                {
                  "bytes": 7624945470,
                  "central_directory_members": 131688,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr2.echtvar.zip",
                  "uncompressed_bytes": 40899531669
                },
                {
                  "bytes": 6263830958,
                  "central_directory_members": 106608,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr3.echtvar.zip",
                  "uncompressed_bytes": 34926259821
                },
                {
                  "bytes": 6074278192,
                  "central_directory_members": 103188,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr4.echtvar.zip",
                  "uncompressed_bytes": 30168583462
                },
                {
                  "bytes": 5622847357,
                  "central_directory_members": 97488,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr5.echtvar.zip",
                  "uncompressed_bytes": 28696987035
                },
                {
                  "bytes": 5333501460,
                  "central_directory_members": 92358,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr6.echtvar.zip",
                  "uncompressed_bytes": 26290904570
                },
                {
                  "bytes": 5284983462,
                  "central_directory_members": 85518,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr7.echtvar.zip",
                  "uncompressed_bytes": 27981956434
                },
                {
                  "bytes": 4867485483,
                  "central_directory_members": 78678,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr8.echtvar.zip",
                  "uncompressed_bytes": 26040552400
                },
                {
                  "bytes": 4063797988,
                  "central_directory_members": 67848,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr9.echtvar.zip",
                  "uncompressed_bytes": 19890246125
                },
                {
                  "bytes": 4299658931,
                  "central_directory_members": 72978,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr10.echtvar.zip",
                  "uncompressed_bytes": 21839763121
                },
                {
                  "bytes": 4415331558,
                  "central_directory_members": 72408,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr11.echtvar.zip",
                  "uncompressed_bytes": 24849737455
                },
                {
                  "bytes": 4324593804,
                  "central_directory_members": 71838,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr12.echtvar.zip",
                  "uncompressed_bytes": 24288159769
                },
                {
                  "bytes": 2982695276,
                  "central_directory_members": 52458,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr13.echtvar.zip",
                  "uncompressed_bytes": 13200835719
                },
                {
                  "bytes": 2978458193,
                  "central_directory_members": 48468,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr14.echtvar.zip",
                  "uncompressed_bytes": 16672855058
                },
                {
                  "bytes": 2821168599,
                  "central_directory_members": 45048,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr15.echtvar.zip",
                  "uncompressed_bytes": 16007926228
                },
                {
                  "bytes": 3147128737,
                  "central_directory_members": 43908,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr16.echtvar.zip",
                  "uncompressed_bytes": 18820464097
                },
                {
                  "bytes": 2804825536,
                  "central_directory_members": 43338,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr17.echtvar.zip",
                  "uncompressed_bytes": 17740105448
                },
                {
                  "bytes": 2434370983,
                  "central_directory_members": 41628,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr18.echtvar.zip",
                  "uncompressed_bytes": 12818068081
                },
                {
                  "bytes": 2393428912,
                  "central_directory_members": 31368,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr19.echtvar.zip",
                  "uncompressed_bytes": 14771573347
                },
                {
                  "bytes": 1939275218,
                  "central_directory_members": 33648,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr20.echtvar.zip",
                  "uncompressed_bytes": 10042845922
                },
                {
                  "bytes": 1259700453,
                  "central_directory_members": 20538,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr21.echtvar.zip",
                  "uncompressed_bytes": 6372577608
                },
                {
                  "bytes": 1362922300,
                  "central_directory_members": 19398,
                  "central_directory_valid": true,
                  "path": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/archives/chr22.echtvar.zip",
                  "uncompressed_bytes": 7774267602
                }
              ],
              "fields_sha256": "28e7720f959b1cf583af444f1290a9bad857e3d733ea8cd6d0b75442b78ebf86",
              "index_bytes": 1367830,
              "library": "/mnt/data6t/tian/immkv_test_hdd/gnomad/immkv",
              "manifests": [
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 1,
                  "manifest_sha256": "e6502a10540c839e768f615df8081d70b4a71e61b2a0a88afa0821d687343e4d",
                  "records": 20370890,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 2,
                  "manifest_sha256": "0be51c08eec7941cd997de60a6fef8e3f72f9f38c3a4af9d889e395a7de99ea7",
                  "records": 21892245,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 3,
                  "manifest_sha256": "4fb8ff096c20282a966b468355676dc8e476b8400d7c200f4a42bd882d587e9c",
                  "records": 17834355,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 4,
                  "manifest_sha256": "85e3fdd517978eefa71480c09b88efbfe179188da72392025e1c995802ff3e37",
                  "records": 17423639,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 5,
                  "manifest_sha256": "88624fbba2bbc84b65c8ef80d78420d0b7a35928ed8ae9ac523b9f2e8f0f91a3",
                  "records": 16290654,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 6,
                  "manifest_sha256": "bdf9562a2110e95c010b344e601f16f8a543a51dca79fa2ced4031c53896fb91",
                  "records": 15190442,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 7,
                  "manifest_sha256": "77e5990e85572cbbbfce7871d79ee2fde5f7fc18cff2fe9a8f9c9900e2d824b5",
                  "records": 14752321,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 8,
                  "manifest_sha256": "56f8047c1da4c09bdb6e15f44acd99c448132ae6230e32ef22cfa20239567b5e",
                  "records": 14073574,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 9,
                  "manifest_sha256": "48b7a1d79c9a4a057bc7513eb930b850d7354f2abe380e3768aadaf87b45714f",
                  "records": 11480614,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 10,
                  "manifest_sha256": "db671e979c2a949f8c0433112f9053bcb02dcf4bc09a30e3407cdffa9378788a",
                  "records": 12118750,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 11,
                  "manifest_sha256": "0c8f3ed0470504385960e748aba1bc8854b806e3dd5044768739ef3d3a19a7f6",
                  "records": 12345492,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 12,
                  "manifest_sha256": "f0d71794035581a52f8a9c49187e7c5d77f79f5300aa76fb147efa0afa65ff88",
                  "records": 11933435,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 13,
                  "manifest_sha256": "563bb2b82c86652d08e862b28befc5be90d24897cf921aaad34621e79744f5d1",
                  "records": 8693573,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 14,
                  "manifest_sha256": "0471ec2e516c76f6b89410c89d8825aebfa36bfc9b15378c0045f9b46f7035c7",
                  "records": 8220638,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 15,
                  "manifest_sha256": "f9210d1c3b5ca3d8b36087dcc025a44f9e9934158ef77f78d85b5bb90c391667",
                  "records": 7723667,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 16,
                  "manifest_sha256": "3cf4786236d60d3c5402f496bd1cd8cfbe2ac3a2773e32f7cb1e680ef96c9a77",
                  "records": 8618657,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 17,
                  "manifest_sha256": "7f18009ed6de25e04fe6781dd2b181ad310ca3ecadeb5d65fd5f881be16eff8b",
                  "records": 7373819,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 18,
                  "manifest_sha256": "e7d6244b50c2e3b2752e25ab393dee0d5135133f087b10a0178d83a0bdb72ccb",
                  "records": 6893751,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 19,
                  "manifest_sha256": "df905df6fec7a78b7a65108551deb1e8c29bd3a3dca06edf68f5488bf8e7d360",
                  "records": 5991870,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 20,
                  "manifest_sha256": "5ffef94f04de22a04f846f3f06f2b33d1c6e828cc099efd15abeb6d6196a1db5",
                  "records": 5516836,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 21,
                  "manifest_sha256": "869bb863c33345bfcedba5de9265f1acac3c5033460ef83cb11d432d7143cdf4",
                  "records": 3483000,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 22,
                  "manifest_sha256": "872946e081799b83f30444581a23e5d7a4d3a4f46dba7e12a5185f8109e77895",
                  "records": 3567085,
                  "schema_hash": "1e91a99722c59d5e"
                }
              ],
              "query": "/mnt/data6t/tian/immkv_test_hdd/test/568_fields/query/hg001.decomp.vcf.gz",
              "query_bytes": 139085616,
              "query_sha256": "4c61188a183910ae4b80fa038139718d0ccf3b3fe3911a7c14917109b3ceee06"
            }
          },
          "nvme": {
            "17": {
              "archives": [
                {
                  "bytes": 244234654,
                  "central_directory_members": 4163,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr1.echtvar.zip",
                  "uncompressed_bytes": 578567882
                },
                {
                  "bytes": 258657713,
                  "central_directory_members": 4391,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr2.echtvar.zip",
                  "uncompressed_bytes": 620538239
                },
                {
                  "bytes": 210659980,
                  "central_directory_members": 3555,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr3.echtvar.zip",
                  "uncompressed_bytes": 505087392
                },
                {
                  "bytes": 208828208,
                  "central_directory_members": 3441,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr4.echtvar.zip",
                  "uncompressed_bytes": 494714168
                },
                {
                  "bytes": 192352387,
                  "central_directory_members": 3251,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr5.echtvar.zip",
                  "uncompressed_bytes": 461270809
                },
                {
                  "bytes": 183934502,
                  "central_directory_members": 3080,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr6.echtvar.zip",
                  "uncompressed_bytes": 432604103
                },
                {
                  "bytes": 178674759,
                  "central_directory_members": 2852,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr7.echtvar.zip",
                  "uncompressed_bytes": 420231928
                },
                {
                  "bytes": 164725132,
                  "central_directory_members": 2624,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr8.echtvar.zip",
                  "uncompressed_bytes": 397702935
                },
                {
                  "bytes": 138388260,
                  "central_directory_members": 2263,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr9.echtvar.zip",
                  "uncompressed_bytes": 325968828
                },
                {
                  "bytes": 147108442,
                  "central_directory_members": 2434,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr10.echtvar.zip",
                  "uncompressed_bytes": 345447116
                },
                {
                  "bytes": 146764897,
                  "central_directory_members": 2415,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr11.echtvar.zip",
                  "uncompressed_bytes": 350042016
                },
                {
                  "bytes": 144193515,
                  "central_directory_members": 2396,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr12.echtvar.zip",
                  "uncompressed_bytes": 339777686
                },
                {
                  "bytes": 104972306,
                  "central_directory_members": 1750,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr13.echtvar.zip",
                  "uncompressed_bytes": 247521444
                },
                {
                  "bytes": 99273295,
                  "central_directory_members": 1617,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr14.echtvar.zip",
                  "uncompressed_bytes": 233719970
                },
                {
                  "bytes": 93455566,
                  "central_directory_members": 1503,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr15.echtvar.zip",
                  "uncompressed_bytes": 219684326
                },
                {
                  "bytes": 102348503,
                  "central_directory_members": 1465,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr16.echtvar.zip",
                  "uncompressed_bytes": 244290082
                },
                {
                  "bytes": 89862907,
                  "central_directory_members": 1446,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr17.echtvar.zip",
                  "uncompressed_bytes": 210528323
                },
                {
                  "bytes": 83163449,
                  "central_directory_members": 1389,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr18.echtvar.zip",
                  "uncompressed_bytes": 196324184
                },
                {
                  "bytes": 75601696,
                  "central_directory_members": 1047,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr19.echtvar.zip",
                  "uncompressed_bytes": 172173473
                },
                {
                  "bytes": 66291586,
                  "central_directory_members": 1123,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr20.echtvar.zip",
                  "uncompressed_bytes": 157162354
                },
                {
                  "bytes": 43014863,
                  "central_directory_members": 686,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr21.echtvar.zip",
                  "uncompressed_bytes": 99628454
                },
                {
                  "bytes": 44732754,
                  "central_directory_members": 648,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/17_fields/archives/chr22.echtvar.zip",
                  "uncompressed_bytes": 102376807
                }
              ],
              "fields_sha256": "db2c8bc29926b5f4265e9077443a19571a55c8f99d2d30906e0a126c6c0db030",
              "index_bytes": 1367854,
              "library": "/home/tian/immkv_test_nvme/test/17_fields/lib17",
              "manifests": [
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 1,
                  "manifest_sha256": "b828cfa17a2df725f18e9907cee4cf0e9907c433d0ab5992ad6a25651f59852d",
                  "records": 20370890,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 2,
                  "manifest_sha256": "c99e8c2b512d651fefdba28fd4bc47ebc1ca96b816393fd8c77d08245a868505",
                  "records": 21892245,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 3,
                  "manifest_sha256": "88c35cffd4cc14950f915a4b39d1f466ccd0bcf1d49890a29beb47436dbf4f55",
                  "records": 17834355,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 4,
                  "manifest_sha256": "e310f6b01626703f969a903e630f029d97d3fdb089de0c8e3a5a1cb4b3f2a23c",
                  "records": 17423639,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 5,
                  "manifest_sha256": "96bfbf63bfb6181918baa760bdb8ff6ec675ef1459e2241ce0eb6eee1f8490f4",
                  "records": 16290654,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 6,
                  "manifest_sha256": "9a77beac2449ac85a8cf028e7769b9568fc7dc168c7f82134761233c8a6983ca",
                  "records": 15190442,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 7,
                  "manifest_sha256": "3e6c4b0b97e2e8acc707f9001fba3c96b7583aa761f285fb9187726b0737467f",
                  "records": 14752321,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 8,
                  "manifest_sha256": "a0de0210e73dbd886181a9c6e3b39e17854efabb53aa95d2f39114ab5de8ed64",
                  "records": 14073574,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 9,
                  "manifest_sha256": "90d4ded4deb27f1c1ffadebe47f5c65cf45455e657191676a4c23c557b61f8a9",
                  "records": 11480614,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 10,
                  "manifest_sha256": "12d27572091e8cb3fbb27386b5e777f73e330643f72d18ac8ea79618f465d547",
                  "records": 12118750,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 11,
                  "manifest_sha256": "c9343209324cedb76d8336d0b162e13adbde8ac07357906e101d816042827b3c",
                  "records": 12345492,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 12,
                  "manifest_sha256": "9c93500cf3fdbdde0991afd4a9a42602419f56be02fc0ede49e2268dec415de8",
                  "records": 11933435,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 13,
                  "manifest_sha256": "ee367aa983c8d31b9ef6485966954bc84c840c308f9ec10673955eea0ba8f052",
                  "records": 8693573,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 14,
                  "manifest_sha256": "2336388c4f0d6324d1d391bcbac4e81f0e32e95a343bb1a6872d1879f8446f2c",
                  "records": 8220638,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 15,
                  "manifest_sha256": "ca21551ad8d8d7b5e92ff40d830393ba374c9aaa8e18a45017433de9ea1c9b29",
                  "records": 7723667,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 16,
                  "manifest_sha256": "43209ec7835c225072e05d34144afba16d6bba7d83b747fccceb6f018d837cc7",
                  "records": 8618657,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 17,
                  "manifest_sha256": "57583441aab447fa2243d6939f679e50098ab45bb207b93820a679c6d36b6488",
                  "records": 7373819,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 18,
                  "manifest_sha256": "9a8bb1ae9a643ee3388533082705d2e26bcd09ab03cda90b64e8965d509cd369",
                  "records": 6893751,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 19,
                  "manifest_sha256": "c905eb1ca428e5fc53d8d67e0c5d4fdc0ee719800d8e089e3f3b794d6c80510a",
                  "records": 5991870,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 20,
                  "manifest_sha256": "fc4513940ce0caed90006e84a126eb2660c11f6e07a28c1cf70d352e89d529a3",
                  "records": 5516836,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 21,
                  "manifest_sha256": "5a064dcd7b265753c4e4e49d31e6aec772357bba315257e7d1b0a8e78d6c1772",
                  "records": 3483000,
                  "schema_hash": "99944e3316871f58"
                },
                {
                  "build_id": "7e9d2daf469b5ec1581da28cfc6b2894",
                  "chromosome": 22,
                  "manifest_sha256": "d5987ea72aaf88492de717ab4f935c2108bec30b957d63357268d66bafff178c",
                  "records": 3567085,
                  "schema_hash": "99944e3316871f58"
                }
              ],
              "query": "/home/tian/immkv_test_nvme/test/17_fields/query/hg001.decomp.vcf.gz",
              "query_bytes": 139085620,
              "query_sha256": "014dc5c07f967a71cfdcc650c57148aeaac34ceb12e7d9f3a513b9720a88a9c6"
            },
            "540": {
              "archives": [
                {
                  "bytes": 5123651180,
                  "central_directory_members": 118700,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr1.echtvar.zip",
                  "uncompressed_bytes": 17527475456
                },
                {
                  "bytes": 5398570861,
                  "central_directory_members": 125204,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr2.echtvar.zip",
                  "uncompressed_bytes": 18815801585
                },
                {
                  "bytes": 4396709608,
                  "central_directory_members": 101356,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr3.echtvar.zip",
                  "uncompressed_bytes": 15334578110
                },
                {
                  "bytes": 4372074758,
                  "central_directory_members": 98104,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr4.echtvar.zip",
                  "uncompressed_bytes": 15000414640
                },
                {
                  "bytes": 4015346448,
                  "central_directory_members": 92684,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr5.echtvar.zip",
                  "uncompressed_bytes": 14002770510
                },
                {
                  "bytes": 3834498842,
                  "central_directory_members": 87806,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr6.echtvar.zip",
                  "uncompressed_bytes": 13088773102
                },
                {
                  "bytes": 3752875084,
                  "central_directory_members": 81302,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr7.echtvar.zip",
                  "uncompressed_bytes": 12699326179
                },
                {
                  "bytes": 3445232398,
                  "central_directory_members": 74798,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr8.echtvar.zip",
                  "uncompressed_bytes": 12093741209
                },
                {
                  "bytes": 2927934457,
                  "central_directory_members": 64500,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr9.echtvar.zip",
                  "uncompressed_bytes": 9883896838
                },
                {
                  "bytes": 3075555846,
                  "central_directory_members": 69378,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr10.echtvar.zip",
                  "uncompressed_bytes": 10438645523
                },
                {
                  "bytes": 3075222135,
                  "central_directory_members": 68836,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr11.echtvar.zip",
                  "uncompressed_bytes": 10620876194
                },
                {
                  "bytes": 3019225963,
                  "central_directory_members": 68294,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr12.echtvar.zip",
                  "uncompressed_bytes": 10271390853
                },
                {
                  "bytes": 2190906161,
                  "central_directory_members": 49866,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr13.echtvar.zip",
                  "uncompressed_bytes": 7487102196
                },
                {
                  "bytes": 2082208118,
                  "central_directory_members": 46072,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr14.echtvar.zip",
                  "uncompressed_bytes": 7077438370
                },
                {
                  "bytes": 1966885592,
                  "central_directory_members": 42820,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr15.echtvar.zip",
                  "uncompressed_bytes": 6648377385
                },
                {
                  "bytes": 2159856477,
                  "central_directory_members": 41736,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr16.echtvar.zip",
                  "uncompressed_bytes": 7407307841
                },
                {
                  "bytes": 1890442581,
                  "central_directory_members": 41194,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr17.echtvar.zip",
                  "uncompressed_bytes": 6342630113
                },
                {
                  "bytes": 1730167927,
                  "central_directory_members": 39568,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr18.echtvar.zip",
                  "uncompressed_bytes": 5936324589
                },
                {
                  "bytes": 1612041534,
                  "central_directory_members": 29812,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr19.echtvar.zip",
                  "uncompressed_bytes": 5165403019
                },
                {
                  "bytes": 1376525969,
                  "central_directory_members": 31980,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr20.echtvar.zip",
                  "uncompressed_bytes": 4746336425
                },
                {
                  "bytes": 902589299,
                  "central_directory_members": 19514,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr21.echtvar.zip",
                  "uncompressed_bytes": 3004608164
                },
                {
                  "bytes": 945741018,
                  "central_directory_members": 18430,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/540_fields/archives/chr22.echtvar.zip",
                  "uncompressed_bytes": 3075770914
                }
              ],
              "fields_sha256": "fa1f21636fcd186308e3d070a65b0c087a8813f6763a78f8db70dc96c65f7643",
              "index_bytes": 1367925,
              "library": "/home/tian/immkv_test_nvme/gnomad/immkv",
              "manifests": [
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 1,
                  "manifest_sha256": "e6502a10540c839e768f615df8081d70b4a71e61b2a0a88afa0821d687343e4d",
                  "records": 20370890,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 2,
                  "manifest_sha256": "0be51c08eec7941cd997de60a6fef8e3f72f9f38c3a4af9d889e395a7de99ea7",
                  "records": 21892245,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 3,
                  "manifest_sha256": "4fb8ff096c20282a966b468355676dc8e476b8400d7c200f4a42bd882d587e9c",
                  "records": 17834355,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 4,
                  "manifest_sha256": "85e3fdd517978eefa71480c09b88efbfe179188da72392025e1c995802ff3e37",
                  "records": 17423639,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 5,
                  "manifest_sha256": "88624fbba2bbc84b65c8ef80d78420d0b7a35928ed8ae9ac523b9f2e8f0f91a3",
                  "records": 16290654,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 6,
                  "manifest_sha256": "bdf9562a2110e95c010b344e601f16f8a543a51dca79fa2ced4031c53896fb91",
                  "records": 15190442,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 7,
                  "manifest_sha256": "77e5990e85572cbbbfce7871d79ee2fde5f7fc18cff2fe9a8f9c9900e2d824b5",
                  "records": 14752321,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 8,
                  "manifest_sha256": "56f8047c1da4c09bdb6e15f44acd99c448132ae6230e32ef22cfa20239567b5e",
                  "records": 14073574,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 9,
                  "manifest_sha256": "48b7a1d79c9a4a057bc7513eb930b850d7354f2abe380e3768aadaf87b45714f",
                  "records": 11480614,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 10,
                  "manifest_sha256": "db671e979c2a949f8c0433112f9053bcb02dcf4bc09a30e3407cdffa9378788a",
                  "records": 12118750,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 11,
                  "manifest_sha256": "0c8f3ed0470504385960e748aba1bc8854b806e3dd5044768739ef3d3a19a7f6",
                  "records": 12345492,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 12,
                  "manifest_sha256": "f0d71794035581a52f8a9c49187e7c5d77f79f5300aa76fb147efa0afa65ff88",
                  "records": 11933435,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 13,
                  "manifest_sha256": "563bb2b82c86652d08e862b28befc5be90d24897cf921aaad34621e79744f5d1",
                  "records": 8693573,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 14,
                  "manifest_sha256": "0471ec2e516c76f6b89410c89d8825aebfa36bfc9b15378c0045f9b46f7035c7",
                  "records": 8220638,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 15,
                  "manifest_sha256": "f9210d1c3b5ca3d8b36087dcc025a44f9e9934158ef77f78d85b5bb90c391667",
                  "records": 7723667,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 16,
                  "manifest_sha256": "3cf4786236d60d3c5402f496bd1cd8cfbe2ac3a2773e32f7cb1e680ef96c9a77",
                  "records": 8618657,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 17,
                  "manifest_sha256": "7f18009ed6de25e04fe6781dd2b181ad310ca3ecadeb5d65fd5f881be16eff8b",
                  "records": 7373819,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 18,
                  "manifest_sha256": "e7d6244b50c2e3b2752e25ab393dee0d5135133f087b10a0178d83a0bdb72ccb",
                  "records": 6893751,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 19,
                  "manifest_sha256": "df905df6fec7a78b7a65108551deb1e8c29bd3a3dca06edf68f5488bf8e7d360",
                  "records": 5991870,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 20,
                  "manifest_sha256": "5ffef94f04de22a04f846f3f06f2b33d1c6e828cc099efd15abeb6d6196a1db5",
                  "records": 5516836,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 21,
                  "manifest_sha256": "869bb863c33345bfcedba5de9265f1acac3c5033460ef83cb11d432d7143cdf4",
                  "records": 3483000,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 22,
                  "manifest_sha256": "872946e081799b83f30444581a23e5d7a4d3a4f46dba7e12a5185f8109e77895",
                  "records": 3567085,
                  "schema_hash": "1e91a99722c59d5e"
                }
              ],
              "query": "/home/tian/immkv_test_nvme/test/540_fields/query/hg001.decomp.vcf.gz",
              "query_bytes": 139085621,
              "query_sha256": "21b02c83b020ef6d97ad39068ce21deb28c496aaec3ba984365b4799bc2f29da"
            },
            "568": {
              "archives": [
                {
                  "bytes": 7257511949,
                  "central_directory_members": 124848,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr1.echtvar.zip",
                  "uncompressed_bytes": 38780286499
                },
                {
                  "bytes": 7624945470,
                  "central_directory_members": 131688,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr2.echtvar.zip",
                  "uncompressed_bytes": 40899531669
                },
                {
                  "bytes": 6263830958,
                  "central_directory_members": 106608,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr3.echtvar.zip",
                  "uncompressed_bytes": 34926259821
                },
                {
                  "bytes": 6074278192,
                  "central_directory_members": 103188,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr4.echtvar.zip",
                  "uncompressed_bytes": 30168583462
                },
                {
                  "bytes": 5622847357,
                  "central_directory_members": 97488,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr5.echtvar.zip",
                  "uncompressed_bytes": 28696987035
                },
                {
                  "bytes": 5333501460,
                  "central_directory_members": 92358,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr6.echtvar.zip",
                  "uncompressed_bytes": 26290904570
                },
                {
                  "bytes": 5284983462,
                  "central_directory_members": 85518,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr7.echtvar.zip",
                  "uncompressed_bytes": 27981956434
                },
                {
                  "bytes": 4867485483,
                  "central_directory_members": 78678,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr8.echtvar.zip",
                  "uncompressed_bytes": 26040552400
                },
                {
                  "bytes": 4063797988,
                  "central_directory_members": 67848,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr9.echtvar.zip",
                  "uncompressed_bytes": 19890246125
                },
                {
                  "bytes": 4299658931,
                  "central_directory_members": 72978,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr10.echtvar.zip",
                  "uncompressed_bytes": 21839763121
                },
                {
                  "bytes": 4415331558,
                  "central_directory_members": 72408,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr11.echtvar.zip",
                  "uncompressed_bytes": 24849737455
                },
                {
                  "bytes": 4324593804,
                  "central_directory_members": 71838,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr12.echtvar.zip",
                  "uncompressed_bytes": 24288159769
                },
                {
                  "bytes": 2982695276,
                  "central_directory_members": 52458,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr13.echtvar.zip",
                  "uncompressed_bytes": 13200835719
                },
                {
                  "bytes": 2978458193,
                  "central_directory_members": 48468,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr14.echtvar.zip",
                  "uncompressed_bytes": 16672855058
                },
                {
                  "bytes": 2821168599,
                  "central_directory_members": 45048,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr15.echtvar.zip",
                  "uncompressed_bytes": 16007926228
                },
                {
                  "bytes": 3147128737,
                  "central_directory_members": 43908,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr16.echtvar.zip",
                  "uncompressed_bytes": 18820464097
                },
                {
                  "bytes": 2804825536,
                  "central_directory_members": 43338,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr17.echtvar.zip",
                  "uncompressed_bytes": 17740105448
                },
                {
                  "bytes": 2434370983,
                  "central_directory_members": 41628,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr18.echtvar.zip",
                  "uncompressed_bytes": 12818068081
                },
                {
                  "bytes": 2393428912,
                  "central_directory_members": 31368,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr19.echtvar.zip",
                  "uncompressed_bytes": 14771573347
                },
                {
                  "bytes": 1939275218,
                  "central_directory_members": 33648,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr20.echtvar.zip",
                  "uncompressed_bytes": 10042845922
                },
                {
                  "bytes": 1259700453,
                  "central_directory_members": 20538,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr21.echtvar.zip",
                  "uncompressed_bytes": 6372577608
                },
                {
                  "bytes": 1362922300,
                  "central_directory_members": 19398,
                  "central_directory_valid": true,
                  "path": "/home/tian/immkv_test_nvme/test/568_fields/archives/chr22.echtvar.zip",
                  "uncompressed_bytes": 7774267602
                }
              ],
              "fields_sha256": "28e7720f959b1cf583af444f1290a9bad857e3d733ea8cd6d0b75442b78ebf86",
              "index_bytes": 1367925,
              "library": "/home/tian/immkv_test_nvme/gnomad/immkv",
              "manifests": [
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 1,
                  "manifest_sha256": "e6502a10540c839e768f615df8081d70b4a71e61b2a0a88afa0821d687343e4d",
                  "records": 20370890,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 2,
                  "manifest_sha256": "0be51c08eec7941cd997de60a6fef8e3f72f9f38c3a4af9d889e395a7de99ea7",
                  "records": 21892245,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 3,
                  "manifest_sha256": "4fb8ff096c20282a966b468355676dc8e476b8400d7c200f4a42bd882d587e9c",
                  "records": 17834355,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 4,
                  "manifest_sha256": "85e3fdd517978eefa71480c09b88efbfe179188da72392025e1c995802ff3e37",
                  "records": 17423639,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 5,
                  "manifest_sha256": "88624fbba2bbc84b65c8ef80d78420d0b7a35928ed8ae9ac523b9f2e8f0f91a3",
                  "records": 16290654,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 6,
                  "manifest_sha256": "bdf9562a2110e95c010b344e601f16f8a543a51dca79fa2ced4031c53896fb91",
                  "records": 15190442,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 7,
                  "manifest_sha256": "77e5990e85572cbbbfce7871d79ee2fde5f7fc18cff2fe9a8f9c9900e2d824b5",
                  "records": 14752321,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 8,
                  "manifest_sha256": "56f8047c1da4c09bdb6e15f44acd99c448132ae6230e32ef22cfa20239567b5e",
                  "records": 14073574,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 9,
                  "manifest_sha256": "48b7a1d79c9a4a057bc7513eb930b850d7354f2abe380e3768aadaf87b45714f",
                  "records": 11480614,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 10,
                  "manifest_sha256": "db671e979c2a949f8c0433112f9053bcb02dcf4bc09a30e3407cdffa9378788a",
                  "records": 12118750,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 11,
                  "manifest_sha256": "0c8f3ed0470504385960e748aba1bc8854b806e3dd5044768739ef3d3a19a7f6",
                  "records": 12345492,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 12,
                  "manifest_sha256": "f0d71794035581a52f8a9c49187e7c5d77f79f5300aa76fb147efa0afa65ff88",
                  "records": 11933435,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 13,
                  "manifest_sha256": "563bb2b82c86652d08e862b28befc5be90d24897cf921aaad34621e79744f5d1",
                  "records": 8693573,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 14,
                  "manifest_sha256": "0471ec2e516c76f6b89410c89d8825aebfa36bfc9b15378c0045f9b46f7035c7",
                  "records": 8220638,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 15,
                  "manifest_sha256": "f9210d1c3b5ca3d8b36087dcc025a44f9e9934158ef77f78d85b5bb90c391667",
                  "records": 7723667,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 16,
                  "manifest_sha256": "3cf4786236d60d3c5402f496bd1cd8cfbe2ac3a2773e32f7cb1e680ef96c9a77",
                  "records": 8618657,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 17,
                  "manifest_sha256": "7f18009ed6de25e04fe6781dd2b181ad310ca3ecadeb5d65fd5f881be16eff8b",
                  "records": 7373819,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 18,
                  "manifest_sha256": "e7d6244b50c2e3b2752e25ab393dee0d5135133f087b10a0178d83a0bdb72ccb",
                  "records": 6893751,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 19,
                  "manifest_sha256": "df905df6fec7a78b7a65108551deb1e8c29bd3a3dca06edf68f5488bf8e7d360",
                  "records": 5991870,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 20,
                  "manifest_sha256": "5ffef94f04de22a04f846f3f06f2b33d1c6e828cc099efd15abeb6d6196a1db5",
                  "records": 5516836,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 21,
                  "manifest_sha256": "869bb863c33345bfcedba5de9265f1acac3c5033460ef83cb11d432d7143cdf4",
                  "records": 3483000,
                  "schema_hash": "1e91a99722c59d5e"
                },
                {
                  "build_id": "a6fa6f081196a359aec12e89ccdf8bcc",
                  "chromosome": 22,
                  "manifest_sha256": "872946e081799b83f30444581a23e5d7a4d3a4f46dba7e12a5185f8109e77895",
                  "records": 3567085,
                  "schema_hash": "1e91a99722c59d5e"
                }
              ],
              "query": "/home/tian/immkv_test_nvme/test/568_fields/query/hg001.decomp.vcf.gz",
              "query_bytes": 139085621,
              "query_sha256": "ab247503db51111482d0b54b5cc76e3c0a94d5b2310746dd85f51d92b48f937b"
            }
          }
        },
        "query_data_sha256_from_readonly_preflight": "0250956380e1f6910a83f92a579810671e1fe4dfeb0b1d855394021a87148fd1"
      },
      "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results/input-preflight.json"
    }
  },
  "required_constraints": {
    "source": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/env_tools.py",
    "values": {
      "GOMAXPROCS": "4",
      "cpus": [
        0,
        1,
        2,
        3
      ],
      "memory.max": "17179869184",
      "memory.swap.max": "0"
    }
  },
  "results": "/home/tian/projects/immkv/test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/results",
  "schema_version": 1
}
```
