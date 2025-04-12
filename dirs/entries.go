package dirs

import (
	"fmt"
	"os"
	"path/filepath"
)

// Entries: return all the files and sub dirs in this current dir
func (dir *Dir) Entries() ([]os.FileInfo, error) {

	var (
		root = dir.path
		mp   = map[string]os.FileInfo{}
	)

	findfiles := func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return fmt.Errorf("error accessing path %s: %v", path, err)
		}

		mp[path] = info

		return nil
	}

	err := filepath.Walk(root, findfiles)
	if err != nil {
		err = fmt.Errorf("error walking the path %s: %v", root, err)
	}

	ls := make([]os.FileInfo, 0, len(mp))
	for _, info := range mp {
		ls = append(ls, info)
	}

	return ls, err
}
