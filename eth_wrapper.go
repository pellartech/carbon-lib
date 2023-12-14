package carbonlib

import (
	"crypto/ecdsa"
	"encoding/hex"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func CreateAccount() (pubKey string, privKey string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privKey = hexutil.Encode(privateKeyBytes)[2:]
	log.Println(privKey)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	pubKey = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return pubKey, privKey, nil
}

func SignMessage(privKey string, message string) (signature string, err error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	sig, err := crypto.Sign(crypto.Keccak256([]byte(message)), privateKey)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(sig), nil
}

func PublicKeyFromPrivateKey(privKey string) (pubKey string, err error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	pubKey = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return pubKey, nil
}
