// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l1

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

// L1MetaData contains all meta data concerning the L1 contract.
var L1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040518060400160405280600d81526020017f4c32437573746f6d4552433230000000000000000000000000000000000000008152506040518060400160405280600381526020017f4c3154000000000000000000000000000000000000000000000000000000000081525081600390805190602001906200009692919062000345565b508060049080519060200190620000af92919062000345565b505050620000d2620000c6620000f460201b60201c565b620000fc60201b60201c565b620000ee3369d3c21bcecceda1000000620001c260201b60201c565b620005a1565b600033905090565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141562000235576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200022c9062000456565b60405180910390fd5b62000249600083836200033b60201b60201c565b80600260008282546200025d9190620004b1565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254620002b49190620004b1565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516200031b91906200051f565b60405180910390a362000337600083836200034060201b60201c565b5050565b505050565b505050565b82805462000353906200056b565b90600052602060002090601f016020900481019282620003775760008555620003c3565b82601f106200039257805160ff1916838001178555620003c3565b82800160010185558215620003c3579182015b82811115620003c2578251825591602001919060010190620003a5565b5b509050620003d29190620003d6565b5090565b5b80821115620003f1576000816000905550600101620003d7565b5090565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b60006200043e601f83620003f5565b91506200044b8262000406565b602082019050919050565b6000602082019050818103600083015262000471816200042f565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000620004be8262000478565b9150620004cb8362000478565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111562000503576200050262000482565b5b828201905092915050565b620005198162000478565b82525050565b60006020820190506200053660008301846200050e565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200058457607f821691505b602082108114156200059b576200059a6200053c565b5b50919050565b611c9280620005b16000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063715018a611610097578063a457c2d711610066578063a457c2d71461029d578063a9059cbb146102cd578063dd62ed3e146102fd578063f2fde38b1461032d57610100565b8063715018a61461023b57806379cc6790146102455780638da5cb5b1461026157806395d89b411461027f57610100565b8063313ce567116100d3578063313ce567146101a157806339509351146101bf57806342966c68146101ef57806370a082311461020b57610100565b806306fdde0314610105578063095ea7b31461012357806318160ddd1461015357806323b872dd14610171575b600080fd5b61010d610349565b60405161011a91906111e7565b60405180910390f35b61013d600480360381019061013891906112a2565b6103db565b60405161014a91906112fd565b60405180910390f35b61015b6103f9565b6040516101689190611327565b60405180910390f35b61018b60048036038101906101869190611342565b610403565b60405161019891906112fd565b60405180910390f35b6101a96104fb565b6040516101b691906113b1565b60405180910390f35b6101d960048036038101906101d491906112a2565b610504565b6040516101e691906112fd565b60405180910390f35b610209600480360381019061020491906113cc565b6105b0565b005b610225600480360381019061022091906113f9565b6105c4565b6040516102329190611327565b60405180910390f35b61024361060c565b005b61025f600480360381019061025a91906112a2565b610694565b005b61026961070f565b6040516102769190611435565b60405180910390f35b610287610739565b60405161029491906111e7565b60405180910390f35b6102b760048036038101906102b291906112a2565b6107cb565b6040516102c491906112fd565b60405180910390f35b6102e760048036038101906102e291906112a2565b6108b6565b6040516102f491906112fd565b60405180910390f35b61031760048036038101906103129190611450565b6108d4565b6040516103249190611327565b60405180910390f35b610347600480360381019061034291906113f9565b61095b565b005b606060038054610358906114bf565b80601f0160208091040260200160405190810160405280929190818152602001828054610384906114bf565b80156103d15780601f106103a6576101008083540402835291602001916103d1565b820191906000526020600020905b8154815290600101906020018083116103b457829003601f168201915b5050505050905090565b60006103ef6103e8610a53565b8484610a5b565b6001905092915050565b6000600254905090565b6000610410848484610c26565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600061045b610a53565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050828110156104db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d290611563565b60405180910390fd5b6104ef856104e7610a53565b858403610a5b565b60019150509392505050565b60006012905090565b60006105a6610511610a53565b84846001600061051f610a53565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546105a191906115b2565b610a5b565b6001905092915050565b6105c16105bb610a53565b82610ea7565b50565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b610614610a53565b73ffffffffffffffffffffffffffffffffffffffff1661063261070f565b73ffffffffffffffffffffffffffffffffffffffff1614610688576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067f90611654565b60405180910390fd5b610692600061107e565b565b60006106a7836106a2610a53565b6108d4565b9050818110156106ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e3906116e6565b60405180910390fd5b610700836106f8610a53565b848403610a5b565b61070a8383610ea7565b505050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b606060048054610748906114bf565b80601f0160208091040260200160405190810160405280929190818152602001828054610774906114bf565b80156107c15780601f10610796576101008083540402835291602001916107c1565b820191906000526020600020905b8154815290600101906020018083116107a457829003601f168201915b5050505050905090565b600080600160006107da610a53565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610897576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088e90611778565b60405180910390fd5b6108ab6108a2610a53565b85858403610a5b565b600191505092915050565b60006108ca6108c3610a53565b8484610c26565b6001905092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610963610a53565b73ffffffffffffffffffffffffffffffffffffffff1661098161070f565b73ffffffffffffffffffffffffffffffffffffffff16146109d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109ce90611654565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610a47576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a3e9061180a565b60405180910390fd5b610a508161107e565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610acb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ac29061189c565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610b3b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b329061192e565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610c199190611327565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610c96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c8d906119c0565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610d06576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cfd90611a52565b60405180910390fd5b610d11838383611144565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610d97576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d8e90611ae4565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610e2a91906115b2565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610e8e9190611327565b60405180910390a3610ea1848484611149565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610f17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f0e90611b76565b60405180910390fd5b610f2382600083611144565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610fa9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fa090611c08565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282546110009190611c28565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516110659190611327565b60405180910390a361107983600084611149565b505050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561118857808201518184015260208101905061116d565b83811115611197576000848401525b50505050565b6000601f19601f8301169050919050565b60006111b98261114e565b6111c38185611159565b93506111d381856020860161116a565b6111dc8161119d565b840191505092915050565b6000602082019050818103600083015261120181846111ae565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006112398261120e565b9050919050565b6112498161122e565b811461125457600080fd5b50565b60008135905061126681611240565b92915050565b6000819050919050565b61127f8161126c565b811461128a57600080fd5b50565b60008135905061129c81611276565b92915050565b600080604083850312156112b9576112b8611209565b5b60006112c785828601611257565b92505060206112d88582860161128d565b9150509250929050565b60008115159050919050565b6112f7816112e2565b82525050565b600060208201905061131260008301846112ee565b92915050565b6113218161126c565b82525050565b600060208201905061133c6000830184611318565b92915050565b60008060006060848603121561135b5761135a611209565b5b600061136986828701611257565b935050602061137a86828701611257565b925050604061138b8682870161128d565b9150509250925092565b600060ff82169050919050565b6113ab81611395565b82525050565b60006020820190506113c660008301846113a2565b92915050565b6000602082840312156113e2576113e1611209565b5b60006113f08482850161128d565b91505092915050565b60006020828403121561140f5761140e611209565b5b600061141d84828501611257565b91505092915050565b61142f8161122e565b82525050565b600060208201905061144a6000830184611426565b92915050565b6000806040838503121561146757611466611209565b5b600061147585828601611257565b925050602061148685828601611257565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806114d757607f821691505b602082108114156114eb576114ea611490565b5b50919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b600061154d602883611159565b9150611558826114f1565b604082019050919050565b6000602082019050818103600083015261157c81611540565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006115bd8261126c565b91506115c88361126c565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156115fd576115fc611583565b5b828201905092915050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b600061163e602083611159565b915061164982611608565b602082019050919050565b6000602082019050818103600083015261166d81611631565b9050919050565b7f45524332303a206275726e20616d6f756e74206578636565647320616c6c6f7760008201527f616e636500000000000000000000000000000000000000000000000000000000602082015250565b60006116d0602483611159565b91506116db82611674565b604082019050919050565b600060208201905081810360008301526116ff816116c3565b9050919050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000611762602583611159565b915061176d82611706565b604082019050919050565b6000602082019050818103600083015261179181611755565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b60006117f4602683611159565b91506117ff82611798565b604082019050919050565b60006020820190508181036000830152611823816117e7565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000611886602483611159565b91506118918261182a565b604082019050919050565b600060208201905081810360008301526118b581611879565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000611918602283611159565b9150611923826118bc565b604082019050919050565b600060208201905081810360008301526119478161190b565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b60006119aa602583611159565b91506119b58261194e565b604082019050919050565b600060208201905081810360008301526119d98161199d565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611a3c602383611159565b9150611a47826119e0565b604082019050919050565b60006020820190508181036000830152611a6b81611a2f565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000611ace602683611159565b9150611ad982611a72565b604082019050919050565b60006020820190508181036000830152611afd81611ac1565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000611b60602183611159565b9150611b6b82611b04565b604082019050919050565b60006020820190508181036000830152611b8f81611b53565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b6000611bf2602283611159565b9150611bfd82611b96565b604082019050919050565b60006020820190508181036000830152611c2181611be5565b9050919050565b6000611c338261126c565b9150611c3e8361126c565b925082821015611c5157611c50611583565b5b82820390509291505056fea2646970667358221220d37eca0d21378217fb3d6f08eca98ba0786633a45b09533f02996f783d23eb6364736f6c63430008090033",
}

