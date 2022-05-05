package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
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
	//file, err := os.Open(*ABIFile)
	//if err != nil {
	//	panic(fmt.Sprintf("File %s not exist\n", *ABIFile))
	//}
	//
	//abi, err := abi.JSON(file)
	//if err != nil {
	//	panic(err)
	//}
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
	json := `[
    { "inputs": [{"internalType": "uint256","name": "a","type": "uint256"	},{	"internalType": "uint256","name": "b","type": "uint256"}],"stateMutability": "nonpayable","type": "constructor"},
    { "type" : "function", "name" : "send", "inputs" : [ { "name" : "amount", "type" : "uint256" } ] }
]`
	//method := abi.NewMethod("", "", abi.Constructor, "nonpayable", false, false, []abi.Argument{{"a", Uint256, false}, {"b", Uint256, false}}, nil)
	// Test from JSON
	Uint256, _ = abi.NewType("uint24", "", nil)
	fmt.Println(Uint256.Size, Uint256.GetType(), Uint256.String())
	Abi, err := abi.JSON(strings.NewReader(json))
	fmt.Println(Abi.Methods["send"].ID, Abi.Methods["send"].Sig)
	if err != nil {
		panic(err)
	}
	output, err := Abi.Pack("send", big.NewInt(1))
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println(output)
}
