package tools

import (
	"Fuzzer/src/pidType"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func InputEncode(inputList pidType.INPUTDATA, methodName string, abiObject abi.ABI) []byte {
	inputPacked, err := abiObject.Pack(methodName, inputList...)
	if err != nil {
		panic(err)
	}
	return inputPacked
}
