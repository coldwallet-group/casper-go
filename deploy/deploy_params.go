package deploy

import (
	"github/casper-go/keys"
	"time"
)

type Params struct {
	accountPublicKey []byte
	keyAlgorithm     keys.SignatureAlgorithm
	chainName        string
	gasPrice         uint64
	ttl              uint64
	dependencies     [][]byte
	timestamp        time.Time
}

func NewParams(rawPublicKey []byte) *Params {
	return &Params{
		accountPublicKey: rawPublicKey,
		chainName:        "delta-11", //从配置文件获取
		gasPrice:         1,          //从配置文件获取
		ttl:              3600000,    //从配置文件获取
		dependencies:     [][]byte{},
		//timestamp:        uint64(time.Now().Unix()),
		timestamp: time.Now(),
	}
}
