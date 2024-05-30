package cloner

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Clone: copy the information at the url and returns it as base64 string,
// it's intended for copying images
func Clone(url string) (_ string, err error) {
	response, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("error getting url: %v", err)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("error reading payload: %v", err)
		return
	}

	res := base64.StdEncoding.EncodeToString(data)
	return res, nil
}
