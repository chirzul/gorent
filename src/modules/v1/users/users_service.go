package users

import (
	"net/http"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/chirzul/gorent/src/interfaces"
)

type user_service struct {
	repo interfaces.UserRepo
}

func NewService(repo interfaces.UserRepo) *user_service {
	return &user_service{repo}
}

func (s *user_service) GetAllUsers() (*models.Users, error) {
	data, err := s.repo.FindAllUsers()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *user_service) AddUser(r *http.Request, data *models.User) (*models.User, error) {
	data, err := s.repo.SaveUser(r, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *user_service) UpdateUser(r *http.Request, data *models.User) (*models.User, error) {
	data, err := s.repo.ChangeUser(r, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *user_service) DeleteUser(r *http.Request, data *models.User) (*models.User, error) {
	data, err := s.repo.RemoveUser(r, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
