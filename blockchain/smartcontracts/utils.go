package smartcontracts

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

func GetMetadataFromPrivateKeyHex(privateKeyHex string) (*ecdsa.PrivateKey, common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, [20]byte{}, errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	return privateKey, crypto.PubkeyToAddress(*publicKeyECDSA), nil
}
