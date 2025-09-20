package auths

import (
	"fmt"

	"github.com/nelsonsaake/go/passport/models"
	"github.com/nelsonsaake/go/passport/tokens"
	"github.com/nelsonsaake/go/spatie"
)

type Issue struct {
	User         models.User `json:"user"`
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
}

func (r *Repo) Issue(user models.User) (*Issue, error) {

	die := func(f string, err error) (*Issue, error) {
		return nil, fmt.Errorf(f, err)
	}

	roles, err := spatie.GetRoles(user.GetID())
	if err != nil {
		return die("error getting roles: %w", err)
	}

	rec, err := r.Create(user.GetID())
	if err != nil {
		return die("error creating auth: %w", err)
	}

	rolesClaim := map[string]any{
		"roles": roles,
	}

	at, err := tokens.New("access", rec.ID, rolesClaim)
	if err != nil {
		return die("error creating access token: %w", err)
	}

	rt, err := tokens.New("refresh", rec.ID, rolesClaim)
	if err != nil {
		return die("error creating refresh token: %w", err)
	}

	res := Issue{
		User:         user,
		AccessToken:  at,
		RefreshToken: rt,
	}

	return &res, nil
}
