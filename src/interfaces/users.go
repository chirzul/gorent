package interfaces

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
)

type UserRepo interface {
	FindAllUsers() (*models.Users, error)
	SaveUser(data *models.User) (*models.User, error)
	ChangeUser(data *models.User, email string) (*models.User, error)
	RemoveUser(data *models.User, email string) (*models.User, error)
}

type UserService interface {
	GetAllUsers() (*models.Users, error)
	AddUser(data *models.User) (*models.User, error)
	UpdateUser(data *models.User, email string) (*models.User, error)
	DeleteUser(data *models.User, email string) (*models.User, error)
}
