package passport

import (
	"fmt"

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

func Issue(user models.User) (*auths.Issue, error) {
	return getRepo().Auths.Issue(user)
}

func Delete(id string) error {
	return getRepo().Auths.Delete(id)
}

func DeleteWhereUserID(userID string) error {
	return getRepo().Auths.DeleteWhereUserID(userID)
}

func Refresh(user models.User) (*auths.Issue, error) {

	err := getRepo().Auths.DeleteWhereUserID(user.GetID())
	if err != nil {
		return nil, fmt.Errorf("error delete existing auths: %w", err)
	}

	auth, err := getRepo().Auths.Issue(user)
	if err != nil {
		return nil, fmt.Errorf("error issuing new auths: %w", err)
	}

	return auth, nil
}
