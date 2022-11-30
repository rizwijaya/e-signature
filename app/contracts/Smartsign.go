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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_creator_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_signers_id\",\"type\":\"string[]\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"}],\"name\":\"getDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"getDocSigned\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_signers_id\",\"type\":\"address\"}],\"name\":\"getSign\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_signers_id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_signers_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"signDoc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060016000556114ab806100256000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80631af2925c146100675780632887590e1461007c5780633d995877146100a5578063a5bde23b146100c9578063c7335804146100f4578063e1ccea4014610117575b600080fd5b61007a610075366004610db5565b61012a565b005b61008f61008a366004610e56565b6102f4565b60405161009c9190610ee0565b60405180910390f35b6100b86100b3366004610efa565b610325565b60405161009c959493929190610f48565b6100dc6100d7366004610e56565b61053f565b60405161009c9c9b9a99989796959493929190610f88565b610107610102366004610e56565b6108d4565b604051901515815260200161009c565b61007a610125366004611157565b61090e565b600061013586610c3a565b6000818152600160208190526040909120600c8101549293509160ff1615151461017a5760405162461bcd60e51b8152600401610171906112aa565b60405180910390fd5b6001600160a01b0386166000908152600e820160205260409020600501546001106101b75760405162461bcd60e51b8152600401610171906112aa565b6001600160a01b0386166000908152600e8201602052604090206004015460ff16156102255760405162461bcd60e51b815260206004820152601c60248201527f596f7520617265207369676e6564207468697320646f63756d656e74000000006044820152606401610171565b600d81018054906000610237836112ec565b90915550506007810161024a8582611386565b50600b8101839055600681016102608682611386565b506001600160a01b0386166000908152600e8201602052604090206003016102888682611386565b506001600160a01b0386166000908152600e82016020526040812060048101805460ff191660011790556005018490556102c186610c3a565b8254600082815260026020526040902055600d830154909150600111156102ea57600260088301555b5050505050505050565b6060600061030183610c3a565b60008181526002602052604081205491925061031c82610c59565b95945050505050565b6000606080600080600061033888610c3a565b6000818152600160208190526040909120600c8101549293509160ff161515146103745760405162461bcd60e51b8152600401610171906112aa565b6001600160a01b038089166000818152600e84016020526040902054909116146103d45760405162461bcd60e51b815260206004820152601160248201527014da59db995c9cc81b9bdd08195e1a5cdd607a1b6044820152606401610171565b6001600160a01b0388166000908152600e8201602052604090206001810154600482015460058301546002840180549394909360039091019260ff169190849061041d90611303565b80601f016020809104026020016040519081016040528092919081815260200182805461044990611303565b80156104965780601f1061046b57610100808354040283529160200191610496565b820191906000526020600020905b81548152906001019060200180831161047957829003601f168201915b505050505093508280546104a990611303565b80601f01602080910402602001604051908101604052809291908181526020018280546104d590611303565b80156105225780601f106104f757610100808354040283529160200191610522565b820191906000526020600020905b81548152906001019060200180831161050557829003601f168201915b505050505092509650965096509650965050509295509295909350565b600080606080606080606060008060008060008061055c8e610c3a565b6000818152600160208190526040909120600c8101549293509160ff161515146105985760405162461bcd60e51b8152600401610171906112aa565b80600101548160030160009054906101000a90046001600160a01b031682600201836004018460050185600601866007018760080154886009015489600a01548a600b01548b600c0160009054906101000a900460ff168980546105fb90611303565b80601f016020809104026020016040519081016040528092919081815260200182805461062790611303565b80156106745780601f1061064957610100808354040283529160200191610674565b820191906000526020600020905b81548152906001019060200180831161065757829003601f168201915b5050505050995088805461068790611303565b80601f01602080910402602001604051908101604052809291908181526020018280546106b390611303565b80156107005780601f106106d557610100808354040283529160200191610700565b820191906000526020600020905b8154815290600101906020018083116106e357829003601f168201915b5050505050985087805461071390611303565b80601f016020809104026020016040519081016040528092919081815260200182805461073f90611303565b801561078c5780601f106107615761010080835404028352916020019161078c565b820191906000526020600020905b81548152906001019060200180831161076f57829003601f168201915b5050505050975086805461079f90611303565b80601f01602080910402602001604051908101604052809291908181526020018280546107cb90611303565b80156108185780601f106107ed57610100808354040283529160200191610818565b820191906000526020600020905b8154815290600101906020018083116107fb57829003601f168201915b5050505050965085805461082b90611303565b80601f016020809104026020016040519081016040528092919081815260200182805461085790611303565b80156108a45780601f10610879576101008083540402835291602001916108a4565b820191906000526020600020905b81548152906001019060200180831161088757829003601f168201915b505050505095509d509d509d509d509d509d509d509d509d509d509d509d50505091939597999b5091939597999b565b6000806108e083610c3a565b6000818152600260205260409020549091506108ff5750600092915050565b50600192915050565b50919050565b60006109198c610c3a565b600081815260016020819052604082209154908201558181556003810180546001600160a01b0319166001600160a01b038f16179055909150600281016109608c82611386565b506004810161096f8b82611386565b506005810161097e8e82611386565b506006810161098d8a82611386565b506007810161099c8982611386565b506008810187905560098101869055600a8101859055600b8101859055600c8101805460ff191660011790558351600d82015560005b8451811015610c16578481815181106109ed576109ed611446565b602002602001015182600e016000878481518110610a0d57610a0d611446565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508082600e016000878481518110610a7557610a75611446565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060010181905550838181518110610ab657610ab6611446565b602002602001015182600e016000878481518110610ad657610ad6611446565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206002019081610b0d9190611386565b508982600e016000878481518110610b2757610b27611446565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206003019081610b5e9190611386565b50600082600e016000878481518110610b7957610b79611446565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060040160006101000a81548160ff0219169083151502179055508582600e016000878481518110610bd457610bd4611446565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600501819055508080610c0e9061145c565b9150506109d2565b50600080549080610c268361145c565b919050555050505050505050505050505050565b805160009082908203610c505750600092915050565b50506020015190565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b6020811015610cdb57838160208110610c9c57610c9c611446565b1a60f81b828281518110610cb257610cb2611446565b60200101906001600160f81b031916908160001a90535080610cd38161145c565b915050610c81565b5092915050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610d2157610d21610ce2565b604052919050565b600082601f830112610d3a57600080fd5b813567ffffffffffffffff811115610d5457610d54610ce2565b610d67601f8201601f1916602001610cf8565b818152846020838601011115610d7c57600080fd5b816020850160208301376000918101602001919091529392505050565b80356001600160a01b0381168114610db057600080fd5b919050565b600080600080600060a08688031215610dcd57600080fd5b853567ffffffffffffffff80821115610de557600080fd5b610df189838a01610d29565b9650610dff60208901610d99565b95506040880135915080821115610e1557600080fd5b610e2189838a01610d29565b94506060880135915080821115610e3757600080fd5b50610e4488828901610d29565b95989497509295608001359392505050565b600060208284031215610e6857600080fd5b813567ffffffffffffffff811115610e7f57600080fd5b610e8b84828501610d29565b949350505050565b6000815180845260005b81811015610eb957602081850181015186830182015201610e9d565b81811115610ecb576000602083870101525b50601f01601f19169290920160200192915050565b602081526000610ef36020830184610e93565b9392505050565b60008060408385031215610f0d57600080fd5b823567ffffffffffffffff811115610f2457600080fd5b610f3085828601610d29565b925050610f3f60208401610d99565b90509250929050565b85815260a060208201526000610f6160a0830187610e93565b8281036040840152610f738187610e93565b94151560608401525050608001529392505050565b8c81526001600160a01b038c16602082015261018060408201819052600090610fb38382018e610e93565b90508281036060840152610fc7818d610e93565b90508281036080840152610fdb818c610e93565b905082810360a0840152610fef818b610e93565b905082810360c0840152611003818a610e93565b9150508660e083015285610100830152846101208301528361014083015261103061016083018415159052565b9d9c50505050505050505050505050565b600067ffffffffffffffff82111561105b5761105b610ce2565b5060051b60200190565b600082601f83011261107657600080fd5b8135602061108b61108683611041565b610cf8565b82815260059290921b840181019181810190868411156110aa57600080fd5b8286015b848110156110cc576110bf81610d99565b83529183019183016110ae565b509695505050505050565b600082601f8301126110e857600080fd5b813560206110f861108683611041565b82815260059290921b8401810191818101908684111561111757600080fd5b8286015b848110156110cc57803567ffffffffffffffff81111561113b5760008081fd5b6111498986838b0101610d29565b84525091830191830161111b565b60008060008060008060008060008060006101608c8e03121561117957600080fd5b67ffffffffffffffff808d35111561119057600080fd5b61119d8e8e358f01610d29565b9b506111ab60208e01610d99565b9a508060408e013511156111be57600080fd5b6111ce8e60408f01358f01610d29565b99508060608e013511156111e157600080fd5b6111f18e60608f01358f01610d29565b98508060808e0135111561120457600080fd5b6112148e60808f01358f01610d29565b97508060a08e0135111561122757600080fd5b6112378e60a08f01358f01610d29565b965060c08d0135955060e08d013594506101008d01359350806101208e0135111561126157600080fd5b6112728e6101208f01358f01611065565b9250806101408e0135111561128657600080fd5b506112988d6101408e01358e016110d7565b90509295989b509295989b9093969950565b602080825260129082015271111bd8dd5b595b9d081b9bdd08195e1a5cdd60721b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6000816112fb576112fb6112d6565b506000190190565b600181811c9082168061131757607f821691505b60208210810361090857634e487b7160e01b600052602260045260246000fd5b601f82111561138157600081815260208120601f850160051c8101602086101561135e5750805b601f850160051c820191505b8181101561137d5782815560010161136a565b5050505b505050565b815167ffffffffffffffff8111156113a0576113a0610ce2565b6113b4816113ae8454611303565b84611337565b602080601f8311600181146113e957600084156113d15750858301515b600019600386901b1c1916600185901b17855561137d565b600085815260208120601f198616915b82811015611418578886015182559484019460019091019084016113f9565b50858210156114365787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b60006001820161146e5761146e6112d6565b506001019056fea2646970667358221220d8c2c288165af0038cefddb4c140224850ca5c6805f242c9613aaea45f9c7ee164736f6c634300080f0033",
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

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string _file) view returns(uint256, address, string, string, string, string, string, uint256, uint256, uint256, uint256, bool)
func (_Api *ApiCaller) GetDoc(opts *bind.CallOpts, _file string) (*big.Int, common.Address, string, string, string, string, string, *big.Int, *big.Int, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getDoc", _file)

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new(string), *new(string), *new(string), *new(string), *new(string), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)
	out5 := *abi.ConvertType(out[5], new(string)).(*string)
	out6 := *abi.ConvertType(out[6], new(string)).(*string)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	out9 := *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	out10 := *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	out11 := *abi.ConvertType(out[11], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, out10, out11, err

}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string _file) view returns(uint256, address, string, string, string, string, string, uint256, uint256, uint256, uint256, bool)
func (_Api *ApiSession) GetDoc(_file string) (*big.Int, common.Address, string, string, string, string, string, *big.Int, *big.Int, *big.Int, *big.Int, bool, error) {
	return _Api.Contract.GetDoc(&_Api.CallOpts, _file)
}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string _file) view returns(uint256, address, string, string, string, string, string, uint256, uint256, uint256, uint256, bool)
func (_Api *ApiCallerSession) GetDoc(_file string) (*big.Int, common.Address, string, string, string, string, string, *big.Int, *big.Int, *big.Int, *big.Int, bool, error) {
	return _Api.Contract.GetDoc(&_Api.CallOpts, _file)
}

