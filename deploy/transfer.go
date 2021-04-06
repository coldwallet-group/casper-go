package deploy

import (
	"bytes"
	cl "github/casper-go/clvalue"
	"math/big"
)

type Transfer struct {
	Tag uint8 `json:"-"`
	RuntimeArgs
}

func NewTransfer(amount *big.Int, accountHash []byte, id *big.Int) (*ExecDeployItem, error) {
	var idCLvalue *cl.CLValue
	if id == nil {
		idCLvalue = cl.NewOptionCLValue(nil, cl.TagU64)
	} else {
		idCLvalue = cl.NewOptionCLValue(cl.NewU64(id), cl.TagU64)
	}

	sessionMap := [][2]interface{}{
		{"amount", cl.NewU512CLValue(amount)},
		{"target", cl.NewByteArrayCLValue(accountHash)},
		{"id", idCLvalue},
	}
	return &ExecDeployItem{
		ItemTrx: &Transfer{
			Tag: 5,
			RuntimeArgs: RuntimeArgs{
				Args: sessionMap,
			},
		},
	}, nil
}

func (t *Transfer) ToBytes() []byte {
	return bytes.Join([][]byte{
		{t.Tag},
		cl.ToBytesBytesArray(t.RuntimeArgs.ToBytes()),
	}, []byte{})
}
