package clvalue

import (
	"math/big"
	"strconv"
)

const (
	hexPrefix = "0x"
)

//type BigNumber struct {
//	hex         string
//	isBigNumber bool
//}
//
//func BigNumberFrom(Value interface{}) (*big.Int, error) {
//	switch v := Value.(type) {
//	case big.Int:
//	case *big.Int:
//		return NewBigNumber(hex.EncodeToString(v.Bytes())), nil
//	case string:
//		if hexutil.Has0xPrefix(v) {
//			v = v[:2]
//		}
//		if hexutil.IsHex(v) {
//			return NewBigNumber(v), nil
//		}
//		if numutil.IsNum(v) {
//
//			n, ok := new(big.Int).SetString(v, 10)
//			if !ok {
//				return nil, errors.New("failed to new BigNumber from string")
//			}
//			return NewBigNumber(hex.EncodeToString(n.Bytes())), nil
//		}
//		return nil, errors.New("invalid BigNumber string")
//	case []byte:
//		return NewBigNumber(hex.EncodeToString(v)), nil
//	case *[]byte:
//		return NewBigNumber(hex.EncodeToString(*v)), nil
//
//	case uint:
//	case uint8:
//	case uint16:
//	case uint32:
//	case uint64:
//	case int:
//	case int8:
//	case int16:
//	case int32:
//	case int64:
//		return NewBigNumber(strconv.Itoa(int(v))), nil
//	}
//	return nil, errors.New("invalid BigNumber Value")
//}
//
//func NewBigNumber(hex string) *big.Int {
//	return &BigNumber{
//		hex:         hexPrefix + hex,
//		isBigNumber: true,
//	}
//}

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
	//cv := CLValue{
	//	Value: e,
	//	CLType: CLType{
	//		Tag: numTag,
	//	},
	//	ByteHex: hex.EncodeToString(toBytes(e.Val, e.BitSize)),
	//}
	return &NumberCoder{
		BitSize: bitSize,
		Tag:     tag,
		Signed:  signed,
		Val:     NumberFrom(val),
		Name:    n + strconv.Itoa(int(bitSize)),
	}
}

func (nc *NumberCoder) GetCLType() CLType {
	//return NumberCoderType{
	//	Tag:  nc.Tag,
	//	Name: nc.Name,
	//}
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
	//case uint, uint8, uint16, uint32, int:
	//	return big.NewInt(int64(v)), nil
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

//{
//	Tag  uint8  `json:"-"`
//	Name string `json:"name"`
//}

func (n NumberCoderType) ToJson() []byte {
	return nil
}
