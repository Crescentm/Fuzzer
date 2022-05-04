package Fuzzer

import (
	"math/big"
)

func Predict(input INPUTDATA, cost COST, inputNew INPUTDATA, costNew COST, choice int) INPUTDATA {
	var predictedInput = make(INPUTDATA, len(input))
	copy(predictedInput, input)
	i := 0

	//todo
	for i = 0; i < len(cost); i++ {
		if cost[i].Cmp(&costNew[i]) != 0 && cost[i].Cmp(big.NewInt(0)) != 0 && costNew[i].Cmp(big.NewInt(0)) != 0 {
			break
		}
	}

	return predictedInput
}
