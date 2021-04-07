# casper-go

CasperLabs client SDK

Go语言实现的CasperLabs客户端SDK

## API

### PutDeploy

可以通过提交deploy的方式来进行转账：

```golang
// 交易的发起人，准备一对秘钥用于签名
sender := NewKeyHolder(<private key>, <public key>, "secp256k1")

// 收款人，只需要公钥
// 转换为AccountHash值
recipient := NewKeyHolder(nil, <public key>, "secp256k1")
recipientAccountHash := recipient.AccountHash()

// deploy需要由payment和session两部分组成
// 创建一个标准的payment数据
payment := NewStandardPayment(big.NewInt(1024))
// 创建一个transfer类型的session数据，转账金额为2500000000
session := NewTransfer(big.NewInt(2500000000), recipientAccountHash, nil)

// 生成deploy数据
deploy := MakeDeploy(NewParams(sender.RawPublicKey()), session, payment)

// 对deploy数据进行签名
deploy.Sign(sender)

// 创建RPC
casper := client.New(<rpcURL>, <eventStoreAPI>)

// 推送deploy提交交易
casper.PutDeploy(deploy)
```
