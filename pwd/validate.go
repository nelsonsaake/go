package pwd

import (
	"fmt"
	"strings"
	"unicode"
)

func count(str string, isT func(rune) bool) (count int) {
	for _, alpha := range str {
		if isT(alpha) {
			count++
		}
	}
	return
}

func empty(v string) bool {
	return len(strings.TrimSpace(v)) == 0
}

// ValidatePassword: Requires at least 8 characters, at least 2 caps, at least 2 digits, at least 2 puncts or symbols.
func ValidatePassword(pwd string) (err error) {

	if empty(pwd) {
		return fmt.Errorf("password is empty")
	}

	if len(pwd) < 8 {
		return fmt.Errorf("password is less than eight")
	}

	if count(pwd, unicode.IsDigit) < 2 {
		return fmt.Errorf("password must include at least two digits")
	}

	if count(pwd, unicode.IsUpper) < 2 {
		return fmt.Errorf("password must include at least two upper case letters")
	}

	if count(pwd, unicode.IsSymbol)+count(pwd, unicode.IsPunct) < 2 {
		return fmt.Errorf("password must include at least two punctuation")
	}

	return
}
