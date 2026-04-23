package urls

import (
	"fmt"
	"net/url"
)

func New(uri string, queries map[string]any) (string, error) {

	u, err := url.Parse(uri)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for k, v := range queries {
		q.Set(k, fmt.Sprint(v))
	}

	u.RawQuery = q.Encode()

	return u.String(), nil
}
