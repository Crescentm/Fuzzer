package parameterType

import (
	"Fuzzer/src/random"
	"math/big"
)

type UintType struct {
	ParameterType
	MaxValue big.Int
	MinValue big.Int
}

func NewUintType(param *ParameterType) *IntType {
	var size = param.Size
	var maxValue, minValue big.Int
	// FIX: wrong maxValue, should be big
	maxValue.Lsh(big.NewInt(1), uint(size+1))
	minValue.SetUint64(0)
	return &IntType{ParameterType: *param, MaxValue: maxValue, MinValue: minValue}
}

func (t *UintType) Generate() interface{} {
	return random.RandBigInt(&t.MinValue, &t.MaxValue)
}
