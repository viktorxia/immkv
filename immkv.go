// Package immkv（immutable sorted KV）是独立发布的存储引擎库：
// 构建一次、不可变、排序 KV。每条染色体的数据是两个文件——键流（有序、
// 带稀疏锚点索引）与值文件（每条值独立成帧）——读侧 pread 显式按需读取，
// 不使用 mmap 缺页路径。
//
// 形态：通用排序键值。键是任意字节串，全序按 bytes.Compare；值是任意字节串。
// 适合「构建后只读、多进程共享同一份数据文件、点查/前缀/区间/批量查询」的场景。
//
// 文件布局（base 是公共前缀，如 …/chr1 → chr1.keys + chr1.vals）：
//
//	chr1.keys: [64B 头][微块 ×N][块索引]
//	  微块（64 条记录）：[keyEnds u32×n][lens u32×n][baseValOff u64][键区]
//	                     lens 三标记位：bit31 裸存 / bit30 打包帧 / bit29 同帧
//	  块索引（常驻内存）：[首键 8B 前缀 u64 大端 ×N][块偏移 u64 ×N]
//	chr1.vals: [64B 头][字典（每值模式）][值帧 ×N]
//
// 值帧三种：短于 RawValueBelow 的值裸存；打包模式下小于 packBelow 的值按
// packTarget 攒组打包成一个 zstd 帧（组头含首记录号与长度表）；其余每条
// 独立成帧。PackTarget<0 进入每值成帧模式（全部值独立成帧 + 构建期训练
// 字典嵌 vals 头）。
//
// 记录按 key 升序；偏移一律是文件内绝对偏移。微块内 valOff 由 baseValOff
// 与 lens 推导（bit29 同帧位不加前条帧长）。
//
// 生命周期契约：
//   - Create 的目标文件必须不存在；已有库不会被覆盖，重建前先删除旧文件。
//   - Finish 走临时文件 + fsync + rename；失败时返回错误并清掉本次的 .tmp。
//   - 同一份数据文件可被任意多进程同时打开（无锁）。
//   - 查询返回的切片只保证到下一次查询前有效，要保留就拷贝。
//
// 平台：纯 Go 可构建（CGO_ENABLED=0）；CGO 可用时每值压缩与解压走 libzstd
// 加速（帧格式与纯 Go 后端标准互通），文件格式本身零地址与内存布局依赖，
// 跨机器、跨架构可直接使用。
//
// 本包为 annoplex `kv-custom` 后端的底层引擎，独立成库发布；
// 前身 immdb（annoplex internal/immdb，块 + 值子块 + mmap 直访）已由本包取代。
package immkv

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/klauspost/compress/zstd"
)

const (
	keysMagic = "IMKVK002" // v2：微块键流（64 条/块 + 常驻块索引）
	valsMagic = "IMKVV003"

	headerSize   = 64
	recsPerBlock = 64      // 每微块记录数（末块可短）
	flagShare    = 1 << 29 // frameLen bit29：与上一条记录同帧（valOff 相同）

	flagRaw      = 1 << 31   // valLen 最高位：1 = 裸存帧
	flagPacked   = 1 << 30   // valLen 次高位：1 = 打包帧（多值一帧）
	frameLenMask = 1<<29 - 1 // valLen 低 29 位 = 帧长（bit31 裸存、bit30 打包、bit29 同帧）
	maxFrameLen  = 1<<29 - 1 // 帧长度语义上限（512MB，高 3 位是标记）
	maxValLen    = 1<<31 - 1 // 单条值的长度上限（u32 存得下，检查在投入时）
	maxKeyLen    = 1<<16 - 1 // keyLen u16 语义上限

	defaultRawValueBelow = 64
	packBelow            = 32 << 10 // 值小于此进打包帧；大值（>32KB，罕见）按值独立成帧
	packTarget           = 64 << 10 // 打包帧的目标未压字节数
	flushChunk           = 4 << 20  // AddSorted 攒集缓冲上限（key+val 字节数）

	// 每值成帧模式（CreateOptions.PackTarget < 0）的字典训练参数
	trainSamples = 5000      // Contents 条数（训熵表用）
	trainHistory = 512 << 10 // History 上限：实测 512KB 时 chr22 为 bgzip 的 82%，128KB 涨到 110% 破体积线
	dictID       = 1         // 字典 ID；帧头记录它，解码侧按 vals 头还原
)

