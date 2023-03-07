package utils

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"fmt"
	"geth-performance-test/contracts/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
	"log"
	"math"
	"math/big"
	"testing"
	"time"
)

func GetRpcClient(rpcUrl string, chainID *big.Int) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return client, nil
}

func BuildAuth(client *ethclient.Client, userPrivateKey string, chainID *big.Int, account *big.Int) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(userPrivateKey)
	if err != nil {
		log.Fatalln(err.Error())
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	publicKeyPublic := privateKey.Public()
	publicKeyECDSA, ok := publicKeyPublic.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("publicKeyECDSA-Error", ok)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(10000000)
	auth.Value = account
	//auth.GasPrice = big.NewInt(0)
	auth.GasPrice = big.NewInt(5000000)
	return auth
}

func BuildAuthL2(client *ethclient.Client, userPrivateKey string, chainID *big.Int, account *big.Int) *bind.TransactOpts {
	privateKey, err := crypto.HexToECDSA(userPrivateKey)
	if err != nil {
		log.Fatalln(err.Error())
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	publicKeyPublic := privateKey.Public()
	publicKeyECDSA, ok := publicKeyPublic.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("publicKeyECDSA-Error", ok)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalln(err.Error())
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.GasLimit = uint64(10000000)
	auth.Value = account
	auth.GasPrice = big.NewInt(0)
	return auth
}

func SetL1Approve(client *ethclient.Client, userPrivateKey string, chainID *big.Int, l1tokenAddress string, l1BridgeAddress string) {
	l1Instance, _ := token.NewToken(common.HexToAddress(l1tokenAddress), client)
	auth := BuildAuth(client, userPrivateKey, chainID, big.NewInt(0))
	n := new(big.Int)
	n, ok := n.SetString("900000000000000000000000000", 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	tx, err := l1Instance.Approve(auth, common.HexToAddress(l1BridgeAddress), n)
	if err != nil {
		log.Fatalln("publicKeyECDSA-Error", err.Error())
	}
	log.Println("setL1Approve tx.Hash:", tx.Hash())
}

func SetL2Approve(client *ethclient.Client, userPrivateKey string, chainID *big.Int, l2tokenAddress string, l2BridgeAddress string) {
	l2Instance, _ := token.NewToken(common.HexToAddress(l2tokenAddress), client)
	auth := BuildAuthL2(client, userPrivateKey, chainID, big.NewInt(0))
	n := new(big.Int)
	n, ok := n.SetString("1000000000000000000", 10)
	if !ok {
		fmt.Println("SetString: error")
		return
	}
	tx, err := l2Instance.Approve(auth, common.HexToAddress(l2BridgeAddress), n)
	if err != nil {
		log.Fatalln("SetL2Approve-error", err.Error())
	}
	log.Println("SetL2Approve tx.Hash:", tx.Hash())
}

func GetBalanceStr(client *ethclient.Client, chainID *big.Int, toAddress string) *big.Float {
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(toAddress), nil)
	if err != nil {
		log.Printf("eth get balance err:%v", err.Error())
		return nil
	}

	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))

	log.Printf("eth-Balance->chainID:%s,balance:%v is eth:", chainID, ethValue)

	return ethValue
}

func GetBalance(client *ethclient.Client, chainID *big.Int, toAddress string, t *testing.T) *big.Float {
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(toAddress), nil)
	if err != nil {
		t.Log(err.Error())
	}

	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	t.Logf("eth-Balance->chainID:%s,balance:%v is eth:", chainID, ethValue)

	return ethValue
}

func GetHash(client *ethclient.Client, hash common.Hash, t *testing.T) {
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	blockByHash, err := client.BlockByHash(context.Background(), receipt.BlockHash)
	require.NoError(t, err, "BlockByHash error")
	t.Log("status:", receipt.Status)
	t.Log("blockHash: ", blockByHash.Hash())
	t.Log("blockNumber: ", blockByHash.Number())
	t.Log("block timestamp: ", blockByHash.Time())
	formatTime := time.Unix(int64(blockByHash.Time()), 0).Format(time.RFC3339)
	t.Log("block timestformatTimeamp: ", formatTime)
	t.Log("----------------------------------------------------------------------------- ")

}

func ClientSimulated(client *ethclient.Client, toAddress common.Address, chainID *big.Int, hexToECDSA string, nonce *big.Int) (*big.Int, error) {
	n := new(big.Int)
	n, ok := n.SetString("5000000000000000000", 10)
	if !ok {
		return nonce, errors.New("SetString: error")
	}
	gasLimit := uint64(210000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nonce, err
	}

	var data []byte
	gasPrice = big.NewInt(1)
	tx := types.NewTransaction(nonce.Uint64(), toAddress, n, gasLimit, gasPrice, data)

	privateKey, err := crypto.HexToECDSA(hexToECDSA)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return nonce, err
	}

	//生成裸交易
	ts := types.Transactions{signedTx}
	rawTxBytes := bytes.NewBuffer([]byte{})
	ts.EncodeIndex(0, rawTxBytes)
	rawTxHex := hex.EncodeToString(rawTxBytes.Bytes())
	log.Printf("nonce->%s,裸交易:%s", nonce, rawTxHex) // f86...772

	//将裸交易转换广播出去
	rawTxHexBytes, err := hex.DecodeString(rawTxHex)

	mewTx := new(types.Transaction)
	rlp.DecodeBytes(rawTxHexBytes, &mewTx)

	err = client.SendTransaction(context.Background(), mewTx)
	if err != nil {
		return nonce, err
	} else {
		fmt.Printf("tx sent: %s,nonce->%s\n", mewTx.Hash().Hex(), nonce) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
		return nonce, nil
	}
}
