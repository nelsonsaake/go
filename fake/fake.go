package fake

import "github.com/brianvoe/gofakeit/v6"

func init() {
	gofakeit.Seed(0)
}

func Word() string {
	return gofakeit.Word()
}
