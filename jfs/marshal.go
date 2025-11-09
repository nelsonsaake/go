package jfs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/nelsonsaake/go/afs"
)

// write: write to file at path
func write(path string, bs []byte) error {

	path = afs.Path(path)

	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return err
	}

	err = os.WriteFile(path, bs, 0644)
	if err != nil {
		return err
	}

	return nil
}

// Marshal: marshal v, and this it to path
func Marshal(path string, v any) error {

	bs, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}

	return write(path, bs)
}
