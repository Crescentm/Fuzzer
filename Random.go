package Fuzzer

import (
    "bytes"
    "crypto/rand"
    "encoding/binary"
)

func Int8Random() (int8, error) {
    randomNumByte := make([]byte, 1)
    var randomNum int8
    _, err := rand.Read(randomNumByte)
    if err != nil {
        return 0, err
    }

    bytesBuffer := bytes.NewBuffer(randomNumByte)
    err = binary.Read(bytesBuffer, binary.LittleEndian, &randomNum)
    if err != nil {
        return 0, err
    }

    return randomNum, nil
}

func Int16Random() (int16, error) {
    randomNumByte := make([]byte, 2)
    var randomNum int16
    _, err := rand.Read(randomNumByte)
    if err != nil {
        return 0, err
    }

    bytesBuffer := bytes.NewBuffer(randomNumByte)
    err = binary.Read(bytesBuffer, binary.LittleEndian, &randomNum)
    if err != nil {
        return 0, err
    }

    return randomNum, nil
}

func Int32Random() (int32, error) {
    randomNumByte := make([]byte, 4)
    var randomNum int32
    _, err := rand.Read(randomNumByte)
    if err != nil {
        return 0, err
    }

    bytesBuffer := bytes.NewBuffer(randomNumByte)
    err = binary.Read(bytesBuffer, binary.LittleEndian, &randomNum)
    if err != nil {
        return 0, err
    }

    return randomNum, nil
}

func Uint8Random() (uint8, error) {
    randomNumByte := make([]byte, 1)
    var randomNum uint8
    _, err := rand.Read(randomNumByte)
    if err != nil {
        return 0, err
    }

    bytesBuffer := bytes.NewBuffer(randomNumByte)
    err = binary.Read(bytesBuffer, binary.LittleEndian, &randomNum)
    if err != nil {
        return 0, err
    }

    return randomNum, nil
}

func Uint16Random() (uint16, error) {
    randomNumByte := make([]byte, 2)
    var randomNum uint16
    _, err := rand.Read(randomNumByte)
    if err != nil {
        return 0, err
    }

    bytesBuffer := bytes.NewBuffer(randomNumByte)
    err = binary.Read(bytesBuffer, binary.LittleEndian, &randomNum)
    if err != nil {
        return 0, err
    }

    return randomNum, nil
}

func Uint32Random() (uint32, error) {
    randomNumByte := make([]byte, 4)
    var randomNum uint32
    _, err := rand.Read(randomNumByte)
    if err != nil {
        return 0, err
    }

    bytesBuffer := bytes.NewBuffer(randomNumByte)
    err = binary.Read(bytesBuffer, binary.LittleEndian, &randomNum)
    if err != nil {
        return 0, err
    }

    return randomNum, nil
}
