package sys

import (
	"os/exec"
)

type builder struct {
	exec.Cmd
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

func (b *builder) Build() *exec.Cmd {
	return &b.Cmd
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

func Command(s string, arg ...any) *builder {
	return &builder{
		Cmd: exec.Cmd{
			Path: s,
			Args: resolveArgs(arg...),
		},
	}
}
