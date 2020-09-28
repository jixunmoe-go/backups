package main

import (
	"encoding/base64"
	"github.com/jixunmoe-go/backups/utils/crypto"
	"github.com/jixunmoe-go/backups/utils/memory"
	"github.com/jixunmoe-go/backups/utils/stream"
	"io"
	"io/ioutil"
)

func printDecryptHelp() {
	println(appName + " decrypt <path to privkey> [input] [output]")
	println("")
	println("Decrypts content passed in stdin. Header & checksum will be verified.")
	println("When the checksum verification failed, the program will return a non-zero code.")
	println("Omit [input] or provide '-' to use stdin (default).")
	println("Omit [output] or provide '-' to use stdout (default).")
	println("")
	println("e.g.")
	println("  " + appName + " decrypt /path/to/private.key crypted.tar.gz.bin decrypted.tar.gz")
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

	privateKeyB64, err := ioutil.ReadFile(privateKey)
	if err != nil {
		println("could not read private key")
		return 1
	}
	return decryptStdin(privateKeyB64, input, output)
}

func decryptStdin(privateKeyB64 []byte, input io.Reader, output io.Writer) int {
	privateKey := make([]byte, crypto.PrivateKeySize)
	n, err := base64.StdEncoding.Decode(privateKey, privateKeyB64)
	memory.SetZero(privateKeyB64)
	if err != nil || n != crypto.PrivateKeySize {
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
