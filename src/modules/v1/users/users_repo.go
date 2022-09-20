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

func (repo *user_repo) FindAllUsers() (*models.Users, error) {
	var data models.Users

	result := repo.db.Order("email ASC").Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed to get data user")
	}
	if len(data) == 0 {
		return nil, errors.New("data user is empty")
	}

	return &data, nil
}

func (repo *user_repo) SaveUser(data *models.User) (*models.User, error) {
	var check models.Users
	repo.db.Where("email = ?", data.Email).Find(&check)
	if len(check) > 0 {
		return nil, errors.New("email already registered")
	}

	result := repo.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("failed to add data user")
	}

	return data, nil
}

func (repo *user_repo) ChangeUser(data *models.User, email string) (*models.User, error) {
	result := repo.db.Model(&data).Where("email = ?", email).Updates(data)
	if result.Error != nil {
		return nil, errors.New("failed to update data user")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("data user is not exist")
	}
	return data, nil
}

func (repo *user_repo) RemoveUser(data *models.User, email string) (*models.User, error) {
	result := repo.db.Where("email = ?", email).Delete(&data)
	if result.Error != nil {
		return nil, errors.New("failed to delete data user")
	}
	if result.RowsAffected == 0 {
		return nil, errors.New("data user is not exist")
	}

	return data, nil
}
