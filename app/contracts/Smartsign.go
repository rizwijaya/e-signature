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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Profiles_private\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"profile_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"profile_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"idsignature\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity_card\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dateregistered\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_profile_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_idsignature\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_identity_card\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_phone\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_dateregistered\",\"type\":\"string\"}],\"name\":\"add_profile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_idsignature\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_password\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_publickey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_role\",\"type\":\"string\"}],\"name\":\"add_profilefirst\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_profile_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_signature\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_datecreated\",\"type\":\"string\"}],\"name\":\"add_signature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash_original\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_signers\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_signer_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"name\":\"addsigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_totalprofile\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"get_totalsignature\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"profile_private\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"profiles_public\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"profile_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"idsignature\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"publickey\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash_original\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_namefile\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_hash_file\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_metadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_file\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_status\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_totalsigned\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_datesigned\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"signing\",\"type\":\"uint256\"}],\"name\":\"signDoc\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signatures\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"signature_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"profile_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"signature\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"datecreated\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sistem\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userAccess\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_hash\",\"type\":\"bytes32\"}],\"name\":\"verify_document\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060008054336001600160a01b0319918216811783556001805490921617815560028190556003819055600455611bd290819061004d90396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c8063aba0bc251161008c578063c05e100111610066578063c05e1001146101e4578063d15de67f146101f7578063ed1c62b61461020a578063ee4b597d1461022d57600080fd5b8063aba0bc2514610197578063afec53f3146101be578063b5fa9cf9146101d157600080fd5b80636ec0610d116100c85780636ec0610d14610123578063839aef5d146101365780638be1019414610161578063aa43577c1461018457600080fd5b80631c87468e146100ef57806337f2678d146101045780633ee141d21461011b575b600080fd5b6101026100fd366004611377565b610251565b005b6003545b6040519081526020015b60405180910390f35b600254610108565b610102610131366004611424565b61031c565b600154610149906001600160a01b031681565b6040516001600160a01b039091168152602001610112565b61017461016f366004611538565b610575565b604051610112949392919061159e565b610149610192366004611538565b610736565b6101aa6101a5366004611604565b610760565b604051610112989796959493929190611626565b600054610149906001600160a01b031681565b6101026101df3660046116d2565b610b63565b6101026101f236600461170d565b610c6c565b610102610205366004611795565b610d1a565b61021d610218366004611538565b610e8f565b6040519015158152602001610112565b61024061023b366004611538565b610efa565b6040516101129594939291906118c1565b6000546001600160a01b0316331461026857600080fd5b6040805160a08101825260025480825260208083018881528385018890526060840187905260808401869052600092835260059091529290208151815591519091829160018201906102ba90826119a9565b50604082015160028201906102cf90826119a9565b50606082015160038201906102e490826119a9565b50608082015160048201906102f990826119a9565b509050506001600260008282546103109190611a7f565b90915550505050505050565b61032589611149565b6103765760405162461bcd60e51b815260206004820152601e60248201527f446f63756d656e74206e6f7420666f756e642f5369676e696e6720616c6c000060448201526064015b60405180910390fd5b61038089846111d4565b61039c5760405162461bcd60e51b815260040161036d90611a97565b6103a58961126b565b156103cc576103b5600182611ae0565b60008a8152600960205260409020600201556103e1565b60008981526009602052604090206002018390555b6000898152600960205260409020600201546103ff90600190611ae0565b60000361042757600089815260096020526040902060080161042185826119a9565b50610465565b604051806040016040528060018152602001600d60fa1b815250600960008b8152602001908152602001600020600801908161046391906119a9565b505b60045460008a81526009602052604090209081556001810180546001600160a01b0319163317905560030161049a89826119a9565b506000898152600960205260409020600481018a90556005016104bd88826119a9565b5060008981526009602052604090206006016104d987826119a9565b5060008981526009602052604090206007016104f586826119a9565b506000898152600960208190526040909120908101849055600a0161051a83826119a9565b50600a80546001818101835560009283527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a89091018b90556004805491929091610565908490611a7f565b9091555050505050505050505050565b6008602052600090815260409020805460018201805491929161059790611920565b80601f01602080910402602001604051908101604052809291908181526020018280546105c390611920565b80156106105780601f106105e557610100808354040283529160200191610610565b820191906000526020600020905b8154815290600101906020018083116105f357829003601f168201915b50505050509080600201805461062590611920565b80601f016020809104026020016040519081016040528092919081815260200182805461065190611920565b801561069e5780601f106106735761010080835404028352916020019161069e565b820191906000526020600020905b81548152906001019060200180831161068157829003601f168201915b5050505050908060030180546106b390611920565b80601f01602080910402602001604051908101604052809291908181526020018280546106df90611920565b801561072c5780601f106107015761010080835404028352916020019161072c565b820191906000526020600020905b81548152906001019060200180831161070f57829003601f168201915b5050505050905084565b6006818154811061074657600080fd5b6000918252602090912001546001600160a01b0316905081565b600760205260009081526040902080546001820180546001600160a01b03909216929161078c90611920565b80601f01602080910402602001604051908101604052809291908181526020018280546107b890611920565b80156108055780601f106107da57610100808354040283529160200191610805565b820191906000526020600020905b8154815290600101906020018083116107e857829003601f168201915b50505050509080600201805461081a90611920565b80601f016020809104026020016040519081016040528092919081815260200182805461084690611920565b80156108935780601f1061086857610100808354040283529160200191610893565b820191906000526020600020905b81548152906001019060200180831161087657829003601f168201915b5050505050908060030180546108a890611920565b80601f01602080910402602001604051908101604052809291908181526020018280546108d490611920565b80156109215780601f106108f657610100808354040283529160200191610921565b820191906000526020600020905b81548152906001019060200180831161090457829003601f168201915b50505050509080600401805461093690611920565b80601f016020809104026020016040519081016040528092919081815260200182805461096290611920565b80156109af5780601f10610984576101008083540402835291602001916109af565b820191906000526020600020905b81548152906001019060200180831161099257829003601f168201915b5050505050908060050180546109c490611920565b80601f01602080910402602001604051908101604052809291908181526020018280546109f090611920565b8015610a3d5780601f10610a1257610100808354040283529160200191610a3d565b820191906000526020600020905b815481529060010190602001808311610a2057829003601f168201915b505050505090806006018054610a5290611920565b80601f0160208091040260200160405190810160405280929190818152602001828054610a7e90611920565b8015610acb5780601f10610aa057610100808354040283529160200191610acb565b820191906000526020600020905b815481529060010190602001808311610aae57829003601f168201915b505050505090806007018054610ae090611920565b80601f0160208091040260200160405190810160405280929190818152602001828054610b0c90611920565b8015610b595780601f10610b2e57610100808354040283529160200191610b59565b820191906000526020600020905b815481529060010190602001808311610b3c57829003601f168201915b5050505050905088565b610b6c84611149565b610bb85760405162461bcd60e51b815260206004820152601e60248201527f446f63756d656e74206e6f7420666f756e642f5369676e696e6720416c6c0000604482015260640161036d565b610bc284826111d4565b610bde5760405162461bcd60e51b815260040161036d90611a97565b6000848152600960205260409020600b01805483919082908110610c0457610c04611af7565b6000918252602080832091909101929092556040805180820182526001600160a01b03968716815280840186815297835260098452818320958352600c90950190925220915182546001600160a01b0319169316929092178155915160019092019190915550565b6001546001600160a01b03163314610c8357600080fd5b6040805160808101825260035480825260208083018781528385018790526060840186905260009283526008909152929020815181559151909182916001820190610cce90826119a9565b5060408201516002820190610ce390826119a9565b5060608201516003820190610cf890826119a9565b50905050600160036000828254610d0f9190611a7f565b909155505050505050565b6000546001600160a01b03163314610d3157600080fd5b60408051610100810182526001600160a01b038a811680835260208084018c81528486018c9052606085018b9052608085018a905260a0850189905260c0850188905260e0850187905260009283526007909152939020825181546001600160a01b03191692169190911781559151909182916001820190610db390826119a9565b5060408201516002820190610dc890826119a9565b5060608201516003820190610ddd90826119a9565b5060808201516004820190610df290826119a9565b5060a08201516005820190610e0790826119a9565b5060c08201516006820190610e1c90826119a9565b5060e08201516007820190610e3190826119a9565b5050600680546001810182556000919091527ff652222313e28459528d920b65115c16c04f3efc82aaedc97be59f3f377c0d3f0180546001600160a01b0319166001600160a01b039b909b169a909a17909955505050505050505050565b604080516020808201835260009182905283825260099052818120915190917fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47091610edd9160070190611b0d565b604051809103902014610ef257506001919050565b506000919050565b60056020526000908152604090208054600182018054919291610f1c90611920565b80601f0160208091040260200160405190810160405280929190818152602001828054610f4890611920565b8015610f955780601f10610f6a57610100808354040283529160200191610f95565b820191906000526020600020905b815481529060010190602001808311610f7857829003601f168201915b505050505090806002018054610faa90611920565b80601f0160208091040260200160405190810160405280929190818152602001828054610fd690611920565b80156110235780601f10610ff857610100808354040283529160200191611023565b820191906000526020600020905b81548152906001019060200180831161100657829003601f168201915b50505050509080600301805461103890611920565b80601f016020809104026020016040519081016040528092919081815260200182805461106490611920565b80156110b15780601f10611086576101008083540402835291602001916110b1565b820191906000526020600020905b81548152906001019060200180831161109457829003601f168201915b5050505050908060040180546110c690611920565b80601f01602080910402602001604051908101604052809291908181526020018280546110f290611920565b801561113f5780601f106111145761010080835404028352916020019161113f565b820191906000526020600020905b81548152906001019060200180831161112257829003601f168201915b5050505050905085565b604080516020808201835260009182905283825260099052818120915190917fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470916111979160080190611b0d565b6040518091039020146111ac57506000919050565b60008281526009602052604081206002015490036111cc57506000919050565b506001919050565b60006111df8361126b565b1561123f5760005b82811015611239576000848152600960209081526040808320848452600c01909152902054336001600160a01b0390911603611227576001915050611265565b8061123181611b83565b9150506111e7565b50611261565b6001546001600160a01b0316330361125957506001611265565b506000611265565b5060005b92915050565b60408051808201825260018152603160f81b60209182015260008381526009909152818120915190917fc89efdaa54c0f20c7adf612882df0950f5a951637e0307cdcb4c672f298b8bc6916112c39160080190611b0d565b604051809103902014159050919050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126112fb57600080fd5b813567ffffffffffffffff80821115611316576113166112d4565b604051601f8301601f19908116603f0116810190828211818310171561133e5761133e6112d4565b8160405283815286602085880101111561135757600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000806080858703121561138d57600080fd5b843567ffffffffffffffff808211156113a557600080fd5b6113b1888389016112ea565b955060208701359150808211156113c757600080fd5b6113d3888389016112ea565b945060408701359150808211156113e957600080fd5b6113f5888389016112ea565b9350606087013591508082111561140b57600080fd5b50611418878288016112ea565b91505092959194509250565b60008060008060008060008060006101208a8c03121561144357600080fd5b8935985060208a013567ffffffffffffffff8082111561146257600080fd5b61146e8d838e016112ea565b995060408c013591508082111561148457600080fd5b6114908d838e016112ea565b985060608c01359150808211156114a657600080fd5b6114b28d838e016112ea565b975060808c01359150808211156114c857600080fd5b6114d48d838e016112ea565b965060a08c01359150808211156114ea57600080fd5b6114f68d838e016112ea565b955060c08c0135945060e08c013591508082111561151357600080fd5b506115208c828d016112ea565b9250506101008a013590509295985092959850929598565b60006020828403121561154a57600080fd5b5035919050565b6000815180845260005b818110156115775760208185018101518683018201520161155b565b81811115611589576000602083870101525b50601f01601f19169290920160200192915050565b8481526080602082015260006115b76080830186611551565b82810360408401526115c98186611551565b905082810360608401526115dd8185611551565b979650505050505050565b80356001600160a01b03811681146115ff57600080fd5b919050565b60006020828403121561161657600080fd5b61161f826115e8565b9392505050565b6001600160a01b03891681526101006020820181905260009061164b8382018b611551565b9050828103604084015261165f818a611551565b905082810360608401526116738189611551565b905082810360808401526116878188611551565b905082810360a084015261169b8187611551565b905082810360c08401526116af8186611551565b905082810360e08401526116c38185611551565b9b9a5050505050505050505050565b600080600080608085870312156116e857600080fd5b843593506116f8602086016115e8565b93969395505050506040820135916060013590565b60008060006060848603121561172257600080fd5b833567ffffffffffffffff8082111561173a57600080fd5b611746878388016112ea565b9450602086013591508082111561175c57600080fd5b611768878388016112ea565b9350604086013591508082111561177e57600080fd5b5061178b868287016112ea565b9150509250925092565b600080600080600080600080610100898b0312156117b257600080fd5b6117bb896115e8565b9750602089013567ffffffffffffffff808211156117d857600080fd5b6117e48c838d016112ea565b985060408b01359150808211156117fa57600080fd5b6118068c838d016112ea565b975060608b013591508082111561181c57600080fd5b6118288c838d016112ea565b965060808b013591508082111561183e57600080fd5b61184a8c838d016112ea565b955060a08b013591508082111561186057600080fd5b61186c8c838d016112ea565b945060c08b013591508082111561188257600080fd5b61188e8c838d016112ea565b935060e08b01359150808211156118a457600080fd5b506118b18b828c016112ea565b9150509295985092959890939650565b85815260a0602082015260006118da60a0830187611551565b82810360408401526118ec8187611551565b905082810360608401526119008186611551565b905082810360808401526119148185611551565b98975050505050505050565b600181811c9082168061193457607f821691505b60208210810361195457634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156119a457600081815260208120601f850160051c810160208610156119815750805b601f850160051c820191505b818110156119a05782815560010161198d565b5050505b505050565b815167ffffffffffffffff8111156119c3576119c36112d4565b6119d7816119d18454611920565b8461195a565b602080601f831160018114611a0c57600084156119f45750858301515b600019600386901b1c1916600185901b1785556119a0565b600085815260208120601f198616915b82811015611a3b57888601518255948401946001909101908401611a1c565b5085821015611a595787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b634e487b7160e01b600052601160045260246000fd5b60008219821115611a9257611a92611a69565b500190565b60208082526029908201527f596f7520617265206e6f7420616c6c6f77656420746f207369676e207468697360408201526808191bd8dd5b595b9d60ba1b606082015260800190565b600082821015611af257611af2611a69565b500390565b634e487b7160e01b600052603260045260246000fd5b6000808354611b1b81611920565b60018281168015611b335760018114611b4857611b77565b60ff1984168752821515830287019450611b77565b8760005260208060002060005b85811015611b6e5781548a820152908401908201611b55565b50505082870194505b50929695505050505050565b600060018201611b9557611b95611a69565b506001019056fea26469706673582212204af5a24265551b82eaf24e98a1cc0e8ac07d35b36d1fd28b5a258207232c3a2964736f6c634300080f0033",
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

