package deploy

import (
	"encoding/hex"
	"math/big"
	"testing"
)

func TestModuleBytes_ToBytes(t *testing.T) {
	payment, err := StandardPayment(big.NewInt(1024))
	if err != nil {
		t.Fatal(err)
	}
	if "00000000000100000006000000616d6f756e740300000002000408" != hex.EncodeToString(payment.ItemModuleBytes.ToBytes()) {
		t.Fatal("moduleBytes toBytes error")
	}
}
