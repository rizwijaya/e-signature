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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Profiles\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"profile_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"profile_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity_card\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateregistered\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_member\",\"type\":\"address\"}],\"name\":\"addPermission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_identity_card\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_dateregistered\",\"type\":\"string\"}],\"name\":\"add_profile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash_original\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signers\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_date\",\"type\":\"string\"}],\"name\":\"addsigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"documentlists\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_totalprofile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"profile\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_sign_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"sign\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"timestamp\",\"type\":\"string\"}],\"name\":\"profilesign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash_original\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_namefile\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash_file\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash_ipfs\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_totalsigned\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signing\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_createdtime\",\"type\":\"uint256\"}],\"name\":\"signDoc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sistem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userAccess\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"verify_document\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060008054336001600160a01b031991821681178355600180549092161781556002819055600381905560045561118590819061004d90396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063839aef5d11610071578063839aef5d146101805780638a4cd19d146101935780639d927d41146101a6578063afec53f3146101b9578063d740ba38146101cc578063ed1c62b6146101df57600080fd5b806321b84ef1146100b95780633ee141d2146100df57806345ffb23c146100e75780634d0a5dbd1461010e578063514b670b1461014257806372cd2b1a14610155575b600080fd5b6100cc6100c7366004610ab0565b610202565b6040519081526020015b60405180910390f35b6002546100cc565b6100fa6100f5366004610ae5565b610223565b6040516100d6989796959493929190610b54565b61014061011c366004610ae5565b6001600160a01b03166000908152600560205260409020805460ff19166001179055565b005b610140610150366004610c8b565b610519565b610168610163366004610ab0565b6105eb565b6040516001600160a01b0390911681526020016100d6565b600154610168906001600160a01b031681565b6101406101a1366004610ce2565b610615565b6101406101b4366004610db4565b6107bc565b600054610168906001600160a01b031681565b6101406101da366004610e17565b610848565b6101f26101ed366004610ab0565b610981565b60405190151581526020016100d6565b6008818154811061021257600080fd5b600091825260209091200154905081565b6007602052600090815260409020805460018201546002830180546001600160a01b0390931693919261025590610ee9565b80601f016020809104026020016040519081016040528092919081815260200182805461028190610ee9565b80156102ce5780601f106102a3576101008083540402835291602001916102ce565b820191906000526020600020905b8154815290600101906020018083116102b157829003601f168201915b5050505050908060030180546102e390610ee9565b80601f016020809104026020016040519081016040528092919081815260200182805461030f90610ee9565b801561035c5780601f106103315761010080835404028352916020019161035c565b820191906000526020600020905b81548152906001019060200180831161033f57829003601f168201915b50505050509080600401805461037190610ee9565b80601f016020809104026020016040519081016040528092919081815260200182805461039d90610ee9565b80156103ea5780601f106103bf576101008083540402835291602001916103ea565b820191906000526020600020905b8154815290600101906020018083116103cd57829003601f168201915b5050505050908060050180546103ff90610ee9565b80601f016020809104026020016040519081016040528092919081815260200182805461042b90610ee9565b80156104785780601f1061044d57610100808354040283529160200191610478565b820191906000526020600020905b81548152906001019060200180831161045b57829003601f168201915b50505050509080600601805461048d90610ee9565b80601f01602080910402602001604051908101604052809291908181526020018280546104b990610ee9565b80156105065780601f106104db57610100808354040283529160200191610506565b820191906000526020600020905b8154815290600101906020018083116104e957829003601f168201915b5050506007909301549192505060ff1688565b610522836109ec565b6105475760405162461bcd60e51b815260040161053e90610f23565b60405180910390fd5b604080516080810182526001600160a01b03808516808352600160208085018281528587018881526060870184905260008b815260098452888120958152600d909501909252959092208451815496511515600160a01b026001600160a81b03199097169416939093179490941782555191929091908201906105ca9082610fbb565b50606091909101516002909101805460ff1916911515919091179055505050565b600681815481106105fb57600080fd5b6000918252602090912001546001600160a01b0316905081565b61061e886109ec565b61063a5760405162461bcd60e51b815260040161053e90610f23565b60008881526009602052604090206003015461065890600190611091565b60000361067a5760008881526009602052604090206002600890910155610691565b600088815260096020526040902060016008909101555b6004805460008a81526009602052604090209081556002810180546001600160a01b03191633179055016106c58882610fbb565b506000888152600960205260409020600181018990556005016106e88782610fbb565b5060008881526009602052604090206006016107048682610fbb565b5060008881526009602052604090206007016107208582610fbb565b5061072c600183611091565b6000898152600960208190526040822060038101939093558201859055600a820183905542600b830155600c909101805460ff191660019081179091556008805480830182559083527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee3018a905560048054919290916107ad9084906110a8565b90915550505050505050505050565b604080516080810182528481526020808201858152828401859052600160608401819052336000908152600784528581208982526008019093529390912082518155905191929091908201906108129082610fbb565b50604082015160028201906108279082610fbb565b50606091909101516003909101805460ff1916911515919091179055505050565b6001546001600160a01b0316331461085f57600080fd5b33600081815260076020526040902080546001600160a01b0319169091178155600280546001830155016108938682610fbb565b503360009081526007602052604090206003016108b08482610fbb565b503360009081526007602052604090206004016108cd8382610fbb565b503360009081526007602052604090206005016108ea8582610fbb565b503360009081526007602052604090206006016109078282610fbb565b50336000818152600760208190526040822001805460ff1916600190811790915560068054918201815582527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0180546001600160a01b0319169092179091556002805491610975836110c0565b91905055505050505050565b604080516020808201835260009182905283825260099052818120915190917fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470916109cf91600501906110d9565b6040518091039020146109e457506001919050565b506000919050565b6000818152600960205260408120600c015460ff1615158103610a1157506000919050565b6000828152600960205260409020600301541580610a4057506000828152600960205260409020600801546002145b15610a4d57506000919050565b6000828152600960209081526040808320338452600d0190915290206002015460ff1615156001036109e4576000828152600960209081526040808320338452600d01909152812054600160a01b900460ff16151590036109e457506001919050565b600060208284031215610ac257600080fd5b5035919050565b80356001600160a01b0381168114610ae057600080fd5b919050565b600060208284031215610af757600080fd5b610b0082610ac9565b9392505050565b6000815180845260005b81811015610b2d57602081850181015186830182015201610b11565b81811115610b3f576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b03891681526020810188905261010060408201819052600090610b808382018a610b07565b90508281036060840152610b948189610b07565b90508281036080840152610ba88188610b07565b905082810360a0840152610bbc8187610b07565b905082810360c0840152610bd08186610b07565b91505082151560e08301529998505050505050505050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610c0f57600080fd5b813567ffffffffffffffff80821115610c2a57610c2a610be8565b604051601f8301601f19908116603f01168101908282118183101715610c5257610c52610be8565b81604052838152866020858801011115610c6b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080600060608486031215610ca057600080fd5b83359250610cb060208501610ac9565b9150604084013567ffffffffffffffff811115610ccc57600080fd5b610cd886828701610bfe565b9150509250925092565b600080600080600080600080610100898b031215610cff57600080fd5b88359750602089013567ffffffffffffffff80821115610d1e57600080fd5b610d2a8c838d01610bfe565b985060408b0135915080821115610d4057600080fd5b610d4c8c838d01610bfe565b975060608b0135915080821115610d6257600080fd5b610d6e8c838d01610bfe565b965060808b0135915080821115610d8457600080fd5b50610d918b828c01610bfe565b989b979a50959894979660a0860135965060c08601359560e00135945092505050565b600080600060608486031215610dc957600080fd5b83359250602084013567ffffffffffffffff80821115610de857600080fd5b610df487838801610bfe565b93506040860135915080821115610e0a57600080fd5b50610cd886828701610bfe565b600080600080600060a08688031215610e2f57600080fd5b853567ffffffffffffffff80821115610e4757600080fd5b610e5389838a01610bfe565b96506020880135915080821115610e6957600080fd5b610e7589838a01610bfe565b95506040880135915080821115610e8b57600080fd5b610e9789838a01610bfe565b94506060880135915080821115610ead57600080fd5b610eb989838a01610bfe565b93506080880135915080821115610ecf57600080fd5b50610edc88828901610bfe565b9150509295509295909350565b600181811c90821680610efd57607f821691505b602082108103610f1d57634e487b7160e01b600052602260045260246000fd5b50919050565b60208082526029908201527f596f7520617265206e6f7420616c6c6f77656420746f207369676e207468697360408201526808191bd8dd5b595b9d60ba1b606082015260800190565b601f821115610fb657600081815260208120601f850160051c81016020861015610f935750805b601f850160051c820191505b81811015610fb257828155600101610f9f565b5050505b505050565b815167ffffffffffffffff811115610fd557610fd5610be8565b610fe981610fe38454610ee9565b84610f6c565b602080601f83116001811461101e57600084156110065750858301515b600019600386901b1c1916600185901b178555610fb2565b600085815260208120601f198616915b8281101561104d5788860151825594840194600190910190840161102e565b508582101561106b5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b6000828210156110a3576110a361107b565b500390565b600082198211156110bb576110bb61107b565b500190565b6000600182016110d2576110d261107b565b5060010190565b60008083546110e781610ee9565b600182811680156110ff576001811461111457611143565b60ff1984168752821515830287019450611143565b8760005260208060002060005b8581101561113a5781548a820152908401908201611121565b50505082870194505b5092969550505050505056fea264697066735822122047668af8e5577e3175b9a00e784b870d332d16b1e72bb5ea05b461207190f2a664736f6c634300080f0033",
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

