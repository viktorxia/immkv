//go:build linux

package immkv

import (
	"fmt"
	"os"
	"sync"
	"sync/atomic"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// io_uring 直连 + O_DIRECT：值文件以 O_DIRECT 打开（绕开页缓存——3GiB cgroup
// 下缓冲读的页缓存记账回收会把单读延迟拖到 ~220µs），读请求经 512 深的提交
// 队列异步下发，单个驱动 goroutine 完成提交与收割——在途读深度与线程数彻底
// 解耦，线程预算仍由 StreamOptions.Threads 约束在解压池上。O_DIRECT 要求 4KB 对齐：
// 每段读 [对齐下界, 对齐上界)，命中视图按段内偏移切出。建环或 O_DIRECT 不可
// 用时 GetBatch 退回 pread 读池。要求内核 ≥5.6。

const (
	ioringOpRead       = 22
	ioringEnterGetEvts = 1
	ioringOffSQRing    = 0
	ioringOffCQRing    = 0x8000000
	ioringOffSQEs      = 0x10000000
	sysIoringSetup     = 425
	sysIoringEnter     = 426

	directAlign = 4096 // O_DIRECT 逻辑块对齐
)

type ioringSQOffsets struct {
	head, tail, ringMask, ringEntries, flags, dropped, array, resv1 uint32
	resv2                                                           uint64
}

type ioringCQOffsets struct {
	head, tail, ringMask, ringEntries, overflow, cqes, flags, resv1 uint32
	resv2                                                           uint64
}

type ioringParams struct {
	sqEntries, cqEntries, flags, sqThreadCPU, sqThreadIdle, features, wqFd uint32
	resv                                                                   [3]uint32
	sqOff                                                                  ioringSQOffsets
	cqOff                                                                  ioringCQOffsets
}

type ioringSQE struct {
	opcode      uint8
	flags       uint8
	ioprio      uint16
	fd          int32
	off         uint64
	addr        uint64
	len         uint32
	rwFlags     uint32
	userData    uint64
	bufIndex    uint16
	personality uint16
	spliceFdIn  int32
	addr3       uint64
	pad         [8]byte
}

type ioringCQE struct {
	userData uint64
	res      int32
	flags    uint32
}

// ioUring 持有一套已 mmap 的 SQ/CQ 环。非并发安全：由单个驱动 goroutine
// 独占（GetBatch 每次调用自建自毁，跨调用天然隔离）。
type ioUring struct {
	fd      int
	sqMask  uint32
	cqMask  uint32
	sqHead  *uint32 // SQ：内核已消费位置（只读）
	sqTail  *uint32 // SQ：本端发布位置（写）
	sqArray []uint32
	cqHead  *uint32 // CQ：本端收割位置（写）
	cqTail  *uint32 // CQ：内核写位置（只读）
	cqes    []ioringCQE
	sqes    []ioringSQE
	sqeTail uint32 // 下一个待发布的 SQE 序号（单调递增，槽 = 序号 & mask）
	mmaps   [3][]byte
}

func newIOUring(entries uint32) (*ioUring, error) {
	var p ioringParams
	fd, _, errno := syscall.Syscall(sysIoringSetup, uintptr(entries), uintptr(unsafe.Pointer(&p)), 0)
	if errno != 0 {
		return nil, errno
	}
	r := &ioUring{fd: int(fd)}
	mk := func(off int64, sz int) ([]byte, error) {
		return unix.Mmap(r.fd, off, sz, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_SHARED|syscall.MAP_POPULATE)
	}
	sqSize := int(p.sqOff.array) + int(p.sqEntries)*4
	cqSize := int(p.cqOff.cqes) + int(p.cqEntries)*int(unsafe.Sizeof(ioringCQE{}))
	sqeSize := int(p.sqEntries) * int(unsafe.Sizeof(ioringSQE{}))
	var err error
	if r.mmaps[0], err = mk(ioringOffSQRing, sqSize); err != nil {
		r.close()
		return nil, err
	}
	if r.mmaps[1], err = mk(ioringOffCQRing, cqSize); err != nil {
		r.close()
		return nil, err
	}
	if r.mmaps[2], err = mk(ioringOffSQEs, sqeSize); err != nil {
		r.close()
		return nil, err
	}
	sq, cq := r.mmaps[0], r.mmaps[1]
	ptr := func(b []byte, off uint32) *uint32 {
		return (*uint32)(unsafe.Add(unsafe.Pointer(&b[0]), int(off)))
	}
	r.sqHead = ptr(sq, p.sqOff.head)
	r.sqTail = ptr(sq, p.sqOff.tail)
	r.sqMask = *ptr(sq, p.sqOff.ringMask)
	r.sqArray = unsafe.Slice((*uint32)(unsafe.Add(unsafe.Pointer(&sq[0]), int(p.sqOff.array))), int(p.sqEntries))
	r.cqHead = ptr(cq, p.cqOff.head)
	r.cqTail = ptr(cq, p.cqOff.tail)
	r.cqMask = *ptr(cq, p.cqOff.ringMask)
	r.cqes = unsafe.Slice((*ioringCQE)(unsafe.Add(unsafe.Pointer(&cq[0]), int(p.cqOff.cqes))), int(p.cqEntries))
	r.sqes = unsafe.Slice((*ioringSQE)(unsafe.Add(unsafe.Pointer(&r.mmaps[2][0]), 0)), int(p.sqEntries))
	return r, nil
}

func (r *ioUring) close() {
	for i, m := range r.mmaps {
		if m != nil {
			syscall.Munmap(m)
			r.mmaps[i] = nil
		}
	}
	if r.fd >= 0 {
		syscall.Close(r.fd)
		r.fd = -1
	}
}

// inFlightSQEs 报已提交未回收槽位的 SQE 数（决定还能不能继续排）。
func (r *ioUring) inFlightSQEs() uint32 {
	return r.sqeTail - atomic.LoadUint32(r.sqHead)
}

// prep 排一个读请求；SQ 满时先 submitWait(0) 腾槽。
func (r *ioUring) prep(f *os.File, off uint64, buf []byte, userData uint64) error {
	if int(r.inFlightSQEs()) >= len(r.sqes) {
		if err := r.submitWait(0); err != nil {
			return err
		}
	}
	idx := r.sqeTail & r.sqMask
	r.sqes[idx] = ioringSQE{
		opcode:   ioringOpRead,
		fd:       int32(f.Fd()),
		off:      off,
		addr:     uint64(uintptr(unsafe.Pointer(&buf[0]))),
		len:      uint32(len(buf)),
		userData: userData,
	}
	r.sqArray[idx] = idx
	r.sqeTail++
	return nil
}

// submitWait 提交全部已排 SQE；minComplete > 0 时阻塞等至少这么多完成事件。
// Go 的异步抢占信号（SIGURG）会以 EINTR 打断 enter，按惯例重试。
func (r *ioUring) submitWait(minComplete uint32) error {
	pending := r.sqeTail - atomic.LoadUint32(r.sqTail)
	if pending > 0 {
		atomic.StoreUint32(r.sqTail, r.sqeTail) // store-release：SQE 与数组项先行
	}
	if pending == 0 && minComplete == 0 {
		return nil
	}
	for {
		_, _, errno := syscall.Syscall6(sysIoringEnter, uintptr(r.fd), uintptr(pending),
			uintptr(minComplete), ioringEnterGetEvts, 0, 0)
		if errno == 0 {
			return nil
		}
		if errno != syscall.EINTR {
			return errno
		}
	}
}

// reap 收割全部已完成 CQE（无阻塞），逐个回调。
func (r *ioUring) reap(fn func(userData uint64, res int32)) {
	head := atomic.LoadUint32(r.cqHead)
	tail := atomic.LoadUint32(r.cqTail)
	n := 0
	for head != tail {
		cqe := &r.cqes[head&r.cqMask]
		fn(cqe.userData, cqe.res)
		head++
		n++
	}
	if n > 0 {
		atomic.StoreUint32(r.cqHead, head)
	}
}

// directBufPools 按 2 的幂分桶的 4KB 对齐读缓冲（8K–256K）。对齐在块创建时
// 一次达成，复用不破坏；块容量即桶大小，归还按容量分桶。
var directBufPools [6]sync.Pool

func directBucket(n int) int {
	b := 8 << 10
	i := 0
	for b < n && i < len(directBufPools)-1 {
		b <<= 1
		i++
	}
	return i
}

const directMaxBucket = 256 << 10

// getDirectBuf 取一块 4KB 对齐、长度 n 的缓冲；超过最大桶（256KB，读段合并
// 的上限是 4MB）就现配现弃，对齐照做。
func getDirectBuf(n int) []byte {
	if n > directMaxBucket {
		raw := make([]byte, n+directAlign)
		off := int((-uintptr(unsafe.Pointer(&raw[0]))) & (directAlign - 1))
		return raw[off : off+n : n]
	}
	i := directBucket(n)
	if v := directBufPools[i].Get(); v != nil {
		return (*v.(*[]byte))[:n]
	}
	size := 8 << 10 << i
	raw := make([]byte, size+directAlign)
	off := int((-uintptr(unsafe.Pointer(&raw[0]))) & (directAlign - 1))
	b := raw[off : off+size : size] // 4KB 对齐、cap == 桶大小
	return b[:n]
}

// putDirectBuf 归还对齐缓冲（按容量分桶；超桶的不收）。
func putDirectBuf(b []byte) {
	if cap(b) > directMaxBucket {
		return
	}
	directBufPools[directBucket(cap(b))].Put(&b)
}

// driveDirect 是 O_DIRECT + io_uring 路径的读驱动（独占 ring 的唯一 goroutine），
// 事件循环：①非阻塞收新块，为每段算对齐区间、取对齐缓冲、排 SQE ②提交
// ③收割 CQE，整块完成的挂进 ready ④ready 非阻塞投给 decCh——投递绝不阻塞
// 驱动，CQ 才能持续收割不溢出；无事可做且在途非空时 enter(WAIT) 睡等完成。
// userData 打包为 块号<<8|段号（块 ≤64 段）。任何错误调 fail 后退出，在途
// 请求随 ring 销毁作废。
func driveDirect(r *ioUring, extCh <-chan []extent, decCh chan<- []extent,
	abort <-chan struct{}, fail func(error)) {

	type cstate struct {
		chunk []extent
		left  int
	}
	pending := map[uint64]*cstate{}
	ready := make([]*cstate, 0, 16)
	var nextUD uint64
	drained := false
	dead := false
	dfail := func(err error) {
		dead = true
		fail(err)
	}

	prep := func(ch []extent) {
		st := &cstate{chunk: ch, left: 0}
		ud := nextUD
		nextUD++
		pending[ud] = st
		for i := range ch {
			e := &ch[i]
			last := e.hits[len(e.hits)-1]
			e.size = last.valOff + uint64(last.frameLen) - e.off
			if e.size == 0 {
				continue // 零长裸存帧独占的段：无需读
			}
			// O_DIRECT 对齐区间与命中视图
			lo := e.off &^ (directAlign - 1)
			hi := (e.off + e.size + directAlign - 1) &^ (directAlign - 1)
			e.dbuf = getDirectBuf(int(hi - lo))
			e.buf = e.dbuf[e.off-lo : e.off-lo+e.size]
			e.dNeed = int32(e.off - lo + e.size)
			e.direct = true
			if err := r.prep(e.s.vfDirect, lo, e.dbuf, ud<<8|uint64(i)); err != nil {
				dfail(err)
				return
			}
			st.left++
		}
		if st.left == 0 {
			delete(pending, ud)
			ready = append(ready, st)
		}
	}
	complete := func(udFull uint64, res int32) {
		st := pending[udFull>>8]
		if st == nil {
			return
		}
		e := &st.chunk[udFull&0xFF]
		// 对齐读可能越过文件尾：EOF 短读合法，覆盖到所需字节即成功
		if res < 0 || res < e.dNeed {
			fail(fmt.Errorf("immkv: O_DIRECT 读失败：res=%d need=%d", res, e.dNeed))
			dead = true
			return
		}
		st.left--
		if st.left == 0 {
			delete(pending, udFull>>8)
			ready = append(ready, st)
		}
	}
	dispatch := func() bool {
		sent := false
		for len(ready) > 0 {
			st := ready[0]
			select {
			case decCh <- st.chunk:
				ready = ready[1:]
				sent = true
			default:
				return sent
			}
		}
		return sent
	}

	for {
		if dead {
			return
		}
		// ① 非阻塞收新块：SQ 有富余槽位才收
		for !drained && int(r.inFlightSQEs())+64 <= len(r.sqes) {
			var ch []extent
			var ok bool
			select {
			case ch, ok = <-extCh:
				if !ok {
					drained = true
					break
				}
				prep(ch)
			default:
				ok = false
				ch = nil
			}
			if ch == nil {
				break
			}
		}
		// ② 提交
		if err := r.submitWait(0); err != nil {
			dfail(err)
			return
		}
		// ③ 收割
		r.reap(complete)
		// ④ 非阻塞投递
		dispatch()
		// ⑤ 推进或等待（任何路径要么有进展要么睡，禁止空转）
		if drained && len(pending) == 0 && len(ready) == 0 {
			return
		}
		if len(pending) > 0 {
			if err := r.submitWait(1); err != nil {
				dfail(err)
				return
			}
			continue
		}
		if len(ready) > 0 {
			st := ready[0]
			select {
			case decCh <- st.chunk:
				ready = ready[1:]
			case <-abort:
				return
			}
			continue
		}
		select {
		case ch, ok := <-extCh:
			if !ok {
				drained = true
				continue
			}
			prep(ch)
		case <-abort:
			return
		}
	}
}
