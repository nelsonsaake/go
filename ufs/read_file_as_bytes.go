package ufs

import "os"

func ReadFileAsBytes(filepath string) ([]byte, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return content, nil
}
