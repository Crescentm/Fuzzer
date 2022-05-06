package fuzz

import (
	"Fuzzer/src/contract"
	. "Fuzzer/src/parameterType"
	. "Fuzzer/src/pidType"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func RunSeeds(program string, seeds any, fcost any, method abi.Method, inputType INPUTTYPE, abiObject abi.ABI, costSize int) PIDS {
	var inputList INPUTDATA
	var pids = make(PIDS)
	for _, arg := range method.Inputs {
		argumentName := arg.Name
		argumentType := arg.Type
		fmt.Println(argumentName, argumentType)
		parameter := NewParameterType(&argumentType)
		switch {
		case parameter.TypeName[:3] == "int":
			varInt := NewIntType(parameter)
			inputList = append(inputList, varInt.Generate())
		case parameter.TypeName[:4] == "uint":
			varUint := NewUintType(parameter)
			inputList = append(inputList, varUint.Generate())
		default:
			continue
		}
	}
	pidNew, path := contract.SendContract(inputList, inputType, method.Name, program, fcost, abiObject, costSize)
	pids[path] = pidNew
	return pids
}
