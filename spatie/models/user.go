package models

type User struct {
	Base
}

func NewUser(id string) *User {
	return &User{Base: Base{ID: id}}
}
