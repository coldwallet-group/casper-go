package deploy

import (
	"fmt"
	"math/big"
	"testing"
)

func TestModuleBytes_ToBytes(t *testing.T) {
	payment, err := StandardPayment(big.NewInt(1024))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(payment.ItemModuleBytes.ToBytes())
	fmt.Println(len(payment.ItemModuleBytes.ToBytes()))
}