// CreateOptions 是 Create 的选项。
type CreateOptions struct {
	// RawValueBelow：值短于该字节数时裸存（不压缩、免解压），0 取默认 64；
	// 设成大值即全裸存（对应旧 immdb 的 raw 档）。读侧按每条记录的裸存
	// 标记识别，不依赖这个参数。
	RawValueBelow int
	// PackTarget：打包帧的目标未压字节数，0 取默认 64KB。小包解码与读量都
	// 小、体积大；大包反之。负值表示不打包——每条值独立成帧，并用首批值
	// 训练 zstd 字典嵌进 vals 头（稀疏点查读得最省，体积比打包大 ~20%）。
	// 读侧不依赖这个参数。
	PackTarget int
	// Schema 是字段表。nil/空表示单 String 字段 value：Add 与 AddSorted
	// 收原始值字节，writer 包成 record。非空表示调用方传入的 value 已经
	// 是 record 字节，writer 只落盘、不校验内容。
	Schema []Field
}

// BuildOptions 是 BuildFromVCFs 与 BuildFromTSVs 的构建选项；schema 由输入
// 格式决定，调用方不能指定。
type BuildOptions struct {
	RawValueBelow int
	PackTarget    int
}

// StreamOptions 是 OpenStreaming 的选项。
type StreamOptions struct {
	// Threads 是并发旋钮：同时决定在途 pread 请求数与并行解压 goroutine
	// 数，默认 1。生产口径由调用方在多个句柄间分配总预算。
	Threads int
	// Columns 是字段选择：nil 或空表示库内全部列，非空是 base.manifest 里
	// 的列号；顺序无关，输出始终按 schema 顺序。选择只决定读出哪些字段，
	// 不决定走记录序还是列存——路径由读侧按估算 I/O 成本选。没有
	// base.col/manifest 的库用非空 Columns 会报错。
	Columns []int
}

// ---- Writer ----

// dictCompressor 是字典逐值压缩的后端：CGO 可用时走 libzstd（快 6 倍），
// 否则回退 klauspost 纯 Go。帧格式标准互通，查询端只认帧不认后端。
type dictCompressor interface {
	encode(src []byte) []byte
}

type kvRec struct {
	key []byte
	val []byte
}

// Writer 收集键值并在 Finish 时一次性落盘；或用 AddSorted 流式逐段落盘。
type Writer struct {
	base          string
	keysPath      string
	valsPath      string
	keysTmp       string
	valsTmp       string
	kf, vf        *os.File
	kw, vw        *bufio.Writer
	opts          CreateOptions
	rawBelow      int
	packTarget    int // <0：每值独立成帧模式（配合训练字典）
	wrapRawValues bool
	wrapBuf       []byte
	enc           *zstd.Encoder
	dictComp      dictCompressor // 每值模式的字典压缩后端（libzstd CGO 或 klauspost）
	dictTrained   bool
	dictLen       uint32
	dictOff       uint64

	mode     int8  // 0 未定，1 任意序（Add），2 组升序（AddSorted）；两种模式不可混用
	finished bool  // Finish 已被调用（无论成败）
	closed   bool  // 资源已释放，Abort/Finish 成功后再调用无操作
	err      error // 首个构建错误；一旦置位，后续 Add/AddSorted/Finish 都返回它

	recs   []kvRec // Add 全量驻存；AddSorted 是当前攒集缓冲
	pend   int     // 攒集缓冲里 key+val 的字节数
	total  uint64
	keysAt int64 // keys 文件逻辑写位置（含 64B 头）
	valsAt int64 // vals 文件逻辑写位置（含 64B 头）

	// 微块组装：攒 recsPerBlock 条（键 + 帧长标记）成块落盘；flagShare 由
	// valOff 与上一条的关系推导，块首条记 baseValOff 绝对值
	blkKeys    []byte
	blkKeyEnds []uint32
	blkLens    []uint32
	blkValOff  uint64 // 本块首条记录的 valOff
	prevValOff uint64 // 上一条已发射记录的 valOff
	prevLen    uint32 // 上一条的帧长
	hasPrev    bool
	nInBlk     int
	blockCount uint32
	blkPrefix  []byte // 块首键 8B 前缀（零填充），Finish 时拼索引
	blkOffs    []byte // 块偏移，与 blkPrefix 平行

	curGroup []byte // AddSorted 当前组键（拷贝；nil = 还没有组）
	lastKey  []byte // AddSorted 已进入攒集缓冲的最大键（拷贝）
	grp      []kvRec
}

