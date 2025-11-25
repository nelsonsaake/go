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

func (res *Response) Bytes() ([]byte, error) {

	die := func(err error) ([]byte, error) {
		return nil, err
	}

	data, err := io.ReadAll(res.resp.Body)
	if err != nil {
		return die(fmt.Errorf("error reading all response body: %v", err))
	}

	// replace content in response body
	res.resp.Body = io.NopCloser(bytes.NewBuffer(data))

	return data, nil
}

func (res *Response) String() (string, error) {

	data, err := res.Bytes()
	return string(data), err
}

func (res *Response) Bind(v any) error {

	data, err := res.Bytes()
	if err != nil {
		return err
	}

	return json.Unmarshal(data, v)
}

func (res *Response) IsArray() bool {

	s, err := res.String()
	if err != nil {
		return false
	}

	return strings.HasPrefix(s, "[")
}

func (res *Response) Array() ([]any, error) {

	v := []any{}
	err := res.Bind(&v)

	return v, err
}

func (res *Response) IsMap() bool {

	s, err := res.String()
	if err != nil {
		return false
	}

	return strings.HasPrefix(s, "{")
}

func (res *Response) Map() (map[string]any, error) {

	v := map[string]any{}
	err := res.Bind(&v)

	return v, err
}

func (res *Response) Json() (any, error) {

	if res.IsArray() {
		return res.Array()
	}

	if res.IsMap() {
		return res.Map()
	}

	return res.String()
}

func (r *Response) Obj() (*objs.Obj, error) {

	res, err := r.Map()
	if err != nil {
		return nil, err
	}

	return objs.FromMap(res), nil
}

func (res *Response) Request() *http.Request {
	return res.resp.Request
}

func (res *Response) Response() *http.Response {
	return res.resp
}

func (res *Response) IsSuccessful() bool {
	r := res.Response()
	return r.StatusCode >= 200 && r.StatusCode < 300
}

func (res *Response) StatusCode() int {
	return res.Response().StatusCode
}

func (res *Response) Status() string {
	return res.Response().Status
}

func (res *Response) Headers() http.Header {
	return res.Response().Header
}

func (res *Response) Header(key string) string {
	return res.Response().Header.Get(key)
}

func (res *Response) Cookies() []*http.Cookie {
	return res.Response().Cookies()
}
