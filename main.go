package main

import (
	"github.com/wintltr/simple-blockchain/models"
)

//Run demonstration
func main() {
	var blockchain []models.Block
	blockchain = models.AddData(blockchain, "test")
	blockchain = models.AddData(blockchain, "123")
	models.PrintBlockChain(blockchain)
}
