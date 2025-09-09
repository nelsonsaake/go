package models

import (
	"time"
)

type Base struct {
	// ID        string    `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	ID        string    `gorm:"primaryKey;type:char(36);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
