# immkv 仓库索引

面向需要在本仓库工作的旗舰模型：先读本文件，再按需打开少数目标文件，避免
整仓通读。索引内容对应 tag `v0.0.6` 的引擎提交 `389697d`（2026-11-27）；
之后只加实验记录或纯文档提交时，源码锚点以 `grep -n` 校正。
§9–§12 是 2026-11-25 增补、2026-11-27 按 `v0.0.6` 复核的审阅层：证据等级、
文档自认的限制、文档矛盾、可攻击面。

- 行号是当前提交的锚点；符号名比行号稳定，改动后以 `grep -n` 校正。
- 未覆盖：`tmp/`（构建产物、JDK、二进制）、`test/*/results/` 的逐文件原始
  证据、`test/*/sources/` 的 echtvar 源码副本。它们是数据，不是设计文档。
- 仓库跟踪文件约 4,559 个，其中设计/源码约 3.2 万行，其余是实验结果 JSON/TSV。

## 0. 最短路径

| 你要做的事 | 读这些 |
|---|---|
| 了解引擎是什么、能做什么 | §1、`README.md` 的 Philosophy/API/Contracts |
| 审阅项目、找潜在问题 | §9–§12，再按 §13 下钻 |
| 改盘上格式或读写实现 | §5、`docs/v003-format.md`、§4.2 的 `shard.go`/`stream.go`/`colread.go` |
| 改建库入口 | §6.3、`docs/ingest-design.md`、`ingest.go` |
| 查一个实验的数字 | §7.3 的表，再开对应 `test/<exp>/README.md` |
| 弄清历史决策为什么这样 | §3 的 handoff 条目，按日期顺序 |
| 排查 `ErrUnsorted` / 坏库 / OOM | §14 |

## 1. 一页事实

| 项 | 值 |
|---|---|
| 模块 | `github.com/viktorxia/immkv`，Go 1.26.0 |
| 形态 | 构建一次、不可变、排序 KV；VCF 查询特化，非通用 KV 存储 |
| 依赖 | `klauspost/compress`（纯 Go zstd）、`valyala/gozstd`（CGO libzstd）、`golang.org/x/sys`；`biogo/hts` 仅 `test/` 使用 |
| 库文件 | `base.keys` + `base.vals`；VCF/TSV 构建另产 `base.col` + `base.manifest` |
| 魔数 | `IMKVK002`（键流）、`IMKVV003`（值）、`IMKVC003`（列存） |
| 句柄 | `OpenEager`（点查/前缀/区间）、`OpenStreaming`（`GetBatch` 批量）、`OpenSimpleKV`（单库 `Get`） |
| 并发 | 句柄内多 goroutine 安全；多进程可同时只读打开，无锁 |
| 平台 | Linux 全功能（io_uring/O_DIRECT/fadvise）；其他平台降级为 pread |
| 当前版本 | `v0.0.6` → `389697d`；`9e25c92` 修复 `io_uring` 在途限流与列流跳读估算 |

## 2. 概念表（README.md 是权威）

| 概念 | 含义 | 出处 |
|---|---|---|
| 微块（micro-block） | 键流里 64 条记录一组：`keyEnds` + `lens` + `baseValOff` + 键区 | README "File layout"；`immkv.go` 顶部注释 |
| 键流（key stream） | `[64B 头][微块×N][块索引]`；块索引打开时全量常驻 | 同上 |
| 值帧（value frame） | `vals` 里一条或多条值的压缩单元，三种形态见 §5.3 | 同上 |
| 打包模式（packed mode） | 小值按 `packTarget` 攒组进一个 zstd 帧，组头带长度表 | `CreateOptions.PackTarget` 默认 64KB |
| 每值成帧（per-value framing） | `PackTarget < 0`：每条值独立成帧，构建期训练 zstd 字典嵌入 `vals` 头 | 同上；gnomAD 大值库用此模式 |
| 点查句柄（point-lookup handle） | `OpenEager`：块索引常驻，按需读微块与值帧，带 LRU | `shard.go:696` |
| 流式句柄（streaming handle） | `OpenStreaming`：整份有序查询从头到尾，`GetBatch` 批量；不缓存 | `stream.go:50` |
| `Threads` 语义 | 每次 `GetBatch` 内解压并行度；io_uring 路径上不决定在途读深度 | `README.md` FAQ；`StreamOptions.Threads` |
| `ErrUnsorted` | `GetBatch` 输入键不按完整编码键字节序升序时返回 | `stream.go:27` |
| 变体键（`VarKey`/`PosKey`/`ChromID`） | 染色体号 + 大端位置 + REF/ALT 的排序键编码 | `key.go` |
| schema | 有序字段表（name/Type/Number），顺序即位图位号 | `docs/v003-format.md` §3 |
| record | `vals` 一条值 = 存在位图 + typed payload | §5.5 |
| 列流（column stream） | `col` 里某列在某个 bin 内的位图 + 值序列 | §5.6 |
| 规范化文本 | 由 typed 值渲染的 INFO 文本，不是源文本 | `RenderRecord` |

## 3. 文档地图

