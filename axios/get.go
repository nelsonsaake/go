package axios

import (
	"fmt"
	"net/http"
)

func (c *Client) Get(url string, config ...any) (*Response, error) {

	die := func(err error) (*Response, error) {
		return nil, err
	}

	url, err := c.Url(url)
	if err != nil {
		return die(fmt.Errorf("error making request url: %v", err))
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return die(fmt.Errorf("error making new request: %v", err))
	}

	req = ResolveConfigArray(req, config...)

	return c.Do(req)
}
