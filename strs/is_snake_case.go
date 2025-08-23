package strs

import "regexp"

func IsSnakeCase(s string) bool {
	re := regexp.MustCompile(`^[a-z]+(_[a-z0-9]+)*$`)
	return re.MatchString(s)
}