| 文件 | 行数 | 内容 | 状态 / 何时读 |
|---|---:|---|---|
| `README.md` | 372 | 理念、API、`GetBatch` 排序 FAQ、染色体拼写、性能摘要、契约、文件布局、读路径、codec | **权威**：概念与公开契约 |
| `README.cn.md` | 293 | 上者的中文版 | 基本同步；`ReadManifest` 一处未跟上 `ReadSchema`（`README.cn.md:58`） |
| `AGENTS.md` / `AGENTS.cn.md` | 199 / 106 | 工程硬约束、`test/` 实验规则、提交与 agent 行为规则 | **强制**：改格式/读路径/实验前必读 |
| `docs/v003-format.md` | 303 | v003 格式契约：schema、record、col、manifest、构建入口、codec、错误、验收 | **权威**：盘上字节 |
| `docs/ingest-design.md` | 162 | `BuildFromVCFs` / `BuildFromTSVs` 行为契约与实现结构 | **权威**：建库语义 |
| `docs/handoff/2026-11-07.md` | 98 | echtvar 三组实验总账、受控口径、数据路径、immkv_anno 立项备忘 | 历史基线；digit 来源 |
| `docs/handoff/2026-11-08.md` | 52 | `OpenSimpleKV` 分支状态；两个待修 bug | 两个 bug 已修，保留决策背景 |
| `docs/handoff/2026-11-10.md` | 53 | 强制 `##INFO` 序、`base.col` 决策、旧库重建 | 部分被 v003 取代；引用的 `docs/base-col-format.md` 已不存在 |
| `docs/handoff/2026-11-11-controlled-4c16g.md` | 84 | 4C16G 矩阵数字、v003 方向与体积抽样 | 历史 |
| `docs/handoff/2026-11-11.md` | 54 | 列存落地、全库重建、`v0.0.2` 发布 | 历史 |
| `docs/handoff/2026-11-12.md` | 13 | v003 里程碑、库体积、点查数字 | 历史 |
| `docs/handoff/2026-11-13.md` | 86 | `ReadSchema`、examples、解析开销测量、未实现的优化 | 优化候选仍有效 |
| `docs/handoff/2026-11-15.md` | 50 | `v0.0.4`、`v0.0.5-perf-untested.1` 状态 | 历史 |
| `docs/handoff/2026-11-15-performance-pending.md` | 150 | 冷缓存协议、数据/工具路径、校验器要求 | 协议部分仍是操作手册 |
| `docs/handoff/2026-11-16-performance-complete.md` | 83 | 4C16G 冷缓存六场景完成、`v0.0.5`、限制 | 历史 |
| `docs/handoff/2026-11-27.md` | 24 | `v0.0.6` 两项修复、六场景验收、提交与标签 | **最新状态**（本地，未入库） |
| `docs/eureka/2026-11-09-hdd-columnar-bundle.md` | 353 | HDD 列存 bundle 的推理与体积账；§12 有实测 | 历史设计；文本列方案已被 typed col 取代，但推理仍可读 |

注意：`docs/handoff/` 与 `docs/eureka/` 被 `.gitignore` 排除，只在本地存在；
它们不是发布物，引用前核对日期与后续结论。

## 4. 源码地图

根目录 24 个非测试源文件、13 个测试文件。

### 4.1 构建与写侧

| 文件:行 | 符号 | 作用 |
|---|---|---|
| `immkv.go:52` | 常量 | 魔数、64B 头、每块 64 条、三种帧标记位、上限 |
| `immkv.go:79` | `CreateOptions` | `RawValueBelow` / `PackTarget` / `Schema` |
| `immkv.go:128` | `Writer` | 攒集、微块组装、字典、发布 |
| `immkv.go:183` | `Create` | 拒绝已存在的 `.keys/.vals/.col/.manifest` 及 `.tmp` |
| `immkv.go:252/283` | `Add` / `AddSorted` | 任意序全量排序 / 组升序流式；两模式不可混用 |
| `immkv.go:399` | `trainDict` | 每值模式训练字典（5000 样本 / 512KB） |
| `immkv.go:451` | `flushPacked` | 打包帧组头与长度表 |
| `immkv.go:555/581` | `prepareValue` / `emitRec` | 裸存/打包/逐值选择；写块内记录 |
| `immkv.go:650/669` | `Finish` / `Abort` | 临时文件 + fsync + rename；失败清理一次 |
| `immkv.go:685` | `finish` | 排序、块索引、回填头、先发布 vals 再 keys |
| `buildfile.go:11` | `buildFiles` | 记录本次创建路径；`publish` 拒绝覆盖；`cleanup` 只消费一次 |
| `col.go:62` | `colWriter` | `base.col` 写入：1Mb bin、每列位图 + 值流 |
| `col.go:357` | `writeColumnar` | `Finish` 后顺序扫 `keys/vals`，不读第二遍 VCF |
| `col.go:478` | `seqValReader` | 顺序读记录序值帧，供列存构建 |

### 4.2 读侧

| 文件:行 | 符号 | 作用 |
|---|---|---|
| `shard.go:123` | `shard` | 单染色体：两个文件、块索引常驻、解码器、可选 LRU |
| `shard.go:137` | `openShard` | 校验魔数/版本/记录数/块索引区；坏头报错不 panic |
| `shard.go:322/351` | `parseBlockView` / `parseBlock` | 微块零分配视图 / 带 `cum` 的随机访问 |
| `shard.go:413` | `blockFor` | 按块首 8B 前缀二分 |
| `shard.go:425` | `seek` | 块间 + 块内二分，返回首个 ≥ key 的位置 |
| `shard.go:443` | `get` | 点查；重复键取排序第一条 |
| `shard.go:470/500` | `getPrefix` / `scan` | 前缀 / 区间回调；`false` 停止不是错误 |
| `shard.go:548` | `decodePack` | 解析打包帧组头并建前缀索引 |
| `shard.go:583` | `value` | 读值帧 / 解压 / LRU；打包帧按记录号取成员 |
| `shard.go:636` | `shardSet` | 25 个染色体槽位；每 shard 独立 LRU |
| `shard.go:696` | `OpenEager` | 打开全部存在的染色体；`cacheMB` 均分 |
| `stream.go:50` | `OpenStreaming` | 打开记录序 + 可选列存；校验 build id/schema 一致 |
| `stream.go:144` | `GetBatch` | 三级流水线；先全量校验升序 |
| `stream.go:398` | `mergeJoin` | 键流窗口顺扫 + 块内二分，产出 `sHit` |
| `stream.go:511` | `buildExtents` | 相邻命中合并（≤4KB，≤4MB/段） |
| `stream.go:540/558` | `readChunk` / `decodeChunk` | 读段读取 / 解压并回填 |
| `stream.go:576` | `decodeExtent` | 单段解码、投影、复制到调用方 slab |
| `stream.go:313` | `preferColumnar` | 成本模型二选一（含 `colRanges * randSec`）；见 §6.2 |
| `colread.go:37` | `openColShard` | 打开 `base.col` + `manifest`，缺一报错 |
| `colread.go:328` | `estimate` | 选中列压缩/原始字节与跳读段数估算 |
| `colread.go:404/482` | `loadBin` / `processBin` | bin 级预取 / 解压装配 |
| `colread.go:662` | `readColBatch` | 列存路径回填；bin 级预取隐藏 I/O |
| `projection.go:39` | `appendRecord` | 投影：保留完整位图，未选字段跳过但校验 |
| `simple.go:22/37` | `OpenSimpleKV` / `Get` | 单库精确点查；解包单 String 字段 |
| `key.go:15/35` | `ChromID` / `ChromName` | 染色体名 ↔ 首字节；自然序 |
| `key.go:50/62` | `VarKey` / `PosKey` | `[cid][pos BE][ref][0x00][alt]` / `[cid][pos BE]` |

