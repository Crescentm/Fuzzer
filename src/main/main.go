package main

import "Fuzzer/src/fuzzer"

func main() {
	fuzzer.Init()
	fuzzer.Fuzz()
}

//flag.Parse()
//var err error
//var jsonData = `[{ "inputs": [{"internalType": "uint256","name": "a","type": "uint256"	},{	"internalType": "uint256","name": "b","type": "uint256"}],"stateMutability": "nonpayable","type": "constructor"},
//    { "type" : "function", "name" : "send", "inputs" : [ { "name" : "amount", "type" : "int256" } ] }]`
//if *ABIFile != "" {
//file, err := os.Open(*ABIFile)
//if err != nil {
//panic(err)
//}
//abiObject, err = abi.JSON(file)
//if err != nil {
//panic(err)
//}
//} else {
//abiObject, err = abi.JSON(strings.NewReader(jsonData))
//if err != nil {
//panic(err)
//}
//}

//import (
//	"flag"
//	"github.com/ethereum/go-ethereum/accounts/abi"
//)
//
//var (
//    ABIFile      *string
//    functionName *string
//)
//
//func init() {
//    ABIFile = flag.String("abi", "", "abi file path")
//    functionName = flag.String("function", "", "function name for test")
//}
//
//var (
//    Uint256, _ = abi.NewType("uint256", "", nil)
//)

//func main() {
//	flag.Parse()
//	var jsonData = `[{ "inputs": [{"internalType": "uint256","name": "a","type": "uint256"	},{	"internalType": "uint256","name": "b","type": "uint256"}],"stateMutability": "nonpayable","type": "constructor"},
//    { "type" : "function", "name" : "send", "inputs" : [ { "name" : "amount", "type" : "int256" } ] }]`
//	var abiObject abi.ABI
//	var err error
//	if *ABIFile != "" {
//		file, err := os.Open(*ABIFile)
//		if err != nil {
//			panic(fmt.Sprintf("File %s not exist\n", *ABIFile))
//		}
//		abiObject, err = abi.JSON(file)
//		if err != nil {
//			panic(err)
//		}
//	} else {
//		abiObject, err = abi.JSON(strings.NewReader(jsonData))
//		if err != nil {
//			panic(err)
//		}
//	}
//	methods := abiObject.Methods
//	for methodName, methodObj := range methods {
//		fmt.Println(methodName, methodObj)
//		var inputList []interface{}
//		for _, arg := range methodObj.Inputs {
//			argumentName := arg.Name
//			argumentType := arg.Type
//			fmt.Println(argumentName, argumentType)
//			parameter := NewParameterType(&argumentType)
//			switch {
//			case parameter.TypeName[:3] == "int":
//				varInt := NewIntType(parameter)
//				//fmt.Println("limit value:", varInt.MaxValue.String(), varInt.MinValue.String())
//				inputList = append(inputList, varInt.Generate())
//			default:
//				continue
//			}
//		}
//		inputPacked, err := abiObject.Pack(methodName, inputList...)
//		if err != nil {
//			panic(err)
//		}
//		fmt.Println(inputPacked)
//	}
//}
