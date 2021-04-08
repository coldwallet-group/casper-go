package keys

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github/casper-go/keys/blake2b"
	"testing"
)

const (
	testSecp256k1 = "secp256k1"
)

func TestSECP256K1_GenerateKey(t *testing.T) {
	holder := NewKeyGenerator(testSecp256k1)

	priv, pub, err := holder.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(pub))
	fmt.Println(hex.EncodeToString(priv))
}

func TestSECP256K1_GenerateKeyBySeed(t *testing.T) {
	holder := NewKeyGenerator(testSecp256k1)
	priv, pub, err := holder.GenerateKeyBySeed(blake2b.Hash([]byte("abcqwer!")))
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(priv))
	fmt.Println(hex.EncodeToString(pub))
}

func TestSECP256K1_AccountHex(t *testing.T) {
	_, pub := getSECP256K1Key()
	holder, err := NewKeyHolder(nil, pub, testSecp256k1)
	if err != nil {
		t.Fatal(err)
	}

	accountHex := holder.AccountHex()
	fmt.Println(accountHex)
	if 68 != len(accountHex) {
		t.Errorf("account len error,actual:%d", len(accountHex))
	}
	if accountHex[:2] != "02" {
		t.Fatal("account prefix[:2] error")
	}
	if "0203447239548b66bdfe334131392dd9db386c054989e2b815fe68fd634c9e4703a1" != accountHex {
		t.Fatal("accountHex error")
	}
}

func TestSECP256K1_AccountHash(t *testing.T) {
	_, pub := getSECP256K1Key()
	holder, err := NewKeyHolder(nil, pub, testSecp256k1)
	if err != nil {
		t.Fatal(err)
	}
	if "e4aa3786a33cac2309989d2b9e4e3c418f6a2861d09a271d969b78819fb77970" != hex.EncodeToString(holder.AccountHash()) {
		t.Fatal("generate accountHash error")
	}
}

func TestSECP256K1_Sign(t *testing.T) {

	priv, pub := getSECP256K1Key()

	msg := []byte("abcde!!")

	holder, err := NewKeyHolder(priv, pub, testSecp256k1)
	if err != nil {
		t.Fatal(err)
	}
	sig, err := holder.Sign(msg)
	if err != nil {
		t.Fatal(err)
	}
	verify := holder.Verify(msg, sig)
	if !verify {
		t.Fatal("signature error")
	}
}

func TestSECP256K1_Sign2(t *testing.T) {

	priv, pub := getSECP256K1Key()
	holder, _ := NewKeyHolder(priv, pub, testSecp256k1)

	msg, _ := hex.DecodeString("c46edf904c9205509e9bf1af9b6028ef9ea05726f04f6caf38cf3fde5ecb2c7e")
	sig, _ := hex.DecodeString("70e306ccb97b7afa56a8bc2d41a1dfdb297784a1cc41205b66768063c83d2fb924fb663dcd4ce0c0c761ce9661f8e9c3684b66fa81060b691ff9de3faaae554b00")

	fmt.Println(holder.Verify(msg, sig))

	d, _ := hex.DecodeString("c46edf904c9205509e9bf1af9b6028ef9ea05726f04f6caf38cf3fde5ecb2c7e")

	a := ethcrypto.Keccak256Hash(d)
	fmt.Println(hex.EncodeToString(a.Bytes()))
	fmt.Println(hex.EncodeToString(blake2b.Hash(d)))

	s256 := sha256.Sum256(d)
	fmt.Println(hex.EncodeToString(s256[:]))

	sign, _ := holder.Sign(d)
	fmt.Println(hex.EncodeToString(sign))

}

func getSECP256K1Key() ([]byte, []byte) {
	priv, _ := hex.DecodeString("be798eee9bb3fa267e0525a7633260c5d2a9512dd2f96b8d621f560dd233d99a")
	pub, _ := hex.DecodeString("03447239548b66bdfe334131392dd9db386c054989e2b815fe68fd634c9e4703a1")
	return priv, pub
}
