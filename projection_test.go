package immkv

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"sync"
	"testing"
)

func projectionTestSchema(n int) Schema {
	fields := []Field{
		{Type: TypeInteger, Number: "1"},
		{Type: TypeFloat, Number: "1"},
		{Type: TypeString, Number: "1"},
		{Type: TypeFlag, Number: "0"},
		{Type: TypeInteger, Number: "R"},
		{Type: TypeFloat, Number: "G"},
		{Type: TypeInteger, Number: "."},
		{Type: TypeFloat, Number: "."},
		{Type: TypeInteger, Number: "A"},
		{Type: TypeFloat, Number: "A"},
		{Type: TypeString, Number: "R"},
		{Type: TypeString, Number: "."},
	}
	schema := make(Schema, n)
	for i := range schema {
		schema[i] = fields[i%len(fields)]
		schema[i].Name = fmt.Sprintf("F%d", i)
	}
	return schema
}

func projectionTestValues(schema Schema, rng *rand.Rand, row int) []Value {
	integers := []int64{math.MinInt64, math.MaxInt64, -8193, -8192, -65, -64, -1, 0, 1, 63, 64, 8191, 8192}
	floats := []float64{0, math.Copysign(0, -1), math.SmallestNonzeroFloat64, -math.SmallestNonzeroFloat64, math.MaxFloat64, -math.MaxFloat64, 0.125, -2.5}
	scalar := func(typ FieldType) Value {
		v := Value{Type: typ}
		switch typ {
		case TypeInteger:
			v.Int = integers[rng.Intn(len(integers))]
			if rng.Intn(2) == 0 {
				v.Int = int64(rng.Uint64())
			}
		case TypeFloat:
			v.Float = floats[rng.Intn(len(floats))]
			if rng.Intn(2) == 0 {
				bits := rng.Uint64()
				if bits&0x7ff0000000000000 == 0x7ff0000000000000 {
					bits ^= 1 << 52
				}
				v.Float = math.Float64frombits(bits)
			}
		case TypeString:
			lengths := []int{0, 1, 2, 127, 128, 255, 1024}
			v.Bytes = make([]byte, lengths[rng.Intn(len(lengths))])
			_, _ = rng.Read(v.Bytes)
		}
		return v
	}
	values := make([]Value, len(schema))
	for i, field := range schema {
		switch row % 8 {
		case 1:
			continue // An entirely missing record.
		case 2:
			if field.Type != TypeFlag {
				continue
			}
		case 3:
			if i != 0 && i != len(schema)-1 && i%64 != 0 && i%64 != 63 {
				continue
			}
		case 4, 5, 6, 7:
			if rng.Intn(3) == 0 {
				continue
			}
		}
		if !field.isArray() {
			values[i] = scalar(field.Type)
			continue
		}
		lengths := []int{0, 1, 2, 5, 17}
		elems := make([]Value, lengths[rng.Intn(len(lengths))])
		for j := range elems {
			if rng.Intn(3) != 0 {
				elems[j] = scalar(field.Type)
			}
		}
		values[i] = Value{Type: field.Type, Elems: elems}
	}
	return values
}

// A nil selection encodes the full record; an empty non-nil selection encodes
// an all-zero bitmap. Expectations come from the encoder, not the projector.
func projectionTestEncode(t testing.TB, schema Schema, values []Value, selected []int) []byte {
	t.Helper()
	var keep []bool
	if selected != nil {
		keep = make([]bool, len(schema))
		for _, i := range selected {
			keep[i] = true
		}
	}
	enc := NewRecordEncoder(schema)
	for i, field := range schema {
		v := values[i]
		if v.Type == 0 || (keep != nil && !keep[i]) {
			continue
		}
		switch {
		case field.isArray():
			enc.Values(field.Name, v.Elems)
		case field.Type == TypeFlag:
			enc.Flag(field.Name)
		case field.Type == TypeInteger:
			enc.Int(field.Name, v.Int)
		case field.Type == TypeFloat:
			enc.Float(field.Name, v.Float)
		case field.Type == TypeString:
			enc.Bytes(field.Name, v.Bytes)
		}
	}
	raw, err := enc.Encode()
	if err != nil {
		t.Fatal(err)
	}
	return bytes.Clone(raw)
}

