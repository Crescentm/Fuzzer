package Fuzzer

import (
	"math/rand"
	"time"
)

func RandomChoose(maxnum int) int {
	rand.Seed(time.Now().Unix())
	v := rand.Intn(maxnum)
	return v
}

// AssignEnergy todo
func AssignEnergy(input INPUTDATA) int {
	return 3
}

// Interrupted todo
func Interrupted() bool {
	return false
}

// PickInput base on map random
func PickInput(pids PIDS) (INPUTDATA, COST) {
	for _, pid := range pids {
		return pid.inputdata, pid.cost
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

func Addpid(pidnew PID, pidPath string, pids PIDS) {
	pids[pidPath] = pidnew
}