// GetDocSigned is a free data retrieval call binding the contract method 0x2887590e.
//
// Solidity: function getDocSigned(string _hash) view returns(string)
func (_Api *ApiCaller) GetDocSigned(opts *bind.CallOpts, _hash string) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getDocSigned", _hash)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDocSigned is a free data retrieval call binding the contract method 0x2887590e.
//
// Solidity: function getDocSigned(string _hash) view returns(string)
func (_Api *ApiSession) GetDocSigned(_hash string) (string, error) {
	return _Api.Contract.GetDocSigned(&_Api.CallOpts, _hash)
}

// GetDocSigned is a free data retrieval call binding the contract method 0x2887590e.
//
// Solidity: function getDocSigned(string _hash) view returns(string)
func (_Api *ApiCallerSession) GetDocSigned(_hash string) (string, error) {
	return _Api.Contract.GetDocSigned(&_Api.CallOpts, _hash)
}

// GetSign is a free data retrieval call binding the contract method 0x3d995877.
//
// Solidity: function getSign(string _file, address _signers_id) view returns(uint256, string, string, bool, uint256)
func (_Api *ApiCaller) GetSign(opts *bind.CallOpts, _file string, _signers_id common.Address) (*big.Int, string, string, bool, *big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getSign", _file, _signers_id)

	if err != nil {
		return *new(*big.Int), *new(string), *new(string), *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetSign is a free data retrieval call binding the contract method 0x3d995877.
//
// Solidity: function getSign(string _file, address _signers_id) view returns(uint256, string, string, bool, uint256)
func (_Api *ApiSession) GetSign(_file string, _signers_id common.Address) (*big.Int, string, string, bool, *big.Int, error) {
	return _Api.Contract.GetSign(&_Api.CallOpts, _file, _signers_id)
}

// GetSign is a free data retrieval call binding the contract method 0x3d995877.
//
// Solidity: function getSign(string _file, address _signers_id) view returns(uint256, string, string, bool, uint256)
func (_Api *ApiCallerSession) GetSign(_file string, _signers_id common.Address) (*big.Int, string, string, bool, *big.Int, error) {
	return _Api.Contract.GetSign(&_Api.CallOpts, _file, _signers_id)
}

// VerifyDoc is a free data retrieval call binding the contract method 0xc7335804.
//
// Solidity: function verifyDoc(string _hash) view returns(bool)
func (_Api *ApiCaller) VerifyDoc(opts *bind.CallOpts, _hash string) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "verifyDoc", _hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyDoc is a free data retrieval call binding the contract method 0xc7335804.
//
// Solidity: function verifyDoc(string _hash) view returns(bool)
func (_Api *ApiSession) VerifyDoc(_hash string) (bool, error) {
	return _Api.Contract.VerifyDoc(&_Api.CallOpts, _hash)
}

// VerifyDoc is a free data retrieval call binding the contract method 0xc7335804.
//
// Solidity: function verifyDoc(string _hash) view returns(bool)
func (_Api *ApiCallerSession) VerifyDoc(_hash string) (bool, error) {
	return _Api.Contract.VerifyDoc(&_Api.CallOpts, _hash)
}

// Create is a paid mutator transaction binding the contract method 0xe1ccea40.
//
// Solidity: function create(string _file, address creator, string _creator_id, string _metadata, string _hash, string _ipfs, uint256 _state, uint256 _mode, uint256 _time, address[] _signers, string[] _signers_id) returns()
func (_Api *ApiTransactor) Create(opts *bind.TransactOpts, _file string, creator common.Address, _creator_id string, _metadata string, _hash string, _ipfs string, _state *big.Int, _mode *big.Int, _time *big.Int, _signers []common.Address, _signers_id []string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "create", _file, creator, _creator_id, _metadata, _hash, _ipfs, _state, _mode, _time, _signers, _signers_id)
}

// Create is a paid mutator transaction binding the contract method 0xe1ccea40.
//
// Solidity: function create(string _file, address creator, string _creator_id, string _metadata, string _hash, string _ipfs, uint256 _state, uint256 _mode, uint256 _time, address[] _signers, string[] _signers_id) returns()
func (_Api *ApiSession) Create(_file string, creator common.Address, _creator_id string, _metadata string, _hash string, _ipfs string, _state *big.Int, _mode *big.Int, _time *big.Int, _signers []common.Address, _signers_id []string) (*types.Transaction, error) {
	return _Api.Contract.Create(&_Api.TransactOpts, _file, creator, _creator_id, _metadata, _hash, _ipfs, _state, _mode, _time, _signers, _signers_id)
}

// Create is a paid mutator transaction binding the contract method 0xe1ccea40.
//
// Solidity: function create(string _file, address creator, string _creator_id, string _metadata, string _hash, string _ipfs, uint256 _state, uint256 _mode, uint256 _time, address[] _signers, string[] _signers_id) returns()
func (_Api *ApiTransactorSession) Create(_file string, creator common.Address, _creator_id string, _metadata string, _hash string, _ipfs string, _state *big.Int, _mode *big.Int, _time *big.Int, _signers []common.Address, _signers_id []string) (*types.Transaction, error) {
	return _Api.Contract.Create(&_Api.TransactOpts, _file, creator, _creator_id, _metadata, _hash, _ipfs, _state, _mode, _time, _signers, _signers_id)
}

// SignDoc is a paid mutator transaction binding the contract method 0x1af2925c.
//
// Solidity: function signDoc(string _file, address _signers_id, string _signers_hash, string _ipfs, uint256 _time) returns()
func (_Api *ApiTransactor) SignDoc(opts *bind.TransactOpts, _file string, _signers_id common.Address, _signers_hash string, _ipfs string, _time *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "signDoc", _file, _signers_id, _signers_hash, _ipfs, _time)
}

// SignDoc is a paid mutator transaction binding the contract method 0x1af2925c.
//
// Solidity: function signDoc(string _file, address _signers_id, string _signers_hash, string _ipfs, uint256 _time) returns()
func (_Api *ApiSession) SignDoc(_file string, _signers_id common.Address, _signers_hash string, _ipfs string, _time *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SignDoc(&_Api.TransactOpts, _file, _signers_id, _signers_hash, _ipfs, _time)
}

// SignDoc is a paid mutator transaction binding the contract method 0x1af2925c.
//
// Solidity: function signDoc(string _file, address _signers_id, string _signers_hash, string _ipfs, uint256 _time) returns()
func (_Api *ApiTransactorSession) SignDoc(_file string, _signers_id common.Address, _signers_hash string, _ipfs string, _time *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SignDoc(&_Api.TransactOpts, _file, _signers_id, _signers_hash, _ipfs, _time)
}