// Create 创建最终文件为 basePath.keys 与 basePath.vals 的构建器。
// 目标文件或 .tmp 文件已存在时返回错误，绝不覆盖已有库；构建失败后由调用方
// 删除残留再重建。Finish 前磁盘上只有两个 .tmp 文件。CreateOptions.Schema
// 决定 Add/AddSorted 收到的是原始值还是已编码 record：空 schema 由 writer
// 包 record，非空 schema 要求调用方自己编码。
func Create(basePath string, opts CreateOptions) (*Writer, error) {
	if len(opts.Schema) > 0 {
		if err := Schema(opts.Schema).Validate(); err != nil {
			return nil, err
		}
	}
	kp, vp := basePath+".keys", basePath+".vals"
	kt, vt := kp+".tmp", vp+".tmp"
	// .col/.manifest 一并检查：避免与 BuildFromTSVs 等入口的残留混建。
	for _, p := range []string{kp, vp, kt, vt, basePath + ".col", basePath + ".manifest"} {
		if _, err := os.Lstat(p); err == nil {
			return nil, fmt.Errorf("immkv: %s already exists; delete it and rebuild", p)
		} else if !os.IsNotExist(err) {
			return nil, err
		}
	}
	kf, err := os.OpenFile(kt, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return nil, err
	}
	vf, err := os.OpenFile(vt, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		kf.Close()
		os.Remove(kt)
		return nil, err
	}
	enc, err := zstd.NewWriter(nil)
	if err != nil {
		kf.Close()
		vf.Close()
		os.Remove(kt)
		os.Remove(vt)
		return nil, err
	}
	w := &Writer{
		base: basePath, keysPath: kp, valsPath: vp, keysTmp: kt, valsTmp: vt,
		kf: kf, vf: vf,
		kw:   bufio.NewWriterSize(kf, 1<<20),
		vw:   bufio.NewWriterSize(vf, 1<<20),
		opts: opts, enc: enc,
		wrapRawValues: len(opts.Schema) == 0,
	}
	if w.opts.RawValueBelow <= 0 {
		w.rawBelow = defaultRawValueBelow
	} else {
		w.rawBelow = w.opts.RawValueBelow
	}
	switch {
	case w.opts.PackTarget < 0:
		w.packTarget = w.opts.PackTarget // 每值成帧模式
	case w.opts.PackTarget > 0:
		w.packTarget = w.opts.PackTarget
	default:
		w.packTarget = packTarget
	}
	// 两个文件各预留 64B 头，Finish 时回填
	if _, err := w.kw.Write(make([]byte, headerSize)); err != nil {
		w.abort()
		return nil, err
	}
	if _, err := w.vw.Write(make([]byte, headerSize)); err != nil {
		w.abort()
		return nil, err
	}
	w.keysAt, w.valsAt = headerSize, headerSize
	return w, nil
}

// Add 投入一条记录，键值内容会被拷贝，调用方可复用自己的缓冲。键可重复，
// Get 返回排序后的第一条，GetPrefix/Scan 能看到全部。与 AddSorted 不可混用。
// CreateOptions.Schema 为空时 val 是原始值，writer 包成单 String 字段 record；
// 非空时 val 必须已经是 record。键值全量驻内存（ClinVar 规模约 GB 级）；
// TB 级用 AddSorted。键或值超出格式上限（u16 / u31）返回错误。
// 任一错误都会粘住 writer：之后的操作与 Finish 都返回该错误。
func (w *Writer) Add(key, val []byte) error {
	if w.finished {
		return errors.New("immkv: Writer already finished")
	}
	if w.err != nil {
		return w.err
	}
	if w.mode == 2 {
		return w.fail(errors.New("immkv: Add after AddSorted"))
	}
	val = w.prepareValue(val)
	if err := checkRecCaps(key, val); err != nil {
		return w.fail(err)
	}
	w.recs = append(w.recs, kvRec{bytes.Clone(key), bytes.Clone(val)})
	w.mode = 1
	return nil
}

