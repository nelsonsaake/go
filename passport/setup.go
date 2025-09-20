package passport

import (
	"context"
	"time"

	"gorm.io/gorm"
)

var (
	cfgInstance *Config
)

type Config struct {
	DB      *gorm.DB
	Context context.Context
	Tokens  map[string]TokenConfig
}

func Setup(cfg Config) {

	cfgInstance = &cfg

	if cfg.Tokens != nil {
		return
	}

	cfg.Tokens = map[string]TokenConfig{
		"access": {
			Name:          "access",
			TTL:           time.Minute * 5, // 5 minutes
			SigningKey:    "a6823c08-65de-4599-8698-64e78da6e2db",
			SigningMethod: "HS256",
		},
		"refresh": {
			Name:          "refresh",
			TTL:           time.Minute * 60 * 24 * 30 * 6, // six months
			SigningKey:    "4c0b1f88-ea83-46c0-8ce1-7da3a2cf5096",
			SigningMethod: "HS256",
		},
		"reset": {
			Name:          "reset",
			TTL:           time.Minute * 30, // 30 minutes
			SigningKey:    "f5884751-dfd0-4054-99de-af33162e7075",
			SigningMethod: "HS256",
		},
	}
}
