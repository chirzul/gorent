package interfaces

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/libs"
)

type UserRepo interface {
	GetAllUsers() (*models.Users, error)
	AddUser(data *models.User) (*models.User, error)
	UpdateUser(data *models.User, username string) (*models.User, error)
	DeleteUser(data *models.User, username string) (*models.User, error)
	CheckUsername(username string) bool
	CheckEmail(email string) bool
}

type UserService interface {
	GetAllUsers() *libs.Response
	AddUser(data *models.User) *libs.Response
	UpdateUser(data *models.User, username string) *libs.Response
	DeleteUser(data *models.User, username string) *libs.Response
}
