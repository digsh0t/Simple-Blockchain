package utils

import (
	"bufio"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func LoadRSAPrivatePemKey(fileName string) *rsa.PrivateKey {
	privateKeyFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}

	pemfileinfo, _ := privateKeyFile.Stat()
	var size int64 = pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(privateKeyFile)
	_, err = buffer.Read(pembytes)

	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	data, _ := pem.Decode([]byte(pembytes))
	privateKeyFile.Close()
	privateKeyImported, err := x509.ParsePKCS1PrivateKey(data.Bytes)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	return privateKeyImported
}

func LoadPublicPemKey(fileName string) *rsa.PublicKey {

	publicKeyFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}

	pemfileinfo, _ := publicKeyFile.Stat()

	size := pemfileinfo.Size()
	pembytes := make([]byte, size)
	buffer := bufio.NewReader(publicKeyFile)
	_, err = buffer.Read(pembytes)
	if err != nil {
		fmt.Println(err)
	}
	data, _ := pem.Decode([]byte(pembytes))
	publicKeyFile.Close()
	publicKeyFileImported, err := x509.ParsePKCS1PublicKey(data.Bytes)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
	return publicKeyFileImported
}
