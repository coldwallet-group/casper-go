package clvalue

import (
	"github/casper-go/common/byteutil"
)

func ToBytesHelper(cltype CLType) []byte {
	switch c := cltype.(type) {
	case *CLT:
		return []byte{uint8(*c)}
	case NumberCoderType:
		return []byte{uint8(numberCoderTransfer(c))}
	case *OptionType:
		return byteutil.Concat([]byte{uint8(c.Tag)}, ToBytesHelper(c.InnerType))
	case *ByteArrayType:
		return byteutil.Concat([]byte{uint8(c.Tag)}, ToBytesU32(c.Size))
	}
	panic("wrong type")
}

func numberCoderTransfer(n NumberCoderType) CLT {
	prefix := string(n)[:1]
	v := string(n)[1:]
	if prefix == "U" {
		switch v {
		case "64":
			return TagU64
		case "512":
			return TagU512
		}
	} else if prefix == "I" {
		//TODO 有符号整型暂时未使用
	}
	panic("wrong numberCoder type")
}
