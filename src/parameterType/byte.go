package parameterType

import (
	"Fuzzer/src/random"
)

type ByteType struct {
	ParameterType
	Length int
}

func NewByteType(param *ParameterType) *ByteType {
	var size = param.Size
	return &ByteType{ParameterType: *param, Length: size}
}

func (t *ByteType) Generate() interface{} {
	return random.RandByte(t.Length)
}
