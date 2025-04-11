package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func Hash(data string) (string, error) {

	secret, err := getSecret()
	if err != nil {
		return "", err
	}

	h := hmac.New(sha256.New, []byte(secret))

	h.Write([]byte(data))

	hashedBytes := h.Sum(nil)

	hashedData := fmt.Sprintf("%x", hashedBytes)

	return hashedData, nil
}
