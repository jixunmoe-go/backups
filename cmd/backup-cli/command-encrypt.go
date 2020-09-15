package main

import (
	"encoding/base64"
	"github.com/jixunmoe-go/backups/utils/crypto"
	"github.com/jixunmoe-go/backups/utils/stream"
	"os"
)

func printEncryptHelp() {
	println(appName + " encrypt <pubkey> [input] [output]")
	println("")
	println("Encrypts content passed in stdin.")
	println("AES encryption key will be derived from the public key.")
	println("Omit [input] or provide '-' to use stdin (default).")
	println("Omit [output] or provide '-' to use stdout (default).")
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

	input, output, err := stream.ParseStream(argv[1:])
	if err != nil {
		println("i/o error")
		return 1
	}

	// Don't care errors here.
	defer input.Close()
	defer output.Close()

	return encryptStdin(publicKey, input, output)
}

func encryptStdin(pubkeyStr string, input *os.File, output *os.File) int {
	publicKey, err := base64.StdEncoding.DecodeString(pubkeyStr)
	if err != nil || len(publicKey) != crypto.PublicKeySize {
		println("err: not a valid public key (not base64 or size mismatch)")
		return 2
	}

	err = crypto.EncryptStream(input, output, publicKey)
	if err != nil {
		println("failed to encrypt: " + err.Error())
		return 3
	}

	return 0
}
