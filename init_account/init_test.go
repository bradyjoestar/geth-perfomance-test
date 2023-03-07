package main

import (
	"fmt"
	"geth-performance-test/deploy"
	"geth-performance-test/utils"
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

func TestInitBitAccount(t *testing.T) {
	fmt.Println("init accounts")

	l2url := "http://localhost:8545"
	l2ChainID := big.NewInt(5003)
	l2client, err := utils.GetRpcClient(l2url, l2ChainID)
	l2BridgeAddress := "0x4200000000000000000000000000000000000010"
	require.NoError(t, err, "l2client error")
	userAddress := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"
	UserPrivateKey := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

	//distribute bit
	utils.DistributeBit(l2client, userAddress, UserPrivateKey, t)

	l1erc20Address := "0xDc64a140Aa3E981100a9becA4E685f962f0cF6C9"
	l2erc20Address, err := deploy.DeployL2CustomERC20(l2client, UserPrivateKey, l2ChainID, l2BridgeAddress, l1erc20Address)

	fmt.Println(l2erc20Address)
}
