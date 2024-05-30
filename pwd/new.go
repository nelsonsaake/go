package pwd

import (
	"fmt"
	"math/rand"
	"strings"
	"unicode"

	"github.com/twinj/uuid"
)

func randletters(count int) (letters string) {

	alpha := "QWERTYUIOPASDFGHJKLZXCVBNMqwertyuiopasdfghjklzxcvbnm"
	for i := 0; i < count; i++ {
		letters += string(alpha[rand.Intn(len(alpha))])
	}
	return
}

func randPunts(count int) (letters string) {

	puncts := "./#$%&*(!"
	for i := 0; i < count; i++ {
		letters += string(puncts[rand.Intn(len(puncts))])
	}
	return
}

func addCaps(str string) string {
	return str + strings.ToUpper(randletters(2))
}

func addPunts(str string) string {
	return str + strings.ToUpper(randPunts(2))
}

func addDigits(str string) string {
	return str + strings.ToUpper(fmt.Sprintf("%d%d", rand.Intn(9), rand.Intn(9)))
}

func washpassword(old string) (new string) {

	for _, ch := range old {

		if !(unicode.IsDigit(ch) || unicode.IsLetter(ch)) {
			ch = rune(randletters(2)[0])
		}

		new += string(ch)
	}

	return
}

func New() (pwd string) {
	pwd = uuid.NewV4().String()
	pwd = washpassword(pwd)
	pwd = addCaps(pwd)
	pwd = addPunts(pwd)
	pwd = addDigits(pwd)
	return
}
