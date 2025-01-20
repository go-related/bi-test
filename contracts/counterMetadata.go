package contracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/pkg/errors"
)

var CounterMetaDataWithBin = &bind.MetaData{
	ABI: "[\n    {\n      \"type\": \"function\",\n      \"name\": \"increment\",\n      \"inputs\": [],\n      \"outputs\": [],\n      \"stateMutability\": \"nonpayable\"\n    },\n    {\n      \"type\": \"function\",\n      \"name\": \"number\",\n      \"inputs\": [],\n      \"outputs\": [\n        {\n          \"name\": \"\",\n          \"type\": \"uint256\",\n          \"internalType\": \"uint256\"\n        }\n      ],\n      \"stateMutability\": \"view\"\n    },\n    {\n      \"type\": \"function\",\n      \"name\": \"setNumber\",\n      \"inputs\": [\n        {\n          \"name\": \"newNumber\",\n          \"type\": \"uint256\",\n          \"internalType\": \"uint256\"\n        }\n      ],\n      \"outputs\": [],\n      \"stateMutability\": \"nonpayable\"\n    }\n  ]\n",
	Bin: "0x6080604052348015600e575f5ffd5b5060ec8061001b5f395ff3fe6080604052348015600e575f5ffd5b5060043610603a575f3560e01c80633fb5c1cb14603e5780638381f58a14604f578063d09de08a146068575b5f5ffd5b604d6049366004607d565b5f55565b005b60565f5481565b60405190815260200160405180910390f35b604d5f805490806076836093565b9190505550565b5f60208284031215608c575f5ffd5b5035919050565b5f6001820160af57634e487b7160e01b5f52601160045260245ffd5b506001019056fea2646970667358221220fa49f4ce8b9cbd3848cd128350606f0eb84071d2b55cc5ed387ef5407001680e64736f6c634300081c0033",
}

func DeployCounterHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	parsed, err := CounterMetaDataWithBin.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterMetaDataWithBin.Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, contract, nil
}
