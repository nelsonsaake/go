package models

type UserPermissionRevoked struct {
	Base
	UserID       string     `gorm:"index;not null"`
	PermissionID string     `gorm:"index;not null"`
	User         User       `gorm:"foreignKey:UserID"`
	Permission   Permission `gorm:"foreignKey:PermissionID"`
}
