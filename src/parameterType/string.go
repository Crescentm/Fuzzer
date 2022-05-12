package parameterType

import (
	"Fuzzer/src/random"
)

type StringType struct {
	ParameterType
	Length int
}

func NewStringType(param *ParameterType, size int) *StringType {
	return &StringType{ParameterType: *param, Length: size}
}

func (t *StringType) Generate() interface{} {
	return random.RandSeq(t.Length)
}
