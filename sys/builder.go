package sys

import (
	"os"
	"os/exec"
)

type builder struct {
	Path string
	Arg  []any

	// wd: working directory
	wd  string
	env string
}

func (b *builder) WithWorkingDirectory(v string) *builder {
	b.wd = v
	return b
}

func (b *builder) WithEnv(k, v string) *builder {
	b.env = k + "=" + v
	return b
}

func (b *builder) Env(k, v string) *builder {
	return b.WithEnv(k, v)
}

func (b *builder) WithArgs(arg ...any) *builder {
	b.Arg = arg
	return b
}

func (b *builder) Command(path string, arg ...any) *builder {
	b.Path, b.Arg = path, arg
	return b
}

func (b *builder) WD(v string) *builder {
	return b.WithWorkingDirectory(v)
}

func (b *builder) NI() *builder {
	return b.WithEnv("DEBIAN_FRONTEND", "noninteractive")
}

func (b *builder) Build() *exec.Cmd {

	cmd := exec.Command(
		b.Path,
		resolveArgs(b.Arg...)...,
	)

	if b.env != "" {
		cmd.Env = append(os.Environ(), b.env)
	}
	if b.wd != "" {
		cmd.Dir = b.wd
	}

	return cmd
}

func (b *builder) Run() (string, error) {
	return runCmd(b.Build())
}

func (b *builder) Ok() bool {
	return b.Build().Run() == nil
}

func New() *builder {
	return &builder{}
}

func Cmd(s string, arg ...any) *builder {
	return &builder{
		Path: s,
		Arg:  arg,
	}
}
