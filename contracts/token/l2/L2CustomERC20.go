// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package l2

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

// L2MetaData contains all meta data concerning the L2 contract.
var L2MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_l2Bridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_l1Token\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l1Token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"l2Bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620023da380380620023da833981810160405281019062000037919062000422565b81816040518060400160405280600f81526020017f437573746f6d204c3220546f6b656e00000000000000000000000000000000008152506040518060400160405280600381526020017f4c3254000000000000000000000000000000000000000000000000000000000081525081818160039080519060200190620000bf92919062000308565b508060049080519060200190620000d892919062000308565b50505082600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050506200017d3369d3c21bcecceda10000006200018560201b60201c565b505062000615565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415620001f8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620001ef90620004ca565b60405180910390fd5b6200020c60008383620002fe60201b60201c565b806002600082825462000220919062000525565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825462000277919062000525565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051620002de919062000593565b60405180910390a3620002fa600083836200030360201b60201c565b5050565b505050565b505050565b8280546200031690620005df565b90600052602060002090601f0160209004810192826200033a576000855562000386565b82601f106200035557805160ff191683800117855562000386565b8280016001018555821562000386579182015b828111156200038557825182559160200191906001019062000368565b5b50905062000395919062000399565b5090565b5b80821115620003b45760008160009055506001016200039a565b5090565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620003ea82620003bd565b9050919050565b620003fc81620003dd565b81146200040857600080fd5b50565b6000815190506200041c81620003f1565b92915050565b600080604083850312156200043c576200043b620003b8565b5b60006200044c858286016200040b565b92505060206200045f858286016200040b565b9150509250929050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000620004b2601f8362000469565b9150620004bf826200047a565b602082019050919050565b60006020820190508181036000830152620004e581620004a3565b9050919050565b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006200053282620004ec565b91506200053f83620004ec565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115620005775762000576620004f6565b5b828201905092915050565b6200058d81620004ec565b82525050565b6000602082019050620005aa600083018462000582565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620005f857607f821691505b602082108114156200060f576200060e620005b0565b5b50919050565b611db580620006256000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806370a0823111610097578063a9059cbb11610066578063a9059cbb146102d5578063ae1f6aaf14610305578063c01e1bd614610323578063dd62ed3e1461034157610100565b806370a082311461023b57806395d89b411461026b5780639dc29fac14610289578063a457c2d7146102a557610100565b806323b872dd116100d357806323b872dd146101a1578063313ce567146101d157806339509351146101ef57806340c10f191461021f57610100565b806301ffc9a71461010557806306fdde0314610135578063095ea7b31461015357806318160ddd14610183575b600080fd5b61011f600480360381019061011a919061132e565b610371565b60405161012c9190611376565b60405180910390f35b61013d610447565b60405161014a919061142a565b60405180910390f35b61016d600480360381019061016891906114e0565b6104d9565b60405161017a9190611376565b60405180910390f35b61018b6104f7565b604051610198919061152f565b60405180910390f35b6101bb60048036038101906101b6919061154a565b610501565b6040516101c89190611376565b60405180910390f35b6101d96105f9565b6040516101e691906115b9565b60405180910390f35b610209600480360381019061020491906114e0565b610602565b6040516102169190611376565b60405180910390f35b610239600480360381019061023491906114e0565b6106ae565b005b610255600480360381019061025091906115d4565b61079a565b604051610262919061152f565b60405180910390f35b6102736107e2565b604051610280919061142a565b60405180910390f35b6102a3600480360381019061029e91906114e0565b610874565b005b6102bf60048036038101906102ba91906114e0565b610960565b6040516102cc9190611376565b60405180910390f35b6102ef60048036038101906102ea91906114e0565b610a4b565b6040516102fc9190611376565b60405180910390f35b61030d610a69565b60405161031a9190611610565b60405180910390f35b61032b610a8f565b6040516103389190611610565b60405180910390f35b61035b6004803603810190610356919061162b565b610ab5565b604051610368919061152f565b60405180910390f35b6000807f01ffc9a7a5cef8baa21ed3c5c0d7e23accb804b619e9333b597f47a0d84076e290506000639dc29fac60e01b6340c10f1960e01b63c01e1bd660e01b18189050817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061043e5750807bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916847bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b92505050919050565b6060600380546104569061169a565b80601f01602080910402602001604051908101604052809291908181526020018280546104829061169a565b80156104cf5780601f106104a4576101008083540402835291602001916104cf565b820191906000526020600020905b8154815290600101906020018083116104b257829003601f168201915b5050505050905090565b60006104ed6104e6610b3c565b8484610b44565b6001905092915050565b6000600254905090565b600061050e848484610d0f565b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610559610b3c565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050828110156105d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105d09061173e565b60405180910390fd5b6105ed856105e5610b3c565b858403610b44565b60019150509392505050565b60006008905090565b60006106a461060f610b3c565b84846001600061061d610b3c565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461069f919061178d565b610b44565b6001905092915050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461073e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107359061182f565b60405180910390fd5b6107488282610f90565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d41213968858260405161078e919061152f565b60405180910390a25050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546107f19061169a565b80601f016020809104026020016040519081016040528092919081815260200182805461081d9061169a565b801561086a5780601f1061083f5761010080835404028352916020019161086a565b820191906000526020600020905b81548152906001019060200180831161084d57829003601f168201915b5050505050905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610904576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108fb9061182f565b60405180910390fd5b61090e82826110f0565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca582604051610954919061152f565b60405180910390a25050565b6000806001600061096f610b3c565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905082811015610a2c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a23906118c1565b60405180910390fd5b610a40610a37610b3c565b85858403610b44565b600191505092915050565b6000610a5f610a58610b3c565b8484610d0f565b6001905092915050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610bb4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bab90611953565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610c24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1b906119e5565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610d02919061152f565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610d7f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d7690611a77565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610def576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610de690611b09565b60405180910390fd5b610dfa8383836112c7565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015610e80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e7790611b9b565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f13919061178d565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610f77919061152f565b60405180910390a3610f8a8484846112cc565b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611000576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff790611c07565b60405180910390fd5b61100c600083836112c7565b806002600082825461101e919061178d565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254611073919061178d565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516110d8919061152f565b60405180910390a36110ec600083836112cc565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415611160576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161115790611c99565b60405180910390fd5b61116c826000836112c7565b60008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156111f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111e990611d2b565b60405180910390fd5b8181036000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600260008282546112499190611d4b565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516112ae919061152f565b60405180910390a36112c2836000846112cc565b505050565b505050565b505050565b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b61130b816112d6565b811461131657600080fd5b50565b60008135905061132881611302565b92915050565b600060208284031215611344576113436112d1565b5b600061135284828501611319565b91505092915050565b60008115159050919050565b6113708161135b565b82525050565b600060208201905061138b6000830184611367565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156113cb5780820151818401526020810190506113b0565b838111156113da576000848401525b50505050565b6000601f19601f8301169050919050565b60006113fc82611391565b611406818561139c565b93506114168185602086016113ad565b61141f816113e0565b840191505092915050565b6000602082019050818103600083015261144481846113f1565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006114778261144c565b9050919050565b6114878161146c565b811461149257600080fd5b50565b6000813590506114a48161147e565b92915050565b6000819050919050565b6114bd816114aa565b81146114c857600080fd5b50565b6000813590506114da816114b4565b92915050565b600080604083850312156114f7576114f66112d1565b5b600061150585828601611495565b9250506020611516858286016114cb565b9150509250929050565b611529816114aa565b82525050565b60006020820190506115446000830184611520565b92915050565b600080600060608486031215611563576115626112d1565b5b600061157186828701611495565b935050602061158286828701611495565b9250506040611593868287016114cb565b9150509250925092565b600060ff82169050919050565b6115b38161159d565b82525050565b60006020820190506115ce60008301846115aa565b92915050565b6000602082840312156115ea576115e96112d1565b5b60006115f884828501611495565b91505092915050565b61160a8161146c565b82525050565b60006020820190506116256000830184611601565b92915050565b60008060408385031215611642576116416112d1565b5b600061165085828601611495565b925050602061166185828601611495565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806116b257607f821691505b602082108114156116c6576116c561166b565b5b50919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206160008201527f6c6c6f77616e6365000000000000000000000000000000000000000000000000602082015250565b600061172860288361139c565b9150611733826116cc565b604082019050919050565b600060208201905081810360008301526117578161171b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611798826114aa565b91506117a3836114aa565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156117d8576117d761175e565b5b828201905092915050565b7f4f6e6c79204c32204272696467652063616e206d696e7420616e64206275726e600082015250565b600061181960208361139c565b9150611824826117e3565b602082019050919050565b600060208201905081810360008301526118488161180c565b9050919050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b60006118ab60258361139c565b91506118b68261184f565b604082019050919050565b600060208201905081810360008301526118da8161189e565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b600061193d60248361139c565b9150611948826118e1565b604082019050919050565b6000602082019050818103600083015261196c81611930565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b60006119cf60228361139c565b91506119da82611973565b604082019050919050565b600060208201905081810360008301526119fe816119c2565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b6000611a6160258361139c565b9150611a6c82611a05565b604082019050919050565b60006020820190508181036000830152611a9081611a54565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611af360238361139c565b9150611afe82611a97565b604082019050919050565b60006020820190508181036000830152611b2281611ae6565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b6000611b8560268361139c565b9150611b9082611b29565b604082019050919050565b60006020820190508181036000830152611bb481611b78565b9050919050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611bf1601f8361139c565b9150611bfc82611bbb565b602082019050919050565b60006020820190508181036000830152611c2081611be4565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000611c8360218361139c565b9150611c8e82611c27565b604082019050919050565b60006020820190508181036000830152611cb281611c76565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b6000611d1560228361139c565b9150611d2082611cb9565b604082019050919050565b60006020820190508181036000830152611d4481611d08565b9050919050565b6000611d56826114aa565b9150611d61836114aa565b925082821015611d7457611d7361175e565b5b82820390509291505056fea2646970667358221220a84edcf53eba2426436d9c10df8b7957f031f8b687ef064d048a087f6e6baa1364736f6c63430008090033",
}

