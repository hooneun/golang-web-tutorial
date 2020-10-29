package models

import (
	"github.com/hooneun/golang-web-tutorial/app/helpers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DBORM db orm
type DBORM struct {
	*gorm.DB
}

// NewORM !
func NewORM() (*DBORM, error) {
	dns := "root:root@tcp(127.0.0.1:4444)/books?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	return &DBORM{
		DB: db,
	}, err
}

// GetUserByID - user data of id
func (db *DBORM) GetUserByID(id int) (user User, err error) {
	return user, db.Find(&user, id).Error
}

// CreateUser - user create
func (db *DBORM) CreateUser(user User) (User, error) {
	helpers.HashPassword(&user.Password)
	return user, db.Create(&user).Error
}
