# immkv

独立存储引擎库：**构建一次、不可变、排序 KV**。
名字取自 immutable sorted KV。它为 annoplex catalog 里 `storage_backend: kv-custom`
的存取实现提供底层引擎，也可在其他项目独立使用。

## 理念

三个决定定义了这个引擎：

1. **显式读，不信 mmap。** 前身 immdb 用 mmap 直访，在「库大 + 值大 +
   查询位点稀疏」时读放大 28 倍（gnomAD 全量注释实测读 263GB，真正需要的
   只有 12GB）——缺页粒度和内核预读都不受引擎控制。immkv 把读取全部换成
   pread/io_uring 显式按需读：读多少、何时读、合并成几段，全部在引擎手里。
   同一台机器同一个负载，92.41 秒变 24.62 秒。
2. **盘上格式为查询模式裁剪。** 键流按 64 条一微块、块索引常驻内存，流式
   批查在块内二分定位，不再逐条解析变长记录；值每条独立成帧（配构建期
   训练的 zstd 字典），点查一条只解压一条、只读那一条的字节。
3. **构建可以慢，查询必须快；CGO 是例外而不是默认。** 默认纯 Go；只在
   zstd 编解码这个热路径上经实测引入 CGO（快 2–6 倍），且保留完整可用的
   纯 Go 回退——`CGO_ENABLED=0` 构建一切照常，只是慢些。

## API

两类句柄对应两种查询形态，同一个库可以同时开（只读无冲突）：

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

`GetBatch` 的契约：输入必须升序（VCF 按染色体自然序 + 位置排序即天然
满足），返回切片只保证到下一次查询前有效；批次大小就是预取窗口（几千到
几十万均可）；`Threads` 是唯一并发旋钮，决定在途读与并行解压数。典型
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

实测与演进史见 [benchmark.md](benchmark.md)，要点（NA12878 全基因组
3,933,369 键注释 gnomAD，3GiB cgroup / 10 线程 / NVMe）：

- 流式 **24.62 秒**（前身 92.41s），每秒 16 万条、每条摊 6.3µs；
- 库体积为原始 bgzip 的 **79.2%**（392GB/495GB）；
- 点查句柄（ClinVar 规模，459 万键）：乱序 p50 291ns，批量有序 702 万条/秒。

## 契约

- **生命周期**：Finish 走 临时文件 + fsync + rename，调用方永远看不到
  半成品；已打开的旧句柄继续读旧数据。同一份数据文件可被任意多进程
  同时打开（无锁），每个句柄内多 goroutine 并发安全。
- **返回切片**：只保证到下一次查询前有效，要保留自己拷。
- **重复键**：允许；`Get` 返回排序后的第一条，`GetPrefix`/`Scan` 看到全部。
- **流式输入**：降序返回 `ErrUnsorted`（防静默错果），等值键允许。

### Threads 怎么花、怎么分（常见疑问）

`Threads` 管的是**每个 GetBatch 调用内部的解压并行度**：每次调用起一条
三级流水线（生产者 1 + io_uring 读驱动 1 + 解压池 N），调用结束即拆。
注意读驱动固定单 goroutine——**在途读深度与 Threads 无关**，这是
io_uring 的意义所在；只在 io_uring 不可用的回退路径上 Threads 才兼作
pread 读者数。

线程扩展实测（NA12878 全量注 gnomAD，3GiB cgroup）：1/2/4/8/10 线程
= 53.8 / 31.1 / **21.2（平台期）** / 21.3 / 24.6 秒。**4 线程即到顶**，
更多线程只多烧 CPU 不提速（解压池争用 + GC 摊薄变差）。

**一个句柄 4 线程还是两个句柄各 2 线程？** 单流形态下吞吐等价（CPU 都
不是瓶颈，4 个解压工拆成几条流水线不改变总活量）；选型看场景：单样本
最快 → 1 句柄 4 线程（21.2s）；多样本并行 → 按编排便利分。**但并发有一
个真实的盘争用坑**：两个进程同时以高队列深度读同一块盘时，DRAM-less 的
低端 NVMe（如本机 Kingston NV3，兼系统盘）会全面塌缩——实测聚合带宽从
单流 0.9GB/s 掉到 0.24→0.09GB/s（每份 31s 恶化到 3m51s→10m12s），CPU
与内存均无压力，iostat 显示 r_await 308ms、队列 666、%util 99%。生产
上同节点多任务共读一块 NVMe 时，「带宽近似均分」对这类盘不成立；对策：
限制节点总在途读深度（每任务线程降到 2、或把 io_uring 深度做成节点级
配额），或节点配带 DRAM 缓存的盘。完整三轮实验与归因更正见
benchmark.md「并发形态对照」。

