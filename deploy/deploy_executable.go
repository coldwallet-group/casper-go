package deploy

import (
	"bytes"
	cl "github/casper-go/clvalue"
)

type ExecutableDeployer interface {
	ToBytes() []byte
}

type ExecDeployItem struct {
	ItemModuleBytes                   *ModuleBytes  `json:"ModuleBytes,omitempty"`
	ItemStoredContractByHash          *StoredParams `json:"StoredContractByHash,omitempty"`
	ItemStoredContractByName          *StoredParams `json:"StoredContractByName,omitempty"`
	ItemStoredVersionedContractByHash *StoredParams `json:"StoredVersionedContractByHash,omitempty"`
	ItemStoredVersionedContractByName *StoredParams `json:"StoredVersionedContractByName,omitempty"`
	ItemTrx                           *Transfer     `json:"Transfer,omitempty"`
}

func (e *ExecDeployItem) ToBytes() []byte {
	if e.ItemModuleBytes != nil {
		return e.ItemModuleBytes.ToBytes()
	}
	if e.ItemTrx != nil {
		return e.ItemTrx.ToBytes()
	}
	panic("failed to serialize ExecDeployItem")
}

type RuntimeArgs struct {
	Args [][2]interface{} `json:"args,omitempty"`
}

func (r *RuntimeArgs) ToBytes() []byte {
	argLen := len(r.Args)
	var bs [][]byte
	bs = append(bs, cl.ToBytesU32(uint32(argLen)))

	for i := 0; i < argLen; i++ {
		k := r.Args[i][0].(string)
		v := r.Args[i][1].(*cl.CLValue)
		bs = append(bs, cl.ToBytesString(k))
		bs = append(bs, v.ToBytes())
	}
	return bytes.Join(bs, []byte{})
}

type StoredParams struct {
	Tag        uint8       `json:"tag,omitempty"`
	Hash       []byte      `json:"hash,omitempty"`
	version    int         `json:"version,omitempty"`
	EntryPoint string      `json:"entryPoint,omitempty"`
	Args       RuntimeArgs `json:"args,omitempty"`
}
