package fuzz

import (
	"Fuzzer/src/contract"
	. "Fuzzer/src/pidType"
	. "Fuzzer/src/tools"
	"flag"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"os"
	"strings"
)

var (
	abiObject abi.ABI
	inputType INPUTTYPE
	costSize  int = 10
	inputSize int
)

var (
	ABIFile *string
	//functionName *string
)

//func Init() {
//	ABIFile = flag.String("abi", "", "abi file path")
//	functionName = flag.String("function", "", "function name for test")
//}

func Fuzz(program string, seeds any, fcost any) {
	flag.Parse()
	var err error
	var jsonData = `[{ "inputs": [{"internalType": "uint256","name": "a","type": "uint256"	},{	"internalType": "uint256","name": "b","type": "uint256"}],"stateMutability": "nonpayable","type": "constructor"},
    { "type" : "function", "name" : "send", "inputs" : [ { "name" : "amount", "type" : "int256" } ] }]`
	if *ABIFile != "" {
		file, err := os.Open(*ABIFile)
		if err != nil {
			panic(err)
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

	for methodName, methodObject := range methods {
		inputSize, inputType = InputParse(methodObject)

		pids := RunSeeds(program, seeds, fcost, methodObject, inputType, abiObject, costSize)
		for !Interrupted() {
			inputData, cost := PickInput(pids)
			energy := 0
			maxEnergy := AssignEnergy(inputData)

			var predictedInputData INPUTDATA

			for energy < maxEnergy || predictedInputData != nil {
				var inputDataNew INPUTDATA
				var choice = 0

				if predictedInputData != nil {
					predictedInputData = nil
				} else {
					inputDataNew, choice, err = RandomInput(inputData, inputType)
					if err != nil {
						panic("can not fuzz input")
					}
				}

				pidNew, pidPath := contract.SendContract(inputDataNew, inputType, methodName, program, fcost, abiObject, costSize)

				if IsNew(pidPath, pids) {
					AddPid(pidNew, pidPath, pids)
				}

				if energy < maxEnergy {
					costNew := pidNew.Cost
					predictedInputData = Predict(inputData, cost, inputDataNew, costNew, choice)
				}

				energy = energy + 1
			}
		}
	}
}