### 4.3 ingest 与 v003 codec

| 文件:行 | 符号 | 作用 |
|---|---|---|
| `ingest.go:44` | `BuildFromTSVs` | 无表头 TSV → 单库；键随便字节升序 |
| `ingest.go:135` | `BuildFromVCFs` | VCF(s) → 分染色体库 + col + manifest |
| `ingest.go:240` | `prepareVCFInputs` | 头解析、跨文件一致性、字段选择、prefix |
| `ingest.go:296/319` | `schemaFromSelectedFields` / `selectValueFields` | `##INFO` 序 → schema |
| `ingest.go:414` | `openVCF` | `.gz/.bgz` 走 bgzf，其余纯文本 |
| `ingest.go:573/598` | `splitVCFLine` / `skipVCFAlt` | 前 8 列解析；跳过符号化 ALT |
| `ingest.go:677` | `vcfShardBuilder.addLine` | 按染色体推进、跳过未知 contig/`REF="."` |
| `ingest.go:759` | `parseInfo` | 保存选中字段原文（不立即解码） |
| `ingest.go:836` | `buildRecord` | `RecordEncoder` 编 typed record |
| `ingest.go:908/925` | `setIntegerArray` / `setFloatArray` | Number=R/G/. 的数组；NaN/Inf 记缺失 |
| `ingest.go:1088` | `buildImmkvHeader` | 渲染参考头，非 schema 权威 |
| `schema.go:65` | `Schema` | `Validate` / `Index` / `Hash`；hash = SHA-256 前 8B |
| `record.go:34` | `NewRecord` | 校验并返回只读视图 |
| `record.go:93` | `Record.Values` | `iter.Seq2` 按 schema 序遍历 |
| `record.go:159` | `RenderRecord` | 规范化 INFO 文本 |
| `record.go:246` | `RecordEncoder` | `Int/Float/Bytes/Flag/Values` + `Encode` |
| `record.go:477` | `validateRecord` | 位图/payload 全量校验 |
| `manifest.go:35` | `Manifest` | JSON 形态；`SetSchema` / `Validate` |
| `manifest.go:149/170` | `ReadManifest` / `ReadSchema` | shard 级诊断 / 库级 schema |
| `manifest.go:207` | `WriteManifest` | 原子更新（构建入口用禁止覆盖的 `writeBuildFile`） |

### 4.4 平台与 I/O

| 文件 | 作用 |
|---|---|
| `disk_linux.go` | `detectIOProfile`：读 `/sys/dev/block/*/queue/rotational`，dm/md 看 slaves |
| `disk_other.go` | 非 Linux 恒定 SSD 画像 |
| `direct_linux.go` | `O_DIRECT` 句柄；`IMMKV_NO_DIRECT` 可禁用 |
| `direct_other.go` | 返回不支持，回退缓冲读 |
| `fadvise_linux.go` | `fadviseDontNeed`：流式读后丢页缓存 |
| `fadvise_other.go` | 空实现 |
| `evict_linux.go` | 测试用 `fadvise(DONTNEED)` 驱逐 shard 页缓存 |
| `iouring_linux.go` | 512 深环、单驱动 goroutine、对齐缓冲池、`driveDirect`、`inFlightIO` 在途限流；内核 ≥5.6 |
| `iouring_other.go` | 桩，返回不支持 |
| `zstdcgo_cgo.go` | libzstd CDict/DDict 路径；字典库解码主路径 |
| `zstdcgo_nocgo.go` | `CGO_ENABLED=0` 时返回错误，调用方回退 klauspost |

### 4.5 examples 与测试工具

| 路径 | 作用 |
|---|---|
| `examples/build` | 小 VCF → 库，逐参数注释 `BuildFromVCFs`，末尾 `ReadSchema` |
| `examples/query` | `OpenEager.Get/Scan` + `OpenStreaming.GetBatch`，只返回 raw |
| `examples/parse` | 同一条 record 的三种解析：`RenderRecord`、`NewRecord+Values`、present 位图；三态判断 |
| `examples/batchformat` | 复用 buffer + 字段索引直接拼 INFO，对应 annotate 的高性能写法 |
| `test/annotate` | 端到端注释：查 VCF → `GetBatch` → 渲染 → 写出；`-fields` 白名单兼输出序 |
| `test/buildfull` | 全字段 gnomAD 库（`BuildFromVCFs` 薄封装） |
| `test/buildnarrow` | 窄库（字段白名单）；`NARROW_ONLY` 限一条染色体 |
| `test/buildvcf` | 通用构建壳；`IMMKV_RAW` / `IMMKV_PACK_KB` 可覆盖帧模式 |
| `test/verifyvalues` | 全量逐行逐字段比对；浮点容差、Flag 存在性、对手缺陷分类 |
| `test/pointlookup` | `OpenEager` 随机点查 p50/p99/QPS；miss 用键尾 0xFF 构造 |
| `test/samplecolumns` | 只读抽样：INFO 顺序、十进制往返、按列体积 |
| `test/genfields.py` | 从 gnomAD 头生成 `fields.tsv` 与 echtvar encode JSON |

