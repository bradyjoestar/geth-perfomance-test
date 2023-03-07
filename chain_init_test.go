package main

import (
	"fmt"
	"geth-performance-test/utils"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestInitBitAccount(t *testing.T) {
	fmt.Println("init accounts")

	l2url := "http://localhost:8545"
	l2ChainID := big.NewInt(17)
	l2client, err := utils.GetRpcClient(l2url, l2ChainID)
	require.NoError(t, err, "l2client error")
	userAddress := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	UserPrivateKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

	//distribute bit
	utils.DistributeBit(l2client, userAddress, UserPrivateKey, t)
}
