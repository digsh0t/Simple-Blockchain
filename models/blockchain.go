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
		newBlock.addHead((*blockchain)[len(*blockchain)-1].Tail)
	}
	*blockchain = append((*blockchain), newBlock)
}

//Print Blockchain
func (blockchain *Blockchain) PrintBlockChain() {
	for pos, block := range *blockchain {
		fmt.Printf("\n========Block #%d=======\n", pos)
		fmt.Printf("<====Head: %x\n", block.Head)
		fmt.Printf("Sender: %s\n", block.Data.Sender)
		fmt.Printf("Receiver: %s\n", block.Data.Receiver)
		fmt.Printf("Amount: %f\n", block.Data.Amount)
		fmt.Printf("Signature: %s\n", block.Signature)
		fmt.Printf("====>Tail: %x\n", block.Tail)
	}
}

func (blockchain *Blockchain) AddNewTransaction(sender string, receiver string, amount float32, privateKey *rsa.PrivateKey) {
	var newBlock Block
	newBlock.Data.InitTransaction(sender, receiver, amount)
	newBlock.Signature = newBlock.Data.SignRSASHA256(*privateKey)
	blockchain.AddBlock(newBlock)
}
