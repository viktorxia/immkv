# immkv

独立发布、人类基因组优先的存储引擎，形态是**构建一次、不可变、排序 KV**
（名字来自 immutable sorted KV）。性能路径为 VCF 查询特化；它不是通用 KV
存储。无染色体/坐标的普通 KV 数据只支持一种形态：单库精确 `Get`，
不含前缀、范围、批量查询。它是 annoplex 仓库 `storage_backend: kv-custom`
访问实现的底层引擎，也可独立用于其他人类基因组项目。

## 理念

三个决定定义了这个引擎：

1. **显式读，不信 mmap。** 前身 immdb 用 mmap 直访，在「库大 + 值大 +
   查询位点稀疏」时读放大严重——缺页粒度和内核预读都不受引擎控制。
   immkv 把读取全部换成 pread/io_uring 显式按需读：读多少、何时读、
   合并成几段，全部在引擎手里。
2. **盘上格式为查询模式裁剪。** 键流按 64 条一微块、块索引常驻内存，流式
   批查在块内二分定位，不再逐条解析变长记录；值每条独立成帧（配构建期
   训练的 zstd 字典），点查一条只解压一条、只读那一条的字节。
3. **构建可以慢，查询必须快；CGO 是例外而不是默认。** 默认纯 Go；只在
   zstd 编解码这个热路径上经实测引入 CGO，且保留完整可用的
   纯 Go 回退——`CGO_ENABLED=0` 构建一切照常，只是慢些。

## API

两类基因组句柄对应两种 VCF 查询形态；`OpenSimpleKV` 是另一套单库普通
KV 句柄。所有句柄只读，可同时打开：

```go
// 构建：AddSorted 流式投入（组升序），Finish 原子落盘
w, _ := immkv.Create(cacheDir+"/chr1", immkv.Options{})
w.AddSorted(immkv.PosKey(0, 1000), key, info)
w.Finish()

// 点查句柄：随时来一条答一条（随机键、在线服务）
db, _ := immkv.OpenEager(cacheDir, 512)          // cacheMB：解压缓存上限
v, ok := db.Get(key)                              // 未命中 ok=false
db.GetPrefix(prefix, func(k, v []byte) bool { … })
db.Scan(lo, hi, func(k, v []byte) bool { … })     // lo/hi 可为 nil
db.Close()

// 流式句柄：整份排好序的 VCF 从头查到尾，跑完退出（批量注释）
sdb, _ := immkv.OpenStreaming(cacheDir, immkv.Options{Threads: 9})
vals, _ := sdb.GetBatch(keys)   // keys 升序；返回等长同序，nil = 未命中
sdb.Close()
```

普通 KV（无染色体/坐标）走独立的单库句柄：

```go
w, _ := immkv.Create("genes", immkv.Options{})
w.Add([]byte("BRCA1"), value)                // 任意字节键值，顺序随意
w.Finish()

kv, _ := immkv.OpenSimpleKV("genes", 64)     // cacheMB：解压缓存上限
sv, ok := kv.Get([]byte("BRCA1"))              // 未命中 ok=false
kv.Close()
```

`OpenSimpleKV` 直接打开 `basePath.keys` / `basePath.vals`；键是任意字节串，
按 `bytes.Compare` 排序；句柄只有 `Get` 和 `Close`，没有 `GetPrefix`、
`Scan`、`GetBatch`。`cacheMB` 是解压缓存上限（0 不缓存）；重复键返回排序
后的第一条。构建用 `Create` + `Add`（顺序随意）或 `AddSorted`（组升序）+
`Finish`。

从输入格式直接建库：

```go
// VCF → 分染色体库（outdir/chrN.keys + chrN.vals）
err := immkv.BuildFromVCFs(outdir, vcfs, fields, "IMMKV", immkv.Options{})

// 无表头 TSV（首列键，其余为值）→ 单库
err := immkv.BuildFromTSVs(basePath, tsvs, immkv.Options{})
```