// AddSorted 以「组升序」流式投入一条记录：键值内容会被拷贝。group 按 bytesCompare
// 升序到来（相等表示同组，同组记录必须连续到达），组内顺序任意；攒够 flushChunk
// （4MB）即时落盘。组结束（group 变化或 Finish）时组内按完整键排序后进入落盘路径。
// group 倒序、或键低于已进入攒集缓冲的最大键（组间倒挂在该条投入时即已注定）返回
// error 且该条不进库。内存有界：攒集缓冲 + 当前组（组大小由数据决定）。与 Add
// 不可混用。CreateOptions.Schema 为空时 val 是原始值，writer 包成单 String
// 字段 record；非空时 val 必须已经是 record。
//
// 基因组数据的标准用法：group 传键的位置前缀（如 [chrom][pos]），同一位置
// (ref, alt) 的乱序由库吸收，升序条件收窄为「本染色体内按位置有序」；输入已
// 全排序时 group 传键本身，行为与逐条升序一致。
// 任一错误都会粘住 writer：之后的操作与 Finish 都返回该错误。
func (w *Writer) AddSorted(group, key, val []byte) error {
	if w.finished {
		return errors.New("immkv: Writer already finished")
	}
	if w.err != nil {
		return w.err
	}
	if w.mode == 1 {
		return w.fail(errors.New("immkv: AddSorted after Add"))
	}
	if w.curGroup == nil {
		w.curGroup = bytes.Clone(group)
	} else {
		c := bytes.Compare(group, w.curGroup)
		if c < 0 {
			return w.fail(fmt.Errorf("immkv: AddSorted group out of order: %q < current %q", group, w.curGroup))
		}
		if c > 0 {
			if err := w.closeGroup(); err != nil {
				return w.fail(err)
			}
			w.curGroup = bytes.Clone(group)
		}
	}
	// 到这里 lastKey 已含先前全部已接受记录；键低于它 ⇒ 本组最小键必然更低，
	// 组间倒挂已定，该条立即拒绝
	if w.lastKey != nil && bytes.Compare(key, w.lastKey) < 0 {
		return w.fail(fmt.Errorf("immkv: AddSorted key %q below last emitted %q", key, w.lastKey))
	}
	val = w.prepareValue(val)
	if err := checkRecCaps(key, val); err != nil {
		return w.fail(err)
	}
	w.grp = append(w.grp, kvRec{bytes.Clone(key), bytes.Clone(val)})
	w.mode = 2
	return nil
}

// closeGroup 把当前组按完整键排序、校验组间不倒挂，然后逐条进入攒集缓冲
// （攒够 flushChunk 即落盘）。重复键允许，照常落盘。
func (w *Writer) closeGroup() error {
	if len(w.grp) == 0 {
		return nil
	}
	sort.Slice(w.grp, func(i, j int) bool { return bytes.Compare(w.grp[i].key, w.grp[j].key) < 0 })
	if w.lastKey != nil && bytes.Compare(w.grp[0].key, w.lastKey) < 0 {
		return fmt.Errorf("immkv: AddSorted key inversion across groups: %q < %q", w.grp[0].key, w.lastKey)
	}
	for i := range w.grp {
		r := w.grp[i]
		w.lastKey = append(w.lastKey[:0], r.key...)
		w.recs = append(w.recs, r)
		w.pend += len(r.key) + len(r.val)
		if w.pend >= flushChunk {
			if err := w.flushPending(); err != nil {
				return err
			}
		}
	}
	w.grp = w.grp[:0]
	return nil
}

