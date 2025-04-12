package dirs

import (
	"os"
)

// SubDirs: return all the dirs in this current dir
func (dir *Dir) SubDirs() ([]string, error) {

	entries, err := os.ReadDir(dir.path)
	if err != nil {
		return nil, err
	}

	res := []string{}
	for _, entry := range entries {
		if entry.IsDir() {
			res = append(res, entry.Name())
		}
	}

	return res, err
}
