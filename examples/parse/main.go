// 解析示例：同一条 raw 的三种解析方式。
// 这里用 RecordEncoder 造一条 raw；实际使用中 raw 来自 Get / GetBatch。
// 运行：go run ./examples/parse
package main

import (
	"flag"
	"fmt"

	"github.com/viktorxia/immkv"
)

// annotationSchema 是这份库实际构建出来的 schema：顺序和类型必须与构建时一致。
// 它决定位图每一位对应哪个字段、payload 按什么顺序切；raw 本身不自描述。
// 本次构建只建了 AF、AN、DB；CLNSIG 在字段表里但被列裁剪裁掉，所以不在 schema 里。
// 实际使用时用 immkv.ReadSchema(libraryDir) 取；这里为了演示，构造了一个实例。
var annotationSchema = immkv.Schema{
	{Name: "AF", Type: immkv.TypeFloat, Number: "A"},
	{Name: "AN", Type: immkv.TypeInteger, Number: "1"},
	{Name: "DB", Type: immkv.TypeFlag, Number: "0"},
}

// ---- 1. INFO 字符串 ----
// 需要文本输出时用 RenderRecord：它按 schema 顺序把字段拼成 INFO 列文本。
// 输出是 INFO 列（例如 AF=0.01;AN=1200），不是完整 VCF 行。
// 批量注释工具为了性能通常不走这里，而是一次 Values() 遍历直接拼输出缓冲。

// renderInfoText 把 raw 渲染成 INFO 文本。
func renderInfoText(schema immkv.Schema, recordBytes []byte) (string, error) {
	infoBytes, err := immkv.RenderRecord(schema, recordBytes)
	if err != nil {
		return "", err
	}
	return string(infoBytes), nil
}

// ---- 2. 普通结构体 ----
// 这里描述一种不区分“字段值存在，但等于零值”和“字段值不存在”的场景；
// 两者在结构体里都表现为零值，适合不需要区分缺失和零值的场景。

// PlainAnnotation 只保存值；它分不清“字段缺失”和“字段为零值”。
type PlainAnnotation struct {
	AF float64
	AN int64
	DB bool
}

// decodePlainAnnotation 用一次 Values() 遍历把 raw 填进普通结构体。
func decodePlainAnnotation(schema immkv.Schema, recordBytes []byte) (PlainAnnotation, error) {
	var decoded PlainAnnotation

	// recordBytes 是 Get/GetBatch 返回的二进制 record，本身不自描述。
	// NewRecord 校验它，并返回绑定 schema 的只读视图 Record；
	// 不复制 recordBytes，也不是转换成另一种格式。
	record, err := immkv.NewRecord(schema, recordBytes)
	if err != nil {
		return decoded, err
	}
	// Values() 只产出有值字段；缺失字段不会进入循环。
	for field, value := range record.Values() {
		switch field.Name {
		case "AF":
			decoded.AF = value.Float
		case "AN":
			decoded.AN = value.Int
		case "DB":
			decoded.DB = true // Flag 出现即为 true
		}
	}
	return decoded, nil
}

// ---- 3. 带有效位图的结构体 ----
// present 位图只能分两态：
//   位 = 1：这条 record 里该字段有值（值即使为 0 也算有值）
//   位 = 0：该字段在这条 record 里没有值，或该字段未参与构建
// 只看 present 位无法区分后两者；用构建字段清单（-lib 模式从 ReadSchema 派生）：
//   清单里有该字段 -> 字段无值
//   清单里没有     -> 字段未构建

// maxFields 是这个示例的 validityBitmap 容量（1024 位）。
// 这不是 immkv 的限制：immkv 的 record 位图按 schema 字段数动态增长。
// 字段超过 1024 时，把 validityBitmap 换成动态长度的 []uint64 即可。
const maxFields = 1024

// validityBitmap 是记录的有效位图：第 i 位表示第 i 个字段是否有值。
type validityBitmap [maxFields / 64]uint64

// fieldBitIndex 是这份存取实现（写这个 decode 函数的人）的字段表：
// 字段名 -> present 位号。
// 它不是 immkv 构建产物的一部分，也不是从 annotationSchema 推导出来的；
// 它和下面的 PresentAnnotation 一起由调用方定义，所以 decode 函数知道
// present 的顺序。
// 参与构建的字段名与 annotationSchema 里的名字一致；CLNSIG 只在字段表里，
// 不参与构建，所以不会出现在 annotationSchema 或 record.Values() 中。
// 这里故意让顺序和 annotationSchema 不同，说明值按 schema 名解析、
// present 位按这张表排：
//
//	annotationSchema：AF, AN, DB
//	fieldBitIndex：   AN, AF, DB, CLNSIG
var fieldBitIndex = map[string]int{
	"AN": 0, "AF": 1, "DB": 2, "CLNSIG": 3,
}