## 5. 盘上格式速查

细节以 `docs/v003-format.md` 为准；这里只列定位用的骨架。

### 5.1 `base.keys`

```text
[64B 头][微块 ×N][块索引]
头：magic IMKVK002(8) | version u32=1 | count u64 | recsPerBlock u16=64 |
    nBlock u32 | indexOff u64 | rawBelow u32 | padding
微块：keyEnds u32×n | lens u32×n | baseValOff u64 | 键区
lens：bit31 裸存 | bit30 打包帧 | bit29 同帧 | 低 29 位帧长
索引：首键 8B 前缀 u64 大端 ×N | 块偏移 u64 小端 ×N
```

`keys` 头里的 `version` 是 1，而魔数是 `IMKVK002`：版本指头布局，魔数指
微块格式，二者不冲突。

### 5.2 `base.vals`

```text
[64B 头][字典（仅每值模式）][值帧 ×N]
头：magic IMKVV003(8) | version u32=3 | count u64 | dictLen u32 | dictOff u64
```

### 5.3 值帧三形态

| 形态 | 触发 | 内容 |
|---|---|---|
| 裸存 | 值长 < `RawValueBelow`（默认 64） | 原字节，无压缩 |
| 打包 | `0 <= PackTarget`，值 < 32KB，按 64KB 目标攒组 | 组头 `firstRec u32` + `n u32` + `lens u32×n` + 值区 |
| 每值 | `PackTarget < 0` | 每条值一个 zstd 帧 + 构建期字典 |

同帧位（bit29）只用于打包组成员：`valOff` 与上一条相同，不重复累加帧长。

### 5.4 `base.col`

```text
[64B 头][bin 目录 32B×N][列目录 32B×(N×列数)][列数据帧]
头：IMKVC003 | version 3 | header_size 64 | build_id 16B | bin_width |
    bin_count | column_count | column_entry_size 32 | bin_dir_offset | schema_hash
bin 目录：first_record u64 | record_count u32 | reserved | column_dir_offset u64 | data_end u64
列目录：bitmap_offset | value_offset | bitmap_len | value_len | present_count | raw_len
列流：位图 ceil(records/8) LSB-first（zstd）+ 有值记录的 payload 流（zstd）
```

Flag 列只有位图；全缺 bin 的两个偏移为 0；`raw_len` 只用于选路估算。

### 5.5 v003 record

```text
bitmap：ceil(N/64) × uint64 小端，字内 LSB-first
payload：按 schema 顺序，只写有值字段
Integer=zigzag uvarint | Float=8B IEEE754 小端 | String=uvarint(len)+bytes
Flag=无 payload | 数组=uvarint(count)+每个 uvarint(len)+element（len 0 = 缺失元素）
```

### 5.6 `base.manifest`

JSON：`format=immkv-v003`、`version=3`、`base`、`build_id`、`bin_width`、
`record_count`、三个字节数、`schema_hash`、`fields[{name,column,type,number}]`、
`source_info_order{ok,bad}`、`non_finite_floats`。打开时校验 magic、版本、
build_id、schema hash、字节数、记录数，任一不符大声失败。

## 6. 调用链

### 6.1 点查 `Get`

`DB.Get` → 按键首字节路由 → `shard.get` → `blockFor`（常驻索引二分）→
`block`（pread 微块，LRU）→ 块内二分 → `value`（读值帧，LRU 解压）。

缓存命中时零系统调用；返回切片只保证到下一次查询前。

### 6.2 批量 `GetBatch`

1. 全量校验 `bytes.Compare(keys[i-1], keys[i]) <= 0`，否则 `ErrUnsorted`。
2. 按染色体切段（键升序 ⇒ 同染色体连续）。
3. `mergeJoin`：256KB 窗口顺扫键流，块内二分，命中产出 `sHit`（valOff/frameLen/raw/packed/rec/dst）。
4. 有 `base.col` 时 `preferColumnar` 估算（`v0.0.6` 起列存计入
   `colRanges × randSec`）：
   - HDD：0.45ms/随机读，170MB/s 顺序；
   - SSD：10µs/随机读，3GB/s 顺序；
   - 记录序解码 250MB/s/线程，列存解码 45MB/s/线程，typed 压缩比 4.3。
5. 列存胜 → `readColBatch`；否则 `buildExtents` → `chunkExtents`。
6. 读：io_uring 单驱动（O_DIRECT 4KB 对齐，512 深）或 Threads 路 pread 池。
7. 解压池：`decodeChunk`，带投影时在 worker 内投影并一次拷入调用方 slab。

调试开关：`IMMKV_FORCE_COLUMNAR`、`IMMKV_FORCE_RECORD`、`IMMKV_DEBUG_PATH`。

### 6.3 建库 `BuildFromVCFs`

1. `prepareVCFInputs`：解析头、校验多文件 `##INFO` 顺序一致、生成 schema。
2. 单文件顺序读；多文件一文件一 goroutine，并发上限 `GOMAXPROCS/2`。
3. 逐行：跳过未知 contig、`REF="."`、符号化 ALT；每个可建 ALT 一条 record。
4. `AddSorted(PosKey, VarKey, record)`：同 POS 成组，组末按完整键排序。
5. `Finish` 后 `writeColumnar` 顺序扫 `keys/vals`，写 `col` + `manifest`。
6. 全部成功后写 `vcf.raw.header` / `vcf.immkv.header`；失败删除本次创建的 `outdir`。

