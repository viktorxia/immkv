package immkv

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func v003Schema() Schema {
	return Schema{
		{Name: "AC", Type: TypeInteger, Number: "A"},
		{Name: "AN", Type: TypeInteger, Number: "1"},
		{Name: "AF", Type: TypeFloat, Number: "A"},
		{Name: "DB", Type: TypeFlag, Number: "0"},
		{Name: "STR", Type: TypeString, Number: "."},
	}
}

func TestSchemaValidate(t *testing.T) {
	cases := []struct {
		name    string
		schema  Schema
		wantErr bool
	}{
		{"valid", v003Schema(), false},
		{"empty", nil, false},
		{"duplicate", Schema{
			{Name: "A", Type: TypeInteger, Number: "1"},
			{Name: "A", Type: TypeFloat, Number: "1"},
		}, true},
		{"empty name", Schema{{Name: "", Type: TypeInteger, Number: "1"}}, true},
		{"unknown type", Schema{{Name: "A", Type: 0, Number: "1"}}, true},
		{"flag number", Schema{{Name: "F", Type: TypeFlag, Number: "1"}}, true},
		{"number zero non-flag", Schema{{Name: "I", Type: TypeInteger, Number: "0"}}, true},
		{"bad number", Schema{{Name: "I", Type: TypeInteger, Number: "X"}}, true},
		{"string R", Schema{{Name: "S", Type: TypeString, Number: "R"}}, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			err := c.schema.Validate()
			if (err != nil) != c.wantErr {
				t.Fatalf("Validate() error = %v, wantErr %v", err, c.wantErr)
			}
		})
	}
}

func TestSchemaHash(t *testing.T) {
	s := Schema{
		{Name: "AC", Type: TypeInteger, Number: "A"},
		{Name: "AF", Type: TypeFloat, Number: "A"},
	}
	sum := sha256.Sum256([]byte("AC\tInteger\tA\nAF\tFloat\tA\n"))
	want := hex.EncodeToString(sum[:8])
	if got := s.HashHex(); got != want {
		t.Fatalf("HashHex() = %q, want %q", got, want)
	}
	if len(s.HashHex()) != 16 {
		t.Fatalf("HashHex length = %d, want 16", len(s.HashHex()))
	}

	swapped := Schema{s[1], s[0]}
	if swapped.HashHex() == want {
		t.Fatal("field order should change the schema hash")
	}
	typeChanged := Schema{s[0], s[1]}
	typeChanged[0].Type = TypeFloat
	if typeChanged.HashHex() == want {
		t.Fatal("field type should change the schema hash")
	}
	numberChanged := Schema{s[0], s[1]}
	numberChanged[0].Number = "1"
	if numberChanged.HashHex() == want {
		t.Fatal("field Number should change the schema hash")
	}
}

func TestSchemaIndex(t *testing.T) {
	s := v003Schema()
	for i := range s {
		if got := s.Index(s[i].Name); got != i {
			t.Fatalf("Index(%q) = %d, want %d", s[i].Name, got, i)
		}
	}
	if got := s.Index("NOPE"); got != -1 {
		t.Fatalf("Index(NOPE) = %d, want -1", got)
	}
}

func TestFieldIsArray(t *testing.T) {
	cases := []struct {
		field Field
		want  bool
	}{
		{Field{Type: TypeInteger, Number: "1"}, false},
		{Field{Type: TypeInteger, Number: "A"}, false},
		{Field{Type: TypeInteger, Number: "R"}, true},
		{Field{Type: TypeFloat, Number: "G"}, true},
		{Field{Type: TypeFloat, Number: "."}, true},
		{Field{Type: TypeString, Number: "R"}, false},
		{Field{Type: TypeFlag, Number: "0"}, false},
	}
	for _, c := range cases {
		if got := c.field.isArray(); got != c.want {
			t.Fatalf("%+v isArray() = %v, want %v", c.field, got, c.want)
		}
	}
}
