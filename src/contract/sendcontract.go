package contract

import (
	"Fuzzer/src/pidType"
	"Fuzzer/src/tools"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func SendContract(inputNew pidType.INPUTDATA, inputType pidType.INPUTTYPE, method string, program string, fcost any, abiObject abi.ABI, costSize int) (pidType.PID, string) {
	fmt.Println(tools.InputEncode(inputNew, method, abiObject))
	cost := make(pidType.COST, costSize)

	var pidNew = pidType.PID{InputData: inputNew, InputType: inputType, Cost: cost}
	return pidNew, ""
}
