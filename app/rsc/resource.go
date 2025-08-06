package rsc

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/nelsonsaake/go/afs"
)

type EnvLoader struct {
	paths []string
}

func (r EnvLoader) load(env string) error {

	cwd, err := afs.Root()
	if err == nil {
		env = filepath.Join(cwd, env)
	}

	err = godotenv.Load(env)
	if err != nil {
		return fmt.Errorf("error loading %s: %v", env, err)
	}

	return nil
}

func (r EnvLoader) Setup() error {

	for _, path := range r.paths {
		r.load(path)
	}

	return nil
}

func NewEnvLoader(paths ...string) *EnvLoader {
	if len(paths) == 0 {
		paths = []string{".env"}
	}
	return &EnvLoader{
		paths: paths,
	}
}
