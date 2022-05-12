package tools

import (
	"Fuzzer/src/pidType"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func InputEncode(inputList pidType.INPUTDATA, methodName string, AbiObject abi.ABI) []byte {
	inputPacked, err := AbiObject.Pack(methodName, inputList...)
	if err != nil {
		panic(err)
	}
	return inputPacked
}
