package utils

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/stretchr/testify/require"
	"log"
	"math/big"
	"os"
	"strings"
	"testing"
	"time"
)

func DistributeBit(l2Client *ethclient.Client, UserAddress string, UserPrivateKey string, t *testing.T) {

	userAddress := common.HexToAddress(UserAddress)

	balancel, err := l2Client.BalanceAt(context.Background(), userAddress, nil)
	require.NoError(t, err, "BalanceOf error")
	t.Log("bit-balance->", balancel)

	//文件读取
	f, err := os.Open("./secrect.csv")
	require.NoError(t, err, "os.Ope error")

	reader := csv.NewReader(f)
	preData, err := reader.ReadAll()
	require.NoError(t, err, "reader.ReadAll error")

	nonce, err := l2Client.PendingNonceAt(context.Background(), userAddress)
	if err != nil {
		log.Fatal("PendingNonceAt->", err)
	}
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
		n, _ = n.SetString("1000000000000000000", 10)

		for {
			fmt.Println("TransferBit: ->nonce->", nonce)
			e := TransferBit(l2Client, privateKey, to, n, nonce)
			if e != nil {
				errMsg := fmt.Sprintf("%v", e)
				if strings.Contains(errMsg, "nonce too low") {
					nonce, err = l2Client.PendingNonceAt(context.Background(), userAddress)
					if err != nil {
						log.Fatal("PendingNonceAt->", err)
					}
				}
				fmt.Println("TransferBit: error->nonce->", errMsg, nonce)
			} else {
				nonce += 1
				break
			}
		}
		time.Sleep(time.Microsecond * time.Duration(400))

	}
}

func TransferBit(client *ethclient.Client, fromPrivateKey *ecdsa.PrivateKey, toAddress common.Address, n *big.Int, nonce uint64) error {

	gasLimit := uint64(21000) // in units
	gasPrice := big.NewInt(1)

	var data []byte
	tx := types.NewTransaction(nonce, toAddress, n, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(17)), fromPrivateKey)
	if err != nil {
		return err
	}

	//生成裸交易
	ts := types.Transactions{signedTx}
	rawTxBytes := bytes.NewBuffer([]byte{})

	ts.EncodeIndex(0, rawTxBytes)
	rawTxHex := hex.EncodeToString(rawTxBytes.Bytes())
	//log.Printf("%v send  nonce->%s,裸交易:%s", crypto.PubkeyToAddress(fromPrivateKey.PublicKey), nonce, rawTxHex) // f86...772

	//将裸交易转换广播出去
	rawTxHexBytes, err := hex.DecodeString(rawTxHex)

	mewTx := new(types.Transaction)
	rlp.DecodeBytes(rawTxHexBytes, &mewTx)

	err = client.SendTransaction(context.Background(), mewTx)
	if err != nil {
		log.Printf("send tx error : %v", err)
		return err
	}
	return nil
}
