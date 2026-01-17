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
	file                   *string
	outWriters             []io.Writer
	errWriters             []io.Writer
	dump                   *string
	disableDefaultWritters bool
	error                  error
}

func (c *Cmd) WithWorkingDirectory(v string) *Cmd {
	c.Dir = v
	return c
}

func (c *Cmd) WithEnv(k, v string) *Cmd {
	c.Cmd.Env = append(c.Cmd.Env, k+"="+v)
	return c
}

func (c *Cmd) Env(k, v string) *Cmd {
	return c.WithEnv(k, v)
}

func (c *Cmd) WithArgs(arg ...any) *Cmd {
	c.Args = ResolveArgs(arg...)
	return c
}

func (c *Cmd) WithDump(v *string) *Cmd {
	c.dump = v
	return c
}

func (c *Cmd) WithFile(v string) *Cmd {

	if !filepath.IsAbs(v) {
		v = afs.Path(v)
	}

	c.file = &v

	return c
}

func (c *Cmd) WD(v string) *Cmd {
	return c.WithWorkingDirectory(v)
}

func (c *Cmd) NI() *Cmd {
	return c.WithEnv("DEBIAN_FRONTEND", "noninteractive")
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

func (c *Cmd) Run() error {

	cmd, err := c.Build()
	if err != nil {
		return err
	}

	var outBuf *strings.Builder

	if !c.disableDefaultWritters {
		c.outWriters = append(c.outWriters, os.Stdout)
		c.errWriters = append(c.errWriters, os.Stderr)
	}

	if c.dump != nil {
		outBuf = &strings.Builder{}
		c.outWriters = append(c.outWriters, outBuf)
		c.errWriters = append(c.errWriters, outBuf)
	}

	cmd.Stdout = io.MultiWriter(c.outWriters...)
	cmd.Stderr = io.MultiWriter(c.errWriters...)

	err = cmd.Run()

	if c.dump != nil {
		out := strings.TrimSpace(outBuf.String())
		*c.dump = out
	}

	return err
}

func (c *Cmd) Runo() (string, error) {
	var dump string
	var err error = c.WithDump(&dump).Run()

	return dump, err
}

func (c *Cmd) Dump(v string) *Dump {

	o, err := c.Runo()

	return &Dump{
		Output: o,
		Error:  err,
	}
}

func (c *Cmd) OK() bool {
	return c.Run() == nil
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
