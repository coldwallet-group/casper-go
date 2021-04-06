package deploy

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github/casper-go/client"
	"testing"
	"time"
)

func TestWallet(t *testing.T) {
	sender, err := mockRecipient()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(sender.AccountHash()))
}

func TestAge(t *testing.T) {

	deploy, err := mockMakeDeploy()
	if err != nil {
		t.Fatal(err)
	}

	sender, err := mockSender()
	if err != nil {
		t.Fatal(err)
	}

	err = deploy.Sign(sender)
	if err != nil {
		t.Fatal(err)
	}

	header := PackageHeader{
		Account:      hex.EncodeToString(deploy.Header.Account.ToBytes()),
		Timestamp:    deploy.Header.Timestamp.Format(time.RFC3339),
		TTL:          "360000000000ms",
		GasPrice:     deploy.Header.GasPrice,
		BodyHash:     hex.EncodeToString(deploy.Header.BodyHash),
		Dependencies: []string{},
		ChainName:    deploy.Header.ChainName,
	}

	i1 := []interface{}{"amount", PackageItem{
		CLType: "U512",
		Bytes:  "020004",
		Parsed: "null",
	}}

	itemMap := [][]interface{}{i1}
	payment := PackagePayment{
		Module: PackageModuleBytes{
			ModuleBytes: "",
			Args:        itemMap,
		},
	}

	m, _ := json.Marshal(payment.Module.Args)
	fmt.Println(string(m))

	se1 := []interface{}{"amount", PackageItem{
		CLType: "U512",
		Bytes:  "0400f90295",
		Parsed: "null",
	}}

	m1 := map[string]int{}
	m1["ByteArray"] = 32
	se2 := []interface{}{"target", PackageItem{
		CLType: m1,
		Bytes:  "23400bdd68d63ffbd3446c4563bf1dd3c7648282ec19b12f0504c6d905bc816d",
		Parsed: "null",
	}}

	se3 := []interface{}{"id", PackageItem{
		CLType: "U64",
		Bytes:  "00",
		Parsed: "null",
	}}

	seMap := [][]interface{}{se1, se2, se3}

	session := PackageTransfer{
		Transfer: PackageTransferArgs{
			Args: seMap,
		},
	}

	approvals := PackageApprovals{
		Signer:    "01d74e5088891f2c938a38e4dbd37d18157bb65ef97a5cdef1aea44a2293d8d2b2",
		Signature: "0119a0fb9ac89f427dcd718c75da3c8985ddc74d97b92403dc5861c6f5b4213af255ea96fecab83ef4910ace5122e02f724f7ac3342efc77153b47d33c6cb4fd05",
	}

	p := PackageDeploy{
		Hash:   "07607c9d604a22a1d96d4a91153636ec5b112325f818232f8220b6b8f07eba9d",
		Header: header,
		Payment:   payment,
		Session:   session,
		Approvals: []PackageApprovals{approvals},
	}

	const (
		eventStoreApi = "https://event-store-api-clarity-delta.make.services"
		RpcUrl        = "https://node-clarity-delta.make.services/rpc"
	)

	//wrap := PackageData{
	//	PackageDeploy: p,
	//}
	//
	d, _ := json.Marshal(p)


	fmt.Println(string(d))

	//m := make(map[string]interface{})
	//j, _ := json.Marshal(p)
	//json.Unmarshal(j, &m)

	casper := client.New(RpcUrl, eventStoreApi)

	result, err := casper.PutDeploy(p)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(result)

}
