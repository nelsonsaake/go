package models

type UserRole struct {
	Base
	UserID string `gorm:"index;not null"`
	RoleID string `gorm:"index;not null"`
	User   User   `gorm:"foreignKey:UserID"`
	Role   Role   `gorm:"foreignKey:RoleID"`
}
