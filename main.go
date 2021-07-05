package main

import (
	"github.com/wintltr/simple-blockchain/models"
	"github.com/wintltr/simple-blockchain/utils"
)

//Run demonstration
func main() {
	var trans models.Transaction
	trans.MakeTransaction("Tri", "Long", 15)
	trans.SignRSASHA256(utils.ReadRSAKeyFromFile("123"))

}
