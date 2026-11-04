//go:build !linux

package immkv

import "os"

// fadviseDontNeed 非 linux 平台为空操作。
func fadviseDontNeed(f *os.File, off, n int64) {}
