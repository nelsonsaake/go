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

type builder struct {
	exec.Cmd
	file                   *string
	outWriters             []io.Writer
	errWriters             []io.Writer
	dump                   *string
	disableDefaultWritters bool
}

func (b *builder) WithWorkingDirectory(v string) *builder {
	b.Dir = v
	return b
}

func (b *builder) WithEnv(k, v string) *builder {
	b.Cmd.Env = append(b.Cmd.Env, k+"="+v)
	return b
}

func (b *builder) Env(k, v string) *builder {
	return b.WithEnv(k, v)
}

func (b *builder) WithArgs(arg ...any) *builder {
	b.Args = resolveArgs(arg...)
	return b
}

func (b *builder) WithDump(v *string) *builder {
	b.dump = v
	return b
}

func (b *builder) WithFile(v string) *builder {

	if !filepath.IsAbs(v) {
		v = afs.Path(v)
	}

	b.file = &v

	return b
}

func (b *builder) Command(path string, arg ...any) *builder {
	b.Path, b.Args = path, resolveArgs(arg...)
	return b
}

func (b *builder) WD(v string) *builder {
	return b.WithWorkingDirectory(v)
}

func (b *builder) NI() *builder {
	return b.WithEnv("DEBIAN_FRONTEND", "noninteractive")
}

func (b *builder) Build() (*exec.Cmd, error) {

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

func (b *builder) Run() error {

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

func (b *builder) Runo() (string, error) {
	var dump string
	var err error = b.WithDump(&dump).Run()

	return dump, err
}

func (b *builder) Ok() bool {
	return b.Run() == nil
}

func New() *builder {
	return &builder{}
}

func Command(s string, arg ...any) *builder {
	p, err := resolvePath(s)

	// Prepend 's' to args so it becomes Args[0]
	// This ensures the command execution matches OS conventions
	args := append([]string{s}, resolveArgs(arg...)...)

	b := &builder{
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
