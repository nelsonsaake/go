package models

type UserRole struct {
	ID     string `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID string `gorm:"index;not null"`
	RoleID string `gorm:"index;not null"`
	User   User   `gorm:"foreignKey:UserID"`
	Role   Role   `gorm:"foreignKey:RoleID"`
}
