package main

import (
	"context"
	"crypto/ecdsa"
	"entdemo/contracts"
	"entdemo/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
)

func main() {
	runContract()
}

func runContract() {
	ethclient := utils.NewEthClient("http://localhost:8545")

	privateKeyHex := "18f9b8f25d49a65b7c2c5c99387fde36e11782d2aa025e25a33d8de991eacf6a"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logrus.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// we get this from node itself
	contractPrivateKeyHex := common.HexToAddress("0xC29c56Dbd04Df6b88a8c8F4167D84Fd9dBaEaefE")

	contractObj, err := contracts.NewCounter(contractPrivateKeyHex, ethclient)
	if err != nil {
		log.Fatal(err)
	}

	// get the tx option
	chainId, err := ethclient.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	auth, err = utils.GetTransactOpts(ethclient, fromAddress, auth)
	if err != nil {
		logrus.Fatal(err)
	}

	//
	input := big.NewInt(400)
	tx, err := contractObj.SetNumber(auth, input)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("tx", tx.Hash().Hex()).Info("set tx")

	tx, err = contractObj.Increment(auth)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("tx", tx.Hash().Hex()).Info("invrement tx")

	callOpts := bind.CallOpts{
		From: fromAddress,
	}
	result, err := contractObj.Number(&callOpts)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("number", result.Int64()).Info("invrement tx")

}
