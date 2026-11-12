# immkv 建库入口契约：BuildFromVCFs / BuildFromTSVs

日期：2026-11-08
修订：2026-11-12，v003 typed record + schema manifest。记录、列流与 manifest 的
字节契约以 [v003-format.md](v003-format.md) 为准，本文只写两个入口的行为。

## 1. API

```go
func BuildFromVCFs(outdir string, vcfs []string, fields []string, prefix string, opts Options) error
func BuildFromTSVs(basePath string, tsvs []string, opts Options) error
```

- `BuildFromVCFs` 产出 `outdir/chrN.keys`、`chrN.vals`、`chrN.col`、
  `chrN.manifest`。
- `BuildFromTSVs` 产出 `basePath.keys`、`basePath.vals`、`basePath.col`、
  `basePath.manifest`；`OpenSimpleKV` 只读前两个。
- `opts` 只用到构建字段：`RawValueBelow`、`PackTarget`；`Threads`、`CacheMB`
  是读侧参数，建库时忽略。`Schema` 由两个入口自己决定，调用方传入的值被覆盖。

## 2. BuildFromTSVs 契约

- 输入是 `\t` 分隔的文本，无表头。
- 第一列是键；第二个字段到行尾原样是值，列间 `\t` 保留。
- `\n` 结尾去掉；CRLF 的 `\r` 也去掉；空行跳过；没有 `\t` 的行视为键 + 空值。
- 多个文件按参数顺序拼接后，第一列必须整体按字节序升序；内部使用
  `AddSorted(key, key, value)`，内存有界；乱序返回错误。
- 值是原始字节：writer 按单 String 字段 schema（`Name="value"`、`Number="1"`）
  包成 record；`OpenSimpleKV.Get` 解包回原字节，调用方语义不变。
- 输出单库，`OpenSimpleKV(basePath, cacheMB)` 打开；`basePath` 的父目录不存在时
  创建。
- 无位置语义：`col` 是单 bin（`bin_width=0`、`bin_count=1`），与 vals 同源。

## 3. BuildFromVCFs 契约

### 3.1 输入形态

两种都接受：

1. 单个 VCF：每条染色体的记录连续一块、同一染色体不重复；块内 POS 升序。
2. 多个 VCF：每个文件只含一条染色体，块内 POS 升序。

两种形态都允许同一 CHROM+POS 内 REF/ALT 乱序：内部用
`AddSorted(PosKey(cid,pos), VarKey(cid,pos,ref,alt), record)`，同 POS 的记录进同
一组，组结束时按完整键排序。

### 3.2 记录过滤

- 未知 contig（`ChromID` 返回 false）的整条记录跳过。
- `REF == "."` 的整条记录跳过。
- ALT 为空、以 `<`、`*`、`.` 开头、或含 `[`/`]` 的 breakend 不建键；被跳过的
  ALT 仍占原始 ALT 列下标，Number=A 取元素时按原始下标算。
- 只读前 8 列；INFO 是第 8 列，有样本列时忽略样本列。
- 一条记录没有任何可建键的 ALT 时整条跳过。
- 少于 8 列、POS 解析失败、CHROM 未知以外的格式错误返回错误，不静默跳过。

### 3.3 schema 与值构造

- schema 由 `##INFO` 头、`fields`、`prefix` 生成：顺序 = `##INFO` 声明顺序，
  `fields` 只做选择，不决定顺序。Flag 判定为 `Number=0` 或 `Type=Flag`。
- `fields` 为 nil 或空：全部 `##INFO` 字段；非空：只存列出的字段。`fields` 中
  头里不存在的名字、重复名字、空名字直接报错。
- `prefix` 非空时字段名写成 `prefix_字段名`；`prefix` 不得含 `;`、`=`、空白。
- 每个可建 ALT 生成一条 record；字段按 `##INFO` 顺序编码：
  - Integer 逐元素 `strconv.ParseInt`；Float 逐元素 `strconv.ParseFloat`；解析
    失败返回构建错误。Float 的 NaN/Inf 按该字段（或该数组元素）缺失处理，并
    在 manifest 的 `non_finite_floats` 计数。
  - Number=A：按原始 ALT 下标取逗号元素；元素不够时该字段在该键上缺位。
  - Number=R / G / `.` 的数值字段：逗号列表整体进数组，元素缺失写 `len 0`。
  - String：原字节，不拆逗号。
  - Flag：只置位；带 `=` 报错。
- 记录级校验：INFO 中出现 `##INFO` 未声明的字段、同一字段重复出现、空段
  （`;;`、首尾 `;`）、Flag 字段带 `=`、非 Flag 字段缺 `=`，构建返回错误，不静默
  丢字段；`##INFO` 声明本身重复也报错。
- INFO 为 `"."` 或空、或选中字段全部缺失时，record 位图全 0、payload 为空。

### 3.4 输出与头文件

- 每条染色体 `outdir/chr<ChromName>.keys` + `.vals` + `.col` + `.manifest`。
- `outdir` 里已有的 `chrN.keys`/`chrN.vals`/`chrN.col`/`chrN.manifest` 视为冲突：
  构建前检查，存在就报错，要求用空目录。
