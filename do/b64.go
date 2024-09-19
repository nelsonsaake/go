package do

import "encoding/base64"

func B64(file string) string {
	return base64.StdEncoding.EncodeToString([]byte(file))
}

func RawToB64(file []byte) string {
	return base64.StdEncoding.EncodeToString(file)
}

func B64ToRaw(b64 string) (bs []byte, err error) {
	bs, err = base64.StdEncoding.DecodeString(b64)
	return
}

func B64Ext(b64string string) (ext string, err error) {
	bs, err := base64.StdEncoding.DecodeString(b64string)
	if err != nil {
		return
	}

	return ExtFromBytes(bs), nil
}
