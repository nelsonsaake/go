package axios

import (
	"fmt"
	"net/http"
)

func (b *Client) Get(url string, config ...any) (*Response, error) {

	die := func(err error) (*Response, error) {
		return nil, err
	}

	url, err := b.Url(url)
	if err != nil {
		return die(fmt.Errorf("error making request url: %v", err))
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return die(fmt.Errorf("error making new request: %v", err))
	}

	req = ResolveConfigArray(req, config...)

	return b.Do(req)
}
