package repo

import (
	"github.com/nelsonsaake/go/spatie/models"
	"github.com/nelsonsaake/go/strs"
	"gorm.io/gorm"
)

func GiveRoleToUser(db *gorm.DB, userId string, roleNames ...string) error {

	// Ensure user exists
	var (
		user    models.User
		newUser = models.NewUser(userId)
	)

	err := db.Model(&models.User{}).FirstOrCreate(&user, *newUser).Error
	if err != nil {
		return err
	}

	for _, roleName := range roleNames {
		var role models.Role
		err := db.Model(&models.Role{}).Where("name = ?", roleName).First(&role).Error
		if err != nil {
			return err
		}

		// Check if already assigned
		var count int64
		err = db.Model(&models.UserRole{}).
			Where("user_id = ? AND role_id = ?", userId, role.ID).
			Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			// Already assigned, skip
			continue
		}

		ur := models.UserRole{UserID: userId, RoleID: role.ID}
		ur.ID = strs.UUID()

		err = db.Create(&ur).Error
		if err != nil {
			return err
		}
	}

	return nil
}
