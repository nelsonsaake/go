package str

import "unicode"

func ToSentenceCase(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(s)                   // Convert to runes to support Unicode characters
	runes[0] = unicode.ToUpper(runes[0]) // Capitalize the first letter
	return string(runes)
}
