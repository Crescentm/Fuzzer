package fuzzer

import (
	. "Fuzzer/src/pidType"
	"Fuzzer/src/tools"
	"math/big"
)

func Predict(input INPUTDATA, cost COST, inputNew INPUTDATA, costNew COST, choice int) INPUTDATA {
	var predictedInput = make(INPUTDATA, len(input))
	copy(predictedInput, input)

	cost1, cost2 := tools.PickDiffCost(cost, costNew)

	input1 := input[choice].(*big.Int)
	input2 := inputNew[choice].(*big.Int)
	argNew := big.NewInt(0)
	numerator := big.NewInt(0)
	denominator := big.NewInt(0)
	mul1 := big.NewInt(0)
	mul2 := big.NewInt(0)
	mul1.Mul(cost1, input2)
	mul2.Mul(cost2, input1)

	numerator.Sub(mul1, mul2)
	denominator.Sub(cost1, cost2)

	argNew.Div(numerator, denominator)

	predictedInput[choice] = argNew
	return predictedInput
}
