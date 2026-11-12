# immkv v003 统一记录格式

日期：2026-11-11
状态：已实现；全库构建与受控矩阵完成（2026-11-12）。

## 1. 目标

目标：

- 盘上只有一套记录格式：`vals` 的每条值 = 存在位图 + typed 字段值；`col` 的每列 =
  存在位图 + 同一套 typed 值流。
- schema（字段表、类型、Number）落 `manifest`；任何构建入口写同一种记录，任何
  读者按 manifest 解码。
- 数值保真：Integer 精确到 int64，Float 精确到 float64（gnomAD 的 6 位有效数字
  全部精确），String 原字节。源文本写法不保证逐字节，文本输出规范化。
- 现有查询路径的机械结构不变：键流、微块、点查、流式批、列存 bin 与目录都沿用。

## 2. 新术语

- **schema**：有序字段表，每项是 name、Type、Number；顺序即记录里字段的位号。
- **record**：`vals` 里的一条值，由存在位图与字段值组成。
- **bitmap**：存在位图，第 i 位 = schema 第 i 个字段在这条记录里有值。
- **列流**：`col` 里某列在某个 bin 内的值序列。
- **规范化文本**：由 typed 值渲染出的 INFO 文本，不是源文本。

## 3. schema

```go
type FieldType uint8

const (
    TypeInteger FieldType = iota + 1
    TypeFloat
    TypeString
    TypeFlag
)

type Field struct {
    Name   string    // 加过 prefix 的字段名
    Type   FieldType
    Number string    // "1"、"A"、"R"、"G"、"."、"0"
}

type Schema []Field
```

- schema 顺序 = 选中字段的 `##INFO` 声明顺序；`fields` 参数只做选择。
- Flag 判定沿用现有规则：`Number == "0"` 或 `Type == "Flag"`。
- schema hash：按顺序每字段一行 `name\ttype\tnumber\n`，UTF-8，取 SHA-256 前
  8 字节；`manifest` 里存十六进制。
- 位号一经分配不再移动：字段表只追加；同一数据资源的历史版本按此纪律维护。

Number 与记录形态：

| Number | 形态 | 编码 |
|---|---|---|
| `1` | 标量 | 按 Type 的标量编码 |
| `A` | 按 ALT 下标取一个元素后的标量 | 同上 |
| `0` / Flag | 无值 | 只占位图 |
| `R`、`G`、`.` 的数值字段 | 数组 | 见 §4.2 |
| String | 总是原字节，不按逗号拆 | 长度前缀 |

## 4. record（vals）

### 4.1 布局

```text
record := bitmap || payload
bitmap := ceil(N/64) 个 uint64，小端；第 i 位 = schema 第 i 个字段有值
         字内 LSB-first：字 = i>>6，位 = i&63
payload := 按 schema 顺序，对每个有值的字段依次写值
```

- N = schema 字段数；不存字段数、不存长度表，记录总长由值帧给出。
- 缺失字段：位为 0，不占 payload。
- 空记录：位图全 0，payload 为空；渲染时写 `.`。
- 不存 schema id：一个分片只对应一份 manifest，打开时校验。

### 4.2 值编码

| Type | 标量编码 |
|---|---|
| Integer | zigzag uvarint（int64） |
| Float | 8 字节 IEEE 754 小端 |
| String | uvarint(len) + 原字节 |
| Flag | 无 payload |

数组（Number 为 `R`、`G`、`.` 的数值字段）：

```text
array := uvarint(count)
         重复 count 次：uvarint(len) + element
```

- `len = 0` 表示该元素缺失（源里的 `.`）；`count = 0` 表示空列表。
- element 按字段 Type 的标量编码；String 数组的元素是原字节。
- gnomAD 的字段只有 `1`、`A`、`.`（String）、Flag，数组路径不在首轮实测范围，
  但格式与解码器一次定好，不留给下一版。

### 4.3 构建期转换

- `##INFO` 的 Integer / Float 逐元素解析；解析失败返回构建错误，不静默跳过。
- Float 拒绝 `nan` / `inf` 进入记录；构建期遇到时按该字段（或该数组元素）缺失
  处理，并在 manifest 的 `non_finite_floats` 计数；VCF 文法外的其余情况报错。
- Number=A：按原始 ALT 下标取逗号元素；元素不够时该字段在该键上缺位。
- Number=R / G：逗号列表整体进数组；元素缺失写 `len 0`。
- String：原字节，不拆逗号（`vep` 这类字段内部自带分隔符）。
- Flag：只置位，不读值；带 `=` 仍报错。
- 其余记录级校验（未知字段、重复字段、空段、非 Flag 缺值）沿用现有规则。

## 5. col

`base.col` 保持现有骨架：64B 头、bin 目录、列目录、每列 zstd 帧；magic 升为
`IMKVC003`。

### 5.1 头与目录

