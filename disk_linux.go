//go:build linux

package immkv

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"golang.org/x/sys/unix"
)

// ioProfile 是读路选用的介质画像。
type ioProfile struct {
	rotational bool // 目录所在块设备是机械盘（或含机械 slave）
}

// detectIOProfile 判定目录所在文件系统的块设备是否机械盘。查不到时按 SSD
// 处理（列存路径仍由读量估算决定，不会因为判错而读错数据）。
func detectIOProfile(dir string) ioProfile {
	var st syscall.Stat_t
	if err := syscall.Stat(dir, &st); err != nil {
		return ioProfile{}
	}
	dev := uint64(st.Dev)
	name := fmt.Sprintf("%d:%d", unix.Major(dev), unix.Minor(dev))
	return ioProfile{rotational: blockRotational("/sys/dev/block/" + name)}
}

// blockRotational 读 /sys/dev/block/<dev> 的 rotational；分区向上找父设备，
// dm/md 等虚拟设备看 slaves（任一机械即机械）。
func blockRotational(path string) bool {
	p, err := filepath.EvalSymlinks(path)
	if err != nil {
		return false
	}
	for _, q := range []string{
		filepath.Join(p, "queue", "rotational"),
		filepath.Join(p, "..", "queue", "rotational"),
	} {
		if b, err := os.ReadFile(q); err == nil {
			return strings.TrimSpace(string(b)) == "1"
		}
	}
	ents, err := os.ReadDir(filepath.Join(p, "slaves"))
	if err != nil {
		return false
	}
	for _, e := range ents {
		if blockRotational(filepath.Join(p, "slaves", e.Name())) {
			return true
		}
	}
	return false
}
