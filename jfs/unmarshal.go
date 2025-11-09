package jfs

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/nelsonsaake/go/afs"
)

// read: read file content from os
func read(path string) ([]byte, error) {
	return os.ReadFile(afs.Path(path))
}

// Unmarshal: read file content at path and unmarshal it into v
func Unmarshal(path string, v any) error {

	bs, err := read(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bs, v)
	if err != nil {
		return fmt.Errorf("error unmarshalling data: %v", err)
	}

	return nil
}
