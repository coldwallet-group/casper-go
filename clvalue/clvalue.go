package clvalue

import (
	"encoding/hex"
	"github/casper-go/common/byteutil"
)

type CLT uint8

const (
	TagBool      CLT = 0
	TagI32       CLT = 1
	TagI64       CLT = 2
	TagU8        CLT = 3
	TagU32       CLT = 4
	TagU64       CLT = 5
	TagU128      CLT = 6
	TagU256      CLT = 7
	TagU512      CLT = 8
	TagUnit      CLT = 9
	TagString    CLT = 10
	TagKey       CLT = 11
	TagURef      CLT = 12
	TagOption    CLT = 13
	TagList      CLT = 14
	TagByteArray CLT = 15
	TagResult    CLT = 16
	TagMap       CLT = 17
	TagTuple1    CLT = 18
	TagTuple2    CLT = 19
	TagTuple3    CLT = 20
	TagAny       CLT = 21
	TagPublicKey CLT = 22
)

func (c CLT) ToJson() []byte {
	return nil
}

type CLType interface {
}

type CLValue struct {
	Value   BytesSerializable `json:"-"`
	Type    CLType            `json:"cl_type,omitempty"`
	ByteHex string            `json:"bytes,omitempty"`
	Parsed  string            `json:"parsed,omitempty"`
}

type BytesSerializable interface {
	GetCLType() CLType
	ToBytes() []byte
}

func NewCLValue(value BytesSerializable) *CLValue {
	return &CLValue{
		Value:   value,
		Type:    value.GetCLType(),
		ByteHex: hex.EncodeToString(value.ToBytes()),
		Parsed:  "null",
	}
}

func (c *CLValue) GetCLType() interface{} {
	return c.Value.GetCLType()
}

func (c *CLValue) ToBytes() []byte {
	helper := ToBytesHelper(c.Type)
	return byteutil.Concat(ToBytesArrayU8(c.Value.ToBytes()), helper)
}

type CLMap struct {
	Name  string
	Value BytesSerializable
}
