package random

import (
	"fmt"
	"testing"
)

func TestRandByte(t *testing.T) {
	for i := 0; i < 10; i++ {
		r := RandByte(12)
		fmt.Println(r)
	}
}

func TestRandSeq(t *testing.T) {
	for i := 0; i < 10; i++ {
		ss := RandSeq(10)
		fmt.Println(ss)
	}
}
