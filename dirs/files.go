package dirs

import (
	"fmt"
	"os"
	"path/filepath"
)

// Files: return all the files in this current dir
func (dir *Dir) Files() ([]string, error) {

	var (
		root  = dir.path
		files = []string{}
	)

	findfiles := func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return fmt.Errorf("error accessing path %s: %v", path, err)
		}

		if info.IsDir() {
			return nil
		}

		files = append(files, clean(path))

		return nil
	}

	err := filepath.Walk(root, findfiles)
	if err != nil {
		err = fmt.Errorf("error walking the path %s: %v", root, err)
	}

	return files, err
}
