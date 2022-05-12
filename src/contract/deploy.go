package contract

import (
	. "Fuzzer/src/jsonData"
	"Fuzzer/src/tools"
	"encoding/hex"
	"encoding/json"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Deploy(BinaryFile string, ABIFile string, AbiObject abi.ABI) (DeployRes, error) {
	var responseData DeployRes
	var abiArray interface{}

	binaryFileStream, err := os.Open(BinaryFile)
	if err != nil {
		return responseData, err
	}

	ABIFileStream, err := os.Open(ABIFile)
	if err != nil {
		return responseData, err
	}

	byteCode, err := ioutil.ReadAll(binaryFileStream)
	if err != nil {
		return responseData, err
	}

	byteAbi, err := ioutil.ReadAll(ABIFileStream)
	if err != nil {
		return responseData, err
	}

	err = binaryFileStream.Close()
	if err != nil {
		return responseData, err
	}
	err = ABIFileStream.Close()
	if err != nil {
		return responseData, err
	}

	err = json.Unmarshal(byteAbi, &abiArray)
	if err != nil {
		return responseData, err
	}

	ConstructorObj := AbiObject.Constructor

	inputData := tools.RandomInput(ConstructorObj)

	inputPacked, err := ConstructorObj.Inputs.PackValues(inputData)
	if err != nil {
		return responseData, err
	}

	dataString := "0x" + hex.EncodeToString(inputPacked)
	byteString := "0x" + hex.EncodeToString(byteCode)

	var requestData = DeployReq{AbiArray: abiArray, Bytecode: byteString, Data: dataString}

	requestJSON, err := json.Marshal(requestData)
	if err != nil {
		return responseData, err
	}

	res, err := http.Post(
		"localhost/api/contracts",
		"application/json;charset=UTF-8",
		strings.NewReader(string(requestJSON)),
	)
	if err != nil {
		return responseData, err
	}
	defer func() { _ = res.Body.Close() }()

	responseJSON, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(responseJSON, &responseData)
	if err != nil {
		return responseData, err
	}

	return responseData, nil

}