// flushPending 把攒集缓冲里的记录依序落盘：短于 rawBelow 的值裸存；打包模式
// 下小于 packBelow 的值攒进打包帧、其余独立成帧；每值模式（PackTarget<0）下
// 除裸存外全部独立成帧，并在首次落盘前用攒集缓冲里的前 trainSamples 条值
// 训练 zstd 字典（正文截 trainHistory），字典嵌进 vals 头，之后全部值帧用它
// 压缩——稀疏点查的读取粒度与解码量都按单值计。
func (w *Writer) flushPending() error {
	if len(w.recs) == 0 {
		return nil
	}
	if w.packTarget < 0 && !w.dictTrained {
		if err := w.trainDict(); err != nil {
			return err
		}
	}
	if w.packTarget >= 0 {
		return w.flushPacked()
	}
	for i := range w.recs {
		r := w.recs[i]
		if len(r.val) < w.rawBelow {
			if err := w.emitRec(r.key, uint64(w.valsAt), uint32(len(r.val))|flagRaw); err != nil {
				return err
			}
			if _, err := w.vw.Write(r.val); err != nil {
				return err
			}
			w.valsAt += int64(len(r.val))
			continue
		}
		var comp []byte
		if w.dictComp != nil {
			comp = w.dictComp.encode(r.val)
		} else {
			comp = w.enc.EncodeAll(r.val, nil)
		}
		if len(comp) > maxFrameLen {
			panic(fmt.Sprintf("immkv: 值帧 %d B 超过上限", len(comp)))
		}
		if err := w.emitRec(r.key, uint64(w.valsAt), uint32(len(comp))); err != nil {
			return err
		}
		if _, err := w.vw.Write(comp); err != nil {
			return err
		}
		w.valsAt += int64(len(comp))
	}
	w.recs = w.recs[:0]
	w.pend = 0
	return nil
}

// trainDict 用攒集缓冲里的首批值训练 zstd 字典，重建带字典的编码器，并把
// 字典写到 vals 头之后（记录 dictOff/dictLen 供读侧还原）。
func (w *Writer) trainDict() error {
	w.dictTrained = true
	n := len(w.recs)
	if n > trainSamples {
		n = trainSamples
	}
	contents := make([][]byte, 0, n)
	var hist []byte
	for i := 0; i < n; i++ {
		v := w.recs[i].val
		contents = append(contents, v)
		if len(hist) < trainHistory {
			take := len(v)
			if rem := trainHistory - len(hist); take > rem {
				take = rem
			}
			hist = append(hist, v[:take]...)
		}
	}
	if len(hist) < 8 || n == 0 {
		return nil // 值太少（或全空）：不写字典，裸 zstd 帧
	}
	dict, err := zstd.BuildDict(zstd.BuildDictOptions{
		ID:       dictID,
		Contents: contents,
		History:  hist,
	})
	if err != nil {
		return err
	}
	w.dictOff = headerSize
	w.dictLen = uint32(len(dict))
	if _, err := w.vw.Write(dict); err != nil {
		return err
	}
	w.valsAt += int64(len(dict))
	w.enc.Close()
	enc, err := zstd.NewWriter(nil, zstd.WithEncoderDict(dict), zstd.WithEncoderCRC(false))
	if err != nil {
		return err
	}
	w.enc = enc
	// 每值压缩优先 libzstd CGO（CDict，实测 630MB/s/核，klauspost 纯 Go
	// 102MB/s）；不可用（CGO_ENABLED=0）或建器失败时 dictComp 留 nil，
	// flushPending 回退 w.enc（klauspost 字典编码器）。帧格式两后端互通。
	if dc, derr := newCGODictCompressor(dict); derr == nil {
		w.dictComp = dc
	}
	return nil
}

