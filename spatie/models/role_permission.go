package models

type RolePermission struct {
	Base
	RoleID       string     `gorm:"index;not null"`
	PermissionID string     `gorm:"index;not null"`
	Role         Role       `gorm:"foreignKey:RoleID"`
	Permission   Permission `gorm:"foreignKey:PermissionID"`
}
