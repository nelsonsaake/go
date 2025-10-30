package axios

import (
	"fmt"
	"net/http"
)

func (b *Client) Patch(url string, body any) (*Response, error) {

	die := func(err error) (*Response, error) {
		return nil, err
	}

	url, err := b.Url(url)
	if err != nil {
		return die(fmt.Errorf("error making request url: %v", err))
	}

	rBody, err := b.Body(body)
	if err != nil {
		return die(fmt.Errorf("error making request body: %v", err))
	}

	req, err := http.NewRequest(http.MethodPatch, url, rBody)
	if err != nil {
		return die(fmt.Errorf("error making new request: %v", err))
	}

	return b.Do(req)
}
