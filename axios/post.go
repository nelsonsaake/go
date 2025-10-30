package axios

import (
	"fmt"
	"net/http"
)

func (c *Client) Post(url string, body any) (*Response, error) {

	die := func(f string, a ...any) (*Response, error) {
		return nil, fmt.Errorf(f, a...)
	}

	url, err := c.Url(url)
	if err != nil {
		return die("error making request url: %v", err)
	}

	rBody, err := c.Body(body)
	if err != nil {
		return die("error making request body: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, rBody)
	if err != nil {
		return die("error making new request: %v", err)
	}

	return c.Do(req)
}
