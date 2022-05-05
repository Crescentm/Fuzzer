package parameterType

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Generator interface {
	Generate() interface{}
}

type ParameterType struct {
	// should be a generator
	TypeName string
	Size     int
}

func NewParameterType(p *abi.Type) *ParameterType {
	return &ParameterType{TypeName: p.String(), Size: p.Size}
}
