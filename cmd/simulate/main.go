package main

import (
	"entdemo/blockchain/smartcontracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
)

func main() {
	runSimulation()
}

func runSimulation() {
	contractAddress := "0x3F0D4B8B33b10c5C723639E27675D862CCE96095"
	ethUrl := "http://localhost:8545"
	adminPrivateKeyHex := "18f9b8f25d49a65b7c2c5c99387fde36e11782d2aa025e25a33d8de991eacf6a"

	ethClient, err := smartcontracts.NewEthClient(ethUrl)
	if err != nil {
		logrus.Fatalf("failed to connect to eth client: %v", err)
	}
	contractAdr := common.HexToAddress(contractAddress)
	adminPk, adminAddress, err := smartcontracts.GetMetadataFromPrivateKeyHex(adminPrivateKeyHex)
	if err != nil {
		logrus.Fatalf("failed to decode admin private key: %v", err)
	}

	err = ethClient.RunSimulation(contractAdr, adminPk, adminAddress)
	if err != nil {
		logrus.Fatalf("failed to run simulation: %v", err)
	}
}
