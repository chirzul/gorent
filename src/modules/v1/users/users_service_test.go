package users

import (
	"fmt"
	"testing"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUsers(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.Users{
		{
			UserID:   "66f10701-b983-496c-8c9c-a99b8b4452b2",
			Username: "chir",
			Email:    "chir@gmail.com",
		},
		{
			UserID:   "66f10701-b983-496c-8c9c-a99b8b4452b3",
			Username: "chirz",
			Email:    "chirz@gmail.com",
		},
	}

	repo.mock.On("GetAllUsers").Return(&dataMock, nil)
	data := service.GetAllUsers()

	result := data.Data.(*models.Users)

	for i, v := range *result {
		assert.Equal(t, dataMock[i].UserID, v.UserID, "expect id from data mock")
	}
}

func TestAddUser(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.User{
		UserID:   "66f10701-b983-496c-8c9c-a99b8b4452b89",
		Username: "a",
		Email:    "a@gmail.com",
	}

	repo.mock.On("AddUser", &dataMock).Return(&dataMock, nil)
	data := service.AddUser(&dataMock)

	result := data.Data
	fmt.Println(result)
	// assert.Equal(t, dataMock.UserID, result.UserID, "expect id from data mock")
}

func TestFindByUsername(t *testing.T) {
	repo := RepoMock{mock.Mock{}}
	service := NewService(&repo)

	var dataMock = models.User{
		UserID:   "66f10701-b983-496c-8c9c-a99b8b4452b2",
		Username: "chir",
		Email:    "chir@gmail.com",
	}

	username := "chir"
	repo.mock.On("FindByUsername", username).Return(&dataMock, nil)
	data := service.FindByUsername(username)

	result := data.Data.(*models.User)

	assert.Equal(t, result.Username, username, "expect id from data mock")
}
