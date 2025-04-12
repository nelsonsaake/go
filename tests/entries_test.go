package tests

import (
	"fmt"
	"testing"

	"github.com/nelsonsaake/go/dirs"
	"github.com/nelsonsaake/go/do"
)

func TestEntries(t *testing.T) {

	path := `C:\Users\user\Desktop\projects\code\go\do`

	dir, err := dirs.New(path)
	do.Die(err)

	entries, err := dir.Entries()
	do.Die(err)

	for k, entry := range entries {
		fmt.Println(k, entry.Name())
	}
}
