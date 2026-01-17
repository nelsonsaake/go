package mix

import (
	"fmt"
	"os"
)

func File(path string, data any) (string, error) {

	die := func(f string, a ...any) (string, error) {
		return "", fmt.Errorf(f, a...)
	}

	bytes, err := os.ReadFile(path)
	if err != nil {
		return die("error reading file %q: %v", path, err)
	}

	return String(string(bytes), data)
}
