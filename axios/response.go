package axios

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/nelsonsaake/go/objs"
)

type Response struct {
	resp *http.Response
}

func NewResponse(resp *http.Response) *Response {
	return &Response{resp: resp}
}

func (resp *Response) Raw() ([]byte, error) {

	die := func(err error) ([]byte, error) {
		return nil, err
	}

	data, err := io.ReadAll(resp.resp.Body)
	if err != nil {
		return die(fmt.Errorf("error reading all response body: %v", err))
	}

	// replace content in response body
	resp.resp.Body = io.NopCloser(bytes.NewBuffer(data))

	return data, nil
}

func (resp *Response) String() (string, error) {

	data, err := resp.Raw()
	return string(data), err
}

func (resp *Response) Bind(v any) error {

	data, err := resp.Raw()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func (resp *Response) Json() (map[string]any, error) {

	v := map[string]any{}
	err := resp.Bind(&v)

	return v, err
}

func (resp *Response) ObjMap() (*objs.Map, error) {

	res, err := resp.Json()
	if err != nil {
		return nil, err
	}

	return objs.From(res), nil
}

func (resp *Response) Request() *http.Request {
	return resp.resp.Request
}

func (resp *Response) Response() *http.Response {
	return resp.resp
}

func (resp *Response) IsSuccessful() bool {
	r := resp.Response()
	return r.StatusCode >= 200 && r.StatusCode < 300
}
