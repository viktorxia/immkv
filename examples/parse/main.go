// 解析示例：同一条 raw 的三种解析方式。
// 这里用 RecordEncoder 造一条 raw；实际使用中 raw 来自 Get / GetBatch。
// 运行：go run ./examples/parse
package main

import (
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
func renderInfoText(recordBytes []byte) (string, error) {
	infoBytes, err := immkv.RenderRecord(annotationSchema, recordBytes)
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
func decodePlainAnnotation(recordBytes []byte) (PlainAnnotation, error) {
	var decoded PlainAnnotation

	// recordBytes 是 Get/GetBatch 返回的二进制 record，本身不自描述。
	// NewRecord 校验它，并返回绑定 schema 的只读视图 Record；
	// 不复制 recordBytes，也不是转换成另一种格式。
	record, err := immkv.NewRecord(annotationSchema, recordBytes)
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
// 只看 present 位无法区分后两者；用 IsBuiltField 查构建字段清单：
//   IsBuiltField = true  -> 字段无值
//   IsBuiltField = false -> 字段未构建

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
// CLNSIG 在字段表里，但这次构建被列裁剪裁掉，所以任何记录里都没有值。
var builtFields = map[string]bool{
	"AN": true, "AF": true, "DB": true,
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

// IsBuiltField 是数据集层的方法：这个字段参与这次构建了吗。
// 字段名不在字段表时 panic。
func (annotation *PresentAnnotation) IsBuiltField(fieldName string) bool {
	if _, ok := fieldBitIndex[fieldName]; !ok {
		panic("unknown field: " + fieldName)
	}
	return builtFields[fieldName]
}

// fieldState 按三态判断：
//   - HasValue 为 true：字段有值；
//   - HasValue 为 false、IsBuiltField 为 true：字段无值；
//   - IsBuiltField 为 false：字段未构建。
func fieldState(annotation PresentAnnotation, fieldName string) string {
	switch {
	case annotation.HasValue(fieldName):
		return "有值"
	case annotation.IsBuiltField(fieldName):
		return "无值"
	default:
		return "未构建"
	}
}

// decodePresentAnnotation 用一次 Values() 遍历同时填值和有效位图。
func decodePresentAnnotation(recordBytes []byte) (PresentAnnotation, error) {
	var decoded PresentAnnotation
	record, err := immkv.NewRecord(annotationSchema, recordBytes)
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
	// 1. 造一条 raw：AF=0.01、AN=1200、DB；CLNSIG 未参与构建。
	recordBytes := encodeRecord(0.01, 1200, true)

	// 2. 解析成 INFO 文本。
	infoText, err := renderInfoText(recordBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("INFO:", infoText)

	// 3. 解析成普通结构体。
	plainAnnotation, err := decodePlainAnnotation(recordBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("plain: AF=%v AN=%d DB=%v\n",
		plainAnnotation.AF, plainAnnotation.AN, plainAnnotation.DB)

	// 4. 解析成带有效位图的结构体。
	presentAnnotation, err := decodePresentAnnotation(recordBytes)
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
	missingAnnotation, err := decodePresentAnnotation(missingBytes)
	if err != nil {
		panic(err)
	}
	fmt.Println("state AF:", fieldState(missingAnnotation, "AF"))         // 有值
	fmt.Println("state AN:", fieldState(missingAnnotation, "AN"))         // 无值
	fmt.Println("state CLNSIG:", fieldState(missingAnnotation, "CLNSIG")) // 未构建
}
