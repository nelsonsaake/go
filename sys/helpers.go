package sys

import (
	"os/exec"
	"path/filepath"
)

func resolvePath(s string) (string, error) {
	if filepath.Base(s) == s {
		if lp, err := exec.LookPath(s); err != nil {
			return s, err
		} else {
			return lp, nil
		}
	}
	return s, nil
}

func ResolveArgs(arg ...any) []string {

	ls := make([]string, 0)
	for _, a := range arg {
		switch v := any(a).(type) {
		case string:
			ls = append(ls, v)
		case []string:
			ls = append(ls, v...)
		case []any:
			ls = append(ls, ResolveArgs(v...)...)
		}
	}

	return ls
}
