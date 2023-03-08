package jobs

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/mantlenetworkio/mantle/l2geth/common"
	"github.com/mantlenetworkio/mantle/l2geth/core/types"
	"github.com/mantlenetworkio/mantle/l2geth/ethclient"
	"github.com/mantlenetworkio/mantle/l2geth/rlp"
	"io"
	"log"

	"math/big"
	"os"
)

type ChainStressStruct struct {
	//设置 起始执行的 wallet address No
	RunStart int
	//设置 执行的 wallet address, 也就是 协程数量
	RunNum int
	//设置 单个 wallet address 请求交易的次数, 也就是 单个协程发起的请求总数
	RunTimes int
	//设置 wallet address's csv
	FilePath string
	//设置 链接的rpc url
	RpcUrl string

	ToAddress string

	TokenAddress string
	LogName      string

	L2Client  *ethclient.Client
	L2ChainID *big.Int

	L2BitAddress    string
	L2BridgeAddress string
}

type hash [32]byte

func InitLog(fileStr string) {
	file := "./logs/" + fileStr + ".log"
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		log.Fatal(err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter) // 将⽂件设置为log输出的⽂件
	log.SetPrefix("[stress]")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return
}

func GetHash(client *ethclient.Client, hash common.Hash) {
	receipt, err := client.TransactionReceipt(context.Background(), hash)
	if err != nil {
		log.Printf("GetHash TransactionReceipt err:%s,\n", err)
	} else {
		blockByHash, err := client.BlockByHash(context.Background(), receipt.BlockHash)
		if err != nil {
			log.Printf("GetHash BlockByHash err:%s,\n", err)
		} else {
			log.Printf("GetHash blockHash:%s,blockNumber: %v,\n", receipt.BlockHash, blockByHash.Number())
		}
	}

}

func TransferBit(client *ethclient.Client, fromPrivateKey *ecdsa.PrivateKey, toAddress common.Address, chainID *big.Int, nonce uint64) error {
	n := big.NewInt(1)
	n.SetString("1000", 10)

	gasLimit := uint64(21000) // in units
	gasPrice := big.NewInt(1)

	var data []byte
	tx := types.NewTransaction(nonce, toAddress, n, gasLimit, gasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), fromPrivateKey)
	if err != nil {
		return err
	}

	//生成裸交易
	ts := types.Transactions{signedTx}
	rawTxBytes := bytes.NewBuffer([]byte{})

	rawTxBytes.Write(ts.GetRlp(0))
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
