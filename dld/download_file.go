package dld

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadFile: download file, save to dir, and return the path and error
func DownloadFile(url, dir string) (string, error) {

	dd := func(err error) (string, error) {
		return "", err
	}

	if isEmpty(url) {
		return dd(fmt.Errorf("url empty"))
	}

	response, err := http.Get(url)
	if err != nil {
		return dd(fmt.Errorf("error getting url: %v", err))
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return dd(fmt.Errorf("error reading payload: %v", err))
	}

	fname := makeName(data)
	fpath := filepath.Join(dir, fname)

	err = makeDir(dir)
	if err != nil {
		return dd(fmt.Errorf("error making dir: %v", err))
	}

	err = os.WriteFile(fpath, data, 0777)
	if err != nil {
		return dd(fmt.Errorf("error saving image: %v", err))
	}

	return fpath, nil
}
