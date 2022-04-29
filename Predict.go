package Fuzzer

import (
	"math/big"
	"reflect"
)

func Predict(input INPUT, cost COST, inputNew INPUT, costNew COST, choice int) INPUT {
	var Predictedinput = make(INPUT, len(input))
	var i int = 0
	for i = 0; i < len(cost); i++ {
		if cost[i].Cmp(&costNew[i]) != 0 && cost[i].Cmp(big.NewInt(0)) != 0 && costNew[i].Cmp(big.NewInt(0)) != 0 {
			break
		}
	}

	copy(Predictedinput, input)
	//var a, b big.Int
	//switch Predictedinput[choice].Type() {
	//case reflect.rt, reflect.Int16, reflect.Int32, reflect.Int64:
	//	{
	//		a.SetInt64(input[choice].Int())
	//		b.SetInt64(inputNew[choice].Int())
	//	}
	//case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
	//	{
	//		a.SetUint64(input[choice].Uint())
	//		b.SetUint64(inputNew[choice].Uint())
	//	}
	//
	//case big.Int:
	//
	//}
	Predictedinput[choice] = reflect.ValueOf()

	return Predictedinput
}
