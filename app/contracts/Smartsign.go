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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_creator_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_state\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_visibility\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_signers_id\",\"type\":\"string[]\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"}],\"name\":\"getDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"getDocSigned\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_signers_id\",\"type\":\"address\"}],\"name\":\"getSign\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_signers_id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_signers_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"signDoc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506001600055611490806100256000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806304205959146100675780631af2925c1461007c5780632887590e1461008f5780633d995877146100b8578063a5bde23b146100dc578063c733580414610107575b600080fd5b61007a610075366004610ec0565b61012a565b005b61007a61008a36600461101a565b61045a565b6100a261009d3660046110bb565b6105f8565b6040516100af9190611145565b60405180910390f35b6100cb6100c636600461115f565b610629565b6040516100af9594939291906111ad565b6100ef6100ea3660046110bb565b610843565b6040516100af9c9b9a999897969594939291906111ed565b61011a6101153660046110bb565b610be5565b60405190151581526020016100af565b60006101358c610c1f565b600081815260016020819052604082209154908201558181556003810180546001600160a01b0319166001600160a01b038f161790559091506002810161017c8c82611331565b506004810161018b8b82611331565b506005810161019a8e82611331565b50600681016101a98a82611331565b50600781016101b88982611331565b506008810187905560098101805487151560ff1991821617909155600a8201869055600b8201869055600c82018054909116600117905560005b84518110156104365784818151811061020d5761020d6113f1565b602002602001015182600d01600087848151811061022d5761022d6113f1565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508082600d016000878481518110610295576102956113f1565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600101819055508381815181106102d6576102d66113f1565b602002602001015182600d0160008784815181106102f6576102f66113f1565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600201908161032d9190611331565b508982600d016000878481518110610347576103476113f1565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600301908161037e9190611331565b50600082600d016000878481518110610399576103996113f1565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060040160006101000a81548160ff0219169083151502179055508582600d0160008784815181106103f4576103f46113f1565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060050181905550808061042e90611407565b9150506101f2565b5060008054908061044683611407565b919050555050505050505050505050505050565b600061046586610c1f565b6000818152600160208190526040909120600c8101549293509160ff161515146104aa5760405162461bcd60e51b81526004016104a19061142e565b60405180910390fd5b6001600160a01b0386166000908152600d820160205260409020600501546001106104e75760405162461bcd60e51b81526004016104a19061142e565b6001600160a01b0386166000908152600d8201602052604090206004015460ff16156105555760405162461bcd60e51b815260206004820152601c60248201527f596f7520617265207369676e6564207468697320646f63756d656e740000000060448201526064016104a1565b600781016105638582611331565b50600b8101839055600681016105798682611331565b506001600160a01b0386166000908152600d8201602052604090206003016105a18682611331565b506001600160a01b0386166000908152600d82016020526040812060048101805460ff191660011790556005018490556105da86610c1f565b91546000928352600260205260409092209190915550505050505050565b6060600061060583610c1f565b60008181526002602052604081205491925061062082610c3e565b95945050505050565b6000606080600080600061063c88610c1f565b6000818152600160208190526040909120600c8101549293509160ff161515146106785760405162461bcd60e51b81526004016104a19061142e565b6001600160a01b038089166000818152600d84016020526040902054909116146106d85760405162461bcd60e51b815260206004820152601160248201527014da59db995c9cc81b9bdd08195e1a5cdd607a1b60448201526064016104a1565b6001600160a01b0388166000908152600d8201602052604090206001810154600482015460058301546002840180549394909360039091019260ff1691908490610721906112ae565b80601f016020809104026020016040519081016040528092919081815260200182805461074d906112ae565b801561079a5780601f1061076f5761010080835404028352916020019161079a565b820191906000526020600020905b81548152906001019060200180831161077d57829003601f168201915b505050505093508280546107ad906112ae565b80601f01602080910402602001604051908101604052809291908181526020018280546107d9906112ae565b80156108265780601f106107fb57610100808354040283529160200191610826565b820191906000526020600020905b81548152906001019060200180831161080957829003601f168201915b505050505092509650965096509650965050509295509295909350565b60008060608060608060606000806000806000806108608e610c1f565b6000818152600160208190526040909120600c8101549293509160ff1615151461089c5760405162461bcd60e51b81526004016104a19061142e565b80600101548160030160009054906101000a90046001600160a01b0316826002018360040184600501856006018660070187600801548860090160009054906101000a900460ff1689600a01548a600b01548b600c0160009054906101000a900460ff1689805461090c906112ae565b80601f0160208091040260200160405190810160405280929190818152602001828054610938906112ae565b80156109855780601f1061095a57610100808354040283529160200191610985565b820191906000526020600020905b81548152906001019060200180831161096857829003601f168201915b50505050509950888054610998906112ae565b80601f01602080910402602001604051908101604052809291908181526020018280546109c4906112ae565b8015610a115780601f106109e657610100808354040283529160200191610a11565b820191906000526020600020905b8154815290600101906020018083116109f457829003601f168201915b50505050509850878054610a24906112ae565b80601f0160208091040260200160405190810160405280929190818152602001828054610a50906112ae565b8015610a9d5780601f10610a7257610100808354040283529160200191610a9d565b820191906000526020600020905b815481529060010190602001808311610a8057829003601f168201915b50505050509750868054610ab0906112ae565b80601f0160208091040260200160405190810160405280929190818152602001828054610adc906112ae565b8015610b295780601f10610afe57610100808354040283529160200191610b29565b820191906000526020600020905b815481529060010190602001808311610b0c57829003601f168201915b50505050509650858054610b3c906112ae565b80601f0160208091040260200160405190810160405280929190818152602001828054610b68906112ae565b8015610bb55780601f10610b8a57610100808354040283529160200191610bb5565b820191906000526020600020905b815481529060010190602001808311610b9857829003601f168201915b505050505095509d509d509d509d509d509d509d509d509d509d509d509d50505091939597999b5091939597999b565b600080610bf183610c1f565b600081815260026020526040902054909150610c105750600092915050565b50600192915050565b50919050565b805160009082908203610c355750600092915050565b50506020015190565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b6020811015610cc057838160208110610c8157610c816113f1565b1a60f81b828281518110610c9757610c976113f1565b60200101906001600160f81b031916908160001a90535080610cb881611407565b915050610c66565b5092915050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610d0657610d06610cc7565b604052919050565b600082601f830112610d1f57600080fd5b813567ffffffffffffffff811115610d3957610d39610cc7565b610d4c601f8201601f1916602001610cdd565b818152846020838601011115610d6157600080fd5b816020850160208301376000918101602001919091529392505050565b80356001600160a01b0381168114610d9557600080fd5b919050565b80358015158114610d9557600080fd5b600067ffffffffffffffff821115610dc457610dc4610cc7565b5060051b60200190565b600082601f830112610ddf57600080fd5b81356020610df4610def83610daa565b610cdd565b82815260059290921b84018101918181019086841115610e1357600080fd5b8286015b84811015610e3557610e2881610d7e565b8352918301918301610e17565b509695505050505050565b600082601f830112610e5157600080fd5b81356020610e61610def83610daa565b82815260059290921b84018101918181019086841115610e8057600080fd5b8286015b84811015610e3557803567ffffffffffffffff811115610ea45760008081fd5b610eb28986838b0101610d0e565b845250918301918301610e84565b60008060008060008060008060008060006101608c8e031215610ee257600080fd5b67ffffffffffffffff808d351115610ef957600080fd5b610f068e8e358f01610d0e565b9b50610f1460208e01610d7e565b9a508060408e01351115610f2757600080fd5b610f378e60408f01358f01610d0e565b99508060608e01351115610f4a57600080fd5b610f5a8e60608f01358f01610d0e565b98508060808e01351115610f6d57600080fd5b610f7d8e60808f01358f01610d0e565b97508060a08e01351115610f9057600080fd5b610fa08e60a08f01358f01610d0e565b965060c08d01359550610fb560e08e01610d9a565b94506101008d01359350806101208e01351115610fd157600080fd5b610fe28e6101208f01358f01610dce565b9250806101408e01351115610ff657600080fd5b506110088d6101408e01358e01610e40565b90509295989b509295989b9093969950565b600080600080600060a0868803121561103257600080fd5b853567ffffffffffffffff8082111561104a57600080fd5b61105689838a01610d0e565b965061106460208901610d7e565b9550604088013591508082111561107a57600080fd5b61108689838a01610d0e565b9450606088013591508082111561109c57600080fd5b506110a988828901610d0e565b95989497509295608001359392505050565b6000602082840312156110cd57600080fd5b813567ffffffffffffffff8111156110e457600080fd5b6110f084828501610d0e565b949350505050565b6000815180845260005b8181101561111e57602081850181015186830182015201611102565b81811115611130576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061115860208301846110f8565b9392505050565b6000806040838503121561117257600080fd5b823567ffffffffffffffff81111561118957600080fd5b61119585828601610d0e565b9250506111a460208401610d7e565b90509250929050565b85815260a0602082015260006111c660a08301876110f8565b82810360408401526111d881876110f8565b94151560608401525050608001529392505050565b8c81526001600160a01b038c166020820152610180604082018190526000906112188382018e6110f8565b9050828103606084015261122c818d6110f8565b90508281036080840152611240818c6110f8565b905082810360a0840152611254818b6110f8565b905082810360c0840152611268818a6110f8565b9150508660e083015261128061010083018715159052565b846101208301528361014083015261129d61016083018415159052565b9d9c50505050505050505050505050565b600181811c908216806112c257607f821691505b602082108103610c1957634e487b7160e01b600052602260045260246000fd5b601f82111561132c57600081815260208120601f850160051c810160208610156113095750805b601f850160051c820191505b8181101561132857828155600101611315565b5050505b505050565b815167ffffffffffffffff81111561134b5761134b610cc7565b61135f8161135984546112ae565b846112e2565b602080601f831160018114611394576000841561137c5750858301515b600019600386901b1c1916600185901b178555611328565b600085815260208120601f198616915b828110156113c3578886015182559484019460019091019084016113a4565b50858210156113e15787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b60006001820161142757634e487b7160e01b600052601160045260246000fd5b5060010190565b602080825260129082015271111bd8dd5b595b9d081b9bdd08195e1a5cdd60721b60408201526060019056fea264697066735822122096d9005dddaa8dc092cef4cf0412d1bd323945d58de553de625f42a2217faf7964736f6c634300080f0033",
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
// Solidity: function getDoc(string _file) view returns(uint256, address, string, string, string, string, string, uint256, bool, uint256, uint256, bool)
func (_Api *ApiCaller) GetDoc(opts *bind.CallOpts, _file string) (*big.Int, common.Address, string, string, string, string, string, *big.Int, bool, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getDoc", _file)

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new(string), *new(string), *new(string), *new(string), *new(string), *new(*big.Int), *new(bool), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)
	out5 := *abi.ConvertType(out[5], new(string)).(*string)
	out6 := *abi.ConvertType(out[6], new(string)).(*string)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new(bool)).(*bool)
	out9 := *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	out10 := *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	out11 := *abi.ConvertType(out[11], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, out10, out11, err

}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string _file) view returns(uint256, address, string, string, string, string, string, uint256, bool, uint256, uint256, bool)
func (_Api *ApiSession) GetDoc(_file string) (*big.Int, common.Address, string, string, string, string, string, *big.Int, bool, *big.Int, *big.Int, bool, error) {
	return _Api.Contract.GetDoc(&_Api.CallOpts, _file)
}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string _file) view returns(uint256, address, string, string, string, string, string, uint256, bool, uint256, uint256, bool)
func (_Api *ApiCallerSession) GetDoc(_file string) (*big.Int, common.Address, string, string, string, string, string, *big.Int, bool, *big.Int, *big.Int, bool, error) {
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

// Create is a paid mutator transaction binding the contract method 0x04205959.
//
// Solidity: function create(string _file, address creator, string _creator_id, string _metadata, string _hash, string _ipfs, uint256 _state, bool _visibility, uint256 _time, address[] _signers, string[] _signers_id) returns()
func (_Api *ApiTransactor) Create(opts *bind.TransactOpts, _file string, creator common.Address, _creator_id string, _metadata string, _hash string, _ipfs string, _state *big.Int, _visibility bool, _time *big.Int, _signers []common.Address, _signers_id []string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "create", _file, creator, _creator_id, _metadata, _hash, _ipfs, _state, _visibility, _time, _signers, _signers_id)
}

