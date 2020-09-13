package main

import (
	"encoding/base64"
	"fmt"
	"github.com/jixunmoe-go/backups/utils/crypto"
	"os"
)

func printDecryptHelp() {
	println(appName + " decrypt <privkey>")
	println("")
	println("Decrypts content passed in stdin. Header & checksum will be verified.")
	println("When the checksum verification failed, the program will return a non-zero code.")
	println("")
	println("e.g.")
	println("  " + appName + " decrypt \"$(cat private.key)\"")
}

func commandDecrypt(argv []string) int {
	if len(argv) < 1 {
		println("Need to specify private key.")
		return 1
	}

	privateKey := argv[0]

	if privateKey == "" {
		println("err: private is empty")
		return 1
	}

	return decryptStdin(privateKey)
}

func decryptStdin(privateKeyStr string) int {
	privateKey, err := base64.StdEncoding.DecodeString(privateKeyStr)
	fmt.Println("---")
	fmt.Println(privateKeyStr, len(privateKey), err)
	fmt.Println("---")
	if err != nil || len(privateKey) != crypto.PrivateKeySize {
		println("err: not a valid private key (not base64 or size mismatch)")
		return 2
	}

	err = crypto.DecryptStream(os.Stdin, os.Stdout, privateKey)
	if err != nil {
		println("failed to decrypt: " + err.Error())
		return 3
	}
	return 0
}
