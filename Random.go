package Fuzzer

import (
	"crypto/rand"
	"math/big"
)

func UintRandom(len uint) (*big.Int, error) {
	max := new(big.Int)
	max.Lsh(big.NewInt(1), len)

	randNum, err := rand.Int(rand.Reader, max)

	if err != nil {
		return nil, err
	}
	return randNum, nil
}

func IntRandom(len uint) (*big.Int, error) {
	max := new(big.Int)
	max.Lsh(big.NewInt(1), len-1)

	min := new(big.Int)
	min.Lsh(big.NewInt(-1), len-1)

	randNum, err := rand.Int(rand.Reader, max.Sub(max, min))
	randNum.Add(randNum, min)

	if err != nil {
		return nil, err
	}
	return randNum, nil

}
