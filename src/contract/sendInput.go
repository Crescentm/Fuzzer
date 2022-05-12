package contract

import (
	"Fuzzer/src/jsonData"
	"Fuzzer/src/pidType"
	"Fuzzer/src/tools"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
)

func SendInput(inputNew pidType.INPUTDATA, method string, ContractAddress string, abi abi.ABI) (pidType.COST, [32]byte) {
	inputPacked := tools.InputEncode(inputNew, method, abi)
	dataString := "0x" + hex.EncodeToString(inputPacked)

	requestJSON, err := json.Marshal(jsonData.TransactionReq{Data: dataString, Address: ContractAddress})
	if err != nil {
		panic(err)
	}

	_, err = http.Post(
		"localhost/api/transactions",
		"application/json;charset=UTF-8",
		strings.NewReader(string(requestJSON)),
	)
	if err != nil {
		panic(err)
	}

	res, err := http.Post(
		"localhost/api/transactions/return",
		"application/json;charset=UTF-8",
		strings.NewReader(string(requestJSON)),
	)
	if err != nil {
		panic(err)
	}

	defer func() { _ = res.Body.Close() }()

	responseJSON, _ := ioutil.ReadAll(res.Body)

	var transaction jsonData.TransactionRes
	err = json.Unmarshal(responseJSON, &transaction)
	if err != nil {
		panic(err)
	}

	var pcInt []int64
	var costValue []*big.Int
	for _, pcstr := range transaction.PC {
		i, _ := strconv.ParseInt(pcstr, 16, 64)
		pcInt = append(pcInt, i)
	}
	for _, coststr := range transaction.Cost {
		var i = big.NewInt(0)
		i.SetString(coststr, 16)
		costValue = append(costValue, i)
	}

	cost := tools.Cost2map(pcInt, costValue)

	pathid := pidType.GeneratePathID(pcInt)
	return cost, pathid
}
