package Fuzzer

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	Uint256, _    = abi.NewType("uint256", "", nil)
	Uint32, _     = abi.NewType("uint32", "", nil)
	Uint16, _     = abi.NewType("uint16", "", nil)
	String, _     = abi.NewType("string", "", nil)
	Bool, _       = abi.NewType("bool", "", nil)
	Bytes, _      = abi.NewType("bytes", "", nil)
	Bytes32, _    = abi.NewType("bytes32", "", nil)
	Address, _    = abi.NewType("address", "", nil)
	Uint64Arr, _  = abi.NewType("uint64[]", "", nil)
	AddressArr, _ = abi.NewType("address[]", "", nil)
	Int8, _       = abi.NewType("int8", "int8", nil)
	Int16, _      = abi.NewType("int16", "int16", nil)
	Int32, _      = abi.NewType("int32", "int32", nil)
	Int256, _     = abi.NewType("int256", "", nil)
)

// RunContract todo
func RunContract(input INPUTDATA, program string, fcost any) (PID, string) {

}
