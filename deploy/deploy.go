package deploy

import (
	"encoding/hex"
	cl "github/casper-go/clvalue"
	"github/casper-go/common/byteutil"
	"github/casper-go/keys"
	"github/casper-go/keys/blake2b"
)

type Deploy struct {
	Hash      []byte          `json:"-"`
	JSONHash  string          `json:"hash,omitempty"`
	Header    *Header         `json:"header,omitempty"`
	Payment   *ExecDeployItem `json:"payment,omitempty"`
	Session   *ExecDeployItem `json:"session,omitempty"`
	Approvals []*Approval     `json:"approvals,omitempty"`
}

type Approval struct {
	Signer    string `json:"signer,omitempty"`
	Signature string `json:"signature,omitempty"`
}

func MakeDeploy(params *Params, session *ExecDeployItem, payment *ExecDeployItem) (*Deploy, error) {
	bodyBytes := serializeBody(payment, session)
	bodyHash := blake2b.Hash(bodyBytes)
	header := NewHeader(cl.NewPublicKey(params.accountPublicKey, params.keyAlgorithm), bodyHash, params)
	headerBytes := serializeHeader(header)
	deployHash := blake2b.Hash(headerBytes)

	return &Deploy{
		Hash:      deployHash,
		JSONHash:  hex.EncodeToString(deployHash),
		Header:    header,
		Payment:   payment,
		Session:   session,
		Approvals: []*Approval{},
	}, nil
}

func (d *Deploy) Sign(signer keys.KeyHolder) error {
	signMsg, err := signer.Sign(d.Hash)
	if err != nil {
		return err
	}
	signature := signer.Prefix() + hex.EncodeToString(signMsg)
	d.Approvals = append(d.Approvals, &Approval{
		Signer:    signer.AccountHex(),
		Signature: signature,
	})
	return nil
}

func serializeBody(payment *ExecDeployItem, session *ExecDeployItem) []byte {
	return byteutil.Concat(payment.ToBytes(), session.ToBytes())
}

func serializeHeader(header *Header) []byte {
	return header.ToBytes()
}
