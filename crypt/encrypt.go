package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func encrypt(data string, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	nonce := make([]byte, 12) // AES-GCM nonce size
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	encryptedBytes := aesGCM.Seal(nil, nonce, []byte(data), nil)

	// Combine nonce and encrypted data, and encode to base64
	return base64.StdEncoding.EncodeToString(append(nonce, encryptedBytes...)), nil
}

func Encrypt(data string) (string, error) {

	secret, err := getSecret()
	if err != nil {
		return "", err
	}

	return encrypt(data, secret)
}
