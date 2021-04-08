package deploy

import (
	"github/casper-go/keys"
	"time"
)

//deploy的相关参数
type Params struct {
	//交易发起人accountHex
	accountPublicKey []byte
	keyAlgorithm     keys.SignatureAlgorithm
	//使用的区块链的名称，测试链使用delta-11
	chainName string
	//设置gas价格
	gasPrice uint64
	//交易的有效时间，单位毫秒
	ttl          uint64
	dependencies [][]byte
	//时间戳
	timestamp time.Time
}

func NewParams(publicKey []byte, algorithm keys.SignatureAlgorithm) *Params {
	return &Params{
		accountPublicKey: publicKey,
		keyAlgorithm:     algorithm,
		chainName:        "delta-11", //从配置文件获取
		gasPrice:         1,          //从配置文件获取
		ttl:              3600000,    //从配置文件获取
		dependencies:     [][]byte{},
		timestamp:        time.Now(), //当前时间
	}
}
