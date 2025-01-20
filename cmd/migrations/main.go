package main

import (
	"context"
	"crypto/ecdsa"
	"entdemo/contracts"
	"entdemo/ent"
	"entdemo/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
)

func main() {
	RunMigrations()
	logrus.Info("finished running migrations")
}

func RunMigrations() {
	client, err := ent.Open("postgres", "host=localhost port=5432 user=blockinvest dbname=test password=Test123 sslmode=disable")
	if err != nil {
		logrus.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		logrus.Fatalf("failed creating schema resources: %v", err)
	}

	// deploy contracts
	ethclient := utils.NewEthClient("http://localhost:8545")

	// Getting the Chain ID
	chainId, err := ethclient.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Init account used to deploy contracts
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

	// Building the transactor object
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}

	auth, err = utils.GetTransactOpts(ethclient, fromAddress, auth)
	if err != nil {
		logrus.Fatal(err)
	}

	fromAddressResult, tx, bindedContract, err := contracts.DeployCounterHelper(auth, ethclient)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("from", fromAddressResult).WithField("tx", tx).WithField("ctr", bindedContract).
		Info("deployed counter helper contract addres")
}
