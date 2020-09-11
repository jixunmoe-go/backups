package main

import (
	"encoding/base64"
	"flag"
	"github.com/jixunmoe-go/backups/utils/crypto"
	"os"
)

func commandEncrypt(argv []string) int {
	command := flag.NewFlagSet("encrypt", flag.ExitOnError)

	var publicKey string
	command.StringVar(&publicKey, "pubkey", "", "Public key in base64")
	if err := command.Parse(argv); err != nil {
		println("err: could not parse args: " + err.Error())
		command.PrintDefaults()
		return 2
	}

	if publicKey == "" {
		println("err: -pubkey is empty")
		command.PrintDefaults()
		return 1
	}

	return encryptStdin(publicKey)
}

func encryptStdin(pubkeyStr string) int {
	if pubkeyStr == "" {
		println("err: public key is empty")
		return 1
	}

	publicKey, err := base64.StdEncoding.DecodeString(pubkeyStr)
	if err != nil || len(publicKey) != crypto.PublicKeySize {
		println("err: not a valid public key (not base64 or size mismatch)")
		return 2
	}

	err = crypto.EncryptStream(os.Stdin, os.Stdout, publicKey)
	if err != nil {
		println("failed to encrypt: " + err.Error())
		return 3
	}
	return 0
}
