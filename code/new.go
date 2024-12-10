package code

import (
	"crypto/rand"
	"io"
)

var charset = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func New(otpLen int) string {
	b := make([]byte, otpLen)
	n, err := io.ReadAtLeast(rand.Reader, b, otpLen)
	if n != otpLen {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = charset[int(b[i])%len(charset)]
	}
	return string(b)
}
