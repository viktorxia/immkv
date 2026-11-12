package immkv

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func v003Manifest(t *testing.T) *Manifest {
	t.Helper()
	m, err := NewManifest("chr1", v003Schema())
	if err != nil {
		t.Fatal(err)
	}
	m.BuildID = strings.Repeat("ab", 16)
	m.BinWidth = 1_000_000
	m.RecordCount = 7
	m.KeysBytes = 100
	m.ValsBytes = 200
	m.ColBytes = 300
	m.SourceInfoOrder = SourceInfoOrder{OK: 6, Bad: 1}
	if err := m.Validate(); err != nil {
		t.Fatal(err)
	}
	return m
}

func TestManifestRoundTrip(t *testing.T) {
	m := v003Manifest(t)
	path := filepath.Join(t.TempDir(), "chr1.manifest")
	if err := WriteManifest(path, m); err != nil {
		t.Fatal(err)
	}
	got, err := ReadManifest(path)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, m) {
		t.Fatalf("manifest round trip mismatch:\n got %+v\nwant %+v", got, m)
	}
	s, err := got.Schema()
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(s, v003Schema()) {
		t.Fatalf("schema = %+v, want %+v", s, v003Schema())
	}
}

func TestReadSchema(t *testing.T) {
	dir := t.TempDir()
	for _, base := range []string{"chr7", "chrM"} {
		m := v003Manifest(t)
		m.Base = base
		if err := WriteManifest(filepath.Join(dir, base+".manifest"), m); err != nil {
			t.Fatal(err)
		}
	}
	got, err := ReadSchema(dir)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, v003Schema()) {
		t.Fatalf("schema = %+v, want %+v", got, v003Schema())
	}
}

func TestReadSchemaNoManifest(t *testing.T) {
	if _, err := ReadSchema(t.TempDir()); err == nil {
		t.Fatal("ReadSchema() = nil error, want no-manifest error")
	}
}

func TestReadSchemaMismatch(t *testing.T) {
	dir := t.TempDir()
	m1 := v003Manifest(t)
	m1.Base = "chr1"
	if err := WriteManifest(filepath.Join(dir, "chr1.manifest"), m1); err != nil {
		t.Fatal(err)
	}
	m2 := v003Manifest(t)
	m2.Base = "chr2"
	if err := m2.SetSchema(Schema{{Name: "AF", Type: TypeFloat, Number: "A"}}); err != nil {
		t.Fatal(err)
	}
	if err := WriteManifest(filepath.Join(dir, "chr2.manifest"), m2); err != nil {
		t.Fatal(err)
	}
	_, err := ReadSchema(dir)
	if err == nil || !strings.Contains(err.Error(), "schema 不一致") {
		t.Fatalf("ReadSchema error = %v, want schema mismatch", err)
	}
}

func TestManifestRejectsOldFormat(t *testing.T) {
	path := filepath.Join(t.TempDir(), "chr1.manifest")
	old := `{"format":"immkv-col-manifest","version":1,"base":"chr1","fields":[{"name":"AF","column":0,"type":"Float","number":"A","encoding":"raw"}]}`
	if err := os.WriteFile(path, []byte(old), 0o644); err != nil {
		t.Fatal(err)
	}
	_, err := ReadManifest(path)
	if err == nil || !strings.Contains(err.Error(), "v003 不兼容旧格式") {
		t.Fatalf("ReadManifest error = %v, want old-format rejection", err)
	}
}

func TestManifestValidation(t *testing.T) {
	cases := []struct {
		name   string
		mutate func(*Manifest)
	}{
		{"schema hash", func(m *Manifest) { m.SchemaHash = strings.Repeat("0", 16) }},
		{"column", func(m *Manifest) { m.Fields[0].Column = 1 }},
		{"duplicate", func(m *Manifest) { m.Fields[1].Name = m.Fields[0].Name }},
		{"build id length", func(m *Manifest) { m.BuildID = "abc" }},
		{"build id hex", func(m *Manifest) { m.BuildID = strings.Repeat("zz", 16) }},
		{"schema hash length", func(m *Manifest) { m.SchemaHash = "abc" }},
		{"schema hash hex", func(m *Manifest) { m.SchemaHash = strings.Repeat("zz", 8) }},
		{"negative bytes", func(m *Manifest) { m.KeysBytes = -1 }},
		{"negative stats", func(m *Manifest) { m.SourceInfoOrder.Bad = -1 }},
		{"negative non-finite", func(m *Manifest) { m.NonFiniteFloats = -1 }},
		{"empty fields", func(m *Manifest) { m.Fields = nil }},
		{"bad base", func(m *Manifest) { m.Base = "dir/chr1" }},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			m := v003Manifest(t)
			c.mutate(m)
			if err := m.Validate(); err == nil {
				t.Fatal("Validate() = nil error, want error")
			}
		})
	}
}

func TestManifestReadInvalidJSON(t *testing.T) {
	path := filepath.Join(t.TempDir(), "chr1.manifest")
	if err := os.WriteFile(path, []byte("{"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := ReadManifest(path); err == nil {
		t.Fatal("ReadManifest() = nil error, want JSON error")
	}
}

func TestManifestWriteRejectsInvalid(t *testing.T) {
	m := v003Manifest(t)
	m.SchemaHash = "bad"
	path := filepath.Join(t.TempDir(), "chr1.manifest")
	if err := WriteManifest(path, m); err == nil {
		t.Fatal("WriteManifest() = nil error, want error")
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("manifest file exists after rejected write: %v", err)
	}
}

func TestManifestSetSchema(t *testing.T) {
	m := &Manifest{}
	if err := m.SetSchema(v003Schema()); err != nil {
		t.Fatal(err)
	}
	if m.Format != manifestFormat || m.Version != manifestVersion {
		t.Fatalf("format/version = %q/%d", m.Format, m.Version)
	}
	if m.SchemaHash != v003Schema().HashHex() {
		t.Fatalf("schema hash = %q, want %q", m.SchemaHash, v003Schema().HashHex())
	}
	for i := range m.Fields {
		if m.Fields[i].Column != i {
			t.Fatalf("field %d column = %d", i, m.Fields[i].Column)
		}
	}
	if _, err := NewManifest("chr1", Schema{{Name: "F", Type: TypeFlag, Number: "1"}}); err == nil {
		t.Fatal("NewManifest() with invalid schema = nil error, want error")
	}
	if _, err := NewManifest("chr1", nil); err == nil {
		t.Fatal("NewManifest() with no fields = nil error, want error")
	}
}
