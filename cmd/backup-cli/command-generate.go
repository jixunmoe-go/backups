package main

import (
	"encoding/base64"
	"fmt"
	"github.com/jixunmoe-go/backups/utils/crypto"
)

func printGenerateHelp() {
	println(appName + " gen")
	println("")
	println("Generates a private key, in base64 encoded format.")
}

func commandGenerate(_ []string) int {
	privateKey, err := crypto.GenPrivateKey()
	if err != nil {
		println("err: could not generate key: " + err.Error())
	}

	fmt.Println(base64.StdEncoding.EncodeToString(privateKey))

	return 0
}
