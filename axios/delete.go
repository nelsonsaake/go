package axios

import (
	"fmt"
	"net/http"
)

func (client *Client) Delete(url string, config ...any) (*Response, error) {

	die := func(err error) (*Response, error) {
		return nil, err
	}

	url, err := client.Url(url)
	if err != nil {
		return die(fmt.Errorf("error making request url: %v", err))
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return die(fmt.Errorf("error making new request: %v", err))
	}

	req = ResolveConfigArray(req, config...)
	return client.Do(req)
}
