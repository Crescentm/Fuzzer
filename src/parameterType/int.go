package parameterType

import (
    "Fuzzer/src/random"
    "math/big"
    "strings"
)

type IntType struct {
    ParameterType
    MaxValue big.Int
    MinValue big.Int
}

func NewIntType(param *ParameterType) *IntType {
    var hexSize = param.Size / 4
    maxString := "0x7" + strings.Repeat("f", hexSize-1)
    minString := "-0x8" + strings.Repeat("0", hexSize-1)
    var maxValue, minValue big.Int
    maxValue.SetString(maxString, 0)
    minValue.SetString(minString, 0)
    return &IntType{ParameterType: *param, MaxValue: maxValue, MinValue: minValue}
}

func (self *IntType) Generate() interface{} {
    return random.RandBigInt(&self.MinValue, &self.MaxValue)
}
