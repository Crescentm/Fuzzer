package contract

import (
    "Fuzzer/src/pidType"
    "Fuzzer/src/tools"
    "fmt"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "math/big"
    "math/rand"
)

func SendInput(inputNew pidType.INPUTDATA, inputType pidType.INPUTTYPE, method string, abiObject abi.ABI, costSize int) (pidType.PID, string) {
    fmt.Println(tools.InputEncode(inputNew, method, abiObject))
    var cost pidType.COST
    for i := 0; i < 10; i++ {
        cost = append(cost, big.NewInt(rand.Int63()))
    }
    var pidNew = pidType.PID{InputData: inputNew, InputType: inputType, Cost: cost}
    return pidNew, ""
}
