// record.go 是 v003 的 record 编解码：record = bitmap + typed payload。
// 读侧公开 Record 视图，按字段名取值；写侧公开 RecordEncoder。两边共用
// 同一套编码，RenderRecord 只负责把 record 规范化为 INFO 文本。
package immkv

import (
	"encoding/binary"
	"fmt"
	"iter"
	"math"
	"strconv"
)

// Value 是 Record.Values 回调的一项。Type 为 0 表示数组里的缺失元素；
// 数组字段用 Elems，标量字段用 Int/Float/Bytes。Bytes 是 raw record 的
// 视图，不复制，有效期由调用方持有的 record 字节决定。
type Value struct {
	Type  FieldType
	Int   int64
	Float float64
	Bytes []byte
	Elems []Value
}

// Record 是 raw record 的只读视图。NewRecord 校验结构；校验通过后，
// HasValue/Int/Float/Bytes 与 Values 不再返回错误。
type Record struct {
	schema Schema
	raw    []byte
}

// NewRecord 校验 raw 是否是 schema 的一条完整 record。raw 不被复制；
// 调用方要保证 raw 在 Record 使用期间有效。
func NewRecord(s Schema, raw []byte) (Record, error) {
	if err := validateRecord(s, raw); err != nil {
		return Record{}, err
	}
	return Record{schema: s, raw: raw}, nil
}

// HasValue 表示这条 record 里该字段是否有值。字段名不在 schema 时 panic。
func (r Record) HasValue(name string) bool {
	i, _ := r.field(name)
	return r.bit(i)
}

// Int 读 Integer 标量。字段缺失返回 0；调用方应先问 HasValue。
// 字段不是 Integer 或属于数组时 panic。
func (r Record) Int(name string) int64 {
	i, f := r.field(name)
	if f.Type != TypeInteger {
		panic("immkv: field " + name + " is not Integer")
	}
	if f.isArray() {
		panic("immkv: field " + name + " is an array; use Values()")
	}
	if !r.bit(i) {
		return 0
	}
	return r.valueAt(i, f).Int
}

// Float 读 Float 标量。字段缺失返回 0；调用方应先问 HasValue。
func (r Record) Float(name string) float64 {
	i, f := r.field(name)
	if f.Type != TypeFloat {
		panic("immkv: field " + name + " is not Float")
	}
	if f.isArray() {
		panic("immkv: field " + name + " is an array; use Values()")
	}
	if !r.bit(i) {
		return 0
	}
	return r.valueAt(i, f).Float
}

// Bytes 读 String 标量的原字节。字段缺失返回 nil；返回的是 raw record
// 的视图，不复制。
func (r Record) Bytes(name string) []byte {
	i, f := r.field(name)
	if f.Type != TypeString {
		panic("immkv: field " + name + " is not String")
	}
	if !r.bit(i) {
		return nil
	}
	return r.valueAt(i, f).Bytes
}

// Values 按 schema 顺序产出有值字段，字段缺失不产出。数组字段的 Value
// 带 Elems；Elems 里的零 Value 表示缺失元素。标量路径不分配。
func (r Record) Values() iter.Seq2[Field, Value] {
	return func(yield func(Field, Value) bool) {
		p := r.raw[recordBitmapLen(len(r.schema)):]
		for i := range r.schema {
			if !r.bit(i) {
				continue
			}
			f := r.schema[i]
			v, rest, err := decodeValue(f, p, true)
			if err != nil {
				panic("immkv: record 已损坏: " + err.Error())
			}
			p = rest
			if !yield(f, v) {
				return
			}
		}
	}
}

func (r Record) field(name string) (int, Field) {
	i := r.schema.Index(name)
	if i < 0 {
		panic("immkv: unknown field " + name)
	}
	return i, r.schema[i]
}

func (r Record) bit(i int) bool {
	if i < 0 || i >= len(r.schema) {
		panic("immkv: field index out of range")
	}
	return bitAt(r.raw, i)
}

