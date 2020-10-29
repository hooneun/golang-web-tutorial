package models

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primaryKey"`
	Email      string `json:"email" gorm:"index:idx_email,unique"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Todos      []Todo
}

// GetUser !
func (u *User) GetUser() User {
	return User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
		Todos: u.Todos,
	}
}

// Todo struct
type Todo struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id" gorm:"primaryKey"`
	UserID     uint   `json:"user_id" gorm:"index:idx_user"`
	Desc       string `json:"desc" gorm:"index:idx_desc"`
}
