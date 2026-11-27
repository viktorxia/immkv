# Cold-cache timing summary (post-fix vs v0.0.5)

| Medium | Fields | Complete / attempts | Post-fix wall (s) | v0.0.5 wall (s) | Post-fix CPU (s) | v0.0.5 CPU (s) | Post-fix cgroup peak (bytes) | v0.0.5 cgroup peak (bytes) | Sources |
|---|---:|---:|---:|---:|---:|---:|---:|---:|---|
| nvme | 17 | 3 / 3 | 11.57 | 13.20 | 35.14 | 33.15 | 4767973376 | 4798017536 | [timing-nvme-17-immkv-r1-cold-postfix](timing-nvme-17-immkv-r1-cold-postfix/status.json), [timing-nvme-17-immkv-r2-cold-postfix](timing-nvme-17-immkv-r2-cold-postfix/status.json), [timing-nvme-17-immkv-r3-cold-postfix](timing-nvme-17-immkv-r3-cold-postfix/status.json) |
| nvme | 540 | 3 / 3 | 171.01 | 158.86 | 660.11 | 616.02 | 6947954688 | 6947467264 | [timing-nvme-540-immkv-r1-cold-postfix](timing-nvme-540-immkv-r1-cold-postfix/status.json), [timing-nvme-540-immkv-r2-cold-postfix](timing-nvme-540-immkv-r2-cold-postfix/status.json), [timing-nvme-540-immkv-r3-cold-postfix](timing-nvme-540-immkv-r3-cold-postfix/status.json) |
| nvme | 568 | 3 / 3 | 192.98 | 172.47 | 731.70 | 670.46 | 7693586432 | 7959490560 | [timing-nvme-568-immkv-r1-cold-postfix](timing-nvme-568-immkv-r1-cold-postfix/status.json), [timing-nvme-568-immkv-r2-cold-postfix](timing-nvme-568-immkv-r2-cold-postfix/status.json), [timing-nvme-568-immkv-r3-cold-postfix](timing-nvme-568-immkv-r3-cold-postfix/status.json) |
| hdd | 17 | 3 / 3 | 99.77 | 96.46 | 215.96 | 204.67 | 7498162176 | 7506931712 | [timing-hdd-17-immkv-r1-cold-postfix](timing-hdd-17-immkv-r1-cold-postfix/status.json), [timing-hdd-17-immkv-r2-cold-postfix](timing-hdd-17-immkv-r2-cold-postfix/status.json), [timing-hdd-17-immkv-r3-cold-postfix](timing-hdd-17-immkv-r3-cold-postfix/status.json) |
| hdd | 540 | 3 / 3 | 1468.42 | 1468.32 | 725.88 | 674.01 | 5346287616 | 5348339712 | [timing-hdd-540-immkv-r1-cold-postfix](timing-hdd-540-immkv-r1-cold-postfix/status.json), [timing-hdd-540-immkv-r2-cold-postfix](timing-hdd-540-immkv-r2-cold-postfix/status.json), [timing-hdd-540-immkv-r3-cold-postfix](timing-hdd-540-immkv-r3-cold-postfix/status.json) |
| hdd | 568 | 3 / 3 | 1469.80 | 1469.52 | 776.11 | 719.20 | 5499441152 | 5525741568 | [timing-hdd-568-immkv-r1-cold-postfix](timing-hdd-568-immkv-r1-cold-postfix/status.json), [timing-hdd-568-immkv-r2-cold-postfix](timing-hdd-568-immkv-r2-cold-postfix/status.json), [timing-hdd-568-immkv-r3-cold-postfix](timing-hdd-568-immkv-r3-cold-postfix/status.json) |

## Raw repetitions

