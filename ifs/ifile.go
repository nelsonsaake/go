package ifs

import (
	"bytes"
	"encoding/base64"
	"io"
	"net/http"
	"path"
)

type ifile struct {
	URL   string
	name  string
	Bytes []byte
}

func (f ifile) Base64() string {
	return base64.StdEncoding.EncodeToString(f.Bytes)
}

// Ext: file extension
func (f ifile) Ext() (ext string) {
	ext = http.DetectContentType(f.Bytes)
	ext = path.Base(ext)
	if len(ext) == 0 {
		return ext
	}
	return "." + ext
}

// (f ifile) NameRaw: this is the timestamp at which the file was created and the file extension
func (f ifile) NameRaw() (ext string) {
	return f.name
}

// (f ifile) Name: this is the timestamp at which the file was created and the file extension
func (f ifile) Name() string {
	return f.name + f.Ext()
}

func (f ifile) NewReader() (reader io.Reader) {
	return bytes.NewReader(f.Bytes)
}
