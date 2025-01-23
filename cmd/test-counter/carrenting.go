package main

import (
	"context"
	"crypto/ecdsa"
	"entdemo/contracts"
	"entdemo/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"log"
	"math/big"
)

func RunCarRenting() {
	// parameters
	ethUrl := "http://localhost:8545"
	privateKeyHex := "18f9b8f25d49a65b7c2c5c99387fde36e11782d2aa025e25a33d8de991eacf6a"
	contractDeployedHex := "0xb7FB4083c32ff2D6b1fEE80BDEb5b256E7B4790e"

	contractPrivateKeyHex := common.HexToAddress(contractDeployedHex)

	ethClient := utils.NewEthClient(ethUrl)
	privateKey, fromAddress := utils.GetMetadataFromPrivateKeyHex(privateKeyHex)
	logrus.WithField("from", fromAddress).Info("admin private key")

	chainId, err := ethClient.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	//TODO ask ema for this
	// what are these other contrctuctor
	//contracts.NewCarRentingFilterer()

	// step 1 add a car from admin but store to owner 1
	owner1Pk, owner1From, err := utils.AddNewUser(ethClient, privateKey, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	owner2Pk, owner2From, err := utils.AddNewUser(ethClient, privateKey, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	renter1Pk, renter1From, err := utils.AddNewUser(ethClient, privateKey, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	renter2Pk, renter2From, err := utils.AddNewUser(ethClient, privateKey, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	balance1, err := ethClient.BalanceAt(context.Background(), owner1From, nil)
	if err != nil {
		log.Fatal(err)
	}
	balance2, err := ethClient.BalanceAt(context.Background(), owner2From, nil)
	if err != nil {
		log.Fatal(err)
	}
	balance3, err := ethClient.BalanceAt(context.Background(), renter1From, nil)
	if err != nil {
		log.Fatal(err)
	}
	balance4, err := ethClient.BalanceAt(context.Background(), renter2From, nil)
	if err != nil {
		log.Fatal(err)
	}
	balance5, err := ethClient.BalanceAt(context.Background(), fromAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	logrus.WithField("owner1PK", owner1Pk).WithField("owner1From", owner1From).WithField("balance", balance1.Int64()).Info("user1 details - owner")
	logrus.WithField("owner2PK", owner2Pk).WithField("owner2From", owner2From).WithField("balance", balance2.Int64()).Info("user2 details - owner")
	logrus.WithField("renter1Pk", renter1Pk).WithField("renter1From", renter1From).WithField("balance", balance3.Int64()).Info("user3 details - renter")
	logrus.WithField("renter2Pk", renter2Pk).WithField("renter2From", renter2From).WithField("balance", balance4.Int64()).Info("user4 details - renter")
	logrus.WithField("balance", balance5.Int64()).Info("admin details")

	logrus.WithField("chainId", chainId).Info("carrenting contract")

	// 1 initialize the smart contract

	contractObj, err := contracts.NewCarRenting(contractPrivateKeyHex, ethClient)
	if err != nil {
		log.Fatal(err)
	}
	AddCar(ethClient, owner1Pk, chainId, owner1From, contractObj, contractPrivateKeyHex)
	AddCar(ethClient, owner2Pk, chainId, owner2From, contractObj, contractPrivateKeyHex)
	//
	RentCar(ethClient, renter1Pk, renter1From, owner1From, chainId, contractObj, contractPrivateKeyHex)

	result, err := contractObj.Rents(&bind.CallOpts{
		From: owner1From,
	}, owner1From)
	logrus.WithField("result", result).Info("current car")

	//RentCar(ethClient, renter2Pk, renter2From, owner1From, chainId, contractObj, contractPrivateKeyHex)

}

func RentCar(ethClient *ethclient.Client, renterPk *ecdsa.PrivateKey, renter, owner common.Address, chainId *big.Int, contractObj *contracts.CarRenting, contractAddress common.Address) {
	transactionOps, err := utils.GetTransaction(ethClient, renterPk, chainId, renter)
	if err != nil {
		logrus.Fatal(err)
	}
	tx, err := contractObj.RentCar(transactionOps, owner)
	if err != nil {
		logrus.WithError(err).WithField("tx", tx).Error("failed to add a car")
		abi, _ := contracts.CarRentingMetaDataWithBin.GetAbi()
		binContract := common.FromHex(contracts.CarRentingMetaDataWithBin.Bin)
		reason, err := utils.GetRevertReason(context.Background(), ethClient, transactionOps, contractAddress, abi, binContract)
		logrus.WithField("reason", reason).WithError(err).Info("result of the get revert reason")
		//
		//reason, err = utils.GetTxRevertReason(context.Background(), ethClient, tx.Hash())
		//logrus.WithField("reason", reason).WithError(err).Info("GetTxRevertReason")

	}
	receipt, err := ethClient.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		logrus.WithError(err).WithField("tx", tx).Fatal("failed to send transaction")
	}
	logrus.WithField("receipt", receipt.Status).Info("car")
}

func AddCar(ethClient *ethclient.Client, owner1Pk *ecdsa.PrivateKey, chainId *big.Int, owner1From common.Address, contractObj *contracts.CarRenting, contractAddress common.Address) {
	transactionOps, err := utils.GetTransaction(ethClient, owner1Pk, chainId, owner1From)
	if err != nil {
		logrus.Fatal(err)
	}
	tx, err := contractObj.AddCar(transactionOps, contracts.CarRentingCar{
		IsRented:  false,
		Hp:        uint16(90),
		FuelLevel: uint16(10),
	})
	if err != nil {
		logrus.WithError(err).WithField("tx", tx).Error("failed to add a car")
		//abi, _ := contracts.CarRentingMetaDataWithBin.GetAbi()
		//binContract := common.FromHex(contracts.CarRentingMetaDataWithBin.Bin)
		//reason, err := utils.GetRevertReason(context.Background(), ethClient, transactionOps, contractAddress, abi, binContract)
		//logrus.WithField("reason", reason).WithError(err).Info("result of the get revert reason")
		//
		//reason, err = utils.GetTxRevertReason(context.Background(), ethClient, tx.Hash())
		//logrus.WithField("reason", reason).WithError(err).Info("GetTxRevertReason")

	}
	//err = ethClient.SendTransaction(context.Background(), tx)
	//if err != nil {
	//	logrus.WithError(err).WithField("tx", tx).Fatal("failed to send transaction")
	//}
}