| Record | Status | Wall (s) | CPU (s) | Max child RSS (KiB) | Cgroup peak (bytes) | Input blocks |
|---|---|---:|---:|---:|---:|---:|
| [timing-hdd-17-immkv-r1-cold-postfix](timing-hdd-17-immkv-r1-cold-postfix/time.txt) | ok | 99.78999999999999 | 215.76000000000002 | 255064 | 7508865024 | 17196952 |
| [timing-hdd-17-immkv-r2-cold-postfix](timing-hdd-17-immkv-r2-cold-postfix/time.txt) | ok | 99.77000000000001 | 216.26 | 255796 | 7497940992 | 17197104 |
| [timing-hdd-17-immkv-r3-cold-postfix](timing-hdd-17-immkv-r3-cold-postfix/time.txt) | ok | 99.47 | 215.95999999999998 | 280780 | 7498162176 | 17196960 |
| [timing-hdd-540-immkv-r1-cold-postfix](timing-hdd-540-immkv-r1-cold-postfix/time.txt) | ok | 1469.02 | 726.94 | 685620 | 5349715968 | 53852128 |
| [timing-hdd-540-immkv-r2-cold-postfix](timing-hdd-540-immkv-r2-cold-postfix/time.txt) | ok | 1468.42 | 725.88 | 679312 | 5342957568 | 53852096 |
| [timing-hdd-540-immkv-r3-cold-postfix](timing-hdd-540-immkv-r3-cold-postfix/time.txt) | ok | 1468.1 | 724.9899999999999 | 682192 | 5346287616 | 53852120 |
| [timing-hdd-568-immkv-r1-cold-postfix](timing-hdd-568-immkv-r1-cold-postfix/time.txt) | ok | 1469.31 | 773.52 | 858900 | 5499441152 | 53852168 |
| [timing-hdd-568-immkv-r2-cold-postfix](timing-hdd-568-immkv-r2-cold-postfix/time.txt) | ok | 1469.8 | 780.87 | 863952 | 5522644992 | 53852168 |
| [timing-hdd-568-immkv-r3-cold-postfix](timing-hdd-568-immkv-r3-cold-postfix/time.txt) | ok | 1470.2 | 776.11 | 866020 | 5497843712 | 53852224 |
| [timing-nvme-17-immkv-r1-cold-postfix](timing-nvme-17-immkv-r1-cold-postfix/time.txt) | ok | 11.68 | 35.93 | 260052 | 4759334912 | 20427760 |
| [timing-nvme-17-immkv-r2-cold-postfix](timing-nvme-17-immkv-r2-cold-postfix/time.txt) | ok | 11.27 | 34.4 | 268964 | 4767973376 | 20427544 |
| [timing-nvme-17-immkv-r3-cold-postfix](timing-nvme-17-immkv-r3-cold-postfix/time.txt) | ok | 11.57 | 35.14 | 287864 | 4787634176 | 20427816 |
| [timing-nvme-540-immkv-r1-cold-postfix](timing-nvme-540-immkv-r1-cold-postfix/time.txt) | ok | 174.0 | 669.68 | 2338604 | 6945189888 | 53638576 |
| [timing-nvme-540-immkv-r2-cold-postfix](timing-nvme-540-immkv-r2-cold-postfix/time.txt) | ok | 171.01 | 660.11 | 2340680 | 6948024320 | 53638600 |
| [timing-nvme-540-immkv-r3-cold-postfix](timing-nvme-540-immkv-r3-cold-postfix/time.txt) | ok | 170.79 | 658.7099999999999 | 2341140 | 6947954688 | 53638600 |
| [timing-nvme-568-immkv-r1-cold-postfix](timing-nvme-568-immkv-r1-cold-postfix/time.txt) | ok | 191.29 | 729.71 | 3053680 | 7679475712 | 53638792 |
| [timing-nvme-568-immkv-r2-cold-postfix](timing-nvme-568-immkv-r2-cold-postfix/time.txt) | ok | 192.98 | 734.9300000000001 | 3373044 | 8007254016 | 53638800 |
| [timing-nvme-568-immkv-r3-cold-postfix](timing-nvme-568-immkv-r3-cold-postfix/time.txt) | ok | 193.16 | 731.6999999999999 | 3067360 | 7693586432 | 53638800 |
