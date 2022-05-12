package fuzzer

import (
	. "Fuzzer/src/contract"
	. "Fuzzer/src/pidType"
	. "Fuzzer/src/tools"
)

func FuctionFuzz(fuctionName string, seeds any) {
	var err error
	methodObject := AbiObject.Methods[fuctionName]

	inputType := InputParse(methodObject)
	pids := RunSeeds(seeds, methodObject, inputType)

	for !Interrupted() {
		inputData, cost := PickInput(pids)
		energy := 0
		maxEnergy := AssignEnergy(inputData)

		var predictedInputData INPUTDATA

		for energy < maxEnergy || predictedInputData != nil {
			var inputDataNew INPUTDATA
			var choice = 0

			if predictedInputData != nil {
				inputDataNew = make(INPUTDATA, len(predictedInputData))
				copy(inputDataNew, predictedInputData)
				predictedInputData = nil
			} else {
				inputDataNew, choice, err = FuzzInput(inputData, inputType)
				if err != nil {
					panic("can not fuzzer input")
				}
			}

			pidCost, pidPath := SendInput(inputDataNew, fuctionName, ContractAddress, AbiObject)
			pidNew := Constructor(inputDataNew, inputType, pidCost)

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
