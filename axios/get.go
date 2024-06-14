package axios

import (
	"fmt"
	"net/http"
)

func QueryParam(q map[string]string) func(req *http.Request) {
	return func(req *http.Request) {

		values := req.URL.Query()

		for k, v := range q {
			values.Add(k, v)
		}

		req.URL.RawQuery = values.Encode()
	}
}

func (axios *Axios) Get(url string, config ...func(*http.Request)) (*Response, error) {

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

	for _, config := range config {
		config(req)
	}

	return axios.Do(req)
}
