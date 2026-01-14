package sys

import (
	"os/exec"
)

func RunCmd(cmd *exec.Cmd) (string, error) {
	return runCmd(cmd)
}

func NewCmd(s string, arg ...any) *exec.Cmd {
	return Cmd(s, arg...).Build()
}

func Run(s string, arg ...any) (string, error) {
	return Cmd(s, arg...).Run()
}

// Ok COMMAND

func Ok(s string, arg ...any) bool {
	return Cmd(s, arg...).Ok()
}