## 实现

### 文件布局

每条染色体两个文件（`Create` 的 basePath 如 `…/chr1`）：

```text
chr1.keys: [64B 头][微块 ×N][块索引]
  微块（64 条记录）：[keyEnds u32×n][lens u32×n][baseValOff u64][键区]
                     lens 三标记位：bit31 裸存 / bit30 打包帧 / bit29 同帧
  块索引（常驻内存）：[首键 8B 前缀 u64 大端 ×N][块偏移 u64 ×N]
chr1.vals: [64B 头][字典（每值模式）][值帧 ×N]
```

- 记录的 valOff 不逐条存储：微块记首条绝对偏移，其后按 lens 推导
  （bit29 同帧位不加前条帧长），键流体积比逐条绝对偏移省 ~40%。
- 值帧三种：短于 `RawValueBelow`（默认 64B）裸存；打包模式下小于
  `packBelow` 的值按 `packTarget`（默认 64KB）攒组一帧（组头含首记录号
  与长度表，小值靠组内上下文取得压缩率）；`Options.PackTarget < 0`
  进入**每值成帧模式**——所有值独立成帧，构建期用首批值训练 zstd 字典
  （512KB 正文、5000 样本、libzstd L9）嵌入 vals 头。gnomAD 类大值库
  用每值模式，ClinVar 类小值库用默认打包模式。
- 打开时：块索引一次读入常驻（gnomAD 全库 ~65MB；点查句柄另配 CacheMB
  的 LRU 缓存微块与解压值），流式句柄不缓存（批次即预取窗口）。

### 读路径

- **流式（GetBatch）**三级流水线：合并连接生产者（顺扫键流微块窗口，
  256KB/窗，块内二分、零分配视图）→ 读段驱动 → Threads 路解压池。
  读段 = 相邻命中合并（间隔 ≤4KB 并段），每命中按 4KB 对齐一次 O_DIRECT
  读，越过文件尾的 EOF 短读按覆盖所需字节判成功。
- **读驱动**：io_uring（512 深，单 goroutine 事件循环提交/收割，在途读
  深度与线程数解耦）；io_uring 或 O_DIRECT 不可用时回退 Threads 路
  pread 读池。读后对区间 fadvise(DONTNEED)——值只读一次，页缓存留着
  只会顶爆 cgroup 限额。
- **点查（Get）**：内存块索引二分 → 读一个微块（~1.5KB，LRU）→ 块内
  二分 → 读值帧解压。热路径零系统调用。

### 编解码（混合档）

| 路径 | 后端 | 速度 | 说明 |
| --- | --- | --- | --- |
| 构建压缩（每值模式） | CGO libzstd L9 | ~120MB/s/核 | klauspost 训字典（一次）+ gozstd CDict 逐值压缩 |
| 查询解压（有字典库） | CGO libzstd DDict | 1.5–2.5GB/s/核 | 帧格式标准，与纯 Go 互通 |
| 查询解压（无字典打包库） | klauspost 纯 Go | 747MB/s/核 | gozstd 解 64KB 打包帧掉 dst 容量慢路径，实测慢 40%，不用 |

`CGO_ENABLED=0` 时全部回退 klauspost（压缩慢 6 倍、解压慢 2–3 倍），
读写格式不变。选择逻辑：构建期看 CGO 可用性；查询期看库头有无字典。

### 平台

linux 全功能（io_uring/O_DIRECT/fadvise 仅 linux，其余平台自动降级到
pread）。文件显式小端、绝对偏移、无地址依赖，跨机器跨架构可直接用。

### 依赖

- `github.com/valyala/gozstd`（CGO，vendored libzstd）：字典库解码主路径；
  `CGO_ENABLED=0` 时回退 `github.com/klauspost/compress`（纯 Go，慢约 6 倍）。
- `golang.org/x/sys`：fadvise/DirectIO 等系统调用的仅 linux 封装。
- `github.com/biogo/hts`：**仅 gitignore 的 `tmp/` 基准工具**读 bgzip VCF
  用，库本体不依赖。因为引用方不在版本库内，`go mod tidy` 会把它移除，
  跑基准前需重新 `go get github.com/biogo/hts@v1.4.5`。
