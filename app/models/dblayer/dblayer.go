package dblayer

import "github.com/hooneun/golang-web-tutorial/app/models"

// DBLayer !
type DBLayer interface {
	GetUserByID(int) (models.User, error)
	CreateUser(models.User) (models.User, error)
}
