package axios

import (
	"crypto/tls"
	"maps"
	"net/http"
)

func (c *Client) DoForm(req *http.Request) (*Response, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Transport: tr,
	}

	headers := req.Header

	reqHeaders := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "multipart/form-data",
	}

	maps.Copy(reqHeaders, c.Headers)

	for k, v := range reqHeaders {
		headers.Add(k, v)
	}

	req.Header = headers

	res, err := client.Do(req)

	return NewResponse(res), err
}
