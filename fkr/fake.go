package fkr

import (
	"math/rand/v2"

	"github.com/brianvoe/gofakeit/v6"
)

func init() {
	gofakeit.Seed(0)
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
