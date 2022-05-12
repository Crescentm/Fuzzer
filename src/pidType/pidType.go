package pidType

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
)

type INPUTDATA []interface{}
type INPUTTYPE []abi.Type
type COST map[int64][]*big.Int
type PIDS map[[32]byte]PID

type PID struct {
	InputData INPUTDATA
	InputType INPUTTYPE
	Cost      COST
}

func Constructor(inputdata INPUTDATA, inputType INPUTTYPE, cost COST) PID {
	pid := PID{InputData: inputdata, InputType: inputType, Cost: cost}
	return pid
}

// PickInput base on map random
func PickInput(pids PIDS) (INPUTDATA, COST) {
	for _, pid := range pids {
		return pid.InputData, pid.Cost
	}
	return nil, nil
}

func IsNew(pidPath [32]byte, pids PIDS) bool {
	_, ok := pids[pidPath]
	if ok {
		return false
	} else {
		return true
	}
}

func AddPid(pidNew PID, pidPath [32]byte, pids PIDS) {
	pids[pidPath] = pidNew
}

func GeneratePathID(w64 []int64) [32]byte {
	wbuf := new(bytes.Buffer)
	err := binary.Write(wbuf, binary.LittleEndian, w64)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	db := wbuf.Bytes()

	sum := sha256.Sum256(db)

	return sum
}