// Keep the reference independent of the compiled mask and scalar fast paths.
// Copy encoded payloads so accepted non-minimal varints remain byte-exact.
func projectionTestReference(schema Schema, raw []byte, selected []int) ([]byte, error) {
	keep := make([]bool, len(schema))
	for _, i := range selected {
		keep[i] = true
	}
	out := make([]byte, recordBitmapLen(len(schema)))
	err := walkRecord(schema, raw, func(i int, payload []byte) error {
		if keep[i] {
			out[i/8] |= 1 << uint(i%8)
			out = append(out, payload...)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return out, nil
}

func projectionTestAllColumns(n int) []int {
	all := make([]int, n)
	for i := range all {
		all[i] = n - i - 1 // Selection order must not become payload order.
	}
	return all
}

func TestRecordProjectionMatrix(t *testing.T) {
	for _, width := range []int{1, 7, 8, 63, 64, 65, 127, 128, 129, 193} {
		t.Run(fmt.Sprintf("fields_%d", width), func(t *testing.T) {
			schema := projectionTestSchema(width)
			rng := rand.New(rand.NewSource(int64(20261115 + width)))
			selections := [][]int{{}, projectionTestAllColumns(width), {}, {}, {}, {}}
			for i := range schema {
				if i%2 == 0 {
					selections[2] = append(selections[2], i)
				}
				if i == 0 || i == width-1 || i%64 == 0 || i%64 == 63 || i == 7 || i == 8 {
					selections[3] = append(selections[3], i)
				}
				if i >= width/3 && i <= 2*width/3 {
					selections[4] = append(selections[4], i)
				}
				if rng.Intn(3) == 0 {
					selections[5] = append(selections[5], i)
				}
			}
			projectors := make([]*recordProjection, len(selections))
			for i, selected := range selections {
				projectors[i] = newRecordProjection(schema, selected)
			}
			for row := 0; row < 24; row++ {
				values := projectionTestValues(schema, rng, row)
				raw := projectionTestEncode(t, schema, values, nil)
				original := bytes.Clone(raw)
				for selection, selected := range selections {
					want := projectionTestEncode(t, schema, values, selected)
					for _, destination := range []string{"nil", "grow", "reuse"} {
						var dst, prefix []byte
						if destination != "nil" {
							prefix = []byte{0xa5, 0, 0xff, 0x5a, 0x7f}
							capacity := len(prefix)
							if destination == "reuse" {
								capacity += len(raw) + 64
							}
							dst = bytes.Repeat([]byte{0xcc}, capacity)[:len(prefix)]
							copy(dst, prefix)
						}
						got, err := projectors[selection].appendRecord(dst, raw)
						if err != nil {
							t.Fatalf("row %d selection %d dst %s: %v", row, selection, destination, err)
						}
						if got == nil || !bytes.Equal(got, append(bytes.Clone(prefix), want...)) {
							t.Fatalf("row %d selection %d dst %s: got %x, want prefix %x + %x", row, selection, destination, got, prefix, want)
						}
						if !bytes.Equal(dst, prefix) {
							t.Fatalf("row %d selection %d: destination prefix changed", row, selection)
						}
						if _, err := NewRecord(schema, got[len(prefix):]); err != nil {
							t.Fatalf("row %d selection %d: projected record is invalid: %v", row, selection, err)
						}
						// Mutation also detects a returned view into the input.
						got[len(prefix)] ^= 0xff
						if !bytes.Equal(raw, original) {
							t.Fatalf("row %d selection %d dst %s: input changed or aliases output", row, selection, destination)
						}
					}
				}
			}
		})
	}
}

func TestRecordProjectionMissingValues(t *testing.T) {
	schema := projectionTestSchema(129)
	values := make([]Value, len(schema))
	values[0] = Value{Type: TypeInteger, Int: 42}
	raw := projectionTestEncode(t, schema, values, nil)
	for _, selected := range [][]int{nil, {}, {64, 128}} {
		got, err := newRecordProjection(schema, selected).appendRecord(nil, raw)
		if err != nil {
			t.Fatal(err)
		}
		if got == nil || !bytes.Equal(got, make([]byte, recordBitmapLen(len(schema)))) {
			t.Fatalf("selected %v: missing fields must produce a non-nil all-zero bitmap, got %x", selected, got)
		}
	}
}

func TestRecordProjectionAppendRetainsRecords(t *testing.T) {
	schema := projectionTestSchema(137)
	selected := []int{136, 130, 128, 127, 65, 64, 63, 8, 7, 3, 2, 1, 0}
	projection := newRecordProjection(schema, selected)
	rng := rand.New(rand.NewSource(991))
	prefix := []byte{0xff, 0, 0xfe}
	dst := bytes.Clone(prefix)
	want := bytes.Clone(prefix)
	var retained, records [][]byte
	for row := 0; row < 12; row++ {
		values := projectionTestValues(schema, rng, row)
		raw := projectionTestEncode(t, schema, values, nil)
		expected := projectionTestEncode(t, schema, values, selected)
		start := len(dst)
		var err error
		dst, err = projection.appendRecord(dst, raw)
		if err != nil {
			t.Fatal(err)
		}
		retained = append(retained, dst[start:len(dst):len(dst)])
		records = append(records, expected)
		want = append(want, expected...)
		if !bytes.Equal(dst, want) {
			t.Fatalf("append %d changed an earlier record or prefix", row)
		}
		for i := range retained {
			if !bytes.Equal(retained[i], records[i]) {
				t.Fatalf("append %d changed retained record %d", row, i)
			}
		}
	}
}

func TestRecordProjectionRejectsUnselectedCorruption(t *testing.T) {
	intField := Field{Type: TypeInteger, Number: "1"}
	floatField := Field{Type: TypeFloat, Number: "A"}
	stringField := Field{Type: TypeString, Number: "R"}
	intArray := Field{Type: TypeInteger, Number: "R"}
	floatArray := Field{Type: TypeFloat, Number: "."}
	overflow := bytes.Repeat([]byte{0xff}, 10)
	floatBytes := func(v float64) []byte {
		return binary.LittleEndian.AppendUint64(nil, math.Float64bits(v))
	}
	cases := []struct {
		name    string
		field   Field
		payload []byte
	}{
		{"integer_missing", intField, nil},
		{"integer_truncated", intField, []byte{0x80}},
		{"integer_overflow", intField, overflow},
		{"integer_tenth_byte_overflow", intField, append(bytes.Repeat([]byte{0x80}, 9), 2)},
		{"float_truncated", floatField, make([]byte, 7)},
		{"float_nan", floatField, floatBytes(math.NaN())},
		{"float_signaling_nan", floatField, binary.LittleEndian.AppendUint64(nil, 0x7ff0000000000001)},
		{"float_positive_inf", floatField, floatBytes(math.Inf(1))},
		{"float_negative_inf", floatField, floatBytes(math.Inf(-1))},
		{"string_length_missing", stringField, nil},
		{"string_length_truncated", stringField, []byte{0x80}},
		{"string_length_overflow", stringField, overflow},
		{"string_length_overrun", stringField, []byte{3, 'a'}},
		{"string_length_max_uint64", stringField, binary.AppendUvarint(nil, math.MaxUint64)},
		{"array_count_missing", intArray, nil},
		{"array_count_truncated", intArray, []byte{0x80}},
		{"array_count_overflow", intArray, overflow},
		{"array_count_overrun", intArray, []byte{3, 0}},
		{"array_count_max_uint64", intArray, binary.AppendUvarint(nil, math.MaxUint64)},
		{"array_element_length_missing", intArray, []byte{1}},
		{"array_element_length_truncated", intArray, []byte{1, 0x80}},
		{"array_element_length_overflow", intArray, append([]byte{1}, overflow...)},
		{"array_element_length_overrun", intArray, []byte{1, 5, 0}},
		{"array_integer_truncated", intArray, []byte{1, 1, 0x80}},
		{"array_integer_overflow", intArray, append([]byte{1, 10}, overflow...)},
		{"array_integer_length_mismatch", intArray, []byte{1, 2, 0, 0}},
		{"array_float_short", floatArray, append([]byte{1, 7}, make([]byte, 7)...)},
		{"array_float_long", floatArray, append([]byte{1, 9}, make([]byte, 9)...)},
		{"array_float_truncated", floatArray, append([]byte{1, 8}, make([]byte, 7)...)},
		{"array_float_nan", floatArray, append([]byte{2, 0, 8}, floatBytes(math.NaN())...)},
		{"array_float_positive_inf", floatArray, append([]byte{2, 0, 8}, floatBytes(math.Inf(1))...)},
		{"array_float_negative_inf", floatArray, append([]byte{2, 0, 8}, floatBytes(math.Inf(-1))...)},
		{"array_trailing", intArray, []byte{0, 0xff}},
		{"flag_trailing", Field{Type: TypeFlag, Number: "0"}, []byte{0xff}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			for _, column := range []int{0, 63, 64, 128} {
				schema := make(Schema, 131)
				for i := range schema {
					schema[i] = Field{Name: fmt.Sprintf("F%d", i), Type: TypeFlag, Number: "0"}
				}
				schema[column] = tc.field
				schema[column].Name = fmt.Sprintf("F%d", column)
				raw := make([]byte, recordBitmapLen(len(schema)))
				var selected []int
				for _, i := range []int{0, 62, 65, 129, 130} {
					raw[i/8] |= 1 << uint(i%8)
					if i != column {
						selected = append(selected, i)
					}
				}
				raw[column/8] |= 1 << uint(column%8)
				raw = append(raw, tc.payload...)
				if _, err := NewRecord(schema, raw); err == nil {
					t.Fatalf("column %d: malformed fixture unexpectedly validates", column)
				}
				original := bytes.Clone(raw)
				for _, selection := range [][]int{nil, selected, projectionTestAllColumns(len(schema))} {
					dst := make([]byte, 3, len(raw)+64)
					copy(dst, []byte{0xa5, 0, 0x5a})
					if _, err := newRecordProjection(schema, selection).appendRecord(dst, raw); err == nil {
						t.Fatalf("column %d selection %v: malformed record accepted", column, selection)
					}
					if !bytes.Equal(dst, []byte{0xa5, 0, 0x5a}) || !bytes.Equal(raw, original) {
						t.Fatalf("column %d: error changed the destination prefix or input", column)
					}
				}
			}
		})
	}
	t.Run("bitmap_truncated", func(t *testing.T) {
		schema := projectionTestSchema(129)
		projection := newRecordProjection(schema, []int{128})
		for _, length := range []int{0, 1, 7, 8, 15, 16, 23} {
			if _, err := projection.appendRecord(nil, make([]byte, length)); err == nil {
				t.Fatalf("accepted %d-byte bitmap for 129 fields", length)
			}
		}
	})
	t.Run("missing_fields_with_trailing_payload", func(t *testing.T) {
		schema := projectionTestSchema(129)
		raw := make([]byte, recordBitmapLen(len(schema))+1)
		if _, err := newRecordProjection(schema, []int{128}).appendRecord(nil, raw); err == nil {
			t.Fatal("accepted trailing payload after an all-zero bitmap")
		}
	})
}

func FuzzRecordProjection(f *testing.F) {
	schema := projectionTestSchema(137)
	rng := rand.New(rand.NewSource(90125))
	all := make([]byte, len(schema))
	for i := range all {
		all[i] = byte(i)
	}
	for row := 0; row < 8; row++ {
		raw := projectionTestEncode(f, schema, projectionTestValues(schema, rng, row), nil)
		f.Add(raw, all)
		f.Add(raw, []byte{0, 3, 7, 8, 63, 64, 127, 128, 136})
		f.Add(raw, []byte{})
		f.Add(raw[:len(raw)-1], []byte{3, 64, 128})
		f.Add(append(bytes.Clone(raw), 0xff), []byte{3})
	}
	f.Add([]byte{}, []byte{})
	// Both encodings are valid to the reader and must remain byte-exact.
	nonminimal := make([]byte, recordBitmapLen(len(schema)))
	nonminimal[0] = 1
	f.Add(append(bytes.Clone(nonminimal), 0x80, 0), []byte{0})
	f.Add(append(bytes.Clone(nonminimal), 0x81, 0), []byte{0, 0})
	padding := make([]byte, recordBitmapLen(len(schema)))
	padding[len(padding)-1] = 0x80 // Bits beyond the schema are ignored by the reader.
	f.Add(padding, all)
	f.Fuzz(func(t *testing.T, raw, columns []byte) {
		if len(raw) > 64<<10 || len(columns) > 1024 {
			return
		}
		seen := make([]bool, len(schema))
		var selected []int
		for _, b := range columns {
			i := int(b) % len(schema)
			if !seen[i] {
				seen[i] = true
				selected = append(selected, i)
			}
		}
		original := bytes.Clone(raw)
		want, wantErr := projectionTestReference(schema, raw, selected)
		prefix := []byte{0xa5, 0, 0x5a}
		dst := make([]byte, len(prefix), len(prefix)+len(raw)+24)
		copy(dst, prefix)
		got, err := newRecordProjection(schema, selected).appendRecord(dst, raw)
		if !bytes.Equal(raw, original) || !bytes.Equal(dst, prefix) {
			t.Fatal("projection changed the input or destination prefix")
		}
		if (err != nil) != (wantErr != nil) {
			t.Fatalf("compiled error %v, reference error %v, raw %x, selection %v", err, wantErr, raw, selected)
		}
		if err == nil && !bytes.Equal(got, append(bytes.Clone(prefix), want...)) {
			t.Fatalf("compiled %x, reference %x, selection %v", got, want, selected)
		}
	})
}

// Build three bins and more than two micro-blocks. Duplicate records 63/64
// straddle a micro-block boundary and carry different selected values.
func projectionTestLibrary(t *testing.T, opts CreateOptions) (string, Schema, [][]byte, map[string][]byte) {
	t.Helper()
	dir := t.TempDir()
	base := filepath.Join(dir, "chr1")
	schema := projectionTestSchema(137)
	opts.Schema = schema
	writer, err := Create(base, opts)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = writer.Abort() })
	rng := rand.New(rand.NewSource(717))
	byKey := make(map[string][]byte)
	var keys [][]byte
	active := []int{0, 1, 2, 3, 4, 5, 7, 8, 11, 63, 64, 65, 127, 128, 129, 130, 131, 136}
	for row := 0; row < 144; row++ {
		values := make([]Value, len(schema))
		switch row % 9 {
		case 0:
		case 1:
			values[11] = Value{Type: TypeString, Bytes: []byte(fmt.Sprintf("unselected row %d", row))}
		case 2:
			values[63] = Value{Type: TypeFlag}
			values[131] = Value{Type: TypeFlag}
		case 3:
			values[2] = Value{Type: TypeString, Bytes: []byte{}}
			values[64] = Value{Type: TypeInteger, Elems: []Value{}}
			values[127] = Value{Type: TypeFloat, Elems: []Value{{}, {}}}
		default:
			generated := projectionTestValues(schema, rng, 0)
			for _, i := range active {
				values[i] = generated[i]
			}
			values[0] = Value{Type: TypeInteger, Int: int64(row)}
		}
		if row == 40 {
			// A single value larger than the pooled decode buffer, between small
			// records, forces growth without permitting earlier outputs to move.
			values[2] = Value{Type: TypeString, Bytes: bytes.Repeat([]byte("large payload\x00\xff"), 8192)}
		}
		if row == 63 || row == 64 {
			values[0] = Value{Type: TypeInteger, Int: int64(row)}
		}
		ordinal := row
		if row == 64 {
			ordinal = 63
		}
		pos := uint32(100 + ordinal/48*1_000_000 + ordinal%48)
		key := VarKey(0, pos, "A", "G")
		raw := projectionTestEncode(t, schema, values, nil)
		if err := writer.AddSorted(PosKey(0, pos), key, raw); err != nil {
			t.Fatal(err)
		}
		if _, exists := byKey[string(key)]; !exists {
			keys = append(keys, key)
			byKey[string(key)] = raw
		}
	}
	if err := writer.Finish(); err != nil {
		t.Fatal(err)
	}
	if err := writeColumnar(base, schema, [16]byte{1}, SourceInfoOrder{}, 0, colBinWidth); err != nil {
		t.Fatal(err)
	}
	return dir, schema, keys, byKey
}

