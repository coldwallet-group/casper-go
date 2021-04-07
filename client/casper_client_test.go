package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"
)

const (
	eventStoreApi = "https://event-store-api-clarity-delta.make.services"
	RpcUrl        = "https://node-clarity-delta.make.services/rpc"
)

var (
	casper = New(RpcUrl, eventStoreApi)
)

func Test_GetStatus(t *testing.T) {

	status, err := casper.GetStatus()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(status.ApiVersion)
}

func Test_GetLatestBlockInfo(t *testing.T) {

	height, err := casper.GetLatestBlockHeight()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(height)

}

func Test_GetBlockInfoByHeight(t *testing.T) {
	block, err := casper.GetBlockInfoByHeight(18584)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(block)
	fmt.Println(string(d))
}

func Test_GetBlockInfoByHash(t *testing.T) {
	block, err := casper.GetBlockInfoByHash("6f168ef1d9bfcca97146b4925924e9594dba03a3fe30952653867ecc5fda5746")
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(block)
	fmt.Println(string(d))
}

func Test_GetDeployByDeployHash(t *testing.T) {
	txid := "4d1ee570091c8dd064ec174daea85eeddc8eda24fc12b037110adfc214abb739"
	casper.GetDeployByHash(txid)
}

func Test_GetBlockTransfer(t *testing.T) {
	height, err := casper.GetBlockTransferByHeight(18584)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(height)
}

func Test_GetSenderBalance(t *testing.T) {
	s := "01d74e5088891f2c938a38e4dbd37d18157bb65ef97a5cdef1aea44a2293d8d2b2"
	ds, _ := hex.DecodeString(s)
	balance, err := casper.GetBalance(ds)
	if err != nil {
		t.Fatal(err)
	}
	//919299770000 - 19299770000(amount) - 10000(gas)=
	//899999990000
	fmt.Println(balance)
}

func Test_GetRecipientBalance(t *testing.T) {
	s := "01a027ac95925adf648e1a8902dab39e7899f919644c625f21cf4eec9d1b2f158f"
	ds, _ := hex.DecodeString(s)
	balance, err := casper.GetBalance(ds)
	if err != nil {
		t.Fatal(err)
	}
	//40000000000
	//59299770000
	fmt.Println(balance)
}
