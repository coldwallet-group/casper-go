package deploy

import (
	"bytes"
	cl "github/casper-go/clvalue"
)

type ModuleBytes struct {
	Tag         uint8  `json:"-"`
	ModuleBytes []byte `json:"module_bytes"`
	RuntimeArgs
}

func NewModuleBytes(moduleBytes []byte, args RuntimeArgs) *ModuleBytes {
	return &ModuleBytes{
		ModuleBytes: moduleBytes,
		RuntimeArgs: args,
		Tag:         0,
	}
}

func (mb *ModuleBytes) ToBytes() []byte {
	return bytes.Join([][]byte{
		{mb.Tag},
		cl.ToBytesArrayU8(mb.ModuleBytes),
		cl.ToBytesBytesArray(mb.RuntimeArgs.ToBytes()),
	}, []byte{})
}
