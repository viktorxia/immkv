package immkv

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"testing"
)

func v003RecordSchema() Schema {
	return Schema{
		{Name: "AC", Type: TypeInteger, Number: "A"},
		{Name: "AN", Type: TypeInteger, Number: "1"},
		{Name: "AF", Type: TypeFloat, Number: "A"},
		{Name: "DB", Type: TypeFlag, Number: "0"},
		{Name: "nhomalt", Type: TypeInteger, Number: "A"},
		{Name: "vep", Type: TypeString, Number: "."},
	}
}

func TestRecordEncodeDecode(t *testing.T) {
	s := v003RecordSchema()
	enc := NewRecordEncoder(s)
	enc.Int("AC", 12).
		Int("AN", 9800).
		Float("AF", 0.001225).
		Int("nhomalt", 3).
		Bytes("vep", []byte("missense_variant"))
	raw, err := enc.Encode()
	if err != nil {
		t.Fatal(err)
	}

	// bitmap：AC/AN/AF/nhomalt/vep 有值，DB 缺失。
	if got := binary.LittleEndian.Uint64(raw[:8]); got != 0x37 {
		t.Fatalf("bitmap = %#x, want 0x37", got)
	}
	if raw[8] != 0x18 { // zigzag(12) = 24
		t.Fatalf("AC byte = %#x, want 0x18", raw[8])
	}
	an := binary.AppendUvarint(nil, zigzagEncode(9800))
	if !bytes.Equal(raw[9:9+len(an)], an) {
		t.Fatalf("AN bytes = %x, want %x", raw[9:9+len(an)], an)
	}
	afOff := 9 + len(an)
	if got := math.Float64frombits(binary.LittleEndian.Uint64(raw[afOff:])); got != 0.001225 {
		t.Fatalf("AF = %v, want 0.001225", got)
	}
	nhOff := afOff + 8
	if raw[nhOff] != 0x06 { // zigzag(3) = 6
		t.Fatalf("nhomalt byte = %#x, want 0x06", raw[nhOff])
	}
	vepOff := nhOff + 1
	if raw[vepOff] != 16 {
		t.Fatalf("vep length = %d, want 16", raw[vepOff])
	}
	if got := string(raw[vepOff+1:]); got != "missense_variant" {
		t.Fatalf("vep = %q", got)
	}
	if len(raw) != 38 {
		t.Fatalf("record length = %d, want 38", len(raw))
	}

	rec, err := NewRecord(s, raw)
	if err != nil {
		t.Fatal(err)
	}
	if !rec.HasValue("AC") || rec.Int("AC") != 12 {
		t.Fatalf("AC = %d", rec.Int("AC"))
	}
	if !rec.HasValue("AN") || rec.Int("AN") != 9800 {
		t.Fatalf("AN = %d", rec.Int("AN"))
	}
	if !rec.HasValue("AF") || rec.Float("AF") != 0.001225 {
		t.Fatalf("AF = %v", rec.Float("AF"))
	}
	if rec.HasValue("DB") {
		t.Fatal("DB should be absent")
	}
	if !rec.HasValue("nhomalt") || rec.Int("nhomalt") != 3 {
		t.Fatalf("nhomalt = %d", rec.Int("nhomalt"))
	}
	if !rec.HasValue("vep") || string(rec.Bytes("vep")) != "missense_variant" {
		t.Fatalf("vep = %q", rec.Bytes("vep"))
	}

	var names []string
	for f, _ := range rec.Values() {
		names = append(names, f.Name)
	}
	want := []string{"AC", "AN", "AF", "nhomalt", "vep"}
	if fmt.Sprint(names) != fmt.Sprint(want) {
		t.Fatalf("Values order = %v, want %v", names, want)
	}
}

