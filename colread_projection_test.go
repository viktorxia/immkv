package immkv

import (
	"bytes"
	"encoding/binary"
	"math"
	"strings"
	"testing"

	"github.com/klauspost/compress/zstd"
)

// colProjectionBin 用真实 zstd 帧构造一个 bin，便于独立改坏解压后的值流，
// 不把帧截断错误误当作 typed payload 校验。
func colProjectionBin(t *testing.T, schema Schema, count uint32, bitmaps, payloads [][]byte) (*colShard, *binLoad) {
	t.Helper()
	dec, err := newCGODecoder()
	if err != nil {
		kd, kerr := zstd.NewReader(nil, zstd.WithDecoderConcurrency(1))
		if kerr != nil {
			t.Fatal(kerr)
		}
		dec = &kpDecoder{dec: kd}
	}
	t.Cleanup(func() { dec.Close() })
	c := &colShard{base: "projection-test", schema: schema, dec: dec}
	ld := &binLoad{
		meta:      &colBinRecord{first: 100, count: count},
		parts:     make([]colPart, len(schema)),
		bitFrames: make([][]byte, len(schema)),
		frames:    make([][]byte, len(schema)),
	}
	for i := range schema {
		ld.parts[i] = colPart{col: i, flag: schema[i].Type == TypeFlag, present: popcountBytes(bitmaps[i])}
		ld.bitFrames[i] = colCompress(bitmaps[i])
		if !ld.parts[i].flag {
			ld.frames[i] = colCompress(payloads[i])
		}
	}
	return c, ld
}

func TestColumnarProjectionSpans(t *testing.T) {
	schema := Schema{
		{Name: "I", Type: TypeInteger, Number: "1"},
		{Name: "S", Type: TypeString, Number: "1"},
		{Name: "F", Type: TypeFlag, Number: "0"},
		{Name: "A", Type: TypeInteger, Number: "R"},
	}
	// 三条记录：第一条有全部字段（含空 String、数组缺失元素），第二条
	// 全缺，第三条只有 String 和 Flag；重复第一条跨越装配分块边界。
	bitmaps := [][]byte{{0x01}, {0x05}, {0x05}, {0x01}}
	payloads := [][]byte{{14}, {0, 3, 'e', 'n', 'd'}, nil, {3, 1, 2, 0, 1, 3}}
	c, ld := colProjectionBin(t, schema, 3, bitmaps, payloads)
	enc := NewRecordEncoder(schema)
	first, err := enc.Int("I", 7).Bytes("S", []byte{}).Flag("F").Values("A", []Value{
		{Type: TypeInteger, Int: 1}, {}, {Type: TypeInteger, Int: -2},
	}).Encode()
	if err != nil {
		t.Fatal(err)
	}
	first = bytes.Clone(first)
	missing, err := enc.Reset().Encode()
	if err != nil {
		t.Fatal(err)
	}
	missing = bytes.Clone(missing)
	last, err := enc.Reset().Bytes("S", []byte("end")).Flag("F").Encode()
	if err != nil {
		t.Fatal(err)
	}
	last = bytes.Clone(last)

	hits := make([]sHit, asmChunk+5)
	want := make([][]byte, len(hits))
	for i := range hits {
		rec, val := int64(100), first
		if i >= asmChunk+1 {
			rec, val = 101, missing
		}
		if i >= asmChunk+3 {
			rec, val = 102, last
		}
		hits[i] = sHit{rec: rec, dst: i}
		want[i] = val
	}
	for _, threads := range []int{1, 3} {
		out := make([][]byte, len(hits))
		if err := c.processBin(ld, hitBinGroup{hits: hits}, out, threads); err != nil {
			t.Fatal(err)
		}
		for i := range out {
			if out[i] == nil || !bytes.Equal(out[i], want[i]) {
				t.Fatalf("threads %d hit %d: got %x, want %x", threads, i, out[i], want[i])
			}
			if cap(out[i]) != len(out[i]) {
				t.Fatalf("hit %d: cap %d != len %d", i, cap(out[i]), len(out[i]))
			}
		}
		// decoded 值流不归结果所有；下一次调用以及 append 都不能改旧结果。
		saved := bytes.Clone(out[0])
		_ = append(out[0], 0xff)
		if !bytes.Equal(out[1], saved) {
			t.Fatal("appending to a result changed a duplicate result")
		}
		again := make([][]byte, len(hits))
		if err := c.processBin(ld, hitBinGroup{hits: hits}, again, threads); err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(out[0], saved) {
			t.Fatal("a later query changed the earlier result")
		}
	}
}

