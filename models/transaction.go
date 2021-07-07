package models

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Transaction struct {
	Sender   string  `json:"sender"`
	Receiver string  `json:"receiver"`
	Amount   float32 `json:"amount"`
}

func (trans *Transaction) InitTransaction(sender string, receiver string, amount float32) {
	trans.Sender = sender
	trans.Receiver = receiver
	trans.Amount = amount
}

func (trans *Transaction) HashSHA256() [32]byte {
	message, err := json.Marshal(trans)
	if err != nil {
		log.Fatalf("Fail to marshal transaction with error: %s", err)
	}
	return sha256.Sum256([]byte(message))
}

func (trans *Transaction) SignRSASHA256(privKey rsa.PrivateKey) string {
	hashed := trans.HashSHA256()
	rng := rand.Reader
	signature, err := rsa.SignPKCS1v15(rng, &privKey, crypto.SHA256, hashed[:])
	sig := base64.StdEncoding.EncodeToString(signature)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from signing: %s\n", err)
	}
	return sig
}
