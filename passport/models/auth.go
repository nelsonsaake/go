package models

import (
	"github.com/nelsonsaake/go/strs"
	"gorm.io/gorm"
)

type Auth struct {
	ID     string `gorm:"primaryKey;type:char(36);not null"`
	UserID string `gorm:"type:char(36);not null;index"`
}

func (a Auth) GetID() string {
	return a.ID
}

func (a *Auth) SetID(id string) {
	a.ID = id
}

func (a *Auth) BeforeCreate(tx *gorm.DB) (err error) {
	if a.ID == "" {
		a.ID = strs.UUID()
	}
	return nil
}