// payloadAt 跳到第 i 个字段的 payload 起点。前面有值字段按编码跳过；
// 数组走 skipValue，不分配 Elems。
func (r Record) payloadAt(i int) ([]byte, error) {
	p := r.raw[recordBitmapLen(len(r.schema)):]
	for j := 0; j < i; j++ {
		if !r.bit(j) {
			continue
		}
		var err error
		p, err = skipValue(r.schema[j], p)
		if err != nil {
			return nil, fmt.Errorf("字段 %q: %w", r.schema[j].Name, err)
		}
	}
	return p, nil
}

func (r Record) valueAt(i int, f Field) Value {
	p, err := r.payloadAt(i)
	if err != nil {
		panic("immkv: record 已损坏: " + err.Error())
	}
	v, _, err := decodeValue(f, p, true)
	if err != nil {
		panic("immkv: record 已损坏: " + err.Error())
	}
	return v
}

// RenderRecord 把 record 规范化为 INFO 文本：字段顺序 = schema 顺序，
// Flag 写 name，数组逗号连接，缺失元素写 '.'，全缺写 '.'。
func RenderRecord(s Schema, raw []byte) ([]byte, error) {
	rec, err := NewRecord(s, raw)
	if err != nil {
		return nil, err
	}
	var out []byte
	first := true
	for f, v := range rec.Values() {
		if !first {
			out = append(out, ';')
		}
		first = false
		out = append(out, f.Name...)
		switch f.Type {
		case TypeFlag:
			continue
		case TypeInteger, TypeFloat, TypeString:
			out = append(out, '=')
			if f.isArray() {
				out = appendArray(out, v)
			} else {
				out = appendScalar(out, f.Type, v)
			}
		}
	}
	if len(out) == 0 {
		out = append(out, '.')
	}
	return out, nil
}

func appendScalar(dst []byte, t FieldType, v Value) []byte {
	switch t {
	case TypeInteger:
		return strconv.AppendInt(dst, v.Int, 10)
	case TypeFloat:
		return strconv.AppendFloat(dst, v.Float, 'g', -1, 64)
	case TypeString:
		return append(dst, v.Bytes...)
	default:
		return dst
	}
}

func appendArray(dst []byte, v Value) []byte {
	if len(v.Elems) == 0 {
		return append(dst, '.')
	}
	for i := range v.Elems {
		if i > 0 {
			dst = append(dst, ',')
		}
		e := v.Elems[i]
		if e.Type == 0 {
			dst = append(dst, '.')
			continue
		}
		dst = appendScalar(dst, e.Type, e)
	}
	return dst
}

// RecordEncoder 按字段名收集一条 record 的值，Encode 按 schema 顺序落成
// bitmap + payload。Encode 返回的 buf 归 encoder 所有，下一次 Encode/Reset
// 会覆盖；要保留就拷贝。
type RecordEncoder struct {
	schema    Schema
	index     map[string]int
	schemaErr error

	set     []bool
	vals    []encValue
	bitmap  []uint64
	buf     []byte
	scratch []byte
	err     error
}

type encValue struct {
	i     int64
	f     float64
	b     []byte
	elems []Value
}

// NewRecordEncoder 为一份 schema 建编码器。schema 的校验错误延迟到 Encode
// 返回；字段名在构造时解析一次。
func NewRecordEncoder(s Schema) *RecordEncoder {
	e := &RecordEncoder{
		schema:    s,
		index:     make(map[string]int, len(s)),
		schemaErr: s.Validate(),
		set:       make([]bool, len(s)),
		vals:      make([]encValue, len(s)),
		bitmap:    make([]uint64, (len(s)+63)/64),
	}
	for i := range s {
		e.index[s[i].Name] = i
	}
	return e
}

// Reset 清空上一条 record 的值，复用 buffer 与 bitmap。
func (e *RecordEncoder) Reset() *RecordEncoder {
	for i := range e.set {
		e.set[i] = false
		e.vals[i] = encValue{}
	}
	e.buf = e.buf[:0]
	e.err = nil
	return e
}

