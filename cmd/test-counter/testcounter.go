package main

import (
	"context"
	"entdemo/contracts"
	"entdemo/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
)

func main() {
	runCounterContract()
}

func runCounterContract() {
	ethUrl := "http://localhost:8545"
	ethclient := utils.NewEthClient(ethUrl)

	privateKeyHex := "18f9b8f25d49a65b7c2c5c99387fde36e11782d2aa025e25a33d8de991eacf6a"
	contractDeployedHex := "0xC29c56Dbd04Df6b88a8c8F4167D84Fd9dBaEaefE"

	privateKey, fromAddress := utils.GetMetadataFromPrivateKeyHex(privateKeyHex)
	contractPrivateKeyHex := common.HexToAddress(contractDeployedHex)
	chainId, err := ethclient.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	contractObj, err := contracts.NewCounter(contractPrivateKeyHex, ethclient)
	if err != nil {
		log.Fatal(err)
	}

	// operation 1

	transactionOps, err := utils.GetTransaction(ethclient, privateKey, chainId, fromAddress)
	if err != nil {
		logrus.Fatal(err)
	}
	input := big.NewInt(400)
	tx, err := contractObj.SetNumber(transactionOps, input)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("tx", tx.Hash().Hex()).Info("operation 1 set tx")

	// operation2
	transactionOps, err = utils.GetTransaction(ethclient, privateKey, chainId, fromAddress)
	if err != nil {
		logrus.Fatal(err)
	}
	tx, err = contractObj.Increment(transactionOps)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("tx", tx.Hash().Hex()).Info("operation 2 increment tx")

	// operation3
	callOpts := bind.CallOpts{
		From: fromAddress,
	}
	result, err := contractObj.Number(&callOpts)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("number", result.Int64()).Info("operation 3 get the number back")

}
