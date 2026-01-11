package sys

import (
	"os"
)

type builder struct {
	// wd: working directory
	wd  string
	env string
}

func (b *builder) WithWorkingDirectory(v string) *builder {
	b.wd = v
	return b
}

func (b *builder) WithEnv(v string) *builder {
	b.env = v
	return b
}

func (b *builder) Run(s string, arg ...any) (string, error) {

	cmd := NewCmd(s, arg...)
	if b.env != "" {
		cmd.Env = append(os.Environ(), b.env)
	}
	if b.wd != "" {
		cmd.Dir = b.wd
	}

	return RunCmd(cmd)
}

func New() *builder {
	return &builder{}
}

// BUILDER HELPERS

func WD(v string) *builder {
	return New().WithWorkingDirectory(v)
}

func UsingEnv(env string) *builder {
	return New().WithEnv(env)
}

func NI() *builder {
	return UsingEnv("DEBIAN_FRONTEND=noninteractive")
}
