# base.col 格式契约（草稿）

日期：2026-11-10
状态：草稿，待用户确认；未实现。`base.col` 是记录序库（`base.keys` / `base.vals`）
的顺序扫描伴生文件；点查、前缀、区间扫描仍走记录序。本文件定义盘上字节与读侧
行为；eureka 记录的是取舍，不是这份契约。改本格式必须升 magic。

## 1. 产物与命名

每个 `base`（如 `chr1`）四件：

| 文件 | 格式 | 作用 |
|---|---|---|
| `base.keys` / `base.vals` | `IMKVK002` / `IMKVV001`，不动 | 记录序；点查与回退路径 |
| `base.col` | 本文，magic `IMKVC001` | 列存扫描件 |
| `base.manifest` | JSON，`format: immkv-col-manifest` | 字段映射、构建代号、配对 |

`outdir` 级另有两份头（每个 outdir 一份，不分染色体）：

| 文件 | 内容 | 作用 |
|---|---|---|
| `vcf.raw.header` | 第一个输入的完整头（所有 `##` 行 + `#CHROM` 行），逐字节 | 源头；provenance / annotate |
| `vcf.immkv.header` | `##fileformat` + 选中字段的 `##INFO`（按 `##INFO` 序，ID 加 prefix）+ 无样本 `#CHROM` | schema；manifest 的 `fields` 必须与它一致 |

这两份头是整个 outdir 构建成功的标记；`base.col` / manifest 是每染色体产物。

- 旧库没有 `base.col` / manifest，照常打开；`OpenEager` / `GetPrefix` / `Scan`
  不受影响。
- `base.col` 只服务 `OpenStreaming` 的批量顺序查；不参与点查。
- `base.col` 按染色体一个文件，与 `base.keys` / `base.vals` 同目录同前缀。
- 不复制键流：列存只存值；记录号、bin、键匹配都走 `base.keys`。

## 2. 构建代号、配对与原子性

- 构建代号 16 字节随机值，写进 `base.col` 头与 `base.manifest`；manifest 里另记
  `base.keys` / `base.vals` / `base.col` 的字节数与 `record_count`。
- 打开 `base.col` 时校验：magic、version、manifest 存在、构建代号一致、schema
  hash 一致、三个文件字节数与记录数一致。任一不一致大声失败，不回退。
- 不改 `base.keys` / `base.vals` 的头与 magic；配对靠 manifest 的字节数/记录数，
  所以旧库与旧读者不受影响。
- 每染色体顺序：记录序 `Finish`（vals 先、keys 后）→ `base.col` temp+rename →
  manifest temp+rename。manifest 缺失时 `base.col` 视为不完整，列存路径报错。
- 全部染色体完成后写 outdir 级的 `vcf.raw.header` / `vcf.immkv.header`
  （temp+rename），它们是整个 outdir 构建成功的最后标记；缺头时库不完整，列存
  路径报错。旧头在构建开始时删除。
- 构建失败：删除未完成的 `base.col` / manifest temp；已 `Finish` 的染色体记录序
  保留，与现有 `BuildFromVCFs` 失败语义一致。

## 3. `base.col` 布局

所有整数小端；所有偏移是文件内绝对偏移。

### 3.1 文件头（64B）

| 偏移 | 长度 | 字段 | 说明 |
|---:|---:|---|---|
| 0 | 8 | magic | `IMKVC001` |
| 8 | 4 | version | 1 |
| 12 | 4 | header_size | 64 |
| 16 | 16 | build_id | 构建代号 |
| 32 | 4 | bin_width | bp，v1 固定 1000000 |
| 36 | 4 | bin_count | |
| 40 | 4 | column_count | 等于 manifest 字段数 |
| 44 | 4 | column_entry_size | 32 |
| 48 | 8 | bin_dir_offset | bin 目录起点 |
| 56 | 8 | schema_hash | schema 字符串 SHA-256 前 8B |

schema 字符串：按 schema 顺序，每字段一行 `name\ttype\tnumber\tencoding\n`，
UTF-8。manifest 里存同一个 hash 的十六进制。

### 3.2 bin 目录

`bin_count` 个 32B 条目：

| 偏移 | 长度 | 字段 |
|---:|---:|---|
| 0 | 8 | first_record（全局记录号，0 起） |
| 8 | 4 | record_count |
| 12 | 4 | reserved，0 |
| 16 | 8 | column_dir_offset |
| 24 | 8 | data_end |

- bin 覆盖位置区间；空 bin 的 `record_count = 0`，列目录与数据区可省略。
- `first_record` 是 `base.keys` 里的记录号；bin 内记录顺序 = 键序。
- `data_end` 是该 bin 数据区终点，读侧用它做边界检查。

### 3.3 列目录

每个 bin 的列目录是 `column_count` 个 32B 条目，顺序 = schema 顺序：

