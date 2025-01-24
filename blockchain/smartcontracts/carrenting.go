package smartcontracts

import (
	"context"
	"crypto/ecdsa"
	"entdemo/blockchain/user"
	"entdemo/contracts"
	"entdemo/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"log"
	"time"
)

func (cl *EthClient) RunSimulation(contractAddress common.Address, adminPk *ecdsa.PrivateKey, adminFrom common.Address) error {

	user1Pk, user1Address, err := cl.addUser(adminPk, adminFrom)
	if err != nil {
		return errors.Wrap(err, "failed to add user1 to the chain")
	}
	user2Pk, user2Address, err := cl.addUser(adminPk, adminFrom)
	if err != nil {
		return errors.Wrap(err, "failed to add user2 to the chain")
	}
	user3Pk, user3Address, err := cl.addUser(adminPk, adminFrom)
	if err != nil {
		return errors.Wrap(err, "failed to add user3 to the chain")
	}
	user4Pk, user4Address, err := cl.addUser(adminPk, adminFrom)
	if err != nil {
		return errors.Wrap(err, "failed to add user4 to the chain")
	}
	time.Sleep(time.Second * 4)
	err = cl.AddCar(user1Pk, user1Address, contractAddress)
	if err != nil {
		return errors.Wrap(err, "failed to store car for the user 1")
	}
	err = cl.AddCar(user2Pk, user2Address, contractAddress)
	if err != nil {
		return errors.Wrap(err, "failed to store car for the user 2")
	}

	err = cl.RentCar(user3Pk, user3Address, user1Address, contractAddress)
	if err != nil {
		return errors.Wrap(err, "failed to rent car for the user 3")
	}

	err = cl.RentCar(user4Pk, user4Address, user1Address, contractAddress)
	if err == nil {
		return errors.Wrap(err, "didn't fail to rent a rented car ")
	}

	return nil
}
func (cl *EthClient) RentCar(renterPk *ecdsa.PrivateKey, renterFrom common.Address, owner common.Address, contractAddress common.Address) error {
	transactionOps, err := cl.CreateTransaction(renterPk, renterFrom, false)
	if err != nil {
		err = errors.Wrap(err, "failed to create transaction")
	}
	contractObj, err := contracts.NewCarRenting(contractAddress, cl.client)
	if err != nil {
		log.Fatal(err)
	}
	carTx, err := contractObj.RentCar(transactionOps, owner)
	if err != nil {
		logrus.WithError(err).Error("failed to rent car")
		//return cl.GetError(carTx, transactionOps, contractAddress)
		return err
	}
	cl.WaitForTx(carTx)

	result, err := contractObj.Rents(&bind.CallOpts{}, owner)
	if err != nil {
		return errors.Wrap(err, "failed to get the car details")
	}
	logrus.WithField("car", result).Info("added car information")
	return err
}

func (cl *EthClient) AddCar(ownerPk *ecdsa.PrivateKey, ownerFrom common.Address, contractAddress common.Address) error {
	transactionOps, err := cl.CreateTransaction(ownerPk, ownerFrom, false)
	if err != nil {
		err = errors.Wrap(err, "failed to create transaction")
	}
	contractObj, err := contracts.NewCarRenting(contractAddress, cl.client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := contractObj.AddCar(transactionOps, contracts.CarRentingCar{
		IsRented:  false,
		Hp:        uint16(90),
		FuelLevel: uint16(10),
	})
	if err != nil {
		return errors.Wrap(err, "failed to add car to the chain")
	}
	cl.WaitForTx(tx)

	// check what we added
	result, err := contractObj.Rents(&bind.CallOpts{}, ownerFrom)
	if err != nil {
		return errors.Wrap(err, "failed to get the car details")
	}
	logrus.WithField("car", result).Info("added car information")
	return err
}

func (cl *EthClient) GetError(tx *types.Transaction, opts *bind.TransactOpts, contractAddress common.Address) error {
	abi, _ := contracts.CarRentingMetaDataWithBin.GetAbi()
	binContract := common.FromHex(contracts.CarRentingMetaDataWithBin.Bin)
	reason, err := utils.GetRevertReason(context.Background(), cl.GetClient(), opts, contractAddress, abi, binContract)
	if err != nil {
		return errors.Wrap(err, "failed to get the error reson")
	}
	if reason == "execution reverted" {
		reason, err = utils.GetTxRevertReason(context.Background(), cl.GetClient(), tx.Hash())
		if err != nil {
			return errors.Wrap(err, "failed to get the revert reason")
		}
		return errors.New("failed reason : " + reason)
	}
	return errors.New("failed reason : " + reason)
}
func (cl *EthClient) WaitForTx(tx *types.Transaction) {
	logrus.Info("started to wait for the publishing to happen")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			logrus.Error("timeout reached while waiting for the transaction")
			return
		case <-ticker.C:
			receipt, err := cl.client.TransactionReceipt(ctx, tx.Hash())
			if err != nil {
				continue
			}
			if receipt != nil && receipt.Status == 1 {
				logrus.Info("transaction confirmed successfully")
				return
			}
		}
	}
}

func (cl *EthClient) addUser(adminPk *ecdsa.PrivateKey, adminFrom common.Address) (*ecdsa.PrivateKey, common.Address, error) {
	user1Pk, user1From, err := user.AddNewUser(cl.GetClient(), adminPk, adminFrom)
	if err != nil {
		return nil, common.Address{}, errors.Wrap(err, "failed to add new user")
	}
	balance1, err := cl.client.BalanceAt(context.Background(), user1From, nil)
	if err != nil {
		return nil, common.Address{}, errors.Wrap(err, "failed to get balance of user")
	}
	logrus.WithField("user1", user1Pk).WithField("address", user1From).
		WithField("balance", balance1.Int64()).Info("user details")
	return user1Pk, user1From, nil
}
