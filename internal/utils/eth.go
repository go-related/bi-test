package utils

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
)

func NewEthClient(host string) *ethclient.Client {
	client, err := ethclient.Dial(host)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Println("connecting to blockchain... ")
	blockNumber, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		logrus.Println("Failed to retrieve block number:", err)
		logrus.Fatal(err)
	}
	logrus.WithField("block", blockNumber.Number().Int64()).Info("blockchain connected ")

	return client
}

func GetTransactOpts(ethclient *ethclient.Client, fromAddress common.Address, auth *bind.TransactOpts) (*bind.TransactOpts, error) {
	nonce, err := ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	auth.GasPrice, err = ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return auth, nil
}
