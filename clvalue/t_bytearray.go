package clvalue

type ByteArray struct {
	RawBytes []byte `json:"-"`
	Size     uint32 `json:"-"`
}

func NewByteArrayCLValue(data []byte) *CLValue {
	return NewCLValue(NewByteArray(data))
}

func NewByteArray(data []byte) *ByteArray {
	return &ByteArray{
		RawBytes: data,
		Size:     uint32(len(data)),
	}
}

func (u *ByteArray) GetCLType() CLType {
	return CLTypeHelperByteArray(uint32(len(u.RawBytes)))
}

func (u *ByteArray) ToBytes() []byte {
	return ToBytesBytesArray(u.RawBytes)
}

type ByteArrayType struct {
	ByteArrayValue interface{} `json:"ByteArray"`
	Tag            CLT         `json:"-"`
	Size           uint32      `json:"-"`
}

func NewByteArrayType(size uint32) *ByteArrayType {
	return &ByteArrayType{
		ByteArrayValue: size,
		Tag:            TagByteArray,
		Size:           size,
	}
}
