package do

import (
	"path/filepath"
)

func rebase(path, newRoot string) string {
	return filepath.Join(string(newRoot), filepath.Base(path))
}

// get the file name from the uri and attack it to the dir
func URLToLocal(url, dir string) string {
	return rebase(url, dir)
}

// get the file name from local and attach it to uri
func MakeURL(local, uri string) string {
	return filepath.ToSlash(rebase(local, uri))
}
