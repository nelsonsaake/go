package sys

import "strings"

type CmdResults struct {
	Dump  string
	Error error
}

func (r *CmdResults) OK() bool {
	return r.Error == nil
}

func (r *CmdResults) Contains(substring string) bool {
	return strings.Contains(r.Dump, substring)
}
