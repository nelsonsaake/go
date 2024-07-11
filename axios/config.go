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

func ResolveConfigArray(req *http.Request, config ...any) *http.Request {

	for _, config := range config {

		var configfunc func(*http.Request)

		switch config := config.(type) {

		case func(*http.Request):
			configfunc = config

		case QueryParam:
			configfunc = AppendQueryParam(config)

		case map[string]any:
			configfunc = AppendQueryParam(config)
		}

		configfunc(req)
	}

	return req
}
