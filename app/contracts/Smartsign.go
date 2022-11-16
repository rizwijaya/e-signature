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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_creator_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_state\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_visibility\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_signers_id\",\"type\":\"string[]\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"}],\"name\":\"getDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_signers_id\",\"type\":\"address\"}],\"name\":\"getSign\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_signers_id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_signers_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"signDoc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506001600055611425806100256000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c8063042059591461005c5780631af2925c146100715780633d99587714610084578063a5bde23b146100b1578063c7335804146100dc575b600080fd5b61006f61006a366004610e56565b6100fc565b005b61006f61007f366004610fb0565b61042c565b610097610092366004611051565b6105ca565b6040516100a89594939291906110ec565b60405180910390f35b6100c46100bf36600461112c565b6107e4565b6040516100a89c9b9a99989796959493929190611169565b6100ef6100ea36600461112c565b610b86565b6040516100a8919061122a565b60006101078c610bb5565b600081815260016020819052604082209154908201558181556003810180546001600160a01b0319166001600160a01b038f161790559091506002810161014e8c826112c6565b506004810161015d8b826112c6565b506005810161016c8e826112c6565b506006810161017b8a826112c6565b506007810161018a89826112c6565b506008810187905560098101805487151560ff1991821617909155600a8201869055600b8201869055600c82018054909116600117905560005b8451811015610408578481815181106101df576101df611386565b602002602001015182600d0160008784815181106101ff576101ff611386565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508082600d01600087848151811061026757610267611386565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600101819055508381815181106102a8576102a8611386565b602002602001015182600d0160008784815181106102c8576102c8611386565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060020190816102ff91906112c6565b508982600d01600087848151811061031957610319611386565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600301908161035091906112c6565b50600082600d01600087848151811061036b5761036b611386565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060040160006101000a81548160ff0219169083151502179055508582600d0160008784815181106103c6576103c6611386565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206005018190555080806104009061139c565b9150506101c4565b506000805490806104188361139c565b919050555050505050505050505050505050565b600061043786610bb5565b6000818152600160208190526040909120600c8101549293509160ff1615151461047c5760405162461bcd60e51b8152600401610473906113c3565b60405180910390fd5b6001600160a01b0386166000908152600d820160205260409020600501546001106104b95760405162461bcd60e51b8152600401610473906113c3565b6001600160a01b0386166000908152600d8201602052604090206004015460ff16156105275760405162461bcd60e51b815260206004820152601c60248201527f596f7520617265207369676e6564207468697320646f63756d656e74000000006044820152606401610473565b6007810161053585826112c6565b50600b81018390556006810161054b86826112c6565b506001600160a01b0386166000908152600d82016020526040902060030161057386826112c6565b506001600160a01b0386166000908152600d82016020526040812060048101805460ff191660011790556005018490556105ac86610bb5565b91546000928352600260205260409092209190915550505050505050565b600060608060008060006105dd88610bb5565b6000818152600160208190526040909120600c8101549293509160ff161515146106195760405162461bcd60e51b8152600401610473906113c3565b6001600160a01b038089166000818152600d84016020526040902054909116146106795760405162461bcd60e51b815260206004820152601160248201527014da59db995c9cc81b9bdd08195e1a5cdd607a1b6044820152606401610473565b6001600160a01b0388166000908152600d8201602052604090206001810154600482015460058301546002840180549394909360039091019260ff16919084906106c29061123d565b80601f01602080910402602001604051908101604052809291908181526020018280546106ee9061123d565b801561073b5780601f106107105761010080835404028352916020019161073b565b820191906000526020600020905b81548152906001019060200180831161071e57829003601f168201915b5050505050935082805461074e9061123d565b80601f016020809104026020016040519081016040528092919081815260200182805461077a9061123d565b80156107c75780601f1061079c576101008083540402835291602001916107c7565b820191906000526020600020905b8154815290600101906020018083116107aa57829003601f168201915b505050505092509650965096509650965050509295509295909350565b60008060608060608060606000806000806000806108018e610bb5565b6000818152600160208190526040909120600c8101549293509160ff1615151461083d5760405162461bcd60e51b8152600401610473906113c3565b80600101548160030160009054906101000a90046001600160a01b0316826002018360040184600501856006018660070187600801548860090160009054906101000a900460ff1689600a01548a600b01548b600c0160009054906101000a900460ff168980546108ad9061123d565b80601f01602080910402602001604051908101604052809291908181526020018280546108d99061123d565b80156109265780601f106108fb57610100808354040283529160200191610926565b820191906000526020600020905b81548152906001019060200180831161090957829003601f168201915b505050505099508880546109399061123d565b80601f01602080910402602001604051908101604052809291908181526020018280546109659061123d565b80156109b25780601f10610987576101008083540402835291602001916109b2565b820191906000526020600020905b81548152906001019060200180831161099557829003601f168201915b505050505098508780546109c59061123d565b80601f01602080910402602001604051908101604052809291908181526020018280546109f19061123d565b8015610a3e5780601f10610a1357610100808354040283529160200191610a3e565b820191906000526020600020905b815481529060010190602001808311610a2157829003601f168201915b50505050509750868054610a519061123d565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7d9061123d565b8015610aca5780601f10610a9f57610100808354040283529160200191610aca565b820191906000526020600020905b815481529060010190602001808311610aad57829003601f168201915b50505050509650858054610add9061123d565b80601f0160208091040260200160405190810160405280929190818152602001828054610b099061123d565b8015610b565780601f10610b2b57610100808354040283529160200191610b56565b820191906000526020600020905b815481529060010190602001808311610b3957829003601f168201915b505050505095509d509d509d509d509d509d509d509d509d509d509d509d50505091939597999b5091939597999b565b60606000610b9383610bb5565b600081815260026020526040902054909150610bae90610bd4565b9392505050565b805160009082908203610bcb5750600092915050565b50506020015190565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b6020811015610c5657838160208110610c1757610c17611386565b1a60f81b828281518110610c2d57610c2d611386565b60200101906001600160f81b031916908160001a90535080610c4e8161139c565b915050610bfc565b5092915050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610c9c57610c9c610c5d565b604052919050565b600082601f830112610cb557600080fd5b813567ffffffffffffffff811115610ccf57610ccf610c5d565b610ce2601f8201601f1916602001610c73565b818152846020838601011115610cf757600080fd5b816020850160208301376000918101602001919091529392505050565b80356001600160a01b0381168114610d2b57600080fd5b919050565b80358015158114610d2b57600080fd5b600067ffffffffffffffff821115610d5a57610d5a610c5d565b5060051b60200190565b600082601f830112610d7557600080fd5b81356020610d8a610d8583610d40565b610c73565b82815260059290921b84018101918181019086841115610da957600080fd5b8286015b84811015610dcb57610dbe81610d14565b8352918301918301610dad565b509695505050505050565b600082601f830112610de757600080fd5b81356020610df7610d8583610d40565b82815260059290921b84018101918181019086841115610e1657600080fd5b8286015b84811015610dcb57803567ffffffffffffffff811115610e3a5760008081fd5b610e488986838b0101610ca4565b845250918301918301610e1a565b60008060008060008060008060008060006101608c8e031215610e7857600080fd5b67ffffffffffffffff808d351115610e8f57600080fd5b610e9c8e8e358f01610ca4565b9b50610eaa60208e01610d14565b9a508060408e01351115610ebd57600080fd5b610ecd8e60408f01358f01610ca4565b99508060608e01351115610ee057600080fd5b610ef08e60608f01358f01610ca4565b98508060808e01351115610f0357600080fd5b610f138e60808f01358f01610ca4565b97508060a08e01351115610f2657600080fd5b610f368e60a08f01358f01610ca4565b965060c08d01359550610f4b60e08e01610d30565b94506101008d01359350806101208e01351115610f6757600080fd5b610f788e6101208f01358f01610d64565b9250806101408e01351115610f8c57600080fd5b50610f9e8d6101408e01358e01610dd6565b90509295989b509295989b9093969950565b600080600080600060a08688031215610fc857600080fd5b853567ffffffffffffffff80821115610fe057600080fd5b610fec89838a01610ca4565b9650610ffa60208901610d14565b9550604088013591508082111561101057600080fd5b61101c89838a01610ca4565b9450606088013591508082111561103257600080fd5b5061103f88828901610ca4565b95989497509295608001359392505050565b6000806040838503121561106457600080fd5b823567ffffffffffffffff81111561107b57600080fd5b61108785828601610ca4565b92505061109660208401610d14565b90509250929050565b6000815180845260005b818110156110c5576020818501810151868301820152016110a9565b818111156110d7576000602083870101525b50601f01601f19169290920160200192915050565b85815260a06020820152600061110560a083018761109f565b8281036040840152611117818761109f565b94151560608401525050608001529392505050565b60006020828403121561113e57600080fd5b813567ffffffffffffffff81111561115557600080fd5b61116184828501610ca4565b949350505050565b8c81526001600160a01b038c166020820152610180604082018190526000906111948382018e61109f565b905082810360608401526111a8818d61109f565b905082810360808401526111bc818c61109f565b905082810360a08401526111d0818b61109f565b905082810360c08401526111e4818a61109f565b9150508660e08301526111fc61010083018715159052565b846101208301528361014083015261121961016083018415159052565b9d9c50505050505050505050505050565b602081526000610bae602083018461109f565b600181811c9082168061125157607f821691505b60208210810361127157634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156112c157600081815260208120601f850160051c8101602086101561129e5750805b601f850160051c820191505b818110156112bd578281556001016112aa565b5050505b505050565b815167ffffffffffffffff8111156112e0576112e0610c5d565b6112f4816112ee845461123d565b84611277565b602080601f83116001811461132957600084156113115750858301515b600019600386901b1c1916600185901b1785556112bd565b600085815260208120601f198616915b8281101561135857888601518255948401946001909101908401611339565b50858210156113765787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b6000600182016113bc57634e487b7160e01b600052601160045260246000fd5b5060010190565b602080825260129082015271111bd8dd5b595b9d081b9bdd08195e1a5cdd60721b60408201526060019056fea264697066735822122056101d14bf51bff9c7acd4953c0f58744472288edc922674a2e74d9c2f2c3df764736f6c634300080f0033",
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
