package models

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
)

type Transaction struct {
	Sender    string  `json:"sender"`
	Receiver  string  `json:"receiver"`
	Amount    float32 `json:"amount"`
	Signature []byte  `json:"signature"`
}

func (trans *Transaction) MakeTransaction(sender string, receiver string, amount float32) {
	trans.Sender = sender
	trans.Receiver = receiver
	trans.Amount = amount
}

func (trans *Transaction) SignRSASHA256(privateKey rsa.PrivateKey) {
	message, err := json.Marshal(trans)
	if err != nil {
		log.Fatalf("Fail to marshal transaction with error: %s", err)
	}
	hashed := sha256.Sum256([]byte(message))
	trans.Signature, err = rsa.SignPKCS1v15(rand.Reader, &privateKey, crypto.SHA256, hashed[:])
	if err != nil {
		fmt.Printf("Error while signing transaction: %s", err)
	}
}
