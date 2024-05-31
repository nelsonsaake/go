package download

import (
	"net/http"
	"path"
)

// extractExt: file extension from bytes in the file
func extractExt(data []byte) string {
	ext := http.DetectContentType(data)
	ext = path.Base(ext)
	if len(ext) == 0 {
		return ext
	}
	return "." + ext
}
