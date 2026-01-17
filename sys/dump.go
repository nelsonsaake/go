package sys

import "strings"

type Results struct {
	Dump  string
	Error error
}

func (r *Results) OK() bool {
	return r.Error == nil
}

func (r *Results) Contains(substring string) bool {
	return strings.Contains(r.Dump, substring)
}
