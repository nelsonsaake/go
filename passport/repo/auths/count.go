package auths

import "github.com/nelsonsaake/go/passport/models"

func (r *Repo) Count(userID string) (int64, error) {

	var count int64

	err := r.DB.
		Model(&models.Auth{}).
		Where("user_id = ?", userID).
		Count(&count).
		Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
