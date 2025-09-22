package spatie

import (
	"context"

	"github.com/nelsonsaake/go/spatie/repo"
	"gorm.io/gorm"
)

func Tx(tx *gorm.DB) *Repo {
	return repo.New(tx, context.Background())
}
