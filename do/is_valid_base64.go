package do

import (
	"encoding/base64"
)

// IsValidBase64 checks if a string is a valid Base64 encoded string
func IsValidBase64(s string) bool {
	_, err := base64.StdEncoding.DecodeString(s)
	return err == nil
}
