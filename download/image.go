package download

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"
)

func empty(v string) bool {
	return len(strings.TrimSpace(v)) == 0
}

func Image(url, fpath string) (err error) {
	if empty(url) {
		return fmt.Errorf("url empty")
	}

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

	ext := filepath.Ext(fpath)
	if len(ext) == 0 {
		ext = http.DetectContentType(data)
		ext = filepath.Base(ext)

		fpath = strings.TrimSuffix(fpath, "/")
		fpath = fmt.Sprintf("%s.%s", fpath, ext)
	}

	return ioutil.WriteFile(fpath, data, 0777)
}
