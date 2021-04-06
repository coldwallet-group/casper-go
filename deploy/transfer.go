package deploy

import (
	"bytes"
	cl "github/casper-go/clvalue"
	"github/casper-go/keys"
	"math/big"
)

type Transfer struct {
	Tag uint8 `json:"-"`
	RuntimeArgs
}

func NewTransfer(amount *big.Int, targetKeyHolder keys.KeyHolder, sourcePurse []byte, id *big.Int) (*ExecDeployItem, error) {

	accountHash := targetKeyHolder.AccountHash()
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
	//
	//var argMap []cl.CLMap
	//argMap = append(argMap, cl.CLMap{
	//	Name:  "amount",
	//	Value: cl.NewU512CLValue(amount),
	//})
	//accountHash := targetKeyHolder.AccountHash()
	//argMap = append(argMap, cl.CLMap{
	//	Name:  "target",
	//	Value: cl.NewByteArrayCLValue(accountHash),
	//})
	//if id == nil {
	//	argMap = append(argMap, cl.CLMap{
	//		Name:  "id",
	//		Value: cl.NewOptionCLValue(nil, cl.TagU64),
	//	})
	//} else {
	//	argMap = append(argMap, cl.CLMap{
	//		Name:  "id",
	//		Value: cl.NewOptionCLValue(cl.NewU64(id), cl.TagU64),
	//	})
	//}

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