`fields` 为 nil 或空取全部 INFO 字段；非空时只存这些字段。值一律按 `##INFO`
声明顺序重建，`fields` 的顺序只做选择、不再决定值序。`Number=A` 按 ALT 下标取
元素。`prefix` 非空时键写成 `prefix_字段名`。源 INFO 的未知/重复字段、空段、
Flag 带值、非 Flag 缺值都报错；没有输出字段的记录值为 `"."`。VCF 支持单个多
染色体文件或多个每文件一条染色体的输入（`##INFO` 声明顺序须一致），同一
CHROM+POS 内 REF/ALT 顺序任意。完整契约见
[docs/ingest-design.md](docs/ingest-design.md)。

`GetBatch` 的契约：输入必须升序（VCF 按染色体自然序 + 位置排序即天然
满足），返回切片只保证到下一次查询前有效；批次大小就是预取窗口（几千到
几十万均可）；`Threads` 是唯一并发旋钮，定解压并行度；io_uring 路径在途读
深度与它无关，回退路径才兼作 pread 读者数。典型
注释循环是「攒一批键 → 对每个库各 GetBatch 一次 → 用完即弃」：

```go
const batch = 4096
keys := make([][]byte, 0, batch)
for vcf.Scan() {
    keys = append(keys, immkv.VarKey(chrom, pos, ref, alt))
    if len(keys) < batch { continue }
    gnomadVals, _ := gnomadDB.GetBatch(keys)   // vals[i] 与 keys[i] 对齐
    clinvarVals, _ := clinvarDB.GetBatch(keys)
    keys = keys[:0]
}
```

键编码三个函数：`ChromID`/`ChromName`（染色体名 ↔ 首字节编码，首字节序 =
染色体自然序）、`VarKey`（变异键：染色体 + 位置大端 + ref + alt）、
`PosKey`（位置前缀，AddSorted 的组键或前缀查询用）。

### GetBatch 的排序要求与染色体命名（常见疑问）

**排序要求是「编码后键的字节序升序」**，落到变异键上等于三条规则：

1. 染色体按 1–22、X、Y、M 的**自然序**，不是字典序——这是最容易踩的坑：
   `ChromID` 的编码保证了首字节大小 = 自然序先后（1→0 … 22→21、X→22、
   Y→23、M→24），但很多工具产出的 VCF 是字典序（chr1, chr10, chr11, …,
   chr19, chr2, …），字典序下 chr10 排在 chr2 前，键却是 9 对 1，喂进
   来就是降序；
2. 同一染色体内按位置从小到大（VCF 的 POS 列天然如此）；
3. 同一位置多条记录（多等位基因拆出的键）顺序随意——等值键允许。

排序不对**不会静默出错**：引擎逐对检查相邻键，降序立刻返回 `ErrUnsorted`，
不会拿乱序输入算出错误的未命中。拿不准直接喂，报错即说明源头要重排。

**怎么把 VCF 排成自然序**：`bcftools sort` 按头部 `##contig` 声明顺序 +
位置排序，主流项目的头部本身就是自然序（gnomAD、1000G、dbSNP 等），
直接排完即用：

```bash
bcftools view -h in.vcf.gz | grep contig   # 先看头部顺序对不对
bcftools sort -O z -o sorted.vcf.gz in.vcf.gz
```

注意少数 VCF 把 MT 声明在头部**第一位**——`bcftools sort` 会忠实照排，
而 M 在我们的编码里排最后，喂进来 `ErrUnsorted`，把头部顺序修了再排。

**染色体写法归一化**：键首字节是规范化编码，不是原始名字，同一份文件里
`chr1` 和 `1` 混写不影响排序。`ChromID` 认的写法：

