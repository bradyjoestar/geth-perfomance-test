package utils

import (
	"context"
	"crypto/ecdsa"
	"encoding/csv"
	"fmt"
	"geth-performance-test/contracts/token/l2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/require"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"
)

/*
*
ERC20AccountTransfer
*/
func DistributeErc20(l2Client *ethclient.Client, UserAddress string,
	L2Erc20Address string, UserPrivateKey string,
	L2ChainID *big.Int, t *testing.T) {
	l2Erc20, err := l2.NewL2(common.HexToAddress(L2Erc20Address), l2Client)
	log.Printf(common.HexToAddress(L2Erc20Address).String())
	require.NoError(t, err, "L1Erc20Address client error")
	balancel, err := l2Erc20.BalanceOf(&bind.CallOpts{}, common.HexToAddress(UserAddress))
	require.NoError(t, err, "BalanceOf error")
	t.Log("l2Erc20-balance->", balancel)

	balancel1 := GetBalance(l2Client, L2ChainID, UserAddress, t)
	t.Log("eth-balance->", balancel1)
	//文件读取
	f, err := os.Open("./secrect.csv")
	require.NoError(t, err, "os.Ope error")

	reader := csv.NewReader(f)
	preData, err := reader.ReadAll()
	require.NoError(t, err, "reader.ReadAll error")

	nonce, err := l2Client.PendingNonceAt(context.Background(), common.HexToAddress(UserAddress))
	if err != nil {
		log.Fatal("PendingNonceAt->", err)
	}

	userBalancel2, err := l2Erc20.BalanceOf(&bind.CallOpts{}, common.HexToAddress(UserAddress))
	log.Printf("====================")

	log.Printf("user erc20 balance is : %s", userBalancel2)

	log.Printf("====================")

	for i := 0; i < len(preData); i++ {
		privateKey, err := crypto.HexToECDSA(UserPrivateKey)
		if err != nil {
			log.Fatalln(err.Error())
		}

		to := common.HexToAddress(preData[i][1])
		if err != nil {
			log.Fatalln(err.Error())
		}
		n := new(big.Int)
		n, _ = n.SetString("10000000000", 10)

		for {
			fmt.Println("TransferErc20: ->nonce->", nonce)
			e := TransferErc20(l2Erc20, privateKey, to, n, big.NewInt(int64(nonce)))
			if e != nil {
				errMsg := fmt.Sprintf("%v", e)
				if strings.Contains(errMsg, "nonce too low") {
					nonce += 1
				}
				fmt.Println("TransferErc20: error->nonce->", errMsg, nonce)
			} else {
				nonce += 1
				break
			}
		}
		time.Sleep(time.Microsecond * time.Duration(400))
		balancel2, err := l2Erc20.BalanceOf(&bind.CallOpts{}, to)
		require.NoError(t, err, "BalanceOf error")
		log.Printf("balanceL2 ： %v", balancel2)
	}
}
func TransferErc20(l2Erc20 *l2.L2, from *ecdsa.PrivateKey, to common.Address, n *big.Int, nonce *big.Int) error {
	auth := bind.NewKeyedTransactor(from)
	auth.Nonce = nonce
	auth.GasLimit = uint64(10000000)
	auth.GasPrice = big.NewInt(1)
	tra, err := l2Erc20.Transfer(auth, to, n)
	if err != nil {
		return err
	} else {
		fmt.Printf("TransferErc20 tx sent: %s\n", tra.Hash().Hex())
		return nil
	}
}
