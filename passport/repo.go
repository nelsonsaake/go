package passport

import (
	"github.com/nelsonsaake/go/passport/models"
	"github.com/nelsonsaake/go/passport/repo"
	"github.com/nelsonsaake/go/passport/repo/auths"
)

type Repo = repo.Repo

type AuthRepo = auths.Repo

type User struct{}

func (u User) GetID() string { return "" }

type UserRepo struct{}

func (r UserRepo) Find(id string) (models.User, error) {
	return User{}, nil
}

// getRepo get repo instance or panic if not initialized
func getRepo() *Repo {
	if cfgInstance == nil {
		panic("passport instance not initialized, call Setup() first")
	}

	return repo.New(cfgInstance.DB)
}

func getAuthRepo() *AuthRepo {
	return getRepo().Auths
}

func getUserRepo() *UserRepo {
	return &UserRepo{}
}
