package users

import (
	"errors"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"gorm.io/gorm"
)

type user_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *user_repo {
	return &user_repo{db}
}

func (repo *user_repo) GetAllUsers() (*models.Users, error) {
	var data models.Users

	result := repo.db.Order("username ASC").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed to get data user")
	}
	if len(data) == 0 {
		return nil, errors.New("data user is empty")
	}

	return &data, nil
}

func (repo *user_repo) FindByUsername(username string) (*models.User, error) {
	var data models.User

	result := repo.db.First(&data, "username = ?", username)

	if result.Error != nil {
		return nil, errors.New("failed to get data user")
	}
	return &data, nil
}

func (repo *user_repo) AddUser(data *models.User) (*models.User, error) {
	result := repo.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("failed to add data user")
	}
	return data, nil
}

func (repo *user_repo) UpdateUser(data *models.User, username string) (*models.User, error) {
	result := repo.db.Model(data).Where("username = ?", username).Updates(data)
	if result.Error != nil {
		return nil, errors.New("failed to update data user")
	}
	return data, nil
}

func (repo *user_repo) DeleteUser(data *models.User, username string) (*models.User, error) {
	result := repo.db.Where("username = ?", username).Delete(&data)
	if result.Error != nil {
		return nil, errors.New("failed to delete data user")
	}
	return data, nil
}

func (repo *user_repo) CheckUsername(username string) bool {
	var data models.User
	result := repo.db.First(&data, "username = ?", username)
	return result.Error == nil
}

func (repo *user_repo) CheckEmail(email string) bool {
	var data models.User
	result := repo.db.First(&data, "email = ?", email)
	return result.Error == nil
}
