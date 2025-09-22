package repo

import (
	"context"

	"github.com/nelsonsaake/go/spatie/models"
	"github.com/nelsonsaake/go/strs"
	"gorm.io/gorm"
)

type Repo struct {
	db  *gorm.DB
	ctx context.Context
}

func New(db *gorm.DB, ctx context.Context) *Repo {
	return &Repo{db: db, ctx: ctx}
}

// do returns a *gorm.DB with the current context
func (r *Repo) do() *gorm.DB {
	return r.db.WithContext(r.ctx)
}

func (r *Repo) WithContext(ctx context.Context) *Repo {
	r.ctx = ctx
	return r
}

// ===================
// Query Methods
// ===================

func (r *Repo) Is(userId string, roles ...string) (bool, error) {
	var userRoles []models.UserRole
	err := r.do().
		Preload("Role").
		Where("user_id = ?", userId).
		Find(&userRoles).
		Error
	if err != nil {
		return false, err
	}

	roleSet := make(map[string]bool)
	for _, ur := range userRoles {
		roleID := ur.Role.ID
		if roleID != "" {
			roleSet[ur.Role.Name] = true
		}
	}

	for _, want := range roles {
		if !roleSet[want] {
			return false, nil
		}
	}
	return true, nil
}

func (r *Repo) IsAny(userId string, roles ...string) (bool, error) {
	var userRoles []models.UserRole
	err := r.do().
		Preload("Role").
		Where("user_id = ?", userId).
		Find(&userRoles).
		Error
	if err != nil {
		return false, err
	}

	for _, ur := range userRoles {
		roleID := ur.Role.ID
		roleName := ur.Role.Name
		if roleID != "" {
			for _, want := range roles {
				if roleName == want {
					return true, nil
				}
			}
		}
	}
	return false, nil
}

func (r *Repo) Can(userId string, perms ...string) (bool, error) {
	permsSet, err := r.getEffectivePermissions(userId)
	if err != nil {
		return false, err
	}
	for _, p := range perms {
		if !permsSet[p] {
			return false, nil
		}
	}
	return true, nil
}

func (r *Repo) CanAny(userId string, perms ...string) (bool, error) {
	permsSet, err := r.getEffectivePermissions(userId)
	if err != nil {
		return false, err
	}
	for _, p := range perms {
		if permsSet[p] {
			return true, nil
		}
	}
	return false, nil
}

func (r *Repo) getEffectivePermissions(userId string) (map[string]bool, error) {
	permsSet := make(map[string]bool)

	// Get roles for user
	var userRoles []models.UserRole
	if err := r.do().
		Preload("Role.Permissions.Permission").
		Where("user_id = ?", userId).
		Find(&userRoles).
		Error; err != nil {
		return nil, err
	}
	for _, ur := range userRoles {
		roleID := ur.Role.ID
		if roleID != "" {
			for _, rp := range ur.Role.Permissions {
				permID := rp.Permission.ID
				permName := rp.Permission.Name
				if permID != "" {
					permsSet[permName] = true
				}
			}
		}
	}

	// Get direct permissions
	var userPerms []models.UserPermission
	if err := r.do().
		Preload("Permission").
		Where("user_id = ?", userId).
		Find(&userPerms).
		Error; err != nil {
		return nil, err
	}
	for _, up := range userPerms {
		permID := up.Permission.ID
		permName := up.Permission.Name
		if permID != "" {
			permsSet[permName] = true
		}
	}

	// Remove revoked
	var revoked []models.UserPermissionRevoked
	if err := r.do().
		Preload("Permission").
		Where("user_id = ?", userId).
		Find(&revoked).
		Error; err != nil {
		return nil, err
	}
	for _, rev := range revoked {
		permID := rev.Permission.ID
		permName := rev.Permission.Name
		if permID != "" {
			delete(permsSet, permName)
		}
	}

	return permsSet, nil
}

