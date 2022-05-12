package tools

import (
	"Fuzzer/src/pidType"
	"math/big"
)

func Cost2map(pc []int64, costValue []*big.Int) pidType.COST {
	var cost pidType.COST
	for i := 0; i < len(pc)/2; i++ {
		cost[pc[i]][0] = costValue[i]
	}
	for i := len(pc) / 2; i < len(pc); i++ {
		cost[pc[i]][1] = costValue[i]
	}
	return cost
}
