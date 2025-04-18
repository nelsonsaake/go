package strs

import "github.com/gertd/go-pluralize"

var pc = pluralize.NewClient()

func PluralCase(v string) string {
	return pc.Plural(v)
}
