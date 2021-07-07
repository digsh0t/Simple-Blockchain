package main

import (
	"fmt"

	"github.com/wintltr/simple-blockchain/models"
	"github.com/wintltr/simple-blockchain/utils"
)

//Run demonstration
func main() {
	var trans models.Transaction
	trans.MakeTransaction("Tri", "Long", 15)
	sig := trans.SignRSASHA256(*utils.ReadRSAKeyFromFile("/home/wintltr/certificates/Blockchain-RSA-Keys/alicepriv.pem"))
	fmt.Println(trans.VerifyTransaction(*utils.ReadRSAPublicKeyFromFile("/home/wintltr/certificates/Blockchain-RSA-Keys/alicepub.pem"), sig))
}
