# immkv 建库入口设计：BuildFromVCFs / BuildFromTSVs

日期：2026-11-08
修订：2026-11-10，值序改为强制 `##INFO` 声明序（`fields` 只做选择）；见 §4.3。
状态：已与用户确认，按本文实现。

## 1. 背景与目标

库目前只有 `immkv.Create` 一个构建入口，VCF/TSV 的解析、键编码、按位置组调用
`AddSorted`、跳过符号化 ALT、并行建库等“最佳构建实现”散落在 `test/buildfull`、
`test/buildnarrow` 和将来的 immkv_anno 里。本次在核心 `immkv` 包新增
`ingest.go`，提供两个从输入格式直接建库的入口：

- `BuildFromVCFs`：VCF → 分染色体库（形态 A）。
- `BuildFromTSVs`：无表头 TSV → 单库（普通 KV）。

## 2. API

```go
func BuildFromVCFs(outdir string, vcfs []string, fields []string, prefix string, opts Options) error
func BuildFromTSVs(basePath string, tsvs []string, opts Options) error
```

- `BuildFromVCFs` 的 `outdir` 是目录，产出
  `outdir/chr1.keys`+`chr1.vals` … `chr22`、`chrX`、`chrY`、`chrM`。
- `BuildFromTSVs` 的 `basePath` 是单库的文件前缀，产出
  `basePath.keys` + `basePath.vals`，与 `Create`、`OpenSimpleKV` 的 `basePath` 一致。
- `opts` 只用到构建相关字段：`RawValueBelow`、`PackTarget`；`Threads`、`CacheMB`
  是读侧参数，建库时忽略。

## 3. BuildFromTSVs 契约

- 输入是 `\t` 分隔的文本，无表头。
- 第一列是键；第二个字段到行尾原样是值，列间 `\t` 保留。
- `\n` 结尾去掉；CRLF 的 `\r` 也去掉；空行跳过；没有 `\t` 的行视为键 + 空值。
- 多个文件按参数顺序拼接后，第一列必须整体按字节序升序；内部使用
  `AddSorted(key, key, value)`，内存有界；乱序返回错误。
- 重复键允许（同组内排序），`OpenSimpleKV` 的 `Get` 返回排序后的第一条。
- 输出单库，`OpenSimpleKV(basePath, cacheMB)` 打开；`basePath` 的父目录不存在时创建。

## 4. BuildFromVCFs 契约

### 4.1 输入形态

两种都接受：

1. 单个 VCF：染色体块按 chr1–22/X/Y/M 自然序（实现上只要求每条染色体的记录
   连续一块、同一染色体不重复；自然序是特例，字典序块序也可），块内 POS 升序。
2. 多个 VCF：每个文件只含一条染色体，块内 POS 升序。

两种形态都允许同一 CHROM+POS 内 REF/ALT 乱序：内部用
`AddSorted(PosKey(cid,pos), VarKey(cid,pos,ref,alt), value)`，同 POS 的记录进同一组，
组结束时按完整键排序。

### 4.2 记录过滤

沿用 `test/buildfull` 口径：

- 未知 contig（`ChromID` 返回 false）的整条记录跳过。
- `REF == "."` 的整条记录跳过。
- ALT 为空、以 `<`、`*`、`.` 开头、或含 `[`/`]` 的 breakend 不建键；被跳过的
  ALT 仍占原始 ALT 列下标，Number=A 取元素时按原始下标算。
- 只读前 8 列；INFO 是第 8 列，有样本列时忽略样本列。
- 一条记录没有任何可建键的 ALT 时整条跳过（不创建空染色体库）。
- 少于 8 列、POS 解析失败、CHROM 未知以外的格式错误返回错误，不静默跳过。

### 4.3 值构造

所有模式的值都按 `##INFO` 声明顺序重建；`fields` 的顺序只做选择，不再决定值序。
源记录的 INFO 字段顺序可以任意。

- `fields` 为 nil 或空：全部 INFO 字段。
  - `prefix == ""`：值按 `##INFO` 序重建为 `键=值;…`；Flag 不带 `=`；INFO 为 `"."`
    或空时值为 `"."`。
  - `prefix != ""`：每个 INFO 键改写为 `prefix_键`，按 `##INFO` 序排列；Flag 保持
    不带 `=`；INFO 为 `"."` 或空时值为 `"."`。
- `fields` 非空：只存列出的字段，按 `##INFO` 声明顺序拼 `键=值`；记录里缺的字段
  省略；Flag 字段只写 `prefix_键`（不带 `=`）；`fields` 中头里不存在的名字、重复
  名字、空名字直接报错；没有输出字段的值写成 `"."`。
- 记录级校验：INFO 中出现 `##INFO` 未声明的字段、同一字段重复出现、空段
  （`;;`、首尾 `;`）、Flag 字段带 `=`、非 Flag 字段缺 `=`，构建返回错误，不静默
  丢字段；`##INFO` 声明本身重复也报错。
- `Number=A`：从 `##INFO` 头读 Number，按原始 ALT 下标取逗号分隔的元素；元素不够
  时该字段在该键上省略。其余 Number（`1`、`.`、`R`、`G`）整段保留。
- Flag 判定：`Number=0` 或 `Type=Flag`。
- `prefix` 非空时键写成 `prefix_字段名`；`prefix` 不得含 `;`、`=`、空白，否则报错。

