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
)

type Client struct {
	BaseUrl             string
	Headers             map[string]string
	Environments        map[string]map[string]string
	SelectedEnvironment string
}

func (c *Client) SetBaseUrl(v string) {
	c.BaseUrl = v
}

// IsEnvironmentSelected returns true when an active environment name is set and corresponding map exists
func (c *Client) IsEnvironmentSelected() bool {
	if c == nil {
		return false
	}
	if c.SelectedEnvironment == "" {
		return false
	}
	if c.Environments == nil {
		return false
	}
	_, ok := c.Environments[c.SelectedEnvironment]
	return ok
}

// GetEnvironment returns the map for the active environment (may be nil)
func (c *Client) GetEnvironment() map[string]string {
	if c == nil || c.Environments == nil {
		return nil
	}
	return c.Environments[c.SelectedEnvironment]
}

func (c *Client) urlMix(elem ...string) []string {

	if !c.IsEnvironmentSelected() {
		return elem
	}

	res := []string{}
	for _, v := range elem {
		v = mix(v, c.GetEnvironment())
		res = append(res, v)
	}

	return res
}

func (c *Client) urlJoin(elem ...string) (string, error) {

	if len(elem[0]) == 0 {
		return elem[1], nil
	}

	v, err := url.JoinPath(elem[0], elem[1])
	if err != nil {
		return "", err
	}

	return v, nil
}

func (c *Client) Url(path string) (string, error) {

	elems := c.urlMix(c.BaseUrl, path)

	v, err := c.urlJoin(elems...)
	if err != nil {
		return "", err
	}

	return v, nil
}

func (c *Client) Body(body any) (io.Reader, error) {

	reader, ok := body.(io.Reader)
	if ok {
		return reader, nil
	}

	if body == nil {
		return nil, nil
	}

	raw, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling request body: %v", raw)
	}

	buf := bytes.NewBuffer(raw)

	return buf, nil
}

func (c *Client) Do(req *http.Request) (*Response, error) {

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

	for k, v := range c.Headers {
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
