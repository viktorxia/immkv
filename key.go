// key.go 是 immkv 的键编码：染色体号 + 位置 + 碱基，保证字节序 = 基因组自然序。
// 编码与 immdb 时期逐字节一致（既有键集文件与全部对照数字都依赖它）。

package immkv

import (
	"encoding/binary"
	"strconv"
	"strings"
)

// ChromID 把 VCF/TSV 里的染色体名映射为键首字节：1..22 → 0..21、X → 22、
// Y → 23、M/MT → 24（带不带 chr 前缀都认）。首字节序 = 染色体自然序。
// 其余 contig（chrUn_*、*_alt、随机片段等）返回 false，调用方自行跳过。
func ChromID(s string) (uint8, bool) {
	name := strings.TrimPrefix(strings.TrimPrefix(s, "chr"), "CHR")
	if n, err := strconv.ParseUint(name, 10, 8); err == nil {
		if n >= 1 && n <= 22 {
			return uint8(n - 1), true
		}
		return 0, false
	}
	switch strings.ToUpper(name) {
	case "X":
		return 22, true
	case "Y":
		return 23, true
	case "M", "MT":
		return 24, true
	}
	return 0, false
}

// ChromName 是 ChromID 的反函数：chromID → 染色体名（文件名与 VCF 列共用）。
func ChromName(id uint8) string {
	switch {
	case id <= 21:
		return strconv.Itoa(int(id) + 1)
	case id == 22:
		return "X"
	case id == 23:
		return "Y"
	case id == 24:
		return "M"
	}
	return "?"
}

// VarKey 编码排序键：[chromID][pos 大端 u32][ref][0x00][alt]。
func VarKey(c uint8, p uint32, ref, alt string) []byte {
	k := make([]byte, 0, 7+len(ref)+1+len(alt))
	k = append(k, c)
	var b [4]byte
	binary.BigEndian.PutUint32(b[:], p)
	k = append(k, b[:]...)
	k = append(k, ref...)
	k = append(k, 0)
	return append(k, alt...)
}

// PosKey 编码位置前缀：[chromID][pos 大端 u32]。
func PosKey(c uint8, p uint32) []byte {
	k := make([]byte, 5)
	k[0] = c
	binary.BigEndian.PutUint32(k[1:], p)
	return k
}
