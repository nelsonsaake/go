package axios

import (
	"fmt"
	"net/http"
)

func (c *Client) Delete(url string, config ...any) (*Response, error) {

	die := func(f string, a ...any) (*Response, error) {
		return nil, fmt.Errorf(f, a...)
	}

	url, err := c.Url(url)
	if err != nil {
		return die("error making request url: %v", err)
	}

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return die("error making new request: %v", err)
	}

	req = ResolveConfigArray(req, config...)
	return c.Do(req)
}