func TestColumnarProjectionSelectedColumnsDuplicates(t *testing.T) {
	lib, schema, keys := buildColTestLib(t)
	db, err := OpenStreaming(lib, StreamOptions{Threads: 2, Columns: []int{0, 3, 5}})
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	db.forceCol = true
	query := [][]byte{keys[0], keys[0], keys[1], keys[2], keys[2], keys[3]}
	out, err := db.GetBatch(query)
	if err != nil {
		t.Fatal(err)
	}
	want := []string{"AF=0.1;DB;AFR=0.01,.", "AF=0.1;DB;AFR=0.01,.", "AF=0.2;DB;AFR=0.01,.", ".", ".", "."}
	for i := range out {
		text, err := RenderRecord(schema, out[i])
		if err != nil || string(text) != want[i] {
			t.Fatalf("hit %d: got %q, %v, want %q", i, text, err, want[i])
		}
		if out[i] == nil || cap(out[i]) != len(out[i]) {
			t.Fatalf("hit %d: nil result or capacity exceeds length", i)
		}
	}
}

func TestColumnarProjectionValidTail(t *testing.T) {
	schema := Schema{{Name: "I", Type: TypeInteger, Number: "1"}}
	c, ld := colProjectionBin(t, schema, 3, [][]byte{{0x05}}, [][]byte{{2, 6}})
	for _, rec := range []int64{100, 101, 102} {
		out := make([][]byte, 1)
		if err := c.processBin(ld, hitBinGroup{hits: []sHit{{rec: rec}}}, out, 1); err != nil {
			t.Fatalf("record %d: %v", rec, err)
		}
		view, err := NewRecord(schema, out[0])
		if err != nil {
			t.Fatal(err)
		}
		if view.HasValue("I") != (rec != 101) || rec != 101 && view.Int("I") != rec-99 {
			t.Fatalf("record %d: unexpected projected value %x", rec, out[0])
		}
	}
}

func TestColumnarProjectionRejectsMalformedTail(t *testing.T) {
	floats := binary.LittleEndian.AppendUint64(nil, math.Float64bits(1))
	floats = binary.LittleEndian.AppendUint64(floats, math.Float64bits(2))
	floats = binary.LittleEndian.AppendUint64(floats, math.Float64bits(math.NaN()))
	for _, tc := range []struct {
		name    string
		field   Field
		payload []byte
		wantErr string
	}{
		{"missing value", Field{"I", TypeInteger, "1"}, []byte{2, 4}, "记录 102"},
		{"truncated integer", Field{"I", TypeInteger, "1"}, []byte{2, 4, 0x80}, "记录 102"},
		{"overflow integer", Field{"I", TypeInteger, "1"}, append([]byte{2, 4}, bytes.Repeat([]byte{0xff}, 10)...), "记录 102"},
		{"extra value", Field{"I", TypeInteger, "1"}, []byte{2, 4, 6, 8}, "值流尾部"},
		{"trailing bytes", Field{"I", TypeInteger, "1"}, []byte{2, 4, 6, 0x80}, "值流尾部"},
		{"truncated string", Field{"S", TypeString, "1"}, []byte{1, 'a', 1, 'b', 2, 'c'}, "记录 102"},
		{"nonfinite float", Field{"F", TypeFloat, "1"}, floats, "NaN"},
		{"truncated float", Field{"F", TypeFloat, "1"}, floats[:len(floats)-1], "记录 102"},
		{"bad array element", Field{"A", TypeInteger, "R"}, []byte{1, 1, 2, 1, 1, 4, 1, 2, 6, 0}, "记录 102"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c, ld := colProjectionBin(t, Schema{tc.field}, 3, [][]byte{{0x07}}, [][]byte{tc.payload})
			// 最后命中在首条记录，损坏处全部位于命中之后，zstd 帧本身完整。
			err := c.processBin(ld, hitBinGroup{hits: []sHit{{rec: 100}}}, make([][]byte, 1), 1)
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("got %v, want error containing %q", err, tc.wantErr)
			}
		})
	}
}

func TestColumnarProjectionBitmapPadding(t *testing.T) {
	schema := Schema{{Name: "F", Type: TypeFlag, Number: "0"}}
	for _, tc := range []struct {
		name    string
		count   uint32
		bitmap  []byte
		wantErr bool
	}{
		{"one valid bit", 1, []byte{0x01}, false},
		{"byte boundary", 8, []byte{0x81}, false},
		{"valid partial byte", 9, []byte{0x01, 0x01}, false},
		{"padding replaces value", 3, []byte{0x81}, true},
		{"second byte padding", 9, []byte{0x01, 0x02}, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			// present_count 包含坏 padding 位但不超过 count，单靠 popcount
			// 比较无法拒绝，Flag 也没有 payload 可用于间接检查。
			c, ld := colProjectionBin(t, schema, tc.count, [][]byte{tc.bitmap}, [][]byte{nil})
			err := c.processBin(ld, hitBinGroup{hits: []sHit{{rec: 100}}}, make([][]byte, 1), 1)
			if tc.wantErr {
				if err == nil || !strings.Contains(err.Error(), "填充位") {
					t.Fatalf("got %v, want bitmap padding error", err)
				}
			} else if err != nil {
				t.Fatal(err)
			}
		})
	}
}
