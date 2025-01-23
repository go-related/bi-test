package utils

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

func GetMetadataFromPrivateKeyHex(privateKeyHex string) (*ecdsa.PrivateKey, common.Address) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		logrus.Println("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return privateKey, crypto.PubkeyToAddress(*publicKeyECDSA)
}

func AddNewUser(ethClient *ethclient.Client, adminPk *ecdsa.PrivateKey, adminFrom common.Address) (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to generate private key: %w", err)
	}
	newUserAddress := crypto.PubkeyToAddress(privateKey.PublicKey)

	// Get sender's nonce
	nonce, err := ethClient.PendingNonceAt(context.Background(), adminFrom)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to get nonce: %w", err)
	}

	// Suggest gas fees for EIP-1559 transaction
	gasTipCap, err := ethClient.SuggestGasTipCap(context.Background())
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to get gas tip cap: %w", err)
	}

	gasFeeCap, err := ethClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to get gas fee cap: %w", err)
	}

	// Set the amount to send (0.01 ETH)
	amount := new(big.Int).Mul(big.NewInt(10000), big.NewInt(1e18))
	gasLimit := uint64(21000) // Standard gas for ETH transfers

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
		return nil, common.Address{}, fmt.Errorf("failed to get chain ID: %w", err)
	}

	// Sign transaction with admin's private key
	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), adminPk)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to sign transaction: %w", err)
	}

	// Broadcast the transaction
	err = ethClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, common.Address{}, fmt.Errorf("failed to send transaction: %w", err)
	}

	log.Printf("Transaction sent! Hash: %s\n", signedTx.Hash().Hex())

	return privateKey, newUserAddress, nil
}

func GetTransaction(ethclient *ethclient.Client, privateKey *ecdsa.PrivateKey, chainId *big.Int, fromAddress common.Address) (*bind.TransactOpts, error) {
	transactionOptions, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}

	auth, err := getTransactOpts(ethclient, fromAddress, transactionOptions)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func getTransactOpts(ethclient *ethclient.Client, fromAddress common.Address, auth *bind.TransactOpts) (*bind.TransactOpts, error) {
	nonce, err := ethclient.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	auth.GasPrice, err = ethclient.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth.GasLimit = uint64(21840)

	//baseFee, err := ethclient.SuggestGasPrice(context.Background())
	//log.Println("BaseFee:", baseFee)
	//if err != nil {
	//	return nil, err
	//}
	//
	//priorityFee, err := ethclient.SuggestGasTipCap(context.Background())
	//log.Println("PriorityFee:", priorityFee)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// Adding 5 gwei to baseFee as margin and priority fee
	//increment := new(big.Int).Mul(big.NewInt(5), big.NewInt(params.GWei))
	//log.Println("Increment:", increment)
	//gasFeeCap := new(big.Int).Add(baseFee, increment)
	//gasFeeCap.Add(gasFeeCap, priorityFee)
	//
	//auth.GasLimit = 0
	//auth.GasFeeCap = gasFeeCap
	//auth.GasTipCap = priorityFee

	return auth, nil
}

func GetRevertReason(ctx context.Context, client *ethclient.Client, signer *bind.TransactOpts, contractAddress common.Address, contractABI *abi.ABI, contractMethod []byte) (string, error) {
	msg := ethereum.CallMsg{
		From:      signer.From,
		To:        &contractAddress,
		Gas:       750000,
		GasFeeCap: signer.GasFeeCap,
		GasTipCap: signer.GasTipCap,
		Value:     signer.Value,
		Data:      contractMethod,
	}

	res, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		log.Printf("call contract error: %s", err)
		return err.Error(), nil
	}

	return string(res), nil
}

func GetTxRevertReason(ctx context.Context, client *ethclient.Client, hash common.Hash) (string, error) {
	tx, _, err := client.TransactionByHash(ctx, hash)
	if err != nil {
		return "", err
	}

	from, err := types.Sender(types.LatestSignerForChainID(tx.ChainId()), tx)
	if err != nil {
		return "", err
	}

	var msg ethereum.CallMsg
	switch tx.Type() {
	case 0:
		msg = ethereum.CallMsg{
			From:     from,
			To:       tx.To(),
			Gas:      tx.Gas(),
			GasPrice: tx.GasPrice(),
			Value:    tx.Value(),
			Data:     tx.Data(),
		}
	case 2:
		msg = ethereum.CallMsg{
			From:      from,
			To:        tx.To(),
			Gas:       tx.Gas(),
			GasFeeCap: tx.GasFeeCap(),
			GasTipCap: tx.GasTipCap(),
			Value:     tx.Value(),
			Data:      tx.Data(),
		}
	}
	res, err := client.CallContract(ctx, msg, nil)
	if err != nil {
		log.Printf("call contract error: %s", err)
		return "", err
	}

	return string(res), nil
}
