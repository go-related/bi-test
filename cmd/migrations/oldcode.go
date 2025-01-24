package main

import (
	"context"
	"entdemo/contracts"
	"entdemo/ent"
	"entdemo/internal/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
)

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
	ethHost := "http://localhost:8545"
	privateKeyHex := "18f9b8f25d49a65b7c2c5c99387fde36e11782d2aa025e25a33d8de991eacf6a"

	ethClient := utils.NewEthClient(ethHost)

	// Getting the Chain ID
	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//DeployCounterContract(privateKeyHex, ethClient, chainId)
	DeployCarRentingContract(privateKeyHex, ethClient, chainId)
}

func DeployCounterContract(privateKeyHex string, ethclient *ethclient.Client, chainId *big.Int) {
	// Init account used to deploy contracts
	privateKey, fromAddress := utils.GetMetadataFromPrivateKeyHex(privateKeyHex)

	// Building the transactor object
	trnOptions, err := utils.GetTransaction(ethclient, privateKey, chainId, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	fromAddressResult, tx, bindedContract, err := contracts.DeployCounterHelper(trnOptions, ethclient)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("from", fromAddressResult).WithField("tx", tx).WithField("ctr", bindedContract).
		Info("deployed counter helper contract address")
	/*
		2025-01-21 12:02:14 INFO [01-21|11:02:14.602] Submitted contract creation hash=0x59530d65fcc0f62c903d3d4ba171a18104de4be512c00852ae120afdc6d69c26
		from=0x872B01f0dd1FC7b5AfEB05610e412E83836C82a6 nonce=59
		contract=0x9CBc233e8067c95C938BDda5666073FC9c084672 value=0
	*/
}

func DeployCarRentingContract(privateKeyHex string, ethclient *ethclient.Client, chainId *big.Int) {
	// Init account used to deploy contracts
	privateKey, fromAddress := utils.GetMetadataFromPrivateKeyHex(privateKeyHex)

	// Building the transactor object
	trnOptions, err := utils.GetTransaction(ethclient, privateKey, chainId, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	fromAddressResult, tx, bindedContract, err := contracts.DeployCarRentingHelper(trnOptions, ethclient)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.WithField("from", fromAddressResult).WithField("tx", tx).WithField("ctr", bindedContract).
		Info("deployed counter helper contract addres")

	/*
			Submitted contract creation              hash=0x59530d65fcc0f62c903d3d4ba171a18104de4be512c00852ae120afdc6d69c26
			from=0x872B01f0dd1FC7b5AfEB05610e412E83836C82a6 nonce=59
			contract=0xC4eA218335D8AbE0b7275a4631eA3F68C4D895C8
		value=0
	*/
}
