package user

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"math/big"
)

func AddNewUser(ethClient *ethclient.Client, adminPk *ecdsa.PrivateKey, adminFrom common.Address) (
	*ecdsa.PrivateKey, common.Address, error) {

	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to generate private key: %w", err)
	}
	newUserAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	err = foundNewUser(ethClient, adminPk, adminFrom, err, newUserAddress)
	if err != nil {
		return nil, common.Address{}, errors.Wrap(err, "failed to found new user")
	}

	// Convert private key to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)
	publicKeyBytes := crypto.FromECDSAPub(&privateKey.PublicKey)

	// Print hex representations
	logrus.WithField("Private Key", hex.EncodeToString(privateKeyBytes)).
		WithField("Public Key", hex.EncodeToString(publicKeyBytes)).
		Info("create new user with data")
	return privateKey, newUserAddress, nil
}

func foundNewUser(ethClient *ethclient.Client, adminPk *ecdsa.PrivateKey,
	adminFrom common.Address, err error, newUserAddress common.Address) error {
	// Get sender's nonce
	nonce, err := ethClient.PendingNonceAt(context.Background(), adminFrom)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %w", err)
	}

	// Suggest gas fees for EIP-1559 transaction
	gasTipCap, err := ethClient.SuggestGasTipCap(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas tip cap: %w", err)
	}

	gasFeeCap, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas fee cap: %w", err)
	}

	// Set the amount to send (0.01 ETH)
	amount := new(big.Int).Mul(big.NewInt(10000), big.NewInt(1e18))
	gasLimit := uint64(21000)

	// Create the EIP-1559 transaction
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		To:        &newUserAddress,
		Value:     amount,
		Gas:       gasLimit,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
	})

	// Get network chain ID
	chainID, err := ethClient.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Sign transaction with admin's private key
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), adminPk)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Broadcast the transaction
	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("failed to send transaction: %w", err)
	}
	return nil
}
