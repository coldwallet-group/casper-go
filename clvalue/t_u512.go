package clvalue

import (
	"math/big"
)

type U512 struct {
	NumberCoder
	value big.Int `json:"value,omitempty"`
}

func NewU512CLValue(value *big.Int) *CLValue {
	return NewCLValue(NewU512(value))
}

func NewU512(value *big.Int) *U512 {
	return &U512{
		NumberCoder: *NewNumberCoder(512, uint8(TagU512), false, value),
	}
}

func (nc *U512) GetCLType() CLType {
	return nc.NumberCoder.GetCLType()
}

func (nc *U512) ToBytes() []byte {
	return nc.NumberCoder.ToBytes()
}
