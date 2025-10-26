package repo

import (
	"strings"

	"github.com/nelsonsaake/go/strs"
)

func getRscName[T any]() string {

	resource := getTypeName(new(T))
	resource = strs.ToWords(resource)
	resource = strings.ToLower(resource)

	return resource
}
