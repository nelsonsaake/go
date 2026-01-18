package paths

import (
	"strings"
)

func (p *pfs) Rel(fullPath string) string {

	uri := strings.TrimPrefix(
		Clean(fullPath),
		Clean(p.root),
	)

	return uri
}
