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
	BaseUrl             string
	Headers             map[string]string
	Environments        map[string]map[string]string
	SelectedEnvironment string
}

func (b *Client) SetBaseUrl(v string) {
	b.BaseUrl = v
}

// IsEnvironmentSelected returns true when an active environment name is set and corresponding map exists
func (b *Client) IsEnvironmentSelected() bool {
	if b == nil {
		return false
	}
	if b.SelectedEnvironment == "" {
		return false
	}
	if b.Environments == nil {
		return false
	}
	_, ok := b.Environments[b.SelectedEnvironment]
	return ok
}

// GetEnvironment returns the map for the active environment (may be nil)
func (b *Client) GetEnvironment() map[string]string {
	if b == nil || b.Environments == nil {
		return nil
	}
	return b.Environments[b.SelectedEnvironment]
}

func (b *Client) Url(path string) (string, error) {

	if arr.IsEmpty(b.BaseUrl) {
		return path, nil
	}

	// build base path
	v, err := url.JoinPath(b.BaseUrl, path)
	if err != nil {
		return "", err
	}

	// if environment is set, apply mix substitutions
	if b.IsEnvironmentSelected() {
		v = mix(v, b.GetEnvironment())
	}

	return v, nil
}

func (b *Client) Body(body any) (io.Reader, error) {

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

func (b *Client) Do(req *http.Request) (*Response, error) {

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

	for k, v := range b.Headers {
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
		Headers:      make(map[string]string),
		Environments: make(map[string]map[string]string),
	}
}
