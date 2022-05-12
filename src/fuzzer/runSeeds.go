package fuzzer

import (
	"Fuzzer/src/contract"
	"Fuzzer/src/pidType"
	"Fuzzer/src/tools"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func RunSeeds(seeds any, method abi.Method, inputType pidType.INPUTTYPE) pidType.PIDS {
	var pids = make(pidType.PIDS)
	inputList := tools.RandomInput(method)

	cost, path := contract.SendInput(inputList, method.Name, ContractAddress, AbiObject)

	pidNew := pidType.Constructor(inputList, inputType, cost)

	pids[path] = pidNew
	return pids
}
