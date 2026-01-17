package sys

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/nelsonsaake/go/afs"
)

type Cmd struct {
	exec.Cmd
	file       *string
	outWriters []io.Writer
	errWriters []io.Writer
	IsNoDump   bool
	error      error
}

func (c *Cmd) WD(v string) *Cmd {
	c.Dir = v
	return c
}

func (c *Cmd) Env(k, v string) *Cmd {
	c.Cmd.Env = append(c.Cmd.Env, k+"="+v)
	return c
}

func (c *Cmd) WithFile(v string) *Cmd {

	if !filepath.IsAbs(v) {
		v = afs.Path(v)
	}

	c.file = &v

	return c
}

func (c *Cmd) NoDump() *Cmd {
	c.IsNoDump = true
	return c
}

func (c *Cmd) IsDump() bool {
	return c.IsNoDump == false
}

func (c *Cmd) NI() *Cmd {
	return c.Env("DEBIAN_FRONTEND", "noninteractive")
}

func (c *Cmd) Build() (*exec.Cmd, error) {

	die := func(f string, a ...any) (*exec.Cmd, error) {
		return nil, fmt.Errorf(f, a...)
	}

	if c.file == nil {
		return &c.Cmd, nil
	}

	// TODO IMPLEMENT FILE SECTION

	f, err := os.Open(*c.file)
	if err != nil {
		return die("failed to open file %q: %w", *c.file, err)
	}

	c.Stdin = f

	return &c.Cmd, nil
}

func (c *Cmd) RunCmd(cmd *exec.Cmd) *CmdResults {

	var outBuf = &strings.Builder{}

	if c.IsDump() {
		c.outWriters = append(c.outWriters, os.Stdout, outBuf)
		c.errWriters = append(c.errWriters, os.Stderr, outBuf)
	}

	cmd.Stdout = io.MultiWriter(c.outWriters...)
	cmd.Stderr = io.MultiWriter(c.errWriters...)

	err := cmd.Run()

	out := strings.TrimSpace(outBuf.String())

	return &CmdResults{
		Dump:  out,
		Error: err,
	}
}

func (c *Cmd) Run() *CmdResults {

	cmd, err := c.Build()
	if err != nil {
		return &CmdResults{
			Error: fmt.Errorf("error building command: %v", err),
		}
	}

	return c.RunCmd(cmd)
}

func New() *Cmd {
	return &Cmd{}
}

func Command(s string, arg ...any) *Cmd {
	p, err := resolvePath(s)

	// Prepend 's' to args so it becomes Args[0]
	// This ensures the command execution matches OS conventions
	args := append([]string{s}, ResolveArgs(arg...)...)

	b := &Cmd{
		Cmd: exec.Cmd{
			Path: p,
			Args: args,
		},
	}

	if err != nil {
		b.Cmd.Err = err
	}

	return b
}
