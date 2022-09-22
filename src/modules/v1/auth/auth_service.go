package auth

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/interfaces"
	"github.com/chirzul/gorent/src/libs"
)

type auth_service struct {
	repo interfaces.UserRepo
}

type token_response struct {
	Token string `json:"token"`
}

func NewService(repo interfaces.UserRepo) *auth_service {
	return &auth_service{repo}
}

func (s *auth_service) Login(data *models.User) *libs.Response {
	user, err := s.repo.FindByUsername(data.Username)
	if err != nil {
		return libs.GetResponse("user is not exist", 401, true)
	}

	if !libs.CheckPassword(user.Password, data.Password) {
		return libs.GetResponse("password false", 401, true)
	}

	token := libs.NewToken(data.Username, user.Role)
	theToken, err := token.Create()
	if err != nil {
		return libs.GetResponse(err.Error(), 401, true)
	}
	return libs.GetResponse(token_response{Token: theToken}, 200, false)
}
