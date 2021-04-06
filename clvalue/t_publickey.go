package clvalue

import (
	"github/casper-go/common/byteutil"
	"github/casper-go/keys"
)

const (
	Ed25519Tag   uint8 = 1
	Secp256k1Tag uint8 = 2
)

type PublicKey struct {
	RawPublicKey []byte
	Algorithm    keys.SignatureAlgorithm
	Tag          uint8
}

func NewPublicKey(rawPubKey []byte, algorithm keys.SignatureAlgorithm) *PublicKey {
	keyTag := Ed25519Tag
	if algorithm == keys.Secp256K1 {
		keyTag = Secp256k1Tag
	}
	return &PublicKey{
		RawPublicKey: rawPubKey,
		Algorithm:    algorithm,
		Tag:          keyTag,
	}
}

func (pk *PublicKey) GetCLType() interface{} {
	return TagPublicKey
}

func (pk *PublicKey) ToBytes() []byte {
	return byteutil.Concat([]byte{pk.Tag}, ToBytesBytesArray(pk.RawPublicKey))
}