// Create is a paid mutator transaction binding the contract method 0x04205959.
//
// Solidity: function create(string _file, address creator, string _creator_id, string _metadata, string _hash, string _ipfs, uint256 _state, bool _visibility, uint256 _time, address[] _signers, string[] _signers_id) returns()
func (_Api *ApiSession) Create(_file string, creator common.Address, _creator_id string, _metadata string, _hash string, _ipfs string, _state *big.Int, _visibility bool, _time *big.Int, _signers []common.Address, _signers_id []string) (*types.Transaction, error) {
	return _Api.Contract.Create(&_Api.TransactOpts, _file, creator, _creator_id, _metadata, _hash, _ipfs, _state, _visibility, _time, _signers, _signers_id)
}

// Create is a paid mutator transaction binding the contract method 0x04205959.
//
// Solidity: function create(string _file, address creator, string _creator_id, string _metadata, string _hash, string _ipfs, uint256 _state, bool _visibility, uint256 _time, address[] _signers, string[] _signers_id) returns()
func (_Api *ApiTransactorSession) Create(_file string, creator common.Address, _creator_id string, _metadata string, _hash string, _ipfs string, _state *big.Int, _visibility bool, _time *big.Int, _signers []common.Address, _signers_id []string) (*types.Transaction, error) {
	return _Api.Contract.Create(&_Api.TransactOpts, _file, creator, _creator_id, _metadata, _hash, _ipfs, _state, _visibility, _time, _signers, _signers_id)
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
