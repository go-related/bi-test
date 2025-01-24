package migration

import (
	"entdemo/blockchain/smartcontracts"
	"entdemo/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func DeployContract(host string, adminPkHex string) error {
	adminPk, adminAddress, err := smartcontracts.GetMetadataFromPrivateKeyHex(adminPkHex)
	if err != nil {
		return errors.Wrap(err, "failed to decode admin pk")
	}

	client, err := smartcontracts.NewEthClient(host)
	if err != nil {
		return errors.Wrap(err, "failed to connect to Ethereum client")
	}
	tx, err := client.CreateTransaction(adminPk, adminAddress, false)
	if err != nil {
		return errors.Wrap(err, "failed to create transaction")
	}

	// Deploy the contract passing the newly created `auth` and `conn` vars
	address, err := deployCarRenting(tx, client.GetClient())
	if err != nil {
		return errors.Wrap(err, "failed to deploy car renting contract")
	}
	logrus.Infof("Deploy contract address: %s", address)
	return nil
}

func deployCarRenting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, error) {
	parsed, err := contracts.CarRentingMetaDataWithBin.GetAbi()
	if err != nil {
		return common.Address{}, err
	}
	if parsed == nil {
		return common.Address{}, errors.New("GetABI returned nil")
	}

	address, tx, _, err := bind.DeployContract(auth, *parsed, common.FromHex(contracts.CarRentingMetaDataWithBin.Bin), backend)
	if err != nil {
		return common.Address{}, err
	}
	logrus.WithField("Contract Address", address.Hex()).WithField("tx", tx).Info("Deployed contract successfully")

	return address, nil
}
