package dirs

import (
	"os"
)

// Files: return all the files in this current dir
func (dir *Dir) Files() ([]string, error) {

	entries, err := os.ReadDir(dir.path)
	if err != nil {
		return nil, err
	}

	res := []string{}
	for _, entry := range entries {
		if !entry.IsDir() {
			res = append(res, entry.Name())
		}
	}

	return res, err
}
