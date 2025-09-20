package passport

import (
	"context"

	"gorm.io/gorm"
)

var (
	cfgInstance *Config
)

type Config struct {
	DB      *gorm.DB
	Context context.Context
	Tokens  []TokenConfig
}

func Setup(cfg Config) {

	cfgInstance = &cfg

	if cfg.Tokens != nil {
		for _, token := range cfg.Tokens {
			cfgsRegister(token.Name, token)
		}
	}
}