## 7. 测试与实验地图

### 7.1 单元测试

| 文件 | 覆盖 |
|---|---|
| `immkv_test.go` | Writer 往返、排序错误、裸存阈值、打包、Abort、拒绝覆盖 |
| `shard_test.go` | 键编码、Get/GetPrefix/Scan、坏魔数/环索引、并发混合读 |
| `stream_test.go` | GetBatch 往返、乱序、重复键、Threads 等价、坏帧 |
| `simple_test.go` | SimpleKV 边界、缓存、重复键、损坏头/帧、per-value |
| `ingest_test.go` | VCF/TSV 契约、字段选择、prefix、数组、bgzf、失败清理 |
| `col_test.go` | 列存构建/读取、列选择、present_count、尾部截断、旧魔数 |
| `colread_projection_test.go` | 列存投影跨段、重复列、合法尾部、位图填充 |
| `projection_test.go` | 投影矩阵、缺失值、未选损坏校验、fuzz、全列 |
| `record_test.go` | 编解码、Flag、位图边界、数组、渲染、损坏、分配 |
| `schema_test.go` | schema 校验/hash/索引/数组判定 |
| `manifest_test.go` | manifest 往返、旧格式拒绝、`ReadSchema` |
| `buildfile_test.go` | 清理只一次、拒绝覆盖、列存失败清理 |
| `iouring_linux_test.go` | 对齐缓冲、mmap 释放、部分提交、读错误 |

### 7.2 共享工具（`test/`）

`annotate` 是端到端注释器：主线程读行/建键/`GetBatch`，值拷入窗口 arena，
再由 `-threads` 个格式化 worker 处理，collector 顺序写出。`-fields` 是白名单
兼输出序；`.gz` 输出走 bgzf。

`verifyvalues` 的规则：Integer/String 精确；Float 容差 1.1e-6（echtvar 量化
截断 + 打印截断）；Flag 只比存在性；对手的 -1 占位、i32 溢出单列计数。

### 7.3 实验目录

命名规律：`vs_<对手>_<参考基因组>_<数据集>_<口径>`。数字以各目录
`README.md` 与 `results/` 为准。

| 目录 | 口径 | 关键结论 |
|---|---|---|
| `sample_columnar_hg19_gnomad_genomes_chr22` | chr22 抽样 | 3,567,085 keys；raw-text 列存 3.77GB（0.879× vals）；Float 全因科学计数法无法无损十进制；长度表占 52% |
| `vs_echtvar_hg19_gnomad_genomes_17_fields` | NVMe，10G | echtvar 17.76/18.04s；immkv 窄库 29.74/35.17s；CPU 42 vs 109s；66,873,138 次比较 0 差异 |
| `..._17_fields_hdd` | HDD，10G | echtvar 23.60/19.53s；全库 31:50；窄库 1:34；列存全库 3:30 |
| `..._540_fields` | NVMe，10G | immkv 5:28/5:15 vs echtvar 6:14/6:52；null 43/47s vs 6:43；808 次对手 i32 溢出 |
| `..._540_fields_hdd` | HDD，10G | echtvar 8:17 vs immkv 45:47；列存 23:37（1.94×）；null 列存 14:58 |
| `..._568_fields` | NVMe，10G | echtvar 在 4/2/1 进程全部 OOM（无限制 24.1GB/进程）；immkv 5:06 全输出、38.8s null、3.9GB；2,234,349,552 次比较 |
| `..._568_fields_hdd` | HDD，10G | 无 echtvar 臂；immkv 44:38/29:49；列存 25:11/17:10 |
| `..._4c16g_gz` | 4C16G，gz 计时 | v0.0.2：NVMe 17 26.83 vs 24.49；540 151 vs 716；HDD 540 2093 vs 664；568 echtvar 不可运行 |
| `..._4c16g_gz_v003` | 4C16G，v003 | NVMe 17 8.72 vs 27.54（帧模式不同，非格式对比）；540 164 vs 649；HDD 540 1640 vs 653 |
| `vs_echtvar_hg19_hg001_4c16g_v003_release_20261115` | v0.0.3，4C16G，六格 | NVMe 17 8.17 vs 24.52；540 157.57 vs 606.40；HDD 17 64.66 vs 24.53；540 1416 vs 622；568 仅 immkv 完成；8,850,856,500 次比较 |
| `vs_echtvar_hg19_hg001_4c16g_cold_v005_20261115` | v0.0.5，4C16G，清缓存 | NVMe 17 13.20 vs 32.27；540 158.86 vs 619.24；568 172.47 vs OOM；HDD 17 96.46 vs 27.60；540 1468 vs 641 |
| `vs_immkv_hg19_hg001_4c16g_cold_postfix_20261125` | v0.0.6，4C16G，清缓存，仅 immkv | NVMe 17 11.57；540 171.01；568 192.98；HDD 17 99.77；540 1468.42；568 1469.80；88.5 亿字段 0 差异 |
| `vs_immkv_hg19_hg001_10g_release_20261114` | 回归测试 | 发现显式 `Columns` 批量回归（full540 +242%）；该候选未发布 |
| `vs_immkv_hg19_hg001_4c16g_projection_20261115` | 投影优化 | chr22 常驻 full568 847.6→140.8ms；WGS narrow17 4.844→3.987s；旧/后输出逐字节一致 |

实验协议要点：
- 默认 10G cgroup + 禁 swap；本仓库后期的 4C16G 实验是用户明确覆盖，不能
  把两套口径的数字混在一起。
- 两轮制：第一轮完整输出证一致，第二轮 `/dev/null` 只计查询。
- 大工件在 `~/immkv_test_nvme/` 与 `/mnt/data6t/tian/immkv_test_hdd/`，不入库。
- `test/` 内容一律英文；旧记录不覆盖，新实验另建目录；`v0.0.6` 的修复后
  对照见 `vs_immkv_hg19_hg001_4c16g_cold_postfix_20261125`。

