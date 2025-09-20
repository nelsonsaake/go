package fctx

import "github.com/nelsonsaake/go/passport/models"

type AuthRepo interface {
	Find(string) (*models.Auth, error)
}

// everyone is going to implement their own user repo
type UserRepo interface {
	Find(string) (models.User, error)
}
