package auths

import "github.com/nelsonsaake/go/passport/models"

func (r *Repo) Delete(id string) error {
	return r.DB.Delete(&models.Auth{}, "id = ?", id).Error
}

func (r *Repo) DeleteWhereUserID(userID string) error {
	return r.DB.Delete(&models.Auth{}, "user_id = ?", userID).Error
}
