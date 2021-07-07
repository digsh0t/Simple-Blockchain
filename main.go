package main

import (
	"fmt"

	"github.com/wintltr/simple-blockchain/models"
	"github.com/wintltr/simple-blockchain/utils"
)

//Run demonstration
func main() {
	var blockchain models.Blockchain
	blockchain.AddNewTransaction("Tri", "Long", 15, utils.ReadRSAKeyFromFile("/home/wintltr/certificates/Blockchain-RSA-Keys/blockchain_private.pem"))
	blockchain.AddNewTransaction("Long", "Ha", 12, utils.ReadRSAKeyFromFile("/home/wintltr/certificates/Blockchain-RSA-Keys/blockchain_private.pem"))
	fmt.Println(blockchain[0].VerifyBlock(*utils.ReadRSAPublicKeyFromFile("/home/wintltr/certificates/Blockchain-RSA-Keys/blockchain_pk.pem")))
	blockchain.PrintBlockChain()
}
