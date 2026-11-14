package immkv

import (
	"encoding/binary"
	"fmt"
	"math/bits"
)

// recordProjection 在打开句柄时编译字段选择；查询时只遍历有值字段，连续
// 选中字段的 payload 一次拷入调用自己的 slab，不改写输入或共享缓冲。
type recordProjection struct {
	schema Schema
	kinds  []FieldType
	mask   []uint64
}

const projectionArray FieldType = 0

func newRecordProjection(schema Schema, selected []int) *recordProjection {
	p := &recordProjection{
		schema: schema,
		kinds:  make([]FieldType, len(schema)),
		mask:   make([]uint64, (len(schema)+63)/64),
	}
	for i, f := range schema {
		p.kinds[i] = f.Type
		if f.isArray() {
			p.kinds[i] = projectionArray
		}
	}
	for _, i := range selected {
		p.mask[i>>6] |= uint64(1) << uint(i&63)
	}
	return p
}

// appendRecord 追加一条投影后的 record。即使没有选中字段有值，也保留完整
// 位图；未选字段仍校验，避免把损坏 payload 当成正常缺失。
func (p *recordProjection) appendRecord(dst, raw []byte) ([]byte, error) {
	bl := len(p.mask) * 8
	if len(raw) < bl {
		return nil, fmt.Errorf("immkv: record 位图长度 %d < %d", len(raw), bl)
	}
	base := len(dst)
	dst = append(dst, make([]byte, bl)...)
	off, run := bl, -1
	for w, mask := range p.mask {
		present := binary.LittleEndian.Uint64(raw[w*8:])
		if n := len(p.schema) - w*64; n < 64 {
			present &= uint64(1)<<uint(n) - 1
		}
		binary.LittleEndian.PutUint64(dst[base+w*8:], present&mask)
		for present != 0 {
			bit := bits.TrailingZeros64(present)
			present &= present - 1
			i := w*64 + bit
			if mask&(uint64(1)<<uint(bit)) != 0 {
				if run < 0 {
					run = off
				}
			} else if run >= 0 {
				dst = append(dst, raw[run:off]...)
				run = -1
			}
			var err error
			switch p.kinds[i] {
			case TypeFlag:
			case TypeInteger:
				_, n := binary.Uvarint(raw[off:])
				if n <= 0 {
					err = fmt.Errorf("integer uvarint 截断或溢出")
				} else {
					off += n
				}
			case TypeFloat:
				if len(raw)-off < 8 {
					err = fmt.Errorf("float 需要 8 字节，剩 %d", len(raw)-off)
				} else if binary.LittleEndian.Uint64(raw[off:])&0x7ff0000000000000 == 0x7ff0000000000000 {
					err = fmt.Errorf("float 是 NaN 或 Inf")
				} else {
					off += 8
				}
			case TypeString:
				l, n := binary.Uvarint(raw[off:])
				if n <= 0 {
					err = fmt.Errorf("uvarint 截断或溢出")
				} else if l > uint64(len(raw)-off-n) {
					err = fmt.Errorf("长度 %d 超过剩余 %d 字节", l, len(raw)-off-n)
				} else {
					off += n + int(l)
				}
			default:
				var rest []byte
				rest, err = skipValue(p.schema[i], raw[off:])
				off = len(raw) - len(rest)
			}
			if err != nil {
				return nil, fmt.Errorf("immkv: 字段 %q: %w", p.schema[i].Name, err)
			}
		}
	}
	if off != len(raw) {
		return nil, fmt.Errorf("immkv: record 尾部有 %d 字节多余数据", len(raw)-off)
	}
	if run >= 0 {
		dst = append(dst, raw[run:off]...)
	}
	return dst, nil
}