// flushPacked 是打包模式的落盘实现。
func (w *Writer) flushPacked() error {
	var (
		packKeys  [][]byte // 打包成员的键（键流记录推迟到帧成形时写）
		packVals  [][]byte
		packFirst int // 首成员的全局记录号（写进组头，读侧换算组内下标）
		packBytes int
	)
	closePack := func() error {
		if len(packVals) == 0 {
			return nil
		}
		frame := make([]byte, 8+4*len(packVals)+packBytes)
		binary.LittleEndian.PutUint32(frame[0:4], uint32(packFirst))
		binary.LittleEndian.PutUint32(frame[4:8], uint32(len(packVals)))
		o := 8
		for _, v := range packVals {
			binary.LittleEndian.PutUint32(frame[o:o+4], uint32(len(v)))
			o += 4
		}
		for _, v := range packVals {
			copy(frame[o:], v)
			o += len(v)
		}
		comp := w.enc.EncodeAll(frame, nil)
		if len(comp) > maxFrameLen {
			panic(fmt.Sprintf("immkv: 打包帧 %d B 超过 u30 上限", len(comp)))
		}
		valOff := uint64(w.valsAt)
		mark := uint32(len(comp)) | flagPacked
		for _, k := range packKeys {
			if err := w.emitRec(k, valOff, mark); err != nil {
				return err
			}
		}
		if _, err := w.vw.Write(comp); err != nil {
			return err
		}
		w.valsAt += int64(len(comp))
		packKeys, packVals, packBytes = packKeys[:0], packVals[:0], 0
		return nil
	}
	for i := range w.recs {
		r := w.recs[i]
		switch {
		case len(r.val) < w.rawBelow:
			if err := closePack(); err != nil {
				return err
			}
			if err := w.emitRec(r.key, uint64(w.valsAt), uint32(len(r.val))|flagRaw); err != nil {
				return err
			}
			if _, err := w.vw.Write(r.val); err != nil {
				return err
			}
			w.valsAt += int64(len(r.val))
		case len(r.val) < packBelow:
			if packBytes > 0 && packBytes+len(r.val) > w.packTarget {
				if err := closePack(); err != nil {
					return err
				}
			}
			if len(packVals) == 0 {
				packFirst = int(w.total)
			}
			packKeys = append(packKeys, r.key)
			packVals = append(packVals, r.val)
			packBytes += len(r.val)
		default:
			if err := closePack(); err != nil {
				return err
			}
			comp := w.enc.EncodeAll(r.val, nil)
			if len(comp) > maxFrameLen {
				panic(fmt.Sprintf("immkv: 值帧 %d B 超过 u30 上限", len(comp)))
			}
			if err := w.emitRec(r.key, uint64(w.valsAt), uint32(len(comp))); err != nil {
				return err
			}
			if _, err := w.vw.Write(comp); err != nil {
				return err
			}
			w.valsAt += int64(len(comp))
		}
	}
	if err := closePack(); err != nil {
		return err
	}
	w.recs = w.recs[:0]
	w.pend = 0
	return nil
}

// fail 记录首个构建错误并返回它；后续操作与 Finish 都返回同一个错误。
func (w *Writer) fail(err error) error {
	if w.err == nil {
		w.err = err
	}
	return w.err
}

// checkRecCaps 检查键值长度是否在格式上限内，超出返回错误。
// prepareValue 在空 schema 的库里把原始值包成单 String 字段的 record：
// 8B bitmap（bit0=1）+ uvarint(len) + 原字节。非空 schema 的库认为调用方
// 已经编码好 record，原样落盘。
func (w *Writer) prepareValue(val []byte) []byte {
	if !w.wrapRawValues {
		return val
	}
	w.wrapBuf = w.wrapBuf[:0]
	var bitmap [8]byte
	binary.LittleEndian.PutUint64(bitmap[:], 1)
	w.wrapBuf = append(w.wrapBuf, bitmap[:]...)
	w.wrapBuf = binary.AppendUvarint(w.wrapBuf, uint64(len(val)))
	w.wrapBuf = append(w.wrapBuf, val...)
	return w.wrapBuf
}

func checkRecCaps(key, val []byte) error {
	if len(key) > maxKeyLen {
		return fmt.Errorf("immkv: key 长度 %d 超过 u16 上限 %d", len(key), maxKeyLen)
	}
	if len(val) > maxValLen {
		return fmt.Errorf("immkv: 值长度 %d 超过 u31 上限 %d", len(val), maxValLen)
	}
	return nil
}

