package random

import (
	"math/big"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

var (
	GlobalRandSeed = int64(time.Now().Nanosecond())
	GlobalRand     = rand.New(rand.NewSource(GlobalRandSeed))
)

type FuzzRand struct {
	Seed      int64
	LastChose int
	R         *rand.Rand
}

func NewFuzzRand() *FuzzRand {
	randomOffset := int64(GlobalRand.Intn(time.Now().Nanosecond()))
	seed := time.Now().UnixNano() + randomOffset
	return &FuzzRand{Seed: seed, LastChose: -1, R: rand.New(rand.NewSource(seed))}
}

func NewRand() *rand.Rand {
	randomOffset := int64(GlobalRand.Intn(time.Now().Nanosecond()))
	seed := time.Now().UnixNano() + randomOffset
	r := rand.New(rand.NewSource(seed))
	return r
}

func RandBigInt(minValue, maxValue *big.Int) *big.Int {
	r := NewRand()
	sub := new(big.Int).Sub(maxValue, minValue)
	return new(big.Int).Add(new(big.Int).Rand(r, sub), minValue)
}

func RandBool() bool {
	r := NewRand()
	return r.Intn(2) == 0
}

func RandByte(size int) []byte {
	r := NewRand()
	var result = make([]byte, size)
	r.Read(result)
	return result
}

func RandSeq(n int) string {
	r := NewRand()
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[r.Intn(len(letters))]
	}
	return string(b)
}
