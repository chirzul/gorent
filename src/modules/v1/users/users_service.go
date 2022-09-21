package users

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/interfaces"
	"github.com/chirzul/gorent/src/libs"
)

type user_service struct {
	repo interfaces.UserRepo
}

func NewService(repo interfaces.UserRepo) *user_service {
	return &user_service{repo}
}

func (s *user_service) GetAllUsers() *libs.Response {
	data, err := s.repo.FindAllUsers()
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *user_service) AddUser(data *models.User) *libs.Response {
	if checkUsername := s.repo.CheckUsername(data.Username); checkUsername {
		return libs.GetResponse("username is already registered", 400, true)
	}
	if checkEmail := s.repo.CheckEmail(data.Email); checkEmail {
		return libs.GetResponse("email is already registered", 400, true)
	}
	data, err := s.repo.SaveUser(data)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *user_service) UpdateUser(data *models.User, username string) *libs.Response {
	if checkUsername := s.repo.CheckUsername(username); !checkUsername {
		return libs.GetResponse("username not found", 404, true)
	}
	data, err := s.repo.ChangeUser(data, username)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}

func (s *user_service) DeleteUser(data *models.User, username string) *libs.Response {
	if checkUsername := s.repo.CheckUsername(username); !checkUsername {
		return libs.GetResponse("username not found", 404, true)
	}
	data, err := s.repo.RemoveUser(data, username)
	if err != nil {
		return libs.GetResponse(err.Error(), 400, true)
	}
	return libs.GetResponse(data, 200, false)
}