func TestRecordFlag(t *testing.T) {
	s := Schema{{Name: "DB", Type: TypeFlag, Number: "0"}}
	raw, err := NewRecordEncoder(s).Flag("DB").Encode()
	if err != nil {
		t.Fatal(err)
	}
	if len(raw) != 8 {
		t.Fatalf("flag record length = %d, want 8", len(raw))
	}
	if got := binary.LittleEndian.Uint64(raw); got != 1 {
		t.Fatalf("bitmap = %#x, want 1", got)
	}
	rec, err := NewRecord(s, raw)
	if err != nil {
		t.Fatal(err)
	}
	if !rec.HasValue("DB") {
		t.Fatal("DB should be present")
	}
	text, err := RenderRecord(s, raw)
	if err != nil {
		t.Fatal(err)
	}
	if string(text) != "DB" {
		t.Fatalf("render = %q, want %q", text, "DB")
	}
}

func TestRecordMissingAndEmpty(t *testing.T) {
	s := v003RecordSchema()
	raw, err := NewRecordEncoder(s).Int("AC", 12).Encode()
	if err != nil {
		t.Fatal(err)
	}
	rec, err := NewRecord(s, raw)
	if err != nil {
		t.Fatal(err)
	}
	if rec.HasValue("AF") || rec.Float("AF") != 0 {
		t.Fatal("AF should be missing")
	}
	if rec.HasValue("vep") || rec.Bytes("vep") != nil {
		t.Fatal("vep should be missing")
	}
	text, err := RenderRecord(s, raw)
	if err != nil {
		t.Fatal(err)
	}
	if string(text) != "AC=12" {
		t.Fatalf("render = %q, want %q", text, "AC=12")
	}

	empty, err := NewRecordEncoder(s).Encode()
	if err != nil {
		t.Fatal(err)
	}
	if len(empty) != 8 {
		t.Fatalf("empty record length = %d, want 8", len(empty))
	}
	if _, err := NewRecord(s, empty); err != nil {
		t.Fatal(err)
	}
	text, err = RenderRecord(s, empty)
	if err != nil {
		t.Fatal(err)
	}
	if string(text) != "." {
		t.Fatalf("empty render = %q, want %q", text, ".")
	}
}

func TestRecordBitmapBoundaries(t *testing.T) {
	for _, n := range []int{1, 64, 65} {
		fields := make(Schema, n)
		for i := range fields {
			fields[i] = Field{Name: fmt.Sprintf("F%d", i), Type: TypeInteger, Number: "1"}
		}
		enc := NewRecordEncoder(fields).Int("F0", 1)
		if n > 1 {
			enc.Int(fmt.Sprintf("F%d", n-1), 2)
		}
		raw, err := enc.Encode()
		if err != nil {
			t.Fatal(err)
		}
		wantLen := recordBitmapLen(n) + 1
		if n > 1 {
			wantLen++
		}
		if len(raw) != wantLen {
			t.Fatalf("n=%d record length = %d, want %d", n, len(raw), wantLen)
		}
		rec, err := NewRecord(fields, raw)
		if err != nil {
			t.Fatal(err)
		}
		if !rec.HasValue("F0") {
			t.Fatalf("n=%d F0 missing", n)
		}
		if n > 1 && !rec.HasValue(fmt.Sprintf("F%d", n-1)) {
			t.Fatalf("n=%d last missing", n)
		}
		if n > 2 && rec.HasValue("F1") {
			t.Fatalf("n=%d F1 should be missing", n)
		}
	}
}

