package do

import (
	"net/http"
	"strings"
)

func IsImage(v []byte) bool {
	return strings.HasPrefix(http.DetectContentType(v), "image")
}
