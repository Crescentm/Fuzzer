package Fuzzer

import (
	"errors"
	"strings"
)

func FuzzInput(input INPUTDATA, inputType INPUTTYPE) (INPUTDATA, int, error) {
	if input == nil {
		return input, 0, errors.New("no input")
	}
	inputNew := make(INPUTDATA, len(input))
	copy(inputNew, input)
	//which one input to variate
	choice := RandomChoose(len(input))

	if strings.HasPrefix(inputType[choice].String(), "int") {
		randNum, err := IntRandom(uint(inputType[choice].Size))
		if err != nil {
			return nil, 0, err
		}
		inputNew[choice] = randNum
	} else if strings.HasPrefix(inputType[choice].String(), "uint") {
		randNum, err := IntRandom(uint(inputType[choice].Size))
		if err != nil {
			return nil, 0, err
		}
		inputNew[choice] = randNum
	} else {
		inputNew[choice] = input[choice]
	}

	return inputNew, choice, nil
}
