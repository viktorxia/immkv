// simple.go 是 immkv 的单库精确点查句柄：直接打开一对 .keys/.vals，
// 不做染色体分片与首字节路由。适用于无染色体/坐标的普通 KV，
// 例如 gene symbol -> 字符串；查询面只有 Get。

package immkv

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

// Get 精确点查；未命中 ok=false。重复键返回排序后的第一条。
func (d *SimpleKVDB) Get(key []byte) ([]byte, bool) {
	return d.s.get(key)
}

// Close 关闭文件句柄与解码器。调用方保证不与查询并发。
func (d *SimpleKVDB) Close() error { return d.s.Close() }