## 8. 当前状态与未了事项

- `v0.0.6` 已同时位于 `dev` 与 `main`（引擎提交 `389697d`）；tag `v0.0.6`
  指向它；`v0.0.5-perf-untested.1` 保留未动。
- `9e25c92` 修复 `io_uring` 在途请求限流与列流跳读选路估算，引擎不再停在
  `d2eac35`；`389697d` 只加实验记录。
- 未创建 GitHub Release；`docs/handoff/` 与 `docs/eureka/` 本地忽略。
- 17 字段打包版库按用户决定未重建；因此 v0.0.2 与 v003 的 17 字段数字不可
  当作格式对比。
- `docs/handoff/2026-11-13.md` 列了四个未实现的解析优化（`Columns` 选字段、
  trusted `Record`、`ValuesSelected`、索引感知遍历），新增 API 需先设计确认。
- `README.cn.md` 的 schema 读取入口仍写 `ReadManifest`，英文版已改
  `ReadSchema`。
- `main` 由 `dev` 快进合并，尚无独立发布记录；`docs/agent-index.md` 从
  `v0.0.6` 起纳入版本库维护。

## 9. 证据等级与主张台账

审阅时先分清“测过”与“推过”。

| 等级 | 含义 |
|---|---|
| A | 受控实测：固定查询、线程、内存上限；数字有 `results/` 源；正确性有逐行逐字段校验 |
| B | 实测但口径不全：单次、缓存未隔离、版本或帧模式混比、无 cgroup，占其一 |
| C | 推断或外推：设计估算、预测，未实测 |
| D | 仓库外事实：外部仓库、对手源码、本机环境；本仓库无法复核 |

| 主张 | 等级 | 依据 | 备注 |
|---|---|---|---|
| v003 record 与 echtvar 值一致（17/540/568，6 组，8,850,856,500 次比较） | A | cold v005 与 v003 release 的 README | Float 容差 1.1e-6；808 次溢出单列 |
| v0.0.6 六场景输出与 v0.0.5 逐字节一致（88.5 亿字段，0 差异） | A | `vs_immkv_hg19_hg001_4c16g_cold_postfix_20261125` README | 4C16G 冷缓存；`data_sha256` 相等 |
| v0.0.6 相对 v0.0.5：NVMe 17 快 12.3%，540/568 慢 7.6%/11.9% | A（限该协议） | 同上 `results/final-summary.md` | 单机、清缓存、仅 immkv |
| 568 字段可在 10G / 16G 内完成 | A | `..._568_fields` README；cold v005 | 10G 峰值 3.9GB；echtvar 两种预算都 OOM |
| v0.0.3：NVMe 540 快 3.85×、HDD 540 慢 2.28× | A（限该协议） | v003 release README | 缓存未隔离，单机单盘 |
| v003 比 v0.0.2：NVMe 慢 8–10%，HDD 快 21–24% | B | `..._4c16g_gz_v003` README | 4C16G；缓存状态未隔离 |
| 17 字段列存约 21s，接近 echtvar 18s | B | `..._17_fields` README 的 Columnar round | 单次强制列存，10G NVMe |
| `Threads` 4 线程是收益上限 | B | `README.md` FAQ | README 未给出该测量的源文件 |
| 选路成本模型能选对记录序/列存 | B | 各实验 README 的 path 统计；`stream.go:361` | `v0.0.6` 已计入列流跳读段数；常量仍来自单次 chr22 标定 |
| HDD 540/568 能到 8–12 分钟、17 子集 15s | C | eureka §10.5；11-10 handoff 验收线 | 实测 23:37 与 3:30，未达标 |
| 列存总体积 60–85GB | C | eureka §9.3/§10.5 | 抽样与最终 116GB 均否定 |
| echtvar i32 溢出 808 次、字符串字典退化 | A（计数）+ D（归因） | `..._568_fields` README；`sources/` | 对手源码副本在实验目录 |
| annoplex `kv-custom` 是消费方；immdb 读放大 | D | `README.md`；`AGENTS.md` | 外部仓库历史，本仓库无法复核 |
| 跨机器/跨架构可直接用 | B | `README.md` Platforms；11-15 handoff | 有跨平台构建，无跨架构读同一库的实测 |

## 10. 文档已承认的限制

只收文档自己写下的限制，不是新判断。

**格式与 API**

- 源 INFO 文本不逐字节保真，数值输出规范化（`docs/v003-format.md` §9、§15）。
- 普通 KV 只有单库精确 `Get`，无前缀、范围、批量（`README.md` "Plain KV"；`simple.go`）。
- 不做外部排序；VCF/TSV 必须预排序；不支持 stdin、BED/GTF（`docs/ingest-design.md` §7）。
- `Create` 建的库没有 `base.col`/`manifest`；传非空 `StreamOptions.Columns` 报错（`README.md`；`stream.go:89-93`）。
- 数组路径（Number=R/G/.）首轮未实测（`docs/v003-format.md` §4.2）。
- 每值模式 Float 体积为文本列的 1.27×，文档设了 1.5× 决策门（`docs/handoff/2026-11-11-controlled-4c16g.md`）。
- 不共享键流、不做字段组 bundle；HDD 双布局就是两份数据（`docs/eureka/2026-11-09-hdd-columnar-bundle.md` §5、§15）。

**性能与实验**

