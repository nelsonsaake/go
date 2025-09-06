package models

type Permission struct {
	ID        string                  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name      string                  `gorm:"uniqueIndex"`
	Roles     []RolePermission        `gorm:"foreignKey:PermissionID"`
	Users     []UserPermission        `gorm:"foreignKey:PermissionID"`
	RevokedBy []UserPermissionRevoked `gorm:"foreignKey:PermissionID"`
}
