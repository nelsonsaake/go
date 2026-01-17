package sys

import "strings"

type Dump struct {
	Output string
	Error  error
}

func (d *Dump) OK() bool {
	return d.Error == nil
}

func (d *Dump) Contains(substring string) bool {
	return strings.Contains(d.Output, substring)
}
