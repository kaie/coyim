package filetransfer

import (
	"crypto/rand"
	"crypto/sha256"
)

type encryptionParameters struct {
	keyType       string
	key           []byte
	iv            []byte
	encryptionKey []byte
	macKey        []byte
}

var encryptionKeyPrefix = []byte("xmpp-encryption-key")
var macKeyPrefix = []byte("xmpp-mac-key")

func deriveKeyWithPrefix(prefix, key []byte, l int) []byte {
	total := append(prefix, key...)
	result := sha256.Sum256(total)
	return result[0:l]
}

func generateSafeRandomBytes(l int) []byte {
	b := make([]byte, l)
	rand.Read(b)
	return b
}

func generateEncryptionParameters(enabled bool, genKey func() []byte, keyType string) *encryptionParameters {
	if !enabled {
		return nil
	}
	key := genKey()
	if len(key) == 256 {
		return &encryptionParameters{
			keyType:       keyType,
			key:           key,
			iv:            generateSafeRandomBytes(128),
			encryptionKey: deriveKeyWithPrefix(encryptionKeyPrefix, key, 128),
			macKey:        deriveKeyWithPrefix(macKeyPrefix, key, 256),
		}
	}
	return nil
}
