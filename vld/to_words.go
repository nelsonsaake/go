package vld

import "github.com/iancoleman/strcase"

func toWords(v string) string {
	return strcase.ToDelimited(v, ' ')
}