// Management methods
func (r *Repo) CreateRole(names ...string) ([]*models.Role, error) {
	var roles []*models.Role

	for _, name := range names {
		role := &models.Role{Name: name}
		role.ID = strs.UUID()

		err := r.do().FirstOrCreate(role, models.Role{Name: name}).Error
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}

func (r *Repo) CreatePermission(names ...string) ([]*models.Permission, error) {
	var perms []*models.Permission

	for _, name := range names {
		perm := &models.Permission{Name: name}
		perm.ID = strs.UUID()

		err := r.do().FirstOrCreate(perm, models.Permission{Name: name}).Error
		if err != nil {
			return nil, err
		}

		perms = append(perms, perm)
	}

	return perms, nil
}

func (r *Repo) AssignPermissionToRole(roleName string, permNames ...string) error {
	var role models.Role
	err := r.do().Where("name = ?", roleName).First(&role).Error
	if err != nil {
		return err
	}

	for _, permName := range permNames {
		var perm models.Permission
		err := r.do().Where("name = ?", permName).First(&perm).Error
		if err != nil {
			return err
		}

		rp := models.RolePermission{RoleID: role.ID, PermissionID: perm.ID}
		rp.ID = strs.UUID()

		err = r.do().FirstOrCreate(&rp, rp).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *Repo) GiveRoleToUser(userId string, roleNames ...string) error {
	return GiveRoleToUser(r.do(), userId, roleNames...)
}

func (r *Repo) GivePermissionToUser(userId string, permNames ...string) error {
	// Ensure user exists
	var user models.User
	err := r.do().FirstOrCreate(&user, models.User{Base: models.Base{ID: userId}}).Error
	if err != nil {
		return err
	}

	for _, permName := range permNames {
		var perm models.Permission
		err := r.do().Where("name = ?", permName).First(&perm).Error
		if err != nil {
			return err
		}

		// Check if already assigned
		var count int64
		err = r.do().Model(&models.UserPermission{}).
			Where("user_id = ? AND permission_id = ?", userId, perm.ID).
			Count(&count).Error
		if err != nil {
			return err
		}
		if count > 0 {
			// Already assigned, skip
			continue
		}

		up := models.UserPermission{UserID: userId, PermissionID: perm.ID}
		up.ID = strs.UUID()

		err = r.do().Create(&up).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// RemoveRoleFromUser removes one or more roles from a user by role name
func (r *Repo) RemoveRoleFromUser(userId string, roleNames ...string) error {
	for _, roleName := range roleNames {
		var role models.Role
		err := r.do().Where("name = ?", roleName).First(&role).Error
		if err != nil {
			return err
		}

		err = r.do().Where("user_id = ? AND role_id = ?", userId, role.ID).Delete(&models.UserRole{}).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repo) RevokePermissionFromUser(userId string, permNames ...string) error {
	for _, permName := range permNames {
		var perm models.Permission
		err := r.do().Where("name = ?", permName).First(&perm).Error
		if err != nil {
			return err
		}

		rev := models.UserPermissionRevoked{UserID: userId, PermissionID: perm.ID}
		rev.ID = strs.UUID()

		err = r.do().FirstOrCreate(&rev, rev).Error
		if err != nil {
			return err
		}
	}

	return nil
}

// GetRoles returns all role names for a given user ID
func (r *Repo) GetRoles(userId string) ([]string, error) {
	var userRoles []models.UserRole
	if err := r.do().Preload("Role").Where("user_id = ?", userId).Find(&userRoles).Error; err != nil {
		return nil, err
	}
	roles := make([]string, 0, len(userRoles))
	for _, ur := range userRoles {
		if ur.Role.ID != "" {
			roles = append(roles, ur.Role.Name)
		}
	}
	return roles, nil
}

// GetPermissionsForUser returns all permission names for a given user ID
// GetDetailedPermissions returns all permission names for a given user ID
// Includes direct permissions, permissions via roles, minus revoked, with details
func (r *Repo) GetDetailedPermissions(userId string) (map[string]struct{ Direct, ViaRole, Revoked bool }, error) {
	result := make(map[string]struct{ Direct, ViaRole, Revoked bool })

	// Direct permissions
	var userPerms []models.UserPermission
	if err := r.do().Preload("Permission").Where("user_id = ?", userId).Find(&userPerms).Error; err != nil {
		return nil, err
	}
	for _, up := range userPerms {
		if up.Permission.ID != "" {
			name := up.Permission.Name
			entry := result[name]
			entry.Direct = true
			result[name] = entry
		}
	}

	// Permissions via roles
	var userRoles []models.UserRole
	if err := r.do().Preload("Role.Permissions.Permission").Where("user_id = ?", userId).Find(&userRoles).Error; err != nil {
		return nil, err
	}
	for _, ur := range userRoles {
		for _, rp := range ur.Role.Permissions {
			if rp.Permission.ID != "" {
				name := rp.Permission.Name
				entry := result[name]
				entry.ViaRole = true
				result[name] = entry
			}
		}
	}

	// Revoked permissions
	var revoked []models.UserPermissionRevoked
	if err := r.do().Preload("Permission").Where("user_id = ?", userId).Find(&revoked).Error; err != nil {
		return nil, err
	}
	for _, rev := range revoked {
		if rev.Permission.ID != "" {
			name := rev.Permission.Name
			entry := result[name]
			entry.Revoked = true
			result[name] = entry
		}
	}

	return result, nil
}

// GetPermissions returns all effective permission names for a given user ID
// Includes direct and via roles, minus revoked
func (r *Repo) GetPermissions(userId string) ([]string, error) {
	permSet := make(map[string]struct{})

	// Direct permissions
	var userPerms []models.UserPermission
	if err := r.do().Preload("Permission").Where("user_id = ?", userId).Find(&userPerms).Error; err != nil {
		return nil, err
	}
	for _, up := range userPerms {
		if up.Permission.ID != "" {
			permSet[up.Permission.Name] = struct{}{}
		}
	}

	// Permissions via roles
	var userRoles []models.UserRole
	if err := r.do().Preload("Role.Permissions.Permission").Where("user_id = ?", userId).Find(&userRoles).Error; err != nil {
		return nil, err
	}
	for _, ur := range userRoles {
		for _, rp := range ur.Role.Permissions {
			if rp.Permission.ID != "" {
				permSet[rp.Permission.Name] = struct{}{}
			}
		}
	}

	// Remove revoked permissions
	var revoked []models.UserPermissionRevoked
	if err := r.do().Preload("Permission").Where("user_id = ?", userId).Find(&revoked).Error; err != nil {
		return nil, err
	}
	for _, rev := range revoked {
		if rev.Permission.ID != "" {
			delete(permSet, rev.Permission.Name)
		}
	}

	// Convert set to slice
	perms := make([]string, 0, len(permSet))
	for name := range permSet {
		perms = append(perms, name)
	}
	return perms, nil
}
