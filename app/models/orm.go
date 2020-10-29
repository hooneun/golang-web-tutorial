package models

import (
	"errors"

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
	err = db.Find(&user, id).Error

	return user.getUser(), err
}

// CreateUser - user create
func (db *DBORM) CreateUser(user User) (User, error) {
	helpers.HashPassword(&user.Password)
	err := db.Create(&user).Error

	return user.getUser(), err
}

// SignInUser user signin
func (db *DBORM) SignInUser(email string, password string) (user User, err error) {
	err = db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}

	if !helpers.CheckPasswordHash(password, user.Password) {
		return user, errors.New("Invalid password")
	}

	return user.getUser(), nil
}
