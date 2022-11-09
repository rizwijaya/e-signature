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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_state\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_visibility\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_signers_id\",\"type\":\"string[]\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"}],\"name\":\"getDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_signers_id\",\"type\":\"address\"}],\"name\":\"getSign\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_signers_id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_signers_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"signDoc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506001600055611266806100256000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80631af2925c1461005c5780633d99587714610071578063a5bde23b1461009e578063c7335804146100c7578063f913c6de146100e7575b600080fd5b61006f61006a366004610bda565b6100fa565b005b61008461007f366004610c7b565b610298565b604051610095959493929190610d16565b60405180910390f35b6100b16100ac366004610d56565b6104b2565b6040516100959a99989796959493929190610d93565b6100da6100d5366004610d56565b61071f565b6040516100959190610e18565b61006f6100f5366004610f51565b61074e565b600061010586610a5f565b6000818152600160208190526040909120600a8101549293509160ff1615151461014a5760405162461bcd60e51b81526004016101419061107e565b60405180910390fd5b6001600160a01b0386166000908152600b820160205260409020600501546001106101875760405162461bcd60e51b81526004016101419061107e565b6001600160a01b0386166000908152600b8201602052604090206004015460ff16156101f55760405162461bcd60e51b815260206004820152601c60248201527f596f7520617265207369676e6564207468697320646f63756d656e74000000006044820152606401610141565b600581016102038582611133565b5060098101839055600481016102198682611133565b506001600160a01b0386166000908152600b8201602052604090206003016102418682611133565b506001600160a01b0386166000908152600b82016020526040812060048101805460ff1916600117905560050184905561027a86610a5f565b91546000928352600260205260409092209190915550505050505050565b600060608060008060006102ab88610a5f565b6000818152600160208190526040909120600a8101549293509160ff161515146102e75760405162461bcd60e51b81526004016101419061107e565b6001600160a01b038089166000818152600b84016020526040902054909116146103475760405162461bcd60e51b815260206004820152601160248201527014da59db995c9cc81b9bdd08195e1a5cdd607a1b6044820152606401610141565b6001600160a01b0388166000908152600b8201602052604090206001810154600482015460058301546002840180549394909360039091019260ff1691908490610390906110aa565b80601f01602080910402602001604051908101604052809291908181526020018280546103bc906110aa565b80156104095780601f106103de57610100808354040283529160200191610409565b820191906000526020600020905b8154815290600101906020018083116103ec57829003601f168201915b5050505050935082805461041c906110aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610448906110aa565b80156104955780601f1061046a57610100808354040283529160200191610495565b820191906000526020600020905b81548152906001019060200180831161047857829003601f168201915b505050505092509650965096509650965050509295509295909350565b60008060608060606000806000806000806104cc8c610a5f565b6000818152600160208190526040909120600a8101549293509160ff161515146105085760405162461bcd60e51b81526004016101419061107e565b600181015460028201546006830154600784015460088501546009860154600a8701546003880180546001600160a01b0390971696909560048a019560058b0195919460ff9182169492939092909116908890610564906110aa565b80601f0160208091040260200160405190810160405280929190818152602001828054610590906110aa565b80156105dd5780601f106105b2576101008083540402835291602001916105dd565b820191906000526020600020905b8154815290600101906020018083116105c057829003601f168201915b505050505097508680546105f0906110aa565b80601f016020809104026020016040519081016040528092919081815260200182805461061c906110aa565b80156106695780601f1061063e57610100808354040283529160200191610669565b820191906000526020600020905b81548152906001019060200180831161064c57829003601f168201915b5050505050965085805461067c906110aa565b80601f01602080910402602001604051908101604052809291908181526020018280546106a8906110aa565b80156106f55780601f106106ca576101008083540402835291602001916106f5565b820191906000526020600020905b8154815290600101906020018083116106d857829003601f168201915b505050505095509b509b509b509b509b509b509b509b509b509b5050509193959799509193959799565b6060600061072c83610a5f565b60008181526002602052604090205490915061074790610a7e565b9392505050565b60006107598b610a5f565b600081815260016020819052604082209154908201558181556002810180546001600160a01b0319166001600160a01b038e16179055909150600381016107a08b82611133565b50600481016107af8a82611133565b50600581016107be8982611133565b506006810187905560078101805487151560ff19918216179091556008820186905560098201869055600a82018054909116600117905560005b8451811015610a3c57848181518110610813576108136111f3565b602002602001015182600b016000878481518110610833576108336111f3565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508082600b01600087848151811061089b5761089b6111f3565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600101819055508381815181106108dc576108dc6111f3565b602002602001015182600b0160008784815181106108fc576108fc6111f3565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060020190816109339190611133565b508982600b01600087848151811061094d5761094d6111f3565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060030190816109849190611133565b50600082600b01600087848151811061099f5761099f6111f3565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060040160006101000a81548160ff0219169083151502179055508582600b0160008784815181106109fa576109fa6111f3565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600501819055508080610a3490611209565b9150506107f8565b50600080549080610a4c83611209565b9190505550505050505050505050505050565b805160009082908203610a755750600092915050565b50506020015190565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b6020811015610b0057838160208110610ac157610ac16111f3565b1a60f81b828281518110610ad757610ad76111f3565b60200101906001600160f81b031916908160001a90535080610af881611209565b915050610aa6565b5092915050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610b4657610b46610b07565b604052919050565b600082601f830112610b5f57600080fd5b813567ffffffffffffffff811115610b7957610b79610b07565b610b8c601f8201601f1916602001610b1d565b818152846020838601011115610ba157600080fd5b816020850160208301376000918101602001919091529392505050565b80356001600160a01b0381168114610bd557600080fd5b919050565b600080600080600060a08688031215610bf257600080fd5b853567ffffffffffffffff80821115610c0a57600080fd5b610c1689838a01610b4e565b9650610c2460208901610bbe565b95506040880135915080821115610c3a57600080fd5b610c4689838a01610b4e565b94506060880135915080821115610c5c57600080fd5b50610c6988828901610b4e565b95989497509295608001359392505050565b60008060408385031215610c8e57600080fd5b823567ffffffffffffffff811115610ca557600080fd5b610cb185828601610b4e565b925050610cc060208401610bbe565b90509250929050565b6000815180845260005b81811015610cef57602081850181015186830182015201610cd3565b81811115610d01576000602083870101525b50601f01601f19169290920160200192915050565b85815260a060208201526000610d2f60a0830187610cc9565b8281036040840152610d418187610cc9565b94151560608401525050608001529392505050565b600060208284031215610d6857600080fd5b813567ffffffffffffffff811115610d7f57600080fd5b610d8b84828501610b4e565b949350505050565b8a81526001600160a01b038a16602082015261014060408201819052600090610dbe8382018c610cc9565b90508281036060840152610dd2818b610cc9565b90508281036080840152610de6818a610cc9565b60a0840198909852505093151560c085015260e084019290925261010083015215156101209091015295945050505050565b6020815260006107476020830184610cc9565b80358015158114610bd557600080fd5b600067ffffffffffffffff821115610e5557610e55610b07565b5060051b60200190565b600082601f830112610e7057600080fd5b81356020610e85610e8083610e3b565b610b1d565b82815260059290921b84018101918181019086841115610ea457600080fd5b8286015b84811015610ec657610eb981610bbe565b8352918301918301610ea8565b509695505050505050565b600082601f830112610ee257600080fd5b81356020610ef2610e8083610e3b565b82815260059290921b84018101918181019086841115610f1157600080fd5b8286015b84811015610ec657803567ffffffffffffffff811115610f355760008081fd5b610f438986838b0101610b4e565b845250918301918301610f15565b6000806000806000806000806000806101408b8d031215610f7157600080fd5b8a3567ffffffffffffffff80821115610f8957600080fd5b610f958e838f01610b4e565b9b50610fa360208e01610bbe565b9a5060408d0135915080821115610fb957600080fd5b610fc58e838f01610b4e565b995060608d0135915080821115610fdb57600080fd5b610fe78e838f01610b4e565b985060808d0135915080821115610ffd57600080fd5b6110098e838f01610b4e565b975060a08d0135965061101e60c08e01610e2b565b955060e08d013594506101008d013591508082111561103c57600080fd5b6110488e838f01610e5f565b93506101208d013591508082111561105f57600080fd5b5061106c8d828e01610ed1565b9150509295989b9194979a5092959850565b602080825260129082015271111bd8dd5b595b9d081b9bdd08195e1a5cdd60721b604082015260600190565b600181811c908216806110be57607f821691505b6020821081036110de57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561112e57600081815260208120601f850160051c8101602086101561110b5750805b601f850160051c820191505b8181101561112a57828155600101611117565b5050505b505050565b815167ffffffffffffffff81111561114d5761114d610b07565b6111618161115b84546110aa565b846110e4565b602080601f831160018114611196576000841561117e5750858301515b600019600386901b1c1916600185901b17855561112a565b600085815260208120601f198616915b828110156111c5578886015182559484019460019091019084016111a6565b50858210156111e35787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b60006001820161122957634e487b7160e01b600052601160045260246000fd5b506001019056fea2646970667358221220f2259a2a1bf2f1871414a965dcaa30ac5c07d00fcf7197ec6b194a3605df471464736f6c634300080f0033",
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
// Solidity: function getDoc(string _file) view returns(uint256, address, string, string, string, uint256, bool, uint256, uint256, bool)
func (_Api *ApiCaller) GetDoc(opts *bind.CallOpts, _file string) (*big.Int, common.Address, string, string, string, *big.Int, bool, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getDoc", _file)

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new(string), *new(string), *new(string), *new(*big.Int), *new(bool), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(bool)).(*bool)
	out7 := *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	out8 := *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	out9 := *abi.ConvertType(out[9], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, out6, out7, out8, out9, err

}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string _file) view returns(uint256, address, string, string, string, uint256, bool, uint256, uint256, bool)
func (_Api *ApiSession) GetDoc(_file string) (*big.Int, common.Address, string, string, string, *big.Int, bool, *big.Int, *big.Int, bool, error) {
	return _Api.Contract.GetDoc(&_Api.CallOpts, _file)
}