func (e *RecordEncoder) fail(format string, args ...any) {
	if e.err == nil {
		e.err = fmt.Errorf(format, args...)
	}
}

func (e *RecordEncoder) field(name string, want FieldType, array bool) (int, bool) {
	i, ok := e.index[name]
	if !ok {
		e.fail("immkv: unknown field %q", name)
		return 0, false
	}
	f := e.schema[i]
	if f.Type != want {
		e.fail("immkv: field %q is %s, not %s", name, f.Type, want)
		return 0, false
	}
	if f.isArray() != array {
		e.fail("immkv: field %q array=%v, not %v", name, f.isArray(), array)
		return 0, false
	}
	return i, true
}

// Int 设置一个 Integer 标量字段。
func (e *RecordEncoder) Int(name string, v int64) *RecordEncoder {
	i, ok := e.field(name, TypeInteger, false)
	if ok {
		e.set[i] = true
		e.vals[i] = encValue{i: v}
	}
	return e
}

// Float 设置一个 Float 标量字段；NaN/Inf 在 Encode 时返回错误。
func (e *RecordEncoder) Float(name string, v float64) *RecordEncoder {
	i, ok := e.field(name, TypeFloat, false)
	if !ok {
		return e
	}
	if math.IsNaN(v) || math.IsInf(v, 0) {
		e.fail("immkv: field %q float is NaN or Inf", name)
		return e
	}
	e.set[i] = true
	e.vals[i] = encValue{f: v}
	return e
}

// Bytes 设置一个 String 字段；Encode 时复制进输出 buffer。
func (e *RecordEncoder) Bytes(name string, v []byte) *RecordEncoder {
	i, ok := e.field(name, TypeString, false)
	if ok {
		e.set[i] = true
		e.vals[i] = encValue{b: v}
	}
	return e
}

// Flag 置位一个 Flag 字段。
func (e *RecordEncoder) Flag(name string) *RecordEncoder {
	i, ok := e.field(name, TypeFlag, false)
	if ok {
		e.set[i] = true
	}
	return e
}

// Values 设置一个 Number 为 R、G、. 的数值数组字段。elems 里的零 Value
// 表示缺失元素；缺整个字段就不要调用。
func (e *RecordEncoder) Values(name string, elems []Value) *RecordEncoder {
	i, ok := e.index[name]
	if !ok {
		e.fail("immkv: unknown field %q", name)
		return e
	}
	f := e.schema[i]
	if !f.isArray() {
		e.fail("immkv: field %q is not an array", name)
		return e
	}
	for k := range elems {
		el := &elems[k]
		if el.Type == 0 {
			if len(el.Elems) != 0 {
				e.fail("immkv: field %q element %d is missing but has Elems", name, k)
				return e
			}
			continue
		}
		if el.Type != f.Type {
			e.fail("immkv: field %q element %d is %s, not %s", name, k, el.Type, f.Type)
			return e
		}
		if len(el.Elems) != 0 {
			e.fail("immkv: field %q element %d has nested Elems", name, k)
			return e
		}
		if f.Type == TypeFloat && (math.IsNaN(el.Float) || math.IsInf(el.Float, 0)) {
			e.fail("immkv: field %q element %d float is NaN or Inf", name, k)
			return e
		}
	}
	e.set[i] = true
	e.vals[i] = encValue{elems: elems}
	return e
}

