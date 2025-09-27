package fctx

import "github.com/nelsonsaake/go/passport/models"

type Repo[T any] interface {
	Find(string) (*T, error)
}

type AuthRepo interface {
	Repo[models.Auth]
}

// everyone is going to implement their own user repo
type UserRepo interface {
	Find(string) (models.User, error)
}
