package ufs

import (
	"fmt"
	"os"
)

func ReadDir(from string) ([]string, error) {

	die := func(err error) ([]string, error) {
		return nil, fmt.Errorf("%v \nfrom: %s", err, from)
	}

	if !IsDir(from) {
		return die(fmt.Errorf("from(source) is not a dir"))
	}

	entries, err := os.ReadDir(from)
	if err != nil {
		return die(err)
	}

	var files []string
	for _, entry := range entries {
		files = append(files, entry.Name())
	}

	return files, nil
}
