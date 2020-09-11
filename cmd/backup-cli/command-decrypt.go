package main

import (
	"encoding/base64"
	"flag"
	"github.com/jixunmoe-go/backups/utils/crypto"
	"os"
)

func commandDecrypt(argv []string) int {
	command := flag.NewFlagSet("decrypt", flag.ExitOnError)

	var privateKey string
	command.StringVar(&privateKey, "privkey", "", "Private key in base64")
	if err := command.Parse(argv); err != nil {
		println("err: could not parse args: " + err.Error())
		command.PrintDefaults()
		return 2
	}

	if privateKey == "" {
		println("err: -privkey is empty")
		command.PrintDefaults()
		return 1
	}

	return decryptStdin(privateKey)
}

func decryptStdin(privateKeyStr string) int {
	if privateKeyStr == "" {
		println("err: private key is empty")
		return 1
	}

	privateKey, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil || len(privateKey) != crypto.PrivateKeySize {
		println("err: not a valid public key (not base64 or size mismatch)")
		return 2
	}

	err = crypto.DecryptStream(os.Stdin, os.Stdout, privateKey)
	if err != nil {
		println("failed to encrypt: " + err.Error())
		return 3
	}
	return 0
}

