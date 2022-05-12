package jsonData

type DeployReq struct {
	AbiArray interface{} `json:"abi"`
	Bytecode string      `json:"bytecodes"`
	Data     string      `json:"data"`
}

type DeployRes struct {
	AbiArray interface{} `json:"abi"`
	Address  string      `json:"address"`
}

type TransactionReq struct {
	Address string `json:"address"`
	Data    string `json:"data"`
}

type TransactionRes struct {
	TransactionHash string   `json:"transaction_hash"`
	PC              []string `json:"pc"`
	Cost            []string `json:"cost"`
	Status          []string `json:"status"`
}
