package dirs

import "strings"

// IsChild: check to see if path, is a direct child of this dir
func (dir *Dir) IsChild(path string) bool {
	return strings.HasPrefix(clean(path), clean(dir.path))
}
