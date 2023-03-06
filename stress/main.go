package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	fmt.Println("stress test")
	pricer := ethereum.CallMsg{
		From:       common.Address{},
		To:         nil,
		Gas:        0,
		GasPrice:   nil,
		GasFeeCap:  nil,
		GasTipCap:  nil,
		Value:      nil,
		Data:       nil,
		AccessList: nil,
	}
	fmt.Println(pricer)
}

func GetRpcClient(rpcUrl string, chainID *big.Int) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return client, nil
}
