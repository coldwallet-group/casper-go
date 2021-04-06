package byteutil

import "bytes"

func Concat(a []byte, b []byte) []byte {
	return bytes.Join([][]byte{
		a,
		b,
	}, []byte{})
}
