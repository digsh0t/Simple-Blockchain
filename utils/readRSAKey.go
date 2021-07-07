package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"strings"
)

func ReadRSAKeyFromFile(pemFile string) *rsa.PrivateKey {
	pemByte, err := ioutil.ReadFile(pemFile)
	pemString := string(pemByte)
	if err != nil {
		log.Fatalf("Fatal when open Private key file with error: %s", err)
	}

	block, _ := pem.Decode([]byte(pemString))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	//key := parseResult.(*rsa.PrivateKey)
	return key
}

func ReadRSAPublicKeyFromFile(pkPemFile string) *rsa.PublicKey {
	pkByte, err := ioutil.ReadFile(pkPemFile)
	pkString := string(pkByte)
	if err != nil {
		log.Fatalf("Fatal when open pk file with error: %s", err)
	}
	var publicKey *rsa.PublicKey
	block, _ := pem.Decode([]byte(pkString))
	if !strings.Contains(pkString, "RSA") {
		parseResult, _ := x509.ParsePKIXPublicKey(block.Bytes)
		publicKey = parseResult.(*rsa.PublicKey)
	} else {
		publicKey, _ = x509.ParsePKCS1PublicKey(block.Bytes)
	}
	return publicKey
}
