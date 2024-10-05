package dirs

import (
	"fmt"
	"os"
	"path/filepath"
)

// SubDirs: return all the dirs in this current dir
func (dir *Dir) SubDirs() ([]string, error) {

	var (
		root = dir.path
		res  = []string{}
	)

	findfiles := func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return fmt.Errorf("error accessing path %s: %v", path, err)
		}

		if info.IsDir() && dir.IsChild(path) {
			res = append(res, clean(path))
		}

		return nil
	}

	err := filepath.Walk(root, findfiles)
	if err != nil {
		err = fmt.Errorf("error walking the path %s: %v", root, err)
	}

	return res, err
}
