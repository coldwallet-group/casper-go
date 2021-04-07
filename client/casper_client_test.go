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
	txid := "20f1190d4ddc06246e07d5fd0454d90f3b509936e3d2584350239104e183a000"
	casper.GetDeployByHash(txid)
}

func Test_GetBlockTransfer(t *testing.T) {
	height, err := casper.GetBlockTransferByHeight(18584)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(height)
}

func Test_GetBalance(t *testing.T) {
	s := "8fee8e44d228eb323e72ff977699dab804d8250868bf6533d3a48b82fe46631d"
	ds, _ := hex.DecodeString(s)
	//19eedb94674ca854fd0483edd5d23120ac3562f34055521ae22282fe3acc4c65
	balance, err := casper.GetBalance(ds)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(balance)
}
