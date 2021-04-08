package keys

import (
	"errors"
	"fmt"
	"github/casper-go/common/byteutil"
	"github/casper-go/common/hexutil"
	"github/casper-go/keys/blake2b"
	"strings"
)

type SignatureAlgorithm string

const (
	Ed25519   = SignatureAlgorithm("ed25519")
	Secp256K1 = SignatureAlgorithm("secp256k1")
)

type KeyHolder interface {
	PrivateToPubKey() ([]byte, error)
	AccountHash() []byte
	//账号十六进制格式
	//注意：这里返回的是accountHex，并非accountHash
	//accountHex代表一个账号的唯一值，从公钥哈希值派生
	//可用作查询账号相关的操作，如余额
	AccountHex() string
	Sign(message []byte) (sig []byte, err error)
	Verify(message, sig []byte) bool
	PublicKey() []byte
	Prefix() string
	Algorithm() SignatureAlgorithm

	//私钥转换成PEM文件格式（加工过的base64格式）例如：
	//-----BEGIN PRIVATE KEY-----
	//MC4CAQAwBQYDK2VwBCIEIBi2p4YSZ58JCjZuKSdKbFB8ixdrJIZHqNMtaJIuhOF5
	//-----END PRIVATE KEY-----
	ParsePrivateKeyToPem() (string, error)

	//公钥转换成PEM文件格式（加工过的base64格式）例如：
	//-----BEGIN PUBLIC KEY-----
	//MCowBQYDK2VwAyEAeKEooE0MhphnznYVBcR+slT22meCiBHH6WYIs6KKHjw=
	//-----END PUBLIC KEY-----
	ParsePublicKeyToPem() (string, error)
}

//根据不同算法构造keyHolder，公钥和私钥入参不一定都需要
//private：私钥
//pub：公钥，不带0x，不带前缀（01/02）
//algorithm：具体算法
func NewKeyHolder(private []byte, pub []byte, algorithm SignatureAlgorithm) (KeyHolder, error) {
	if algorithm == Secp256K1 {
		return NewSECP256K1(private, pub)
	} else {
		return NewED25519(private, pub)
	}
}

//从accountHex生成keyHolder
func NewKeyHolderFromAccountHex(accountHexBytes []byte) (KeyHolder, error) {
	if accountHexBytes[0] == 2 {
		return NewSECP256K1(nil, accountHexBytes[1:])
	} else if accountHexBytes[0] == 1 {
		return NewED25519(nil, accountHexBytes[1:])
	}
	return nil, errors.New("invalid prefix, accountHex prefix must 01 or 02")
}

func IsAccount(addr string) bool {
	if hexutil.Has0xPrefix(addr) {
		addr = addr[2:]
	}
	if !hexutil.IsHex(addr) {
		return false
	}
	prefix := addr[:2]
	addrLen := 0
	if prefix == "01" {
		addrLen = 66
	} else if prefix == "02" {
		addrLen = 68
	} else {
		return false
	}
	return addrLen == len(addr)
}

func CheckPrivKey(priv []byte, l int) error {
	if priv == nil {
		return errors.New("CheckPrivKey:privKey require")
	}
	if len(priv) != l {
		return errors.New(fmt.Sprintf("CheckPrivKey:invalid privKey len"))
	}
	return nil
}

//根据公钥数据生成accountHash
func AccountHash(pub []byte, sa SignatureAlgorithm) []byte {
	separator := []byte{0}
	prefix := byteutil.Concat([]byte(strings.ToLower(string(sa))), separator)
	return blake2b.Hash(byteutil.Concat(prefix, pub))
}
