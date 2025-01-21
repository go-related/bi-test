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

var CarRentingMetaDataWithBin = &bind.MetaData{
	ABI: "[\n    {\n      \"type\": \"constructor\",\n      \"inputs\": [],\n      \"stateMutability\": \"nonpayable\"\n    },\n    {\n      \"type\": \"function\",\n      \"name\": \"addCar\",\n      \"inputs\": [\n        {\n          \"name\": \"car\",\n          \"type\": \"tuple\",\n          \"internalType\": \"struct CarRenting.Car\",\n          \"components\": [\n            {\n              \"name\": \"hp\",\n              \"type\": \"uint16\",\n              \"internalType\": \"uint16\"\n            },\n            {\n              \"name\": \"fuelLevel\",\n              \"type\": \"uint16\",\n              \"internalType\": \"uint16\"\n            },\n            {\n              \"name\": \"isRented\",\n              \"type\": \"bool\",\n              \"internalType\": \"bool\"\n            },\n            {\n              \"name\": \"rentedTo\",\n              \"type\": \"address\",\n              \"internalType\": \"address\"\n            }\n          ]\n        }\n      ],\n      \"outputs\": [],\n      \"stateMutability\": \"nonpayable\"\n    },\n    {\n      \"type\": \"function\",\n      \"name\": \"rentCar\",\n      \"inputs\": [\n        {\n          \"name\": \"owner\",\n          \"type\": \"address\",\n          \"internalType\": \"address\"\n        }\n      ],\n      \"outputs\": [],\n      \"stateMutability\": \"nonpayable\"\n    },\n    {\n      \"type\": \"function\",\n      \"name\": \"rents\",\n      \"inputs\": [\n        {\n          \"name\": \"\",\n          \"type\": \"address\",\n          \"internalType\": \"address\"\n        }\n      ],\n      \"outputs\": [\n        {\n          \"name\": \"hp\",\n          \"type\": \"uint16\",\n          \"internalType\": \"uint16\"\n        },\n        {\n          \"name\": \"fuelLevel\",\n          \"type\": \"uint16\",\n          \"internalType\": \"uint16\"\n        },\n        {\n          \"name\": \"isRented\",\n          \"type\": \"bool\",\n          \"internalType\": \"bool\"\n        },\n        {\n          \"name\": \"rentedTo\",\n          \"type\": \"address\",\n          \"internalType\": \"address\"\n        }\n      ],\n      \"stateMutability\": \"view\"\n    }\n  ]",
	Bin: "0x6080604052348015600e575f5ffd5b5061032f8061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061003f575f3560e01c80634f1d221f146100435780639fdc8f051461009a578063c70676451461011e575b5f5ffd5b61009861005136600461024a565b335f90815260208181526040909120825181549390920151640100000000600160c81b031961ffff918216620100000263ffffffff19909516919093161792909217169055565b005b6100e86100a83660046102d9565b5f6020819052908152604090205461ffff8082169162010000810490911690640100000000810460ff16906501000000000090046001600160a01b031684565b6040805161ffff9586168152939094166020840152901515828401526001600160a01b0316606082015290519081900360800190f35b61009861012c3660046102d9565b6001600160a01b0381165f9081526020819052604090208054640100000000900460ff16156101975760405162461bcd60e51b815260206004820152601260248201527110d85c88185b1c9958591e481c995b9d195960721b60448201526064015b60405180910390fd5b805461ffff16158015906101b55750805462010000900461ffff1615155b6101f65760405162461bcd60e51b815260206004820152601260248201527110d85c88191bd95cc81b9bdd08195e1a5cdd60721b604482015260640161018e565b8054650100000000003302640100000000600160c81b03199091161764010000000017905550565b803561ffff8116811461022f575f5ffd5b919050565b80356001600160a01b038116811461022f575f5ffd5b5f608082840312801561025b575f5ffd5b506040516080810167ffffffffffffffff8111828210171561028b57634e487b7160e01b5f52604160045260245ffd5b6040526102978361021e565b81526102a56020840161021e565b6020820152604083013580151581146102bc575f5ffd5b60408201526102cd60608401610234565b60608201529392505050565b5f602082840312156102e9575f5ffd5b6102f282610234565b939250505056fea26469706673582212203202b0649349b07120088f9e282bc0fa86d2255293d02c6c331b4b3d5d55f9da64736f6c634300081c0033",
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

func DeployCarRentingHelper(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *bind.BoundContract, error) {
	parsed, err := CarRentingMetaDataWithBin.GetAbi()
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
