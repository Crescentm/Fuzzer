package Fuzzer

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
)

type INPUTDATA []interface{}
type INPUTTYPE []abi.Type
type COST []big.Int
type PIDS map[string]PID

type PID struct {
	inputData INPUTDATA
	inputType INPUTTYPE
	cost      COST
}

func Fuzz(program string, seeds any, fcost any) {
	var pids = make(PIDS)
	//pids = RunSeeds(seeds, program, fcost)
	for !Interrupted() {
		input, cost := PickInput(pids)
		energy := 0
		maxEnergy := AssignEnergy(input)

		var predictedInput INPUTDATA

		for energy < maxEnergy || predictedInput != nil {
			var inputNew INPUTDATA
			var err error
			var choice = 0

			if predictedInput != nil {
				predictedInput = nil
			} else {
				inputNew, choice, err = FuzzInput(input)
				if err != nil {
					panic("can not fuzz input")
				}
			}

			pidNew, pidPath := RunContract(inputNew, program, fcost)
			if IsNew(pidPath, pids) {
				Addpid(pidNew, pidPath, pids)
			}

			if energy < maxEnergy {
				costNew := pidNew.cost
				predictedInput = Predict(input, cost, inputNew, costNew, choice)
			}

			energy = energy + 1
		}
	}
}
