package strs

import "github.com/iancoleman/strcase"

func ToWordCase(v string) string {
	return strcase.ToDelimited(v, ' ')
}
