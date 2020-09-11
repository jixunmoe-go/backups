package crypto

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/curve25519"
)

type PrivateKey = []byte
type PublicKey = []byte

const PublicKeySize = curve25519.PointSize
const PrivateKeySize = curve25519.ScalarSize

// GenKeyPair generates a pair of keys that is suitable for
func GenKeyPair() (privateKey PrivateKey, publicKey PublicKey, err error) {
	privateKey, err = GenPrivateKey()
	if err != nil {
		return nil, nil, err
	}

	publicKey, err = GetPublicKey(privateKey)
	return privateKey, publicKey, err
}

// GenPrivateKey generates a private key.
func GenPrivateKey() (privateKey PrivateKey, err error) {
	privateKey = make([]byte, PrivateKeySize)
	_, err = rand.Read(privateKey)
	if err != nil {
		return nil, fmt.Errorf("could not generate key: %w", err)
	}

	return privateKey, nil
}

func GetPublicKey(privateKey PrivateKey) (publicKey PublicKey, err error) {
	publicKey, err = curve25519.X25519(privateKey, curve25519.Basepoint)
	if err != nil {
		return nil, fmt.Errorf("could not derive public key: %w", err)
	}

	return publicKey, nil
}

// DeriveEncryptionKey derives a "shared key" from two key pair (Private and Public key from different party).
// Each of those combination will be able to derive the same "shared key" suitable for encryption (e.g. AES)
// The return value should be a 32-byte (256-bit) key.
func DeriveEncryptionKey(privateKy PrivateKey, publicKey PublicKey) (encryptionKey []byte, err error) {
	return curve25519.X25519(privateKy, publicKey)
}