- 头：magic `IMKVC003`、version 3、header_size 64、build_id、bin_width、bin_count、
  column_count、column_entry_size 32、bin_dir_offset、schema_hash。
- bin 目录 32B：first_record、record_count、reserved、column_dir_offset、data_end。
- 列目录 32B：bitmap_offset、value_offset、bitmap_len、value_len、present_count、
  raw_len。`raw_len` = 该列该 bin 有值记录的 payload 字节合计，用于选路估算。

### 5.2 列流

- 位图：`ceil(record_count/8)` 字节，LSB-first，第 i 位 = bin 内第 i 条记录该字段
  有值；zstd 帧。
- 值流：按 bin 内记录顺序，对有值的记录写 §4.2 的值编码；zstd 帧。
- Flag 列：value_offset / value_len 为 0，只有位图。
- 全缺 bin：两项偏移为 0。
- 每列独立 zstd，压缩级别与 bin_width 沿用现有值，实测后再调。

### 5.3 装配

- 读侧只读选中列；装配出的 record 使用**库 schema 的完整位图**，未选中的字段位
  为 0。调用方知道自己传了 `Options.Columns`，不会把“未选中”误判成“无值”。
- 记录序路径与列存路径产出的 record 字节由同一个 codec 定义，渲染后必须逐字节
  一致。

## 6. manifest 与配对

`base.manifest` 是 UTF-8 JSON：

```json
{
  "format": "immkv-v003",
  "version": 3,
  "base": "chr1",
  "build_id": "hex32",
  "bin_width": 1000000,
  "record_count": 123,
  "keys_bytes": 123,
  "vals_bytes": 123,
  "col_bytes": 123,
  "schema_hash": "hex16",
  "fields": [
    {"name": "AF", "column": 0, "type": "Float", "number": "A"}
  ],
  "source_info_order": {"ok": 123, "bad": 0},
  "non_finite_floats": 0
}
```

- `fields` 顺序 = schema 顺序；`column` 即位号。
- 打开时校验 magic、version、build_id、schema_hash、三个文件字节数与记录数；
  任一不一致大声失败。

## 7. 构建入口

### 7.1 BuildFromVCFs

- schema 由 `##INFO` 头 + `fields` + `prefix` 生成；`fields` 为空取全部 `##INFO`。
- 每条记录解析成 typed 值，按 §4.1 编码后交给 `Create` writer 的 `AddSorted`。
- 记录序 `Finish` 后，顺序扫一遍刚写完的 `keys` / `vals`，解码 record，按 bin 写
  typed 列流；VCF 不读第二遍。构建结构与现在相同，只把“拆文本”换成“解 record”。
- `vcf.raw.header` 保留；`vcf.immkv.header` 保留为渲染产物，不再作为 schema 权威。
- 构建入口要求空目录：`outdir` 里已有任何分片/col/manifest 即报错。

### 7.2 BuildFromTSVs

- schema = 单个 String 字段，`Name = "value"`，`Number = "1"`。
- 每行的值编码成 record；`OpenSimpleKV.Get` 返回该字段的原字节（解包），调用方
  语义不变。
- 无位置语义：产一份单 bin 的 `col`（`bin_width = 0`、`bin_count = 1`），与 vals
  同源；读侧可不用，但盘上产物保持同一套。

### 7.3 Create / AddSorted

- `Options` 新增 `Schema []Field`。
- 空 schema 等价于单 String 字段 `value`；`AddSorted` 收原始值字节，writer 负责
  包成 record。
- 非空 schema：调用方用公开的 `EncodeRecord` 产出 record 字节再 `AddSorted`；
  writer 不校验内容，格式由 codec 契约保证。
- `OpenSimpleKV` 只接受单 String 字段的库；其他 schema 报错，引导改用
  `OpenEager` / `OpenStreaming` + codec。

## 8. 读路径与 codec

- `Get` / `GetBatch` / `Scan` 的签名与语义不变，值仍是 `[]byte`，内容变成 record。
- 公开 codec（已实现，与 annoplex 的存取写法对齐）：

```go
// 校验并返回只读视图；HasValue/Int/Float/Bytes 按字段名取值。
func NewRecord(s Schema, raw []byte) (Record, error)

// 按 schema 顺序遍历有值字段；标量路径零分配，数组字段的 Value 带 Elems。
func (r Record) Values() iter.Seq2[Field, Value]

// 规范化 INFO 文本；用于 test/annotate 与 verifyvalues。
func RenderRecord(s Schema, raw []byte) ([]byte, error)

// 写侧：按字段名设置值，Encode 产出 bitmap + payload。
func NewRecordEncoder(s Schema) *RecordEncoder
func (e *RecordEncoder) Int(name string, v int64) *RecordEncoder
func (e *RecordEncoder) Float(name string, v float64) *RecordEncoder
func (e *RecordEncoder) Bytes(name string, v []byte) *RecordEncoder
func (e *RecordEncoder) Flag(name string) *RecordEncoder
func (e *RecordEncoder) Values(name string, elems []Value) *RecordEncoder
func (e *RecordEncoder) Encode() ([]byte, error)
```