| 写法 | 识别 | 编码 |
| --- | --- | --- |
| `1`…`22`、`chr1`…`chr22`、`CHR1` | ✓ | 0…21 |
| `X`/`x`/`chrX`/`CHRX`，Y 同理 | ✓（大小写不敏感） | 22、23 |
| `M`/`MT`/`chrM`/`chrMT`（大小写不限） | ✓ | 24 |
| `Chr1`（首字母大写的混合形态） | ✗ | — |
| `NC_000001.11` 等 accession | ✗ | — |
| `chrUn_*`、`*_alt`、`HLA-*` 等 | ✗ | — |

不认识的 contig 返回 false，**调用方跳过**——alt/decoy 序列没有注释意义，
这是设计口径，不是缺陷。

## 性能

受控对比实验在 test/（对手 echtvar v0.2.4，hg19 / gnomAD v2.1.1 genomes，
10G cgroup，两轮协议，17→540→568 字段剂量系列），要点：

- **17 字段轻载**（echtvar 主场，archive 全在页缓存）：墙钟 echtvar 18s
  对 immkv 窄库 30s；CPU 42s 对 109s（省约 2.6 倍）；66,873,138 次字段
  比较零不一致；
- **568 全字段重载**（16 String + 12 Flag）：echtvar 在 10G 预算下任何
  并发都 OOM（无约束参照峰值 24.1GB RSS/进程）；immkv 全程 10G 内，
  全输出约 5 分钟、null 轮 39–48s、峰值 RSS 3.9GB；2,234,349,552 次字段
  比较零值不一致、字符串逐字节一致；
- **字段剂量对照**（540→568，加字符串字段）：echtvar 从 10G 内 6.5 分钟
  变跑不动；immkv 5:28→5:07、null 43→39s——字符串字段近乎零成本
  （裸字节存储，不逐字段解码）。

## 契约

- **生命周期**：Finish 走 临时文件 + fsync + rename，调用方永远看不到
  半成品；已打开的旧句柄继续读旧数据。同一份数据文件可被任意多进程
  同时打开（无锁），每个句柄内多 goroutine 并发安全。
- **返回切片**：只保证到下一次查询前有效，要保留自己拷。
- **重复键**：允许；`Get` 返回排序后的第一条，`GetPrefix`/`Scan` 看到全部。
- **普通 KV**：`OpenSimpleKV` 打开一对 `basePath.keys`/`basePath.vals`；
  键是任意字节串，按 `bytes.Compare` 排序；句柄只有 `Get`，没有前缀/
  范围/批量查询；重复键返回排序后的第一条。
- **流式输入**：降序返回 `ErrUnsorted`（防静默错果），等值键允许。

### Threads 怎么花、怎么分（常见疑问）

`Threads` 管的是**每个 GetBatch 调用内部的解压并行度**：每次调用起一条
三级流水线（生产者 1 + io_uring 读驱动 1 + 解压池 N），调用结束即拆。
注意读驱动固定单 goroutine——**在途读深度与 Threads 无关**，这是
io_uring 的意义所在；只在 io_uring 不可用的回退路径上 Threads 才兼作
pread 读者数。

线程扩展实测结论：**4 线程即到顶**，更多线程只多烧 CPU 不提速
（解压池争用 + GC 摊薄变差）。

**一个句柄 4 线程还是两个句柄各 2 线程？** 单流形态下吞吐等价（CPU 都
不是瓶颈，4 个解压工拆成几条流水线不改变总活量）；选型看场景：单样本
最快 → 1 句柄 4 线程；多样本并行 → 按编排便利分。**但并发有一
个真实的盘争用坑**：两个进程同时以高队列深度读同一块盘时，DRAM-less 的
低端 NVMe（如本机 Kingston NV3，兼系统盘）会全面塌缩——聚合带宽跌到
远低于单流水平，而 CPU 与内存均无压力。生产
上同节点多任务共读一块 NVMe 时，「带宽近似均分」对这类盘不成立；对策：
限制节点总在途读深度（每任务线程降到 2、或把 io_uring 深度做成节点级
配额），或节点配带 DRAM 缓存的盘。

## 实现

### 文件布局

