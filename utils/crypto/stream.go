package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"github.com/jixunmoe-go/backups/utils/checksum"
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
	verifiedOutput := checksum.NewWriter(output)
	_, _ = rand.Read(iv)
	_, _ = verifiedOutput.Write(encryptHeader)
	_, _ = verifiedOutput.Write(pub2)
	_, _ = verifiedOutput.Write(iv)
	stream := cipher.NewOFB(aesBlock, iv)
	writer := &cipher.StreamWriter{S: stream, W: verifiedOutput}
	if _, err := io.Copy(writer, input); err != nil {
		return fmt.Errorf("failed to perform aes encrypt: %w", err)
	}

	_, _ = verifiedOutput.WriteChecksum()

	return writer.Close()
}

func DecryptStream(input io.Reader, output io.Writer, privateKey PrivateKey) error {
	verifyInput := checksum.NewReader(input)

	header := make([]byte, len(encryptHeader))
	n, err := verifyInput.Read(header)
	if err != nil || n != len(encryptHeader) || !bytes.Equal(header, encryptHeader) {
		return fmt.Errorf("input does not contain a valid header")
	}

	publicKey := make([]byte, PublicKeySize)
	n, err = verifyInput.Read(publicKey)
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
	n, err = verifyInput.Read(iv)
	if err != nil || n != len(iv) {
		return fmt.Errorf("could not read iv")
	}

	stream := cipher.NewOFB(aesBlock, iv)
	writer := &cipher.StreamWriter{S: stream, W: output}
	if _, err := io.Copy(writer, verifyInput); err != nil {
		return fmt.Errorf("failed to perform aes decrypt: %w", err)
	}

	_ = writer.Close()

	if !verifyInput.Verify() {
		return fmt.Errorf("could not verify the data read")
	}

	return nil
}
