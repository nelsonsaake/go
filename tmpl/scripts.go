package tmpl

import (
	"fmt"
	"os"
	"path/filepath"
)

func Scripts() ([]string, error) {

	var (
		root  = RootDir
		files = []string{}
	)

	findFiles := func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return fmt.Errorf("error accessing path %s: %v", path, err)
		}

		if !info.IsDir() && isDirectChildOfRootDir(path) {
			files = append(files, CleanPath(path))
		}

		return nil
	}

	err := filepath.Walk(root, findFiles)
	if err != nil {
		err = fmt.Errorf("error walking the path %s: %v", root, err)
	}

	return files, err
}