// builtFields 是这份构建实际保留的字段，也就是构建字段清单。
// 单库演示里它从 annotationSchema 派生；-lib 模式从 immkv.ReadSchema 派生。
// CLNSIG 只在候选字段表里，这次构建被列裁剪裁掉，所以任何记录里都没有值。
var builtFields = builtFieldsFromSchema(annotationSchema)

// candidateFieldTypes 是候选字段表里的类型，用于 -lib 模式启动时校验库 schema。
var candidateFieldTypes = map[string]immkv.FieldType{
	"AF": immkv.TypeFloat, "AN": immkv.TypeInteger,
	"DB": immkv.TypeFlag, "CLNSIG": immkv.TypeString,
}

func builtFieldsFromSchema(schema immkv.Schema) map[string]bool {
	out := make(map[string]bool, len(schema))
	for _, field := range schema {
		out[field.Name] = true
	}
	return out
}

// validateCandidateFields 校验库 schema 的每个字段都在候选字段表里且类型一致。
// 数组字段（Number 为 R/G/. 的数值字段）单独报错：本示例只演示标量解析，
// 数组要用 Record.Values 的 Elems。
// 真实使用中候选表由调用方维护；库新增字段时这里会报错，提示补表。
func validateCandidateFields(schema immkv.Schema) error {
	for _, field := range schema {
		if field.Type == immkv.TypeInteger || field.Type == immkv.TypeFloat {
			switch field.Number {
			case "R", "G", ".":
				return fmt.Errorf("库字段 %q 是数组字段（Number=%s）；本示例只演示标量，数组请遍历 Value.Elems",
					field.Name, field.Number)
			}
		}
		want, ok := candidateFieldTypes[field.Name]
		if !ok {
			return fmt.Errorf("库字段 %q 不在候选字段表里", field.Name)
		}
		if field.Type != want {
			return fmt.Errorf("库字段 %q 类型 %s，候选表期望 %s", field.Name, field.Type, want)
		}
	}
	return nil
}

// PresentAnnotation 保存值，并用有效位图记录每个字段是否有值。
// 位号由上面的 fieldBitIndex 决定，结构体里不存这张表。
type PresentAnnotation struct {
	present validityBitmap
	AF      float64
	AN      int64
}

// HasValue 是数据记录结构体的方法：这条 record 里该字段有值吗。
// 字段名不在字段表时 panic。
func (annotation *PresentAnnotation) HasValue(fieldName string) bool {
	bit, ok := fieldBitIndex[fieldName]
	if !ok {
		panic("unknown field: " + fieldName)
	}
	return annotation.present[bit/64]&(1<<(bit%64)) != 0
}

// fieldState 按三态判断：
//   - HasValue 为 true：字段有值；
//   - HasValue 为 false、构建清单里有该字段：字段无值；
//   - 构建清单里没有：字段未构建。
func fieldState(annotation PresentAnnotation, built map[string]bool, fieldName string) string {
	switch {
	case annotation.HasValue(fieldName):
		return "有值"
	case built[fieldName]:
		return "无值"
	default:
		return "未构建"
	}
}

// decodePresentAnnotation 用一次 Values() 遍历同时填值和有效位图。
func decodePresentAnnotation(schema immkv.Schema, recordBytes []byte) (PresentAnnotation, error) {
	var decoded PresentAnnotation
	record, err := immkv.NewRecord(schema, recordBytes)
	if err != nil {
		return decoded, err
	}
	// Values() 只产出有值字段；没产出的字段位保持 0。
	for field, value := range record.Values() {

		// 第一步：按 schema 字段名把值填进结构体。
		switch field.Name {
		case "AF":
			decoded.AF = value.Float
		case "AN":
			decoded.AN = value.Int
		}

		// 第二步：迁移只看字段名。Values() 给出的是有值字段的名字，
		// fieldBitIndex 把它翻译成本结构体 present 的位号；
		// annotationSchema 的顺序不参与。真正被写的是 decoded.present。
		bit, ok := fieldBitIndex[field.Name]
		if !ok {
			panic("field not in field table: " + field.Name)
		}
		decoded.present[bit/64] |= 1 << (bit % 64)
	}
	return decoded, nil
}