// emitRec 发射一条键流记录进微块组装器；frameMark 含裸存/打包标记（低 29 位
// 是帧长）。valOff 必须等于「上一条的 valOff + 上一条的帧长」或与上一条相同
// （打包帧成员）——写侧不变式，违反即 panic。攒满 recsPerBlock 条即落一块。
func (w *Writer) emitRec(key []byte, valOff uint64, frameMark uint32) error {
	mark := frameMark &^ flagShare
	if w.hasPrev {
		if valOff == w.prevValOff {
			mark |= flagShare
		} else if valOff != w.prevValOff+uint64(w.prevLen) {
			panic(fmt.Sprintf("immkv: 写侧不变式破坏：valOff %d ≠ 前条 %d+%d", valOff, w.prevValOff, w.prevLen))
		}
	} else {
		w.blkValOff = valOff
	}
	if w.nInBlk == 0 {
		w.blkValOff = valOff
	}
	w.blkKeys = append(w.blkKeys, key...)
	w.blkKeyEnds = append(w.blkKeyEnds, uint32(len(w.blkKeys)))
	w.blkLens = append(w.blkLens, mark)
	w.prevValOff, w.prevLen, w.hasPrev = valOff, mark&frameLenMask, true
	w.nInBlk++
	w.total++
	if w.nInBlk == recsPerBlock {
		return w.flushBlock()
	}
	return nil
}

// flushBlock 落一个微块：[keyEnds u32×n][lens u32×n][baseValOff u64][keyData]，
// 并把块首键 8B 前缀与块偏移记进索引材料。u32 的 keyEnd 使单块 keyData 上限
// 4GB，远超 64 条键的规模。
func (w *Writer) flushBlock() error {
	if w.nInBlk == 0 {
		return nil
	}
	// 记索引材料：块偏移 + 首键前 8B（零填充）
	w.blkOffs = binary.LittleEndian.AppendUint64(w.blkOffs, uint64(w.keysAt))
	var pfx [8]byte
	copy(pfx[:], w.blkKeys[:w.blkKeyEnds[0]])
	w.blkPrefix = append(w.blkPrefix, pfx[:]...)
	// 块体：[keyEnds u32×n][lens u32×n][baseValOff u64][keyData]
	buf := make([]byte, 0, 8*w.nInBlk+8+len(w.blkKeys))
	for _, e := range w.blkKeyEnds {
		buf = binary.LittleEndian.AppendUint32(buf, e)
	}
	buf = append(buf, w.blkLensToBytes()...)
	buf = binary.LittleEndian.AppendUint64(buf, w.blkValOff)
	buf = append(buf, w.blkKeys...)
	if _, err := w.kw.Write(buf); err != nil {
		return err
	}
	w.keysAt += int64(len(buf))
	w.blockCount++
	w.blkKeys = w.blkKeys[:0]
	w.blkKeyEnds = w.blkKeyEnds[:0]
	w.blkLens = w.blkLens[:0]
	w.nInBlk = 0
	return nil
}

func (w *Writer) blkLensToBytes() []byte {
	out := make([]byte, 4*len(w.blkLens))
	for i, l := range w.blkLens {
		binary.LittleEndian.PutUint32(out[4*i:], l)
	}
	return out
}

// Finish 排序（Add 模式）、落盘锚点、回填两个文件头、fsync、rename，然后
// 关闭构建器。再次调用返回错误。构建期间出现过错误时，Finish 返回该错误且
// 不发布任何文件。任何一步失败都会清掉本次构建的全部产物（.tmp 与已
// rename 出去的正式文件）。
func (w *Writer) Finish() error {
	if w.finished {
		return errors.New("immkv: Writer already finished")
	}
	w.finished = true
	if w.err != nil {
		w.closeEnc()
		cerr := w.abort()
		w.closed = true // 资源已释放；后续 Abort 不再执行清理，避免误删同路径新 writer 的文件
		if cerr != nil {
			return fmt.Errorf("%w; cleanup: %v", w.err, cerr)
		}
		return w.err
	}
	err := w.finish()
	w.closeEnc()
	if err != nil {
		cerr := w.abort()
		w.closed = true
		if cerr != nil {
			return fmt.Errorf("%w; cleanup: %v", err, cerr)
		}
		return err
	}
	w.closed = true
	return nil
}

