package models

type Role struct {
	ID          string           `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	Name        string           `gorm:"uniqueIndex"`
	Permissions []RolePermission `gorm:"foreignKey:RoleID"`
	Users       []UserRole       `gorm:"foreignKey:RoleID"`
}
