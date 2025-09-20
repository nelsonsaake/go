package auths

import "github.com/nelsonsaake/go/passport/models"

func (r *Repo) Find(id string) (*models.Auth, error) {
	var auth models.Auth
	if err := r.DB.First(&auth, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}
