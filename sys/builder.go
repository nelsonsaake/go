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

type Builder struct {
	exec.Cmd
	file                   *string
	outWriters             []io.Writer
	errWriters             []io.Writer
	dump                   *string
	disableDefaultWritters bool
}

func (b *Builder) WithWorkingDirectory(v string) *Builder {
	b.Dir = v
	return b
}

func (b *Builder) WithEnv(k, v string) *Builder {
	b.Cmd.Env = append(b.Cmd.Env, k+"="+v)
	return b
}

func (b *Builder) Env(k, v string) *Builder {
	return b.WithEnv(k, v)
}

func (b *Builder) WithArgs(arg ...any) *Builder {
	b.Args = ResolveArgs(arg...)
	return b
}

func (b *Builder) WithDump(v *string) *Builder {
	b.dump = v
	return b
}

func (b *Builder) WithFile(v string) *Builder {

	if !filepath.IsAbs(v) {
		v = afs.Path(v)
	}

	b.file = &v

	return b
}

func (b *Builder) Command(path string, arg ...any) *Builder {
	b.Path, b.Args = path, ResolveArgs(arg...)
	return b
}

func (b *Builder) WD(v string) *Builder {
	return b.WithWorkingDirectory(v)
}

func (b *Builder) NI() *Builder {
	return b.WithEnv("DEBIAN_FRONTEND", "noninteractive")
}

func (b *Builder) Build() (*exec.Cmd, error) {

	die := func(f string, a ...any) (*exec.Cmd, error) {
		return nil, fmt.Errorf(f, a...)
	}

	if b.file == nil {
		return &b.Cmd, nil
	}

	// TODO IMPLEMENT FILE SECTION

	f, err := os.Open(*b.file)
	if err != nil {
		return die("failed to open file %q: %w", *b.file, err)
	}

	b.Stdin = f

	return &b.Cmd, nil
}

func (b *Builder) Run() error {

	cmd, err := b.Build()
	if err != nil {
		return err
	}

	var outBuf *strings.Builder

	if !b.disableDefaultWritters {
		b.outWriters = append(b.outWriters, os.Stdout)
		b.errWriters = append(b.errWriters, os.Stderr)
	}

	if b.dump != nil {
		outBuf = &strings.Builder{}
		b.outWriters = append(b.outWriters, outBuf)
		b.errWriters = append(b.errWriters, outBuf)
	}

	cmd.Stdout = io.MultiWriter(b.outWriters...)
	cmd.Stderr = io.MultiWriter(b.errWriters...)

	err = cmd.Run()

	if b.dump != nil {
		out := strings.TrimSpace(outBuf.String())
		*b.dump = out
	}

	return err
}

func (b *Builder) Runo() (string, error) {
	var dump string
	var err error = b.WithDump(&dump).Run()

	return dump, err
}

func (b *Builder) Ok() bool {
	return b.Run() == nil
}

func New() *Builder {
	return &Builder{}
}

func Command(s string, arg ...any) *Builder {
	p, err := resolvePath(s)

	// Prepend 's' to args so it becomes Args[0]
	// This ensures the command execution matches OS conventions
	args := append([]string{s}, ResolveArgs(arg...)...)

	b := &Builder{
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
