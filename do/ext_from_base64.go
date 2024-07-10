package do

import "encoding/base64"

func ExtFromBase64(file string) (ext string, err error) {
	bs, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		return
	}

	return ExtFromBytes(bs), nil
}
