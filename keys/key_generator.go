package keys

//密钥生成接口
//包含密钥生成相关的功能
type KeyGenerator interface {
	//生成密钥对
	//return:私钥 公钥
	GenerateKey() ([]byte, []byte, error)

	//根据种子生成密钥对
	//return:私钥 公钥
	GenerateKeyBySeed(seed []byte) ([]byte, []byte, error)
}

//根据不同签名算法，生成对应的keyGenerator
//目前支持ed2519/secp256k1
func NewKeyGenerator(algorithm SignatureAlgorithm) KeyGenerator {
	//由于生成密钥操作只是单纯的New对应的实例
	//并不会执行其他逻辑，所以这里是可以忽略返回的error
	if algorithm == Secp256K1 {
		k, _ := NewSECP256K1(nil, nil)
		return k
	} else {
		k, _ := NewED25519(nil, nil)
		return k
	}
}
