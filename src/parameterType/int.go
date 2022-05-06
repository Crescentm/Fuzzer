package parameterType

import (
	"Fuzzer/src/random"
	"math/big"
)

type IntType struct {
	ParameterType
	MaxValue big.Int
	MinValue big.Int
}

func NewIntType(param *ParameterType) *IntType {
	var size = param.Size
	var maxValue, minValue big.Int
	maxValue.Lsh(big.NewInt(1), uint(size-1))
	minValue.Lsh(big.NewInt(-1), uint(size-1))
	return &IntType{ParameterType: *param, MaxValue: maxValue, MinValue: minValue}
}

func (t *IntType) Generate() interface{} {
	return random.RandBigInt(&t.MinValue, &t.MaxValue)
}