// GetDoc is a free data retrieval call binding the contract method 0xa5bde23b.
//
// Solidity: function getDoc(string _file) view returns(uint256, address, string, string, string, uint256, bool, uint256, uint256, bool)
func (_Api *ApiCallerSession) GetDoc(_file string) (*big.Int, common.Address, string, string, string, *big.Int, bool, *big.Int, *big.Int, bool, error) {
	return _Api.Contract.GetDoc(&_Api.CallOpts, _file)
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
// Solidity: function verifyDoc(string _hash) view returns(string)
func (_Api *ApiCaller) VerifyDoc(opts *bind.CallOpts, _hash string) (string, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "verifyDoc", _hash)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VerifyDoc is a free data retrieval call binding the contract method 0xc7335804.
//
// Solidity: function verifyDoc(string _hash) view returns(string)
func (_Api *ApiSession) VerifyDoc(_hash string) (string, error) {
	return _Api.Contract.VerifyDoc(&_Api.CallOpts, _hash)
}

// VerifyDoc is a free data retrieval call binding the contract method 0xc7335804.
//
// Solidity: function verifyDoc(string _hash) view returns(string)
func (_Api *ApiCallerSession) VerifyDoc(_hash string) (string, error) {
	return _Api.Contract.VerifyDoc(&_Api.CallOpts, _hash)
}

// Create is a paid mutator transaction binding the contract method 0xf913c6de.
//
// Solidity: function create(string _file, address creator, string _metadata, string _hash, string _ipfs, uint256 _state, bool _visibility, uint256 _time, address[] _signers, string[] _signers_id) returns()
func (_Api *ApiTransactor) Create(opts *bind.TransactOpts, _file string, creator common.Address, _metadata string, _hash string, _ipfs string, _state *big.Int, _visibility bool, _time *big.Int, _signers []common.Address, _signers_id []string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "create", _file, creator, _metadata, _hash, _ipfs, _state, _visibility, _time, _signers, _signers_id)
}