- v0.0.2 与 v003 的 17 字段数字不是格式对比：前者打包帧，后者每值帧（`test/vs_echtvar_hg19_gnomad_genomes_4c16g_gz_v003/README.md`）。
- v0.0.3 六格矩阵缓存未隔离，不是冷缓存或隔离设备基准（`test/vs_echtvar_hg19_hg001_4c16g_v003_release_20261115/README.md`）。
- cold v005：原始 `results/` 未入库；`io.stat` 不可用；GNU RSS 与 cgroup peak 语义不同；OOM 元数据有缺口；一次 0.6s 重叠（同目录 README 的 Recorded limitations）。
- v0.0.6：相对 v0.0.5，NVMe 540/568 墙钟慢 7.6%/11.9%，CPU 同比例增加；NVMe 17 快 12.3%，HDD 540/568 差异 < 0.02%（`test/vs_immkv_hg19_hg001_4c16g_cold_postfix_20261125/README.md`）。
- projection 实验：宽库全基因组 NVMe wall 不稳定（21.9/52.5/130.3s），后段中断；不作 HDD 主张（同目录 README）。
- 10G release 候选未发布；显式 `Columns` 批量回归 +242%；full568 的 I/O regime 发生变化（同目录 README）。
- `samplecolumns` 只覆盖 chr22；估算不含二进制值流与压缩长度表（`test/sample_columnar_hg19_gnomad_genomes_chr22/README.md`）。
- 568 字段的 echtvar 无受限输出，正确性参考来自无限制运行（`test/vs_echtvar_hg19_gnomad_genomes_568_fields/README.md`）。
- HDD 下限由选中列字节 + 输出文本 + keys 决定；调度优化空间有限（`docs/handoff/2026-11-11.md`；540 HDD README）。

**运行环境**

- 机器会随机重启，长任务用 nohup；`github.com:443` 不通（`docs/handoff/2026-11-07.md`）。
- cgroup 测试必须禁 swap，否则限制被换页绕过（同上）。
- 无 DRAM 的低端 NVMe 多进程同时高队列深读会整体塌陷（`README.md` FAQ）。
- cgroup 未委派 io controller，拿不到 `io.stat`（cold v005 README）。

## 11. 文档矛盾与漂移

| 位置 | 现象 | 影响 |
|---|---|---|
| `README.cn.md:58` vs `README.md:71` | 中文版仍用 `ReadManifest(...).Schema()`；英文版已用 `ReadSchema(dir)` | 照中文版写会多一步，且未排序多染色体输入会退化 |
| `docs/v003-format.md:192` | 写“公开的 `EncodeRecord`”；实际导出是 `RecordEncoder`（`record.go:246`） | 按文档找函数会找不到；`EncodeRecord`/`DecodeRecord` 只在 11-11 handoff 出现，未落地 |
| `AGENTS.md:58` vs `test/*4c16g*` | 规则写“所有计时 run 一律 10G”；后期实验用 16G，实验 README 称用户明确覆盖 | 新 session 按 AGENTS 复现会得到不可比数字 |
| `docs/handoff/2026-11-10.md:17` | 引用已删除的 `docs/base-col-format.md` 与 magic `IMKVC001` | 与现行 `IMKVC003` 冲突 |
| `docs/eureka/2026-11-09-hdd-columnar-bundle.md:211-219` | 预测 HDD 540/568 8–12 分钟、17 子集 15s | 实测 23:37/3:30；修正埋在 §12，开头没有撤回标记 |
| 同一 eureka 文档 §4/§10/§15 | 先主张 bundle 不新增格式，后改单一 col，最后把 bundle 列为不做 | 三层决策，必须读到末尾 |
| `README.md:213` 性能段 | 混用 10G/16G、不同版本（568 约 5 分钟 vs 172s） | 不带口径会误读 |
| 多个实验 README | 同为“17 字段”出现 20.66s / 8.72s / 8.17s / 13.20s | 差异来自版本、帧模式、预算、缓存；不可直接比较 |

## 12. 可攻击面

以下不是结论，而是证据最薄或已被承认却缺回归的位置；审阅从这里下钻，每条
应能转成一个失败用例或一个对照实验。标注“文档没有给出”的，表示仓库里找不到
对应证据，不代表一定不存在。

建议路线：先读 §1–§2 与 §5–§6 建立模型，再从本节挑 3–5 条下钻到源码或实验，
只对能给出失败用例或对照方案的条目下结论。

**格式与生命周期**

- `Finish` 先发布 `vals` 再发布 `keys`；两次 rename 之间崩溃会留下同记录数的新旧混用。只有 eureka §7 记录，`buildfile_test.go` 覆盖的是覆盖拒绝与清理，没有该窗口的故障注入。
- `Create` 只写 keys/vals；col/manifest 由 `BuildFromVCFs` 追加，Create 库无法后补列存，只能重建。
- 没有库级 manifest；`ReadSchema` 扫 chr1..M 取第一份并校验一致，部分构建或混放目录靠 shard 级校验兜底。
- manifest 的 `source_info_order` 与 `non_finite_floats` 只作诊断，不参与拒绝；NaN/Inf 静默按缺失处理。

**读路径**

- `preferColumnar` 的跳读段数已计入 `colSec`（`v0.0.6`，`TestPreferColumnarCountsColumnRanges`）；成本常数仍只有一次 chr22 标定（`stream.go:361` 注释），typed 压缩比 4.3 是常数；介质判定依赖 `/sys/block` rotational，USB 桥接与 dm/md 未单列实验。
- `fadvise(DONTNEED)` 读后驱逐与多进程共享页缓存的交互，文档没有给出实验；注释论证的是单进程 cgroup。
- 返回值生命周期靠契约：点查结果到下次查询、`GetBatch` 结果到本次调用；误用无检测。
- io_uring 在途请求已限制在环深内并及时收割（`v0.0.6`，`TestDriveDirectInFlightRespectsRingDepth`）；512 深 + 4MB 读段 + `Threads` 解压池的批内内存峰值仍无推导，实验只报 cgroup peak。
- 打包帧 `pfx` 用 int32、`raw_len` 饱和到 u32；单测覆盖溢出边界，真实库上的压力特征文档没有给出。

**Ingest**

