// manifest.go 是 v003 的 base.manifest 读写层。manifest 是 schema 的权威
// 副本，也是 vals/col/keys 三个文件与记录数的配对凭据；旧 v001 manifest
// 直接拒绝，不做双解码。
package immkv

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	manifestFormat  = "immkv-v003"
	manifestVersion = 3
)

// ManifestField 是 manifest 里的一项 schema 字段；Column 必须等于下标。
type ManifestField struct {
	Name   string `json:"name"`
	Column int    `json:"column"`
	Type   string `json:"type"`
	Number string `json:"number"`
}

// SourceInfoOrder 是构建期源 INFO 顺序的统计，只作诊断。
type SourceInfoOrder struct {
	OK  int64 `json:"ok"`
	Bad int64 `json:"bad"`
}

// Manifest 是 base.manifest 的 JSON 形态。
type Manifest struct {
	Format          string          `json:"format"`
	Version         int             `json:"version"`
	Base            string          `json:"base"`
	BuildID         string          `json:"build_id"`
	BinWidth        uint32          `json:"bin_width"`
	RecordCount     uint64          `json:"record_count"`
	KeysBytes       int64           `json:"keys_bytes"`
	ValsBytes       int64           `json:"vals_bytes"`
	ColBytes        int64           `json:"col_bytes"`
	SchemaHash      string          `json:"schema_hash"`
	Fields          []ManifestField `json:"fields"`
	SourceInfoOrder SourceInfoOrder `json:"source_info_order"`
	// NonFiniteFloats 是构建期遇到 NaN/Inf 并按缺失处理的 Float 次数；只作诊断。
	NonFiniteFloats int64 `json:"non_finite_floats"`
}

// NewManifest 建一份 v003 manifest 骨架；Fields、Column 与 schema hash 由
// SetSchema 写入。Base 是文件名前缀（如 chr1），不是路径。
func NewManifest(base string, fields []Field) (*Manifest, error) {
	if len(fields) == 0 {
		return nil, fmt.Errorf("immkv: manifest 没有字段")
	}
	m := &Manifest{Base: base}
	if err := m.SetSchema(fields); err != nil {
		return nil, err
	}
	return m, nil
}

// SetSchema 用字段表重写 manifest 的 Fields、Column 与 schema hash。
func (m *Manifest) SetSchema(fields []Field) error {
	s := Schema(fields)
	if err := s.Validate(); err != nil {
		return err
	}
	m.Format = manifestFormat
	m.Version = manifestVersion
	m.Fields = make([]ManifestField, len(fields))
	for i := range fields {
		m.Fields[i] = ManifestField{
			Name:   fields[i].Name,
			Column: i,
			Type:   fields[i].Type.String(),
			Number: fields[i].Number,
		}
	}
	m.SchemaHash = s.HashHex()
	return nil
}

// Schema 把 manifest 的字段条目还原成 schema，并校验 Column/Type/Number。
func (m *Manifest) Schema() (Schema, error) {
	s := make(Schema, len(m.Fields))
	for i := range m.Fields {
		mf := &m.Fields[i]
		if mf.Column != i {
			return nil, fmt.Errorf("immkv: manifest 字段 %q 的 column %d ≠ %d", mf.Name, mf.Column, i)
		}
		t, ok := parseFieldType(mf.Type)
		if !ok {
			return nil, fmt.Errorf("immkv: manifest 字段 %q 的类型 %q 未知", mf.Name, mf.Type)
		}
		s[i] = Field{Name: mf.Name, Type: t, Number: mf.Number}
	}
	if err := s.Validate(); err != nil {
		return nil, err
	}
	return s, nil
}

