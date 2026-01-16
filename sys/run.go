package sys

import (
	"os/exec"
)

func RunCmd(cmd *exec.Cmd) (string, error) {
	return runCmd(cmd)
}

func NewCmd(s string, arg ...any) *exec.Cmd {
	return Command(s, arg...).Build()
}

func Run(s string, arg ...any) error {
	return Command(s, arg...).Run()
}

func Runo(s string, arg ...any) (string, error) {
	return Command(s, arg...).Runo()
}

// Ok COMMAND

func Ok(s string, arg ...any) bool {
	return Command(s, arg...).Ok()
}
