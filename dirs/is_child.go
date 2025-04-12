package dirs

import (
	"path/filepath"
	"strings"
)

func IsParentDir(parentDir, childDir string) bool {
	parentDir = filepath.Clean(parentDir)
	childDir = filepath.Clean(childDir)

	// Normalize the paths to ensure consistent comparison
	parentDir = filepath.ToSlash(parentDir)
	childDir = filepath.ToSlash(childDir)

	return strings.HasPrefix(childDir, parentDir+"/")
}

func IsChildOfDir(file, dir string) bool {
	absFilePath, err := filepath.Abs(file)
	if err != nil {
		return false
	}

	absDirPath, err := filepath.Abs(dir)
	if err != nil {
		return false
	}

	return strings.HasPrefix(absFilePath, absDirPath)
}

// IsChild: check to see if path, is a direct child of this dir
func (dir *Dir) IsChild(path string) bool {
	return IsParentDir(dir.path, path)
}