// Create is a paid mutator transaction binding the contract method 0xf913c6de.
//
// Solidity: function create(string _file, address creator, string _metadata, string _hash, string _ipfs, uint256 _state, bool _visibility, uint256 _time, address[] _signers, string[] _signers_id) returns()
func (_Api *ApiSession) Create(_file string, creator common.Address, _metadata string, _hash string, _ipfs string, _state *big.Int, _visibility bool, _time *big.Int, _signers []common.Address, _signers_id []string) (*types.Transaction, error) {
	return _Api.Contract.Create(&_Api.TransactOpts, _file, creator, _metadata, _hash, _ipfs, _state, _visibility, _time, _signers, _signers_id)
}

// Create is a paid mutator transaction binding the contract method 0xf913c6de.
//
// Solidity: function create(string _file, address creator, string _metadata, string _hash, string _ipfs, uint256 _state, bool _visibility, uint256 _time, address[] _signers, string[] _signers_id) returns()
func (_Api *ApiTransactorSession) Create(_file string, creator common.Address, _metadata string, _hash string, _ipfs string, _state *big.Int, _visibility bool, _time *big.Int, _signers []common.Address, _signers_id []string) (*types.Transaction, error) {
	return _Api.Contract.Create(&_Api.TransactOpts, _file, creator, _metadata, _hash, _ipfs, _state, _visibility, _time, _signers, _signers_id)
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
