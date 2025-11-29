package resources

import (
	"slices"

	"github.com/nelsonsaake/go/afs"
	"github.com/nelsonsaake/go/storage/local"
	"github.com/nelsonsaake/go/ufs"
)

var prefix = "/resources/"

func ls() *local.Storage {
	return local.New(Root())
}

func Store(b64content string) (string, error) {
	return ls().Save(b64content)
}

func SaveAs(path, b64content string) error {
	return ls().SaveAs(path, b64content)
}

func Delete(file string) error {
	return ls().Delete(file)
}

func Path(elem ...string) string {
	return afs.Path(slices.Concat([]string{prefix}, elem)...)
}

func Root() string {
	return afs.Path(prefix)
}

func MakeDir(elem ...string) (string, error) {
	dir := Path(elem...)
	err := ufs.MakeDir(dir)
	return dir, err
}
