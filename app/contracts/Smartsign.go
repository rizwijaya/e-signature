// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package api

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"addPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_signature_data\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_signature_nodata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_latin_data\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_latin_nodata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_date_updated\",\"type\":\"string\"}],\"name\":\"addmysignatures\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash_original\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signers\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"addsigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"documentlists\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mySign\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signers\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"signature_data\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"signature_nodata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"latin_data\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"latin_nodata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"date_updated\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mySignature\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash_original\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_namefile\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash_file\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash_ipfs\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_totalsigned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_createdtime\",\"type\":\"uint256\"}],\"name\":\"signDoc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sistem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userAccess\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"verify_document\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060008054336001600160a01b0319918216811783556001805490921617815560028190556003819055600481905560055561101890819061005290396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80638a4cd19d116100665780638a4cd19d146101605780638bb5396014610173578063a355176414610186578063afec53f314610199578063ed1c62b6146101ac57600080fd5b806321b84ef1146100a35780634d0a5dbd146100c9578063514b670b146100fd57806377abc16514610110578063839aef5d14610135575b600080fd5b6100b66100b13660046109d4565b6101cf565b6040519081526020015b60405180910390f35b6100fb6100d7366004610a09565b6001600160a01b03166000908152600660205260409020805460ff19166001179055565b005b6100fb61010b366004610ace565b6101f0565b61012361011e366004610a09565b6102c2565b6040516100c096959493929190610b72565b600154610148906001600160a01b031681565b6040516001600160a01b0390911681526020016100c0565b6100fb61016e366004610bf1565b6105a9565b6101486101813660046109d4565b610751565b6100fb610194366004610cc3565b61077b565b600054610148906001600160a01b031681565b6101bf6101ba3660046109d4565b6108a5565b60405190151581526020016100c0565b600981815481106101df57600080fd5b600091825260209091200154905081565b6101f983610910565b61021e5760405162461bcd60e51b815260040161021590610d95565b60405180910390fd5b604080516080810182526001600160a01b03808516808352600160208085018281528587018881526060870184905260008b8152600a8452888120958152600d909501909252959092208451815496511515600160a01b026001600160a81b03199097169416939093179490941782555191929091908201906102a19082610e67565b50606091909101516002909101805460ff1916911515919091179055505050565b600860205260009081526040902080546001820180546001600160a01b0390921692916102ee90610dde565b80601f016020809104026020016040519081016040528092919081815260200182805461031a90610dde565b80156103675780601f1061033c57610100808354040283529160200191610367565b820191906000526020600020905b81548152906001019060200180831161034a57829003601f168201915b50505050509080600201805461037c90610dde565b80601f01602080910402602001604051908101604052809291908181526020018280546103a890610dde565b80156103f55780601f106103ca576101008083540402835291602001916103f5565b820191906000526020600020905b8154815290600101906020018083116103d857829003601f168201915b50505050509080600301805461040a90610dde565b80601f016020809104026020016040519081016040528092919081815260200182805461043690610dde565b80156104835780601f1061045857610100808354040283529160200191610483565b820191906000526020600020905b81548152906001019060200180831161046657829003601f168201915b50505050509080600401805461049890610dde565b80601f01602080910402602001604051908101604052809291908181526020018280546104c490610dde565b80156105115780601f106104e657610100808354040283529160200191610511565b820191906000526020600020905b8154815290600101906020018083116104f457829003601f168201915b50505050509080600501805461052690610dde565b80601f016020809104026020016040519081016040528092919081815260200182805461055290610dde565b801561059f5780601f106105745761010080835404028352916020019161059f565b820191906000526020600020905b81548152906001019060200180831161058257829003601f168201915b5050505050905086565b6105b288610910565b6105ce5760405162461bcd60e51b815260040161021590610d95565b6000888152600a60205260409020600301546105ec90600190610f3d565b60000361060e576000888152600a602052604090206002600890910155610625565b6000888152600a6020526040902060016008909101555b6004805460008a8152600a602052604090209081556002810180546001600160a01b03191633179055016106598882610e67565b506000888152600a602052604090206001810189905560050161067c8782610e67565b506000888152600a602052604090206007016106988682610e67565b506000888152600a602052604090206006016106b48582610e67565b506106c0600183610f3d565b6000898152600a6020819052604082206003810193909355600980840187905590830184905542600b840155600c909201805460ff19166001908117909155825480820184559282527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af9092018a905560048054909190610742908490610f54565b90915550505050505050505050565b6007818154811061076157600080fd5b6000918252602090912001546001600160a01b0316905081565b6040805160c081018252338082526020808301898152838501899052606084018890526080840187905260a0840186905260009283526008909152929020815181546001600160a01b0319166001600160a01b0390911617815591519091829160018201906107ea9082610e67565b50604082015160028201906107ff9082610e67565b50606082015160038201906108149082610e67565b50608082015160048201906108299082610e67565b5060a0820151600582019061083e9082610e67565b5050600780546001818101835560009283527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c68890910180546001600160a01b03191633179055600580549193509190610898908490610f54565b9091555050505050505050565b6040805160208082018352600091829052838252600a9052818120915190917fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470916108f39160050190610f6c565b60405180910390201461090857506001919050565b506000919050565b6000818152600a60205260408120600c015460ff161515810361093557506000919050565b6000828152600a6020526040902060030154158061096457506000828152600a60205260409020600801546002145b1561097157506000919050565b6000828152600a60209081526040808320338452600d0190915290206002015460ff161515600103610908576000828152600a60209081526040808320338452600d01909152812054600160a01b900460ff161515900361090857506001919050565b6000602082840312156109e657600080fd5b5035919050565b80356001600160a01b0381168114610a0457600080fd5b919050565b600060208284031215610a1b57600080fd5b610a24826109ed565b9392505050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610a5257600080fd5b813567ffffffffffffffff80821115610a6d57610a6d610a2b565b604051601f8301601f19908116603f01168101908282118183101715610a9557610a95610a2b565b81604052838152866020858801011115610aae57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600060608486031215610ae357600080fd5b83359250610af3602085016109ed565b9150604084013567ffffffffffffffff811115610b0f57600080fd5b610b1b86828701610a41565b9150509250925092565b6000815180845260005b81811015610b4b57602081850181015186830182015201610b2f565b81811115610b5d576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b038716815260c060208201819052600090610b9690830188610b25565b8281036040840152610ba88188610b25565b90508281036060840152610bbc8187610b25565b90508281036080840152610bd08186610b25565b905082810360a0840152610be48185610b25565b9998505050505050505050565b600080600080600080600080610100898b031215610c0e57600080fd5b88359750602089013567ffffffffffffffff80821115610c2d57600080fd5b610c398c838d01610a41565b985060408b0135915080821115610c4f57600080fd5b610c5b8c838d01610a41565b975060608b0135915080821115610c7157600080fd5b610c7d8c838d01610a41565b965060808b0135915080821115610c9357600080fd5b50610ca08b828c01610a41565b989b979a50959894979660a0860135965060c08601359560e00135945092505050565b600080600080600060a08688031215610cdb57600080fd5b853567ffffffffffffffff80821115610cf357600080fd5b610cff89838a01610a41565b96506020880135915080821115610d1557600080fd5b610d2189838a01610a41565b95506040880135915080821115610d3757600080fd5b610d4389838a01610a41565b94506060880135915080821115610d5957600080fd5b610d6589838a01610a41565b93506080880135915080821115610d7b57600080fd5b50610d8888828901610a41565b9150509295509295909350565b60208082526029908201527f596f7520617265206e6f7420616c6c6f77656420746f207369676e207468697360408201526808191bd8dd5b595b9d60ba1b606082015260800190565b600181811c90821680610df257607f821691505b602082108103610e1257634e487b7160e01b600052602260045260246000fd5b50919050565b601f821115610e6257600081815260208120601f850160051c81016020861015610e3f5750805b601f850160051c820191505b81811015610e5e57828155600101610e4b565b5050505b505050565b815167ffffffffffffffff811115610e8157610e81610a2b565b610e9581610e8f8454610dde565b84610e18565b602080601f831160018114610eca5760008415610eb25750858301515b600019600386901b1c1916600185901b178555610e5e565b600085815260208120601f198616915b82811015610ef957888601518255948401946001909101908401610eda565b5085821015610f175787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b600082821015610f4f57610f4f610f27565b500390565b60008219821115610f6757610f67610f27565b500190565b6000808354610f7a81610dde565b60018281168015610f925760018114610fa757610fd6565b60ff1984168752821515830287019450610fd6565b8760005260208060002060005b85811015610fcd5781548a820152908401908201610fb4565b50505082870194505b5092969550505050505056fea2646970667358221220b26fa0b882c37ba1fa412e52c63a3969fb5112a08171833e00554a486201393d64736f6c634300080f0033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ApiABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// Documentlists is a free data retrieval call binding the contract method 0x21b84ef1.
//
// Solidity: function documentlists(uint256 ) view returns(bytes32)
func (_Api *ApiCaller) Documentlists(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "documentlists", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Documentlists is a free data retrieval call binding the contract method 0x21b84ef1.
//
// Solidity: function documentlists(uint256 ) view returns(bytes32)
func (_Api *ApiSession) Documentlists(arg0 *big.Int) ([32]byte, error) {
	return _Api.Contract.Documentlists(&_Api.CallOpts, arg0)
}

// Documentlists is a free data retrieval call binding the contract method 0x21b84ef1.
//
// Solidity: function documentlists(uint256 ) view returns(bytes32)
func (_Api *ApiCallerSession) Documentlists(arg0 *big.Int) ([32]byte, error) {
	return _Api.Contract.Documentlists(&_Api.CallOpts, arg0)
}

// MySign is a free data retrieval call binding the contract method 0x77abc165.
//
// Solidity: function mySign(address ) view returns(address signers, string signature_data, string signature_nodata, string latin_data, string latin_nodata, string date_updated)
func (_Api *ApiCaller) MySign(opts *bind.CallOpts, arg0 common.Address) (struct {
	Signers         common.Address
	SignatureData   string
	SignatureNodata string
	LatinData       string
	LatinNodata     string
	DateUpdated     string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "mySign", arg0)

	outstruct := new(struct {
		Signers         common.Address
		SignatureData   string
		SignatureNodata string
		LatinData       string
		LatinNodata     string
		DateUpdated     string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Signers = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SignatureData = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.SignatureNodata = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.LatinData = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.LatinNodata = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.DateUpdated = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// MySign is a free data retrieval call binding the contract method 0x77abc165.
//
// Solidity: function mySign(address ) view returns(address signers, string signature_data, string signature_nodata, string latin_data, string latin_nodata, string date_updated)
func (_Api *ApiSession) MySign(arg0 common.Address) (struct {
	Signers         common.Address
	SignatureData   string
	SignatureNodata string
	LatinData       string
	LatinNodata     string
	DateUpdated     string
}, error) {
	return _Api.Contract.MySign(&_Api.CallOpts, arg0)
}

// MySign is a free data retrieval call binding the contract method 0x77abc165.
//
// Solidity: function mySign(address ) view returns(address signers, string signature_data, string signature_nodata, string latin_data, string latin_nodata, string date_updated)
func (_Api *ApiCallerSession) MySign(arg0 common.Address) (struct {
	Signers         common.Address
	SignatureData   string
	SignatureNodata string
	LatinData       string
	LatinNodata     string
	DateUpdated     string
}, error) {
	return _Api.Contract.MySign(&_Api.CallOpts, arg0)
}

// MySignature is a free data retrieval call binding the contract method 0x8bb53960.
//
// Solidity: function mySignature(uint256 ) view returns(address)
func (_Api *ApiCaller) MySignature(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "mySignature", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MySignature is a free data retrieval call binding the contract method 0x8bb53960.
//
// Solidity: function mySignature(uint256 ) view returns(address)
func (_Api *ApiSession) MySignature(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.MySignature(&_Api.CallOpts, arg0)
}

// MySignature is a free data retrieval call binding the contract method 0x8bb53960.
//
// Solidity: function mySignature(uint256 ) view returns(address)
func (_Api *ApiCallerSession) MySignature(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.MySignature(&_Api.CallOpts, arg0)
}

// Sistem is a free data retrieval call binding the contract method 0xafec53f3.
//
// Solidity: function sistem() view returns(address)
func (_Api *ApiCaller) Sistem(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "sistem")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sistem is a free data retrieval call binding the contract method 0xafec53f3.
//
// Solidity: function sistem() view returns(address)
func (_Api *ApiSession) Sistem() (common.Address, error) {
	return _Api.Contract.Sistem(&_Api.CallOpts)
}

// Sistem is a free data retrieval call binding the contract method 0xafec53f3.
//
// Solidity: function sistem() view returns(address)
func (_Api *ApiCallerSession) Sistem() (common.Address, error) {
	return _Api.Contract.Sistem(&_Api.CallOpts)
}

// UserAccess is a free data retrieval call binding the contract method 0x839aef5d.
//
// Solidity: function userAccess() view returns(address)
func (_Api *ApiCaller) UserAccess(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "userAccess")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserAccess is a free data retrieval call binding the contract method 0x839aef5d.
//
// Solidity: function userAccess() view returns(address)
func (_Api *ApiSession) UserAccess() (common.Address, error) {
	return _Api.Contract.UserAccess(&_Api.CallOpts)
}

// UserAccess is a free data retrieval call binding the contract method 0x839aef5d.
//
// Solidity: function userAccess() view returns(address)
func (_Api *ApiCallerSession) UserAccess() (common.Address, error) {
	return _Api.Contract.UserAccess(&_Api.CallOpts)
}

// VerifyDocument is a free data retrieval call binding the contract method 0xed1c62b6.
//
// Solidity: function verify_document(bytes32 _hash) view returns(bool)
func (_Api *ApiCaller) VerifyDocument(opts *bind.CallOpts, _hash [32]byte) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "verify_document", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDocument is a free data retrieval call binding the contract method 0xed1c62b6.
//
// Solidity: function verify_document(bytes32 _hash) view returns(bool)
func (_Api *ApiSession) VerifyDocument(_hash [32]byte) (bool, error) {
	return _Api.Contract.VerifyDocument(&_Api.CallOpts, _hash)
}

// VerifyDocument is a free data retrieval call binding the contract method 0xed1c62b6.
//
// Solidity: function verify_document(bytes32 _hash) view returns(bool)
func (_Api *ApiCallerSession) VerifyDocument(_hash [32]byte) (bool, error) {
	return _Api.Contract.VerifyDocument(&_Api.CallOpts, _hash)
}

// AddPermission is a paid mutator transaction binding the contract method 0x4d0a5dbd.
//
// Solidity: function addPermission(address _member) returns()
func (_Api *ApiTransactor) AddPermission(opts *bind.TransactOpts, _member common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addPermission", _member)
}

// AddPermission is a paid mutator transaction binding the contract method 0x4d0a5dbd.
//
// Solidity: function addPermission(address _member) returns()
func (_Api *ApiSession) AddPermission(_member common.Address) (*types.Transaction, error) {
	return _Api.Contract.AddPermission(&_Api.TransactOpts, _member)
}

// AddPermission is a paid mutator transaction binding the contract method 0x4d0a5dbd.
//
// Solidity: function addPermission(address _member) returns()
func (_Api *ApiTransactorSession) AddPermission(_member common.Address) (*types.Transaction, error) {
	return _Api.Contract.AddPermission(&_Api.TransactOpts, _member)
}

// Addmysignatures is a paid mutator transaction binding the contract method 0xa3551764.
//
// Solidity: function addmysignatures(string _signature_data, string _signature_nodata, string _latin_data, string _latin_nodata, string _date_updated) returns()
func (_Api *ApiTransactor) Addmysignatures(opts *bind.TransactOpts, _signature_data string, _signature_nodata string, _latin_data string, _latin_nodata string, _date_updated string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addmysignatures", _signature_data, _signature_nodata, _latin_data, _latin_nodata, _date_updated)
}

// Addmysignatures is a paid mutator transaction binding the contract method 0xa3551764.
//
// Solidity: function addmysignatures(string _signature_data, string _signature_nodata, string _latin_data, string _latin_nodata, string _date_updated) returns()
func (_Api *ApiSession) Addmysignatures(_signature_data string, _signature_nodata string, _latin_data string, _latin_nodata string, _date_updated string) (*types.Transaction, error) {
	return _Api.Contract.Addmysignatures(&_Api.TransactOpts, _signature_data, _signature_nodata, _latin_data, _latin_nodata, _date_updated)
}

// Addmysignatures is a paid mutator transaction binding the contract method 0xa3551764.
//
// Solidity: function addmysignatures(string _signature_data, string _signature_nodata, string _latin_data, string _latin_nodata, string _date_updated) returns()
func (_Api *ApiTransactorSession) Addmysignatures(_signature_data string, _signature_nodata string, _latin_data string, _latin_nodata string, _date_updated string) (*types.Transaction, error) {
	return _Api.Contract.Addmysignatures(&_Api.TransactOpts, _signature_data, _signature_nodata, _latin_data, _latin_nodata, _date_updated)
}

// Addsigners is a paid mutator transaction binding the contract method 0x514b670b.
//
// Solidity: function addsigners(bytes32 _hash_original, address _signers, string _date) returns()
func (_Api *ApiTransactor) Addsigners(opts *bind.TransactOpts, _hash_original [32]byte, _signers common.Address, _date string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addsigners", _hash_original, _signers, _date)
}

// Addsigners is a paid mutator transaction binding the contract method 0x514b670b.
//
// Solidity: function addsigners(bytes32 _hash_original, address _signers, string _date) returns()
func (_Api *ApiSession) Addsigners(_hash_original [32]byte, _signers common.Address, _date string) (*types.Transaction, error) {
	return _Api.Contract.Addsigners(&_Api.TransactOpts, _hash_original, _signers, _date)
}

// Addsigners is a paid mutator transaction binding the contract method 0x514b670b.
//
// Solidity: function addsigners(bytes32 _hash_original, address _signers, string _date) returns()
func (_Api *ApiTransactorSession) Addsigners(_hash_original [32]byte, _signers common.Address, _date string) (*types.Transaction, error) {
	return _Api.Contract.Addsigners(&_Api.TransactOpts, _hash_original, _signers, _date)
}

// SignDoc is a paid mutator transaction binding the contract method 0x8a4cd19d.
//
// Solidity: function signDoc(bytes32 _hash_original, string _namefile, string _hash_file, string _metadata, string _hash_ipfs, uint256 _totalsigned, uint256 signing, uint256 _createdtime) returns()
func (_Api *ApiTransactor) SignDoc(opts *bind.TransactOpts, _hash_original [32]byte, _namefile string, _hash_file string, _metadata string, _hash_ipfs string, _totalsigned *big.Int, signing *big.Int, _createdtime *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "signDoc", _hash_original, _namefile, _hash_file, _metadata, _hash_ipfs, _totalsigned, signing, _createdtime)
}

// SignDoc is a paid mutator transaction binding the contract method 0x8a4cd19d.
//
// Solidity: function signDoc(bytes32 _hash_original, string _namefile, string _hash_file, string _metadata, string _hash_ipfs, uint256 _totalsigned, uint256 signing, uint256 _createdtime) returns()
func (_Api *ApiSession) SignDoc(_hash_original [32]byte, _namefile string, _hash_file string, _metadata string, _hash_ipfs string, _totalsigned *big.Int, signing *big.Int, _createdtime *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SignDoc(&_Api.TransactOpts, _hash_original, _namefile, _hash_file, _metadata, _hash_ipfs, _totalsigned, signing, _createdtime)
}

// SignDoc is a paid mutator transaction binding the contract method 0x8a4cd19d.
//
// Solidity: function signDoc(bytes32 _hash_original, string _namefile, string _hash_file, string _metadata, string _hash_ipfs, uint256 _totalsigned, uint256 signing, uint256 _createdtime) returns()
func (_Api *ApiTransactorSession) SignDoc(_hash_original [32]byte, _namefile string, _hash_file string, _metadata string, _hash_ipfs string, _totalsigned *big.Int, signing *big.Int, _createdtime *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SignDoc(&_Api.TransactOpts, _hash_original, _namefile, _hash_file, _metadata, _hash_ipfs, _totalsigned, signing, _createdtime)
}
