//go:build linux

package immkv

import (
	"testing"
	"unsafe"
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
