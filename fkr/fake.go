package fkr

import (
	"math/rand/v2"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/nelsonsaake/go/strs"
)

func init() {
	gofakeit.Seed(0)
}

// todo: break the dependency on strs package
func UUID() string {
	return strs.UUID()
}

func Word() string {
	return gofakeit.Word()
}

func Name() string {
	return gofakeit.Name()
}

func Bool() bool {
	return rand.Float32() >= 0.5
}

func Sentence() string {
	return gofakeit.Sentence(10)
}

func IPv4() string {
	return gofakeit.IPv4Address()
}

func PhoneNumber() string {
	// Ghanaian mobile prefixes
	prefixes := []string{"024", "054", "055", "056", "057", "059", "020", "050"}
	prefix := prefixes[rand.IntN(len(prefixes))]
	// Generate 7 random digits
	num := ""
	for i := 0; i < 7; i++ {
		num += gofakeit.Digit()
	}
	return "+233" + prefix[1:] + num
}

func Email() string {
	return gofakeit.Email()
}
