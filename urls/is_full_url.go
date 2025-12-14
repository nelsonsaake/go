package urls

import (
	"net/url"
)

func IsAbs(v string) bool {
	p, err := url.Parse(v)
	if err == nil && p.Scheme != "" && p.Host != "" {
		return true
	}
	return false
}
