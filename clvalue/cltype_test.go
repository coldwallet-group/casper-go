package clvalue

import (
	"encoding/hex"
	"fmt"
	"github/casper-go/keys"
	"testing"
)

func TestByteArray_GetCLType(t *testing.T) {
	ds, _ := hex.DecodeString("447239548b66bdfe334131392dd9db386c054989e2b815fe68fd634c9e4703a1")

	holder, err := keys.NewKeyHolder(nil, ds, "ed25519")
	if err != nil {
		t.Fatal(err)
	}

	accountHex, err := holder.AccountHash()
	if err != nil {
		t.Fatal(err)
	}
	ds2, err := hex.DecodeString(accountHex)
	d := NewByteArray(ds2)

	fmt.Println(d)
}
