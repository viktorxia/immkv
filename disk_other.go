//go:build !linux

package immkv

// ioProfile 是读路选用的介质画像。非 Linux 平台不探测，按 SSD 处理。
type ioProfile struct {
	rotational bool
}

func detectIOProfile(dir string) ioProfile { return ioProfile{} }
