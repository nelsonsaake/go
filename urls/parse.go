package urls

import (
	"net/url"
)

func Parse(uri string) (*url.URL, error) {
	return url.Parse(uri)
}
