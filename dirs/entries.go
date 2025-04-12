package dirs

import (
	"os"
)

// Entries: return all the files and sub dirs in this current dir
func (dir *Dir) Entries() ([]os.DirEntry, error) {
	return os.ReadDir(dir.path)
}
