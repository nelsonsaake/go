package crypt

type Config struct {
	Secret string
}

func Setup(cfg Config) {
	secret = cfg.Secret
}
