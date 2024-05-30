package ifile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func New(url string) (_ *ifile, err error) {
	
	response, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("error getting url: %v", err)
		return
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		err = fmt.Errorf("error reading payload from(%s): %v", url, err)
		return
	}

	if len(data) == 0 {
		err = fmt.Errorf("empty response body")
		return
	}

	file := &ifile{
		URL:   url,
		name:  fmt.Sprint(time.Now().UnixNano()),
		Bytes: data,
	}

	return file, nil
}
