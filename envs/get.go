package envs

import "os"

func Get(k string) string {
	return os.Getenv(k)
}
