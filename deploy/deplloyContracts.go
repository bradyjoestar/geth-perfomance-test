package deploy

import (
	"context"
	"crypto/ecdsa"
	"geth-performance-test/contracts/token"
	"geth-performance-test/contracts/token/l1"
	"geth-performance-test/contracts/token/l2"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

/**
deployErc20
*/
func DeployErc20(client *ethclient.Client, userPrivateKey string, chainID *big.Int, tokenName string, sym string) (string, error) {
	privateKey, err := crypto.HexToECDSA(userPrivateKey)
	if err != nil {
		log.Fatalln(err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalln(err.Error())
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err == nil {
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)      // in wei
		auth.GasLimit = uint64(3000000) // in units
		auth.GasPrice = gasPrice
		_initialAmount := big.NewInt(1000000000000000000) //4
		address, tx, instance, err := token.DeployToken(auth, client, _initialAmount, tokenName, 18, sym)
		if err != nil {
			log.Fatalln(err.Error())
			return "", err
		}
		log.Printf("tx.Hash:%s", tx.Hash().Hex())
		_ = instance
		return address.Hex(), nil
	} else {
		log.Fatalln("err", err)
		return "", err
	}
}

func DeployL1CustomERC20(client *ethclient.Client, userPrivateKey string, chainID *big.Int) (string, error) {
	privateKey, err := crypto.HexToECDSA(userPrivateKey)
	if err != nil {
		log.Fatalln(err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalln(err.Error())
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err == nil {
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)      // in wei
		auth.GasLimit = uint64(3000000) // in units
		auth.GasPrice = gasPrice
		address, tx, instance, err := l1.DeployL1(auth, client)
		if err != nil {
			log.Fatalln(err.Error())
			return "", err
		}
		log.Printf("tx.Hash:%s", tx.Hash().Hex())
		_ = instance
		return address.Hex(), nil
	} else {
		log.Fatalln("err", err)
		return "", err
	}
}

func DeployL2CustomERC20(client *ethclient.Client, userPrivateKey string, chainID *big.Int, l2bridgeAddress string, l1tokenAddress string) (string, error) {
	privateKey, err := crypto.HexToECDSA(userPrivateKey)
	if err != nil {
		log.Fatalln(err.Error())
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalln("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalln(err.Error())
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalln(err.Error())
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err == nil {
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)      // in wei
		auth.GasLimit = uint64(3000000) // in units
		auth.GasPrice = gasPrice
		address, tx, instance, err := l2.DeployL2(auth, client, common.HexToAddress(l2bridgeAddress), common.HexToAddress(l1tokenAddress))
		if err != nil {
			log.Fatalln(err.Error())
			return "", err
		}
		log.Printf("tx.Hash:%s", tx.Hash().Hex())
		_ = instance
		return address.Hex(), nil
	} else {
		log.Fatalln("err", err)
		return "", err
	}
}
