package urls

import "net/url"

func Join(base string, elem ...string) (string, error) {
	return url.JoinPath(base, elem...)
}
