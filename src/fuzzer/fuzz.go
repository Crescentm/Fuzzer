package fuzzer

import (
	"Fuzzer/src/contract"
	"flag"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"os"
)

var (
	AbiObject       abi.ABI
	ABIFile         *string
	SeedFile        *string
	BinaryFile      *string
	ContractAddress string
)

func Init() {
	ABIFile = flag.String("abi", "abi.json", "abi file path")
	SeedFile = flag.String("seed", "seed.json", "seed file path")
	BinaryFile = flag.String("binary", "codefile.bin", "binary file path")
}

func Fuzz() {
	flag.Parse()

	ABIFileStream, err := os.Open(*ABIFile)
	if err != nil {
		panic("No such abi file")
	}

	AbiObject, err = abi.JSON(ABIFileStream)
	if err != nil {
		panic("Can not parse abi file")
	}

	err = ABIFileStream.Close()
	if err != nil {
		panic(err)
	}

	res, err := contract.Deploy(*BinaryFile, *ABIFile, AbiObject)
	if err != nil {
		panic(err)
	}
	ContractAddress = res.Address

	for fuctionName := range AbiObject.Methods {
		FuctionFuzz(fuctionName, SeedFile)
	}

}
