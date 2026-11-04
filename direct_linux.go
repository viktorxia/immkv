//go:build linux

package immkv

import (
	"os"
	"syscall"
)

// openDirect 打开 O_DIRECT 只读句柄；内核或文件系统不支持时返回错误，
// 调用方退回缓冲读。
func openDirect(path string) (*os.File, error) {
	if os.Getenv("IMMKV_NO_DIRECT") != "" {
		return nil, os.ErrInvalid
	}
	return os.OpenFile(path, os.O_RDONLY|syscall.O_DIRECT, 0)
}
