package strs

import "strings"

func ToCSV(ls []string) string {
	return strings.Join(ls, ",")
}
