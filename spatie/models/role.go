package models

type Role struct {
	Base
	Name        string           `gorm:"uniqueIndex;size:191"`
	Permissions []RolePermission `gorm:"foreignKey:RoleID"`
	Users       []UserRole       `gorm:"foreignKey:RoleID"`
}