// Validate 校验 manifest 的格式、schema hash、配对字段与字节数。
func (m *Manifest) Validate() error {
	if m.Format != manifestFormat || m.Version != manifestVersion {
		return fmt.Errorf("immkv: manifest format/version 不合法")
	}
	if m.Base == "" || m.Base == "." || m.Base == ".." || filepath.Base(m.Base) != m.Base {
		return fmt.Errorf("immkv: manifest base %q 不合法", m.Base)
	}
	if len(m.Fields) == 0 {
		return fmt.Errorf("immkv: manifest 没有字段")
	}
	if len(m.BuildID) != 32 {
		return fmt.Errorf("immkv: manifest build_id %q 长度不是 32", m.BuildID)
	}
	if _, err := hex.DecodeString(m.BuildID); err != nil {
		return fmt.Errorf("immkv: manifest build_id %q 不是十六进制", m.BuildID)
	}
	s, err := m.Schema()
	if err != nil {
		return err
	}
	if len(m.SchemaHash) != 16 {
		return fmt.Errorf("immkv: manifest schema_hash %q 长度不是 16", m.SchemaHash)
	}
	if _, err := hex.DecodeString(m.SchemaHash); err != nil {
		return fmt.Errorf("immkv: manifest schema_hash %q 不是十六进制", m.SchemaHash)
	}
	if !strings.EqualFold(m.SchemaHash, s.HashHex()) {
		return fmt.Errorf("immkv: manifest schema hash 不匹配")
	}
	if m.KeysBytes < 0 || m.ValsBytes < 0 || m.ColBytes < 0 {
		return fmt.Errorf("immkv: manifest 文件字节数为负")
	}
	if m.SourceInfoOrder.OK < 0 || m.SourceInfoOrder.Bad < 0 {
		return fmt.Errorf("immkv: manifest source_info_order 为负")
	}
	if m.NonFiniteFloats < 0 {
		return fmt.Errorf("immkv: manifest non_finite_floats 为负")
	}
	return nil
}

// ReadManifest 读 base.manifest 并校验；旧 v001/v002 manifest 直接拒绝。
func ReadManifest(path string) (*Manifest, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var m Manifest
	if err := json.Unmarshal(raw, &m); err != nil {
		return nil, fmt.Errorf("immkv: %s: 解析 manifest: %w", path, err)
	}
	if m.Format != manifestFormat || m.Version != manifestVersion {
		return nil, fmt.Errorf("immkv: %s: v003 不兼容旧格式，请重建", path)
	}
	if err := m.Validate(); err != nil {
		return nil, fmt.Errorf("immkv: %s: %w", path, err)
	}
	return &m, nil
}

// ReadSchema 读库目录里任意一个 shard 的 manifest，返回整库 schema。
// 所有 shard 的 schema 必须一致；一个 manifest 都找不到时报错。
// 需要 shard 级诊断信息（build id、记录数、字节数）时用 ReadManifest。
func ReadSchema(dir string) (Schema, error) {
	if dir == "" {
		return nil, fmt.Errorf("immkv: ReadSchema: empty dir")
	}
	var (
		schema Schema
		hash   string
		first  string
	)
	for id := uint8(0); id <= 24; id++ {
		path := filepath.Join(dir, "chr"+ChromName(id)+".manifest")
		m, err := ReadManifest(path)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, err
		}
		if first == "" {
			s, err := m.Schema()
			if err != nil {
				return nil, err
			}
			schema, hash, first = s, m.SchemaHash, path
			continue
		}
		if !strings.EqualFold(hash, m.SchemaHash) {
			return nil, fmt.Errorf("immkv: %s 与 %s 的 schema 不一致", first, path)
		}
	}
	if first == "" {
		return nil, fmt.Errorf("immkv: %s 下没有 chr*.manifest", dir)
	}
	return schema, nil
}

// WriteManifest 校验后原子写 base.manifest。
func WriteManifest(path string, m *Manifest) error {
	if err := m.Validate(); err != nil {
		return err
	}
	raw, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return err
	}
	raw = append(raw, '\n')
	return writeFileAtomic(path, raw)
}
