package auths

import (
	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *Repo {
	return &Repo{DB: db}
}
