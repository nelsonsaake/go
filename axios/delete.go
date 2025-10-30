package axios

import (
	"fmt"
	"net/http"
)

func (b *Client) Delete(url string, config ...any) (*Response, error) {

	die := func(f string, a ...any) (*Response, error) {
		return nil, fmt.Errorf(f, a...)
	}

	url, err := b.Url(url)
	if err != nil {
		return die("error making request url: %v", err)
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return die("error making new request: %v", err)
	}

	req = ResolveConfigArray(req, config...)
	return b.Do(req)
}
