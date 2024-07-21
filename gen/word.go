package gen

import "github.com/brianvoe/gofakeit/v6"

func init() {
	gofakeit.Seed(0)
}

func Word() string {
	return gofakeit.Word()
}

func Sentence(wordCount ...int) string {
	if len(wordCount) == 0 {
		return gofakeit.Sentence(9)
	}
	return gofakeit.Sentence(wordCount[0])
}
