package main

import (
	"encoding/base64"
	"fmt"
	"github.com/jixunmoe-go/backups/utils/crypto"
	"io/ioutil"
	"os"
)

func printPubKeyHelp() {
	println(appName + " pubkey")
	println("")
	println("Derives corresponding public key for a given private key piped in.")
	println("")
	println("e.g.")
	println("  cat private.key | " + appName + " pubkey > public.key")
}

func commandPubKey(_ []string) int {
	privKeyStr, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		println("could not get private key from stdin: " + err.Error())
	}

	privateKey, err := base64.StdEncoding.DecodeString(string(privKeyStr))
	if err != nil {
		println("err: could not parse private key: " + err.Error())
	}

	publicKey, err := crypto.GetPublicKey(privateKey)
	fmt.Println(base64.StdEncoding.EncodeToString(publicKey))
	return 0
}
