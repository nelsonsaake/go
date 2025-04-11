package dirs

import (
	"fmt"
	"os"
	"path/filepath"
)

// Entries: return all the files and sub dirs in this current dir
func (dir *Dir) Entries() ([]os.FileInfo, error) {

	var (
		root    = dir.path
		entries = []os.FileInfo{}
	)

	findfiles := func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return fmt.Errorf("error accessing path %s: %v", path, err)
		}

		entries = append(entries, info)

		return nil
	}

	err := filepath.Walk(root, findfiles)
	if err != nil {
		err = fmt.Errorf("error walking the path %s: %v", root, err)
	}

	return entries, err
}
