// simple.go 是 immkv 的单库精确点查句柄：直接打开一对 .keys/.vals，
// 不做染色体分片与首字节路由。适用于无染色体/坐标的普通 KV，
// 例如 gene symbol -> 字符串；查询面只有 Get。

package immkv

import (
	"encoding/binary"
	"fmt"
)

// SimpleKVDB 是单库精确点查句柄：打开 basePath.keys / basePath.vals 这一对
// 文件，键可以是任意字节串。多 goroutine 并发安全；返回切片只保证到下一次
// 查询前有效。
type SimpleKVDB struct {
	s *shard
}

// OpenSimpleKV 打开 basePath.keys 与 basePath.vals 这一对文件。cacheMB 是
// 解压缓存预算（MB），0 表示不缓存。与 OpenEager 不同，它不做染色体分片，
// 也不按首字节路由。
func OpenSimpleKV(basePath string, cacheMB int) (*SimpleKVDB, error) {
	var cache *lruCache
	if cacheMB > 0 {
		cache = newLRU(cacheMB << 20)
	}
	s, err := openShard(basePath, cache)
	if err != nil {
		return nil, err
	}
	return &SimpleKVDB{s: s}, nil
}

// Get 精确点查；未命中返回 (nil, false, nil)，读取、解压或 record 形状不符
// 返回 (nil, false, err)。重复键返回排序后的第一条。库里的值是单 String
// 字段 record，这里解包后返回原始值字节。
func (d *SimpleKVDB) Get(key []byte) ([]byte, bool, error) {
	raw, ok, err := d.s.get(key)
	if err != nil || !ok {
		return nil, ok, err
	}
	return decodeSimpleValue(raw)
}

// decodeSimpleValue 解空 schema 库的 record：8B bitmap（只允许 bit0）
// + uvarint(len) + 原字节。record 形状不符时返回错误。
func decodeSimpleValue(raw []byte) ([]byte, bool, error) {
	if len(raw) < 8 {
		return nil, false, fmt.Errorf("immkv: 简单库 record 长度 %d < 8", len(raw))
	}
	if binary.LittleEndian.Uint64(raw[:8]) != 1 {
		return nil, false, fmt.Errorf("immkv: 简单库 record 位图不是单 String 字段")
	}
	length, n := binary.Uvarint(raw[8:])
	if n <= 0 {
		return nil, false, fmt.Errorf("immkv: 简单库 record 长度字段无效")
	}
	start := 8 + n
	if length != uint64(len(raw)-start) {
		return nil, false, fmt.Errorf("immkv: 简单库 record 长度 %d ≠ 剩余 %d", length, len(raw)-start)
	}
	return raw[start:], true, nil
}

// Close 关闭文件句柄与解码器。调用方保证不与查询并发。
func (d *SimpleKVDB) Close() error { return d.s.Close() }
