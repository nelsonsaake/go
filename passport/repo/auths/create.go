package auths

import (
	"github.com/nelsonsaake/go/passport/models"
)

func (r *Repo) create(userID string) (*models.Auth, error) {

	auth := &models.Auth{UserID: userID}

	err := r.DB.Create(auth).Error
	if err != nil {
		return nil, err
	}

	return auth, nil
}

func (r *Repo) Create(userID string) (*models.Auth, error) {

	// KEEP USER WITHIN THE LIMITS OF ALLOWABLE SESSIONS

	// count, err := r.Count(userID)
	// if err != nil {
	// 	return nil, err
	// }

	// if count >= env.MaxSessionsPerUserAsInt() {

	// 	err := r.DeleteAllUserAuth(userID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// }

	// ACTIVATE AUTH BY SAVING IT

	auth, err := r.create(userID)
	if err != nil {
		return nil, err
	}

	return auth, nil
}
