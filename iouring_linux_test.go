//go:build linux

package immkv

import (
	"os"
	"runtime"
	"sync/atomic"
	"syscall"
	"testing"
	"time"
	"unsafe"

	"golang.org/x/sys/unix"
)

func TestDirectBufferAlignmentAndCapacity(t *testing.T) {
	for _, n := range []int{1, 4096, 8192, 8193, 65536, directMaxBucket, directMaxBucket + 4096} {
		for j := 0; j < 4; j++ {
			buf := getDirectBuf(n)
			if len(buf) != n || cap(buf) < n {
				t.Fatalf("size %d: len=%d cap=%d", n, len(buf), cap(buf))
			}
			if uintptr(unsafe.Pointer(&buf[0]))%directAlign != 0 {
				t.Fatalf("size %d: buffer is not aligned", n)
			}
			wantCap := n
			if n <= directMaxBucket {
				wantCap = 8 << 10 << directBucket(n)
			}
			if cap(buf) != wantCap {
				t.Fatalf("size %d: cap=%d, want %d", n, cap(buf), wantCap)
			}
			putDirectBuf(buf)
		}
	}
}

func TestIOUringCloseUnmaps(t *testing.T) {
	r := &ioUring{fd: -1}
	for i := range r.mmaps {
		b, err := unix.Mmap(-1, 0, 4096, unix.PROT_READ|unix.PROT_WRITE, unix.MAP_PRIVATE|unix.MAP_ANON)
		if err != nil {
			r.close()
			t.Fatal(err)
		}
		r.mmaps[i] = b
	}
	mmaps := r.mmaps
	r.close()
	r.close()
	for i, b := range mmaps {
		if err := unix.Munmap(b); err != syscall.EINVAL {
			t.Errorf("mapping %d still registered after close: second unmap=%v", i, err)
		}
		if r.mmaps[i] != nil {
			t.Errorf("mapping %d retained after close", i)
		}
	}
}

func TestIOUringPartialSubmission(t *testing.T) {
	r, err := newIOUring(8)
	if err != nil {
		t.Skipf("io_uring unavailable: %v", err)
	}
	defer r.close()
	f, err := os.Open("/dev/zero")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	bufs := [][]byte{make([]byte, 16), make([]byte, 16), make([]byte, 16)}
	defer runtime.KeepAlive(bufs)
	for i, b := range bufs {
		if err := r.prep(f, 0, b, uint64(i)); err != nil {
			t.Fatal(err)
		}
	}
	// 中间的非法操作强制内核提前结束一次提交，后面的合法读取仍须提交。
	r.sqes[1].opcode = 255
	if err := r.submitWait(0); err != nil {
		t.Fatal(err)
	}
	if got := atomic.LoadUint32(r.sqHead); got != r.sqeTail {
		t.Fatalf("submitted %d of %d SQEs", got, r.sqeTail)
	}
	results := make(map[uint64]int32)
	for len(results) < len(bufs) {
		if err := r.submitWait(1); err != nil {
			t.Fatal(err)
		}
		r.reap(func(id uint64, res int32) { results[id] = res })
	}
	if results[0] != 16 || results[1] != -int32(syscall.EINVAL) || results[2] != 16 {
		t.Fatalf("completion results=%v", results)
	}
}

func TestDirectReadErrorReturns(t *testing.T) {
	for _, chunks := range []int{1, 3} {
		r, err := newIOUring(64)
		if err != nil {
			t.Skipf("io_uring unavailable: %v", err)
		}
		f, err := os.Open("/dev/null")
		if err != nil {
			r.close()
			t.Fatal(err)
		}
		extCh := make(chan []extent, chunks)
		decCh := make(chan []extent, chunks)
		abort := make(chan struct{})
		errs := make(chan error, 1)
		done := make(chan struct{})
		for i := 0; i < chunks; i++ {
			extCh <- []extent{{s: &shard{vfDirect: f}, hits: []sHit{{frameLen: 1}}}}
		}
		close(extCh)
		go func() {
			defer close(done)
			driveDirect(r, extCh, decCh, abort, func(err error) {
				select {
				case errs <- err:
				default:
				}
			})
		}()
		select {
		case <-done:
		case <-time.After(2 * time.Second):
			t.Fatal("read driver did not return after short read")
		}
		r.close()
		f.Close()
		select {
		case err := <-errs:
			if err == nil {
				t.Fatal("short read returned nil error")
			}
		default:
			t.Fatal("short read was not reported")
		}
		if len(decCh) != 0 {
			t.Fatal("failed read dispatched to decompression")
		}
	}
}

func TestDriveDirectInFlightRespectsRingDepth(t *testing.T) {
	r, err := newIOUring(64)
	if err != nil {
		t.Skipf("io_uring unavailable: %v", err)
	}
	dir, err := os.MkdirTemp(".", ".tmp-iouring-test-*")
	if err != nil {
		r.close()
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)
	path := dir + "/vals"
	buf := make([]byte, 4<<20)
	if err := os.WriteFile(path, buf, 0o600); err != nil {
		r.close()
		t.Fatal(err)
	}
	f, err := openDirect(path)
	if err != nil {
		r.close()
		t.Skipf("O_DIRECT unavailable: %v", err)
	}
	defer f.Close()

	const (
		numChunks      = 8
		extentsPerChunk = 64
	)
	s := &shard{vfDirect: f}
	extCh := make(chan []extent, numChunks)
	decCh := make(chan []extent)
	abort := make(chan struct{})

	for c := 0; c < numChunks; c++ {
		ch := make([]extent, extentsPerChunk)
		for i := 0; i < extentsPerChunk; i++ {
			off := uint64(((c*extentsPerChunk + i) * directAlign) % len(buf))
			ch[i] = extent{
				s:    s,
				off:  off,
				hits: []sHit{{valOff: off, frameLen: directAlign}},
			}
		}
		extCh <- ch
	}
	close(extCh)

	var maxInFlight uint32
	stopSample := make(chan struct{})
	sampled := make(chan struct{})
	go func() {
		defer close(sampled)
		for {
			select {
			case <-stopSample:
				return
			default:
				submitted := atomic.LoadUint32(r.sqHead)
				reaped := atomic.LoadUint32(r.cqHead)
				inFlight := submitted - reaped
				if inFlight > maxInFlight {
					maxInFlight = inFlight
				}
			}
		}
	}()

	done := make(chan struct{})
	var derr error
	go func() {
		defer close(done)
		driveDirect(r, extCh, decCh, abort, func(err error) {
			derr = err
		})
	}()

	for c := 0; c < numChunks; c++ {
		ch := <-decCh
		if c == 0 {
			close(stopSample)
			<-sampled
		}
		for i := range ch {
			putDirectBuf(ch[i].dbuf)
		}
	}
	<-done
	if derr != nil {
		t.Fatalf("driveDirect failed: %v", derr)
	}
	if maxInFlight > uint32(len(r.sqes)) {
		t.Fatalf("kernel in-flight requests (sqHead - cqHead) reached %d, exceeding ring depth %d", maxInFlight, len(r.sqes))
	}
}
