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
)

type Client struct {
	BaseUrl string
	Headers map[string]string
}

func (client *Client) SetBaseUrl(v string) {
	client.BaseUrl = v
}

func (client *Client) AddHeaders(headers map[string]string) {
	for k, v := range headers {
		client.AddHeader(k, v)
	}
}

func (client *Client) AddHeader(key, value string) {
	client.Headers[key] = value
}

func (client *Client) Url(path string) (string, error) {

	if arr.IsEmpty(client.BaseUrl) {
		return path, nil
	}

	return url.JoinPath(client.BaseUrl, path)
}

func (client *Client) Body(body any) (io.Reader, error) {

	reader, ok := body.(io.Reader)
	if ok {
		return reader, nil
	}

	raw, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %v", raw)
	}

	buf := bytes.NewBuffer(raw)

	return buf, nil
}

func (axiosClient *Client) Do(req *http.Request) (*Response, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}

	headers := req.Header

	reqHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
	}

	for k, v := range axiosClient.Headers {
		reqHeaders[k] = v
	}

	for k, v := range reqHeaders {
		headers.Add(k, v)
	}

	req.Header = headers

	res, err := client.Do(req)

	return NewResponse(res), err
}

func NewClient() *Client {
	return &Client{
		Headers: make(map[string]string),
	}
}
