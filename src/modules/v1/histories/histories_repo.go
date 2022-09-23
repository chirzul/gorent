package histories

import (
	"errors"

	"github.com/chirzul/gorent/src/databases/orm/models"
	"gorm.io/gorm"
)

type history_repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *history_repo {
	return &history_repo{db}
}

func (repo *history_repo) GetAllHistories() (*models.Histories, error) {
	var data models.Histories
	result := repo.db.
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("user_id, name, created_at, updated_at")
		}).
		Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
			return db.Select("vehicle_id, name, created_at, updated_at")
		}).
		Find(&data)

	if result.Error != nil {
		return nil, errors.New("failed to get data history")
	}
	if len(data) == 0 {
		return nil, errors.New("data history is empty")
	}

	return &data, nil
}

func (repo *history_repo) SearchHistory(query map[string]string) (*models.Histories, error) {
	var data models.Histories
	var result *gorm.DB

	if query["vehicle_id"] != "" {
		result = repo.db.
			Preload("User", func(db *gorm.DB) *gorm.DB {
				return db.Select("user_id, name, created_at, updated_at")
			}).
			Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
				return db.Select("vehicle_id, name, created_at, updated_at")
			}).
			Order("created_at desc").
			Where("vehicle_id = ?", query["vehicle_id"]).Find(&data)
	}
	if query["user_id"] != "" {
		result = repo.db.
			Preload("User", func(db *gorm.DB) *gorm.DB {
				return db.Select("user_id, name, created_at, updated_at")
			}).
			Preload("Vehicle", func(db *gorm.DB) *gorm.DB {
				return db.Select("vehicle_id, name, created_at, updated_at")
			}).
			Order("created_at desc").
			Where("user_id = ?", query["user_id"]).Find(&data)
	}
	if result.Error != nil {
		return nil, errors.New("failed to get data vehicle")
	}
	if len(data) == 0 {
		return nil, errors.New("history not found")
	}

	return &data, nil
}

func (repo *history_repo) AddHistory(data *models.History) (*models.History, error) {
	var checkUser models.Users
	repo.db.Where("user_id = ?", data.UserID).Find(&checkUser)
	if len(checkUser) == 0 {
		return nil, errors.New("user does not exist")
	}

	var checkVehicle models.Vehicles
	repo.db.Where("vehicle_id = ?", data.VehicleID).Find(&checkVehicle)
	if len(checkVehicle) == 0 {
		return nil, errors.New("vehicle does not exist")
	}

	result := repo.db.Create(data)
	if result.Error != nil {
		return nil, errors.New("failed to add data history")
	}

	var dataVehicle models.Vehicle
	repo.db.First(&dataVehicle, "vehicle_id = ?", data.VehicleID)
	repo.db.Model(&dataVehicle).Where("vehicle_id = ?", data.VehicleID).Update("total_rented", dataVehicle.TotalRented+1)

	return data, nil
}

func (repo *history_repo) UpdateHistory(data *models.History, id string) (*models.History, error) {
	result := repo.db.Model(&data).Where("history_id = ?", id).Updates(data)
	if result.Error != nil {
		return nil, errors.New("failed to update data")
	}
	return data, nil
}

func (repo *history_repo) DeleteHistory(data *models.History, id string) (*models.History, error) {
	result := repo.db.Delete(&data, id)
	if result.Error != nil {
		return nil, errors.New("failed to delete data history")
	}
	return data, nil
}

func (repo *history_repo) CheckHistory(id string) bool {
	var data models.History
	result := repo.db.First(&data, "history_id = ?", id)
	return result.Error == nil
}