// Encode 按 schema 顺序写 bitmap + payload。返回的 buffer 由 encoder 持有。
func (e *RecordEncoder) Encode() ([]byte, error) {
	if e.schemaErr != nil {
		return nil, e.schemaErr
	}
	if e.err != nil {
		return nil, e.err
	}
	for i := range e.bitmap {
		e.bitmap[i] = 0
	}
	for i := range e.set {
		if e.set[i] {
			e.bitmap[i>>6] |= uint64(1) << uint(i&63)
		}
	}
	e.buf = e.buf[:0]
	for _, w := range e.bitmap {
		e.buf = binary.LittleEndian.AppendUint64(e.buf, w)
	}
	for i := range e.schema {
		if !e.set[i] {
			continue
		}
		f := e.schema[i]
		v := &e.vals[i]
		switch f.Type {
		case TypeFlag:
			// 只占位图，没有 payload。
		case TypeString:
			e.buf = binary.AppendUvarint(e.buf, uint64(len(v.b)))
			e.buf = append(e.buf, v.b...)
		case TypeInteger:
			if f.isArray() {
				e.buf = e.appendIntArray(e.buf, v.elems)
			} else {
				e.buf = binary.AppendUvarint(e.buf, zigzagEncode(v.i))
			}
		case TypeFloat:
			if f.isArray() {
				e.buf = e.appendFloatArray(e.buf, v.elems)
			} else {
				e.buf = binary.LittleEndian.AppendUint64(e.buf, math.Float64bits(v.f))
			}
		default:
			return nil, fmt.Errorf("immkv: field %q has invalid type %d", f.Name, f.Type)
		}
	}
	return e.buf, nil
}

func (e *RecordEncoder) appendIntArray(dst []byte, elems []Value) []byte {
	dst = binary.AppendUvarint(dst, uint64(len(elems)))
	for k := range elems {
		el := &elems[k]
		if el.Type == 0 {
			dst = binary.AppendUvarint(dst, 0)
			continue
		}
		e.scratch = binary.AppendUvarint(e.scratch[:0], zigzagEncode(el.Int))
		dst = binary.AppendUvarint(dst, uint64(len(e.scratch)))
		dst = append(dst, e.scratch...)
	}
	return dst
}

func (e *RecordEncoder) appendFloatArray(dst []byte, elems []Value) []byte {
	dst = binary.AppendUvarint(dst, uint64(len(elems)))
	for k := range elems {
		el := &elems[k]
		if el.Type == 0 {
			dst = binary.AppendUvarint(dst, 0)
			continue
		}
		dst = binary.AppendUvarint(dst, 8)
		dst = binary.LittleEndian.AppendUint64(dst, math.Float64bits(el.Float))
	}
	return dst
}

func zigzagEncode(v int64) uint64 {
	return uint64(v)<<1 ^ uint64(v>>63)
}

func zigzagDecode(u uint64) int64 {
	return int64(u>>1) ^ -int64(u&1)
}

func recordBitmapLen(n int) int {
	return ((n + 63) / 64) * 8
}

func bitAt(raw []byte, i int) bool {
	w := i >> 6
	return binary.LittleEndian.Uint64(raw[w*8:])&(uint64(1)<<uint(i&63)) != 0
}

func validateRecord(s Schema, raw []byte) error {
	return walkRecord(s, raw, nil)
}

// walkRecord 校验 record 结构，并按 schema 顺序把每个有值字段的编码 payload
// 交给 fn。payload 是 raw 的视图，数组不展开；fn 为 nil 时只校验。
func walkRecord(s Schema, raw []byte, fn func(index int, payload []byte) error) error {
	bl := recordBitmapLen(len(s))
	if len(raw) < bl {
		return fmt.Errorf("immkv: record 位图长度 %d < %d", len(raw), bl)
	}
	p := raw[bl:]
	for i := range s {
		if !bitAt(raw, i) {
			continue
		}
		before := len(p)
		rest, err := skipValue(s[i], p)
		if err != nil {
			return fmt.Errorf("immkv: 字段 %q: %w", s[i].Name, err)
		}
		if fn != nil {
			if err := fn(i, p[:before-len(rest)]); err != nil {
				return err
			}
		}
		p = rest
	}
	if len(p) != 0 {
		return fmt.Errorf("immkv: record 尾部有 %d 字节多余数据", len(p))
	}
	return nil
}

func skipValue(f Field, p []byte) ([]byte, error) {
	_, rest, err := decodeValue(f, p, false)
	return rest, err
}

