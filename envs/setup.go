package envs

type Config struct {
	Paths []string
}

func Setup(cfg Config) error {
	return Load(cfg.Paths...)
}
