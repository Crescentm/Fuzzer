package tools

import (
	"Fuzzer/src/parameterType"
	"Fuzzer/src/pidType"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func RandomInput(method abi.Method) pidType.INPUTDATA {
	var inputList pidType.INPUTDATA
	for _, arg := range method.Inputs {
		argumentType := arg.Type
		parameter := parameterType.NewParameterType(&argumentType)
		switch {
		case parameter.TypeName[:3] == "int":
			varInt := parameterType.NewIntType(parameter)
			inputList = append(inputList, varInt.Generate())
		case parameter.TypeName[:4] == "uint":
			varUint := parameterType.NewUintType(parameter)
			inputList = append(inputList, varUint.Generate())
		case parameter.TypeName[:4] == "bool":
			varBool := parameterType.NewBoolType(parameter)
			inputList = append(inputList, varBool.Generate())
		case parameter.TypeName[:4] == "byte":
			varByte := parameterType.NewByteType(parameter)
			inputList = append(inputList, varByte.Generate())
		case parameter.TypeName[:6] == "string":
			varString := parameterType.NewStringType(parameter, 24)
			inputList = append(inputList, varString.Generate())
		default:
			continue
		}
	}
	return inputList
}
