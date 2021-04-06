package deploy

import (
	"fmt"
	cl "github/casper-go/clvalue"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println("Starting test...")
	ml := NewMapList()
	var a, b, c Keyer
	a = &Elements{"Alice",&cl.CLValue{ByteHex: "VALUE Alice"}}
	b = &Elements{"Bob",&cl.CLValue{ByteHex: "VALUE Bob"}}
	c = &Elements{"Conrad",&cl.CLValue{ByteHex: "CVALUE onrad"}}
	ml.Push(a)
	ml.Push(b)
	ml.Push(c)
	cb := func(data Keyer) {
		fmt.Println(ml.dataMap[data.GetKey()].Value.(*Elements).value.ByteHex)
	}
	fmt.Println("Print elements in the order of pushing:")
	ml.Walk(cb)
	fmt.Printf("Size of MapList: %d \n", ml.Size())
	ml.Remove(b)
	fmt.Println("After removing b:")
	ml.Walk(cb)
	fmt.Printf("Size of MapList: %d \n", ml.Size())
}
