package b64

import (
	"encoding/base64"
	"net/http"
	"path"
)

func Encode(rawstr string) string {
	return base64.StdEncoding.EncodeToString([]byte(rawstr))
}

func Decode(b64str string) (bs []byte, err error) {
	return base64.StdEncoding.DecodeString(b64str)
}

func FromBytes(file []byte) string {
	return base64.StdEncoding.EncodeToString(file)
}

func Ext(b64str string) (ext string, err error) {
	bs, err := base64.StdEncoding.DecodeString(b64str)
	if err != nil {
		return
	}

	return extFromBytes(bs), nil
}

// ext: file extension from bytes in the file
func extFromBytes(bs []byte) (ext string) {
	ext = http.DetectContentType(bs)
	ext = path.Base(ext)
	if len(ext) == 0 {
		return ext
	}
	return "." + ext
}
