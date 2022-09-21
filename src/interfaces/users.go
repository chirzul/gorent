package interfaces

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/libs"
)

type UserRepo interface {
	FindAllUsers() (*models.Users, error)
	SaveUser(data *models.User) (*models.User, error)
	ChangeUser(data *models.User, username string) (*models.User, error)
	RemoveUser(data *models.User, username string) (*models.User, error)
	CheckUsername(username string) bool
	CheckEmail(email string) bool
}

type UserService interface {
	GetAllUsers() *libs.Response
	AddUser(data *models.User) *libs.Response
	UpdateUser(data *models.User, username string) *libs.Response
	DeleteUser(data *models.User, username string) *libs.Response
}
