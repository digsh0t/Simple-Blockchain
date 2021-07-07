package models

import (
	"crypto/rsa"
	"fmt"
)

type Blockchain []Block

//Append Block to Blockchain
func (blockchain *Blockchain) AddBlock(newBlock Block) {
	newBlock.Mine()
	if len(*blockchain) > 0 {
		newBlock.addPrevious((*blockchain)[len(*blockchain)-1].Hash)
	}
	*blockchain = append((*blockchain), newBlock)
}

//Print Blockchain
func (blockchain *Blockchain) PrintBlockChain() {
	for pos, block := range *blockchain {
		fmt.Printf("\n========Block #%d=======\n", pos)
		fmt.Printf("<====Previous: %x\n", block.Previous)
		fmt.Printf("Sender: %s\n", block.Data.Sender)
		fmt.Printf("Receiver: %s\n", block.Data.Receiver)
		fmt.Printf("Amount: %f\n", block.Data.Amount)
		fmt.Printf("Signature: %s\n", block.Signature)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}

//Create a new transaction and put into blockchain
func (blockchain *Blockchain) AddNewTransaction(sender string, receiver string, amount float32, privateKey *rsa.PrivateKey) {
	var newBlock Block
	newBlock.Data.InitTransaction(sender, receiver, amount)
	newBlock.Signature = newBlock.Data.SignRSASHA256(*privateKey)
	blockchain.AddBlock(newBlock)
}
