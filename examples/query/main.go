// 查询示例：打开库，做点查、Scan、批量查；只返回 raw，不解析。
// 运行：go run ./examples/query -lib <库目录>
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/viktorxia/immkv"
)

func main() {
	// 1. 从命令行取库目录。
	libraryDir := flag.String("lib", "", "immkv library directory")
	flag.Parse()
	if *libraryDir == "" {
		fmt.Fprintln(os.Stderr, "usage: query -lib <dir>")
		os.Exit(2)
	}

	// 2. 把 "chr1" 转成键首字节。
	chromID, ok := immkv.ChromID("chr1")
	if !ok {
		panic("unknown chromosome")
	}

	// 3. 打开点查句柄；256 是解压缓存预算（MB）。
	eagerDB, err := immkv.OpenEager(*libraryDir, 256)
	if err != nil {
		panic(err)
	}
	defer eagerDB.Close()

	// 4. 点查：键 = 染色体 + 位置 + ref + alt。
	// recordBytes 是二进制 record，不是文本，也不是完整 VCF 行：
	// 开头是 presence 位图（每个字段一位，1 = 这条记录有值），
	// 后面是按 schema 顺序编码的字段值。
	// 它不自描述，必须配 schema 才能读；解析方式见 examples/parse/main.go。
	queryKey := immkv.VarKey(chromID, 1000, "A", "G")
	if recordBytes, ok := eagerDB.Get(queryKey); ok {
		fmt.Printf("Get: hit, %d bytes\n", len(recordBytes))
	} else {
		fmt.Println("Get: miss")
	}

	// 5. 区间扫描：[lo, hi)，hi 独占；回调返回 false 提前结束。
	eagerDB.Scan(immkv.PosKey(chromID, 0), immkv.PosKey(chromID, 2001), func(key, recordBytes []byte) bool {
		fmt.Printf("Scan: key=%x bytes=%d\n", key, len(recordBytes))
		return true
	})

	// 6. 打开批量句柄。Threads 决定并行解压的 goroutine 数；
	//    io_uring 路径下在途读深由队列控制，与 Threads 无关；
	//    只有 io_uring/O_DIRECT 不可用的兜底路径，Threads 才兼作 pread 读线程数。
	//    机械盘上单磁头，Threads 不起作用。
	streamDB, err := immkv.OpenStreaming(*libraryDir, immkv.Options{Threads: 4})
	if err != nil {
		panic(err)
	}
	defer streamDB.Close()

	// 7. 批量查的键必须升序（非降序，允许重复）；顺序错会返回 immkv.ErrUnsorted。
	queryKeys := [][]byte{
		immkv.VarKey(chromID, 1000, "A", "G"),
		immkv.VarKey(chromID, 1000, "A", "T"),
		immkv.VarKey(chromID, 2000, "C", "T"),
	}
	rawValues, err := streamDB.GetBatch(queryKeys)
	if err != nil {
		panic(err)
	}

	// 8. 结果与输入一一对应；nil 表示未命中。
	for i, recordBytes := range rawValues {
		if recordBytes == nil {
			fmt.Printf("GetBatch: miss %x\n", queryKeys[i])
			continue
		}
		fmt.Printf("GetBatch: hit key=%x bytes=%d\n", queryKeys[i], len(recordBytes))
	}
}