// ProfilesPrivate is a free data retrieval call binding the contract method 0xaba0bc25.
//
// Solidity: function Profiles_private(address ) view returns(address profile_address, string profile_id, string idsignature, string name, string identity_card, string email, string phone, string dateregistered)
func (_Api *ApiCaller) ProfilesPrivate(opts *bind.CallOpts, arg0 common.Address) (struct {
	ProfileAddress common.Address
	ProfileId      string
	Idsignature    string
	Name           string
	IdentityCard   string
	Email          string
	Phone          string
	Dateregistered string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "Profiles_private", arg0)

	outstruct := new(struct {
		ProfileAddress common.Address
		ProfileId      string
		Idsignature    string
		Name           string
		IdentityCard   string
		Email          string
		Phone          string
		Dateregistered string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProfileAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ProfileId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Idsignature = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Name = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.IdentityCard = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Email = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Phone = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Dateregistered = *abi.ConvertType(out[7], new(string)).(*string)

	return *outstruct, err

}

// ProfilesPrivate is a free data retrieval call binding the contract method 0xaba0bc25.
//
// Solidity: function Profiles_private(address ) view returns(address profile_address, string profile_id, string idsignature, string name, string identity_card, string email, string phone, string dateregistered)
func (_Api *ApiSession) ProfilesPrivate(arg0 common.Address) (struct {
	ProfileAddress common.Address
	ProfileId      string
	Idsignature    string
	Name           string
	IdentityCard   string
	Email          string
	Phone          string
	Dateregistered string
}, error) {
	return _Api.Contract.ProfilesPrivate(&_Api.CallOpts, arg0)
}

// ProfilesPrivate is a free data retrieval call binding the contract method 0xaba0bc25.
//
// Solidity: function Profiles_private(address ) view returns(address profile_address, string profile_id, string idsignature, string name, string identity_card, string email, string phone, string dateregistered)
func (_Api *ApiCallerSession) ProfilesPrivate(arg0 common.Address) (struct {
	ProfileAddress common.Address
	ProfileId      string
	Idsignature    string
	Name           string
	IdentityCard   string
	Email          string
	Phone          string
	Dateregistered string
}, error) {
	return _Api.Contract.ProfilesPrivate(&_Api.CallOpts, arg0)
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

// GetTotalsignature is a free data retrieval call binding the contract method 0x37f2678d.
//
// Solidity: function get_totalsignature() view returns(uint256)
func (_Api *ApiCaller) GetTotalsignature(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "get_totalsignature")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalsignature is a free data retrieval call binding the contract method 0x37f2678d.
//
// Solidity: function get_totalsignature() view returns(uint256)
func (_Api *ApiSession) GetTotalsignature() (*big.Int, error) {
	return _Api.Contract.GetTotalsignature(&_Api.CallOpts)
}

// GetTotalsignature is a free data retrieval call binding the contract method 0x37f2678d.
//
// Solidity: function get_totalsignature() view returns(uint256)
func (_Api *ApiCallerSession) GetTotalsignature() (*big.Int, error) {
	return _Api.Contract.GetTotalsignature(&_Api.CallOpts)
}

// ProfilePrivate is a free data retrieval call binding the contract method 0xaa43577c.
//
// Solidity: function profile_private(uint256 ) view returns(address)
func (_Api *ApiCaller) ProfilePrivate(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "profile_private", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProfilePrivate is a free data retrieval call binding the contract method 0xaa43577c.
//
// Solidity: function profile_private(uint256 ) view returns(address)
func (_Api *ApiSession) ProfilePrivate(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.ProfilePrivate(&_Api.CallOpts, arg0)
}

// ProfilePrivate is a free data retrieval call binding the contract method 0xaa43577c.
//
// Solidity: function profile_private(uint256 ) view returns(address)
func (_Api *ApiCallerSession) ProfilePrivate(arg0 *big.Int) (common.Address, error) {
	return _Api.Contract.ProfilePrivate(&_Api.CallOpts, arg0)
}

// ProfilesPublic is a free data retrieval call binding the contract method 0xee4b597d.
//
// Solidity: function profiles_public(uint256 ) view returns(uint256 profile_id, string idsignature, string password, string publickey, string role)
func (_Api *ApiCaller) ProfilesPublic(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ProfileId   *big.Int
	Idsignature string
	Password    string
	Publickey   string
	Role        string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "profiles_public", arg0)

	outstruct := new(struct {
		ProfileId   *big.Int
		Idsignature string
		Password    string
		Publickey   string
		Role        string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProfileId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Idsignature = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Password = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Publickey = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Role = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// ProfilesPublic is a free data retrieval call binding the contract method 0xee4b597d.
//
// Solidity: function profiles_public(uint256 ) view returns(uint256 profile_id, string idsignature, string password, string publickey, string role)
func (_Api *ApiSession) ProfilesPublic(arg0 *big.Int) (struct {
	ProfileId   *big.Int
	Idsignature string
	Password    string
	Publickey   string
	Role        string
}, error) {
	return _Api.Contract.ProfilesPublic(&_Api.CallOpts, arg0)
}

// ProfilesPublic is a free data retrieval call binding the contract method 0xee4b597d.
//
// Solidity: function profiles_public(uint256 ) view returns(uint256 profile_id, string idsignature, string password, string publickey, string role)
func (_Api *ApiCallerSession) ProfilesPublic(arg0 *big.Int) (struct {
	ProfileId   *big.Int
	Idsignature string
	Password    string
	Publickey   string
	Role        string
}, error) {
	return _Api.Contract.ProfilesPublic(&_Api.CallOpts, arg0)
}

// Signatures is a free data retrieval call binding the contract method 0x8be10194.
//
// Solidity: function signatures(uint256 ) view returns(uint256 signature_id, string profile_id, string signature, string datecreated)
func (_Api *ApiCaller) Signatures(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SignatureId *big.Int
	ProfileId   string
	Signature   string
	Datecreated string
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "signatures", arg0)

	outstruct := new(struct {
		SignatureId *big.Int
		ProfileId   string
		Signature   string
		Datecreated string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SignatureId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ProfileId = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Signature = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Datecreated = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// Signatures is a free data retrieval call binding the contract method 0x8be10194.
//
// Solidity: function signatures(uint256 ) view returns(uint256 signature_id, string profile_id, string signature, string datecreated)
func (_Api *ApiSession) Signatures(arg0 *big.Int) (struct {
	SignatureId *big.Int
	ProfileId   string
	Signature   string
	Datecreated string
}, error) {
	return _Api.Contract.Signatures(&_Api.CallOpts, arg0)
}

// Signatures is a free data retrieval call binding the contract method 0x8be10194.
//
// Solidity: function signatures(uint256 ) view returns(uint256 signature_id, string profile_id, string signature, string datecreated)
func (_Api *ApiCallerSession) Signatures(arg0 *big.Int) (struct {
	SignatureId *big.Int
	ProfileId   string
	Signature   string
	Datecreated string
}, error) {
	return _Api.Contract.Signatures(&_Api.CallOpts, arg0)
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

// AddProfile is a paid mutator transaction binding the contract method 0xd15de67f.
//
// Solidity: function add_profile(address _address, string _profile_id, string _idsignature, string _name, string _identity_card, string _email, string _phone, string _dateregistered) returns()
func (_Api *ApiTransactor) AddProfile(opts *bind.TransactOpts, _address common.Address, _profile_id string, _idsignature string, _name string, _identity_card string, _email string, _phone string, _dateregistered string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "add_profile", _address, _profile_id, _idsignature, _name, _identity_card, _email, _phone, _dateregistered)
}

// AddProfile is a paid mutator transaction binding the contract method 0xd15de67f.
//
// Solidity: function add_profile(address _address, string _profile_id, string _idsignature, string _name, string _identity_card, string _email, string _phone, string _dateregistered) returns()
func (_Api *ApiSession) AddProfile(_address common.Address, _profile_id string, _idsignature string, _name string, _identity_card string, _email string, _phone string, _dateregistered string) (*types.Transaction, error) {
	return _Api.Contract.AddProfile(&_Api.TransactOpts, _address, _profile_id, _idsignature, _name, _identity_card, _email, _phone, _dateregistered)
}

// AddProfile is a paid mutator transaction binding the contract method 0xd15de67f.
//
// Solidity: function add_profile(address _address, string _profile_id, string _idsignature, string _name, string _identity_card, string _email, string _phone, string _dateregistered) returns()
func (_Api *ApiTransactorSession) AddProfile(_address common.Address, _profile_id string, _idsignature string, _name string, _identity_card string, _email string, _phone string, _dateregistered string) (*types.Transaction, error) {
	return _Api.Contract.AddProfile(&_Api.TransactOpts, _address, _profile_id, _idsignature, _name, _identity_card, _email, _phone, _dateregistered)
}

// AddProfilefirst is a paid mutator transaction binding the contract method 0x1c87468e.
//
// Solidity: function add_profilefirst(string _idsignature, string _password, string _publickey, string _role) returns()
func (_Api *ApiTransactor) AddProfilefirst(opts *bind.TransactOpts, _idsignature string, _password string, _publickey string, _role string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "add_profilefirst", _idsignature, _password, _publickey, _role)
}

// AddProfilefirst is a paid mutator transaction binding the contract method 0x1c87468e.
//
// Solidity: function add_profilefirst(string _idsignature, string _password, string _publickey, string _role) returns()
func (_Api *ApiSession) AddProfilefirst(_idsignature string, _password string, _publickey string, _role string) (*types.Transaction, error) {
	return _Api.Contract.AddProfilefirst(&_Api.TransactOpts, _idsignature, _password, _publickey, _role)
}

// AddProfilefirst is a paid mutator transaction binding the contract method 0x1c87468e.
//
// Solidity: function add_profilefirst(string _idsignature, string _password, string _publickey, string _role) returns()
func (_Api *ApiTransactorSession) AddProfilefirst(_idsignature string, _password string, _publickey string, _role string) (*types.Transaction, error) {
	return _Api.Contract.AddProfilefirst(&_Api.TransactOpts, _idsignature, _password, _publickey, _role)
}

// AddSignature is a paid mutator transaction binding the contract method 0xc05e1001.
//
// Solidity: function add_signature(string _profile_id, string _signature, string _datecreated) returns()
func (_Api *ApiTransactor) AddSignature(opts *bind.TransactOpts, _profile_id string, _signature string, _datecreated string) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "add_signature", _profile_id, _signature, _datecreated)
}

// AddSignature is a paid mutator transaction binding the contract method 0xc05e1001.
//
// Solidity: function add_signature(string _profile_id, string _signature, string _datecreated) returns()
func (_Api *ApiSession) AddSignature(_profile_id string, _signature string, _datecreated string) (*types.Transaction, error) {
	return _Api.Contract.AddSignature(&_Api.TransactOpts, _profile_id, _signature, _datecreated)
}

// AddSignature is a paid mutator transaction binding the contract method 0xc05e1001.
//
// Solidity: function add_signature(string _profile_id, string _signature, string _datecreated) returns()
func (_Api *ApiTransactorSession) AddSignature(_profile_id string, _signature string, _datecreated string) (*types.Transaction, error) {
	return _Api.Contract.AddSignature(&_Api.TransactOpts, _profile_id, _signature, _datecreated)
}

// Addsigners is a paid mutator transaction binding the contract method 0xb5fa9cf9.
//
// Solidity: function addsigners(bytes32 _hash_original, address _signers, uint256 _signer_id, uint256 total) returns()
func (_Api *ApiTransactor) Addsigners(opts *bind.TransactOpts, _hash_original [32]byte, _signers common.Address, _signer_id *big.Int, total *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addsigners", _hash_original, _signers, _signer_id, total)
}

// Addsigners is a paid mutator transaction binding the contract method 0xb5fa9cf9.
//
// Solidity: function addsigners(bytes32 _hash_original, address _signers, uint256 _signer_id, uint256 total) returns()
func (_Api *ApiSession) Addsigners(_hash_original [32]byte, _signers common.Address, _signer_id *big.Int, total *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Addsigners(&_Api.TransactOpts, _hash_original, _signers, _signer_id, total)
}

// Addsigners is a paid mutator transaction binding the contract method 0xb5fa9cf9.
//
// Solidity: function addsigners(bytes32 _hash_original, address _signers, uint256 _signer_id, uint256 total) returns()
func (_Api *ApiTransactorSession) Addsigners(_hash_original [32]byte, _signers common.Address, _signer_id *big.Int, total *big.Int) (*types.Transaction, error) {
	return _Api.Contract.Addsigners(&_Api.TransactOpts, _hash_original, _signers, _signer_id, total)
}

// SignDoc is a paid mutator transaction binding the contract method 0x6ec0610d.
//
// Solidity: function signDoc(bytes32 _hash_original, string _namefile, string _hash_file, string _metadata, string _file, string _status, uint256 _totalsigned, string _datesigned, uint256 signing) returns()
func (_Api *ApiTransactor) SignDoc(opts *bind.TransactOpts, _hash_original [32]byte, _namefile string, _hash_file string, _metadata string, _file string, _status string, _totalsigned *big.Int, _datesigned string, signing *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "signDoc", _hash_original, _namefile, _hash_file, _metadata, _file, _status, _totalsigned, _datesigned, signing)
}

// SignDoc is a paid mutator transaction binding the contract method 0x6ec0610d.
//
// Solidity: function signDoc(bytes32 _hash_original, string _namefile, string _hash_file, string _metadata, string _file, string _status, uint256 _totalsigned, string _datesigned, uint256 signing) returns()
func (_Api *ApiSession) SignDoc(_hash_original [32]byte, _namefile string, _hash_file string, _metadata string, _file string, _status string, _totalsigned *big.Int, _datesigned string, signing *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SignDoc(&_Api.TransactOpts, _hash_original, _namefile, _hash_file, _metadata, _file, _status, _totalsigned, _datesigned, signing)
}

// SignDoc is a paid mutator transaction binding the contract method 0x6ec0610d.
//
// Solidity: function signDoc(bytes32 _hash_original, string _namefile, string _hash_file, string _metadata, string _file, string _status, uint256 _totalsigned, string _datesigned, uint256 signing) returns()
func (_Api *ApiTransactorSession) SignDoc(_hash_original [32]byte, _namefile string, _hash_file string, _metadata string, _file string, _status string, _totalsigned *big.Int, _datesigned string, signing *big.Int) (*types.Transaction, error) {
	return _Api.Contract.SignDoc(&_Api.TransactOpts, _hash_original, _namefile, _hash_file, _metadata, _file, _status, _totalsigned, _datesigned, signing)
}
