package do

import (
	"bytes"
	"encoding/json"
	"io"
)

// JSONReader: convert v, to json, and convert the json to io.Reader
func JSONReader(v interface{}) (_ io.Reader, err error) {

	b, err := json.Marshal(v)
	if err != nil {
		return
	}

	return bytes.NewBuffer(b), nil
}
