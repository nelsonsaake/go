package axios

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/nelsonsaake/go/arr"
	"github.com/nelsonsaake/go/str"
)

type axios struct {
	baseUrl string
	headers map[string]string
}

func (axios *axios) SetBaseUrl(v string) {
	axios.baseUrl = v
}

func (axios *axios) AddHeaders(headers map[string]string) {
	for k, v := range headers {
		axios.AddHeader(k, v)
	}
}

func (axios *axios) AddHeader(key, value string) {
	axios.headers[key] = value
}

func (axios *axios) Url(path string) (string, error) {

	if arr.IsEmpty(axios.baseUrl) {
		return path, nil
	}

	return url.JoinPath(axios.baseUrl, path)
}

func (axios *axios) Body(body any) (io.Reader, error) {

	raw, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %v", raw)
	}

	buf := bytes.NewBuffer(raw)

	return buf, nil
}

func (axios *axios) Do(req *http.Request) (*Response, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}

	hdr := req.Header
	for k, v := range axios.headers {
		hdr.Add(k, v)
	}
	req.Header = hdr

	res, err := client.Do(req)

	return NewResponse(res), err
}

type Options struct {
	BaseUrl string
	Headers map[string]string
}

func New(options ...Options) *axios {

	client := axios{
		headers: map[string]string{
			"Accept":       "application/json",
			"Content-Type": "application/json",
		},
	}

	for _, opt := range options {
		if str.IsEmpty(opt.BaseUrl) {
			client.baseUrl = opt.BaseUrl
		}
		for k, v := range opt.Headers {
			client.headers[k] = v
		}
	}

	return &client
}
