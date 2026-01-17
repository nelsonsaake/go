package sys

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func runCmd(cmd *exec.Cmd) (string, error) {

	outBuf := strings.Builder{}

	mw := io.MultiWriter(&outBuf, os.Stdout)
	cmd.Stdout, cmd.Stderr = mw, mw

	err := cmd.Run()
	out := strings.TrimSpace(outBuf.String())

	return out, err
}
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
