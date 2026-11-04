//go:build !linux

package immkv

import "errors"

// ioUring 在非 linux 平台不可用，GetBatch 走 pread 读池。
type ioUring struct{}

var errIOUringUnsupported = errors.New("immkv: io_uring 仅 linux 可用")

func newIOUring(entries uint32) (*ioUring, error) { return nil, errIOUringUnsupported }

// driveDirect 非 linux 不可达（newIOUring 恒失败）。
func driveDirect(r *ioUring, extCh <-chan []extent, decCh chan<- []extent,
	abort <-chan struct{}, fail func(error)) {
}
