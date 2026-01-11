package sys

import (
	"io"
	"os"
	"os/exec"
	"strings"
)

func RunCmd(cmd *exec.Cmd) (string, error) {

	outBuf := strings.Builder{}

	mw := io.MultiWriter(&outBuf, os.Stdout)
	cmd.Stdout, cmd.Stderr = mw, mw

	err := cmd.Run()
	out := strings.TrimSpace(outBuf.String())

	return out, err
}

func NewCmd(s string, arg ...any) *exec.Cmd {

	ls := make([]string, 0)
	for _, a := range arg {
		switch v := any(a).(type) {
		case string:
			ls = append(ls, v)
		case []string:
			ls = append(ls, v...)
		}
	}

	return exec.Command(s, ls...)
}

func Run(s string, arg ...any) (string, error) {

	cmd := NewCmd(s, arg...)
	return RunCmd(cmd)
}

// Ok COMMAND

func Ok(s string, arg ...any) bool {
	return NewCmd(s, arg...).Run() == nil
}
