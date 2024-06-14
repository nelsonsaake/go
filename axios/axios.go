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

type Axios struct {
	base string
}

func (axios *Axios) SetBase(v string) {
	axios.base = v
}

func (axios *Axios) Url(path string) (string, error) {

	if arr.IsEmpty(axios.base) {
		return path, nil
	}

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

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := http.Client{
		Transport: tr,
		Timeout:   5 * time.Second,
	}

	hdr := req.Header
	hdr.Add("Accept", "application/json")
	hdr.Add("Content-Type", "application/json")
	req.Header = hdr

	res, err := client.Do(req)

	return NewResponse(res), err
}

func New() *Axios {

	return new(Axios)
}
