package dir

import (
	"path/filepath"
)

// IsChild: check to see if path, is a direct child of this dir
func (dir *Dir) IsChild(path string) bool {
	return clean(dir.path) == clean(filepath.Dir(path))
}
