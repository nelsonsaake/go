package spatie

import (
	"context"

	"gorm.io/gorm"
)

type Config struct {
	DB      *gorm.DB
	Context context.Context
}

func Setup(cfg Config) *Repo {
	return connect(cfg.DB, cfg.Context)
}
