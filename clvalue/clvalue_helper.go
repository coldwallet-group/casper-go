package clvalue

import (
	"github/casper-go/common/byteutil"
)

func ToBytesHelper(cltype CLType) []byte {
	switch c := cltype.(type) {
	case *CLT:
		return []byte{uint8(*c)}
	case NumberCoderType:
		return []byte{8}
	case *OptionType:
		return byteutil.Concat([]byte{uint8(c.Tag)}, ToBytesHelper(c.InnerType))
	case *ByteArrayType:
		return byteutil.Concat([]byte{uint8(c.Tag)}, ToBytesU32(c.Size))
	}
	panic("wrong type")
}
