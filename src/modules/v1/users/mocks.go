package users

import (
	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/stretchr/testify/mock"
)

type RepoMock struct {
	mock mock.Mock
}

func (m *RepoMock) GetAllUsers() (*models.Users, error) {
	args := m.mock.Called()
	return args.Get(0).(*models.Users), nil
}

func (m *RepoMock) FindByUsername(username string) (*models.User, error) {
	args := m.mock.Called(username)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) AddUser(data *models.User) (*models.User, error) {
	args := m.mock.Called(data)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) UpdateUser(data *models.User, username string) (*models.User, error) {
	args := m.mock.Called(data, username)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) DeleteUser(data *models.User, username string) (*models.User, error) {
	args := m.mock.Called(data, username)
	return args.Get(0).(*models.User), nil
}

func (m *RepoMock) CheckUsername(username string) (result bool) {
	return
}

func (m *RepoMock) CheckEmail(email string) (result bool) {
	return
}
