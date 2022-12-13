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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_creator_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_mode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"_signers_id\",\"type\":\"string[]\"}],\"name\":\"create\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"}],\"name\":\"getDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"getDocSigned\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"}],\"name\":\"getListSign\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_signers_id\",\"type\":\"address\"}],\"name\":\"getSign\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_signers_id\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_signers_hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_ipfs\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"signDoc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_hash\",\"type\":\"string\"}],\"name\":\"verifyDoc\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506001600055611679806100256000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063492bdf491161005b578063492bdf49146100e4578063a5bde23b14610104578063c73358041461012f578063e1ccea401461015257600080fd5b80631af2925c146100825780632887590e146100975780633d995877146100c0575b600080fd5b610095610090366004610f36565b610165565b005b6100aa6100a5366004610fd7565b61032f565b6040516100b79190611061565b60405180910390f35b6100d36100ce36600461107b565b610360565b6040516100b79594939291906110c9565b6100f76100f2366004610fd7565b61057a565b6040516100b79190611109565b610117610112366004610fd7565b610629565b6040516100b79c9b9a99989796959493929190611156565b61014261013d366004610fd7565b6109be565b60405190151581526020016100b7565b610095610160366004611325565b6109f8565b600061017086610d41565b6000818152600160208190526040909120600c8101549293509160ff161515146101b55760405162461bcd60e51b81526004016101ac90611478565b60405180910390fd5b6001600160a01b03861660009081526010820160205260409020600501546001106101f25760405162461bcd60e51b81526004016101ac90611478565b6001600160a01b038616600090815260108201602052604090206004015460ff16156102605760405162461bcd60e51b815260206004820152601c60248201527f596f7520617265207369676e6564207468697320646f63756d656e740000000060448201526064016101ac565b600d81018054906000610272836114ba565b9091555050600781016102858582611554565b50600b81018390556006810161029b8682611554565b506001600160a01b038616600090815260108201602052604090206003016102c38682611554565b506001600160a01b0386166000908152601082016020526040812060048101805460ff191660011790556005018490556102fc86610d41565b8254600082815260026020526040902055600d8301549091506001111561032557600260088301555b5050505050505050565b6060600061033c83610d41565b60008181526002602052604081205491925061035782610d60565b95945050505050565b6000606080600080600061037388610d41565b6000818152600160208190526040909120600c8101549293509160ff161515146103af5760405162461bcd60e51b81526004016101ac90611478565b6001600160a01b0380891660008181526010840160205260409020549091161461040f5760405162461bcd60e51b815260206004820152601160248201527014da59db995c9cc81b9bdd08195e1a5cdd607a1b60448201526064016101ac565b6001600160a01b038816600090815260108201602052604090206001810154600482015460058301546002840180549394909360039091019260ff1691908490610458906114d1565b80601f0160208091040260200160405190810160405280929190818152602001828054610484906114d1565b80156104d15780601f106104a6576101008083540402835291602001916104d1565b820191906000526020600020905b8154815290600101906020018083116104b457829003601f168201915b505050505093508280546104e4906114d1565b80601f0160208091040260200160405190810160405280929190818152602001828054610510906114d1565b801561055d5780601f106105325761010080835404028352916020019161055d565b820191906000526020600020905b81548152906001019060200180831161054057829003601f168201915b505050505092509650965096509650965050509295509295909350565b6060600061058783610d41565b6000818152600160208190526040909120600c8101549293509160ff161515146105c35760405162461bcd60e51b81526004016101ac90611478565b80600f0180548060200260200160405190810160405280929190818152602001828054801561061b57602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116105fd575b505050505092505050919050565b60008060608060608060606000806000806000806106468e610d41565b6000818152600160208190526040909120600c8101549293509160ff161515146106825760405162461bcd60e51b81526004016101ac90611478565b80600101548160030160009054906101000a90046001600160a01b031682600201836004018460050185600601866007018760080154886009015489600a01548a600b01548b600c0160009054906101000a900460ff168980546106e5906114d1565b80601f0160208091040260200160405190810160405280929190818152602001828054610711906114d1565b801561075e5780601f106107335761010080835404028352916020019161075e565b820191906000526020600020905b81548152906001019060200180831161074157829003601f168201915b50505050509950888054610771906114d1565b80601f016020809104026020016040519081016040528092919081815260200182805461079d906114d1565b80156107ea5780601f106107bf576101008083540402835291602001916107ea565b820191906000526020600020905b8154815290600101906020018083116107cd57829003601f168201915b505050505098508780546107fd906114d1565b80601f0160208091040260200160405190810160405280929190818152602001828054610829906114d1565b80156108765780601f1061084b57610100808354040283529160200191610876565b820191906000526020600020905b81548152906001019060200180831161085957829003601f168201915b50505050509750868054610889906114d1565b80601f01602080910402602001604051908101604052809291908181526020018280546108b5906114d1565b80156109025780601f106108d757610100808354040283529160200191610902565b820191906000526020600020905b8154815290600101906020018083116108e557829003601f168201915b50505050509650858054610915906114d1565b80601f0160208091040260200160405190810160405280929190818152602001828054610941906114d1565b801561098e5780601f106109635761010080835404028352916020019161098e565b820191906000526020600020905b81548152906001019060200180831161097157829003601f168201915b505050505095509d509d509d509d509d509d509d509d509d509d509d509d50505091939597999b5091939597999b565b6000806109ca83610d41565b6000818152600260205260409020549091506109e95750600092915050565b50600192915050565b50919050565b6000610a038c610d41565b600081815260016020819052604082209154908201558181556003810180546001600160a01b0319166001600160a01b038f1617905590915060028101610a4a8c82611554565b5060048101610a598b82611554565b5060058101610a688e82611554565b5060068101610a778a82611554565b5060078101610a868982611554565b506008810187905560098101869055600a8101859055600b8101859055600c8101805460ff191660011790558351600d8201819055600e8201819055610ad590600f8301906020870190610de9565b5060005b8451811015610d1d57848181518110610af457610af4611614565b6020026020010151826010016000878481518110610b1457610b14611614565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060000160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080826010016000878481518110610b7c57610b7c611614565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060010181905550838181518110610bbd57610bbd611614565b6020026020010151826010016000878481518110610bdd57610bdd611614565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206002019081610c149190611554565b5089826010016000878481518110610c2e57610c2e611614565b60200260200101516001600160a01b03166001600160a01b031681526020019081526020016000206003019081610c659190611554565b506000826010016000878481518110610c8057610c80611614565b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060040160006101000a81548160ff02191690831515021790555085826010016000878481518110610cdb57610cdb611614565b60200260200101516001600160a01b03166001600160a01b03168152602001908152602001600020600501819055508080610d159061162a565b915050610ad9565b50600080549080610d2d8361162a565b919050555050505050505050505050505050565b805160009082908203610d575750600092915050565b50506020015190565b6040805160208082528183019092526060916000919060208201818036833701905050905060005b6020811015610de257838160208110610da357610da3611614565b1a60f81b828281518110610db957610db9611614565b60200101906001600160f81b031916908160001a90535080610dda8161162a565b915050610d88565b5092915050565b828054828255906000526020600020908101928215610e3e579160200282015b82811115610e3e57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190610e09565b50610e4a929150610e4e565b5090565b5b80821115610e4a5760008155600101610e4f565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff81118282101715610ea257610ea2610e63565b604052919050565b600082601f830112610ebb57600080fd5b813567ffffffffffffffff811115610ed557610ed5610e63565b610ee8601f8201601f1916602001610e79565b818152846020838601011115610efd57600080fd5b816020850160208301376000918101602001919091529392505050565b80356001600160a01b0381168114610f3157600080fd5b919050565b600080600080600060a08688031215610f4e57600080fd5b853567ffffffffffffffff80821115610f6657600080fd5b610f7289838a01610eaa565b9650610f8060208901610f1a565b95506040880135915080821115610f9657600080fd5b610fa289838a01610eaa565b94506060880135915080821115610fb857600080fd5b50610fc588828901610eaa565b95989497509295608001359392505050565b600060208284031215610fe957600080fd5b813567ffffffffffffffff81111561100057600080fd5b61100c84828501610eaa565b949350505050565b6000815180845260005b8181101561103a5760208185018101518683018201520161101e565b8181111561104c576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006110746020830184611014565b9392505050565b6000806040838503121561108e57600080fd5b823567ffffffffffffffff8111156110a557600080fd5b6110b185828601610eaa565b9250506110c060208401610f1a565b90509250929050565b85815260a0602082015260006110e260a0830187611014565b82810360408401526110f48187611014565b94151560608401525050608001529392505050565b6020808252825182820181905260009190848201906040850190845b8181101561114a5783516001600160a01b031683529284019291840191600101611125565b50909695505050505050565b8c81526001600160a01b038c166020820152610180604082018190526000906111818382018e611014565b90508281036060840152611195818d611014565b905082810360808401526111a9818c611014565b905082810360a08401526111bd818b611014565b905082810360c08401526111d1818a611014565b9150508660e08301528561010083015284610120830152836101408301526111fe61016083018415159052565b9d9c50505050505050505050505050565b600067ffffffffffffffff82111561122957611229610e63565b5060051b60200190565b600082601f83011261124457600080fd5b813560206112596112548361120f565b610e79565b82815260059290921b8401810191818101908684111561127857600080fd5b8286015b8481101561129a5761128d81610f1a565b835291830191830161127c565b509695505050505050565b600082601f8301126112b657600080fd5b813560206112c66112548361120f565b82815260059290921b840181019181810190868411156112e557600080fd5b8286015b8481101561129a57803567ffffffffffffffff8111156113095760008081fd5b6113178986838b0101610eaa565b8452509183019183016112e9565b60008060008060008060008060008060006101608c8e03121561134757600080fd5b67ffffffffffffffff808d35111561135e57600080fd5b61136b8e8e358f01610eaa565b9b5061137960208e01610f1a565b9a508060408e0135111561138c57600080fd5b61139c8e60408f01358f01610eaa565b99508060608e013511156113af57600080fd5b6113bf8e60608f01358f01610eaa565b98508060808e013511156113d257600080fd5b6113e28e60808f01358f01610eaa565b97508060a08e013511156113f557600080fd5b6114058e60a08f01358f01610eaa565b965060c08d0135955060e08d013594506101008d01359350806101208e0135111561142f57600080fd5b6114408e6101208f01358f01611233565b9250806101408e0135111561145457600080fd5b506114668d6101408e01358e016112a5565b90509295989b509295989b9093969950565b602080825260129082015271111bd8dd5b595b9d081b9bdd08195e1a5cdd60721b604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6000816114c9576114c96114a4565b506000190190565b600181811c908216806114e557607f821691505b6020821081036109f257634e487b7160e01b600052602260045260246000fd5b601f82111561154f57600081815260208120601f850160051c8101602086101561152c5750805b601f850160051c820191505b8181101561154b57828155600101611538565b5050505b505050565b815167ffffffffffffffff81111561156e5761156e610e63565b6115828161157c84546114d1565b84611505565b602080601f8311600181146115b7576000841561159f5750858301515b600019600386901b1c1916600185901b17855561154b565b600085815260208120601f198616915b828110156115e6578886015182559484019460019091019084016115c7565b50858210156116045787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052603260045260246000fd5b60006001820161163c5761163c6114a4565b506001019056fea2646970667358221220dea0e5523f3dd661ad512bd55b1afeee1b58b3f22ca227771cf53838fa29bd7964736f6c634300080f0033",
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

// GetListSign is a free data retrieval call binding the contract method 0x492bdf49.
//
// Solidity: function getListSign(string _file) view returns(address[])
func (_Api *ApiCaller) GetListSign(opts *bind.CallOpts, _file string) ([]common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getListSign", _file)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetListSign is a free data retrieval call binding the contract method 0x492bdf49.
//
// Solidity: function getListSign(string _file) view returns(address[])
func (_Api *ApiSession) GetListSign(_file string) ([]common.Address, error) {
	return _Api.Contract.GetListSign(&_Api.CallOpts, _file)
}

// GetListSign is a free data retrieval call binding the contract method 0x492bdf49.
//
// Solidity: function getListSign(string _file) view returns(address[])
func (_Api *ApiCallerSession) GetListSign(_file string) ([]common.Address, error) {
	return _Api.Contract.GetListSign(&_Api.CallOpts, _file)
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
