// 解析示例：同一条 raw 的三种解析方式。
// 这里用 RecordEncoder 造一条 raw；实际使用中 raw 来自 Get / GetBatch。
// 运行：go run ./examples/parse
package main

import (
	"fmt"

	"github.com/viktorxia/immkv"
)

// annotationSchema 是 raw 的字段表：顺序和类型必须与构建时一致。
// 它决定位图每一位对应哪个字段、payload 按什么顺序切；raw 本身不自描述。
// 实际使用时用 immkv.ReadSchema(libraryDir) 取；这里为了演示，构造了一个实例。
var annotationSchema = immkv.Schema{
	{Name: "AF", Type: immkv.TypeFloat, Number: "A"},
	{Name: "AN", Type: immkv.TypeInteger, Number: "1"},
	{Name: "CLNSIG", Type: immkv.TypeString, Number: "1"},
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
	AF      float64
	AN      int64
	ClinSig string
	DB      bool
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
		case "CLNSIG":
			decoded.ClinSig = string(value.Bytes)
		case "DB":
			decoded.DB = true // Flag 出现即为 true
		}
	}
	return decoded, nil
}

// ---- 3. 带有效位图的结构体 ----
// present 位图能区分两态：
//   位 = 1：字段在 schema 里，且这条 record 有值（值即使为 0 也算有值）
//   位 = 0：字段在 schema 里，但这条 record 没给值
// “库内没有这个字段”（schema 里没有）需要另查 fieldBitIndex；
// 只看 present 位时，它和“记录没给值”一样都是 0，无法区分。

// maxFields 是这个示例的 validityBitmap 容量（1024 位）。
// 这不是 immkv 的限制：immkv 的 record 位图按 schema 字段数动态增长。
// 字段超过 1024 时，把 validityBitmap 换成动态长度的 []uint64 即可。
const maxFields = 1024

// validityBitmap 是记录的有效位图：第 i 位表示第 i 个字段是否有值。
type validityBitmap [maxFields / 64]uint64

// PresentAnnotation 保存值，并用有效位图记录每个字段是否有值。
type PresentAnnotation struct {
	// fieldBitIndex 是 schema 字段名到 present 位号的映射。
	// 键必须是 schema 里的名字，不是 Go 结构体字段名：
	//   map[string]int{"AF": 0, "AN": 1, "CLNSIG": 2}
	// 例如结构体字段叫 ClinSig，键仍要写 "CLNSIG"。
	fieldBitIndex map[string]int
	present       validityBitmap
	AF            float64
	AN            int64
	ClinSig       string
}

// IsFieldBuilt 查字段是否在这份库的 schema 里。
func (annotation *PresentAnnotation) IsFieldBuilt(fieldName string) bool {
	_, ok := annotation.fieldBitIndex[fieldName]
	return ok
}

// IsFieldPresent 查字段在这条 raw 里是否有值。
// 字段不在 fieldBitIndex 里时返回 false；要区分“库内没有这个字段”，
// 先用 IsFieldBuilt 查。
func (annotation *PresentAnnotation) IsFieldPresent(fieldName string) bool {
	bit, ok := annotation.fieldBitIndex[fieldName]
	return ok && annotation.present[bit/64]&(1<<(bit%64)) != 0
}

// decodePresentAnnotation 用一次 Values() 遍历同时填值和有效位图。
func decodePresentAnnotation(fieldBitIndex map[string]int, recordBytes []byte) (PresentAnnotation, error) {
	decoded := PresentAnnotation{fieldBitIndex: fieldBitIndex}
	record, err := immkv.NewRecord(annotationSchema, recordBytes)
	if err != nil {
		return decoded, err
	}
	// Values() 只产出有值字段；没产出的字段位保持 0。
	for field, value := range record.Values() {

		// 第一步：按字段名把值填进结构体。
		switch field.Name {
		case "AF":
			decoded.AF = value.Float
		case "AN":
			decoded.AN = value.Int
		case "CLNSIG":
			decoded.ClinSig = string(value.Bytes)
		}

		// 第二步：字段有值，就置 present 位。
		if bit, ok := fieldBitIndex[field.Name]; ok {
			decoded.present[bit/64] |= 1 << (bit % 64)
		}
	}
	return decoded, nil
}

// buildFieldBitIndex 按 schema 顺序生成“字段名 -> 位号”表。
func buildFieldBitIndex() map[string]int {
	if len(annotationSchema) > maxFields {
		panic("schema has more than 1024 fields")
	}
	fieldBitIndex := make(map[string]int, len(annotationSchema))
	for fieldIndex, field := range annotationSchema {
		fieldBitIndex[field.Name] = fieldIndex
	}
	return fieldBitIndex
}

func main() {
	// 1. 造一条 raw：AF=0.01、AN=1200、CLNSIG=Benign、DB。
	encoder := immkv.NewRecordEncoder(annotationSchema)
	recordBytes, err := encoder.Reset().
		Float("AF", 0.01).
		Int("AN", 1200).
		Bytes("CLNSIG", []byte("Benign")).
		Flag("DB").
		Encode()
	if err != nil {
		panic(err)
	}

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
	fmt.Printf("plain: AF=%v AN=%d CLNSIG=%q DB=%v\n",
		plainAnnotation.AF, plainAnnotation.AN, plainAnnotation.ClinSig, plainAnnotation.DB)

	// 4. 解析成带有效位图的结构体。
	presentAnnotation, err := decodePresentAnnotation(buildFieldBitIndex(), recordBytes)
	if err != nil {
		panic(err)
	}
	fmt.Printf("present: AF=%v(%v) AN=%d(%v) CLNSIG=%q(%v) DB=%v\n",
		presentAnnotation.AF, presentAnnotation.IsFieldPresent("AF"),
		presentAnnotation.AN, presentAnnotation.IsFieldPresent("AN"),
		presentAnnotation.ClinSig, presentAnnotation.IsFieldPresent("CLNSIG"),
		presentAnnotation.IsFieldPresent("DB"))
}