func TestRecordArray(t *testing.T) {
	s := Schema{
		{Name: "AD", Type: TypeInteger, Number: "R"},
		{Name: "GL", Type: TypeFloat, Number: "G"},
		{Name: "XX", Type: TypeInteger, Number: "."},
	}
	enc := NewRecordEncoder(s)
	enc.Values("AD", []Value{
		{Type: TypeInteger, Int: 1},
		{},
		{Type: TypeInteger, Int: 3},
	})
	enc.Values("GL", []Value{
		{Type: TypeFloat, Float: 0.5},
		{},
	})
	enc.Values("XX", nil)
	raw, err := enc.Encode()
	if err != nil {
		t.Fatal(err)
	}
	if len(raw) != 26 {
		t.Fatalf("array record length = %d, want 26", len(raw))
	}
	if got := binary.LittleEndian.Uint64(raw[:8]); got != 0x07 {
		t.Fatalf("bitmap = %#x, want 0x07", got)
	}
	wantAD := []byte{3, 1, 0x02, 0, 1, 0x06}
	if !bytes.Equal(raw[8:14], wantAD) {
		t.Fatalf("AD bytes = %x, want %x", raw[8:14], wantAD)
	}
	if raw[14] != 2 || raw[15] != 8 {
		t.Fatalf("GL prefix = %x, want count=2 len=8", raw[14:16])
	}
	if got := math.Float64frombits(binary.LittleEndian.Uint64(raw[16:24])); got != 0.5 {
		t.Fatalf("GL[0] = %v, want 0.5", got)
	}
	if raw[24] != 0 || raw[25] != 0 {
		t.Fatalf("GL missing/XX count = %x", raw[24:26])
	}

	rec, err := NewRecord(s, raw)
	if err != nil {
		t.Fatal(err)
	}
	got := map[string]Value{}
	for f, v := range rec.Values() {
		got[f.Name] = v
	}
	ad := got["AD"]
	if ad.Type != TypeInteger || len(ad.Elems) != 3 {
		t.Fatalf("AD = %+v", ad)
	}
	if ad.Elems[0].Int != 1 || ad.Elems[1].Type != 0 || ad.Elems[2].Int != 3 {
		t.Fatalf("AD elems = %+v", ad.Elems)
	}
	gl := got["GL"]
	if gl.Type != TypeFloat || len(gl.Elems) != 2 || gl.Elems[0].Float != 0.5 || gl.Elems[1].Type != 0 {
		t.Fatalf("GL = %+v", gl)
	}
	if xx := got["XX"]; xx.Type != TypeInteger || len(xx.Elems) != 0 {
		t.Fatalf("XX = %+v", xx)
	}
	text, err := RenderRecord(s, raw)
	if err != nil {
		t.Fatal(err)
	}
	if string(text) != "AD=1,.,3;GL=0.5,.;XX=." {
		t.Fatalf("render = %q", text)
	}
	mustPanic(t, func() { _ = rec.Int("AD") })
}

func TestRecordArrayZeroVsMissing(t *testing.T) {
	s := Schema{{Name: "AD", Type: TypeInteger, Number: "R"}}
	raw, err := NewRecordEncoder(s).Values("AD", []Value{
		{Type: TypeInteger, Int: 0},
		{},
	}).Encode()
	if err != nil {
		t.Fatal(err)
	}
	rec, err := NewRecord(s, raw)
	if err != nil {
		t.Fatal(err)
	}
	var ad Value
	for f, v := range rec.Values() {
		if f.Name == "AD" {
			ad = v
		}
	}
	if len(ad.Elems) != 2 {
		t.Fatalf("AD elems = %+v", ad.Elems)
	}
	if ad.Elems[0].Type != TypeInteger || ad.Elems[0].Int != 0 {
		t.Fatalf("AD[0] = %+v, want present 0", ad.Elems[0])
	}
	if ad.Elems[1].Type != 0 {
		t.Fatalf("AD[1] = %+v, want missing", ad.Elems[1])
	}
	text, err := RenderRecord(s, raw)
	if err != nil {
		t.Fatal(err)
	}
	if string(text) != "AD=0,." {
		t.Fatalf("render = %q, want %q", text, "AD=0,.")
	}
}

