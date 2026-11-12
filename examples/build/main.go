// 构建示例：把一个小 VCF 构建成 immkv 库。
// 运行：go run ./examples/build
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/viktorxia/immkv"
)

// demoVCF 是构建输入：一条染色体、两个位置、三个变异。
const demoVCF = `##fileformat=VCFv4.2
##INFO=<ID=AF,Number=A,Type=Float,Description="Allele frequency">
##INFO=<ID=AN,Number=1,Type=Integer,Description="Allele number">
##INFO=<ID=CLNSIG,Number=1,Type=String,Description="Clinical significance">
##INFO=<ID=DB,Number=0,Type=Flag,Description="dbSNP membership">
#CHROM	POS	ID	REF	ALT	QUAL	FILTER	INFO
1	1000	.	A	G	.	.	AF=0.01;AN=1200;CLNSIG=Benign;DB
1	1000	.	A	T	.	.	AF=0.02;AN=1200;CLNSIG=Likely_benign
1	2000	.	C	T	.	.	AF=0.5;AN=800;CLNSIG=Pathogenic;DB
`

func main() {
	// 1. 建临时工作目录，存放 VCF 和构建出的库。
	workDir, err := os.MkdirTemp("", "immkv-build-")
	if err != nil {
		panic(err)
	}

	// 2. 把演示 VCF 写到工作目录。
	vcfPath := filepath.Join(workDir, "input.vcf")
	if err := os.WriteFile(vcfPath, []byte(demoVCF), 0o644); err != nil {
		panic(err)
	}

	// 3. 输出目录：构建成功后里面是 chrN.keys/.vals/.col/.manifest。
	//    目录不必为空，但不能已有这些 shard 文件；构建会覆盖 vcf.*.header。
	libraryDir := filepath.Join(workDir, "library")

	// 4. 输入 VCF 列表：这里只有一个文件；多个文件时每个文件只能含一条染色体。
	inputVCFPaths := []string{vcfPath}

	// 5. 字段选择：nil 或空表示存 ##INFO 里的全部字段；非空时只存列出的字段。
	//    这个切片的顺序只用于选择，schema 顺序始终按 ##INFO 声明顺序。
	var selectedFieldNames []string

	// 6. 字段名前缀：空表示 INFO 键原名落盘；非空时写成 prefix_键。
	fieldNamePrefix := ""

	// 7. 构建选项：Options{} 全是默认值——短值裸存阈值 64 字节，值帧按 64KB 打包。
	//    Schema 由 VCF 头生成；Threads/CacheMB/Columns 是读侧参数，构建不用。
	//    点查库可设 PackTarget: -1 改成每值独立成帧 + 训练字典，
	//    但需要真实规模语料，这种三行 VCF 训练不出字典。
	buildOptions := immkv.Options{}

	// 8. 构建，五个实参依次是 outdir、vcfs、fields、prefix、opts。
	if err := immkv.BuildFromVCFs(
		libraryDir,         // 输出库目录
		inputVCFPaths,      // 输入 VCF 列表
		selectedFieldNames, // nil = 存全部 INFO 字段
		fieldNamePrefix,    // "" = 不改字段名
		buildOptions,       // 构建选项
	); err != nil {
		panic(err)
	}

	// 9. 读回库级 schema，确认构建出的字段表。
	librarySchema, err := immkv.ReadSchema(libraryDir)
	if err != nil {
		panic(err)
	}

	// 10. 打印库目录和字段表。
	fmt.Println("library:", libraryDir)
	fmt.Print("schema:")
	for _, field := range librarySchema {
		fmt.Printf(" %s(%s)", field.Name, field.Type)
	}
	fmt.Println()
}
