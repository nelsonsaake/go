package axios

import (
	"fmt"
	"net/http"
)

type QueryParam map[string]any

func AppendQueryParam(q QueryParam) func(req *http.Request) {
	return func(req *http.Request) {

		values := req.URL.Query()

		for k, v := range q {
			values.Add(k, fmt.Sprint(v))
		}

		req.URL.RawQuery = values.Encode()
	}
}

func (client *Client) Get(url string, config ...any) (*Response, error) {

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

	for _, config := range config {
		switch config := config.(type) {
		case func(*http.Request):
			configfunc := config
			configfunc(req)
		case QueryParam:
			configfunc := AppendQueryParam(config)
			configfunc(req)
		}
	}

	return client.Do(req)
}