### 4.4 输出与头文件

- 每条染色体 `outdir/chr<ChromName>.keys` + `.vals`，由 `Create` + `Finish`
  原子替换。
- `outdir` 里已有的 `chrN.keys`/`chrN.vals` 视为冲突：构建前检查，存在就报错，
  要求用空目录，避免新旧染色体混在一起。旧的两个 header 文件会被删除。
- 全部染色体 `Finish` 成功后写：
  - `vcf.raw.header`：第一个输入文件的完整头（所有 `##` 行 + `#CHROM` 行），
    逐字节保留。
  - `vcf.immkv.header`：`##fileformat` 行 + 选中字段的 `##INFO` 行（按 `##INFO`
    声明顺序排列；ID 改为 `prefix_字段名`，Number/Type/Description 不变）+
    无样本的 `#CHROM\tPOS\tID\tREF\tALT\tQUAL\tFILTER\tINFO` 行。
- 多输入要求 `##fileformat` 与全部 `##INFO` 定义一致、`##INFO` 声明顺序也一致，
  不一致报错。
- 两个 header 用临时文件 + rename 落盘；失败时删除临时文件。

### 4.5 并行与失败清理

- 单文件模式：顺序读；bgzf 解压用 `bgzf.NewReader(f, 0)`（内部并行）。
- 多文件模式：一个文件一个 goroutine、一个 writer；`bgzf.NewReader(f, 1)`，
  并行度来自文件而不是解压线程。
- 仅 `.gz`/`.bgz` 后缀走 bgzf 读取；其余按纯文本。
- 失败时当前及未完成 writer 的 `.tmp` 删除；已 `Finish` 的染色体保留；两个
  header 不写。多文件模式首个错误取消其他 worker，返回首个错误。
- 输入只有 `.gz`/`.bgz`/纯文本文件路径，不支持 stdin。

## 5. 实现结构（ingest.go）

- `BuildFromTSVs` / `addTSVFile`
- `BuildFromVCFs` / `prepareVCFInputs` / `buildVCFShardsSingle` /
  `buildVCFShardsParallel` / `buildOneVCFShard`
- `vcfShardBuilder`：按记录推进、管理当前 writer、组键与值构造；`all` 模式由
  `sortInfo` 按 `##INFO` 序重建值，`fields` 模式由 `parseInfo` 处理选中字段并做
  记录级校验；`fixed >= 0` 表示多文件模式只允许一条染色体，`-1` 表示单文件模式的
  CHROM 块切换。
- `vcfHeader` / `infoDef`：头解析、`##INFO` 属性解析、跨输入一致性检查（含声明
  顺序）、字段声明顺序。
- `valueField`：选中字段的 id、加前缀后的 id、Number、Flag、原始 `##INFO` 行、
  声明顺序。
- `openVCF` / `vcfReader` / `readVCFHeader` / `splitVCFLine` / `skipVCFAlt`。
- `buildImmkvHeader` / `prefixInfoLine` / `writeFileAtomic`；`buildImmkvHeader`
  按 `valueField` 的声明顺序输出。
- `buildState`：多文件模式的取消标志与首个错误。

## 6. 测试计划（ingest_test.go）

- TSV：基本键值、列间 tab 保留、空值、空行、重复键、多文件拼接、乱序报错、
  `OpenSimpleKV` 回查、失败后不留正式文件。
- VCF 单文件：多染色体块、同 POS REF/ALT 乱序、Number=A 按 ALT 下标取元素
  （含符号化 ALT 占位）、Flag、`fields`/`prefix`、整条 INFO 按 `##INFO` 序重建
  （源 INFO 反序也一致）、`fields` 顺序只做选择、无字段写 `.`、样本列忽略、
  `REF="."`/符号化 ALT 跳过、未知 contig 跳过、`vcf.raw.header`/
  `vcf.immkv.header` 内容（后者按声明序）。
- 记录级错误：未知字段、重复字段、空段、空字段名、Flag 带值、非 Flag 缺值。
- 多文件：按文件并行、跨文件重复染色体报错、文件内 CHROM 变化报错、`##INFO`
  声明顺序不一致报错。
- 错误路径：POS 倒序、染色体块不连续、未知字段、重复字段、非法 prefix、
  outdir 已有 chrN 文件、空输入。
- bgzf：用 `bgzf.NewWriter` 写一个 `.vcf.gz` 验证读取路径。

## 7. 验收

- `go build ./...`、`go test ./...` 在默认与 `CGO_ENABLED=0` 下全过。
- `gofmt -l` 对本次改动的文件无输出；`test/verifyvalues/main.go` 的既有格式问题
  不在本次范围。
- 值序强制 `##INFO` 序；`fields` 只做选择；未知/重复/空段/Flag/非 Flag 错误；
  无字段写 `.`。
- 既有测试不回退；核心读路径、Writer、盘上格式与 magic 不变。
- README 中英同步补两个入口的 API 与契约。

## 8. 不做

- 不做外部排序：VCF 与 TSV 都必须预排序。
- 不修 `docs/handoff/2026-11-08.md` 里记录的两个 bug。
- 不重构 `test/` 下的工具；`buildfull` 随值序规则同步，其余保持原样作为实验
  参考实现。
- 不支持 stdin、字段裁剪之外的 INFO 变换（前缀除外）、BED/GTF 等格式。
