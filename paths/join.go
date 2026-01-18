package paths

import "path/filepath"

func toStringSlice(arg ...any) []string {

	ls := make([]string, 0)
	for _, a := range arg {
		switch v := any(a).(type) {
		case string:
			ls = append(ls, v)
		case []string:
			ls = append(ls, v...)
		case []any:
			ls = append(ls, toStringSlice(v...)...)
		}
	}

	return ls
}

func Join(paths ...any) string {
	return filepath.Join(toStringSlice(paths...)...)
}

func (p *pfs) Join(paths ...any) string {

	joinedPath := Join(paths...)

	if p.unix {
		joinedPath = filepath.ToSlash(joinedPath)
	}

	return joinedPath
}
