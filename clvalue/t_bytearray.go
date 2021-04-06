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

//func (u *ByteArray) GetJSONType() interface{} {
//	return map[string]uint32{"ByteArray": u.Size}
//}

type ByteArrayType struct {
	ByteArrayValue interface{} `json:"ByteArray"`
	Tag            CLT         `json:"-"`
	Size           uint32      `json:"-"`
}

//type JSONByteArrayType struct {
//	ValueMap map[string]int
//}

func NewByteArrayType(size uint32) *ByteArrayType {
	return &ByteArrayType{
		ByteArrayValue: size,
		Tag:            TagByteArray,
		Size:           size,
	}
}

//
func (o *ByteArrayType) ToJson() []byte {
	return nil
}

//func (b ByteArrayType) MarshalJSON() ([]byte, error) {
//	return json.Marshal(struct {
//		Option interface{} `json:"ByteArray"`
//	}{
//		Option: map[string]interface{}{b.Type: b.Size},
//	})
//}
