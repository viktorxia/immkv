// schema.go 定义 v003 的 schema：有序字段表、字段类型、Number 校验与
// schema hash。schema 的权威副本落在 base.manifest；这里只提供类型和
// 纯函数，不碰盘上读写。
package immkv

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// FieldType 是一个 schema 字段的物理类型。
type FieldType uint8

const (
	// TypeInteger 是 int64，盘上 zigzag uvarint。
	TypeInteger FieldType = iota + 1
	// TypeFloat 是 float64，盘上 8 字节 IEEE 754 小端。
	TypeFloat
	// TypeString 是原字节，盘上 uvarint(len) + bytes。
	TypeString
	// TypeFlag 没有 payload，只占 record 位图的一位。
	TypeFlag
)

// String 返回 manifest 里的类型名。
func (t FieldType) String() string {
	switch t {
	case TypeInteger:
		return "Integer"
	case TypeFloat:
		return "Float"
	case TypeString:
		return "String"
	case TypeFlag:
		return "Flag"
	default:
		return fmt.Sprintf("FieldType(%d)", uint8(t))
	}
}

func parseFieldType(s string) (FieldType, bool) {
	switch s {
	case "Integer":
		return TypeInteger, true
	case "Float":
		return TypeFloat, true
	case "String":
		return TypeString, true
	case "Flag":
		return TypeFlag, true
	default:
		return 0, false
	}
}

// Field 是 schema 的一项；Name 已加过 prefix，Number 是 VCF 的原文。
type Field struct {
	Name   string
	Type   FieldType
	Number string
}

// Schema 是有序字段表；顺序即 record 位图的位号。
type Schema []Field

// Validate 校验字段表。空 schema 合法（对应零字段 record），但 manifest
// 会另外要求至少一个字段。
func (s Schema) Validate() error {
	seen := make(map[string]bool, len(s))
	for i := range s {
		f := &s[i]
		if f.Name == "" {
			return fmt.Errorf("immkv: schema 第 %d 个字段名为空", i)
		}
		if seen[f.Name] {
			return fmt.Errorf("immkv: schema 字段 %q 重复", f.Name)
		}
		seen[f.Name] = true

		switch f.Type {
		case TypeFlag:
			if f.Number != "0" {
				return fmt.Errorf("immkv: Flag 字段 %q 的 Number 必须是 0，得到 %q", f.Name, f.Number)
			}
		case TypeInteger, TypeFloat, TypeString:
			switch f.Number {
			case "1", "A", "R", "G", ".":
			case "0":
				return fmt.Errorf("immkv: 字段 %q 的 Number=0 只能是 Flag", f.Name)
			default:
				return fmt.Errorf("immkv: 字段 %q 的 Number %q 不合法", f.Name, f.Number)
			}
		default:
			return fmt.Errorf("immkv: 字段 %q 的类型 %d 不合法", f.Name, f.Type)
		}
	}
	return nil
}

// Index 返回字段名的位号；不存在返回 -1。
func (s Schema) Index(name string) int {
	for i := range s {
		if s[i].Name == name {
			return i
		}
	}
	return -1
}

// Hash 按 schema 顺序对 "name\tType\tNumber\n" 取 SHA-256 前 8 字节。
func (s Schema) Hash() [8]byte {
	h := sha256.New()
	for i := range s {
		fmt.Fprintf(h, "%s\t%s\t%s\n", s[i].Name, s[i].Type.String(), s[i].Number)
	}
	var out [8]byte
	copy(out[:], h.Sum(nil))
	return out
}

// HashHex 返回 Hash 的十六进制形态，manifest 里存这个。
func (s Schema) HashHex() string {
	h := s.Hash()
	return hex.EncodeToString(h[:])
}

// isArray 表示该字段在 record 里是否按数组编码。只有 Number 为 R、G、.
// 的数值字段是数组；String 永远是原字节，A 已按 ALT 下标取成标量。
func (f Field) isArray() bool {
	if f.Type != TypeInteger && f.Type != TypeFloat {
		return false
	}
	switch f.Number {
	case "R", "G", ".":
		return true
	}
	return false
}