func TestRecordRenderNormalization(t *testing.T) {
	s := Schema{
		{Name: "I", Type: TypeInteger, Number: "1"},
		{Name: "F", Type: TypeFloat, Number: "1"},
		{Name: "S", Type: TypeString, Number: "1"},
		{Name: "B", Type: TypeFlag, Number: "0"},
	}
	raw, err := NewRecordEncoder(s).
		Int("I", -5).
		Float("F", 1142.0).
		Bytes("S", []byte("x")).
		Flag("B").
		Encode()
	if err != nil {
		t.Fatal(err)
	}
	text, err := RenderRecord(s, raw)
	if err != nil {
		t.Fatal(err)
	}
	if string(text) != "I=-5;F=1142;S=x;B" {
		t.Fatalf("render = %q", text)
	}

	emptyString, err := NewRecordEncoder(Schema{
		{Name: "S", Type: TypeString, Number: "1"},
	}).Bytes("S", nil).Encode()
	if err != nil {
		t.Fatal(err)
	}
	text, err = RenderRecord(Schema{{Name: "S", Type: TypeString, Number: "1"}}, emptyString)
	if err != nil {
		t.Fatal(err)
	}
	if string(text) != "S=" {
		t.Fatalf("empty string render = %q, want %q", text, "S=")
	}
}

func TestRecordEncoderErrors(t *testing.T) {
	s := v003RecordSchema()
	cases := []struct {
		name string
		enc  *RecordEncoder
	}{
		{"unknown field", NewRecordEncoder(s).Int("NOPE", 1)},
		{"int on float", NewRecordEncoder(s).Int("AF", 1)},
		{"float on int", NewRecordEncoder(s).Float("AC", 1)},
		{"bytes on int", NewRecordEncoder(s).Bytes("AC", []byte("x"))},
		{"flag on int", NewRecordEncoder(s).Flag("AC")},
		{"float NaN", NewRecordEncoder(s).Float("AF", math.NaN())},
		{"float Inf", NewRecordEncoder(s).Float("AF", math.Inf(1))},
		{"array via int", NewRecordEncoder(Schema{
			{Name: "AD", Type: TypeInteger, Number: "R"},
		}).Int("AD", 1)},
		{"scalar via values", NewRecordEncoder(s).Values("AC", []Value{{Type: TypeInteger, Int: 1}})},
		{"array element type", NewRecordEncoder(Schema{
			{Name: "GL", Type: TypeFloat, Number: "G"},
		}).Values("GL", []Value{{Type: TypeInteger, Int: 1}})},
		{"invalid schema", NewRecordEncoder(Schema{
			{Name: "F", Type: TypeFlag, Number: "1"},
		})},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if _, err := c.enc.Encode(); err == nil {
				t.Fatal("Encode() = nil error, want error")
			}
		})
	}
}

func TestRecordCorruption(t *testing.T) {
	intSchema := Schema{{Name: "I", Type: TypeInteger, Number: "1"}}
	floatSchema := Schema{{Name: "F", Type: TypeFloat, Number: "1"}}
	strSchema := Schema{{Name: "S", Type: TypeString, Number: "1"}}
	arrSchema := Schema{{Name: "AD", Type: TypeInteger, Number: "R"}}
	floatArrSchema := Schema{{Name: "GL", Type: TypeFloat, Number: "G"}}

	floatNaN := make([]byte, 8)
	binary.LittleEndian.PutUint64(floatNaN, math.Float64bits(math.NaN()))

	cases := []struct {
		name   string
		schema Schema
		raw    []byte
	}{
		{"bitmap short", intSchema, []byte{1, 0, 0}},
		{"payload short", intSchema, v003Raw(1)},
		{"trailing", intSchema, v003Raw(1, 0x00, 0xff)},
		{"integer overflow", intSchema, v003Raw(1, bytes.Repeat([]byte{0xff}, 10)...)},
		{"float NaN", floatSchema, v003Raw(1, floatNaN...)},
		{"string len overrun", strSchema, v003Raw(1, append(binary.AppendUvarint(nil, 5), 'a', 'b')...)},
		{"array count overrun", arrSchema, v003Raw(1, append(binary.AppendUvarint(nil, 5), 0)...)},
		{"array len overrun", arrSchema, v003Raw(1, append(binary.AppendUvarint(nil, 1), append(binary.AppendUvarint(nil, 5), 'a', 'b')...)...)},
		{"array int len mismatch", arrSchema, v003Raw(1, append(binary.AppendUvarint(nil, 1), append(binary.AppendUvarint(nil, 2), 0x01, 0x00)...)...)},
		{"array float len", floatArrSchema, v003Raw(1, append(binary.AppendUvarint(nil, 1), append(binary.AppendUvarint(nil, 4), 0, 0, 0, 0)...)...)},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if _, err := NewRecord(c.schema, c.raw); err == nil {
				t.Fatal("NewRecord() = nil error, want error")
			}
		})
	}
}

