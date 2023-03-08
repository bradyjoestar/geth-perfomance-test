package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/csv"
	"flag"
	"geth-performance-test/jobs"
	"github.com/mantlenetworkio/mantle/l2geth/crypto"
	"github.com/mantlenetworkio/mantle/l2geth/ethclient"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"sync"

	"time"
)

func main() {
	s := new(http.Server)
	// 增加服务器读写超时时间
	s.ReadTimeout = 10 * time.Hour
	s.WriteTimeout = 10 * time.Hour
	logFileName := "log" + time.Now().String()
	filePath := flag.String("f", "./secrect.csv", "设置 wallet address's csv")
	logName := flag.String("l", logFileName, "生成log文件名称")

	flag.Parse()
	SecondSendNum := 200
	AllSendNum := 10000

	l2url := "http://localhost:8545"
	l2chainID := big.NewInt(17)

	//初始化log
	jobs.InitLog(*logName)
	starTime := time.Now()

	//文件读取
	f, err1 := os.Open(*filePath)
	if err1 != nil {
		log.Fatal("os.Open", err1)
	}
	reader := csv.NewReader(f)
	accountData, err2 := reader.ReadAll()
	if err2 != nil {
		log.Fatal("reader.ReadAll", err2)
	}
	fromAccountList := [10000]*ecdsa.PrivateKey{}
	for i, list := range accountData {
		hex := strings.TrimPrefix(list[0], "0x")
		privateKey, err := crypto.HexToECDSA(hex)
		if err != nil {
			log.Fatal("private key transform error :", err)
		}
		fromAccountList[i] = privateKey
	}
	to := crypto.PubkeyToAddress(fromAccountList[9999].PublicKey)

	//Start Test

	var wg = new(sync.WaitGroup)
	n := AllSendNum / SecondSendNum
	wg.Add(int(n) * SecondSendNum)
	l2client, err := ethclient.Dial(l2url)

	nonce, err := l2client.PendingNonceAt(context.Background(), crypto.PubkeyToAddress(fromAccountList[10].PublicKey))

	for j := 0; j < n; j++ {
		for i := j * SecondSendNum; i < (j+1)*SecondSendNum; i++ {
			go func(num int, wg *sync.WaitGroup) {
				defer wg.Done()

				client, err := ethclient.Dial(l2url)
				if err != nil {
					log.Printf("ethclient.Dial,%v\n", err)
				}
				defer client.Close()

				//query eth balance
				if fromAccountList[num] == nil {
					return
				}
				timeS := time.Now()
				err = jobs.TransferBit(client, fromAccountList[num], to, l2chainID, nonce)
				if err != nil {
					log.Printf("transfer error: %v \n", err)
					return
				}
				timeE := time.Now().UnixMilli() - timeS.UnixMilli()
				log.Printf("forTransferTime: %vms", timeE)

			}(i, wg)

		}
		time.Sleep(time.Millisecond * time.Duration(50))

	}

	wg.Wait()
	log.Printf("任务整体耗时：%v\n", time.Since(starTime))
	client, err := ethclient.Dial(l2url)
	if err != nil {
		log.Printf("ethclient.Dial,%v\n", err)
	}
	defer client.Close()
	balance, err := client.BalanceAt(context.Background(), to, nil)
	if err != nil {
		log.Printf("get balance error ,%v\n", err)

	}
	log.Printf("balance = %v", balance)

}
