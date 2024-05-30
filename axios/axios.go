package axios

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Axios struct {
	base string
}

func (axios *Axios) SetBase(v string) {
	axios.base = v
}

func (axios *Axios) Url(path string) (string, error) {
	return url.JoinPath(axios.base, path)
}

func (axios *Axios) Body(body any) (io.Reader, error) {

	raw, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %v", raw)
	}

	buf := bytes.NewBuffer(raw)

	return buf, nil
}

func (axios *Axios) Do(req *http.Request) (*Response, error) {

	res, err := http.DefaultClient.Do(req)
	return NewResponse(res), err
}

func New() *Axios {

	return new(Axios)
}
