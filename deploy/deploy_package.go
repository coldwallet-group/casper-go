package deploy

type PackageData struct {
	PackageDeploy `json:"deploy,omitempty"`
}

type PackageHeader struct {
	Account      string   `json:"account,omitempty"`
	Timestamp    string   `json:"timestamp,omitempty"`
	TTL          string   `json:"ttl,omitempty"`
	GasPrice     uint64   `json:"gas_price,omitempty"`
	BodyHash     string   `json:"body_hash,omitempty"`
	Dependencies []string `json:"dependencies"`
	ChainName    string   `json:"chain_name,omitempty"`
}

type PackageDeploy struct {
	Hash      string             `json:"hash,omitempty"`
	Header    PackageHeader      `json:"header,omitempty"`
	Payment   PackagePayment  `json:"payment,omitempty"`
	Session   PackageTransfer  `json:"session,omitempty"`
	Approvals []PackageApprovals `json:"approvals,omitempty"`
}

type PackageDeployItem struct {
	Module   PackageModuleBytes  `json:"ModuleBytes,omitempty"`
	Transfer PackageTransferArgs `json:"Transfer,omitempty"`
}

type PackagePayment struct {
	Module PackageModuleBytes `json:"ModuleBytes,omitempty"`
}

type PackageModuleBytes struct {
	ModuleBytes string          `json:"module_bytes"`
	Args        [][]interface{} `json:"args,omitempty"`
}

type PackageItem struct {
	CLType interface{} `json:"cl_type,omitempty"`
	Bytes  string      `json:"bytes,omitempty"`
	Parsed string      `json:"parsed,omitempty"`
}

type PackageTransfer struct {
	Transfer PackageTransferArgs `json:"Transfer,omitempty"`
}

type PackageTransferArgs struct {
	Args [][]interface{} `json:"args,omitempty"`
}

type PackageApprovals struct {
	Signer    string `json:"signer,omitempty"`
	Signature string `json:"signature,omitempty"`
}
