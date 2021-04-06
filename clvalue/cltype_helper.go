package clvalue

func CLTypeHelperOption(innerType CLType) *OptionType {
	return NewOptionType(innerType)
}

func CLTypeHelperByteArray(l uint32) *ByteArrayType {
	return NewByteArrayType(l)
}
