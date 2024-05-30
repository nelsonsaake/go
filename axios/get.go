package axios

import (
	"fmt"
	"net/http"
)

func (axios *Axios) Get(url string) (*Response, error) {

	die := func(err error) (*Response, error) {
		return nil, err
	}

	url, err := axios.Url(url)
	if err != nil {
		return die(fmt.Errorf("error making request url: %v", err))
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return die(fmt.Errorf("error making new request: %v", err))
	}

	return axios.Do(req)
}
