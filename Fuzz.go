package Fuzzer

import (
	"math/big"
	"reflect"
)

type INPUT []reflect.Value
type COST []big.Int
type PIDS map[string]PID

type PID struct {
	input INPUT
	cost  COST
}

func Fuzz(program string, seeds any, fcost any) {
	var pids = make(PIDS)
	//pids = RunSeeds(seeds, program, fcost)
	for !Interrupted() {
		input, cost := PickInput(pids)
		energy := 0
		maxenergy := AssignEnergy(input)

		var predictedInput []reflect.Value = nil

		for energy < maxenergy || predictedInput != nil {
			var inputNew INPUT = nil
			var err error
			var choice int = 0

			if predictedInput != nil {
				copy(inputNew, predictedInput)
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

			if energy < maxenergy {
				costNew := pidNew.cost
				predictedInput = Predict(input, cost, inputNew, costNew, choice)
			}

			energy = energy + 1
		}
	}
}
