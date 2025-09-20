package crypt

import (
	"fmt"

	"github.com/nelsonsaake/go/strs"
)

var secret string

func getSecret() (string, error) {
	if strs.IsEmpty(secret) {
		return "", fmt.Errorf("HASH_SECRET is not set in the environment")
	}
	return secret, nil
}