- `Value` 用 Type 区分 int64 / float64 / 原字节 / 数组；缺失由位图表达。
- `OpenEager` / `OpenStreaming` 不解码、不认识结构体；annoplex 的存取实现用
  `NewRecord` 填自己的结构体，`test/annotate` 用 `RenderRecord` 出文本。
- 列存 `GetBatch` 的合并、bin 预取、并行解压结构不变；只是列值流换成 typed
  解码，装配出的 record 使用库 schema 的完整位图。

## 9. 规范化文本

- 字段顺序 = schema 顺序；`name=value`，Flag 写 `name`，全缺写 `.`。
- Integer：十进制；Float：`strconv.FormatFloat(v, 'g', -1, 64)`；String：原字节；
  数组：逗号连接，缺失元素写 `.`。
- 与源文本的差异只在数值写法（`1.14200e+03` → `1142`、`0.00001` → `1e-05`）；
  数值语义一致。验收按数值与字符串比较，不按源字节。

## 10. 选路

- 记录序与列存的成本模型结构不变；`raw_len` 改为 typed payload 字节。
- typed 解码比文本解析便宜，`decodePerThread` 按实测重新标定；列存在 HDD 上会
  比现在更早胜出，NVMe 上可能仍选记录序。
- `IMMKV_DEBUG_PATH` 的成本行保留，字段含义更新为 typed 字节。

## 11. 错误语义

以下一律报错，不静默回退：

- magic / manifest / schema_hash 不一致 / build_id 不一致 / 字节数或记录数
  不一致。
- record 位图长度不足、payload 越界、uvarint 截断、Float NaN/Inf、Integer 溢出、
  数组 count 或 len 越界。
- 列目录偏移越界、zstd 解压失败、present_count 与位图不符、值流条数不符。
- `Options.Columns` 越界/负数/重复；没有 col 的库传非空 Columns。
- `GetBatch` 键不升序 → `ErrUnsorted`。

## 12. 体积与性能

chr22 抽样（20 万值，zstd -9）：

| 列 | typed 压缩 |
|---|---:|
| AF（Float） | 3.67 B/值（float64） |
| AC（Integer） | 1.05 B/值（int64 固定） |
| AC（Integer） | 0.71 B/值（zigzag varint） |

- Integer 用 zigzag varint，Float 用 float64。
- 2026-11-12 chr22 实测（默认构建参数）：`vals` 2.82GB，`col` 1.656GB；
  17 字段 `col` 56.45MB。
- 性能验收：受控 4 核 16G 矩阵，点查 p50/p99、流式批、540/568 全量注释。

## 13. 验收

- 单元测试：schema hash、位图边界（N=1/64/65）、缺字段、Flag、数组、Integer
  溢出、Float 非法、record 截断、列流截断、present_count 不符、未知 magic 拒绝。
- chr22 最小闭环：
  - `chr22.keys` 的键序列、微块分组、`keyEnds`、键区与块索引保持不变；
    值帧 `lens` 与 `baseValOff` 由 record 长度决定，不做逐字节比较。
  - 同一批键分别走记录序与列存，`RenderRecord` 输出逐字节一致。
  - 与源 VCF 逐字段比较：Integer/Float 数值相等，String 原字节相等；Flag 只比
    存在性。
  - 全列与 17 列子集各跑一遍。
- 构建：`BuildFromVCFs` 与 `BuildFromTSVs` 都产 record；TSV 的
  `OpenSimpleKV.Get` 返回原字节。
- 平台：默认与 `CGO_ENABLED=0` 下 `go build ./...` / `go test ./...` 全过；
  zstd 帧两种后端互通。
- 损坏用例全部大声失败。

## 14. 分阶段

1. schema + codec（encode/decode/render）+ manifest 读写 + 单测。**已完成**
2. `BuildFromVCFs` 写 typed vals；点查与流式读仍返回字节；chr22 解码校验。
   **已完成**
3. typed col 写入与读取；选路成本重新标定；chr22 记录序/列存闭环。**已完成**
4. `test/annotate`、`test/verifyvalues`、`test/buildvcf`、`test/buildfull`
   适配 typed record。**已完成**
5. 重建全库与 17/540/568 窄库，重跑 4 核 16G 受控矩阵并逐字段校验。**进行中**

## 15. 明确不做

- 源 INFO 文本逐字节保真。
- 把 annoplex 的 1024 位字段表或具体结构体带进引擎。
- mmap 直访；`pread` / `io_uring` 的读路径不变。
- 数值 Number=R/G 之外的 VCF 语义扩展（基因型、样本列）。
- 共享键流、字段组 bundle、只留列存。
