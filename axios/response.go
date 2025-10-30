package axios

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/nelsonsaake/go/objs"
)

type Response struct {
	resp *http.Response
}

func NewResponse(resp *http.Response) *Response {
	return &Response{resp: resp}
}

func (r *Response) Raw() ([]byte, error) {

	die := func(err error) ([]byte, error) {
		return nil, err
	}

	data, err := io.ReadAll(r.resp.Body)
	if err != nil {
		return die(fmt.Errorf("error reading all response body: %v", err))
	}

	// replace content in response body
	r.resp.Body = io.NopCloser(bytes.NewBuffer(data))

	return data, nil
}

func (r *Response) String() (string, error) {

	data, err := r.Raw()
	return string(data), err
}

func (r *Response) Bind(v any) error {

	data, err := r.Raw()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func (r *Response) IsMap() bool {

	s, err := r.String()
	if err != nil {
		return false
	}

	return strings.HasPrefix(s, "{")
}

func (r *Response) Map() (map[string]any, error) {

	v := map[string]any{}
	err := r.Bind(&v)

	return v, err
}

func (r *Response) Obj() (*objs.Obj, error) {

	res, err := r.Map()
	if err != nil {
		return nil, err
	}

	return objs.FromMap(res), nil
}

func (r *Response) Request() *http.Request {
	return r.resp.Request
}

func (r *Response) Response() *http.Response {
	return r.resp
}

func (resp *Response) IsSuccessful() bool {
	r := resp.Response()
	return r.StatusCode >= 200 && r.StatusCode < 300
}
