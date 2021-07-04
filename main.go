package main

import (
	"github.com/wintltr/simple-blockchain/models"
)

func main() {
	var blockchain []models.Block
	blockchain = models.AddData(blockchain, "test")
	blockchain = models.AddData(blockchain, "123")
	models.PrintBlockChain(blockchain)
}
