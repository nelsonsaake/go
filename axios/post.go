package axios

import (
	"fmt"
	"net/http"
)

func (axios *axios) Post(url string, body any) (*Response, error) {

	die := func(err error) (*Response, error) {
		return nil, err
	}

	url, err := axios.Url(url)
	if err != nil {
		return die(fmt.Errorf("error making request url: %v", err))
	}

	rBody, err := axios.Body(body)
	if err != nil {
		return die(fmt.Errorf("error making request body: %v", err))
	}

	req, err := http.NewRequest(http.MethodPost, url, rBody)
	if err != nil {
		return die(fmt.Errorf("error making new request: %v", err))
	}

	return axios.Do(req)
}
