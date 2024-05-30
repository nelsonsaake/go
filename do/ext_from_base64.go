package do

import "encoding/base64"

func ExtFromBase64(file string) (ext string, err error) {
	bs, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		return
	}

	return ExtFromBytes(bs), nil
}

// old
func getExtension(file string) string {
	if len(file) == 0 {
		return ""
	}
	switch file[0] {
	case '/':
		return ".jpeg"
	case 'i':
		return ".png"
	case 'R':
		return ".gif"
	case 'U':
		return ".webp"
	case 'J':
		return ".pdf"
	default:
		return ""
	}
}