// decodeValue 解一个字段。want=false 时不收集数组元素，只校验并跳过；
// 标量路径不分配。
func decodeValue(f Field, p []byte, want bool) (Value, []byte, error) {
	switch f.Type {
	case TypeFlag:
		return Value{Type: TypeFlag}, p, nil
	case TypeString:
		l, rest, err := readLen(p)
		if err != nil {
			return Value{}, nil, err
		}
		return Value{Type: TypeString, Bytes: rest[:l]}, rest[l:], nil
	case TypeInteger, TypeFloat:
		if f.isArray() {
			return decodeArray(f, p, want)
		}
		return decodeScalar(f, p)
	default:
		return Value{}, nil, fmt.Errorf("未知字段类型 %d", f.Type)
	}
}

func decodeScalar(f Field, p []byte) (Value, []byte, error) {
	switch f.Type {
	case TypeInteger:
		u, n := binary.Uvarint(p)
		if n <= 0 {
			return Value{}, nil, fmt.Errorf("integer uvarint 截断或溢出")
		}
		return Value{Type: TypeInteger, Int: zigzagDecode(u)}, p[n:], nil
	case TypeFloat:
		if len(p) < 8 {
			return Value{}, nil, fmt.Errorf("float 需要 8 字节，剩 %d", len(p))
		}
		v := math.Float64frombits(binary.LittleEndian.Uint64(p))
		if math.IsNaN(v) || math.IsInf(v, 0) {
			return Value{}, nil, fmt.Errorf("float 是 NaN 或 Inf")
		}
		return Value{Type: TypeFloat, Float: v}, p[8:], nil
	default:
		return Value{}, nil, fmt.Errorf("字段 %q 不是标量类型", f.Name)
	}
}

func decodeArray(f Field, p []byte, want bool) (Value, []byte, error) {
	count, n := binary.Uvarint(p)
	if n <= 0 {
		return Value{}, nil, fmt.Errorf("数组 count 截断或溢出")
	}
	p = p[n:]
	if count > uint64(len(p)) {
		return Value{}, nil, fmt.Errorf("数组 count %d 超过剩余 %d 字节", count, len(p))
	}
	var elems []Value
	if want {
		elems = make([]Value, int(count))
	}
	for k := 0; k < int(count); k++ {
		l, rest, err := readLen(p)
		if err != nil {
			return Value{}, nil, err
		}
		if l > 0 {
			elem := rest[:l]
			switch f.Type {
			case TypeInteger:
				u, en := binary.Uvarint(elem)
				if en <= 0 {
					return Value{}, nil, fmt.Errorf("数组元素 %d integer 截断或溢出", k)
				}
				if en != len(elem) {
					return Value{}, nil, fmt.Errorf("数组元素 %d integer 长度 %d ≠ %d", k, en, len(elem))
				}
				if want {
					elems[k] = Value{Type: TypeInteger, Int: zigzagDecode(u)}
				}
			case TypeFloat:
				if l != 8 {
					return Value{}, nil, fmt.Errorf("数组元素 %d float 长度 %d ≠ 8", k, l)
				}
				v := math.Float64frombits(binary.LittleEndian.Uint64(elem))
				if math.IsNaN(v) || math.IsInf(v, 0) {
					return Value{}, nil, fmt.Errorf("数组元素 %d float 是 NaN 或 Inf", k)
				}
				if want {
					elems[k] = Value{Type: TypeFloat, Float: v}
				}
			default:
				return Value{}, nil, fmt.Errorf("数组元素类型 %s 不支持", f.Type)
			}
			_ = elem
		}
		p = rest[l:]
	}
	return Value{Type: f.Type, Elems: elems}, p, nil
}

func readLen(p []byte) (int, []byte, error) {
	l, n := binary.Uvarint(p)
	if n <= 0 {
		return 0, nil, fmt.Errorf("uvarint 截断或溢出")
	}
	if l > uint64(len(p)-n) {
		return 0, nil, fmt.Errorf("长度 %d 超过剩余 %d 字节", l, len(p)-n)
	}
	return int(l), p[n:], nil
}
