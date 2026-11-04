//go:build cgo

package immkv

import (
	gz "github.com/valyala/gozstd"
)

// cgoDictCompressor 用 libzstd（经 valyala/gozstd，vendored C 源码）的 CDict
// 做逐值压缩。产出的帧是标准 zstd 格式，查询端 klauspost/gozstd 带同一字典
// 都能解（已实测互通）。档位 9：实测 chr1 11.1%、chr7 12.9%、小染色体 ~8.7%
// （每值 ratio），96-137MB/s/核；3 档全量投影 598GB（bgzip 的 121%，越线），
// 19 档仅 3.25MB/s 不可用，9 档全量投影 426GB（最差兜底 480GB）< bgzip 495GB。
// CGO_ENABLED=0 构建的查询二进制不含本路径。
type cgoDictCompressor struct {
	cd    *gz.CDict
	dst   []byte
	level int
}

func newCGODictCompressor(dict []byte) (dictCompressor, error) {
	cd, err := gz.NewCDictLevel(dict, 9)
	if err != nil {
		return nil, err
	}
	return &cgoDictCompressor{cd: cd, dst: make([]byte, 0, 64<<10), level: 9}, nil
}

func (c *cgoDictCompressor) encode(src []byte) []byte {
	return gz.CompressDict(c.dst[:0], src, c.cd)
}

// cgoDecoder 是 frameDecoder 的 libzstd 实现：DDict 常驻（可被多 goroutine
// 并发使用），DCtx 走 gozstd 内部的 sync.Pool，解码调用本身并发安全。
type cgoDecoder struct {
	dd *gz.DDict // 无字典（v2 打包库）为 nil：gozstd 的 nil DDict 即裸解
}

// newCGODictDecoder 建带字典的 CGO 解码器；dict 为空返回 error（语义上没有
// 「空字典」，无字典场景用 newCGODecoder）。
func newCGODictDecoder(dict []byte) (frameDecoder, error) {
	dd, err := gz.NewDDict(dict)
	if err != nil {
		return nil, err
	}
	return &cgoDecoder{dd: dd}, nil
}

// newCGODecoder 建无字典的 CGO 解码器（v2 打包帧库：帧自带全部上下文）。
func newCGODecoder() (frameDecoder, error) {
	return &cgoDecoder{}, nil
}

func (c *cgoDecoder) decode(src, dst []byte) ([]byte, error) {
	return gz.DecompressDict(dst, src, c.dd)
}

func (c *cgoDecoder) Close() error {
	if c.dd != nil {
		c.dd.Release()
		c.dd = nil
	}
	return nil
}