func projectionTestBatchError(keys, got [][]byte, want map[string][]byte) error {
	if len(got) != len(keys) {
		return fmt.Errorf("GetBatch length %d, want %d", len(got), len(keys))
	}
	for i, key := range keys {
		expected := want[string(key)]
		if (got[i] == nil) != (expected == nil) || !bytes.Equal(got[i], expected) {
			return fmt.Errorf("key %x at %d: got %x (nil=%v), want %x (nil=%v)", key, i, got[i], got[i] == nil, expected, expected == nil)
		}
	}
	return nil
}

func projectionTestConcurrent(t *testing.T, db *StreamDB, keys [][]byte, want map[string][]byte) {
	t.Helper()
	const workers = 4
	start := make(chan struct{})
	errs := make(chan error, workers)
	var wg sync.WaitGroup
	for worker := 0; worker < workers; worker++ {
		wg.Add(1)
		go func(worker int) {
			defer wg.Done()
			<-start
			query := keys[worker : len(keys)-worker]
			retained, err := db.GetBatch(query)
			if err != nil {
				errs <- err
				return
			}
			for round := 0; round < 4; round++ {
				subset := keys[round : len(keys)-round]
				got, err := db.GetBatch(subset)
				if err == nil {
					err = projectionTestBatchError(subset, got, want)
				}
				if err == nil {
					err = projectionTestBatchError(query, retained, want)
				}
				if err != nil {
					errs <- fmt.Errorf("worker %d round %d: %w", worker, round, err)
					return
				}
			}
		}(worker)
	}
	close(start)
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Error(err)
	}
}

