package sys

import "strings"

type CmdResults struct {
	Dump  string
	Error error
}

func (r *CmdResults) OK() bool {
	return r.Error == nil
}

func (r *CmdResults) Contains(ls ...string) bool {
	for _, s := range ls {
		if !strings.Contains(r.Dump, s) {
			return false
		}
	}
	return true
}

func (r *CmdResults) Result() (string, error) {
	return r.Dump, r.Error
}
