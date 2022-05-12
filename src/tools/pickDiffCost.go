package tools

import (
	"Fuzzer/src/pidType"
	"Fuzzer/src/random"
	"math/big"
)

func PickDiffCost(cost1 pidType.COST, cost2 pidType.COST) (*big.Int, *big.Int) {
	for costID1, costValue1 := range cost1 {
		for costID2, costValue2 := range cost2 {
			if costID1 == costID2 {
				if costValue1[0].Cmp(costValue2[0]) != 0 && costValue1[0].Cmp(big.NewInt(0)) != 0 && costValue2[0].Cmp(big.NewInt(0)) != 0 {
					return costValue1[0], costValue2[0]
				}

				if costValue1[1].Cmp(costValue2[1]) != 0 && costValue1[1].Cmp(big.NewInt(0)) != 0 && costValue2[1].Cmp(big.NewInt(0)) != 0 {
					return costValue1[1], costValue2[1]
				}
			}
		}
	}
	return nil, nil
}

func PickRDiffCost(cost1 pidType.COST, cost2 pidType.COST) (*big.Int, *big.Int) {
	var Temp [][2]*big.Int
	for costID1, costValue1 := range cost1 {
		for costID2, costValue2 := range cost2 {
			if costID1 == costID2 {
				if costValue1[0].Cmp(costValue2[0]) != 0 && costValue1[0].Cmp(big.NewInt(0)) != 0 && costValue2[0].Cmp(big.NewInt(0)) != 0 {

					Temp = append(Temp, [2]*big.Int{costValue1[0], costValue2[0]})

				}

				if costValue1[1].Cmp(costValue2[1]) != 0 && costValue1[1].Cmp(big.NewInt(0)) != 0 && costValue2[1].Cmp(big.NewInt(0)) != 0 {

					Temp = append(Temp, [2]*big.Int{costValue1[1], costValue2[1]})

				}
			}
		}
	}

	r := random.NewRand()
	N := r.Intn(len(Temp))
	return Temp[N][0], Temp[N][1]
}
