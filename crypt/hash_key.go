package crypt

import (
	"fmt"
	"os"

	"github.com/nelsonsaake/go/str"
)

func getSecret() (string, error) {
	secret := os.Getenv("HASH_SECRET")
	if str.IsEmpty(secret) {
		return "", fmt.Errorf("HASH_SECRET is not set in the environment")
	}
	return secret, nil
}