// L1ABI is the input ABI used to generate the binding from.
// Deprecated: Use L1MetaData.ABI instead.
var L1ABI = L1MetaData.ABI

// L1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L1MetaData.Bin instead.
var L1Bin = L1MetaData.Bin

// DeployL1 deploys a new Ethereum contract, binding an instance of L1 to it.
func DeployL1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *L1, error) {
	parsed, err := L1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L1{L1Caller: L1Caller{contract: contract}, L1Transactor: L1Transactor{contract: contract}, L1Filterer: L1Filterer{contract: contract}}, nil
}

// L1 is an auto generated Go binding around an Ethereum contract.
type L1 struct {
	L1Caller     // Read-only binding to the contract
	L1Transactor // Write-only binding to the contract
	L1Filterer   // Log filterer for contract events
}

// L1Caller is an auto generated read-only Go binding around an Ethereum contract.
type L1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type L1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L1Session struct {
	Contract     *L1               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L1CallerSession struct {
	Contract *L1Caller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// L1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L1TransactorSession struct {
	Contract     *L1Transactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L1Raw is an auto generated low-level Go binding around an Ethereum contract.
type L1Raw struct {
	Contract *L1 // Generic contract binding to access the raw methods on
}

// L1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L1CallerRaw struct {
	Contract *L1Caller // Generic read-only contract binding to access the raw methods on
}

// L1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L1TransactorRaw struct {
	Contract *L1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewL1 creates a new instance of L1, bound to a specific deployed contract.
func NewL1(address common.Address, backend bind.ContractBackend) (*L1, error) {
	contract, err := bindL1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L1{L1Caller: L1Caller{contract: contract}, L1Transactor: L1Transactor{contract: contract}, L1Filterer: L1Filterer{contract: contract}}, nil
}

// NewL1Caller creates a new read-only instance of L1, bound to a specific deployed contract.
func NewL1Caller(address common.Address, caller bind.ContractCaller) (*L1Caller, error) {
	contract, err := bindL1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L1Caller{contract: contract}, nil
}

// NewL1Transactor creates a new write-only instance of L1, bound to a specific deployed contract.
func NewL1Transactor(address common.Address, transactor bind.ContractTransactor) (*L1Transactor, error) {
	contract, err := bindL1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L1Transactor{contract: contract}, nil
}

// NewL1Filterer creates a new log filterer instance of L1, bound to a specific deployed contract.
func NewL1Filterer(address common.Address, filterer bind.ContractFilterer) (*L1Filterer, error) {
	contract, err := bindL1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L1Filterer{contract: contract}, nil
}

// bindL1 binds a generic wrapper to an already deployed contract.
func bindL1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L1ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1 *L1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1.Contract.L1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1 *L1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1.Contract.L1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1 *L1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1.Contract.L1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L1 *L1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L1 *L1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L1 *L1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L1.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_L1 *L1Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_L1 *L1Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _L1.Contract.Allowance(&_L1.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_L1 *L1CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _L1.Contract.Allowance(&_L1.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_L1 *L1Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_L1 *L1Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _L1.Contract.BalanceOf(&_L1.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_L1 *L1CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _L1.Contract.BalanceOf(&_L1.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L1 *L1Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L1 *L1Session) Decimals() (uint8, error) {
	return _L1.Contract.Decimals(&_L1.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L1 *L1CallerSession) Decimals() (uint8, error) {
	return _L1.Contract.Decimals(&_L1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L1 *L1Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L1 *L1Session) Name() (string, error) {
	return _L1.Contract.Name(&_L1.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L1 *L1CallerSession) Name() (string, error) {
	return _L1.Contract.Name(&_L1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1 *L1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1 *L1Session) Owner() (common.Address, error) {
	return _L1.Contract.Owner(&_L1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_L1 *L1CallerSession) Owner() (common.Address, error) {
	return _L1.Contract.Owner(&_L1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L1 *L1Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L1 *L1Session) Symbol() (string, error) {
	return _L1.Contract.Symbol(&_L1.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L1 *L1CallerSession) Symbol() (string, error) {
	return _L1.Contract.Symbol(&_L1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L1 *L1Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L1.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L1 *L1Session) TotalSupply() (*big.Int, error) {
	return _L1.Contract.TotalSupply(&_L1.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L1 *L1CallerSession) TotalSupply() (*big.Int, error) {
	return _L1.Contract.TotalSupply(&_L1.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L1 *L1Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L1 *L1Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1.Contract.Approve(&_L1.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L1 *L1TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1.Contract.Approve(&_L1.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_L1 *L1Transactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_L1 *L1Session) Burn(amount *big.Int) (*types.Transaction, error) {
	return _L1.Contract.Burn(&_L1.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_L1 *L1TransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _L1.Contract.Burn(&_L1.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_L1 *L1Transactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_L1 *L1Session) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1.Contract.BurnFrom(&_L1.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_L1 *L1TransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1.Contract.BurnFrom(&_L1.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_L1 *L1Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_L1 *L1Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _L1.Contract.DecreaseAllowance(&_L1.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_L1 *L1TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _L1.Contract.DecreaseAllowance(&_L1.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_L1 *L1Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_L1 *L1Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _L1.Contract.IncreaseAllowance(&_L1.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_L1 *L1TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _L1.Contract.IncreaseAllowance(&_L1.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1 *L1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1 *L1Session) RenounceOwnership() (*types.Transaction, error) {
	return _L1.Contract.RenounceOwnership(&_L1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_L1 *L1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _L1.Contract.RenounceOwnership(&_L1.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_L1 *L1Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_L1 *L1Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1.Contract.Transfer(&_L1.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_L1 *L1TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1.Contract.Transfer(&_L1.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_L1 *L1Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_L1 *L1Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1.Contract.TransferFrom(&_L1.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_L1 *L1TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L1.Contract.TransferFrom(&_L1.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1 *L1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _L1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1 *L1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1.Contract.TransferOwnership(&_L1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_L1 *L1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _L1.Contract.TransferOwnership(&_L1.TransactOpts, newOwner)
}

// L1ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the L1 contract.
type L1ApprovalIterator struct {
	Event *L1Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L1ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L1Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L1ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1Approval represents a Approval event raised by the L1 contract.
type L1Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_L1 *L1Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*L1ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _L1.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &L1ApprovalIterator{contract: _L1.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_L1 *L1Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *L1Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _L1.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1Approval)
				if err := _L1.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_L1 *L1Filterer) ParseApproval(log types.Log) (*L1Approval, error) {
	event := new(L1Approval)
	if err := _L1.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the L1 contract.
type L1OwnershipTransferredIterator struct {
	Event *L1OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L1OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1OwnershipTransferred represents a OwnershipTransferred event raised by the L1 contract.
type L1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1 *L1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*L1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &L1OwnershipTransferredIterator{contract: _L1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1 *L1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *L1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _L1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1OwnershipTransferred)
				if err := _L1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_L1 *L1Filterer) ParseOwnershipTransferred(log types.Log) (*L1OwnershipTransferred, error) {
	event := new(L1OwnershipTransferred)
	if err := _L1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L1TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the L1 contract.
type L1TransferIterator struct {
	Event *L1Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *L1TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L1Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(L1Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *L1TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L1TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L1Transfer represents a Transfer event raised by the L1 contract.
type L1Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L1 *L1Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L1TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L1TransferIterator{contract: _L1.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L1 *L1Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *L1Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L1.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L1Transfer)
				if err := _L1.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L1 *L1Filterer) ParseTransfer(log types.Log) (*L1Transfer, error) {
	event := new(L1Transfer)
	if err := _L1.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
