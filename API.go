package Fuzzer

import (
	"math/big"
)

func Baz(a *big.Int, b *big.Int, c *big.Int) string {
	cost := make(COST, 26)
	d := new(big.Int).Add(c, b)
	var res big.Int
	res.Sub(big.NewInt(1), d)
	cost[4].Abs(&res)
	if d.Cmp(big.NewInt(1)) == -1 {
		if b.Cmp(big.NewInt(3)) == -1 {
			return "1"
		}

		if a.Cmp(big.NewInt(42)) == 0 {
			return "2"
		}
		return "3"
	} else {
		if c.Cmp(big.NewInt(42)) == -1 {
			return "4"
		}
		return "5"
	}
}
