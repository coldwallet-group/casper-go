package clvalue

import (
	"math/big"
)

type U64 struct {
	NumberCoder
	value big.Int `json:"value,omitempty"`
}

func NewU64CLValue(value *big.Int) *CLValue {
	return NewCLValue(NewU64(value))
}

func NewU64(value *big.Int) *U64 {
	return &U64{
		NumberCoder: *NewNumberCoder(64, uint8(TagU64), false, value),
	}
}

func (nc *U64) GetCLType() CLType {
	return nc.NumberCoder.GetCLType()
}

func (nc *U64) ToBytes() []byte {
	return nc.NumberCoder.ToBytes()
}
