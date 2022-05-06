package tools

import (
	"Fuzzer/src/pidType"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func InputParse(methodObject abi.Method) (int, pidType.INPUTTYPE) {
	inputSize := len(methodObject.Inputs)
	var inputType pidType.INPUTTYPE
	for _, arg := range methodObject.Inputs {
		inputType = append(inputType, arg.Type)
	}
	return inputSize, inputType
}