// encodeRecord 造一条演示 record；withAN=false 时 AN 缺失。
func encodeRecord(af float64, an int64, withAN bool) []byte {
	encoder := immkv.NewRecordEncoder(annotationSchema)
	encoder.Reset().Float("AF", af)
	if withAN {
		encoder.Int("AN", an)
	}
	recordBytes, err := encoder.Flag("DB").Encode()
	if err != nil {
		panic(err)
	}
	return append([]byte(nil), recordBytes...)
}

func main() {
	lib := flag.String("lib", "", "immkv library directory; empty runs the self-contained demo")
	flag.Parse()
	if *lib != "" {
		runLibraryDemo(*lib)
		return
	}
	runStandaloneDemo()
}

// runStandaloneDemo 不依赖外部库：用 RecordEncoder 造 record，演示三种解析方式。
func runStandaloneDemo() {
	// 1. 造一条 raw：AF=0.01、AN=1200、DB；CLNSIG 未参与构建。
	recordBytes := encodeRecord(0.01, 1200, true)

	// 2. 解析成 INFO 文本。
	infoText, err := renderInfoText(annotationSchema, recordBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("INFO:", infoText)

	// 3. 解析成普通结构体。
	plainAnnotation, err := decodePlainAnnotation(annotationSchema, recordBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("plain: AF=%v AN=%d DB=%v\n",
		plainAnnotation.AF, plainAnnotation.AN, plainAnnotation.DB)

	// 4. 解析成带有效位图的结构体。
	presentAnnotation, err := decodePresentAnnotation(annotationSchema, recordBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("present: AF=%v(%v) AN=%d(%v) DB=%v\n",
		presentAnnotation.AF, presentAnnotation.HasValue("AF"),
		presentAnnotation.AN, presentAnnotation.HasValue("AN"),
		presentAnnotation.HasValue("DB"))

	// present 位号顺序故意与 schema 顺序不同：值按 schema 名解析，位按字段表排。
	fmt.Printf("present bits: AN=%d AF=%d DB=%d CLNSIG=%d\n",
		fieldBitIndex["AN"], fieldBitIndex["AF"], fieldBitIndex["DB"], fieldBitIndex["CLNSIG"])

	// 5. 三态判断：有值 / 无值 / 未构建。
	missingBytes := encodeRecord(0.50, 0, false) // 这条 record 没有 AN
	missingAnnotation, err := decodePresentAnnotation(annotationSchema, missingBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("state AF:", fieldState(missingAnnotation, builtFields, "AF"))         // 有值
	fmt.Println("state AN:", fieldState(missingAnnotation, builtFields, "AN"))         // 无值
	fmt.Println("state CLNSIG:", fieldState(missingAnnotation, builtFields, "CLNSIG")) // 未构建
}

// runLibraryDemo 从真实库读取 schema 和第一条 record，再走同一套解析路径。
func runLibraryDemo(dir string) {
	schema, err := immkv.ReadSchema(dir)
	if err != nil {
		panic(err)
	}
	if err := validateCandidateFields(schema); err != nil {
		panic(err)
	}
	built := builtFieldsFromSchema(schema)

	db, err := immkv.OpenEager(dir, 0)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var key, recordBytes []byte
	if err := db.Scan(nil, nil, func(k, v []byte) bool {
		key = append([]byte(nil), k...)
		recordBytes = append([]byte(nil), v...)
		return false
	}); err != nil {
		panic(err)
	}
	if recordBytes == nil {
		fmt.Println("library has no records:", dir)
		return
	}

	fmt.Println("library:", dir)
	fmt.Print("schema:")
	for _, field := range schema {
		fmt.Printf(" %s(%s)", field.Name, field.Type)
	}
	fmt.Println()
	fmt.Printf("record: key=%x\n", key)

	infoText, err := renderInfoText(schema, recordBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("INFO:", infoText)

	present, err := decodePresentAnnotation(schema, recordBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("present: AF=%v(%v) AN=%d(%v) DB=%v\n",
		present.AF, present.HasValue("AF"),
		present.AN, present.HasValue("AN"),
		present.HasValue("DB"))
	fmt.Println("state AF:", fieldState(present, built, "AF"))
	fmt.Println("state AN:", fieldState(present, built, "AN"))
	fmt.Println("state CLNSIG:", fieldState(present, built, "CLNSIG"))
}
