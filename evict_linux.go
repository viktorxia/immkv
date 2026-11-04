//go:build linux

package immkv

import "golang.org/x/sys/unix"

// evictPageCache 用 posix_fadvise(DONTNEED) 驱逐本 shard 两个文件的页缓存。
// 只给基准测试用；POSIX_FADV_DONTNEED = 4，x/sys/unix 在 linux 上未导出该常量。
func (s *shard) evictPageCache() error {
	if err := unix.Fadvise(int(s.kf.Fd()), 0, 0, 4); err != nil {
		return err
	}
	return unix.Fadvise(int(s.vf.Fd()), 0, 0, 4)
}
