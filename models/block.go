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
	Previous  [32]byte
	Hash      [32]byte
	Data      Transaction
	Nounce    int
	Signature string
}

//Block Constructor
func (block *Block) createBlock(hash [32]byte, data Transaction, nounce int) {
	block.Hash = hash
	block.Nounce = nounce
	block.Data = data
}

//Add Block Previous when put into Blockchain
func (block *Block) addPrevious(previous [32]byte) {
	(*block).Previous = previous
}

//Check if Block is a Valid one
func isValidBlock(hash [32]byte) bool {
	if hash[0] == 0 && hash[1] == 0 {
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
