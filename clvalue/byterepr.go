package clvalue

import (
	"bytes"
	"github/casper-go/common/byteutil"
	"math/big"
)

func toByteNumber(bitSize uint32, signed bool, vb []byte) []byte {
	var val big.Int
	val.SetBytes(vb)
	if val.Cmp(big.NewInt(0)) >= 0 {
		// for positive number, we had to deal with paddings
		if bitSize > 64 {
			// for u128, u256, u512, we have to and append extra byte for length
			r := byteutil.Concat(vb, []byte{byte(len(vb))})
			byteReverse(&r)
			return r
		} else {
			// for other types, we have to add padding 0s
			byteLen := bitSize / 8
			var b []byte
			for i := 0; i < int(byteLen)-len(vb); i++ {
				b = append(b, 0)
			}
			byteReverse(&vb)
			return byteutil.Concat(vb, b[:])
		}
	}
	byteReverse(&vb)
	return vb
}

//Serializes an array of u8, equal to Vec<u8> in rust.
func ToBytesArrayU8(v []byte) []byte {
	return byteutil.Concat(ToBytesU32(uint32(len(v))), v)
}

func ToBytesU32(u32 uint32) []byte {
	return toByteNumber(32, false, NumberFrom(u32).Bytes())
}

func ToBytesU64(u64 uint64) []byte {
	return toByteNumber(64, false, NumberFrom(u64).Bytes())
}

func ToBytesBytesArray(arr []byte) []byte {
	return arr
}

func ToBytesString(val string) []byte {
	bs := []byte(val)
	return byteutil.Concat(ToBytesU32(uint32(len(bs))), bs)
}

func ToByteSlice(list [][]byte) []byte {
	listLen := len(list)
	var bs [][]byte
	bs = append(bs, ToBytesU32(uint32(listLen)))

	for i := 0; i < listLen; i++ {
		bs = append(bs, list[i])
	}
	return bytes.Join(bs, []byte{})
}
