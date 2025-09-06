package engine

import (
	"context"

	"github.com/nelsonsaake/go/spatie/models"
	"gorm.io/gorm"
)

type Spatie struct {
	db  *gorm.DB
	ctx context.Context
}

func New(db *gorm.DB, ctx context.Context) *Spatie {
	return &Spatie{db: db, ctx: ctx}
}

// do returns a *gorm.DB with the current context
func (s *Spatie) do() *gorm.DB {
	return s.db.WithContext(s.ctx)
}

func (s *Spatie) WithContext(ctx context.Context) *Spatie {
	s.ctx = ctx
	return s
}

// ===================
// Query Methods
// ===================

func (s *Spatie) Is(userId string, roles ...string) (bool, error) {
	var userRoles []models.UserRole
	err := s.do().
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

func (s *Spatie) IsAny(userId string, roles ...string) (bool, error) {
	var userRoles []models.UserRole
	err := s.do().
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

func (s *Spatie) Can(userId string, perms ...string) (bool, error) {
	permsSet, err := s.getEffectivePermissions(userId)
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

func (s *Spatie) CanAny(userId string, perms ...string) (bool, error) {
	permsSet, err := s.getEffectivePermissions(userId)
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

func (s *Spatie) getEffectivePermissions(userId string) (map[string]bool, error) {
	permsSet := make(map[string]bool)

	// Get roles for user
	var userRoles []models.UserRole
	if err := s.do().
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
	if err := s.do().
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
	if err := s.do().
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
func (s *Spatie) CreateRole(names ...string) ([]*models.Role, error) {
	var roles []*models.Role
	for _, name := range names {
		role := &models.Role{Name: name}
		if err := s.do().
			FirstOrCreate(role, models.Role{Name: name}).
			Error; err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (s *Spatie) CreatePermission(names ...string) ([]*models.Permission, error) {
	var perms []*models.Permission
	for _, name := range names {
		perm := &models.Permission{Name: name}
		if err := s.do().
			FirstOrCreate(perm, models.Permission{Name: name}).
			Error; err != nil {
			return nil, err
		}
		perms = append(perms, perm)
	}
	return perms, nil
}

func (s *Spatie) AssignPermissionToRole(roleName string, permNames ...string) error {
	var role models.Role
	if err := s.do().
		Where("name = ?", roleName).
		First(&role).
		Error; err != nil {
		return err
	}
	for _, permName := range permNames {
		var perm models.Permission
		if err := s.do().
			Where("name = ?", permName).
			First(&perm).
			Error; err != nil {
			return err
		}
		rp := models.RolePermission{RoleID: role.ID, PermissionID: perm.ID}
		if err := s.do().
			FirstOrCreate(&rp, rp).
			Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *Spatie) GiveRoleToUser(userId string, roleNames ...string) error {
	for _, roleName := range roleNames {
		var role models.Role
		if err := s.do().
			Where("name = ?", roleName).
			First(&role).
			Error; err != nil {
			return err
		}
		ur := models.UserRole{UserID: userId, RoleID: role.ID}
		if err := s.do().
			FirstOrCreate(&ur, ur).
			Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *Spatie) GivePermissionToUser(userId string, permNames ...string) error {
	for _, permName := range permNames {
		var perm models.Permission
		if err := s.do().
			Where("name = ?", permName).
			First(&perm).
			Error; err != nil {
			return err
		}
		up := models.UserPermission{UserID: userId, PermissionID: perm.ID}
		if err := s.do().
			FirstOrCreate(&up, up).
			Error; err != nil {
			return err
		}
	}
	return nil
}

func (s *Spatie) RevokePermissionFromUser(userId string, permNames ...string) error {
	for _, permName := range permNames {
		var perm models.Permission
		if err := s.do().
			Where("name = ?", permName).
			First(&perm).
			Error; err != nil {
			return err
		}
		rev := models.UserPermissionRevoked{UserID: userId, PermissionID: perm.ID}
		if err := s.do().
			FirstOrCreate(&rev, rev).
			Error; err != nil {
			return err
		}
	}
	return nil
}

// ...existing code...
