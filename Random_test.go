package Fuzzer

import (
    "fmt"
    "testing"
)

func TestInt8Random(t *testing.T) {
    var out []int8
    for i := 0; i < 10; i++ {
        num, err := Int8Random()
        if err != nil {
            t.Fail()
        }
        out = append(out, num)
    }
    for _, i := range out {
        if i < -128 || i > 127 {
            t.Fail()
        }
        fmt.Printf("%d, ", i)

    }
    fmt.Println()
}
