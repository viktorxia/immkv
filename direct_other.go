//go:build !linux

package immkv

import "os"

// openDirect 非 linux 平台不可用。
func openDirect(path string) (*os.File, error) {
	return nil, os.ErrInvalid
}
