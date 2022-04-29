package Fuzzer

import (
	"math/big"
	"testing"
)

func TestBaz(t *testing.T) {
	a := big.NewInt(42)
	b := big.NewInt(2)
	c := big.NewInt(-11)

	println(Baz(a, b, c))

}
