package clvalue

import (
	"math/big"
	"strconv"
)

type NumberCoder struct {
	BitSize uint32
	Signed  bool
	Tag     uint8
	Val     *big.Int
	Name    string
}

func NewNumberCoder(bitSize uint32, tag uint8, signed bool, val interface{}) *NumberCoder {
	n := "U"
	if signed {
		n = "I"
	}
	return &NumberCoder{
		BitSize: bitSize,
		Tag:     tag,
		Signed:  signed,
		Val:     NumberFrom(val),
		Name:    n + strconv.Itoa(int(bitSize)),
	}
}

func (nc *NumberCoder) GetCLType() CLType {
	return NumberCoderType(nc.Name)
}

func (nc *NumberCoder) ToBytes() []byte {
	return toBytes(nc.Val, nc.BitSize)
}

func toBytes(val *big.Int, bitSize uint32) []byte {
	return toByteNumber(bitSize, false, val.Bytes())
}

func byteReverse(s *[]byte) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

func NumberFrom(value interface{}) *big.Int {
	switch v := value.(type) {
	case big.Int:
		return &v
	case *big.Int:
		return v
	case int8:
		return big.NewInt(int64(v))
	case int16:
		return big.NewInt(int64(v))
	case int32:
		return big.NewInt(int64(v))
	case uint32:
		return big.NewInt(int64(v))
	case uint64:
		return big.NewInt(int64(v))
	case int64:
		return big.NewInt(v)
	}
	return nil
}

type NumberCoderType string
