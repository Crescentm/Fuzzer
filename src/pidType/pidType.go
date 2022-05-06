package pidType

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
	"math/big"
)

type INPUTDATA []interface{}
type INPUTTYPE []abi.Type
type COST []big.Int
type PIDS map[string]PID

type PID struct {
	InputData INPUTDATA
	InputType INPUTTYPE
	Cost      COST
}

// PickInput base on map random
func PickInput(pids PIDS) (INPUTDATA, COST) {
	for _, pid := range pids {
		return pid.InputData, pid.Cost
	}
	return nil, nil
}

func IsNew(pidPath string, pids PIDS) bool {
	_, ok := pids[pidPath]
	if ok {
		return true
	} else {
		return false
	}
}

func AddPid(pidNew PID, pidPath string, pids PIDS) {
	pids[pidPath] = pidNew
}
