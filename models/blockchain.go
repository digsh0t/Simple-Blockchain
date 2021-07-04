package models

import "fmt"

func AddData(blockchain []Block, data string) []Block {
	var newBlock Block
	newBlock = newBlock.Mine(data)
	if len(blockchain) > 0 {
		newBlock.addHead(blockchain[len(blockchain)-1].Tail)
	}
	blockchain = append(blockchain, newBlock)
	return blockchain
}

func PrintBlockChain(blockchain []Block) {
	for pos, block := range blockchain {
		fmt.Printf("========Block #%d=======\n", pos)
		fmt.Printf("<====Head: %x\n", block.Head)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("====>Tail: %x\n", block.Tail)
	}
}