func TestRecordPanics(t *testing.T) {
	s := Schema{
		{Name: "I", Type: TypeInteger, Number: "1"},
		{Name: "F", Type: TypeFloat, Number: "1"},
		{Name: "S", Type: TypeString, Number: "1"},
		{Name: "B", Type: TypeFlag, Number: "0"},
		{Name: "AD", Type: TypeInteger, Number: "R"},
	}
	raw, err := NewRecordEncoder(s).Int("I", 1).Encode()
	if err != nil {
		t.Fatal(err)
	}
	rec, err := NewRecord(s, raw)
	if err != nil {
		t.Fatal(err)
	}
	mustPanic(t, func() { rec.HasValue("NOPE") })
	mustPanic(t, func() { _ = rec.Int("F") })
	mustPanic(t, func() { _ = rec.Float("I") })
	mustPanic(t, func() { _ = rec.Bytes("I") })
	mustPanic(t, func() { _ = rec.Int("AD") })
}

func TestRecordDecodeAllocs(t *testing.T) {
	s := v003RecordSchema()
	raw, err := NewRecordEncoder(s).
		Int("AC", 12).
		Int("AN", 9800).
		Float("AF", 0.001225).
		Int("nhomalt", 3).
		Bytes("vep", []byte("missense_variant")).
		Encode()
	if err != nil {
		t.Fatal(err)
	}
	rec, err := NewRecord(s, raw)
	if err != nil {
		t.Fatal(err)
	}
	if got := testing.AllocsPerRun(1000, func() {
		var n int64
		if rec.HasValue("AC") {
			n += rec.Int("AC")
		}
		if rec.HasValue("AN") {
			n += rec.Int("AN")
		}
		if rec.HasValue("nhomalt") {
			n += rec.Int("nhomalt")
		}
		_ = n
	}); got != 0 {
		t.Fatalf("scalar access allocations = %v, want 0", got)
	}
	if got := testing.AllocsPerRun(1000, func() {
		for f, v := range rec.Values() {
			_, _ = f, v
		}
	}); got != 0 {
		t.Fatalf("Values allocations = %v, want 0", got)
	}
}

func mustPanic(t *testing.T, fn func()) {
	t.Helper()
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic")
		}
	}()
	fn()
}

func v003Raw(bits uint64, payload ...byte) []byte {
	raw := make([]byte, 8+len(payload))
	binary.LittleEndian.PutUint64(raw, bits)
	copy(raw[8:], payload)
	return raw
}

// TestProjectRecordWideSchema 验证投影位图按字节寻址：跨 64 位字边界的字段
// （第 8 列及以后）位不丢、不写错字节（回归：曾按 i>>6/i&63 寻址 []byte）。
func TestProjectRecordWideSchema(t *testing.T) {
	for _, col := range []int{0, 7, 8, 63, 64, 71} {
		schema := make(Schema, 72)
		for i := range schema {
			schema[i] = Field{Name: fmt.Sprintf("F%d", i), Type: TypeFlag, Number: "0"}
		}
		raw, err := NewRecordEncoder(schema).Flag(schema[col].Name).Encode()
		if err != nil {
			t.Fatal(err)
		}
		selected := make([]bool, len(schema))
		selected[col] = true
		got, err := projectRecord(schema, raw, selected)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(got, raw) {
			t.Errorf("column %d: projection changed the presence bitmap: got=%x want=%x", col, got, raw)
		}
	}
}
