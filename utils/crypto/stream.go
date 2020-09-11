package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

var encryptHeader = []byte{'E', '!', 'J', 'X'}

// EncryptStream encrypts a given stream with ECC key (curve25519 public key) and AES.
func EncryptStream(input io.Reader, output io.Writer, publicKey PublicKey) error {
	priv2, pub2, err := GenKeyPair()
	if err != nil {
		return fmt.Errorf("could not generate key pair: %w", err)
	}

	aesKey, err := DeriveEncryptionKey(priv2, publicKey)
	if err != nil {
		return fmt.Errorf("could not derive aes key: %w", err)
	}

	aesBlock, err := aes.NewCipher(aesKey)
	if err != nil {
		return fmt.Errorf("could not init aes: %w", err)
	}

	iv := make([]byte, aesBlock.BlockSize())
	_, _ = rand.Read(iv)
	_, _ = output.Write(encryptHeader)
	_, _ = output.Write(pub2)
	_, _ = output.Write(iv)
	stream := cipher.NewOFB(aesBlock, iv)
	writer := &cipher.StreamWriter{S: stream, W: output}
	if _, err := io.Copy(writer, input); err != nil {
		return fmt.Errorf("failed to perform aes encrypt: %w", err)
	}

	return writer.Close()
}

func DecryptStream(input io.Reader, output io.Writer, privateKey PrivateKey) error {
	header := make([]byte, len(encryptHeader))
	n, err := input.Read(header)
	if err != nil || n != len(encryptHeader) || !bytes.Equal(header, encryptHeader) {
		return fmt.Errorf("input does not contain a valid header")
	}

	publicKey := make([]byte, PublicKeySize)
	n, err = input.Read(publicKey)
	if err != nil || n != PublicKeySize {
		return fmt.Errorf("could not read public key")
	}

	aesKey, err := DeriveEncryptionKey(privateKey, publicKey)
	if err != nil {
		return fmt.Errorf("could not derive aes key: %w", err)
	}

	aesBlock, err := aes.NewCipher(aesKey)
	if err != nil {
		return fmt.Errorf("could not init aes: %w", err)
	}

	iv := make([]byte, aesBlock.BlockSize())
	n, err = input.Read(iv)
	if err != nil || n != len(iv) {
		return fmt.Errorf("could not read iv")
	}

	stream := cipher.NewOFB(aesBlock, iv)
	writer := &cipher.StreamWriter{S: stream, W: output}
	if _, err := io.Copy(writer, input); err != nil {
		return fmt.Errorf("failed to perform aes decrypt: %w", err)
	}

	return writer.Close()
}
