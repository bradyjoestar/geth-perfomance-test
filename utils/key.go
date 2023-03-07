package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
	"testing"
)

/**
 * 根据私钥获取公钥匙
 */
func GetPubKey(priKeyList []string, t *testing.T) {
	for i := 0; i < len(priKeyList); i++ {
		// 创建私钥对象，上面私钥没有钱哦
		priKey, err := crypto.HexToECDSA(priKeyList[i])
		if err != nil {
			panic(err)
		}
		priKeyBytes := crypto.FromECDSA(priKey)
		t.Logf("私钥为: %s\n", hex.EncodeToString(priKeyBytes))

		pubKey := priKey.Public().(*ecdsa.PublicKey)
		// 获取公钥并去除头部0x04
		pubKeyBytes := crypto.FromECDSAPub(pubKey)[1:]
		t.Logf("去0x04公钥为: %s\n", hex.EncodeToString(pubKeyBytes))
		compressed := crypto.CompressPubkey(pubKey)
		t.Logf("33压缩公钥为: %s\n", hex.EncodeToString(compressed))
		t.Log("--------------------------------------------------------------")

	}
}

func HexToPrivateKey(str string) (*ecdsa.PrivateKey, error) {
	hex := strings.TrimPrefix(str, "0x")
	privateKey, err := crypto.HexToECDSA(hex)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
