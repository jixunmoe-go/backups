package main

import (
	"encoding/base64"
	"github.com/jixunmoe-go/backups/utils/crypto"
	"github.com/jixunmoe-go/backups/utils/stream"
	"io"
)

func printDecryptHelp() {
	println(appName + " decrypt <privkey> [input] [output]")
	println("")
	println("Decrypts content passed in stdin. Header & checksum will be verified.")
	println("When the checksum verification failed, the program will return a non-zero code.")
	println("Omit [input] or provide '-' to use stdin (default).")
	println("Omit [output] or provide '-' to use stdout (default).")
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

	input, output, err := stream.ParseStream(argv[1:])
	if err != nil {
		println("i/o error")
		return 1
	}

	// Don't care errors here.
	defer input.Close()
	defer output.Close()

	return decryptStdin(privateKey, input, output)
}

func decryptStdin(privateKeyStr string, input io.Reader, output io.Writer) int {
	privateKey, err := base64.StdEncoding.DecodeString(privateKeyStr)
	if err != nil || len(privateKey) != crypto.PrivateKeySize {
		println("err: not a valid private key (not base64 or size mismatch)")
		return 2
	}

	err = crypto.DecryptStream(input, output, privateKey)
	if err != nil {
		println("failed to decrypt: " + err.Error())
		return 3
	}
	return 0
}
