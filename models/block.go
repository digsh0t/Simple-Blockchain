package models

import (
	"crypto/sha256"
	"strconv"
)

//Create Block Struct
type Block struct {
	Head   [32]byte
	Tail   [32]byte
	Data   string
	Nounce int
}

//Block Constructor
func (block Block) createBlock(tail [32]byte, data string, nounce int) Block {
	block.Tail = tail
	block.Nounce = nounce
	block.Data = data
	return block
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
func (block Block) Mine(data string) Block {
	var newBlock Block
	nounce := 0
	for {
		computedHash := sha256.Sum256([]byte(data + strconv.Itoa(nounce)))
		if isValidBlock(computedHash) {
			newBlock = newBlock.createBlock(computedHash, data, nounce)
			return newBlock
		}
		nounce++
	}
}
