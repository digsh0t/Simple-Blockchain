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
	sig := trans.SignRSASHA256(*utils.LoadRSAPrivatePemKey("/home/wintltr/Projects/test-rsa-signing/alicepriv.pem"))
	fmt.Println(trans.VerifyTransaction(*utils.LoadPublicPemKey("/home/wintltr/Projects/test-rsa-signing/alicepub.pem"), sig))
}
