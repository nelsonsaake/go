package models

type UserPermissionRevoked struct {
	ID           string     `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	UserID       string     `gorm:"index;not null"`
	PermissionID string     `gorm:"index;not null"`
	User         User       `gorm:"foreignKey:UserID"`
	Permission   Permission `gorm:"foreignKey:PermissionID"`
}
