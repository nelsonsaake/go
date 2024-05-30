package do

import (
	"net/http"
	"path"
)

// Ext: file extension from bytes in the file
func ExtFromBytes(bs []byte) (ext string) {
	ext = http.DetectContentType(bs)
	ext = path.Base(ext)
	if len(ext) == 0 {
		return ext
	}
	return "." + ext
}
