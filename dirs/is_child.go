package dirs

import "strings"

// IsChild: check to see if path, is a direct child of this dir
func (dir *Dir) IsChild(path string) bool {
	child, parent := clean(path), clean(dir.path)
	return strings.HasPrefix(child, parent) && child != parent
}
