package parameterType

import (
	"Fuzzer/src/random"
)

type BoolType struct {
	ParameterType
}

func NewBoolType(param *ParameterType) *BoolType {
	return &BoolType{ParameterType: *param}
}

func (t *BoolType) Generate() interface{} {
	return random.RandBool()
}
