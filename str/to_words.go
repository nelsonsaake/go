package str

import "github.com/iancoleman/strcase"

func ToWords(v string) string {
	return strcase.ToDelimited(v, ' ')
}
