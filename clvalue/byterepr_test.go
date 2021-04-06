package clvalue

import (
	"fmt"
	"testing"
)

func TestToBytesU32(t *testing.T) {
	u8 := ToBytesArrayU8([]byte{12})
	fmt.Println(u8)
}


func TestToBytesString(t *testing.T) {
	fmt.Println(ToBytesString("123abc"))
}