- 全部染色体 `Finish` 成功后写：
  - `vcf.raw.header`：第一个输入文件的完整头（所有 `##` 行 + `#CHROM` 行），
    逐字节保留。
  - `vcf.immkv.header`：`##fileformat` 行 + 选中字段的 `##INFO` 行（按 `##INFO`
    声明顺序；ID 改为 `prefix_字段名`，Number/Type/Description 不变）+
    无样本的 `#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO` 行。它是渲染参考，
    不是 schema 权威；schema 以 manifest 为准。
- 多输入要求 `##fileformat` 与全部 `##INFO` 定义一致、`##INFO` 声明顺序也一致，
  不一致报错。
- 两个 header 用临时文件 + rename 落盘；失败时删除临时文件。

### 3.5 并行与失败清理

- 单文件模式：顺序读；bgzf 解压用 `bgzf.NewReader(f, 0)`（内部并行）。
- 多文件模式：一个文件一个 goroutine、一个 writer；`bgzf.NewReader(f, 1)`，
  并行度来自文件而不是解压线程。文件级并发上限 `GOMAXPROCS/2`。
- 仅 `.gz`/`.bgz` 后缀走 bgzf 读取；其余按纯文本。
- 失败时当前及未完成 writer 的 `.tmp` 删除；已 `Finish` 的染色体保留；两个
  header 不写。多文件模式首个错误取消其他 worker，返回首个错误。
- 输入只有 `.gz`/`.bgz`/纯文本文件路径，不支持 stdin。

### 3.6 列存伴生文件

- 每条染色体记录序 `Finish` 之后，顺序扫一遍刚写完的 `chrN.keys`/`chrN.vals`，
  写 `chrN.col` + `chrN.manifest`；VCF 不读第二遍。
- 列流直接取 record 里该字段的编码 payload；位图、值流、目录与 manifest 的字节
  契约见 [v003-format.md](v003-format.md) §5–§6。
- 同一构建的全部染色体共用一个 16B 随机 `build_id`，写进 col 头与 manifest；
  打开时用 build_id、schema hash、文件字节数与记录数配对校验，不一致报错。
- 列存写失败时删除该染色体的 col/manifest 半成品；已 `Finish` 的记录序保留。

## 4. 实现结构（ingest.go）

- `BuildFromTSVs` / `addTSVFile`；writer 的空 schema 包装路径负责单 String 字段。
- `BuildFromVCFs` / `prepareVCFInputs` / `buildVCFShardsSingle` /
  `buildVCFShardsParallel` / `buildOneVCFShard`。
- `vcfShardBuilder`：按记录推进、管理当前 writer；`parseInfo` 保存选中字段的
  原始文本，`buildRecord` 用 `RecordEncoder` 生成 typed record；`fixed >= 0`
  表示多文件模式只允许一条染色体，`-1` 表示单文件模式的 CHROM 块切换。
- `schemaFromSelectedFields`：`valueField` → `Schema`；Flag 落 `TypeFlag`。
- `vcfHeader` / `infoDef`：头解析、`##INFO` 属性解析、跨输入一致性检查（含声明
  顺序）、字段声明顺序。
- `openVCF` / `vcfReader` / `readVCFHeader` / `splitVCFLine` / `skipVCFAlt`。
- `buildImmkvHeader` / `prefixInfoLine` / `writeFileAtomic`。
- `buildState`：多文件模式的取消标志与首个错误。

## 5. 测试（ingest_test.go / col_test.go / record_test.go）

- TSV：基本键值、列间 tab 保留、空值、空行、重复键、多文件拼接、乱序报错、
  `OpenSimpleKV` 回查、单 bin col + manifest、失败后不留正式文件。
- VCF 单文件：多染色体块、同 POS REF/ALT 乱序、Number=A 按 ALT 下标取元素
  （含符号化 ALT 占位）、Flag、`fields`/`prefix`、源 INFO 反序仍按 `##INFO`
  序编码、`fields` 顺序只做选择、无字段位图全 0、样本列忽略、`REF="."`/
  符号化 ALT 跳过、未知 contig 跳过、两个 header 内容。
- 值类型：Integer/Float/String/Flag、数组与缺失元素、Float NaN/Inf 拒绝、
  Integer 溢出拒绝、缺字段、N=1/64/65 位图边界。
- 记录级错误：未知字段、重复字段、空段、空字段名、Flag 带值、非 Flag 缺值。
- 多文件：按文件并行、跨文件重复染色体报错、文件内 CHROM 变化报错、`##INFO`
  声明顺序不一致报错。
- 错误路径：POS 倒序、染色体块不连续、未知字段、重复字段、非法 prefix、
  outdir 已有 chrN 文件、空输入。
- 列存：present_count 不符、值流截断、未知 magic、col/manifest 缺一、
  Columns 越界/重复。
- bgzf：用 `bgzf.NewWriter` 写一个 `.vcf.gz` 验证读取路径。

## 6. 验收

- `go build ./...`、`go test ./...` 在默认与 `CGO_ENABLED=0` 下全过。
- schema 顺序强制 `##INFO` 序；`fields` 只做选择；未知/重复/空段/Flag/非 Flag
  错误；无字段位图全 0。
- chr22 全字段与 17 字段：记录序/列存输出逐字节一致，源 VCF 逐字段 0 差异。
- README 中英同步两个入口与 typed record 的 API。

## 7. 不做

- 不做外部排序：VCF 与 TSV 都必须预排序。
- 不支持 stdin、字段裁剪之外的 INFO 变换（前缀除外）、BED/GTF 等格式。
