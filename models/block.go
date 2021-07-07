package models

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

//Create Block Struct
type Block struct {
	Head      [32]byte
	Tail      [32]byte
	Data      Transaction
	Nounce    int
	Signature string
}

//Block Constructor
func (block *Block) createBlock(tail [32]byte, data Transaction, nounce int) {
	block.Tail = tail
	block.Nounce = nounce
	block.Data = data
}

//Add Block Head when put into Blockchain
func (block *Block) addHead(head [32]byte) {
	(*block).Head = head
}

//Check if Block is a Valid one
func isValidBlock(tail [32]byte) bool {
	if tail[0] == 0 && tail[1] == 0 {
		return true
	} else {
		return false
	}
}

//Find the right nounce that makes isValidBlock return true
func (newBlock *Block) Mine() {
	nounce := 0
	trans, err := json.Marshal(newBlock.Data)
	if err != nil {
		fmt.Printf("Error when marshal data: %s", err)
	}
	for {
		computedHash := sha256.Sum256([]byte(string(trans) + strconv.Itoa(nounce)))
		if isValidBlock(computedHash) {
			newBlock.createBlock(computedHash, newBlock.Data, nounce)
			return
		}
		nounce++
	}
}

//Verify if block match signature
func (block *Block) VerifyBlock(publicKey rsa.PublicKey) string {
	sig, _ := base64.StdEncoding.DecodeString(block.Signature)
	message, err := json.Marshal(block.Data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from marshalling message: %s\n", err)
	}
	hashed := sha256.Sum256([]byte(message))
	err = rsa.VerifyPKCS1v15(&publicKey, crypto.SHA256, hashed[:], sig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from verification: %s\n", err)
		return "Error from verification:"
	}
	return "Signature Verification Passed"
}
