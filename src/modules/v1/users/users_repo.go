package users

import (
	"errors"
	"net/http"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *user_repo {
	return &user_repo{db}
}

func (repo *user_repo) FindAllUsers() (*models.Users, error) {
	var data models.Users

	result := repo.db.Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed to get data user")
	}
	if len(data) == 0 {
		return nil, errors.New("data user is empty")
	}

	return &data, nil
}

func (repo *user_repo) SaveUser(r *http.Request, data *models.User) (*models.User, error) {
	var check models.Users
	repo.db.Where("name = ?", data.Name).Find(&check)
	if len(check) > 0 {
		return nil, errors.New("username already exist")
	}

	result := repo.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("failed to add data user")
	}

	return data, nil
}

func (repo *user_repo) ChangeUser(r *http.Request, data *models.User) (*models.User, error) {
	vars := mux.Vars(r)

	result := repo.db.Model(&data).Where("name = ?", vars["name"]).Updates(data)
	if result.Error != nil {
		return nil, errors.New("failed to update data user")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("data user is not exist")
	}

	return data, nil
}

func (repo *user_repo) RemoveUser(r *http.Request, data *models.User) (*models.User, error) {
	vars := mux.Vars(r)
	result := repo.db.Where("name = ?", vars["name"]).Delete(data)

	if result.Error != nil {
		return nil, errors.New("failed to delete data user")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("data user is not exist")
	}

	return data, nil
}
