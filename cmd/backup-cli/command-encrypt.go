package main

import (
	"encoding/base64"
	"github.com/jixunmoe-go/backups/utils/crypto"
	"os"
)

func printEncryptHelp() {
	println(appName + " encrypt <pubkey>")
	println("")
	println("Encrypts content passed in stdin.")
	println("AES encryption key will be derived from the public key.")
	println("")
	println("e.g.")
	println("  " + appName + " encrypt \"$(cat public.key)\"")
}

func commandEncrypt(argv []string) int {
	if len(argv) < 1 {
		println("Need to specify public key.")
		return 1
	}

	publicKey := argv[0]

	if publicKey == "" {
		println("err: pubkey is empty")
		return 1
	}

	return encryptStdin(publicKey)
}

func encryptStdin(pubkeyStr string) int {
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
