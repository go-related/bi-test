package smartcontracts

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"math/big"
)

type EthClient struct {
	client *ethclient.Client
}

func NewEthClient(url string) (*EthClient, error) {
	client, err := ethclient.Dial(url)
	return &EthClient{client}, err
}

func (cl *EthClient) GetClient() *ethclient.Client {
	return cl.client
}

func (cl *EthClient) CreateTransaction(userPrivateKey *ecdsa.PrivateKey, userAddress common.Address, addCostSettings bool) (*bind.TransactOpts, error) {
	chainID, err := cl.client.ChainID(context.Background())
	if err != nil {
		return nil, errors.Wrap(err, "failed to get chain ID")
	}

	transactionOptions, err := bind.NewKeyedTransactorWithChainID(userPrivateKey, chainID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create transactor")
	}
	if addCostSettings {
		tipCap, _ := cl.client.SuggestGasTipCap(context.Background())
		feeCap, _ := cl.client.SuggestGasPrice(context.Background())
		nonce, _ := cl.client.PendingNonceAt(context.Background(), userAddress)

		transactionOptions.Nonce = big.NewInt(int64(nonce))
		transactionOptions.GasFeeCap = feeCap
		transactionOptions.GasTipCap = tipCap
		transactionOptions.GasLimit = uint64(21840)
	}

	return transactionOptions, nil
}
