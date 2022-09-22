package interfaces

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/libs"
)

type AuthService interface {
	Login(data *models.User) *libs.Response
}
