package carbonlib

import (
	"log"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	publicKey, privateKey, err := CreateAccount()
	if err != nil {
		t.Fatalf("CreateAccount failed: %v", err)
	}

	if len(privateKey) == 0 || len(publicKey) == 0 {
		t.Error("Expected non-empty public and private keys")
	}
}

func TestSignMessage(t *testing.T) {
	pubKey, privateKey, err := CreateAccount()
	log.Println(pubKey)
	if err != nil {
		t.Fatalf("CreateAccount failed: %v", err)
	}

	message := "Test Message"
	signature, err := SignMessage(privateKey, message)
	if err != nil {
		t.Fatalf("SignMessage failed: %v", err)
	}

	if len(signature) == 0 {
		t.Error("Expected a non-empty signature")
	}
}