- 多文件要求 `##INFO` 声明顺序完全一致；多来源 VCF 不满足时在构建期报错。
- 未知 contig、`REF="."`、符号化 ALT 的跳过是设计行为；调用方拿不到“跳过了多少”的返回值。
- Number=A 按原始 ALT 下标取元素，被跳过的 ALT 仍占位；“字段缺位”与“源里没有”在输出上相同。
- 64MB 单行上限、`outdir` 必须不存在（含空目录）都是硬失败。

**实验方法**

- 共享 `verifyvalues` 带对手专用例外；它证明两工具输出一致，不证明源 VCF 语义正确（cold v005 明说）。
- 10G/16G、冷/热缓存、v0.0.2/v003/优化后版本的数字常在同一段出现；实验目录没有机器可读的口径字段。
- cold v005 原始证据不入库，`final-summary.*` 是唯一提交记录；复现依赖本机残留。
- v0.0.6 只测 immkv 单边，不重跑 echtvar；v0.0.5 与 v0.0.6 的数字为同协议不同轮次，不是同轮配对。
- v0.0.3 矩阵缓存未隔离；17 字段 echtvar 报 0 个输入块，不能当冷 HDD vs 冷 SSD。

**API 与消费方**

- 公开 API 被 annoplex `kv-custom` 钉住（eureka §8）；`Options` 已改名 `CreateOptions`/`StreamOptions`，旧文字仍出现 `immkv.Options`。
- `StreamOptions.Columns` 只收列号；名称到列号的映射留在调用方，跨重建缓存列号会错。
- `OpenSimpleKV` 只接受单 String 字段；其他 schema 只能换句柄或重建。

## 13. 任务 → 读什么

| 任务 | 最小阅读集 |
|---|---|
| 加字段类型 / 改 record 编码 | `docs/v003-format.md` §3–§4、`schema.go`、`record.go`、`record_test.go` |
| 改键流/微块/值帧 | §5.1–§5.3、`immkv.go:52`、`shard.go:266-420`、`stream.go:398` |
| 改列存格式 | `docs/v003-format.md` §5–§6、`col.go`、`colread.go`、`col_test.go` |
| 改选路成本 | `stream.go:304-377`、`stream.go:313`、`colread.go:328`、对应实验 README 的 path 统计 |
| 改建库过滤/编码 | `docs/ingest-design.md` §3、`ingest.go:677-985`、`ingest_test.go` |
| 加公开 API | `README.md` API/Contracts、`immkv.go` 句柄定义、`shard.go:688`、`stream.go:31`、`simple.go:15` |
| 排查 `ErrUnsorted` | `README.md` FAQ、`stream.go:144-151`、`key.go` |
| 排查坏库/panic | `shard.go:137-247`、`manifest.go:107`、`record.go:477` |
| 跑对照实验 | `AGENTS.md` 实验规则、`test/<exp>/notebook.md`、`common.sh` |
| 复现某一组数字 | 对应 `test/<exp>/README.md` 的 source 列，再开 `results/` 单个文件 |

## 14. 硬约束与陷阱

1. **盘上格式变更必须换魔数**（`IMKVK002` / `IMKVV003` / `IMKVC003`）；旧库
   要么照常打开，要么明确报错。
2. **`CGO_ENABLED=0` 必须可构建**，且纯 Go 与 CGO 的 zstd 帧格式互通。
3. **读路径不引入 mmap 直访**；显式 pread/io_uring 是根基，推翻需要实测。
4. **Linux 专属调用必须有非 Linux 回退**；新增同类代码同理。
5. `GetBatch` 的升序是**完整编码键的字节序**：染色体自然序（1–22, X, Y, M），
   不是字典序；同 POS 的 REF/ALT 也要按字节序。`bcftools sort` 按 `##contig`
   顺序排，先检查头里 MT 是否被放在最前。
6. `ChromID` 接受 `1`–`22`、`chrN`、`CHRN`、`X/Y/M/MT`（大小写不敏感）；
   不接受 `Chr1` 混合写法、`NC_*`、`chrUn_*`、`*_alt`。不认识的 contig 由
   调用方跳过，这是设计行为。
7. 构建目标必须不存在；`Finish` 用临时文件 + fsync + rename，失败清理一次；
   `Abort` 后残留由调用方删除。调用方在构建期间独占目标。
8. `Finish` 先发布 `vals` 再发布 `keys`；两步之间崩溃会留下同记录数的新
   `vals` 与旧 `keys`（eureka §7 记录过这个风险），只能靠调用方保证构建期间
   不打开。
9. `StreamOptions.Columns` 只选字段、不选路径；没有 `base.col/manifest` 的库
   传非空 `Columns` 报错；全选等价于 nil。
10. `GetBatch` 返回值由该次调用持有；点查/前缀/区间返回值只保证到下一次查询。
11. `Threads` 只控制每次 `GetBatch` 的解压并行度；io_uring 路径的在途读深度
    与它无关。实测 4 线程是收益上限；`v0.0.6` 起在途未收割请求上限为环深
    （512），`TestDriveDirectInFlightRespectsRingDepth` 守住该不变量。
12. 无 DRAM 的低端 NVMe 在多进程同时高队列深读时会整体塌陷；多任务同盘时
    要限制节点总在途深度（降每任务线程或做节点级配额）。
13. 实验数字必须能追到源文件；`test/` 全英文；旧记录不覆盖；大工件放仓库外。
14. `docs/handoff/*` 与 `docs/eureka/*` 是本地、gitignored、按日期演进的记录；
    与源码冲突时以源码和 `docs/v003-format.md` 为准。

## 15. 维护

改动以下任一项时更新本索引：新增/删除源文件、公开 API 变化、魔数或格式变化、
新增实验目录、版本 tag 推进。更新时保持“每条目一行 + 锚点”，不要把设计文档
整段搬进来——索引只负责指路。