// L2ABI is the input ABI used to generate the binding from.
// Deprecated: Use L2MetaData.ABI instead.
var L2ABI = L2MetaData.ABI

// L2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use L2MetaData.Bin instead.
var L2Bin = L2MetaData.Bin

// DeployL2 deploys a new Ethereum contract, binding an instance of L2 to it.
func DeployL2(auth *bind.TransactOpts, backend bind.ContractBackend, _l2Bridge common.Address, _l1Token common.Address) (common.Address, *types.Transaction, *L2, error) {
	parsed, err := L2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(L2Bin), backend, _l2Bridge, _l1Token)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &L2{L2Caller: L2Caller{contract: contract}, L2Transactor: L2Transactor{contract: contract}, L2Filterer: L2Filterer{contract: contract}}, nil
}

// L2 is an auto generated Go binding around an Ethereum contract.
type L2 struct {
	L2Caller     // Read-only binding to the contract
	L2Transactor // Write-only binding to the contract
	L2Filterer   // Log filterer for contract events
}

// L2Caller is an auto generated read-only Go binding around an Ethereum contract.
type L2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type L2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type L2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// L2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type L2Session struct {
	Contract     *L2               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type L2CallerSession struct {
	Contract *L2Caller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// L2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type L2TransactorSession struct {
	Contract     *L2Transactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// L2Raw is an auto generated low-level Go binding around an Ethereum contract.
type L2Raw struct {
	Contract *L2 // Generic contract binding to access the raw methods on
}

// L2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type L2CallerRaw struct {
	Contract *L2Caller // Generic read-only contract binding to access the raw methods on
}

