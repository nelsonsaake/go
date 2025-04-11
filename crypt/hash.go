package crypt

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"server/src/env"
)

func Hash(data string) string {

	h := hmac.New(sha256.New, []byte(env.HashSecret()))

	h.Write([]byte(data))

	hashedBytes := h.Sum(nil)

	return fmt.Sprintf("%x", hashedBytes)
}
