//go:build linux

package immkv

import (
	"os"

	"golang.org/x/sys/unix"
)

// fadviseDontNeed 对文件的一段区间做 POSIX_FADV_DONTNEED（丢弃干净页缓存）。
// 流式路径的值只读一次，页缓存留着只会顶爆 cgroup 限额、拖慢后续读。
// POSIX_FADV_DONTNEED = 4，x/sys/unix 在 linux 上未导出该常量。
func fadviseDontNeed(f *os.File, off, n int64) {
	_ = unix.Fadvise(int(f.Fd()), off, n, 4)
}
