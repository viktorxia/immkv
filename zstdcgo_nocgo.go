//go:build !cgo

package immkv

import "errors"

// newCGODictCompressor 在纯 Go 构建下不可用；调用方回退 klauspost 字典编码器。
func newCGODictCompressor(dict []byte) (dictCompressor, error) {
	return nil, errors.New("immkv: CGO 不可用")
}

// newCGODictDecoder / newCGODecoder 在纯 Go 构建下不可用；调用方回退 klauspost 解码器。
func newCGODictDecoder(dict []byte) (frameDecoder, error) {
	return nil, errors.New("immkv: CGO 不可用")
}

func newCGODecoder() (frameDecoder, error) {
	return nil, errors.New("immkv: CGO 不可用")
}
