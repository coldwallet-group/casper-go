package deploy

import (
	cl "github/casper-go/clvalue"
	"math/big"
)

func StandardPayment(paymentAmount *big.Int) (*ExecDeployItem, error) {
	ra := RuntimeArgs{
		Args: [][2]interface{}{{"amount", cl.NewU512CLValue(paymentAmount)}},
	}
	payment := &ExecDeployItem{
		ItemModuleBytes: NewModuleBytes([]byte{}, ra),
	}
	return payment, nil
}
