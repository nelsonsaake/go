package models

type Permission struct {
	Base
	Name      string                  `gorm:"uniqueIndex"`
	Roles     []RolePermission        `gorm:"foreignKey:PermissionID"`
	Users     []UserPermission        `gorm:"foreignKey:PermissionID"`
	RevokedBy []UserPermissionRevoked `gorm:"foreignKey:PermissionID"`
}