func TestStreamingProjectionFraming(t *testing.T) {
	for _, mode := range []struct {
		name string
		opts CreateOptions
	}{
		{"raw", CreateOptions{RawValueBelow: 1 << 20}},
		{"per_value", CreateOptions{RawValueBelow: 1, PackTarget: -1}},
		{"packed", CreateOptions{RawValueBelow: 1, PackTarget: 16 << 10}},
	} {
		t.Run(mode.name, func(t *testing.T) {
			dir, schema, keys, records := projectionTestLibrary(t, mode.opts)
			recs, _, _ := parseKeys(t, filepath.Join(dir, "chr1.keys"))
			shared := false
			for i, record := range recs {
				wantPacked := mode.name == "packed" && len(records[string(record.key)]) < packBelow
				if record.raw != (mode.name == "raw") || record.packed != wantPacked {
					t.Fatalf("record %d: raw=%v packed=%v does not exercise %s", i, record.raw, record.packed, mode.name)
				}
				if i > 0 && recs[i-1].valOff == record.valOff {
					shared = true
				}
			}
			if mode.name == "packed" && !shared {
				t.Fatal("packed fixture has no shared value frame")
			}
			var batch [][]byte
			batch = append(batch, nil)
			for i, key := range keys {
				batch = append(batch, key)
				if i%7 == 0 {
					batch = append(batch, bytes.Clone(key), key)
				}
				batch = append(batch, append(bytes.Clone(key), 0xff))
			}
			batch = append(batch, VarKey(24, 100, "A", "G"))
			for _, selection := range []struct {
				name    string
				columns []int
			}{
				{"default", nil},
				{"all", projectionTestAllColumns(len(schema))},
				{"mixed", []int{136, 130, 129, 128, 127, 100, 65, 64, 63, 8, 7, 3, 2, 1, 0}},
				{"flags", []int{131, 63, 3}},
				{"missing", []int{100}},
			} {
				t.Run(selection.name, func(t *testing.T) {
					want := make(map[string][]byte, len(records))
					for key, raw := range records {
						if selection.name == "default" || selection.name == "all" {
							want[key] = raw
							continue
						}
						projected, err := projectionTestReference(schema, raw, selection.columns)
						if err != nil {
							t.Fatal(err)
						}
						want[key] = projected
					}
					var previous [][]byte
					for _, path := range []string{"record", "columnar"} {
						t.Run(path, func(t *testing.T) {
							t.Setenv("IMMKV_FORCE_RECORD", "")
							t.Setenv("IMMKV_FORCE_COLUMNAR", "")
							if path == "record" {
								t.Setenv("IMMKV_FORCE_RECORD", "1")
							} else {
								t.Setenv("IMMKV_FORCE_COLUMNAR", "1")
							}
							db, err := OpenStreaming(dir, StreamOptions{Threads: 3, Columns: selection.columns})
							if err != nil {
								t.Fatal(err)
							}
							t.Cleanup(func() { _ = db.Close() })
							retained, err := db.GetBatch(batch)
							if err != nil {
								t.Fatal(err)
							}
							if err := projectionTestBatchError(batch, retained, want); err != nil {
								t.Fatal(err)
							}
							if previous != nil {
								for i := range retained {
									if !bytes.Equal(retained[i], previous[i]) {
										t.Fatalf("key %x: record and columnar paths differ", batch[i])
									}
								}
							}
							previous = make([][]byte, len(retained))
							for i := range retained {
								previous[i] = bytes.Clone(retained[i])
							}
							for _, query := range [][][]byte{batch[10:80], batch[len(batch)/2:], {keys[0], keys[0]}} {
								got, err := db.GetBatch(query)
								if err != nil {
									t.Fatal(err)
								}
								if err := projectionTestBatchError(query, got, want); err != nil {
									t.Fatal(err)
								}
								if err := projectionTestBatchError(batch, retained, want); err != nil {
									t.Fatalf("later query changed retained results: %v", err)
								}
							}
							if selection.name == "mixed" || selection.name == "default" || selection.name == "all" {
								projectionTestConcurrent(t, db, batch, want)
								if err := projectionTestBatchError(batch, retained, want); err != nil {
									t.Fatalf("concurrent queries changed retained results: %v", err)
								}
							}
						})
					}
				})
			}
		})
	}
}

