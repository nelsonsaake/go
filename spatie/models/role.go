package models

type Role struct {
	Base
	Name        string           `gorm:"uniqueIndex"`
	Permissions []RolePermission `gorm:"foreignKey:RoleID"`
	Users       []UserRole       `gorm:"foreignKey:RoleID"`
}