每个库两个文件（`Create` 的 basePath 如 `…/chr1` 或 `…/genes`）：

```text
chr1.keys: [64B 头][微块 ×N][块索引]
  微块（64 条记录）：[keyEnds u32×n][lens u32×n][baseValOff u64][键区]
                     lens 三标记位：bit31 裸存 / bit30 打包帧 / bit29 同帧
  块索引（常驻内存）：[首键 8B 前缀 u64 大端 ×N][块偏移 u64 ×N]
chr1.vals: [64B 头][字典（每值模式）][值帧 ×N]
```

- 记录的 valOff 不逐条存储：微块记首条绝对偏移，其后按 lens 推导
  （bit29 同帧位不加前条帧长），键流体积明显小于逐条绝对偏移方案。
- 值帧三种：短于 `RawValueBelow`（默认 64B）裸存；打包模式下小于
  `packBelow` 的值按 `packTarget`（默认 64KB）攒组一帧（组头含首记录号
  与长度表，小值靠组内上下文取得压缩率）；`Options.PackTarget < 0`
  进入**每值成帧模式**——所有值独立成帧，构建期用首批值训练 zstd 字典
  （512KB 正文、5000 样本、libzstd L9）嵌入 vals 头。gnomAD 类大值库
  用每值模式，ClinVar 类小值库用默认打包模式。
- 打开时：块索引一次读入常驻（`OpenEager` 与 `OpenSimpleKV` 两个点查
  句柄各配 cacheMB 的 LRU 缓存微块与解压值），流式句柄不缓存（批次即
  预取窗口）。

### 读路径

- **流式（GetBatch）**三级流水线：合并连接生产者（顺扫键流微块窗口，
  256KB/窗，块内二分、零分配视图）→ 读段驱动 → Threads 路解压池。
  读段 = 相邻命中合并（间隔 ≤4KB 并段），每读段按 4KB 对齐一次 O_DIRECT
  读，越过文件尾的 EOF 短读按覆盖所需字节判成功。
- **读驱动**：io_uring（512 深，单 goroutine 事件循环提交/收割，在途读
  深度与线程数解耦）；io_uring 或 O_DIRECT 不可用时回退 Threads 路
  pread 读池。读后对区间 fadvise(DONTNEED)——值只读一次，页缓存留着
  只会顶爆 cgroup 限额。
- **点查（Get）**：内存块索引二分 → 读一个微块（LRU）→ 块内
  二分 → 读值帧解压。缓存命中时零系统调用。

### 编解码（混合档）

| 路径 | 后端 | 说明 |
| --- | --- | --- |
| 构建压缩（每值模式） | CGO libzstd L9 | klauspost 训字典（一次）+ gozstd CDict 逐值压缩 |
| 查询解压（有字典库） | CGO libzstd DDict | 帧格式标准，与纯 Go 互通 |
| 查询解压（无字典打包库） | klauspost 纯 Go | gozstd 解打包帧即使 dst 给足也不快过 klauspost，还多占常驻解码缓冲，不用 |

`CGO_ENABLED=0` 时全部回退 klauspost，读写格式不变。选择逻辑：构建期
看 CGO 可用性；查询期看库头有无字典。

### 平台

linux 全功能（io_uring/O_DIRECT/fadvise 仅 linux，其余平台自动降级到
pread）。文件显式小端、绝对偏移、无地址依赖，跨机器跨架构可直接用。

### 依赖

- `github.com/valyala/gozstd`（CGO，vendored libzstd）：字典库解码主路径；
  `CGO_ENABLED=0` 时回退 `github.com/klauspost/compress`（纯 Go）。
- `golang.org/x/sys`：fadvise/DirectIO 等系统调用的仅 linux 封装。
- `github.com/biogo/hts`：**`test/` 对照工具**（annotate/buildnarrow/buildfull）
  读 bgzip VCF 用，库本体不依赖；引用方在版本库内，`go mod tidy` 会保留它。