func TestStreamingProjectionAllColumns(t *testing.T) {
	dir, schema, keys := buildColTestLib(t)
	batch := [][]byte{keys[0], keys[0], keys[1], keys[2], keys[3], VarKey(24, 100, "A", "G")}
	all := projectionTestAllColumns(len(schema))
	for _, path := range []string{"record", "columnar"} {
		t.Run(path, func(t *testing.T) {
			t.Setenv("IMMKV_FORCE_RECORD", "")
			t.Setenv("IMMKV_FORCE_COLUMNAR", "")
			if path == "record" {
				t.Setenv("IMMKV_FORCE_RECORD", "1")
			} else {
				t.Setenv("IMMKV_FORCE_COLUMNAR", "1")
			}
			var want [][]byte
			for _, columns := range [][]int{nil, {}, all} {
				original := append([]int(nil), columns...)
				db, err := OpenStreaming(dir, StreamOptions{Threads: 2, Columns: columns})
				if err != nil {
					t.Fatal(err)
				}
				t.Cleanup(func() { _ = db.Close() })
				for i := range columns {
					if columns[i] != original[i] {
						t.Fatal("OpenStreaming changed StreamOptions.Columns")
					}
				}
				got, err := db.GetBatch(batch)
				if err != nil {
					t.Fatal(err)
				}
				if len(got) != len(batch) {
					t.Fatalf("GetBatch length %d, want %d", len(got), len(batch))
				}
				if want == nil {
					want = make([][]byte, len(got))
					for i := range got {
						want[i] = bytes.Clone(got[i])
					}
					continue
				}
				for i := range got {
					if (got[i] == nil) != (want[i] == nil) || !bytes.Equal(got[i], want[i]) {
						t.Fatalf("columns %v key %x: got %x, default %x", columns, batch[i], got[i], want[i])
					}
				}
			}
		})
	}
}

