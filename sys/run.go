package sys

import (
	"os/exec"
)

func RunCmd(cmd *exec.Cmd) *Results {
	return New().RunCmd(cmd)
}

func NewCmd(s string, arg ...any) *exec.Cmd {
	c, _ := Command(s, arg...).Build()
	return c
}

func Run(s string, arg ...any) *Results {
	return Command(s, arg...).Run()
}

// Ok COMMAND

func Ok(s string, arg ...any) bool {
	return Command(s, arg...).Run().OK()
}
