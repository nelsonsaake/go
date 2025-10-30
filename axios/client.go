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
	BaseUrl      string
	Headers      map[string]string
	Environments map[string]map[string]string
	Environment  string
}

func (client *Client) SetBaseUrl(v string) {
	client.BaseUrl = v
}

// Add builder: set multiple environments and return client for chaining
func (client *Client) WithEnvironments(envs map[string]map[string]string) *Client {
	client.Environments = envs
	return client
}

// Add builder: set active environment name and return client for chaining
func (client *Client) WithEnvironment(name string) *Client {
	client.Environment = name
	return client
}

// IsEnvironmentSet returns true when an active environment name is set and corresponding map exists
func (client *Client) IsEnvironmentSet() bool {
	if client == nil {
		return false
	}
	if client.Environment == "" {
		return false
	}
	if client.Environments == nil {
		return false
	}
	_, ok := client.Environments[client.Environment]
	return ok
}

// GetEnvironment returns the map for the active environment (may be nil)
func (client *Client) GetEnvironment() map[string]string {
	if client == nil || client.Environments == nil {
		return nil
	}
	return client.Environments[client.Environment]
}

func (client *Client) Url(path string) (string, error) {

	if arr.IsEmpty(client.BaseUrl) {
		return path, nil
	}

	// build base path
	v, err := url.JoinPath(client.BaseUrl, path)
	if err != nil {
		return "", err
	}

	// if environment is set, apply mix substitutions
	if client.IsEnvironmentSet() {
		v = mix(v, client.GetEnvironment())
	}

	return v, nil
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
