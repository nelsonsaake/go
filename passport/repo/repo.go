package repo

import (
	"github.com/nelsonsaake/go/passport/repo/auths"
	"gorm.io/gorm"
)

type Repo struct {
	Auths *auths.Repo
}

func New(db *gorm.DB) *Repo {
	return &Repo{
		Auths: auths.New(db),
	}
}