// Profiles is a free data retrieval call binding the contract method 0x45ffb23c.
//
// Solidity: function Profiles(address ) view returns(address profile_address, uint256 profile_id, string name, string email, string phone, string identity_card, string dateregistered, bool exist)
func (_Api *ApiCaller) Profiles(opts *bind.CallOpts, arg0 common.Address) (struct {
	ProfileAddress common.Address
	ProfileId      *big.Int
	Name           string
	Email          string
	Phone          string
	IdentityCard   string
	Dateregistered string
	Exist          bool
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "Profiles", arg0)

	outstruct := new(struct {
		ProfileAddress common.Address
		ProfileId      *big.Int
		Name           string
		Email          string
		Phone          string
		IdentityCard   string
		Dateregistered string
		Exist          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProfileAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ProfileId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Email = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Phone = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.IdentityCard = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Dateregistered = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Exist = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Profiles is a free data retrieval call binding the contract method 0x45ffb23c.
//
// Solidity: function Profiles(address ) view returns(address profile_address, uint256 profile_id, string name, string email, string phone, string identity_card, string dateregistered, bool exist)
func (_Api *ApiSession) Profiles(arg0 common.Address) (struct {
	ProfileAddress common.Address
	ProfileId      *big.Int
	Name           string
	Email          string
	Phone          string
	IdentityCard   string
	Dateregistered string
	Exist          bool
}, error) {
	return _Api.Contract.Profiles(&_Api.CallOpts, arg0)
}

// Profiles is a free data retrieval call binding the contract method 0x45ffb23c.
//
// Solidity: function Profiles(address ) view returns(address profile_address, uint256 profile_id, string name, string email, string phone, string identity_card, string dateregistered, bool exist)
func (_Api *ApiCallerSession) Profiles(arg0 common.Address) (struct {
	ProfileAddress common.Address
	ProfileId      *big.Int
	Name           string
	Email          string
	Phone          string
	IdentityCard   string
	Dateregistered string
	Exist          bool
}, error) {
	return _Api.Contract.Profiles(&_Api.CallOpts, arg0)
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

// GetTotalprofile is a free data retrieval call binding the contract method 0x3ee141d2.
//
// Solidity: function get_totalprofile() view returns(uint256)
func (_Api *ApiCaller) GetTotalprofile(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_totalprofile")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalprofile is a free data retrieval call binding the contract method 0x3ee141d2.
//
// Solidity: function get_totalprofile() view returns(uint256)
func (_Api *ApiSession) GetTotalprofile() (*big.Int, error) {
	return _Api.Contract.GetTotalprofile(&_Api.CallOpts)
}

// GetTotalprofile is a free data retrieval call binding the contract method 0x3ee141d2.
//
// Solidity: function get_totalprofile() view returns(uint256)
func (_Api *ApiCallerSession) GetTotalprofile() (*big.Int, error) {
	return _Api.Contract.GetTotalprofile(&_Api.CallOpts)
}

// Profile is a free data retrieval call binding the contract method 0x72cd2b1a.
//
// Solidity: function profile(uint256 ) view returns(address)
func (_Api *ApiCaller) Profile(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "profile", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Profile is a free data retrieval call binding the contract method 0x72cd2b1a.
//
// Solidity: function profile(uint256 ) view returns(address)
func (_Api *ApiSession) Profile(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.Profile(&_Api.CallOpts, arg0)
}

// Profile is a free data retrieval call binding the contract method 0x72cd2b1a.
//
// Solidity: function profile(uint256 ) view returns(address)
func (_Api *ApiCallerSession) Profile(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.Profile(&_Api.CallOpts, arg0)
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

// AddProfile is a paid mutator transaction binding the contract method 0xd740ba38.
//
// Solidity: function add_profile(string _name, string _identity_card, string _email, string _phone, string _dateregistered) returns()
func (_Api *ApiTransactor) AddProfile(opts *bind.TransactOpts, _name string, _identity_card string, _email string, _phone string, _dateregistered string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "add_profile", _name, _identity_card, _email, _phone, _dateregistered)
}

// AddProfile is a paid mutator transaction binding the contract method 0xd740ba38.
//
// Solidity: function add_profile(string _name, string _identity_card, string _email, string _phone, string _dateregistered) returns()
func (_Api *ApiSession) AddProfile(_name string, _identity_card string, _email string, _phone string, _dateregistered string) (*types.Transaction, error) {
	return _Api.Contract.AddProfile(&_Api.TransactOpts, _name, _identity_card, _email, _phone, _dateregistered)
}

// AddProfile is a paid mutator transaction binding the contract method 0xd740ba38.
//
// Solidity: function add_profile(string _name, string _identity_card, string _email, string _phone, string _dateregistered) returns()
func (_Api *ApiTransactorSession) AddProfile(_name string, _identity_card string, _email string, _phone string, _dateregistered string) (*types.Transaction, error) {
	return _Api.Contract.AddProfile(&_Api.TransactOpts, _name, _identity_card, _email, _phone, _dateregistered)
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

// Profilesign is a paid mutator transaction binding the contract method 0x9d927d41.
//
// Solidity: function profilesign(uint256 _sign_id, string sign, string timestamp) returns()
func (_Api *ApiTransactor) Profilesign(opts *bind.TransactOpts, _sign_id *big.Int, sign string, timestamp string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "profilesign", _sign_id, sign, timestamp)
}

// Profilesign is a paid mutator transaction binding the contract method 0x9d927d41.
//
// Solidity: function profilesign(uint256 _sign_id, string sign, string timestamp) returns()
func (_Api *ApiSession) Profilesign(_sign_id *big.Int, sign string, timestamp string) (*types.Transaction, error) {
	return _Api.Contract.Profilesign(&_Api.TransactOpts, _sign_id, sign, timestamp)
}

// Profilesign is a paid mutator transaction binding the contract method 0x9d927d41.
//
// Solidity: function profilesign(uint256 _sign_id, string sign, string timestamp) returns()
func (_Api *ApiTransactorSession) Profilesign(_sign_id *big.Int, sign string, timestamp string) (*types.Transaction, error) {
	return _Api.Contract.Profilesign(&_Api.TransactOpts, _sign_id, sign, timestamp)
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
