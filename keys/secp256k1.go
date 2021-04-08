package keys

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"
	"math/big"
)

type SECP256K1 struct {
	//使用此算法生成的密钥对应的账号前缀
	prefix string
	//使用的算法
	algorithm SignatureAlgorithm
	//压缩的公钥字节长度
	pubCompressByteLen int
	//生成的私钥字节长度，注意这里是原始的私钥长度
	privByteLen int

	//保存的私钥数据
	privateKey     []byte
	publicKeyEcdsa *ecdsa.PublicKey
}

func NewSECP256K1(private []byte, public []byte) (*SECP256K1, error) {
	//传入的公钥数据可以是两种不同的格式：压缩和非压缩
	//压缩的长度是33，非压缩是65
	//如果是压缩的公钥，需要转换为完整的格式
	//最后只保留完整格式公钥
	var pub *ecdsa.PublicKey
	if public != nil {
		if len(public) == 33 {
			rawPublic, err := ethcrypto.DecompressPubkey(public)
			if err != nil {
				return nil, err
			}
			pub = rawPublic
		} else if len(public) == 65 {
			rawPublic, err := ethcrypto.UnmarshalPubkey(public)
			if err != nil {
				return nil, err
			}
			pub = rawPublic
		} else {
			return nil, errors.New("failed to new secp256k1:invalid public key len")
		}
	}

	return &SECP256K1{
		prefix:             "02",
		algorithm:          Secp256K1,
		pubCompressByteLen: 33,
		privByteLen:        32,
		privateKey:         private,
		publicKeyEcdsa:     pub,
	}, nil
}

//注意：这里返回的私钥是经过压缩的02或03开头33字节长度
func (s *SECP256K1) GenerateKey() ([]byte, []byte, error) {
	priv, err := ethcrypto.GenerateKey()
	if err != nil {
		return nil, nil, err
	}
	l := len(priv.D.Bytes())
	if l != s.privByteLen {
		return nil, nil, errors.New(fmt.Sprintf("%s genrateKey:invalid key len", s.algorithm))
	}

	pub := priv.PublicKey
	return priv.D.Bytes(), ethcrypto.CompressPubkey(&pub), nil
}

//注意：这里返回的公钥是经过压缩的02或03开头33字节长度
func (s *SECP256K1) GenerateKeyBySeed(seed []byte) ([]byte, []byte, error) {
	cruve := secp256k1.S256()
	priv := new(ecdsa.PrivateKey)
	priv.D = new(big.Int).SetBytes(seed)
	priv.X, priv.Y = cruve.ScalarBaseMult(priv.D.Bytes())
	pub := priv.PublicKey
	return priv.D.Bytes(), ethcrypto.CompressPubkey(&pub), nil
}

//注意：这里返回的公钥是经过压缩的02或03开头33字节长度
func (s *SECP256K1) PrivateToPubKey() ([]byte, error) {
	if err := CheckPrivKey(s.privateKey, s.privByteLen); err != nil {
		return nil, err
	}
	if len(s.privateKey) != s.privByteLen {
		return nil, errors.New(fmt.Sprintf("%s PrivateToPubKey:invalid key len", s.algorithm))
	}
	priv, err := ethcrypto.HexToECDSA(hex.EncodeToString(s.privateKey))
	if err != nil {
		return nil, err
	}
	return ethcrypto.CompressPubkey(&priv.PublicKey), nil
}

func (s *SECP256K1) AccountHash() []byte {
	return AccountHash(ethcrypto.CompressPubkey(s.publicKeyEcdsa), s.algorithm)
}

func (s *SECP256K1) AccountHex() string {
	return s.prefix + hex.EncodeToString(ethcrypto.CompressPubkey(s.publicKeyEcdsa))
}

//返回不带`V`的签名数据
func (s *SECP256K1) Sign(message []byte) ([]byte, error) {
	if err := CheckPrivKey(s.privateKey, s.privByteLen); err != nil {
		return nil, err
	}
	priv, err := ethcrypto.ToECDSA(s.privateKey)
	if err != nil {
		return nil, err
	}
	digestHash := sha256.Sum256(message)
	sig, err := ethcrypto.Sign(digestHash[:], priv)
	if err != nil {
		return nil, err
	}
	return sig[:len(sig)-1], nil
}

func (s *SECP256K1) Verify(message, sig []byte) bool {
	digestHash := sha256.Sum256(message)
	return ethcrypto.VerifySignature(ethcrypto.CompressPubkey(s.publicKeyEcdsa), digestHash[:], sig)
}

func (s *SECP256K1) PublicKey() []byte {
	return ethcrypto.CompressPubkey(s.publicKeyEcdsa)
}

func (s *SECP256K1) Prefix() string {
	return s.prefix
}

func (s *SECP256K1) Algorithm() SignatureAlgorithm {
	return Secp256K1
}

func (s *SECP256K1) ParsePrivateKeyToPem() (string, error) {
	//TODO 尚未从typescript翻译
	panic("TODO")
	return "", nil
}

func (s *SECP256K1) ParsePublicKeyToPem() (string, error) {
	//TODO 尚未从typescript翻译
	panic("TODO")
	return "", nil
}
