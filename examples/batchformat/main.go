// 批量格式化示例：一次 Values() 遍历，把字段值直接写进复用缓冲。
//
// 和 examples/parse 的 RenderRecord 方式不同，这里不生成中间 INFO 字符串，
// 也不按字段名逐个调用 getter：
//  1. 每条 record 只遍历一次 Values()。
//  2. 用“字段名 -> 输出下标”的表过滤请求字段，O(1)。
//  3. 值的文本直接 append 到复用的 rendered 缓冲，记下 start/end。
//  4. 按请求字段顺序从 rendered 里 slice，拼成最终 INFO 文本。
//
// 实际使用中 recordBytes 来自 GetBatch；这里用 RecordEncoder 造数据。
// 运行：go run ./examples/batchformat
package main

import (
	"fmt"
	"strconv"

	"github.com/viktorxia/immkv"
)

// annotationSchema 与要处理的库一致。
var annotationSchema = immkv.Schema{
	{Name: "AF", Type: immkv.TypeFloat, Number: "A"},
	{Name: "AN", Type: immkv.TypeInteger, Number: "1"},
	{Name: "CLNSIG", Type: immkv.TypeString, Number: "1"},
	{Name: "DB", Type: immkv.TypeFlag, Number: "0"},
}

// requestedFields 是调用方要输出的字段；顺序就是最终 INFO 文本的顺序。
var requestedFields = []string{"AF", "AN"}

// fieldOutputIndex 是“字段索引表”：schema 字段名 -> requestedFields 下标。
// 不在表里的字段在遍历时直接跳过。
var fieldOutputIndex = buildFieldOutputIndex(requestedFields)

func buildFieldOutputIndex(fieldNames []string) map[string]int {
	index := make(map[string]int, len(fieldNames))
	for outputIndex, fieldName := range fieldNames {
		index[fieldName] = outputIndex
	}
	return index
}

// valueSpan 是一个字段值在 rendered 缓冲里的起止位置。
type valueSpan struct {
	start int
	end   int
}

// formatter 复用三块内存，避免每行分配：
//
//	rendered：这一行所有值的文本
//	spans：每个请求字段在 rendered 里的位置
//	output：最终拼好的 INFO 文本
type formatter struct {
	rendered []byte
	spans    []valueSpan
	output   []byte
}

func newFormatter() *formatter {
	return &formatter{spans: make([]valueSpan, len(requestedFields))}
}

// format 把一条 raw record 格式化成 "AF=...;AN=..."。
// 返回的 []byte 在下一次 format 调用前有效。
func (f *formatter) format(recordBytes []byte) ([]byte, error) {
	// 1. 清空复用缓冲；spans.start = -1 表示该字段没值。
	f.rendered = f.rendered[:0]
	f.output = f.output[:0]
	for i := range f.spans {
		f.spans[i] = valueSpan{start: -1, end: -1}
	}

	// 2. 校验 raw 并绑定 schema；raw 本身不自描述。
	record, err := immkv.NewRecord(annotationSchema, recordBytes)
	if err != nil {
		return nil, err
	}

	// 3. 一次 Values() 遍历所有有值字段；只把请求字段写进缓冲。
	for field, value := range record.Values() {
		outputIndex, requested := fieldOutputIndex[field.Name]
		if !requested {
			continue
		}
		start := len(f.rendered)
		f.rendered = appendFieldValue(f.rendered, field, value)
		f.spans[outputIndex] = valueSpan{start: start, end: len(f.rendered)}
	}

	// 4. 按请求字段顺序从 rendered 里取片段，拼成 INFO 文本。
	for outputIndex, fieldName := range requestedFields {
		span := f.spans[outputIndex]
		if span.start < 0 {
			continue
		}
		if len(f.output) > 0 {
			f.output = append(f.output, ';')
		}
		f.output = append(f.output, fieldName...)
		f.output = append(f.output, '=')
		f.output = append(f.output, f.rendered[span.start:span.end]...)
	}
	return f.output, nil
}

// appendFieldValue 把一个标量值写成 INFO 文本，写进 dst。
// 这里只实现标量；数组字段要遍历 value.Elems，Flag 要单独决定是否写 "="。
func appendFieldValue(dst []byte, field immkv.Field, value immkv.Value) []byte {
	switch field.Type {
	case immkv.TypeInteger:
		return strconv.AppendInt(dst, value.Int, 10)
	case immkv.TypeFloat:
		return strconv.AppendFloat(dst, value.Float, 'g', -1, 64)
	case immkv.TypeString:
		return append(dst, value.Bytes...)
	default:
		return dst
	}
}

// encodeRecord 造一条演示 record；实际使用中 recordBytes 来自 GetBatch。
func encodeRecord(af float64, an int64) []byte {
	encoder := immkv.NewRecordEncoder(annotationSchema)
	recordBytes, err := encoder.Reset().
		Float("AF", af).
		Int("AN", an).
		Bytes("CLNSIG", []byte("Benign")).
		Flag("DB").
		Encode()
	if err != nil {
		panic(err)
	}
	return append([]byte(nil), recordBytes...)
}

func main() {
	// 两条演示 record。
	recordBytesList := [][]byte{
		encodeRecord(0.01, 1200),
		encodeRecord(0.50, 800),
	}

	// 复用同一个 formatter 处理整批 record。
	formatter := newFormatter()
	for _, recordBytes := range recordBytesList {
		infoText, err := formatter.format(recordBytes)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(infoText))
	}
}
