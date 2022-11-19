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
	Bin: "0x608060405234801561001057600080fd5b5060016000556114e2806100256000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806304205959146100675780631af2925c1461007c5780632887590e1461008f5780633d995877146100b8578063a5bde23b146100dc578063c733580414610107575b600080fd5b61007a610075366004610ef3565b61012a565b005b61007a61008a36600461104d565b610461565b6100a261009d3660046110ee565b61062b565b6040516100af9190611178565b60405180910390f35b6100cb6100c6366004611192565b61065c565b6040516100af9594939291906111e0565b6100ef6100ea3660046110ee565b610876565b6040516100af9c9b9a99989796959493929190611220565b61011a6101153660046110ee565b610c18565b60405190151581526020016100af565b60006101358c610c52565b600081815260016020819052604082209154908201558181556003810180546001600160a01b0319166001600160a01b038f161790559091506002810161017c8c82611364565b506004810161018b8b82611364565b506005810161019a8e82611364565b50600681016101a98a82611364565b50600781016101b88982611364565b506008810187905560098101805487151560ff1991821617909155600a8201869055600b8201869055600c8201805490911660011790558351600d82015560005b845181101561043d5784818151811061021457610214611424565b602002602001015182600e01600087848151811061023457610234611424565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160006101000a8154816001600160a01b0302191690836001600160a01b031602179055508082600e01600087848151811061029c5761029c611424565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600101819055508381815181106102dd576102dd611424565b602002602001015182600e0160008784815181106102fd576102fd611424565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060020190816103349190611364565b508982600e01600087848151811061034e5761034e611424565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060030190816103859190611364565b50600082600e0160008784815181106103a0576103a0611424565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060040160006101000a81548160ff0219169083151502179055508582600e0160008784815181106103fb576103fb611424565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060050181905550808061043590611450565b9150506101f9565b5060008054908061044d83611450565b919050555050505050505050505050505050565b600061046c86610c52565b6000818152600160208190526040909120600c8101549293509160ff161515146104b15760405162461bcd60e51b81526004016104a890611469565b60405180910390fd5b6001600160a01b0386166000908152600e820160205260409020600501546001106104ee5760405162461bcd60e51b81526004016104a890611469565b6001600160a01b0386166000908152600e8201602052604090206004015460ff161561055c5760405162461bcd60e51b815260206004820152601c60248201527f596f7520617265207369676e6564207468697320646f63756d656e740000000060448201526064016104a8565b600d8101805490600061056e83611495565b9091555050600781016105818582611364565b50600b8101839055600681016105978682611364565b506001600160a01b0386166000908152600e8201602052604090206003016105bf8682611364565b506001600160a01b0386166000908152600e82016020526040812060048101805460ff191660011790556005018490556105f886610c52565b8254600082815260026020526040902055600d8301549091506001111561062157600260088301555b5050505050505050565b6060600061063883610c52565b60008181526002602052604081205491925061065382610c71565b95945050505050565b6000606080600080600061066f88610c52565b6000818152600160208190526040909120600c8101549293509160ff161515146106ab5760405162461bcd60e51b81526004016104a890611469565b6001600160a01b038089166000818152600e840160205260409020549091161461070b5760405162461bcd60e51b815260206004820152601160248201527014da59db995c9cc81b9bdd08195e1a5cdd607a1b60448201526064016104a8565b6001600160a01b0388166000908152600e8201602052604090206001810154600482015460058301546002840180549394909360039091019260ff1691908490610754906112e1565b80601f0160208091040260200160405190810160405280929190818152602001828054610780906112e1565b80156107cd5780601f106107a2576101008083540402835291602001916107cd565b820191906000526020600020905b8154815290600101906020018083116107b057829003601f168201915b505050505093508280546107e0906112e1565b80601f016020809104026020016040519081016040528092919081815260200182805461080c906112e1565b80156108595780601f1061082e57610100808354040283529160200191610859565b820191906000526020600020905b81548152906001019060200180831161083c57829003601f168201915b505050505092509650965096509650965050509295509295909350565b60008060608060608060606000806000806000806108938e610c52565b6000818152600160208190526040909120600c8101549293509160ff161515146108cf5760405162461bcd60e51b81526004016104a890611469565b80600101548160030160009054906101000a90046001600160a01b0316826002018360040184600501856006018660070187600801548860090160009054906101000a900460ff1689600a01548a600b01548b600c0160009054906101000a900460ff1689805461093f906112e1565b80601f016020809104026020016040519081016040528092919081815260200182805461096b906112e1565b80156109b85780601f1061098d576101008083540402835291602001916109b8565b820191906000526020600020905b81548152906001019060200180831161099b57829003601f168201915b505050505099508880546109cb906112e1565b80601f01602080910402602001604051908101604052809291908181526020018280546109f7906112e1565b8015610a445780601f10610a1957610100808354040283529160200191610a44565b820191906000526020600020905b815481529060010190602001808311610a2757829003601f168201915b50505050509850878054610a57906112e1565b80601f0160208091040260200160405190810160405280929190818152602001828054610a83906112e1565b8015610ad05780601f10610aa557610100808354040283529160200191610ad0565b820191906000526020600020905b815481529060010190602001808311610ab357829003601f168201915b50505050509750868054610ae3906112e1565b80601f0160208091040260200160405190810160405280929190818152602001828054610b0f906112e1565b8015610b5c5780601f10610b3157610100808354040283529160200191610b5c565b820191906000526020600020905b815481529060010190602001808311610b3f57829003601f168201915b50505050509650858054610b6f906112e1565b80601f0160208091040260200160405190810160405280929190818152602001828054610b9b906112e1565b8015610be85780601f10610bbd57610100808354040283529160200191610be8565b820191906000526020600020905b815481529060010190602001808311610bcb57829003601f168201915b505050505095509d509d509d509d509d509d509d509d509d509d509d509d50505091939597999b5091939597999b565b600080610c2483610c52565b600081815260026020526040902054909150610c435750600092915050565b50600192915050565b50919050565b805160009082908203610c685750600092915050565b50506020015190565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b6020811015610cf357838160208110610cb457610cb4611424565b1a60f81b828281518110610cca57610cca611424565b60200101906001600160f81b031916908160001a90535080610ceb81611450565b915050610c99565b5092915050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610d3957610d39610cfa565b604052919050565b600082601f830112610d5257600080fd5b813567ffffffffffffffff811115610d6c57610d6c610cfa565b610d7f601f8201601f1916602001610d10565b818152846020838601011115610d9457600080fd5b816020850160208301376000918101602001919091529392505050565b80356001600160a01b0381168114610dc857600080fd5b919050565b80358015158114610dc857600080fd5b600067ffffffffffffffff821115610df757610df7610cfa565b5060051b60200190565b600082601f830112610e1257600080fd5b81356020610e27610e2283610ddd565b610d10565b82815260059290921b84018101918181019086841115610e4657600080fd5b8286015b84811015610e6857610e5b81610db1565b8352918301918301610e4a565b509695505050505050565b600082601f830112610e8457600080fd5b81356020610e94610e2283610ddd565b82815260059290921b84018101918181019086841115610eb357600080fd5b8286015b84811015610e6857803567ffffffffffffffff811115610ed75760008081fd5b610ee58986838b0101610d41565b845250918301918301610eb7565b60008060008060008060008060008060006101608c8e031215610f1557600080fd5b67ffffffffffffffff808d351115610f2c57600080fd5b610f398e8e358f01610d41565b9b50610f4760208e01610db1565b9a508060408e01351115610f5a57600080fd5b610f6a8e60408f01358f01610d41565b99508060608e01351115610f7d57600080fd5b610f8d8e60608f01358f01610d41565b98508060808e01351115610fa057600080fd5b610fb08e60808f01358f01610d41565b97508060a08e01351115610fc357600080fd5b610fd38e60a08f01358f01610d41565b965060c08d01359550610fe860e08e01610dcd565b94506101008d01359350806101208e0135111561100457600080fd5b6110158e6101208f01358f01610e01565b9250806101408e0135111561102957600080fd5b5061103b8d6101408e01358e01610e73565b90509295989b509295989b9093969950565b600080600080600060a0868803121561106557600080fd5b853567ffffffffffffffff8082111561107d57600080fd5b61108989838a01610d41565b965061109760208901610db1565b955060408801359150808211156110ad57600080fd5b6110b989838a01610d41565b945060608801359150808211156110cf57600080fd5b506110dc88828901610d41565b95989497509295608001359392505050565b60006020828403121561110057600080fd5b813567ffffffffffffffff81111561111757600080fd5b61112384828501610d41565b949350505050565b6000815180845260005b8181101561115157602081850181015186830182015201611135565b81811115611163576000602083870101525b50601f01601f19169290920160200192915050565b60208152600061118b602083018461112b565b9392505050565b600080604083850312156111a557600080fd5b823567ffffffffffffffff8111156111bc57600080fd5b6111c885828601610d41565b9250506111d760208401610db1565b90509250929050565b85815260a0602082015260006111f960a083018761112b565b828103604084015261120b818761112b565b94151560608401525050608001529392505050565b8c81526001600160a01b038c1660208201526101806040820181905260009061124b8382018e61112b565b9050828103606084015261125f818d61112b565b90508281036080840152611273818c61112b565b905082810360a0840152611287818b61112b565b905082810360c084015261129b818a61112b565b9150508660e08301526112b361010083018715159052565b84610120830152836101408301526112d061016083018415159052565b9d9c50505050505050505050505050565b600181811c908216806112f557607f821691505b602082108103610c4c57634e487b7160e01b600052602260045260246000fd5b601f82111561135f57600081815260208120601f850160051c8101602086101561133c5750805b601f850160051c820191505b8181101561135b57828155600101611348565b5050505b505050565b815167ffffffffffffffff81111561137e5761137e610cfa565b6113928161138c84546112e1565b84611315565b602080601f8311600181146113c757600084156113af5750858301515b600019600386901b1c1916600185901b17855561135b565b600085815260208120601f198616915b828110156113f6578886015182559484019460019091019084016113d7565b50858210156114145787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016114625761146261143a565b5060010190565b602080825260129082015271111bd8dd5b595b9d081b9bdd08195e1a5cdd60721b604082015260600190565b6000816114a4576114a461143a565b50600019019056fea264697066735822122058a10c3a54c905cb5f2ff70e7829f70ebe327c31b18337d3913dd5fdbe13eede64736f6c634300080f0033",
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