// Abort 放弃构建：关闭文件与编码器、删除本次构建的产物。幂等；Finish
// （无论成败）之后调用是无操作并返回 nil——清理只做一次，重试删除只会
// 误删同路径新 writer 的文件。
func (w *Writer) Abort() error {
	if w.closed {
		return nil
	}
	w.finished = true
	w.closeEnc()
	if err := w.abort(); err != nil {
		return err
	}
	w.closed = true
	return nil
}

func (w *Writer) closeEnc() {
	if w.enc != nil {
		w.enc.Close()
		w.enc = nil
	}
}

func (w *Writer) finish() error {
	if w.mode == 2 {
		// 流式：组与攒集缓冲已在 AddSorted 里逐段落盘，只剩尾部一组一段
		if err := w.closeGroup(); err != nil {
			return err
		}
		if err := w.flushPending(); err != nil {
			return err
		}
	} else if w.mode == 1 {
		sort.Slice(w.recs, func(i, j int) bool {
			return bytes.Compare(w.recs[i].key, w.recs[j].key) < 0
		})
		if err := w.flushPending(); err != nil {
			return err
		}
	}
	if err := w.flushBlock(); err != nil { // 尾部不足一块的记录
		return err
	}
	indexOff := w.keysAt
	if _, err := w.kw.Write(w.blkPrefix); err != nil {
		return err
	}
	if _, err := w.kw.Write(w.blkOffs); err != nil {
		return err
	}
	if err := w.kw.Flush(); err != nil {
		return err
	}
	if err := w.vw.Flush(); err != nil {
		return err
	}
	if err := w.writeKeysHeader(uint64(indexOff)); err != nil {
		return err
	}
	if err := w.writeValsHeader(); err != nil {
		return err
	}
	if err := w.kf.Sync(); err != nil {
		return err
	}
	if err := w.vf.Sync(); err != nil {
		return err
	}
	if err := w.kf.Close(); err != nil {
		return err
	}
	w.kf = nil
	if err := w.vf.Close(); err != nil {
		return err
	}
	w.vf = nil
	// vals 先换、keys 后换：keys 文件的 rename 是构建可见性的分界
	if err := os.Rename(w.valsTmp, w.valsPath); err != nil {
		return err
	}
	return os.Rename(w.keysTmp, w.keysPath)
}

func (w *Writer) writeKeysHeader(indexOff uint64) error {
	h := make([]byte, headerSize)
	copy(h[0:8], keysMagic)
	binary.LittleEndian.PutUint32(h[8:12], 1) // version
	binary.LittleEndian.PutUint64(h[12:20], w.total)
	binary.LittleEndian.PutUint16(h[20:22], recsPerBlock)
	binary.LittleEndian.PutUint32(h[24:28], w.blockCount)
	binary.LittleEndian.PutUint64(h[28:36], indexOff)
	binary.LittleEndian.PutUint32(h[36:40], uint32(w.rawBelow))
	_, err := w.kf.WriteAt(h, 0)
	return err
}

func (w *Writer) writeValsHeader() error {
	h := make([]byte, headerSize)
	copy(h[0:8], valsMagic)
	binary.LittleEndian.PutUint32(h[8:12], 3) // version
	binary.LittleEndian.PutUint64(h[12:20], w.total)
	binary.LittleEndian.PutUint32(h[20:24], w.dictLen)
	binary.LittleEndian.PutUint64(h[24:32], w.dictOff)
	_, err := w.vf.WriteAt(h, 0)
	return err
}

// abort 清理失败路径：关闭文件、删除本次构建的产物，返回首个清理错误。
// 正式路径若存在必为本次 rename 出去的文件（Create 的独占检查保证），
// 删除它们不会碰到别人的库。
func (w *Writer) abort() error {
	var first error
	if w.kf != nil {
		if err := w.kf.Close(); err != nil {
			first = err
		}
		w.kf = nil
	}
	if w.vf != nil {
		if err := w.vf.Close(); err != nil && first == nil {
			first = err
		}
		w.vf = nil
	}
	for _, p := range []string{w.keysTmp, w.valsTmp, w.keysPath, w.valsPath} {
		if err := os.Remove(p); err != nil && !os.IsNotExist(err) && first == nil {
			first = err
		}
	}
	return first
}
