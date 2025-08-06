package app

import (
	"fmt"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/nelsonsaake/go/afs"
)

var (
	env = ".env"
)

func Boot() error {

	cwd, err := afs.Root()
	if err == nil {
		env = filepath.Join(cwd, env)
	}

	err = godotenv.Load(env)
	if err != nil {
		return fmt.Errorf("error loading %s: %v", env, err)
	}

	for name, setup := range setups() {
		fmt.Println("setting up", name, "...")
		err = setup.Setup()
		if err != nil {
			return fmt.Errorf("error setting up %s: %v", name, err)
		}
		fmt.Println("setyp", name, "completed successfully")
	}

	return nil
}
