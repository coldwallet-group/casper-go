package clvalue

import (
	"bytes"
	"github/casper-go/common/byteutil"
)

const (
	OptionTagNone = 0
	OptionTagSome = 1
)

type Option struct {
	T         *BytesSerializable `json:"t,omitempty""`
	InnerType CLT                `json:"innerType,omitempty"`
}

func NewOptionCLValue(t BytesSerializable, innerType CLT) *CLValue {
	return NewCLValue(NewOption(t, innerType))
}

func NewOption(t BytesSerializable, innerType CLT) *Option {
	var tInstantiation *BytesSerializable
	if t == nil {
		tInstantiation = nil
	} else {
		tInstantiation = &t
	}

	return &Option{
		T:         tInstantiation,
		InnerType: innerType,
	}
}

func (op *Option) GetCLType() CLType {
	return CLTypeHelperOption(&op.InnerType)
}

func (op *Option) ToJson() []byte {
	return nil
}

func (op *Option) ToWrapBytes() []byte {
	return byteutil.Concat(op.ToBytes(), ToBytesHelper(TagOption))
}

func (op *Option) ToBytes() []byte {
	if op.T == nil {
		return []byte{OptionTagNone}
	}
	t := *op.T
	return bytes.Join([][]byte{
		{OptionTagSome},
		t.ToBytes(),
	}, []byte{})
}

type OptionType struct {
	OptionValue interface{} `json:"Option"`
	Type        string      `json:"-"`
	Tag         CLT         `json:"-"`
	InnerType   CLType      `json:"-"`
}

func NewOptionType(innerType CLType) *OptionType {
	return &OptionType{
		OptionValue: "U64",
		Type:        "Option",
		Tag:         TagOption,
		InnerType:   innerType,
	}
}
