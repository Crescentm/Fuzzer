package main

import (
	. "Fuzzer/src/parameterType"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"os"
	"strings"
)

var (
	ABIFile      *string
	functionName *string
)

func init() {
	ABIFile = flag.String("abi", "", "abi file path")
	functionName = flag.String("function", "", "function name for test")
}

var (
	Uint256, _ = abi.NewType("uint256", "", nil)
)

func main() {
	flag.Parse()
	var jsonData = `[{ "inputs": [{"internalType": "uint256","name": "a","type": "uint256"	},{	"internalType": "uint256","name": "b","type": "uint256"}],"stateMutability": "nonpayable","type": "constructor"},
    { "type" : "function", "name" : "send", "inputs" : [ { "name" : "amount", "type" : "int256" } ] }]`
	var abiObject abi.ABI
	var err error
	if *ABIFile != "" {
		file, err := os.Open(*ABIFile)
		if err != nil {
			panic(fmt.Sprintf("File %s not exist\n", *ABIFile))
		}
		abiObject, err = abi.JSON(file)
		if err != nil {
			panic(err)
		}
	} else {
		abiObject, err = abi.JSON(strings.NewReader(jsonData))
		if err != nil {
			panic(err)
		}
	}
	methods := abiObject.Methods
	for methodName, methodObj := range methods {
		fmt.Println(methodName, methodObj)
		var inputList []interface{}
		for _, arg := range methodObj.Inputs {
			argumentName := arg.Name
			argumentType := arg.Type
			fmt.Println(argumentName, argumentType)
			parameter := NewParameterType(&argumentType)
			switch {
			case parameter.TypeName[:3] == "int":
				varInt := NewIntType(parameter)
				//fmt.Println("limit value:", varInt.MaxValue.String(), varInt.MinValue.String())
				inputList = append(inputList, varInt.Generate())
			default:
				continue
			}
		}
		inputPacked, err := abiObject.Pack(methodName, inputList...)
		if err != nil {
			panic(err)
		}
		fmt.Println(inputPacked)
	}
	//var args [][32]byte
	//for i := 0; i < flag.NArg(); i++ {
	//	s := flag.Arg(i)
	//	fmt.Println("args:", i, s)
	//	if len(s) > 2 && s[:2] == "0x" {
	//		argLeftJust := s[2:]
	//		if len(s)&1 == 1 {
	//			argLeftJust = "0" + argLeftJust
	//		}
	//		fmt.Println(s[2:], argLeftJust)
	//		argBytes := common.Hex2BytesFixed(argLeftJust, 32)
	//		fmt.Println(argBytes)
	//		argBytes = common.LeftPadBytes(argBytes, 32)
	//		fmt.Println(argBytes)
	//		//args = append(args, argBytes)
	//	}
	//}
	//method := abi.NewMethod("", "", abi.Constructor, "nonpayable", false, false, []abi.Argument{{"a", Uint256, false}, {"b", Uint256, false}}, nil)
	// Test from JSON
	//Uint256, _ = abi.NewType("uint24", "", nil)
	//fmt.Println(Uint256.Size, Uint256.GetType(), Uint256.String())
	//fmt.Println(Abi.Methods["send"].ID, Abi.Methods["send"].Sig)
	//if err != nil {
	//    panic(err)
	//}
	//output, err := Abi.Pack("send", big.NewInt(1))
	//if err != nil {
	//    fmt.Printf("%s\n", err)
	//}
	//fmt.Println(output)
}