func TestExtentBufPoolRejectsWrongCapacity(t *testing.T) {
	// Keep Put/Get on one P without a GC clearing the pools between them.
	// Under -race Put may still discard entries; never require a valid Put
	// to survive, and check every retained entry before requesting a full buffer.
	procs := runtime.GOMAXPROCS(1)
	gc := debug.SetGCPercent(-1)
	extentBufPools = [2]sync.Pool{}
	t.Cleanup(func() {
		extentBufPools = [2]sync.Pool{}
		debug.SetGCPercent(gc)
		runtime.GOMAXPROCS(procs)
	})
	for _, capacity := range []int{0, 1, (16 << 10) - 1, (16 << 10) + 1, 20 << 10, (64 << 10) - 1, (64 << 10) + 1} {
		var b []byte
		if capacity != 0 {
			b = make([]byte, 0, capacity)
		}
		putExtentBuf(b)
		for bucket := range extentBufPools {
			if value := extentBufPools[bucket].Get(); value != nil {
				t.Fatalf("capacity %d entered buffer pool %d", capacity, bucket)
			}
		}
		for _, size := range []int{16 << 10, 64 << 10} {
			got := getExtentBuf(size)
			if len(got) != size || cap(got) < size {
				t.Fatalf("after rejecting capacity %d, getExtentBuf(%d) returned len=%d cap=%d", capacity, size, len(got), cap(got))
			}
		}
	}
	for _, size := range []int{16 << 10, 64 << 10} {
		putExtentBuf(make([]byte, 0, size))
		got := getExtentBuf(size)
		if len(got) != size || cap(got) != size {
			t.Fatalf("valid capacity %d: len=%d cap=%d", size, len(got), cap(got))
		}
	}
}

func TestStreamingProjectionValidatesAllColumnsBeforeNormalization(t *testing.T) {
	dir, schema, _ := buildColTestLib(t)
	all := projectionTestAllColumns(len(schema))
	for _, tc := range []struct {
		name  string
		value int
		extra bool
	}{
		{"same_length_duplicate", all[0], false},
		{"same_length_negative", -1, false},
		{"same_length_out_of_range", len(schema), false},
		{"all_plus_duplicate", all[0], true},
		{"all_plus_negative", -1, true},
		{"all_plus_out_of_range", len(schema), true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			columns := append([]int(nil), all...)
			if tc.extra {
				columns = append(columns, tc.value)
			} else {
				columns[len(columns)-1] = tc.value
			}
			db, err := OpenStreaming(dir, StreamOptions{Threads: 1, Columns: columns})
			if db != nil {
				_ = db.Close()
			}
			if err == nil {
				t.Fatalf("invalid Columns %v accepted as all columns", columns)
			}
		})
	}
}