// L2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type L2TransactorRaw struct {
	Contract *L2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewL2 creates a new instance of L2, bound to a specific deployed contract.
func NewL2(address common.Address, backend bind.ContractBackend) (*L2, error) {
	contract, err := bindL2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &L2{L2Caller: L2Caller{contract: contract}, L2Transactor: L2Transactor{contract: contract}, L2Filterer: L2Filterer{contract: contract}}, nil
}

// NewL2Caller creates a new read-only instance of L2, bound to a specific deployed contract.
func NewL2Caller(address common.Address, caller bind.ContractCaller) (*L2Caller, error) {
	contract, err := bindL2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &L2Caller{contract: contract}, nil
}

// NewL2Transactor creates a new write-only instance of L2, bound to a specific deployed contract.
func NewL2Transactor(address common.Address, transactor bind.ContractTransactor) (*L2Transactor, error) {
	contract, err := bindL2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &L2Transactor{contract: contract}, nil
}

// NewL2Filterer creates a new log filterer instance of L2, bound to a specific deployed contract.
func NewL2Filterer(address common.Address, filterer bind.ContractFilterer) (*L2Filterer, error) {
	contract, err := bindL2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &L2Filterer{contract: contract}, nil
}

// bindL2 binds a generic wrapper to an already deployed contract.
func bindL2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(L2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2 *L2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2.Contract.L2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2 *L2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2.Contract.L2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2 *L2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2.Contract.L2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_L2 *L2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _L2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_L2 *L2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _L2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_L2 *L2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _L2.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_L2 *L2Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_L2 *L2Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _L2.Contract.Allowance(&_L2.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_L2 *L2CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _L2.Contract.Allowance(&_L2.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_L2 *L2Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_L2 *L2Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _L2.Contract.BalanceOf(&_L2.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_L2 *L2CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _L2.Contract.BalanceOf(&_L2.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L2 *L2Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L2 *L2Session) Decimals() (uint8, error) {
	return _L2.Contract.Decimals(&_L2.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_L2 *L2CallerSession) Decimals() (uint8, error) {
	return _L2.Contract.Decimals(&_L2.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_L2 *L2Caller) L1Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "l1Token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_L2 *L2Session) L1Token() (common.Address, error) {
	return _L2.Contract.L1Token(&_L2.CallOpts)
}

// L1Token is a free data retrieval call binding the contract method 0xc01e1bd6.
//
// Solidity: function l1Token() view returns(address)
func (_L2 *L2CallerSession) L1Token() (common.Address, error) {
	return _L2.Contract.L1Token(&_L2.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L2 *L2Caller) L2Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "l2Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L2 *L2Session) L2Bridge() (common.Address, error) {
	return _L2.Contract.L2Bridge(&_L2.CallOpts)
}

// L2Bridge is a free data retrieval call binding the contract method 0xae1f6aaf.
//
// Solidity: function l2Bridge() view returns(address)
func (_L2 *L2CallerSession) L2Bridge() (common.Address, error) {
	return _L2.Contract.L2Bridge(&_L2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L2 *L2Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L2 *L2Session) Name() (string, error) {
	return _L2.Contract.Name(&_L2.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_L2 *L2CallerSession) Name() (string, error) {
	return _L2.Contract.Name(&_L2.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_L2 *L2Caller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "supportsInterface", _interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_L2 *L2Session) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _L2.Contract.SupportsInterface(&_L2.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceId) pure returns(bool)
func (_L2 *L2CallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _L2.Contract.SupportsInterface(&_L2.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L2 *L2Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L2 *L2Session) Symbol() (string, error) {
	return _L2.Contract.Symbol(&_L2.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_L2 *L2CallerSession) Symbol() (string, error) {
	return _L2.Contract.Symbol(&_L2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L2 *L2Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _L2.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L2 *L2Session) TotalSupply() (*big.Int, error) {
	return _L2.Contract.TotalSupply(&_L2.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_L2 *L2CallerSession) TotalSupply() (*big.Int, error) {
	return _L2.Contract.TotalSupply(&_L2.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L2 *L2Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L2 *L2Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.Approve(&_L2.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_L2 *L2TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.Approve(&_L2.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_L2 *L2Transactor) Burn(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "burn", _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_L2 *L2Session) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.Burn(&_L2.TransactOpts, _from, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address _from, uint256 _amount) returns()
func (_L2 *L2TransactorSession) Burn(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.Burn(&_L2.TransactOpts, _from, _amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_L2 *L2Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_L2 *L2Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _L2.Contract.DecreaseAllowance(&_L2.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_L2 *L2TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _L2.Contract.DecreaseAllowance(&_L2.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_L2 *L2Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_L2 *L2Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _L2.Contract.IncreaseAllowance(&_L2.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_L2 *L2TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _L2.Contract.IncreaseAllowance(&_L2.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_L2 *L2Transactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_L2 *L2Session) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.Mint(&_L2.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_L2 *L2TransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.Mint(&_L2.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_L2 *L2Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_L2 *L2Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.Transfer(&_L2.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_L2 *L2TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.Transfer(&_L2.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_L2 *L2Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_L2 *L2Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.TransferFrom(&_L2.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_L2 *L2TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _L2.Contract.TransferFrom(&_L2.TransactOpts, sender, recipient, amount)
}

// L2ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the L2 contract.
type L2ApprovalIterator struct {
	Event *L2Approval // Event containing the contract specifics and raw log

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
func (it *L2ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2Approval)
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
		it.Event = new(L2Approval)
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
func (it *L2ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2Approval represents a Approval event raised by the L2 contract.
type L2Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_L2 *L2Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*L2ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _L2.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &L2ApprovalIterator{contract: _L2.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_L2 *L2Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *L2Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _L2.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2Approval)
				if err := _L2.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_L2 *L2Filterer) ParseApproval(log types.Log) (*L2Approval, error) {
	event := new(L2Approval)
	if err := _L2.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the L2 contract.
type L2BurnIterator struct {
	Event *L2Burn // Event containing the contract specifics and raw log

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
func (it *L2BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2Burn)
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
		it.Event = new(L2Burn)
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
func (it *L2BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2Burn represents a Burn event raised by the L2 contract.
type L2Burn struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed _account, uint256 _amount)
func (_L2 *L2Filterer) FilterBurn(opts *bind.FilterOpts, _account []common.Address) (*L2BurnIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _L2.contract.FilterLogs(opts, "Burn", _accountRule)
	if err != nil {
		return nil, err
	}
	return &L2BurnIterator{contract: _L2.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed _account, uint256 _amount)
func (_L2 *L2Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *L2Burn, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _L2.contract.WatchLogs(opts, "Burn", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2Burn)
				if err := _L2.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed _account, uint256 _amount)
func (_L2 *L2Filterer) ParseBurn(log types.Log) (*L2Burn, error) {
	event := new(L2Burn)
	if err := _L2.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the L2 contract.
type L2MintIterator struct {
	Event *L2Mint // Event containing the contract specifics and raw log

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
func (it *L2MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2Mint)
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
		it.Event = new(L2Mint)
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
func (it *L2MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2Mint represents a Mint event raised by the L2 contract.
type L2Mint struct {
	Account common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _account, uint256 _amount)
func (_L2 *L2Filterer) FilterMint(opts *bind.FilterOpts, _account []common.Address) (*L2MintIterator, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _L2.contract.FilterLogs(opts, "Mint", _accountRule)
	if err != nil {
		return nil, err
	}
	return &L2MintIterator{contract: _L2.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _account, uint256 _amount)
func (_L2 *L2Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *L2Mint, _account []common.Address) (event.Subscription, error) {

	var _accountRule []interface{}
	for _, _accountItem := range _account {
		_accountRule = append(_accountRule, _accountItem)
	}

	logs, sub, err := _L2.contract.WatchLogs(opts, "Mint", _accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2Mint)
				if err := _L2.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: event Mint(address indexed _account, uint256 _amount)
func (_L2 *L2Filterer) ParseMint(log types.Log) (*L2Mint, error) {
	event := new(L2Mint)
	if err := _L2.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// L2TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the L2 contract.
type L2TransferIterator struct {
	Event *L2Transfer // Event containing the contract specifics and raw log

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
func (it *L2TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(L2Transfer)
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
		it.Event = new(L2Transfer)
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
func (it *L2TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *L2TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// L2Transfer represents a Transfer event raised by the L2 contract.
type L2Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L2 *L2Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*L2TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &L2TransferIterator{contract: _L2.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_L2 *L2Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *L2Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _L2.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(L2Transfer)
				if err := _L2.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_L2 *L2Filterer) ParseTransfer(log types.Log) (*L2Transfer, error) {
	event := new(L2Transfer)
	if err := _L2.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