| 偏移 | 长度 | 字段 |
|---:|---:|---|
| 0 | 8 | bitmap_offset |
| 8 | 8 | length_table_offset |
| 16 | 8 | value_stream_offset |
| 24 | 4 | length_table_len（压缩后） |
| 28 | 4 | value_stream_len（压缩后） |

- 存在位图长度 = `ceil(record_count/8)`，不存长度；bit i（LSB-first）对应 bin 内
  第 i 条记录，1 = 字段存在。
- `length_table` 是 zstd 帧，解压后为 uvarint 序列，每个存在值一个长度。
- `value_stream` 是 zstd 帧，解压后为按记录顺序拼接的值（`raw` 或 `dec` 编码）。
- Flag 列：`length_table_offset` / `value_stream_offset` 为 0，长度字段为 0；
  只有位图。
- 字段全部缺失的 bin：该列三项偏移为 0；读侧按位图全 0 处理。

### 3.4 值编码

manifest 的 `fields[].encoding` 三选一：

- `raw`：字段值原文字节；长度表给每个值的字节数。
- `bitmap`：仅 Flag；不存值；输出时写字段名（带 prefix）。
- `dec`：Integer/Float 的无损十进制编码；见 §3.5。

构建时按列决定：`dec` 只在整列所有值（含逗号列表的每个元素）都能无损往返时
启用，否则整列 `raw`。String 固定 `raw`，Flag 固定 `bitmap`。

### 3.5 `dec` 编码

支持的十进制文法：

```text
value   := element ("," element)*
element := [+-]? ( digits ( "." digits? )? | "." digits ) ( [eE] [+-]? digits )?
digits  := [0-9]+
```

至少一个尾数数字；指数可带 `+` / `-`；保留 `e` / `E` 的大小写、指数前导零、
尾数前导/尾随零、`.5`、`5.`、`+`、`-0`。空元素、`.`、`nan`、`inf`、空白、
多小数点、多指数都失败，整列回退 `raw`。

每个 element 的字节：

| 字段 | 编码 |
|---|---|
| flags | u8：bit0-1 尾数符号（0 无、1 `+`、2 `-`）；bit2-3 指数符号（同）；bit4 hasPoint；bit5 hasExp；bit6 指数标记 `E`=1 / `e`=0；bit7 保留 0 |
| mantissa_digits | uvarint，尾数数字个数（含前后的零） |
| integer_digits | uvarint，小数点前数字个数（0 = `.5`；等于 mantissa_digits = `5.`） |
| exp_digits | uvarint，仅 hasExp；指数数字个数 |
| mantissa | BCD，`ceil(mantissa_digits/2)` 字节，高 nibble 在前，末 nibble 补 0 |
| exponent | BCD，`ceil(exp_digits/2)` 字节，同上 |

整个字段值先写 `element_count`（uvarint，1 起），再逐 element 写上述字节；逗号
由读侧按 element_count 恢复。`dec` 值长度进长度表，与 `raw` 相同。

### 3.6 manifest

`base.manifest` 是 UTF-8 JSON：

```json
{
  "format": "immkv-col-manifest",
  "version": 1,
  "base": "chr1",
  "build_id": "hex32",
  "bin_width": 1000000,
  "record_count": 123,
  "keys_bytes": 123,
  "vals_bytes": 123,
  "col_bytes": 123,
  "schema_hash": "hex16",
  "source_info_order": {
    "ok": 123, "bad": 0, "inversions": 0,
    "unknown": 0, "duplicate": 0, "empty_segments": 0
  },
  "fields": [
    {"name": "AF", "column": 0, "type": "Float", "number": "A", "encoding": "dec"}
  ]
}
```

- `fields` 顺序 = schema 顺序 = `vcf.immkv.header` 的 `##INFO` 顺序；`column` 即
  列号。构建/打开时校验 manifest 字段与 `vcf.immkv.header` 的 ID、Type、Number、
  顺序一致。
- `source_info_order` 是构建时源 INFO 顺序的统计，只作诊断；入库已强制 header 序。
- `keys_bytes` / `vals_bytes` / `col_bytes` 与 `record_count` 用于 §2 的配对校验。

## 4. 构建路径

- `BuildFromVCFs` 一次解析同时写记录序、`base.col`、manifest；记录序与列存都按
  `##INFO` 声明序重建值（`fields` 只做选择），所以两条路径的字段顺序天然一致。
- `outdir` 里已有的 `chrN.keys` / `chrN.vals` / `chrN.col` / `chrN.manifest` 都算
  冲突；要求空目录，避免新旧产物混在一起。
- 列存写入按 bin 攒：每 bin 每字段的存在位图、长度表、值流；bin 结束时压缩并写
  列目录；bin 目录最后回填。
- bin_width v1 固定 1,000,000 bp；manifest/头记录；读侧按头。
- `base.col` 与 manifest 只由 `BuildFromVCFs` 产出；`BuildFromTSVs` 和
  `OpenSimpleKV` 没有列存。`test/buildfull` 仍是记录序全负载工具；需要列存全
  负载时用 `BuildFromVCFs` 建。
