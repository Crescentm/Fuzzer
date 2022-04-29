package Fuzzer

import (
	"errors"
	"reflect"
)

func FuzzInput(input INPUT) (INPUT, int, error) {
	if input == nil {
		return nil, 0, errors.New("no input")
	}

	var inputNew = make(INPUT, len(input))
	copy(inputNew, input)

	//which one input to variate
	choice := RandomChoose(len(input))

	switch input[choice].Kind() {

	//todo: add more type
	case reflect.Int8:
		random, err := Int8Random()
		if err != nil {
			return input, 0, err
		}
		inputNew[choice] = reflect.ValueOf(random)

	case reflect.Int16:
		random, err := Int16Random()
		if err != nil {
			return input, 0, err
		}
		inputNew[choice] = reflect.ValueOf(random)

	case reflect.Int32:
		random, err := Int32Random()
		if err != nil {
			return input, 0, err
		}
		inputNew[choice] = reflect.ValueOf(random)

	case reflect.Uint8:
		random, err := Uint8Random()
		if err != nil {
			return input, 0, err
		}
		inputNew[choice] = reflect.ValueOf(random)

	case reflect.Uint16:
		random, err := Uint16Random()
		if err != nil {
			return input, 0, err
		}
		inputNew[choice] = reflect.ValueOf(random)

	case reflect.Uint32:
		random, err := Uint32Random()
		if err != nil {
			return input, 0, err
		}
		inputNew[choice] = reflect.ValueOf(random)
	}

	return inputNew, choice, nil
}
