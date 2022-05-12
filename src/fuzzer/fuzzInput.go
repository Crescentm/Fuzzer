package fuzzer

import (
    "Fuzzer/src/parameterType"
    . "Fuzzer/src/pidType"
    "errors"
    "math/rand"
    "time"
)

func FuzzInput(input INPUTDATA, inputType INPUTTYPE) (INPUTDATA, int, error) {
    if input == nil {
        return input, 0, errors.New("no input")
    }
    inputNew := make(INPUTDATA, len(input))
    copy(inputNew, input)

    //which one input to variate
    choice := RandomChoose(len(input))

    parameter := parameterType.NewParameterType(&inputType[choice])

    switch {
    case parameter.TypeName[:3] == "int":
        varInt := parameterType.NewIntType(parameter)
        inputNew[choice] = varInt.Generate()
    case parameter.TypeName[:4] == "uint":
        varUint := parameterType.NewUintType(parameter)
        inputNew[choice] = varUint.Generate()
    }

    return inputNew, choice, nil
}

func RandomChoose(maxNum int) int {
    rand.Seed(time.Now().Unix())
    v := rand.Intn(maxNum)
    return v
}
