package sys

import (
	"io"
	"os"
	"os/exec"
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

func resolveArgs(arg ...any) []string {

	ls := make([]string, 0)
	for _, a := range arg {
		switch v := any(a).(type) {
		case string:
			ls = append(ls, v)
		case []string:
			ls = append(ls, v...)
		}
	}

	return ls
}