- 始终产列存：构建不按行长预判查询模式；小值库也产；读侧决定是否使用。

## 5. 读路径

- `OpenStreaming` 新增 `Options.Columns []int`：
  - `nil` 或空：全部列，走列存（eureka §10.3 的“不传 = 全部列”）。
  - 非空：走列存，元素是 manifest 里的列号；越界、负数报错。
  - 没有 `base.col` / manifest：报错；旧库可删除，不再回退。
- 列存 `GetBatch`：
  1. 校验升序（`ErrUnsorted`）。
  2. 对 `base.keys` 做现有合并连接，得到每个键的全局记录号与命中/未命中。
  3. 按 bin 聚合命中；读该 bin 的列目录与选中列的位图、长度表、值流。
  4. 按记录号在 bin 内的下标查位图；1 才取值；按长度表切值；`raw` 直接用，
     `dec` 解码，`bitmap` 只写字段名。
  5. 按 schema 顺序拼值：`name=value` / `name` / `.`；返回与 keys 等长同序。
- 点查、`GetPrefix`、`Scan`、`OpenSimpleKV` 全部保持记录序；列存不参与。
- `Threads` 仍只控解压并行度；列存读复用现有 io_uring / pread 路径与 fadvise
  策略。v1 不按介质自动选路；`OpenStreaming` 默认列存全列，`Options.Columns`
  只选子集。
- 未选中的列不读；缺失值不占值空间；Flag 只读位图。

## 6. 错误语义

以下情况一律返回错误，不静默回退：

- `base.col` magic / version / header_size / bin_width / bin_count / column_count
  不合法；schema_hash 不匹配。
- manifest 缺失、JSON 非法、build_id / 字节数 / 记录数 / schema 不一致。
- bin / 列目录偏移越界；位图、长度表、值流解压失败；长度表和超过值流长度；
  记录号不在 bin 范围；`dec` 解码失败。
- `Options.Columns` 越界、负数；库没有 `base.col` 或 manifest。
- `GetBatch` 键不升序 → `ErrUnsorted`（与现有语义一致）。

## 7. 平台与编码

- 所有整数小端；偏移绝对；无地址依赖；跨机器/架构可直接读。
- zstd 帧标准格式；纯 Go 与 CGO 后端互通；`CGO_ENABLED=0` 必须能写能读。
- 读侧继续显式按需 pread / io_uring，不用 mmap 直访；Linux-only 系统调用有非
  Linux 回退。
- `base.col` 的帧校验和默认关闭（文件完整性由构建/配对保证），与现有值帧一致。

## 8. 验收

- chr22 最小闭环：`BuildFromVCFs` 产 `chr22.col` / `chr22.manifest`；同一批键
  分别走记录序与列存（全列 + 17 列子集）；`verifyvalues` 字段级全量一致；强制
  排序后两条路径逐字节一致；键数、命中/未命中一致。
- 现有 `go build ./...` / `go test ./...` 在默认与 `CGO_ENABLED=0` 下全过；没有
  `base.col` 的旧库行为不变；点查/前缀/区间测试不回退。
- `vcf.raw.header` 与输入首头逐字节一致；`vcf.immkv.header` 按 schema 序；
  manifest 的 `fields` 与 `vcf.immkv.header` 的 ID/Type/Number/顺序一致；失败
  时两个头都不留。
- HDD 540/568 重跑：列存路径的 `verifyvalues`、墙钟、峰值内存、体积；与记录序
  和 echtvar 对照。
- 损坏用例：错 magic、错 manifest、错 build_id、错 schema、截断的 bin/列、坏
  zstd、`dec` 坏值，全部大声失败。
- 格式变更：magic 从 `IMKVC001` 升位；旧文件要么照旧打开，要么大声失败。

## 9. 分阶段与待定

1. 本契约定稿。
2. `base.col` 写入 + manifest + `BuildFromVCFs`；`raw` 列；chr22 最小闭环；
   `Options.Columns` 读路径；单测。
3. `dec` / Float 指数编码、长度表压缩、体积与 HDD 实验。
4. 自动选路 / 介质判定 / 生产化。

待用户拍板：

1. `dec` 是否进 v1：草案把格式定义完整、chr22 最小闭环先跑 `raw`，`dec` 在
   第 3 阶段启用；还是 v1 就带上 `dec`？
2. bin_width：v1 固定 1Mb（采样 chr22 的压缩值 43.1M/bin，高于 20–30M 目标）；
   先按 1Mb 实测再调，还是 v1 就改小？
3. 构建代号：16B 随机值（草案）还是内容哈希？
4. 长度表：zstd 压缩的 uvarint（草案）还是 u32？

## 10. 不做

- 点查 / 前缀 / 区间扫描走列存。
- 共享键流、字段组 bundle、`B` 紧凑记录序值、`C` 只留列存。
- v1 的自动介质判定与按命中密度自动选路。
- `BuildFromTSVs` / `OpenSimpleKV` 的列存。
- 把列存值编码预设成某个具体字段集或查询模式。
