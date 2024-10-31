package dirs

import (
	"path/filepath"
	"strings"
)

func isParentDir(parentDir, childDir string) bool {
	parentDir = filepath.Clean(parentDir)
	childDir = filepath.Clean(childDir)

	// Normalize the paths to ensure consistent comparison
	parentDir = filepath.ToSlash(parentDir)
	childDir = filepath.ToSlash(childDir)

	return strings.HasPrefix(childDir, parentDir+"/")
}

// IsChild: check to see if path, is a direct child of this dir
func (dir *Dir) IsChild(path string) bool {
	return isParentDir(dir.path, path)
}
