package gen

import (
	"crypto/rand"
	"io"
)

var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}

func Code(codelen int) string {
	b := make([]byte, codelen)
	n, err := io.ReadAtLeast(rand.Reader, b, codelen)
	if n != codelen {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var alphanumtable = [...]byte{
	'1', '2', '3', '4', '5', '6', '7', '8', '9', '0',
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z',
}

func AlphaNum(codelen int) string {
	table := alphanumtable
	b := make([]byte, codelen)
	n, err := io.ReadAtLeast(rand.Reader, b, codelen)
	if n != codelen {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}

var alphtable = [...]byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
	'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T',
	'U', 'V', 'W', 'X', 'Y', 'Z',
}

func Alpha(codelen int) string {
	table := alphtable
	b := make([]byte, codelen)
	n, err := io.ReadAtLeast(rand.Reader, b, codelen)
	if n != codelen {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
