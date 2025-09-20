package passport

import (
	"context"

	"gorm.io/gorm"
)

var (
	cfgInstance *Config
)

type Config struct {
	DB           *gorm.DB
	Context      context.Context
	TokenConfigs map[string]TokenConfig
}

func Setup(cfg Config) {
	cfgInstance = &cfg
}
